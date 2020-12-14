package integrate

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/crea-asia/kala/api"
	kalaclient "github.com/crea-asia/kala/client"
	"github.com/crea-asia/kala/clock"
	"github.com/crea-asia/kala/job"
)

func TestIntegrationTest(t *testing.T) {

	jobDB := &job.MockDB{}
	cache := job.NewLockFreeJobCache(jobDB)

	clk := clock.NewMockClock()
	cache.Clock.SetClock(clk)

	addr := newLocalListener(t)

	kalaApi := api.MakeServer(addr, cache, "default@example.com", false)
	go kalaApi.ListenAndServe()
	runtime.Gosched()
	kalaClient := kalaclient.NewAPIClient(&kalaclient.Configuration{BasePath: "http://" + addr})

	hit := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hit <- struct{}{}
		w.WriteHeader(200)
	}))

	scheduleTime := clk.Now().Add(time.Minute * 3)
	parsedTime := scheduleTime.Format(time.RFC3339)
	delay := "PT5M"
	scheduleStr := fmt.Sprintf("R/%s/%s", parsedTime, delay)
	body := kalaclient.Job{
		Name:     "Increment HitCount",
		Schedule: scheduleStr,
		Type:     int64(job.RemoteJob),
		RemoteProperties: kalaclient.RemoteProperties{
			Url:                   "http://" + srv.Listener.Addr().String(),
			Method:                http.MethodGet,
			ExpectedResponseCodes: []int64{200},
		},
	}
	_, _, err := kalaClient.DefaultApi.CreateJob(context.Background(), &kalaclient.CreateJobOpts{
		CreateJob: optional.NewInterface(body),
	})
	if err != nil {
		t.Fatal(err)
	}

	clk.AddTime(time.Minute)
	select {
	case <-hit:
		t.Fatalf("Did not expect job to have run yet.")
	case <-time.After(time.Second):
	}

	clk.AddTime(time.Minute)
	select {
	case <-hit:
		t.Fatalf("Still did not expect job to have run yet.")
	case <-time.After(time.Second):
	}

	clk.AddTime(time.Minute * 3)
	select {
	case <-hit:
	case <-time.After(time.Second * 5):
		t.Fatalf("Expected job to have run.")
	}

	clk.AddTime(time.Minute * 5)
	select {
	case <-hit:
	case <-time.After(time.Second * 5):
		t.Fatalf("Expected job to have run again.")
	}

}

func newLocalListener(t *testing.T) string {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			t.Fatalf("failed to listen on a port: %v", err)
		}
	}
	defer func() {
		if err := l.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	return l.Addr().String()
}
