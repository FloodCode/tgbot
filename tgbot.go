package tgbot

import (
  "time"
)

type Bot struct {
  apiKey string
  poolDelay int
  lastUpdateId int
  updatesCallback func([]Update)
}

func New(apiKey string) Bot {
  var bot = Bot{
    apiKey: apiKey,
    poolDelay: 100,
    lastUpdateId: -1,
    updatesCallback: func([]Update) { },
  }
  return bot
}

func (b *Bot) SetPoolDelay(delay int) {
  b.poolDelay = delay
}

func (b *Bot) SetUpdatesCallback(callback func([]Update)) {
  b.updatesCallback = callback
}

func (b *Bot) Poll() {
  for true {
    var updates, err = b.GetUpdates(ParamsGetUpdates{Offset: b.lastUpdateId + 1})
    if err == nil && len(updates) != 0 {
      var lastUpdate = updates[len(updates) - 1]
      b.updatesCallback(updates)
      b.lastUpdateId = lastUpdate.UpdateId
    }

    time.Sleep(time.Duration(b.poolDelay) * time.Millisecond)
  }
}

func (b Bot) GetMe() (me User, err error) {
  err = b.sendResuest("getme", nil, &me)
  return me, err
}

func (b Bot) GetUpdates(params ParamsGetUpdates) (updates []Update, err error) {
  err = b.sendResuest("getUpdates", params, &updates)
  return updates, err
}

func (b Bot) SendMessage(params ParamsSendMessage) (message Message, err error) {
  err = b.sendResuest("sendMessage", params, &message)
  return message, err
}

func (b Bot) ForwardMessage(params ParamsForwardMessage) (message Message, err error) {
  err = b.sendResuest("forwardMessage", params, &message)
  return message, err
}

func (b Bot) SendPhoto(params ParamsSendPhoto) (message Message, err error) {
  err = b.sendResuest("sendPhoto", params, &message)
  return message, err
}

func (b Bot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
  return sendResuest(method, b.apiKey, paramsObject, &t)
}
