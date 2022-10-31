package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	log.SetFlags(log.Lmicroseconds)

	c := cron.New()
	c.AddFunc("* * * * *", func() { log.Println("every 1m") })
	c.Start()
	time.Sleep(3 * time.Minute)

	return

}
