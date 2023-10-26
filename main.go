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

// // Assignment

// // Run the program. You'll see that it deadlocks and never exists. The
// // filterOldEmails function is trying to send on a channel and there is no other
// // goroutine running that can accept the value from the channel.

// // Fix the deadlock by spawning an anonymous goroutine to write to the
// // isOldChan channel isntead of using the same goroutine that's reading from it.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func filterOldEmails(emails []email) {
// 	isOldChan := make(chan bool)

// 	go func() {
// 		for _, e := range emails {
// 			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
// 				isOldChan <- true
// 				continue
// 			}
// 			isOldChan <- false
// 		}
// 	}()

// 	isOld := <-isOldChan
// 	fmt.Println("email 1 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 2 is old:", isOld)
// 	isOld = <-isOldChan
// 	fmt.Println("email 3 is old:", isOld)
// }

// // TEST SUITE

// type email struct {
// 	body string
// 	date time.Time
// }

// func test(emails []email) {
// 	filterOldEmails(emails)
// 	fmt.Println("==========================================")
// }

// func main() {
// 	test([]email{
// 		{
// 			body: "Are you going to make it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "I need a break",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What were you thinking?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Yo are you okay?",
// 			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Have you heard of that website Boot.dev?",
// 			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "It's awesome honestly.",
// 			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Today is the day!",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "What do you want for lunch?",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Why are you the way that you are?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// 	test([]email{
// 		{
// 			body: "Did we do it?",
// 			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Letsa Go!",
// 			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 		{
// 			body: "Okay...?",
// 			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
// 		},
// 	})
// }

//// 2023/10/26 // 17:40 //
//// Channels
//// 3 channels send

// Empty structs are often used as tokes in Go programs. In this context, a token is a
// unary value. In other words, we don't care what is passsed through the channel. We care
// when and if it is passed.

// // We can block and wait until something is sent on a channel using the following syntax
// <-ch

// This will block until it pops a single item off the channel, then continue, discarding the
// item.

// //// Assignment

// // Out Mailio server isn't able to boot up until it receives the singal that its databases are
// // all online, and it learns about them being online by waiting for tokens (empty structs)
// // on a channe.

// // Complete the waitForDbs function. It should block until it receives numDBs tokens on
// // the dbChan channel. Each time it reads a token the getDatabasesChannel goroutine will
// // print a message to the console for you.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func waitForDbs(numDBs int, dbChan chan struct{}) {
// 	for i := 0; i < numDBs; i++ {
// 		<-dbChan
// 	}
// }

// // don't touch below this line

// func test(numDBs int) {
// 	dbChan := getDatabasesChannel(numDBs)
// 	fmt.Printf("Waiting for %v databases...\n", numDBs)
// 	waitForDbs(numDBs, dbChan)
// 	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
// 	fmt.Println("All databases are online!")
// 	fmt.Println("=====================================")
// }

// func main() {
// 	test(3)
// 	test(4)
// 	test(5)
// }

// func getDatabasesChannel(numDBs int) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		for i := 0; i < numDBs; i++ {
// 			ch <- struct{}{}
// 			fmt.Printf("Database %v is online\n", i+1)
// 		}
// 	}()
// 	return ch
// }

//// 17:53
//// Buffered Channels

// Channels can optinally be buffered.

// Creating a channel with a buffer

// // You can provide a buffer length as the second argument to make() to create a buffered
// // channel:
// ch := make(chan int, 100)

// Sending on a buffered channel only blocks when the buffer is full.

// Receiving blocks only when the buffer is empty.

// Assignment

// We want to be able to send emails in batches. A writing goroutine will write an entire
// batch of email messages to a buffered channel, and later, once the channel is full, a
// reading goroutine will read all of the messages from the channel and send them out to
// out clients.

// // Complete the addEmailsToQueue function. It should create a buffered channel with a
// // buffer large enough to store all of the emails it's given. It should then write the emails
// // to the chnnel in order, and finally return the channel.

// package main

// import "fmt"

// func addEmailsToQueue(emails []string) chan string {
// 	emailsToSend := make(chan string, len(emails))
// 	for _, email := range emails {
// 		emailsToSend <- email
// 	}
// 	return emailsToSend
// }

// // TEST SUITE

// func sendEmails(batchSize int, ch chan string) {
// 	for i := 0; i < batchSize; i++ {
// 		email := <-ch
// 		fmt.Println("Sending email:", email)
// 	}
// }

