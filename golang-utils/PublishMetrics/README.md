Below is the curl command to post the metric

curl -X POST \
  http://127.0.0.1:8085/metrics/publish \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 8e8f19d8-7655-1400-447e-4513f6f1e54a' \
  -d '{
  "receiver": "pubsub",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "CPULoad",
        "env": "poc",
        "instance": "test_instance ",
        "severity": "critical"
      },
      "annotations": {
        "description": "CPU load is > 80%\n VALUE = 0.6761904761905129\n LABELS = map[instance:prometheus-server]",
        "summary": "Host high CPU load (instance prometheus-server)"
      },
      "startsAt": "2022-03-11T12:06:08.884Z",
      "endsAt": "0001-01-01T00:00:00Z",
      "generatorURL": "http://prometheus-server:9090/graph?g0.expr=100+-+%28avg+by%28instance%29+%28rate%28node_cpu_seconds_total%7Bmode%3D%22idle%22%7D%5B2m%5D%29%29+%2A+100%29+%3C+80&g0.tab=1",
      "fingerprint": "14b92061ab1aed4a"
    }
  ],
  "groupLabels": {
    "alertname": "CPULoad",
    "env": "poc"
  },
  "commonLabels": {
    "alertname": "CPULoad",
    "env": "poc",
    "instance": "prometheus-server",
    "severity": "critical"
  },
  "commonAnnotations": {
    "description": "CPU load is > 80%\n VALUE = 0.6761904761905129\n LABELS = map[instance:prometheus-server]",
    "summary": "Host high CPU load (instance prometheus-server)"
  },
  "externalURL": "http://prometheus-server:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"CPULoad\", env=\"poc\"}",
  "truncatedAlerts": 0
}'