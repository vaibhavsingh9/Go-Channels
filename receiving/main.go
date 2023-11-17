package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	/*for {
	//infinite execution of pinging websites
	go checkLink(<-c /*string*/ /*, c)
	}*/
	//refactoring
	//for l := range c {
	//	//time.Sleep(time.Second /*1second_pause*/ /*int64*/)
	//	/*not appropriate place to declare as will stop the
	//	main routine and many messages will get queued in buffer.*/
	//	go func() {
	//		time.Sleep(5 * time.Second)
	//		checkLink(l, c)
	//	}() //function literal
	//}only fetches one value from slice after the first fetch.
	//same variable inside two different routines requires us
	//to refactor the program.
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //creates a copy in memory,
		//even if the value of l changes
		//the previous values won't be effected.
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
