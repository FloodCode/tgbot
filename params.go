package tgbot

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// API method parameters

// GetUpdatesConfig represents parameters for GetUpdates method
type GetUpdatesConfig struct {
	Offset  int `option:"offset"`
	Limit   int `option:"limit"`
	Timeout int `option:"timeout"`
}

// SetWebhookConfig represents parameters for SetWebhook method
type SetWebhookConfig struct {
	URL            string     `option:"url"`
	Certificate    *InputFile `option:"certificate"`
	MaxConnections int        `option:"max_connections"`
	AllowedUpdates []string   `option:"allowed_updates"`
}

// SendMessageConfig represents parameters for SendMessage method
type SendMessageConfig struct {
	ChatID                *ChatIdentifier `option:"chat_id"`
	Text                  string          `option:"text"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	DisableNotification   bool            `option:"disable_notification"`
	ReplyToMessageID      int             `option:"reply_to_message_id"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

// ForwardMessageConfig represents parameters for ForwardMessage method
type ForwardMessageConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	FromChatID          *ChatIdentifier `option:"from_chat_id"`
	DisableNotification bool            `option:"disable_notification"`
	MessageID           int             `option:"message_id"`
}

// SendPhotoConfig represents parameters for SendPhoto method
type SendPhotoConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Photo               *InputFile      `option:"photo"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendAudioConfig represents parameters for SendAudio method
type SendAudioConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Audio               *InputFile      `option:"audio"`
	Caption             string          `option:"caption"`
	Duration            int             `option:"duration"`
	Performer           string          `option:"performer"`
	Title               string          `option:"title"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendDocumentConfig represents parameters for SendDocument method
type SendDocumentConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Document            *InputFile      `option:"document"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendVideoConfig represents parameters for SendVideo method
type SendVideoConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Video               *InputFile      `option:"video"`
	Duration            int             `option:"duration"`
	Width               int             `option:"width"`
	Height              int             `option:"height"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendVoiceConfig represents parameters for SendVoice method
type SendVoiceConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Voice               *InputFile      `option:"voice"`
	Duration            int             `option:"duration"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendVideoNoteConfig represents parameters for SendVideoNote method
type SendVideoNoteConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	VideoNote           *InputFile      `option:"video_note"`
	Duration            int             `option:"duration"`
	Length              int             `option:"length"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendLocationConfig represents parameters for SendLocation method
type SendLocationConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Latitude            float64         `option:"latitude"`
	Longitude           float64         `option:"longitude"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendVenueConfig represents parameters for SendVenue method
type SendVenueConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Latitude            float64         `option:"latitude"`
	Longitude           float64         `option:"longitude"`
	Title               string          `option:"title"`
	Address             string          `option:"address"`
	FoursquareID        string          `option:"foursquare_id"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendContactConfig represents parameters for SendContact method
type SendContactConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	PhoneNumber         string          `option:"phone_number"`
	FirstName           string          `option:"first_name"`
	LastName            string          `option:"last_name"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendStickerConfig represents parameters for SendSticker method
type SendStickerConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	Sticker             *InputFile      `option:"sticker"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// SendChatActionConfig represents parameters for SendChatAction method
type SendChatActionConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
	Action *ChatAction     `option:"action"`
}

// GetUserProfilePhotosConfig represents parameters for GetUserProfilePhotos method
type GetUserProfilePhotosConfig struct {
	UserID int `option:"user_id"`
	Offset int `option:"offset"`
	Limit  int `option:"limit"`
}

// GetFileConfig represents parameters for GetFile method
type GetFileConfig struct {
	FileID string `option:"file_id"`
}

// KickChatMemberConfig represents parameters for KickChatMember method
type KickChatMemberConfig struct {
	ChatID    *ChatIdentifier `option:"chat_id"`
	UserID    int             `option:"user_id"`
	UntilDate int             `option:"until_date"`
}

// UnbanChatMemberConfig represents parameters for KickChatMember method
type UnbanChatMemberConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
	UserID int             `option:"user_id"`
}

// RestrictChatMemberConfig represents parameters for RestrictChatMember method
type RestrictChatMemberConfig struct {
	ChatID                *ChatIdentifier `option:"chat_id"`
	UserID                int             `option:"user_id"`
	UntilDate             int             `option:"until_date"`
	CanSendMessages       bool            `option:"can_send_messages"`
	CanSendMediaMessages  bool            `option:"can_send_media_messages"`
	CanSendOtherMessages  bool            `option:"can_send_other_messages"`
	CanAddWebPagePreviews bool            `option:"can_add_web_page_previews"`
}

// PromoteChatMemberConfig represents parameters for PromoteChatMember method
type PromoteChatMemberConfig struct {
	ChatID             *ChatIdentifier `option:"chat_id"`
	UserID             int             `option:"user_id"`
	CanChangeInfo      bool            `option:"can_change_info"`
	CanPostMessages    bool            `option:"can_post_messages"`
	CanEditMessages    bool            `option:"can_edit_messages"`
	CanDeleteMessages  bool            `option:"can_delete_messages"`
	CanInviteUsers     bool            `option:"can_invite_users"`
	CanRestrictMembers bool            `option:"can_restrict_members"`
	CanPinMessages     bool            `option:"can_pin_messages"`
	CanPromoteMembers  bool            `option:"can_promote_members"`
}

