package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// SendMessageToSlack - send message to slack webhook and show formatted information on channel
func SendMessageToSlack(info WebhookInfo) (err error) {
	if info.Payload == nil {
		return fmt.Errorf("Payload is required field on WebhookInfo")
	}
	fillDefaultValues(&info)
	webhookPayload, err := getPayload(&info)
	if err != nil {
		return fmt.Errorf("Get Slack json payload failed: %+v", info)
	}
	slackWebhookURL := GetSlackWebhookURL()

	req, err := http.NewRequest("POST", slackWebhookURL, strings.NewReader(webhookPayload))
	if err != nil {
		return fmt.Errorf("Create new request failed. %s", err.Error())
	}
	proxyClient := getProxyClient(&info)
	res, err := proxyClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error getting response. %s", err.Error())
	}
	defer res.Body.Close()
	// Slack respond "HTTP 200 OK" for successful request
	// See "Handling errors" on https://api.slack.com/incoming-webhooks
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Fail to send notification to Slack by %s. %s, body: %+v", slackWebhookURL, res.Status, res.Body)
	}
	return nil
}

func getProxyClient(info *WebhookInfo) *http.Client {
	proxyClient := &http.Client{
		Timeout: time.Duration(info.Timeout) * time.Second,
	}

	if len(info.ProxyURL) != 0 {
		proxy, err := url.Parse(info.ProxyURL)
		if err != nil {
			return nil
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		proxyClient.Transport = transport
	}
	return proxyClient
}

func fillDefaultValues(info *WebhookInfo) {
	payload := info.Payload
	if len(payload.Channel) == 0 {
		payload.Channel = defaultChannel
	}
	if len(payload.Text) == 0 {
		payload.Text = defaultText
	}
	if info.Timeout == 0 {
		info.Timeout = defaultTimeout
	}
}

func getPayload(info *WebhookInfo) (payloadStr string, err error) {
	payload := info.Payload
	timeFormat := time.Now().UTC().Format("2006-01-02 15:04:05Z")

	for i := 0; i < len(payload.Attachments); i++ {
		payload.Attachments[i].Fields = append([]*Field{{Title: "Time", Value: timeFormat, Short: false}}, payload.Attachments[i].Fields...)
	}

	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Fail to marshal json payload. %s", err.Error())
	}
	return string(payloadByte), nil
}

// AddField - Add field and sort by field titles ASC
func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	sort.Slice(attachment.Fields, func(i, j int) bool {
		return attachment.Fields[i].Title < attachment.Fields[j].Title
	})
	return attachment
}

// GetSlackWebhookURL - Get the webhook URL with default path
// Can be used for httpmock of unit test
func GetSlackWebhookURL() string {
	return fmt.Sprintf("https://%s:%s%s", host, port, defaultWebhookPath)
}
