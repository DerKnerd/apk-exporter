package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ApkCollector struct {
	InstalledPackages  *prometheus.Desc
	UpgradablePackages *prometheus.Desc
}

func (ApkCollector) Describe(channel chan<- *prometheus.Desc) {
	var collector = NewApkCollector()

	channel <- collector.InstalledPackages
	channel <- collector.UpgradablePackages
}

func (ApkCollector) Collect(channel chan<- prometheus.Metric) {
	var collector = NewApkCollector()

	channel <- getInstalledPackages(collector.InstalledPackages)
	channel <- getUpgradablePackages(collector.UpgradablePackages)
}

var _ prometheus.Collector = &ApkCollector{}

func NewApkCollector() *ApkCollector {
	const (
		subsystem = "apk"
		namespace = "alpine"
	)

	return &ApkCollector{
		InstalledPackages: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "installed_packages"),
			"The amount of installed packages",
			nil,
			nil,
		),
		UpgradablePackages: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "upgradable_packages"),
			"The amount of upgradable packages",
			nil,
			nil,
		),
	}
}
