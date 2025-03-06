package Mutex

import (
	"sync/atomic"
)

type Mymutex struct {
	state uint32 // 0 = unlocked, 1 = locked
}


func (mut * Mymutex) Reset() {

	atomic.StoreUint32(&mut.state, 0)

}

func NMutex () *Mymutex {
	var tempmutex Mymutex
	tempmutex.Reset()
	return &tempmutex
}

func (mut * Mymutex) Lock() {

	for !atomic.CompareAndSwapUint32(&mut.state , 0 , 1) {
              // wait
	}
	//fmt.Println(mut.state)
}
func (mut *Mymutex) Unlock() {
	
	atomic.CompareAndSwapUint32(&mut.state , 1 , 0)// set the state to unlocked
	//fmt.Println(mut.state)
}

