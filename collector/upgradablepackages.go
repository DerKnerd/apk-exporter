package collector

import (
	"../utils"
	"github.com/prometheus/client_golang/prometheus"
)

func getUpgradablePackages(desc *prometheus.Desc) prometheus.Metric {
	result, err := utils.ExecuteApk("list", "--upgradable")

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(desc, prometheus.CounterValue, utils.LineCounter(result))
}
