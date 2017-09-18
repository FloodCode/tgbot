package tgbot

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// API method parameters

// ParamsGetUpdates represents parameters for GetUpdates method.
type ParamsGetUpdates struct {
	Offset  int `option:"offset"`
	Limit   int `option:"limit"`
	Timeout int `option:"timeout"`
}

// ParamsSendMessage represents parameters for SendMessage method.
type ParamsSendMessage struct {
	ChatID                *ChatIdentifier `option:"chat_id"                 required:"true"`
	Text                  string          `option:"text"                    required:"true"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	DisableNotification   bool            `option:"disable_notification"`
	ReplyToMessageID      int             `option:"reply_to_message_id"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

// ParamsForwardMessage represents parameters for ForwardMessage method.
type ParamsForwardMessage struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	FromChatID          *ChatIdentifier `option:"from_chat_id"              required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	MessageID           int             `option:"message_id"                required:"true"`
}

// ParamsSendPhoto represents parameters for SendPhoto method.
type ParamsSendPhoto struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Photo               *InputFile      `option:"photo"                     required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendAudio represents parameters for SendAudio method.
type ParamsSendAudio struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Audio               *InputFile      `option:"audio"                     required:"true"`
	Caption             string          `option:"caption"`
	Duration            int             `option:"duration"`
	Performer           string          `option:"performer"`
	Title               string          `option:"title"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendDocument represents parameters for SendDocument method.
type ParamsSendDocument struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Document            *InputFile      `option:"document"                  required:"true"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVideo represents parameters for SendVideo method.
type ParamsSendVideo struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Video               *InputFile      `option:"video"                     required:"true"`
	Duration            int             `option:"duration"`
	Width               int             `option:"width"`
	Height              int             `option:"height"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVoice represents parameters for SendVoice method.
type ParamsSendVoice struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Voice               *InputFile      `option:"voice"                     required:"true"`
	Duration            int             `option:"duration"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVideoNote represents parameters for SendVideoNote method.
type ParamsSendVideoNote struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	VideoNote           *InputFile      `option:"video_note"                required:"true"`
	Duration            int             `option:"duration"`
	Length              int             `option:"length"`
	Caption             string          `option:"caption"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendLocation represents parameters for SendLocation method.
type ParamsSendLocation struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Latitude            float64         `option:"latitude"                  required:"true"`
	Longitude           float64         `option:"longitude"                 required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendVenue represents parameters for SendVenue method.
type ParamsSendVenue struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Latitude            float64         `option:"latitude"                  required:"true"`
	Longitude           float64         `option:"longitude"                 required:"true"`
	Title               string          `option:"title"                     required:"true"`
	Address             string          `option:"address"                   required:"true"`
	FoursquareID        string          `option:"foursquare_id"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendContact represents parameters for SendContact method.
type ParamsSendContact struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	PhoneNumber         string          `option:"phone_number"              required:"true"`
	FirstName           string          `option:"first_name"                required:"true"`
	LastName            string          `option:"last_name"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendSticker represents parameters for SendSticker method.
type ParamsSendSticker struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	Sticker             *InputFile      `option:"sticker"                   required:"true"`
	DisableNotification bool            `option:"disable_notification"`
	ReplyToMessageID    int             `option:"reply_to_message_id"`
	ReplyMarkup         *ReplyMarkup    `option:"reply_markup"`
}

// ParamsSendChatAction represents parameters for SendChatAction method.
type ParamsSendChatAction struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
	Action *ChatAction     `option:"action"                                 required:"true"`
}

// ParamsGetUserProfilePhotos represents parameters for GetUserProfilePhotos method.
type ParamsGetUserProfilePhotos struct {
	UserID int `option:"user_id"                                            required:"true"`
	Offset int `option:"offset"`
	Limit  int `option:"limit"`
}

// ParamsGetFile represents parameters for GetFile method.
type ParamsGetFile struct {
	FileID string `option:"file_id"                                         required:"true"`
}

