// collector.go主要实现对收集器(prometheus.Collector)的组合.
package collector

import "github.com/prometheus/client_golang/prometheus"

// SNC 是我们需要的收集器工厂函数切片，
// 它存放了我们需要组测的收集器，它是核心.
var SNC []func() prometheus.Collector

// MakeSNC 我们的收集器添加到SNC中.
func MakeSNC(fName func() prometheus.Collector) {
	SNC = append(SNC, fName)
}
