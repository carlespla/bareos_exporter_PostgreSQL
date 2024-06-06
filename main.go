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
	postgresUser     = flag.String("u", "root", "Bareos PostgreSQL username")
	postgresPassword = flag.String("p", "./auth", "PostgreSQL password file path")
	postgresHostname = flag.String("h", "127.0.0.1", "PostgreSQL hostname")
	postgresPort     = flag.String("P", "5432", "PostgreSQL port")
	postgresDb       = flag.String("db", "bareos", "PostgreSQL database name")
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

	pass, err := ioutil.ReadFile(*postgresPassword)
	error.Check(err)
	connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgresHostname, postgresPort, postgresUser, strings.TrimSpace(string(pass)), postgresDb)
	//connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *postgresUser, strings.TrimSpace(string(pass)), *postgresHostname, *postgresPort, *postgresDb)
	log.Info("Starting ...", connectionString)
	collector := bareosCollector()
	log.Info("Conecting ...", collector)
	prometheus.MustRegister(collector)

	http.Handle(*exporterEndpoint, promhttp.Handler())
	log.Info("Beginning to server on port ", *exporterPort)

	addr := fmt.Sprintf(":%d", *exporterPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}
