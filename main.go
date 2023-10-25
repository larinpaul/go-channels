//// 2023/10/25 // 19:17 //

//// Ping Pong

// Many tech companies play ping pong in the office during break time. At Mailio,
// we play virtual ping pong using channels.

// // Assignment

// // Fix the bug in the pingPong function. It shouldn't return until the entire game
// // of ping pong is complete.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func pingPong(numPings int) {
// 	pings := make(chan struct{})
// 	pongs := make(chan struct{})
// 	go ponger(pings, pongs)
// 	pinger(pings, pongs, numPings)
// }

// // TEST SUITE:

// func pinger(pings, pongs chan struct{}, numPings int) {
// 	go func() {
// 		sleepTime := 50 * time.Millisecond
// 		for i := 0; i < numPings; i++ {
// 			fmt.Println("ping", i, "sent")
// 			pings <- struct{}{}
// 			time.Sleep(sleepTime)
// 			sleepTime *= 2
// 		}
// 		close(pings)
// 	}()
// 	i := 0
// 	for range pongs {
// 		fmt.Println("pong", i, "got")
// 		i++
// 	}
// 	fmt.Println("pongs done")
// }

// func ponger(pings, pongs chan struct{}) {
// 	i := 0
// 	for range pings {
// 		fmt.Println("ping", i, "got", "pong", i, "sent")
// 		pongs <- struct{}{}
// 		i++
// 	}
// 	fmt.Println("pings done")
// 	close(pongs)
// }

// func test(numPings int) {
// 	fmt.Println("Starting game...")
// 	pingPong(numPings)
// 	fmt.Println("===== Game over =====")
// }

// func main() {
// 	test(4)
// 	test(3)
// 	test(2)
// }

//// 19:23
//// Concurrency

// What is concurrency?

// Concurrency is the ability to perform multiple tasks at the same time. Typically,
// out code is executed one line at a time, one after the other. This is called
// sequential exectuion or synchronous exection.

// If the compulet we're running our code on has multipe cores, we can even
// execute multiple tasks at exactly the same time. If we're running on a single
// core, a single code executes code at almost the same time by switching
// between tasks very quickly. Either way, the code we write loks the same in Go
// and takes advantage of whatever resourvces are available.

// How does concurrency work in Go?

// Go was designed to be concurrent, which is a trai fairly unique to Go. It excels
// at performing many tasks simultaneously safely using a simpole syntax.

// There isn't a popilar programming language in existence where spawning
// concurrent execution is quite as elegant, at least in my opinion.

// // Concurrency is as simple as using the go keyword when calling a function:
// go doSomething()

// In the example above, doSomething() will be executed concurrently with the
// rest of the code in the function. The go keyword is used to spawn a new
// goroutine.

// //// Assignment

// // At Mailio we send a lot of network requests. Each email we sent must go out
// // over the internet. To server our millions of customers, we need a single Go
// // program to be capable of sending thousands of emails at once.

// // Edit the sendEmail() function to execute its anonymous function concurrently
// // so that the "received" message prints after the "sent" message.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendEmail(message string) {
// 	go func() {
// 		time.Sleep(time.Millisecond * 250)
// 		fmt.Printf("Email received: '%s'\n", message)
// 	}()
// 	fmt.Printf("Email sent: '%s'\n", message)
// }

// func test(message string) {
// 	sendEmail(message)
// 	time.Sleep(time.Millisecond * 500)
// 	fmt.Println("==========================")
// }

// func main() {
// 	test("Hello there Stacy!")
// 	test("Hi there John!")
// 	test("Hey there Jane!")
// }

//// 19:30
//// Channels (channels deadlock)

// Channels are a typed, thread-safe queue. Channels allow different goroutines to
// communicate with each other.

// Create a channel

// // Like maps and slices, channels must be created before use. They also use the
// // same make keyword:
// ch := make(chan int)

// // Send data to a channel
// cn <- 69

// The <- operator is called the channel operator. Data flows in the direction of
// the arrow. This operation will block until another goroutine is ready to receive
// the value.

// // Receive fata from a channel
// v := <-ch

// This reads and removes a value from the channel and saves it into the variable
// v. This operation will block until there is a value in the channel to be read.

// Blocking and deadlocks
// A deadlock is when a group of goroutines are all blocking so none of them can
// continue. This is a common bug that you need to watch out for in concurrent
// programming.

// Assignment

// Run the program. You'll see that it deadlocks and never exists. The
// filterOldEmails function is trying to send on a channel and there is no other
// goroutine running that can accept the value from the channel.

// Fix the deadlock by spawning an anonymous goroutine to write to the
// isOldChan channel isntead of using the same goroutine that's reading from it.

package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

// TEST SUITE

type email struct {
	body string
	date time.Time
}

func test(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

func main() {
	test([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Today is the day!",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Letsa Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}