// ExportChatInviteLinkConfig represents parameters for ExportChatInviteLink method
type ExportChatInviteLinkConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// SetChatPhotoConfig represents parameters for SetChatPhoto method
type SetChatPhotoConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
	Photo  *InputFile      `option:"photo"`
}

// DeleteChatPhotoConfig represents parameters for DeleteChatPhoto method
type DeleteChatPhotoConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// SetChatTitleConfig represents parameters for SetChatTitle method
type SetChatTitleConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
	Title  string          `option:"title"`
}

// SetChatDescriptionConfig represents parameters for SetChatDescription method
type SetChatDescriptionConfig struct {
	ChatID      *ChatIdentifier `option:"chat_id"`
	Description string          `option:"description"`
}

// PinChatMessageConfig represents parameters for PinChatMessage method
type PinChatMessageConfig struct {
	ChatID              *ChatIdentifier `option:"chat_id"`
	MessageID           int             `option:"message_id"`
	DisableNotification bool            `option:"disable_notification"`
}

// UnpinChatMessageConfig represents parameters for UnpinChatMessage method
type UnpinChatMessageConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// LeaveChatConfig represents parameters for LeaveChat method
type LeaveChatConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// GetChatConfig represents parameters for GetChat method
type GetChatConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// GetChatAdministratorsConfig represents parameters for GetChatAdministrators method
type GetChatAdministratorsConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// GetChatMembersCountConfig represents parameters for GetChatMembersCount method
type GetChatMembersCountConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
}

// GetChatMemberConfig represents parameters for GetChatMember method
type GetChatMemberConfig struct {
	ChatID *ChatIdentifier `option:"chat_id"`
	UserID int             `option:"user_id"`
}

// EditMessageTextConfig represents parameters for EditMessageText method
type EditMessageTextConfig struct {
	ChatID                *ChatIdentifier `option:"chat_id"`
	MessageID             int             `option:"message_id"`
	InlineMessageID       string          `option:"inline_message_id"`
	Text                  string          `option:"text"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

// EditMessageCaptionConfig represents parameters for EditMessageCaption method
type EditMessageCaptionConfig struct {
	ChatID          *ChatIdentifier `option:"chat_id"`
	MessageID       int             `option:"message_id"`
	InlineMessageID string          `option:"inline_message_id"`
	Caption         string          `option:"caption"`
	ReplyMarkup     *ReplyMarkup    `option:"reply_markup"`
}

// GetStickerSetConfig represents parameters for GetStickerSet method
type GetStickerSetConfig struct {
	Name string `option:"name"`
}

// UploadStickerFileConfig represents parameters for UploadStickerFile method
type UploadStickerFileConfig struct {
	UserID     int        `option:"user_id"`
	PNGSticker *InputFile `option:"png_sticker"`
}

// CreateNewStickerSetConfig represents parameters for CreateNewStickerSet method
type CreateNewStickerSetConfig struct {
	UserID       int           `option:"user_id"`
	Name         string        `option:"name"`
	Title        string        `option:"title"`
	PNGSticker   *InputFile    `option:"png_sticker"`
	Emojis       string        `option:"emojis"`
	ContainsMask bool          `option:"contains_masks"`
	MaskPosition *MaskPosition `option:"mask_position"`
}

// AddStickerToSetConfig represents parameters for AddStickerToSet method
type AddStickerToSetConfig struct {
	UserID       int           `option:"user_id"`
	Name         string        `option:"name"`
	PNGSticker   *InputFile    `option:"png_sticker"`
	Emojis       string        `option:"emojis"`
	MaskPosition *MaskPosition `option:"mask_position"`
}

// SetStickerPositionInSetConfig represents parameters for SetStickerPositionInSet method
type SetStickerPositionInSetConfig struct {
	Sticker  string `option:"sticker"`
	Position int    `option:"position"`
}

// DeleteStickerFromSetConfig represents parameters for DeleteStickerFromSet method
type DeleteStickerFromSetConfig struct {
	Sticker string `option:"sticker"`
}

// AnswerInlineQueryConfig represents parameters for AnswerInlineQuery method
type AnswerInlineQueryConfig struct {
	InlineQueryID     string              `option:"inline_query_id"`
	Results           []InlineQueryResult `option:"results"`
	CacheTime         int                 `option:"cache_time"`
	IsPersonal        bool                `option:"is_personal"`
	NextOffset        string              `option:"next_offset"`
	SwitchPmText      string              `option:"switch_pm_text"`
	SwitchPmParameter string              `option:"switch_pm_parameter"`
}

// AnswerCallbackQueryConfig represents parameters for AnswerCallbackQuery method
type AnswerCallbackQueryConfig struct {
	CallbackQueryID string `option:"callback_query_id"`
	Text            string `option:"text"`
	ShowAlert       bool   `option:"show_alert"`
	URL             string `option:"url"`
	CacheTime       int    `option:"cache_time"`
}

type stringConfig interface {
	getString() string
}

// API method option ChatIdentifier

// ChatIdentifier represents unique identifier of chat
type ChatIdentifier struct {
	stringConfig
	chatID string
}

func (c *ChatIdentifier) getString() string {
	return c.chatID
}

// ChatID creates new ChatIdentifier by chat ID
func ChatID(id int64) *ChatIdentifier {
	return &ChatIdentifier{chatID: strconv.FormatInt(id, 10)}
}

// Username creates new ChatIdentifier by chat username
func Username(username string) *ChatIdentifier {
	if strings.HasPrefix(username, "@") {
		return &ChatIdentifier{chatID: username}
	}

	return &ChatIdentifier{chatID: "@" + username}
}

// API method option ParseMode

// ParseMode represents message parse mode
type ParseMode struct {
	stringConfig
	parseMode string
}

func (p *ParseMode) getString() string {
	return p.parseMode
}

// ParseModeMarkdown creates new ParseMode with Markdown option
func ParseModeMarkdown() *ParseMode {
	return &ParseMode{parseMode: "Markdown"}
}

// ParseModeHTML creates new ParseMode with HTML option
func ParseModeHTML() *ParseMode {
	return &ParseMode{parseMode: "HTML"}
}

// API method option ReplyMarkup

// ReplyMarkup represents message reply markup
type ReplyMarkup struct {
	stringConfig
	markup string
}

func (m *ReplyMarkup) getString() string {
	return m.markup
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to
func InlineKeyboardMarkup(inlineKeyboard [][]InlineKeyboardButton) *ReplyMarkup {
	keyboardJSON, _ := json.Marshal(map[string]interface{}{"inline_keyboard": inlineKeyboard})
	return &ReplyMarkup{markup: string(keyboardJSON)}
}

// ReplyKeyboardMarkupConfig represents parameters for ReplyKeyboardMarkup method
type ReplyKeyboardMarkupConfig struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options
func ReplyKeyboardMarkup(config ReplyKeyboardMarkupConfig) *ReplyMarkup {
	markupJSON, _ := json.Marshal(config)
	return &ReplyMarkup{markup: string(markupJSON)}
}

// ReplyKeyboardRemove represents reply markup with removal option
func ReplyKeyboardRemove() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"remove_keyboard":true}`}
}

