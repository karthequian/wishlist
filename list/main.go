package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/karthequian/wishlist/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	helloCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_calls",
			Help: "Number of authentication calls.",
		},
		[]string{"url"},
	)
)

var tokenMap map[string]common.User

func init() {
	tokenMap = make(map[string]common.User)

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(helloCounter)
	for _, user := range common.Userlist {
		tokenMap[user.Token] = user
	}

}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Version handler was called")
	helloCounter.With(prometheus.Labels{"url": "/version"}).Inc()
	fmt.Fprintf(w, "{'version':'1.0'}")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Status handler was called")
	helloCounter.With(prometheus.Labels{"url": "/status"}).Inc()
	fmt.Fprintf(w, "{'status':'ok'}")
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("List handler was called")
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", vars["user"])

	helloCounter.With(prometheus.Labels{"url": "/list"}).Inc()
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("newhandler was called")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the wishlist API. Valid endpoints are /wishlist/{user}, /version, /status, /metrics")

	helloCounter.With(prometheus.Labels{"url": "/list"}).Inc()
}

func main() {
	log.Info(os.Environ())
	port := os.Getenv("PORT")
	log.Infof("Port: %v", port)
	if len(port) == 0 {
		log.Fatalf("Port wasn't passed. An env variable for port must be passed")
	}

	r := mux.NewRouter()

	r.HandleFunc("/wishlist/{user}", ListHandler)
	r.HandleFunc("/version", versionHandler)
	r.HandleFunc("/status", statusHandler)
	r.HandleFunc("/", newHandler)
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("Starting up server")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