// func test(emails ...string) {
// 	fmt.Printf("Adding %v emails to queue...\n", len(emails))
// 	ch := addEmailsToQueue(emails)
// 	fmt.Println("Sending emails...")
// 	sendEmails(len(emails), ch)
// 	fmt.Println("==========================================")
// }

// func main() {
// 	test("Hello John, tell Kathy I said hi", "Whazzup bruther")
// 	test("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
// 	test("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
// }

//// 18:10

//// Closing channels in Go

// // Channels can be explicitly closed by a sende:
// ch := make(chan int)
// // do some stuff with the channel
// close(ch)

// // Checking if a channel is closed

// // Similar to the ok value when accessing data in a map, receivers can check the ok
// // value when receiving form a channel to test if a channel was closed.

// v, ok := <-ch

// // ok is false if the channel is empty and closed.

// // Don't send on a closed channel

// Sending on a closed channel will cause a panic. A panic on the main goroutine will
// cause the entire program to crash, and a panic in any other goroutine will cause that
// goroutine to crash.

// Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still
// be garbage collected if they're unused. You should close channels to indicate explicitly
// to a receiver that nothing else is going to come across.

//// Assignment

// At Mailio we're all about keeping track of what our systems are up to with great logging
// and telemetry.

// The sendReports function sends out a batch of reports to our clients and reports back
// how many were sent across a channel. It closes the channel when it's doen.

