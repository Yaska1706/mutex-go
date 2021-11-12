package main

import (
	"runtime"
	"sync/atomic"
)

const (
	unlocked int32 = iota
	locked
)

type Locker interface {
	Lock()
	UnLock()
}

type Lock struct {
	state int32
}

func (l *Lock) Lock() {
	for !atomic.CompareAndSwapInt32(&l.state, unlocked, locked) {
		runtime.Gosched()
	}
}
func (l *Lock) UnLock() {
	atomic.StoreInt32(&l.state, unlocked)
}
