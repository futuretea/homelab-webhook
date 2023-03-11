package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HarvesterNotify struct {
	Hostname string `json:"hostname"`
	IPAddrV4 string `json:"ipv4"`
	IPAddrV6 string `json:"ipv6"`
	MACAddr  string `json:"mac"`
	VIP      string `json:"vip"`
	Version  string `json:"version"`
	Msg      string `json:"msg"`
}

const harvesterServerChanPath = "/harvester-serverchan/:event"

func harvesterServerChanHandler(c *gin.Context) {
	event := c.Param("event")
	var notify HarvesterNotify
	err := c.ShouldBind(&notify)
	if err != nil {
		log.Error().Err(err)
	}
	if err = harvesterToServerchan(&notify, event); err != nil {
		log.Error().Err(err)
	}
	c.String(http.StatusOK, "")
	return
}

func harvesterToServerchan(notify *HarvesterNotify, event string) error {
	title := fmt.Sprintf("[%s install %s]", notify.Hostname, event)
	description := fmt.Sprintf(`
hostname: %s
event: %s
ipv4: %s
ipv6: %s
mac: %s`, notify.Hostname, event, notify.IPAddrV4, notify.IPAddrV6, notify.MACAddr)
	log.Debug().Str("title", title).Msg(description)
	return sendToServerChan(title, description)
}