// ParamsKickChatMember represents parameters for KickChatMember method.
type ParamsKickChatMember struct {
	ChatID    *ChatIdentifier `option:"chat_id"                             required:"true"`
	UserID    int             `option:"user_id"                             required:"true"`
	UntilDate int             `option:"until_date"`
}

// ParamsUnbanChatMember represents parameters for KickChatMember method.
type ParamsUnbanChatMember struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
	UserID int             `option:"user_id"                                required:"true"`
}

// ParamsRestrictChatMember represents parameters for RestrictChatMember method.
type ParamsRestrictChatMember struct {
	ChatID                *ChatIdentifier `option:"chat_id"                 required:"true"`
	UserID                int             `option:"user_id"                 required:"true"`
	UntilDate             int             `option:"until_date"`
	CanSendMessages       bool            `option:"can_send_messages"`
	CanSendMediaMessages  bool            `option:"can_send_media_messages"`
	CanSendOtherMessages  bool            `option:"can_send_other_messages"`
	CanAddWebPagePreviews bool            `option:"can_add_web_page_previews"`
}

// ParamsPromoteChatMember represents parameters for PromoteChatMember method.
type ParamsPromoteChatMember struct {
	ChatID             *ChatIdentifier `option:"chat_id"                    required:"true"`
	UserID             int             `option:"user_id"                    required:"true"`
	CanChangeInfo      bool            `option:"can_change_info"`
	CanPostMessages    bool            `option:"can_post_messages"`
	CanEditMessages    bool            `option:"can_edit_messages"`
	CanDeleteMessages  bool            `option:"can_delete_messages"`
	CanInviteUsers     bool            `option:"can_invite_users"`
	CanRestrictMembers bool            `option:"can_restrict_members"`
	CanPinMessages     bool            `option:"can_pin_messages"`
	CanPromoteMembers  bool            `option:"can_promote_members"`
}

// ParamsExportChatInviteLink represents parameters for ExportChatInviteLink method.
type ParamsExportChatInviteLink struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsSetChatPhoto represents parameters for SetChatPhoto method.
type ParamsSetChatPhoto struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
	Photo  *InputFile      `option:"photo"                                  required:"true"`
}

// ParamsDeleteChatPhoto represents parameters for DeleteChatPhoto method.
type ParamsDeleteChatPhoto struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsSetChatTitle represents parameters for SetChatTitle method.
type ParamsSetChatTitle struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
	Title  string          `option:"title"                                  required:"true"`
}

// ParamsSetChatDescription represents parameters for SetChatDescription method.
type ParamsSetChatDescription struct {
	ChatID      *ChatIdentifier `option:"chat_id"                           required:"true"`
	Description string          `option:"description"                       required:"true"`
}

// ParamsPinChatMessage represents parameters for PinChatMessage method.
type ParamsPinChatMessage struct {
	ChatID              *ChatIdentifier `option:"chat_id"                   required:"true"`
	MessageID           int             `option:"message_id"                required:"true"`
	DisableNotification bool            `option:"disable_notification"`
}

// ParamsUnpinChatMessage represents parameters for UnpinChatMessage method.
type ParamsUnpinChatMessage struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsLeaveChat represents parameters for LeaveChat method.
type ParamsLeaveChat struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsGetChat represents parameters for GetChat method.
type ParamsGetChat struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsGetChatAdministrators represents parameters for GetChatAdministrators method.
type ParamsGetChatAdministrators struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsGetChatMembersCount represents parameters for GetChatMembersCount method.
type ParamsGetChatMembersCount struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
}

// ParamsGetChatMember represents parameters for GetChatMember method.
type ParamsGetChatMember struct {
	ChatID *ChatIdentifier `option:"chat_id"                                required:"true"`
	UserID int             `option:"user_id"`
}

// ParamsEditMessageText represents parameters for EditMessageText method.
type ParamsEditMessageText struct {
	ChatID                *ChatIdentifier `option:"chat_id"`
	MessageID             int             `option:"message_id"`
	InlineMessageID       string          `option:"inline_message_id"`
	Text                  string          `option:"text"                    required:"true"`
	ParseMode             *ParseMode      `option:"parse_mode"`
	DisableWebPagePreview bool            `option:"disable_web_page_preview"`
	ReplyMarkup           *ReplyMarkup    `option:"reply_markup"`
}

