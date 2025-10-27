package main

import (
	"fmt"
	"sync"
	"time"
)

func Download(filename string, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	result <- fmt.Sprint(filename, " 下载完成")
}

func main() {
	var files []string = []string{"file1,zip", "file2,pdf", "file3.mp4"}
	wg := new(sync.WaitGroup)
	result := make(chan string, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Download(files[i], wg, result)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for massage := range result {
		fmt.Println(massage)
	}
}
