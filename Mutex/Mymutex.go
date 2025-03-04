package Mutex

import (
	"runtime"
	"sync/atomic"
)

type Mymutex struct {
	state uint32 // 0 = unlocked, 1 = locked
}

func (m * Mymutex) Lock() {
	for {
		if atomic.CompareAndSwapUint32(&m.state, 0, 1) {
			// Successfully acquired the lock
			return
		}
		// Lock is currently held.  Yield the processor to allow other
		// goroutines to run.  This prevents a tight loop that consumes
		// CPU unnecessarily.
		runtime.Gosched()
	}
}
func (m *Mymutex) Unlock() {
	atomic.StoreUint32(&m.state, 0) // set the state to unlocked
}

