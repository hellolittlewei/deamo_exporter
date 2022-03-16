// collector.go主要实现对收集器(prometheus.Collector)的组合.
// 它是非常必要的,如果您要使用它请看下面的离职.

// 先创建您的收集器NewCollector()
// func NewCollector() prometheus.Collector {
//		当然您的MyMetrics必须实现Collector接口
//  	return &MyMetrics{
//  	my_metrics: prometheus.NewDesc("you_metrics_name", "", []string{}, prometheus.Labels{}),
//  	}
//  }
//	调用生成器函数(MakeSNC)进行注册.
//	func init()  {
//		MakeSNC(NewCollector)
//	}

package collector

import "github.com/prometheus/client_golang/prometheus"

// SNC 是我们需要的收集器工厂函数切片，
// 它存放了我们需要组测的收集器，它是核心.
var SNC []func() prometheus.Collector

// MakeSNC 我们的收集器添加到SNC中.
func MakeSNC(fName func() prometheus.Collector) {
	SNC = append(SNC, fName)
}
