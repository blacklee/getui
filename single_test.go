package getui

import (
	"fmt"
	"testing"
)

func TestPushSingle(t *testing.T) {
	cfg := loadAppConfig()
	hash := loadTestInfo()
	token := hash["token"].(string)
	cids := hash["cid"].([]string)
	cid := cids[0]
	notification := GeTuiNotification{
		Title:     "Hello title",
		Body:      "Hello body",
		ClickType: "none",
	}
	ios := GeTuiChannel_iOS{
		Type: "notify",
		Aps: &GeTui_iOS_Aps{
			Sound: "default",
			Alert: &GeTui_iOS_Alert{
				Title: "hello title",
				Body:  "hello body",
			},
		},
	}
	channel := GeTuiPushChannel{Ios: &ios}
	msg := GeTuiPushMessage{Notification: &notification}
	err := PushSingle(token, "", cid, cfg, msg, channel)
	fmt.Println(err)
}
