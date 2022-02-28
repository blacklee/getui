package getui

import (
	"fmt"
	"testing"
	"time"
)

func TestPushSingle(t *testing.T) {
	setup()
	cid := testCids[0]
	reqid := newRequestID("", testConfig)

	title := "Hello Go-GeTui"
	body := fmt.Sprintf("Hello, %s", time.Now().Format("2006-01-02 15:04:05"))

	singleMessage := CreateGeTuiSingleMessage(reqid, cid, testConfig, title, body, "none")
	err := PushSingle(testAuthToken, cid, testConfig, *singleMessage)
	fmt.Println(err)
}
