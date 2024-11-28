# otel-workshop-go

This repository contains two APIs developed in Golang, with the objetive to learn how could we instrument applications with OpenTelemetry using his SDK to generate traces, metrics, and logs. These APIs will be deployed on Kubernetes (K8s), along with several monitoring providers such as Prometheus for metrics and Jaeger for traces.

## Goals
- Understand OpenTelemetry:
  - Learn how to generate traces, metrics, and logs using Golang's OpenTelemetry SDK.
- Deploy and Configure OTel Collector:
    Use the OpenTelemetry Collector to gather, process, and export telemetry data.
- Integrate Monitoring Providers:
  - Visualize metrics in Prometheus.
  - Analyze traces in Jaeger.
  - Explore logs with a logging solution (e.g., Fluentd/Elasticsearch).
  - Create dashboards in Grafana.
- Containerize services created in this repository 
- Hands-on Experience with Kubernetes:
    Deploy APIs, OTel Collector, and monitoring tools on a Kubernetes cluster.

## Getting Started
### Prerequisites
 - Golang (v1.23+)
 - Docker
 - kubectl 
 - Helm 
  
## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

Happy Observing! 🚀