package cgroups

// Metrics represents container stats structure
type Metrics struct {
	Hugetlb []*HugetlbStat `json:"hugetlb,omitempty"`
	Pids    *PidsStat      `json:"pids,omitempty"`
	CPU     *CPUStat       `json:"cpu,omitempty"`
	Memory  *MemoryStat    `json:"memory,omitempty"`
	Blkio   *BlkIOStat     `json:"blkio,omitempty"`
	Rdma    *RdmaStat      `json:"rdma,omitempty"`
}

// HugetlbStat represents huge TLB container stats structure
type HugetlbStat struct {
	Usage    uint64 `json:"usage,omitempty"`
	Max      uint64 `json:"max,omitempty"`
	Failcnt  uint64 `json:"failcnt,omitempty"`
	Pagesize string `json:"pagesize,omitempty"`
}

// PidsStat represents pid container stats structure
type PidsStat struct {
	Current uint64 `json:"current,omitempty"`
	Limit   uint64 `json:"limit,omitempty"`
}

// CPUStat represents cpu container stats structure
type CPUStat struct {
	Usage      *CPUUsage `json:"usage,omitempty"`
	Throttling *Throttle `json:"throttling,omitempty"`
}

// CPUUsage represents cpu usage container stats structure
type CPUUsage struct {
	// values in nanoseconds
	Total  uint64   `json:"total,omitempty"`
	Kernel uint64   `json:"kernel,omitempty"`
	User   uint64   `json:"user,omitempty"`
	PerCPU []uint64 `json:"per_cpu,omitempty"`
}

// Throttle represents throttle container stats structure
type Throttle struct {
	Periods          uint64 `json:"periods,omitempty"`
	ThrottledPeriods uint64 `json:"throttled_periods,omitempty"`
	ThrottledTime    uint64 `json:"throttled_time,omitempty"`
}

// MemoryStat represents memory container stats structure
type MemoryStat struct {
	Cache                   uint64       `json:"cache,omitempty"`
	RSS                     uint64       `json:"rss,omitempty"`
	RSSHuge                 uint64       `json:"rss_huge,omitempty"`
	MappedFile              uint64       `json:"mapped_file,omitempty"`
	Dirty                   uint64       `json:"dirty,omitempty"`
	Writeback               uint64       `json:"writeback,omitempty"`
	PgPgIn                  uint64       `json:"pg_pg_in,omitempty"`
	PgPgOut                 uint64       `json:"pg_pg_out,omitempty"`
	PgFault                 uint64       `json:"pg_fault,omitempty"`
	PgMajFault              uint64       `json:"pg_maj_fault,omitempty"`
	InactiveAnon            uint64       `json:"inactive_anon,omitempty"`
	ActiveAnon              uint64       `json:"active_anon,omitempty"`
	InactiveFile            uint64       `json:"inactive_file,omitempty"`
	ActiveFile              uint64       `json:"active_file,omitempty"`
	Unevictable             uint64       `json:"unevictable,omitempty"`
	HierarchicalMemoryLimit uint64       `json:"hierarchical_memory_limit,omitempty"`
	HierarchicalSwapLimit   uint64       `json:"hierarchical_swap_limit,omitempty"`
	TotalCache              uint64       `json:"total_cache,omitempty"`
	TotalRSS                uint64       `json:"total_rss,omitempty"`
	TotalRSSHuge            uint64       `json:"total_rss_huge,omitempty"`
	TotalMappedFile         uint64       `json:"total_mapped_file,omitempty"`
	TotalDirty              uint64       `json:"total_dirty,omitempty"`
	TotalWriteback          uint64       `json:"total_writeback,omitempty"`
	TotalPgPgIn             uint64       `json:"total_pg_pg_in,omitempty"`
	TotalPgPgOut            uint64       `json:"total_pg_pg_out,omitempty"`
	TotalPgFault            uint64       `json:"total_pg_fault,omitempty"`
	TotalPgMajFault         uint64       `json:"total_pg_maj_fault,omitempty"`
	TotalInactiveAnon       uint64       `json:"total_inactive_anon,omitempty"`
	TotalActiveAnon         uint64       `json:"total_active_anon,omitempty"`
	TotalInactiveFile       uint64       `json:"total_inactive_file,omitempty"`
	TotalActiveFile         uint64       `json:"total_active_file,omitempty"`
	TotalUnevictable        uint64       `json:"total_unevictable,omitempty"`
	Usage                   *MemoryEntry `json:"usage,omitempty"`
	Swap                    *MemoryEntry `json:"swap,omitempty"`
	Kernel                  *MemoryEntry `json:"kernel,omitempty"`
	KernelTCP               *MemoryEntry `json:"kernel_tcp,omitempty"`
}

// MemoryEntry represents memory entry container stats structure
type MemoryEntry struct {
	Limit   uint64 `json:"limit,omitempty"`
	Usage   uint64 `json:"usage,omitempty"`
	Max     uint64 `json:"max,omitempty"`
	Failcnt uint64 `json:"failcnt,omitempty"`
}

// BlkIOStat represents block IO container stats structure
type BlkIOStat struct {
	IoServiceBytesRecursive []*BlkIOEntry `json:"io_service_bytes_recursive,omitempty"`
	IoServicedRecursive     []*BlkIOEntry `json:"io_serviced_recursive,omitempty"`
	IoQueuedRecursive       []*BlkIOEntry `json:"io_queued_recursive,omitempty"`
	IoServiceTimeRecursive  []*BlkIOEntry `json:"io_service_time_recursive,omitempty"`
	IoWaitTimeRecursive     []*BlkIOEntry `json:"io_wait_time_recursive,omitempty"`
	IoMergedRecursive       []*BlkIOEntry `json:"io_merged_recursive,omitempty"`
	IoTimeRecursive         []*BlkIOEntry `json:"io_time_recursive,omitempty"`
	SectorsRecursive        []*BlkIOEntry `json:"sectors_recursive,omitempty"`
}

// BlkIOEntry represents block IO entry container stats structure
type BlkIOEntry struct {
	Op     string `json:"op,omitempty"`
	Device string `json:"device,omitempty"`
	Major  uint64 `json:"major,omitempty"`
	Minor  uint64 `json:"minor,omitempty"`
	Value  uint64 `json:"value,omitempty"`
}

// RdmaStat represents RDMA container stats structure
type RdmaStat struct {
	Current []*RdmaEntry `json:"current,omitempty"`
	Limit   []*RdmaEntry `json:"limit,omitempty"`
}

// RdmaEntry represents RDMA entry container stats structure
type RdmaEntry struct {
	Device     string `json:"device,omitempty"`
	HcaHandles uint32 `json:"hca_handles,omitempty"`
	HcaObjects uint32 `json:"hca_objects,omitempty"`
}
