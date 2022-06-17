package main

import (
	"fmt"
	"os"

	"github.com/guonaihong/gout"
)

func sendToServerChan(title, description string) error {
	url := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", os.Getenv("SERVER_CHAN_SEND_KEY"))
	return gout.GET(url).Debug(false).SetQuery(gout.H{"title": title, "desp": description}).Do()
}
