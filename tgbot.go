package tgbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TelegramBot allows to interact with Telegram Bot API
type TelegramBot struct {
	apiKey       string
	lastUpdateID int
}

// New returns a new TelegramBot instance
func New(apiKey string) (TelegramBot, error) {
	var err error
	var bot = TelegramBot{
		apiKey:       apiKey,
		lastUpdateID: -1,
	}

	_, err = bot.GetMe()
	if err != nil {
		return TelegramBot{}, err
	}

	return bot, nil
}

// PollConfig represents bot's polling configuration
type PollConfig struct {
	Callback func([]Update)
	Delay    int
}

// Poll starts updates polling
func (b *TelegramBot) Poll(config PollConfig) error {
	b.DeleteWebhook()

	for {
		var updates, err = b.GetUpdates(GetUpdatesConfig{Offset: b.lastUpdateID + 1})
		if err == nil && len(updates) != 0 {
			b.lastUpdateID = updates[len(updates)-1].UpdateID
			go config.Callback(updates)
		}

		time.Sleep(time.Duration(config.Delay) * time.Millisecond)
	}
}

// ListenConfig represents bot's webhook configuration
type ListenConfig struct {
	Callback       func([]Update)
	Host           string
	Port           uint16
	KeyFilename    string
	CertFilename   string
	MaxConnections int
	AllowedUpdates []string
}

// Listen starts HTTPS server to receive updates
func (b *TelegramBot) Listen(config ListenConfig) error {
	http.HandleFunc("/"+b.apiKey, func(w http.ResponseWriter, req *http.Request) {
		var update Update
		err := json.NewDecoder(req.Body).Decode(&update)
		if err != nil {
			return
		}

		defer req.Body.Close()
		go config.Callback([]Update{update})
	})

	_, err := b.SetWebhook(SetWebhookConfig{
		URL:            fmt.Sprintf("https://%s:%d/%s", config.Host, config.Port, b.apiKey),
		Certificate:    FilePath(config.CertFilename),
		AllowedUpdates: config.AllowedUpdates,
	})

	if err != nil {
		return err
	}

	return http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Port), config.CertFilename, config.KeyFilename, nil)
}

// GetMe returns basic information about the bot
func (b TelegramBot) GetMe() (user User, err error) {
	return user, b.sendResuest("getMe", nil, &user)
}

// GetUpdates allows to get new updates
func (b TelegramBot) GetUpdates(config GetUpdatesConfig) (updates []Update, err error) {
	return updates, b.sendResuest("getUpdates", config, &updates)
}

// SetWebhook used to specify url and receive incoming updates via an outgoing webhook
func (b TelegramBot) SetWebhook(config SetWebhookConfig) (success bool, err error) {
	return success, b.sendResuest("setWebhook", config, &success)
}

// DeleteWebhook used to remove webhook integration
func (b TelegramBot) DeleteWebhook() (success bool, err error) {
	return success, b.sendResuest("deleteWebhook", nil, &success)
}

// GetWebhookInfo user to get current webhook status
func (b TelegramBot) GetWebhookInfo() (info WebhookInfo, err error) {
	return info, b.sendResuest("getWebhookInfo", nil, &info)
}

// SendMessage sends text message
func (b TelegramBot) SendMessage(config SendMessageConfig) (message Message, err error) {
	return message, b.sendResuest("sendMessage", config, &message)
}

// ForwardMessage re-sends message of any type
func (b TelegramBot) ForwardMessage(config ForwardMessageConfig) (message Message, err error) {
	return message, b.sendResuest("forwardMessage", config, &message)
}

// SendPhoto sends photo message
func (b TelegramBot) SendPhoto(config SendPhotoConfig) (message Message, err error) {
	return message, b.sendResuest("sendPhoto", config, &message)
}

// SendAudio sends audio message
func (b TelegramBot) SendAudio(config SendAudioConfig) (message Message, err error) {
	return message, b.sendResuest("sendAudio", config, &message)
}

// SendDocument sends document message
func (b TelegramBot) SendDocument(config SendDocumentConfig) (message Message, err error) {
	return message, b.sendResuest("sendDocument", config, &message)
}

