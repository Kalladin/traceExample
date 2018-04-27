package slack

const (
	host               = "hooks.slack.com"
	port               = "443"
	defaultWebhookPath = "/services/T0BL2BW3G/B3ZUSL677/X9yHXvWoEm3ocqKyqNEPgviH"
	defaultText        = "Droi Slack Webhook"
	defaultChannel     = "bot_home"
	defaultTimeout     = 15
)

// WebhookInfo - basic information for slack webhook
// Required fields: Payload and Payload.Channel
type WebhookInfo struct {
	ProxyURL string
	Timeout  int
	Payload  *Payload
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Fallback   string   `json:"fallback"`
	Color      string   `json:"color"`
	PreText    string   `json:"pretext"`
	AuthorName string   `json:"author_name"`
	AuthorLink string   `json:"author_link"`
	AuthorIcon string   `json:"author_icon"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	ImageUrl   string   `json:"image_url"`
	Fields     []*Field `json:"fields"`
	Footer     string   `json:"footer"`
	FooterIcon string   `json:"footer_icon"`
}

type Field struct {
	Title string      `json:"title"`
	Value interface{} `json:"value"`
	Short bool        `json:"short"`
}
