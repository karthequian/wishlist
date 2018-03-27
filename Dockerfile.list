FROM golang
MAINTAINER Karthik Gaekwad

RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN go get github.com/gorilla/mux
COPY . /go/src/github.com/karthequian/wishlist
WORKDIR /go/src/github.com/karthequian/wishlist/list
RUN go get && go build -o /bin/list

ENTRYPOINT ["/bin/list"]