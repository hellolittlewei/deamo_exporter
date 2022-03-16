// deamo_exporter.go是prometheus指标插件,
// 运行它会按照prometheus格式暴露指标数据.

package main

import (
	ac "deamo_expoeter/collector"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/exporter-toolkit/web"
	"github.com/prometheus/exporter-toolkit/web"
	"net/http"
)

// registry 是我们自定义的注册器，而不是全局的.
// prometheus通过NewGoCollector()和NewProcessCollector()函数
// 创建Go运行时数据指标的NewGoCollector和进程数据指标的NewProcessCollector. 而这
// 两个Collector已在默认的注册表DefaultRegisterer中注册.使用自定
// 义注册表，您可以控制并自行决定要注册的 Collector.
var registry *prometheus.Registry

// 程序运行前必须先注册我们的收集器
func init() {
	registry = prometheus.NewRegistry()
	for _, collector := range ac.SNC {
		registry.MustRegister(collector())
	}
}

func main() {
	// 定义一个我们的日志配置
	promLogConfig := &promlog.Config{}
	logger := promlog.New(promLogConfig)

	// 暴露自定义指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	// 定义默认访问页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<head><title>Node Exporter</title></head>
			<body>
			<h1>Node Exporter</h1>
			<p><a href="` + "/metrics" + `">Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			panic(err)
		}
	})
	err := level.Info(logger).Log("msg", "Listening on", "address", "9900")
	if err != nil {
		panic(err)
	}
	server := &http.Server{Addr: ":9900"}
	if err = web.ListenAndServe(server, "", logger); err != nil {
		panic(err)
	}
}