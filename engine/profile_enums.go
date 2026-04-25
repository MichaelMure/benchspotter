package engine

type Profile int

const (
	ProfileBench Profile = iota
	ProfileCPU
	ProfileMem
	ProfileMutex
	ProfileBlock
	ProfileEscape
	ProfileInline
)

// MemMetric selects which pprof SampleType to aggregate for a mem profile.
type MemMetric int

const (
	MetricAllocSpace   MemMetric = iota // total bytes allocated — best for benchmark allocation pressure
	MetricAllocObjects                  // total allocation events
	MetricInuseSpace                    // bytes live at profile time
	MetricInuseObjects                  // objects live at profile time
	memMetricCount
)

var memMetricNames = [memMetricCount]string{
	"alloc_space", "alloc_objects", "inuse_space", "inuse_objects",
}

func (m MemMetric) String() string  { return memMetricNames[m] }
func (m MemMetric) Next() MemMetric { return (m + 1) % memMetricCount }
func (m MemMetric) IsCount() bool   { return m == MetricAllocObjects || m == MetricInuseObjects }

// SortOrder controls how AggregateFuncs sorts its results.
type SortOrder int

const (
	SortFlat       SortOrder = iota // flat descending, then cumulative descending (default)
	SortCumulative                  // cumulative descending, then flat descending
	SortName                        // function name ascending
	sortOrderCount
)

var sortOrderNames = [sortOrderCount]string{"flat", "cumulative", "name"}

func (s SortOrder) String() string  { return sortOrderNames[s] }
func (s SortOrder) Next() SortOrder { return (s + 1) % sortOrderCount }
