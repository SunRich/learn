package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	minDelay time.Duration = time.Second
	avgDelay time.Duration
	maxDelay time.Duration

	thStart int
	thStop  int

	nM = fmt.Sprintf("ps M %d | wc -l", os.Getpid())
	nG = 100
)

func main() {
	thStart, err := getThreads()
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	// 让所有 goroutine 都休眠到当前的五秒钟以后
	t := time.Now()
	run := t.Add(time.Second * 5)

	wg := sync.WaitGroup{}
	wg.Add(nG)
	for i := 0; i < nG; i++ {
		go func() {
			// 计算时间误差，将 timer 设置为修正后的时间
			remain := run.Sub(time.Now())
			if remain > 0 {
				timer := time.NewTimer(remain)
				<-timer.C
			}

			// 统计
			delay := time.Now().Sub(run)
			if delay < minDelay {
				minDelay = delay
			}
			if delay > maxDelay {
				maxDelay = delay
			}
			avgDelay += delay
			wg.Done()
		}()
	}
	wg.Wait()

	thStop, err = getThreads()
	if err != nil {
		panic(err)
	}
	fmt.Printf("nG: %v, nM: %v->%v, min delay: %v, max delay: %v, avg delay: %v\n", nG, thStart, thStop, minDelay, maxDelay, avgDelay/time.Duration(nG))
}

// getThreads returns the number of running threads
func getThreads() (int, error) {
	out, err := exec.Command("bash", "-c", nM).Output()
	if err != nil {
		return 0, fmt.Errorf("M: failed to fetch #threads: %v", err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, fmt.Errorf("M: failed to parse #threads: %v", err)
	}
	return n, nil
}
