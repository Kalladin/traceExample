package slack

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	mockInfo     WebhookInfo
	mockChannel  = "ci"
	mockText     = "droipkg unit test"
	mockUsername = "merei"
	mockPayload  = Payload{
		Username: mockUsername,
		Text:     mockText,
	}
	mockAttachment Attachment
)

func BeforeTest() {
	mockInfo = WebhookInfo{
		Payload: &mockPayload,
	}
	httpmock.ActivateNonDefault(getProxyClient(&mockInfo))
}

func Test_getPayload(t *testing.T) {
	payload, err := getPayload(&mockInfo)
	assert.Nil(t, err)
	assert.Equal(t, "{\"username\":\"merei\",\"text\":\"droipkg unit test\"}", payload)
}

func Test_AddField(t *testing.T) {
	mockAttachment := Attachment{}
	fields := map[string]string{
		"b":   "",
		"123": "",
		"a":   "",
	}
	for k, v := range fields {
		field := Field{
			Title: k,
			Value: v,
		}
		mockAttachment.AddField(field)
	}
	assert.Equal(t, len(fields), len(mockAttachment.Fields))
	assert.Equal(t, "123", mockAttachment.Fields[0].Title)
	assert.Equal(t, "a", mockAttachment.Fields[1].Title)
	assert.Equal(t, "b", mockAttachment.Fields[2].Title)
}

func Test_fillDefaultValues(t *testing.T) {
	fillDefaultValues(&mockInfo)
	assert.Equal(t, defaultChannel, mockInfo.Payload.Channel)
	assert.Equal(t, "", mockInfo.ProxyURL)
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
