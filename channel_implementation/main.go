package main

import (
	"fmt"
	"net/http"
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

	fmt.Println(<-c) //only prints google
	//receiving a message from a channel is a blocking like of code.
	/*As the main routine (sleeping) is waiting for a message from the
	channel, first resolved link will execute the full function,
	completed go routine ask whether any channel waiting to print
	then this executes and program exits.*/
	//so to solve it you write len(links) times 'fmt.Println(<-c)'

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
