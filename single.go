package getui

import (
	"encoding/json"
	"fmt"
)

type GeTuiSingleMessage struct {
	RequestID   string            `json:"request_id"`
	Settings    *GeTuiSettings    `json:"settings"`
	Audience    *GeTuiAudience    `json:"audience"`
	PushMessage *GeTuiPushMessage `json:"push_message"`
	PushChannel *GeTuiPushChannel `json:"push_channel"`
}

func CreateGeTuiSingleMessage(reqid, cid string, cfg GeTuiConfig, title, body, clickType string) *GeTuiSingleMessage {
	reqid = newRequestID(reqid, cfg)
	singleMsg := &GeTuiSingleMessage{RequestID: reqid}
	singleMsg.Settings = defaultSettings()
	singleMsg.Audience = singleAudience(cid)

	notification := GeTuiNotification{
		Title:     title,
		Body:      body,
		ClickType: clickType,
	}
	ios := GeTuiChannel_iOS{
		Type: "notify",
		Aps: &GeTui_iOS_Aps{
			Sound: "default",
			Alert: &GeTui_iOS_Alert{
				Title: title,
				Body:  body,
			},
		},
	}
	channel := GeTuiPushChannel{Ios: &ios}
	msg := &GeTuiPushMessage{Notification: &notification}

	singleMsg.PushMessage = msg
	singleMsg.PushChannel = &channel

	return singleMsg
}

func PushSingle(authToken, reqid, cid string, cfg GeTuiConfig, singleMessage GeTuiSingleMessage) error {
	reqid = newRequestID(reqid, cfg)

	bodyByte, err := json.Marshal(singleMessage)
	if err != nil {
		return err
	}

	pushSingleUrl := fmt.Sprintf("%s%s/push/single/cid", baseURL, cfg.AppID)
	result, err := sendPost(pushSingleUrl, authToken, bodyByte)
	if GTV2_DEBUG {
		fmt.Println(reqid, "push single body:", pushSingleUrl, "\n", string(bodyByte))
		fmt.Println(reqid, "push single result:\n", result)
	}
	if err != nil {
		return err
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		return err
	}
	if resultMap["code"].(float64) != 0 {
		err = fmt.Errorf("GeTui.PushSingle error: %v", resultMap)
		return err
	}
	return nil
}
