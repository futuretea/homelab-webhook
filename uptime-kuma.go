package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UptimeKumaNotify struct {
	Heartbeat Heartbeat `json:"heartbeat"`
	Monitor   Monitor   `json:"monitor"`
	Msg       string    `json:"msg"`
}

type Heartbeat struct {
	MonitorID int64  `json:"monitorID"`
	Status    int64  `json:"status"`
	Time      string `json:"time"`
	Msg       string `json:"msg"`
	Important bool   `json:"important"`
	Duration  int64  `json:"duration"`
}

type Monitor struct {
	ID                  int64           `json:"id"`
	Name                string          `json:"name"`
	URL                 string          `json:"url"`
	Method              string          `json:"method"`
	Hostname            string          `json:"hostname"`
	Port                string          `json:"port"`
	Maxretries          int64           `json:"maxretries"`
	Weight              int64           `json:"weight"`
	Active              int64           `json:"active"`
	Type                string          `json:"type"`
	Interval            int64           `json:"interval"`
	RetryInterval       int64           `json:"retryInterval"`
	Keyword             interface{}     `json:"keyword"`
	ExpiryNotification  bool            `json:"expiryNotification"`
	IgnoreTLS           bool            `json:"ignoreTls"`
	UpsideDown          bool            `json:"upsideDown"`
	Maxredirects        int64           `json:"maxredirects"`
	AcceptedStatuscodes []string        `json:"accepted_statuscodes"`
	DNSResolveType      string          `json:"dns_resolve_type"`
	DNSResolveServer    string          `json:"dns_resolve_server"`
	DNSLastResult       interface{}     `json:"dns_last_result"`
	ProxyID             interface{}     `json:"proxyId"`
	NotificationIDList  map[string]bool `json:"notificationIDList"`
	Tags                []interface{}   `json:"tags"`
	MqttUsername        string          `json:"mqttUsername"`
	MqttPassword        string          `json:"mqttPassword"`
	MqttTopic           string          `json:"mqttTopic"`
	MqttSuccessMessage  string          `json:"mqttSuccessMessage"`
}

func uptimeKumaToServerchan(notify *UptimeKumaNotify) error {
	title := fmt.Sprintf("[%s]", notify.Monitor.Name)
	description := fmt.Sprintf(`
type: %s
hostname: %s
name: %s
msg:%s`, notify.Monitor.Type, notify.Monitor.Hostname, notify.Monitor.Name, notify.Msg)
	log.Debug().Str("title", title).Msg(notify.Msg)
	return sendToServerChan(title, description)
}

func uptimeKumaServerChanHandler(c *gin.Context) {
	var notify UptimeKumaNotify
	err := c.ShouldBind(&notify)
	if err != nil {
		log.Error().Err(err)
	}
	if err = uptimeKumaToServerchan(&notify); err != nil {
		log.Error().Err(err)
	}
	c.String(http.StatusOK, "")
	return
}
