package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/crea-asia/kala/job"
)

var (
	TableName = "jobs"
)

type DB struct {
	conn *sql.DB
}

// New instantiates a new DB.
func New(dsn string) *DB {
	var dbConn *sql.DB
	var err error
	for i := 0; i < 30; i++ {
		dbConn, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("unable to connect to DB, retrying... error=%v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if  err := dbConn.Ping(); err != nil {
			log.Printf("unable to connect to DB, retrying... error=%v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}
	if err != nil {
		log.Printf("unable to connect to DB, error=%v\n", err)
	}


	// passive attempt to create table
	_, _ = dbConn.Exec(fmt.Sprintf(`create table %s (id varchar(120) unique, job jsonb);`, TableName))
	return &DB{
		conn: dbConn,
	}
}

// GetAll returns all persisted Jobs.
func (d DB) GetAll() ([]*job.Job, error) {
	query := fmt.Sprintf(`select coalesce(json_agg(j.job), '[]'::json) from (select * from %[1]s) as j;`, TableName)
	var r sql.NullString
	err := d.conn.QueryRow(query).Scan(&r)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	err = nil
	jobs := []*job.Job{}
	if r.Valid {
		err = json.Unmarshal([]byte(r.String), &jobs)
	}
	return jobs, err
}

// Get returns a persisted Job.
func (d DB) Get(id string) (*job.Job, error) {
	template := `select to_jsonb(j.job) from (select * from %[1]s where id = $1) as j;`
	query := fmt.Sprintf(template, TableName)
	var r sql.NullString
	err := d.conn.QueryRow(query, id).Scan(&r)
	if err != nil {
		return nil, err
	}
	result := &job.Job{}
	if r.Valid {
		err = json.Unmarshal([]byte(r.String), &result)
	}
	return result, err
}

// Delete deletes a persisted Job.
func (d DB) Delete(id string) error {
	query := fmt.Sprintf(`delete from %v where id = $1;`, TableName)
	_, err := d.conn.Exec(query, id)
	return err
}

// Save persists a Job.
func (d DB) Save(j *job.Job) error {
	template := `insert into %[1]s (id, job) values($1, $2) ON CONFLICT ON CONSTRAINT jobs_id_key
		DO UPDATE SET job = $2`
	query := fmt.Sprintf(template, TableName)
	r, err := json.Marshal(j)
	if err != nil {
		return err
	}
	transaction, err := d.conn.Begin()
	if err != nil {
		return err
	}
	statement, err := transaction.Prepare(query)
	if err != nil {
		transaction.Rollback() //nolint:errcheck // adding insult to injury
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(j.Id, string(r))
	if err != nil {
		transaction.Rollback() //nolint:errcheck // adding insult to injury
		return err
	}
	return transaction.Commit()
}

// Close closes the connection to Postgres.
func (d DB) Close() error {
	return d.conn.Close()
}