// SendVideo sends video message
func (b TelegramBot) SendVideo(config SendVideoConfig) (message Message, err error) {
	return message, b.sendResuest("sendVideo", config, &message)
}

// SendVoice sends voice note message
func (b TelegramBot) SendVoice(config SendVoiceConfig) (message Message, err error) {
	return message, b.sendResuest("sendVoice", config, &message)
}

// SendVideoNote sends video note message
func (b TelegramBot) SendVideoNote(config SendVideoNoteConfig) (message Message, err error) {
	return message, b.sendResuest("sendVideoNote", config, &message)
}

// SendLocation sends location message
func (b TelegramBot) SendLocation(config SendLocationConfig) (message Message, err error) {
	return message, b.sendResuest("sendLocation", config, &message)
}

// SendVenue sends information about a venue
func (b TelegramBot) SendVenue(config SendVenueConfig) (message Message, err error) {
	return message, b.sendResuest("sendVenue", config, &message)
}

// SendContact sends phone contact
func (b TelegramBot) SendContact(config SendContactConfig) (message Message, err error) {
	return message, b.sendResuest("sendContact", config, &message)
}

// SendSticker sends sticker
func (b TelegramBot) SendSticker(config SendStickerConfig) (message Message, err error) {
	return message, b.sendResuest("sendSticker", config, &message)
}

// SendChatAction sends phone contact
func (b TelegramBot) SendChatAction(config SendChatActionConfig) (success bool, err error) {
	return success, b.sendResuest("sendChatAction", config, &success)
}

// GetUserProfilePhotos returns user profile photos
func (b TelegramBot) GetUserProfilePhotos(config GetUserProfilePhotosConfig) (photos UserProfilePhotos, err error) {
	return photos, b.sendResuest("getUserProfilePhotos", config, &photos)
}

// GetFile allows to get basic info about a file and prepare it for downloading
func (b TelegramBot) GetFile(config GetFileConfig) (file File, err error) {
	return file, b.sendResuest("getFile", config, &file)
}

// KickChatMember allows to kick user from a group, a supergroup or a channel
func (b TelegramBot) KickChatMember(config KickChatMemberConfig) (success bool, err error) {
	return success, b.sendResuest("kickChatMember", config, &success)
}

// UnbanChatMember allows to unban user from a group, a supergroup or a channel
func (b TelegramBot) UnbanChatMember(config UnbanChatMemberConfig) (success bool, err error) {
	return success, b.sendResuest("unbanChatMember", config, &success)
}

// RestrictChatMember allows to restrict a user in a supergroup
func (b TelegramBot) RestrictChatMember(config RestrictChatMemberConfig) (success bool, err error) {
	return success, b.sendResuest("restrictChatMember", config, &success)
}

// PromoteChatMember allows to promote or demote a user in a supergroup or a channel
func (b TelegramBot) PromoteChatMember(config PromoteChatMemberConfig) (success bool, err error) {
	return success, b.sendResuest("promoteChatMember", config, &success)
}

// ExportChatInviteLink allows to export an invite link to a supergroup or a channel
func (b TelegramBot) ExportChatInviteLink(config ExportChatInviteLinkConfig) (link string, err error) {
	return link, b.sendResuest("exportChatInviteLink", config, &link)
}

// SetChatPhoto allows to set a new profile photo for the chat
func (b TelegramBot) SetChatPhoto(config SetChatPhotoConfig) (success bool, err error) {
	return success, b.sendResuest("setChatPhoto", config, &success)
}

// DeleteChatPhoto allows to delete a new profile photo for the chat
func (b TelegramBot) DeleteChatPhoto(config DeleteChatPhotoConfig) (success bool, err error) {
	return success, b.sendResuest("deleteChatPhoto", config, &success)
}

// SetChatTitle allows to change the title of a chat
func (b TelegramBot) SetChatTitle(config SetChatTitleConfig) (success bool, err error) {
	return success, b.sendResuest("setChatTitle", config, &success)
}

// SetChatDescription allows to change the description of a chat
func (b TelegramBot) SetChatDescription(config SetChatDescriptionConfig) (success bool, err error) {
	return success, b.sendResuest("setChatDescription", config, &success)
}

