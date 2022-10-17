package domain

type Notification struct {
	AndroidChannelId string `json:"android_channel_id,omitempty"`
	BodyLocKey       string `json:"body_loc_key,omitempty"`
	BodyLocArgs      string `json:"body_loc_args,omitempty"`
	Body             string `json:"body,omitempty"`
	Badge            string `json:"badge,omitempty"`
	Color            string `json:"color,omitempty"`
	ClickAction      string `json:"click_action,omitempty"`
	Icon             string `json:"icon,omitempty"`
	Subtitle         string `json:"subtitle,omitempty"`
	Sound            string `json:"sound,omitempty"`
	Tag              string `json:"tag,omitempty"`
	Title            string `json:"title,omitempty"`
	TitleLocKey      string `json:"title_loc_key,omitempty"`
	TitleLocArgs     string `json:"title_loc_args,omitempty"`
}
