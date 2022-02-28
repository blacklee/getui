package getui

import (
	"fmt"
	"testing"
)

func TestPushSingle(t *testing.T) {
	cfg := loadAppConfig()
	hash := loadTestInfo()
	token := hash["token"].(string)
	cids := hash["cid"].([]interface{})
	cid := cids[0].(string)
	reqid := newRequestID("", cfg)

	title := "hello title---"
	body := "hello msg body..."

	singleMessage := CreateGeTuiSingleMessage(reqid, cid, cfg, title, body, "none")
	err := PushSingle(token, reqid, cid, cfg, *singleMessage)
	fmt.Println(err)
}
