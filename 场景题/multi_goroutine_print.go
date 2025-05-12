package 场景题

import (
	"fmt"
	"sync"
	"testing"
)

// TestGChanPrint 两个协程依次打印121212
// 思路：管道控制2个携程的打印信号，记得close管道
func TestGChanPrint(t *testing.T) {
	// 1.创建无缓冲区的管道
	// 控制奇数打印管道
	c1 := make(chan int)
	// 控制偶数打印管道
	c2 := make(chan int)
	// 控制打印结果
	n := 100

	// 开2个g
	var wg sync.WaitGroup
	wg.Add(2)

	// 2.启动g1
	go func() {
		defer wg.Done()
		for i := 1; i < n; i += 2 {
			// g1接受打印指令
			<-c1
			t.Logf("g1=%v", i)
			// 告诉g2可以打印了
			c2 <- 1
		}
		// 关闭c1
		close(c1)
	}()
	// 3.启动g2
	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			// g2接受打印指令
			<-c2
			t.Logf("g2=%v", i)
			// 告诉g1可以打印了
			if i < n {
				// c1已关闭了，不用写了
				c1 <- 1
			}
		}
		// 关闭c2
		close(c2)
	}()

	// g1先写
	c1 <- 1

	wg.Wait()
}

// TestGChanPrintAB 两个协程交替打印ABABAB。。。
// 思路：还是管道，记得close管道
func TestGChanPrintAB(t *testing.T) {
	// 1.创建无缓冲区的管道
	// 控制奇数打印管道
	c1 := make(chan int)
	// 控制偶数打印管道
	c2 := make(chan int)
	// 控制打印结果
	n := 100

	// 开2个g
	var wg sync.WaitGroup
	wg.Add(2)

	// 2.启动g1
	go func() {
		defer wg.Done()
		for i := 0; i <= n; i++ {
			// g1接受打印信号
			<-c1
			fmt.Printf("A")
			// 告诉g2可以打印了
			c2 <- 1
		}
		// 关闭c1
		close(c1)
	}()
	// 3.启动g2
	go func() {
		defer wg.Done()
		for i := 0; i <= n; i++ {
			// 接受g2可以打印的信号
			<-c2
			fmt.Printf("B")
			// 告诉g1可以打印了
			if i < n {
				c1 <- 1
			}
		}
		// 关闭c2
		close(c2)
	}()

	// g1先写
	c1 <- 1

	wg.Wait()
}
