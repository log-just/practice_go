package util

// Init : 초기화!
func Init() bool {
	//return fmt.Errorf("math: square root of negative number %s", "oh")
	ch := make(chan bool, 1)
	go func() { ch <- logInit() }()
	//go func() { ch <- redisHealth() }()
	//go func() { ch <- gcacheHealth() }()
	// // wait for goroutines to finish
	for i := 0; i < 1; i++ {
		if !<-ch {
			return false
		}
	}
	return true
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
