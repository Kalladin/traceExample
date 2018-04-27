package slack

const (
	Host               = "hooks.slack.com"
	Port               = "443"
	DefaultWebhookPath = "/services/T0BL2BW3G/B3ZUSL677/X9yHXvWoEm3ocqKyqNEPgviH"
	DefaultUsername    = "Satomi"
	DefaultChannel     = "random"
	DefaultPretext     = "Droi Slack Webhook"
	DefaultColor       = "warning"
	DefaultTimeout     = 15
)

// WebhookInfo - basic information for slack webhook
// Required fields: Channel and Contents
type WebhookInfo struct {
	Path      string
	Username  string
	Channel   string
	Pretext   string
	Color     string // See https://api.slack.com/docs/message-attachments
	ProxyURL  string
	Timestamp int64
	Timeout   int
	Contents  map[string]interface{}
}

type webhookPayload struct {
	Channel     string               `json:"channel"`
	UserName    string               `json:"username"`
	Attachments []payloadAttachments `json:"attachments"`
}

type payloadAttachments struct {
	Pretext   string           `json:"pretext"`
	Color     string           `json:"color"`
	Fields    []payloadFields  `json:"fields"`
	Actions   []payloadActions `json:"actions"`
	Timestamp int64            `json:"ts"`
}

type payloadFields struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
	Short bool        `json:"short"`
}

type payloadActions struct {
	Name    string                `json:"name"`
	Text    string                `json:"text"`
	Type    string                `json:"type"`
	Value   string                `json:"value"`
	Confirm payloadActionsConfirm `json:"confirm"`
}

type payloadActionsConfirm struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	OkText      string `json:"ok_text"`
	DismissText string `json:"dismiss_text"`
}
