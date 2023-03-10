package main

import (
	"fmt"
	api "nikita/ass1/pkg/api"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tables := []struct {
		desc  string
		id    string
		views int
		itr   int
	}{
		{desc: "test-case-1",
			id:    "123",
			views: 0,
			itr:   0,
		},
		{desc: "test-case-2",
			id:    "1234",
			views: 10,
			itr:   10,
		},
		{desc: "test-case-3",
			id:    "12345",
			views: 5,
			itr:   20,
		},
	}
	wt := &sync.WaitGroup{}
	lc := &sync.Mutex{}
	for _, test := range tables {
		video := api.NewVideo(test.id)
		t.Run(test.desc, func(t *testing.T) {
			for i := 0; i < test.itr; i++ {
				wt.Add(1)
				go video.IncViews(wt, lc)
			}
			wt.Wait()

			ch := make(chan int)
			wt.Add(1)
			go video.GetViews(wt, lc, ch)
			views := <-ch
			if views != test.views {
				t.Fatalf("expected %d views but getting %d", test.views, views)
			} else {
				fmt.Println("Right !! ")
			}
		})
	}

}