// PinChatMessage allows to pin a message in a supergroup
func (b TelegramBot) PinChatMessage(config PinChatMessageConfig) (success bool, err error) {
	return success, b.sendResuest("pinChatMessage", config, &success)
}

// UnpinChatMessage allows to unpin a message in a supergroup
func (b TelegramBot) UnpinChatMessage(config UnpinChatMessageConfig) (success bool, err error) {
	return success, b.sendResuest("unpinChatMessage", config, &success)
}

// LeaveChat allows to leave a group, supergroup or channel
func (b TelegramBot) LeaveChat(config LeaveChatConfig) (success bool, err error) {
	return success, b.sendResuest("leaveChat", config, &success)
}

// GetChat allows to get up to date information about the chat
func (b TelegramBot) GetChat(config GetChatConfig) (chat Chat, err error) {
	return chat, b.sendResuest("getChat", config, &chat)
}

// GetChatAdministrators allows to get a list of administrators in a chat
func (b TelegramBot) GetChatAdministrators(config GetChatAdministratorsConfig) (members []ChatMember, err error) {
	return members, b.sendResuest("getChatAdministrators", config, &members)
}

// GetChatMembersCount allows get the number of members in a chat
func (b TelegramBot) GetChatMembersCount(config GetChatMembersCountConfig) (count int, err error) {
	return count, b.sendResuest("getChatMembersCount", config, &count)
}

// GetChatMember allows to get information about a member of a chat
func (b TelegramBot) GetChatMember(config GetChatMemberConfig) (member ChatMember, err error) {
	return member, b.sendResuest("getChatMember", config, &member)
}

// EditMessageText allows to edit text and game messages sent by the bot or via the bot (for inline bots)
func (b TelegramBot) EditMessageText(config EditMessageTextConfig) (message Message, err error) {
	return message, b.sendResuest("editMessageText", config, &message)
}

// EditMessageCaption allows to edit captions of messages sent by the bot or via the bot (for inline bots)
func (b TelegramBot) EditMessageCaption(config EditMessageCaptionConfig) (message Message, err error) {
	return message, b.sendResuest("editMessageCaption", config, &message)
}

// GetStickerSet allows to get a sticker set
func (b TelegramBot) GetStickerSet(config GetStickerSetConfig) (stickerSet StickerSet, err error) {
	return stickerSet, b.sendResuest("getStickerSet", config, &stickerSet)
}

// UploadStickerFile allows to upload a .png file with a sticker for
// later use in CreateNewStickerSet and AddStickerToSet methods
func (b TelegramBot) UploadStickerFile(config UploadStickerFileConfig) (file File, err error) {
	return file, b.sendResuest("uploadStickerFile", config, &file)
}

// CreateNewStickerSet allows to create new sticker set owned by a user
func (b TelegramBot) CreateNewStickerSet(config CreateNewStickerSetConfig) (success bool, err error) {
	return success, b.sendResuest("createNewStickerSet", config, &success)
}

// AddStickerToSet allows to add a new sticker to a set created by the bot
func (b TelegramBot) AddStickerToSet(config AddStickerToSetConfig) (success bool, err error) {
	return success, b.sendResuest("addStickerToSet", config, &success)
}

// SetStickerPositionInSet allows to move a sticker in a set created by the bot to a specific position
func (b TelegramBot) SetStickerPositionInSet(config SetStickerPositionInSetConfig) (success bool, err error) {
	return success, b.sendResuest("setStickerPositionInSet", config, &success)
}

// DeleteStickerFromSet allows to delete a sticker from a set created by the bot
func (b TelegramBot) DeleteStickerFromSet(config DeleteStickerFromSetConfig) (success bool, err error) {
	return success, b.sendResuest("deleteStickerFromSet", config, &success)
}

// AnswerInlineQuery allows to send answers to an inline query
func (b TelegramBot) AnswerInlineQuery(config AnswerInlineQueryConfig) (success bool, err error) {
	return success, b.sendResuest("answerInlineQuery", config, &success)
}

func (b TelegramBot) sendResuest(method string, paramsObject interface{}, t interface{}) error {
	return sendResuest(method, b.apiKey, paramsObject, &t)
}
