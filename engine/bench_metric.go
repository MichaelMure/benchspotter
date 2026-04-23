package engine

// BenchMetric identifies a standard Go benchmark output metric.
type BenchMetric int

const (
	MetricNsPerOp    BenchMetric = iota // ns/op
	MetricBytesPerOp                    // B/op
	MetricAllocsPerOp                   // allocs/op
	benchMetricCount
)

var benchMetricKeys = [benchMetricCount]string{"ns/op", "B/op", "allocs/op"}
var benchMetricLabels = [benchMetricCount]string{
	"time per operation",
	"bytes allocated per operation",
	"allocations per operation",
}

func (m BenchMetric) String() string { return benchMetricKeys[m] }
func (m BenchMetric) Label() string  { return benchMetricLabels[m] }

// KnownBenchMetrics lists all standard benchmark metrics in display order.
var KnownBenchMetrics = func() []BenchMetric {
	ms := make([]BenchMetric, benchMetricCount)
	for i := range ms {
		ms[i] = BenchMetric(i)
	}
	return ms
}()
