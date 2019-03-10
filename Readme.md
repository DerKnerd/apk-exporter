# Alpine apk exporter for Prometheus
This is a simple wrapper around the `apk` executable on Alpine Linux. It will return a metric for installed and upgradable packages.

## Configuration
The only config is the listen address, simple pass it as a flag.