// ParamsEditMessageCaption represents parameters for EditMessageCaption method.
type ParamsEditMessageCaption struct {
	ChatID          *ChatIdentifier `option:"chat_id"`
	MessageID       int             `option:"message_id"`
	InlineMessageID string          `option:"inline_message_id"`
	Caption         string          `option:"caption"                       required:"true"`
	ReplyMarkup     *ReplyMarkup    `option:"reply_markup"`
}

// ParamsGetStickerSet represents parameters for GetStickerSet method.
type ParamsGetStickerSet struct {
	Name string `option:"name"                                              required:"true"`
}

// ParamsUploadStickerFile represents parameters for UploadStickerFile method.
type ParamsUploadStickerFile struct {
	UserID     int        `option:"user_id"                                 required:"true"`
	PNGSticker *InputFile `option:"png_sticker"                             required:"true"`
}

// ParamsCreateNewStickerSet represents parameters for CreateNewStickerSet method.
type ParamsCreateNewStickerSet struct {
	UserID       int           `option:"user_id"                            required:"true"`
	Name         string        `option:"name"                               required:"true"`
	Title        string        `option:"title"                              required:"true"`
	PNGSticker   *InputFile    `option:"png_sticker"                        required:"true"`
	Emojis       string        `option:"emojis"                             required:"true"`
	ContainsMask bool          `option:"contains_masks"`
	MaskPosition *MaskPosition `option:"mask_position"`
}

// ParamsAddStickerToSet represents parameters for AddStickerToSet method.
type ParamsAddStickerToSet struct {
	UserID       int           `option:"user_id"                            required:"true"`
	Name         string        `option:"name"                               required:"true"`
	PNGSticker   *InputFile    `option:"png_sticker"                        required:"true"`
	Emojis       string        `option:"emojis"                             required:"true"`
	MaskPosition *MaskPosition `option:"mask_position"`
}

// ParamsSetStickerPositionInSet represents parameters for SetStickerPositionInSet method.
type ParamsSetStickerPositionInSet struct {
	Sticker  string `option:"sticker"                                       required:"true"`
	Position int    `option:"position"                                      required:"true"`
}

// ParamsDeleteStickerFromSet represents parameters for DeleteStickerFromSet method.
type ParamsDeleteStickerFromSet struct {
	Sticker string `option:"sticker"                                        required:"true"`
}

// StringParam is an interface for string method parameters
type StringParam interface {
	get() string
}

// API method option ChatIdentifier

// ChatIdentifier represents unique identifier of chat.
type ChatIdentifier struct {
	chatID string
}

// Get returns string representation of chat ID.
func (c *ChatIdentifier) get() string {
	return c.chatID
}

// ChatID creates new ChatIdentifier by chat ID.
func ChatID(id int64) *ChatIdentifier {
	return &ChatIdentifier{chatID: strconv.FormatInt(id, 10)}
}

// Username creates new ChatIdentifier by chat username.
func Username(username string) *ChatIdentifier {
	if strings.HasPrefix(username, "@") {
		return &ChatIdentifier{chatID: username}
	}

	return &ChatIdentifier{chatID: "@" + username}
}

// API method option ParseMode

// ParseMode represents message parse mode.
type ParseMode struct {
	parseMode string
}

// Get returns string representation of parse mode.
func (p *ParseMode) get() string {
	return p.parseMode
}

// ParseModeMarkdown creates new ParseMode with Markdown option.
func ParseModeMarkdown() *ParseMode {
	return &ParseMode{parseMode: "Markdown"}
}

// ParseModeHTML creates new ParseMode with HTML option.
func ParseModeHTML() *ParseMode {
	return &ParseMode{parseMode: "HTML"}
}

// API method option ReplyMarkup

// ReplyMarkup represents message reply markup
type ReplyMarkup struct {
	markup string
}

