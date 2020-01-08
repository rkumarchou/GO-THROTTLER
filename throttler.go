package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)

var taskList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func throttler(throttleLimit int) {
  var wg sync.WaitGroup
  for count := 0; count <= throttleLimit; count++ {
    wg.Add(1)
    go toDoTask(taskList[count], &wg)
  }
  wg.Wait()
}

func main() {
  throttler(rand.Intn(10))
}

func toDoTask(id int, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Printf("Task %d starting\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Task %d done\n", id)
}