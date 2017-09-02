package tgbot

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// API method option ChatIdentifier

// ChatIdentifier represents unique identifier of chat.
type ChatIdentifier struct {
	chatID string
}

// Get returns string representation of chat ID.
func (c *ChatIdentifier) Get() string {
	return c.chatID
}

// ChatID creates new ChatIdentifier by chat ID.
func ChatID(id int64) *ChatIdentifier {
	return &ChatIdentifier{chatID: strconv.FormatInt(id, 10)}
}

// Username creates new ChatIdentifier by chat username.
func Username(username string) *ChatIdentifier {
	if strings.HasPrefix(username, "@") {
		return &ChatIdentifier{chatID: username}
	}

	return &ChatIdentifier{chatID: "@" + username}
}

// API method option ParseMode

// ParseMode represents message parse mode.
type ParseMode struct {
	parseMode string
}

// Get returns string representation of parse mode.
func (p *ParseMode) Get() string {
	return p.parseMode
}

// ParseModeMarkdown creates new ParseMode with Markdown option.
func ParseModeMarkdown() *ParseMode {
	return &ParseMode{parseMode: "Markdown"}
}

// ParseModeHTML creates new ParseMode with HTML option.
func ParseModeHTML() *ParseMode {
	return &ParseMode{parseMode: "HTML"}
}

// API method option ReplyMarkup
// TODO: Implement logic

// ReplyMarkup represents message reply markup
type ReplyMarkup struct {
	markup string
}

// API method option InputFile

// InputFile represents file to send.
type InputFile struct {
	fileData interface{}
	filename string
}

// Get returns file data.
func (f *InputFile) Get() interface{} {
	return f.fileData
}

// GetBytes returns file data byte array.
func (f *InputFile) GetBytes() []byte {
	if v, ok := f.fileData.([]byte); ok {
		return v
	}

	return []byte{}
}

// GetFilename returns filename.
func (f *InputFile) GetFilename() string {
	return f.filename
}

// FileID creates new InputFile by file id.
func FileID(fileID string) *InputFile {
	return &InputFile{fileData: fileID}
}

// FileBytes creates new InputFile by array of bytes and filename.
func FileBytes(fileData []byte, filename string) *InputFile {
	return &InputFile{fileData: fileData, filename: filename}
}

// FilePath creates new InputFile by file path.
func FilePath(filePath string) *InputFile {
	bytes, _ := ioutil.ReadFile(filePath)
	return &InputFile{fileData: bytes, filename: filepath.Base(filePath)}
}

// API method parameters

// ParamsGetUpdates represents parameters for GetUpdates method.
type ParamsGetUpdates struct {
	Offset  int `option:"offset"`
	Limit   int `option:"limit"`
	Timeout int `option:"timeout"`
}

// ParamsSendMessage represents parameters for SendMessage method.
type ParamsSendMessage struct {
	ChatID                *ChatIdentifier `option:"chat_id"                   required:"true"`
	Text                  string          `option:"text"                      required:"true"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	DisableNotification   bool            `option:"disable_notification"`
	ReplyToMessageID      int             `option:"reply_to_message_id"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

// ParamsForwardMessage represents parameters for ForwardMessage method.
type ParamsForwardMessage struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	FromChatID          *ChatIdentifier `option:"from_chat_id"              required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	MessageID           int             `option:"message_id"                required:"true"`
}

// ParamsSendPhoto represents parameters for SendPhoto method.
type ParamsSendPhoto struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Photo               *InputFile      `option:"photo"                     required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendAudio represents parameters for SendAudio method.
type ParamsSendAudio struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Audio               *InputFile      `option:"audio"                     required:"true"`
	Caption             string          `option:"caption"`
	Duration            int             `option:"duration"`
	Performer           string          `option:"performer"`
	Title               string          `option:"title"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendDocument represents parameters for SendDocument method.
type ParamsSendDocument struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Document            *InputFile      `option:"document"                  required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVideo represents parameters for SendVideo method.
type ParamsSendVideo struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Video               *InputFile      `option:"video"                     required:"true"`
	Duration            int             `option:"duration"`
	Width               int             `option:"width"`
	Height              int             `option:"height"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVoice represents parameters for SendVoice method.
type ParamsSendVoice struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Voice               *InputFile      `option:"voice"                     required:"true"`
	Duration            int             `option:"duration"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVideoNote represents parameters for SendVideoNote method.
type ParamsSendVideoNote struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	VideoNote           *InputFile      `option:"video_note"                required:"true"`
	Duration            int             `option:"duration"`
	Length              int             `option:"length"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendLocation represents parameters for SendLocation method.
type ParamsSendLocation struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Latitude            float64         `option:"latitude"                  required:"true"`
	Longitude           float64         `option:"longitude"                 required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}
