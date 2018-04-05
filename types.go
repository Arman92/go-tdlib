package tdlib

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type tdCommon struct {
	Type  string `json:"@type"`
	Extra string `json:"@extra"`
}

// JSONInt64 alias for int64, in order to deal with json big number problem
type JSONInt64 int64

// MarshalJSON marshals to json
func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
	intStr := "\"" + strconv.FormatInt(int64(*jsonInt), 10) + "\""
	return []byte(intStr), nil
}

// UnmarshalJSON unmarshals from json
func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
	intStr := string(b)
	intStr = strings.Replace(intStr, "\"", "", 2)
	jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return err
	}
	*jsonInt = JSONInt64(jsonBigInt)
	return nil
}

// TdMessage is the interface for all messages send and received to/from tdlib
type TdMessage interface {
	MessageType() string
}

// AuthenticationCodeTypeEnum Alias for abstract AuthenticationCodeType 'Sub-Classes', used as constant-enum here
type AuthenticationCodeTypeEnum string

// AuthenticationCodeType enums
const (
	AuthenticationCodeTypeTelegramMessageType AuthenticationCodeTypeEnum = "authenticationCodeTypeTelegramMessage"
	AuthenticationCodeTypeSmsType             AuthenticationCodeTypeEnum = "authenticationCodeTypeSms"
	AuthenticationCodeTypeCallType            AuthenticationCodeTypeEnum = "authenticationCodeTypeCall"
	AuthenticationCodeTypeFlashCallType       AuthenticationCodeTypeEnum = "authenticationCodeTypeFlashCall"
)

// AuthorizationStateEnum Alias for abstract AuthorizationState 'Sub-Classes', used as constant-enum here
type AuthorizationStateEnum string

// AuthorizationState enums
const (
	AuthorizationStateWaitTdlibParametersType AuthorizationStateEnum = "authorizationStateWaitTdlibParameters"
	AuthorizationStateWaitEncryptionKeyType   AuthorizationStateEnum = "authorizationStateWaitEncryptionKey"
	AuthorizationStateWaitPhoneNumberType     AuthorizationStateEnum = "authorizationStateWaitPhoneNumber"
	AuthorizationStateWaitCodeType            AuthorizationStateEnum = "authorizationStateWaitCode"
	AuthorizationStateWaitPasswordType        AuthorizationStateEnum = "authorizationStateWaitPassword"
	AuthorizationStateReadyType               AuthorizationStateEnum = "authorizationStateReady"
	AuthorizationStateLoggingOutType          AuthorizationStateEnum = "authorizationStateLoggingOut"
	AuthorizationStateClosingType             AuthorizationStateEnum = "authorizationStateClosing"
	AuthorizationStateClosedType              AuthorizationStateEnum = "authorizationStateClosed"
)

// InputFileEnum Alias for abstract InputFile 'Sub-Classes', used as constant-enum here
type InputFileEnum string

// InputFile enums
const (
	InputFileIDType        InputFileEnum = "inputFileID"
	InputFileRemoteType    InputFileEnum = "inputFileRemote"
	InputFileLocalType     InputFileEnum = "inputFileLocal"
	InputFileGeneratedType InputFileEnum = "inputFileGenerated"
)

// MaskPointEnum Alias for abstract MaskPoint 'Sub-Classes', used as constant-enum here
type MaskPointEnum string

// MaskPoint enums
const (
	MaskPointForeheadType MaskPointEnum = "maskPointForehead"
	MaskPointEyesType     MaskPointEnum = "maskPointEyes"
	MaskPointMouthType    MaskPointEnum = "maskPointMouth"
	MaskPointChinType     MaskPointEnum = "maskPointChin"
)

// LinkStateEnum Alias for abstract LinkState 'Sub-Classes', used as constant-enum here
type LinkStateEnum string

// LinkState enums
const (
	LinkStateNoneType             LinkStateEnum = "linkStateNone"
	LinkStateKnowsPhoneNumberType LinkStateEnum = "linkStateKnowsPhoneNumber"
	LinkStateIsContactType        LinkStateEnum = "linkStateIsContact"
)

// UserTypeEnum Alias for abstract UserType 'Sub-Classes', used as constant-enum here
type UserTypeEnum string

// UserType enums
const (
	UserTypeRegularType UserTypeEnum = "userTypeRegular"
	UserTypeDeletedType UserTypeEnum = "userTypeDeleted"
	UserTypeBotType     UserTypeEnum = "userTypeBot"
	UserTypeUnknownType UserTypeEnum = "userTypeUnknown"
)

// ChatMemberStatusEnum Alias for abstract ChatMemberStatus 'Sub-Classes', used as constant-enum here
type ChatMemberStatusEnum string

// ChatMemberStatus enums
const (
	ChatMemberStatusCreatorType       ChatMemberStatusEnum = "chatMemberStatusCreator"
	ChatMemberStatusAdministratorType ChatMemberStatusEnum = "chatMemberStatusAdministrator"
	ChatMemberStatusMemberType        ChatMemberStatusEnum = "chatMemberStatusMember"
	ChatMemberStatusRestrictedType    ChatMemberStatusEnum = "chatMemberStatusRestricted"
	ChatMemberStatusLeftType          ChatMemberStatusEnum = "chatMemberStatusLeft"
	ChatMemberStatusBannedType        ChatMemberStatusEnum = "chatMemberStatusBanned"
)

// SupergroupMembersFilterEnum Alias for abstract SupergroupMembersFilter 'Sub-Classes', used as constant-enum here
type SupergroupMembersFilterEnum string

// SupergroupMembersFilter enums
const (
	SupergroupMembersFilterRecentType         SupergroupMembersFilterEnum = "supergroupMembersFilterRecent"
	SupergroupMembersFilterAdministratorsType SupergroupMembersFilterEnum = "supergroupMembersFilterAdministrators"
	SupergroupMembersFilterSearchType         SupergroupMembersFilterEnum = "supergroupMembersFilterSearch"
	SupergroupMembersFilterRestrictedType     SupergroupMembersFilterEnum = "supergroupMembersFilterRestricted"
	SupergroupMembersFilterBannedType         SupergroupMembersFilterEnum = "supergroupMembersFilterBanned"
	SupergroupMembersFilterBotsType           SupergroupMembersFilterEnum = "supergroupMembersFilterBots"
)

// SecretChatStateEnum Alias for abstract SecretChatState 'Sub-Classes', used as constant-enum here
type SecretChatStateEnum string

// SecretChatState enums
const (
	SecretChatStatePendingType SecretChatStateEnum = "secretChatStatePending"
	SecretChatStateReadyType   SecretChatStateEnum = "secretChatStateReady"
	SecretChatStateClosedType  SecretChatStateEnum = "secretChatStateClosed"
)

// MessageForwardInfoEnum Alias for abstract MessageForwardInfo 'Sub-Classes', used as constant-enum here
type MessageForwardInfoEnum string

// MessageForwardInfo enums
const (
	MessageForwardedFromUserType MessageForwardInfoEnum = "messageForwardedFromUser"
	MessageForwardedPostType     MessageForwardInfoEnum = "messageForwardedPost"
)

// MessageSendingStateEnum Alias for abstract MessageSendingState 'Sub-Classes', used as constant-enum here
type MessageSendingStateEnum string

// MessageSendingState enums
const (
	MessageSendingStatePendingType MessageSendingStateEnum = "messageSendingStatePending"
	MessageSendingStateFailedType  MessageSendingStateEnum = "messageSendingStateFailed"
)

// NotificationSettingsScopeEnum Alias for abstract NotificationSettingsScope 'Sub-Classes', used as constant-enum here
type NotificationSettingsScopeEnum string

// NotificationSettingsScope enums
const (
	NotificationSettingsScopeChatType            NotificationSettingsScopeEnum = "notificationSettingsScopeChat"
	NotificationSettingsScopePrivateChatsType    NotificationSettingsScopeEnum = "notificationSettingsScopePrivateChats"
	NotificationSettingsScopeBasicGroupChatsType NotificationSettingsScopeEnum = "notificationSettingsScopeBasicGroupChats"
	NotificationSettingsScopeAllChatsType        NotificationSettingsScopeEnum = "notificationSettingsScopeAllChats"
)

// ChatTypeEnum Alias for abstract ChatType 'Sub-Classes', used as constant-enum here
type ChatTypeEnum string

// ChatType enums
const (
	ChatTypePrivateType    ChatTypeEnum = "chatTypePrivate"
	ChatTypeBasicGroupType ChatTypeEnum = "chatTypeBasicGroup"
	ChatTypeSupergroupType ChatTypeEnum = "chatTypeSupergroup"
	ChatTypeSecretType     ChatTypeEnum = "chatTypeSecret"
)

// KeyboardButtonTypeEnum Alias for abstract KeyboardButtonType 'Sub-Classes', used as constant-enum here
type KeyboardButtonTypeEnum string

// KeyboardButtonType enums
const (
	KeyboardButtonTypeTextType               KeyboardButtonTypeEnum = "keyboardButtonTypeText"
	KeyboardButtonTypeRequestPhoneNumberType KeyboardButtonTypeEnum = "keyboardButtonTypeRequestPhoneNumber"
	KeyboardButtonTypeRequestLocationType    KeyboardButtonTypeEnum = "keyboardButtonTypeRequestLocation"
)

// InlineKeyboardButtonTypeEnum Alias for abstract InlineKeyboardButtonType 'Sub-Classes', used as constant-enum here
type InlineKeyboardButtonTypeEnum string

// InlineKeyboardButtonType enums
const (
	InlineKeyboardButtonTypeURLType          InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeURL"
	InlineKeyboardButtonTypeCallbackType     InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeCallback"
	InlineKeyboardButtonTypeCallbackGameType InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeCallbackGame"
	InlineKeyboardButtonTypeSwitchInlineType InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeSwitchInline"
	InlineKeyboardButtonTypeBuyType          InlineKeyboardButtonTypeEnum = "inlineKeyboardButtonTypeBuy"
)

// ReplyMarkupEnum Alias for abstract ReplyMarkup 'Sub-Classes', used as constant-enum here
type ReplyMarkupEnum string

// ReplyMarkup enums
const (
	ReplyMarkupRemoveKeyboardType ReplyMarkupEnum = "replyMarkupRemoveKeyboard"
	ReplyMarkupForceReplyType     ReplyMarkupEnum = "replyMarkupForceReply"
	ReplyMarkupShowKeyboardType   ReplyMarkupEnum = "replyMarkupShowKeyboard"
	ReplyMarkupInlineKeyboardType ReplyMarkupEnum = "replyMarkupInlineKeyboard"
)

// RichTextEnum Alias for abstract RichText 'Sub-Classes', used as constant-enum here
type RichTextEnum string

// RichText enums
const (
	RichTextPlainType         RichTextEnum = "richTextPlain"
	RichTextBoldType          RichTextEnum = "richTextBold"
	RichTextItalicType        RichTextEnum = "richTextItalic"
	RichTextUnderlineType     RichTextEnum = "richTextUnderline"
	RichTextStrikethroughType RichTextEnum = "richTextStrikethrough"
	RichTextFixedType         RichTextEnum = "richTextFixed"
	RichTextURLType           RichTextEnum = "richTextURL"
	RichTextEmailAddressType  RichTextEnum = "richTextEmailAddress"
	RichTextsType             RichTextEnum = "richTexts"
)

// PageBlockEnum Alias for abstract PageBlock 'Sub-Classes', used as constant-enum here
type PageBlockEnum string

// PageBlock enums
const (
	PageBlockTitleType        PageBlockEnum = "pageBlockTitle"
	PageBlockSubtitleType     PageBlockEnum = "pageBlockSubtitle"
	PageBlockAuthorDateType   PageBlockEnum = "pageBlockAuthorDate"
	PageBlockHeaderType       PageBlockEnum = "pageBlockHeader"
	PageBlockSubheaderType    PageBlockEnum = "pageBlockSubheader"
	PageBlockParagraphType    PageBlockEnum = "pageBlockParagraph"
	PageBlockPreformattedType PageBlockEnum = "pageBlockPreformatted"
	PageBlockFooterType       PageBlockEnum = "pageBlockFooter"
	PageBlockDividerType      PageBlockEnum = "pageBlockDivider"
	PageBlockAnchorType       PageBlockEnum = "pageBlockAnchor"
	PageBlockListType         PageBlockEnum = "pageBlockList"
	PageBlockBlockQuoteType   PageBlockEnum = "pageBlockBlockQuote"
	PageBlockPullQuoteType    PageBlockEnum = "pageBlockPullQuote"
	PageBlockAnimationType    PageBlockEnum = "pageBlockAnimation"
	PageBlockAudioType        PageBlockEnum = "pageBlockAudio"
	PageBlockPhotoType        PageBlockEnum = "pageBlockPhoto"
	PageBlockVideoType        PageBlockEnum = "pageBlockVideo"
	PageBlockCoverType        PageBlockEnum = "pageBlockCover"
	PageBlockEmbeddedType     PageBlockEnum = "pageBlockEmbedded"
	PageBlockEmbeddedPostType PageBlockEnum = "pageBlockEmbeddedPost"
	PageBlockCollageType      PageBlockEnum = "pageBlockCollage"
	PageBlockSlideshowType    PageBlockEnum = "pageBlockSlideshow"
	PageBlockChatLinkType     PageBlockEnum = "pageBlockChatLink"
)

// InputCredentialsEnum Alias for abstract InputCredentials 'Sub-Classes', used as constant-enum here
type InputCredentialsEnum string

// InputCredentials enums
const (
	InputCredentialsSavedType      InputCredentialsEnum = "inputCredentialsSaved"
	InputCredentialsNewType        InputCredentialsEnum = "inputCredentialsNew"
	InputCredentialsAndroidPayType InputCredentialsEnum = "inputCredentialsAndroidPay"
	InputCredentialsApplePayType   InputCredentialsEnum = "inputCredentialsApplePay"
)

// MessageContentEnum Alias for abstract MessageContent 'Sub-Classes', used as constant-enum here
type MessageContentEnum string

// MessageContent enums
const (
	MessageTextType                 MessageContentEnum = "messageText"
	MessageAnimationType            MessageContentEnum = "messageAnimation"
	MessageAudioType                MessageContentEnum = "messageAudio"
	MessageDocumentType             MessageContentEnum = "messageDocument"
	MessagePhotoType                MessageContentEnum = "messagePhoto"
	MessageExpiredPhotoType         MessageContentEnum = "messageExpiredPhoto"
	MessageStickerType              MessageContentEnum = "messageSticker"
	MessageVideoType                MessageContentEnum = "messageVideo"
	MessageExpiredVideoType         MessageContentEnum = "messageExpiredVideo"
	MessageVideoNoteType            MessageContentEnum = "messageVideoNote"
	MessageVoiceNoteType            MessageContentEnum = "messageVoiceNote"
	MessageLocationType             MessageContentEnum = "messageLocation"
	MessageVenueType                MessageContentEnum = "messageVenue"
	MessageContactType              MessageContentEnum = "messageContact"
	MessageGameType                 MessageContentEnum = "messageGame"
	MessageInvoiceType              MessageContentEnum = "messageInvoice"
	MessageCallType                 MessageContentEnum = "messageCall"
	MessageBasicGroupChatCreateType MessageContentEnum = "messageBasicGroupChatCreate"
	MessageSupergroupChatCreateType MessageContentEnum = "messageSupergroupChatCreate"
	MessageChatChangeTitleType      MessageContentEnum = "messageChatChangeTitle"
	MessageChatChangePhotoType      MessageContentEnum = "messageChatChangePhoto"
	MessageChatDeletePhotoType      MessageContentEnum = "messageChatDeletePhoto"
	MessageChatAddMembersType       MessageContentEnum = "messageChatAddMembers"
	MessageChatJoinByLinkType       MessageContentEnum = "messageChatJoinByLink"
	MessageChatDeleteMemberType     MessageContentEnum = "messageChatDeleteMember"
	MessageChatUpgradeToType        MessageContentEnum = "messageChatUpgradeTo"
	MessageChatUpgradeFromType      MessageContentEnum = "messageChatUpgradeFrom"
	MessagePinMessageType           MessageContentEnum = "messagePinMessage"
	MessageScreenshotTakenType      MessageContentEnum = "messageScreenshotTaken"
	MessageChatSetTTLType           MessageContentEnum = "messageChatSetTTL"
	MessageCustomServiceActionType  MessageContentEnum = "messageCustomServiceAction"
	MessageGameScoreType            MessageContentEnum = "messageGameScore"
	MessagePaymentSuccessfulType    MessageContentEnum = "messagePaymentSuccessful"
	MessagePaymentSuccessfulBotType MessageContentEnum = "messagePaymentSuccessfulBot"
	MessageContactRegisteredType    MessageContentEnum = "messageContactRegistered"
	MessageWebsiteConnectedType     MessageContentEnum = "messageWebsiteConnected"
	MessageUnsupportedType          MessageContentEnum = "messageUnsupported"
)

// TextEntityTypeEnum Alias for abstract TextEntityType 'Sub-Classes', used as constant-enum here
type TextEntityTypeEnum string

// TextEntityType enums
const (
	TextEntityTypeMentionType      TextEntityTypeEnum = "textEntityTypeMention"
	TextEntityTypeHashtagType      TextEntityTypeEnum = "textEntityTypeHashtag"
	TextEntityTypeCashtagType      TextEntityTypeEnum = "textEntityTypeCashtag"
	TextEntityTypeBotCommandType   TextEntityTypeEnum = "textEntityTypeBotCommand"
	TextEntityTypeURLType          TextEntityTypeEnum = "textEntityTypeURL"
	TextEntityTypeEmailAddressType TextEntityTypeEnum = "textEntityTypeEmailAddress"
	TextEntityTypeBoldType         TextEntityTypeEnum = "textEntityTypeBold"
	TextEntityTypeItalicType       TextEntityTypeEnum = "textEntityTypeItalic"
	TextEntityTypeCodeType         TextEntityTypeEnum = "textEntityTypeCode"
	TextEntityTypePreType          TextEntityTypeEnum = "textEntityTypePre"
	TextEntityTypePreCodeType      TextEntityTypeEnum = "textEntityTypePreCode"
	TextEntityTypeTextURLType      TextEntityTypeEnum = "textEntityTypeTextURL"
	TextEntityTypeMentionNameType  TextEntityTypeEnum = "textEntityTypeMentionName"
	TextEntityTypePhoneNumberType  TextEntityTypeEnum = "textEntityTypePhoneNumber"
)

// InputMessageContentEnum Alias for abstract InputMessageContent 'Sub-Classes', used as constant-enum here
type InputMessageContentEnum string

// InputMessageContent enums
const (
	InputMessageTextType      InputMessageContentEnum = "inputMessageText"
	InputMessageAnimationType InputMessageContentEnum = "inputMessageAnimation"
	InputMessageAudioType     InputMessageContentEnum = "inputMessageAudio"
	InputMessageDocumentType  InputMessageContentEnum = "inputMessageDocument"
	InputMessagePhotoType     InputMessageContentEnum = "inputMessagePhoto"
	InputMessageStickerType   InputMessageContentEnum = "inputMessageSticker"
	InputMessageVideoType     InputMessageContentEnum = "inputMessageVideo"
	InputMessageVideoNoteType InputMessageContentEnum = "inputMessageVideoNote"
	InputMessageVoiceNoteType InputMessageContentEnum = "inputMessageVoiceNote"
	InputMessageLocationType  InputMessageContentEnum = "inputMessageLocation"
	InputMessageVenueType     InputMessageContentEnum = "inputMessageVenue"
	InputMessageContactType   InputMessageContentEnum = "inputMessageContact"
	InputMessageGameType      InputMessageContentEnum = "inputMessageGame"
	InputMessageInvoiceType   InputMessageContentEnum = "inputMessageInvoice"
	InputMessageForwardedType InputMessageContentEnum = "inputMessageForwarded"
)

// SearchMessagesFilterEnum Alias for abstract SearchMessagesFilter 'Sub-Classes', used as constant-enum here
type SearchMessagesFilterEnum string

// SearchMessagesFilter enums
const (
	SearchMessagesFilterEmptyType             SearchMessagesFilterEnum = "searchMessagesFilterEmpty"
	SearchMessagesFilterAnimationType         SearchMessagesFilterEnum = "searchMessagesFilterAnimation"
	SearchMessagesFilterAudioType             SearchMessagesFilterEnum = "searchMessagesFilterAudio"
	SearchMessagesFilterDocumentType          SearchMessagesFilterEnum = "searchMessagesFilterDocument"
	SearchMessagesFilterPhotoType             SearchMessagesFilterEnum = "searchMessagesFilterPhoto"
	SearchMessagesFilterVideoType             SearchMessagesFilterEnum = "searchMessagesFilterVideo"
	SearchMessagesFilterVoiceNoteType         SearchMessagesFilterEnum = "searchMessagesFilterVoiceNote"
	SearchMessagesFilterPhotoAndVideoType     SearchMessagesFilterEnum = "searchMessagesFilterPhotoAndVideo"
	SearchMessagesFilterURLType               SearchMessagesFilterEnum = "searchMessagesFilterURL"
	SearchMessagesFilterChatPhotoType         SearchMessagesFilterEnum = "searchMessagesFilterChatPhoto"
	SearchMessagesFilterCallType              SearchMessagesFilterEnum = "searchMessagesFilterCall"
	SearchMessagesFilterMissedCallType        SearchMessagesFilterEnum = "searchMessagesFilterMissedCall"
	SearchMessagesFilterVideoNoteType         SearchMessagesFilterEnum = "searchMessagesFilterVideoNote"
	SearchMessagesFilterVoiceAndVideoNoteType SearchMessagesFilterEnum = "searchMessagesFilterVoiceAndVideoNote"
	SearchMessagesFilterMentionType           SearchMessagesFilterEnum = "searchMessagesFilterMention"
	SearchMessagesFilterUnreadMentionType     SearchMessagesFilterEnum = "searchMessagesFilterUnreadMention"
)

// ChatActionEnum Alias for abstract ChatAction 'Sub-Classes', used as constant-enum here
type ChatActionEnum string

// ChatAction enums
const (
	ChatActionTypingType             ChatActionEnum = "chatActionTyping"
	ChatActionRecordingVideoType     ChatActionEnum = "chatActionRecordingVideo"
	ChatActionUploadingVideoType     ChatActionEnum = "chatActionUploadingVideo"
	ChatActionRecordingVoiceNoteType ChatActionEnum = "chatActionRecordingVoiceNote"
	ChatActionUploadingVoiceNoteType ChatActionEnum = "chatActionUploadingVoiceNote"
	ChatActionUploadingPhotoType     ChatActionEnum = "chatActionUploadingPhoto"
	ChatActionUploadingDocumentType  ChatActionEnum = "chatActionUploadingDocument"
	ChatActionChoosingLocationType   ChatActionEnum = "chatActionChoosingLocation"
	ChatActionChoosingContactType    ChatActionEnum = "chatActionChoosingContact"
	ChatActionStartPlayingGameType   ChatActionEnum = "chatActionStartPlayingGame"
	ChatActionRecordingVideoNoteType ChatActionEnum = "chatActionRecordingVideoNote"
	ChatActionUploadingVideoNoteType ChatActionEnum = "chatActionUploadingVideoNote"
	ChatActionCancelType             ChatActionEnum = "chatActionCancel"
)

// UserStatusEnum Alias for abstract UserStatus 'Sub-Classes', used as constant-enum here
type UserStatusEnum string

// UserStatus enums
const (
	UserStatusEmptyType     UserStatusEnum = "userStatusEmpty"
	UserStatusOnlineType    UserStatusEnum = "userStatusOnline"
	UserStatusOfflineType   UserStatusEnum = "userStatusOffline"
	UserStatusRecentlyType  UserStatusEnum = "userStatusRecently"
	UserStatusLastWeekType  UserStatusEnum = "userStatusLastWeek"
	UserStatusLastMonthType UserStatusEnum = "userStatusLastMonth"
)

// CallDiscardReasonEnum Alias for abstract CallDiscardReason 'Sub-Classes', used as constant-enum here
type CallDiscardReasonEnum string

// CallDiscardReason enums
const (
	CallDiscardReasonEmptyType        CallDiscardReasonEnum = "callDiscardReasonEmpty"
	CallDiscardReasonMissedType       CallDiscardReasonEnum = "callDiscardReasonMissed"
	CallDiscardReasonDeclinedType     CallDiscardReasonEnum = "callDiscardReasonDeclined"
	CallDiscardReasonDisconnectedType CallDiscardReasonEnum = "callDiscardReasonDisconnected"
	CallDiscardReasonHungUpType       CallDiscardReasonEnum = "callDiscardReasonHungUp"
)

// CallStateEnum Alias for abstract CallState 'Sub-Classes', used as constant-enum here
type CallStateEnum string

// CallState enums
const (
	CallStatePendingType        CallStateEnum = "callStatePending"
	CallStateExchangingKeysType CallStateEnum = "callStateExchangingKeys"
	CallStateReadyType          CallStateEnum = "callStateReady"
	CallStateHangingUpType      CallStateEnum = "callStateHangingUp"
	CallStateDiscardedType      CallStateEnum = "callStateDiscarded"
	CallStateErrorType          CallStateEnum = "callStateError"
)

// InputInlineQueryResultEnum Alias for abstract InputInlineQueryResult 'Sub-Classes', used as constant-enum here
type InputInlineQueryResultEnum string

// InputInlineQueryResult enums
const (
	InputInlineQueryResultAnimatedGifType   InputInlineQueryResultEnum = "inputInlineQueryResultAnimatedGif"
	InputInlineQueryResultAnimatedMpeg4Type InputInlineQueryResultEnum = "inputInlineQueryResultAnimatedMpeg4"
	InputInlineQueryResultArticleType       InputInlineQueryResultEnum = "inputInlineQueryResultArticle"
	InputInlineQueryResultAudioType         InputInlineQueryResultEnum = "inputInlineQueryResultAudio"
	InputInlineQueryResultContactType       InputInlineQueryResultEnum = "inputInlineQueryResultContact"
	InputInlineQueryResultDocumentType      InputInlineQueryResultEnum = "inputInlineQueryResultDocument"
	InputInlineQueryResultGameType          InputInlineQueryResultEnum = "inputInlineQueryResultGame"
	InputInlineQueryResultLocationType      InputInlineQueryResultEnum = "inputInlineQueryResultLocation"
	InputInlineQueryResultPhotoType         InputInlineQueryResultEnum = "inputInlineQueryResultPhoto"
	InputInlineQueryResultStickerType       InputInlineQueryResultEnum = "inputInlineQueryResultSticker"
	InputInlineQueryResultVenueType         InputInlineQueryResultEnum = "inputInlineQueryResultVenue"
	InputInlineQueryResultVideoType         InputInlineQueryResultEnum = "inputInlineQueryResultVideo"
	InputInlineQueryResultVoiceNoteType     InputInlineQueryResultEnum = "inputInlineQueryResultVoiceNote"
)

// InlineQueryResultEnum Alias for abstract InlineQueryResult 'Sub-Classes', used as constant-enum here
type InlineQueryResultEnum string

// InlineQueryResult enums
const (
	InlineQueryResultArticleType   InlineQueryResultEnum = "inlineQueryResultArticle"
	InlineQueryResultContactType   InlineQueryResultEnum = "inlineQueryResultContact"
	InlineQueryResultLocationType  InlineQueryResultEnum = "inlineQueryResultLocation"
	InlineQueryResultVenueType     InlineQueryResultEnum = "inlineQueryResultVenue"
	InlineQueryResultGameType      InlineQueryResultEnum = "inlineQueryResultGame"
	InlineQueryResultAnimationType InlineQueryResultEnum = "inlineQueryResultAnimation"
	InlineQueryResultAudioType     InlineQueryResultEnum = "inlineQueryResultAudio"
	InlineQueryResultDocumentType  InlineQueryResultEnum = "inlineQueryResultDocument"
	InlineQueryResultPhotoType     InlineQueryResultEnum = "inlineQueryResultPhoto"
	InlineQueryResultStickerType   InlineQueryResultEnum = "inlineQueryResultSticker"
	InlineQueryResultVideoType     InlineQueryResultEnum = "inlineQueryResultVideo"
	InlineQueryResultVoiceNoteType InlineQueryResultEnum = "inlineQueryResultVoiceNote"
)

// CallbackQueryPayloadEnum Alias for abstract CallbackQueryPayload 'Sub-Classes', used as constant-enum here
type CallbackQueryPayloadEnum string

// CallbackQueryPayload enums
const (
	CallbackQueryPayloadDataType CallbackQueryPayloadEnum = "callbackQueryPayloadData"
	CallbackQueryPayloadGameType CallbackQueryPayloadEnum = "callbackQueryPayloadGame"
)

// ChatEventActionEnum Alias for abstract ChatEventAction 'Sub-Classes', used as constant-enum here
type ChatEventActionEnum string

// ChatEventAction enums
const (
	ChatEventMessageEditedType                ChatEventActionEnum = "chatEventMessageEdited"
	ChatEventMessageDeletedType               ChatEventActionEnum = "chatEventMessageDeleted"
	ChatEventMessagePinnedType                ChatEventActionEnum = "chatEventMessagePinned"
	ChatEventMessageUnpinnedType              ChatEventActionEnum = "chatEventMessageUnpinned"
	ChatEventMemberJoinedType                 ChatEventActionEnum = "chatEventMemberJoined"
	ChatEventMemberLeftType                   ChatEventActionEnum = "chatEventMemberLeft"
	ChatEventMemberInvitedType                ChatEventActionEnum = "chatEventMemberInvited"
	ChatEventMemberPromotedType               ChatEventActionEnum = "chatEventMemberPromoted"
	ChatEventMemberRestrictedType             ChatEventActionEnum = "chatEventMemberRestricted"
	ChatEventTitleChangedType                 ChatEventActionEnum = "chatEventTitleChanged"
	ChatEventDescriptionChangedType           ChatEventActionEnum = "chatEventDescriptionChanged"
	ChatEventUsernameChangedType              ChatEventActionEnum = "chatEventUsernameChanged"
	ChatEventPhotoChangedType                 ChatEventActionEnum = "chatEventPhotoChanged"
	ChatEventInvitesToggledType               ChatEventActionEnum = "chatEventInvitesToggled"
	ChatEventSignMessagesToggledType          ChatEventActionEnum = "chatEventSignMessagesToggled"
	ChatEventStickerSetChangedType            ChatEventActionEnum = "chatEventStickerSetChanged"
	ChatEventIsAllHistoryAvailableToggledType ChatEventActionEnum = "chatEventIsAllHistoryAvailableToggled"
)

// DeviceTokenEnum Alias for abstract DeviceToken 'Sub-Classes', used as constant-enum here
type DeviceTokenEnum string

// DeviceToken enums
const (
	DeviceTokenGoogleCloudMessagingType DeviceTokenEnum = "deviceTokenGoogleCloudMessaging"
	DeviceTokenApplePushType            DeviceTokenEnum = "deviceTokenApplePush"
	DeviceTokenApplePushVoIPType        DeviceTokenEnum = "deviceTokenApplePushVoIP"
	DeviceTokenWindowsPushType          DeviceTokenEnum = "deviceTokenWindowsPush"
	DeviceTokenMicrosoftPushType        DeviceTokenEnum = "deviceTokenMicrosoftPush"
	DeviceTokenMicrosoftPushVoIPType    DeviceTokenEnum = "deviceTokenMicrosoftPushVoIP"
	DeviceTokenWebPushType              DeviceTokenEnum = "deviceTokenWebPush"
	DeviceTokenSimplePushType           DeviceTokenEnum = "deviceTokenSimplePush"
	DeviceTokenUbuntuPushType           DeviceTokenEnum = "deviceTokenUbuntuPush"
	DeviceTokenBlackBerryPushType       DeviceTokenEnum = "deviceTokenBlackBerryPush"
	DeviceTokenTizenPushType            DeviceTokenEnum = "deviceTokenTizenPush"
)

// CheckChatUsernameResultEnum Alias for abstract CheckChatUsernameResult 'Sub-Classes', used as constant-enum here
type CheckChatUsernameResultEnum string

// CheckChatUsernameResult enums
const (
	CheckChatUsernameResultOkType                      CheckChatUsernameResultEnum = "checkChatUsernameResultOk"
	CheckChatUsernameResultUsernameInvalidType         CheckChatUsernameResultEnum = "checkChatUsernameResultUsernameInvalid"
	CheckChatUsernameResultUsernameOccupiedType        CheckChatUsernameResultEnum = "checkChatUsernameResultUsernameOccupied"
	CheckChatUsernameResultPublicChatsTooMuchType      CheckChatUsernameResultEnum = "checkChatUsernameResultPublicChatsTooMuch"
	CheckChatUsernameResultPublicGroupsUnavailableType CheckChatUsernameResultEnum = "checkChatUsernameResultPublicGroupsUnavailable"
)

// OptionValueEnum Alias for abstract OptionValue 'Sub-Classes', used as constant-enum here
type OptionValueEnum string

// OptionValue enums
const (
	OptionValueBooleanType OptionValueEnum = "optionValueBoolean"
	OptionValueEmptyType   OptionValueEnum = "optionValueEmpty"
	OptionValueIntegerType OptionValueEnum = "optionValueInteger"
	OptionValueStringType  OptionValueEnum = "optionValueString"
)

// UserPrivacySettingRuleEnum Alias for abstract UserPrivacySettingRule 'Sub-Classes', used as constant-enum here
type UserPrivacySettingRuleEnum string

// UserPrivacySettingRule enums
const (
	UserPrivacySettingRuleAllowAllType         UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowAll"
	UserPrivacySettingRuleAllowContactsType    UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowContacts"
	UserPrivacySettingRuleAllowUsersType       UserPrivacySettingRuleEnum = "userPrivacySettingRuleAllowUsers"
	UserPrivacySettingRuleRestrictAllType      UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictAll"
	UserPrivacySettingRuleRestrictContactsType UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictContacts"
	UserPrivacySettingRuleRestrictUsersType    UserPrivacySettingRuleEnum = "userPrivacySettingRuleRestrictUsers"
)

// UserPrivacySettingEnum Alias for abstract UserPrivacySetting 'Sub-Classes', used as constant-enum here
type UserPrivacySettingEnum string

// UserPrivacySetting enums
const (
	UserPrivacySettingShowStatusType       UserPrivacySettingEnum = "userPrivacySettingShowStatus"
	UserPrivacySettingAllowChatInvitesType UserPrivacySettingEnum = "userPrivacySettingAllowChatInvites"
	UserPrivacySettingAllowCallsType       UserPrivacySettingEnum = "userPrivacySettingAllowCalls"
)

// ChatReportReasonEnum Alias for abstract ChatReportReason 'Sub-Classes', used as constant-enum here
type ChatReportReasonEnum string

// ChatReportReason enums
const (
	ChatReportReasonSpamType        ChatReportReasonEnum = "chatReportReasonSpam"
	ChatReportReasonViolenceType    ChatReportReasonEnum = "chatReportReasonViolence"
	ChatReportReasonPornographyType ChatReportReasonEnum = "chatReportReasonPornography"
	ChatReportReasonCustomType      ChatReportReasonEnum = "chatReportReasonCustom"
)

// FileTypeEnum Alias for abstract FileType 'Sub-Classes', used as constant-enum here
type FileTypeEnum string

// FileType enums
const (
	FileTypeNoneType            FileTypeEnum = "fileTypeNone"
	FileTypeAnimationType       FileTypeEnum = "fileTypeAnimation"
	FileTypeAudioType           FileTypeEnum = "fileTypeAudio"
	FileTypeDocumentType        FileTypeEnum = "fileTypeDocument"
	FileTypePhotoType           FileTypeEnum = "fileTypePhoto"
	FileTypeProfilePhotoType    FileTypeEnum = "fileTypeProfilePhoto"
	FileTypeSecretType          FileTypeEnum = "fileTypeSecret"
	FileTypeStickerType         FileTypeEnum = "fileTypeSticker"
	FileTypeThumbnailType       FileTypeEnum = "fileTypeThumbnail"
	FileTypeUnknownType         FileTypeEnum = "fileTypeUnknown"
	FileTypeVideoType           FileTypeEnum = "fileTypeVideo"
	FileTypeVideoNoteType       FileTypeEnum = "fileTypeVideoNote"
	FileTypeVoiceNoteType       FileTypeEnum = "fileTypeVoiceNote"
	FileTypeWallpaperType       FileTypeEnum = "fileTypeWallpaper"
	FileTypeSecretThumbnailType FileTypeEnum = "fileTypeSecretThumbnail"
)

// NetworkTypeEnum Alias for abstract NetworkType 'Sub-Classes', used as constant-enum here
type NetworkTypeEnum string

// NetworkType enums
const (
	NetworkTypeNoneType          NetworkTypeEnum = "networkTypeNone"
	NetworkTypeMobileType        NetworkTypeEnum = "networkTypeMobile"
	NetworkTypeMobileRoamingType NetworkTypeEnum = "networkTypeMobileRoaming"
	NetworkTypeWiFiType          NetworkTypeEnum = "networkTypeWiFi"
	NetworkTypeOtherType         NetworkTypeEnum = "networkTypeOther"
)

// NetworkStatisticsEntryEnum Alias for abstract NetworkStatisticsEntry 'Sub-Classes', used as constant-enum here
type NetworkStatisticsEntryEnum string

// NetworkStatisticsEntry enums
const (
	NetworkStatisticsEntryFileType NetworkStatisticsEntryEnum = "networkStatisticsEntryFile"
	NetworkStatisticsEntryCallType NetworkStatisticsEntryEnum = "networkStatisticsEntryCall"
)

// ConnectionStateEnum Alias for abstract ConnectionState 'Sub-Classes', used as constant-enum here
type ConnectionStateEnum string

// ConnectionState enums
const (
	ConnectionStateWaitingForNetworkType ConnectionStateEnum = "connectionStateWaitingForNetwork"
	ConnectionStateConnectingToProxyType ConnectionStateEnum = "connectionStateConnectingToProxy"
	ConnectionStateConnectingType        ConnectionStateEnum = "connectionStateConnecting"
	ConnectionStateUpdatingType          ConnectionStateEnum = "connectionStateUpdating"
	ConnectionStateReadyType             ConnectionStateEnum = "connectionStateReady"
)

// TopChatCategoryEnum Alias for abstract TopChatCategory 'Sub-Classes', used as constant-enum here
type TopChatCategoryEnum string

// TopChatCategory enums
const (
	TopChatCategoryUsersType      TopChatCategoryEnum = "topChatCategoryUsers"
	TopChatCategoryBotsType       TopChatCategoryEnum = "topChatCategoryBots"
	TopChatCategoryGroupsType     TopChatCategoryEnum = "topChatCategoryGroups"
	TopChatCategoryChannelsType   TopChatCategoryEnum = "topChatCategoryChannels"
	TopChatCategoryInlineBotsType TopChatCategoryEnum = "topChatCategoryInlineBots"
	TopChatCategoryCallsType      TopChatCategoryEnum = "topChatCategoryCalls"
)

// TMeURLTypeEnum Alias for abstract TMeURLType 'Sub-Classes', used as constant-enum here
type TMeURLTypeEnum string

// TMeURLType enums
const (
	TMeURLTypeUserType       TMeURLTypeEnum = "tMeURLTypeUser"
	TMeURLTypeSupergroupType TMeURLTypeEnum = "tMeURLTypeSupergroup"
	TMeURLTypeChatInviteType TMeURLTypeEnum = "tMeURLTypeChatInvite"
	TMeURLTypeStickerSetType TMeURLTypeEnum = "tMeURLTypeStickerSet"
)

// TextParseModeEnum Alias for abstract TextParseMode 'Sub-Classes', used as constant-enum here
type TextParseModeEnum string

// TextParseMode enums
const (
	TextParseModeMarkdownType TextParseModeEnum = "textParseModeMarkdown"
	TextParseModeHTMLType     TextParseModeEnum = "textParseModeHTML"
)

// ProxyEnum Alias for abstract Proxy 'Sub-Classes', used as constant-enum here
type ProxyEnum string

// Proxy enums
const (
	ProxyEmptyType  ProxyEnum = "proxyEmpty"
	ProxySocks5Type ProxyEnum = "proxySocks5"
)

// UpdateEnum Alias for abstract Update 'Sub-Classes', used as constant-enum here
type UpdateEnum string

// Update enums
const (
	UpdateAuthorizationStateType      UpdateEnum = "updateAuthorizationState"
	UpdateNewMessageType              UpdateEnum = "updateNewMessage"
	UpdateMessageSendAcknowledgedType UpdateEnum = "updateMessageSendAcknowledged"
	UpdateMessageSendSucceededType    UpdateEnum = "updateMessageSendSucceeded"
	UpdateMessageSendFailedType       UpdateEnum = "updateMessageSendFailed"
	UpdateMessageContentType          UpdateEnum = "updateMessageContent"
	UpdateMessageEditedType           UpdateEnum = "updateMessageEdited"
	UpdateMessageViewsType            UpdateEnum = "updateMessageViews"
	UpdateMessageContentOpenedType    UpdateEnum = "updateMessageContentOpened"
	UpdateMessageMentionReadType      UpdateEnum = "updateMessageMentionRead"
	UpdateNewChatType                 UpdateEnum = "updateNewChat"
	UpdateChatTitleType               UpdateEnum = "updateChatTitle"
	UpdateChatPhotoType               UpdateEnum = "updateChatPhoto"
	UpdateChatLastMessageType         UpdateEnum = "updateChatLastMessage"
	UpdateChatOrderType               UpdateEnum = "updateChatOrder"
	UpdateChatIsPinnedType            UpdateEnum = "updateChatIsPinned"
	UpdateChatReadInboxType           UpdateEnum = "updateChatReadInbox"
	UpdateChatReadOutboxType          UpdateEnum = "updateChatReadOutbox"
	UpdateChatUnreadMentionCountType  UpdateEnum = "updateChatUnreadMentionCount"
	UpdateNotificationSettingsType    UpdateEnum = "updateNotificationSettings"
	UpdateChatReplyMarkupType         UpdateEnum = "updateChatReplyMarkup"
	UpdateChatDraftMessageType        UpdateEnum = "updateChatDraftMessage"
	UpdateDeleteMessagesType          UpdateEnum = "updateDeleteMessages"
	UpdateUserChatActionType          UpdateEnum = "updateUserChatAction"
	UpdateUserStatusType              UpdateEnum = "updateUserStatus"
	UpdateUserType                    UpdateEnum = "updateUser"
	UpdateBasicGroupType              UpdateEnum = "updateBasicGroup"
	UpdateSupergroupType              UpdateEnum = "updateSupergroup"
	UpdateSecretChatType              UpdateEnum = "updateSecretChat"
	UpdateUserFullInfoType            UpdateEnum = "updateUserFullInfo"
	UpdateBasicGroupFullInfoType      UpdateEnum = "updateBasicGroupFullInfo"
	UpdateSupergroupFullInfoType      UpdateEnum = "updateSupergroupFullInfo"
	UpdateServiceNotificationType     UpdateEnum = "updateServiceNotification"
	UpdateFileType                    UpdateEnum = "updateFile"
	UpdateFileGenerationStartType     UpdateEnum = "updateFileGenerationStart"
	UpdateFileGenerationStopType      UpdateEnum = "updateFileGenerationStop"
	UpdateCallType                    UpdateEnum = "updateCall"
	UpdateUserPrivacySettingRulesType UpdateEnum = "updateUserPrivacySettingRules"
	UpdateUnreadMessageCountType      UpdateEnum = "updateUnreadMessageCount"
	UpdateOptionType                  UpdateEnum = "updateOption"
	UpdateInstalledStickerSetsType    UpdateEnum = "updateInstalledStickerSets"
	UpdateTrendingStickerSetsType     UpdateEnum = "updateTrendingStickerSets"
	UpdateRecentStickersType          UpdateEnum = "updateRecentStickers"
	UpdateFavoriteStickersType        UpdateEnum = "updateFavoriteStickers"
	UpdateSavedAnimationsType         UpdateEnum = "updateSavedAnimations"
	UpdateConnectionStateType         UpdateEnum = "updateConnectionState"
	UpdateNewInlineQueryType          UpdateEnum = "updateNewInlineQuery"
	UpdateNewChosenInlineResultType   UpdateEnum = "updateNewChosenInlineResult"
	UpdateNewCallbackQueryType        UpdateEnum = "updateNewCallbackQuery"
	UpdateNewInlineCallbackQueryType  UpdateEnum = "updateNewInlineCallbackQuery"
	UpdateNewShippingQueryType        UpdateEnum = "updateNewShippingQuery"
	UpdateNewPreCheckoutQueryType     UpdateEnum = "updateNewPreCheckoutQuery"
	UpdateNewCustomEventType          UpdateEnum = "updateNewCustomEvent"
	UpdateNewCustomQueryType          UpdateEnum = "updateNewCustomQuery"
) // AuthenticationCodeType Provides information about the method by which an authentication code is delivered to the user
type AuthenticationCodeType interface {
	GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum
}

// AuthorizationState Represents the current authorization state of the client
type AuthorizationState interface {
	GetAuthorizationStateEnum() AuthorizationStateEnum
}

// InputFile Points to a file
type InputFile interface {
	GetInputFileEnum() InputFileEnum
}

// MaskPoint Part of the face, relative to which a mask should be placed
type MaskPoint interface {
	GetMaskPointEnum() MaskPointEnum
}

// LinkState Represents the relationship between user A and user B. For incoming_link, user A is the current user; for outgoing_link, user B is the current user
type LinkState interface {
	GetLinkStateEnum() LinkStateEnum
}

// UserType Represents the type of the user. The following types are possible: regular users, deleted users and bots
type UserType interface {
	GetUserTypeEnum() UserTypeEnum
}

// ChatMemberStatus Provides information about the status of a member in a chat
type ChatMemberStatus interface {
	GetChatMemberStatusEnum() ChatMemberStatusEnum
}

// SupergroupMembersFilter Specifies the kind of chat members to return in getSupergroupMembers
type SupergroupMembersFilter interface {
	GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum
}

// SecretChatState Describes the current secret chat state
type SecretChatState interface {
	GetSecretChatStateEnum() SecretChatStateEnum
}

// MessageForwardInfo Contains information about the initial sender of a forwarded message
type MessageForwardInfo interface {
	GetMessageForwardInfoEnum() MessageForwardInfoEnum
}

// MessageSendingState Contains information about the sending state of the message
type MessageSendingState interface {
	GetMessageSendingStateEnum() MessageSendingStateEnum
}

// NotificationSettingsScope Describes the types of chats for which notification settings are applied
type NotificationSettingsScope interface {
	GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum
}

// ChatType Describes the type of a chat
type ChatType interface {
	GetChatTypeEnum() ChatTypeEnum
}

// KeyboardButtonType Describes a keyboard button type
type KeyboardButtonType interface {
	GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum
}

// InlineKeyboardButtonType Describes the type of an inline keyboard button
type InlineKeyboardButtonType interface {
	GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum
}

// ReplyMarkup Contains a description of a custom keyboard and actions that can be done with it to quickly reply to bots
type ReplyMarkup interface {
	GetReplyMarkupEnum() ReplyMarkupEnum
}

// RichText Describes a text object inside an instant-view web page
type RichText interface {
	GetRichTextEnum() RichTextEnum
}

// PageBlock Describes a block of an instant view web page
type PageBlock interface {
	GetPageBlockEnum() PageBlockEnum
}

// InputCredentials Contains information about the payment method chosen by the user
type InputCredentials interface {
	GetInputCredentialsEnum() InputCredentialsEnum
}

// MessageContent Contains the content of a message
type MessageContent interface {
	GetMessageContentEnum() MessageContentEnum
}

// TextEntityType Represents a part of the text which must be formatted differently
type TextEntityType interface {
	GetTextEntityTypeEnum() TextEntityTypeEnum
}

// InputMessageContent The content of a message to send
type InputMessageContent interface {
	GetInputMessageContentEnum() InputMessageContentEnum
}

// SearchMessagesFilter Represents a filter for message search results
type SearchMessagesFilter interface {
	GetSearchMessagesFilterEnum() SearchMessagesFilterEnum
}

// ChatAction Describes the different types of activity in a chat
type ChatAction interface {
	GetChatActionEnum() ChatActionEnum
}

// UserStatus Describes the last time the user was online
type UserStatus interface {
	GetUserStatusEnum() UserStatusEnum
}

// CallDiscardReason Describes the reason why a call was discarded
type CallDiscardReason interface {
	GetCallDiscardReasonEnum() CallDiscardReasonEnum
}

// CallState Describes the current call state
type CallState interface {
	GetCallStateEnum() CallStateEnum
}

// InputInlineQueryResult Represents a single result of an inline query; for bots only
type InputInlineQueryResult interface {
	GetInputInlineQueryResultEnum() InputInlineQueryResultEnum
}

// InlineQueryResult Represents a single result of an inline query
type InlineQueryResult interface {
	GetInlineQueryResultEnum() InlineQueryResultEnum
}

// CallbackQueryPayload Represents a payload of a callback query
type CallbackQueryPayload interface {
	GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum
}

// ChatEventAction Represents a chat event
type ChatEventAction interface {
	GetChatEventActionEnum() ChatEventActionEnum
}

// DeviceToken Represents a data needed to subscribe for push notifications. To use specific push notification service, you must specify the correct application platform and upload valid server authentication data at https://my.telegram.org
type DeviceToken interface {
	GetDeviceTokenEnum() DeviceTokenEnum
}

// CheckChatUsernameResult Represents result of checking whether a username can be set for a chat
type CheckChatUsernameResult interface {
	GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum
}

// OptionValue Represents the value of an option
type OptionValue interface {
	GetOptionValueEnum() OptionValueEnum
}

// UserPrivacySettingRule Represents a single rule for managing privacy settings
type UserPrivacySettingRule interface {
	GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum
}

// UserPrivacySetting Describes available user privacy settings
type UserPrivacySetting interface {
	GetUserPrivacySettingEnum() UserPrivacySettingEnum
}

// ChatReportReason Describes the reason why a chat is reported
type ChatReportReason interface {
	GetChatReportReasonEnum() ChatReportReasonEnum
}

// FileType Represents the type of a file
type FileType interface {
	GetFileTypeEnum() FileTypeEnum
}

// NetworkType Represents the type of a network
type NetworkType interface {
	GetNetworkTypeEnum() NetworkTypeEnum
}

// NetworkStatisticsEntry Contains statistics about network usage
type NetworkStatisticsEntry interface {
	GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum
}

// ConnectionState Describes the current state of the connection to Telegram servers
type ConnectionState interface {
	GetConnectionStateEnum() ConnectionStateEnum
}

// TopChatCategory Represents the categories of chats for which a list of frequently used chats can be retrieved
type TopChatCategory interface {
	GetTopChatCategoryEnum() TopChatCategoryEnum
}

// TMeURLType Describes the type of a URL linking to an internal Telegram entity
type TMeURLType interface {
	GetTMeURLTypeEnum() TMeURLTypeEnum
}

// TextParseMode Describes the way the text should be parsed for TextEntities
type TextParseMode interface {
	GetTextParseModeEnum() TextParseModeEnum
}

// Proxy Contains information about a proxy server
type Proxy interface {
	GetProxyEnum() ProxyEnum
}

// Update Contains notifications about data changes
type Update interface {
	GetUpdateEnum() UpdateEnum
}

// Error An object of this type can be returned on every function call, in case of an error
type Error struct {
	tdCommon
	Code    int32  `json:"code"`    // Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
	Message string `json:"message"` // Error message; subject to future changes
}

// MessageType return the string telegram-type of Error
func (error *Error) MessageType() string {
	return "error"
}

// Ok An object of this type is returned on a successful function call for certain functions
type Ok struct {
	tdCommon
}

// MessageType return the string telegram-type of Ok
func (ok *Ok) MessageType() string {
	return "ok"
}

// TdlibParameters Contains parameters for TDLib initialization
type TdlibParameters struct {
	tdCommon
	UseTestDc              bool   `json:"use_test_dc"`              // If set to true, the Telegram test environment will be used instead of the production environment
	DatabaseDirectory      string `json:"database_directory"`       // The path to the directory for the persistent database; if empty, the current working directory will be used
	FilesDirectory         string `json:"files_directory"`          // The path to the directory for storing files; if empty, database_directory will be used
	UseFileDatabase        bool   `json:"use_file_database"`        // If set to true, information about downloaded and uploaded files will be saved between application restarts
	UseChatInfoDatabase    bool   `json:"use_chat_info_database"`   // If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database
	UseMessageDatabase     bool   `json:"use_message_database"`     // If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database
	UseSecretChats         bool   `json:"use_secret_chats"`         // If set to true, support for secret chats will be enabled
	APIID                  int32  `json:"api_id"`                   // Application identifier for Telegram API access, which can be obtained at https://my.telegram.org
	APIHash                string `json:"api_hash"`                 // Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org
	SystemLanguageCode     string `json:"system_language_code"`     // IETF language tag of the user's operating system language; must be non-empty
	DeviceModel            string `json:"device_model"`             // Model of the device the application is being run on; must be non-empty
	SystemVersion          string `json:"system_version"`           // Version of the operating system the application is being run on; must be non-empty
	ApplicationVersion     string `json:"application_version"`      // Application version; must be non-empty
	EnableStorageOptimizer bool   `json:"enable_storage_optimizer"` // If set to true, old files will automatically be deleted
	IgnoreFileNames        bool   `json:"ignore_file_names"`        // If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name
}

// MessageType return the string telegram-type of TdlibParameters
func (tdlibParameters *TdlibParameters) MessageType() string {
	return "tdlibParameters"
}

// AuthenticationCodeTypeTelegramMessage An authentication code is delivered via a private Telegram message, which can be viewed in another client
type AuthenticationCodeTypeTelegramMessage struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeTelegramMessage
func (authenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage) MessageType() string {
	return "authenticationCodeTypeTelegramMessage"
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeTelegramMessageType
}

// AuthenticationCodeTypeSms An authentication code is delivered via an SMS message to the specified phone number
type AuthenticationCodeTypeSms struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeSms
func (authenticationCodeTypeSms *AuthenticationCodeTypeSms) MessageType() string {
	return "authenticationCodeTypeSms"
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeSms *AuthenticationCodeTypeSms) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeSmsType
}

// AuthenticationCodeTypeCall An authentication code is delivered via a phone call to the specified phone number
type AuthenticationCodeTypeCall struct {
	tdCommon
	Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeCall
func (authenticationCodeTypeCall *AuthenticationCodeTypeCall) MessageType() string {
	return "authenticationCodeTypeCall"
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeCall *AuthenticationCodeTypeCall) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeCallType
}

// AuthenticationCodeTypeFlashCall An authentication code is delivered by an immediately cancelled call to the specified phone number. The number from which the call was made is the code
type AuthenticationCodeTypeFlashCall struct {
	tdCommon
	Pattern string `json:"pattern"` // Pattern of the phone number from which the call will be made
}

// MessageType return the string telegram-type of AuthenticationCodeTypeFlashCall
func (authenticationCodeTypeFlashCall *AuthenticationCodeTypeFlashCall) MessageType() string {
	return "authenticationCodeTypeFlashCall"
}

// GetAuthenticationCodeTypeEnum return the enum type of this object
func (authenticationCodeTypeFlashCall *AuthenticationCodeTypeFlashCall) GetAuthenticationCodeTypeEnum() AuthenticationCodeTypeEnum {
	return AuthenticationCodeTypeFlashCallType
}

// AuthenticationCodeInfo Information about the authentication code that was sent
type AuthenticationCodeInfo struct {
	tdCommon
	PhoneNumber string                 `json:"phone_number"` // A phone number that is being authenticated
	Type        AuthenticationCodeType `json:"type"`         // Describes the way the code was sent to the user
	NextType    AuthenticationCodeType `json:"next_type"`    // Describes the way the next code will be sent to the user; may be null
	Timeout     int32                  `json:"timeout"`      // Timeout before the code should be re-sent, in seconds
}

// MessageType return the string telegram-type of AuthenticationCodeInfo
func (authenticationCodeInfo *AuthenticationCodeInfo) MessageType() string {
	return "authenticationCodeInfo"
}

// UnmarshalJSON unmarshal to json
func (authenticationCodeInfo *AuthenticationCodeInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PhoneNumber string `json:"phone_number"` // A phone number that is being authenticated
		Timeout     int32  `json:"timeout"`      // Timeout before the code should be re-sent, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	authenticationCodeInfo.tdCommon = tempObj.tdCommon
	authenticationCodeInfo.PhoneNumber = tempObj.PhoneNumber
	authenticationCodeInfo.Timeout = tempObj.Timeout

	fieldType, _ := unmarshalAuthenticationCodeType(objMap["type"])
	authenticationCodeInfo.Type = fieldType

	fieldNextType, _ := unmarshalAuthenticationCodeType(objMap["next_type"])
	authenticationCodeInfo.NextType = fieldNextType

	return nil
}

// AuthorizationStateWaitTdlibParameters TDLib needs TdlibParameters for initialization
type AuthorizationStateWaitTdlibParameters struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateWaitTdlibParameters
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) MessageType() string {
	return "authorizationStateWaitTdlibParameters"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitTdlibParametersType
}

// AuthorizationStateWaitEncryptionKey TDLib needs an encryption key to decrypt the local database
type AuthorizationStateWaitEncryptionKey struct {
	tdCommon
	IsEncrypted bool `json:"is_encrypted"` // True, if the database is currently encrypted
}

// MessageType return the string telegram-type of AuthorizationStateWaitEncryptionKey
func (authorizationStateWaitEncryptionKey *AuthorizationStateWaitEncryptionKey) MessageType() string {
	return "authorizationStateWaitEncryptionKey"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitEncryptionKey *AuthorizationStateWaitEncryptionKey) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitEncryptionKeyType
}

// AuthorizationStateWaitPhoneNumber TDLib needs the user's phone number to authorize
type AuthorizationStateWaitPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateWaitPhoneNumber
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) MessageType() string {
	return "authorizationStateWaitPhoneNumber"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitPhoneNumberType
}

// AuthorizationStateWaitCode TDLib needs the user's authentication code to finalize authorization
type AuthorizationStateWaitCode struct {
	tdCommon
	IsRegistered bool                   `json:"is_registered"` // True, if the user is already registered
	CodeInfo     AuthenticationCodeInfo `json:"code_info"`     // Information about the authorization code that was sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitCode
func (authorizationStateWaitCode *AuthorizationStateWaitCode) MessageType() string {
	return "authorizationStateWaitCode"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitCode *AuthorizationStateWaitCode) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitCodeType
}

// AuthorizationStateWaitPassword The user has been authorized, but needs to enter a password to start using the application
type AuthorizationStateWaitPassword struct {
	tdCommon
	PasswordHint                string `json:"password_hint"`                  // Hint for the password; can be empty
	HasRecoveryEmailAddress     bool   `json:"has_recovery_email_address"`     // True if a recovery email address has been set up
	RecoveryEmailAddressPattern string `json:"recovery_email_address_pattern"` // Pattern of the email address to which the recovery email was sent; empty until a recovery email has been sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitPassword
func (authorizationStateWaitPassword *AuthorizationStateWaitPassword) MessageType() string {
	return "authorizationStateWaitPassword"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitPassword *AuthorizationStateWaitPassword) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitPasswordType
}

// AuthorizationStateReady The user has been successfully authorized. TDLib is now ready to answer queries
type AuthorizationStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateReady
func (authorizationStateReady *AuthorizationStateReady) MessageType() string {
	return "authorizationStateReady"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateReady *AuthorizationStateReady) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateReadyType
}

// AuthorizationStateLoggingOut The user is currently logging out
type AuthorizationStateLoggingOut struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateLoggingOut
func (authorizationStateLoggingOut *AuthorizationStateLoggingOut) MessageType() string {
	return "authorizationStateLoggingOut"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateLoggingOut *AuthorizationStateLoggingOut) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateLoggingOutType
}

// AuthorizationStateClosing TDLib is closing, all subsequent queries will be answered with the error 500. Note that closing TDLib can take a while. All resources will be freed only after authorizationStateClosed has been received
type AuthorizationStateClosing struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateClosing
func (authorizationStateClosing *AuthorizationStateClosing) MessageType() string {
	return "authorizationStateClosing"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateClosing *AuthorizationStateClosing) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateClosingType
}

// AuthorizationStateClosed TDLib client is in its final state. All databases are closed and all resources are released. No other updates will be received after this. All queries will be responded to
type AuthorizationStateClosed struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateClosed
func (authorizationStateClosed *AuthorizationStateClosed) MessageType() string {
	return "authorizationStateClosed"
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateClosed *AuthorizationStateClosed) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateClosedType
}

// PasswordState Represents the current state of 2-step verification
type PasswordState struct {
	tdCommon
	HasPassword                            bool   `json:"has_password"`                               // True if a 2-step verification password has been set up
	PasswordHint                           string `json:"password_hint"`                              // Hint for the password; can be empty
	HasRecoveryEmailAddress                bool   `json:"has_recovery_email_address"`                 // True if a recovery email has been set up
	UnconfirmedRecoveryEmailAddressPattern string `json:"unconfirmed_recovery_email_address_pattern"` // Pattern of the email address to which a confirmation email was sent
}

// MessageType return the string telegram-type of PasswordState
func (passwordState *PasswordState) MessageType() string {
	return "passwordState"
}

// PasswordRecoveryInfo Contains information available to the user after requesting password recovery
type PasswordRecoveryInfo struct {
	tdCommon
	RecoveryEmailAddressPattern string `json:"recovery_email_address_pattern"` // Pattern of the email address to which a recovery email was sent
}

// MessageType return the string telegram-type of PasswordRecoveryInfo
func (passwordRecoveryInfo *PasswordRecoveryInfo) MessageType() string {
	return "passwordRecoveryInfo"
}

// RecoveryEmailAddress Contains information about the current recovery email address
type RecoveryEmailAddress struct {
	tdCommon
	RecoveryEmailAddress string `json:"recovery_email_address"` // Recovery email address
}

// MessageType return the string telegram-type of RecoveryEmailAddress
func (recoveryEmailAddress *RecoveryEmailAddress) MessageType() string {
	return "recoveryEmailAddress"
}

// TemporaryPasswordState Returns information about the availability of a temporary password, which can be used for payments
type TemporaryPasswordState struct {
	tdCommon
	HasPassword bool  `json:"has_password"` // True, if a temporary password is available
	ValidFor    int32 `json:"valid_for"`    // Time left before the temporary password expires, in seconds
}

// MessageType return the string telegram-type of TemporaryPasswordState
func (temporaryPasswordState *TemporaryPasswordState) MessageType() string {
	return "temporaryPasswordState"
}

// LocalFile Represents a local file
type LocalFile struct {
	tdCommon
	Path                   string `json:"path"`                     // Local path to the locally available file part; may be empty
	CanBeDownloaded        bool   `json:"can_be_downloaded"`        // True, if it is possible to try to download or generate the file
	CanBeDeleted           bool   `json:"can_be_deleted"`           // True, if the file can be deleted
	IsDownloadingActive    bool   `json:"is_downloading_active"`    // True, if the file is currently being downloaded (or a local copy is being generated by some other means)
	IsDownloadingCompleted bool   `json:"is_downloading_completed"` // True, if the local copy is fully available
	DownloadedPrefixSize   int32  `json:"downloaded_prefix_size"`   // If is_downloading_completed is false, then only some prefix of the file is ready to be read. downloaded_prefix_size is the size of that prefix
	DownloadedSize         int32  `json:"downloaded_size"`          // Total downloaded file bytes. Should be used only for calculating download progress. The actual file size may be bigger, and some parts of it may contain garbage
}

// MessageType return the string telegram-type of LocalFile
func (localFile *LocalFile) MessageType() string {
	return "localFile"
}

// RemoteFile Represents a remote file
type RemoteFile struct {
	tdCommon
	ID                   string `json:"id"`                     // Remote file identifier, may be empty. Can be used across application restarts or even from other devices for the current user. If the ID starts with "http://" or "https://", it represents the HTTP URL of the file. TDLib is currently unable to download files if only their URL is known.
	IsUploadingActive    bool   `json:"is_uploading_active"`    // True, if the file is currently being uploaded (or a remote copy is being generated by some other means)
	IsUploadingCompleted bool   `json:"is_uploading_completed"` // True, if a remote copy is fully available
	UploadedSize         int32  `json:"uploaded_size"`          // Size of the remote available part of the file; 0 if unknown
}

// MessageType return the string telegram-type of RemoteFile
func (remoteFile *RemoteFile) MessageType() string {
	return "remoteFile"
}

// File Represents a file
type File struct {
	tdCommon
	ID           int32      `json:"id"`            // Unique file identifier
	Size         int32      `json:"size"`          // File size; 0 if unknown
	ExpectedSize int32      `json:"expected_size"` // Expected file size in case the exact file size is unknown, but an approximate size is known. Can be used to show download/upload progress
	Local        LocalFile  `json:"local"`         // Information about the local copy of the file
	Remote       RemoteFile `json:"remote"`        // Information about the remote copy of the file
}

// MessageType return the string telegram-type of File
func (file *File) MessageType() string {
	return "file"
}

// InputFileID A file defined by its unique ID
type InputFileID struct {
	tdCommon
	ID int32 `json:"id"` // Unique file identifier
}

// MessageType return the string telegram-type of InputFileID
func (inputFileID *InputFileID) MessageType() string {
	return "inputFileId"
}

// GetInputFileEnum return the enum type of this object
func (inputFileID *InputFileID) GetInputFileEnum() InputFileEnum {
	return InputFileIDType
}

// InputFileRemote A file defined by its remote ID
type InputFileRemote struct {
	tdCommon
	ID string `json:"id"` // Remote file identifier
}

// MessageType return the string telegram-type of InputFileRemote
func (inputFileRemote *InputFileRemote) MessageType() string {
	return "inputFileRemote"
}

// GetInputFileEnum return the enum type of this object
func (inputFileRemote *InputFileRemote) GetInputFileEnum() InputFileEnum {
	return InputFileRemoteType
}

// InputFileLocal A file defined by a local path
type InputFileLocal struct {
	tdCommon
	Path string `json:"path"` // Local path to the file
}

// MessageType return the string telegram-type of InputFileLocal
func (inputFileLocal *InputFileLocal) MessageType() string {
	return "inputFileLocal"
}

// GetInputFileEnum return the enum type of this object
func (inputFileLocal *InputFileLocal) GetInputFileEnum() InputFileEnum {
	return InputFileLocalType
}

// InputFileGenerated A file generated by the client
type InputFileGenerated struct {
	tdCommon
	OriginalPath string `json:"original_path"` // Local path to a file from which the file is generated, may be empty if there is no such file
	Conversion   string `json:"conversion"`    // String specifying the conversion applied to the original file; should be persistent across application restarts
	ExpectedSize int32  `json:"expected_size"` // Expected size of the generated file; 0 if unknown
}

// MessageType return the string telegram-type of InputFileGenerated
func (inputFileGenerated *InputFileGenerated) MessageType() string {
	return "inputFileGenerated"
}

// GetInputFileEnum return the enum type of this object
func (inputFileGenerated *InputFileGenerated) GetInputFileEnum() InputFileEnum {
	return InputFileGeneratedType
}

// PhotoSize Photo description
type PhotoSize struct {
	tdCommon
	Type   string `json:"type"`   // Thumbnail type (see https://core.telegram.org/constructor/photoSize)
	Photo  File   `json:"photo"`  // Information about the photo file
	Width  int32  `json:"width"`  // Photo width
	Height int32  `json:"height"` // Photo height
}

// MessageType return the string telegram-type of PhotoSize
func (photoSize *PhotoSize) MessageType() string {
	return "photoSize"
}

// MaskPointForehead A mask should be placed relatively to the forehead
type MaskPointForehead struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointForehead
func (maskPointForehead *MaskPointForehead) MessageType() string {
	return "maskPointForehead"
}

// GetMaskPointEnum return the enum type of this object
func (maskPointForehead *MaskPointForehead) GetMaskPointEnum() MaskPointEnum {
	return MaskPointForeheadType
}

// MaskPointEyes A mask should be placed relatively to the eyes
type MaskPointEyes struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointEyes
func (maskPointEyes *MaskPointEyes) MessageType() string {
	return "maskPointEyes"
}

// GetMaskPointEnum return the enum type of this object
func (maskPointEyes *MaskPointEyes) GetMaskPointEnum() MaskPointEnum {
	return MaskPointEyesType
}

// MaskPointMouth A mask should be placed relatively to the mouth
type MaskPointMouth struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointMouth
func (maskPointMouth *MaskPointMouth) MessageType() string {
	return "maskPointMouth"
}

// GetMaskPointEnum return the enum type of this object
func (maskPointMouth *MaskPointMouth) GetMaskPointEnum() MaskPointEnum {
	return MaskPointMouthType
}

// MaskPointChin A mask should be placed relatively to the chin
type MaskPointChin struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointChin
func (maskPointChin *MaskPointChin) MessageType() string {
	return "maskPointChin"
}

// GetMaskPointEnum return the enum type of this object
func (maskPointChin *MaskPointChin) GetMaskPointEnum() MaskPointEnum {
	return MaskPointChinType
}

// MaskPosition Position on a photo where a mask should be placed
type MaskPosition struct {
	tdCommon
	Point  MaskPoint `json:"point"`   // Part of the face, relative to which the mask should be placed
	XShift float64   `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	YShift float64   `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
	Scale  float64   `json:"scale"`   // Mask scaling coefficient. (For example, 2.0 means a doubled size)
}

// MessageType return the string telegram-type of MaskPosition
func (maskPosition *MaskPosition) MessageType() string {
	return "maskPosition"
}

// UnmarshalJSON unmarshal to json
func (maskPosition *MaskPosition) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
		YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
		Scale  float64 `json:"scale"`   // Mask scaling coefficient. (For example, 2.0 means a doubled size)
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	maskPosition.tdCommon = tempObj.tdCommon
	maskPosition.XShift = tempObj.XShift
	maskPosition.YShift = tempObj.YShift
	maskPosition.Scale = tempObj.Scale

	fieldPoint, _ := unmarshalMaskPoint(objMap["point"])
	maskPosition.Point = fieldPoint

	return nil
}

// TextEntity Represents a part of the text that needs to be formatted in some unusual way
type TextEntity struct {
	tdCommon
	Offset int32          `json:"offset"` // Offset of the entity in UTF-16 code points
	Length int32          `json:"length"` // Length of the entity, in UTF-16 code points
	Type   TextEntityType `json:"type"`   // Type of the entity
}

// MessageType return the string telegram-type of TextEntity
func (textEntity *TextEntity) MessageType() string {
	return "textEntity"
}

// UnmarshalJSON unmarshal to json
func (textEntity *TextEntity) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Offset int32 `json:"offset"` // Offset of the entity in UTF-16 code points
		Length int32 `json:"length"` // Length of the entity, in UTF-16 code points

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	textEntity.tdCommon = tempObj.tdCommon
	textEntity.Offset = tempObj.Offset
	textEntity.Length = tempObj.Length

	fieldType, _ := unmarshalTextEntityType(objMap["type"])
	textEntity.Type = fieldType

	return nil
}

// TextEntities Contains a list of text entities
type TextEntities struct {
	tdCommon
	Entities []TextEntity `json:"entities"` // List of text entities
}

// MessageType return the string telegram-type of TextEntities
func (textEntities *TextEntities) MessageType() string {
	return "textEntities"
}

// FormattedText A text with some entities
type FormattedText struct {
	tdCommon
	Text     string       `json:"text"`     // The text
	Entities []TextEntity `json:"entities"` // Entities contained in the text
}

// MessageType return the string telegram-type of FormattedText
func (formattedText *FormattedText) MessageType() string {
	return "formattedText"
}

// Animation Describes an animation file. The animation must be encoded in GIF or MPEG4 format
type Animation struct {
	tdCommon
	Duration  int32     `json:"duration"`  // Duration of the animation, in seconds; as defined by the sender
	Width     int32     `json:"width"`     // Width of the animation
	Height    int32     `json:"height"`    // Height of the animation
	FileName  string    `json:"file_name"` // Original name of the file; as defined by the sender
	MimeType  string    `json:"mime_type"` // MIME type of the file, usually "image/gif" or "video/mp4"
	Thumbnail PhotoSize `json:"thumbnail"` // Animation thumbnail; may be null
	Animation File      `json:"animation"` // File containing the animation
}

// MessageType return the string telegram-type of Animation
func (animation *Animation) MessageType() string {
	return "animation"
}

// Audio Describes an audio file. Audio is usually in MP3 format
type Audio struct {
	tdCommon
	Duration            int32     `json:"duration"`              // Duration of the audio, in seconds; as defined by the sender
	Title               string    `json:"title"`                 // Title of the audio; as defined by the sender
	Performer           string    `json:"performer"`             // Performer of the audio; as defined by the sender
	FileName            string    `json:"file_name"`             // Original name of the file; as defined by the sender
	MimeType            string    `json:"mime_type"`             // The MIME type of the file; as defined by the sender
	AlbumCoverThumbnail PhotoSize `json:"album_cover_thumbnail"` // The thumbnail of the album cover; as defined by the sender. The full size thumbnail should be extracted from the downloaded file; may be null
	Audio               File      `json:"audio"`                 // File containing the audio
}

// MessageType return the string telegram-type of Audio
func (audio *Audio) MessageType() string {
	return "audio"
}

// Document Describes a document of any type
type Document struct {
	tdCommon
	FileName  string    `json:"file_name"` // Original name of the file; as defined by the sender
	MimeType  string    `json:"mime_type"` // MIME type of the file; as defined by the sender
	Thumbnail PhotoSize `json:"thumbnail"` // Document thumbnail; as defined by the sender; may be null
	Document  File      `json:"document"`  // File containing the document
}

// MessageType return the string telegram-type of Document
func (document *Document) MessageType() string {
	return "document"
}

// Photo Describes a photo
type Photo struct {
	tdCommon
	ID          JSONInt64   `json:"id"`           // Photo identifier; 0 for deleted photos
	HasStickers bool        `json:"has_stickers"` // True, if stickers were added to the photo
	Sizes       []PhotoSize `json:"sizes"`        // Available variants of the photo, in different sizes
}

// MessageType return the string telegram-type of Photo
func (photo *Photo) MessageType() string {
	return "photo"
}

// Sticker Describes a sticker
type Sticker struct {
	tdCommon
	SetID        JSONInt64    `json:"set_id"`        // The identifier of the sticker set to which the sticker belongs; 0 if none
	Width        int32        `json:"width"`         // Sticker width; as defined by the sender
	Height       int32        `json:"height"`        // Sticker height; as defined by the sender
	Emoji        string       `json:"emoji"`         // Emoji corresponding to the sticker
	IsMask       bool         `json:"is_mask"`       // True, if the sticker is a mask
	MaskPosition MaskPosition `json:"mask_position"` // Position where the mask should be placed; may be null
	Thumbnail    PhotoSize    `json:"thumbnail"`     // Sticker thumbnail in WEBP or JPEG format; may be null
	Sticker      File         `json:"sticker"`       // File containing the sticker
}

// MessageType return the string telegram-type of Sticker
func (sticker *Sticker) MessageType() string {
	return "sticker"
}

// Video Describes a video file
type Video struct {
	tdCommon
	Duration          int32     `json:"duration"`           // Duration of the video, in seconds; as defined by the sender
	Width             int32     `json:"width"`              // Video width; as defined by the sender
	Height            int32     `json:"height"`             // Video height; as defined by the sender
	FileName          string    `json:"file_name"`          // Original name of the file; as defined by the sender
	MimeType          string    `json:"mime_type"`          // MIME type of the file; as defined by the sender
	HasStickers       bool      `json:"has_stickers"`       // True, if stickers were added to the photo
	SupportsStreaming bool      `json:"supports_streaming"` // True, if the video should be tried to be streamed
	Thumbnail         PhotoSize `json:"thumbnail"`          // Video thumbnail; as defined by the sender; may be null
	Video             File      `json:"video"`              // File containing the video
}

// MessageType return the string telegram-type of Video
func (video *Video) MessageType() string {
	return "video"
}

// VideoNote Describes a video note. The video must be equal in width and height, cropped to a circle, and stored in MPEG4 format
type VideoNote struct {
	tdCommon
	Duration  int32     `json:"duration"`  // Duration of the video, in seconds; as defined by the sender
	Length    int32     `json:"length"`    // Video width and height; as defined by the sender
	Thumbnail PhotoSize `json:"thumbnail"` // Video thumbnail; as defined by the sender; may be null
	Video     File      `json:"video"`     // File containing the video
}

// MessageType return the string telegram-type of VideoNote
func (videoNote *VideoNote) MessageType() string {
	return "videoNote"
}

// VoiceNote Describes a voice note. The voice note must be encoded with the Opus codec, and stored inside an OGG container. Voice notes can have only a single audio channel
type VoiceNote struct {
	tdCommon
	Duration int32  `json:"duration"`  // Duration of the voice note, in seconds; as defined by the sender
	Waveform []byte `json:"waveform"`  // A waveform representation of the voice note in 5-bit format
	MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender
	Voice    File   `json:"voice"`     // File containing the voice note
}

// MessageType return the string telegram-type of VoiceNote
func (voiceNote *VoiceNote) MessageType() string {
	return "voiceNote"
}

// Contact Describes a user contact
type Contact struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // Phone number of the user
	FirstName   string `json:"first_name"`   // First name of the user; 1-255 characters in length
	LastName    string `json:"last_name"`    // Last name of the user
	UserID      int32  `json:"user_id"`      // Identifier of the user, if known; otherwise 0
}

// MessageType return the string telegram-type of Contact
func (contact *Contact) MessageType() string {
	return "contact"
}

// Location Describes a location on planet Earth
type Location struct {
	tdCommon
	Latitude  float64 `json:"latitude"`  // Latitude of the location in degrees; as defined by the sender
	Longitude float64 `json:"longitude"` // Longitude of the location, in degrees; as defined by the sender
}

// MessageType return the string telegram-type of Location
func (location *Location) MessageType() string {
	return "location"
}

// Venue Describes a venue
type Venue struct {
	tdCommon
	Location Location `json:"location"` // Venue location; as defined by the sender
	Title    string   `json:"title"`    // Venue name; as defined by the sender
	Address  string   `json:"address"`  // Venue address; as defined by the sender
	Provider string   `json:"provider"` // Provider of the venue database; as defined by the sender. Currently only "foursquare" needs to be supported
	ID       string   `json:"id"`       // Identifier of the venue in the provider database; as defined by the sender
}

// MessageType return the string telegram-type of Venue
func (venue *Venue) MessageType() string {
	return "venue"
}

// Game Describes a game
type Game struct {
	tdCommon
	ID          JSONInt64     `json:"id"`          // Game ID
	ShortName   string        `json:"short_name"`  // Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name}
	Title       string        `json:"title"`       // Game title
	Text        FormattedText `json:"text"`        // Game text, usually containing scoreboards for a game
	Description string        `json:"description"` //
	Photo       Photo         `json:"photo"`       // Game photo
	Animation   Animation     `json:"animation"`   // Game animation; may be null
}

// MessageType return the string telegram-type of Game
func (game *Game) MessageType() string {
	return "game"
}

// ProfilePhoto Describes a user profile photo
type ProfilePhoto struct {
	tdCommon
	ID    JSONInt64 `json:"id"`    // Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of userProfilePhotos
	Small File      `json:"small"` // A small (160x160) user profile photo
	Big   File      `json:"big"`   // A big (640x640) user profile photo
}

// MessageType return the string telegram-type of ProfilePhoto
func (profilePhoto *ProfilePhoto) MessageType() string {
	return "profilePhoto"
}

// ChatPhoto Describes the photo of a chat
type ChatPhoto struct {
	tdCommon
	Small File `json:"small"` // A small (160x160) chat photo
	Big   File `json:"big"`   // A big (640x640) chat photo
}

// MessageType return the string telegram-type of ChatPhoto
func (chatPhoto *ChatPhoto) MessageType() string {
	return "chatPhoto"
}

// LinkStateNone The phone number of user A is not known to user B
type LinkStateNone struct {
	tdCommon
}

// MessageType return the string telegram-type of LinkStateNone
func (linkStateNone *LinkStateNone) MessageType() string {
	return "linkStateNone"
}

// GetLinkStateEnum return the enum type of this object
func (linkStateNone *LinkStateNone) GetLinkStateEnum() LinkStateEnum {
	return LinkStateNoneType
}

// LinkStateKnowsPhoneNumber The phone number of user A is known but that number has not been saved to the contacts list of user B
type LinkStateKnowsPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of LinkStateKnowsPhoneNumber
func (linkStateKnowsPhoneNumber *LinkStateKnowsPhoneNumber) MessageType() string {
	return "linkStateKnowsPhoneNumber"
}

// GetLinkStateEnum return the enum type of this object
func (linkStateKnowsPhoneNumber *LinkStateKnowsPhoneNumber) GetLinkStateEnum() LinkStateEnum {
	return LinkStateKnowsPhoneNumberType
}

// LinkStateIsContact The phone number of user A has been saved to the contacts list of user B
type LinkStateIsContact struct {
	tdCommon
}

// MessageType return the string telegram-type of LinkStateIsContact
func (linkStateIsContact *LinkStateIsContact) MessageType() string {
	return "linkStateIsContact"
}

// GetLinkStateEnum return the enum type of this object
func (linkStateIsContact *LinkStateIsContact) GetLinkStateEnum() LinkStateEnum {
	return LinkStateIsContactType
}

// UserTypeRegular A regular user
type UserTypeRegular struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeRegular
func (userTypeRegular *UserTypeRegular) MessageType() string {
	return "userTypeRegular"
}

// GetUserTypeEnum return the enum type of this object
func (userTypeRegular *UserTypeRegular) GetUserTypeEnum() UserTypeEnum {
	return UserTypeRegularType
}

// UserTypeDeleted A deleted user or deleted bot. No information on the user besides the user_id is available. It is not possible to perform any active actions on this type of user
type UserTypeDeleted struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeDeleted
func (userTypeDeleted *UserTypeDeleted) MessageType() string {
	return "userTypeDeleted"
}

// GetUserTypeEnum return the enum type of this object
func (userTypeDeleted *UserTypeDeleted) GetUserTypeEnum() UserTypeEnum {
	return UserTypeDeletedType
}

// UserTypeBot A bot (see https://core.telegram.org/bots)
type UserTypeBot struct {
	tdCommon
	CanJoinGroups           bool   `json:"can_join_groups"`             // True, if the bot can be invited to basic group and supergroup chats
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // True, if the bot can read all messages in basic group or supergroup chats and not just those addressed to the bot. In private and channel chats a bot can always read all messages
	IsInline                bool   `json:"is_inline"`                   // True, if the bot supports inline queries
	InlineQueryPlaceholder  string `json:"inline_query_placeholder"`    // Placeholder for inline queries (displayed on the client input field)
	NeedLocation            bool   `json:"need_location"`               // True, if the location of the user should be sent with every inline query to this bot
}

// MessageType return the string telegram-type of UserTypeBot
func (userTypeBot *UserTypeBot) MessageType() string {
	return "userTypeBot"
}

// GetUserTypeEnum return the enum type of this object
func (userTypeBot *UserTypeBot) GetUserTypeEnum() UserTypeEnum {
	return UserTypeBotType
}

// UserTypeUnknown No information on the user besides the user_id is available, yet this user has not been deleted. This object is extremely rare and must be handled like a deleted user. It is not possible to perform any actions on users of this type
type UserTypeUnknown struct {
	tdCommon
}

// MessageType return the string telegram-type of UserTypeUnknown
func (userTypeUnknown *UserTypeUnknown) MessageType() string {
	return "userTypeUnknown"
}

// GetUserTypeEnum return the enum type of this object
func (userTypeUnknown *UserTypeUnknown) GetUserTypeEnum() UserTypeEnum {
	return UserTypeUnknownType
}

// BotCommand Represents commands supported by a bot
type BotCommand struct {
	tdCommon
	Command     string `json:"command"`     // Text of the bot command
	Description string `json:"description"` //
}

// MessageType return the string telegram-type of BotCommand
func (botCommand *BotCommand) MessageType() string {
	return "botCommand"
}

// BotInfo Provides information about a bot and its supported commands
type BotInfo struct {
	tdCommon
	Description string       `json:"description"` //
	Commands    []BotCommand `json:"commands"`    // A list of commands supported by the bot
}

// MessageType return the string telegram-type of BotInfo
func (botInfo *BotInfo) MessageType() string {
	return "botInfo"
}

// User Represents a user
type User struct {
	tdCommon
	ID                int32        `json:"id"`                 // User identifier
	FirstName         string       `json:"first_name"`         // First name of the user
	LastName          string       `json:"last_name"`          // Last name of the user
	Username          string       `json:"username"`           // Username of the user
	PhoneNumber       string       `json:"phone_number"`       // Phone number of the user
	Status            UserStatus   `json:"status"`             // Current online status of the user
	ProfilePhoto      ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
	OutgoingLink      LinkState    `json:"outgoing_link"`      // Relationship from the current user to the other user
	IncomingLink      LinkState    `json:"incoming_link"`      // Relationship from the other user to the current user
	IsVerified        bool         `json:"is_verified"`        // True, if the user is verified
	RestrictionReason string       `json:"restriction_reason"` // If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
	HaveAccess        bool         `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
	Type              UserType     `json:"type"`               // Type of the user
	LanguageCode      string       `json:"language_code"`      // IETF language tag of the user's language; only available to bots
}

// MessageType return the string telegram-type of User
func (user *User) MessageType() string {
	return "user"
}

// UnmarshalJSON unmarshal to json
func (user *User) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                int32        `json:"id"`                 // User identifier
		FirstName         string       `json:"first_name"`         // First name of the user
		LastName          string       `json:"last_name"`          // Last name of the user
		Username          string       `json:"username"`           // Username of the user
		PhoneNumber       string       `json:"phone_number"`       // Phone number of the user
		ProfilePhoto      ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
		IsVerified        bool         `json:"is_verified"`        // True, if the user is verified
		RestrictionReason string       `json:"restriction_reason"` // If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
		HaveAccess        bool         `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
		LanguageCode      string       `json:"language_code"`      // IETF language tag of the user's language; only available to bots
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	user.tdCommon = tempObj.tdCommon
	user.ID = tempObj.ID
	user.FirstName = tempObj.FirstName
	user.LastName = tempObj.LastName
	user.Username = tempObj.Username
	user.PhoneNumber = tempObj.PhoneNumber
	user.ProfilePhoto = tempObj.ProfilePhoto
	user.IsVerified = tempObj.IsVerified
	user.RestrictionReason = tempObj.RestrictionReason
	user.HaveAccess = tempObj.HaveAccess
	user.LanguageCode = tempObj.LanguageCode

	fieldStatus, _ := unmarshalUserStatus(objMap["status"])
	user.Status = fieldStatus

	fieldOutgoingLink, _ := unmarshalLinkState(objMap["outgoing_link"])
	user.OutgoingLink = fieldOutgoingLink

	fieldIncomingLink, _ := unmarshalLinkState(objMap["incoming_link"])
	user.IncomingLink = fieldIncomingLink

	fieldType, _ := unmarshalUserType(objMap["type"])
	user.Type = fieldType

	return nil
}

// UserFullInfo Contains full information about a user (except the full list of profile photos)
type UserFullInfo struct {
	tdCommon
	IsBlocked          bool    `json:"is_blocked"`            // True, if the user is blacklisted by the current user
	CanBeCalled        bool    `json:"can_be_called"`         // True, if the user can be called
	HasPrivateCalls    bool    `json:"has_private_calls"`     // True, if the user can't be called due to their privacy settings
	Bio                string  `json:"bio"`                   // A short user bio
	ShareText          string  `json:"share_text"`            // For bots, the text that is included with the link when users share the bot
	GroupInCommonCount int32   `json:"group_in_common_count"` // Number of group chats where both the other user and the current user are a member; 0 for the current user
	BotInfo            BotInfo `json:"bot_info"`              // If the user is a bot, information about the bot; may be null
}

// MessageType return the string telegram-type of UserFullInfo
func (userFullInfo *UserFullInfo) MessageType() string {
	return "userFullInfo"
}

// UserProfilePhotos Contains part of the list of user photos
type UserProfilePhotos struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Total number of user profile photos
	Photos     []Photo `json:"photos"`      // A list of photos
}

// MessageType return the string telegram-type of UserProfilePhotos
func (userProfilePhotos *UserProfilePhotos) MessageType() string {
	return "userProfilePhotos"
}

// Users Represents a list of users
type Users struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Approximate total count of users found
	UserIDs    []int32 `json:"user_ids"`    // A list of user identifiers
}

// MessageType return the string telegram-type of Users
func (users *Users) MessageType() string {
	return "users"
}

// ChatMemberStatusCreator The user is the creator of a chat and has all the administrator privileges
type ChatMemberStatusCreator struct {
	tdCommon
	IsMember bool `json:"is_member"` // True, if the user is a member of the chat
}

// MessageType return the string telegram-type of ChatMemberStatusCreator
func (chatMemberStatusCreator *ChatMemberStatusCreator) MessageType() string {
	return "chatMemberStatusCreator"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusCreator *ChatMemberStatusCreator) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusCreatorType
}

// ChatMemberStatusAdministrator The user is a member of a chat and has some additional privileges. In basic groups, administrators can edit and delete messages sent by others, add new members, and ban unprivileged members. In supergroups and channels, there are more detailed options for administrator privileges
type ChatMemberStatusAdministrator struct {
	tdCommon
	CanBeEdited        bool `json:"can_be_edited"`        // True, if the current user can edit the administrator privileges for the called user
	CanChangeInfo      bool `json:"can_change_info"`      // True, if the administrator can change the chat title, photo, and other settings
	CanPostMessages    bool `json:"can_post_messages"`    // True, if the administrator can create channel posts; applicable to channels only
	CanEditMessages    bool `json:"can_edit_messages"`    // True, if the administrator can edit messages of other users and pin messages; applicable to channels only
	CanDeleteMessages  bool `json:"can_delete_messages"`  // True, if the administrator can delete messages of other users
	CanInviteUsers     bool `json:"can_invite_users"`     // True, if the administrator can invite new users to the chat
	CanRestrictMembers bool `json:"can_restrict_members"` // True, if the administrator can restrict, ban, or unban chat members
	CanPinMessages     bool `json:"can_pin_messages"`     // True, if the administrator can pin messages; applicable to supergroups only
	CanPromoteMembers  bool `json:"can_promote_members"`  // True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that were directly or indirectly promoted by him
}

// MessageType return the string telegram-type of ChatMemberStatusAdministrator
func (chatMemberStatusAdministrator *ChatMemberStatusAdministrator) MessageType() string {
	return "chatMemberStatusAdministrator"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusAdministrator *ChatMemberStatusAdministrator) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusAdministratorType
}

// ChatMemberStatusMember The user is a member of a chat, without any additional privileges or restrictions
type ChatMemberStatusMember struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMemberStatusMember
func (chatMemberStatusMember *ChatMemberStatusMember) MessageType() string {
	return "chatMemberStatusMember"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusMember *ChatMemberStatusMember) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusMemberType
}

// ChatMemberStatusRestricted The user is under certain restrictions in the chat. Not supported in basic groups and channels
type ChatMemberStatusRestricted struct {
	tdCommon
	IsMember              bool  `json:"is_member"`                 // True, if the user is a member of the chat
	RestrictedUntilDate   int32 `json:"restricted_until_date"`     // Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if never. If the user is restricted for more than 366 days or for less than 30 seconds from the current time, the user is considered to be restricted forever
	CanSendMessages       bool  `json:"can_send_messages"`         // True, if the user can send text messages, contacts, locations, and venues
	CanSendMediaMessages  bool  `json:"can_send_media_messages"`   // True, if the user can send audio files, documents, photos, videos, video notes, and voice notes. Implies can_send_messages permissions
	CanSendOtherMessages  bool  `json:"can_send_other_messages"`   // True, if the user can send animations, games, and stickers and use inline bots. Implies can_send_media_messages permissions
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews"` // True, if the user may add a web page preview to his messages. Implies can_send_messages permissions
}

// MessageType return the string telegram-type of ChatMemberStatusRestricted
func (chatMemberStatusRestricted *ChatMemberStatusRestricted) MessageType() string {
	return "chatMemberStatusRestricted"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusRestricted *ChatMemberStatusRestricted) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusRestrictedType
}

// ChatMemberStatusLeft The user is not a chat member
type ChatMemberStatusLeft struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMemberStatusLeft
func (chatMemberStatusLeft *ChatMemberStatusLeft) MessageType() string {
	return "chatMemberStatusLeft"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusLeft *ChatMemberStatusLeft) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusLeftType
}

// ChatMemberStatusBanned The user was banned (and hence is not a member of the chat). Implies the user can't return to the chat or view messages
type ChatMemberStatusBanned struct {
	tdCommon
	BannedUntilDate int32 `json:"banned_until_date"` // Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever
}

// MessageType return the string telegram-type of ChatMemberStatusBanned
func (chatMemberStatusBanned *ChatMemberStatusBanned) MessageType() string {
	return "chatMemberStatusBanned"
}

// GetChatMemberStatusEnum return the enum type of this object
func (chatMemberStatusBanned *ChatMemberStatusBanned) GetChatMemberStatusEnum() ChatMemberStatusEnum {
	return ChatMemberStatusBannedType
}

// ChatMember A user with information about joining/leaving a chat
type ChatMember struct {
	tdCommon
	UserID         int32            `json:"user_id"`          // User identifier of the chat member
	InviterUserID  int32            `json:"inviter_user_id"`  // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
	JoinedChatDate int32            `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined a chat
	Status         ChatMemberStatus `json:"status"`           // Status of the member in the chat
	BotInfo        BotInfo          `json:"bot_info"`         // If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
}

// MessageType return the string telegram-type of ChatMember
func (chatMember *ChatMember) MessageType() string {
	return "chatMember"
}

// UnmarshalJSON unmarshal to json
func (chatMember *ChatMember) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserID         int32   `json:"user_id"`          // User identifier of the chat member
		InviterUserID  int32   `json:"inviter_user_id"`  // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
		JoinedChatDate int32   `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined a chat
		BotInfo        BotInfo `json:"bot_info"`         // If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatMember.tdCommon = tempObj.tdCommon
	chatMember.UserID = tempObj.UserID
	chatMember.InviterUserID = tempObj.InviterUserID
	chatMember.JoinedChatDate = tempObj.JoinedChatDate
	chatMember.BotInfo = tempObj.BotInfo

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	chatMember.Status = fieldStatus

	return nil
}

// ChatMembers Contains a list of chat members
type ChatMembers struct {
	tdCommon
	TotalCount int32        `json:"total_count"` // Approximate total count of chat members found
	Members    []ChatMember `json:"members"`     // A list of chat members
}

// MessageType return the string telegram-type of ChatMembers
func (chatMembers *ChatMembers) MessageType() string {
	return "chatMembers"
}

// SupergroupMembersFilterRecent Returns recently active users in reverse chronological order
type SupergroupMembersFilterRecent struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterRecent
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) MessageType() string {
	return "supergroupMembersFilterRecent"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterRecentType
}

// SupergroupMembersFilterAdministrators Returns the creator and administrators
type SupergroupMembersFilterAdministrators struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterAdministrators
func (supergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators) MessageType() string {
	return "supergroupMembersFilterAdministrators"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterAdministratorsType
}

// SupergroupMembersFilterSearch Used to search for supergroup or channel members via a (string) query
type SupergroupMembersFilterSearch struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterSearch
func (supergroupMembersFilterSearch *SupergroupMembersFilterSearch) MessageType() string {
	return "supergroupMembersFilterSearch"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterSearch *SupergroupMembersFilterSearch) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterSearchType
}

// SupergroupMembersFilterRestricted Returns restricted supergroup members; can be used only by administrators
type SupergroupMembersFilterRestricted struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterRestricted
func (supergroupMembersFilterRestricted *SupergroupMembersFilterRestricted) MessageType() string {
	return "supergroupMembersFilterRestricted"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterRestricted *SupergroupMembersFilterRestricted) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterRestrictedType
}

// SupergroupMembersFilterBanned Returns users banned from the supergroup or channel; can be used only by administrators
type SupergroupMembersFilterBanned struct {
	tdCommon
	Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterBanned
func (supergroupMembersFilterBanned *SupergroupMembersFilterBanned) MessageType() string {
	return "supergroupMembersFilterBanned"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterBanned *SupergroupMembersFilterBanned) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterBannedType
}

// SupergroupMembersFilterBots Returns bot members of the supergroup or channel
type SupergroupMembersFilterBots struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterBots
func (supergroupMembersFilterBots *SupergroupMembersFilterBots) MessageType() string {
	return "supergroupMembersFilterBots"
}

// GetSupergroupMembersFilterEnum return the enum type of this object
func (supergroupMembersFilterBots *SupergroupMembersFilterBots) GetSupergroupMembersFilterEnum() SupergroupMembersFilterEnum {
	return SupergroupMembersFilterBotsType
}

// BasicGroup Represents a basic group of 0-200 users (must be upgraded to a supergroup to accommodate more than 200 users)
type BasicGroup struct {
	tdCommon
	ID                      int32            `json:"id"`                        // Group identifier
	MemberCount             int32            `json:"member_count"`              // Number of members in the group
	Status                  ChatMemberStatus `json:"status"`                    // Status of the current user in the group
	EveryoneIsAdministrator bool             `json:"everyone_is_administrator"` // True, if all members have been granted administrator rights in the group
	IsActive                bool             `json:"is_active"`                 // True, if the group is active
	UpgradedToSupergroupID  int32            `json:"upgraded_to_supergroup_id"` // Identifier of the supergroup to which this group was upgraded; 0 if none
}

// MessageType return the string telegram-type of BasicGroup
func (basicGroup *BasicGroup) MessageType() string {
	return "basicGroup"
}

// UnmarshalJSON unmarshal to json
func (basicGroup *BasicGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                      int32 `json:"id"`                        // Group identifier
		MemberCount             int32 `json:"member_count"`              // Number of members in the group
		EveryoneIsAdministrator bool  `json:"everyone_is_administrator"` // True, if all members have been granted administrator rights in the group
		IsActive                bool  `json:"is_active"`                 // True, if the group is active
		UpgradedToSupergroupID  int32 `json:"upgraded_to_supergroup_id"` // Identifier of the supergroup to which this group was upgraded; 0 if none
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	basicGroup.tdCommon = tempObj.tdCommon
	basicGroup.ID = tempObj.ID
	basicGroup.MemberCount = tempObj.MemberCount
	basicGroup.EveryoneIsAdministrator = tempObj.EveryoneIsAdministrator
	basicGroup.IsActive = tempObj.IsActive
	basicGroup.UpgradedToSupergroupID = tempObj.UpgradedToSupergroupID

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	basicGroup.Status = fieldStatus

	return nil
}

// BasicGroupFullInfo Contains full information about a basic group
type BasicGroupFullInfo struct {
	tdCommon
	CreatorUserID int32        `json:"creator_user_id"` // User identifier of the creator of the group; 0 if unknown
	Members       []ChatMember `json:"members"`         // Group members
	InviteLink    string       `json:"invite_link"`     // Invite link for this group; available only for the group creator and only after it has been generated at least once
}

// MessageType return the string telegram-type of BasicGroupFullInfo
func (basicGroupFullInfo *BasicGroupFullInfo) MessageType() string {
	return "basicGroupFullInfo"
}

// Supergroup Represents a supergroup or channel with zero or more members (subscribers in the case of channels). From the point of view of the system, a channel is a special kind of a supergroup: only administrators can post and see the list of members, and posts from all administrators use the name and photo of the channel instead of individual names and profile photos. Unlike supergroups, channels can have an unlimited number of subscribers
type Supergroup struct {
	tdCommon
	ID                int32            `json:"id"`                 // Supergroup or channel identifier
	Username          string           `json:"username"`           // Username of the supergroup or channel; empty for private supergroups or channels
	Date              int32            `json:"date"`               // Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
	Status            ChatMemberStatus `json:"status"`             // Status of the current user in the supergroup or channel
	MemberCount       int32            `json:"member_count"`       // Member count; 0 if unknown. Currently it is guaranteed to be known only if the supergroup or channel was found through SearchPublicChats
	AnyoneCanInvite   bool             `json:"anyone_can_invite"`  // True, if any member of the supergroup can invite other members. This field has no meaning for channels
	SignMessages      bool             `json:"sign_messages"`      // True, if messages sent to the channel should contain information about the sender. This field is only applicable to channels
	IsChannel         bool             `json:"is_channel"`         // True, if the supergroup is a channel
	IsVerified        bool             `json:"is_verified"`        // True, if the supergroup or channel is verified
	RestrictionReason string           `json:"restriction_reason"` // If non-empty, contains the reason why access to this supergroup or channel must be restricted. Format of the string is "{type}: {description}".
}

// MessageType return the string telegram-type of Supergroup
func (supergroup *Supergroup) MessageType() string {
	return "supergroup"
}

// UnmarshalJSON unmarshal to json
func (supergroup *Supergroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                int32  `json:"id"`                 // Supergroup or channel identifier
		Username          string `json:"username"`           // Username of the supergroup or channel; empty for private supergroups or channels
		Date              int32  `json:"date"`               // Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
		MemberCount       int32  `json:"member_count"`       // Member count; 0 if unknown. Currently it is guaranteed to be known only if the supergroup or channel was found through SearchPublicChats
		AnyoneCanInvite   bool   `json:"anyone_can_invite"`  // True, if any member of the supergroup can invite other members. This field has no meaning for channels
		SignMessages      bool   `json:"sign_messages"`      // True, if messages sent to the channel should contain information about the sender. This field is only applicable to channels
		IsChannel         bool   `json:"is_channel"`         // True, if the supergroup is a channel
		IsVerified        bool   `json:"is_verified"`        // True, if the supergroup or channel is verified
		RestrictionReason string `json:"restriction_reason"` // If non-empty, contains the reason why access to this supergroup or channel must be restricted. Format of the string is "{type}: {description}".
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	supergroup.tdCommon = tempObj.tdCommon
	supergroup.ID = tempObj.ID
	supergroup.Username = tempObj.Username
	supergroup.Date = tempObj.Date
	supergroup.MemberCount = tempObj.MemberCount
	supergroup.AnyoneCanInvite = tempObj.AnyoneCanInvite
	supergroup.SignMessages = tempObj.SignMessages
	supergroup.IsChannel = tempObj.IsChannel
	supergroup.IsVerified = tempObj.IsVerified
	supergroup.RestrictionReason = tempObj.RestrictionReason

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	supergroup.Status = fieldStatus

	return nil
}

// SupergroupFullInfo Contains full information about a supergroup or channel
type SupergroupFullInfo struct {
	tdCommon
	Description              string    `json:"description"`                  //
	MemberCount              int32     `json:"member_count"`                 // Number of members in the supergroup or channel; 0 if unknown
	AdministratorCount       int32     `json:"administrator_count"`          // Number of privileged users in the supergroup or channel; 0 if unknown
	RestrictedCount          int32     `json:"restricted_count"`             // Number of restricted users in the supergroup; 0 if unknown
	BannedCount              int32     `json:"banned_count"`                 // Number of users banned from chat; 0 if unknown
	CanGetMembers            bool      `json:"can_get_members"`              // True, if members of the chat can be retrieved
	CanSetUsername           bool      `json:"can_set_username"`             // True, if the chat can be made public
	CanSetStickerSet         bool      `json:"can_set_sticker_set"`          // True, if the supergroup sticker set can be changed
	IsAllHistoryAvailable    bool      `json:"is_all_history_available"`     // True, if new chat members will have access to old messages. In public supergroups and both public and private channels, old messages are always available, so this option affects only private supergroups. The value of this field is only available for chat administrators
	StickerSetID             JSONInt64 `json:"sticker_set_id"`               // Identifier of the supergroup sticker set; 0 if none
	InviteLink               string    `json:"invite_link"`                  // Invite link for this chat
	PinnedMessageID          int64     `json:"pinned_message_id"`            // Identifier of the pinned message in the chat; 0 if none
	UpgradedFromBasicGroupID int32     `json:"upgraded_from_basic_group_id"` // Identifier of the basic group from which supergroup was upgraded; 0 if none
	UpgradedFromMaxMessageID int64     `json:"upgraded_from_max_message_id"` // Identifier of the last message in the basic group from which supergroup was upgraded; 0 if none
}

// MessageType return the string telegram-type of SupergroupFullInfo
func (supergroupFullInfo *SupergroupFullInfo) MessageType() string {
	return "supergroupFullInfo"
}

// SecretChatStatePending The secret chat is not yet created; waiting for the other user to get online
type SecretChatStatePending struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStatePending
func (secretChatStatePending *SecretChatStatePending) MessageType() string {
	return "secretChatStatePending"
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStatePending *SecretChatStatePending) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStatePendingType
}

// SecretChatStateReady The secret chat is ready to use
type SecretChatStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStateReady
func (secretChatStateReady *SecretChatStateReady) MessageType() string {
	return "secretChatStateReady"
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStateReady *SecretChatStateReady) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStateReadyType
}

// SecretChatStateClosed The secret chat is closed
type SecretChatStateClosed struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStateClosed
func (secretChatStateClosed *SecretChatStateClosed) MessageType() string {
	return "secretChatStateClosed"
}

// GetSecretChatStateEnum return the enum type of this object
func (secretChatStateClosed *SecretChatStateClosed) GetSecretChatStateEnum() SecretChatStateEnum {
	return SecretChatStateClosedType
}

// SecretChat Represents a secret chat
type SecretChat struct {
	tdCommon
	ID         int32           `json:"id"`          // Secret chat identifier
	UserID     int32           `json:"user_id"`     // Identifier of the chat partner
	State      SecretChatState `json:"state"`       // State of the secret chat
	IsOutbound bool            `json:"is_outbound"` // True, if the chat was created by the current user; otherwise false
	TTL        int32           `json:"ttl"`         // Current message Time To Live setting (self-destruct timer) for the chat, in seconds
	KeyHash    []byte          `json:"key_hash"`    // Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 bytes, which must be used to make a 12x12 square image with a color depth of 4. The first 16 bytes should be used to make a central 8x8 square, while the remaining 20 bytes should be used to construct a 2-pixel-wide border around that square.
	Layer      int32           `json:"layer"`       // Secret chat layer; determines features supported by the other client. Video notes are supported if the layer >= 66
}

// MessageType return the string telegram-type of SecretChat
func (secretChat *SecretChat) MessageType() string {
	return "secretChat"
}

// UnmarshalJSON unmarshal to json
func (secretChat *SecretChat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID         int32  `json:"id"`          // Secret chat identifier
		UserID     int32  `json:"user_id"`     // Identifier of the chat partner
		IsOutbound bool   `json:"is_outbound"` // True, if the chat was created by the current user; otherwise false
		TTL        int32  `json:"ttl"`         // Current message Time To Live setting (self-destruct timer) for the chat, in seconds
		KeyHash    []byte `json:"key_hash"`    // Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 bytes, which must be used to make a 12x12 square image with a color depth of 4. The first 16 bytes should be used to make a central 8x8 square, while the remaining 20 bytes should be used to construct a 2-pixel-wide border around that square.
		Layer      int32  `json:"layer"`       // Secret chat layer; determines features supported by the other client. Video notes are supported if the layer >= 66
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	secretChat.tdCommon = tempObj.tdCommon
	secretChat.ID = tempObj.ID
	secretChat.UserID = tempObj.UserID
	secretChat.IsOutbound = tempObj.IsOutbound
	secretChat.TTL = tempObj.TTL
	secretChat.KeyHash = tempObj.KeyHash
	secretChat.Layer = tempObj.Layer

	fieldState, _ := unmarshalSecretChatState(objMap["state"])
	secretChat.State = fieldState

	return nil
}

// MessageForwardedFromUser The message was originally written by a known user
type MessageForwardedFromUser struct {
	tdCommon
	SenderUserID           int32 `json:"sender_user_id"`            // Identifier of the user that originally sent this message
	Date                   int32 `json:"date"`                      // Point in time (Unix timestamp) when the message was originally sent
	ForwardedFromChatID    int64 `json:"forwarded_from_chat_id"`    // For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
	ForwardedFromMessageID int64 `json:"forwarded_from_message_id"` // For messages forwarded to the chat with the current user (saved messages) the identifier of the original message from which the new message was forwarded; 0 if unknown
}

// MessageType return the string telegram-type of MessageForwardedFromUser
func (messageForwardedFromUser *MessageForwardedFromUser) MessageType() string {
	return "messageForwardedFromUser"
}

// GetMessageForwardInfoEnum return the enum type of this object
func (messageForwardedFromUser *MessageForwardedFromUser) GetMessageForwardInfoEnum() MessageForwardInfoEnum {
	return MessageForwardedFromUserType
}

// MessageForwardedPost The message was originally a post in a channel
type MessageForwardedPost struct {
	tdCommon
	ChatID                 int64  `json:"chat_id"`                   // Identifier of the chat from which the message was forwarded
	AuthorSignature        string `json:"author_signature"`          // Post author signature
	Date                   int32  `json:"date"`                      // Point in time (Unix timestamp) when the message was originally sent
	MessageID              int64  `json:"message_id"`                // Message identifier of the original message from which the new message was forwarded; 0 if unknown
	ForwardedFromChatID    int64  `json:"forwarded_from_chat_id"`    // For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
	ForwardedFromMessageID int64  `json:"forwarded_from_message_id"` // For messages forwarded to the chat with the current user (saved messages), the identifier of the original message from which the new message was forwarded; 0 if unknown
}

// MessageType return the string telegram-type of MessageForwardedPost
func (messageForwardedPost *MessageForwardedPost) MessageType() string {
	return "messageForwardedPost"
}

// GetMessageForwardInfoEnum return the enum type of this object
func (messageForwardedPost *MessageForwardedPost) GetMessageForwardInfoEnum() MessageForwardInfoEnum {
	return MessageForwardedPostType
}

// MessageSendingStatePending The message is being sent now, but has not yet been delivered to the server
type MessageSendingStatePending struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageSendingStatePending
func (messageSendingStatePending *MessageSendingStatePending) MessageType() string {
	return "messageSendingStatePending"
}

// GetMessageSendingStateEnum return the enum type of this object
func (messageSendingStatePending *MessageSendingStatePending) GetMessageSendingStateEnum() MessageSendingStateEnum {
	return MessageSendingStatePendingType
}

// MessageSendingStateFailed The message failed to be sent
type MessageSendingStateFailed struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageSendingStateFailed
func (messageSendingStateFailed *MessageSendingStateFailed) MessageType() string {
	return "messageSendingStateFailed"
}

// GetMessageSendingStateEnum return the enum type of this object
func (messageSendingStateFailed *MessageSendingStateFailed) GetMessageSendingStateEnum() MessageSendingStateEnum {
	return MessageSendingStateFailedType
}

// Message Describes a message
type Message struct {
	tdCommon
	ID                      int64               `json:"id"`                           // Unique message identifier
	SenderUserID            int32               `json:"sender_user_id"`               // Identifier of the user who sent the message; 0 if unknown. It is unknown for channel posts
	ChatID                  int64               `json:"chat_id"`                      // Chat identifier
	SendingState            MessageSendingState `json:"sending_state"`                // Information about the sending state of the message; may be null
	IsOutgoing              bool                `json:"is_outgoing"`                  // True, if the message is outgoing
	CanBeEdited             bool                `json:"can_be_edited"`                // True, if the message can be edited
	CanBeForwarded          bool                `json:"can_be_forwarded"`             // True, if the message can be forwarded
	CanBeDeletedOnlyForSelf bool                `json:"can_be_deleted_only_for_self"` // True, if the message can be deleted only for the current user while other users will continue to see it
	CanBeDeletedForAllUsers bool                `json:"can_be_deleted_for_all_users"` // True, if the message can be deleted for all users
	IsChannelPost           bool                `json:"is_channel_post"`              // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
	ContainsUnreadMention   bool                `json:"contains_unread_mention"`      // True, if the message contains an unread mention for the current user
	Date                    int32               `json:"date"`                         // Point in time (Unix timestamp) when the message was sent
	EditDate                int32               `json:"edit_date"`                    // Point in time (Unix timestamp) when the message was last edited
	ForwardInfo             MessageForwardInfo  `json:"forward_info"`                 // Information about the initial message sender; may be null
	ReplyToMessageID        int64               `json:"reply_to_message_id"`          // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
	TTL                     int32               `json:"ttl"`                          // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
	TTLExpiresIn            float64             `json:"ttl_expires_in"`               // Time left before the message expires, in seconds
	ViaBotUserID            int32               `json:"via_bot_user_id"`              // If non-zero, the user identifier of the bot through which this message was sent
	AuthorSignature         string              `json:"author_signature"`             // For channel posts, optional author signature
	Views                   int32               `json:"views"`                        // Number of times this message was viewed
	MediaAlbumID            JSONInt64           `json:"media_album_id"`               // Unique identifier of an album this message belongs to. Only photos and videos can be grouped together in albums
	Content                 MessageContent      `json:"content"`                      // Content of the message
	ReplyMarkup             ReplyMarkup         `json:"reply_markup"`                 // Reply markup for the message; may be null
}

// MessageType return the string telegram-type of Message
func (message *Message) MessageType() string {
	return "message"
}

// UnmarshalJSON unmarshal to json
func (message *Message) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                      int64     `json:"id"`                           // Unique message identifier
		SenderUserID            int32     `json:"sender_user_id"`               // Identifier of the user who sent the message; 0 if unknown. It is unknown for channel posts
		ChatID                  int64     `json:"chat_id"`                      // Chat identifier
		IsOutgoing              bool      `json:"is_outgoing"`                  // True, if the message is outgoing
		CanBeEdited             bool      `json:"can_be_edited"`                // True, if the message can be edited
		CanBeForwarded          bool      `json:"can_be_forwarded"`             // True, if the message can be forwarded
		CanBeDeletedOnlyForSelf bool      `json:"can_be_deleted_only_for_self"` // True, if the message can be deleted only for the current user while other users will continue to see it
		CanBeDeletedForAllUsers bool      `json:"can_be_deleted_for_all_users"` // True, if the message can be deleted for all users
		IsChannelPost           bool      `json:"is_channel_post"`              // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
		ContainsUnreadMention   bool      `json:"contains_unread_mention"`      // True, if the message contains an unread mention for the current user
		Date                    int32     `json:"date"`                         // Point in time (Unix timestamp) when the message was sent
		EditDate                int32     `json:"edit_date"`                    // Point in time (Unix timestamp) when the message was last edited
		ReplyToMessageID        int64     `json:"reply_to_message_id"`          // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
		TTL                     int32     `json:"ttl"`                          // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
		TTLExpiresIn            float64   `json:"ttl_expires_in"`               // Time left before the message expires, in seconds
		ViaBotUserID            int32     `json:"via_bot_user_id"`              // If non-zero, the user identifier of the bot through which this message was sent
		AuthorSignature         string    `json:"author_signature"`             // For channel posts, optional author signature
		Views                   int32     `json:"views"`                        // Number of times this message was viewed
		MediaAlbumID            JSONInt64 `json:"media_album_id"`               // Unique identifier of an album this message belongs to. Only photos and videos can be grouped together in albums

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	message.tdCommon = tempObj.tdCommon
	message.ID = tempObj.ID
	message.SenderUserID = tempObj.SenderUserID
	message.ChatID = tempObj.ChatID
	message.IsOutgoing = tempObj.IsOutgoing
	message.CanBeEdited = tempObj.CanBeEdited
	message.CanBeForwarded = tempObj.CanBeForwarded
	message.CanBeDeletedOnlyForSelf = tempObj.CanBeDeletedOnlyForSelf
	message.CanBeDeletedForAllUsers = tempObj.CanBeDeletedForAllUsers
	message.IsChannelPost = tempObj.IsChannelPost
	message.ContainsUnreadMention = tempObj.ContainsUnreadMention
	message.Date = tempObj.Date
	message.EditDate = tempObj.EditDate
	message.ReplyToMessageID = tempObj.ReplyToMessageID
	message.TTL = tempObj.TTL
	message.TTLExpiresIn = tempObj.TTLExpiresIn
	message.ViaBotUserID = tempObj.ViaBotUserID
	message.AuthorSignature = tempObj.AuthorSignature
	message.Views = tempObj.Views
	message.MediaAlbumID = tempObj.MediaAlbumID

	fieldSendingState, _ := unmarshalMessageSendingState(objMap["sending_state"])
	message.SendingState = fieldSendingState

	fieldForwardInfo, _ := unmarshalMessageForwardInfo(objMap["forward_info"])
	message.ForwardInfo = fieldForwardInfo

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	message.Content = fieldContent

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	message.ReplyMarkup = fieldReplyMarkup

	return nil
}

// Messages Contains a list of messages
type Messages struct {
	tdCommon
	TotalCount int32     `json:"total_count"` // Approximate total count of messages found
	Messages   []Message `json:"messages"`    // List of messages; messages may be null
}

// MessageType return the string telegram-type of Messages
func (messages *Messages) MessageType() string {
	return "messages"
}

// FoundMessages Contains a list of messages found by a search
type FoundMessages struct {
	tdCommon
	Messages         []Message `json:"messages"`            // List of messages
	NextFromSearchID JSONInt64 `json:"next_from_search_id"` // Value to pass as from_search_id to get more results
}

// MessageType return the string telegram-type of FoundMessages
func (foundMessages *FoundMessages) MessageType() string {
	return "foundMessages"
}

// NotificationSettingsScopeChat Notification settings applied to a particular chat
type NotificationSettingsScopeChat struct {
	tdCommon
	ChatID int64 `json:"chat_id"` // Chat identifier
}

// MessageType return the string telegram-type of NotificationSettingsScopeChat
func (notificationSettingsScopeChat *NotificationSettingsScopeChat) MessageType() string {
	return "notificationSettingsScopeChat"
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeChat *NotificationSettingsScopeChat) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeChatType
}

// NotificationSettingsScopePrivateChats Notification settings applied to all private chats
type NotificationSettingsScopePrivateChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopePrivateChats
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) MessageType() string {
	return "notificationSettingsScopePrivateChats"
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopePrivateChatsType
}

// NotificationSettingsScopeBasicGroupChats Notification settings applied to all basic groups and channels. (Supergroups have no common settings)
type NotificationSettingsScopeBasicGroupChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopeBasicGroupChats
func (notificationSettingsScopeBasicGroupChats *NotificationSettingsScopeBasicGroupChats) MessageType() string {
	return "notificationSettingsScopeBasicGroupChats"
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeBasicGroupChats *NotificationSettingsScopeBasicGroupChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeBasicGroupChatsType
}

// NotificationSettingsScopeAllChats Notification settings applied to all chats
type NotificationSettingsScopeAllChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopeAllChats
func (notificationSettingsScopeAllChats *NotificationSettingsScopeAllChats) MessageType() string {
	return "notificationSettingsScopeAllChats"
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeAllChats *NotificationSettingsScopeAllChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeAllChatsType
}

// NotificationSettings Contains information about notification settings for a chat or several chats
type NotificationSettings struct {
	tdCommon
	MuteFor     int32  `json:"mute_for"`     // Time left before notifications will be unmuted, in seconds
	Sound       string `json:"sound"`        // An audio file name for notification sounds; only applies to iOS applications
	ShowPreview bool   `json:"show_preview"` // True, if message content should be displayed in notifications
}

// MessageType return the string telegram-type of NotificationSettings
func (notificationSettings *NotificationSettings) MessageType() string {
	return "notificationSettings"
}

// DraftMessage Contains information about a message draft
type DraftMessage struct {
	tdCommon
	ReplyToMessageID int64               `json:"reply_to_message_id"` // Identifier of the message to reply to; 0 if none
	InputMessageText InputMessageContent `json:"input_message_text"`  // Content of the message draft; this should always be of type inputMessageText
}

// MessageType return the string telegram-type of DraftMessage
func (draftMessage *DraftMessage) MessageType() string {
	return "draftMessage"
}

// UnmarshalJSON unmarshal to json
func (draftMessage *DraftMessage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ReplyToMessageID int64 `json:"reply_to_message_id"` // Identifier of the message to reply to; 0 if none

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	draftMessage.tdCommon = tempObj.tdCommon
	draftMessage.ReplyToMessageID = tempObj.ReplyToMessageID

	fieldInputMessageText, _ := unmarshalInputMessageContent(objMap["input_message_text"])
	draftMessage.InputMessageText = fieldInputMessageText

	return nil
}

// ChatTypePrivate An ordinary chat with a user
type ChatTypePrivate struct {
	tdCommon
	UserID int32 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of ChatTypePrivate
func (chatTypePrivate *ChatTypePrivate) MessageType() string {
	return "chatTypePrivate"
}

// GetChatTypeEnum return the enum type of this object
func (chatTypePrivate *ChatTypePrivate) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypePrivateType
}

// ChatTypeBasicGroup A basic group (i.e., a chat with 0-200 other users)
type ChatTypeBasicGroup struct {
	tdCommon
	BasicGroupID int32 `json:"basic_group_id"` // Basic group identifier
}

// MessageType return the string telegram-type of ChatTypeBasicGroup
func (chatTypeBasicGroup *ChatTypeBasicGroup) MessageType() string {
	return "chatTypeBasicGroup"
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeBasicGroup *ChatTypeBasicGroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeBasicGroupType
}

// ChatTypeSupergroup A supergroup (i.e. a chat with up to GetOption("supergroup_max_size") other users), or channel (with unlimited members)
type ChatTypeSupergroup struct {
	tdCommon
	SupergroupID int32 `json:"supergroup_id"` // Supergroup or channel identifier
	IsChannel    bool  `json:"is_channel"`    // True, if the supergroup is a channel
}

// MessageType return the string telegram-type of ChatTypeSupergroup
func (chatTypeSupergroup *ChatTypeSupergroup) MessageType() string {
	return "chatTypeSupergroup"
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSupergroup *ChatTypeSupergroup) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSupergroupType
}

// ChatTypeSecret A secret chat with a user
type ChatTypeSecret struct {
	tdCommon
	SecretChatID int32 `json:"secret_chat_id"` // Secret chat identifier
	UserID       int32 `json:"user_id"`        // User identifier of the secret chat peer
}

// MessageType return the string telegram-type of ChatTypeSecret
func (chatTypeSecret *ChatTypeSecret) MessageType() string {
	return "chatTypeSecret"
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSecret *ChatTypeSecret) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSecretType
}

// Chat A chat. (Can be a private chat, basic group, supergroup, or secret chat)
type Chat struct {
	tdCommon
	ID                      int64                `json:"id"`                          // Chat unique identifier
	Type                    ChatType             `json:"type"`                        // Type of the chat
	Title                   string               `json:"title"`                       // Chat title
	Photo                   ChatPhoto            `json:"photo"`                       // Chat photo; may be null
	LastMessage             Message              `json:"last_message"`                // Last message in the chat; may be null
	Order                   JSONInt64            `json:"order"`                       // Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
	IsPinned                bool                 `json:"is_pinned"`                   // True, if the chat is pinned
	CanBeReported           bool                 `json:"can_be_reported"`             // True, if the chat can be reported to Telegram moderators through reportChat
	UnreadCount             int32                `json:"unread_count"`                // Number of unread messages in the chat
	LastReadInboxMessageID  int64                `json:"last_read_inbox_message_id"`  // Identifier of the last read incoming message
	LastReadOutboxMessageID int64                `json:"last_read_outbox_message_id"` // Identifier of the last read outgoing message
	UnreadMentionCount      int32                `json:"unread_mention_count"`        // Number of unread messages with a mention/reply in the chat
	NotificationSettings    NotificationSettings `json:"notification_settings"`       // Notification settings for this chat
	ReplyMarkupMessageID    int64                `json:"reply_markup_message_id"`     // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
	DraftMessage            DraftMessage         `json:"draft_message"`               // A draft of a message in the chat; may be null
	ClientData              string               `json:"client_data"`                 // Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
}

// MessageType return the string telegram-type of Chat
func (chat *Chat) MessageType() string {
	return "chat"
}

// UnmarshalJSON unmarshal to json
func (chat *Chat) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                      int64                `json:"id"`                          // Chat unique identifier
		Title                   string               `json:"title"`                       // Chat title
		Photo                   ChatPhoto            `json:"photo"`                       // Chat photo; may be null
		LastMessage             Message              `json:"last_message"`                // Last message in the chat; may be null
		Order                   JSONInt64            `json:"order"`                       // Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
		IsPinned                bool                 `json:"is_pinned"`                   // True, if the chat is pinned
		CanBeReported           bool                 `json:"can_be_reported"`             // True, if the chat can be reported to Telegram moderators through reportChat
		UnreadCount             int32                `json:"unread_count"`                // Number of unread messages in the chat
		LastReadInboxMessageID  int64                `json:"last_read_inbox_message_id"`  // Identifier of the last read incoming message
		LastReadOutboxMessageID int64                `json:"last_read_outbox_message_id"` // Identifier of the last read outgoing message
		UnreadMentionCount      int32                `json:"unread_mention_count"`        // Number of unread messages with a mention/reply in the chat
		NotificationSettings    NotificationSettings `json:"notification_settings"`       // Notification settings for this chat
		ReplyMarkupMessageID    int64                `json:"reply_markup_message_id"`     // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
		DraftMessage            DraftMessage         `json:"draft_message"`               // A draft of a message in the chat; may be null
		ClientData              string               `json:"client_data"`                 // Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chat.tdCommon = tempObj.tdCommon
	chat.ID = tempObj.ID
	chat.Title = tempObj.Title
	chat.Photo = tempObj.Photo
	chat.LastMessage = tempObj.LastMessage
	chat.Order = tempObj.Order
	chat.IsPinned = tempObj.IsPinned
	chat.CanBeReported = tempObj.CanBeReported
	chat.UnreadCount = tempObj.UnreadCount
	chat.LastReadInboxMessageID = tempObj.LastReadInboxMessageID
	chat.LastReadOutboxMessageID = tempObj.LastReadOutboxMessageID
	chat.UnreadMentionCount = tempObj.UnreadMentionCount
	chat.NotificationSettings = tempObj.NotificationSettings
	chat.ReplyMarkupMessageID = tempObj.ReplyMarkupMessageID
	chat.DraftMessage = tempObj.DraftMessage
	chat.ClientData = tempObj.ClientData

	fieldType, _ := unmarshalChatType(objMap["type"])
	chat.Type = fieldType

	return nil
}

// Chats Represents a list of chats
type Chats struct {
	tdCommon
	ChatIDs []int64 `json:"chat_ids"` // List of chat identifiers
}

// MessageType return the string telegram-type of Chats
func (chats *Chats) MessageType() string {
	return "chats"
}

// ChatInviteLink Contains a chat invite link
type ChatInviteLink struct {
	tdCommon
	InviteLink string `json:"invite_link"` // Chat invite link
}

// MessageType return the string telegram-type of ChatInviteLink
func (chatInviteLink *ChatInviteLink) MessageType() string {
	return "chatInviteLink"
}

// ChatInviteLinkInfo Contains information about a chat invite link
type ChatInviteLinkInfo struct {
	tdCommon
	ChatID        int64     `json:"chat_id"`         // Chat identifier of the invite link; 0 if the user is not a member of this chat
	Type          ChatType  `json:"type"`            // Contains information about the type of the chat
	Title         string    `json:"title"`           // Title of the chat
	Photo         ChatPhoto `json:"photo"`           // Chat photo; may be null
	MemberCount   int32     `json:"member_count"`    // Number of members
	MemberUserIDs []int32   `json:"member_user_ids"` // User identifiers of some chat members that may be known to the current user
	IsPublic      bool      `json:"is_public"`       // True, if the chat is a public supergroup or channel with a username
}

// MessageType return the string telegram-type of ChatInviteLinkInfo
func (chatInviteLinkInfo *ChatInviteLinkInfo) MessageType() string {
	return "chatInviteLinkInfo"
}

// UnmarshalJSON unmarshal to json
func (chatInviteLinkInfo *ChatInviteLinkInfo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatID        int64     `json:"chat_id"`         // Chat identifier of the invite link; 0 if the user is not a member of this chat
		Title         string    `json:"title"`           // Title of the chat
		Photo         ChatPhoto `json:"photo"`           // Chat photo; may be null
		MemberCount   int32     `json:"member_count"`    // Number of members
		MemberUserIDs []int32   `json:"member_user_ids"` // User identifiers of some chat members that may be known to the current user
		IsPublic      bool      `json:"is_public"`       // True, if the chat is a public supergroup or channel with a username
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatInviteLinkInfo.tdCommon = tempObj.tdCommon
	chatInviteLinkInfo.ChatID = tempObj.ChatID
	chatInviteLinkInfo.Title = tempObj.Title
	chatInviteLinkInfo.Photo = tempObj.Photo
	chatInviteLinkInfo.MemberCount = tempObj.MemberCount
	chatInviteLinkInfo.MemberUserIDs = tempObj.MemberUserIDs
	chatInviteLinkInfo.IsPublic = tempObj.IsPublic

	fieldType, _ := unmarshalChatType(objMap["type"])
	chatInviteLinkInfo.Type = fieldType

	return nil
}

// KeyboardButtonTypeText A simple button, with text that should be sent when the button is pressed
type KeyboardButtonTypeText struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeText
func (keyboardButtonTypeText *KeyboardButtonTypeText) MessageType() string {
	return "keyboardButtonTypeText"
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeText *KeyboardButtonTypeText) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeTextType
}

// KeyboardButtonTypeRequestPhoneNumber A button that sends the user's phone number when pressed; available only in private chats
type KeyboardButtonTypeRequestPhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestPhoneNumber
func (keyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber) MessageType() string {
	return "keyboardButtonTypeRequestPhoneNumber"
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeRequestPhoneNumberType
}

// KeyboardButtonTypeRequestLocation A button that sends the user's location when pressed; available only in private chats
type KeyboardButtonTypeRequestLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestLocation
func (keyboardButtonTypeRequestLocation *KeyboardButtonTypeRequestLocation) MessageType() string {
	return "keyboardButtonTypeRequestLocation"
}

// GetKeyboardButtonTypeEnum return the enum type of this object
func (keyboardButtonTypeRequestLocation *KeyboardButtonTypeRequestLocation) GetKeyboardButtonTypeEnum() KeyboardButtonTypeEnum {
	return KeyboardButtonTypeRequestLocationType
}

// KeyboardButton Represents a single button in a bot keyboard
type KeyboardButton struct {
	tdCommon
	Text string             `json:"text"` // Text of the button
	Type KeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of KeyboardButton
func (keyboardButton *KeyboardButton) MessageType() string {
	return "keyboardButton"
}

// UnmarshalJSON unmarshal to json
func (keyboardButton *KeyboardButton) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Text string `json:"text"` // Text of the button

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	keyboardButton.tdCommon = tempObj.tdCommon
	keyboardButton.Text = tempObj.Text

	fieldType, _ := unmarshalKeyboardButtonType(objMap["type"])
	keyboardButton.Type = fieldType

	return nil
}

// InlineKeyboardButtonTypeURL A button that opens a specified URL
type InlineKeyboardButtonTypeURL struct {
	tdCommon
	URL string `json:"url"` // URL to open
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeURL
func (inlineKeyboardButtonTypeURL *InlineKeyboardButtonTypeURL) MessageType() string {
	return "inlineKeyboardButtonTypeUrl"
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeURL *InlineKeyboardButtonTypeURL) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeURLType
}

// InlineKeyboardButtonTypeCallback A button that sends a special callback query to a bot
type InlineKeyboardButtonTypeCallback struct {
	tdCommon
	Data []byte `json:"data"` // Data to be sent to the bot via a callback query
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallback
func (inlineKeyboardButtonTypeCallback *InlineKeyboardButtonTypeCallback) MessageType() string {
	return "inlineKeyboardButtonTypeCallback"
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeCallback *InlineKeyboardButtonTypeCallback) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeCallbackType
}

// InlineKeyboardButtonTypeCallbackGame A button with a game that sends a special callback query to a bot. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageGame
type InlineKeyboardButtonTypeCallbackGame struct {
	tdCommon
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallbackGame
func (inlineKeyboardButtonTypeCallbackGame *InlineKeyboardButtonTypeCallbackGame) MessageType() string {
	return "inlineKeyboardButtonTypeCallbackGame"
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeCallbackGame *InlineKeyboardButtonTypeCallbackGame) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeCallbackGameType
}

// InlineKeyboardButtonTypeSwitchInline A button that forces an inline query to the bot to be inserted in the input field
type InlineKeyboardButtonTypeSwitchInline struct {
	tdCommon
	Query         string `json:"query"`           // Inline query to be sent to the bot
	InCurrentChat bool   `json:"in_current_chat"` // True, if the inline query should be sent from the current chat
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeSwitchInline
func (inlineKeyboardButtonTypeSwitchInline *InlineKeyboardButtonTypeSwitchInline) MessageType() string {
	return "inlineKeyboardButtonTypeSwitchInline"
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeSwitchInline *InlineKeyboardButtonTypeSwitchInline) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeSwitchInlineType
}

// InlineKeyboardButtonTypeBuy A button to buy something. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageInvoice
type InlineKeyboardButtonTypeBuy struct {
	tdCommon
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeBuy
func (inlineKeyboardButtonTypeBuy *InlineKeyboardButtonTypeBuy) MessageType() string {
	return "inlineKeyboardButtonTypeBuy"
}

// GetInlineKeyboardButtonTypeEnum return the enum type of this object
func (inlineKeyboardButtonTypeBuy *InlineKeyboardButtonTypeBuy) GetInlineKeyboardButtonTypeEnum() InlineKeyboardButtonTypeEnum {
	return InlineKeyboardButtonTypeBuyType
}

// InlineKeyboardButton Represents a single button in an inline keyboard
type InlineKeyboardButton struct {
	tdCommon
	Text string                   `json:"text"` // Text of the button
	Type InlineKeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of InlineKeyboardButton
func (inlineKeyboardButton *InlineKeyboardButton) MessageType() string {
	return "inlineKeyboardButton"
}

// UnmarshalJSON unmarshal to json
func (inlineKeyboardButton *InlineKeyboardButton) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Text string `json:"text"` // Text of the button

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inlineKeyboardButton.tdCommon = tempObj.tdCommon
	inlineKeyboardButton.Text = tempObj.Text

	fieldType, _ := unmarshalInlineKeyboardButtonType(objMap["type"])
	inlineKeyboardButton.Type = fieldType

	return nil
}

// ReplyMarkupRemoveKeyboard Instructs clients to remove the keyboard once this message has been received. This kind of keyboard can't be received in an incoming message; instead, UpdateChatReplyMarkup with message_id == 0 will be sent
type ReplyMarkupRemoveKeyboard struct {
	tdCommon
	IsPersonal bool `json:"is_personal"` // True, if the keyboard is removed only for the mentioned users or the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupRemoveKeyboard
func (replyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard) MessageType() string {
	return "replyMarkupRemoveKeyboard"
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupRemoveKeyboardType
}

// ReplyMarkupForceReply Instructs clients to force a reply to this message
type ReplyMarkupForceReply struct {
	tdCommon
	IsPersonal bool `json:"is_personal"` // True, if a forced reply must automatically be shown to the current user. For outgoing messages, specify true to show the forced reply only for the mentioned users and for the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupForceReply
func (replyMarkupForceReply *ReplyMarkupForceReply) MessageType() string {
	return "replyMarkupForceReply"
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupForceReply *ReplyMarkupForceReply) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupForceReplyType
}

// ReplyMarkupShowKeyboard Contains a custom keyboard layout to quickly reply to bots
type ReplyMarkupShowKeyboard struct {
	tdCommon
	Rows           [][]KeyboardButton `json:"rows"`            // A list of rows of bot keyboard buttons
	ResizeKeyboard bool               `json:"resize_keyboard"` // True, if the client needs to resize the keyboard vertically
	OneTime        bool               `json:"one_time"`        // True, if the client needs to hide the keyboard after use
	IsPersonal     bool               `json:"is_personal"`     // True, if the keyboard must automatically be shown to the current user. For outgoing messages, specify true to show the keyboard only for the mentioned users and for the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupShowKeyboard
func (replyMarkupShowKeyboard *ReplyMarkupShowKeyboard) MessageType() string {
	return "replyMarkupShowKeyboard"
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupShowKeyboard *ReplyMarkupShowKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupShowKeyboardType
}

// ReplyMarkupInlineKeyboard Contains an inline keyboard layout
type ReplyMarkupInlineKeyboard struct {
	tdCommon
	Rows [][]InlineKeyboardButton `json:"rows"` // A list of rows of inline keyboard buttons
}

// MessageType return the string telegram-type of ReplyMarkupInlineKeyboard
func (replyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard) MessageType() string {
	return "replyMarkupInlineKeyboard"
}

// GetReplyMarkupEnum return the enum type of this object
func (replyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard) GetReplyMarkupEnum() ReplyMarkupEnum {
	return ReplyMarkupInlineKeyboardType
}

// RichTextPlain A plain text
type RichTextPlain struct {
	tdCommon
	Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextPlain
func (richTextPlain *RichTextPlain) MessageType() string {
	return "richTextPlain"
}

// GetRichTextEnum return the enum type of this object
func (richTextPlain *RichTextPlain) GetRichTextEnum() RichTextEnum {
	return RichTextPlainType
}

// RichTextBold A bold rich text
type RichTextBold struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextBold
func (richTextBold *RichTextBold) MessageType() string {
	return "richTextBold"
}

// UnmarshalJSON unmarshal to json
func (richTextBold *RichTextBold) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextBold.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextBold.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextBold *RichTextBold) GetRichTextEnum() RichTextEnum {
	return RichTextBoldType
}

// RichTextItalic An italicized rich text
type RichTextItalic struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextItalic
func (richTextItalic *RichTextItalic) MessageType() string {
	return "richTextItalic"
}

// UnmarshalJSON unmarshal to json
func (richTextItalic *RichTextItalic) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextItalic.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextItalic.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextItalic *RichTextItalic) GetRichTextEnum() RichTextEnum {
	return RichTextItalicType
}

// RichTextUnderline An underlined rich text
type RichTextUnderline struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextUnderline
func (richTextUnderline *RichTextUnderline) MessageType() string {
	return "richTextUnderline"
}

// UnmarshalJSON unmarshal to json
func (richTextUnderline *RichTextUnderline) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextUnderline.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextUnderline.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextUnderline *RichTextUnderline) GetRichTextEnum() RichTextEnum {
	return RichTextUnderlineType
}

// RichTextStrikethrough A strike-through rich text
type RichTextStrikethrough struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextStrikethrough
func (richTextStrikethrough *RichTextStrikethrough) MessageType() string {
	return "richTextStrikethrough"
}

// UnmarshalJSON unmarshal to json
func (richTextStrikethrough *RichTextStrikethrough) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextStrikethrough.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextStrikethrough.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextStrikethrough *RichTextStrikethrough) GetRichTextEnum() RichTextEnum {
	return RichTextStrikethroughType
}

// RichTextFixed A fixed-width rich text
type RichTextFixed struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextFixed
func (richTextFixed *RichTextFixed) MessageType() string {
	return "richTextFixed"
}

// UnmarshalJSON unmarshal to json
func (richTextFixed *RichTextFixed) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextFixed.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextFixed.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextFixed *RichTextFixed) GetRichTextEnum() RichTextEnum {
	return RichTextFixedType
}

// RichTextURL A rich text URL link
type RichTextURL struct {
	tdCommon
	Text RichText `json:"text"` // Text
	URL  string   `json:"url"`  // URL
}

// MessageType return the string telegram-type of RichTextURL
func (richTextURL *RichTextURL) MessageType() string {
	return "richTextUrl"
}

// UnmarshalJSON unmarshal to json
func (richTextURL *RichTextURL) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		URL string `json:"url"` // URL
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextURL.tdCommon = tempObj.tdCommon
	richTextURL.URL = tempObj.URL

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextURL.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextURL *RichTextURL) GetRichTextEnum() RichTextEnum {
	return RichTextURLType
}

// RichTextEmailAddress A rich text email link
type RichTextEmailAddress struct {
	tdCommon
	Text         RichText `json:"text"`          // Text
	EmailAddress string   `json:"email_address"` // Email address
}

// MessageType return the string telegram-type of RichTextEmailAddress
func (richTextEmailAddress *RichTextEmailAddress) MessageType() string {
	return "richTextEmailAddress"
}

// UnmarshalJSON unmarshal to json
func (richTextEmailAddress *RichTextEmailAddress) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		EmailAddress string `json:"email_address"` // Email address
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextEmailAddress.tdCommon = tempObj.tdCommon
	richTextEmailAddress.EmailAddress = tempObj.EmailAddress

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextEmailAddress.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextEmailAddress *RichTextEmailAddress) GetRichTextEnum() RichTextEnum {
	return RichTextEmailAddressType
}

// RichTexts A concatenation of rich texts
type RichTexts struct {
	tdCommon
	Texts []RichText `json:"texts"` // Texts
}

// MessageType return the string telegram-type of RichTexts
func (richTexts *RichTexts) MessageType() string {
	return "richTexts"
}

// GetRichTextEnum return the enum type of this object
func (richTexts *RichTexts) GetRichTextEnum() RichTextEnum {
	return RichTextsType
}

// PageBlockTitle The title of a page
type PageBlockTitle struct {
	tdCommon
	Title RichText `json:"title"` // Title
}

// MessageType return the string telegram-type of PageBlockTitle
func (pageBlockTitle *PageBlockTitle) MessageType() string {
	return "pageBlockTitle"
}

// UnmarshalJSON unmarshal to json
func (pageBlockTitle *PageBlockTitle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockTitle.tdCommon = tempObj.tdCommon

	fieldTitle, _ := unmarshalRichText(objMap["title"])
	pageBlockTitle.Title = fieldTitle

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockTitle *PageBlockTitle) GetPageBlockEnum() PageBlockEnum {
	return PageBlockTitleType
}

// PageBlockSubtitle The subtitle of a page
type PageBlockSubtitle struct {
	tdCommon
	Subtitle RichText `json:"subtitle"` // Subtitle
}

// MessageType return the string telegram-type of PageBlockSubtitle
func (pageBlockSubtitle *PageBlockSubtitle) MessageType() string {
	return "pageBlockSubtitle"
}

// UnmarshalJSON unmarshal to json
func (pageBlockSubtitle *PageBlockSubtitle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockSubtitle.tdCommon = tempObj.tdCommon

	fieldSubtitle, _ := unmarshalRichText(objMap["subtitle"])
	pageBlockSubtitle.Subtitle = fieldSubtitle

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSubtitle *PageBlockSubtitle) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSubtitleType
}

// PageBlockAuthorDate The author and publishing date of a page
type PageBlockAuthorDate struct {
	tdCommon
	Author      RichText `json:"author"`       // Author
	PublishDate int32    `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
}

// MessageType return the string telegram-type of PageBlockAuthorDate
func (pageBlockAuthorDate *PageBlockAuthorDate) MessageType() string {
	return "pageBlockAuthorDate"
}

// UnmarshalJSON unmarshal to json
func (pageBlockAuthorDate *PageBlockAuthorDate) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PublishDate int32 `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockAuthorDate.tdCommon = tempObj.tdCommon
	pageBlockAuthorDate.PublishDate = tempObj.PublishDate

	fieldAuthor, _ := unmarshalRichText(objMap["author"])
	pageBlockAuthorDate.Author = fieldAuthor

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAuthorDate *PageBlockAuthorDate) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAuthorDateType
}

// PageBlockHeader A header
type PageBlockHeader struct {
	tdCommon
	Header RichText `json:"header"` // Header
}

// MessageType return the string telegram-type of PageBlockHeader
func (pageBlockHeader *PageBlockHeader) MessageType() string {
	return "pageBlockHeader"
}

// UnmarshalJSON unmarshal to json
func (pageBlockHeader *PageBlockHeader) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockHeader.tdCommon = tempObj.tdCommon

	fieldHeader, _ := unmarshalRichText(objMap["header"])
	pageBlockHeader.Header = fieldHeader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockHeader *PageBlockHeader) GetPageBlockEnum() PageBlockEnum {
	return PageBlockHeaderType
}

// PageBlockSubheader A subheader
type PageBlockSubheader struct {
	tdCommon
	Subheader RichText `json:"subheader"` // Subheader
}

// MessageType return the string telegram-type of PageBlockSubheader
func (pageBlockSubheader *PageBlockSubheader) MessageType() string {
	return "pageBlockSubheader"
}

// UnmarshalJSON unmarshal to json
func (pageBlockSubheader *PageBlockSubheader) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockSubheader.tdCommon = tempObj.tdCommon

	fieldSubheader, _ := unmarshalRichText(objMap["subheader"])
	pageBlockSubheader.Subheader = fieldSubheader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSubheader *PageBlockSubheader) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSubheaderType
}

// PageBlockParagraph A text paragraph
type PageBlockParagraph struct {
	tdCommon
	Text RichText `json:"text"` // Paragraph text
}

// MessageType return the string telegram-type of PageBlockParagraph
func (pageBlockParagraph *PageBlockParagraph) MessageType() string {
	return "pageBlockParagraph"
}

// UnmarshalJSON unmarshal to json
func (pageBlockParagraph *PageBlockParagraph) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockParagraph.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockParagraph.Text = fieldText

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockParagraph *PageBlockParagraph) GetPageBlockEnum() PageBlockEnum {
	return PageBlockParagraphType
}

// PageBlockPreformatted A preformatted text paragraph
type PageBlockPreformatted struct {
	tdCommon
	Text     RichText `json:"text"`     // Paragraph text
	Language string   `json:"language"` // Programming language for which the text should be formatted
}

// MessageType return the string telegram-type of PageBlockPreformatted
func (pageBlockPreformatted *PageBlockPreformatted) MessageType() string {
	return "pageBlockPreformatted"
}

// UnmarshalJSON unmarshal to json
func (pageBlockPreformatted *PageBlockPreformatted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Language string `json:"language"` // Programming language for which the text should be formatted
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockPreformatted.tdCommon = tempObj.tdCommon
	pageBlockPreformatted.Language = tempObj.Language

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockPreformatted.Text = fieldText

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPreformatted *PageBlockPreformatted) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPreformattedType
}

// PageBlockFooter The footer of a page
type PageBlockFooter struct {
	tdCommon
	Footer RichText `json:"footer"` // Footer
}

// MessageType return the string telegram-type of PageBlockFooter
func (pageBlockFooter *PageBlockFooter) MessageType() string {
	return "pageBlockFooter"
}

// UnmarshalJSON unmarshal to json
func (pageBlockFooter *PageBlockFooter) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockFooter.tdCommon = tempObj.tdCommon

	fieldFooter, _ := unmarshalRichText(objMap["footer"])
	pageBlockFooter.Footer = fieldFooter

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockFooter *PageBlockFooter) GetPageBlockEnum() PageBlockEnum {
	return PageBlockFooterType
}

// PageBlockDivider An empty block separating a page
type PageBlockDivider struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockDivider
func (pageBlockDivider *PageBlockDivider) MessageType() string {
	return "pageBlockDivider"
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockDivider *PageBlockDivider) GetPageBlockEnum() PageBlockEnum {
	return PageBlockDividerType
}

// PageBlockAnchor An invisible anchor on a page, which can be used in a URL to open the page from the specified anchor
type PageBlockAnchor struct {
	tdCommon
	Name string `json:"name"` // Name of the anchor
}

// MessageType return the string telegram-type of PageBlockAnchor
func (pageBlockAnchor *PageBlockAnchor) MessageType() string {
	return "pageBlockAnchor"
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAnchor *PageBlockAnchor) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAnchorType
}

// PageBlockList A list of texts
type PageBlockList struct {
	tdCommon
	Items     []RichText `json:"items"`      // Texts
	IsOrdered bool       `json:"is_ordered"` // True, if the items should be marked with numbers
}

// MessageType return the string telegram-type of PageBlockList
func (pageBlockList *PageBlockList) MessageType() string {
	return "pageBlockList"
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockList *PageBlockList) GetPageBlockEnum() PageBlockEnum {
	return PageBlockListType
}

// PageBlockBlockQuote A block quote
type PageBlockBlockQuote struct {
	tdCommon
	Text    RichText `json:"text"`    // Quote text
	Caption RichText `json:"caption"` // Quote caption
}

// MessageType return the string telegram-type of PageBlockBlockQuote
func (pageBlockBlockQuote *PageBlockBlockQuote) MessageType() string {
	return "pageBlockBlockQuote"
}

// UnmarshalJSON unmarshal to json
func (pageBlockBlockQuote *PageBlockBlockQuote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockBlockQuote.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockBlockQuote.Text = fieldText

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockBlockQuote.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockBlockQuote *PageBlockBlockQuote) GetPageBlockEnum() PageBlockEnum {
	return PageBlockBlockQuoteType
}

// PageBlockPullQuote A pull quote
type PageBlockPullQuote struct {
	tdCommon
	Text    RichText `json:"text"`    // Quote text
	Caption RichText `json:"caption"` // Quote caption
}

// MessageType return the string telegram-type of PageBlockPullQuote
func (pageBlockPullQuote *PageBlockPullQuote) MessageType() string {
	return "pageBlockPullQuote"
}

// UnmarshalJSON unmarshal to json
func (pageBlockPullQuote *PageBlockPullQuote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockPullQuote.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockPullQuote.Text = fieldText

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockPullQuote.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPullQuote *PageBlockPullQuote) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPullQuoteType
}

// PageBlockAnimation An animation
type PageBlockAnimation struct {
	tdCommon
	Animation    Animation `json:"animation"`     // Animation file; may be null
	Caption      RichText  `json:"caption"`       // Animation caption
	NeedAutoplay bool      `json:"need_autoplay"` // True, if the animation should be played automatically
}

// MessageType return the string telegram-type of PageBlockAnimation
func (pageBlockAnimation *PageBlockAnimation) MessageType() string {
	return "pageBlockAnimation"
}

// UnmarshalJSON unmarshal to json
func (pageBlockAnimation *PageBlockAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Animation    Animation `json:"animation"`     // Animation file; may be null
		NeedAutoplay bool      `json:"need_autoplay"` // True, if the animation should be played automatically
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockAnimation.tdCommon = tempObj.tdCommon
	pageBlockAnimation.Animation = tempObj.Animation
	pageBlockAnimation.NeedAutoplay = tempObj.NeedAutoplay

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockAnimation.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAnimation *PageBlockAnimation) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAnimationType
}

// PageBlockAudio An audio file
type PageBlockAudio struct {
	tdCommon
	Audio   Audio    `json:"audio"`   // Audio file; may be null
	Caption RichText `json:"caption"` // Audio file caption
}

// MessageType return the string telegram-type of PageBlockAudio
func (pageBlockAudio *PageBlockAudio) MessageType() string {
	return "pageBlockAudio"
}

// UnmarshalJSON unmarshal to json
func (pageBlockAudio *PageBlockAudio) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Audio Audio `json:"audio"` // Audio file; may be null

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockAudio.tdCommon = tempObj.tdCommon
	pageBlockAudio.Audio = tempObj.Audio

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockAudio.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAudio *PageBlockAudio) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAudioType
}

// PageBlockPhoto A photo
type PageBlockPhoto struct {
	tdCommon
	Photo   Photo    `json:"photo"`   // Photo file; may be null
	Caption RichText `json:"caption"` // Photo caption
}

// MessageType return the string telegram-type of PageBlockPhoto
func (pageBlockPhoto *PageBlockPhoto) MessageType() string {
	return "pageBlockPhoto"
}

// UnmarshalJSON unmarshal to json
func (pageBlockPhoto *PageBlockPhoto) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Photo Photo `json:"photo"` // Photo file; may be null

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockPhoto.tdCommon = tempObj.tdCommon
	pageBlockPhoto.Photo = tempObj.Photo

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockPhoto.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPhoto *PageBlockPhoto) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPhotoType
}

// PageBlockVideo A video
type PageBlockVideo struct {
	tdCommon
	Video        Video    `json:"video"`         // Video file; may be null
	Caption      RichText `json:"caption"`       // Video caption
	NeedAutoplay bool     `json:"need_autoplay"` // True, if the video should be played automatically
	IsLooped     bool     `json:"is_looped"`     // True, if the video should be looped
}

// MessageType return the string telegram-type of PageBlockVideo
func (pageBlockVideo *PageBlockVideo) MessageType() string {
	return "pageBlockVideo"
}

// UnmarshalJSON unmarshal to json
func (pageBlockVideo *PageBlockVideo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Video        Video `json:"video"`         // Video file; may be null
		NeedAutoplay bool  `json:"need_autoplay"` // True, if the video should be played automatically
		IsLooped     bool  `json:"is_looped"`     // True, if the video should be looped
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockVideo.tdCommon = tempObj.tdCommon
	pageBlockVideo.Video = tempObj.Video
	pageBlockVideo.NeedAutoplay = tempObj.NeedAutoplay
	pageBlockVideo.IsLooped = tempObj.IsLooped

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockVideo.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockVideo *PageBlockVideo) GetPageBlockEnum() PageBlockEnum {
	return PageBlockVideoType
}

// PageBlockCover A page cover
type PageBlockCover struct {
	tdCommon
	Cover PageBlock `json:"cover"` // Cover
}

// MessageType return the string telegram-type of PageBlockCover
func (pageBlockCover *PageBlockCover) MessageType() string {
	return "pageBlockCover"
}

// UnmarshalJSON unmarshal to json
func (pageBlockCover *PageBlockCover) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockCover.tdCommon = tempObj.tdCommon

	fieldCover, _ := unmarshalPageBlock(objMap["cover"])
	pageBlockCover.Cover = fieldCover

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockCover *PageBlockCover) GetPageBlockEnum() PageBlockEnum {
	return PageBlockCoverType
}

// PageBlockEmbedded An embedded web page
type PageBlockEmbedded struct {
	tdCommon
	URL            string   `json:"url"`             // Web page URL, if available
	HTML           string   `json:"html"`            // HTML-markup of the embedded page
	PosterPhoto    Photo    `json:"poster_photo"`    // Poster photo, if available; may be null
	Width          int32    `json:"width"`           // Block width
	Height         int32    `json:"height"`          // Block height
	Caption        RichText `json:"caption"`         // Block caption
	IsFullWidth    bool     `json:"is_full_width"`   // True, if the block should be full width
	AllowScrolling bool     `json:"allow_scrolling"` // True, if scrolling should be allowed
}

// MessageType return the string telegram-type of PageBlockEmbedded
func (pageBlockEmbedded *PageBlockEmbedded) MessageType() string {
	return "pageBlockEmbedded"
}

// UnmarshalJSON unmarshal to json
func (pageBlockEmbedded *PageBlockEmbedded) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		URL            string `json:"url"`             // Web page URL, if available
		HTML           string `json:"html"`            // HTML-markup of the embedded page
		PosterPhoto    Photo  `json:"poster_photo"`    // Poster photo, if available; may be null
		Width          int32  `json:"width"`           // Block width
		Height         int32  `json:"height"`          // Block height
		IsFullWidth    bool   `json:"is_full_width"`   // True, if the block should be full width
		AllowScrolling bool   `json:"allow_scrolling"` // True, if scrolling should be allowed
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockEmbedded.tdCommon = tempObj.tdCommon
	pageBlockEmbedded.URL = tempObj.URL
	pageBlockEmbedded.HTML = tempObj.HTML
	pageBlockEmbedded.PosterPhoto = tempObj.PosterPhoto
	pageBlockEmbedded.Width = tempObj.Width
	pageBlockEmbedded.Height = tempObj.Height
	pageBlockEmbedded.IsFullWidth = tempObj.IsFullWidth
	pageBlockEmbedded.AllowScrolling = tempObj.AllowScrolling

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockEmbedded.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockEmbedded *PageBlockEmbedded) GetPageBlockEnum() PageBlockEnum {
	return PageBlockEmbeddedType
}

// PageBlockEmbeddedPost An embedded post
type PageBlockEmbeddedPost struct {
	tdCommon
	URL         string      `json:"url"`          // Web page URL
	Author      string      `json:"author"`       // Post author
	AuthorPhoto Photo       `json:"author_photo"` // Post author photo
	Date        int32       `json:"date"`         // Point in time (Unix timestamp) when the post was created; 0 if unknown
	PageBlocks  []PageBlock `json:"page_blocks"`  // Post content
	Caption     RichText    `json:"caption"`      // Post caption
}

// MessageType return the string telegram-type of PageBlockEmbeddedPost
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) MessageType() string {
	return "pageBlockEmbeddedPost"
}

// UnmarshalJSON unmarshal to json
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		URL         string      `json:"url"`          // Web page URL
		Author      string      `json:"author"`       // Post author
		AuthorPhoto Photo       `json:"author_photo"` // Post author photo
		Date        int32       `json:"date"`         // Point in time (Unix timestamp) when the post was created; 0 if unknown
		PageBlocks  []PageBlock `json:"page_blocks"`  // Post content

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockEmbeddedPost.tdCommon = tempObj.tdCommon
	pageBlockEmbeddedPost.URL = tempObj.URL
	pageBlockEmbeddedPost.Author = tempObj.Author
	pageBlockEmbeddedPost.AuthorPhoto = tempObj.AuthorPhoto
	pageBlockEmbeddedPost.Date = tempObj.Date
	pageBlockEmbeddedPost.PageBlocks = tempObj.PageBlocks

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockEmbeddedPost.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) GetPageBlockEnum() PageBlockEnum {
	return PageBlockEmbeddedPostType
}

// PageBlockCollage A collage
type PageBlockCollage struct {
	tdCommon
	PageBlocks []PageBlock `json:"page_blocks"` // Collage item contents
	Caption    RichText    `json:"caption"`     // Block caption
}

// MessageType return the string telegram-type of PageBlockCollage
func (pageBlockCollage *PageBlockCollage) MessageType() string {
	return "pageBlockCollage"
}

// UnmarshalJSON unmarshal to json
func (pageBlockCollage *PageBlockCollage) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PageBlocks []PageBlock `json:"page_blocks"` // Collage item contents

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockCollage.tdCommon = tempObj.tdCommon
	pageBlockCollage.PageBlocks = tempObj.PageBlocks

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockCollage.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockCollage *PageBlockCollage) GetPageBlockEnum() PageBlockEnum {
	return PageBlockCollageType
}

// PageBlockSlideshow A slideshow
type PageBlockSlideshow struct {
	tdCommon
	PageBlocks []PageBlock `json:"page_blocks"` // Slideshow item contents
	Caption    RichText    `json:"caption"`     // Block caption
}

// MessageType return the string telegram-type of PageBlockSlideshow
func (pageBlockSlideshow *PageBlockSlideshow) MessageType() string {
	return "pageBlockSlideshow"
}

// UnmarshalJSON unmarshal to json
func (pageBlockSlideshow *PageBlockSlideshow) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PageBlocks []PageBlock `json:"page_blocks"` // Slideshow item contents

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockSlideshow.tdCommon = tempObj.tdCommon
	pageBlockSlideshow.PageBlocks = tempObj.PageBlocks

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockSlideshow.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSlideshow *PageBlockSlideshow) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSlideshowType
}

// PageBlockChatLink A link to a chat
type PageBlockChatLink struct {
	tdCommon
	Title    string    `json:"title"`    // Chat title
	Photo    ChatPhoto `json:"photo"`    // Chat photo; may be null
	Username string    `json:"username"` // Chat username, by which all other information about the chat should be resolved
}

// MessageType return the string telegram-type of PageBlockChatLink
func (pageBlockChatLink *PageBlockChatLink) MessageType() string {
	return "pageBlockChatLink"
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockChatLink *PageBlockChatLink) GetPageBlockEnum() PageBlockEnum {
	return PageBlockChatLinkType
}

// WebPageInstantView Describes an instant view page for a web page
type WebPageInstantView struct {
	tdCommon
	PageBlocks []PageBlock `json:"page_blocks"` // Content of the web page
	IsFull     bool        `json:"is_full"`     // True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
}

// MessageType return the string telegram-type of WebPageInstantView
func (webPageInstantView *WebPageInstantView) MessageType() string {
	return "webPageInstantView"
}

// WebPage Describes a web page preview
type WebPage struct {
	tdCommon
	URL            string    `json:"url"`              // Original URL of the link
	DisplayURL     string    `json:"display_url"`      // URL to display
	Type           string    `json:"type"`             // Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
	SiteName       string    `json:"site_name"`        // Short name of the site (e.g., Google Docs, App Store)
	Title          string    `json:"title"`            // Title of the content
	Description    string    `json:"description"`      //
	Photo          Photo     `json:"photo"`            // Image representing the content; may be null
	EmbedURL       string    `json:"embed_url"`        // URL to show in the embedded preview
	EmbedType      string    `json:"embed_type"`       // MIME type of the embedded preview, (e.g., text/html or video/mp4)
	EmbedWidth     int32     `json:"embed_width"`      // Width of the embedded preview
	EmbedHeight    int32     `json:"embed_height"`     // Height of the embedded preview
	Duration       int32     `json:"duration"`         // Duration of the content, in seconds
	Author         string    `json:"author"`           // Author of the content
	Animation      Animation `json:"animation"`        // Preview of the content as an animation, if available; may be null
	Audio          Audio     `json:"audio"`            // Preview of the content as an audio file, if available; may be null
	Document       Document  `json:"document"`         // Preview of the content as a document, if available (currently only available for small PDF files and ZIP archives); may be null
	Sticker        Sticker   `json:"sticker"`          // Preview of the content as a sticker for small WEBP files, if available; may be null
	Video          Video     `json:"video"`            // Preview of the content as a video, if available; may be null
	VideoNote      VideoNote `json:"video_note"`       // Preview of the content as a video note, if available; may be null
	VoiceNote      VoiceNote `json:"voice_note"`       // Preview of the content as a voice note, if available; may be null
	HasInstantView bool      `json:"has_instant_view"` // True, if the web page has an instant view
}

// MessageType return the string telegram-type of WebPage
func (webPage *WebPage) MessageType() string {
	return "webPage"
}

// LabeledPricePart Portion of the price of a product (e.g., "delivery cost", "tax amount")
type LabeledPricePart struct {
	tdCommon
	Label  string `json:"label"`  // Label for this portion of the product price
	Amount int64  `json:"amount"` // Currency amount in minimal quantity of the currency
}

// MessageType return the string telegram-type of LabeledPricePart
func (labeledPricePart *LabeledPricePart) MessageType() string {
	return "labeledPricePart"
}

// Invoice Product invoice
type Invoice struct {
	tdCommon
	Currency                   string             `json:"currency"`                       // ISO 4217 currency code
	PriceParts                 []LabeledPricePart `json:"price_parts"`                    // A list of objects used to calculate the total price of the product
	IsTest                     bool               `json:"is_test"`                        // True, if the payment is a test payment
	NeedName                   bool               `json:"need_name"`                      // True, if the user's name is needed for payment
	NeedPhoneNumber            bool               `json:"need_phone_number"`              // True, if the user's phone number is needed for payment
	NeedEmailAddress           bool               `json:"need_email_address"`             // True, if the user's email address is needed for payment
	NeedShippingAddress        bool               `json:"need_shipping_address"`          // True, if the user's shipping address is needed for payment
	SendPhoneNumberToProvider  bool               `json:"send_phone_number_to_provider"`  // True, if the user's phone number will be sent to the provider
	SendEmailAddressToProvider bool               `json:"send_email_address_to_provider"` // True, if the user's email address will be sent to the provider
	IsFlexible                 bool               `json:"is_flexible"`                    // True, if the total price depends on the shipping method
}

// MessageType return the string telegram-type of Invoice
func (invoice *Invoice) MessageType() string {
	return "invoice"
}

// ShippingAddress Describes a shipping address
type ShippingAddress struct {
	tdCommon
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line of the address
	StreetLine2 string `json:"street_line2"` // Second line of the address
	PostalCode  string `json:"postal_code"`  // Address postal code
}

// MessageType return the string telegram-type of ShippingAddress
func (shippingAddress *ShippingAddress) MessageType() string {
	return "shippingAddress"
}

// OrderInfo Order information
type OrderInfo struct {
	tdCommon
	Name            string          `json:"name"`             // Name of the user
	PhoneNumber     string          `json:"phone_number"`     // Phone number of the user
	EmailAddress    string          `json:"email_address"`    // Email address of the user
	ShippingAddress ShippingAddress `json:"shipping_address"` // Shipping address for this order; may be null
}

// MessageType return the string telegram-type of OrderInfo
func (orderInfo *OrderInfo) MessageType() string {
	return "orderInfo"
}

// ShippingOption One shipping option
type ShippingOption struct {
	tdCommon
	ID         string             `json:"id"`          // Shipping option identifier
	Title      string             `json:"title"`       // Option title
	PriceParts []LabeledPricePart `json:"price_parts"` // A list of objects used to calculate the total shipping costs
}

// MessageType return the string telegram-type of ShippingOption
func (shippingOption *ShippingOption) MessageType() string {
	return "shippingOption"
}

// SavedCredentials Contains information about saved card credentials
type SavedCredentials struct {
	tdCommon
	ID    string `json:"id"`    // Unique identifier of the saved credentials
	Title string `json:"title"` // Title of the saved credentials
}

// MessageType return the string telegram-type of SavedCredentials
func (savedCredentials *SavedCredentials) MessageType() string {
	return "savedCredentials"
}

// InputCredentialsSaved Applies if a user chooses some previously saved payment credentials. To use their previously saved credentials, the user must have a valid temporary password
type InputCredentialsSaved struct {
	tdCommon
	SavedCredentialsID string `json:"saved_credentials_id"` // Identifier of the saved credentials
}

// MessageType return the string telegram-type of InputCredentialsSaved
func (inputCredentialsSaved *InputCredentialsSaved) MessageType() string {
	return "inputCredentialsSaved"
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsSaved *InputCredentialsSaved) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsSavedType
}

// InputCredentialsNew Applies if a user enters new credentials on a payment provider website
type InputCredentialsNew struct {
	tdCommon
	Data      string `json:"data"`       // Contains JSON-encoded data with a credential identifier from the payment provider
	AllowSave bool   `json:"allow_save"` // True, if the credential identifier can be saved on the server side
}

// MessageType return the string telegram-type of InputCredentialsNew
func (inputCredentialsNew *InputCredentialsNew) MessageType() string {
	return "inputCredentialsNew"
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsNew *InputCredentialsNew) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsNewType
}

// InputCredentialsAndroidPay Applies if a user enters new credentials using Android Pay
type InputCredentialsAndroidPay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsAndroidPay
func (inputCredentialsAndroidPay *InputCredentialsAndroidPay) MessageType() string {
	return "inputCredentialsAndroidPay"
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsAndroidPay *InputCredentialsAndroidPay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsAndroidPayType
}

// InputCredentialsApplePay Applies if a user enters new credentials using Apple Pay
type InputCredentialsApplePay struct {
	tdCommon
	Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsApplePay
func (inputCredentialsApplePay *InputCredentialsApplePay) MessageType() string {
	return "inputCredentialsApplePay"
}

// GetInputCredentialsEnum return the enum type of this object
func (inputCredentialsApplePay *InputCredentialsApplePay) GetInputCredentialsEnum() InputCredentialsEnum {
	return InputCredentialsApplePayType
}

// PaymentsProviderStripe Stripe payment provider
type PaymentsProviderStripe struct {
	tdCommon
	PublishableKey     string `json:"publishable_key"`      // Stripe API publishable key
	NeedCountry        bool   `json:"need_country"`         // True, if the user country must be provided
	NeedPostalCode     bool   `json:"need_postal_code"`     // True, if the user ZIP/postal code must be provided
	NeedCardholderName bool   `json:"need_cardholder_name"` // True, if the cardholder name must be provided
}

// MessageType return the string telegram-type of PaymentsProviderStripe
func (paymentsProviderStripe *PaymentsProviderStripe) MessageType() string {
	return "paymentsProviderStripe"
}

// PaymentForm Contains information about an invoice payment form
type PaymentForm struct {
	tdCommon
	Invoice            Invoice                `json:"invoice"`              // Full information of the invoice
	URL                string                 `json:"url"`                  // Payment form URL
	PaymentsProvider   PaymentsProviderStripe `json:"payments_provider"`    // Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
	SavedOrderInfo     OrderInfo              `json:"saved_order_info"`     // Saved server-side order information; may be null
	SavedCredentials   SavedCredentials       `json:"saved_credentials"`    // Contains information about saved card credentials; may be null
	CanSaveCredentials bool                   `json:"can_save_credentials"` // True, if the user can choose to save credentials
	NeedPassword       bool                   `json:"need_password"`        // True, if the user will be able to save credentials protected by a password they set up
}

// MessageType return the string telegram-type of PaymentForm
func (paymentForm *PaymentForm) MessageType() string {
	return "paymentForm"
}

// ValidatedOrderInfo Contains a temporary identifier of validated order information, which is stored for one hour. Also contains the available shipping options
type ValidatedOrderInfo struct {
	tdCommon
	OrderInfoID     string           `json:"order_info_id"`    // Temporary identifier of the order information
	ShippingOptions []ShippingOption `json:"shipping_options"` // Available shipping options
}

// MessageType return the string telegram-type of ValidatedOrderInfo
func (validatedOrderInfo *ValidatedOrderInfo) MessageType() string {
	return "validatedOrderInfo"
}

// PaymentResult Contains the result of a payment request
type PaymentResult struct {
	tdCommon
	Success         bool   `json:"success"`          // True, if the payment request was successful; otherwise the verification_url will be not empty
	VerificationURL string `json:"verification_url"` // URL for additional payment credentials verification
}

// MessageType return the string telegram-type of PaymentResult
func (paymentResult *PaymentResult) MessageType() string {
	return "paymentResult"
}

// PaymentReceipt Contains information about a successful payment
type PaymentReceipt struct {
	tdCommon
	Date                   int32          `json:"date"`                      // Point in time (Unix timestamp) when the payment was made
	PaymentsProviderUserID int32          `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	Invoice                Invoice        `json:"invoice"`                   // Contains information about the invoice
	OrderInfo              OrderInfo      `json:"order_info"`                // Contains order information; may be null
	ShippingOption         ShippingOption `json:"shipping_option"`           // Chosen shipping option; may be null
	CredentialsTitle       string         `json:"credentials_title"`         // Title of the saved credentials
}

// MessageType return the string telegram-type of PaymentReceipt
func (paymentReceipt *PaymentReceipt) MessageType() string {
	return "paymentReceipt"
}

// MessageText A text message
type MessageText struct {
	tdCommon
	Text    FormattedText `json:"text"`     // Text of the message
	WebPage WebPage       `json:"web_page"` // A preview of the web page that's mentioned in the text; may be null
}

// MessageType return the string telegram-type of MessageText
func (messageText *MessageText) MessageType() string {
	return "messageText"
}

// GetMessageContentEnum return the enum type of this object
func (messageText *MessageText) GetMessageContentEnum() MessageContentEnum {
	return MessageTextType
}

// MessageAnimation An animation message (GIF-style).
type MessageAnimation struct {
	tdCommon
	Animation Animation     `json:"animation"` // Message content
	Caption   FormattedText `json:"caption"`   // Animation caption
	IsSecret  bool          `json:"is_secret"` // True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
}

// MessageType return the string telegram-type of MessageAnimation
func (messageAnimation *MessageAnimation) MessageType() string {
	return "messageAnimation"
}

// GetMessageContentEnum return the enum type of this object
func (messageAnimation *MessageAnimation) GetMessageContentEnum() MessageContentEnum {
	return MessageAnimationType
}

// MessageAudio An audio message
type MessageAudio struct {
	tdCommon
	Audio   Audio         `json:"audio"`   // Message content
	Caption FormattedText `json:"caption"` // Audio caption
}

// MessageType return the string telegram-type of MessageAudio
func (messageAudio *MessageAudio) MessageType() string {
	return "messageAudio"
}

// GetMessageContentEnum return the enum type of this object
func (messageAudio *MessageAudio) GetMessageContentEnum() MessageContentEnum {
	return MessageAudioType
}

// MessageDocument A document message (general file)
type MessageDocument struct {
	tdCommon
	Document Document      `json:"document"` // Message content
	Caption  FormattedText `json:"caption"`  // Document caption
}

// MessageType return the string telegram-type of MessageDocument
func (messageDocument *MessageDocument) MessageType() string {
	return "messageDocument"
}

// GetMessageContentEnum return the enum type of this object
func (messageDocument *MessageDocument) GetMessageContentEnum() MessageContentEnum {
	return MessageDocumentType
}

// MessagePhoto A photo message
type MessagePhoto struct {
	tdCommon
	Photo    Photo         `json:"photo"`     // Message content
	Caption  FormattedText `json:"caption"`   // Photo caption
	IsSecret bool          `json:"is_secret"` // True, if the photo must be blurred and must be shown only while tapped
}

// MessageType return the string telegram-type of MessagePhoto
func (messagePhoto *MessagePhoto) MessageType() string {
	return "messagePhoto"
}

// GetMessageContentEnum return the enum type of this object
func (messagePhoto *MessagePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessagePhotoType
}

// MessageExpiredPhoto An expired photo message (self-destructed after TTL has elapsed)
type MessageExpiredPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageExpiredPhoto
func (messageExpiredPhoto *MessageExpiredPhoto) MessageType() string {
	return "messageExpiredPhoto"
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredPhoto *MessageExpiredPhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredPhotoType
}

// MessageSticker A sticker message
type MessageSticker struct {
	tdCommon
	Sticker Sticker `json:"sticker"` // Message content
}

// MessageType return the string telegram-type of MessageSticker
func (messageSticker *MessageSticker) MessageType() string {
	return "messageSticker"
}

// GetMessageContentEnum return the enum type of this object
func (messageSticker *MessageSticker) GetMessageContentEnum() MessageContentEnum {
	return MessageStickerType
}

// MessageVideo A video message
type MessageVideo struct {
	tdCommon
	Video    Video         `json:"video"`     // Message content
	Caption  FormattedText `json:"caption"`   // Video caption
	IsSecret bool          `json:"is_secret"` // True, if the video thumbnail must be blurred and the video must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideo
func (messageVideo *MessageVideo) MessageType() string {
	return "messageVideo"
}

// GetMessageContentEnum return the enum type of this object
func (messageVideo *MessageVideo) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoType
}

// MessageExpiredVideo An expired video message (self-destructed after TTL has elapsed)
type MessageExpiredVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageExpiredVideo
func (messageExpiredVideo *MessageExpiredVideo) MessageType() string {
	return "messageExpiredVideo"
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredVideo *MessageExpiredVideo) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredVideoType
}

// MessageVideoNote A video note message
type MessageVideoNote struct {
	tdCommon
	VideoNote VideoNote `json:"video_note"` // Message content
	IsViewed  bool      `json:"is_viewed"`  // True, if at least one of the recipients has viewed the video note
	IsSecret  bool      `json:"is_secret"`  // True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideoNote
func (messageVideoNote *MessageVideoNote) MessageType() string {
	return "messageVideoNote"
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoNote *MessageVideoNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoNoteType
}

// MessageVoiceNote A voice note message
type MessageVoiceNote struct {
	tdCommon
	VoiceNote  VoiceNote     `json:"voice_note"`  // Message content
	Caption    FormattedText `json:"caption"`     // Voice note caption
	IsListened bool          `json:"is_listened"` // True, if at least one of the recipients has listened to the voice note
}

// MessageType return the string telegram-type of MessageVoiceNote
func (messageVoiceNote *MessageVoiceNote) MessageType() string {
	return "messageVoiceNote"
}

// GetMessageContentEnum return the enum type of this object
func (messageVoiceNote *MessageVoiceNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVoiceNoteType
}

// MessageLocation A message with a location
type MessageLocation struct {
	tdCommon
	Location   Location `json:"location"`    // Message content
	LivePeriod int32    `json:"live_period"` // Time relative to the message sent date until which the location can be updated, in seconds
	ExpiresIn  int32    `json:"expires_in"`  // Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
}

// MessageType return the string telegram-type of MessageLocation
func (messageLocation *MessageLocation) MessageType() string {
	return "messageLocation"
}

// GetMessageContentEnum return the enum type of this object
func (messageLocation *MessageLocation) GetMessageContentEnum() MessageContentEnum {
	return MessageLocationType
}

// MessageVenue A message with information about a venue
type MessageVenue struct {
	tdCommon
	Venue Venue `json:"venue"` // Message content
}

// MessageType return the string telegram-type of MessageVenue
func (messageVenue *MessageVenue) MessageType() string {
	return "messageVenue"
}

// GetMessageContentEnum return the enum type of this object
func (messageVenue *MessageVenue) GetMessageContentEnum() MessageContentEnum {
	return MessageVenueType
}

// MessageContact A message with a user contact
type MessageContact struct {
	tdCommon
	Contact Contact `json:"contact"` // Message content
}

// MessageType return the string telegram-type of MessageContact
func (messageContact *MessageContact) MessageType() string {
	return "messageContact"
}

// GetMessageContentEnum return the enum type of this object
func (messageContact *MessageContact) GetMessageContentEnum() MessageContentEnum {
	return MessageContactType
}

// MessageGame A message with a game
type MessageGame struct {
	tdCommon
	Game Game `json:"game"` // Game
}

// MessageType return the string telegram-type of MessageGame
func (messageGame *MessageGame) MessageType() string {
	return "messageGame"
}

// GetMessageContentEnum return the enum type of this object
func (messageGame *MessageGame) GetMessageContentEnum() MessageContentEnum {
	return MessageGameType
}

// MessageInvoice A message with an invoice from a bot
type MessageInvoice struct {
	tdCommon
	Title               string `json:"title"`                 // Product title
	Description         string `json:"description"`           //
	Photo               Photo  `json:"photo"`                 // Product photo; may be null
	Currency            string `json:"currency"`              // Currency for the product price
	TotalAmount         int64  `json:"total_amount"`          // Product total price in the minimal quantity of the currency
	StartParameter      string `json:"start_parameter"`       // Unique invoice bot start_parameter. To share an invoice use the URL https://t.me/{bot_username}?start={start_parameter}
	IsTest              bool   `json:"is_test"`               // True, if the invoice is a test invoice
	NeedShippingAddress bool   `json:"need_shipping_address"` // True, if the shipping address should be specified
	ReceiptMessageID    int64  `json:"receipt_message_id"`    // The identifier of the message with the receipt, after the product has been purchased
}

// MessageType return the string telegram-type of MessageInvoice
func (messageInvoice *MessageInvoice) MessageType() string {
	return "messageInvoice"
}

// GetMessageContentEnum return the enum type of this object
func (messageInvoice *MessageInvoice) GetMessageContentEnum() MessageContentEnum {
	return MessageInvoiceType
}

// MessageCall A message with information about an ended call
type MessageCall struct {
	tdCommon
	DiscardReason CallDiscardReason `json:"discard_reason"` // Reason why the call was discarded
	Duration      int32             `json:"duration"`       // Call duration, in seconds
}

// MessageType return the string telegram-type of MessageCall
func (messageCall *MessageCall) MessageType() string {
	return "messageCall"
}

// UnmarshalJSON unmarshal to json
func (messageCall *MessageCall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Duration int32 `json:"duration"` // Call duration, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageCall.tdCommon = tempObj.tdCommon
	messageCall.Duration = tempObj.Duration

	fieldDiscardReason, _ := unmarshalCallDiscardReason(objMap["discard_reason"])
	messageCall.DiscardReason = fieldDiscardReason

	return nil
}

// GetMessageContentEnum return the enum type of this object
func (messageCall *MessageCall) GetMessageContentEnum() MessageContentEnum {
	return MessageCallType
}

// MessageBasicGroupChatCreate A newly created basic group
type MessageBasicGroupChatCreate struct {
	tdCommon
	Title         string  `json:"title"`           // Title of the basic group
	MemberUserIDs []int32 `json:"member_user_ids"` // User identifiers of members in the basic group
}

// MessageType return the string telegram-type of MessageBasicGroupChatCreate
func (messageBasicGroupChatCreate *MessageBasicGroupChatCreate) MessageType() string {
	return "messageBasicGroupChatCreate"
}

// GetMessageContentEnum return the enum type of this object
func (messageBasicGroupChatCreate *MessageBasicGroupChatCreate) GetMessageContentEnum() MessageContentEnum {
	return MessageBasicGroupChatCreateType
}

// MessageSupergroupChatCreate A newly created supergroup or channel
type MessageSupergroupChatCreate struct {
	tdCommon
	Title string `json:"title"` // Title of the supergroup or channel
}

// MessageType return the string telegram-type of MessageSupergroupChatCreate
func (messageSupergroupChatCreate *MessageSupergroupChatCreate) MessageType() string {
	return "messageSupergroupChatCreate"
}

// GetMessageContentEnum return the enum type of this object
func (messageSupergroupChatCreate *MessageSupergroupChatCreate) GetMessageContentEnum() MessageContentEnum {
	return MessageSupergroupChatCreateType
}

// MessageChatChangeTitle An updated chat title
type MessageChatChangeTitle struct {
	tdCommon
	Title string `json:"title"` // New chat title
}

// MessageType return the string telegram-type of MessageChatChangeTitle
func (messageChatChangeTitle *MessageChatChangeTitle) MessageType() string {
	return "messageChatChangeTitle"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatChangeTitle *MessageChatChangeTitle) GetMessageContentEnum() MessageContentEnum {
	return MessageChatChangeTitleType
}

// MessageChatChangePhoto An updated chat photo
type MessageChatChangePhoto struct {
	tdCommon
	Photo Photo `json:"photo"` // New chat photo
}

// MessageType return the string telegram-type of MessageChatChangePhoto
func (messageChatChangePhoto *MessageChatChangePhoto) MessageType() string {
	return "messageChatChangePhoto"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatChangePhoto *MessageChatChangePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageChatChangePhotoType
}

// MessageChatDeletePhoto A deleted chat photo
type MessageChatDeletePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageChatDeletePhoto
func (messageChatDeletePhoto *MessageChatDeletePhoto) MessageType() string {
	return "messageChatDeletePhoto"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatDeletePhoto *MessageChatDeletePhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageChatDeletePhotoType
}

// MessageChatAddMembers New chat members were added
type MessageChatAddMembers struct {
	tdCommon
	MemberUserIDs []int32 `json:"member_user_ids"` // User identifiers of the new members
}

// MessageType return the string telegram-type of MessageChatAddMembers
func (messageChatAddMembers *MessageChatAddMembers) MessageType() string {
	return "messageChatAddMembers"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatAddMembers *MessageChatAddMembers) GetMessageContentEnum() MessageContentEnum {
	return MessageChatAddMembersType
}

// MessageChatJoinByLink A new member joined the chat by invite link
type MessageChatJoinByLink struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageChatJoinByLink
func (messageChatJoinByLink *MessageChatJoinByLink) MessageType() string {
	return "messageChatJoinByLink"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatJoinByLink *MessageChatJoinByLink) GetMessageContentEnum() MessageContentEnum {
	return MessageChatJoinByLinkType
}

// MessageChatDeleteMember A chat member was deleted
type MessageChatDeleteMember struct {
	tdCommon
	UserID int32 `json:"user_id"` // User identifier of the deleted chat member
}

// MessageType return the string telegram-type of MessageChatDeleteMember
func (messageChatDeleteMember *MessageChatDeleteMember) MessageType() string {
	return "messageChatDeleteMember"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatDeleteMember *MessageChatDeleteMember) GetMessageContentEnum() MessageContentEnum {
	return MessageChatDeleteMemberType
}

// MessageChatUpgradeTo A basic group was upgraded to a supergroup and was deactivated as the result
type MessageChatUpgradeTo struct {
	tdCommon
	SupergroupID int32 `json:"supergroup_id"` // Identifier of the supergroup to which the basic group was upgraded
}

// MessageType return the string telegram-type of MessageChatUpgradeTo
func (messageChatUpgradeTo *MessageChatUpgradeTo) MessageType() string {
	return "messageChatUpgradeTo"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatUpgradeTo *MessageChatUpgradeTo) GetMessageContentEnum() MessageContentEnum {
	return MessageChatUpgradeToType
}

// MessageChatUpgradeFrom A supergroup has been created from a basic group
type MessageChatUpgradeFrom struct {
	tdCommon
	Title        string `json:"title"`          // Title of the newly created supergroup
	BasicGroupID int32  `json:"basic_group_id"` // The identifier of the original basic group
}

// MessageType return the string telegram-type of MessageChatUpgradeFrom
func (messageChatUpgradeFrom *MessageChatUpgradeFrom) MessageType() string {
	return "messageChatUpgradeFrom"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatUpgradeFrom *MessageChatUpgradeFrom) GetMessageContentEnum() MessageContentEnum {
	return MessageChatUpgradeFromType
}

// MessagePinMessage A message has been pinned
type MessagePinMessage struct {
	tdCommon
	MessageID int64 `json:"message_id"` // Identifier of the pinned message, can be an identifier of a deleted message
}

// MessageType return the string telegram-type of MessagePinMessage
func (messagePinMessage *MessagePinMessage) MessageType() string {
	return "messagePinMessage"
}

// GetMessageContentEnum return the enum type of this object
func (messagePinMessage *MessagePinMessage) GetMessageContentEnum() MessageContentEnum {
	return MessagePinMessageType
}

// MessageScreenshotTaken A screenshot of a message in the chat has been taken
type MessageScreenshotTaken struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageScreenshotTaken
func (messageScreenshotTaken *MessageScreenshotTaken) MessageType() string {
	return "messageScreenshotTaken"
}

// GetMessageContentEnum return the enum type of this object
func (messageScreenshotTaken *MessageScreenshotTaken) GetMessageContentEnum() MessageContentEnum {
	return MessageScreenshotTakenType
}

// MessageChatSetTTL The TTL (Time To Live) setting messages in a secret chat has been changed
type MessageChatSetTTL struct {
	tdCommon
	TTL int32 `json:"ttl"` // New TTL
}

// MessageType return the string telegram-type of MessageChatSetTTL
func (messageChatSetTTL *MessageChatSetTTL) MessageType() string {
	return "messageChatSetTtl"
}

// GetMessageContentEnum return the enum type of this object
func (messageChatSetTTL *MessageChatSetTTL) GetMessageContentEnum() MessageContentEnum {
	return MessageChatSetTTLType
}

// MessageCustomServiceAction A non-standard action has happened in the chat
type MessageCustomServiceAction struct {
	tdCommon
	Text string `json:"text"` // Message text to be shown in the chat
}

// MessageType return the string telegram-type of MessageCustomServiceAction
func (messageCustomServiceAction *MessageCustomServiceAction) MessageType() string {
	return "messageCustomServiceAction"
}

// GetMessageContentEnum return the enum type of this object
func (messageCustomServiceAction *MessageCustomServiceAction) GetMessageContentEnum() MessageContentEnum {
	return MessageCustomServiceActionType
}

// MessageGameScore A new high score was achieved in a game
type MessageGameScore struct {
	tdCommon
	GameMessageID int64     `json:"game_message_id"` // Identifier of the message with the game, can be an identifier of a deleted message
	GameID        JSONInt64 `json:"game_id"`         // Identifier of the game, may be different from the games presented in the message with the game
	Score         int32     `json:"score"`           // New score
}

// MessageType return the string telegram-type of MessageGameScore
func (messageGameScore *MessageGameScore) MessageType() string {
	return "messageGameScore"
}

// GetMessageContentEnum return the enum type of this object
func (messageGameScore *MessageGameScore) GetMessageContentEnum() MessageContentEnum {
	return MessageGameScoreType
}

// MessagePaymentSuccessful A payment has been completed
type MessagePaymentSuccessful struct {
	tdCommon
	InvoiceMessageID int64  `json:"invoice_message_id"` // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
	Currency         string `json:"currency"`           // Currency for the price of the product
	TotalAmount      int64  `json:"total_amount"`       // Total price for the product, in the minimal quantity of the currency
}

// MessageType return the string telegram-type of MessagePaymentSuccessful
func (messagePaymentSuccessful *MessagePaymentSuccessful) MessageType() string {
	return "messagePaymentSuccessful"
}

// GetMessageContentEnum return the enum type of this object
func (messagePaymentSuccessful *MessagePaymentSuccessful) GetMessageContentEnum() MessageContentEnum {
	return MessagePaymentSuccessfulType
}

// MessagePaymentSuccessfulBot A payment has been completed; for bots only
type MessagePaymentSuccessfulBot struct {
	tdCommon
	InvoiceMessageID        int64     `json:"invoice_message_id"`         // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
	Currency                string    `json:"currency"`                   // Currency for price of the product
	TotalAmount             int64     `json:"total_amount"`               // Total price for the product, in the minimal quantity of the currency
	InvoicePayload          []byte    `json:"invoice_payload"`            // Invoice payload
	ShippingOptionID        string    `json:"shipping_option_id"`         // Identifier of the shipping option chosen by the user, may be empty if not applicable
	OrderInfo               OrderInfo `json:"order_info"`                 // Information about the order; may be null
	TelegramPaymentChargeID string    `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeID string    `json:"provider_payment_charge_id"` // Provider payment identifier
}

// MessageType return the string telegram-type of MessagePaymentSuccessfulBot
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) MessageType() string {
	return "messagePaymentSuccessfulBot"
}

// GetMessageContentEnum return the enum type of this object
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) GetMessageContentEnum() MessageContentEnum {
	return MessagePaymentSuccessfulBotType
}

// MessageContactRegistered A contact has registered with Telegram
type MessageContactRegistered struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageContactRegistered
func (messageContactRegistered *MessageContactRegistered) MessageType() string {
	return "messageContactRegistered"
}

// GetMessageContentEnum return the enum type of this object
func (messageContactRegistered *MessageContactRegistered) GetMessageContentEnum() MessageContentEnum {
	return MessageContactRegisteredType
}

// MessageWebsiteConnected The current user has connected a website by logging in using Telegram Login Widget on it
type MessageWebsiteConnected struct {
	tdCommon
	DomainName string `json:"domain_name"` // Domain name of the connected website
}

// MessageType return the string telegram-type of MessageWebsiteConnected
func (messageWebsiteConnected *MessageWebsiteConnected) MessageType() string {
	return "messageWebsiteConnected"
}

// GetMessageContentEnum return the enum type of this object
func (messageWebsiteConnected *MessageWebsiteConnected) GetMessageContentEnum() MessageContentEnum {
	return MessageWebsiteConnectedType
}

// MessageUnsupported Message content that is not supported by the client
type MessageUnsupported struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageUnsupported
func (messageUnsupported *MessageUnsupported) MessageType() string {
	return "messageUnsupported"
}

// GetMessageContentEnum return the enum type of this object
func (messageUnsupported *MessageUnsupported) GetMessageContentEnum() MessageContentEnum {
	return MessageUnsupportedType
}

// TextEntityTypeMention A mention of a user by their username
type TextEntityTypeMention struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeMention
func (textEntityTypeMention *TextEntityTypeMention) MessageType() string {
	return "textEntityTypeMention"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeMention *TextEntityTypeMention) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeMentionType
}

// TextEntityTypeHashtag A hashtag text, beginning with "#"
type TextEntityTypeHashtag struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeHashtag
func (textEntityTypeHashtag *TextEntityTypeHashtag) MessageType() string {
	return "textEntityTypeHashtag"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeHashtag *TextEntityTypeHashtag) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeHashtagType
}

// TextEntityTypeCashtag A cashtag text, beginning with "$" and consisting of capital english letters (i.e. "$USD")
type TextEntityTypeCashtag struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeCashtag
func (textEntityTypeCashtag *TextEntityTypeCashtag) MessageType() string {
	return "textEntityTypeCashtag"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeCashtag *TextEntityTypeCashtag) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeCashtagType
}

// TextEntityTypeBotCommand A bot command, beginning with "/". This shouldn't be highlighted if there are no bots in the chat
type TextEntityTypeBotCommand struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeBotCommand
func (textEntityTypeBotCommand *TextEntityTypeBotCommand) MessageType() string {
	return "textEntityTypeBotCommand"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeBotCommand *TextEntityTypeBotCommand) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeBotCommandType
}

// TextEntityTypeURL An HTTP URL
type TextEntityTypeURL struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeURL
func (textEntityTypeURL *TextEntityTypeURL) MessageType() string {
	return "textEntityTypeUrl"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeURL *TextEntityTypeURL) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeURLType
}

// TextEntityTypeEmailAddress An email address
type TextEntityTypeEmailAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeEmailAddress
func (textEntityTypeEmailAddress *TextEntityTypeEmailAddress) MessageType() string {
	return "textEntityTypeEmailAddress"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeEmailAddress *TextEntityTypeEmailAddress) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeEmailAddressType
}

// TextEntityTypeBold A bold text
type TextEntityTypeBold struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeBold
func (textEntityTypeBold *TextEntityTypeBold) MessageType() string {
	return "textEntityTypeBold"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeBold *TextEntityTypeBold) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeBoldType
}

// TextEntityTypeItalic An italic text
type TextEntityTypeItalic struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeItalic
func (textEntityTypeItalic *TextEntityTypeItalic) MessageType() string {
	return "textEntityTypeItalic"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeItalic *TextEntityTypeItalic) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeItalicType
}

// TextEntityTypeCode Text that must be formatted as if inside a code HTML tag
type TextEntityTypeCode struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypeCode
func (textEntityTypeCode *TextEntityTypeCode) MessageType() string {
	return "textEntityTypeCode"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeCode *TextEntityTypeCode) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeCodeType
}

// TextEntityTypePre Text that must be formatted as if inside a pre HTML tag
type TextEntityTypePre struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypePre
func (textEntityTypePre *TextEntityTypePre) MessageType() string {
	return "textEntityTypePre"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePre *TextEntityTypePre) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePreType
}

// TextEntityTypePreCode Text that must be formatted as if inside pre, and code HTML tags
type TextEntityTypePreCode struct {
	tdCommon
	Language string `json:"language"` // Programming language of the code; as defined by the sender
}

// MessageType return the string telegram-type of TextEntityTypePreCode
func (textEntityTypePreCode *TextEntityTypePreCode) MessageType() string {
	return "textEntityTypePreCode"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePreCode *TextEntityTypePreCode) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePreCodeType
}

// TextEntityTypeTextURL A text description shown instead of a raw URL
type TextEntityTypeTextURL struct {
	tdCommon
	URL string `json:"url"` // URL to be opened when the link is clicked
}

// MessageType return the string telegram-type of TextEntityTypeTextURL
func (textEntityTypeTextURL *TextEntityTypeTextURL) MessageType() string {
	return "textEntityTypeTextUrl"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeTextURL *TextEntityTypeTextURL) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeTextURLType
}

// TextEntityTypeMentionName A text shows instead of a raw mention of the user (e.g., when the user has no username)
type TextEntityTypeMentionName struct {
	tdCommon
	UserID int32 `json:"user_id"` // Identifier of the mentioned user
}

// MessageType return the string telegram-type of TextEntityTypeMentionName
func (textEntityTypeMentionName *TextEntityTypeMentionName) MessageType() string {
	return "textEntityTypeMentionName"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypeMentionName *TextEntityTypeMentionName) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypeMentionNameType
}

// TextEntityTypePhoneNumber A phone number
type TextEntityTypePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of TextEntityTypePhoneNumber
func (textEntityTypePhoneNumber *TextEntityTypePhoneNumber) MessageType() string {
	return "textEntityTypePhoneNumber"
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePhoneNumber *TextEntityTypePhoneNumber) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePhoneNumberType
}

// InputThumbnail A thumbnail to be sent along with a file; should be in JPEG or WEBP format for stickers, and less than 200 kB in size
type InputThumbnail struct {
	tdCommon
	Thumbnail InputFile `json:"thumbnail"` // Thumbnail file to send. Sending thumbnails by file_id is currently not supported
	Width     int32     `json:"width"`     // Thumbnail width, usually shouldn't exceed 90. Use 0 if unknown
	Height    int32     `json:"height"`    // Thumbnail height, usually shouldn't exceed 90. Use 0 if unknown
}

// MessageType return the string telegram-type of InputThumbnail
func (inputThumbnail *InputThumbnail) MessageType() string {
	return "inputThumbnail"
}

// UnmarshalJSON unmarshal to json
func (inputThumbnail *InputThumbnail) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Width  int32 `json:"width"`  // Thumbnail width, usually shouldn't exceed 90. Use 0 if unknown
		Height int32 `json:"height"` // Thumbnail height, usually shouldn't exceed 90. Use 0 if unknown
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputThumbnail.tdCommon = tempObj.tdCommon
	inputThumbnail.Width = tempObj.Width
	inputThumbnail.Height = tempObj.Height

	fieldThumbnail, _ := unmarshalInputFile(objMap["thumbnail"])
	inputThumbnail.Thumbnail = fieldThumbnail

	return nil
}

// InputMessageText A text message
type InputMessageText struct {
	tdCommon
	Text                  FormattedText `json:"text"`                     // Formatted text to be sent. Only Bold, Italic, Code, Pre, PreCode and TextUrl entities are allowed to be specified manually
	DisableWebPagePreview bool          `json:"disable_web_page_preview"` // True, if rich web page previews for URLs in the message text should be disabled
	ClearDraft            bool          `json:"clear_draft"`              // True, if a chat message draft should be deleted
}

// MessageType return the string telegram-type of InputMessageText
func (inputMessageText *InputMessageText) MessageType() string {
	return "inputMessageText"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageText *InputMessageText) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageTextType
}

// InputMessageAnimation An animation message (GIF-style).
type InputMessageAnimation struct {
	tdCommon
	Animation InputFile      `json:"animation"` // Animation file to be sent
	Thumbnail InputThumbnail `json:"thumbnail"` // Animation thumbnail, if available
	Duration  int32          `json:"duration"`  // Duration of the animation, in seconds
	Width     int32          `json:"width"`     // Width of the animation; may be replaced by the server
	Height    int32          `json:"height"`    // Height of the animation; may be replaced by the server
	Caption   FormattedText  `json:"caption"`   // Animation caption; 0-200 characters
}

// MessageType return the string telegram-type of InputMessageAnimation
func (inputMessageAnimation *InputMessageAnimation) MessageType() string {
	return "inputMessageAnimation"
}

// UnmarshalJSON unmarshal to json
func (inputMessageAnimation *InputMessageAnimation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail InputThumbnail `json:"thumbnail"` // Animation thumbnail, if available
		Duration  int32          `json:"duration"`  // Duration of the animation, in seconds
		Width     int32          `json:"width"`     // Width of the animation; may be replaced by the server
		Height    int32          `json:"height"`    // Height of the animation; may be replaced by the server
		Caption   FormattedText  `json:"caption"`   // Animation caption; 0-200 characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageAnimation.tdCommon = tempObj.tdCommon
	inputMessageAnimation.Thumbnail = tempObj.Thumbnail
	inputMessageAnimation.Duration = tempObj.Duration
	inputMessageAnimation.Width = tempObj.Width
	inputMessageAnimation.Height = tempObj.Height
	inputMessageAnimation.Caption = tempObj.Caption

	fieldAnimation, _ := unmarshalInputFile(objMap["animation"])
	inputMessageAnimation.Animation = fieldAnimation

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageAnimation *InputMessageAnimation) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageAnimationType
}

// InputMessageAudio An audio message
type InputMessageAudio struct {
	tdCommon
	Audio               InputFile      `json:"audio"`                 // Audio file to be sent
	AlbumCoverThumbnail InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album, if available
	Duration            int32          `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
	Title               string         `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
	Performer           string         `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
	Caption             FormattedText  `json:"caption"`               // Audio caption; 0-200 characters
}

// MessageType return the string telegram-type of InputMessageAudio
func (inputMessageAudio *InputMessageAudio) MessageType() string {
	return "inputMessageAudio"
}

// UnmarshalJSON unmarshal to json
func (inputMessageAudio *InputMessageAudio) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AlbumCoverThumbnail InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album, if available
		Duration            int32          `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
		Title               string         `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
		Performer           string         `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
		Caption             FormattedText  `json:"caption"`               // Audio caption; 0-200 characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageAudio.tdCommon = tempObj.tdCommon
	inputMessageAudio.AlbumCoverThumbnail = tempObj.AlbumCoverThumbnail
	inputMessageAudio.Duration = tempObj.Duration
	inputMessageAudio.Title = tempObj.Title
	inputMessageAudio.Performer = tempObj.Performer
	inputMessageAudio.Caption = tempObj.Caption

	fieldAudio, _ := unmarshalInputFile(objMap["audio"])
	inputMessageAudio.Audio = fieldAudio

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageAudio *InputMessageAudio) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageAudioType
}

// InputMessageDocument A document message (general file)
type InputMessageDocument struct {
	tdCommon
	Document  InputFile      `json:"document"`  // Document to be sent
	Thumbnail InputThumbnail `json:"thumbnail"` // Document thumbnail, if available
	Caption   FormattedText  `json:"caption"`   // Document caption; 0-200 characters
}

// MessageType return the string telegram-type of InputMessageDocument
func (inputMessageDocument *InputMessageDocument) MessageType() string {
	return "inputMessageDocument"
}

// UnmarshalJSON unmarshal to json
func (inputMessageDocument *InputMessageDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail InputThumbnail `json:"thumbnail"` // Document thumbnail, if available
		Caption   FormattedText  `json:"caption"`   // Document caption; 0-200 characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageDocument.tdCommon = tempObj.tdCommon
	inputMessageDocument.Thumbnail = tempObj.Thumbnail
	inputMessageDocument.Caption = tempObj.Caption

	fieldDocument, _ := unmarshalInputFile(objMap["document"])
	inputMessageDocument.Document = fieldDocument

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageDocument *InputMessageDocument) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageDocumentType
}

// InputMessagePhoto A photo message
type InputMessagePhoto struct {
	tdCommon
	Photo               InputFile      `json:"photo"`                  // Photo to send
	Thumbnail           InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent, this is sent to the other party in secret chats only
	AddedStickerFileIDs []int32        `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
	Width               int32          `json:"width"`                  // Photo width
	Height              int32          `json:"height"`                 // Photo height
	Caption             FormattedText  `json:"caption"`                // Photo caption; 0-200 characters
	TTL                 int32          `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessagePhoto
func (inputMessagePhoto *InputMessagePhoto) MessageType() string {
	return "inputMessagePhoto"
}

// UnmarshalJSON unmarshal to json
func (inputMessagePhoto *InputMessagePhoto) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail           InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent, this is sent to the other party in secret chats only
		AddedStickerFileIDs []int32        `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
		Width               int32          `json:"width"`                  // Photo width
		Height              int32          `json:"height"`                 // Photo height
		Caption             FormattedText  `json:"caption"`                // Photo caption; 0-200 characters
		TTL                 int32          `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessagePhoto.tdCommon = tempObj.tdCommon
	inputMessagePhoto.Thumbnail = tempObj.Thumbnail
	inputMessagePhoto.AddedStickerFileIDs = tempObj.AddedStickerFileIDs
	inputMessagePhoto.Width = tempObj.Width
	inputMessagePhoto.Height = tempObj.Height
	inputMessagePhoto.Caption = tempObj.Caption
	inputMessagePhoto.TTL = tempObj.TTL

	fieldPhoto, _ := unmarshalInputFile(objMap["photo"])
	inputMessagePhoto.Photo = fieldPhoto

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessagePhoto *InputMessagePhoto) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessagePhotoType
}

// InputMessageSticker A sticker message
type InputMessageSticker struct {
	tdCommon
	Sticker   InputFile      `json:"sticker"`   // Sticker to be sent
	Thumbnail InputThumbnail `json:"thumbnail"` // Sticker thumbnail, if available
	Width     int32          `json:"width"`     // Sticker width
	Height    int32          `json:"height"`    // Sticker height
}

// MessageType return the string telegram-type of InputMessageSticker
func (inputMessageSticker *InputMessageSticker) MessageType() string {
	return "inputMessageSticker"
}

// UnmarshalJSON unmarshal to json
func (inputMessageSticker *InputMessageSticker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail InputThumbnail `json:"thumbnail"` // Sticker thumbnail, if available
		Width     int32          `json:"width"`     // Sticker width
		Height    int32          `json:"height"`    // Sticker height
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageSticker.tdCommon = tempObj.tdCommon
	inputMessageSticker.Thumbnail = tempObj.Thumbnail
	inputMessageSticker.Width = tempObj.Width
	inputMessageSticker.Height = tempObj.Height

	fieldSticker, _ := unmarshalInputFile(objMap["sticker"])
	inputMessageSticker.Sticker = fieldSticker

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageSticker *InputMessageSticker) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageStickerType
}

// InputMessageVideo A video message
type InputMessageVideo struct {
	tdCommon
	Video               InputFile      `json:"video"`                  // Video to be sent
	Thumbnail           InputThumbnail `json:"thumbnail"`              // Video thumbnail, if available
	AddedStickerFileIDs []int32        `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
	Duration            int32          `json:"duration"`               // Duration of the video, in seconds
	Width               int32          `json:"width"`                  // Video width
	Height              int32          `json:"height"`                 // Video height
	SupportsStreaming   bool           `json:"supports_streaming"`     // True, if the video should be tried to be streamed
	Caption             FormattedText  `json:"caption"`                // Video caption; 0-200 characters
	TTL                 int32          `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessageVideo
func (inputMessageVideo *InputMessageVideo) MessageType() string {
	return "inputMessageVideo"
}

// UnmarshalJSON unmarshal to json
func (inputMessageVideo *InputMessageVideo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail           InputThumbnail `json:"thumbnail"`              // Video thumbnail, if available
		AddedStickerFileIDs []int32        `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
		Duration            int32          `json:"duration"`               // Duration of the video, in seconds
		Width               int32          `json:"width"`                  // Video width
		Height              int32          `json:"height"`                 // Video height
		SupportsStreaming   bool           `json:"supports_streaming"`     // True, if the video should be tried to be streamed
		Caption             FormattedText  `json:"caption"`                // Video caption; 0-200 characters
		TTL                 int32          `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVideo.tdCommon = tempObj.tdCommon
	inputMessageVideo.Thumbnail = tempObj.Thumbnail
	inputMessageVideo.AddedStickerFileIDs = tempObj.AddedStickerFileIDs
	inputMessageVideo.Duration = tempObj.Duration
	inputMessageVideo.Width = tempObj.Width
	inputMessageVideo.Height = tempObj.Height
	inputMessageVideo.SupportsStreaming = tempObj.SupportsStreaming
	inputMessageVideo.Caption = tempObj.Caption
	inputMessageVideo.TTL = tempObj.TTL

	fieldVideo, _ := unmarshalInputFile(objMap["video"])
	inputMessageVideo.Video = fieldVideo

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVideo *InputMessageVideo) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVideoType
}

// InputMessageVideoNote A video note message
type InputMessageVideoNote struct {
	tdCommon
	VideoNote InputFile      `json:"video_note"` // Video note to be sent
	Thumbnail InputThumbnail `json:"thumbnail"`  // Video thumbnail, if available
	Duration  int32          `json:"duration"`   // Duration of the video, in seconds
	Length    int32          `json:"length"`     // Video width and height; must be positive and not greater than 640
}

// MessageType return the string telegram-type of InputMessageVideoNote
func (inputMessageVideoNote *InputMessageVideoNote) MessageType() string {
	return "inputMessageVideoNote"
}

// UnmarshalJSON unmarshal to json
func (inputMessageVideoNote *InputMessageVideoNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Thumbnail InputThumbnail `json:"thumbnail"` // Video thumbnail, if available
		Duration  int32          `json:"duration"`  // Duration of the video, in seconds
		Length    int32          `json:"length"`    // Video width and height; must be positive and not greater than 640
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVideoNote.tdCommon = tempObj.tdCommon
	inputMessageVideoNote.Thumbnail = tempObj.Thumbnail
	inputMessageVideoNote.Duration = tempObj.Duration
	inputMessageVideoNote.Length = tempObj.Length

	fieldVideoNote, _ := unmarshalInputFile(objMap["video_note"])
	inputMessageVideoNote.VideoNote = fieldVideoNote

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVideoNote *InputMessageVideoNote) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVideoNoteType
}

// InputMessageVoiceNote A voice note message
type InputMessageVoiceNote struct {
	tdCommon
	VoiceNote InputFile     `json:"voice_note"` // Voice note to be sent
	Duration  int32         `json:"duration"`   // Duration of the voice note, in seconds
	Waveform  []byte        `json:"waveform"`   // Waveform representation of the voice note, in 5-bit format
	Caption   FormattedText `json:"caption"`    // Voice note caption; 0-200 characters
}

// MessageType return the string telegram-type of InputMessageVoiceNote
func (inputMessageVoiceNote *InputMessageVoiceNote) MessageType() string {
	return "inputMessageVoiceNote"
}

// UnmarshalJSON unmarshal to json
func (inputMessageVoiceNote *InputMessageVoiceNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Duration int32         `json:"duration"` // Duration of the voice note, in seconds
		Waveform []byte        `json:"waveform"` // Waveform representation of the voice note, in 5-bit format
		Caption  FormattedText `json:"caption"`  // Voice note caption; 0-200 characters
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputMessageVoiceNote.tdCommon = tempObj.tdCommon
	inputMessageVoiceNote.Duration = tempObj.Duration
	inputMessageVoiceNote.Waveform = tempObj.Waveform
	inputMessageVoiceNote.Caption = tempObj.Caption

	fieldVoiceNote, _ := unmarshalInputFile(objMap["voice_note"])
	inputMessageVoiceNote.VoiceNote = fieldVoiceNote

	return nil
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVoiceNote *InputMessageVoiceNote) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVoiceNoteType
}

// InputMessageLocation A message with a location
type InputMessageLocation struct {
	tdCommon
	Location   Location `json:"location"`    // Location to be sent
	LivePeriod int32    `json:"live_period"` // Period for which the location can be updated, in seconds; should bebetween 60 and 86400 for a live location and 0 otherwise
}

// MessageType return the string telegram-type of InputMessageLocation
func (inputMessageLocation *InputMessageLocation) MessageType() string {
	return "inputMessageLocation"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageLocation *InputMessageLocation) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageLocationType
}

// InputMessageVenue A message with information about a venue
type InputMessageVenue struct {
	tdCommon
	Venue Venue `json:"venue"` // Venue to send
}

// MessageType return the string telegram-type of InputMessageVenue
func (inputMessageVenue *InputMessageVenue) MessageType() string {
	return "inputMessageVenue"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVenue *InputMessageVenue) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVenueType
}

// InputMessageContact A message containing a user contact
type InputMessageContact struct {
	tdCommon
	Contact Contact `json:"contact"` // Contact to send
}

// MessageType return the string telegram-type of InputMessageContact
func (inputMessageContact *InputMessageContact) MessageType() string {
	return "inputMessageContact"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageContact *InputMessageContact) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageContactType
}

// InputMessageGame A message with a game; not supported for channels or secret chats
type InputMessageGame struct {
	tdCommon
	BotUserID     int32  `json:"bot_user_id"`     // User identifier of the bot that owns the game
	GameShortName string `json:"game_short_name"` // Short name of the game
}

// MessageType return the string telegram-type of InputMessageGame
func (inputMessageGame *InputMessageGame) MessageType() string {
	return "inputMessageGame"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageGame *InputMessageGame) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageGameType
}

// InputMessageInvoice A message with an invoice; can be used only by bots and only in private chats
type InputMessageInvoice struct {
	tdCommon
	Invoice        Invoice `json:"invoice"`         // Invoice
	Title          string  `json:"title"`           // Product title; 1-32 characters
	Description    string  `json:"description"`     //
	PhotoURL       string  `json:"photo_url"`       // Product photo URL; optional
	PhotoSize      int32   `json:"photo_size"`      // Product photo size
	PhotoWidth     int32   `json:"photo_width"`     // Product photo width
	PhotoHeight    int32   `json:"photo_height"`    // Product photo height
	Payload        []byte  `json:"payload"`         // The invoice payload
	ProviderToken  string  `json:"provider_token"`  // Payment provider token
	ProviderData   string  `json:"provider_data"`   // JSON-encoded data about the invoice, which will be shared with the payment provider
	StartParameter string  `json:"start_parameter"` // Unique invoice bot start_parameter for the generation of this invoice
}

// MessageType return the string telegram-type of InputMessageInvoice
func (inputMessageInvoice *InputMessageInvoice) MessageType() string {
	return "inputMessageInvoice"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageInvoice *InputMessageInvoice) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageInvoiceType
}

// InputMessageForwarded A forwarded message
type InputMessageForwarded struct {
	tdCommon
	FromChatID  int64 `json:"from_chat_id"`  // Identifier for the chat this forwarded message came from
	MessageID   int64 `json:"message_id"`    // Identifier of the message to forward
	InGameShare bool  `json:"in_game_share"` // True, if a game message should be shared within a launched game; applies only to game messages
}

// MessageType return the string telegram-type of InputMessageForwarded
func (inputMessageForwarded *InputMessageForwarded) MessageType() string {
	return "inputMessageForwarded"
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageForwarded *InputMessageForwarded) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageForwardedType
}

// SearchMessagesFilterEmpty Returns all found messages, no filter is applied
type SearchMessagesFilterEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterEmpty
func (searchMessagesFilterEmpty *SearchMessagesFilterEmpty) MessageType() string {
	return "searchMessagesFilterEmpty"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterEmpty *SearchMessagesFilterEmpty) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterEmptyType
}

// SearchMessagesFilterAnimation Returns only animation messages
type SearchMessagesFilterAnimation struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterAnimation
func (searchMessagesFilterAnimation *SearchMessagesFilterAnimation) MessageType() string {
	return "searchMessagesFilterAnimation"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterAnimation *SearchMessagesFilterAnimation) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterAnimationType
}

// SearchMessagesFilterAudio Returns only audio messages
type SearchMessagesFilterAudio struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterAudio
func (searchMessagesFilterAudio *SearchMessagesFilterAudio) MessageType() string {
	return "searchMessagesFilterAudio"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterAudio *SearchMessagesFilterAudio) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterAudioType
}

// SearchMessagesFilterDocument Returns only document messages
type SearchMessagesFilterDocument struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterDocument
func (searchMessagesFilterDocument *SearchMessagesFilterDocument) MessageType() string {
	return "searchMessagesFilterDocument"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterDocument *SearchMessagesFilterDocument) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterDocumentType
}

// SearchMessagesFilterPhoto Returns only photo messages
type SearchMessagesFilterPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterPhoto
func (searchMessagesFilterPhoto *SearchMessagesFilterPhoto) MessageType() string {
	return "searchMessagesFilterPhoto"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterPhoto *SearchMessagesFilterPhoto) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterPhotoType
}

// SearchMessagesFilterVideo Returns only video messages
type SearchMessagesFilterVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVideo
func (searchMessagesFilterVideo *SearchMessagesFilterVideo) MessageType() string {
	return "searchMessagesFilterVideo"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVideo *SearchMessagesFilterVideo) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVideoType
}

// SearchMessagesFilterVoiceNote Returns only voice note messages
type SearchMessagesFilterVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceNote
func (searchMessagesFilterVoiceNote *SearchMessagesFilterVoiceNote) MessageType() string {
	return "searchMessagesFilterVoiceNote"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVoiceNote *SearchMessagesFilterVoiceNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVoiceNoteType
}

// SearchMessagesFilterPhotoAndVideo Returns only photo and video messages
type SearchMessagesFilterPhotoAndVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterPhotoAndVideo
func (searchMessagesFilterPhotoAndVideo *SearchMessagesFilterPhotoAndVideo) MessageType() string {
	return "searchMessagesFilterPhotoAndVideo"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterPhotoAndVideo *SearchMessagesFilterPhotoAndVideo) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterPhotoAndVideoType
}

// SearchMessagesFilterURL Returns only messages containing URLs
type SearchMessagesFilterURL struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterURL
func (searchMessagesFilterURL *SearchMessagesFilterURL) MessageType() string {
	return "searchMessagesFilterUrl"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterURL *SearchMessagesFilterURL) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterURLType
}

// SearchMessagesFilterChatPhoto Returns only messages containing chat photos
type SearchMessagesFilterChatPhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterChatPhoto
func (searchMessagesFilterChatPhoto *SearchMessagesFilterChatPhoto) MessageType() string {
	return "searchMessagesFilterChatPhoto"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterChatPhoto *SearchMessagesFilterChatPhoto) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterChatPhotoType
}

// SearchMessagesFilterCall Returns only call messages
type SearchMessagesFilterCall struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterCall
func (searchMessagesFilterCall *SearchMessagesFilterCall) MessageType() string {
	return "searchMessagesFilterCall"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterCall *SearchMessagesFilterCall) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterCallType
}

// SearchMessagesFilterMissedCall Returns only incoming call messages with missed/declined discard reasons
type SearchMessagesFilterMissedCall struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterMissedCall
func (searchMessagesFilterMissedCall *SearchMessagesFilterMissedCall) MessageType() string {
	return "searchMessagesFilterMissedCall"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterMissedCall *SearchMessagesFilterMissedCall) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterMissedCallType
}

// SearchMessagesFilterVideoNote Returns only video note messages
type SearchMessagesFilterVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVideoNote
func (searchMessagesFilterVideoNote *SearchMessagesFilterVideoNote) MessageType() string {
	return "searchMessagesFilterVideoNote"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVideoNote *SearchMessagesFilterVideoNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVideoNoteType
}

// SearchMessagesFilterVoiceAndVideoNote Returns only voice and video note messages
type SearchMessagesFilterVoiceAndVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceAndVideoNote
func (searchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote) MessageType() string {
	return "searchMessagesFilterVoiceAndVideoNote"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterVoiceAndVideoNoteType
}

// SearchMessagesFilterMention Returns only messages with mentions of the current user, or messages that are replies to their messages
type SearchMessagesFilterMention struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterMention
func (searchMessagesFilterMention *SearchMessagesFilterMention) MessageType() string {
	return "searchMessagesFilterMention"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterMention *SearchMessagesFilterMention) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterMentionType
}

// SearchMessagesFilterUnreadMention Returns only messages with unread mentions of the current user or messages that are replies to their messages. When using this filter the results can't be additionally filtered by a query or by the sending user
type SearchMessagesFilterUnreadMention struct {
	tdCommon
}

// MessageType return the string telegram-type of SearchMessagesFilterUnreadMention
func (searchMessagesFilterUnreadMention *SearchMessagesFilterUnreadMention) MessageType() string {
	return "searchMessagesFilterUnreadMention"
}

// GetSearchMessagesFilterEnum return the enum type of this object
func (searchMessagesFilterUnreadMention *SearchMessagesFilterUnreadMention) GetSearchMessagesFilterEnum() SearchMessagesFilterEnum {
	return SearchMessagesFilterUnreadMentionType
}

// ChatActionTyping The user is typing a message
type ChatActionTyping struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionTyping
func (chatActionTyping *ChatActionTyping) MessageType() string {
	return "chatActionTyping"
}

// GetChatActionEnum return the enum type of this object
func (chatActionTyping *ChatActionTyping) GetChatActionEnum() ChatActionEnum {
	return ChatActionTypingType
}

// ChatActionRecordingVideo The user is recording a video
type ChatActionRecordingVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVideo
func (chatActionRecordingVideo *ChatActionRecordingVideo) MessageType() string {
	return "chatActionRecordingVideo"
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVideo *ChatActionRecordingVideo) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVideoType
}

// ChatActionUploadingVideo The user is uploading a video
type ChatActionUploadingVideo struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideo
func (chatActionUploadingVideo *ChatActionUploadingVideo) MessageType() string {
	return "chatActionUploadingVideo"
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVideo *ChatActionUploadingVideo) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVideoType
}

// ChatActionRecordingVoiceNote The user is recording a voice note
type ChatActionRecordingVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVoiceNote
func (chatActionRecordingVoiceNote *ChatActionRecordingVoiceNote) MessageType() string {
	return "chatActionRecordingVoiceNote"
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVoiceNote *ChatActionRecordingVoiceNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVoiceNoteType
}

// ChatActionUploadingVoiceNote The user is uploading a voice note
type ChatActionUploadingVoiceNote struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVoiceNote
func (chatActionUploadingVoiceNote *ChatActionUploadingVoiceNote) MessageType() string {
	return "chatActionUploadingVoiceNote"
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVoiceNote *ChatActionUploadingVoiceNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVoiceNoteType
}

// ChatActionUploadingPhoto The user is uploading a photo
type ChatActionUploadingPhoto struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingPhoto
func (chatActionUploadingPhoto *ChatActionUploadingPhoto) MessageType() string {
	return "chatActionUploadingPhoto"
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingPhoto *ChatActionUploadingPhoto) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingPhotoType
}

// ChatActionUploadingDocument The user is uploading a document
type ChatActionUploadingDocument struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingDocument
func (chatActionUploadingDocument *ChatActionUploadingDocument) MessageType() string {
	return "chatActionUploadingDocument"
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingDocument *ChatActionUploadingDocument) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingDocumentType
}

// ChatActionChoosingLocation The user is picking a location or venue to send
type ChatActionChoosingLocation struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionChoosingLocation
func (chatActionChoosingLocation *ChatActionChoosingLocation) MessageType() string {
	return "chatActionChoosingLocation"
}

// GetChatActionEnum return the enum type of this object
func (chatActionChoosingLocation *ChatActionChoosingLocation) GetChatActionEnum() ChatActionEnum {
	return ChatActionChoosingLocationType
}

// ChatActionChoosingContact The user is picking a contact to send
type ChatActionChoosingContact struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionChoosingContact
func (chatActionChoosingContact *ChatActionChoosingContact) MessageType() string {
	return "chatActionChoosingContact"
}

// GetChatActionEnum return the enum type of this object
func (chatActionChoosingContact *ChatActionChoosingContact) GetChatActionEnum() ChatActionEnum {
	return ChatActionChoosingContactType
}

// ChatActionStartPlayingGame The user has started to play a game
type ChatActionStartPlayingGame struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionStartPlayingGame
func (chatActionStartPlayingGame *ChatActionStartPlayingGame) MessageType() string {
	return "chatActionStartPlayingGame"
}

// GetChatActionEnum return the enum type of this object
func (chatActionStartPlayingGame *ChatActionStartPlayingGame) GetChatActionEnum() ChatActionEnum {
	return ChatActionStartPlayingGameType
}

// ChatActionRecordingVideoNote The user is recording a video note
type ChatActionRecordingVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionRecordingVideoNote
func (chatActionRecordingVideoNote *ChatActionRecordingVideoNote) MessageType() string {
	return "chatActionRecordingVideoNote"
}

// GetChatActionEnum return the enum type of this object
func (chatActionRecordingVideoNote *ChatActionRecordingVideoNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionRecordingVideoNoteType
}

// ChatActionUploadingVideoNote The user is uploading a video note
type ChatActionUploadingVideoNote struct {
	tdCommon
	Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideoNote
func (chatActionUploadingVideoNote *ChatActionUploadingVideoNote) MessageType() string {
	return "chatActionUploadingVideoNote"
}

// GetChatActionEnum return the enum type of this object
func (chatActionUploadingVideoNote *ChatActionUploadingVideoNote) GetChatActionEnum() ChatActionEnum {
	return ChatActionUploadingVideoNoteType
}

// ChatActionCancel The user has cancelled the previous action
type ChatActionCancel struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatActionCancel
func (chatActionCancel *ChatActionCancel) MessageType() string {
	return "chatActionCancel"
}

// GetChatActionEnum return the enum type of this object
func (chatActionCancel *ChatActionCancel) GetChatActionEnum() ChatActionEnum {
	return ChatActionCancelType
}

// UserStatusEmpty The user status was never changed
type UserStatusEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusEmpty
func (userStatusEmpty *UserStatusEmpty) MessageType() string {
	return "userStatusEmpty"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusEmpty *UserStatusEmpty) GetUserStatusEnum() UserStatusEnum {
	return UserStatusEmptyType
}

// UserStatusOnline The user is online
type UserStatusOnline struct {
	tdCommon
	Expires int32 `json:"expires"` // Point in time (Unix timestamp) when the user's online status will expire
}

// MessageType return the string telegram-type of UserStatusOnline
func (userStatusOnline *UserStatusOnline) MessageType() string {
	return "userStatusOnline"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusOnline *UserStatusOnline) GetUserStatusEnum() UserStatusEnum {
	return UserStatusOnlineType
}

// UserStatusOffline The user is offline
type UserStatusOffline struct {
	tdCommon
	WasOnline int32 `json:"was_online"` // Point in time (Unix timestamp) when the user was last online
}

// MessageType return the string telegram-type of UserStatusOffline
func (userStatusOffline *UserStatusOffline) MessageType() string {
	return "userStatusOffline"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusOffline *UserStatusOffline) GetUserStatusEnum() UserStatusEnum {
	return UserStatusOfflineType
}

// UserStatusRecently The user was online recently
type UserStatusRecently struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusRecently
func (userStatusRecently *UserStatusRecently) MessageType() string {
	return "userStatusRecently"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusRecently *UserStatusRecently) GetUserStatusEnum() UserStatusEnum {
	return UserStatusRecentlyType
}

// UserStatusLastWeek The user is offline, but was online last week
type UserStatusLastWeek struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusLastWeek
func (userStatusLastWeek *UserStatusLastWeek) MessageType() string {
	return "userStatusLastWeek"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusLastWeek *UserStatusLastWeek) GetUserStatusEnum() UserStatusEnum {
	return UserStatusLastWeekType
}

// UserStatusLastMonth The user is offline, but was online last month
type UserStatusLastMonth struct {
	tdCommon
}

// MessageType return the string telegram-type of UserStatusLastMonth
func (userStatusLastMonth *UserStatusLastMonth) MessageType() string {
	return "userStatusLastMonth"
}

// GetUserStatusEnum return the enum type of this object
func (userStatusLastMonth *UserStatusLastMonth) GetUserStatusEnum() UserStatusEnum {
	return UserStatusLastMonthType
}

// Stickers Represents a list of stickers
type Stickers struct {
	tdCommon
	Stickers []Sticker `json:"stickers"` // List of stickers
}

// MessageType return the string telegram-type of Stickers
func (stickers *Stickers) MessageType() string {
	return "stickers"
}

// StickerEmojis Represents a list of all emoji corresponding to a sticker in a sticker set. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
type StickerEmojis struct {
	tdCommon
	Emojis []string `json:"emojis"` // List of emojis
}

// MessageType return the string telegram-type of StickerEmojis
func (stickerEmojis *StickerEmojis) MessageType() string {
	return "stickerEmojis"
}

// StickerSet Represents a sticker set
type StickerSet struct {
	tdCommon
	ID          JSONInt64       `json:"id"`           // Identifier of the sticker set
	Title       string          `json:"title"`        // Title of the sticker set
	Name        string          `json:"name"`         // Name of the sticker set
	IsInstalled bool            `json:"is_installed"` // True, if the sticker set has been installed by the current user
	IsArchived  bool            `json:"is_archived"`  // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
	IsOfficial  bool            `json:"is_official"`  // True, if the sticker set is official
	IsMasks     bool            `json:"is_masks"`     // True, if the stickers in the set are masks
	IsViewed    bool            `json:"is_viewed"`    // True for already viewed trending sticker sets
	Stickers    []Sticker       `json:"stickers"`     // List of stickers in this set
	Emojis      []StickerEmojis `json:"emojis"`       // A list of emoji corresponding to the stickers in the same order
}

// MessageType return the string telegram-type of StickerSet
func (stickerSet *StickerSet) MessageType() string {
	return "stickerSet"
}

// StickerSetInfo Represents short information about a sticker set
type StickerSetInfo struct {
	tdCommon
	ID          JSONInt64 `json:"id"`           // Identifier of the sticker set
	Title       string    `json:"title"`        // Title of the sticker set
	Name        string    `json:"name"`         // Name of the sticker set
	IsInstalled bool      `json:"is_installed"` // True, if the sticker set has been installed by current user
	IsArchived  bool      `json:"is_archived"`  // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
	IsOfficial  bool      `json:"is_official"`  // True, if the sticker set is official
	IsMasks     bool      `json:"is_masks"`     // True, if the stickers in the set are masks
	IsViewed    bool      `json:"is_viewed"`    // True for already viewed trending sticker sets
	Size        int32     `json:"size"`         // Total number of stickers in the set
	Covers      []Sticker `json:"covers"`       // Contains up to the first 5 stickers from the set, depending on the context. If the client needs more stickers the full set should be requested
}

// MessageType return the string telegram-type of StickerSetInfo
func (stickerSetInfo *StickerSetInfo) MessageType() string {
	return "stickerSetInfo"
}

// StickerSets Represents a list of sticker sets
type StickerSets struct {
	tdCommon
	TotalCount int32            `json:"total_count"` // Approximate total number of sticker sets found
	Sets       []StickerSetInfo `json:"sets"`        // List of sticker sets
}

// MessageType return the string telegram-type of StickerSets
func (stickerSets *StickerSets) MessageType() string {
	return "stickerSets"
}

// CallDiscardReasonEmpty The call wasn't discarded, or the reason is unknown
type CallDiscardReasonEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonEmpty
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) MessageType() string {
	return "callDiscardReasonEmpty"
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonEmptyType
}

// CallDiscardReasonMissed The call was ended before the conversation started. It was cancelled by the caller or missed by the other party
type CallDiscardReasonMissed struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonMissed
func (callDiscardReasonMissed *CallDiscardReasonMissed) MessageType() string {
	return "callDiscardReasonMissed"
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonMissed *CallDiscardReasonMissed) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonMissedType
}

// CallDiscardReasonDeclined The call was ended before the conversation started. It was declined by the other party
type CallDiscardReasonDeclined struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonDeclined
func (callDiscardReasonDeclined *CallDiscardReasonDeclined) MessageType() string {
	return "callDiscardReasonDeclined"
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonDeclined *CallDiscardReasonDeclined) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonDeclinedType
}

// CallDiscardReasonDisconnected The call was ended during the conversation because the users were disconnected
type CallDiscardReasonDisconnected struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonDisconnected
func (callDiscardReasonDisconnected *CallDiscardReasonDisconnected) MessageType() string {
	return "callDiscardReasonDisconnected"
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonDisconnected *CallDiscardReasonDisconnected) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonDisconnectedType
}

// CallDiscardReasonHungUp The call was ended because one of the parties hung up
type CallDiscardReasonHungUp struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonHungUp
func (callDiscardReasonHungUp *CallDiscardReasonHungUp) MessageType() string {
	return "callDiscardReasonHungUp"
}

// GetCallDiscardReasonEnum return the enum type of this object
func (callDiscardReasonHungUp *CallDiscardReasonHungUp) GetCallDiscardReasonEnum() CallDiscardReasonEnum {
	return CallDiscardReasonHungUpType
}

// CallProtocol Specifies the supported call protocols
type CallProtocol struct {
	tdCommon
	UDPP2p       bool  `json:"udp_p2p"`       // True, if UDP peer-to-peer connections are supported
	UDPReflector bool  `json:"udp_reflector"` // True, if connection through UDP reflectors is supported
	MinLayer     int32 `json:"min_layer"`     // Minimum supported API layer; use 65
	MaxLayer     int32 `json:"max_layer"`     // Maximum supported API layer; use 65
}

// MessageType return the string telegram-type of CallProtocol
func (callProtocol *CallProtocol) MessageType() string {
	return "callProtocol"
}

// CallConnection Describes the address of UDP reflectors
type CallConnection struct {
	tdCommon
	ID      JSONInt64 `json:"id"`       // Reflector identifier
	IP      string    `json:"ip"`       // IPv4 reflector address
	IPv6    string    `json:"ipv6"`     // IPv6 reflector address
	Port    int32     `json:"port"`     // Reflector port number
	PeerTag []byte    `json:"peer_tag"` // Connection peer tag
}

// MessageType return the string telegram-type of CallConnection
func (callConnection *CallConnection) MessageType() string {
	return "callConnection"
}

// CallID Contains the call identifier
type CallID struct {
	tdCommon
	ID int32 `json:"id"` // Call identifier
}

// MessageType return the string telegram-type of CallID
func (callID *CallID) MessageType() string {
	return "callId"
}

// CallStatePending The call is pending, waiting to be accepted by a user
type CallStatePending struct {
	tdCommon
	IsCreated  bool `json:"is_created"`  // True, if the call has already been created by the server
	IsReceived bool `json:"is_received"` // True, if the call has already been received by the other party
}

// MessageType return the string telegram-type of CallStatePending
func (callStatePending *CallStatePending) MessageType() string {
	return "callStatePending"
}

// GetCallStateEnum return the enum type of this object
func (callStatePending *CallStatePending) GetCallStateEnum() CallStateEnum {
	return CallStatePendingType
}

// CallStateExchangingKeys The call has been answered and encryption keys are being exchanged
type CallStateExchangingKeys struct {
	tdCommon
}

// MessageType return the string telegram-type of CallStateExchangingKeys
func (callStateExchangingKeys *CallStateExchangingKeys) MessageType() string {
	return "callStateExchangingKeys"
}

// GetCallStateEnum return the enum type of this object
func (callStateExchangingKeys *CallStateExchangingKeys) GetCallStateEnum() CallStateEnum {
	return CallStateExchangingKeysType
}

// CallStateReady The call is ready to use
type CallStateReady struct {
	tdCommon
	Protocol      CallProtocol     `json:"protocol"`       // Call protocols supported by the peer
	Connections   []CallConnection `json:"connections"`    // Available UDP reflectors
	Config        string           `json:"config"`         // A JSON-encoded call config
	EncryptionKey []byte           `json:"encryption_key"` // Call encryption key
	Emojis        []string         `json:"emojis"`         // Encryption key emojis fingerprint
}

// MessageType return the string telegram-type of CallStateReady
func (callStateReady *CallStateReady) MessageType() string {
	return "callStateReady"
}

// GetCallStateEnum return the enum type of this object
func (callStateReady *CallStateReady) GetCallStateEnum() CallStateEnum {
	return CallStateReadyType
}

// CallStateHangingUp The call is hanging up after discardCall has been called
type CallStateHangingUp struct {
	tdCommon
}

// MessageType return the string telegram-type of CallStateHangingUp
func (callStateHangingUp *CallStateHangingUp) MessageType() string {
	return "callStateHangingUp"
}

// GetCallStateEnum return the enum type of this object
func (callStateHangingUp *CallStateHangingUp) GetCallStateEnum() CallStateEnum {
	return CallStateHangingUpType
}

// CallStateDiscarded The call has ended successfully
type CallStateDiscarded struct {
	tdCommon
	Reason               CallDiscardReason `json:"reason"`                 // The reason, why the call has ended
	NeedRating           bool              `json:"need_rating"`            // True, if the call rating should be sent to the server
	NeedDebugInformation bool              `json:"need_debug_information"` // True, if the call debug information should be sent to the server
}

// MessageType return the string telegram-type of CallStateDiscarded
func (callStateDiscarded *CallStateDiscarded) MessageType() string {
	return "callStateDiscarded"
}

// UnmarshalJSON unmarshal to json
func (callStateDiscarded *CallStateDiscarded) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NeedRating           bool `json:"need_rating"`            // True, if the call rating should be sent to the server
		NeedDebugInformation bool `json:"need_debug_information"` // True, if the call debug information should be sent to the server
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	callStateDiscarded.tdCommon = tempObj.tdCommon
	callStateDiscarded.NeedRating = tempObj.NeedRating
	callStateDiscarded.NeedDebugInformation = tempObj.NeedDebugInformation

	fieldReason, _ := unmarshalCallDiscardReason(objMap["reason"])
	callStateDiscarded.Reason = fieldReason

	return nil
}

// GetCallStateEnum return the enum type of this object
func (callStateDiscarded *CallStateDiscarded) GetCallStateEnum() CallStateEnum {
	return CallStateDiscardedType
}

// CallStateError The call has ended with an error
type CallStateError struct {
	tdCommon
	Error Error `json:"error"` // Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
}

// MessageType return the string telegram-type of CallStateError
func (callStateError *CallStateError) MessageType() string {
	return "callStateError"
}

// GetCallStateEnum return the enum type of this object
func (callStateError *CallStateError) GetCallStateEnum() CallStateEnum {
	return CallStateErrorType
}

// Call Describes a call
type Call struct {
	tdCommon
	ID         int32     `json:"id"`          // Call identifier, not persistent
	UserID     int32     `json:"user_id"`     // Peer user identifier
	IsOutgoing bool      `json:"is_outgoing"` // True, if the call is outgoing
	State      CallState `json:"state"`       // Call state
}

// MessageType return the string telegram-type of Call
func (call *Call) MessageType() string {
	return "call"
}

// UnmarshalJSON unmarshal to json
func (call *Call) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID         int32 `json:"id"`          // Call identifier, not persistent
		UserID     int32 `json:"user_id"`     // Peer user identifier
		IsOutgoing bool  `json:"is_outgoing"` // True, if the call is outgoing

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	call.tdCommon = tempObj.tdCommon
	call.ID = tempObj.ID
	call.UserID = tempObj.UserID
	call.IsOutgoing = tempObj.IsOutgoing

	fieldState, _ := unmarshalCallState(objMap["state"])
	call.State = fieldState

	return nil
}

// Animations Represents a list of animations
type Animations struct {
	tdCommon
	Animations []Animation `json:"animations"` // List of animations
}

// MessageType return the string telegram-type of Animations
func (animations *Animations) MessageType() string {
	return "animations"
}

// ImportedContacts Represents the result of an ImportContacts request
type ImportedContacts struct {
	tdCommon
	UserIDs       []int32 `json:"user_ids"`       // User identifiers of the imported contacts in the same order as they were specified in the request; 0 if the contact is not yet a registered user
	ImporterCount []int32 `json:"importer_count"` // The number of users that imported the corresponding contact; 0 for already registered users or if unavailable
}

// MessageType return the string telegram-type of ImportedContacts
func (importedContacts *ImportedContacts) MessageType() string {
	return "importedContacts"
}

// InputInlineQueryResultAnimatedGif Represents a link to an animated GIF
type InputInlineQueryResultAnimatedGif struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the query result
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the static result thumbnail (JPEG or GIF), if it exists
	GifURL              string              `json:"gif_url"`               // The URL of the GIF-file (file size must not exceed 1MB)
	GifDuration         int32               `json:"gif_duration"`          // Duration of the GIF, in seconds
	GifWidth            int32               `json:"gif_width"`             // Width of the GIF
	GifHeight           int32               `json:"gif_height"`            // Height of the GIF
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAnimatedGif
func (inputInlineQueryResultAnimatedGif *InputInlineQueryResultAnimatedGif) MessageType() string {
	return "inputInlineQueryResultAnimatedGif"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultAnimatedGif *InputInlineQueryResultAnimatedGif) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID           string `json:"id"`            // Unique identifier of the query result
		Title        string `json:"title"`         // Title of the query result
		ThumbnailURL string `json:"thumbnail_url"` // URL of the static result thumbnail (JPEG or GIF), if it exists
		GifURL       string `json:"gif_url"`       // The URL of the GIF-file (file size must not exceed 1MB)
		GifDuration  int32  `json:"gif_duration"`  // Duration of the GIF, in seconds
		GifWidth     int32  `json:"gif_width"`     // Width of the GIF
		GifHeight    int32  `json:"gif_height"`    // Height of the GIF

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultAnimatedGif.tdCommon = tempObj.tdCommon
	inputInlineQueryResultAnimatedGif.ID = tempObj.ID
	inputInlineQueryResultAnimatedGif.Title = tempObj.Title
	inputInlineQueryResultAnimatedGif.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultAnimatedGif.GifURL = tempObj.GifURL
	inputInlineQueryResultAnimatedGif.GifDuration = tempObj.GifDuration
	inputInlineQueryResultAnimatedGif.GifWidth = tempObj.GifWidth
	inputInlineQueryResultAnimatedGif.GifHeight = tempObj.GifHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultAnimatedGif.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultAnimatedGif.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultAnimatedGif *InputInlineQueryResultAnimatedGif) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultAnimatedGifType
}

// InputInlineQueryResultAnimatedMpeg4 Represents a link to an animated (i.e. without sound) H.264/MPEG-4 AVC video
type InputInlineQueryResultAnimatedMpeg4 struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the result
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the static result thumbnail (JPEG or GIF), if it exists
	Mpeg4URL            string              `json:"mpeg4_url"`             // The URL of the MPEG4-file (file size must not exceed 1MB)
	Mpeg4Duration       int32               `json:"mpeg4_duration"`        // Duration of the video, in seconds
	Mpeg4Width          int32               `json:"mpeg4_width"`           // Width of the video
	Mpeg4Height         int32               `json:"mpeg4_height"`          // Height of the video
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAnimatedMpeg4
func (inputInlineQueryResultAnimatedMpeg4 *InputInlineQueryResultAnimatedMpeg4) MessageType() string {
	return "inputInlineQueryResultAnimatedMpeg4"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultAnimatedMpeg4 *InputInlineQueryResultAnimatedMpeg4) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID            string `json:"id"`             // Unique identifier of the query result
		Title         string `json:"title"`          // Title of the result
		ThumbnailURL  string `json:"thumbnail_url"`  // URL of the static result thumbnail (JPEG or GIF), if it exists
		Mpeg4URL      string `json:"mpeg4_url"`      // The URL of the MPEG4-file (file size must not exceed 1MB)
		Mpeg4Duration int32  `json:"mpeg4_duration"` // Duration of the video, in seconds
		Mpeg4Width    int32  `json:"mpeg4_width"`    // Width of the video
		Mpeg4Height   int32  `json:"mpeg4_height"`   // Height of the video

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultAnimatedMpeg4.tdCommon = tempObj.tdCommon
	inputInlineQueryResultAnimatedMpeg4.ID = tempObj.ID
	inputInlineQueryResultAnimatedMpeg4.Title = tempObj.Title
	inputInlineQueryResultAnimatedMpeg4.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultAnimatedMpeg4.Mpeg4URL = tempObj.Mpeg4URL
	inputInlineQueryResultAnimatedMpeg4.Mpeg4Duration = tempObj.Mpeg4Duration
	inputInlineQueryResultAnimatedMpeg4.Mpeg4Width = tempObj.Mpeg4Width
	inputInlineQueryResultAnimatedMpeg4.Mpeg4Height = tempObj.Mpeg4Height

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultAnimatedMpeg4.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultAnimatedMpeg4.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultAnimatedMpeg4 *InputInlineQueryResultAnimatedMpeg4) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultAnimatedMpeg4Type
}

// InputInlineQueryResultArticle Represents a link to an article or web page
type InputInlineQueryResultArticle struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	URL                 string              `json:"url"`                   // URL of the result, if it exists
	HideURL             bool                `json:"hide_url"`              // True, if the URL must be not shown
	Title               string              `json:"title"`                 // Title of the result
	Description         string              `json:"description"`           //
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultArticle
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) MessageType() string {
	return "inputInlineQueryResultArticle"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              string `json:"id"`               // Unique identifier of the query result
		URL             string `json:"url"`              // URL of the result, if it exists
		HideURL         bool   `json:"hide_url"`         // True, if the URL must be not shown
		Title           string `json:"title"`            // Title of the result
		Description     string `json:"description"`      //
		ThumbnailURL    string `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32  `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultArticle.tdCommon = tempObj.tdCommon
	inputInlineQueryResultArticle.ID = tempObj.ID
	inputInlineQueryResultArticle.URL = tempObj.URL
	inputInlineQueryResultArticle.HideURL = tempObj.HideURL
	inputInlineQueryResultArticle.Title = tempObj.Title
	inputInlineQueryResultArticle.Description = tempObj.Description
	inputInlineQueryResultArticle.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultArticle.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultArticle.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultArticle.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultArticle.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultArticleType
}

// InputInlineQueryResultAudio Represents a link to an MP3 audio file
type InputInlineQueryResultAudio struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the audio file
	Performer           string              `json:"performer"`             // Performer of the audio file
	AudioURL            string              `json:"audio_url"`             // The URL of the audio file
	AudioDuration       int32               `json:"audio_duration"`        // Audio file duration, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAudio, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAudio
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) MessageType() string {
	return "inputInlineQueryResultAudio"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID            string `json:"id"`             // Unique identifier of the query result
		Title         string `json:"title"`          // Title of the audio file
		Performer     string `json:"performer"`      // Performer of the audio file
		AudioURL      string `json:"audio_url"`      // The URL of the audio file
		AudioDuration int32  `json:"audio_duration"` // Audio file duration, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultAudio.tdCommon = tempObj.tdCommon
	inputInlineQueryResultAudio.ID = tempObj.ID
	inputInlineQueryResultAudio.Title = tempObj.Title
	inputInlineQueryResultAudio.Performer = tempObj.Performer
	inputInlineQueryResultAudio.AudioURL = tempObj.AudioURL
	inputInlineQueryResultAudio.AudioDuration = tempObj.AudioDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultAudio.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultAudio.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultAudioType
}

// InputInlineQueryResultContact Represents a user contact
type InputInlineQueryResultContact struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Contact             Contact             `json:"contact"`               // User contact
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultContact
func (inputInlineQueryResultContact *InputInlineQueryResultContact) MessageType() string {
	return "inputInlineQueryResultContact"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultContact *InputInlineQueryResultContact) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              string  `json:"id"`               // Unique identifier of the query result
		Contact         Contact `json:"contact"`          // User contact
		ThumbnailURL    string  `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32   `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32   `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultContact.tdCommon = tempObj.tdCommon
	inputInlineQueryResultContact.ID = tempObj.ID
	inputInlineQueryResultContact.Contact = tempObj.Contact
	inputInlineQueryResultContact.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultContact.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultContact.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultContact.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultContact.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultContact *InputInlineQueryResultContact) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultContactType
}

// InputInlineQueryResultDocument Represents a link to a file
type InputInlineQueryResultDocument struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the resulting file
	Description         string              `json:"description"`           //
	DocumentURL         string              `json:"document_url"`          // URL of the file
	MimeType            string              `json:"mime_type"`             // MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
	ThumbnailURL        string              `json:"thumbnail_url"`         // The URL of the file thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Width of the thumbnail
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Height of the thumbnail
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageDocument, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultDocument
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) MessageType() string {
	return "inputInlineQueryResultDocument"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              string `json:"id"`               // Unique identifier of the query result
		Title           string `json:"title"`            // Title of the resulting file
		Description     string `json:"description"`      //
		DocumentURL     string `json:"document_url"`     // URL of the file
		MimeType        string `json:"mime_type"`        // MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
		ThumbnailURL    string `json:"thumbnail_url"`    // The URL of the file thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Width of the thumbnail
		ThumbnailHeight int32  `json:"thumbnail_height"` // Height of the thumbnail

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultDocument.tdCommon = tempObj.tdCommon
	inputInlineQueryResultDocument.ID = tempObj.ID
	inputInlineQueryResultDocument.Title = tempObj.Title
	inputInlineQueryResultDocument.Description = tempObj.Description
	inputInlineQueryResultDocument.DocumentURL = tempObj.DocumentURL
	inputInlineQueryResultDocument.MimeType = tempObj.MimeType
	inputInlineQueryResultDocument.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultDocument.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultDocument.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultDocument.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultDocument.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultDocumentType
}

// InputInlineQueryResultGame Represents a game
type InputInlineQueryResultGame struct {
	tdCommon
	ID            string      `json:"id"`              // Unique identifier of the query result
	GameShortName string      `json:"game_short_name"` // Short name of the game
	ReplyMarkup   ReplyMarkup `json:"reply_markup"`    // Message reply markup. Must be of type replyMarkupInlineKeyboard or null
}

// MessageType return the string telegram-type of InputInlineQueryResultGame
func (inputInlineQueryResultGame *InputInlineQueryResultGame) MessageType() string {
	return "inputInlineQueryResultGame"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultGame *InputInlineQueryResultGame) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID            string `json:"id"`              // Unique identifier of the query result
		GameShortName string `json:"game_short_name"` // Short name of the game

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultGame.tdCommon = tempObj.tdCommon
	inputInlineQueryResultGame.ID = tempObj.ID
	inputInlineQueryResultGame.GameShortName = tempObj.GameShortName

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultGame.ReplyMarkup = fieldReplyMarkup

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultGame *InputInlineQueryResultGame) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultGameType
}

// InputInlineQueryResultLocation Represents a point on the map
type InputInlineQueryResultLocation struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Location            Location            `json:"location"`              // Location result
	LivePeriod          int32               `json:"live_period"`           // Amount of time relative to the message sent time until the location can be updated, in seconds
	Title               string              `json:"title"`                 // Title of the result
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultLocation
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) MessageType() string {
	return "inputInlineQueryResultLocation"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              string   `json:"id"`               // Unique identifier of the query result
		Location        Location `json:"location"`         // Location result
		LivePeriod      int32    `json:"live_period"`      // Amount of time relative to the message sent time until the location can be updated, in seconds
		Title           string   `json:"title"`            // Title of the result
		ThumbnailURL    string   `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32    `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32    `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultLocation.tdCommon = tempObj.tdCommon
	inputInlineQueryResultLocation.ID = tempObj.ID
	inputInlineQueryResultLocation.Location = tempObj.Location
	inputInlineQueryResultLocation.LivePeriod = tempObj.LivePeriod
	inputInlineQueryResultLocation.Title = tempObj.Title
	inputInlineQueryResultLocation.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultLocation.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultLocation.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultLocation.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultLocation.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultLocationType
}

// InputInlineQueryResultPhoto Represents link to a JPEG image
type InputInlineQueryResultPhoto struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the result, if known
	Description         string              `json:"description"`           //
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the photo thumbnail, if it exists
	PhotoURL            string              `json:"photo_url"`             // The URL of the JPEG photo (photo size must not exceed 5MB)
	PhotoWidth          int32               `json:"photo_width"`           // Width of the photo
	PhotoHeight         int32               `json:"photo_height"`          // Height of the photo
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessagePhoto, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultPhoto
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) MessageType() string {
	return "inputInlineQueryResultPhoto"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID           string `json:"id"`            // Unique identifier of the query result
		Title        string `json:"title"`         // Title of the result, if known
		Description  string `json:"description"`   //
		ThumbnailURL string `json:"thumbnail_url"` // URL of the photo thumbnail, if it exists
		PhotoURL     string `json:"photo_url"`     // The URL of the JPEG photo (photo size must not exceed 5MB)
		PhotoWidth   int32  `json:"photo_width"`   // Width of the photo
		PhotoHeight  int32  `json:"photo_height"`  // Height of the photo

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultPhoto.tdCommon = tempObj.tdCommon
	inputInlineQueryResultPhoto.ID = tempObj.ID
	inputInlineQueryResultPhoto.Title = tempObj.Title
	inputInlineQueryResultPhoto.Description = tempObj.Description
	inputInlineQueryResultPhoto.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultPhoto.PhotoURL = tempObj.PhotoURL
	inputInlineQueryResultPhoto.PhotoWidth = tempObj.PhotoWidth
	inputInlineQueryResultPhoto.PhotoHeight = tempObj.PhotoHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultPhoto.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultPhoto.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultPhotoType
}

// InputInlineQueryResultSticker Represents a link to a WEBP sticker
type InputInlineQueryResultSticker struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the sticker thumbnail, if it exists
	StickerURL          string              `json:"sticker_url"`           // The URL of the WEBP sticker (sticker file size must not exceed 5MB)
	StickerWidth        int32               `json:"sticker_width"`         // Width of the sticker
	StickerHeight       int32               `json:"sticker_height"`        // Height of the sticker
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, inputMessageSticker, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultSticker
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) MessageType() string {
	return "inputInlineQueryResultSticker"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID            string `json:"id"`             // Unique identifier of the query result
		ThumbnailURL  string `json:"thumbnail_url"`  // URL of the sticker thumbnail, if it exists
		StickerURL    string `json:"sticker_url"`    // The URL of the WEBP sticker (sticker file size must not exceed 5MB)
		StickerWidth  int32  `json:"sticker_width"`  // Width of the sticker
		StickerHeight int32  `json:"sticker_height"` // Height of the sticker

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultSticker.tdCommon = tempObj.tdCommon
	inputInlineQueryResultSticker.ID = tempObj.ID
	inputInlineQueryResultSticker.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultSticker.StickerURL = tempObj.StickerURL
	inputInlineQueryResultSticker.StickerWidth = tempObj.StickerWidth
	inputInlineQueryResultSticker.StickerHeight = tempObj.StickerHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultSticker.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultSticker.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultStickerType
}

// InputInlineQueryResultVenue Represents information about a venue
type InputInlineQueryResultVenue struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Venue               Venue               `json:"venue"`                 // Venue result
	ThumbnailURL        string              `json:"thumbnail_url"`         // URL of the result thumbnail, if it exists
	ThumbnailWidth      int32               `json:"thumbnail_width"`       // Thumbnail width, if known
	ThumbnailHeight     int32               `json:"thumbnail_height"`      // Thumbnail height, if known
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVenue
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) MessageType() string {
	return "inputInlineQueryResultVenue"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              string `json:"id"`               // Unique identifier of the query result
		Venue           Venue  `json:"venue"`            // Venue result
		ThumbnailURL    string `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32  `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32  `json:"thumbnail_height"` // Thumbnail height, if known

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVenue.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVenue.ID = tempObj.ID
	inputInlineQueryResultVenue.Venue = tempObj.Venue
	inputInlineQueryResultVenue.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultVenue.ThumbnailWidth = tempObj.ThumbnailWidth
	inputInlineQueryResultVenue.ThumbnailHeight = tempObj.ThumbnailHeight

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVenue.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVenue.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVenueType
}

// InputInlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file
type InputInlineQueryResultVideo struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the result
	Description         string              `json:"description"`           //
	ThumbnailURL        string              `json:"thumbnail_url"`         // The URL of the video thumbnail (JPEG), if it exists
	VideoURL            string              `json:"video_url"`             // URL of the embedded video player or video file
	MimeType            string              `json:"mime_type"`             // MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
	VideoWidth          int32               `json:"video_width"`           // Width of the video
	VideoHeight         int32               `json:"video_height"`          // Height of the video
	VideoDuration       int32               `json:"video_duration"`        // Video duration, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVideo, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVideo
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) MessageType() string {
	return "inputInlineQueryResultVideo"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID            string `json:"id"`             // Unique identifier of the query result
		Title         string `json:"title"`          // Title of the result
		Description   string `json:"description"`    //
		ThumbnailURL  string `json:"thumbnail_url"`  // The URL of the video thumbnail (JPEG), if it exists
		VideoURL      string `json:"video_url"`      // URL of the embedded video player or video file
		MimeType      string `json:"mime_type"`      // MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
		VideoWidth    int32  `json:"video_width"`    // Width of the video
		VideoHeight   int32  `json:"video_height"`   // Height of the video
		VideoDuration int32  `json:"video_duration"` // Video duration, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVideo.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVideo.ID = tempObj.ID
	inputInlineQueryResultVideo.Title = tempObj.Title
	inputInlineQueryResultVideo.Description = tempObj.Description
	inputInlineQueryResultVideo.ThumbnailURL = tempObj.ThumbnailURL
	inputInlineQueryResultVideo.VideoURL = tempObj.VideoURL
	inputInlineQueryResultVideo.MimeType = tempObj.MimeType
	inputInlineQueryResultVideo.VideoWidth = tempObj.VideoWidth
	inputInlineQueryResultVideo.VideoHeight = tempObj.VideoHeight
	inputInlineQueryResultVideo.VideoDuration = tempObj.VideoDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVideo.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVideo.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVideoType
}

// InputInlineQueryResultVoiceNote Represents a link to an opus-encoded audio file within an OGG container, single channel audio
type InputInlineQueryResultVoiceNote struct {
	tdCommon
	ID                  string              `json:"id"`                    // Unique identifier of the query result
	Title               string              `json:"title"`                 // Title of the voice note
	VoiceNoteURL        string              `json:"voice_note_url"`        // The URL of the voice note file
	VoiceNoteDuration   int32               `json:"voice_note_duration"`   // Duration of the voice note, in seconds
	ReplyMarkup         ReplyMarkup         `json:"reply_markup"`          // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
	InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVoiceNote, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVoiceNote
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) MessageType() string {
	return "inputInlineQueryResultVoiceNote"
}

// UnmarshalJSON unmarshal to json
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                string `json:"id"`                  // Unique identifier of the query result
		Title             string `json:"title"`               // Title of the voice note
		VoiceNoteURL      string `json:"voice_note_url"`      // The URL of the voice note file
		VoiceNoteDuration int32  `json:"voice_note_duration"` // Duration of the voice note, in seconds

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputInlineQueryResultVoiceNote.tdCommon = tempObj.tdCommon
	inputInlineQueryResultVoiceNote.ID = tempObj.ID
	inputInlineQueryResultVoiceNote.Title = tempObj.Title
	inputInlineQueryResultVoiceNote.VoiceNoteURL = tempObj.VoiceNoteURL
	inputInlineQueryResultVoiceNote.VoiceNoteDuration = tempObj.VoiceNoteDuration

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	inputInlineQueryResultVoiceNote.ReplyMarkup = fieldReplyMarkup

	fieldInputMessageContent, _ := unmarshalInputMessageContent(objMap["input_message_content"])
	inputInlineQueryResultVoiceNote.InputMessageContent = fieldInputMessageContent

	return nil
}

// GetInputInlineQueryResultEnum return the enum type of this object
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) GetInputInlineQueryResultEnum() InputInlineQueryResultEnum {
	return InputInlineQueryResultVoiceNoteType
}

// InlineQueryResultArticle Represents a link to an article or web page
type InlineQueryResultArticle struct {
	tdCommon
	ID          string    `json:"id"`          // Unique identifier of the query result
	URL         string    `json:"url"`         // URL of the result, if it exists
	HideURL     bool      `json:"hide_url"`    // True, if the URL must be not shown
	Title       string    `json:"title"`       // Title of the result
	Description string    `json:"description"` //
	Thumbnail   PhotoSize `json:"thumbnail"`   // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultArticle
func (inlineQueryResultArticle *InlineQueryResultArticle) MessageType() string {
	return "inlineQueryResultArticle"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultArticle *InlineQueryResultArticle) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultArticleType
}

// InlineQueryResultContact Represents a user contact
type InlineQueryResultContact struct {
	tdCommon
	ID        string    `json:"id"`        // Unique identifier of the query result
	Contact   Contact   `json:"contact"`   // A user contact
	Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultContact
func (inlineQueryResultContact *InlineQueryResultContact) MessageType() string {
	return "inlineQueryResultContact"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultContact *InlineQueryResultContact) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultContactType
}

// InlineQueryResultLocation Represents a point on the map
type InlineQueryResultLocation struct {
	tdCommon
	ID        string    `json:"id"`        // Unique identifier of the query result
	Location  Location  `json:"location"`  // Location result
	Title     string    `json:"title"`     // Title of the result
	Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultLocation
func (inlineQueryResultLocation *InlineQueryResultLocation) MessageType() string {
	return "inlineQueryResultLocation"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultLocation *InlineQueryResultLocation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultLocationType
}

// InlineQueryResultVenue Represents information about a venue
type InlineQueryResultVenue struct {
	tdCommon
	ID        string    `json:"id"`        // Unique identifier of the query result
	Venue     Venue     `json:"venue"`     // Venue result
	Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultVenue
func (inlineQueryResultVenue *InlineQueryResultVenue) MessageType() string {
	return "inlineQueryResultVenue"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVenue *InlineQueryResultVenue) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVenueType
}

// InlineQueryResultGame Represents information about a game
type InlineQueryResultGame struct {
	tdCommon
	ID   string `json:"id"`   // Unique identifier of the query result
	Game Game   `json:"game"` // Game result
}

// MessageType return the string telegram-type of InlineQueryResultGame
func (inlineQueryResultGame *InlineQueryResultGame) MessageType() string {
	return "inlineQueryResultGame"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultGame *InlineQueryResultGame) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultGameType
}

// InlineQueryResultAnimation Represents an animation file
type InlineQueryResultAnimation struct {
	tdCommon
	ID        string    `json:"id"`        // Unique identifier of the query result
	Animation Animation `json:"animation"` // Animation file
	Title     string    `json:"title"`     // Animation title
}

// MessageType return the string telegram-type of InlineQueryResultAnimation
func (inlineQueryResultAnimation *InlineQueryResultAnimation) MessageType() string {
	return "inlineQueryResultAnimation"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAnimation *InlineQueryResultAnimation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAnimationType
}

// InlineQueryResultAudio Represents an audio file
type InlineQueryResultAudio struct {
	tdCommon
	ID    string `json:"id"`    // Unique identifier of the query result
	Audio Audio  `json:"audio"` // Audio file
}

// MessageType return the string telegram-type of InlineQueryResultAudio
func (inlineQueryResultAudio *InlineQueryResultAudio) MessageType() string {
	return "inlineQueryResultAudio"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAudio *InlineQueryResultAudio) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAudioType
}

// InlineQueryResultDocument Represents a document
type InlineQueryResultDocument struct {
	tdCommon
	ID          string   `json:"id"`          // Unique identifier of the query result
	Document    Document `json:"document"`    // Document
	Title       string   `json:"title"`       // Document title
	Description string   `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultDocument
func (inlineQueryResultDocument *InlineQueryResultDocument) MessageType() string {
	return "inlineQueryResultDocument"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultDocument *InlineQueryResultDocument) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultDocumentType
}

// InlineQueryResultPhoto Represents a photo
type InlineQueryResultPhoto struct {
	tdCommon
	ID          string `json:"id"`          // Unique identifier of the query result
	Photo       Photo  `json:"photo"`       // Photo
	Title       string `json:"title"`       // Title of the result, if known
	Description string `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultPhoto
func (inlineQueryResultPhoto *InlineQueryResultPhoto) MessageType() string {
	return "inlineQueryResultPhoto"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultPhoto *InlineQueryResultPhoto) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultPhotoType
}

// InlineQueryResultSticker Represents a sticker
type InlineQueryResultSticker struct {
	tdCommon
	ID      string  `json:"id"`      // Unique identifier of the query result
	Sticker Sticker `json:"sticker"` // Sticker
}

// MessageType return the string telegram-type of InlineQueryResultSticker
func (inlineQueryResultSticker *InlineQueryResultSticker) MessageType() string {
	return "inlineQueryResultSticker"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultSticker *InlineQueryResultSticker) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultStickerType
}

// InlineQueryResultVideo Represents a video
type InlineQueryResultVideo struct {
	tdCommon
	ID          string `json:"id"`          // Unique identifier of the query result
	Video       Video  `json:"video"`       // Video
	Title       string `json:"title"`       // Title of the video
	Description string `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultVideo
func (inlineQueryResultVideo *InlineQueryResultVideo) MessageType() string {
	return "inlineQueryResultVideo"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVideo *InlineQueryResultVideo) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVideoType
}

// InlineQueryResultVoiceNote Represents a voice note
type InlineQueryResultVoiceNote struct {
	tdCommon
	ID        string    `json:"id"`         // Unique identifier of the query result
	VoiceNote VoiceNote `json:"voice_note"` // Voice note
	Title     string    `json:"title"`      // Title of the voice note
}

// MessageType return the string telegram-type of InlineQueryResultVoiceNote
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) MessageType() string {
	return "inlineQueryResultVoiceNote"
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVoiceNoteType
}

// InlineQueryResults Represents the results of the inline query. Use sendInlineQueryResultMessage to send the result of the query
type InlineQueryResults struct {
	tdCommon
	InlineQueryID     JSONInt64           `json:"inline_query_id"`     // Unique identifier of the inline query
	NextOffset        string              `json:"next_offset"`         // The offset for the next request. If empty, there are no more results
	Results           []InlineQueryResult `json:"results"`             // Results of the query
	SwitchPmText      string              `json:"switch_pm_text"`      // If non-empty, this text should be shown on the button, which opens a private chat with the bot and sends the bot a start message with the switch_pm_parameter
	SwitchPmParameter string              `json:"switch_pm_parameter"` // Parameter for the bot start message
}

// MessageType return the string telegram-type of InlineQueryResults
func (inlineQueryResults *InlineQueryResults) MessageType() string {
	return "inlineQueryResults"
}

// CallbackQueryPayloadData The payload from a general callback button
type CallbackQueryPayloadData struct {
	tdCommon
	Data []byte `json:"data"` // Data that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadData
func (callbackQueryPayloadData *CallbackQueryPayloadData) MessageType() string {
	return "callbackQueryPayloadData"
}

// GetCallbackQueryPayloadEnum return the enum type of this object
func (callbackQueryPayloadData *CallbackQueryPayloadData) GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum {
	return CallbackQueryPayloadDataType
}

// CallbackQueryPayloadGame The payload from a game callback button
type CallbackQueryPayloadGame struct {
	tdCommon
	GameShortName string `json:"game_short_name"` // A short name of the game that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadGame
func (callbackQueryPayloadGame *CallbackQueryPayloadGame) MessageType() string {
	return "callbackQueryPayloadGame"
}

// GetCallbackQueryPayloadEnum return the enum type of this object
func (callbackQueryPayloadGame *CallbackQueryPayloadGame) GetCallbackQueryPayloadEnum() CallbackQueryPayloadEnum {
	return CallbackQueryPayloadGameType
}

// CallbackQueryAnswer Contains a bot's answer to a callback query
type CallbackQueryAnswer struct {
	tdCommon
	Text      string `json:"text"`       // Text of the answer
	ShowAlert bool   `json:"show_alert"` // True, if an alert should be shown to the user instead of a toast notification
	URL       string `json:"url"`        // URL to be opened
}

// MessageType return the string telegram-type of CallbackQueryAnswer
func (callbackQueryAnswer *CallbackQueryAnswer) MessageType() string {
	return "callbackQueryAnswer"
}

// CustomRequestResult Contains the result of a custom request
type CustomRequestResult struct {
	tdCommon
	Result string `json:"result"` // A JSON-serialized result
}

// MessageType return the string telegram-type of CustomRequestResult
func (customRequestResult *CustomRequestResult) MessageType() string {
	return "customRequestResult"
}

// GameHighScore Contains one row of the game high score table
type GameHighScore struct {
	tdCommon
	Position int32 `json:"position"` // Position in the high score table
	UserID   int32 `json:"user_id"`  // User identifier
	Score    int32 `json:"score"`    // User score
}

// MessageType return the string telegram-type of GameHighScore
func (gameHighScore *GameHighScore) MessageType() string {
	return "gameHighScore"
}

// GameHighScores Contains a list of game high scores
type GameHighScores struct {
	tdCommon
	Scores []GameHighScore `json:"scores"` // A list of game high scores
}

// MessageType return the string telegram-type of GameHighScores
func (gameHighScores *GameHighScores) MessageType() string {
	return "gameHighScores"
}

// ChatEventMessageEdited A message was edited
type ChatEventMessageEdited struct {
	tdCommon
	OldMessage Message `json:"old_message"` // The original message before the edit
	NewMessage Message `json:"new_message"` // The message after it was edited
}

// MessageType return the string telegram-type of ChatEventMessageEdited
func (chatEventMessageEdited *ChatEventMessageEdited) MessageType() string {
	return "chatEventMessageEdited"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageEdited *ChatEventMessageEdited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageEditedType
}

// ChatEventMessageDeleted A message was deleted
type ChatEventMessageDeleted struct {
	tdCommon
	Message Message `json:"message"` // Deleted message
}

// MessageType return the string telegram-type of ChatEventMessageDeleted
func (chatEventMessageDeleted *ChatEventMessageDeleted) MessageType() string {
	return "chatEventMessageDeleted"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageDeleted *ChatEventMessageDeleted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageDeletedType
}

// ChatEventMessagePinned A message was pinned
type ChatEventMessagePinned struct {
	tdCommon
	Message Message `json:"message"` // Pinned message
}

// MessageType return the string telegram-type of ChatEventMessagePinned
func (chatEventMessagePinned *ChatEventMessagePinned) MessageType() string {
	return "chatEventMessagePinned"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessagePinned *ChatEventMessagePinned) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessagePinnedType
}

// ChatEventMessageUnpinned A message was unpinned
type ChatEventMessageUnpinned struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatEventMessageUnpinned
func (chatEventMessageUnpinned *ChatEventMessageUnpinned) MessageType() string {
	return "chatEventMessageUnpinned"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageUnpinned *ChatEventMessageUnpinned) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageUnpinnedType
}

// ChatEventMemberJoined A new member joined the chat
type ChatEventMemberJoined struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatEventMemberJoined
func (chatEventMemberJoined *ChatEventMemberJoined) MessageType() string {
	return "chatEventMemberJoined"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberJoined *ChatEventMemberJoined) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberJoinedType
}

// ChatEventMemberLeft A member left the chat
type ChatEventMemberLeft struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatEventMemberLeft
func (chatEventMemberLeft *ChatEventMemberLeft) MessageType() string {
	return "chatEventMemberLeft"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberLeft *ChatEventMemberLeft) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberLeftType
}

// ChatEventMemberInvited A new chat member was invited
type ChatEventMemberInvited struct {
	tdCommon
	UserID int32            `json:"user_id"` // New member user identifier
	Status ChatMemberStatus `json:"status"`  // New member status
}

// MessageType return the string telegram-type of ChatEventMemberInvited
func (chatEventMemberInvited *ChatEventMemberInvited) MessageType() string {
	return "chatEventMemberInvited"
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberInvited *ChatEventMemberInvited) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserID int32 `json:"user_id"` // New member user identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberInvited.tdCommon = tempObj.tdCommon
	chatEventMemberInvited.UserID = tempObj.UserID

	fieldStatus, _ := unmarshalChatMemberStatus(objMap["status"])
	chatEventMemberInvited.Status = fieldStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberInvited *ChatEventMemberInvited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberInvitedType
}

// ChatEventMemberPromoted A chat member has gained/lost administrator status, or the list of their administrator privileges has changed
type ChatEventMemberPromoted struct {
	tdCommon
	UserID    int32            `json:"user_id"`    // Chat member user identifier
	OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member
	NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
}

// MessageType return the string telegram-type of ChatEventMemberPromoted
func (chatEventMemberPromoted *ChatEventMemberPromoted) MessageType() string {
	return "chatEventMemberPromoted"
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberPromoted *ChatEventMemberPromoted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserID int32 `json:"user_id"` // Chat member user identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberPromoted.tdCommon = tempObj.tdCommon
	chatEventMemberPromoted.UserID = tempObj.UserID

	fieldOldStatus, _ := unmarshalChatMemberStatus(objMap["old_status"])
	chatEventMemberPromoted.OldStatus = fieldOldStatus

	fieldNewStatus, _ := unmarshalChatMemberStatus(objMap["new_status"])
	chatEventMemberPromoted.NewStatus = fieldNewStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberPromoted *ChatEventMemberPromoted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberPromotedType
}

// ChatEventMemberRestricted A chat member was restricted/unrestricted or banned/unbanned, or the list of their restrictions has changed
type ChatEventMemberRestricted struct {
	tdCommon
	UserID    int32            `json:"user_id"`    // Chat member user identifier
	OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member
	NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
}

// MessageType return the string telegram-type of ChatEventMemberRestricted
func (chatEventMemberRestricted *ChatEventMemberRestricted) MessageType() string {
	return "chatEventMemberRestricted"
}

// UnmarshalJSON unmarshal to json
func (chatEventMemberRestricted *ChatEventMemberRestricted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserID int32 `json:"user_id"` // Chat member user identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEventMemberRestricted.tdCommon = tempObj.tdCommon
	chatEventMemberRestricted.UserID = tempObj.UserID

	fieldOldStatus, _ := unmarshalChatMemberStatus(objMap["old_status"])
	chatEventMemberRestricted.OldStatus = fieldOldStatus

	fieldNewStatus, _ := unmarshalChatMemberStatus(objMap["new_status"])
	chatEventMemberRestricted.NewStatus = fieldNewStatus

	return nil
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMemberRestricted *ChatEventMemberRestricted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMemberRestrictedType
}

// ChatEventTitleChanged The chat title was changed
type ChatEventTitleChanged struct {
	tdCommon
	OldTitle string `json:"old_title"` // Previous chat title
	NewTitle string `json:"new_title"` // New chat title
}

// MessageType return the string telegram-type of ChatEventTitleChanged
func (chatEventTitleChanged *ChatEventTitleChanged) MessageType() string {
	return "chatEventTitleChanged"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventTitleChanged *ChatEventTitleChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventTitleChangedType
}

// ChatEventDescriptionChanged The chat description was changed
type ChatEventDescriptionChanged struct {
	tdCommon
	OldDescription string `json:"old_description"` // Previous chat description
	NewDescription string `json:"new_description"` // New chat description
}

// MessageType return the string telegram-type of ChatEventDescriptionChanged
func (chatEventDescriptionChanged *ChatEventDescriptionChanged) MessageType() string {
	return "chatEventDescriptionChanged"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventDescriptionChanged *ChatEventDescriptionChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventDescriptionChangedType
}

// ChatEventUsernameChanged The chat username was changed
type ChatEventUsernameChanged struct {
	tdCommon
	OldUsername string `json:"old_username"` // Previous chat username
	NewUsername string `json:"new_username"` // New chat username
}

// MessageType return the string telegram-type of ChatEventUsernameChanged
func (chatEventUsernameChanged *ChatEventUsernameChanged) MessageType() string {
	return "chatEventUsernameChanged"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventUsernameChanged *ChatEventUsernameChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventUsernameChangedType
}

// ChatEventPhotoChanged The chat photo was changed
type ChatEventPhotoChanged struct {
	tdCommon
	OldPhoto ChatPhoto `json:"old_photo"` // Previous chat photo value; may be null
	NewPhoto ChatPhoto `json:"new_photo"` // New chat photo value; may be null
}

// MessageType return the string telegram-type of ChatEventPhotoChanged
func (chatEventPhotoChanged *ChatEventPhotoChanged) MessageType() string {
	return "chatEventPhotoChanged"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventPhotoChanged *ChatEventPhotoChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventPhotoChangedType
}

// ChatEventInvitesToggled The anyone_can_invite setting of a supergroup chat was toggled
type ChatEventInvitesToggled struct {
	tdCommon
	AnyoneCanInvite bool `json:"anyone_can_invite"` // New value of anyone_can_invite
}

// MessageType return the string telegram-type of ChatEventInvitesToggled
func (chatEventInvitesToggled *ChatEventInvitesToggled) MessageType() string {
	return "chatEventInvitesToggled"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventInvitesToggled *ChatEventInvitesToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventInvitesToggledType
}

// ChatEventSignMessagesToggled The sign_messages setting of a channel was toggled
type ChatEventSignMessagesToggled struct {
	tdCommon
	SignMessages bool `json:"sign_messages"` // New value of sign_messages
}

// MessageType return the string telegram-type of ChatEventSignMessagesToggled
func (chatEventSignMessagesToggled *ChatEventSignMessagesToggled) MessageType() string {
	return "chatEventSignMessagesToggled"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventSignMessagesToggled *ChatEventSignMessagesToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventSignMessagesToggledType
}

// ChatEventStickerSetChanged The supergroup sticker set was changed
type ChatEventStickerSetChanged struct {
	tdCommon
	OldStickerSetID JSONInt64 `json:"old_sticker_set_id"` // Previous identifier of the chat sticker set; 0 if none
	NewStickerSetID JSONInt64 `json:"new_sticker_set_id"` // New identifier of the chat sticker set; 0 if none
}

// MessageType return the string telegram-type of ChatEventStickerSetChanged
func (chatEventStickerSetChanged *ChatEventStickerSetChanged) MessageType() string {
	return "chatEventStickerSetChanged"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventStickerSetChanged *ChatEventStickerSetChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventStickerSetChangedType
}

// ChatEventIsAllHistoryAvailableToggled The is_all_history_available setting of a supergroup was toggled
type ChatEventIsAllHistoryAvailableToggled struct {
	tdCommon
	IsAllHistoryAvailable bool `json:"is_all_history_available"` // New value of is_all_history_available
}

// MessageType return the string telegram-type of ChatEventIsAllHistoryAvailableToggled
func (chatEventIsAllHistoryAvailableToggled *ChatEventIsAllHistoryAvailableToggled) MessageType() string {
	return "chatEventIsAllHistoryAvailableToggled"
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventIsAllHistoryAvailableToggled *ChatEventIsAllHistoryAvailableToggled) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventIsAllHistoryAvailableToggledType
}

// ChatEvent Represents a chat event
type ChatEvent struct {
	tdCommon
	ID     JSONInt64       `json:"id"`      // Chat event identifier
	Date   int32           `json:"date"`    // Point in time (Unix timestamp) when the event happened
	UserID int32           `json:"user_id"` // Identifier of the user who performed the action that triggered the event
	Action ChatEventAction `json:"action"`  // Action performed by the user
}

// MessageType return the string telegram-type of ChatEvent
func (chatEvent *ChatEvent) MessageType() string {
	return "chatEvent"
}

// UnmarshalJSON unmarshal to json
func (chatEvent *ChatEvent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID     JSONInt64 `json:"id"`      // Chat event identifier
		Date   int32     `json:"date"`    // Point in time (Unix timestamp) when the event happened
		UserID int32     `json:"user_id"` // Identifier of the user who performed the action that triggered the event

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEvent.tdCommon = tempObj.tdCommon
	chatEvent.ID = tempObj.ID
	chatEvent.Date = tempObj.Date
	chatEvent.UserID = tempObj.UserID

	fieldAction, _ := unmarshalChatEventAction(objMap["action"])
	chatEvent.Action = fieldAction

	return nil
}

// ChatEvents Contains a list of chat events
type ChatEvents struct {
	tdCommon
	Events []ChatEvent `json:"events"` // List of events
}

// MessageType return the string telegram-type of ChatEvents
func (chatEvents *ChatEvents) MessageType() string {
	return "chatEvents"
}

// ChatEventLogFilters Represents a set of filters used to obtain a chat event log
type ChatEventLogFilters struct {
	tdCommon
	MessageEdits       bool `json:"message_edits"`       // True, if message edits should be returned
	MessageDeletions   bool `json:"message_deletions"`   // True, if message deletions should be returned
	MessagePins        bool `json:"message_pins"`        // True, if pin/unpin events should be returned
	MemberJoins        bool `json:"member_joins"`        // True, if members joining events should be returned
	MemberLeaves       bool `json:"member_leaves"`       // True, if members leaving events should be returned
	MemberInvites      bool `json:"member_invites"`      // True, if invited member events should be returned
	MemberPromotions   bool `json:"member_promotions"`   // True, if member promotion/demotion events should be returned
	MemberRestrictions bool `json:"member_restrictions"` // True, if member restricted/unrestricted/banned/unbanned events should be returned
	InfoChanges        bool `json:"info_changes"`        // True, if changes in chat information should be returned
	SettingChanges     bool `json:"setting_changes"`     // True, if changes in chat settings should be returned
}

// MessageType return the string telegram-type of ChatEventLogFilters
func (chatEventLogFilters *ChatEventLogFilters) MessageType() string {
	return "chatEventLogFilters"
}

// DeviceTokenGoogleCloudMessaging A token for Google Cloud Messaging
type DeviceTokenGoogleCloudMessaging struct {
	tdCommon
	Token string `json:"token"` // Device registration token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenGoogleCloudMessaging
func (deviceTokenGoogleCloudMessaging *DeviceTokenGoogleCloudMessaging) MessageType() string {
	return "deviceTokenGoogleCloudMessaging"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenGoogleCloudMessaging *DeviceTokenGoogleCloudMessaging) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenGoogleCloudMessagingType
}

// DeviceTokenApplePush A token for Apple Push Notification service
type DeviceTokenApplePush struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token, may be empty to de-register a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePush
func (deviceTokenApplePush *DeviceTokenApplePush) MessageType() string {
	return "deviceTokenApplePush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePush *DeviceTokenApplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushType
}

// DeviceTokenApplePushVoIP A token for Apple Push Notification service VoIP notifications
type DeviceTokenApplePushVoIP struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token, may be empty to de-register a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePushVoIP
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) MessageType() string {
	return "deviceTokenApplePushVoIP"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushVoIPType
}

// DeviceTokenWindowsPush A token for Windows Push Notification Services
type DeviceTokenWindowsPush struct {
	tdCommon
	AccessToken string `json:"access_token"` // The access token that will be used to send notifications, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenWindowsPush
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) MessageType() string {
	return "deviceTokenWindowsPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWindowsPushType
}

// DeviceTokenMicrosoftPush A token for Microsoft Push Notification Service
type DeviceTokenMicrosoftPush struct {
	tdCommon
	ChannelURI string `json:"channel_uri"` // Push notification channel URI, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPush
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) MessageType() string {
	return "deviceTokenMicrosoftPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushType
}

// DeviceTokenMicrosoftPushVoIP A token for Microsoft Push Notification Service VoIP channel
type DeviceTokenMicrosoftPushVoIP struct {
	tdCommon
	ChannelURI string `json:"channel_uri"` // Push notification channel URI, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPushVoIP
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) MessageType() string {
	return "deviceTokenMicrosoftPushVoIP"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushVoIPType
}

// DeviceTokenWebPush A token for web Push API
type DeviceTokenWebPush struct {
	tdCommon
	Endpoint        string `json:"endpoint"`         // Absolute URL exposed by the push service where the application server can send push messages, may be empty to de-register a device
	P256dhBase64url string `json:"p256dh_base64url"` // Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
	AuthBase64url   string `json:"auth_base64url"`   // Base64url-encoded authentication secret
}

// MessageType return the string telegram-type of DeviceTokenWebPush
func (deviceTokenWebPush *DeviceTokenWebPush) MessageType() string {
	return "deviceTokenWebPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWebPush *DeviceTokenWebPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWebPushType
}

// DeviceTokenSimplePush A token for Simple Push API for Firefox OS
type DeviceTokenSimplePush struct {
	tdCommon
	Endpoint string `json:"endpoint"` // Absolute URL exposed by the push service where the application server can send push messages, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenSimplePush
func (deviceTokenSimplePush *DeviceTokenSimplePush) MessageType() string {
	return "deviceTokenSimplePush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenSimplePush *DeviceTokenSimplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenSimplePushType
}

// DeviceTokenUbuntuPush A token for Ubuntu Push Client service
type DeviceTokenUbuntuPush struct {
	tdCommon
	Token string `json:"token"` // Token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenUbuntuPush
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) MessageType() string {
	return "deviceTokenUbuntuPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenUbuntuPushType
}

// DeviceTokenBlackBerryPush A token for BlackBerry Push Service
type DeviceTokenBlackBerryPush struct {
	tdCommon
	Token string `json:"token"` // Token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenBlackBerryPush
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) MessageType() string {
	return "deviceTokenBlackBerryPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenBlackBerryPushType
}

// DeviceTokenTizenPush A token for Tizen Push Service
type DeviceTokenTizenPush struct {
	tdCommon
	RegID string `json:"reg_id"` // Push service registration identifier, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenTizenPush
func (deviceTokenTizenPush *DeviceTokenTizenPush) MessageType() string {
	return "deviceTokenTizenPush"
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenTizenPush *DeviceTokenTizenPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenTizenPushType
}

// Wallpaper Contains information about a wallpaper
type Wallpaper struct {
	tdCommon
	ID    int32       `json:"id"`    // Unique persistent wallpaper identifier
	Sizes []PhotoSize `json:"sizes"` // Available variants of the wallpaper in different sizes. These photos can only be downloaded; they can't be sent in a message
	Color int32       `json:"color"` // Main color of the wallpaper in RGB24 format; should be treated as background color if no photos are specified
}

// MessageType return the string telegram-type of Wallpaper
func (wallpaper *Wallpaper) MessageType() string {
	return "wallpaper"
}

// Wallpapers Contains a list of wallpapers
type Wallpapers struct {
	tdCommon
	Wallpapers []Wallpaper `json:"wallpapers"` // A list of wallpapers
}

// MessageType return the string telegram-type of Wallpapers
func (wallpapers *Wallpapers) MessageType() string {
	return "wallpapers"
}

// Hashtags Contains a list of hashtags
type Hashtags struct {
	tdCommon
	Hashtags []string `json:"hashtags"` // A list of hashtags
}

// MessageType return the string telegram-type of Hashtags
func (hashtags *Hashtags) MessageType() string {
	return "hashtags"
}

// CheckChatUsernameResultOk The username can be set
type CheckChatUsernameResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultOk
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) MessageType() string {
	return "checkChatUsernameResultOk"
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultOkType
}

// CheckChatUsernameResultUsernameInvalid The username is invalid
type CheckChatUsernameResultUsernameInvalid struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameInvalid
func (checkChatUsernameResultUsernameInvalid *CheckChatUsernameResultUsernameInvalid) MessageType() string {
	return "checkChatUsernameResultUsernameInvalid"
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultUsernameInvalid *CheckChatUsernameResultUsernameInvalid) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultUsernameInvalidType
}

// CheckChatUsernameResultUsernameOccupied The username is occupied
type CheckChatUsernameResultUsernameOccupied struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameOccupied
func (checkChatUsernameResultUsernameOccupied *CheckChatUsernameResultUsernameOccupied) MessageType() string {
	return "checkChatUsernameResultUsernameOccupied"
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultUsernameOccupied *CheckChatUsernameResultUsernameOccupied) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultUsernameOccupiedType
}

// CheckChatUsernameResultPublicChatsTooMuch The user has too much public chats, one of them should be made private first
type CheckChatUsernameResultPublicChatsTooMuch struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicChatsTooMuch
func (checkChatUsernameResultPublicChatsTooMuch *CheckChatUsernameResultPublicChatsTooMuch) MessageType() string {
	return "checkChatUsernameResultPublicChatsTooMuch"
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultPublicChatsTooMuch *CheckChatUsernameResultPublicChatsTooMuch) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultPublicChatsTooMuchType
}

// CheckChatUsernameResultPublicGroupsUnavailable The user can't be a member of a public supergroup
type CheckChatUsernameResultPublicGroupsUnavailable struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicGroupsUnavailable
func (checkChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable) MessageType() string {
	return "checkChatUsernameResultPublicGroupsUnavailable"
}

// GetCheckChatUsernameResultEnum return the enum type of this object
func (checkChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable) GetCheckChatUsernameResultEnum() CheckChatUsernameResultEnum {
	return CheckChatUsernameResultPublicGroupsUnavailableType
}

// OptionValueBoolean Boolean option
type OptionValueBoolean struct {
	tdCommon
	Value bool `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueBoolean
func (optionValueBoolean *OptionValueBoolean) MessageType() string {
	return "optionValueBoolean"
}

// GetOptionValueEnum return the enum type of this object
func (optionValueBoolean *OptionValueBoolean) GetOptionValueEnum() OptionValueEnum {
	return OptionValueBooleanType
}

// OptionValueEmpty An unknown option or an option which has a default value
type OptionValueEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of OptionValueEmpty
func (optionValueEmpty *OptionValueEmpty) MessageType() string {
	return "optionValueEmpty"
}

// GetOptionValueEnum return the enum type of this object
func (optionValueEmpty *OptionValueEmpty) GetOptionValueEnum() OptionValueEnum {
	return OptionValueEmptyType
}

// OptionValueInteger An integer option
type OptionValueInteger struct {
	tdCommon
	Value int32 `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueInteger
func (optionValueInteger *OptionValueInteger) MessageType() string {
	return "optionValueInteger"
}

// GetOptionValueEnum return the enum type of this object
func (optionValueInteger *OptionValueInteger) GetOptionValueEnum() OptionValueEnum {
	return OptionValueIntegerType
}

// OptionValueString A string option
type OptionValueString struct {
	tdCommon
	Value string `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueString
func (optionValueString *OptionValueString) MessageType() string {
	return "optionValueString"
}

// GetOptionValueEnum return the enum type of this object
func (optionValueString *OptionValueString) GetOptionValueEnum() OptionValueEnum {
	return OptionValueStringType
}

// UserPrivacySettingRuleAllowAll A rule to allow all users to do something
type UserPrivacySettingRuleAllowAll struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowAll
func (userPrivacySettingRuleAllowAll *UserPrivacySettingRuleAllowAll) MessageType() string {
	return "userPrivacySettingRuleAllowAll"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowAll *UserPrivacySettingRuleAllowAll) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowAllType
}

// UserPrivacySettingRuleAllowContacts A rule to allow all of a user's contacts to do something
type UserPrivacySettingRuleAllowContacts struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowContacts
func (userPrivacySettingRuleAllowContacts *UserPrivacySettingRuleAllowContacts) MessageType() string {
	return "userPrivacySettingRuleAllowContacts"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowContacts *UserPrivacySettingRuleAllowContacts) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowContactsType
}

// UserPrivacySettingRuleAllowUsers A rule to allow certain specified users to do something
type UserPrivacySettingRuleAllowUsers struct {
	tdCommon
	UserIDs []int32 `json:"user_ids"` // The user identifiers
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowUsers
func (userPrivacySettingRuleAllowUsers *UserPrivacySettingRuleAllowUsers) MessageType() string {
	return "userPrivacySettingRuleAllowUsers"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleAllowUsers *UserPrivacySettingRuleAllowUsers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleAllowUsersType
}

// UserPrivacySettingRuleRestrictAll A rule to restrict all users from doing something
type UserPrivacySettingRuleRestrictAll struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictAll
func (userPrivacySettingRuleRestrictAll *UserPrivacySettingRuleRestrictAll) MessageType() string {
	return "userPrivacySettingRuleRestrictAll"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictAll *UserPrivacySettingRuleRestrictAll) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictAllType
}

// UserPrivacySettingRuleRestrictContacts A rule to restrict all contacts of a user from doing something
type UserPrivacySettingRuleRestrictContacts struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictContacts
func (userPrivacySettingRuleRestrictContacts *UserPrivacySettingRuleRestrictContacts) MessageType() string {
	return "userPrivacySettingRuleRestrictContacts"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictContacts *UserPrivacySettingRuleRestrictContacts) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictContactsType
}

// UserPrivacySettingRuleRestrictUsers A rule to restrict all specified users from doing something
type UserPrivacySettingRuleRestrictUsers struct {
	tdCommon
	UserIDs []int32 `json:"user_ids"` // The user identifiers
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictUsers
func (userPrivacySettingRuleRestrictUsers *UserPrivacySettingRuleRestrictUsers) MessageType() string {
	return "userPrivacySettingRuleRestrictUsers"
}

// GetUserPrivacySettingRuleEnum return the enum type of this object
func (userPrivacySettingRuleRestrictUsers *UserPrivacySettingRuleRestrictUsers) GetUserPrivacySettingRuleEnum() UserPrivacySettingRuleEnum {
	return UserPrivacySettingRuleRestrictUsersType
}

// UserPrivacySettingRules A list of privacy rules. Rules are matched in the specified order. The first matched rule defines the privacy setting for a given user. If no rule matches, the action is not allowed
type UserPrivacySettingRules struct {
	tdCommon
	Rules []UserPrivacySettingRule `json:"rules"` // A list of rules
}

// MessageType return the string telegram-type of UserPrivacySettingRules
func (userPrivacySettingRules *UserPrivacySettingRules) MessageType() string {
	return "userPrivacySettingRules"
}

// UserPrivacySettingShowStatus A privacy setting for managing whether the user's online status is visible
type UserPrivacySettingShowStatus struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowStatus
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) MessageType() string {
	return "userPrivacySettingShowStatus"
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingShowStatusType
}

// UserPrivacySettingAllowChatInvites A privacy setting for managing whether the user can be invited to chats
type UserPrivacySettingAllowChatInvites struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowChatInvites
func (userPrivacySettingAllowChatInvites *UserPrivacySettingAllowChatInvites) MessageType() string {
	return "userPrivacySettingAllowChatInvites"
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowChatInvites *UserPrivacySettingAllowChatInvites) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowChatInvitesType
}

// UserPrivacySettingAllowCalls A privacy setting for managing whether the user can be called
type UserPrivacySettingAllowCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingAllowCalls
func (userPrivacySettingAllowCalls *UserPrivacySettingAllowCalls) MessageType() string {
	return "userPrivacySettingAllowCalls"
}

// GetUserPrivacySettingEnum return the enum type of this object
func (userPrivacySettingAllowCalls *UserPrivacySettingAllowCalls) GetUserPrivacySettingEnum() UserPrivacySettingEnum {
	return UserPrivacySettingAllowCallsType
}

// AccountTTL Contains information about the period of inactivity after which the current user's account will automatically be deleted
type AccountTTL struct {
	tdCommon
	Days int32 `json:"days"` // Number of days of inactivity before the account will be flagged for deletion; should range from 30-366 days
}

// MessageType return the string telegram-type of AccountTTL
func (accountTTL *AccountTTL) MessageType() string {
	return "accountTtl"
}

// Session Contains information about one session in a Telegram application used by the current user
type Session struct {
	tdCommon
	ID                    JSONInt64 `json:"id"`                      // Session identifier
	IsCurrent             bool      `json:"is_current"`              // True, if this session is the current session
	APIID                 int32     `json:"api_id"`                  // Telegram API identifier, as provided by the application
	ApplicationName       string    `json:"application_name"`        // Name of the application, as provided by the application
	ApplicationVersion    string    `json:"application_version"`     // The version of the application, as provided by the application
	IsOfficialApplication bool      `json:"is_official_application"` // True, if the application is an official application or uses the api_id of an official application
	DeviceModel           string    `json:"device_model"`            // Model of the device the application has been run or is running on, as provided by the application
	Platform              string    `json:"platform"`                // Operating system the application has been run or is running on, as provided by the application
	SystemVersion         string    `json:"system_version"`          // Version of the operating system the application has been run or is running on, as provided by the application
	LogInDate             int32     `json:"log_in_date"`             // Point in time (Unix timestamp) when the user has logged in
	LastActiveDate        int32     `json:"last_active_date"`        // Point in time (Unix timestamp) when the session was last used
	IP                    string    `json:"ip"`                      // IP address from which the session was created, in human-readable format
	Country               string    `json:"country"`                 // A two-letter country code for the country from which the session was created, based on the IP address
	Region                string    `json:"region"`                  // Region code from which the session was created, based on the IP address
}

// MessageType return the string telegram-type of Session
func (session *Session) MessageType() string {
	return "session"
}

// Sessions Contains a list of sessions
type Sessions struct {
	tdCommon
	Sessions []Session `json:"sessions"` // List of sessions
}

// MessageType return the string telegram-type of Sessions
func (sessions *Sessions) MessageType() string {
	return "sessions"
}

// ConnectedWebsite Contains information about one website the current user is logged in with Telegram
type ConnectedWebsite struct {
	tdCommon
	ID             JSONInt64 `json:"id"`               // Website identifier
	DomainName     string    `json:"domain_name"`      // The domain name of the website
	BotUserID      int32     `json:"bot_user_id"`      // User identifier of a bot linked with the website
	Browser        string    `json:"browser"`          // The version of a browser used to log in
	Platform       string    `json:"platform"`         // Operating system the browser is running on
	LogInDate      int32     `json:"log_in_date"`      // Point in time (Unix timestamp) when the user was logged in
	LastActiveDate int32     `json:"last_active_date"` // Point in time (Unix timestamp) when obtained authorization was last used
	IP             string    `json:"ip"`               // IP address from which the user was logged in, in human-readable format
	Location       string    `json:"location"`         // Human-readable description of a country and a region, from which the user was logged in, based on the IP address
}

// MessageType return the string telegram-type of ConnectedWebsite
func (connectedWebsite *ConnectedWebsite) MessageType() string {
	return "connectedWebsite"
}

// ConnectedWebsites Contains a list of websites the current user is logged in with Telegram
type ConnectedWebsites struct {
	tdCommon
	Websites []ConnectedWebsite `json:"websites"` // List of connected websites
}

// MessageType return the string telegram-type of ConnectedWebsites
func (connectedWebsites *ConnectedWebsites) MessageType() string {
	return "connectedWebsites"
}

// ChatReportSpamState Contains information about the availability of the "Report spam" action for a chat
type ChatReportSpamState struct {
	tdCommon
	CanReportSpam bool `json:"can_report_spam"` // True, if a prompt with the "Report spam" action should be shown to the user
}

// MessageType return the string telegram-type of ChatReportSpamState
func (chatReportSpamState *ChatReportSpamState) MessageType() string {
	return "chatReportSpamState"
}

// ChatReportReasonSpam The chat contains spam messages
type ChatReportReasonSpam struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonSpam
func (chatReportReasonSpam *ChatReportReasonSpam) MessageType() string {
	return "chatReportReasonSpam"
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonSpam *ChatReportReasonSpam) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonSpamType
}

// ChatReportReasonViolence The chat promotes violence
type ChatReportReasonViolence struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonViolence
func (chatReportReasonViolence *ChatReportReasonViolence) MessageType() string {
	return "chatReportReasonViolence"
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonViolence *ChatReportReasonViolence) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonViolenceType
}

// ChatReportReasonPornography The chat contains pornographic messages
type ChatReportReasonPornography struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonPornography
func (chatReportReasonPornography *ChatReportReasonPornography) MessageType() string {
	return "chatReportReasonPornography"
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonPornography *ChatReportReasonPornography) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonPornographyType
}

// ChatReportReasonCustom A custom reason provided by the user
type ChatReportReasonCustom struct {
	tdCommon
	Text string `json:"text"` // Report text
}

// MessageType return the string telegram-type of ChatReportReasonCustom
func (chatReportReasonCustom *ChatReportReasonCustom) MessageType() string {
	return "chatReportReasonCustom"
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonCustom *ChatReportReasonCustom) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonCustomType
}

// PublicMessageLink Contains a public HTTPS link to a message in a public supergroup or channel
type PublicMessageLink struct {
	tdCommon
	Link string `json:"link"` // Message link
	HTML string `json:"html"` // HTML-code for embedding the message
}

// MessageType return the string telegram-type of PublicMessageLink
func (publicMessageLink *PublicMessageLink) MessageType() string {
	return "publicMessageLink"
}

// FileTypeNone The data is not a file
type FileTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeNone
func (fileTypeNone *FileTypeNone) MessageType() string {
	return "fileTypeNone"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeNone *FileTypeNone) GetFileTypeEnum() FileTypeEnum {
	return FileTypeNoneType
}

// FileTypeAnimation The file is an animation
type FileTypeAnimation struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeAnimation
func (fileTypeAnimation *FileTypeAnimation) MessageType() string {
	return "fileTypeAnimation"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeAnimation *FileTypeAnimation) GetFileTypeEnum() FileTypeEnum {
	return FileTypeAnimationType
}

// FileTypeAudio The file is an audio file
type FileTypeAudio struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeAudio
func (fileTypeAudio *FileTypeAudio) MessageType() string {
	return "fileTypeAudio"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeAudio *FileTypeAudio) GetFileTypeEnum() FileTypeEnum {
	return FileTypeAudioType
}

// FileTypeDocument The file is a document
type FileTypeDocument struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeDocument
func (fileTypeDocument *FileTypeDocument) MessageType() string {
	return "fileTypeDocument"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeDocument *FileTypeDocument) GetFileTypeEnum() FileTypeEnum {
	return FileTypeDocumentType
}

// FileTypePhoto The file is a photo
type FileTypePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypePhoto
func (fileTypePhoto *FileTypePhoto) MessageType() string {
	return "fileTypePhoto"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypePhoto *FileTypePhoto) GetFileTypeEnum() FileTypeEnum {
	return FileTypePhotoType
}

// FileTypeProfilePhoto The file is a profile photo
type FileTypeProfilePhoto struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeProfilePhoto
func (fileTypeProfilePhoto *FileTypeProfilePhoto) MessageType() string {
	return "fileTypeProfilePhoto"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeProfilePhoto *FileTypeProfilePhoto) GetFileTypeEnum() FileTypeEnum {
	return FileTypeProfilePhotoType
}

// FileTypeSecret The file was sent to a secret chat (the file type is not known to the server)
type FileTypeSecret struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecret
func (fileTypeSecret *FileTypeSecret) MessageType() string {
	return "fileTypeSecret"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecret *FileTypeSecret) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretType
}

// FileTypeSticker The file is a sticker
type FileTypeSticker struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSticker
func (fileTypeSticker *FileTypeSticker) MessageType() string {
	return "fileTypeSticker"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSticker *FileTypeSticker) GetFileTypeEnum() FileTypeEnum {
	return FileTypeStickerType
}

// FileTypeThumbnail The file is a thumbnail of another file
type FileTypeThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeThumbnail
func (fileTypeThumbnail *FileTypeThumbnail) MessageType() string {
	return "fileTypeThumbnail"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeThumbnail *FileTypeThumbnail) GetFileTypeEnum() FileTypeEnum {
	return FileTypeThumbnailType
}

// FileTypeUnknown The file type is not yet known
type FileTypeUnknown struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeUnknown
func (fileTypeUnknown *FileTypeUnknown) MessageType() string {
	return "fileTypeUnknown"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeUnknown *FileTypeUnknown) GetFileTypeEnum() FileTypeEnum {
	return FileTypeUnknownType
}

// FileTypeVideo The file is a video
type FileTypeVideo struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVideo
func (fileTypeVideo *FileTypeVideo) MessageType() string {
	return "fileTypeVideo"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVideo *FileTypeVideo) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVideoType
}

// FileTypeVideoNote The file is a video note
type FileTypeVideoNote struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVideoNote
func (fileTypeVideoNote *FileTypeVideoNote) MessageType() string {
	return "fileTypeVideoNote"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVideoNote *FileTypeVideoNote) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVideoNoteType
}

// FileTypeVoiceNote The file is a voice note
type FileTypeVoiceNote struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeVoiceNote
func (fileTypeVoiceNote *FileTypeVoiceNote) MessageType() string {
	return "fileTypeVoiceNote"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeVoiceNote *FileTypeVoiceNote) GetFileTypeEnum() FileTypeEnum {
	return FileTypeVoiceNoteType
}

// FileTypeWallpaper The file is a wallpaper
type FileTypeWallpaper struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeWallpaper
func (fileTypeWallpaper *FileTypeWallpaper) MessageType() string {
	return "fileTypeWallpaper"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeWallpaper *FileTypeWallpaper) GetFileTypeEnum() FileTypeEnum {
	return FileTypeWallpaperType
}

// FileTypeSecretThumbnail The file is a thumbnail of a file from a secret chat
type FileTypeSecretThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecretThumbnail
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) MessageType() string {
	return "fileTypeSecretThumbnail"
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretThumbnailType
}

// StorageStatisticsByFileType Contains the storage usage statistics for a specific file type
type StorageStatisticsByFileType struct {
	tdCommon
	FileType FileType `json:"file_type"` // File type
	Size     int64    `json:"size"`      // Total size of the files
	Count    int32    `json:"count"`     // Total number of files
}

// MessageType return the string telegram-type of StorageStatisticsByFileType
func (storageStatisticsByFileType *StorageStatisticsByFileType) MessageType() string {
	return "storageStatisticsByFileType"
}

// UnmarshalJSON unmarshal to json
func (storageStatisticsByFileType *StorageStatisticsByFileType) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Size  int64 `json:"size"`  // Total size of the files
		Count int32 `json:"count"` // Total number of files
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	storageStatisticsByFileType.tdCommon = tempObj.tdCommon
	storageStatisticsByFileType.Size = tempObj.Size
	storageStatisticsByFileType.Count = tempObj.Count

	fieldFileType, _ := unmarshalFileType(objMap["file_type"])
	storageStatisticsByFileType.FileType = fieldFileType

	return nil
}

// StorageStatisticsByChat Contains the storage usage statistics for a specific chat
type StorageStatisticsByChat struct {
	tdCommon
	ChatID     int64                         `json:"chat_id"`      // Chat identifier; 0 if none
	Size       int64                         `json:"size"`         // Total size of the files in the chat
	Count      int32                         `json:"count"`        // Total number of files in the chat
	ByFileType []StorageStatisticsByFileType `json:"by_file_type"` // Statistics split by file types
}

// MessageType return the string telegram-type of StorageStatisticsByChat
func (storageStatisticsByChat *StorageStatisticsByChat) MessageType() string {
	return "storageStatisticsByChat"
}

// StorageStatistics Contains the exact storage usage statistics split by chats and file type
type StorageStatistics struct {
	tdCommon
	Size   int64                     `json:"size"`    // Total size of files
	Count  int32                     `json:"count"`   // Total number of files
	ByChat []StorageStatisticsByChat `json:"by_chat"` // Statistics split by chats
}

// MessageType return the string telegram-type of StorageStatistics
func (storageStatistics *StorageStatistics) MessageType() string {
	return "storageStatistics"
}

// StorageStatisticsFast Contains approximate storage usage statistics, excluding files of unknown file type
type StorageStatisticsFast struct {
	tdCommon
	FilesSize    int64 `json:"files_size"`    // Approximate total size of files
	FileCount    int32 `json:"file_count"`    // Approximate number of files
	DatabaseSize int64 `json:"database_size"` // Size of the database
}

// MessageType return the string telegram-type of StorageStatisticsFast
func (storageStatisticsFast *StorageStatisticsFast) MessageType() string {
	return "storageStatisticsFast"
}

// NetworkTypeNone The network is not available
type NetworkTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeNone
func (networkTypeNone *NetworkTypeNone) MessageType() string {
	return "networkTypeNone"
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeNone *NetworkTypeNone) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeNoneType
}

// NetworkTypeMobile A mobile network
type NetworkTypeMobile struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeMobile
func (networkTypeMobile *NetworkTypeMobile) MessageType() string {
	return "networkTypeMobile"
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeMobile *NetworkTypeMobile) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeMobileType
}

// NetworkTypeMobileRoaming A mobile roaming network
type NetworkTypeMobileRoaming struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeMobileRoaming
func (networkTypeMobileRoaming *NetworkTypeMobileRoaming) MessageType() string {
	return "networkTypeMobileRoaming"
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeMobileRoaming *NetworkTypeMobileRoaming) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeMobileRoamingType
}

// NetworkTypeWiFi A Wi-Fi network
type NetworkTypeWiFi struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeWiFi
func (networkTypeWiFi *NetworkTypeWiFi) MessageType() string {
	return "networkTypeWiFi"
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeWiFi *NetworkTypeWiFi) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeWiFiType
}

// NetworkTypeOther A different network type (e.g., Ethernet network)
type NetworkTypeOther struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeOther
func (networkTypeOther *NetworkTypeOther) MessageType() string {
	return "networkTypeOther"
}

// GetNetworkTypeEnum return the enum type of this object
func (networkTypeOther *NetworkTypeOther) GetNetworkTypeEnum() NetworkTypeEnum {
	return NetworkTypeOtherType
}

// NetworkStatisticsEntryFile Contains information about the total amount of data that was used to send and receive files
type NetworkStatisticsEntryFile struct {
	tdCommon
	FileType      FileType    `json:"file_type"`      // Type of the file the data is part of
	NetworkType   NetworkType `json:"network_type"`   // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
	SentBytes     int64       `json:"sent_bytes"`     // Total number of bytes sent
	ReceivedBytes int64       `json:"received_bytes"` // Total number of bytes received
}

// MessageType return the string telegram-type of NetworkStatisticsEntryFile
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) MessageType() string {
	return "networkStatisticsEntryFile"
}

// UnmarshalJSON unmarshal to json
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		SentBytes     int64 `json:"sent_bytes"`     // Total number of bytes sent
		ReceivedBytes int64 `json:"received_bytes"` // Total number of bytes received
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	networkStatisticsEntryFile.tdCommon = tempObj.tdCommon
	networkStatisticsEntryFile.SentBytes = tempObj.SentBytes
	networkStatisticsEntryFile.ReceivedBytes = tempObj.ReceivedBytes

	fieldFileType, _ := unmarshalFileType(objMap["file_type"])
	networkStatisticsEntryFile.FileType = fieldFileType

	fieldNetworkType, _ := unmarshalNetworkType(objMap["network_type"])
	networkStatisticsEntryFile.NetworkType = fieldNetworkType

	return nil
}

// GetNetworkStatisticsEntryEnum return the enum type of this object
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum {
	return NetworkStatisticsEntryFileType
}

// NetworkStatisticsEntryCall Contains information about the total amount of data that was used for calls
type NetworkStatisticsEntryCall struct {
	tdCommon
	NetworkType   NetworkType `json:"network_type"`   // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
	SentBytes     int64       `json:"sent_bytes"`     // Total number of bytes sent
	ReceivedBytes int64       `json:"received_bytes"` // Total number of bytes received
	Duration      float64     `json:"duration"`       // Total call duration, in seconds
}

// MessageType return the string telegram-type of NetworkStatisticsEntryCall
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) MessageType() string {
	return "networkStatisticsEntryCall"
}

// UnmarshalJSON unmarshal to json
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		SentBytes     int64   `json:"sent_bytes"`     // Total number of bytes sent
		ReceivedBytes int64   `json:"received_bytes"` // Total number of bytes received
		Duration      float64 `json:"duration"`       // Total call duration, in seconds
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	networkStatisticsEntryCall.tdCommon = tempObj.tdCommon
	networkStatisticsEntryCall.SentBytes = tempObj.SentBytes
	networkStatisticsEntryCall.ReceivedBytes = tempObj.ReceivedBytes
	networkStatisticsEntryCall.Duration = tempObj.Duration

	fieldNetworkType, _ := unmarshalNetworkType(objMap["network_type"])
	networkStatisticsEntryCall.NetworkType = fieldNetworkType

	return nil
}

// GetNetworkStatisticsEntryEnum return the enum type of this object
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) GetNetworkStatisticsEntryEnum() NetworkStatisticsEntryEnum {
	return NetworkStatisticsEntryCallType
}

// NetworkStatistics A full list of available network statistic entries
type NetworkStatistics struct {
	tdCommon
	SinceDate int32                    `json:"since_date"` // Point in time (Unix timestamp) when the app began collecting statistics
	Entries   []NetworkStatisticsEntry `json:"entries"`    // Network statistics entries
}

// MessageType return the string telegram-type of NetworkStatistics
func (networkStatistics *NetworkStatistics) MessageType() string {
	return "networkStatistics"
}

// ConnectionStateWaitingForNetwork Currently waiting for the network to become available. Use SetNetworkType to change the available network type
type ConnectionStateWaitingForNetwork struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateWaitingForNetwork
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) MessageType() string {
	return "connectionStateWaitingForNetwork"
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateWaitingForNetworkType
}

// ConnectionStateConnectingToProxy Currently establishing a connection with a proxy server
type ConnectionStateConnectingToProxy struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateConnectingToProxy
func (connectionStateConnectingToProxy *ConnectionStateConnectingToProxy) MessageType() string {
	return "connectionStateConnectingToProxy"
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateConnectingToProxy *ConnectionStateConnectingToProxy) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateConnectingToProxyType
}

// ConnectionStateConnecting Currently establishing a connection to the Telegram servers
type ConnectionStateConnecting struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateConnecting
func (connectionStateConnecting *ConnectionStateConnecting) MessageType() string {
	return "connectionStateConnecting"
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateConnecting *ConnectionStateConnecting) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateConnectingType
}

// ConnectionStateUpdating Downloading data received while the client was offline
type ConnectionStateUpdating struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateUpdating
func (connectionStateUpdating *ConnectionStateUpdating) MessageType() string {
	return "connectionStateUpdating"
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateUpdating *ConnectionStateUpdating) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateUpdatingType
}

// ConnectionStateReady There is a working connection to the Telegram servers
type ConnectionStateReady struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateReady
func (connectionStateReady *ConnectionStateReady) MessageType() string {
	return "connectionStateReady"
}

// GetConnectionStateEnum return the enum type of this object
func (connectionStateReady *ConnectionStateReady) GetConnectionStateEnum() ConnectionStateEnum {
	return ConnectionStateReadyType
}

// TopChatCategoryUsers A category containing frequently used private chats with non-bot users
type TopChatCategoryUsers struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryUsers
func (topChatCategoryUsers *TopChatCategoryUsers) MessageType() string {
	return "topChatCategoryUsers"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryUsers *TopChatCategoryUsers) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryUsersType
}

// TopChatCategoryBots A category containing frequently used private chats with bot users
type TopChatCategoryBots struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryBots
func (topChatCategoryBots *TopChatCategoryBots) MessageType() string {
	return "topChatCategoryBots"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryBots *TopChatCategoryBots) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryBotsType
}

// TopChatCategoryGroups A category containing frequently used basic groups and supergroups
type TopChatCategoryGroups struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryGroups
func (topChatCategoryGroups *TopChatCategoryGroups) MessageType() string {
	return "topChatCategoryGroups"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryGroups *TopChatCategoryGroups) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryGroupsType
}

// TopChatCategoryChannels A category containing frequently used channels
type TopChatCategoryChannels struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryChannels
func (topChatCategoryChannels *TopChatCategoryChannels) MessageType() string {
	return "topChatCategoryChannels"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryChannels *TopChatCategoryChannels) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryChannelsType
}

// TopChatCategoryInlineBots A category containing frequently used chats with inline bots sorted by their usage in inline mode
type TopChatCategoryInlineBots struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryInlineBots
func (topChatCategoryInlineBots *TopChatCategoryInlineBots) MessageType() string {
	return "topChatCategoryInlineBots"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryInlineBots *TopChatCategoryInlineBots) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryInlineBotsType
}

// TopChatCategoryCalls A category containing frequently used chats used for calls
type TopChatCategoryCalls struct {
	tdCommon
}

// MessageType return the string telegram-type of TopChatCategoryCalls
func (topChatCategoryCalls *TopChatCategoryCalls) MessageType() string {
	return "topChatCategoryCalls"
}

// GetTopChatCategoryEnum return the enum type of this object
func (topChatCategoryCalls *TopChatCategoryCalls) GetTopChatCategoryEnum() TopChatCategoryEnum {
	return TopChatCategoryCallsType
}

// TMeURLTypeUser A URL linking to a user
type TMeURLTypeUser struct {
	tdCommon
	UserID int32 `json:"user_id"` // Identifier of the user
}

// MessageType return the string telegram-type of TMeURLTypeUser
func (tMeURLTypeUser *TMeURLTypeUser) MessageType() string {
	return "tMeUrlTypeUser"
}

// GetTMeURLTypeEnum return the enum type of this object
func (tMeURLTypeUser *TMeURLTypeUser) GetTMeURLTypeEnum() TMeURLTypeEnum {
	return TMeURLTypeUserType
}

// TMeURLTypeSupergroup A URL linking to a public supergroup or channel
type TMeURLTypeSupergroup struct {
	tdCommon
	SupergroupID int64 `json:"supergroup_id"` // Identifier of the supergroup or channel
}

// MessageType return the string telegram-type of TMeURLTypeSupergroup
func (tMeURLTypeSupergroup *TMeURLTypeSupergroup) MessageType() string {
	return "tMeUrlTypeSupergroup"
}

// GetTMeURLTypeEnum return the enum type of this object
func (tMeURLTypeSupergroup *TMeURLTypeSupergroup) GetTMeURLTypeEnum() TMeURLTypeEnum {
	return TMeURLTypeSupergroupType
}

// TMeURLTypeChatInvite A chat invite link
type TMeURLTypeChatInvite struct {
	tdCommon
	Info ChatInviteLinkInfo `json:"info"` // Chat invite link info
}

// MessageType return the string telegram-type of TMeURLTypeChatInvite
func (tMeURLTypeChatInvite *TMeURLTypeChatInvite) MessageType() string {
	return "tMeUrlTypeChatInvite"
}

// GetTMeURLTypeEnum return the enum type of this object
func (tMeURLTypeChatInvite *TMeURLTypeChatInvite) GetTMeURLTypeEnum() TMeURLTypeEnum {
	return TMeURLTypeChatInviteType
}

// TMeURLTypeStickerSet A URL linking to a sticker set
type TMeURLTypeStickerSet struct {
	tdCommon
	StickerSetID JSONInt64 `json:"sticker_set_id"` // Identifier of the sticker set
}

// MessageType return the string telegram-type of TMeURLTypeStickerSet
func (tMeURLTypeStickerSet *TMeURLTypeStickerSet) MessageType() string {
	return "tMeUrlTypeStickerSet"
}

// GetTMeURLTypeEnum return the enum type of this object
func (tMeURLTypeStickerSet *TMeURLTypeStickerSet) GetTMeURLTypeEnum() TMeURLTypeEnum {
	return TMeURLTypeStickerSetType
}

// TMeURL Represents a URL linking to an internal Telegram entity
type TMeURL struct {
	tdCommon
	URL  string     `json:"url"`  // URL
	Type TMeURLType `json:"type"` // Type of the URL
}

// MessageType return the string telegram-type of TMeURL
func (tMeURL *TMeURL) MessageType() string {
	return "tMeUrl"
}

// UnmarshalJSON unmarshal to json
func (tMeURL *TMeURL) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		URL string `json:"url"` // URL

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	tMeURL.tdCommon = tempObj.tdCommon
	tMeURL.URL = tempObj.URL

	fieldType, _ := unmarshalTMeURLType(objMap["type"])
	tMeURL.Type = fieldType

	return nil
}

// TMeURLs Contains a list of t.me URLs
type TMeURLs struct {
	tdCommon
	URLs []TMeURL `json:"urls"` // List of URLs
}

// MessageType return the string telegram-type of TMeURLs
func (tMeURLs *TMeURLs) MessageType() string {
	return "tMeUrls"
}

// Count Contains a counter
type Count struct {
	tdCommon
	Count int32 `json:"count"` // Count
}

// MessageType return the string telegram-type of Count
func (count *Count) MessageType() string {
	return "count"
}

// Text Contains some text
type Text struct {
	tdCommon
	Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of Text
func (text *Text) MessageType() string {
	return "text"
}

// TextParseModeMarkdown The text should be parsed in markdown-style
type TextParseModeMarkdown struct {
	tdCommon
}

// MessageType return the string telegram-type of TextParseModeMarkdown
func (textParseModeMarkdown *TextParseModeMarkdown) MessageType() string {
	return "textParseModeMarkdown"
}

// GetTextParseModeEnum return the enum type of this object
func (textParseModeMarkdown *TextParseModeMarkdown) GetTextParseModeEnum() TextParseModeEnum {
	return TextParseModeMarkdownType
}

// TextParseModeHTML The text should be parsed in HTML-style
type TextParseModeHTML struct {
	tdCommon
}

// MessageType return the string telegram-type of TextParseModeHTML
func (textParseModeHTML *TextParseModeHTML) MessageType() string {
	return "textParseModeHTML"
}

// GetTextParseModeEnum return the enum type of this object
func (textParseModeHTML *TextParseModeHTML) GetTextParseModeEnum() TextParseModeEnum {
	return TextParseModeHTMLType
}

// ProxyEmpty An empty proxy server
type ProxyEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of ProxyEmpty
func (proxyEmpty *ProxyEmpty) MessageType() string {
	return "proxyEmpty"
}

// GetProxyEnum return the enum type of this object
func (proxyEmpty *ProxyEmpty) GetProxyEnum() ProxyEnum {
	return ProxyEmptyType
}

// ProxySocks5 A SOCKS5 proxy server
type ProxySocks5 struct {
	tdCommon
	Server   string `json:"server"`   // Proxy server IP address
	Port     int32  `json:"port"`     // Proxy server port
	Username string `json:"username"` // Username for logging in
	Password string `json:"password"` // Password for logging in
}

// MessageType return the string telegram-type of ProxySocks5
func (proxySocks5 *ProxySocks5) MessageType() string {
	return "proxySocks5"
}

// GetProxyEnum return the enum type of this object
func (proxySocks5 *ProxySocks5) GetProxyEnum() ProxyEnum {
	return ProxySocks5Type
}

// InputSticker Describes a sticker that should be added to a sticker set
type InputSticker struct {
	tdCommon
	PngSticker   InputFile    `json:"png_sticker"`   // PNG image with the sticker; must be up to 512 kB in size and fit in a 512x512 square
	Emojis       string       `json:"emojis"`        // Emoji corresponding to the sticker
	MaskPosition MaskPosition `json:"mask_position"` // For masks, position where the mask should be placed; may be null
}

// MessageType return the string telegram-type of InputSticker
func (inputSticker *InputSticker) MessageType() string {
	return "inputSticker"
}

// UnmarshalJSON unmarshal to json
func (inputSticker *InputSticker) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Emojis       string       `json:"emojis"`        // Emoji corresponding to the sticker
		MaskPosition MaskPosition `json:"mask_position"` // For masks, position where the mask should be placed; may be null
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputSticker.tdCommon = tempObj.tdCommon
	inputSticker.Emojis = tempObj.Emojis
	inputSticker.MaskPosition = tempObj.MaskPosition

	fieldPngSticker, _ := unmarshalInputFile(objMap["png_sticker"])
	inputSticker.PngSticker = fieldPngSticker

	return nil
}

// UpdateAuthorizationState The user authorization state has changed
type UpdateAuthorizationState struct {
	tdCommon
	AuthorizationState AuthorizationState `json:"authorization_state"` // New authorization state
}

// MessageType return the string telegram-type of UpdateAuthorizationState
func (updateAuthorizationState *UpdateAuthorizationState) MessageType() string {
	return "updateAuthorizationState"
}

// UnmarshalJSON unmarshal to json
func (updateAuthorizationState *UpdateAuthorizationState) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateAuthorizationState.tdCommon = tempObj.tdCommon

	fieldAuthorizationState, _ := unmarshalAuthorizationState(objMap["authorization_state"])
	updateAuthorizationState.AuthorizationState = fieldAuthorizationState

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateAuthorizationState *UpdateAuthorizationState) GetUpdateEnum() UpdateEnum {
	return UpdateAuthorizationStateType
}

// UpdateNewMessage A new message was received; can also be an outgoing message
type UpdateNewMessage struct {
	tdCommon
	Message             Message `json:"message"`              // The new message
	DisableNotification bool    `json:"disable_notification"` // True, if this message must not generate a notification
	ContainsMention     bool    `json:"contains_mention"`     // True, if the message contains a mention of the current user
}

// MessageType return the string telegram-type of UpdateNewMessage
func (updateNewMessage *UpdateNewMessage) MessageType() string {
	return "updateNewMessage"
}

// GetUpdateEnum return the enum type of this object
func (updateNewMessage *UpdateNewMessage) GetUpdateEnum() UpdateEnum {
	return UpdateNewMessageType
}

// UpdateMessageSendAcknowledged A request to send a message has reached the Telegram server. This doesn't mean that the message will be sent successfully or even that the send message request will be processed. This update will be sent only if the option "use_quick_ack" is set to true. This update may be sent multiple times for the same message
type UpdateMessageSendAcknowledged struct {
	tdCommon
	ChatID    int64 `json:"chat_id"`    // The chat identifier of the sent message
	MessageID int64 `json:"message_id"` // A temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendAcknowledged
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) MessageType() string {
	return "updateMessageSendAcknowledged"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendAcknowledgedType
}

// UpdateMessageSendSucceeded A message has been successfully sent
type UpdateMessageSendSucceeded struct {
	tdCommon
	Message      Message `json:"message"`        // Information about the sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change
	OldMessageID int64   `json:"old_message_id"` // The previous temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendSucceeded
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) MessageType() string {
	return "updateMessageSendSucceeded"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendSucceededType
}

// UpdateMessageSendFailed A message failed to send. Be aware that some messages being sent can be irrecoverably deleted, in which case updateDeleteMessages will be received instead of this update
type UpdateMessageSendFailed struct {
	tdCommon
	Message      Message `json:"message"`        // Contains information about the message that failed to send
	OldMessageID int64   `json:"old_message_id"` // The previous temporary message identifier
	ErrorCode    int32   `json:"error_code"`     // An error code
	ErrorMessage string  `json:"error_message"`  // Error message
}

// MessageType return the string telegram-type of UpdateMessageSendFailed
func (updateMessageSendFailed *UpdateMessageSendFailed) MessageType() string {
	return "updateMessageSendFailed"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendFailed *UpdateMessageSendFailed) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendFailedType
}

// UpdateMessageContent The message content has changed
type UpdateMessageContent struct {
	tdCommon
	ChatID     int64          `json:"chat_id"`     // Chat identifier
	MessageID  int64          `json:"message_id"`  // Message identifier
	NewContent MessageContent `json:"new_content"` // New message content
}

// MessageType return the string telegram-type of UpdateMessageContent
func (updateMessageContent *UpdateMessageContent) MessageType() string {
	return "updateMessageContent"
}

// UnmarshalJSON unmarshal to json
func (updateMessageContent *UpdateMessageContent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatID    int64 `json:"chat_id"`    // Chat identifier
		MessageID int64 `json:"message_id"` // Message identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateMessageContent.tdCommon = tempObj.tdCommon
	updateMessageContent.ChatID = tempObj.ChatID
	updateMessageContent.MessageID = tempObj.MessageID

	fieldNewContent, _ := unmarshalMessageContent(objMap["new_content"])
	updateMessageContent.NewContent = fieldNewContent

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateMessageContent *UpdateMessageContent) GetUpdateEnum() UpdateEnum {
	return UpdateMessageContentType
}

// UpdateMessageEdited A message was edited. Changes in the message content will come in a separate updateMessageContent
type UpdateMessageEdited struct {
	tdCommon
	ChatID      int64       `json:"chat_id"`      // Chat identifier
	MessageID   int64       `json:"message_id"`   // Message identifier
	EditDate    int32       `json:"edit_date"`    // Point in time (Unix timestamp) when the message was edited
	ReplyMarkup ReplyMarkup `json:"reply_markup"` // New message reply markup; may be null
}

// MessageType return the string telegram-type of UpdateMessageEdited
func (updateMessageEdited *UpdateMessageEdited) MessageType() string {
	return "updateMessageEdited"
}

// UnmarshalJSON unmarshal to json
func (updateMessageEdited *UpdateMessageEdited) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatID    int64 `json:"chat_id"`    // Chat identifier
		MessageID int64 `json:"message_id"` // Message identifier
		EditDate  int32 `json:"edit_date"`  // Point in time (Unix timestamp) when the message was edited

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateMessageEdited.tdCommon = tempObj.tdCommon
	updateMessageEdited.ChatID = tempObj.ChatID
	updateMessageEdited.MessageID = tempObj.MessageID
	updateMessageEdited.EditDate = tempObj.EditDate

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	updateMessageEdited.ReplyMarkup = fieldReplyMarkup

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateMessageEdited *UpdateMessageEdited) GetUpdateEnum() UpdateEnum {
	return UpdateMessageEditedType
}

// UpdateMessageViews The view count of the message has changed
type UpdateMessageViews struct {
	tdCommon
	ChatID    int64 `json:"chat_id"`    // Chat identifier
	MessageID int64 `json:"message_id"` // Message identifier
	Views     int32 `json:"views"`      // New value of the view count
}

// MessageType return the string telegram-type of UpdateMessageViews
func (updateMessageViews *UpdateMessageViews) MessageType() string {
	return "updateMessageViews"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageViews *UpdateMessageViews) GetUpdateEnum() UpdateEnum {
	return UpdateMessageViewsType
}

// UpdateMessageContentOpened The message content was opened. Updates voice note messages to "listened", video note messages to "viewed" and starts the TTL timer for self-destructing messages
type UpdateMessageContentOpened struct {
	tdCommon
	ChatID    int64 `json:"chat_id"`    // Chat identifier
	MessageID int64 `json:"message_id"` // Message identifier
}

// MessageType return the string telegram-type of UpdateMessageContentOpened
func (updateMessageContentOpened *UpdateMessageContentOpened) MessageType() string {
	return "updateMessageContentOpened"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageContentOpened *UpdateMessageContentOpened) GetUpdateEnum() UpdateEnum {
	return UpdateMessageContentOpenedType
}

// UpdateMessageMentionRead A message with an unread mention was read
type UpdateMessageMentionRead struct {
	tdCommon
	ChatID             int64 `json:"chat_id"`              // Chat identifier
	MessageID          int64 `json:"message_id"`           // Message identifier
	UnreadMentionCount int32 `json:"unread_mention_count"` // The new number of unread mention messages left in the chat
}

// MessageType return the string telegram-type of UpdateMessageMentionRead
func (updateMessageMentionRead *UpdateMessageMentionRead) MessageType() string {
	return "updateMessageMentionRead"
}

// GetUpdateEnum return the enum type of this object
func (updateMessageMentionRead *UpdateMessageMentionRead) GetUpdateEnum() UpdateEnum {
	return UpdateMessageMentionReadType
}

// UpdateNewChat A new chat has been loaded/created. This update is guaranteed to come before the chat identifier is returned to the client. The chat field changes will be reported through separate updates
type UpdateNewChat struct {
	tdCommon
	Chat Chat `json:"chat"` // The chat
}

// MessageType return the string telegram-type of UpdateNewChat
func (updateNewChat *UpdateNewChat) MessageType() string {
	return "updateNewChat"
}

// GetUpdateEnum return the enum type of this object
func (updateNewChat *UpdateNewChat) GetUpdateEnum() UpdateEnum {
	return UpdateNewChatType
}

// UpdateChatTitle The title of a chat was changed
type UpdateChatTitle struct {
	tdCommon
	ChatID int64  `json:"chat_id"` // Chat identifier
	Title  string `json:"title"`   // The new chat title
}

// MessageType return the string telegram-type of UpdateChatTitle
func (updateChatTitle *UpdateChatTitle) MessageType() string {
	return "updateChatTitle"
}

// GetUpdateEnum return the enum type of this object
func (updateChatTitle *UpdateChatTitle) GetUpdateEnum() UpdateEnum {
	return UpdateChatTitleType
}

// UpdateChatPhoto A chat photo was changed
type UpdateChatPhoto struct {
	tdCommon
	ChatID int64     `json:"chat_id"` // Chat identifier
	Photo  ChatPhoto `json:"photo"`   // The new chat photo; may be null
}

// MessageType return the string telegram-type of UpdateChatPhoto
func (updateChatPhoto *UpdateChatPhoto) MessageType() string {
	return "updateChatPhoto"
}

// GetUpdateEnum return the enum type of this object
func (updateChatPhoto *UpdateChatPhoto) GetUpdateEnum() UpdateEnum {
	return UpdateChatPhotoType
}

// UpdateChatLastMessage The last message of a chat was changed. If last_message is null then the last message in the chat became unknown. Some new unknown messages might be added to the chat in this case
type UpdateChatLastMessage struct {
	tdCommon
	ChatID      int64     `json:"chat_id"`      // Chat identifier
	LastMessage Message   `json:"last_message"` // The new last message in the chat; may be null
	Order       JSONInt64 `json:"order"`        // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatLastMessage
func (updateChatLastMessage *UpdateChatLastMessage) MessageType() string {
	return "updateChatLastMessage"
}

// GetUpdateEnum return the enum type of this object
func (updateChatLastMessage *UpdateChatLastMessage) GetUpdateEnum() UpdateEnum {
	return UpdateChatLastMessageType
}

// UpdateChatOrder The order of the chat in the chats list has changed. Instead of this update updateChatLastMessage, updateChatIsPinned or updateChatDraftMessage might be sent
type UpdateChatOrder struct {
	tdCommon
	ChatID int64     `json:"chat_id"` // Chat identifier
	Order  JSONInt64 `json:"order"`   // New value of the order
}

// MessageType return the string telegram-type of UpdateChatOrder
func (updateChatOrder *UpdateChatOrder) MessageType() string {
	return "updateChatOrder"
}

// GetUpdateEnum return the enum type of this object
func (updateChatOrder *UpdateChatOrder) GetUpdateEnum() UpdateEnum {
	return UpdateChatOrderType
}

// UpdateChatIsPinned A chat was pinned or unpinned
type UpdateChatIsPinned struct {
	tdCommon
	ChatID   int64     `json:"chat_id"`   // Chat identifier
	IsPinned bool      `json:"is_pinned"` // New value of is_pinned
	Order    JSONInt64 `json:"order"`     // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatIsPinned
func (updateChatIsPinned *UpdateChatIsPinned) MessageType() string {
	return "updateChatIsPinned"
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsPinned *UpdateChatIsPinned) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsPinnedType
}

// UpdateChatReadInbox Incoming messages were read or number of unread messages has been changed
type UpdateChatReadInbox struct {
	tdCommon
	ChatID                 int64 `json:"chat_id"`                    // Chat identifier
	LastReadInboxMessageID int64 `json:"last_read_inbox_message_id"` // Identifier of the last read incoming message
	UnreadCount            int32 `json:"unread_count"`               // The number of unread messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatReadInbox
func (updateChatReadInbox *UpdateChatReadInbox) MessageType() string {
	return "updateChatReadInbox"
}

// GetUpdateEnum return the enum type of this object
func (updateChatReadInbox *UpdateChatReadInbox) GetUpdateEnum() UpdateEnum {
	return UpdateChatReadInboxType
}

// UpdateChatReadOutbox Outgoing messages were read
type UpdateChatReadOutbox struct {
	tdCommon
	ChatID                  int64 `json:"chat_id"`                     // Chat identifier
	LastReadOutboxMessageID int64 `json:"last_read_outbox_message_id"` // Identifier of last read outgoing message
}

// MessageType return the string telegram-type of UpdateChatReadOutbox
func (updateChatReadOutbox *UpdateChatReadOutbox) MessageType() string {
	return "updateChatReadOutbox"
}

// GetUpdateEnum return the enum type of this object
func (updateChatReadOutbox *UpdateChatReadOutbox) GetUpdateEnum() UpdateEnum {
	return UpdateChatReadOutboxType
}

// UpdateChatUnreadMentionCount The chat unread_mention_count has changed
type UpdateChatUnreadMentionCount struct {
	tdCommon
	ChatID             int64 `json:"chat_id"`              // Chat identifier
	UnreadMentionCount int32 `json:"unread_mention_count"` // The number of unread mention messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatUnreadMentionCount
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) MessageType() string {
	return "updateChatUnreadMentionCount"
}

// GetUpdateEnum return the enum type of this object
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) GetUpdateEnum() UpdateEnum {
	return UpdateChatUnreadMentionCountType
}

// UpdateNotificationSettings Notification settings for some chats were updated
type UpdateNotificationSettings struct {
	tdCommon
	Scope                NotificationSettingsScope `json:"scope"`                 // Types of chats for which notification settings were updated
	NotificationSettings NotificationSettings      `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateNotificationSettings
func (updateNotificationSettings *UpdateNotificationSettings) MessageType() string {
	return "updateNotificationSettings"
}

// UnmarshalJSON unmarshal to json
func (updateNotificationSettings *UpdateNotificationSettings) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NotificationSettings NotificationSettings `json:"notification_settings"` // The new notification settings
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNotificationSettings.tdCommon = tempObj.tdCommon
	updateNotificationSettings.NotificationSettings = tempObj.NotificationSettings

	fieldScope, _ := unmarshalNotificationSettingsScope(objMap["scope"])
	updateNotificationSettings.Scope = fieldScope

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNotificationSettings *UpdateNotificationSettings) GetUpdateEnum() UpdateEnum {
	return UpdateNotificationSettingsType
}

// UpdateChatReplyMarkup The default chat reply markup was changed. Can occur because new messages with reply markup were received or because an old reply markup was hidden by the user
type UpdateChatReplyMarkup struct {
	tdCommon
	ChatID               int64 `json:"chat_id"`                 // Chat identifier
	ReplyMarkupMessageID int64 `json:"reply_markup_message_id"` // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
}

// MessageType return the string telegram-type of UpdateChatReplyMarkup
func (updateChatReplyMarkup *UpdateChatReplyMarkup) MessageType() string {
	return "updateChatReplyMarkup"
}

// GetUpdateEnum return the enum type of this object
func (updateChatReplyMarkup *UpdateChatReplyMarkup) GetUpdateEnum() UpdateEnum {
	return UpdateChatReplyMarkupType
}

// UpdateChatDraftMessage A draft has changed. Be aware that the update may come in the currently opened chat but with old content of the draft. If the user has changed the content of the draft, this update shouldn't be applied
type UpdateChatDraftMessage struct {
	tdCommon
	ChatID       int64        `json:"chat_id"`       // Chat identifier
	DraftMessage DraftMessage `json:"draft_message"` // The new draft message; may be null
	Order        JSONInt64    `json:"order"`         // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatDraftMessage
func (updateChatDraftMessage *UpdateChatDraftMessage) MessageType() string {
	return "updateChatDraftMessage"
}

// GetUpdateEnum return the enum type of this object
func (updateChatDraftMessage *UpdateChatDraftMessage) GetUpdateEnum() UpdateEnum {
	return UpdateChatDraftMessageType
}

// UpdateDeleteMessages Some messages were deleted
type UpdateDeleteMessages struct {
	tdCommon
	ChatID      int64   `json:"chat_id"`      // Chat identifier
	MessageIDs  []int64 `json:"message_ids"`  // Identifiers of the deleted messages
	IsPermanent bool    `json:"is_permanent"` // True, if the messages are permanently deleted by a user (as opposed to just becoming unaccessible)
	FromCache   bool    `json:"from_cache"`   // True, if the messages are deleted only from the cache and can possibly be retrieved again in the future
}

// MessageType return the string telegram-type of UpdateDeleteMessages
func (updateDeleteMessages *UpdateDeleteMessages) MessageType() string {
	return "updateDeleteMessages"
}

// GetUpdateEnum return the enum type of this object
func (updateDeleteMessages *UpdateDeleteMessages) GetUpdateEnum() UpdateEnum {
	return UpdateDeleteMessagesType
}

// UpdateUserChatAction User activity in the chat has changed
type UpdateUserChatAction struct {
	tdCommon
	ChatID int64      `json:"chat_id"` // Chat identifier
	UserID int32      `json:"user_id"` // Identifier of a user performing an action
	Action ChatAction `json:"action"`  // The action description
}

// MessageType return the string telegram-type of UpdateUserChatAction
func (updateUserChatAction *UpdateUserChatAction) MessageType() string {
	return "updateUserChatAction"
}

// UnmarshalJSON unmarshal to json
func (updateUserChatAction *UpdateUserChatAction) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ChatID int64 `json:"chat_id"` // Chat identifier
		UserID int32 `json:"user_id"` // Identifier of a user performing an action

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserChatAction.tdCommon = tempObj.tdCommon
	updateUserChatAction.ChatID = tempObj.ChatID
	updateUserChatAction.UserID = tempObj.UserID

	fieldAction, _ := unmarshalChatAction(objMap["action"])
	updateUserChatAction.Action = fieldAction

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserChatAction *UpdateUserChatAction) GetUpdateEnum() UpdateEnum {
	return UpdateUserChatActionType
}

// UpdateUserStatus The user went online or offline
type UpdateUserStatus struct {
	tdCommon
	UserID int32      `json:"user_id"` // User identifier
	Status UserStatus `json:"status"`  // New status of the user
}

// MessageType return the string telegram-type of UpdateUserStatus
func (updateUserStatus *UpdateUserStatus) MessageType() string {
	return "updateUserStatus"
}

// UnmarshalJSON unmarshal to json
func (updateUserStatus *UpdateUserStatus) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		UserID int32 `json:"user_id"` // User identifier

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserStatus.tdCommon = tempObj.tdCommon
	updateUserStatus.UserID = tempObj.UserID

	fieldStatus, _ := unmarshalUserStatus(objMap["status"])
	updateUserStatus.Status = fieldStatus

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserStatus *UpdateUserStatus) GetUpdateEnum() UpdateEnum {
	return UpdateUserStatusType
}

// UpdateUser Some data of a user has changed. This update is guaranteed to come before the user identifier is returned to the client
type UpdateUser struct {
	tdCommon
	User User `json:"user"` // New data about the user
}

// MessageType return the string telegram-type of UpdateUser
func (updateUser *UpdateUser) MessageType() string {
	return "updateUser"
}

// GetUpdateEnum return the enum type of this object
func (updateUser *UpdateUser) GetUpdateEnum() UpdateEnum {
	return UpdateUserType
}

// UpdateBasicGroup Some data of a basic group has changed. This update is guaranteed to come before the basic group identifier is returned to the client
type UpdateBasicGroup struct {
	tdCommon
	BasicGroup BasicGroup `json:"basic_group"` // New data about the group
}

// MessageType return the string telegram-type of UpdateBasicGroup
func (updateBasicGroup *UpdateBasicGroup) MessageType() string {
	return "updateBasicGroup"
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroup *UpdateBasicGroup) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupType
}

// UpdateSupergroup Some data of a supergroup or a channel has changed. This update is guaranteed to come before the supergroup identifier is returned to the client
type UpdateSupergroup struct {
	tdCommon
	Supergroup Supergroup `json:"supergroup"` // New data about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroup
func (updateSupergroup *UpdateSupergroup) MessageType() string {
	return "updateSupergroup"
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroup *UpdateSupergroup) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupType
}

// UpdateSecretChat Some data of a secret chat has changed. This update is guaranteed to come before the secret chat identifier is returned to the client
type UpdateSecretChat struct {
	tdCommon
	SecretChat SecretChat `json:"secret_chat"` // New data about the secret chat
}

// MessageType return the string telegram-type of UpdateSecretChat
func (updateSecretChat *UpdateSecretChat) MessageType() string {
	return "updateSecretChat"
}

// GetUpdateEnum return the enum type of this object
func (updateSecretChat *UpdateSecretChat) GetUpdateEnum() UpdateEnum {
	return UpdateSecretChatType
}

// UpdateUserFullInfo Some data from userFullInfo has been changed
type UpdateUserFullInfo struct {
	tdCommon
	UserID       int32        `json:"user_id"`        // User identifier
	UserFullInfo UserFullInfo `json:"user_full_info"` // New full information about the user
}

// MessageType return the string telegram-type of UpdateUserFullInfo
func (updateUserFullInfo *UpdateUserFullInfo) MessageType() string {
	return "updateUserFullInfo"
}

// GetUpdateEnum return the enum type of this object
func (updateUserFullInfo *UpdateUserFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateUserFullInfoType
}

// UpdateBasicGroupFullInfo Some data from basicGroupFullInfo has been changed
type UpdateBasicGroupFullInfo struct {
	tdCommon
	BasicGroupID       int32              `json:"basic_group_id"`        // Identifier of a basic group
	BasicGroupFullInfo BasicGroupFullInfo `json:"basic_group_full_info"` // New full information about the group
}

// MessageType return the string telegram-type of UpdateBasicGroupFullInfo
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) MessageType() string {
	return "updateBasicGroupFullInfo"
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupFullInfoType
}

// UpdateSupergroupFullInfo Some data from supergroupFullInfo has been changed
type UpdateSupergroupFullInfo struct {
	tdCommon
	SupergroupID       int32              `json:"supergroup_id"`        // Identifier of the supergroup or channel
	SupergroupFullInfo SupergroupFullInfo `json:"supergroup_full_info"` // New full information about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroupFullInfo
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) MessageType() string {
	return "updateSupergroupFullInfo"
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupFullInfoType
}

// UpdateServiceNotification Service notification from the server. Upon receiving this the client must show a popup with the content of the notification
type UpdateServiceNotification struct {
	tdCommon
	Type    string         `json:"type"`    // Notification type
	Content MessageContent `json:"content"` // Notification content
}

// MessageType return the string telegram-type of UpdateServiceNotification
func (updateServiceNotification *UpdateServiceNotification) MessageType() string {
	return "updateServiceNotification"
}

// UnmarshalJSON unmarshal to json
func (updateServiceNotification *UpdateServiceNotification) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Type string `json:"type"` // Notification type

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateServiceNotification.tdCommon = tempObj.tdCommon
	updateServiceNotification.Type = tempObj.Type

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	updateServiceNotification.Content = fieldContent

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateServiceNotification *UpdateServiceNotification) GetUpdateEnum() UpdateEnum {
	return UpdateServiceNotificationType
}

// UpdateFile Information about a file was updated
type UpdateFile struct {
	tdCommon
	File File `json:"file"` // New data about the file
}

// MessageType return the string telegram-type of UpdateFile
func (updateFile *UpdateFile) MessageType() string {
	return "updateFile"
}

// GetUpdateEnum return the enum type of this object
func (updateFile *UpdateFile) GetUpdateEnum() UpdateEnum {
	return UpdateFileType
}

// UpdateFileGenerationStart The file generation process needs to be started by the client
type UpdateFileGenerationStart struct {
	tdCommon
	GenerationID    JSONInt64 `json:"generation_id"`    // Unique identifier for the generation process
	OriginalPath    string    `json:"original_path"`    // The path to a file from which a new file is generated, may be empty
	DestinationPath string    `json:"destination_path"` // The path to a file that should be created and where the new file should be generated
	Conversion      string    `json:"conversion"`       // String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains a HTTP/HTTPS URL of a file, which should be downloaded by the client
}

// MessageType return the string telegram-type of UpdateFileGenerationStart
func (updateFileGenerationStart *UpdateFileGenerationStart) MessageType() string {
	return "updateFileGenerationStart"
}

// GetUpdateEnum return the enum type of this object
func (updateFileGenerationStart *UpdateFileGenerationStart) GetUpdateEnum() UpdateEnum {
	return UpdateFileGenerationStartType
}

// UpdateFileGenerationStop File generation is no longer needed
type UpdateFileGenerationStop struct {
	tdCommon
	GenerationID JSONInt64 `json:"generation_id"` // Unique identifier for the generation process
}

// MessageType return the string telegram-type of UpdateFileGenerationStop
func (updateFileGenerationStop *UpdateFileGenerationStop) MessageType() string {
	return "updateFileGenerationStop"
}

// GetUpdateEnum return the enum type of this object
func (updateFileGenerationStop *UpdateFileGenerationStop) GetUpdateEnum() UpdateEnum {
	return UpdateFileGenerationStopType
}

// UpdateCall New call was created or information about a call was updated
type UpdateCall struct {
	tdCommon
	Call Call `json:"call"` // New data about a call
}

// MessageType return the string telegram-type of UpdateCall
func (updateCall *UpdateCall) MessageType() string {
	return "updateCall"
}

// GetUpdateEnum return the enum type of this object
func (updateCall *UpdateCall) GetUpdateEnum() UpdateEnum {
	return UpdateCallType
}

// UpdateUserPrivacySettingRules Some privacy setting rules have been changed
type UpdateUserPrivacySettingRules struct {
	tdCommon
	Setting UserPrivacySetting      `json:"setting"` // The privacy setting
	Rules   UserPrivacySettingRules `json:"rules"`   // New privacy rules
}

// MessageType return the string telegram-type of UpdateUserPrivacySettingRules
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) MessageType() string {
	return "updateUserPrivacySettingRules"
}

// UnmarshalJSON unmarshal to json
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Rules UserPrivacySettingRules `json:"rules"` // New privacy rules
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateUserPrivacySettingRules.tdCommon = tempObj.tdCommon
	updateUserPrivacySettingRules.Rules = tempObj.Rules

	fieldSetting, _ := unmarshalUserPrivacySetting(objMap["setting"])
	updateUserPrivacySettingRules.Setting = fieldSetting

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) GetUpdateEnum() UpdateEnum {
	return UpdateUserPrivacySettingRulesType
}

// UpdateUnreadMessageCount Number of unread messages has changed. This update is sent only if a message database is used
type UpdateUnreadMessageCount struct {
	tdCommon
	UnreadCount        int32 `json:"unread_count"`         // Total number of unread messages
	UnreadUnmutedCount int32 `json:"unread_unmuted_count"` // Total number of unread messages in unmuted chats
}

// MessageType return the string telegram-type of UpdateUnreadMessageCount
func (updateUnreadMessageCount *UpdateUnreadMessageCount) MessageType() string {
	return "updateUnreadMessageCount"
}

// GetUpdateEnum return the enum type of this object
func (updateUnreadMessageCount *UpdateUnreadMessageCount) GetUpdateEnum() UpdateEnum {
	return UpdateUnreadMessageCountType
}

// UpdateOption An option changed its value
type UpdateOption struct {
	tdCommon
	Name  string      `json:"name"`  // The option name
	Value OptionValue `json:"value"` // The new option value
}

// MessageType return the string telegram-type of UpdateOption
func (updateOption *UpdateOption) MessageType() string {
	return "updateOption"
}

// UnmarshalJSON unmarshal to json
func (updateOption *UpdateOption) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Name string `json:"name"` // The option name

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateOption.tdCommon = tempObj.tdCommon
	updateOption.Name = tempObj.Name

	fieldValue, _ := unmarshalOptionValue(objMap["value"])
	updateOption.Value = fieldValue

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateOption *UpdateOption) GetUpdateEnum() UpdateEnum {
	return UpdateOptionType
}

// UpdateInstalledStickerSets The list of installed sticker sets was updated
type UpdateInstalledStickerSets struct {
	tdCommon
	IsMasks       bool        `json:"is_masks"`        // True, if the list of installed mask sticker sets was updated
	StickerSetIDs []JSONInt64 `json:"sticker_set_ids"` // The new list of installed ordinary sticker sets
}

// MessageType return the string telegram-type of UpdateInstalledStickerSets
func (updateInstalledStickerSets *UpdateInstalledStickerSets) MessageType() string {
	return "updateInstalledStickerSets"
}

// GetUpdateEnum return the enum type of this object
func (updateInstalledStickerSets *UpdateInstalledStickerSets) GetUpdateEnum() UpdateEnum {
	return UpdateInstalledStickerSetsType
}

// UpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed
type UpdateTrendingStickerSets struct {
	tdCommon
	StickerSets StickerSets `json:"sticker_sets"` // The new list of trending sticker sets
}

// MessageType return the string telegram-type of UpdateTrendingStickerSets
func (updateTrendingStickerSets *UpdateTrendingStickerSets) MessageType() string {
	return "updateTrendingStickerSets"
}

// GetUpdateEnum return the enum type of this object
func (updateTrendingStickerSets *UpdateTrendingStickerSets) GetUpdateEnum() UpdateEnum {
	return UpdateTrendingStickerSetsType
}

// UpdateRecentStickers The list of recently used stickers was updated
type UpdateRecentStickers struct {
	tdCommon
	IsAttached bool    `json:"is_attached"` // True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated
	StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of recently used stickers
}

// MessageType return the string telegram-type of UpdateRecentStickers
func (updateRecentStickers *UpdateRecentStickers) MessageType() string {
	return "updateRecentStickers"
}

// GetUpdateEnum return the enum type of this object
func (updateRecentStickers *UpdateRecentStickers) GetUpdateEnum() UpdateEnum {
	return UpdateRecentStickersType
}

// UpdateFavoriteStickers The list of favorite stickers was updated
type UpdateFavoriteStickers struct {
	tdCommon
	StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of favorite stickers
}

// MessageType return the string telegram-type of UpdateFavoriteStickers
func (updateFavoriteStickers *UpdateFavoriteStickers) MessageType() string {
	return "updateFavoriteStickers"
}

// GetUpdateEnum return the enum type of this object
func (updateFavoriteStickers *UpdateFavoriteStickers) GetUpdateEnum() UpdateEnum {
	return UpdateFavoriteStickersType
}

// UpdateSavedAnimations The list of saved animations was updated
type UpdateSavedAnimations struct {
	tdCommon
	AnimationIDs []int32 `json:"animation_ids"` // The new list of file identifiers of saved animations
}

// MessageType return the string telegram-type of UpdateSavedAnimations
func (updateSavedAnimations *UpdateSavedAnimations) MessageType() string {
	return "updateSavedAnimations"
}

// GetUpdateEnum return the enum type of this object
func (updateSavedAnimations *UpdateSavedAnimations) GetUpdateEnum() UpdateEnum {
	return UpdateSavedAnimationsType
}

// UpdateConnectionState The connection state has changed
type UpdateConnectionState struct {
	tdCommon
	State ConnectionState `json:"state"` // The new connection state
}

// MessageType return the string telegram-type of UpdateConnectionState
func (updateConnectionState *UpdateConnectionState) MessageType() string {
	return "updateConnectionState"
}

// UnmarshalJSON unmarshal to json
func (updateConnectionState *UpdateConnectionState) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateConnectionState.tdCommon = tempObj.tdCommon

	fieldState, _ := unmarshalConnectionState(objMap["state"])
	updateConnectionState.State = fieldState

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateConnectionState *UpdateConnectionState) GetUpdateEnum() UpdateEnum {
	return UpdateConnectionStateType
}

// UpdateNewInlineQuery A new incoming inline query; for bots only
type UpdateNewInlineQuery struct {
	tdCommon
	ID           JSONInt64 `json:"id"`             // Unique query identifier
	SenderUserID int32     `json:"sender_user_id"` // Identifier of the user who sent the query
	UserLocation Location  `json:"user_location"`  // User location, provided by the client; may be null
	Query        string    `json:"query"`          // Text of the query
	Offset       string    `json:"offset"`         // Offset of the first entry to return
}

// MessageType return the string telegram-type of UpdateNewInlineQuery
func (updateNewInlineQuery *UpdateNewInlineQuery) MessageType() string {
	return "updateNewInlineQuery"
}

// GetUpdateEnum return the enum type of this object
func (updateNewInlineQuery *UpdateNewInlineQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewInlineQueryType
}

// UpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only
type UpdateNewChosenInlineResult struct {
	tdCommon
	SenderUserID    int32    `json:"sender_user_id"`    // Identifier of the user who sent the query
	UserLocation    Location `json:"user_location"`     // User location, provided by the client; may be null
	Query           string   `json:"query"`             // Text of the query
	ResultID        string   `json:"result_id"`         // Identifier of the chosen result
	InlineMessageID string   `json:"inline_message_id"` // Identifier of the sent inline message, if known
}

// MessageType return the string telegram-type of UpdateNewChosenInlineResult
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) MessageType() string {
	return "updateNewChosenInlineResult"
}

// GetUpdateEnum return the enum type of this object
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) GetUpdateEnum() UpdateEnum {
	return UpdateNewChosenInlineResultType
}

// UpdateNewCallbackQuery A new incoming callback query; for bots only
type UpdateNewCallbackQuery struct {
	tdCommon
	ID           JSONInt64            `json:"id"`             // Unique query identifier
	SenderUserID int32                `json:"sender_user_id"` // Identifier of the user who sent the query
	ChatID       int64                `json:"chat_id"`        // Identifier of the chat, in which the query was sent
	MessageID    int64                `json:"message_id"`     // Identifier of the message, from which the query originated
	ChatInstance JSONInt64            `json:"chat_instance"`  // Identifier that uniquely corresponds to the chat to which the message was sent
	Payload      CallbackQueryPayload `json:"payload"`        // Query payload
}

// MessageType return the string telegram-type of UpdateNewCallbackQuery
func (updateNewCallbackQuery *UpdateNewCallbackQuery) MessageType() string {
	return "updateNewCallbackQuery"
}

// UnmarshalJSON unmarshal to json
func (updateNewCallbackQuery *UpdateNewCallbackQuery) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID           JSONInt64 `json:"id"`             // Unique query identifier
		SenderUserID int32     `json:"sender_user_id"` // Identifier of the user who sent the query
		ChatID       int64     `json:"chat_id"`        // Identifier of the chat, in which the query was sent
		MessageID    int64     `json:"message_id"`     // Identifier of the message, from which the query originated
		ChatInstance JSONInt64 `json:"chat_instance"`  // Identifier that uniquely corresponds to the chat to which the message was sent

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNewCallbackQuery.tdCommon = tempObj.tdCommon
	updateNewCallbackQuery.ID = tempObj.ID
	updateNewCallbackQuery.SenderUserID = tempObj.SenderUserID
	updateNewCallbackQuery.ChatID = tempObj.ChatID
	updateNewCallbackQuery.MessageID = tempObj.MessageID
	updateNewCallbackQuery.ChatInstance = tempObj.ChatInstance

	fieldPayload, _ := unmarshalCallbackQueryPayload(objMap["payload"])
	updateNewCallbackQuery.Payload = fieldPayload

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNewCallbackQuery *UpdateNewCallbackQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewCallbackQueryType
}

// UpdateNewInlineCallbackQuery A new incoming callback query from a message sent via a bot; for bots only
type UpdateNewInlineCallbackQuery struct {
	tdCommon
	ID              JSONInt64            `json:"id"`                // Unique query identifier
	SenderUserID    int32                `json:"sender_user_id"`    // Identifier of the user who sent the query
	InlineMessageID string               `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
	ChatInstance    JSONInt64            `json:"chat_instance"`     // An identifier uniquely corresponding to the chat a message was sent to
	Payload         CallbackQueryPayload `json:"payload"`           // Query payload
}

// MessageType return the string telegram-type of UpdateNewInlineCallbackQuery
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) MessageType() string {
	return "updateNewInlineCallbackQuery"
}

// UnmarshalJSON unmarshal to json
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID              JSONInt64 `json:"id"`                // Unique query identifier
		SenderUserID    int32     `json:"sender_user_id"`    // Identifier of the user who sent the query
		InlineMessageID string    `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
		ChatInstance    JSONInt64 `json:"chat_instance"`     // An identifier uniquely corresponding to the chat a message was sent to

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateNewInlineCallbackQuery.tdCommon = tempObj.tdCommon
	updateNewInlineCallbackQuery.ID = tempObj.ID
	updateNewInlineCallbackQuery.SenderUserID = tempObj.SenderUserID
	updateNewInlineCallbackQuery.InlineMessageID = tempObj.InlineMessageID
	updateNewInlineCallbackQuery.ChatInstance = tempObj.ChatInstance

	fieldPayload, _ := unmarshalCallbackQueryPayload(objMap["payload"])
	updateNewInlineCallbackQuery.Payload = fieldPayload

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewInlineCallbackQueryType
}

// UpdateNewShippingQuery A new incoming shipping query; for bots only. Only for invoices with flexible price
type UpdateNewShippingQuery struct {
	tdCommon
	ID              JSONInt64       `json:"id"`               // Unique query identifier
	SenderUserID    int32           `json:"sender_user_id"`   // Identifier of the user who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User shipping address
}

// MessageType return the string telegram-type of UpdateNewShippingQuery
func (updateNewShippingQuery *UpdateNewShippingQuery) MessageType() string {
	return "updateNewShippingQuery"
}

// GetUpdateEnum return the enum type of this object
func (updateNewShippingQuery *UpdateNewShippingQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewShippingQueryType
}

// UpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout
type UpdateNewPreCheckoutQuery struct {
	tdCommon
	ID               JSONInt64 `json:"id"`                 // Unique query identifier
	SenderUserID     int32     `json:"sender_user_id"`     // Identifier of the user who sent the query
	Currency         string    `json:"currency"`           // Currency for the product price
	TotalAmount      int64     `json:"total_amount"`       // Total price for the product, in the minimal quantity of the currency
	InvoicePayload   []byte    `json:"invoice_payload"`    // Invoice payload
	ShippingOptionID string    `json:"shipping_option_id"` // Identifier of a shipping option chosen by the user; may be empty if not applicable
	OrderInfo        OrderInfo `json:"order_info"`         // Information about the order; may be null
}

// MessageType return the string telegram-type of UpdateNewPreCheckoutQuery
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) MessageType() string {
	return "updateNewPreCheckoutQuery"
}

// GetUpdateEnum return the enum type of this object
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewPreCheckoutQueryType
}

// UpdateNewCustomEvent A new incoming event; for bots only
type UpdateNewCustomEvent struct {
	tdCommon
	Event string `json:"event"` // A JSON-serialized event
}

// MessageType return the string telegram-type of UpdateNewCustomEvent
func (updateNewCustomEvent *UpdateNewCustomEvent) MessageType() string {
	return "updateNewCustomEvent"
}

// GetUpdateEnum return the enum type of this object
func (updateNewCustomEvent *UpdateNewCustomEvent) GetUpdateEnum() UpdateEnum {
	return UpdateNewCustomEventType
}

// UpdateNewCustomQuery A new incoming query; for bots only
type UpdateNewCustomQuery struct {
	tdCommon
	ID      JSONInt64 `json:"id"`      // The query identifier
	Data    string    `json:"data"`    // JSON-serialized query data
	Timeout int32     `json:"timeout"` // Query timeout
}

// MessageType return the string telegram-type of UpdateNewCustomQuery
func (updateNewCustomQuery *UpdateNewCustomQuery) MessageType() string {
	return "updateNewCustomQuery"
}

// GetUpdateEnum return the enum type of this object
func (updateNewCustomQuery *UpdateNewCustomQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewCustomQueryType
}

// TestInt A simple object containing a number; for testing only
type TestInt struct {
	tdCommon
	Value int32 `json:"value"` // Number
}

// MessageType return the string telegram-type of TestInt
func (testInt *TestInt) MessageType() string {
	return "testInt"
}

// TestString A simple object containing a string; for testing only
type TestString struct {
	tdCommon
	Value string `json:"value"` // String
}

// MessageType return the string telegram-type of TestString
func (testString *TestString) MessageType() string {
	return "testString"
}

// TestBytes A simple object containing a sequence of bytes; for testing only
type TestBytes struct {
	tdCommon
	Value []byte `json:"value"` // Bytes
}

// MessageType return the string telegram-type of TestBytes
func (testBytes *TestBytes) MessageType() string {
	return "testBytes"
}

// TestVectorInt A simple object containing a vector of numbers; for testing only
type TestVectorInt struct {
	tdCommon
	Value []int32 `json:"value"` // Vector of numbers
}

// MessageType return the string telegram-type of TestVectorInt
func (testVectorInt *TestVectorInt) MessageType() string {
	return "testVectorInt"
}

// TestVectorIntObject A simple object containing a vector of objects that hold a number; for testing only
type TestVectorIntObject struct {
	tdCommon
	Value []TestInt `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorIntObject
func (testVectorIntObject *TestVectorIntObject) MessageType() string {
	return "testVectorIntObject"
}

// TestVectorString A simple object containing a vector of strings; for testing only
type TestVectorString struct {
	tdCommon
	Value []string `json:"value"` // Vector of strings
}

// MessageType return the string telegram-type of TestVectorString
func (testVectorString *TestVectorString) MessageType() string {
	return "testVectorString"
}

// TestVectorStringObject A simple object containing a vector of objects that hold a string; for testing only
type TestVectorStringObject struct {
	tdCommon
	Value []TestString `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorStringObject
func (testVectorStringObject *TestVectorStringObject) MessageType() string {
	return "testVectorStringObject"
}

func unmarshalAuthenticationCodeType(rawMsg *json.RawMessage) (AuthenticationCodeType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch AuthenticationCodeTypeEnum(objMap["@type"].(string)) {
	case AuthenticationCodeTypeTelegramMessageType:
		var authenticationCodeTypeTelegramMessage AuthenticationCodeTypeTelegramMessage
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeTelegramMessage)
		return &authenticationCodeTypeTelegramMessage, err

	case AuthenticationCodeTypeSmsType:
		var authenticationCodeTypeSms AuthenticationCodeTypeSms
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeSms)
		return &authenticationCodeTypeSms, err

	case AuthenticationCodeTypeCallType:
		var authenticationCodeTypeCall AuthenticationCodeTypeCall
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeCall)
		return &authenticationCodeTypeCall, err

	case AuthenticationCodeTypeFlashCallType:
		var authenticationCodeTypeFlashCall AuthenticationCodeTypeFlashCall
		err := json.Unmarshal(*rawMsg, &authenticationCodeTypeFlashCall)
		return &authenticationCodeTypeFlashCall, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalAuthorizationState(rawMsg *json.RawMessage) (AuthorizationState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch AuthorizationStateEnum(objMap["@type"].(string)) {
	case AuthorizationStateWaitTdlibParametersType:
		var authorizationStateWaitTdlibParameters AuthorizationStateWaitTdlibParameters
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitTdlibParameters)
		return &authorizationStateWaitTdlibParameters, err

	case AuthorizationStateWaitEncryptionKeyType:
		var authorizationStateWaitEncryptionKey AuthorizationStateWaitEncryptionKey
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitEncryptionKey)
		return &authorizationStateWaitEncryptionKey, err

	case AuthorizationStateWaitPhoneNumberType:
		var authorizationStateWaitPhoneNumber AuthorizationStateWaitPhoneNumber
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitPhoneNumber)
		return &authorizationStateWaitPhoneNumber, err

	case AuthorizationStateWaitCodeType:
		var authorizationStateWaitCode AuthorizationStateWaitCode
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitCode)
		return &authorizationStateWaitCode, err

	case AuthorizationStateWaitPasswordType:
		var authorizationStateWaitPassword AuthorizationStateWaitPassword
		err := json.Unmarshal(*rawMsg, &authorizationStateWaitPassword)
		return &authorizationStateWaitPassword, err

	case AuthorizationStateReadyType:
		var authorizationStateReady AuthorizationStateReady
		err := json.Unmarshal(*rawMsg, &authorizationStateReady)
		return &authorizationStateReady, err

	case AuthorizationStateLoggingOutType:
		var authorizationStateLoggingOut AuthorizationStateLoggingOut
		err := json.Unmarshal(*rawMsg, &authorizationStateLoggingOut)
		return &authorizationStateLoggingOut, err

	case AuthorizationStateClosingType:
		var authorizationStateClosing AuthorizationStateClosing
		err := json.Unmarshal(*rawMsg, &authorizationStateClosing)
		return &authorizationStateClosing, err

	case AuthorizationStateClosedType:
		var authorizationStateClosed AuthorizationStateClosed
		err := json.Unmarshal(*rawMsg, &authorizationStateClosed)
		return &authorizationStateClosed, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputFile(rawMsg *json.RawMessage) (InputFile, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputFileEnum(objMap["@type"].(string)) {
	case InputFileIDType:
		var inputFileID InputFileID
		err := json.Unmarshal(*rawMsg, &inputFileID)
		return &inputFileID, err

	case InputFileRemoteType:
		var inputFileRemote InputFileRemote
		err := json.Unmarshal(*rawMsg, &inputFileRemote)
		return &inputFileRemote, err

	case InputFileLocalType:
		var inputFileLocal InputFileLocal
		err := json.Unmarshal(*rawMsg, &inputFileLocal)
		return &inputFileLocal, err

	case InputFileGeneratedType:
		var inputFileGenerated InputFileGenerated
		err := json.Unmarshal(*rawMsg, &inputFileGenerated)
		return &inputFileGenerated, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalMaskPoint(rawMsg *json.RawMessage) (MaskPoint, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MaskPointEnum(objMap["@type"].(string)) {
	case MaskPointForeheadType:
		var maskPointForehead MaskPointForehead
		err := json.Unmarshal(*rawMsg, &maskPointForehead)
		return &maskPointForehead, err

	case MaskPointEyesType:
		var maskPointEyes MaskPointEyes
		err := json.Unmarshal(*rawMsg, &maskPointEyes)
		return &maskPointEyes, err

	case MaskPointMouthType:
		var maskPointMouth MaskPointMouth
		err := json.Unmarshal(*rawMsg, &maskPointMouth)
		return &maskPointMouth, err

	case MaskPointChinType:
		var maskPointChin MaskPointChin
		err := json.Unmarshal(*rawMsg, &maskPointChin)
		return &maskPointChin, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalLinkState(rawMsg *json.RawMessage) (LinkState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LinkStateEnum(objMap["@type"].(string)) {
	case LinkStateNoneType:
		var linkStateNone LinkStateNone
		err := json.Unmarshal(*rawMsg, &linkStateNone)
		return &linkStateNone, err

	case LinkStateKnowsPhoneNumberType:
		var linkStateKnowsPhoneNumber LinkStateKnowsPhoneNumber
		err := json.Unmarshal(*rawMsg, &linkStateKnowsPhoneNumber)
		return &linkStateKnowsPhoneNumber, err

	case LinkStateIsContactType:
		var linkStateIsContact LinkStateIsContact
		err := json.Unmarshal(*rawMsg, &linkStateIsContact)
		return &linkStateIsContact, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalUserType(rawMsg *json.RawMessage) (UserType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserTypeEnum(objMap["@type"].(string)) {
	case UserTypeRegularType:
		var userTypeRegular UserTypeRegular
		err := json.Unmarshal(*rawMsg, &userTypeRegular)
		return &userTypeRegular, err

	case UserTypeDeletedType:
		var userTypeDeleted UserTypeDeleted
		err := json.Unmarshal(*rawMsg, &userTypeDeleted)
		return &userTypeDeleted, err

	case UserTypeBotType:
		var userTypeBot UserTypeBot
		err := json.Unmarshal(*rawMsg, &userTypeBot)
		return &userTypeBot, err

	case UserTypeUnknownType:
		var userTypeUnknown UserTypeUnknown
		err := json.Unmarshal(*rawMsg, &userTypeUnknown)
		return &userTypeUnknown, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalChatMemberStatus(rawMsg *json.RawMessage) (ChatMemberStatus, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatMemberStatusEnum(objMap["@type"].(string)) {
	case ChatMemberStatusCreatorType:
		var chatMemberStatusCreator ChatMemberStatusCreator
		err := json.Unmarshal(*rawMsg, &chatMemberStatusCreator)
		return &chatMemberStatusCreator, err

	case ChatMemberStatusAdministratorType:
		var chatMemberStatusAdministrator ChatMemberStatusAdministrator
		err := json.Unmarshal(*rawMsg, &chatMemberStatusAdministrator)
		return &chatMemberStatusAdministrator, err

	case ChatMemberStatusMemberType:
		var chatMemberStatusMember ChatMemberStatusMember
		err := json.Unmarshal(*rawMsg, &chatMemberStatusMember)
		return &chatMemberStatusMember, err

	case ChatMemberStatusRestrictedType:
		var chatMemberStatusRestricted ChatMemberStatusRestricted
		err := json.Unmarshal(*rawMsg, &chatMemberStatusRestricted)
		return &chatMemberStatusRestricted, err

	case ChatMemberStatusLeftType:
		var chatMemberStatusLeft ChatMemberStatusLeft
		err := json.Unmarshal(*rawMsg, &chatMemberStatusLeft)
		return &chatMemberStatusLeft, err

	case ChatMemberStatusBannedType:
		var chatMemberStatusBanned ChatMemberStatusBanned
		err := json.Unmarshal(*rawMsg, &chatMemberStatusBanned)
		return &chatMemberStatusBanned, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalSupergroupMembersFilter(rawMsg *json.RawMessage) (SupergroupMembersFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SupergroupMembersFilterEnum(objMap["@type"].(string)) {
	case SupergroupMembersFilterRecentType:
		var supergroupMembersFilterRecent SupergroupMembersFilterRecent
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterRecent)
		return &supergroupMembersFilterRecent, err

	case SupergroupMembersFilterAdministratorsType:
		var supergroupMembersFilterAdministrators SupergroupMembersFilterAdministrators
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterAdministrators)
		return &supergroupMembersFilterAdministrators, err

	case SupergroupMembersFilterSearchType:
		var supergroupMembersFilterSearch SupergroupMembersFilterSearch
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterSearch)
		return &supergroupMembersFilterSearch, err

	case SupergroupMembersFilterRestrictedType:
		var supergroupMembersFilterRestricted SupergroupMembersFilterRestricted
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterRestricted)
		return &supergroupMembersFilterRestricted, err

	case SupergroupMembersFilterBannedType:
		var supergroupMembersFilterBanned SupergroupMembersFilterBanned
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterBanned)
		return &supergroupMembersFilterBanned, err

	case SupergroupMembersFilterBotsType:
		var supergroupMembersFilterBots SupergroupMembersFilterBots
		err := json.Unmarshal(*rawMsg, &supergroupMembersFilterBots)
		return &supergroupMembersFilterBots, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalSecretChatState(rawMsg *json.RawMessage) (SecretChatState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SecretChatStateEnum(objMap["@type"].(string)) {
	case SecretChatStatePendingType:
		var secretChatStatePending SecretChatStatePending
		err := json.Unmarshal(*rawMsg, &secretChatStatePending)
		return &secretChatStatePending, err

	case SecretChatStateReadyType:
		var secretChatStateReady SecretChatStateReady
		err := json.Unmarshal(*rawMsg, &secretChatStateReady)
		return &secretChatStateReady, err

	case SecretChatStateClosedType:
		var secretChatStateClosed SecretChatStateClosed
		err := json.Unmarshal(*rawMsg, &secretChatStateClosed)
		return &secretChatStateClosed, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalMessageForwardInfo(rawMsg *json.RawMessage) (MessageForwardInfo, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageForwardInfoEnum(objMap["@type"].(string)) {
	case MessageForwardedFromUserType:
		var messageForwardedFromUser MessageForwardedFromUser
		err := json.Unmarshal(*rawMsg, &messageForwardedFromUser)
		return &messageForwardedFromUser, err

	case MessageForwardedPostType:
		var messageForwardedPost MessageForwardedPost
		err := json.Unmarshal(*rawMsg, &messageForwardedPost)
		return &messageForwardedPost, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalMessageSendingState(rawMsg *json.RawMessage) (MessageSendingState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageSendingStateEnum(objMap["@type"].(string)) {
	case MessageSendingStatePendingType:
		var messageSendingStatePending MessageSendingStatePending
		err := json.Unmarshal(*rawMsg, &messageSendingStatePending)
		return &messageSendingStatePending, err

	case MessageSendingStateFailedType:
		var messageSendingStateFailed MessageSendingStateFailed
		err := json.Unmarshal(*rawMsg, &messageSendingStateFailed)
		return &messageSendingStateFailed, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalNotificationSettingsScope(rawMsg *json.RawMessage) (NotificationSettingsScope, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NotificationSettingsScopeEnum(objMap["@type"].(string)) {
	case NotificationSettingsScopeChatType:
		var notificationSettingsScopeChat NotificationSettingsScopeChat
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeChat)
		return &notificationSettingsScopeChat, err

	case NotificationSettingsScopePrivateChatsType:
		var notificationSettingsScopePrivateChats NotificationSettingsScopePrivateChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopePrivateChats)
		return &notificationSettingsScopePrivateChats, err

	case NotificationSettingsScopeBasicGroupChatsType:
		var notificationSettingsScopeBasicGroupChats NotificationSettingsScopeBasicGroupChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeBasicGroupChats)
		return &notificationSettingsScopeBasicGroupChats, err

	case NotificationSettingsScopeAllChatsType:
		var notificationSettingsScopeAllChats NotificationSettingsScopeAllChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeAllChats)
		return &notificationSettingsScopeAllChats, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalChatType(rawMsg *json.RawMessage) (ChatType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatTypeEnum(objMap["@type"].(string)) {
	case ChatTypePrivateType:
		var chatTypePrivate ChatTypePrivate
		err := json.Unmarshal(*rawMsg, &chatTypePrivate)
		return &chatTypePrivate, err

	case ChatTypeBasicGroupType:
		var chatTypeBasicGroup ChatTypeBasicGroup
		err := json.Unmarshal(*rawMsg, &chatTypeBasicGroup)
		return &chatTypeBasicGroup, err

	case ChatTypeSupergroupType:
		var chatTypeSupergroup ChatTypeSupergroup
		err := json.Unmarshal(*rawMsg, &chatTypeSupergroup)
		return &chatTypeSupergroup, err

	case ChatTypeSecretType:
		var chatTypeSecret ChatTypeSecret
		err := json.Unmarshal(*rawMsg, &chatTypeSecret)
		return &chatTypeSecret, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalKeyboardButtonType(rawMsg *json.RawMessage) (KeyboardButtonType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch KeyboardButtonTypeEnum(objMap["@type"].(string)) {
	case KeyboardButtonTypeTextType:
		var keyboardButtonTypeText KeyboardButtonTypeText
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeText)
		return &keyboardButtonTypeText, err

	case KeyboardButtonTypeRequestPhoneNumberType:
		var keyboardButtonTypeRequestPhoneNumber KeyboardButtonTypeRequestPhoneNumber
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeRequestPhoneNumber)
		return &keyboardButtonTypeRequestPhoneNumber, err

	case KeyboardButtonTypeRequestLocationType:
		var keyboardButtonTypeRequestLocation KeyboardButtonTypeRequestLocation
		err := json.Unmarshal(*rawMsg, &keyboardButtonTypeRequestLocation)
		return &keyboardButtonTypeRequestLocation, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInlineKeyboardButtonType(rawMsg *json.RawMessage) (InlineKeyboardButtonType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InlineKeyboardButtonTypeEnum(objMap["@type"].(string)) {
	case InlineKeyboardButtonTypeURLType:
		var inlineKeyboardButtonTypeURL InlineKeyboardButtonTypeURL
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeURL)
		return &inlineKeyboardButtonTypeURL, err

	case InlineKeyboardButtonTypeCallbackType:
		var inlineKeyboardButtonTypeCallback InlineKeyboardButtonTypeCallback
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeCallback)
		return &inlineKeyboardButtonTypeCallback, err

	case InlineKeyboardButtonTypeCallbackGameType:
		var inlineKeyboardButtonTypeCallbackGame InlineKeyboardButtonTypeCallbackGame
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeCallbackGame)
		return &inlineKeyboardButtonTypeCallbackGame, err

	case InlineKeyboardButtonTypeSwitchInlineType:
		var inlineKeyboardButtonTypeSwitchInline InlineKeyboardButtonTypeSwitchInline
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeSwitchInline)
		return &inlineKeyboardButtonTypeSwitchInline, err

	case InlineKeyboardButtonTypeBuyType:
		var inlineKeyboardButtonTypeBuy InlineKeyboardButtonTypeBuy
		err := json.Unmarshal(*rawMsg, &inlineKeyboardButtonTypeBuy)
		return &inlineKeyboardButtonTypeBuy, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalReplyMarkup(rawMsg *json.RawMessage) (ReplyMarkup, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ReplyMarkupEnum(objMap["@type"].(string)) {
	case ReplyMarkupRemoveKeyboardType:
		var replyMarkupRemoveKeyboard ReplyMarkupRemoveKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupRemoveKeyboard)
		return &replyMarkupRemoveKeyboard, err

	case ReplyMarkupForceReplyType:
		var replyMarkupForceReply ReplyMarkupForceReply
		err := json.Unmarshal(*rawMsg, &replyMarkupForceReply)
		return &replyMarkupForceReply, err

	case ReplyMarkupShowKeyboardType:
		var replyMarkupShowKeyboard ReplyMarkupShowKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupShowKeyboard)
		return &replyMarkupShowKeyboard, err

	case ReplyMarkupInlineKeyboardType:
		var replyMarkupInlineKeyboard ReplyMarkupInlineKeyboard
		err := json.Unmarshal(*rawMsg, &replyMarkupInlineKeyboard)
		return &replyMarkupInlineKeyboard, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalRichText(rawMsg *json.RawMessage) (RichText, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch RichTextEnum(objMap["@type"].(string)) {
	case RichTextPlainType:
		var richTextPlain RichTextPlain
		err := json.Unmarshal(*rawMsg, &richTextPlain)
		return &richTextPlain, err

	case RichTextBoldType:
		var richTextBold RichTextBold
		err := json.Unmarshal(*rawMsg, &richTextBold)
		return &richTextBold, err

	case RichTextItalicType:
		var richTextItalic RichTextItalic
		err := json.Unmarshal(*rawMsg, &richTextItalic)
		return &richTextItalic, err

	case RichTextUnderlineType:
		var richTextUnderline RichTextUnderline
		err := json.Unmarshal(*rawMsg, &richTextUnderline)
		return &richTextUnderline, err

	case RichTextStrikethroughType:
		var richTextStrikethrough RichTextStrikethrough
		err := json.Unmarshal(*rawMsg, &richTextStrikethrough)
		return &richTextStrikethrough, err

	case RichTextFixedType:
		var richTextFixed RichTextFixed
		err := json.Unmarshal(*rawMsg, &richTextFixed)
		return &richTextFixed, err

	case RichTextURLType:
		var richTextURL RichTextURL
		err := json.Unmarshal(*rawMsg, &richTextURL)
		return &richTextURL, err

	case RichTextEmailAddressType:
		var richTextEmailAddress RichTextEmailAddress
		err := json.Unmarshal(*rawMsg, &richTextEmailAddress)
		return &richTextEmailAddress, err

	case RichTextsType:
		var richTexts RichTexts
		err := json.Unmarshal(*rawMsg, &richTexts)
		return &richTexts, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalPageBlock(rawMsg *json.RawMessage) (PageBlock, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PageBlockEnum(objMap["@type"].(string)) {
	case PageBlockTitleType:
		var pageBlockTitle PageBlockTitle
		err := json.Unmarshal(*rawMsg, &pageBlockTitle)
		return &pageBlockTitle, err

	case PageBlockSubtitleType:
		var pageBlockSubtitle PageBlockSubtitle
		err := json.Unmarshal(*rawMsg, &pageBlockSubtitle)
		return &pageBlockSubtitle, err

	case PageBlockAuthorDateType:
		var pageBlockAuthorDate PageBlockAuthorDate
		err := json.Unmarshal(*rawMsg, &pageBlockAuthorDate)
		return &pageBlockAuthorDate, err

	case PageBlockHeaderType:
		var pageBlockHeader PageBlockHeader
		err := json.Unmarshal(*rawMsg, &pageBlockHeader)
		return &pageBlockHeader, err

	case PageBlockSubheaderType:
		var pageBlockSubheader PageBlockSubheader
		err := json.Unmarshal(*rawMsg, &pageBlockSubheader)
		return &pageBlockSubheader, err

	case PageBlockParagraphType:
		var pageBlockParagraph PageBlockParagraph
		err := json.Unmarshal(*rawMsg, &pageBlockParagraph)
		return &pageBlockParagraph, err

	case PageBlockPreformattedType:
		var pageBlockPreformatted PageBlockPreformatted
		err := json.Unmarshal(*rawMsg, &pageBlockPreformatted)
		return &pageBlockPreformatted, err

	case PageBlockFooterType:
		var pageBlockFooter PageBlockFooter
		err := json.Unmarshal(*rawMsg, &pageBlockFooter)
		return &pageBlockFooter, err

	case PageBlockDividerType:
		var pageBlockDivider PageBlockDivider
		err := json.Unmarshal(*rawMsg, &pageBlockDivider)
		return &pageBlockDivider, err

	case PageBlockAnchorType:
		var pageBlockAnchor PageBlockAnchor
		err := json.Unmarshal(*rawMsg, &pageBlockAnchor)
		return &pageBlockAnchor, err

	case PageBlockListType:
		var pageBlockList PageBlockList
		err := json.Unmarshal(*rawMsg, &pageBlockList)
		return &pageBlockList, err

	case PageBlockBlockQuoteType:
		var pageBlockBlockQuote PageBlockBlockQuote
		err := json.Unmarshal(*rawMsg, &pageBlockBlockQuote)
		return &pageBlockBlockQuote, err

	case PageBlockPullQuoteType:
		var pageBlockPullQuote PageBlockPullQuote
		err := json.Unmarshal(*rawMsg, &pageBlockPullQuote)
		return &pageBlockPullQuote, err

	case PageBlockAnimationType:
		var pageBlockAnimation PageBlockAnimation
		err := json.Unmarshal(*rawMsg, &pageBlockAnimation)
		return &pageBlockAnimation, err

	case PageBlockAudioType:
		var pageBlockAudio PageBlockAudio
		err := json.Unmarshal(*rawMsg, &pageBlockAudio)
		return &pageBlockAudio, err

	case PageBlockPhotoType:
		var pageBlockPhoto PageBlockPhoto
		err := json.Unmarshal(*rawMsg, &pageBlockPhoto)
		return &pageBlockPhoto, err

	case PageBlockVideoType:
		var pageBlockVideo PageBlockVideo
		err := json.Unmarshal(*rawMsg, &pageBlockVideo)
		return &pageBlockVideo, err

	case PageBlockCoverType:
		var pageBlockCover PageBlockCover
		err := json.Unmarshal(*rawMsg, &pageBlockCover)
		return &pageBlockCover, err

	case PageBlockEmbeddedType:
		var pageBlockEmbedded PageBlockEmbedded
		err := json.Unmarshal(*rawMsg, &pageBlockEmbedded)
		return &pageBlockEmbedded, err

	case PageBlockEmbeddedPostType:
		var pageBlockEmbeddedPost PageBlockEmbeddedPost
		err := json.Unmarshal(*rawMsg, &pageBlockEmbeddedPost)
		return &pageBlockEmbeddedPost, err

	case PageBlockCollageType:
		var pageBlockCollage PageBlockCollage
		err := json.Unmarshal(*rawMsg, &pageBlockCollage)
		return &pageBlockCollage, err

	case PageBlockSlideshowType:
		var pageBlockSlideshow PageBlockSlideshow
		err := json.Unmarshal(*rawMsg, &pageBlockSlideshow)
		return &pageBlockSlideshow, err

	case PageBlockChatLinkType:
		var pageBlockChatLink PageBlockChatLink
		err := json.Unmarshal(*rawMsg, &pageBlockChatLink)
		return &pageBlockChatLink, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputCredentials(rawMsg *json.RawMessage) (InputCredentials, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputCredentialsEnum(objMap["@type"].(string)) {
	case InputCredentialsSavedType:
		var inputCredentialsSaved InputCredentialsSaved
		err := json.Unmarshal(*rawMsg, &inputCredentialsSaved)
		return &inputCredentialsSaved, err

	case InputCredentialsNewType:
		var inputCredentialsNew InputCredentialsNew
		err := json.Unmarshal(*rawMsg, &inputCredentialsNew)
		return &inputCredentialsNew, err

	case InputCredentialsAndroidPayType:
		var inputCredentialsAndroidPay InputCredentialsAndroidPay
		err := json.Unmarshal(*rawMsg, &inputCredentialsAndroidPay)
		return &inputCredentialsAndroidPay, err

	case InputCredentialsApplePayType:
		var inputCredentialsApplePay InputCredentialsApplePay
		err := json.Unmarshal(*rawMsg, &inputCredentialsApplePay)
		return &inputCredentialsApplePay, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalMessageContent(rawMsg *json.RawMessage) (MessageContent, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch MessageContentEnum(objMap["@type"].(string)) {
	case MessageTextType:
		var messageText MessageText
		err := json.Unmarshal(*rawMsg, &messageText)
		return &messageText, err

	case MessageAnimationType:
		var messageAnimation MessageAnimation
		err := json.Unmarshal(*rawMsg, &messageAnimation)
		return &messageAnimation, err

	case MessageAudioType:
		var messageAudio MessageAudio
		err := json.Unmarshal(*rawMsg, &messageAudio)
		return &messageAudio, err

	case MessageDocumentType:
		var messageDocument MessageDocument
		err := json.Unmarshal(*rawMsg, &messageDocument)
		return &messageDocument, err

	case MessagePhotoType:
		var messagePhoto MessagePhoto
		err := json.Unmarshal(*rawMsg, &messagePhoto)
		return &messagePhoto, err

	case MessageExpiredPhotoType:
		var messageExpiredPhoto MessageExpiredPhoto
		err := json.Unmarshal(*rawMsg, &messageExpiredPhoto)
		return &messageExpiredPhoto, err

	case MessageStickerType:
		var messageSticker MessageSticker
		err := json.Unmarshal(*rawMsg, &messageSticker)
		return &messageSticker, err

	case MessageVideoType:
		var messageVideo MessageVideo
		err := json.Unmarshal(*rawMsg, &messageVideo)
		return &messageVideo, err

	case MessageExpiredVideoType:
		var messageExpiredVideo MessageExpiredVideo
		err := json.Unmarshal(*rawMsg, &messageExpiredVideo)
		return &messageExpiredVideo, err

	case MessageVideoNoteType:
		var messageVideoNote MessageVideoNote
		err := json.Unmarshal(*rawMsg, &messageVideoNote)
		return &messageVideoNote, err

	case MessageVoiceNoteType:
		var messageVoiceNote MessageVoiceNote
		err := json.Unmarshal(*rawMsg, &messageVoiceNote)
		return &messageVoiceNote, err

	case MessageLocationType:
		var messageLocation MessageLocation
		err := json.Unmarshal(*rawMsg, &messageLocation)
		return &messageLocation, err

	case MessageVenueType:
		var messageVenue MessageVenue
		err := json.Unmarshal(*rawMsg, &messageVenue)
		return &messageVenue, err

	case MessageContactType:
		var messageContact MessageContact
		err := json.Unmarshal(*rawMsg, &messageContact)
		return &messageContact, err

	case MessageGameType:
		var messageGame MessageGame
		err := json.Unmarshal(*rawMsg, &messageGame)
		return &messageGame, err

	case MessageInvoiceType:
		var messageInvoice MessageInvoice
		err := json.Unmarshal(*rawMsg, &messageInvoice)
		return &messageInvoice, err

	case MessageCallType:
		var messageCall MessageCall
		err := json.Unmarshal(*rawMsg, &messageCall)
		return &messageCall, err

	case MessageBasicGroupChatCreateType:
		var messageBasicGroupChatCreate MessageBasicGroupChatCreate
		err := json.Unmarshal(*rawMsg, &messageBasicGroupChatCreate)
		return &messageBasicGroupChatCreate, err

	case MessageSupergroupChatCreateType:
		var messageSupergroupChatCreate MessageSupergroupChatCreate
		err := json.Unmarshal(*rawMsg, &messageSupergroupChatCreate)
		return &messageSupergroupChatCreate, err

	case MessageChatChangeTitleType:
		var messageChatChangeTitle MessageChatChangeTitle
		err := json.Unmarshal(*rawMsg, &messageChatChangeTitle)
		return &messageChatChangeTitle, err

	case MessageChatChangePhotoType:
		var messageChatChangePhoto MessageChatChangePhoto
		err := json.Unmarshal(*rawMsg, &messageChatChangePhoto)
		return &messageChatChangePhoto, err

	case MessageChatDeletePhotoType:
		var messageChatDeletePhoto MessageChatDeletePhoto
		err := json.Unmarshal(*rawMsg, &messageChatDeletePhoto)
		return &messageChatDeletePhoto, err

	case MessageChatAddMembersType:
		var messageChatAddMembers MessageChatAddMembers
		err := json.Unmarshal(*rawMsg, &messageChatAddMembers)
		return &messageChatAddMembers, err

	case MessageChatJoinByLinkType:
		var messageChatJoinByLink MessageChatJoinByLink
		err := json.Unmarshal(*rawMsg, &messageChatJoinByLink)
		return &messageChatJoinByLink, err

	case MessageChatDeleteMemberType:
		var messageChatDeleteMember MessageChatDeleteMember
		err := json.Unmarshal(*rawMsg, &messageChatDeleteMember)
		return &messageChatDeleteMember, err

	case MessageChatUpgradeToType:
		var messageChatUpgradeTo MessageChatUpgradeTo
		err := json.Unmarshal(*rawMsg, &messageChatUpgradeTo)
		return &messageChatUpgradeTo, err

	case MessageChatUpgradeFromType:
		var messageChatUpgradeFrom MessageChatUpgradeFrom
		err := json.Unmarshal(*rawMsg, &messageChatUpgradeFrom)
		return &messageChatUpgradeFrom, err

	case MessagePinMessageType:
		var messagePinMessage MessagePinMessage
		err := json.Unmarshal(*rawMsg, &messagePinMessage)
		return &messagePinMessage, err

	case MessageScreenshotTakenType:
		var messageScreenshotTaken MessageScreenshotTaken
		err := json.Unmarshal(*rawMsg, &messageScreenshotTaken)
		return &messageScreenshotTaken, err

	case MessageChatSetTTLType:
		var messageChatSetTTL MessageChatSetTTL
		err := json.Unmarshal(*rawMsg, &messageChatSetTTL)
		return &messageChatSetTTL, err

	case MessageCustomServiceActionType:
		var messageCustomServiceAction MessageCustomServiceAction
		err := json.Unmarshal(*rawMsg, &messageCustomServiceAction)
		return &messageCustomServiceAction, err

	case MessageGameScoreType:
		var messageGameScore MessageGameScore
		err := json.Unmarshal(*rawMsg, &messageGameScore)
		return &messageGameScore, err

	case MessagePaymentSuccessfulType:
		var messagePaymentSuccessful MessagePaymentSuccessful
		err := json.Unmarshal(*rawMsg, &messagePaymentSuccessful)
		return &messagePaymentSuccessful, err

	case MessagePaymentSuccessfulBotType:
		var messagePaymentSuccessfulBot MessagePaymentSuccessfulBot
		err := json.Unmarshal(*rawMsg, &messagePaymentSuccessfulBot)
		return &messagePaymentSuccessfulBot, err

	case MessageContactRegisteredType:
		var messageContactRegistered MessageContactRegistered
		err := json.Unmarshal(*rawMsg, &messageContactRegistered)
		return &messageContactRegistered, err

	case MessageWebsiteConnectedType:
		var messageWebsiteConnected MessageWebsiteConnected
		err := json.Unmarshal(*rawMsg, &messageWebsiteConnected)
		return &messageWebsiteConnected, err

	case MessageUnsupportedType:
		var messageUnsupported MessageUnsupported
		err := json.Unmarshal(*rawMsg, &messageUnsupported)
		return &messageUnsupported, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalTextEntityType(rawMsg *json.RawMessage) (TextEntityType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TextEntityTypeEnum(objMap["@type"].(string)) {
	case TextEntityTypeMentionType:
		var textEntityTypeMention TextEntityTypeMention
		err := json.Unmarshal(*rawMsg, &textEntityTypeMention)
		return &textEntityTypeMention, err

	case TextEntityTypeHashtagType:
		var textEntityTypeHashtag TextEntityTypeHashtag
		err := json.Unmarshal(*rawMsg, &textEntityTypeHashtag)
		return &textEntityTypeHashtag, err

	case TextEntityTypeCashtagType:
		var textEntityTypeCashtag TextEntityTypeCashtag
		err := json.Unmarshal(*rawMsg, &textEntityTypeCashtag)
		return &textEntityTypeCashtag, err

	case TextEntityTypeBotCommandType:
		var textEntityTypeBotCommand TextEntityTypeBotCommand
		err := json.Unmarshal(*rawMsg, &textEntityTypeBotCommand)
		return &textEntityTypeBotCommand, err

	case TextEntityTypeURLType:
		var textEntityTypeURL TextEntityTypeURL
		err := json.Unmarshal(*rawMsg, &textEntityTypeURL)
		return &textEntityTypeURL, err

	case TextEntityTypeEmailAddressType:
		var textEntityTypeEmailAddress TextEntityTypeEmailAddress
		err := json.Unmarshal(*rawMsg, &textEntityTypeEmailAddress)
		return &textEntityTypeEmailAddress, err

	case TextEntityTypeBoldType:
		var textEntityTypeBold TextEntityTypeBold
		err := json.Unmarshal(*rawMsg, &textEntityTypeBold)
		return &textEntityTypeBold, err

	case TextEntityTypeItalicType:
		var textEntityTypeItalic TextEntityTypeItalic
		err := json.Unmarshal(*rawMsg, &textEntityTypeItalic)
		return &textEntityTypeItalic, err

	case TextEntityTypeCodeType:
		var textEntityTypeCode TextEntityTypeCode
		err := json.Unmarshal(*rawMsg, &textEntityTypeCode)
		return &textEntityTypeCode, err

	case TextEntityTypePreType:
		var textEntityTypePre TextEntityTypePre
		err := json.Unmarshal(*rawMsg, &textEntityTypePre)
		return &textEntityTypePre, err

	case TextEntityTypePreCodeType:
		var textEntityTypePreCode TextEntityTypePreCode
		err := json.Unmarshal(*rawMsg, &textEntityTypePreCode)
		return &textEntityTypePreCode, err

	case TextEntityTypeTextURLType:
		var textEntityTypeTextURL TextEntityTypeTextURL
		err := json.Unmarshal(*rawMsg, &textEntityTypeTextURL)
		return &textEntityTypeTextURL, err

	case TextEntityTypeMentionNameType:
		var textEntityTypeMentionName TextEntityTypeMentionName
		err := json.Unmarshal(*rawMsg, &textEntityTypeMentionName)
		return &textEntityTypeMentionName, err

	case TextEntityTypePhoneNumberType:
		var textEntityTypePhoneNumber TextEntityTypePhoneNumber
		err := json.Unmarshal(*rawMsg, &textEntityTypePhoneNumber)
		return &textEntityTypePhoneNumber, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputMessageContent(rawMsg *json.RawMessage) (InputMessageContent, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputMessageContentEnum(objMap["@type"].(string)) {
	case InputMessageTextType:
		var inputMessageText InputMessageText
		err := json.Unmarshal(*rawMsg, &inputMessageText)
		return &inputMessageText, err

	case InputMessageAnimationType:
		var inputMessageAnimation InputMessageAnimation
		err := json.Unmarshal(*rawMsg, &inputMessageAnimation)
		return &inputMessageAnimation, err

	case InputMessageAudioType:
		var inputMessageAudio InputMessageAudio
		err := json.Unmarshal(*rawMsg, &inputMessageAudio)
		return &inputMessageAudio, err

	case InputMessageDocumentType:
		var inputMessageDocument InputMessageDocument
		err := json.Unmarshal(*rawMsg, &inputMessageDocument)
		return &inputMessageDocument, err

	case InputMessagePhotoType:
		var inputMessagePhoto InputMessagePhoto
		err := json.Unmarshal(*rawMsg, &inputMessagePhoto)
		return &inputMessagePhoto, err

	case InputMessageStickerType:
		var inputMessageSticker InputMessageSticker
		err := json.Unmarshal(*rawMsg, &inputMessageSticker)
		return &inputMessageSticker, err

	case InputMessageVideoType:
		var inputMessageVideo InputMessageVideo
		err := json.Unmarshal(*rawMsg, &inputMessageVideo)
		return &inputMessageVideo, err

	case InputMessageVideoNoteType:
		var inputMessageVideoNote InputMessageVideoNote
		err := json.Unmarshal(*rawMsg, &inputMessageVideoNote)
		return &inputMessageVideoNote, err

	case InputMessageVoiceNoteType:
		var inputMessageVoiceNote InputMessageVoiceNote
		err := json.Unmarshal(*rawMsg, &inputMessageVoiceNote)
		return &inputMessageVoiceNote, err

	case InputMessageLocationType:
		var inputMessageLocation InputMessageLocation
		err := json.Unmarshal(*rawMsg, &inputMessageLocation)
		return &inputMessageLocation, err

	case InputMessageVenueType:
		var inputMessageVenue InputMessageVenue
		err := json.Unmarshal(*rawMsg, &inputMessageVenue)
		return &inputMessageVenue, err

	case InputMessageContactType:
		var inputMessageContact InputMessageContact
		err := json.Unmarshal(*rawMsg, &inputMessageContact)
		return &inputMessageContact, err

	case InputMessageGameType:
		var inputMessageGame InputMessageGame
		err := json.Unmarshal(*rawMsg, &inputMessageGame)
		return &inputMessageGame, err

	case InputMessageInvoiceType:
		var inputMessageInvoice InputMessageInvoice
		err := json.Unmarshal(*rawMsg, &inputMessageInvoice)
		return &inputMessageInvoice, err

	case InputMessageForwardedType:
		var inputMessageForwarded InputMessageForwarded
		err := json.Unmarshal(*rawMsg, &inputMessageForwarded)
		return &inputMessageForwarded, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalSearchMessagesFilter(rawMsg *json.RawMessage) (SearchMessagesFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch SearchMessagesFilterEnum(objMap["@type"].(string)) {
	case SearchMessagesFilterEmptyType:
		var searchMessagesFilterEmpty SearchMessagesFilterEmpty
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterEmpty)
		return &searchMessagesFilterEmpty, err

	case SearchMessagesFilterAnimationType:
		var searchMessagesFilterAnimation SearchMessagesFilterAnimation
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterAnimation)
		return &searchMessagesFilterAnimation, err

	case SearchMessagesFilterAudioType:
		var searchMessagesFilterAudio SearchMessagesFilterAudio
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterAudio)
		return &searchMessagesFilterAudio, err

	case SearchMessagesFilterDocumentType:
		var searchMessagesFilterDocument SearchMessagesFilterDocument
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterDocument)
		return &searchMessagesFilterDocument, err

	case SearchMessagesFilterPhotoType:
		var searchMessagesFilterPhoto SearchMessagesFilterPhoto
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterPhoto)
		return &searchMessagesFilterPhoto, err

	case SearchMessagesFilterVideoType:
		var searchMessagesFilterVideo SearchMessagesFilterVideo
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVideo)
		return &searchMessagesFilterVideo, err

	case SearchMessagesFilterVoiceNoteType:
		var searchMessagesFilterVoiceNote SearchMessagesFilterVoiceNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVoiceNote)
		return &searchMessagesFilterVoiceNote, err

	case SearchMessagesFilterPhotoAndVideoType:
		var searchMessagesFilterPhotoAndVideo SearchMessagesFilterPhotoAndVideo
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterPhotoAndVideo)
		return &searchMessagesFilterPhotoAndVideo, err

	case SearchMessagesFilterURLType:
		var searchMessagesFilterURL SearchMessagesFilterURL
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterURL)
		return &searchMessagesFilterURL, err

	case SearchMessagesFilterChatPhotoType:
		var searchMessagesFilterChatPhoto SearchMessagesFilterChatPhoto
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterChatPhoto)
		return &searchMessagesFilterChatPhoto, err

	case SearchMessagesFilterCallType:
		var searchMessagesFilterCall SearchMessagesFilterCall
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterCall)
		return &searchMessagesFilterCall, err

	case SearchMessagesFilterMissedCallType:
		var searchMessagesFilterMissedCall SearchMessagesFilterMissedCall
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterMissedCall)
		return &searchMessagesFilterMissedCall, err

	case SearchMessagesFilterVideoNoteType:
		var searchMessagesFilterVideoNote SearchMessagesFilterVideoNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVideoNote)
		return &searchMessagesFilterVideoNote, err

	case SearchMessagesFilterVoiceAndVideoNoteType:
		var searchMessagesFilterVoiceAndVideoNote SearchMessagesFilterVoiceAndVideoNote
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterVoiceAndVideoNote)
		return &searchMessagesFilterVoiceAndVideoNote, err

	case SearchMessagesFilterMentionType:
		var searchMessagesFilterMention SearchMessagesFilterMention
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterMention)
		return &searchMessagesFilterMention, err

	case SearchMessagesFilterUnreadMentionType:
		var searchMessagesFilterUnreadMention SearchMessagesFilterUnreadMention
		err := json.Unmarshal(*rawMsg, &searchMessagesFilterUnreadMention)
		return &searchMessagesFilterUnreadMention, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalChatAction(rawMsg *json.RawMessage) (ChatAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatActionEnum(objMap["@type"].(string)) {
	case ChatActionTypingType:
		var chatActionTyping ChatActionTyping
		err := json.Unmarshal(*rawMsg, &chatActionTyping)
		return &chatActionTyping, err

	case ChatActionRecordingVideoType:
		var chatActionRecordingVideo ChatActionRecordingVideo
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVideo)
		return &chatActionRecordingVideo, err

	case ChatActionUploadingVideoType:
		var chatActionUploadingVideo ChatActionUploadingVideo
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVideo)
		return &chatActionUploadingVideo, err

	case ChatActionRecordingVoiceNoteType:
		var chatActionRecordingVoiceNote ChatActionRecordingVoiceNote
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVoiceNote)
		return &chatActionRecordingVoiceNote, err

	case ChatActionUploadingVoiceNoteType:
		var chatActionUploadingVoiceNote ChatActionUploadingVoiceNote
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVoiceNote)
		return &chatActionUploadingVoiceNote, err

	case ChatActionUploadingPhotoType:
		var chatActionUploadingPhoto ChatActionUploadingPhoto
		err := json.Unmarshal(*rawMsg, &chatActionUploadingPhoto)
		return &chatActionUploadingPhoto, err

	case ChatActionUploadingDocumentType:
		var chatActionUploadingDocument ChatActionUploadingDocument
		err := json.Unmarshal(*rawMsg, &chatActionUploadingDocument)
		return &chatActionUploadingDocument, err

	case ChatActionChoosingLocationType:
		var chatActionChoosingLocation ChatActionChoosingLocation
		err := json.Unmarshal(*rawMsg, &chatActionChoosingLocation)
		return &chatActionChoosingLocation, err

	case ChatActionChoosingContactType:
		var chatActionChoosingContact ChatActionChoosingContact
		err := json.Unmarshal(*rawMsg, &chatActionChoosingContact)
		return &chatActionChoosingContact, err

	case ChatActionStartPlayingGameType:
		var chatActionStartPlayingGame ChatActionStartPlayingGame
		err := json.Unmarshal(*rawMsg, &chatActionStartPlayingGame)
		return &chatActionStartPlayingGame, err

	case ChatActionRecordingVideoNoteType:
		var chatActionRecordingVideoNote ChatActionRecordingVideoNote
		err := json.Unmarshal(*rawMsg, &chatActionRecordingVideoNote)
		return &chatActionRecordingVideoNote, err

	case ChatActionUploadingVideoNoteType:
		var chatActionUploadingVideoNote ChatActionUploadingVideoNote
		err := json.Unmarshal(*rawMsg, &chatActionUploadingVideoNote)
		return &chatActionUploadingVideoNote, err

	case ChatActionCancelType:
		var chatActionCancel ChatActionCancel
		err := json.Unmarshal(*rawMsg, &chatActionCancel)
		return &chatActionCancel, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalUserStatus(rawMsg *json.RawMessage) (UserStatus, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserStatusEnum(objMap["@type"].(string)) {
	case UserStatusEmptyType:
		var userStatusEmpty UserStatusEmpty
		err := json.Unmarshal(*rawMsg, &userStatusEmpty)
		return &userStatusEmpty, err

	case UserStatusOnlineType:
		var userStatusOnline UserStatusOnline
		err := json.Unmarshal(*rawMsg, &userStatusOnline)
		return &userStatusOnline, err

	case UserStatusOfflineType:
		var userStatusOffline UserStatusOffline
		err := json.Unmarshal(*rawMsg, &userStatusOffline)
		return &userStatusOffline, err

	case UserStatusRecentlyType:
		var userStatusRecently UserStatusRecently
		err := json.Unmarshal(*rawMsg, &userStatusRecently)
		return &userStatusRecently, err

	case UserStatusLastWeekType:
		var userStatusLastWeek UserStatusLastWeek
		err := json.Unmarshal(*rawMsg, &userStatusLastWeek)
		return &userStatusLastWeek, err

	case UserStatusLastMonthType:
		var userStatusLastMonth UserStatusLastMonth
		err := json.Unmarshal(*rawMsg, &userStatusLastMonth)
		return &userStatusLastMonth, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalCallDiscardReason(rawMsg *json.RawMessage) (CallDiscardReason, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallDiscardReasonEnum(objMap["@type"].(string)) {
	case CallDiscardReasonEmptyType:
		var callDiscardReasonEmpty CallDiscardReasonEmpty
		err := json.Unmarshal(*rawMsg, &callDiscardReasonEmpty)
		return &callDiscardReasonEmpty, err

	case CallDiscardReasonMissedType:
		var callDiscardReasonMissed CallDiscardReasonMissed
		err := json.Unmarshal(*rawMsg, &callDiscardReasonMissed)
		return &callDiscardReasonMissed, err

	case CallDiscardReasonDeclinedType:
		var callDiscardReasonDeclined CallDiscardReasonDeclined
		err := json.Unmarshal(*rawMsg, &callDiscardReasonDeclined)
		return &callDiscardReasonDeclined, err

	case CallDiscardReasonDisconnectedType:
		var callDiscardReasonDisconnected CallDiscardReasonDisconnected
		err := json.Unmarshal(*rawMsg, &callDiscardReasonDisconnected)
		return &callDiscardReasonDisconnected, err

	case CallDiscardReasonHungUpType:
		var callDiscardReasonHungUp CallDiscardReasonHungUp
		err := json.Unmarshal(*rawMsg, &callDiscardReasonHungUp)
		return &callDiscardReasonHungUp, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalCallState(rawMsg *json.RawMessage) (CallState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallStateEnum(objMap["@type"].(string)) {
	case CallStatePendingType:
		var callStatePending CallStatePending
		err := json.Unmarshal(*rawMsg, &callStatePending)
		return &callStatePending, err

	case CallStateExchangingKeysType:
		var callStateExchangingKeys CallStateExchangingKeys
		err := json.Unmarshal(*rawMsg, &callStateExchangingKeys)
		return &callStateExchangingKeys, err

	case CallStateReadyType:
		var callStateReady CallStateReady
		err := json.Unmarshal(*rawMsg, &callStateReady)
		return &callStateReady, err

	case CallStateHangingUpType:
		var callStateHangingUp CallStateHangingUp
		err := json.Unmarshal(*rawMsg, &callStateHangingUp)
		return &callStateHangingUp, err

	case CallStateDiscardedType:
		var callStateDiscarded CallStateDiscarded
		err := json.Unmarshal(*rawMsg, &callStateDiscarded)
		return &callStateDiscarded, err

	case CallStateErrorType:
		var callStateError CallStateError
		err := json.Unmarshal(*rawMsg, &callStateError)
		return &callStateError, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputInlineQueryResult(rawMsg *json.RawMessage) (InputInlineQueryResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputInlineQueryResultEnum(objMap["@type"].(string)) {
	case InputInlineQueryResultAnimatedGifType:
		var inputInlineQueryResultAnimatedGif InputInlineQueryResultAnimatedGif
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultAnimatedGif)
		return &inputInlineQueryResultAnimatedGif, err

	case InputInlineQueryResultAnimatedMpeg4Type:
		var inputInlineQueryResultAnimatedMpeg4 InputInlineQueryResultAnimatedMpeg4
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultAnimatedMpeg4)
		return &inputInlineQueryResultAnimatedMpeg4, err

	case InputInlineQueryResultArticleType:
		var inputInlineQueryResultArticle InputInlineQueryResultArticle
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultArticle)
		return &inputInlineQueryResultArticle, err

	case InputInlineQueryResultAudioType:
		var inputInlineQueryResultAudio InputInlineQueryResultAudio
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultAudio)
		return &inputInlineQueryResultAudio, err

	case InputInlineQueryResultContactType:
		var inputInlineQueryResultContact InputInlineQueryResultContact
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultContact)
		return &inputInlineQueryResultContact, err

	case InputInlineQueryResultDocumentType:
		var inputInlineQueryResultDocument InputInlineQueryResultDocument
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultDocument)
		return &inputInlineQueryResultDocument, err

	case InputInlineQueryResultGameType:
		var inputInlineQueryResultGame InputInlineQueryResultGame
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultGame)
		return &inputInlineQueryResultGame, err

	case InputInlineQueryResultLocationType:
		var inputInlineQueryResultLocation InputInlineQueryResultLocation
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultLocation)
		return &inputInlineQueryResultLocation, err

	case InputInlineQueryResultPhotoType:
		var inputInlineQueryResultPhoto InputInlineQueryResultPhoto
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultPhoto)
		return &inputInlineQueryResultPhoto, err

	case InputInlineQueryResultStickerType:
		var inputInlineQueryResultSticker InputInlineQueryResultSticker
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultSticker)
		return &inputInlineQueryResultSticker, err

	case InputInlineQueryResultVenueType:
		var inputInlineQueryResultVenue InputInlineQueryResultVenue
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVenue)
		return &inputInlineQueryResultVenue, err

	case InputInlineQueryResultVideoType:
		var inputInlineQueryResultVideo InputInlineQueryResultVideo
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVideo)
		return &inputInlineQueryResultVideo, err

	case InputInlineQueryResultVoiceNoteType:
		var inputInlineQueryResultVoiceNote InputInlineQueryResultVoiceNote
		err := json.Unmarshal(*rawMsg, &inputInlineQueryResultVoiceNote)
		return &inputInlineQueryResultVoiceNote, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInlineQueryResult(rawMsg *json.RawMessage) (InlineQueryResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InlineQueryResultEnum(objMap["@type"].(string)) {
	case InlineQueryResultArticleType:
		var inlineQueryResultArticle InlineQueryResultArticle
		err := json.Unmarshal(*rawMsg, &inlineQueryResultArticle)
		return &inlineQueryResultArticle, err

	case InlineQueryResultContactType:
		var inlineQueryResultContact InlineQueryResultContact
		err := json.Unmarshal(*rawMsg, &inlineQueryResultContact)
		return &inlineQueryResultContact, err

	case InlineQueryResultLocationType:
		var inlineQueryResultLocation InlineQueryResultLocation
		err := json.Unmarshal(*rawMsg, &inlineQueryResultLocation)
		return &inlineQueryResultLocation, err

	case InlineQueryResultVenueType:
		var inlineQueryResultVenue InlineQueryResultVenue
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVenue)
		return &inlineQueryResultVenue, err

	case InlineQueryResultGameType:
		var inlineQueryResultGame InlineQueryResultGame
		err := json.Unmarshal(*rawMsg, &inlineQueryResultGame)
		return &inlineQueryResultGame, err

	case InlineQueryResultAnimationType:
		var inlineQueryResultAnimation InlineQueryResultAnimation
		err := json.Unmarshal(*rawMsg, &inlineQueryResultAnimation)
		return &inlineQueryResultAnimation, err

	case InlineQueryResultAudioType:
		var inlineQueryResultAudio InlineQueryResultAudio
		err := json.Unmarshal(*rawMsg, &inlineQueryResultAudio)
		return &inlineQueryResultAudio, err

	case InlineQueryResultDocumentType:
		var inlineQueryResultDocument InlineQueryResultDocument
		err := json.Unmarshal(*rawMsg, &inlineQueryResultDocument)
		return &inlineQueryResultDocument, err

	case InlineQueryResultPhotoType:
		var inlineQueryResultPhoto InlineQueryResultPhoto
		err := json.Unmarshal(*rawMsg, &inlineQueryResultPhoto)
		return &inlineQueryResultPhoto, err

	case InlineQueryResultStickerType:
		var inlineQueryResultSticker InlineQueryResultSticker
		err := json.Unmarshal(*rawMsg, &inlineQueryResultSticker)
		return &inlineQueryResultSticker, err

	case InlineQueryResultVideoType:
		var inlineQueryResultVideo InlineQueryResultVideo
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVideo)
		return &inlineQueryResultVideo, err

	case InlineQueryResultVoiceNoteType:
		var inlineQueryResultVoiceNote InlineQueryResultVoiceNote
		err := json.Unmarshal(*rawMsg, &inlineQueryResultVoiceNote)
		return &inlineQueryResultVoiceNote, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalCallbackQueryPayload(rawMsg *json.RawMessage) (CallbackQueryPayload, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CallbackQueryPayloadEnum(objMap["@type"].(string)) {
	case CallbackQueryPayloadDataType:
		var callbackQueryPayloadData CallbackQueryPayloadData
		err := json.Unmarshal(*rawMsg, &callbackQueryPayloadData)
		return &callbackQueryPayloadData, err

	case CallbackQueryPayloadGameType:
		var callbackQueryPayloadGame CallbackQueryPayloadGame
		err := json.Unmarshal(*rawMsg, &callbackQueryPayloadGame)
		return &callbackQueryPayloadGame, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalChatEventAction(rawMsg *json.RawMessage) (ChatEventAction, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatEventActionEnum(objMap["@type"].(string)) {
	case ChatEventMessageEditedType:
		var chatEventMessageEdited ChatEventMessageEdited
		err := json.Unmarshal(*rawMsg, &chatEventMessageEdited)
		return &chatEventMessageEdited, err

	case ChatEventMessageDeletedType:
		var chatEventMessageDeleted ChatEventMessageDeleted
		err := json.Unmarshal(*rawMsg, &chatEventMessageDeleted)
		return &chatEventMessageDeleted, err

	case ChatEventMessagePinnedType:
		var chatEventMessagePinned ChatEventMessagePinned
		err := json.Unmarshal(*rawMsg, &chatEventMessagePinned)
		return &chatEventMessagePinned, err

	case ChatEventMessageUnpinnedType:
		var chatEventMessageUnpinned ChatEventMessageUnpinned
		err := json.Unmarshal(*rawMsg, &chatEventMessageUnpinned)
		return &chatEventMessageUnpinned, err

	case ChatEventMemberJoinedType:
		var chatEventMemberJoined ChatEventMemberJoined
		err := json.Unmarshal(*rawMsg, &chatEventMemberJoined)
		return &chatEventMemberJoined, err

	case ChatEventMemberLeftType:
		var chatEventMemberLeft ChatEventMemberLeft
		err := json.Unmarshal(*rawMsg, &chatEventMemberLeft)
		return &chatEventMemberLeft, err

	case ChatEventMemberInvitedType:
		var chatEventMemberInvited ChatEventMemberInvited
		err := json.Unmarshal(*rawMsg, &chatEventMemberInvited)
		return &chatEventMemberInvited, err

	case ChatEventMemberPromotedType:
		var chatEventMemberPromoted ChatEventMemberPromoted
		err := json.Unmarshal(*rawMsg, &chatEventMemberPromoted)
		return &chatEventMemberPromoted, err

	case ChatEventMemberRestrictedType:
		var chatEventMemberRestricted ChatEventMemberRestricted
		err := json.Unmarshal(*rawMsg, &chatEventMemberRestricted)
		return &chatEventMemberRestricted, err

	case ChatEventTitleChangedType:
		var chatEventTitleChanged ChatEventTitleChanged
		err := json.Unmarshal(*rawMsg, &chatEventTitleChanged)
		return &chatEventTitleChanged, err

	case ChatEventDescriptionChangedType:
		var chatEventDescriptionChanged ChatEventDescriptionChanged
		err := json.Unmarshal(*rawMsg, &chatEventDescriptionChanged)
		return &chatEventDescriptionChanged, err

	case ChatEventUsernameChangedType:
		var chatEventUsernameChanged ChatEventUsernameChanged
		err := json.Unmarshal(*rawMsg, &chatEventUsernameChanged)
		return &chatEventUsernameChanged, err

	case ChatEventPhotoChangedType:
		var chatEventPhotoChanged ChatEventPhotoChanged
		err := json.Unmarshal(*rawMsg, &chatEventPhotoChanged)
		return &chatEventPhotoChanged, err

	case ChatEventInvitesToggledType:
		var chatEventInvitesToggled ChatEventInvitesToggled
		err := json.Unmarshal(*rawMsg, &chatEventInvitesToggled)
		return &chatEventInvitesToggled, err

	case ChatEventSignMessagesToggledType:
		var chatEventSignMessagesToggled ChatEventSignMessagesToggled
		err := json.Unmarshal(*rawMsg, &chatEventSignMessagesToggled)
		return &chatEventSignMessagesToggled, err

	case ChatEventStickerSetChangedType:
		var chatEventStickerSetChanged ChatEventStickerSetChanged
		err := json.Unmarshal(*rawMsg, &chatEventStickerSetChanged)
		return &chatEventStickerSetChanged, err

	case ChatEventIsAllHistoryAvailableToggledType:
		var chatEventIsAllHistoryAvailableToggled ChatEventIsAllHistoryAvailableToggled
		err := json.Unmarshal(*rawMsg, &chatEventIsAllHistoryAvailableToggled)
		return &chatEventIsAllHistoryAvailableToggled, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalDeviceToken(rawMsg *json.RawMessage) (DeviceToken, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch DeviceTokenEnum(objMap["@type"].(string)) {
	case DeviceTokenGoogleCloudMessagingType:
		var deviceTokenGoogleCloudMessaging DeviceTokenGoogleCloudMessaging
		err := json.Unmarshal(*rawMsg, &deviceTokenGoogleCloudMessaging)
		return &deviceTokenGoogleCloudMessaging, err

	case DeviceTokenApplePushType:
		var deviceTokenApplePush DeviceTokenApplePush
		err := json.Unmarshal(*rawMsg, &deviceTokenApplePush)
		return &deviceTokenApplePush, err

	case DeviceTokenApplePushVoIPType:
		var deviceTokenApplePushVoIP DeviceTokenApplePushVoIP
		err := json.Unmarshal(*rawMsg, &deviceTokenApplePushVoIP)
		return &deviceTokenApplePushVoIP, err

	case DeviceTokenWindowsPushType:
		var deviceTokenWindowsPush DeviceTokenWindowsPush
		err := json.Unmarshal(*rawMsg, &deviceTokenWindowsPush)
		return &deviceTokenWindowsPush, err

	case DeviceTokenMicrosoftPushType:
		var deviceTokenMicrosoftPush DeviceTokenMicrosoftPush
		err := json.Unmarshal(*rawMsg, &deviceTokenMicrosoftPush)
		return &deviceTokenMicrosoftPush, err

	case DeviceTokenMicrosoftPushVoIPType:
		var deviceTokenMicrosoftPushVoIP DeviceTokenMicrosoftPushVoIP
		err := json.Unmarshal(*rawMsg, &deviceTokenMicrosoftPushVoIP)
		return &deviceTokenMicrosoftPushVoIP, err

	case DeviceTokenWebPushType:
		var deviceTokenWebPush DeviceTokenWebPush
		err := json.Unmarshal(*rawMsg, &deviceTokenWebPush)
		return &deviceTokenWebPush, err

	case DeviceTokenSimplePushType:
		var deviceTokenSimplePush DeviceTokenSimplePush
		err := json.Unmarshal(*rawMsg, &deviceTokenSimplePush)
		return &deviceTokenSimplePush, err

	case DeviceTokenUbuntuPushType:
		var deviceTokenUbuntuPush DeviceTokenUbuntuPush
		err := json.Unmarshal(*rawMsg, &deviceTokenUbuntuPush)
		return &deviceTokenUbuntuPush, err

	case DeviceTokenBlackBerryPushType:
		var deviceTokenBlackBerryPush DeviceTokenBlackBerryPush
		err := json.Unmarshal(*rawMsg, &deviceTokenBlackBerryPush)
		return &deviceTokenBlackBerryPush, err

	case DeviceTokenTizenPushType:
		var deviceTokenTizenPush DeviceTokenTizenPush
		err := json.Unmarshal(*rawMsg, &deviceTokenTizenPush)
		return &deviceTokenTizenPush, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalCheckChatUsernameResult(rawMsg *json.RawMessage) (CheckChatUsernameResult, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch CheckChatUsernameResultEnum(objMap["@type"].(string)) {
	case CheckChatUsernameResultOkType:
		var checkChatUsernameResultOk CheckChatUsernameResultOk
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultOk)
		return &checkChatUsernameResultOk, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResultUsernameInvalid CheckChatUsernameResultUsernameInvalid
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultUsernameInvalid)
		return &checkChatUsernameResultUsernameInvalid, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResultUsernameOccupied CheckChatUsernameResultUsernameOccupied
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultUsernameOccupied)
		return &checkChatUsernameResultUsernameOccupied, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResultPublicChatsTooMuch CheckChatUsernameResultPublicChatsTooMuch
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultPublicChatsTooMuch)
		return &checkChatUsernameResultPublicChatsTooMuch, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResultPublicGroupsUnavailable CheckChatUsernameResultPublicGroupsUnavailable
		err := json.Unmarshal(*rawMsg, &checkChatUsernameResultPublicGroupsUnavailable)
		return &checkChatUsernameResultPublicGroupsUnavailable, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalOptionValue(rawMsg *json.RawMessage) (OptionValue, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch OptionValueEnum(objMap["@type"].(string)) {
	case OptionValueBooleanType:
		var optionValueBoolean OptionValueBoolean
		err := json.Unmarshal(*rawMsg, &optionValueBoolean)
		return &optionValueBoolean, err

	case OptionValueEmptyType:
		var optionValueEmpty OptionValueEmpty
		err := json.Unmarshal(*rawMsg, &optionValueEmpty)
		return &optionValueEmpty, err

	case OptionValueIntegerType:
		var optionValueInteger OptionValueInteger
		err := json.Unmarshal(*rawMsg, &optionValueInteger)
		return &optionValueInteger, err

	case OptionValueStringType:
		var optionValueString OptionValueString
		err := json.Unmarshal(*rawMsg, &optionValueString)
		return &optionValueString, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalUserPrivacySettingRule(rawMsg *json.RawMessage) (UserPrivacySettingRule, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserPrivacySettingRuleEnum(objMap["@type"].(string)) {
	case UserPrivacySettingRuleAllowAllType:
		var userPrivacySettingRuleAllowAll UserPrivacySettingRuleAllowAll
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowAll)
		return &userPrivacySettingRuleAllowAll, err

	case UserPrivacySettingRuleAllowContactsType:
		var userPrivacySettingRuleAllowContacts UserPrivacySettingRuleAllowContacts
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowContacts)
		return &userPrivacySettingRuleAllowContacts, err

	case UserPrivacySettingRuleAllowUsersType:
		var userPrivacySettingRuleAllowUsers UserPrivacySettingRuleAllowUsers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleAllowUsers)
		return &userPrivacySettingRuleAllowUsers, err

	case UserPrivacySettingRuleRestrictAllType:
		var userPrivacySettingRuleRestrictAll UserPrivacySettingRuleRestrictAll
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictAll)
		return &userPrivacySettingRuleRestrictAll, err

	case UserPrivacySettingRuleRestrictContactsType:
		var userPrivacySettingRuleRestrictContacts UserPrivacySettingRuleRestrictContacts
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictContacts)
		return &userPrivacySettingRuleRestrictContacts, err

	case UserPrivacySettingRuleRestrictUsersType:
		var userPrivacySettingRuleRestrictUsers UserPrivacySettingRuleRestrictUsers
		err := json.Unmarshal(*rawMsg, &userPrivacySettingRuleRestrictUsers)
		return &userPrivacySettingRuleRestrictUsers, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalUserPrivacySetting(rawMsg *json.RawMessage) (UserPrivacySetting, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UserPrivacySettingEnum(objMap["@type"].(string)) {
	case UserPrivacySettingShowStatusType:
		var userPrivacySettingShowStatus UserPrivacySettingShowStatus
		err := json.Unmarshal(*rawMsg, &userPrivacySettingShowStatus)
		return &userPrivacySettingShowStatus, err

	case UserPrivacySettingAllowChatInvitesType:
		var userPrivacySettingAllowChatInvites UserPrivacySettingAllowChatInvites
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowChatInvites)
		return &userPrivacySettingAllowChatInvites, err

	case UserPrivacySettingAllowCallsType:
		var userPrivacySettingAllowCalls UserPrivacySettingAllowCalls
		err := json.Unmarshal(*rawMsg, &userPrivacySettingAllowCalls)
		return &userPrivacySettingAllowCalls, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalChatReportReason(rawMsg *json.RawMessage) (ChatReportReason, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatReportReasonEnum(objMap["@type"].(string)) {
	case ChatReportReasonSpamType:
		var chatReportReasonSpam ChatReportReasonSpam
		err := json.Unmarshal(*rawMsg, &chatReportReasonSpam)
		return &chatReportReasonSpam, err

	case ChatReportReasonViolenceType:
		var chatReportReasonViolence ChatReportReasonViolence
		err := json.Unmarshal(*rawMsg, &chatReportReasonViolence)
		return &chatReportReasonViolence, err

	case ChatReportReasonPornographyType:
		var chatReportReasonPornography ChatReportReasonPornography
		err := json.Unmarshal(*rawMsg, &chatReportReasonPornography)
		return &chatReportReasonPornography, err

	case ChatReportReasonCustomType:
		var chatReportReasonCustom ChatReportReasonCustom
		err := json.Unmarshal(*rawMsg, &chatReportReasonCustom)
		return &chatReportReasonCustom, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalFileType(rawMsg *json.RawMessage) (FileType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch FileTypeEnum(objMap["@type"].(string)) {
	case FileTypeNoneType:
		var fileTypeNone FileTypeNone
		err := json.Unmarshal(*rawMsg, &fileTypeNone)
		return &fileTypeNone, err

	case FileTypeAnimationType:
		var fileTypeAnimation FileTypeAnimation
		err := json.Unmarshal(*rawMsg, &fileTypeAnimation)
		return &fileTypeAnimation, err

	case FileTypeAudioType:
		var fileTypeAudio FileTypeAudio
		err := json.Unmarshal(*rawMsg, &fileTypeAudio)
		return &fileTypeAudio, err

	case FileTypeDocumentType:
		var fileTypeDocument FileTypeDocument
		err := json.Unmarshal(*rawMsg, &fileTypeDocument)
		return &fileTypeDocument, err

	case FileTypePhotoType:
		var fileTypePhoto FileTypePhoto
		err := json.Unmarshal(*rawMsg, &fileTypePhoto)
		return &fileTypePhoto, err

	case FileTypeProfilePhotoType:
		var fileTypeProfilePhoto FileTypeProfilePhoto
		err := json.Unmarshal(*rawMsg, &fileTypeProfilePhoto)
		return &fileTypeProfilePhoto, err

	case FileTypeSecretType:
		var fileTypeSecret FileTypeSecret
		err := json.Unmarshal(*rawMsg, &fileTypeSecret)
		return &fileTypeSecret, err

	case FileTypeStickerType:
		var fileTypeSticker FileTypeSticker
		err := json.Unmarshal(*rawMsg, &fileTypeSticker)
		return &fileTypeSticker, err

	case FileTypeThumbnailType:
		var fileTypeThumbnail FileTypeThumbnail
		err := json.Unmarshal(*rawMsg, &fileTypeThumbnail)
		return &fileTypeThumbnail, err

	case FileTypeUnknownType:
		var fileTypeUnknown FileTypeUnknown
		err := json.Unmarshal(*rawMsg, &fileTypeUnknown)
		return &fileTypeUnknown, err

	case FileTypeVideoType:
		var fileTypeVideo FileTypeVideo
		err := json.Unmarshal(*rawMsg, &fileTypeVideo)
		return &fileTypeVideo, err

	case FileTypeVideoNoteType:
		var fileTypeVideoNote FileTypeVideoNote
		err := json.Unmarshal(*rawMsg, &fileTypeVideoNote)
		return &fileTypeVideoNote, err

	case FileTypeVoiceNoteType:
		var fileTypeVoiceNote FileTypeVoiceNote
		err := json.Unmarshal(*rawMsg, &fileTypeVoiceNote)
		return &fileTypeVoiceNote, err

	case FileTypeWallpaperType:
		var fileTypeWallpaper FileTypeWallpaper
		err := json.Unmarshal(*rawMsg, &fileTypeWallpaper)
		return &fileTypeWallpaper, err

	case FileTypeSecretThumbnailType:
		var fileTypeSecretThumbnail FileTypeSecretThumbnail
		err := json.Unmarshal(*rawMsg, &fileTypeSecretThumbnail)
		return &fileTypeSecretThumbnail, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalNetworkType(rawMsg *json.RawMessage) (NetworkType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NetworkTypeEnum(objMap["@type"].(string)) {
	case NetworkTypeNoneType:
		var networkTypeNone NetworkTypeNone
		err := json.Unmarshal(*rawMsg, &networkTypeNone)
		return &networkTypeNone, err

	case NetworkTypeMobileType:
		var networkTypeMobile NetworkTypeMobile
		err := json.Unmarshal(*rawMsg, &networkTypeMobile)
		return &networkTypeMobile, err

	case NetworkTypeMobileRoamingType:
		var networkTypeMobileRoaming NetworkTypeMobileRoaming
		err := json.Unmarshal(*rawMsg, &networkTypeMobileRoaming)
		return &networkTypeMobileRoaming, err

	case NetworkTypeWiFiType:
		var networkTypeWiFi NetworkTypeWiFi
		err := json.Unmarshal(*rawMsg, &networkTypeWiFi)
		return &networkTypeWiFi, err

	case NetworkTypeOtherType:
		var networkTypeOther NetworkTypeOther
		err := json.Unmarshal(*rawMsg, &networkTypeOther)
		return &networkTypeOther, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalNetworkStatisticsEntry(rawMsg *json.RawMessage) (NetworkStatisticsEntry, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch NetworkStatisticsEntryEnum(objMap["@type"].(string)) {
	case NetworkStatisticsEntryFileType:
		var networkStatisticsEntryFile NetworkStatisticsEntryFile
		err := json.Unmarshal(*rawMsg, &networkStatisticsEntryFile)
		return &networkStatisticsEntryFile, err

	case NetworkStatisticsEntryCallType:
		var networkStatisticsEntryCall NetworkStatisticsEntryCall
		err := json.Unmarshal(*rawMsg, &networkStatisticsEntryCall)
		return &networkStatisticsEntryCall, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalConnectionState(rawMsg *json.RawMessage) (ConnectionState, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ConnectionStateEnum(objMap["@type"].(string)) {
	case ConnectionStateWaitingForNetworkType:
		var connectionStateWaitingForNetwork ConnectionStateWaitingForNetwork
		err := json.Unmarshal(*rawMsg, &connectionStateWaitingForNetwork)
		return &connectionStateWaitingForNetwork, err

	case ConnectionStateConnectingToProxyType:
		var connectionStateConnectingToProxy ConnectionStateConnectingToProxy
		err := json.Unmarshal(*rawMsg, &connectionStateConnectingToProxy)
		return &connectionStateConnectingToProxy, err

	case ConnectionStateConnectingType:
		var connectionStateConnecting ConnectionStateConnecting
		err := json.Unmarshal(*rawMsg, &connectionStateConnecting)
		return &connectionStateConnecting, err

	case ConnectionStateUpdatingType:
		var connectionStateUpdating ConnectionStateUpdating
		err := json.Unmarshal(*rawMsg, &connectionStateUpdating)
		return &connectionStateUpdating, err

	case ConnectionStateReadyType:
		var connectionStateReady ConnectionStateReady
		err := json.Unmarshal(*rawMsg, &connectionStateReady)
		return &connectionStateReady, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalTopChatCategory(rawMsg *json.RawMessage) (TopChatCategory, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TopChatCategoryEnum(objMap["@type"].(string)) {
	case TopChatCategoryUsersType:
		var topChatCategoryUsers TopChatCategoryUsers
		err := json.Unmarshal(*rawMsg, &topChatCategoryUsers)
		return &topChatCategoryUsers, err

	case TopChatCategoryBotsType:
		var topChatCategoryBots TopChatCategoryBots
		err := json.Unmarshal(*rawMsg, &topChatCategoryBots)
		return &topChatCategoryBots, err

	case TopChatCategoryGroupsType:
		var topChatCategoryGroups TopChatCategoryGroups
		err := json.Unmarshal(*rawMsg, &topChatCategoryGroups)
		return &topChatCategoryGroups, err

	case TopChatCategoryChannelsType:
		var topChatCategoryChannels TopChatCategoryChannels
		err := json.Unmarshal(*rawMsg, &topChatCategoryChannels)
		return &topChatCategoryChannels, err

	case TopChatCategoryInlineBotsType:
		var topChatCategoryInlineBots TopChatCategoryInlineBots
		err := json.Unmarshal(*rawMsg, &topChatCategoryInlineBots)
		return &topChatCategoryInlineBots, err

	case TopChatCategoryCallsType:
		var topChatCategoryCalls TopChatCategoryCalls
		err := json.Unmarshal(*rawMsg, &topChatCategoryCalls)
		return &topChatCategoryCalls, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalTMeURLType(rawMsg *json.RawMessage) (TMeURLType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TMeURLTypeEnum(objMap["@type"].(string)) {
	case TMeURLTypeUserType:
		var tMeURLTypeUser TMeURLTypeUser
		err := json.Unmarshal(*rawMsg, &tMeURLTypeUser)
		return &tMeURLTypeUser, err

	case TMeURLTypeSupergroupType:
		var tMeURLTypeSupergroup TMeURLTypeSupergroup
		err := json.Unmarshal(*rawMsg, &tMeURLTypeSupergroup)
		return &tMeURLTypeSupergroup, err

	case TMeURLTypeChatInviteType:
		var tMeURLTypeChatInvite TMeURLTypeChatInvite
		err := json.Unmarshal(*rawMsg, &tMeURLTypeChatInvite)
		return &tMeURLTypeChatInvite, err

	case TMeURLTypeStickerSetType:
		var tMeURLTypeStickerSet TMeURLTypeStickerSet
		err := json.Unmarshal(*rawMsg, &tMeURLTypeStickerSet)
		return &tMeURLTypeStickerSet, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalTextParseMode(rawMsg *json.RawMessage) (TextParseMode, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch TextParseModeEnum(objMap["@type"].(string)) {
	case TextParseModeMarkdownType:
		var textParseModeMarkdown TextParseModeMarkdown
		err := json.Unmarshal(*rawMsg, &textParseModeMarkdown)
		return &textParseModeMarkdown, err

	case TextParseModeHTMLType:
		var textParseModeHTML TextParseModeHTML
		err := json.Unmarshal(*rawMsg, &textParseModeHTML)
		return &textParseModeHTML, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalProxy(rawMsg *json.RawMessage) (Proxy, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ProxyEnum(objMap["@type"].(string)) {
	case ProxyEmptyType:
		var proxyEmpty ProxyEmpty
		err := json.Unmarshal(*rawMsg, &proxyEmpty)
		return &proxyEmpty, err

	case ProxySocks5Type:
		var proxySocks5 ProxySocks5
		err := json.Unmarshal(*rawMsg, &proxySocks5)
		return &proxySocks5, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalUpdate(rawMsg *json.RawMessage) (Update, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch UpdateEnum(objMap["@type"].(string)) {
	case UpdateAuthorizationStateType:
		var updateAuthorizationState UpdateAuthorizationState
		err := json.Unmarshal(*rawMsg, &updateAuthorizationState)
		return &updateAuthorizationState, err

	case UpdateNewMessageType:
		var updateNewMessage UpdateNewMessage
		err := json.Unmarshal(*rawMsg, &updateNewMessage)
		return &updateNewMessage, err

	case UpdateMessageSendAcknowledgedType:
		var updateMessageSendAcknowledged UpdateMessageSendAcknowledged
		err := json.Unmarshal(*rawMsg, &updateMessageSendAcknowledged)
		return &updateMessageSendAcknowledged, err

	case UpdateMessageSendSucceededType:
		var updateMessageSendSucceeded UpdateMessageSendSucceeded
		err := json.Unmarshal(*rawMsg, &updateMessageSendSucceeded)
		return &updateMessageSendSucceeded, err

	case UpdateMessageSendFailedType:
		var updateMessageSendFailed UpdateMessageSendFailed
		err := json.Unmarshal(*rawMsg, &updateMessageSendFailed)
		return &updateMessageSendFailed, err

	case UpdateMessageContentType:
		var updateMessageContent UpdateMessageContent
		err := json.Unmarshal(*rawMsg, &updateMessageContent)
		return &updateMessageContent, err

	case UpdateMessageEditedType:
		var updateMessageEdited UpdateMessageEdited
		err := json.Unmarshal(*rawMsg, &updateMessageEdited)
		return &updateMessageEdited, err

	case UpdateMessageViewsType:
		var updateMessageViews UpdateMessageViews
		err := json.Unmarshal(*rawMsg, &updateMessageViews)
		return &updateMessageViews, err

	case UpdateMessageContentOpenedType:
		var updateMessageContentOpened UpdateMessageContentOpened
		err := json.Unmarshal(*rawMsg, &updateMessageContentOpened)
		return &updateMessageContentOpened, err

	case UpdateMessageMentionReadType:
		var updateMessageMentionRead UpdateMessageMentionRead
		err := json.Unmarshal(*rawMsg, &updateMessageMentionRead)
		return &updateMessageMentionRead, err

	case UpdateNewChatType:
		var updateNewChat UpdateNewChat
		err := json.Unmarshal(*rawMsg, &updateNewChat)
		return &updateNewChat, err

	case UpdateChatTitleType:
		var updateChatTitle UpdateChatTitle
		err := json.Unmarshal(*rawMsg, &updateChatTitle)
		return &updateChatTitle, err

	case UpdateChatPhotoType:
		var updateChatPhoto UpdateChatPhoto
		err := json.Unmarshal(*rawMsg, &updateChatPhoto)
		return &updateChatPhoto, err

	case UpdateChatLastMessageType:
		var updateChatLastMessage UpdateChatLastMessage
		err := json.Unmarshal(*rawMsg, &updateChatLastMessage)
		return &updateChatLastMessage, err

	case UpdateChatOrderType:
		var updateChatOrder UpdateChatOrder
		err := json.Unmarshal(*rawMsg, &updateChatOrder)
		return &updateChatOrder, err

	case UpdateChatIsPinnedType:
		var updateChatIsPinned UpdateChatIsPinned
		err := json.Unmarshal(*rawMsg, &updateChatIsPinned)
		return &updateChatIsPinned, err

	case UpdateChatReadInboxType:
		var updateChatReadInbox UpdateChatReadInbox
		err := json.Unmarshal(*rawMsg, &updateChatReadInbox)
		return &updateChatReadInbox, err

	case UpdateChatReadOutboxType:
		var updateChatReadOutbox UpdateChatReadOutbox
		err := json.Unmarshal(*rawMsg, &updateChatReadOutbox)
		return &updateChatReadOutbox, err

	case UpdateChatUnreadMentionCountType:
		var updateChatUnreadMentionCount UpdateChatUnreadMentionCount
		err := json.Unmarshal(*rawMsg, &updateChatUnreadMentionCount)
		return &updateChatUnreadMentionCount, err

	case UpdateNotificationSettingsType:
		var updateNotificationSettings UpdateNotificationSettings
		err := json.Unmarshal(*rawMsg, &updateNotificationSettings)
		return &updateNotificationSettings, err

	case UpdateChatReplyMarkupType:
		var updateChatReplyMarkup UpdateChatReplyMarkup
		err := json.Unmarshal(*rawMsg, &updateChatReplyMarkup)
		return &updateChatReplyMarkup, err

	case UpdateChatDraftMessageType:
		var updateChatDraftMessage UpdateChatDraftMessage
		err := json.Unmarshal(*rawMsg, &updateChatDraftMessage)
		return &updateChatDraftMessage, err

	case UpdateDeleteMessagesType:
		var updateDeleteMessages UpdateDeleteMessages
		err := json.Unmarshal(*rawMsg, &updateDeleteMessages)
		return &updateDeleteMessages, err

	case UpdateUserChatActionType:
		var updateUserChatAction UpdateUserChatAction
		err := json.Unmarshal(*rawMsg, &updateUserChatAction)
		return &updateUserChatAction, err

	case UpdateUserStatusType:
		var updateUserStatus UpdateUserStatus
		err := json.Unmarshal(*rawMsg, &updateUserStatus)
		return &updateUserStatus, err

	case UpdateUserType:
		var updateUser UpdateUser
		err := json.Unmarshal(*rawMsg, &updateUser)
		return &updateUser, err

	case UpdateBasicGroupType:
		var updateBasicGroup UpdateBasicGroup
		err := json.Unmarshal(*rawMsg, &updateBasicGroup)
		return &updateBasicGroup, err

	case UpdateSupergroupType:
		var updateSupergroup UpdateSupergroup
		err := json.Unmarshal(*rawMsg, &updateSupergroup)
		return &updateSupergroup, err

	case UpdateSecretChatType:
		var updateSecretChat UpdateSecretChat
		err := json.Unmarshal(*rawMsg, &updateSecretChat)
		return &updateSecretChat, err

	case UpdateUserFullInfoType:
		var updateUserFullInfo UpdateUserFullInfo
		err := json.Unmarshal(*rawMsg, &updateUserFullInfo)
		return &updateUserFullInfo, err

	case UpdateBasicGroupFullInfoType:
		var updateBasicGroupFullInfo UpdateBasicGroupFullInfo
		err := json.Unmarshal(*rawMsg, &updateBasicGroupFullInfo)
		return &updateBasicGroupFullInfo, err

	case UpdateSupergroupFullInfoType:
		var updateSupergroupFullInfo UpdateSupergroupFullInfo
		err := json.Unmarshal(*rawMsg, &updateSupergroupFullInfo)
		return &updateSupergroupFullInfo, err

	case UpdateServiceNotificationType:
		var updateServiceNotification UpdateServiceNotification
		err := json.Unmarshal(*rawMsg, &updateServiceNotification)
		return &updateServiceNotification, err

	case UpdateFileType:
		var updateFile UpdateFile
		err := json.Unmarshal(*rawMsg, &updateFile)
		return &updateFile, err

	case UpdateFileGenerationStartType:
		var updateFileGenerationStart UpdateFileGenerationStart
		err := json.Unmarshal(*rawMsg, &updateFileGenerationStart)
		return &updateFileGenerationStart, err

	case UpdateFileGenerationStopType:
		var updateFileGenerationStop UpdateFileGenerationStop
		err := json.Unmarshal(*rawMsg, &updateFileGenerationStop)
		return &updateFileGenerationStop, err

	case UpdateCallType:
		var updateCall UpdateCall
		err := json.Unmarshal(*rawMsg, &updateCall)
		return &updateCall, err

	case UpdateUserPrivacySettingRulesType:
		var updateUserPrivacySettingRules UpdateUserPrivacySettingRules
		err := json.Unmarshal(*rawMsg, &updateUserPrivacySettingRules)
		return &updateUserPrivacySettingRules, err

	case UpdateUnreadMessageCountType:
		var updateUnreadMessageCount UpdateUnreadMessageCount
		err := json.Unmarshal(*rawMsg, &updateUnreadMessageCount)
		return &updateUnreadMessageCount, err

	case UpdateOptionType:
		var updateOption UpdateOption
		err := json.Unmarshal(*rawMsg, &updateOption)
		return &updateOption, err

	case UpdateInstalledStickerSetsType:
		var updateInstalledStickerSets UpdateInstalledStickerSets
		err := json.Unmarshal(*rawMsg, &updateInstalledStickerSets)
		return &updateInstalledStickerSets, err

	case UpdateTrendingStickerSetsType:
		var updateTrendingStickerSets UpdateTrendingStickerSets
		err := json.Unmarshal(*rawMsg, &updateTrendingStickerSets)
		return &updateTrendingStickerSets, err

	case UpdateRecentStickersType:
		var updateRecentStickers UpdateRecentStickers
		err := json.Unmarshal(*rawMsg, &updateRecentStickers)
		return &updateRecentStickers, err

	case UpdateFavoriteStickersType:
		var updateFavoriteStickers UpdateFavoriteStickers
		err := json.Unmarshal(*rawMsg, &updateFavoriteStickers)
		return &updateFavoriteStickers, err

	case UpdateSavedAnimationsType:
		var updateSavedAnimations UpdateSavedAnimations
		err := json.Unmarshal(*rawMsg, &updateSavedAnimations)
		return &updateSavedAnimations, err

	case UpdateConnectionStateType:
		var updateConnectionState UpdateConnectionState
		err := json.Unmarshal(*rawMsg, &updateConnectionState)
		return &updateConnectionState, err

	case UpdateNewInlineQueryType:
		var updateNewInlineQuery UpdateNewInlineQuery
		err := json.Unmarshal(*rawMsg, &updateNewInlineQuery)
		return &updateNewInlineQuery, err

	case UpdateNewChosenInlineResultType:
		var updateNewChosenInlineResult UpdateNewChosenInlineResult
		err := json.Unmarshal(*rawMsg, &updateNewChosenInlineResult)
		return &updateNewChosenInlineResult, err

	case UpdateNewCallbackQueryType:
		var updateNewCallbackQuery UpdateNewCallbackQuery
		err := json.Unmarshal(*rawMsg, &updateNewCallbackQuery)
		return &updateNewCallbackQuery, err

	case UpdateNewInlineCallbackQueryType:
		var updateNewInlineCallbackQuery UpdateNewInlineCallbackQuery
		err := json.Unmarshal(*rawMsg, &updateNewInlineCallbackQuery)
		return &updateNewInlineCallbackQuery, err

	case UpdateNewShippingQueryType:
		var updateNewShippingQuery UpdateNewShippingQuery
		err := json.Unmarshal(*rawMsg, &updateNewShippingQuery)
		return &updateNewShippingQuery, err

	case UpdateNewPreCheckoutQueryType:
		var updateNewPreCheckoutQuery UpdateNewPreCheckoutQuery
		err := json.Unmarshal(*rawMsg, &updateNewPreCheckoutQuery)
		return &updateNewPreCheckoutQuery, err

	case UpdateNewCustomEventType:
		var updateNewCustomEvent UpdateNewCustomEvent
		err := json.Unmarshal(*rawMsg, &updateNewCustomEvent)
		return &updateNewCustomEvent, err

	case UpdateNewCustomQueryType:
		var updateNewCustomQuery UpdateNewCustomQuery
		err := json.Unmarshal(*rawMsg, &updateNewCustomQuery)
		return &updateNewCustomQuery, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}
