package resources

type Element struct {
	Tag  string `json:"tag"`
	Text struct {
		Tag     string `json:"tag"`
		Content string `json:"content"`
	} `json:"text"`
}

type CardMessage struct {
	MsgType string `json:"msg_type"`
	Card    struct {
		Config struct {
			WideScreenMode bool `json:"wide_screen_mode"`
		} `json:"config"`
		Header struct {
			Title struct {
				Tag     string `json:"tag"`
				Content string `json:"content"`
			} `json:"title"`
			Template string `json:"template"`
		} `json:"header"`
		Elements []Element `json:"elements"`
	} `json:"card"`
}
