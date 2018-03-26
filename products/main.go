package main

import (
	"encoding/json"
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
			Name: "product_calls",
			Help: "Number of product calls.",
		},
		[]string{"url"},
	)
)

func init() {

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(helloCounter)

	common.CreateProductMap()

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

func productsHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Products handler was called")
	productsJson, _ := json.Marshal(common.ProductList)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(productsJson))

	helloCounter.With(prometheus.Labels{"url": "/products"}).Inc()
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Product handler was called")
	w.Header().Set("Content-Type", "application/json")

	log.Infof("%v", common.ProductMap)

	vars := mux.Vars(r)
	log.Infof("Product key: %v\n", vars["key"])
	prd := common.ProductMap[vars["key"]]
	helloCounter.With(prometheus.Labels{"url": "/products"}).Inc()

	if len(prd.ID) == 0 {
		fmt.Fprintf(w, string("{}"))
		return
	}

	productJson, _ := json.Marshal(prd)
	fmt.Fprintf(w, string(productJson))
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("newhandler was called")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the Wishlist Catalog API. Valid endpoints are /products, /products/{productID}, /version, /status, /metrics")

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
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/products/{key}", productHandler)
	r.HandleFunc("/version", versionHandler)
	r.HandleFunc("/status", statusHandler)
	r.HandleFunc("/", newHandler)
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("Starting up server")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
