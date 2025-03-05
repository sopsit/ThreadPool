package Mutex

import (
	//"runtime"
	//"fmt"
	"sync/atomic"
)

type Mymutex struct {
	state uint32 // 0 = unlocked, 1 = locked
}

func (mut * Mymutex) Reset() {

	atomic.StoreUint32(&mut.state, 0)

}

func (mut * Mymutex) Lock() {
	//for {
	//	if atomic.CompareAndSwapUint32(&m.state, 0, 1) {
			// Successfully acquired the lock
		//	return
	//	}
		// Lock is currently held.  Yield the processor to allow other
		// goroutines to run.  This prevents a tight loop that consumes
		// CPU unnecessarily.
		//runtime.Gosched()
	//}

	for !atomic.CompareAndSwapUint32(&mut.state , 0 , 1) {
              // waite
	}
	//fmt.Println(mut.state)
}
func (mut *Mymutex) Unlock() {
	
	atomic.CompareAndSwapUint32(&mut.state , 1 , 0)// set the state to unlocked
	//fmt.Println(mut.state)
}

