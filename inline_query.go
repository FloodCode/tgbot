package tgbot

// InlineQueryResult represents one result of an inline query
type InlineQueryResult interface {
	getType() string
}

// InlineQueryResultArticle represents a link to an article or web page
type InlineQueryResultArticle struct {
	ID                  string               `json:"id"`
	Title               string               `json:"title"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	URL                 string               `json:"url"`
	HideURL             bool                 `json:"hide_url"`
	Description         string               `json:"description"`
	ThumbURL            string               `json:"thumb_url"`
	ThumbWidth          string               `json:"thumb_width"`
	ThumbHeight         string               `json:"thumb_height"`
}

func (r InlineQueryResultArticle) getType() string {
	return "article"
}

// InlineQueryResultPhoto represents a link to a photo. By default, this photo will be sent by the user
// with optional caption. Alternatively, you can use InputMessageContent to send a message with the specified content
// instead of the photo
type InlineQueryResultPhoto struct {
	ID                  string               `json:"id"`
	PhotoURL            string               `json:"photo_url"`
	ThumbURL            string               `json:"thumb_url"`
	PhotoWidth          int                  `json:"photo_width"`
	PhotoHeight         int                  `json:"photo_height"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultPhoto) getType() string {
	return "photo"
}

// InputMessageContent represents the content of a message to be sent as a result of an inline query
type InputMessageContent interface {
	_dummyInputMessageContent()
}

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

func (c InputTextMessageContent) _dummyInputMessageContent() {}

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period"`
}

func (c InputLocationMessageContent) _dummyInputMessageContent() {}

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query
type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id"`
}

func (c InputVenueMessageContent) _dummyInputMessageContent() {}

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query
type InputContactMessageContent struct {
	InputMessageContent
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

func (c InputContactMessageContent) _dummyInputMessageContent() {}
