package main

import (
	"fmt"
	"sync"
)

type Stats struct {
	mu       sync.Mutex
	counters map[string]int
}

// Snapshot returns the current stats.
func (s *Stats) SnapshotBad() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}

func (s *Stats) SnapshotGood() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

func main() {

	fmt.Println("start test bad case")
	sBad := Stats{
		counters: map[string]int{
			"a": 1,
			"b": 2,
		},
	}

	// snapshot is no longer protected by the mutex, so any
	// access to the snapshot is subject to data races.
	snapshotBad := sBad.SnapshotBad()
	fmt.Printf("snapshotBad: %v, Stats.counter : %v \n ", snapshotBad, sBad.counters)
	//change snapshotBad
	snapshotBad["a"] = 2
	fmt.Printf("snapshotBad: %v, Stats.counter : %v \n", snapshotBad, sBad.counters)
	fmt.Println("this is bad, both variable are changed")

	fmt.Println("start test good case")
	sGood := Stats{
		counters: map[string]int{
			"a": 1,
			"b": 2,
		},
	}

	// snapshot is no longer protected by the mutex, so any
	// access to the snapshot is subject to data races.
	snapshotGood := sGood.SnapshotGood()
	fmt.Printf("snapshotBad: %v, Stats.counter : %v \n ", snapshotGood, sGood.counters)
	//change snapshotBad
	snapshotGood["a"] = 2
	fmt.Printf("snapshotBad: %v, Stats.counter : %v \n", snapshotGood, sGood.counters)
	fmt.Println("this is good case")
}
