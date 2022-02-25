package getui

import (
	"encoding/json"
	"fmt"
	"time"
)

type GeTuiSingleMessage struct {
	RequestID   string           `json:"request_id"`
	Settings    GeTuiSettings    `json:"settings"`
	Audience    GeTuiAudience    `json:"audience"`
	PushMessage GeTuiPushMessage `json:"push_message"`
	PushChannel GeTuiPushChannel `json:"push_channel"`
}

func PushSingle(authToken, reqid, cid string, cfg GeTuiConfig, pushMsg GeTuiPushMessage, channel GeTuiPushChannel) error {
	if len(reqid) == 0 {
		reqid = fmt.Sprintf("gtreq-%s-%d", cfg.AppID, time.Now().Unix())
	}
	params := &GeTuiSingleMessage{RequestID: reqid}
	params.Settings = defaultSettings()
	params.Audience = singleAudience(cid)
	params.PushMessage = pushMsg
	params.PushChannel = channel

	bodyByte, err := json.Marshal(params)
	if err != nil {
		return err
	}

	pushSingleUrl := fmt.Sprintf("%s%s/push/single/cid", baseURL, cfg.AppID)
	result, err := sendPost(pushSingleUrl, authToken, bodyByte)
	if err != nil {
		return err
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		return err
	}
	if resultMap["code"].(float64) != 0 {
		err = fmt.Errorf("GeTui.pushSingle error: %v", resultMap)
		return err
	}
	return nil
}
