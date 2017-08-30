package tgbot

import (
  "strconv"
  "strings"
  "io/ioutil"
)

// API method option ChatIdentifier

type ChatIdentifier struct {
  chatId string
}

func (c *ChatIdentifier) Get() string {
  return c.chatId
}

func ChatId(id int64) *ChatIdentifier {
  return &ChatIdentifier{chatId: strconv.FormatInt(id, 10)}
}

func Username(username string) *ChatIdentifier {
  if strings.HasPrefix(username, "@") {
    return &ChatIdentifier{chatId: username}
  } else {
    return &ChatIdentifier{chatId: "@" + username}
  }
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
}

func (f *InputFile) Get() interface{} {
  return f.fileData
}

func FileId(fileId string) *InputFile {
  return &InputFile{fileData: fileId}
}

func FileBytes(fileData []byte) *InputFile {
  return &InputFile{fileData: fileData}
}

func FilePath(filePath string) *InputFile {
  bytes, _ := ioutil.ReadFile(filePath)
  return &InputFile{fileData: bytes}
}

// API method parameters

type ParamsGetUpdates struct {
  Offset  int `option:"offset"`
  Limit   int `option:"limit"`
  Timeout int `option:"timeout"`
}

type ParamsSendMessage struct {
  ChatId                *ChatIdentifier `option:"chat_id"                   required:"true"`
  Text                  string          `option:"text"                      required:"true"`
  ParseMode             *ParseMode      `option:"parse_mode"`
  DisableWebPagePreview bool            `option:"disable_web_page_preview"`
  DisableNotification   bool            `option:"disable_notification"`
  ReplyToMessageId      int             `option:"reply_to_message_id"`
  ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

type ParamsForwardMessage struct {
  ChatId                *ChatIdentifier `option:"chat_id"                   required:"true"`
  FromChatId            *ChatIdentifier `option:"from_chat_id"              required:"true"`
  DisableNotification   bool            `option:"disable_notification"`
  MessageId             int             `option:"message_id"                required:"true"`
}

type ParamsSendPhoto struct {
  ChatId                *ChatIdentifier `option:"chat_id" required:"true"`
  Photo                 *InputFile      `option:"photo"   required:"true"`
  Caption               string          `option:"caption"`
  DisableNotification   bool            `option:"disable_notification"`
  ReplyToMessageId      int             `option:"reply_to_message_id"`
  ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}
