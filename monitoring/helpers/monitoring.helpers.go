package helpers

import "runtime"

type MemoryStat struct {
	RSSBytes          uint64 // Rough equivalent to Resident Set Size
	HeapUsedBytes     uint64
	HeapTotalBytes    uint64
	HeapSysBytes      uint64 // Total bytes obtained from system (similar to heapTotal)
	ExternalBytes     uint64 // Go doesn't track this separately; placeholder
	ArrayBuffersBytes uint64 // Go doesn't track this; placeholder
}

func GetServerStatistic() MemoryStat {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return MemoryStat{
		RSSBytes:          m.Sys,       // Total memory obtained from OS
		HeapUsedBytes:     m.HeapAlloc, // Heap memory in use
		HeapTotalBytes:    m.HeapInuse, // Heap memory in use by allocator
		HeapSysBytes:      m.HeapSys,   // Heap memory obtained from the system
		ExternalBytes:     0,           // Not separately tracked in Go
		ArrayBuffersBytes: 0,           // Not applicable in Go
	}
}