// // Complete the countReports function. It should:
// // * Use an infinite for loop to read from the channel:
// // * If the channel is closed, break out of the loop
// // * Otherwise, keep a running total of the number of reports sent
// // * Return the total number of reports sent

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func countReports(numSentCh chan int) int {
// 	total := 0
// 	for {
// 		numSent, ok := <-numSentCh
// 		if !ok {
// 			break
// 		}
// 		total += numSent
// 	}
// 	return total
// }

// // TEST SUITE

// func test(numBatches int) {
// 	numSentCh := make(chan int)
// 	go sendReports(numBatches, numSentCh)

// 	fmt.Println("Start counting...")
// 	numReports := countReports(numSentCh)
// 	fmt.Printf("%v reports sent!\n", numReports)
// 	fmt.Println("=========================")
// }

// func main() {
// 	test(3)
// 	test(4)
// 	test(5)
// 	test(6)
// }

// func sendReports(numBatches int, ch chan int) {
// 	for i := 0; i < numBatches; i++ {
// 		numReports := i*23 + 32%17
// 		ch <- numReports
// 		fmt.Printf("Sent batch of %v reports\n", numReports)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	close(ch)
// }

//// 18:28
//// Range

// // Similar to slices and maps, channels can be ranged over.

// for item := range ch {
// 	// item is the next value received from the channel
// }

// This example will receive values over the channel (blocking at each iteration of nothing
// new is there) and will exit only when the channel is closed.

// Assignment

// It's that time again, Mailio is hiring and we've been assigned to do the interview. For
// some reason, the Fibonacci sequence is Mailio's interview problem of choice. We've
// been tasked with building a small toy program we can use in the interview.

// // Complete the concurrentFib function. It should:
// // * Create a new channel of ints
// // * Call fibonacci in a goroutine, passing it the channel and the number of Fibonacci
// // numbers to generate, n
// // * Use a range loop to read from the channel and print out the numbers one by one,
// // each on a new line

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func concurrentFib(n int) {
// 	ch := make(chan int)
// 	go fibonacci(n, ch)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }

// // TEST SUITE

// func test(n int) {
// 	fmt.Printf("Printing %v numbers...\n", n)
// 	concurrentFib(n)
// 	fmt.Println("===============================")
// }

// func main() {
// 	test(10)
// 	test(5)
// 	test(20)
// 	test(13)
// }

// func fibonacci(n int, ch chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 		time.Sleep(time.Millisecond * 10)
// 	}
// 	close(ch)
// }

//// 18:40
//// Select

// Sometiemes we have a single goroutine listening to multiple channels and want to
// process data in the order it coms through each channel.

// // A select statement is used to listen to multiple channels at the same time. It is similar
// // to a switch statement but for channels.
// select {
// case i, ok := <-chInts:
// 	fmt.Println(i)
// case s, ok := <-chStrings:
// 	fmt.Println(s)
// }

// The first channel with a value ready to be received will fire and its body will execute. If
// multiple channels are ready at the same time one is chosen randomly. The ok variable
// in the example above refers to whether or not the channel has been closed by the
// sender yet.

// Assignment

// // Complete the logMessages function.

// // Use an infinite for loop and a select statement to log the emails and sms messages as
// // they come in order across the two channels. Add a condition to return from the
// // function when one of the two channels closes, whichever is first.

// // Use the logSms and logEmail functions to log the messages.

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func logMessages(chEmails, chSms chan string) {
// 	for {
// 		select {
// 		case email, ok := <-chEmails:
// 			if !ok {
// 				return
// 			}
// 			logEmail(email)
// 		case sms, ok := <-chSms:
// 			if !ok {
// 				return
// 			}
// 			logSms(sms)
// 		}
// 	}
// }

// // TEST SUITE - Don't touch below this line

// func logSms(sms string) {
// 	fmt.Println("SMS:", sms)
// }

// func logEmail(email string) {
// 	fmt.Println("Email:", email)
// }

// func test(sms []string, emails []string) {
// 	fmt.Println("Starting...")

// 	chSms, chEmails := sendToLogger(sms, emails)

// 	logMessages(chEmails, chSms)
// 	fmt.Println("===============================")
// }

// func main() {
// 	rand.Seed(0)
// 	test(
// 		[]string{
// 			"hi friend",
// 			"What's going on?",
// 			"Welcome to the business",
// 			"I'll pay you to be my friend",
// 		},
// 		[]string{
// 			"Will you make your appointment?",
// 			"Let's be friends",
// 			"What are you doing?",
// 			"I can't believe you've done this.",
// 		},
// 	)
// 	test(
// 		[]string{
// 			"this song slaps hard",
// 			"yooo hoooo",
// 			"i'm a big fan",
// 		},
// 		[]string{
// 			"What do you think of this song?",
// 			"I hate this band",
// 			"Can you believe this song?",
// 		},
// 	)
// }

// func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
// 	chSms = make(chan string)
// 	chEmails = make(chan string)
// 	go func() {
// 		for i := 0; i < len(sms) && i < len(emails); i++ {
// 			done := make(chan struct{})
// 			s := sms[i]
// 			e := emails[i]
// 			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
// 			go func() {
// 				time.Sleep(t1)
// 				chSms <- s
// 				done <- struct{}{}
// 			}()
// 			go func() {
// 				time.Sleep(t2)
// 				chEmails <- e
// 				done <- struct{}{}
// 			}()
// 			<-done
// 			<-done
// 			time.Sleep(10 * time.Millisecond)
// 		}
// 		close(chSms)
// 		close(chEmails)
// 	}()
// 	return chSms, chEmails
// }

//// 19:07
//// Select Default Case

// // The default case in a select statement executes immediately if no other channel has
// // a value ready. A default case stops the select statement from blocking.

// select {
// 	case v := <-ch:
// 		// use v
// 	default:
// 		// receiving form ch would block
// 		// so do something else
// }

// Tickets
// * time.Tick() is a standard library function that returns a channel that sends a value
// on a given interval.
// * time.After() sends a value once after the duration has passed.
// * time.Sleep() blocks the current goroutine for the specific amount of time.

// // Read-only Channels

// // A channel can be marked as read-only by casting it from a chan to a <-chan type.
// // For example:
// func main() {
// 	ch := make(chan int)
// 	readCh(ch)
// }
// func readCh(ch <-chan int) {
// 	// ch can only be read from
// 	// in this function
// }

// // Write-only Channels

// // The same goes for write-only channels, but the arrow's position moves.
// func writeCh(ch chan<- int) {
// 	// ch can only be written to
// 	// in this function
// }

//// Assignment

// Like all good back-end engineers, we frequently save backup snapshots of the Mailio
// database.

// Complete the saveBackup function.

// It should read values from the snapshotTicker and saveAfter channels
// simultaneously.
// * If a value is received from snapshotTicker, call takeSnapshot()
// * If a value is received from saveAfter, call saveSnapshot() and return from the
// function: you're done.
// * If neither channel has a value ready, call waitForData() and then time.Sleep() for
// 500 milliseconds. After all, we want to show in the logs that the snapshot service is
// running.

package main

import (
	"fmt"
	"time"
)

func saveBackup(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// TEST SUITE

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backup saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackup(snapshotTicker, saveAfter)
	fmt.Println("==========================")
}

func main() {
	test()
}
