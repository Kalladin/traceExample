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
func SendMessageToSlack(info WebhookInfo) error {
	if len(info.Contents) == 0 {
		return fmt.Errorf("Contents field is required. Slack Webhook Info: %#v", info)
	}
	fillDefaultValues(&info)
	slackWebhookURL := fmt.Sprintf("https://%s:%s%s", Host, Port, info.Path)

	webhookPayload, err := genWebhookPayload(&info)
	if err != nil {
		return fmt.Errorf("Generate Slack json payload failed. %s", err.Error())
	}

	req, err := http.NewRequest("POST", slackWebhookURL, strings.NewReader(webhookPayload))
	if err != nil {
		return fmt.Errorf("Create new request failed. %s", err.Error())
	}
	DefaultClient := &http.Client{
		Timeout: time.Duration(info.Timeout) * time.Second,
	}

	if len(info.ProxyURL) != 0 {
		proxy, err := url.Parse(info.ProxyURL)
		if err != nil {
			return fmt.Errorf("Error on parsing proxy url: %s. %s", info.ProxyURL, err.Error())
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		DefaultClient.Transport = transport
	}

	res, err := DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error getting response. %s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Fail to send notification to Slack by %s. %s, body: %+v", slackWebhookURL, res.Status, res.Body)
	}
	return nil
}

func fillDefaultValues(info *WebhookInfo) {
	if len(info.Channel) == 0 {
		info.Channel = DefaultChannel
	}
	if len(info.Username) == 0 {
		info.Username = DefaultUsername
	}
	if len(info.Pretext) == 0 {
		info.Pretext = DefaultPretext
	}
	if len(info.Color) == 0 {
		info.Color = DefaultColor
	}
	if info.Timeout == 0 {
		info.Timeout = DefaultTimeout
	}
	if info.Timestamp == 0 {
		info.Timestamp = time.Now().UTC().Unix()
	}
	if len(info.Path) == 0 {
		info.Path = DefaultWebhookPath
	}
}

func genWebhookPayload(info *WebhookInfo) (payloadStr string, err error) {

	var payload webhookPayload
	var paloadAttachments payloadAttachments
	timeFormat := time.Unix(int64(info.Timestamp), 0).String()

	payload.Channel = info.Channel
	payload.UserName = info.Username

	payload.Attachments = append(payload.Attachments, paloadAttachments)
	payload.Attachments[0].Pretext = info.Pretext
	payload.Attachments[0].Color = info.Color
	payload.Attachments[0].Timestamp = info.Timestamp

	payload.Attachments[0].Fields = append(payload.Attachments[0].Fields,
		payloadFields{
			Title: "Time",
			Value: timeFormat,
			Short: false,
		})
	keys := make([]string, len(info.Contents))
	for k := range info.Contents {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		payload.Attachments[0].Fields = append(payload.Attachments[0].Fields,
			payloadFields{
				Title: key,
				Value: info.Contents[key],
				Short: true,
			})
	}

	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Fail to marshal json payload. %s", err.Error())
	}
	return string(payloadByte), nil
}
