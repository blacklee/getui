package getui

import (
	"fmt"
	"math/rand"
	"time"
)

var GTV2_DEBUG bool = false

type GeTuiSettings struct {
	Tts int64 `json:"tts"`
}

type GeTuiAudience struct {
	Cid []string `json:"cid"`
}

type GeTuiNotification struct {
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	ClickType string `json:"click_type,omitempty"`
	URL       string `json:"url,omitempty"`
}

type GeTuiPushMessage struct {
	Notification *GeTuiNotification `json:"notification"`
}

type GeTui_iOS_Alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GeTui_iOS_Aps struct {
	Alert *GeTui_iOS_Alert `json:"alert"`
	Sound string           `json:"sound"`
}

type GeTuiChannel_iOS struct {
	Type    string         `json:"type"`
	Payload string         `json:"payload,omitempty"`
	Aps     *GeTui_iOS_Aps `json:"aps"`
}

type GeTuiPushChannel struct {
	Ios *GeTuiChannel_iOS `json:"ios"`
}

func defaultSettings() *GeTuiSettings {
	setting := &GeTuiSettings{Tts: 360000}
	return setting
}

func singleAudience(cid string) *GeTuiAudience {
	audience := &GeTuiAudience{}
	audience.Cid = []string{cid}
	return audience
}

func newRequestID(reqid string, cfg GeTuiConfig) string {
	if len(reqid) > 10 {
		return reqid
	}
	return fmt.Sprintf("gtreq-%s-%d-%d", cfg.AppID[0:10], time.Now().UnixNano()%100000000, rand.Int63()%1000000)
}
