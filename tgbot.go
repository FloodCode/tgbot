package tgbot

import (
	"time"
)

// TelegramBot allows to interact with Telegram TelegramBot API.
type TelegramBot struct {
	apiKey          string
	poolDelay       int
	lastUpdateID    int
	updatesCallback func([]Update)
}

// New returns a new TelegramBot instance.
func New(apiKey string) TelegramBot {
	var bot = TelegramBot{
		apiKey:          apiKey,
		poolDelay:       100,
		lastUpdateID:    -1,
		updatesCallback: func([]Update) {},
	}
	return bot
}

// SetPollDelay used to specify updates polling delay.
func (b *TelegramBot) SetPollDelay(delay int) {
	b.poolDelay = delay
}

// SetUpdatesCallback used to specify callback for new updates.
func (b *TelegramBot) SetUpdatesCallback(callback func([]Update)) {
	b.updatesCallback = callback
}

// Poll starts updates polling.
func (b *TelegramBot) Poll() {
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

// GetMe returns basic information about the bot.
func (b TelegramBot) GetMe() (me User, err error) {
	err = b.sendResuest("getme", nil, &me)
	return me, err
}

// GetUpdates allows to get new updates.
func (b TelegramBot) GetUpdates(params ParamsGetUpdates) (updates []Update, err error) {
	err = b.sendResuest("getUpdates", params, &updates)
	return updates, err
}

// SendMessage sends text message.
func (b TelegramBot) SendMessage(params ParamsSendMessage) (message Message, err error) {
	err = b.sendResuest("sendMessage", params, &message)
	return message, err
}

// ForwardMessage re-sends message of any type.
func (b TelegramBot) ForwardMessage(params ParamsForwardMessage) (message Message, err error) {
	err = b.sendResuest("forwardMessage", params, &message)
	return message, err
}

// SendPhoto sends photo message.
func (b TelegramBot) SendPhoto(params ParamsSendPhoto) (message Message, err error) {
	err = b.sendResuest("sendPhoto", params, &message)
	return message, err
}

// SendAudio sends audio message.
func (b TelegramBot) SendAudio(params ParamsSendAudio) (message Message, err error) {
	err = b.sendResuest("sendAudio", params, &message)
	return message, err
}

// SendDocument sends document message.
func (b TelegramBot) SendDocument(params ParamsSendDocument) (message Message, err error) {
	err = b.sendResuest("sendDocument", params, &message)
	return message, err
}

// SendVideo sends video message.
func (b TelegramBot) SendVideo(params ParamsSendVideo) (message Message, err error) {
	err = b.sendResuest("sendVideo", params, &message)
	return message, err
}

// SendVoice sends voice note message.
func (b TelegramBot) SendVoice(params ParamsSendVoice) (message Message, err error) {
	err = b.sendResuest("sendVoice", params, &message)
	return message, err
}

// SendVideoNote sends video note message.
func (b TelegramBot) SendVideoNote(params ParamsSendVideoNote) (message Message, err error) {
	err = b.sendResuest("sendVideoNote", params, &message)
	return message, err
}

// SendLocation sends location message.
func (b TelegramBot) SendLocation(params ParamsSendLocation) (message Message, err error) {
	err = b.sendResuest("sendLocation", params, &message)
	return message, err
}

// SendVenue sends information about a venue.
func (b TelegramBot) SendVenue(params ParamsSendVenue) (message Message, err error) {
	err = b.sendResuest("sendVenue", params, &message)
	return message, err
}

// SendContact sends phone contact.
func (b TelegramBot) SendContact(params ParamsSendContact) (message Message, err error) {
	err = b.sendResuest("sendContact", params, &message)
	return message, err
}

// SendChatAction sends phone contact.
func (b TelegramBot) SendChatAction(params ParamsSendChatAction) (success bool, err error) {
	err = b.sendResuest("sendChatAction", params, &success)
	return success, err
}

func (b TelegramBot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
	return sendResuest(method, b.apiKey, paramsObject, &t)
}
