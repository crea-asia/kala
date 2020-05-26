package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/ajvb/kala/api"
	"github.com/ajvb/kala/job"
	"github.com/ajvb/kala/job/storage/boltdb"
	"github.com/ajvb/kala/job/storage/consul"
	"github.com/ajvb/kala/job/storage/mongo"
	"github.com/ajvb/kala/job/storage/mysql"
	"github.com/ajvb/kala/job/storage/postgres"
	"github.com/ajvb/kala/job/storage/redis"

	"github.com/codegangsta/cli"
	redislib "github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var db job.JobDB
var App *cli.App

func init() {
	App = cli.NewApp()
	App.Name = "Kala"
	App.Usage = "Modern job scheduler"
	App.Commands = []cli.Command{
		{
			Name:  "run_command",
			Usage: "Run a command as if it was being run by Kala",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					log.Fatal("Must include a command")
				} else if len(c.Args()) > 1 {
					log.Fatal("Must only include a command")
				}

				cmd := c.Args()[0]

				j := &job.Job{
					Command: cmd,
				}

				err := j.RunCmd()
				if err != nil {
					log.Fatalf("Command Failed with err: %s", err)
				} else {
					fmt.Println("Command Succeeded!")
				}
			},
		},
		{
			Name:  "run",
			Usage: "run kala",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port, p",
					EnvVar: "KALA_PORT",
					Value:  ":8000",
					Usage:  "Port for Kala to run on.",
				},
				cli.BoolFlag{
					Name:  "no-persist, np",
					Usage: "No Persistence Mode - In this mode no data will be saved to the database. Perfect for testing.",
				},
				cli.StringFlag{
					Name:   "interface, i",
					EnvVar: "KALA_INTERFACE",
					Value:  "",
					Usage:  "Interface to listen on, default is all.",
				},
				cli.StringFlag{
					Name:  "default-owner, do",
					Value: "",
					Usage: "Default owner. The inputted email will be attached to any job missing an owner",
				},
				cli.StringFlag{
					Name:  "jobDB",
					Value: "boltdb",
					Usage: "Implementation of job database, either 'boltdb', 'redis', 'mongo', 'consul', 'postgres', 'mariadb', or 'mysql'.",
				},
				cli.StringFlag{
					Name:  "boltpath",
					Value: "",
					Usage: "Path to the bolt database file, default is current directory.",
				},
				cli.StringFlag{
					Name:  "jobDBAddress",
					Value: "",
					Usage: "Network address for the job database, in 'host:port' format.",
				},
				cli.StringFlag{
					Name:  "jobDBUsername",
					Value: "",
					Usage: "Username for the job database, in 'username' format.",
				},
				cli.StringFlag{
					Name:  "jobDBPassword",
					Value: "",
					Usage: "Password for the job database, in 'password' format.",
				},
				cli.StringFlag{
					Name:  "jobDBTlsCAPath",
					Value: "",
					Usage: "Path to tls server CA file for the job database.",
				},
				cli.StringFlag{
					Name:  "jobDBTlsCertPath",
					Value: "",
					Usage: "Path to tls client cert file for the job database.",
				},
				cli.StringFlag{
					Name:  "jobDBTlsKeyPath",
					Value: "",
					Usage: "Path to tls client key file for the job database.",
				},
				cli.StringFlag{
					Name:  "jobDBTlsServerName",
					Value: "",
					Usage: "Server name to verify cert for the job database.",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "Set for verbose logging.",
				},
				cli.IntFlag{
					Name:  "persist-every",
					Value: 5,
					Usage: "Sets the persisWaitTime in seconds",
				},
				cli.IntFlag{
					Name:  "jobstat-ttl",
					Value: -1,
					Usage: "Sets the jobstat-ttl in minutes. The default -1 value indicates JobStat entries will be kept forever",
				},
				cli.BoolFlag{
					Name:  "profile",
					Usage: "Activate pprof handlers",
				},
			},
			Action: func(c *cli.Context) {
				if c.Bool("v") {
					log.SetLevel(log.DebugLevel)
				}

				var parsedPort string
				port := c.String("port")
				if port != "" {
					if strings.Contains(port, ":") {
						parsedPort = port
					} else {
						parsedPort = ":" + port
					}
				} else {
					parsedPort = ":8000"
				}

				var connectionString string
				if c.String("interface") != "" {
					connectionString = c.String("interface") + parsedPort
				} else {
					connectionString = parsedPort
				}

				switch c.String("jobDB") {
				case "boltdb":
					db = boltdb.GetBoltDB(c.String("boltpath"))
				case "redis":
					if c.String("jobDBPassword") != "" {
						option := redislib.DialPassword(c.String("jobDBPassword"))
						db = redis.New(c.String("jobDBAddress"), option, true)
					} else {
						db = redis.New(c.String("jobDBAddress"), redislib.DialOption{}, false)
					}
				case "mongo":
					if c.String("jobDBUsername") != "" {
						cred := &mgo.Credential{
							Username: c.String("jobDBUsername"),
							Password: c.String("jobDBPassword")}
						db = mongo.New(c.String("jobDBAddress"), cred)
					} else {
						db = mongo.New(c.String("jobDBAddress"), &mgo.Credential{})
					}
				case "consul":
					db = consul.New(c.String("jobDBAddress"))
				case "postgres":
					dsn := fmt.Sprintf("postgres://%s:%s@%s", c.String("jobDBUsername"), c.String("jobDBPassword"), c.String("jobDBAddress"))
					db = postgres.New(dsn)
				case "mysql", "mariadb":
					dsn := fmt.Sprintf("%s:%s@%s", c.String("jobDBUsername"), c.String("jobDBPassword"), c.String("jobDBAddress"))
					log.Debug("Mysql/Maria DSN: ", dsn)
					if c.String("jobDBTlsCAPath") != "" {
						// https://godoc.org/github.com/go-sql-driver/mysql#RegisterTLSConfig
						rootCertPool := x509.NewCertPool()
						pem, err := ioutil.ReadFile(c.String("jobDBTlsCAPath"))
						if err != nil {
							log.Fatal(err)
						}
						if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
							log.Fatal("Failed to append PEM.")
						}
						clientCert := make([]tls.Certificate, 0, 1)
						certs, err := tls.LoadX509KeyPair(c.String("jobDBTlsCertPath"), c.String("jobDBTlsKeyPath"))
						if err != nil {
							log.Fatal(err)
						}
						clientCert = append(clientCert, certs)
						db = mysql.New(dsn, &tls.Config{
							RootCAs:      rootCertPool,
							Certificates: clientCert,
							ServerName:   c.String("jobDBTlsServerName"),
						})
					} else {
						db = mysql.New(dsn, nil)
					}
				default:
					log.Fatalf("Unknown Job DB implementation '%s'", c.String("jobDB"))
				}

				if c.Bool("no-persist") {
					db = &job.MockDB{}
				}

				// Create cache
				cache := job.NewLockFreeJobCache(db)
				log.Infof("Preparing cache")
				cache.Start(time.Duration(c.Int("persist-every"))*time.Second, time.Duration(c.Int("jobstat-ttl"))*time.Minute)

				log.Infof("Starting server on port %s", connectionString)
				srv := api.MakeServer(connectionString, cache, db, c.String("default-owner"), c.Bool("profile"))
				log.Fatal(srv.ListenAndServe())
			},
		},
	}
}
