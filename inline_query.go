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

// InlineQueryResultPhoto represents a link to a photo
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

// InlineQueryResultGif represents a link to an animated GIF file
type InlineQueryResultGif struct {
	ID                  string               `json:"id"`
	GifURL              string               `json:"gif_url"`
	GifWidth            int                  `json:"gif_width"`
	GifHeight           int                  `json:"gif_height"`
	GifDuration         int                  `json:"gif_duration"`
	ThumbURL            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultGif) getType() string {
	return "gif"
}

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound)
type InlineQueryResultMpeg4Gif struct {
	ID                  string               `json:"id"`
	Mpeg4URL            string               `json:"mpeg4_url"`
	Mpeg4Width          int                  `json:"mpeg4_width"`
	Mpeg4Height         int                  `json:"mpeg4_height"`
	Mpeg4Duration       int                  `json:"mpeg4_duration"`
	ThumbURL            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultMpeg4Gif) getType() string {
	return "mpeg4_gif"
}

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file
type InlineQueryResultVideo struct {
	ID                  string               `json:"id"`
	VideoURL            string               `json:"video_url"`
	MimeType            string               `json:"mime_type"`
	ThumbURL            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	VideoWidth          int                  `json:"video_width"`
	VideoHeight         int                  `json:"video_height"`
	VideoDuration       int                  `json:"video_duration"`
	Description         string               `json:"description"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultVideo) getType() string {
	return "video"
}

// InlineQueryResultAudio represents a link to a voice recording in an .ogg container encoded with OPUS
type InlineQueryResultAudio struct {
	ID                  string               `json:"id"`
	AudioURL            string               `json:"audio_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	Performer           string               `json:"performer"`
	AudioDuration       int                  `json:"audio_duration"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultAudio) getType() string {
	return "audio"
}

// InlineQueryResultVoice represents a link to a voice recording in an .ogg container encoded with OPUS
type InlineQueryResultVoice struct {
	ID                  string               `json:"id"`
	VoiceURL            string               `json:"voice_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	VoiceDuration       int                  `json:"voice_duration"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}

func (r InlineQueryResultVoice) getType() string {
	return "voice"
}

// InlineQueryResultDocument represents a link to a file
type InlineQueryResultDocument struct {
	ID                  string               `json:"id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	DocumentURL         string               `json:"document_url"`
	MimeType            string               `json:"mime_type"`
	Description         string               `json:"description"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	ThumbURL            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (r InlineQueryResultDocument) getType() string {
	return "document"
}

// InlineQueryResultLocation represents a location on a map
type InlineQueryResultLocation struct {
	ID                  string               `json:"id"`
	Latitude            float64              `json:"latitude"`
	Longitude           float64              `json:"longitude"`
	Title               string               `json:"title"`
	LivePeriod          int                  `json:"live_period"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	ThumbURL            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (r InlineQueryResultLocation) getType() string {
	return "location"
}

// InlineQueryResultVenue represents a venue
type InlineQueryResultVenue struct {
	ID                  string               `json:"id"`
	Latitude            float64              `json:"latitude"`
	Longitude           float64              `json:"longitude"`
	Title               string               `json:"title"`
	Address             string               `json:"address"`
	FoursquareID        string               `json:"foursquare_id"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	ThumbURL            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (r InlineQueryResultVenue) getType() string {
	return "venue"
}

// InlineQueryResultContact represents a contact with a phone number
type InlineQueryResultContact struct {
	ID                  string               `json:"id"`
	PhoneNumber         string               `json:"phone_number"`
	FirstName           string               `json:"first_name"`
	LastName            string               `json:"last_name"`
	ReplyMarkup         *ReplyMarkup         `json:"reply_markup"`
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	ThumbURL            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (r InlineQueryResultContact) getType() string {
	return "contact"
}

// InlineQueryResultGame represents a Game
type InlineQueryResultGame struct {
	ID            string `json:"id"`
	GameShortName string `json:"game_short_name"`
}

func (r InlineQueryResultGame) getType() string {
	return "game"
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
