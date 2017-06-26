package poloniex

import (
	"log"
	"time"

	"fmt"

	"gopkg.in/jcelliott/turnpike.v2"
)

func PoloTest() {
	turnpike.Debug()

	fmt.Println("start")
	c, err := turnpike.NewWebsocketClient(turnpike.JSON, "wss://api.poloniex.com", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("1")

	_, err = c.JoinRealm("realm1", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("2")

	if err := c.Subscribe("BTC_ETH", nil, func(args []interface{}, kwargs map[string]interface{}) {
		log.Println(args)
	}); err != nil {
		log.Fatalln("Error subscribing to chat channel:", err)
	}

	fmt.Println("3")
	go func() {
		c.Receive()
	}()

	time.Sleep(10 * 1)
	c.ReceiveDone <- true

	fmt.Println("fin")
	c.LeaveRealm()
	c.Close()

}
