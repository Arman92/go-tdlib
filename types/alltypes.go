package types

import "encoding/json"

// GetChatsReq ...
type GetChatsReq struct {
	TdClass
	OffsetOrder  int64 `json:"offset_order"`
	OffsetChatID int64 `json:"offset_chat_id"`
	Limit        int32 `json:"limit"`
}

// Chats ...
type Chats struct {
	ChatIDs []int64 `json:"chat_ids"`
}

// Chat ...
type Chat struct {
	ID                      int64                `json:"id"`
	Type                    json.RawMessage      `json:"type"`
	Title                   string               `json:"title"`
	Photo                   ChatPhoto            `json:"photo"`
	LastMessage             Message              `json:"last_message"`
	Order                   int64                `json:"order"`
	IsPinned                bool                 `json:"is_pinned"`
	CanBeReported           bool                 `json:"can_be_reported"`
	UnreadCount             int32                `json:"unread_count"`
	LastReadInboxMessageID  int64                `json:"last_read_inbox_message_id"`
	LastReadOutboxMessageID int64                `json:"last_read_outbox_message_id"`
	UnreadMentionCount      int32                `json:"unread_mention_count"`
	NotificationSettings    NotificationSettings `json:"notification_settings"`
	ReplyMarkupMessageID    int64                `json:"reply_markup_message_id"`
	DraftMessage            json.RawMessage      `json:"draft_message"`
	ClientData              string               `json:"client_data"`
}

// Message ...
type Message struct {
	ID                      int64           `json:"id"`
	SenderUserID            int32           `json:"sender_user_id"`
	ChatID                  int64           `json:"chat_id"`
	SendingState            json.RawMessage `json:"sending_state"`
	IsOutgoing              bool            `json:"is_outgoing"`
	CanBeEdited             bool            `json:"can_be_edited"`
	CanBeForwarded          bool            `json:"can_be_forwarded"`
	CanBeDeletedForSelf     bool            `json:"can_be_deleted_only_for_self"`
	CanBeDeletedForAllUsers bool            `json:"can_be_deleted_for_all_users"`
	IsChannelPost           bool            `json:"is_channel_post"`
	ContainsUnreadMention   bool            `json:"contains_unread_mention"`
	Date                    int32           `json:"date"`
	EditDate                int32           `json:"edit_date"`
	ForwardInfo             json.RawMessage `json:"forward_info"`
	ReplyToMessageID        int64           `json:"reply_to_message_id"`
	TTL                     int32           `json:"ttl"`
	ExpiresIn               float64         `json:"expires_in"`
	ViaBotUserID            int32           `json:"via_bot_user_id"`
	AuthorSignature         string          `json:"author_signature"`
	Views                   int32           `json:"views"`
	MediaAlbumID            int64           `json:"media_album_id"`
	Content                 int64           `json:"content"`
	ReplyMarkup             json.RawMessage `json:"reply_markup"`
}

// ChatPhoto ...
type ChatPhoto struct {
	Small string `json:"small"`
	Big   string `json:"big"`
}

// NotificationSettings ...
type NotificationSettings struct {
	MuteFor     int32  `json:"mute_for"`
	Sound       string `json:"sound"`
	ShowPreview bool   `json:"show_preview"`
}

// ChatMember ...
type ChatMember struct {
	UserID         int32            `json:"user_id"`
	InviterUserID  int32            `json:"inviter_user_id"`
	JoinedChatDate int32            `json:"joined_chat_date"`
	Status         ChatMemberStatus `json:"status"`
	BotInfo        BotInfo          `json:"bot_info"`
}

// BotInfo ...
type BotInfo struct {
	Description string        `json:"description"`
	Commands    []BotCommands `json:"inviter_user_id"`
}

// BotCommands ..
type BotCommands struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ChatMemberStatus ..
type ChatMemberStatus struct {
	Type                string `json:"@type"`
	StatusAdministrator ChatMemberStatusAdministrator
	StatusBanned        ChatMemberStatusBanned
	StatusCreator       ChatMemberStatusCreator
	StatusRestricted    ChatMemberStatusRestricted
}

// ChatMemberStatusAdministrator ..
type ChatMemberStatusAdministrator struct {
	CanBeEdited        bool `json:"can_be_edited"`
	CanChangeInfo      bool `json:"can_change_info"`
	CanPostMessages    bool `json:"can_post_messages"`
	CanEditMessages    bool `json:"	can_edit_messages"`
	CanDeleteMessages  bool `json:"can_delete_messages"`
	CanInviteUsers     bool `json:"can_invite_users"`
	CanRestrictMembers bool `json:"can_restrict_members"`
	CanPinMessages     bool `json:"can_pin_messages"`
	CanPromoteMembers  bool `json:"can_promote_members"`
}

// ChatMemberStatusBanned ..
type ChatMemberStatusBanned struct {
	BannedUntilDate int32 `json:"banned_until_date"`
}

// ChatMemberStatusCreator ..
type ChatMemberStatusCreator struct {
	IsMember bool `json:"is_member"`
}

// ChatMemberStatusRestricted ..
type ChatMemberStatusRestricted struct {
	RestrictedUntilDate   int32 `json:"restricted_until_date"`
	CanSendMessages       bool  `json:"can_send_messages"`
	CanSendMediaMessages  bool  `json:"can_send_media_messages"`
	CanSendOtherMessages  bool  `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews"`
}

// User ..
type User struct {
	ID               int32           `json:"id"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	UserName         string          `json:"username"`
	PhoneNumber      string          `json:"phone_number"`
	Status           json.RawMessage `json:"status"`
	ProfilePhoto     ProfilePhoto    `json:"profile_photo"`
	OutgoingLink     json.RawMessage `json:"outgoing_link"`
	IncomingLink     json.RawMessage `json:"incoming_link"`
	IsVerifed        bool            `json:"is_verified"`
	RestricionReason string          `json:"restriction_reason"`
	HaveAccess       bool            `json:"have_access"`
	Type             json.RawMessage `json:"type"`
	LanguageCode     string          `json:"language_code"`
}

// ProfilePhoto ...
type ProfilePhoto struct {
	Small string `json:"small"`
	Big   string `json:"big"`
}
