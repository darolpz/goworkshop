package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	channel := make(chan string,5)
	var waitGroup,wg2 sync.WaitGroup
	waitGroup.Add(5)
	wg2.Add(5)
	go func (job <- chan string){
		for word := range job{
			for _,letter := range word{
				fmt.Printf("%c\n",letter)
				time.Sleep(2 * time.Second)
			}
			waitGroup.Done()
		}
	}(channel)

	go func(){
		for{
			var msg string
			fmt.Scanln(&msg)
			channel<-msg
			wg2.Done()
		}
	}()
	/* var msg string
	for i := 0; i<5;i++{
		fmt.Scanln(&msg)
			channel<-msg
	} */
	wg2.Wait()
	fmt.Printf("Fin del loop\n")
	waitGroup.Wait()
	close(channel)
	fmt.Printf("Fin del programa\n")
}