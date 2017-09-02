package tgbot

import (
	"time"
)

type Bot struct {
	apiKey          string
	poolDelay       int
	lastUpdateID    int
	updatesCallback func([]Update)
}

func New(apiKey string) Bot {
	var bot = Bot{
		apiKey:          apiKey,
		poolDelay:       100,
		lastUpdateID:    -1,
		updatesCallback: func([]Update) {},
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
		var updates, err = b.GetUpdates(ParamsGetUpdates{Offset: b.lastUpdateID + 1})
		if err == nil && len(updates) != 0 {
			var lastUpdate = updates[len(updates)-1]
			b.updatesCallback(updates)
			b.lastUpdateID = lastUpdate.UpdateID
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

func (b Bot) SendAudio(params ParamsSendAudio) (message Message, err error) {
	err = b.sendResuest("sendAudio", params, &message)
	return message, err
}

func (b Bot) SendDocument(params ParamsSendDocument) (message Message, err error) {
	err = b.sendResuest("sendDocument", params, &message)
	return message, err
}

func (b Bot) SendVideo(params ParamsSendVideo) (message Message, err error) {
	err = b.sendResuest("sendVideo", params, &message)
	return message, err
}

func (b Bot) SendVoice(params ParamsSendVoice) (message Message, err error) {
	err = b.sendResuest("sendVoice", params, &message)
	return message, err
}

func (b Bot) SendVideoNote(params ParamsSendVideoNote) (message Message, err error) {
	err = b.sendResuest("sendVideoNote", params, &message)
	return message, err
}

func (b Bot) SendLocation(params ParamsSendLocation) (message Message, err error) {
	err = b.sendResuest("sendLocation", params, &message)
	return message, err
}

func (b Bot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
	return sendResuest(method, b.apiKey, paramsObject, &t)
}
