package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	log.SetFlags(log.Lmicroseconds)

	c := cron.New()
	c.AddFunc("* * * * *", func() { go httpClient() })
	c.Start()
	time.Sleep(time.Second * 70)

	return

}

func httpClient() {
	url := "https://us-central1-striped-proxy-187410.cloudfunctions.net/get-time1"

	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	//fmt.Println("gorutine complete")
}
