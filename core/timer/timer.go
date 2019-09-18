package timer

import "time"

//定时器
func SetTimer(duration time.Duration, f func()) {
	timer := time.NewTimer(duration)
	go func() {
		for {
			select {
			case <-timer.C:
				f()
				return //todo 定时器推出
			}
		}
	}()
}

func SetTicker(duration time.Duration, f func(), n int) {
	var i int
	go func(i int) {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ticker.C:
				f()
				if n > 0 {
					if n == i {
						ticker.Stop()
						return //退出携程
					}
					i++
				}
			}
		}
	}(i)
}
