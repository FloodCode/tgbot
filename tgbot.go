package tgbot

import (
  "time"
)

var apiUrl = ""
var poolDelay = 300
var lastUpdateId int = -1
var updatesCallback = func([]Update) { }

func SetAPIKey(key string) {
  apiUrl = "https://api.telegram.org/bot" + key + "/"
}

func SetPoolDelay(delay int) {
  poolDelay = delay
}

func SetUpdatesCallback(callback func([]Update)) {
  updatesCallback = callback
}

func Poll() {
  for true {
    var updates, err = GetUpdates(ParamsGetUpdates{Offset: lastUpdateId + 1})
    if err == nil && len(updates) != 0 {
      var lastUpdate = updates[len(updates) - 1]
      updatesCallback(updates)
      lastUpdateId = lastUpdate.UpdateId
    }

    time.Sleep(time.Duration(poolDelay) * time.Millisecond)
  }
}

func GetUpdates(params ParamsGetUpdates) (updates []Update, err error) {
  err = sendResuest("getUpdates", params, &updates)
  return updates, err
}

func SendMessage(params ParamsSendMessage) (message Message, err error) {
  err = sendResuest("sendMessage", params, &message)
  return message, err
}
