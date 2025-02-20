# fiber-metrics-telegraf-influxdb
Golang fiber based app integrated with grafana, telegraf, influxdb to showcase monitoring.

### How this works?

1. `git clone` the repository.
2. `docker-compose up` to start services (app, influxdb, telegraf, grafana)
3. visit http://localhost:3030 to login to grafana (use the credentials specified in docker-compose.yml)
4. Change the InfluxDB datasource to match.