// Get returns markup data in JSON format.
func (m *ReplyMarkup) get() string {
	return m.markup
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
func InlineKeyboardMarkup(inlineKeyboard [][]InlineKeyboardButton) *ReplyMarkup {
	keyboardJSON, _ := json.Marshal(inlineKeyboard)
	return &ReplyMarkup{markup: `{"inline_keyboard":` + string(keyboardJSON) + `}`}
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
func ReplyKeyboardMarkup(keyboard [][]KeyboardButton, resizekeyboard, oneTimeKeyboard, selective bool) *ReplyMarkup {
	keyboardJSON, _ := json.Marshal(keyboard)
	resultJSON := `{"keyboard":` + string(keyboardJSON)

	if resizekeyboard {
		resultJSON += `,"resize_keyboard":true`
	}

	if oneTimeKeyboard {
		resultJSON += `,"one_time_keyboard":true`
	}

	if selective {
		resultJSON += `,"selective":true`
	}

	resultJSON += `}`
	return &ReplyMarkup{markup: resultJSON}
}

// ReplyKeyboardRemove represents reply markup with removal option.
func ReplyKeyboardRemove() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"remove_keyboard":true}`}
}

// ReplyKeyboardRemoveSelective represents reply markup with selective removal option.
func ReplyKeyboardRemoveSelective() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"remove_keyboard":true,"selective":true}`}
}

// ForceReply shows reply interface to user,
// as if they manually selected the bot‘s message and tapped ’Reply'.
func ForceReply() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"force_reply":true}`}
}

// ForceReplySelective selectively shows reply interface to user,
// as if they manually selected the bot‘s message and tapped ’Reply'.
func ForceReplySelective() *ReplyMarkup {
	return &ReplyMarkup{markup: `{"force_reply":true,"selective":true}`}
}

// API method option InputFile

// InputFile represents file to send.
type InputFile struct {
	fileData interface{}
	filename string
}

// Get returns file data.
func (f *InputFile) Get() interface{} {
	return f.fileData
}

// GetBytes returns file data byte array.
func (f *InputFile) GetBytes() []byte {
	if v, ok := f.fileData.([]byte); ok {
		return v
	}

	return []byte{}
}

// GetFilename returns filename.
func (f *InputFile) GetFilename() string {
	return f.filename
}

// FileID creates new InputFile by file id.
func FileID(fileID string) *InputFile {
	return &InputFile{fileData: fileID}
}

// FileBytes creates new InputFile by array of bytes and filename.
func FileBytes(fileData []byte, filename string) *InputFile {
	return &InputFile{fileData: fileData, filename: filename}
}

// FilePath creates new InputFile by file path.
func FilePath(filePath string) *InputFile {
	bytes, _ := ioutil.ReadFile(filePath)
	return &InputFile{fileData: bytes, filename: filepath.Base(filePath)}
}

// API method option ChatAction

// ChatAction represents chat action.
type ChatAction struct {
	action string
}

// Get returns string representation of chat action.
func (p *ChatAction) get() string {
	return p.action
}

// ChatActionTyping creates ChatAction with "typing" option.
func ChatActionTyping() *ChatAction {
	return &ChatAction{action: "typing"}
}

// ChatActionUploadPhoto creates ChatAction with "upload_photo" option.
func ChatActionUploadPhoto() *ChatAction {
	return &ChatAction{action: "upload_photo"}
}

// ChatActionUploadVideo creates ChatAction with "upload_video" option.
func ChatActionUploadVideo() *ChatAction {
	return &ChatAction{action: "upload_video"}
}

// ChatActionUploadAudio creates ChatAction with "upload_audio" option.
func ChatActionUploadAudio() *ChatAction {
	return &ChatAction{action: "upload_audio"}
}

// ChatActionUploadDocument creates ChatAction with "upload_document" option.
func ChatActionUploadDocument() *ChatAction {
	return &ChatAction{action: "upload_document"}
}

// ChatActionFindLocation creates ChatAction with "find_location" option.
func ChatActionFindLocation() *ChatAction {
	return &ChatAction{action: "find_location"}
}

// ChatActionUploadVideoNote creates ChatAction with "upload_video_note" option.
func ChatActionUploadVideoNote() *ChatAction {
	return &ChatAction{action: "upload_video_note"}
}
