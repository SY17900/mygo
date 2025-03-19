package exam

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AddNum() {
	var count int32
	var wg sync.WaitGroup
	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddInt32(&count, int32(1))
			}
		}()
	}

	wg.Wait()
	fmt.Println("expected: 10000, get:", count)
}

func PrintString(count int) {
	str1, str2 := "shima", "kaze"
	ch12, ch21 := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for range count {
			<-ch21
			fmt.Print(str1)
			ch12 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := range count {
			<-ch12
			fmt.Println(str2)
			if i < count-1 {
				ch21 <- struct{}{}
			}
		}
	}()

	ch21 <- struct{}{}
	wg.Wait()
}
