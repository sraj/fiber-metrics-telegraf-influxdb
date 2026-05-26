# fiber-metrics-telegraf-influxdb

Golang Fiber web app with a full monitoring stack: Prometheus metrics → Telegraf → InfluxDB → Grafana.

## Architecture

```
┌─────────────┐     ┌──────────┐     ┌──────────┐     ┌─────────┐
│  Fiber App  │────▶│ Telegraf │────▶│ InfluxDB │◀────│ Grafana │
│  (:3000)    │     │          │     │ (:8086)  │     │ (:3030) │
└─────────────┘     └──────────┘     └──────────┘     └─────────┘
       │
       ▼
  Prometheus
  /metrics
```

## Services

| Service   | Port  | Description                                      |
|-----------|-------|--------------------------------------------------|
| **app**   | 3000  | Go Fiber app with Prometheus metrics middleware   |
| **telegraf** | - | Scrapes `/metrics`, collects system stats, writes to InfluxDB |
| **influxdb** | 8086 | Time-series database (v2.7, bucket: `metrics`) |
| **grafana** | 3030 | Pre-provisioned dashboards querying InfluxDB |

## Tech Stack

- **Go 1.23** with [Fiber v2](https://github.com/gofiber/fiber) web framework
- **Prometheus** client library for HTTP request duration histograms (`http_response_time_seconds`)
- **Telegraf** scrapes Prometheus metrics and system stats (CPU, disk, memory, network)
- **InfluxDB v2** as the time-series storage backend
- **Grafana** with auto-provisioned datasource and dashboard
- **Air** for hot-reloading during development

## Quick Start

```bash
git clone <repo>
cd fiber-metrics-telegraf-influxdb
docker-compose up
```

Then:
- **App**: http://localhost:3000
- **Grafana**: http://localhost:3030 (login: `admin` / `admin123`)
- **InfluxDB**: http://localhost:8086

## Project Structure

```
.
├── main.go                    # Fiber app entrypoint
├── Dockerfile                 # Multi-stage dev build with Air hot-reload
├── docker-compose.yml         # Orchestrates all 4 services
├── .air.toml                  # Air hot-reload config
├── telegraf/
│   └── telegraf.conf          # Telegraf config: prometheus input + influxdb output
├── grafana/
│   └── provisioning/
│       ├── datasources/
│       │   └── influxdb.yml   # Auto-provisioned InfluxDB datasource
│       └── dashboards/
│           ├── dashboard.yml  # Dashboard provider config
│           └── go-dashboard.json  # Pre-built Go metrics dashboard
└── README.md
```

## What Gets Measured

The Fiber app exposes a **Prometheus histogram** `http_response_time_seconds` tagged by `method` and `path`. Telegraf scrapes this at `/metrics` every 10s, alongside system metrics (CPU, disk, memory, network, load).

## Grafana

Grafana is pre-configured with:
- An **InfluxDB datasource** pointing at the `myorg` organization
- A **Go dashboard** dashboard auto-loaded from provisioning files

Login at http://localhost:3030 with `admin` / `admin123`.
