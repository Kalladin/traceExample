package slack

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	mockInfo    WebhookInfo
	mockChannel = "ci"
	mockPretext = "Pretext for droipkg unit test"
	mockContent map[string]interface{}
)

func BeforeTest() {
	httpmock.Activate()
	mockContent = make(map[string]interface{})
	mockContent["f1"] = 123
	mockContent["f2"] = "foo"

	mockInfo = WebhookInfo{
		Pretext:  mockPretext,
		Channel:  mockChannel,
		Contents: mockContent,
	}
}

func Test_SendMessageToSlack(t *testing.T) {
	fillDefaultValues(&mockInfo)
	url := fmt.Sprintf("https://%s:%s%s", Host, Port, mockInfo.Path)
	httpmock.RegisterResponder("POST", url,
		httpmock.NewStringResponder(201, "POST message to slack webhook"))
	err := SendMessageToSlack(mockInfo)
	assert.NoError(t, err)
}

func Test_fillDefaultValues(t *testing.T) {
	fillDefaultValues(&mockInfo)
	fmt.Printf("mockInfo: %+v\n", mockInfo)

	assert.Equal(t, DefaultUsername, mockInfo.Username)
	assert.Equal(t, mockChannel, mockInfo.Channel)
	assert.Equal(t, "", mockInfo.ProxyURL)
}

func Test_genWebhookPayload(t *testing.T) {
	var payloadObj webhookPayload
	fillDefaultValues(&mockInfo)

	payload, err := genWebhookPayload(&mockInfo)
	assert.NoError(t, err)

	err = json.Unmarshal([]byte(payload), &payloadObj)
	assert.NoError(t, err)
	assert.Equal(t, DefaultUsername, payloadObj.UserName)
	assert.Equal(t, mockChannel, payloadObj.Channel)
	attachments := payloadObj.Attachments
	for _, att := range attachments {
		assert.Equal(t, DefaultColor, att.Color)
		assert.Equal(t, mockPretext, att.Pretext)
		assert.NotEqual(t, 0, att.Timestamp)
	}
}

// Do somethings after all test cases
func AfterTest() {
	httpmock.DeactivateAndReset()
}

func TestMain(m *testing.M) {
	BeforeTest()
	retCode := m.Run()
	AfterTest()
	os.Exit(retCode)
}
