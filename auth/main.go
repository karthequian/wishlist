package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/karthequian/wishlist/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

//Userlist is the list of users
var userlist = []common.User{
	common.User{
		Name:     "Karthik",
		Password: "helloworld",
		Username: "karthik",
		Token:    "20d15fed-42f4-4a71-b9a3-7c7fee78d38d",
	},
	common.User{
		Name:     "Tracey",
		Password: "helloworld",
		Username: "tracey",
		Token:    "4d76c945-a946-4d2b-95a8-281aff55404f",
	},
	common.User{
		Name:     "Carisa",
		Password: "helloworld",
		Username: "carisa",
		Token:    "224b3200-d09b-4881-8a7b-d69d6d8ba543",
	},
	common.User{
		Name:     "Ernest",
		Password: "helloworld",
		Username: "ernest",
		Token:    "2115274e-34bc-4456-a1ee-1c4c171231a9",
	},
	common.User{
		Name:     "Amy",
		Password: "helloworld",
		Username: "amy",
		Token:    "5343b1d3-dfa3-4823-b544-e5907c3585f5",
	},
}

var tokenMap map[string]common.User

func init() {
	tokenMap = make(map[string]common.User)

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(helloCounter)
	for _, user := range userlist {
		tokenMap[user.Token] = user
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	helloCounter.With(prometheus.Labels{"url": "/hello"}).Inc()
	fmt.Fprintf(w, "12 clouds demo at CloudAustin 2017!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	helloCounter.With(prometheus.Labels{"url": "/login"}).Inc()
	u := r.URL.Query().Get("u")
	p := r.URL.Query().Get("p")

	for _, user := range userlist {
		if user.Username == u {
			if p == user.Password {
				returnUser := user
				returnUser.Password = ""
				jsonuser, _ := json.Marshal(returnUser)
				fmt.Fprintf(w, string(jsonuser))
				return
			}
		}
	}
	http.Error(w, "Invalid username or password", 500)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	helloCounter.With(prometheus.Labels{"url": "/token"}).Inc()
	t := r.URL.Query().Get("t")

	foundUser := tokenMap[t]
	if len(foundUser.Token) > 0 {
		foundUser.Password = ""
		jsonuser, _ := json.Marshal(foundUser)
		fmt.Fprintf(w, string(jsonuser))
		return
	}
	http.Error(w, "Invalid token", 500)

}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("Need to pass port as an argument")
		return
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/token", TokenHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+argsWithoutProg[0], nil)
}
