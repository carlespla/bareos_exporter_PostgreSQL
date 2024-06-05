## bareos_exporter
[![Go Report Card](https://goreportcard.com/badge/github.com/carlespla/bareos_exporter2)](https://goreportcard.com/report/github.com/carlespla/bareos_exporter2)

[Prometheus](https://github.com/prometheus) exporter for [bareos](https://github.com/bareos) data recovery system

### [`Dockerfile`](https://github.com/carlespla/bareos_exporter2/blob/master/Dockerfile)

### Usage with [docker](https://hub.docker.com/r/carlespla/bareos_exporter2)
1. Create a file containing your mysql password and mount it inside `/bareos_exporter/pw/auth`
2. **(optional)** [Overwrite](https://docs.docker.com/engine/reference/run/#env-environment-variables) default args using ENV variables
3. Run docker image as follows
```bash
docker run --name bareos_exporter2 -p 9625:9625 -v /bareos_exporter/pw/auth:/bareos_exporter2/pw/auth -d carlespla/bareos_exporter2:latest
```
### Metrics

- Total amout of bytes and files saved
- Latest executed job metrics (level, errors, execution time, bytes and files saved)
- Latest full job (level = F) metrics
- Amount of scheduled jobs

### Flags

Name    | Description                                                                                 | Default
--------|---------------------------------------------------------------------------------------------|----------------------
port    | Bareos exporter port                                                                        | 9625
endpoint| Bareos exporter endpoint.                                                                   | "/metrics"
u       | Username used to access Bareos MySQL Database                                               | "root"
p       | Path to file containing your MySQL password. Written inside a file to prevent from leaking. | "./auth"
h       | MySQL instance hostname.                                                                    | "127.0.0.1"
P       | MySQL instance port.                                                                        | "3306"
db      | MySQL database name.                                                                        | "bareos"
