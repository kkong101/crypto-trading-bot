package webhook

type TradingViewWebhook struct {
	Pair    string `json:"pair"`
	Content string `json:"content"`
	// TODO: Add more fields
}
