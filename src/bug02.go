package main

import (
	"fmt"
	"time"
)

// // This program should go to 11, but it seemingly only prints 1 to 10.
// func main() {
// 	ch := make(chan int)
// 	go Print(ch)
// 	for i := 1; i <= 11; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// // Print prints all numbers sent on the channel.
// // The function returns when the channel is closed.
// func Print(ch <-chan int) {
// 	for n := range ch { // reads from channel until it's closed
// 		time.Sleep(10 * time.Millisecond) // simulate processing time
// 		fmt.Println(n)
// 	}
// }

/*
1. The bug in the code is that the Print function will only print the first 10 numbers sent on the
channel, instead of all 11 as expected. This is because the main function closes the channel
after sending 11 numbers, causing the Print function to exit before it can process the final value sent on the channel.
*/

// 2. To fix the bug, we can modify the Print function to check if the channel is closed before reading from it. Here's the modified code:
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(100 * time.Millisecond) // wait for Print to finish
}

func Print(ch <-chan int) {
	for {
		n, ok := <-ch
		if !ok { // channel is closed
			return
		}
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
	}
}

/*
3. In the fixed code, we use a for loop to continually read from the channel until it's closed, rather than
relying on the range syntax. We also use a boolean variable ok to check if the channel is closed, and return
from the function if it is. Finally, we add a time.Sleep statement in the main function to wait for the Print
function to finish before exiting the program.
*/
