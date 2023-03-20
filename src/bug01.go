package main

import "fmt"

/*
1. The program is using an unbuffered channel ch created using make(chan string), and immediately
tries to send the string "Hello world!" into the channel using the send operator <-. However, because
the channel is unbuffered, the send operation will block until another goroutine reads from the channel.
But in this case, there are no other goroutines to read from the channel, so the program will
deadlock and never reach the fmt.Println(<-ch) statement.

2. One way to fix the bug is to create a separate goroutine to read from the channel
before attempting to send to it. Here's the modified code:
*/
func main() {
	ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- "Hello world!"
}

/*
3. In the fixed code, a goroutine is spawned using the go keyword to read from the channel ch using fmt.Println(<-ch)
before the ch channel is written to using ch <- "Hello world!". This allows the program to avoid deadlock by ensuring
that there is a separate goroutine to read from the channel and unblock the send operation.
*/
