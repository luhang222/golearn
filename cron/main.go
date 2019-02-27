package main

import (
	"fmt"
	"time"
)

//检查间隔(秒)
const INTERVAL = 1

func main() {
	timer(run)
}

func timer(timer func()) {
	timer()
	ticker := time.NewTicker(INTERVAL * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			timer()
		}
	}
}

func run() {
	fmt.Println(time.Now().Second(), "1111111111")
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now().Second(), "2222222")
}
