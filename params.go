package tgbot

import (
  "strconv"
  "strings"
)

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

type ReplyMarkup struct {
  markup string
}

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
