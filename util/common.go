package util

import "golang.org/x/sync/errgroup"

// Init : 초기화!
func Init() bool {

	// errgroup을 사용해보자..!
	// https://godoc.org/golang.org/x/sync/errgroup
	g := new(errgroup.Group)
	g.Go(initLog)
	if err := g.Wait(); err != nil {
		return false
	}
	return true

	// ch := make(chan bool, 1)
	// go func() { ch <- initLog() }()
	// for i := 0; i < 1; i++ {
	// 	if !<-ch {
	// 		return false
	// 	}
	// }
	// return true
}

// Health : 헬스체크 from 헬스체크 api
func Health() bool {
	//ch := make(chan bool, 1)
	//go func() { ch <- logInit() }()
	//go func() { ch <- redisHealth() }()
	//go func() { ch <- gcacheHealth() }()

	// wait for goroutines to finish
	//for i := 0; i < 1; i++ {
	//	if !<-ch {
	//		return false
	//	}
	//}
	return true
}
