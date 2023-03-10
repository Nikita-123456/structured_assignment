package myinterface 

import "sync"

type Api interface {
	IncViews(*sync.WaitGroup, *sync.Mutex)
	GetViews(*sync.WaitGroup, *sync.Mutex, chan int)
}
