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
			b.lastUpdateID = updates[len(updates)-1].UpdateID
			go b.updatesCallback(updates)
		}

		time.Sleep(time.Duration(b.poolDelay) * time.Millisecond)
	}
}

// GetMe returns basic information about the bot.
func (b TelegramBot) GetMe() (me User, err error) {
	return me, b.sendResuest("getme", nil, &me)
}

// GetUpdates allows to get new updates.
func (b TelegramBot) GetUpdates(params ParamsGetUpdates) (updates []Update, err error) {
	return updates, b.sendResuest("getUpdates", params, &updates)
}

// SendMessage sends text message.
func (b TelegramBot) SendMessage(params ParamsSendMessage) (message Message, err error) {
	return message, b.sendResuest("sendMessage", params, &message)
}

// ForwardMessage re-sends message of any type.
func (b TelegramBot) ForwardMessage(params ParamsForwardMessage) (message Message, err error) {
	return message, b.sendResuest("forwardMessage", params, &message)
}

// SendPhoto sends photo message.
func (b TelegramBot) SendPhoto(params ParamsSendPhoto) (message Message, err error) {
	return message, b.sendResuest("sendPhoto", params, &message)
}

// SendAudio sends audio message.
func (b TelegramBot) SendAudio(params ParamsSendAudio) (message Message, err error) {
	return message, b.sendResuest("sendAudio", params, &message)
}

// SendDocument sends document message.
func (b TelegramBot) SendDocument(params ParamsSendDocument) (message Message, err error) {
	return message, b.sendResuest("sendDocument", params, &message)
}

// SendVideo sends video message.
func (b TelegramBot) SendVideo(params ParamsSendVideo) (message Message, err error) {
	return message, b.sendResuest("sendVideo", params, &message)
}

// SendVoice sends voice note message.
func (b TelegramBot) SendVoice(params ParamsSendVoice) (message Message, err error) {
	return message, b.sendResuest("sendVoice", params, &message)
}

// SendVideoNote sends video note message.
func (b TelegramBot) SendVideoNote(params ParamsSendVideoNote) (message Message, err error) {
	return message, b.sendResuest("sendVideoNote", params, &message)
}

// SendLocation sends location message.
func (b TelegramBot) SendLocation(params ParamsSendLocation) (message Message, err error) {
	return message, b.sendResuest("sendLocation", params, &message)
}

// SendVenue sends information about a venue.
func (b TelegramBot) SendVenue(params ParamsSendVenue) (message Message, err error) {
	return message, b.sendResuest("sendVenue", params, &message)
}

// SendContact sends phone contact.
func (b TelegramBot) SendContact(params ParamsSendContact) (message Message, err error) {
	return message, b.sendResuest("sendContact", params, &message)
}

// SendChatAction sends phone contact.
func (b TelegramBot) SendChatAction(params ParamsSendChatAction) (success bool, err error) {
	return success, b.sendResuest("sendChatAction", params, &success)
}

// GetUserProfilePhotos returns user profile photos.
func (b TelegramBot) GetUserProfilePhotos(params ParamsGetUserProfilePhotos) (photos UserProfilePhotos, err error) {
	return photos, b.sendResuest("getUserProfilePhotos", params, &photos)
}

// GetFile allows to get basic info about a file and prepare it for downloading.
func (b TelegramBot) GetFile(params ParamsGetFile) (file File, err error) {
	return file, b.sendResuest("getFile", params, &file)
}

// KickChatMember allows to kick user from a group, a supergroup or a channel.
func (b TelegramBot) KickChatMember(params ParamsKickChatMember) (success bool, err error) {
	return success, b.sendResuest("kickChatMember", params, &success)
}

// UnbanChatMember allows to unban user from a group, a supergroup or a channel.
func (b TelegramBot) UnbanChatMember(params ParamsUnbanChatMember) (success bool, err error) {
	return success, b.sendResuest("unbanChatMember", params, &success)
}

func (b TelegramBot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
	return sendResuest(method, b.apiKey, paramsObject, &t)
}
