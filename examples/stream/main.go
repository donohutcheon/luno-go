package main

import (
	"log"
	"time"

	"github.com/luno/luno-go/streaming"
)

func printOrderBook(c *streaming.Conn) {
	ss := c.Snapshot()
	log.Printf("%d: %s %v %v\n", ss.Sequence, ss.Status, ss.Bids[0], ss.Asks[0])
}

func printUpdate(u streaming.Update) {
	log.Printf("%d: %s %v %v\n", u.Sequence, u.)
}

func main() {
	keyID := "key here"
	keySecret := "key secret here"

	c, err := streaming.Dial(keyID, keySecret, "XBTZAR", streaming.WithConnectCallback(printOrderBook),
		streaming.WithUpdateCallback(printUpdate))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		time.Sleep(time.Minute)
		printOrderBook(c)
	}
}
