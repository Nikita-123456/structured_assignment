package api

import "sync"

type Videos struct {
	Id    string
	Views int
}

func NewVideo(id string) *Videos {
	return &Videos{
		Id:    id,
		Views: 0,
	}
}

// increase the views
func (v *Videos) IncViews(wt *sync.WaitGroup, lc *sync.Mutex) {
	defer wt.Done()
	lc.Lock()
	v.Views++
	lc.Unlock()
}

// get the views
func (v *Videos) GetViews(wt *sync.WaitGroup, lc *sync.Mutex, c chan int) {
	defer wt.Done()
	lc.Lock()
	c <- v.Views
	lc.Unlock()
}