// ReplyKeyboardRemoveSelective represents reply markup with selective removal option
func ReplyKeyboardRemoveSelective() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"remove_keyboard":true,"selective":true}`}
}

// ForceReply shows reply interface to user,
// as if they manually selected the bot‘s message and tapped ’Reply'
func ForceReply() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"force_reply":true}`}
}

// ForceReplySelective selectively shows reply interface to user,
// as if they manually selected the bot‘s message and tapped ’Reply'
func ForceReplySelective() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"force_reply":true,"selective":true}`}
}

// API method option InputFile

// InputFile represents file to send
type InputFile struct {
	fileData interface{}
	filename string
}

func (f *InputFile) getData() interface{} {
	return f.fileData
}

func (f *InputFile) getBytes() []byte {
	if v, ok := f.fileData.([]byte); ok {
		return v
	}

	return []byte{}
}

func (f *InputFile) getFilename() string {
	return f.filename
}

// FileID creates new InputFile by file id
func FileID(fileID string) *InputFile {
	return &InputFile{fileData: fileID}
}

// FileBytes creates new InputFile by array of bytes and filename
func FileBytes(fileData []byte, filename string) *InputFile {
	return &InputFile{fileData: fileData, filename: filename}
}

// FilePath creates new InputFile by file path
func FilePath(filePath string) *InputFile {
	bytes, _ := ioutil.ReadFile(filePath)
	return &InputFile{fileData: bytes, filename: filepath.Base(filePath)}
}

// API method option ChatAction

// ChatAction represents chat action
type ChatAction struct {
	stringConfig
	action string
}

func (p *ChatAction) getString() string {
	return p.action
}

// ChatActionTyping creates ChatAction with "typing" option
func ChatActionTyping() *ChatAction {
	return &ChatAction{action: "typing"}
}

// ChatActionUploadPhoto creates ChatAction with "upload_photo" option
func ChatActionUploadPhoto() *ChatAction {
	return &ChatAction{action: "upload_photo"}
}

// ChatActionUploadVideo creates ChatAction with "upload_video" option
func ChatActionUploadVideo() *ChatAction {
	return &ChatAction{action: "upload_video"}
}

// ChatActionUploadAudio creates ChatAction with "upload_audio" option
func ChatActionUploadAudio() *ChatAction {
	return &ChatAction{action: "upload_audio"}
}

// ChatActionUploadDocument creates ChatAction with "upload_document" option
func ChatActionUploadDocument() *ChatAction {
	return &ChatAction{action: "upload_document"}
}

// ChatActionFindLocation creates ChatAction with "find_location" option
func ChatActionFindLocation() *ChatAction {
	return &ChatAction{action: "find_location"}
}

// ChatActionUploadVideoNote creates ChatAction with "upload_video_note" option
func ChatActionUploadVideoNote() *ChatAction {
	return &ChatAction{action: "upload_video_note"}
}
