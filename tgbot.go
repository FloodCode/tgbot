package tgbot

import (
	"time"
)

// TelegramBot allows to interact with Telegram Bot API.
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
func (b TelegramBot) GetMe() (user User, err error) {
	return user, b.sendResuest("getMe", nil, &user)
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

// SendSticker sends sticker.
func (b TelegramBot) SendSticker(params ParamsSendSticker) (message Message, err error) {
	return message, b.sendResuest("sendSticker", params, &message)
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

// RestrictChatMember allows to restrict a user in a supergroup.
func (b TelegramBot) RestrictChatMember(params ParamsRestrictChatMember) (success bool, err error) {
	return success, b.sendResuest("restrictChatMember", params, &success)
}

// PromoteChatMember allows to promote or demote a user in a supergroup or a channel.
func (b TelegramBot) PromoteChatMember(params ParamsPromoteChatMember) (success bool, err error) {
	return success, b.sendResuest("promoteChatMember", params, &success)
}

// ExportChatInviteLink allows to export an invite link to a supergroup or a channel.
func (b TelegramBot) ExportChatInviteLink(params ParamsExportChatInviteLink) (link string, err error) {
	return link, b.sendResuest("exportChatInviteLink", params, &link)
}

// SetChatPhoto allows to set a new profile photo for the chat.
func (b TelegramBot) SetChatPhoto(params ParamsSetChatPhoto) (success bool, err error) {
	return success, b.sendResuest("setChatPhoto", params, &success)
}

// DeleteChatPhoto allows to delete a new profile photo for the chat.
func (b TelegramBot) DeleteChatPhoto(params ParamsDeleteChatPhoto) (success bool, err error) {
	return success, b.sendResuest("deleteChatPhoto", params, &success)
}

// SetChatTitle allows to change the title of a chat.
func (b TelegramBot) SetChatTitle(params ParamsSetChatTitle) (success bool, err error) {
	return success, b.sendResuest("setChatTitle", params, &success)
}

// SetChatDescription allows to change the description of a chat.
func (b TelegramBot) SetChatDescription(params ParamsSetChatDescription) (success bool, err error) {
	return success, b.sendResuest("setChatDescription", params, &success)
}

// PinChatMessage allows to pin a message in a supergroup.
func (b TelegramBot) PinChatMessage(params ParamsPinChatMessage) (success bool, err error) {
	return success, b.sendResuest("pinChatMessage", params, &success)
}

// UnpinChatMessage allows to unpin a message in a supergroup.
func (b TelegramBot) UnpinChatMessage(params ParamsUnpinChatMessage) (success bool, err error) {
	return success, b.sendResuest("unpinChatMessage", params, &success)
}

// LeaveChat allows to leave a group, supergroup or channel.
func (b TelegramBot) LeaveChat(params ParamsLeaveChat) (success bool, err error) {
	return success, b.sendResuest("leaveChat", params, &success)
}

// GetChat allows to get up to date information about the chat.
func (b TelegramBot) GetChat(params ParamsGetChat) (chat Chat, err error) {
	return chat, b.sendResuest("getChat", params, &chat)
}

// GetChatAdministrators allows to get a list of administrators in a chat.
func (b TelegramBot) GetChatAdministrators(params ParamsGetChatAdministrators) (members []ChatMember, err error) {
	return members, b.sendResuest("getChatAdministrators", params, &members)
}

// GetChatMembersCount allows get the number of members in a chat.
func (b TelegramBot) GetChatMembersCount(params ParamsGetChatMembersCount) (count int, err error) {
	return count, b.sendResuest("getChatMembersCount", params, &count)
}

// GetChatMember allows to get information about a member of a chat.
func (b TelegramBot) GetChatMember(params ParamsGetChatMember) (member ChatMember, err error) {
	return member, b.sendResuest("getChatMember", params, &member)
}

// EditMessageText allows to edit text and game messages sent by the bot or via the bot (for inline bots).
func (b TelegramBot) EditMessageText(params ParamsEditMessageText) (message Message, err error) {
	return message, b.sendResuest("editMessageText", params, &message)
}

// EditMessageCaption allows to edit captions of messages sent by the bot or via the bot (for inline bots).
func (b TelegramBot) EditMessageCaption(params ParamsEditMessageCaption) (message Message, err error) {
	return message, b.sendResuest("editMessageCaption", params, &message)
}

// GetStickerSet allows to get a sticker set.
func (b TelegramBot) GetStickerSet(params ParamsGetStickerSet) (stickerSet StickerSet, err error) {
	return stickerSet, b.sendResuest("getStickerSet", params, &stickerSet)
}

// UploadStickerFile allows to upload a .png file with a sticker for
// later use in CreateNewStickerSet and AddStickerToSet methods.
func (b TelegramBot) UploadStickerFile(params ParamsUploadStickerFile) (file File, err error) {
	return file, b.sendResuest("uploadStickerFile", params, &file)
}

// CreateNewStickerSet allows to create new sticker set owned by a user.
func (b TelegramBot) CreateNewStickerSet(params ParamsCreateNewStickerSet) (success bool, err error) {
	return success, b.sendResuest("createNewStickerSet", params, &success)
}

// AddStickerToSet allows to add a new sticker to a set created by the bot.
func (b TelegramBot) AddStickerToSet(params ParamsAddStickerToSet) (success bool, err error) {
	return success, b.sendResuest("addStickerToSet", params, &success)
}

// SetStickerPositionInSet allows to move a sticker in a set created by the bot to a specific position.
func (b TelegramBot) SetStickerPositionInSet(params ParamsSetStickerPositionInSet) (success bool, err error) {
	return success, b.sendResuest("setStickerPositionInSet", params, &success)
}

// DeleteStickerFromSet allows to delete a sticker from a set created by the bot.
func (b TelegramBot) DeleteStickerFromSet(params ParamsDeleteStickerFromSet) (success bool, err error) {
	return success, b.sendResuest("deleteStickerFromSet", params, &success)
}

func (b TelegramBot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
	return sendResuest(method, b.apiKey, paramsObject, &t)
}
