package main

import (
	"github.com/carlespla/bareos_exporter_PostgreSQL/error"

	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var connectionString string

var (
	exporterPort     = flag.Int("port", 9625, "Bareos exporter port")
	exporterEndpoint = flag.String("endpoint", "/metrics", "Bareos exporter endpoint")
	mysqlUser        = flag.String("u", "root", "Bareos PostgreSQL username")
	mysqlAuthFile    = flag.String("p", "./auth", "PostgreSQL password file path")
	mysqlHostname    = flag.String("h", "127.0.0.1", "PostgreSQL hostname")
	mysqlPort        = flag.String("P", "5432", "PostgreSQL port")
	mysqlDb          = flag.String("db", "bareos", "PostgreSQL database name")
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: bareos_exporter_PostgreSQL [ ... ]\n\nParameters:")
		fmt.Println()
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	pass, err := ioutil.ReadFile(*mysqlAuthFile)
	error.Check(err)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", *mysqlUser, strings.TrimSpace(string(pass)), *mysqlHostname, *mysqlPort, *mysqlDb)
	log.Info("Starting ...")
	collector := bareosCollector()
	log.Info("Conecting ...", collector)
	prometheus.MustRegister(collector)

	http.Handle(*exporterEndpoint, promhttp.Handler())
	log.Info("Beginning to server on port ", *exporterPort)

	addr := fmt.Sprintf(":%d", *exporterPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}
