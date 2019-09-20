package main

func aboutChannels() {
	ch := make(chan int) // ch has type 'chan int'

	// A channel has two principal operations - send and receive
	// also supports close, which sets a flag indicating that
	// no more values will ever be sent on this channel.
	ch <- x // a send statement

	x = <-ch  // a receive expression in an assignment statement
	<-ch      // a receive statement with discarded result
	close(ch) // closes channel
}

func bufOrUnbufChannel() {
	// channel created with a simple call to make is called an unbuffered channel
	// but make(chan accepts an optional second argument, an integer called channel's capacity
	ch = make(chan int)    // unbuffered channel
	ch = make(chan int, 0) // unbuffered channel
	ch = make(chan int, 3) // buffered channel with capacity 3
}
