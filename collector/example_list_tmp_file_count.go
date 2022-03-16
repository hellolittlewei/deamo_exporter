// 这是一个测试收集器，它将显示/tmp目录下的目录
// 和文件个数，当然它是一个模板，您可以参考它实
// 现自己的Collector，祝您好运 <^^>

package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
)

// TmpFileCount 对应为该指标需要收集的数据
// 例如我们收集文件个数以及目录个数
type TmpFileCount struct {
	fileCount *prometheus.Desc
	dirCount  *prometheus.Desc
}

// Describe 实现Collector接口的Describe方法
// 它是必须的
func (t *TmpFileCount) Describe(descs chan<- *prometheus.Desc) {
	descs <- t.dirCount
	descs <- t.fileCount
}

// Collect 实现Collector接口的Collect方法
// 它是必须的，当然这里面也可以直接写你的收集实现
// 这样你可以省略Update函数，都一样.
func (t *TmpFileCount) Collect(metrics chan<- prometheus.Metric) {
	err := t.Update(metrics)
	if err != nil {
		panic(err)
	}
}

// TmpFileCountStatistics 对应的统计数据
type TmpFileCountStatistics struct {
	fileRootN uint64
	dirRootN  uint64
}

// Update 实现收集metrics的动作,这里就是收集器的具体实现
func (t *TmpFileCount) Update(ch chan<- prometheus.Metric) error {
	var dirPath string = "/tmp"
	fc, dc, err := goToCountNumber(dirPath)

	tmpFile := TmpFileCountStatistics{}
	tmpFile.fileRootN = fc
	tmpFile.dirRootN = dc
	ch <- prometheus.MustNewConstMetric(t.fileCount, prometheus.GaugeValue, float64(tmpFile.fileRootN))
	ch <- prometheus.MustNewConstMetric(t.dirCount, prometheus.GaugeValue, float64(tmpFile.dirRootN))
	if err != nil {
		return err
	}
	return nil
}

// MyCounterCollector 是我们的收集器
// 这一步非常重要
func MyCounterCollector() prometheus.Collector {
	return &TmpFileCount{
		fileCount: prometheus.NewDesc("file_tmp_n", "", []string{}, prometheus.Labels{"file_tmp_number": "tmp"}),
		dirCount:  prometheus.NewDesc("dir_tmp_n", "", []string{}, prometheus.Labels{}),
	}
}

// 注册我们的收集器到自定义注册表
func init() {
	MakeSNC(MyCounterCollector)
}

func goToCountNumber(path string) (fC uint64, dC uint64, err error) {
	var dirCount uint64 = 0
	var fileCount uint64 = 0
	files, err := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
	}
	return fileCount, dirCount, err
}
