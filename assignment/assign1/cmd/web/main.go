package main

import (
	"fmt"
	api "nikita/ass1/pkg/api"
	myinterface "nikita/ass1/pkg/interfaces"
	"sync"
)

func main() {

	var vid myinterface.Api
	vid = api.NewVideo("1234")

	lc := &sync.Mutex{}
	wt := &sync.WaitGroup{}

	//calling go routines for  video 100 times
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go vid.IncViews(wt, lc)

	}

	wt.Wait()

	c := make(chan int)

	wt.Add(1)

	go vid.GetViews(wt, lc, c)

	fmt.Println("Views on video1 - ", <-c)

	wt.Wait()

}
