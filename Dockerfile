FROM golang as builder
RUN mkdir /git
WORKDIR /git
RUN git clone https://github.com/carlespla/bareos_exporter_PostgreSQL
WORKDIR /git/bareos_exporter_PostgreSQL
RUN rm -f go.mod go.sum bareos_exporter_PostgreSQL
RUN go mod init github.com/carlespla/bareos_exporter_PostgreSQL
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bareos_exporter_PostgreSQL .

FROM busybox:latest

ENV sql_port 5432
ENV sql_server 127.0.0.1
ENV sql_username root
ENV endpoint /metrics
ENV port 9625

WORKDIR /bareos_exporter_PostgreSQL
COPY --from=builder /git/bareos_exporter_PostgreSQL bareos_exporter_PostgreSQL
RUN chmod +x /bareos_exporter_PostgreSQL/bareos_exporter_PostgreSQL/bareos_exporter_PostgreSQL


CMD ./bareos_exporter_PostgreSQL/bareos_exporter_PostgreSQL -port $port -endpoint $endpoint -u $sql_username -h $sql_server -P $sql_port -p pw/auth
EXPOSE $port
