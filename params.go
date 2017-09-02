package tgbot

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// API method option ChatIdentifier

type ChatIdentifier struct {
	chatID string
}

func (c *ChatIdentifier) Get() string {
	return c.chatID
}

func ChatID(id int64) *ChatIdentifier {
	return &ChatIdentifier{chatID: strconv.FormatInt(id, 10)}
}

func Username(username string) *ChatIdentifier {
	if strings.HasPrefix(username, "@") {
		return &ChatIdentifier{chatID: username}
	}

	return &ChatIdentifier{chatID: "@" + username}
}

// API method option ParseMode

type ParseMode struct {
	parseMode string
}

func (p *ParseMode) Get() string {
	return p.parseMode
}

func ParseModeMarkdown() *ParseMode {
	return &ParseMode{parseMode: "Markdown"}
}

func ParseModeHTML() *ParseMode {
	return &ParseMode{parseMode: "HTML"}
}

// API method option ReplyMarkup
// TODO: Implement logic

type ReplyMarkup struct {
	markup string
}

// API method option InputFile

type InputFile struct {
	fileData interface{}
	filename string
}

func (f *InputFile) Get() interface{} {
	return f.fileData
}

func (f *InputFile) GetBytes() []byte {
	if v, ok := f.fileData.([]byte); ok {
		return v
	}

	return []byte{}
}

func (f *InputFile) GetFilename() string {
	return f.filename
}

func FileID(fileID string) *InputFile {
	return &InputFile{fileData: fileID}
}

func FileBytes(fileData []byte, filename string) *InputFile {
	return &InputFile{fileData: fileData, filename: filename}
}

func FilePath(filePath string) *InputFile {
	bytes, _ := ioutil.ReadFile(filePath)
	return &InputFile{fileData: bytes, filename: filepath.Base(filePath)}
}

// API method parameters

type ParamsGetUpdates struct {
	Offset  int `option:"offset"`
	Limit   int `option:"limit"`
	Timeout int `option:"timeout"`
}

type ParamsSendMessage struct {
	ChatID                *ChatIdentifier `option:"chat_id"                   required:"true"`
	Text                  string          `option:"text"                      required:"true"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	DisableNotification   bool            `option:"disable_notification"`
	ReplyToMessageID      int             `option:"reply_to_message_id"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

type ParamsForwardMessage struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	FromChatID          *ChatIdentifier `option:"from_chat_id"              required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	MessageID           int             `option:"message_id"                required:"true"`
}

type ParamsSendPhoto struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Photo               *InputFile      `option:"photo"                     required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

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

type ParamsSendDocument struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Document            *InputFile      `option:"document"                  required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

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

type ParamsSendVoice struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Voice               *InputFile      `option:"voice"                     required:"true"`
	Duration            int             `option:"duration"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

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

type ParamsSendLocation struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Latitude            float64         `option:"latitude"                  required:"true"`
	Longitude           float64         `option:"longitude"                 required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}
