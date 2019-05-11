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
	intStr := strconv.FormatInt(int64(*jsonInt), 10)
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

// ChatMembersFilterEnum Alias for abstract ChatMembersFilter 'Sub-Classes', used as constant-enum here
type ChatMembersFilterEnum string

// ChatMembersFilter enums
const (
	ChatMembersFilterAdministratorsType ChatMembersFilterEnum = "chatMembersFilterAdministrators"
	ChatMembersFilterMembersType        ChatMembersFilterEnum = "chatMembersFilterMembers"
	ChatMembersFilterRestrictedType     ChatMembersFilterEnum = "chatMembersFilterRestricted"
	ChatMembersFilterBannedType         ChatMembersFilterEnum = "chatMembersFilterBanned"
	ChatMembersFilterBotsType           ChatMembersFilterEnum = "chatMembersFilterBots"
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
	NotificationSettingsScopePrivateChatsType NotificationSettingsScopeEnum = "notificationSettingsScopePrivateChats"
	NotificationSettingsScopeGroupChatsType   NotificationSettingsScopeEnum = "notificationSettingsScopeGroupChats"
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

// PassportElementTypeEnum Alias for abstract PassportElementType 'Sub-Classes', used as constant-enum here
type PassportElementTypeEnum string

// PassportElementType enums
const (
	PassportElementTypePersonalDetailsType       PassportElementTypeEnum = "passportElementTypePersonalDetails"
	PassportElementTypePassportType              PassportElementTypeEnum = "passportElementTypePassport"
	PassportElementTypeDriverLicenseType         PassportElementTypeEnum = "passportElementTypeDriverLicense"
	PassportElementTypeIDentityCardType          PassportElementTypeEnum = "passportElementTypeIDentityCard"
	PassportElementTypeInternalPassportType      PassportElementTypeEnum = "passportElementTypeInternalPassport"
	PassportElementTypeAddressType               PassportElementTypeEnum = "passportElementTypeAddress"
	PassportElementTypeUtilityBillType           PassportElementTypeEnum = "passportElementTypeUtilityBill"
	PassportElementTypeBankStatementType         PassportElementTypeEnum = "passportElementTypeBankStatement"
	PassportElementTypeRentalAgreementType       PassportElementTypeEnum = "passportElementTypeRentalAgreement"
	PassportElementTypePassportRegistrationType  PassportElementTypeEnum = "passportElementTypePassportRegistration"
	PassportElementTypeTemporaryRegistrationType PassportElementTypeEnum = "passportElementTypeTemporaryRegistration"
	PassportElementTypePhoneNumberType           PassportElementTypeEnum = "passportElementTypePhoneNumber"
	PassportElementTypeEmailAddressType          PassportElementTypeEnum = "passportElementTypeEmailAddress"
)

// PassportElementEnum Alias for abstract PassportElement 'Sub-Classes', used as constant-enum here
type PassportElementEnum string

// PassportElement enums
const (
	PassportElementPersonalDetailsType       PassportElementEnum = "passportElementPersonalDetails"
	PassportElementPassportType              PassportElementEnum = "passportElementPassport"
	PassportElementDriverLicenseType         PassportElementEnum = "passportElementDriverLicense"
	PassportElementIDentityCardType          PassportElementEnum = "passportElementIDentityCard"
	PassportElementInternalPassportType      PassportElementEnum = "passportElementInternalPassport"
	PassportElementAddressType               PassportElementEnum = "passportElementAddress"
	PassportElementUtilityBillType           PassportElementEnum = "passportElementUtilityBill"
	PassportElementBankStatementType         PassportElementEnum = "passportElementBankStatement"
	PassportElementRentalAgreementType       PassportElementEnum = "passportElementRentalAgreement"
	PassportElementPassportRegistrationType  PassportElementEnum = "passportElementPassportRegistration"
	PassportElementTemporaryRegistrationType PassportElementEnum = "passportElementTemporaryRegistration"
	PassportElementPhoneNumberType           PassportElementEnum = "passportElementPhoneNumber"
	PassportElementEmailAddressType          PassportElementEnum = "passportElementEmailAddress"
)

// InputPassportElementEnum Alias for abstract InputPassportElement 'Sub-Classes', used as constant-enum here
type InputPassportElementEnum string

// InputPassportElement enums
const (
	InputPassportElementPersonalDetailsType       InputPassportElementEnum = "inputPassportElementPersonalDetails"
	InputPassportElementPassportType              InputPassportElementEnum = "inputPassportElementPassport"
	InputPassportElementDriverLicenseType         InputPassportElementEnum = "inputPassportElementDriverLicense"
	InputPassportElementIDentityCardType          InputPassportElementEnum = "inputPassportElementIDentityCard"
	InputPassportElementInternalPassportType      InputPassportElementEnum = "inputPassportElementInternalPassport"
	InputPassportElementAddressType               InputPassportElementEnum = "inputPassportElementAddress"
	InputPassportElementUtilityBillType           InputPassportElementEnum = "inputPassportElementUtilityBill"
	InputPassportElementBankStatementType         InputPassportElementEnum = "inputPassportElementBankStatement"
	InputPassportElementRentalAgreementType       InputPassportElementEnum = "inputPassportElementRentalAgreement"
	InputPassportElementPassportRegistrationType  InputPassportElementEnum = "inputPassportElementPassportRegistration"
	InputPassportElementTemporaryRegistrationType InputPassportElementEnum = "inputPassportElementTemporaryRegistration"
	InputPassportElementPhoneNumberType           InputPassportElementEnum = "inputPassportElementPhoneNumber"
	InputPassportElementEmailAddressType          InputPassportElementEnum = "inputPassportElementEmailAddress"
)

// PassportElementErrorSourceEnum Alias for abstract PassportElementErrorSource 'Sub-Classes', used as constant-enum here
type PassportElementErrorSourceEnum string

// PassportElementErrorSource enums
const (
	PassportElementErrorSourceUnspecifiedType      PassportElementErrorSourceEnum = "passportElementErrorSourceUnspecified"
	PassportElementErrorSourceDataFieldType        PassportElementErrorSourceEnum = "passportElementErrorSourceDataField"
	PassportElementErrorSourceFrontSideType        PassportElementErrorSourceEnum = "passportElementErrorSourceFrontSide"
	PassportElementErrorSourceReverseSideType      PassportElementErrorSourceEnum = "passportElementErrorSourceReverseSide"
	PassportElementErrorSourceSelfieType           PassportElementErrorSourceEnum = "passportElementErrorSourceSelfie"
	PassportElementErrorSourceTranslationFileType  PassportElementErrorSourceEnum = "passportElementErrorSourceTranslationFile"
	PassportElementErrorSourceTranslationFilesType PassportElementErrorSourceEnum = "passportElementErrorSourceTranslationFiles"
	PassportElementErrorSourceFileType             PassportElementErrorSourceEnum = "passportElementErrorSourceFile"
	PassportElementErrorSourceFilesType            PassportElementErrorSourceEnum = "passportElementErrorSourceFiles"
)

// InputPassportElementErrorSourceEnum Alias for abstract InputPassportElementErrorSource 'Sub-Classes', used as constant-enum here
type InputPassportElementErrorSourceEnum string

// InputPassportElementErrorSource enums
const (
	InputPassportElementErrorSourceUnspecifiedType      InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceUnspecified"
	InputPassportElementErrorSourceDataFieldType        InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceDataField"
	InputPassportElementErrorSourceFrontSideType        InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFrontSide"
	InputPassportElementErrorSourceReverseSideType      InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceReverseSide"
	InputPassportElementErrorSourceSelfieType           InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceSelfie"
	InputPassportElementErrorSourceTranslationFileType  InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceTranslationFile"
	InputPassportElementErrorSourceTranslationFilesType InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceTranslationFiles"
	InputPassportElementErrorSourceFileType             InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFile"
	InputPassportElementErrorSourceFilesType            InputPassportElementErrorSourceEnum = "inputPassportElementErrorSourceFiles"
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
	MessagePassportDataSentType     MessageContentEnum = "messagePassportDataSent"
	MessagePassportDataReceivedType MessageContentEnum = "messagePassportDataReceived"
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
	TextEntityTypeTextURLType      TextEntityTypeEnum = "textEntityTypeTextUrl"
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

// LanguagePackStringValueEnum Alias for abstract LanguagePackStringValue 'Sub-Classes', used as constant-enum here
type LanguagePackStringValueEnum string

// LanguagePackStringValue enums
const (
	LanguagePackStringValueOrdinaryType   LanguagePackStringValueEnum = "languagePackStringValueOrdinary"
	LanguagePackStringValuePluralizedType LanguagePackStringValueEnum = "languagePackStringValuePluralized"
	LanguagePackStringValueDeletedType    LanguagePackStringValueEnum = "languagePackStringValueDeleted"
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
	ChatReportReasonCopyrightType   ChatReportReasonEnum = "chatReportReasonCopyright"
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
	FileTypeSecretThumbnailType FileTypeEnum = "fileTypeSecretThumbnail"
	FileTypeSecureType          FileTypeEnum = "fileTypeSecure"
	FileTypeStickerType         FileTypeEnum = "fileTypeSticker"
	FileTypeThumbnailType       FileTypeEnum = "fileTypeThumbnail"
	FileTypeUnknownType         FileTypeEnum = "fileTypeUnknown"
	FileTypeVideoType           FileTypeEnum = "fileTypeVideo"
	FileTypeVideoNoteType       FileTypeEnum = "fileTypeVideoNote"
	FileTypeVoiceNoteType       FileTypeEnum = "fileTypeVoiceNote"
	FileTypeWallpaperType       FileTypeEnum = "fileTypeWallpaper"
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

// ProxyTypeEnum Alias for abstract ProxyType 'Sub-Classes', used as constant-enum here
type ProxyTypeEnum string

// ProxyType enums
const (
	ProxyTypeSocks5Type  ProxyTypeEnum = "proxyTypeSocks5"
	ProxyTypeHttpType    ProxyTypeEnum = "proxyTypeHttp"
	ProxyTypeMtprotoType ProxyTypeEnum = "proxyTypeMtproto"
)

// UpdateEnum Alias for abstract Update 'Sub-Classes', used as constant-enum here
type UpdateEnum string

// Update enums
const (
	UpdateAuthorizationStateType             UpdateEnum = "updateAuthorizationState"
	UpdateNewMessageType                     UpdateEnum = "updateNewMessage"
	UpdateMessageSendAcknowledgedType        UpdateEnum = "updateMessageSendAcknowledged"
	UpdateMessageSendSucceededType           UpdateEnum = "updateMessageSendSucceeded"
	UpdateMessageSendFailedType              UpdateEnum = "updateMessageSendFailed"
	UpdateMessageContentType                 UpdateEnum = "updateMessageContent"
	UpdateMessageEditedType                  UpdateEnum = "updateMessageEdited"
	UpdateMessageViewsType                   UpdateEnum = "updateMessageViews"
	UpdateMessageContentOpenedType           UpdateEnum = "updateMessageContentOpened"
	UpdateMessageMentionReadType             UpdateEnum = "updateMessageMentionRead"
	UpdateNewChatType                        UpdateEnum = "updateNewChat"
	UpdateChatTitleType                      UpdateEnum = "updateChatTitle"
	UpdateChatPhotoType                      UpdateEnum = "updateChatPhoto"
	UpdateChatLastMessageType                UpdateEnum = "updateChatLastMessage"
	UpdateChatOrderType                      UpdateEnum = "updateChatOrder"
	UpdateChatIsPinnedType                   UpdateEnum = "updateChatIsPinned"
	UpdateChatIsMarkedAsUnreadType           UpdateEnum = "updateChatIsMarkedAsUnread"
	UpdateChatIsSponsoredType                UpdateEnum = "updateChatIsSponsored"
	UpdateChatDefaultDisableNotificationType UpdateEnum = "updateChatDefaultDisableNotification"
	UpdateChatReadInboxType                  UpdateEnum = "updateChatReadInbox"
	UpdateChatReadOutboxType                 UpdateEnum = "updateChatReadOutbox"
	UpdateChatUnreadMentionCountType         UpdateEnum = "updateChatUnreadMentionCount"
	UpdateChatNotificationSettingsType       UpdateEnum = "updateChatNotificationSettings"
	UpdateScopeNotificationSettingsType      UpdateEnum = "updateScopeNotificationSettings"
	UpdateChatReplyMarkupType                UpdateEnum = "updateChatReplyMarkup"
	UpdateChatDraftMessageType               UpdateEnum = "updateChatDraftMessage"
	UpdateDeleteMessagesType                 UpdateEnum = "updateDeleteMessages"
	UpdateUserChatActionType                 UpdateEnum = "updateUserChatAction"
	UpdateUserStatusType                     UpdateEnum = "updateUserStatus"
	UpdateUserType                           UpdateEnum = "updateUser"
	UpdateBasicGroupType                     UpdateEnum = "updateBasicGroup"
	UpdateSupergroupType                     UpdateEnum = "updateSupergroup"
	UpdateSecretChatType                     UpdateEnum = "updateSecretChat"
	UpdateUserFullInfoType                   UpdateEnum = "updateUserFullInfo"
	UpdateBasicGroupFullInfoType             UpdateEnum = "updateBasicGroupFullInfo"
	UpdateSupergroupFullInfoType             UpdateEnum = "updateSupergroupFullInfo"
	UpdateServiceNotificationType            UpdateEnum = "updateServiceNotification"
	UpdateFileType                           UpdateEnum = "updateFile"
	UpdateFileGenerationStartType            UpdateEnum = "updateFileGenerationStart"
	UpdateFileGenerationStopType             UpdateEnum = "updateFileGenerationStop"
	UpdateCallType                           UpdateEnum = "updateCall"
	UpdateUserPrivacySettingRulesType        UpdateEnum = "updateUserPrivacySettingRules"
	UpdateUnreadMessageCountType             UpdateEnum = "updateUnreadMessageCount"
	UpdateUnreadChatCountType                UpdateEnum = "updateUnreadChatCount"
	UpdateOptionType                         UpdateEnum = "updateOption"
	UpdateInstalledStickerSetsType           UpdateEnum = "updateInstalledStickerSets"
	UpdateTrendingStickerSetsType            UpdateEnum = "updateTrendingStickerSets"
	UpdateRecentStickersType                 UpdateEnum = "updateRecentStickers"
	UpdateFavoriteStickersType               UpdateEnum = "updateFavoriteStickers"
	UpdateSavedAnimationsType                UpdateEnum = "updateSavedAnimations"
	UpdateLanguagePackStringsType            UpdateEnum = "updateLanguagePackStrings"
	UpdateConnectionStateType                UpdateEnum = "updateConnectionState"
	UpdateTermsOfServiceType                 UpdateEnum = "updateTermsOfService"
	UpdateNewInlineQueryType                 UpdateEnum = "updateNewInlineQuery"
	UpdateNewChosenInlineResultType          UpdateEnum = "updateNewChosenInlineResult"
	UpdateNewCallbackQueryType               UpdateEnum = "updateNewCallbackQuery"
	UpdateNewInlineCallbackQueryType         UpdateEnum = "updateNewInlineCallbackQuery"
	UpdateNewShippingQueryType               UpdateEnum = "updateNewShippingQuery"
	UpdateNewPreCheckoutQueryType            UpdateEnum = "updateNewPreCheckoutQuery"
	UpdateNewCustomEventType                 UpdateEnum = "updateNewCustomEvent"
	UpdateNewCustomQueryType                 UpdateEnum = "updateNewCustomQuery"
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

// ChatMembersFilter Specifies the kind of chat members to return in searchChatMembers
type ChatMembersFilter interface {
	GetChatMembersFilterEnum() ChatMembersFilterEnum
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

// NotificationSettingsScope Describes the types of chats to which notification settings are applied
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

// PassportElementType Contains the type of a Telegram Passport element
type PassportElementType interface {
	GetPassportElementTypeEnum() PassportElementTypeEnum
}

// PassportElement Contains information about a Telegram Passport element
type PassportElement interface {
	GetPassportElementEnum() PassportElementEnum
}

// InputPassportElement Contains information about a Telegram Passport element to be saved
type InputPassportElement interface {
	GetInputPassportElementEnum() InputPassportElementEnum
}

// PassportElementErrorSource Contains the description of an error in a Telegram Passport element
type PassportElementErrorSource interface {
	GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum
}

// InputPassportElementErrorSource Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementErrorSource interface {
	GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum
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

// LanguagePackStringValue Represents the value of a string in a language pack
type LanguagePackStringValue interface {
	GetLanguagePackStringValueEnum() LanguagePackStringValueEnum
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

// ProxyType Describes the type of the proxy server
type ProxyType interface {
	GetProxyTypeEnum() ProxyTypeEnum
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

// NewError creates a new Error
//
// @param code Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
// @param message Error message; subject to future changes
func NewError(code int32, message string) *Error {
	errorTemp := Error{
		tdCommon: tdCommon{Type: "error"},
		Code:     code,
		Message:  message,
	}

	return &errorTemp
}

// Ok An object of this type is returned on a successful function call for certain functions
type Ok struct {
	tdCommon
}

// MessageType return the string telegram-type of Ok
func (ok *Ok) MessageType() string {
	return "ok"
}

// NewOk creates a new Ok
//
func NewOk() *Ok {
	okTemp := Ok{
		tdCommon: tdCommon{Type: "ok"},
	}

	return &okTemp
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

// NewTdlibParameters creates a new TdlibParameters
//
// @param useTestDc If set to true, the Telegram test environment will be used instead of the production environment
// @param databaseDirectory The path to the directory for the persistent database; if empty, the current working directory will be used
// @param filesDirectory The path to the directory for storing files; if empty, database_directory will be used
// @param useFileDatabase If set to true, information about downloaded and uploaded files will be saved between application restarts
// @param useChatInfoDatabase If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database
// @param useMessageDatabase If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database
// @param useSecretChats If set to true, support for secret chats will be enabled
// @param aPIID Application identifier for Telegram API access, which can be obtained at https://my.telegram.org
// @param aPIHash Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org
// @param systemLanguageCode IETF language tag of the user's operating system language; must be non-empty
// @param deviceModel Model of the device the application is being run on; must be non-empty
// @param systemVersion Version of the operating system the application is being run on; must be non-empty
// @param applicationVersion Application version; must be non-empty
// @param enableStorageOptimizer If set to true, old files will automatically be deleted
// @param ignoreFileNames If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name
func NewTdlibParameters(useTestDc bool, databaseDirectory string, filesDirectory string, useFileDatabase bool, useChatInfoDatabase bool, useMessageDatabase bool, useSecretChats bool, aPIID int32, aPIHash string, systemLanguageCode string, deviceModel string, systemVersion string, applicationVersion string, enableStorageOptimizer bool, ignoreFileNames bool) *TdlibParameters {
	tdlibParametersTemp := TdlibParameters{
		tdCommon:               tdCommon{Type: "tdlibParameters"},
		UseTestDc:              useTestDc,
		DatabaseDirectory:      databaseDirectory,
		FilesDirectory:         filesDirectory,
		UseFileDatabase:        useFileDatabase,
		UseChatInfoDatabase:    useChatInfoDatabase,
		UseMessageDatabase:     useMessageDatabase,
		UseSecretChats:         useSecretChats,
		APIID:                  aPIID,
		APIHash:                aPIHash,
		SystemLanguageCode:     systemLanguageCode,
		DeviceModel:            deviceModel,
		SystemVersion:          systemVersion,
		ApplicationVersion:     applicationVersion,
		EnableStorageOptimizer: enableStorageOptimizer,
		IgnoreFileNames:        ignoreFileNames,
	}

	return &tdlibParametersTemp
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

// NewAuthenticationCodeTypeTelegramMessage creates a new AuthenticationCodeTypeTelegramMessage
//
// @param length Length of the code
func NewAuthenticationCodeTypeTelegramMessage(length int32) *AuthenticationCodeTypeTelegramMessage {
	authenticationCodeTypeTelegramMessageTemp := AuthenticationCodeTypeTelegramMessage{
		tdCommon: tdCommon{Type: "authenticationCodeTypeTelegramMessage"},
		Length:   length,
	}

	return &authenticationCodeTypeTelegramMessageTemp
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

// NewAuthenticationCodeTypeSms creates a new AuthenticationCodeTypeSms
//
// @param length Length of the code
func NewAuthenticationCodeTypeSms(length int32) *AuthenticationCodeTypeSms {
	authenticationCodeTypeSmsTemp := AuthenticationCodeTypeSms{
		tdCommon: tdCommon{Type: "authenticationCodeTypeSms"},
		Length:   length,
	}

	return &authenticationCodeTypeSmsTemp
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

// NewAuthenticationCodeTypeCall creates a new AuthenticationCodeTypeCall
//
// @param length Length of the code
func NewAuthenticationCodeTypeCall(length int32) *AuthenticationCodeTypeCall {
	authenticationCodeTypeCallTemp := AuthenticationCodeTypeCall{
		tdCommon: tdCommon{Type: "authenticationCodeTypeCall"},
		Length:   length,
	}

	return &authenticationCodeTypeCallTemp
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

// NewAuthenticationCodeTypeFlashCall creates a new AuthenticationCodeTypeFlashCall
//
// @param pattern Pattern of the phone number from which the call will be made
func NewAuthenticationCodeTypeFlashCall(pattern string) *AuthenticationCodeTypeFlashCall {
	authenticationCodeTypeFlashCallTemp := AuthenticationCodeTypeFlashCall{
		tdCommon: tdCommon{Type: "authenticationCodeTypeFlashCall"},
		Pattern:  pattern,
	}

	return &authenticationCodeTypeFlashCallTemp
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

// NewAuthenticationCodeInfo creates a new AuthenticationCodeInfo
//
// @param phoneNumber A phone number that is being authenticated
// @param typeParam Describes the way the code was sent to the user
// @param nextType Describes the way the next code will be sent to the user; may be null
// @param timeout Timeout before the code should be re-sent, in seconds
func NewAuthenticationCodeInfo(phoneNumber string, typeParam AuthenticationCodeType, nextType AuthenticationCodeType, timeout int32) *AuthenticationCodeInfo {
	authenticationCodeInfoTemp := AuthenticationCodeInfo{
		tdCommon:    tdCommon{Type: "authenticationCodeInfo"},
		PhoneNumber: phoneNumber,
		Type:        typeParam,
		NextType:    nextType,
		Timeout:     timeout,
	}

	return &authenticationCodeInfoTemp
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

// EmailAddressAuthenticationCodeInfo Information about the email address authentication code that was sent
type EmailAddressAuthenticationCodeInfo struct {
	tdCommon
	EmailAddressPattern string `json:"email_address_pattern"` // Pattern of the email address to which an authentication code was sent
	Length              int32  `json:"length"`                // Length of the code; 0 if unknown
}

// MessageType return the string telegram-type of EmailAddressAuthenticationCodeInfo
func (emailAddressAuthenticationCodeInfo *EmailAddressAuthenticationCodeInfo) MessageType() string {
	return "emailAddressAuthenticationCodeInfo"
}

// NewEmailAddressAuthenticationCodeInfo creates a new EmailAddressAuthenticationCodeInfo
//
// @param emailAddressPattern Pattern of the email address to which an authentication code was sent
// @param length Length of the code; 0 if unknown
func NewEmailAddressAuthenticationCodeInfo(emailAddressPattern string, length int32) *EmailAddressAuthenticationCodeInfo {
	emailAddressAuthenticationCodeInfoTemp := EmailAddressAuthenticationCodeInfo{
		tdCommon:            tdCommon{Type: "emailAddressAuthenticationCodeInfo"},
		EmailAddressPattern: emailAddressPattern,
		Length:              length,
	}

	return &emailAddressAuthenticationCodeInfoTemp
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

// NewTextEntity creates a new TextEntity
//
// @param offset Offset of the entity in UTF-16 code points
// @param length Length of the entity, in UTF-16 code points
// @param typeParam Type of the entity
func NewTextEntity(offset int32, length int32, typeParam TextEntityType) *TextEntity {
	textEntityTemp := TextEntity{
		tdCommon: tdCommon{Type: "textEntity"},
		Offset:   offset,
		Length:   length,
		Type:     typeParam,
	}

	return &textEntityTemp
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

// NewTextEntities creates a new TextEntities
//
// @param entities List of text entities
func NewTextEntities(entities []TextEntity) *TextEntities {
	textEntitiesTemp := TextEntities{
		tdCommon: tdCommon{Type: "textEntities"},
		Entities: entities,
	}

	return &textEntitiesTemp
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

// NewFormattedText creates a new FormattedText
//
// @param text The text
// @param entities Entities contained in the text
func NewFormattedText(text string, entities []TextEntity) *FormattedText {
	formattedTextTemp := FormattedText{
		tdCommon: tdCommon{Type: "formattedText"},
		Text:     text,
		Entities: entities,
	}

	return &formattedTextTemp
}

// TermsOfService Contains Telegram terms of service
type TermsOfService struct {
	tdCommon
	Text       *FormattedText `json:"text"`         // Text of the terms of service
	MinUserAge int32          `json:"min_user_age"` // Mininum age of a user to be able to accept the terms; 0 if any
	ShowPopup  bool           `json:"show_popup"`   // True, if a blocking popup with terms of service must be shown to the user
}

// MessageType return the string telegram-type of TermsOfService
func (termsOfService *TermsOfService) MessageType() string {
	return "termsOfService"
}

// NewTermsOfService creates a new TermsOfService
//
// @param text Text of the terms of service
// @param minUserAge Mininum age of a user to be able to accept the terms; 0 if any
// @param showPopup True, if a blocking popup with terms of service must be shown to the user
func NewTermsOfService(text *FormattedText, minUserAge int32, showPopup bool) *TermsOfService {
	termsOfServiceTemp := TermsOfService{
		tdCommon:   tdCommon{Type: "termsOfService"},
		Text:       text,
		MinUserAge: minUserAge,
		ShowPopup:  showPopup,
	}

	return &termsOfServiceTemp
}

// AuthorizationStateWaitTdlibParameters TDLib needs TdlibParameters for initialization
type AuthorizationStateWaitTdlibParameters struct {
	tdCommon
}

// MessageType return the string telegram-type of AuthorizationStateWaitTdlibParameters
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) MessageType() string {
	return "authorizationStateWaitTdlibParameters"
}

// NewAuthorizationStateWaitTdlibParameters creates a new AuthorizationStateWaitTdlibParameters
//
func NewAuthorizationStateWaitTdlibParameters() *AuthorizationStateWaitTdlibParameters {
	authorizationStateWaitTdlibParametersTemp := AuthorizationStateWaitTdlibParameters{
		tdCommon: tdCommon{Type: "authorizationStateWaitTdlibParameters"},
	}

	return &authorizationStateWaitTdlibParametersTemp
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

// NewAuthorizationStateWaitEncryptionKey creates a new AuthorizationStateWaitEncryptionKey
//
// @param isEncrypted True, if the database is currently encrypted
func NewAuthorizationStateWaitEncryptionKey(isEncrypted bool) *AuthorizationStateWaitEncryptionKey {
	authorizationStateWaitEncryptionKeyTemp := AuthorizationStateWaitEncryptionKey{
		tdCommon:    tdCommon{Type: "authorizationStateWaitEncryptionKey"},
		IsEncrypted: isEncrypted,
	}

	return &authorizationStateWaitEncryptionKeyTemp
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

// NewAuthorizationStateWaitPhoneNumber creates a new AuthorizationStateWaitPhoneNumber
//
func NewAuthorizationStateWaitPhoneNumber() *AuthorizationStateWaitPhoneNumber {
	authorizationStateWaitPhoneNumberTemp := AuthorizationStateWaitPhoneNumber{
		tdCommon: tdCommon{Type: "authorizationStateWaitPhoneNumber"},
	}

	return &authorizationStateWaitPhoneNumberTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateWaitPhoneNumberType
}

// AuthorizationStateWaitCode TDLib needs the user's authentication code to finalize authorization
type AuthorizationStateWaitCode struct {
	tdCommon
	IsRegistered   bool                    `json:"is_registered"`    // True, if the user is already registered
	TermsOfService *TermsOfService         `json:"terms_of_service"` // Telegram terms of service, which should be accepted before user can continue registration; may be null
	CodeInfo       *AuthenticationCodeInfo `json:"code_info"`        // Information about the authorization code that was sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitCode
func (authorizationStateWaitCode *AuthorizationStateWaitCode) MessageType() string {
	return "authorizationStateWaitCode"
}

// NewAuthorizationStateWaitCode creates a new AuthorizationStateWaitCode
//
// @param isRegistered True, if the user is already registered
// @param termsOfService Telegram terms of service, which should be accepted before user can continue registration; may be null
// @param codeInfo Information about the authorization code that was sent
func NewAuthorizationStateWaitCode(isRegistered bool, termsOfService *TermsOfService, codeInfo *AuthenticationCodeInfo) *AuthorizationStateWaitCode {
	authorizationStateWaitCodeTemp := AuthorizationStateWaitCode{
		tdCommon:       tdCommon{Type: "authorizationStateWaitCode"},
		IsRegistered:   isRegistered,
		TermsOfService: termsOfService,
		CodeInfo:       codeInfo,
	}

	return &authorizationStateWaitCodeTemp
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

// NewAuthorizationStateWaitPassword creates a new AuthorizationStateWaitPassword
//
// @param passwordHint Hint for the password; can be empty
// @param hasRecoveryEmailAddress True if a recovery email address has been set up
// @param recoveryEmailAddressPattern Pattern of the email address to which the recovery email was sent; empty until a recovery email has been sent
func NewAuthorizationStateWaitPassword(passwordHint string, hasRecoveryEmailAddress bool, recoveryEmailAddressPattern string) *AuthorizationStateWaitPassword {
	authorizationStateWaitPasswordTemp := AuthorizationStateWaitPassword{
		tdCommon:                    tdCommon{Type: "authorizationStateWaitPassword"},
		PasswordHint:                passwordHint,
		HasRecoveryEmailAddress:     hasRecoveryEmailAddress,
		RecoveryEmailAddressPattern: recoveryEmailAddressPattern,
	}

	return &authorizationStateWaitPasswordTemp
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

// NewAuthorizationStateReady creates a new AuthorizationStateReady
//
func NewAuthorizationStateReady() *AuthorizationStateReady {
	authorizationStateReadyTemp := AuthorizationStateReady{
		tdCommon: tdCommon{Type: "authorizationStateReady"},
	}

	return &authorizationStateReadyTemp
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

// NewAuthorizationStateLoggingOut creates a new AuthorizationStateLoggingOut
//
func NewAuthorizationStateLoggingOut() *AuthorizationStateLoggingOut {
	authorizationStateLoggingOutTemp := AuthorizationStateLoggingOut{
		tdCommon: tdCommon{Type: "authorizationStateLoggingOut"},
	}

	return &authorizationStateLoggingOutTemp
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

// NewAuthorizationStateClosing creates a new AuthorizationStateClosing
//
func NewAuthorizationStateClosing() *AuthorizationStateClosing {
	authorizationStateClosingTemp := AuthorizationStateClosing{
		tdCommon: tdCommon{Type: "authorizationStateClosing"},
	}

	return &authorizationStateClosingTemp
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

// NewAuthorizationStateClosed creates a new AuthorizationStateClosed
//
func NewAuthorizationStateClosed() *AuthorizationStateClosed {
	authorizationStateClosedTemp := AuthorizationStateClosed{
		tdCommon: tdCommon{Type: "authorizationStateClosed"},
	}

	return &authorizationStateClosedTemp
}

// GetAuthorizationStateEnum return the enum type of this object
func (authorizationStateClosed *AuthorizationStateClosed) GetAuthorizationStateEnum() AuthorizationStateEnum {
	return AuthorizationStateClosedType
}

// PasswordState Represents the current state of 2-step verification
type PasswordState struct {
	tdCommon
	HasPassword                            bool   `json:"has_password"`                               // True if a 2-step verification password is set
	PasswordHint                           string `json:"password_hint"`                              // Hint for the password; can be empty
	HasRecoveryEmailAddress                bool   `json:"has_recovery_email_address"`                 // True if a recovery email is set
	HasPassportData                        bool   `json:"has_passport_data"`                          // True if some Telegram Passport elements were saved
	UnconfirmedRecoveryEmailAddressPattern string `json:"unconfirmed_recovery_email_address_pattern"` // Pattern of the email address to which the confirmation email was sent
}

// MessageType return the string telegram-type of PasswordState
func (passwordState *PasswordState) MessageType() string {
	return "passwordState"
}

// NewPasswordState creates a new PasswordState
//
// @param hasPassword True if a 2-step verification password is set
// @param passwordHint Hint for the password; can be empty
// @param hasRecoveryEmailAddress True if a recovery email is set
// @param hasPassportData True if some Telegram Passport elements were saved
// @param unconfirmedRecoveryEmailAddressPattern Pattern of the email address to which the confirmation email was sent
func NewPasswordState(hasPassword bool, passwordHint string, hasRecoveryEmailAddress bool, hasPassportData bool, unconfirmedRecoveryEmailAddressPattern string) *PasswordState {
	passwordStateTemp := PasswordState{
		tdCommon:                               tdCommon{Type: "passwordState"},
		HasPassword:                            hasPassword,
		PasswordHint:                           passwordHint,
		HasRecoveryEmailAddress:                hasRecoveryEmailAddress,
		HasPassportData:                        hasPassportData,
		UnconfirmedRecoveryEmailAddressPattern: unconfirmedRecoveryEmailAddressPattern,
	}

	return &passwordStateTemp
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

// NewRecoveryEmailAddress creates a new RecoveryEmailAddress
//
// @param recoveryEmailAddress Recovery email address
func NewRecoveryEmailAddress(recoveryEmailAddress string) *RecoveryEmailAddress {
	recoveryEmailAddressTemp := RecoveryEmailAddress{
		tdCommon:             tdCommon{Type: "recoveryEmailAddress"},
		RecoveryEmailAddress: recoveryEmailAddress,
	}

	return &recoveryEmailAddressTemp
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

// NewTemporaryPasswordState creates a new TemporaryPasswordState
//
// @param hasPassword True, if a temporary password is available
// @param validFor Time left before the temporary password expires, in seconds
func NewTemporaryPasswordState(hasPassword bool, validFor int32) *TemporaryPasswordState {
	temporaryPasswordStateTemp := TemporaryPasswordState{
		tdCommon:    tdCommon{Type: "temporaryPasswordState"},
		HasPassword: hasPassword,
		ValidFor:    validFor,
	}

	return &temporaryPasswordStateTemp
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

// NewLocalFile creates a new LocalFile
//
// @param path Local path to the locally available file part; may be empty
// @param canBeDownloaded True, if it is possible to try to download or generate the file
// @param canBeDeleted True, if the file can be deleted
// @param isDownloadingActive True, if the file is currently being downloaded (or a local copy is being generated by some other means)
// @param isDownloadingCompleted True, if the local copy is fully available
// @param downloadedPrefixSize If is_downloading_completed is false, then only some prefix of the file is ready to be read. downloaded_prefix_size is the size of that prefix
// @param downloadedSize Total downloaded file bytes. Should be used only for calculating download progress. The actual file size may be bigger, and some parts of it may contain garbage
func NewLocalFile(path string, canBeDownloaded bool, canBeDeleted bool, isDownloadingActive bool, isDownloadingCompleted bool, downloadedPrefixSize int32, downloadedSize int32) *LocalFile {
	localFileTemp := LocalFile{
		tdCommon:               tdCommon{Type: "localFile"},
		Path:                   path,
		CanBeDownloaded:        canBeDownloaded,
		CanBeDeleted:           canBeDeleted,
		IsDownloadingActive:    isDownloadingActive,
		IsDownloadingCompleted: isDownloadingCompleted,
		DownloadedPrefixSize:   downloadedPrefixSize,
		DownloadedSize:         downloadedSize,
	}

	return &localFileTemp
}

// RemoteFile Represents a remote file
type RemoteFile struct {
	tdCommon
	ID                   string `json:"id"`                     // Remote file identifier; may be empty. Can be used across application restarts or even from other devices for the current user. If the ID starts with "http://" or "https://", it represents the HTTP URL of the file. TDLib is currently unable to download files if only their URL is known.
	IsUploadingActive    bool   `json:"is_uploading_active"`    // True, if the file is currently being uploaded (or a remote copy is being generated by some other means)
	IsUploadingCompleted bool   `json:"is_uploading_completed"` // True, if a remote copy is fully available
	UploadedSize         int32  `json:"uploaded_size"`          // Size of the remote available part of the file; 0 if unknown
}

// MessageType return the string telegram-type of RemoteFile
func (remoteFile *RemoteFile) MessageType() string {
	return "remoteFile"
}

// NewRemoteFile creates a new RemoteFile
//
// @param iD Remote file identifier; may be empty. Can be used across application restarts or even from other devices for the current user. If the ID starts with "http://" or "https://", it represents the HTTP URL of the file. TDLib is currently unable to download files if only their URL is known.
// @param isUploadingActive True, if the file is currently being uploaded (or a remote copy is being generated by some other means)
// @param isUploadingCompleted True, if a remote copy is fully available
// @param uploadedSize Size of the remote available part of the file; 0 if unknown
func NewRemoteFile(iD string, isUploadingActive bool, isUploadingCompleted bool, uploadedSize int32) *RemoteFile {
	remoteFileTemp := RemoteFile{
		tdCommon:             tdCommon{Type: "remoteFile"},
		ID:                   iD,
		IsUploadingActive:    isUploadingActive,
		IsUploadingCompleted: isUploadingCompleted,
		UploadedSize:         uploadedSize,
	}

	return &remoteFileTemp
}

// File Represents a file
type File struct {
	tdCommon
	ID           int32       `json:"id"`            // Unique file identifier
	Size         int32       `json:"size"`          // File size; 0 if unknown
	ExpectedSize int32       `json:"expected_size"` // Expected file size in case the exact file size is unknown, but an approximate size is known. Can be used to show download/upload progress
	Local        *LocalFile  `json:"local"`         // Information about the local copy of the file
	Remote       *RemoteFile `json:"remote"`        // Information about the remote copy of the file
}

// MessageType return the string telegram-type of File
func (file *File) MessageType() string {
	return "file"
}

// NewFile creates a new File
//
// @param iD Unique file identifier
// @param size File size; 0 if unknown
// @param expectedSize Expected file size in case the exact file size is unknown, but an approximate size is known. Can be used to show download/upload progress
// @param local Information about the local copy of the file
// @param remote Information about the remote copy of the file
func NewFile(iD int32, size int32, expectedSize int32, local *LocalFile, remote *RemoteFile) *File {
	fileTemp := File{
		tdCommon:     tdCommon{Type: "file"},
		ID:           iD,
		Size:         size,
		ExpectedSize: expectedSize,
		Local:        local,
		Remote:       remote,
	}

	return &fileTemp
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

// NewInputFileID creates a new InputFileID
//
// @param iD Unique file identifier
func NewInputFileID(iD int32) *InputFileID {
	inputFileIDTemp := InputFileID{
		tdCommon: tdCommon{Type: "inputFileId"},
		ID:       iD,
	}

	return &inputFileIDTemp
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

// NewInputFileRemote creates a new InputFileRemote
//
// @param iD Remote file identifier
func NewInputFileRemote(iD string) *InputFileRemote {
	inputFileRemoteTemp := InputFileRemote{
		tdCommon: tdCommon{Type: "inputFileRemote"},
		ID:       iD,
	}

	return &inputFileRemoteTemp
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

// NewInputFileLocal creates a new InputFileLocal
//
// @param path Local path to the file
func NewInputFileLocal(path string) *InputFileLocal {
	inputFileLocalTemp := InputFileLocal{
		tdCommon: tdCommon{Type: "inputFileLocal"},
		Path:     path,
	}

	return &inputFileLocalTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileLocal *InputFileLocal) GetInputFileEnum() InputFileEnum {
	return InputFileLocalType
}

// InputFileGenerated A file generated by the client
type InputFileGenerated struct {
	tdCommon
	OriginalPath string `json:"original_path"` // Local path to a file from which the file is generated; may be empty if there is no such file
	Conversion   string `json:"conversion"`    // String specifying the conversion applied to the original file; should be persistent across application restarts
	ExpectedSize int32  `json:"expected_size"` // Expected size of the generated file; 0 if unknown
}

// MessageType return the string telegram-type of InputFileGenerated
func (inputFileGenerated *InputFileGenerated) MessageType() string {
	return "inputFileGenerated"
}

// NewInputFileGenerated creates a new InputFileGenerated
//
// @param originalPath Local path to a file from which the file is generated; may be empty if there is no such file
// @param conversion String specifying the conversion applied to the original file; should be persistent across application restarts
// @param expectedSize Expected size of the generated file; 0 if unknown
func NewInputFileGenerated(originalPath string, conversion string, expectedSize int32) *InputFileGenerated {
	inputFileGeneratedTemp := InputFileGenerated{
		tdCommon:     tdCommon{Type: "inputFileGenerated"},
		OriginalPath: originalPath,
		Conversion:   conversion,
		ExpectedSize: expectedSize,
	}

	return &inputFileGeneratedTemp
}

// GetInputFileEnum return the enum type of this object
func (inputFileGenerated *InputFileGenerated) GetInputFileEnum() InputFileEnum {
	return InputFileGeneratedType
}

// PhotoSize Photo description
type PhotoSize struct {
	tdCommon
	Type   string `json:"type"`   // Thumbnail type (see https://core.telegram.org/constructor/photoSize)
	Photo  *File  `json:"photo"`  // Information about the photo file
	Width  int32  `json:"width"`  // Photo width
	Height int32  `json:"height"` // Photo height
}

// MessageType return the string telegram-type of PhotoSize
func (photoSize *PhotoSize) MessageType() string {
	return "photoSize"
}

// NewPhotoSize creates a new PhotoSize
//
// @param typeParam Thumbnail type (see https://core.telegram.org/constructor/photoSize)
// @param photo Information about the photo file
// @param width Photo width
// @param height Photo height
func NewPhotoSize(typeParam string, photo *File, width int32, height int32) *PhotoSize {
	photoSizeTemp := PhotoSize{
		tdCommon: tdCommon{Type: "photoSize"},
		Type:     typeParam,
		Photo:    photo,
		Width:    width,
		Height:   height,
	}

	return &photoSizeTemp
}

// MaskPointForehead A mask should be placed relatively to the forehead
type MaskPointForehead struct {
	tdCommon
}

// MessageType return the string telegram-type of MaskPointForehead
func (maskPointForehead *MaskPointForehead) MessageType() string {
	return "maskPointForehead"
}

// NewMaskPointForehead creates a new MaskPointForehead
//
func NewMaskPointForehead() *MaskPointForehead {
	maskPointForeheadTemp := MaskPointForehead{
		tdCommon: tdCommon{Type: "maskPointForehead"},
	}

	return &maskPointForeheadTemp
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

// NewMaskPointEyes creates a new MaskPointEyes
//
func NewMaskPointEyes() *MaskPointEyes {
	maskPointEyesTemp := MaskPointEyes{
		tdCommon: tdCommon{Type: "maskPointEyes"},
	}

	return &maskPointEyesTemp
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

// NewMaskPointMouth creates a new MaskPointMouth
//
func NewMaskPointMouth() *MaskPointMouth {
	maskPointMouthTemp := MaskPointMouth{
		tdCommon: tdCommon{Type: "maskPointMouth"},
	}

	return &maskPointMouthTemp
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

// NewMaskPointChin creates a new MaskPointChin
//
func NewMaskPointChin() *MaskPointChin {
	maskPointChinTemp := MaskPointChin{
		tdCommon: tdCommon{Type: "maskPointChin"},
	}

	return &maskPointChinTemp
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

// NewMaskPosition creates a new MaskPosition
//
// @param point Part of the face, relative to which the mask should be placed
// @param xShift Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
// @param yShift Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
// @param scale Mask scaling coefficient. (For example, 2.0 means a doubled size)
func NewMaskPosition(point MaskPoint, xShift float64, yShift float64, scale float64) *MaskPosition {
	maskPositionTemp := MaskPosition{
		tdCommon: tdCommon{Type: "maskPosition"},
		Point:    point,
		XShift:   xShift,
		YShift:   yShift,
		Scale:    scale,
	}

	return &maskPositionTemp
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

// Animation Describes an animation file. The animation must be encoded in GIF or MPEG4 format
type Animation struct {
	tdCommon
	Duration  int32      `json:"duration"`  // Duration of the animation, in seconds; as defined by the sender
	Width     int32      `json:"width"`     // Width of the animation
	Height    int32      `json:"height"`    // Height of the animation
	FileName  string     `json:"file_name"` // Original name of the file; as defined by the sender
	MimeType  string     `json:"mime_type"` // MIME type of the file, usually "image/gif" or "video/mp4"
	Thumbnail *PhotoSize `json:"thumbnail"` // Animation thumbnail; may be null
	Animation *File      `json:"animation"` // File containing the animation
}

// MessageType return the string telegram-type of Animation
func (animation *Animation) MessageType() string {
	return "animation"
}

// NewAnimation creates a new Animation
//
// @param duration Duration of the animation, in seconds; as defined by the sender
// @param width Width of the animation
// @param height Height of the animation
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file, usually "image/gif" or "video/mp4"
// @param thumbnail Animation thumbnail; may be null
// @param animation File containing the animation
func NewAnimation(duration int32, width int32, height int32, fileName string, mimeType string, thumbnail *PhotoSize, animation *File) *Animation {
	animationTemp := Animation{
		tdCommon:  tdCommon{Type: "animation"},
		Duration:  duration,
		Width:     width,
		Height:    height,
		FileName:  fileName,
		MimeType:  mimeType,
		Thumbnail: thumbnail,
		Animation: animation,
	}

	return &animationTemp
}

// Audio Describes an audio file. Audio is usually in MP3 format
type Audio struct {
	tdCommon
	Duration            int32      `json:"duration"`              // Duration of the audio, in seconds; as defined by the sender
	Title               string     `json:"title"`                 // Title of the audio; as defined by the sender
	Performer           string     `json:"performer"`             // Performer of the audio; as defined by the sender
	FileName            string     `json:"file_name"`             // Original name of the file; as defined by the sender
	MimeType            string     `json:"mime_type"`             // The MIME type of the file; as defined by the sender
	AlbumCoverThumbnail *PhotoSize `json:"album_cover_thumbnail"` // The thumbnail of the album cover; as defined by the sender. The full size thumbnail should be extracted from the downloaded file; may be null
	Audio               *File      `json:"audio"`                 // File containing the audio
}

// MessageType return the string telegram-type of Audio
func (audio *Audio) MessageType() string {
	return "audio"
}

// NewAudio creates a new Audio
//
// @param duration Duration of the audio, in seconds; as defined by the sender
// @param title Title of the audio; as defined by the sender
// @param performer Performer of the audio; as defined by the sender
// @param fileName Original name of the file; as defined by the sender
// @param mimeType The MIME type of the file; as defined by the sender
// @param albumCoverThumbnail The thumbnail of the album cover; as defined by the sender. The full size thumbnail should be extracted from the downloaded file; may be null
// @param audio File containing the audio
func NewAudio(duration int32, title string, performer string, fileName string, mimeType string, albumCoverThumbnail *PhotoSize, audio *File) *Audio {
	audioTemp := Audio{
		tdCommon:            tdCommon{Type: "audio"},
		Duration:            duration,
		Title:               title,
		Performer:           performer,
		FileName:            fileName,
		MimeType:            mimeType,
		AlbumCoverThumbnail: albumCoverThumbnail,
		Audio:               audio,
	}

	return &audioTemp
}

// Document Describes a document of any type
type Document struct {
	tdCommon
	FileName  string     `json:"file_name"` // Original name of the file; as defined by the sender
	MimeType  string     `json:"mime_type"` // MIME type of the file; as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail"` // Document thumbnail; as defined by the sender; may be null
	Document  *File      `json:"document"`  // File containing the document
}

// MessageType return the string telegram-type of Document
func (document *Document) MessageType() string {
	return "document"
}

// NewDocument creates a new Document
//
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file; as defined by the sender
// @param thumbnail Document thumbnail; as defined by the sender; may be null
// @param document File containing the document
func NewDocument(fileName string, mimeType string, thumbnail *PhotoSize, document *File) *Document {
	documentTemp := Document{
		tdCommon:  tdCommon{Type: "document"},
		FileName:  fileName,
		MimeType:  mimeType,
		Thumbnail: thumbnail,
		Document:  document,
	}

	return &documentTemp
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

// NewPhoto creates a new Photo
//
// @param iD Photo identifier; 0 for deleted photos
// @param hasStickers True, if stickers were added to the photo
// @param sizes Available variants of the photo, in different sizes
func NewPhoto(iD JSONInt64, hasStickers bool, sizes []PhotoSize) *Photo {
	photoTemp := Photo{
		tdCommon:    tdCommon{Type: "photo"},
		ID:          iD,
		HasStickers: hasStickers,
		Sizes:       sizes,
	}

	return &photoTemp
}

// Sticker Describes a sticker
type Sticker struct {
	tdCommon
	SetID        JSONInt64     `json:"set_id"`        // The identifier of the sticker set to which the sticker belongs; 0 if none
	Width        int32         `json:"width"`         // Sticker width; as defined by the sender
	Height       int32         `json:"height"`        // Sticker height; as defined by the sender
	Emoji        string        `json:"emoji"`         // Emoji corresponding to the sticker
	IsMask       bool          `json:"is_mask"`       // True, if the sticker is a mask
	MaskPosition *MaskPosition `json:"mask_position"` // Position where the mask should be placed; may be null
	Thumbnail    *PhotoSize    `json:"thumbnail"`     // Sticker thumbnail in WEBP or JPEG format; may be null
	Sticker      *File         `json:"sticker"`       // File containing the sticker
}

// MessageType return the string telegram-type of Sticker
func (sticker *Sticker) MessageType() string {
	return "sticker"
}

// NewSticker creates a new Sticker
//
// @param setID The identifier of the sticker set to which the sticker belongs; 0 if none
// @param width Sticker width; as defined by the sender
// @param height Sticker height; as defined by the sender
// @param emoji Emoji corresponding to the sticker
// @param isMask True, if the sticker is a mask
// @param maskPosition Position where the mask should be placed; may be null
// @param thumbnail Sticker thumbnail in WEBP or JPEG format; may be null
// @param sticker File containing the sticker
func NewSticker(setID JSONInt64, width int32, height int32, emoji string, isMask bool, maskPosition *MaskPosition, thumbnail *PhotoSize, sticker *File) *Sticker {
	stickerTemp := Sticker{
		tdCommon:     tdCommon{Type: "sticker"},
		SetID:        setID,
		Width:        width,
		Height:       height,
		Emoji:        emoji,
		IsMask:       isMask,
		MaskPosition: maskPosition,
		Thumbnail:    thumbnail,
		Sticker:      sticker,
	}

	return &stickerTemp
}

// Video Describes a video file
type Video struct {
	tdCommon
	Duration          int32      `json:"duration"`           // Duration of the video, in seconds; as defined by the sender
	Width             int32      `json:"width"`              // Video width; as defined by the sender
	Height            int32      `json:"height"`             // Video height; as defined by the sender
	FileName          string     `json:"file_name"`          // Original name of the file; as defined by the sender
	MimeType          string     `json:"mime_type"`          // MIME type of the file; as defined by the sender
	HasStickers       bool       `json:"has_stickers"`       // True, if stickers were added to the photo
	SupportsStreaming bool       `json:"supports_streaming"` // True, if the video should be tried to be streamed
	Thumbnail         *PhotoSize `json:"thumbnail"`          // Video thumbnail; as defined by the sender; may be null
	Video             *File      `json:"video"`              // File containing the video
}

// MessageType return the string telegram-type of Video
func (video *Video) MessageType() string {
	return "video"
}

// NewVideo creates a new Video
//
// @param duration Duration of the video, in seconds; as defined by the sender
// @param width Video width; as defined by the sender
// @param height Video height; as defined by the sender
// @param fileName Original name of the file; as defined by the sender
// @param mimeType MIME type of the file; as defined by the sender
// @param hasStickers True, if stickers were added to the photo
// @param supportsStreaming True, if the video should be tried to be streamed
// @param thumbnail Video thumbnail; as defined by the sender; may be null
// @param video File containing the video
func NewVideo(duration int32, width int32, height int32, fileName string, mimeType string, hasStickers bool, supportsStreaming bool, thumbnail *PhotoSize, video *File) *Video {
	videoTemp := Video{
		tdCommon:          tdCommon{Type: "video"},
		Duration:          duration,
		Width:             width,
		Height:            height,
		FileName:          fileName,
		MimeType:          mimeType,
		HasStickers:       hasStickers,
		SupportsStreaming: supportsStreaming,
		Thumbnail:         thumbnail,
		Video:             video,
	}

	return &videoTemp
}

// VideoNote Describes a video note. The video must be equal in width and height, cropped to a circle, and stored in MPEG4 format
type VideoNote struct {
	tdCommon
	Duration  int32      `json:"duration"`  // Duration of the video, in seconds; as defined by the sender
	Length    int32      `json:"length"`    // Video width and height; as defined by the sender
	Thumbnail *PhotoSize `json:"thumbnail"` // Video thumbnail; as defined by the sender; may be null
	Video     *File      `json:"video"`     // File containing the video
}

// MessageType return the string telegram-type of VideoNote
func (videoNote *VideoNote) MessageType() string {
	return "videoNote"
}

// NewVideoNote creates a new VideoNote
//
// @param duration Duration of the video, in seconds; as defined by the sender
// @param length Video width and height; as defined by the sender
// @param thumbnail Video thumbnail; as defined by the sender; may be null
// @param video File containing the video
func NewVideoNote(duration int32, length int32, thumbnail *PhotoSize, video *File) *VideoNote {
	videoNoteTemp := VideoNote{
		tdCommon:  tdCommon{Type: "videoNote"},
		Duration:  duration,
		Length:    length,
		Thumbnail: thumbnail,
		Video:     video,
	}

	return &videoNoteTemp
}

// VoiceNote Describes a voice note. The voice note must be encoded with the Opus codec, and stored inside an OGG container. Voice notes can have only a single audio channel
type VoiceNote struct {
	tdCommon
	Duration int32  `json:"duration"`  // Duration of the voice note, in seconds; as defined by the sender
	Waveform []byte `json:"waveform"`  // A waveform representation of the voice note in 5-bit format
	MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender
	Voice    *File  `json:"voice"`     // File containing the voice note
}

// MessageType return the string telegram-type of VoiceNote
func (voiceNote *VoiceNote) MessageType() string {
	return "voiceNote"
}

// NewVoiceNote creates a new VoiceNote
//
// @param duration Duration of the voice note, in seconds; as defined by the sender
// @param waveform A waveform representation of the voice note in 5-bit format
// @param mimeType MIME type of the file; as defined by the sender
// @param voice File containing the voice note
func NewVoiceNote(duration int32, waveform []byte, mimeType string, voice *File) *VoiceNote {
	voiceNoteTemp := VoiceNote{
		tdCommon: tdCommon{Type: "voiceNote"},
		Duration: duration,
		Waveform: waveform,
		MimeType: mimeType,
		Voice:    voice,
	}

	return &voiceNoteTemp
}

// Contact Describes a user contact
type Contact struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // Phone number of the user
	FirstName   string `json:"first_name"`   // First name of the user; 1-255 characters in length
	LastName    string `json:"last_name"`    // Last name of the user
	Vcard       string `json:"vcard"`        // Additional data about the user in a form of vCard; 0-2048 bytes in length
	UserID      int32  `json:"user_id"`      // Identifier of the user, if known; otherwise 0
}

// MessageType return the string telegram-type of Contact
func (contact *Contact) MessageType() string {
	return "contact"
}

// NewContact creates a new Contact
//
// @param phoneNumber Phone number of the user
// @param firstName First name of the user; 1-255 characters in length
// @param lastName Last name of the user
// @param vcard Additional data about the user in a form of vCard; 0-2048 bytes in length
// @param userID Identifier of the user, if known; otherwise 0
func NewContact(phoneNumber string, firstName string, lastName string, vcard string, userID int32) *Contact {
	contactTemp := Contact{
		tdCommon:    tdCommon{Type: "contact"},
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
		Vcard:       vcard,
		UserID:      userID,
	}

	return &contactTemp
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

// NewLocation creates a new Location
//
// @param latitude Latitude of the location in degrees; as defined by the sender
// @param longitude Longitude of the location, in degrees; as defined by the sender
func NewLocation(latitude float64, longitude float64) *Location {
	locationTemp := Location{
		tdCommon:  tdCommon{Type: "location"},
		Latitude:  latitude,
		Longitude: longitude,
	}

	return &locationTemp
}

// Venue Describes a venue
type Venue struct {
	tdCommon
	Location *Location `json:"location"` // Venue location; as defined by the sender
	Title    string    `json:"title"`    // Venue name; as defined by the sender
	Address  string    `json:"address"`  // Venue address; as defined by the sender
	Provider string    `json:"provider"` // Provider of the venue database; as defined by the sender. Currently only "foursquare" needs to be supported
	ID       string    `json:"id"`       // Identifier of the venue in the provider database; as defined by the sender
	Type     string    `json:"type"`     // Type of the venue in the provider database; as defined by the sender
}

// MessageType return the string telegram-type of Venue
func (venue *Venue) MessageType() string {
	return "venue"
}

// NewVenue creates a new Venue
//
// @param location Venue location; as defined by the sender
// @param title Venue name; as defined by the sender
// @param address Venue address; as defined by the sender
// @param provider Provider of the venue database; as defined by the sender. Currently only "foursquare" needs to be supported
// @param iD Identifier of the venue in the provider database; as defined by the sender
// @param typeParam Type of the venue in the provider database; as defined by the sender
func NewVenue(location *Location, title string, address string, provider string, iD string, typeParam string) *Venue {
	venueTemp := Venue{
		tdCommon: tdCommon{Type: "venue"},
		Location: location,
		Title:    title,
		Address:  address,
		Provider: provider,
		ID:       iD,
		Type:     typeParam,
	}

	return &venueTemp
}

// Game Describes a game
type Game struct {
	tdCommon
	ID          JSONInt64      `json:"id"`          // Game ID
	ShortName   string         `json:"short_name"`  // Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name}
	Title       string         `json:"title"`       // Game title
	Text        *FormattedText `json:"text"`        // Game text, usually containing scoreboards for a game
	Description string         `json:"description"` //
	Photo       *Photo         `json:"photo"`       // Game photo
	Animation   *Animation     `json:"animation"`   // Game animation; may be null
}

// MessageType return the string telegram-type of Game
func (game *Game) MessageType() string {
	return "game"
}

// NewGame creates a new Game
//
// @param iD Game ID
// @param shortName Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name}
// @param title Game title
// @param text Game text, usually containing scoreboards for a game
// @param description
// @param photo Game photo
// @param animation Game animation; may be null
func NewGame(iD JSONInt64, shortName string, title string, text *FormattedText, description string, photo *Photo, animation *Animation) *Game {
	gameTemp := Game{
		tdCommon:    tdCommon{Type: "game"},
		ID:          iD,
		ShortName:   shortName,
		Title:       title,
		Text:        text,
		Description: description,
		Photo:       photo,
		Animation:   animation,
	}

	return &gameTemp
}

// ProfilePhoto Describes a user profile photo
type ProfilePhoto struct {
	tdCommon
	ID    JSONInt64 `json:"id"`    // Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of userProfilePhotos
	Small *File     `json:"small"` // A small (160x160) user profile photo
	Big   *File     `json:"big"`   // A big (640x640) user profile photo
}

// MessageType return the string telegram-type of ProfilePhoto
func (profilePhoto *ProfilePhoto) MessageType() string {
	return "profilePhoto"
}

// NewProfilePhoto creates a new ProfilePhoto
//
// @param iD Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of userProfilePhotos
// @param small A small (160x160) user profile photo
// @param big A big (640x640) user profile photo
func NewProfilePhoto(iD JSONInt64, small *File, big *File) *ProfilePhoto {
	profilePhotoTemp := ProfilePhoto{
		tdCommon: tdCommon{Type: "profilePhoto"},
		ID:       iD,
		Small:    small,
		Big:      big,
	}

	return &profilePhotoTemp
}

// ChatPhoto Describes the photo of a chat
type ChatPhoto struct {
	tdCommon
	Small *File `json:"small"` // A small (160x160) chat photo
	Big   *File `json:"big"`   // A big (640x640) chat photo
}

// MessageType return the string telegram-type of ChatPhoto
func (chatPhoto *ChatPhoto) MessageType() string {
	return "chatPhoto"
}

// NewChatPhoto creates a new ChatPhoto
//
// @param small A small (160x160) chat photo
// @param big A big (640x640) chat photo
func NewChatPhoto(small *File, big *File) *ChatPhoto {
	chatPhotoTemp := ChatPhoto{
		tdCommon: tdCommon{Type: "chatPhoto"},
		Small:    small,
		Big:      big,
	}

	return &chatPhotoTemp
}

// LinkStateNone The phone number of user A is not known to user B
type LinkStateNone struct {
	tdCommon
}

// MessageType return the string telegram-type of LinkStateNone
func (linkStateNone *LinkStateNone) MessageType() string {
	return "linkStateNone"
}

// NewLinkStateNone creates a new LinkStateNone
//
func NewLinkStateNone() *LinkStateNone {
	linkStateNoneTemp := LinkStateNone{
		tdCommon: tdCommon{Type: "linkStateNone"},
	}

	return &linkStateNoneTemp
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

// NewLinkStateKnowsPhoneNumber creates a new LinkStateKnowsPhoneNumber
//
func NewLinkStateKnowsPhoneNumber() *LinkStateKnowsPhoneNumber {
	linkStateKnowsPhoneNumberTemp := LinkStateKnowsPhoneNumber{
		tdCommon: tdCommon{Type: "linkStateKnowsPhoneNumber"},
	}

	return &linkStateKnowsPhoneNumberTemp
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

// NewLinkStateIsContact creates a new LinkStateIsContact
//
func NewLinkStateIsContact() *LinkStateIsContact {
	linkStateIsContactTemp := LinkStateIsContact{
		tdCommon: tdCommon{Type: "linkStateIsContact"},
	}

	return &linkStateIsContactTemp
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

// NewUserTypeRegular creates a new UserTypeRegular
//
func NewUserTypeRegular() *UserTypeRegular {
	userTypeRegularTemp := UserTypeRegular{
		tdCommon: tdCommon{Type: "userTypeRegular"},
	}

	return &userTypeRegularTemp
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

// NewUserTypeDeleted creates a new UserTypeDeleted
//
func NewUserTypeDeleted() *UserTypeDeleted {
	userTypeDeletedTemp := UserTypeDeleted{
		tdCommon: tdCommon{Type: "userTypeDeleted"},
	}

	return &userTypeDeletedTemp
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

// NewUserTypeBot creates a new UserTypeBot
//
// @param canJoinGroups True, if the bot can be invited to basic group and supergroup chats
// @param canReadAllGroupMessages True, if the bot can read all messages in basic group or supergroup chats and not just those addressed to the bot. In private and channel chats a bot can always read all messages
// @param isInline True, if the bot supports inline queries
// @param inlineQueryPlaceholder Placeholder for inline queries (displayed on the client input field)
// @param needLocation True, if the location of the user should be sent with every inline query to this bot
func NewUserTypeBot(canJoinGroups bool, canReadAllGroupMessages bool, isInline bool, inlineQueryPlaceholder string, needLocation bool) *UserTypeBot {
	userTypeBotTemp := UserTypeBot{
		tdCommon:                tdCommon{Type: "userTypeBot"},
		CanJoinGroups:           canJoinGroups,
		CanReadAllGroupMessages: canReadAllGroupMessages,
		IsInline:                isInline,
		InlineQueryPlaceholder:  inlineQueryPlaceholder,
		NeedLocation:            needLocation,
	}

	return &userTypeBotTemp
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

// NewUserTypeUnknown creates a new UserTypeUnknown
//
func NewUserTypeUnknown() *UserTypeUnknown {
	userTypeUnknownTemp := UserTypeUnknown{
		tdCommon: tdCommon{Type: "userTypeUnknown"},
	}

	return &userTypeUnknownTemp
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

// NewBotCommand creates a new BotCommand
//
// @param command Text of the bot command
// @param description
func NewBotCommand(command string, description string) *BotCommand {
	botCommandTemp := BotCommand{
		tdCommon:    tdCommon{Type: "botCommand"},
		Command:     command,
		Description: description,
	}

	return &botCommandTemp
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

// NewBotInfo creates a new BotInfo
//
// @param description
// @param commands A list of commands supported by the bot
func NewBotInfo(description string, commands []BotCommand) *BotInfo {
	botInfoTemp := BotInfo{
		tdCommon:    tdCommon{Type: "botInfo"},
		Description: description,
		Commands:    commands,
	}

	return &botInfoTemp
}

// User Represents a user
type User struct {
	tdCommon
	ID                int32         `json:"id"`                 // User identifier
	FirstName         string        `json:"first_name"`         // First name of the user
	LastName          string        `json:"last_name"`          // Last name of the user
	Username          string        `json:"username"`           // Username of the user
	PhoneNumber       string        `json:"phone_number"`       // Phone number of the user
	Status            UserStatus    `json:"status"`             // Current online status of the user
	ProfilePhoto      *ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
	OutgoingLink      LinkState     `json:"outgoing_link"`      // Relationship from the current user to the other user
	IncomingLink      LinkState     `json:"incoming_link"`      // Relationship from the other user to the current user
	IsVerified        bool          `json:"is_verified"`        // True, if the user is verified
	RestrictionReason string        `json:"restriction_reason"` // If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
	HaveAccess        bool          `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
	Type              UserType      `json:"type"`               // Type of the user
	LanguageCode      string        `json:"language_code"`      // IETF language tag of the user's language; only available to bots
}

// MessageType return the string telegram-type of User
func (user *User) MessageType() string {
	return "user"
}

// NewUser creates a new User
//
// @param iD User identifier
// @param firstName First name of the user
// @param lastName Last name of the user
// @param username Username of the user
// @param phoneNumber Phone number of the user
// @param status Current online status of the user
// @param profilePhoto Profile photo of the user; may be null
// @param outgoingLink Relationship from the current user to the other user
// @param incomingLink Relationship from the other user to the current user
// @param isVerified True, if the user is verified
// @param restrictionReason If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
// @param haveAccess If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
// @param typeParam Type of the user
// @param languageCode IETF language tag of the user's language; only available to bots
func NewUser(iD int32, firstName string, lastName string, username string, phoneNumber string, status UserStatus, profilePhoto *ProfilePhoto, outgoingLink LinkState, incomingLink LinkState, isVerified bool, restrictionReason string, haveAccess bool, typeParam UserType, languageCode string) *User {
	userTemp := User{
		tdCommon:          tdCommon{Type: "user"},
		ID:                iD,
		FirstName:         firstName,
		LastName:          lastName,
		Username:          username,
		PhoneNumber:       phoneNumber,
		Status:            status,
		ProfilePhoto:      profilePhoto,
		OutgoingLink:      outgoingLink,
		IncomingLink:      incomingLink,
		IsVerified:        isVerified,
		RestrictionReason: restrictionReason,
		HaveAccess:        haveAccess,
		Type:              typeParam,
		LanguageCode:      languageCode,
	}

	return &userTemp
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
		ID                int32         `json:"id"`                 // User identifier
		FirstName         string        `json:"first_name"`         // First name of the user
		LastName          string        `json:"last_name"`          // Last name of the user
		Username          string        `json:"username"`           // Username of the user
		PhoneNumber       string        `json:"phone_number"`       // Phone number of the user
		ProfilePhoto      *ProfilePhoto `json:"profile_photo"`      // Profile photo of the user; may be null
		IsVerified        bool          `json:"is_verified"`        // True, if the user is verified
		RestrictionReason string        `json:"restriction_reason"` // If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
		HaveAccess        bool          `json:"have_access"`        // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser
		LanguageCode      string        `json:"language_code"`      // IETF language tag of the user's language; only available to bots
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
	IsBlocked          bool     `json:"is_blocked"`            // True, if the user is blacklisted by the current user
	CanBeCalled        bool     `json:"can_be_called"`         // True, if the user can be called
	HasPrivateCalls    bool     `json:"has_private_calls"`     // True, if the user can't be called due to their privacy settings
	Bio                string   `json:"bio"`                   // A short user bio
	ShareText          string   `json:"share_text"`            // For bots, the text that is included with the link when users share the bot
	GroupInCommonCount int32    `json:"group_in_common_count"` // Number of group chats where both the other user and the current user are a member; 0 for the current user
	BotInfo            *BotInfo `json:"bot_info"`              // If the user is a bot, information about the bot; may be null
}

// MessageType return the string telegram-type of UserFullInfo
func (userFullInfo *UserFullInfo) MessageType() string {
	return "userFullInfo"
}

// NewUserFullInfo creates a new UserFullInfo
//
// @param isBlocked True, if the user is blacklisted by the current user
// @param canBeCalled True, if the user can be called
// @param hasPrivateCalls True, if the user can't be called due to their privacy settings
// @param bio A short user bio
// @param shareText For bots, the text that is included with the link when users share the bot
// @param groupInCommonCount Number of group chats where both the other user and the current user are a member; 0 for the current user
// @param botInfo If the user is a bot, information about the bot; may be null
func NewUserFullInfo(isBlocked bool, canBeCalled bool, hasPrivateCalls bool, bio string, shareText string, groupInCommonCount int32, botInfo *BotInfo) *UserFullInfo {
	userFullInfoTemp := UserFullInfo{
		tdCommon:           tdCommon{Type: "userFullInfo"},
		IsBlocked:          isBlocked,
		CanBeCalled:        canBeCalled,
		HasPrivateCalls:    hasPrivateCalls,
		Bio:                bio,
		ShareText:          shareText,
		GroupInCommonCount: groupInCommonCount,
		BotInfo:            botInfo,
	}

	return &userFullInfoTemp
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

// NewUserProfilePhotos creates a new UserProfilePhotos
//
// @param totalCount Total number of user profile photos
// @param photos A list of photos
func NewUserProfilePhotos(totalCount int32, photos []Photo) *UserProfilePhotos {
	userProfilePhotosTemp := UserProfilePhotos{
		tdCommon:   tdCommon{Type: "userProfilePhotos"},
		TotalCount: totalCount,
		Photos:     photos,
	}

	return &userProfilePhotosTemp
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

// NewUsers creates a new Users
//
// @param totalCount Approximate total count of users found
// @param userIDs A list of user identifiers
func NewUsers(totalCount int32, userIDs []int32) *Users {
	usersTemp := Users{
		tdCommon:   tdCommon{Type: "users"},
		TotalCount: totalCount,
		UserIDs:    userIDs,
	}

	return &usersTemp
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

// NewChatMemberStatusCreator creates a new ChatMemberStatusCreator
//
// @param isMember True, if the user is a member of the chat
func NewChatMemberStatusCreator(isMember bool) *ChatMemberStatusCreator {
	chatMemberStatusCreatorTemp := ChatMemberStatusCreator{
		tdCommon: tdCommon{Type: "chatMemberStatusCreator"},
		IsMember: isMember,
	}

	return &chatMemberStatusCreatorTemp
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

// NewChatMemberStatusAdministrator creates a new ChatMemberStatusAdministrator
//
// @param canBeEdited True, if the current user can edit the administrator privileges for the called user
// @param canChangeInfo True, if the administrator can change the chat title, photo, and other settings
// @param canPostMessages True, if the administrator can create channel posts; applicable to channels only
// @param canEditMessages True, if the administrator can edit messages of other users and pin messages; applicable to channels only
// @param canDeleteMessages True, if the administrator can delete messages of other users
// @param canInviteUsers True, if the administrator can invite new users to the chat
// @param canRestrictMembers True, if the administrator can restrict, ban, or unban chat members
// @param canPinMessages True, if the administrator can pin messages; applicable to supergroups only
// @param canPromoteMembers True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that were directly or indirectly promoted by him
func NewChatMemberStatusAdministrator(canBeEdited bool, canChangeInfo bool, canPostMessages bool, canEditMessages bool, canDeleteMessages bool, canInviteUsers bool, canRestrictMembers bool, canPinMessages bool, canPromoteMembers bool) *ChatMemberStatusAdministrator {
	chatMemberStatusAdministratorTemp := ChatMemberStatusAdministrator{
		tdCommon:           tdCommon{Type: "chatMemberStatusAdministrator"},
		CanBeEdited:        canBeEdited,
		CanChangeInfo:      canChangeInfo,
		CanPostMessages:    canPostMessages,
		CanEditMessages:    canEditMessages,
		CanDeleteMessages:  canDeleteMessages,
		CanInviteUsers:     canInviteUsers,
		CanRestrictMembers: canRestrictMembers,
		CanPinMessages:     canPinMessages,
		CanPromoteMembers:  canPromoteMembers,
	}

	return &chatMemberStatusAdministratorTemp
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

// NewChatMemberStatusMember creates a new ChatMemberStatusMember
//
func NewChatMemberStatusMember() *ChatMemberStatusMember {
	chatMemberStatusMemberTemp := ChatMemberStatusMember{
		tdCommon: tdCommon{Type: "chatMemberStatusMember"},
	}

	return &chatMemberStatusMemberTemp
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

// NewChatMemberStatusRestricted creates a new ChatMemberStatusRestricted
//
// @param isMember True, if the user is a member of the chat
// @param restrictedUntilDate Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if never. If the user is restricted for more than 366 days or for less than 30 seconds from the current time, the user is considered to be restricted forever
// @param canSendMessages True, if the user can send text messages, contacts, locations, and venues
// @param canSendMediaMessages True, if the user can send audio files, documents, photos, videos, video notes, and voice notes. Implies can_send_messages permissions
// @param canSendOtherMessages True, if the user can send animations, games, and stickers and use inline bots. Implies can_send_media_messages permissions
// @param canAddWebPagePreviews True, if the user may add a web page preview to his messages. Implies can_send_messages permissions
func NewChatMemberStatusRestricted(isMember bool, restrictedUntilDate int32, canSendMessages bool, canSendMediaMessages bool, canSendOtherMessages bool, canAddWebPagePreviews bool) *ChatMemberStatusRestricted {
	chatMemberStatusRestrictedTemp := ChatMemberStatusRestricted{
		tdCommon:              tdCommon{Type: "chatMemberStatusRestricted"},
		IsMember:              isMember,
		RestrictedUntilDate:   restrictedUntilDate,
		CanSendMessages:       canSendMessages,
		CanSendMediaMessages:  canSendMediaMessages,
		CanSendOtherMessages:  canSendOtherMessages,
		CanAddWebPagePreviews: canAddWebPagePreviews,
	}

	return &chatMemberStatusRestrictedTemp
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

// NewChatMemberStatusLeft creates a new ChatMemberStatusLeft
//
func NewChatMemberStatusLeft() *ChatMemberStatusLeft {
	chatMemberStatusLeftTemp := ChatMemberStatusLeft{
		tdCommon: tdCommon{Type: "chatMemberStatusLeft"},
	}

	return &chatMemberStatusLeftTemp
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

// NewChatMemberStatusBanned creates a new ChatMemberStatusBanned
//
// @param bannedUntilDate Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever
func NewChatMemberStatusBanned(bannedUntilDate int32) *ChatMemberStatusBanned {
	chatMemberStatusBannedTemp := ChatMemberStatusBanned{
		tdCommon:        tdCommon{Type: "chatMemberStatusBanned"},
		BannedUntilDate: bannedUntilDate,
	}

	return &chatMemberStatusBannedTemp
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
	BotInfo        *BotInfo         `json:"bot_info"`         // If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
}

// MessageType return the string telegram-type of ChatMember
func (chatMember *ChatMember) MessageType() string {
	return "chatMember"
}

// NewChatMember creates a new ChatMember
//
// @param userID User identifier of the chat member
// @param inviterUserID Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
// @param joinedChatDate Point in time (Unix timestamp) when the user joined a chat
// @param status Status of the member in the chat
// @param botInfo If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
func NewChatMember(userID int32, inviterUserID int32, joinedChatDate int32, status ChatMemberStatus, botInfo *BotInfo) *ChatMember {
	chatMemberTemp := ChatMember{
		tdCommon:       tdCommon{Type: "chatMember"},
		UserID:         userID,
		InviterUserID:  inviterUserID,
		JoinedChatDate: joinedChatDate,
		Status:         status,
		BotInfo:        botInfo,
	}

	return &chatMemberTemp
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
		UserID         int32    `json:"user_id"`          // User identifier of the chat member
		InviterUserID  int32    `json:"inviter_user_id"`  // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
		JoinedChatDate int32    `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined a chat
		BotInfo        *BotInfo `json:"bot_info"`         // If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
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

// NewChatMembers creates a new ChatMembers
//
// @param totalCount Approximate total count of chat members found
// @param members A list of chat members
func NewChatMembers(totalCount int32, members []ChatMember) *ChatMembers {
	chatMembersTemp := ChatMembers{
		tdCommon:   tdCommon{Type: "chatMembers"},
		TotalCount: totalCount,
		Members:    members,
	}

	return &chatMembersTemp
}

// ChatMembersFilterAdministrators Returns the creator and administrators
type ChatMembersFilterAdministrators struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterAdministrators
func (chatMembersFilterAdministrators *ChatMembersFilterAdministrators) MessageType() string {
	return "chatMembersFilterAdministrators"
}

// NewChatMembersFilterAdministrators creates a new ChatMembersFilterAdministrators
//
func NewChatMembersFilterAdministrators() *ChatMembersFilterAdministrators {
	chatMembersFilterAdministratorsTemp := ChatMembersFilterAdministrators{
		tdCommon: tdCommon{Type: "chatMembersFilterAdministrators"},
	}

	return &chatMembersFilterAdministratorsTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterAdministrators *ChatMembersFilterAdministrators) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterAdministratorsType
}

// ChatMembersFilterMembers Returns all chat members, including restricted chat members
type ChatMembersFilterMembers struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterMembers
func (chatMembersFilterMembers *ChatMembersFilterMembers) MessageType() string {
	return "chatMembersFilterMembers"
}

// NewChatMembersFilterMembers creates a new ChatMembersFilterMembers
//
func NewChatMembersFilterMembers() *ChatMembersFilterMembers {
	chatMembersFilterMembersTemp := ChatMembersFilterMembers{
		tdCommon: tdCommon{Type: "chatMembersFilterMembers"},
	}

	return &chatMembersFilterMembersTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterMembers *ChatMembersFilterMembers) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterMembersType
}

// ChatMembersFilterRestricted Returns users under certain restrictions in the chat; can be used only by administrators in a supergroup
type ChatMembersFilterRestricted struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterRestricted
func (chatMembersFilterRestricted *ChatMembersFilterRestricted) MessageType() string {
	return "chatMembersFilterRestricted"
}

// NewChatMembersFilterRestricted creates a new ChatMembersFilterRestricted
//
func NewChatMembersFilterRestricted() *ChatMembersFilterRestricted {
	chatMembersFilterRestrictedTemp := ChatMembersFilterRestricted{
		tdCommon: tdCommon{Type: "chatMembersFilterRestricted"},
	}

	return &chatMembersFilterRestrictedTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterRestricted *ChatMembersFilterRestricted) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterRestrictedType
}

// ChatMembersFilterBanned Returns users banned from the chat; can be used only by administrators in a supergroup or in a channel
type ChatMembersFilterBanned struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterBanned
func (chatMembersFilterBanned *ChatMembersFilterBanned) MessageType() string {
	return "chatMembersFilterBanned"
}

// NewChatMembersFilterBanned creates a new ChatMembersFilterBanned
//
func NewChatMembersFilterBanned() *ChatMembersFilterBanned {
	chatMembersFilterBannedTemp := ChatMembersFilterBanned{
		tdCommon: tdCommon{Type: "chatMembersFilterBanned"},
	}

	return &chatMembersFilterBannedTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterBanned *ChatMembersFilterBanned) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterBannedType
}

// ChatMembersFilterBots Returns bot members of the chat
type ChatMembersFilterBots struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatMembersFilterBots
func (chatMembersFilterBots *ChatMembersFilterBots) MessageType() string {
	return "chatMembersFilterBots"
}

// NewChatMembersFilterBots creates a new ChatMembersFilterBots
//
func NewChatMembersFilterBots() *ChatMembersFilterBots {
	chatMembersFilterBotsTemp := ChatMembersFilterBots{
		tdCommon: tdCommon{Type: "chatMembersFilterBots"},
	}

	return &chatMembersFilterBotsTemp
}

// GetChatMembersFilterEnum return the enum type of this object
func (chatMembersFilterBots *ChatMembersFilterBots) GetChatMembersFilterEnum() ChatMembersFilterEnum {
	return ChatMembersFilterBotsType
}

// SupergroupMembersFilterRecent Returns recently active users in reverse chronological order
type SupergroupMembersFilterRecent struct {
	tdCommon
}

// MessageType return the string telegram-type of SupergroupMembersFilterRecent
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) MessageType() string {
	return "supergroupMembersFilterRecent"
}

// NewSupergroupMembersFilterRecent creates a new SupergroupMembersFilterRecent
//
func NewSupergroupMembersFilterRecent() *SupergroupMembersFilterRecent {
	supergroupMembersFilterRecentTemp := SupergroupMembersFilterRecent{
		tdCommon: tdCommon{Type: "supergroupMembersFilterRecent"},
	}

	return &supergroupMembersFilterRecentTemp
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

// NewSupergroupMembersFilterAdministrators creates a new SupergroupMembersFilterAdministrators
//
func NewSupergroupMembersFilterAdministrators() *SupergroupMembersFilterAdministrators {
	supergroupMembersFilterAdministratorsTemp := SupergroupMembersFilterAdministrators{
		tdCommon: tdCommon{Type: "supergroupMembersFilterAdministrators"},
	}

	return &supergroupMembersFilterAdministratorsTemp
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

// NewSupergroupMembersFilterSearch creates a new SupergroupMembersFilterSearch
//
// @param query Query to search for
func NewSupergroupMembersFilterSearch(query string) *SupergroupMembersFilterSearch {
	supergroupMembersFilterSearchTemp := SupergroupMembersFilterSearch{
		tdCommon: tdCommon{Type: "supergroupMembersFilterSearch"},
		Query:    query,
	}

	return &supergroupMembersFilterSearchTemp
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

// NewSupergroupMembersFilterRestricted creates a new SupergroupMembersFilterRestricted
//
// @param query Query to search for
func NewSupergroupMembersFilterRestricted(query string) *SupergroupMembersFilterRestricted {
	supergroupMembersFilterRestrictedTemp := SupergroupMembersFilterRestricted{
		tdCommon: tdCommon{Type: "supergroupMembersFilterRestricted"},
		Query:    query,
	}

	return &supergroupMembersFilterRestrictedTemp
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

// NewSupergroupMembersFilterBanned creates a new SupergroupMembersFilterBanned
//
// @param query Query to search for
func NewSupergroupMembersFilterBanned(query string) *SupergroupMembersFilterBanned {
	supergroupMembersFilterBannedTemp := SupergroupMembersFilterBanned{
		tdCommon: tdCommon{Type: "supergroupMembersFilterBanned"},
		Query:    query,
	}

	return &supergroupMembersFilterBannedTemp
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

// NewSupergroupMembersFilterBots creates a new SupergroupMembersFilterBots
//
func NewSupergroupMembersFilterBots() *SupergroupMembersFilterBots {
	supergroupMembersFilterBotsTemp := SupergroupMembersFilterBots{
		tdCommon: tdCommon{Type: "supergroupMembersFilterBots"},
	}

	return &supergroupMembersFilterBotsTemp
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

// NewBasicGroup creates a new BasicGroup
//
// @param iD Group identifier
// @param memberCount Number of members in the group
// @param status Status of the current user in the group
// @param everyoneIsAdministrator True, if all members have been granted administrator rights in the group
// @param isActive True, if the group is active
// @param upgradedToSupergroupID Identifier of the supergroup to which this group was upgraded; 0 if none
func NewBasicGroup(iD int32, memberCount int32, status ChatMemberStatus, everyoneIsAdministrator bool, isActive bool, upgradedToSupergroupID int32) *BasicGroup {
	basicGroupTemp := BasicGroup{
		tdCommon:                tdCommon{Type: "basicGroup"},
		ID:                      iD,
		MemberCount:             memberCount,
		Status:                  status,
		EveryoneIsAdministrator: everyoneIsAdministrator,
		IsActive:                isActive,
		UpgradedToSupergroupID:  upgradedToSupergroupID,
	}

	return &basicGroupTemp
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

// NewBasicGroupFullInfo creates a new BasicGroupFullInfo
//
// @param creatorUserID User identifier of the creator of the group; 0 if unknown
// @param members Group members
// @param inviteLink Invite link for this group; available only for the group creator and only after it has been generated at least once
func NewBasicGroupFullInfo(creatorUserID int32, members []ChatMember, inviteLink string) *BasicGroupFullInfo {
	basicGroupFullInfoTemp := BasicGroupFullInfo{
		tdCommon:      tdCommon{Type: "basicGroupFullInfo"},
		CreatorUserID: creatorUserID,
		Members:       members,
		InviteLink:    inviteLink,
	}

	return &basicGroupFullInfoTemp
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

// NewSupergroup creates a new Supergroup
//
// @param iD Supergroup or channel identifier
// @param username Username of the supergroup or channel; empty for private supergroups or channels
// @param date Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
// @param status Status of the current user in the supergroup or channel
// @param memberCount Member count; 0 if unknown. Currently it is guaranteed to be known only if the supergroup or channel was found through SearchPublicChats
// @param anyoneCanInvite True, if any member of the supergroup can invite other members. This field has no meaning for channels
// @param signMessages True, if messages sent to the channel should contain information about the sender. This field is only applicable to channels
// @param isChannel True, if the supergroup is a channel
// @param isVerified True, if the supergroup or channel is verified
// @param restrictionReason If non-empty, contains the reason why access to this supergroup or channel must be restricted. Format of the string is "{type}: {description}".
func NewSupergroup(iD int32, username string, date int32, status ChatMemberStatus, memberCount int32, anyoneCanInvite bool, signMessages bool, isChannel bool, isVerified bool, restrictionReason string) *Supergroup {
	supergroupTemp := Supergroup{
		tdCommon:          tdCommon{Type: "supergroup"},
		ID:                iD,
		Username:          username,
		Date:              date,
		Status:            status,
		MemberCount:       memberCount,
		AnyoneCanInvite:   anyoneCanInvite,
		SignMessages:      signMessages,
		IsChannel:         isChannel,
		IsVerified:        isVerified,
		RestrictionReason: restrictionReason,
	}

	return &supergroupTemp
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

// NewSupergroupFullInfo creates a new SupergroupFullInfo
//
// @param description
// @param memberCount Number of members in the supergroup or channel; 0 if unknown
// @param administratorCount Number of privileged users in the supergroup or channel; 0 if unknown
// @param restrictedCount Number of restricted users in the supergroup; 0 if unknown
// @param bannedCount Number of users banned from chat; 0 if unknown
// @param canGetMembers True, if members of the chat can be retrieved
// @param canSetUsername True, if the chat can be made public
// @param canSetStickerSet True, if the supergroup sticker set can be changed
// @param isAllHistoryAvailable True, if new chat members will have access to old messages. In public supergroups and both public and private channels, old messages are always available, so this option affects only private supergroups. The value of this field is only available for chat administrators
// @param stickerSetID Identifier of the supergroup sticker set; 0 if none
// @param inviteLink Invite link for this chat
// @param pinnedMessageID Identifier of the pinned message in the chat; 0 if none
// @param upgradedFromBasicGroupID Identifier of the basic group from which supergroup was upgraded; 0 if none
// @param upgradedFromMaxMessageID Identifier of the last message in the basic group from which supergroup was upgraded; 0 if none
func NewSupergroupFullInfo(description string, memberCount int32, administratorCount int32, restrictedCount int32, bannedCount int32, canGetMembers bool, canSetUsername bool, canSetStickerSet bool, isAllHistoryAvailable bool, stickerSetID JSONInt64, inviteLink string, pinnedMessageID int64, upgradedFromBasicGroupID int32, upgradedFromMaxMessageID int64) *SupergroupFullInfo {
	supergroupFullInfoTemp := SupergroupFullInfo{
		tdCommon:                 tdCommon{Type: "supergroupFullInfo"},
		Description:              description,
		MemberCount:              memberCount,
		AdministratorCount:       administratorCount,
		RestrictedCount:          restrictedCount,
		BannedCount:              bannedCount,
		CanGetMembers:            canGetMembers,
		CanSetUsername:           canSetUsername,
		CanSetStickerSet:         canSetStickerSet,
		IsAllHistoryAvailable:    isAllHistoryAvailable,
		StickerSetID:             stickerSetID,
		InviteLink:               inviteLink,
		PinnedMessageID:          pinnedMessageID,
		UpgradedFromBasicGroupID: upgradedFromBasicGroupID,
		UpgradedFromMaxMessageID: upgradedFromMaxMessageID,
	}

	return &supergroupFullInfoTemp
}

// SecretChatStatePending The secret chat is not yet created; waiting for the other user to get online
type SecretChatStatePending struct {
	tdCommon
}

// MessageType return the string telegram-type of SecretChatStatePending
func (secretChatStatePending *SecretChatStatePending) MessageType() string {
	return "secretChatStatePending"
}

// NewSecretChatStatePending creates a new SecretChatStatePending
//
func NewSecretChatStatePending() *SecretChatStatePending {
	secretChatStatePendingTemp := SecretChatStatePending{
		tdCommon: tdCommon{Type: "secretChatStatePending"},
	}

	return &secretChatStatePendingTemp
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

// NewSecretChatStateReady creates a new SecretChatStateReady
//
func NewSecretChatStateReady() *SecretChatStateReady {
	secretChatStateReadyTemp := SecretChatStateReady{
		tdCommon: tdCommon{Type: "secretChatStateReady"},
	}

	return &secretChatStateReadyTemp
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

// NewSecretChatStateClosed creates a new SecretChatStateClosed
//
func NewSecretChatStateClosed() *SecretChatStateClosed {
	secretChatStateClosedTemp := SecretChatStateClosed{
		tdCommon: tdCommon{Type: "secretChatStateClosed"},
	}

	return &secretChatStateClosedTemp
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

// NewSecretChat creates a new SecretChat
//
// @param iD Secret chat identifier
// @param userID Identifier of the chat partner
// @param state State of the secret chat
// @param isOutbound True, if the chat was created by the current user; otherwise false
// @param tTL Current message Time To Live setting (self-destruct timer) for the chat, in seconds
// @param keyHash Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 bytes, which must be used to make a 12x12 square image with a color depth of 4. The first 16 bytes should be used to make a central 8x8 square, while the remaining 20 bytes should be used to construct a 2-pixel-wide border around that square.
// @param layer Secret chat layer; determines features supported by the other client. Video notes are supported if the layer >= 66
func NewSecretChat(iD int32, userID int32, state SecretChatState, isOutbound bool, tTL int32, keyHash []byte, layer int32) *SecretChat {
	secretChatTemp := SecretChat{
		tdCommon:   tdCommon{Type: "secretChat"},
		ID:         iD,
		UserID:     userID,
		State:      state,
		IsOutbound: isOutbound,
		TTL:        tTL,
		KeyHash:    keyHash,
		Layer:      layer,
	}

	return &secretChatTemp
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

// NewMessageForwardedFromUser creates a new MessageForwardedFromUser
//
// @param senderUserID Identifier of the user that originally sent this message
// @param date Point in time (Unix timestamp) when the message was originally sent
// @param forwardedFromChatID For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
// @param forwardedFromMessageID For messages forwarded to the chat with the current user (saved messages) the identifier of the original message from which the new message was forwarded; 0 if unknown
func NewMessageForwardedFromUser(senderUserID int32, date int32, forwardedFromChatID int64, forwardedFromMessageID int64) *MessageForwardedFromUser {
	messageForwardedFromUserTemp := MessageForwardedFromUser{
		tdCommon:               tdCommon{Type: "messageForwardedFromUser"},
		SenderUserID:           senderUserID,
		Date:                   date,
		ForwardedFromChatID:    forwardedFromChatID,
		ForwardedFromMessageID: forwardedFromMessageID,
	}

	return &messageForwardedFromUserTemp
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

// NewMessageForwardedPost creates a new MessageForwardedPost
//
// @param chatID Identifier of the chat from which the message was forwarded
// @param authorSignature Post author signature
// @param date Point in time (Unix timestamp) when the message was originally sent
// @param messageID Message identifier of the original message from which the new message was forwarded; 0 if unknown
// @param forwardedFromChatID For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
// @param forwardedFromMessageID For messages forwarded to the chat with the current user (saved messages), the identifier of the original message from which the new message was forwarded; 0 if unknown
func NewMessageForwardedPost(chatID int64, authorSignature string, date int32, messageID int64, forwardedFromChatID int64, forwardedFromMessageID int64) *MessageForwardedPost {
	messageForwardedPostTemp := MessageForwardedPost{
		tdCommon:               tdCommon{Type: "messageForwardedPost"},
		ChatID:                 chatID,
		AuthorSignature:        authorSignature,
		Date:                   date,
		MessageID:              messageID,
		ForwardedFromChatID:    forwardedFromChatID,
		ForwardedFromMessageID: forwardedFromMessageID,
	}

	return &messageForwardedPostTemp
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

// NewMessageSendingStatePending creates a new MessageSendingStatePending
//
func NewMessageSendingStatePending() *MessageSendingStatePending {
	messageSendingStatePendingTemp := MessageSendingStatePending{
		tdCommon: tdCommon{Type: "messageSendingStatePending"},
	}

	return &messageSendingStatePendingTemp
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

// NewMessageSendingStateFailed creates a new MessageSendingStateFailed
//
func NewMessageSendingStateFailed() *MessageSendingStateFailed {
	messageSendingStateFailedTemp := MessageSendingStateFailed{
		tdCommon: tdCommon{Type: "messageSendingStateFailed"},
	}

	return &messageSendingStateFailedTemp
}

// GetMessageSendingStateEnum return the enum type of this object
func (messageSendingStateFailed *MessageSendingStateFailed) GetMessageSendingStateEnum() MessageSendingStateEnum {
	return MessageSendingStateFailedType
}

// Message Describes a message
type Message struct {
	tdCommon
	ID                      int64               `json:"id"`                           // Message identifier, unique for the chat to which the message belongs
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

// NewMessage creates a new Message
//
// @param iD Message identifier, unique for the chat to which the message belongs
// @param senderUserID Identifier of the user who sent the message; 0 if unknown. It is unknown for channel posts
// @param chatID Chat identifier
// @param sendingState Information about the sending state of the message; may be null
// @param isOutgoing True, if the message is outgoing
// @param canBeEdited True, if the message can be edited
// @param canBeForwarded True, if the message can be forwarded
// @param canBeDeletedOnlyForSelf True, if the message can be deleted only for the current user while other users will continue to see it
// @param canBeDeletedForAllUsers True, if the message can be deleted for all users
// @param isChannelPost True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
// @param containsUnreadMention True, if the message contains an unread mention for the current user
// @param date Point in time (Unix timestamp) when the message was sent
// @param editDate Point in time (Unix timestamp) when the message was last edited
// @param forwardInfo Information about the initial message sender; may be null
// @param replyToMessageID If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
// @param tTL For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
// @param tTLExpiresIn Time left before the message expires, in seconds
// @param viaBotUserID If non-zero, the user identifier of the bot through which this message was sent
// @param authorSignature For channel posts, optional author signature
// @param views Number of times this message was viewed
// @param mediaAlbumID Unique identifier of an album this message belongs to. Only photos and videos can be grouped together in albums
// @param content Content of the message
// @param replyMarkup Reply markup for the message; may be null
func NewMessage(iD int64, senderUserID int32, chatID int64, sendingState MessageSendingState, isOutgoing bool, canBeEdited bool, canBeForwarded bool, canBeDeletedOnlyForSelf bool, canBeDeletedForAllUsers bool, isChannelPost bool, containsUnreadMention bool, date int32, editDate int32, forwardInfo MessageForwardInfo, replyToMessageID int64, tTL int32, tTLExpiresIn float64, viaBotUserID int32, authorSignature string, views int32, mediaAlbumID JSONInt64, content MessageContent, replyMarkup ReplyMarkup) *Message {
	messageTemp := Message{
		tdCommon:                tdCommon{Type: "message"},
		ID:                      iD,
		SenderUserID:            senderUserID,
		ChatID:                  chatID,
		SendingState:            sendingState,
		IsOutgoing:              isOutgoing,
		CanBeEdited:             canBeEdited,
		CanBeForwarded:          canBeForwarded,
		CanBeDeletedOnlyForSelf: canBeDeletedOnlyForSelf,
		CanBeDeletedForAllUsers: canBeDeletedForAllUsers,
		IsChannelPost:           isChannelPost,
		ContainsUnreadMention:   containsUnreadMention,
		Date:                    date,
		EditDate:                editDate,
		ForwardInfo:             forwardInfo,
		ReplyToMessageID:        replyToMessageID,
		TTL:                     tTL,
		TTLExpiresIn:            tTLExpiresIn,
		ViaBotUserID:            viaBotUserID,
		AuthorSignature:         authorSignature,
		Views:                   views,
		MediaAlbumID:            mediaAlbumID,
		Content:                 content,
		ReplyMarkup:             replyMarkup,
	}

	return &messageTemp
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
		ID                      int64     `json:"id"`                           // Message identifier, unique for the chat to which the message belongs
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

// NewMessages creates a new Messages
//
// @param totalCount Approximate total count of messages found
// @param messages List of messages; messages may be null
func NewMessages(totalCount int32, messages []Message) *Messages {
	messagesTemp := Messages{
		tdCommon:   tdCommon{Type: "messages"},
		TotalCount: totalCount,
		Messages:   messages,
	}

	return &messagesTemp
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

// NewFoundMessages creates a new FoundMessages
//
// @param messages List of messages
// @param nextFromSearchID Value to pass as from_search_id to get more results
func NewFoundMessages(messages []Message, nextFromSearchID JSONInt64) *FoundMessages {
	foundMessagesTemp := FoundMessages{
		tdCommon:         tdCommon{Type: "foundMessages"},
		Messages:         messages,
		NextFromSearchID: nextFromSearchID,
	}

	return &foundMessagesTemp
}

// NotificationSettingsScopePrivateChats Notification settings applied to all private and secret chats when the corresponding chat setting has a default value
type NotificationSettingsScopePrivateChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopePrivateChats
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) MessageType() string {
	return "notificationSettingsScopePrivateChats"
}

// NewNotificationSettingsScopePrivateChats creates a new NotificationSettingsScopePrivateChats
//
func NewNotificationSettingsScopePrivateChats() *NotificationSettingsScopePrivateChats {
	notificationSettingsScopePrivateChatsTemp := NotificationSettingsScopePrivateChats{
		tdCommon: tdCommon{Type: "notificationSettingsScopePrivateChats"},
	}

	return &notificationSettingsScopePrivateChatsTemp
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopePrivateChatsType
}

// NotificationSettingsScopeGroupChats Notification settings applied to all basic groups, supergroups and channels when the corresponding chat setting has a default value
type NotificationSettingsScopeGroupChats struct {
	tdCommon
}

// MessageType return the string telegram-type of NotificationSettingsScopeGroupChats
func (notificationSettingsScopeGroupChats *NotificationSettingsScopeGroupChats) MessageType() string {
	return "notificationSettingsScopeGroupChats"
}

// NewNotificationSettingsScopeGroupChats creates a new NotificationSettingsScopeGroupChats
//
func NewNotificationSettingsScopeGroupChats() *NotificationSettingsScopeGroupChats {
	notificationSettingsScopeGroupChatsTemp := NotificationSettingsScopeGroupChats{
		tdCommon: tdCommon{Type: "notificationSettingsScopeGroupChats"},
	}

	return &notificationSettingsScopeGroupChatsTemp
}

// GetNotificationSettingsScopeEnum return the enum type of this object
func (notificationSettingsScopeGroupChats *NotificationSettingsScopeGroupChats) GetNotificationSettingsScopeEnum() NotificationSettingsScopeEnum {
	return NotificationSettingsScopeGroupChatsType
}

// ChatNotificationSettings Contains information about notification settings for a chat
type ChatNotificationSettings struct {
	tdCommon
	UseDefaultMuteFor     bool   `json:"use_default_mute_for"`     // If true, mute_for is ignored and the value for the relevant type of chat is used instead
	MuteFor               int32  `json:"mute_for"`                 // Time left before notifications will be unmuted, in seconds
	UseDefaultSound       bool   `json:"use_default_sound"`        // If true, sound is ignored and the value for the relevant type of chat is used instead
	Sound                 string `json:"sound"`                    // The name of an audio file to be used for notification sounds; only applies to iOS applications
	UseDefaultShowPreview bool   `json:"use_default_show_preview"` // If true, show_preview is ignored and the value for the relevant type of chat is used instead
	ShowPreview           bool   `json:"show_preview"`             // True, if message content should be displayed in notifications
}

// MessageType return the string telegram-type of ChatNotificationSettings
func (chatNotificationSettings *ChatNotificationSettings) MessageType() string {
	return "chatNotificationSettings"
}

// NewChatNotificationSettings creates a new ChatNotificationSettings
//
// @param useDefaultMuteFor If true, mute_for is ignored and the value for the relevant type of chat is used instead
// @param muteFor Time left before notifications will be unmuted, in seconds
// @param useDefaultSound If true, sound is ignored and the value for the relevant type of chat is used instead
// @param sound The name of an audio file to be used for notification sounds; only applies to iOS applications
// @param useDefaultShowPreview If true, show_preview is ignored and the value for the relevant type of chat is used instead
// @param showPreview True, if message content should be displayed in notifications
func NewChatNotificationSettings(useDefaultMuteFor bool, muteFor int32, useDefaultSound bool, sound string, useDefaultShowPreview bool, showPreview bool) *ChatNotificationSettings {
	chatNotificationSettingsTemp := ChatNotificationSettings{
		tdCommon:              tdCommon{Type: "chatNotificationSettings"},
		UseDefaultMuteFor:     useDefaultMuteFor,
		MuteFor:               muteFor,
		UseDefaultSound:       useDefaultSound,
		Sound:                 sound,
		UseDefaultShowPreview: useDefaultShowPreview,
		ShowPreview:           showPreview,
	}

	return &chatNotificationSettingsTemp
}

// ScopeNotificationSettings Contains information about notification settings for several chats
type ScopeNotificationSettings struct {
	tdCommon
	MuteFor     int32  `json:"mute_for"`     // Time left before notifications will be unmuted, in seconds
	Sound       string `json:"sound"`        // The name of an audio file to be used for notification sounds; only applies to iOS applications
	ShowPreview bool   `json:"show_preview"` // True, if message content should be displayed in notifications
}

// MessageType return the string telegram-type of ScopeNotificationSettings
func (scopeNotificationSettings *ScopeNotificationSettings) MessageType() string {
	return "scopeNotificationSettings"
}

// NewScopeNotificationSettings creates a new ScopeNotificationSettings
//
// @param muteFor Time left before notifications will be unmuted, in seconds
// @param sound The name of an audio file to be used for notification sounds; only applies to iOS applications
// @param showPreview True, if message content should be displayed in notifications
func NewScopeNotificationSettings(muteFor int32, sound string, showPreview bool) *ScopeNotificationSettings {
	scopeNotificationSettingsTemp := ScopeNotificationSettings{
		tdCommon:    tdCommon{Type: "scopeNotificationSettings"},
		MuteFor:     muteFor,
		Sound:       sound,
		ShowPreview: showPreview,
	}

	return &scopeNotificationSettingsTemp
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

// NewDraftMessage creates a new DraftMessage
//
// @param replyToMessageID Identifier of the message to reply to; 0 if none
// @param inputMessageText Content of the message draft; this should always be of type inputMessageText
func NewDraftMessage(replyToMessageID int64, inputMessageText InputMessageContent) *DraftMessage {
	draftMessageTemp := DraftMessage{
		tdCommon:         tdCommon{Type: "draftMessage"},
		ReplyToMessageID: replyToMessageID,
		InputMessageText: inputMessageText,
	}

	return &draftMessageTemp
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

// NewChatTypePrivate creates a new ChatTypePrivate
//
// @param userID User identifier
func NewChatTypePrivate(userID int32) *ChatTypePrivate {
	chatTypePrivateTemp := ChatTypePrivate{
		tdCommon: tdCommon{Type: "chatTypePrivate"},
		UserID:   userID,
	}

	return &chatTypePrivateTemp
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

// NewChatTypeBasicGroup creates a new ChatTypeBasicGroup
//
// @param basicGroupID Basic group identifier
func NewChatTypeBasicGroup(basicGroupID int32) *ChatTypeBasicGroup {
	chatTypeBasicGroupTemp := ChatTypeBasicGroup{
		tdCommon:     tdCommon{Type: "chatTypeBasicGroup"},
		BasicGroupID: basicGroupID,
	}

	return &chatTypeBasicGroupTemp
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

// NewChatTypeSupergroup creates a new ChatTypeSupergroup
//
// @param supergroupID Supergroup or channel identifier
// @param isChannel True, if the supergroup is a channel
func NewChatTypeSupergroup(supergroupID int32, isChannel bool) *ChatTypeSupergroup {
	chatTypeSupergroupTemp := ChatTypeSupergroup{
		tdCommon:     tdCommon{Type: "chatTypeSupergroup"},
		SupergroupID: supergroupID,
		IsChannel:    isChannel,
	}

	return &chatTypeSupergroupTemp
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

// NewChatTypeSecret creates a new ChatTypeSecret
//
// @param secretChatID Secret chat identifier
// @param userID User identifier of the secret chat peer
func NewChatTypeSecret(secretChatID int32, userID int32) *ChatTypeSecret {
	chatTypeSecretTemp := ChatTypeSecret{
		tdCommon:     tdCommon{Type: "chatTypeSecret"},
		SecretChatID: secretChatID,
		UserID:       userID,
	}

	return &chatTypeSecretTemp
}

// GetChatTypeEnum return the enum type of this object
func (chatTypeSecret *ChatTypeSecret) GetChatTypeEnum() ChatTypeEnum {
	return ChatTypeSecretType
}

// Chat A chat. (Can be a private chat, basic group, supergroup, or secret chat)
type Chat struct {
	tdCommon
	ID                         int64                     `json:"id"`                           // Chat unique identifier
	Type                       ChatType                  `json:"type"`                         // Type of the chat
	Title                      string                    `json:"title"`                        // Chat title
	Photo                      *ChatPhoto                `json:"photo"`                        // Chat photo; may be null
	LastMessage                *Message                  `json:"last_message"`                 // Last message in the chat; may be null
	Order                      JSONInt64                 `json:"order"`                        // Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
	IsPinned                   bool                      `json:"is_pinned"`                    // True, if the chat is pinned
	IsMarkedAsUnread           bool                      `json:"is_marked_as_unread"`          // True, if the chat is marked as unread
	IsSponsored                bool                      `json:"is_sponsored"`                 // True, if the chat is sponsored by the user's MTProxy server
	CanBeReported              bool                      `json:"can_be_reported"`              // True, if the chat can be reported to Telegram moderators through reportChat
	DefaultDisableNotification bool                      `json:"default_disable_notification"` // Default value of the disable_notification parameter, used when a message is sent to the chat
	UnreadCount                int32                     `json:"unread_count"`                 // Number of unread messages in the chat
	LastReadInboxMessageID     int64                     `json:"last_read_inbox_message_id"`   // Identifier of the last read incoming message
	LastReadOutboxMessageID    int64                     `json:"last_read_outbox_message_id"`  // Identifier of the last read outgoing message
	UnreadMentionCount         int32                     `json:"unread_mention_count"`         // Number of unread messages with a mention/reply in the chat
	NotificationSettings       *ChatNotificationSettings `json:"notification_settings"`        // Notification settings for this chat
	ReplyMarkupMessageID       int64                     `json:"reply_markup_message_id"`      // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
	DraftMessage               *DraftMessage             `json:"draft_message"`                // A draft of a message in the chat; may be null
	ClientData                 string                    `json:"client_data"`                  // Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
}

// MessageType return the string telegram-type of Chat
func (chat *Chat) MessageType() string {
	return "chat"
}

// NewChat creates a new Chat
//
// @param iD Chat unique identifier
// @param typeParam Type of the chat
// @param title Chat title
// @param photo Chat photo; may be null
// @param lastMessage Last message in the chat; may be null
// @param order Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
// @param isPinned True, if the chat is pinned
// @param isMarkedAsUnread True, if the chat is marked as unread
// @param isSponsored True, if the chat is sponsored by the user's MTProxy server
// @param canBeReported True, if the chat can be reported to Telegram moderators through reportChat
// @param defaultDisableNotification Default value of the disable_notification parameter, used when a message is sent to the chat
// @param unreadCount Number of unread messages in the chat
// @param lastReadInboxMessageID Identifier of the last read incoming message
// @param lastReadOutboxMessageID Identifier of the last read outgoing message
// @param unreadMentionCount Number of unread messages with a mention/reply in the chat
// @param notificationSettings Notification settings for this chat
// @param replyMarkupMessageID Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
// @param draftMessage A draft of a message in the chat; may be null
// @param clientData Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
func NewChat(iD int64, typeParam ChatType, title string, photo *ChatPhoto, lastMessage *Message, order JSONInt64, isPinned bool, isMarkedAsUnread bool, isSponsored bool, canBeReported bool, defaultDisableNotification bool, unreadCount int32, lastReadInboxMessageID int64, lastReadOutboxMessageID int64, unreadMentionCount int32, notificationSettings *ChatNotificationSettings, replyMarkupMessageID int64, draftMessage *DraftMessage, clientData string) *Chat {
	chatTemp := Chat{
		tdCommon:                   tdCommon{Type: "chat"},
		ID:                         iD,
		Type:                       typeParam,
		Title:                      title,
		Photo:                      photo,
		LastMessage:                lastMessage,
		Order:                      order,
		IsPinned:                   isPinned,
		IsMarkedAsUnread:           isMarkedAsUnread,
		IsSponsored:                isSponsored,
		CanBeReported:              canBeReported,
		DefaultDisableNotification: defaultDisableNotification,
		UnreadCount:                unreadCount,
		LastReadInboxMessageID:     lastReadInboxMessageID,
		LastReadOutboxMessageID:    lastReadOutboxMessageID,
		UnreadMentionCount:         unreadMentionCount,
		NotificationSettings:       notificationSettings,
		ReplyMarkupMessageID:       replyMarkupMessageID,
		DraftMessage:               draftMessage,
		ClientData:                 clientData,
	}

	return &chatTemp
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
		ID                         int64                     `json:"id"`                           // Chat unique identifier
		Title                      string                    `json:"title"`                        // Chat title
		Photo                      *ChatPhoto                `json:"photo"`                        // Chat photo; may be null
		LastMessage                *Message                  `json:"last_message"`                 // Last message in the chat; may be null
		Order                      JSONInt64                 `json:"order"`                        // Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
		IsPinned                   bool                      `json:"is_pinned"`                    // True, if the chat is pinned
		IsMarkedAsUnread           bool                      `json:"is_marked_as_unread"`          // True, if the chat is marked as unread
		IsSponsored                bool                      `json:"is_sponsored"`                 // True, if the chat is sponsored by the user's MTProxy server
		CanBeReported              bool                      `json:"can_be_reported"`              // True, if the chat can be reported to Telegram moderators through reportChat
		DefaultDisableNotification bool                      `json:"default_disable_notification"` // Default value of the disable_notification parameter, used when a message is sent to the chat
		UnreadCount                int32                     `json:"unread_count"`                 // Number of unread messages in the chat
		LastReadInboxMessageID     int64                     `json:"last_read_inbox_message_id"`   // Identifier of the last read incoming message
		LastReadOutboxMessageID    int64                     `json:"last_read_outbox_message_id"`  // Identifier of the last read outgoing message
		UnreadMentionCount         int32                     `json:"unread_mention_count"`         // Number of unread messages with a mention/reply in the chat
		NotificationSettings       *ChatNotificationSettings `json:"notification_settings"`        // Notification settings for this chat
		ReplyMarkupMessageID       int64                     `json:"reply_markup_message_id"`      // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
		DraftMessage               *DraftMessage             `json:"draft_message"`                // A draft of a message in the chat; may be null
		ClientData                 string                    `json:"client_data"`                  // Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
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
	chat.IsMarkedAsUnread = tempObj.IsMarkedAsUnread
	chat.IsSponsored = tempObj.IsSponsored
	chat.CanBeReported = tempObj.CanBeReported
	chat.DefaultDisableNotification = tempObj.DefaultDisableNotification
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

// NewChats creates a new Chats
//
// @param chatIDs List of chat identifiers
func NewChats(chatIDs []int64) *Chats {
	chatsTemp := Chats{
		tdCommon: tdCommon{Type: "chats"},
		ChatIDs:  chatIDs,
	}

	return &chatsTemp
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

// NewChatInviteLink creates a new ChatInviteLink
//
// @param inviteLink Chat invite link
func NewChatInviteLink(inviteLink string) *ChatInviteLink {
	chatInviteLinkTemp := ChatInviteLink{
		tdCommon:   tdCommon{Type: "chatInviteLink"},
		InviteLink: inviteLink,
	}

	return &chatInviteLinkTemp
}

// ChatInviteLinkInfo Contains information about a chat invite link
type ChatInviteLinkInfo struct {
	tdCommon
	ChatID        int64      `json:"chat_id"`         // Chat identifier of the invite link; 0 if the user is not a member of this chat
	Type          ChatType   `json:"type"`            // Contains information about the type of the chat
	Title         string     `json:"title"`           // Title of the chat
	Photo         *ChatPhoto `json:"photo"`           // Chat photo; may be null
	MemberCount   int32      `json:"member_count"`    // Number of members
	MemberUserIDs []int32    `json:"member_user_ids"` // User identifiers of some chat members that may be known to the current user
	IsPublic      bool       `json:"is_public"`       // True, if the chat is a public supergroup or channel with a username
}

// MessageType return the string telegram-type of ChatInviteLinkInfo
func (chatInviteLinkInfo *ChatInviteLinkInfo) MessageType() string {
	return "chatInviteLinkInfo"
}

// NewChatInviteLinkInfo creates a new ChatInviteLinkInfo
//
// @param chatID Chat identifier of the invite link; 0 if the user is not a member of this chat
// @param typeParam Contains information about the type of the chat
// @param title Title of the chat
// @param photo Chat photo; may be null
// @param memberCount Number of members
// @param memberUserIDs User identifiers of some chat members that may be known to the current user
// @param isPublic True, if the chat is a public supergroup or channel with a username
func NewChatInviteLinkInfo(chatID int64, typeParam ChatType, title string, photo *ChatPhoto, memberCount int32, memberUserIDs []int32, isPublic bool) *ChatInviteLinkInfo {
	chatInviteLinkInfoTemp := ChatInviteLinkInfo{
		tdCommon:      tdCommon{Type: "chatInviteLinkInfo"},
		ChatID:        chatID,
		Type:          typeParam,
		Title:         title,
		Photo:         photo,
		MemberCount:   memberCount,
		MemberUserIDs: memberUserIDs,
		IsPublic:      isPublic,
	}

	return &chatInviteLinkInfoTemp
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
		ChatID        int64      `json:"chat_id"`         // Chat identifier of the invite link; 0 if the user is not a member of this chat
		Title         string     `json:"title"`           // Title of the chat
		Photo         *ChatPhoto `json:"photo"`           // Chat photo; may be null
		MemberCount   int32      `json:"member_count"`    // Number of members
		MemberUserIDs []int32    `json:"member_user_ids"` // User identifiers of some chat members that may be known to the current user
		IsPublic      bool       `json:"is_public"`       // True, if the chat is a public supergroup or channel with a username
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

// NewKeyboardButtonTypeText creates a new KeyboardButtonTypeText
//
func NewKeyboardButtonTypeText() *KeyboardButtonTypeText {
	keyboardButtonTypeTextTemp := KeyboardButtonTypeText{
		tdCommon: tdCommon{Type: "keyboardButtonTypeText"},
	}

	return &keyboardButtonTypeTextTemp
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

// NewKeyboardButtonTypeRequestPhoneNumber creates a new KeyboardButtonTypeRequestPhoneNumber
//
func NewKeyboardButtonTypeRequestPhoneNumber() *KeyboardButtonTypeRequestPhoneNumber {
	keyboardButtonTypeRequestPhoneNumberTemp := KeyboardButtonTypeRequestPhoneNumber{
		tdCommon: tdCommon{Type: "keyboardButtonTypeRequestPhoneNumber"},
	}

	return &keyboardButtonTypeRequestPhoneNumberTemp
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

// NewKeyboardButtonTypeRequestLocation creates a new KeyboardButtonTypeRequestLocation
//
func NewKeyboardButtonTypeRequestLocation() *KeyboardButtonTypeRequestLocation {
	keyboardButtonTypeRequestLocationTemp := KeyboardButtonTypeRequestLocation{
		tdCommon: tdCommon{Type: "keyboardButtonTypeRequestLocation"},
	}

	return &keyboardButtonTypeRequestLocationTemp
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

// NewKeyboardButton creates a new KeyboardButton
//
// @param text Text of the button
// @param typeParam Type of the button
func NewKeyboardButton(text string, typeParam KeyboardButtonType) *KeyboardButton {
	keyboardButtonTemp := KeyboardButton{
		tdCommon: tdCommon{Type: "keyboardButton"},
		Text:     text,
		Type:     typeParam,
	}

	return &keyboardButtonTemp
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
	URL string `json:"url"` // HTTP or tg:// URL to open
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeURL
func (inlineKeyboardButtonTypeURL *InlineKeyboardButtonTypeURL) MessageType() string {
	return "inlineKeyboardButtonTypeUrl"
}

// NewInlineKeyboardButtonTypeURL creates a new InlineKeyboardButtonTypeURL
//
// @param uRL HTTP or tg:// URL to open
func NewInlineKeyboardButtonTypeURL(uRL string) *InlineKeyboardButtonTypeURL {
	inlineKeyboardButtonTypeURLTemp := InlineKeyboardButtonTypeURL{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeUrl"},
		URL:      uRL,
	}

	return &inlineKeyboardButtonTypeURLTemp
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

// NewInlineKeyboardButtonTypeCallback creates a new InlineKeyboardButtonTypeCallback
//
// @param data Data to be sent to the bot via a callback query
func NewInlineKeyboardButtonTypeCallback(data []byte) *InlineKeyboardButtonTypeCallback {
	inlineKeyboardButtonTypeCallbackTemp := InlineKeyboardButtonTypeCallback{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeCallback"},
		Data:     data,
	}

	return &inlineKeyboardButtonTypeCallbackTemp
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

// NewInlineKeyboardButtonTypeCallbackGame creates a new InlineKeyboardButtonTypeCallbackGame
//
func NewInlineKeyboardButtonTypeCallbackGame() *InlineKeyboardButtonTypeCallbackGame {
	inlineKeyboardButtonTypeCallbackGameTemp := InlineKeyboardButtonTypeCallbackGame{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeCallbackGame"},
	}

	return &inlineKeyboardButtonTypeCallbackGameTemp
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

// NewInlineKeyboardButtonTypeSwitchInline creates a new InlineKeyboardButtonTypeSwitchInline
//
// @param query Inline query to be sent to the bot
// @param inCurrentChat True, if the inline query should be sent from the current chat
func NewInlineKeyboardButtonTypeSwitchInline(query string, inCurrentChat bool) *InlineKeyboardButtonTypeSwitchInline {
	inlineKeyboardButtonTypeSwitchInlineTemp := InlineKeyboardButtonTypeSwitchInline{
		tdCommon:      tdCommon{Type: "inlineKeyboardButtonTypeSwitchInline"},
		Query:         query,
		InCurrentChat: inCurrentChat,
	}

	return &inlineKeyboardButtonTypeSwitchInlineTemp
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

// NewInlineKeyboardButtonTypeBuy creates a new InlineKeyboardButtonTypeBuy
//
func NewInlineKeyboardButtonTypeBuy() *InlineKeyboardButtonTypeBuy {
	inlineKeyboardButtonTypeBuyTemp := InlineKeyboardButtonTypeBuy{
		tdCommon: tdCommon{Type: "inlineKeyboardButtonTypeBuy"},
	}

	return &inlineKeyboardButtonTypeBuyTemp
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

// NewInlineKeyboardButton creates a new InlineKeyboardButton
//
// @param text Text of the button
// @param typeParam Type of the button
func NewInlineKeyboardButton(text string, typeParam InlineKeyboardButtonType) *InlineKeyboardButton {
	inlineKeyboardButtonTemp := InlineKeyboardButton{
		tdCommon: tdCommon{Type: "inlineKeyboardButton"},
		Text:     text,
		Type:     typeParam,
	}

	return &inlineKeyboardButtonTemp
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

// NewReplyMarkupRemoveKeyboard creates a new ReplyMarkupRemoveKeyboard
//
// @param isPersonal True, if the keyboard is removed only for the mentioned users or the target user of a reply
func NewReplyMarkupRemoveKeyboard(isPersonal bool) *ReplyMarkupRemoveKeyboard {
	replyMarkupRemoveKeyboardTemp := ReplyMarkupRemoveKeyboard{
		tdCommon:   tdCommon{Type: "replyMarkupRemoveKeyboard"},
		IsPersonal: isPersonal,
	}

	return &replyMarkupRemoveKeyboardTemp
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

// NewReplyMarkupForceReply creates a new ReplyMarkupForceReply
//
// @param isPersonal True, if a forced reply must automatically be shown to the current user. For outgoing messages, specify true to show the forced reply only for the mentioned users and for the target user of a reply
func NewReplyMarkupForceReply(isPersonal bool) *ReplyMarkupForceReply {
	replyMarkupForceReplyTemp := ReplyMarkupForceReply{
		tdCommon:   tdCommon{Type: "replyMarkupForceReply"},
		IsPersonal: isPersonal,
	}

	return &replyMarkupForceReplyTemp
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

// NewReplyMarkupShowKeyboard creates a new ReplyMarkupShowKeyboard
//
// @param rows A list of rows of bot keyboard buttons
// @param resizeKeyboard True, if the client needs to resize the keyboard vertically
// @param oneTime True, if the client needs to hide the keyboard after use
// @param isPersonal True, if the keyboard must automatically be shown to the current user. For outgoing messages, specify true to show the keyboard only for the mentioned users and for the target user of a reply
func NewReplyMarkupShowKeyboard(rows [][]KeyboardButton, resizeKeyboard bool, oneTime bool, isPersonal bool) *ReplyMarkupShowKeyboard {
	replyMarkupShowKeyboardTemp := ReplyMarkupShowKeyboard{
		tdCommon:       tdCommon{Type: "replyMarkupShowKeyboard"},
		Rows:           rows,
		ResizeKeyboard: resizeKeyboard,
		OneTime:        oneTime,
		IsPersonal:     isPersonal,
	}

	return &replyMarkupShowKeyboardTemp
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

// NewReplyMarkupInlineKeyboard creates a new ReplyMarkupInlineKeyboard
//
// @param rows A list of rows of inline keyboard buttons
func NewReplyMarkupInlineKeyboard(rows [][]InlineKeyboardButton) *ReplyMarkupInlineKeyboard {
	replyMarkupInlineKeyboardTemp := ReplyMarkupInlineKeyboard{
		tdCommon: tdCommon{Type: "replyMarkupInlineKeyboard"},
		Rows:     rows,
	}

	return &replyMarkupInlineKeyboardTemp
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

// NewRichTextPlain creates a new RichTextPlain
//
// @param text Text
func NewRichTextPlain(text string) *RichTextPlain {
	richTextPlainTemp := RichTextPlain{
		tdCommon: tdCommon{Type: "richTextPlain"},
		Text:     text,
	}

	return &richTextPlainTemp
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

// NewRichTextBold creates a new RichTextBold
//
// @param text Text
func NewRichTextBold(text RichText) *RichTextBold {
	richTextBoldTemp := RichTextBold{
		tdCommon: tdCommon{Type: "richTextBold"},
		Text:     text,
	}

	return &richTextBoldTemp
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

// NewRichTextItalic creates a new RichTextItalic
//
// @param text Text
func NewRichTextItalic(text RichText) *RichTextItalic {
	richTextItalicTemp := RichTextItalic{
		tdCommon: tdCommon{Type: "richTextItalic"},
		Text:     text,
	}

	return &richTextItalicTemp
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

// NewRichTextUnderline creates a new RichTextUnderline
//
// @param text Text
func NewRichTextUnderline(text RichText) *RichTextUnderline {
	richTextUnderlineTemp := RichTextUnderline{
		tdCommon: tdCommon{Type: "richTextUnderline"},
		Text:     text,
	}

	return &richTextUnderlineTemp
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

// NewRichTextStrikethrough creates a new RichTextStrikethrough
//
// @param text Text
func NewRichTextStrikethrough(text RichText) *RichTextStrikethrough {
	richTextStrikethroughTemp := RichTextStrikethrough{
		tdCommon: tdCommon{Type: "richTextStrikethrough"},
		Text:     text,
	}

	return &richTextStrikethroughTemp
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

// NewRichTextFixed creates a new RichTextFixed
//
// @param text Text
func NewRichTextFixed(text RichText) *RichTextFixed {
	richTextFixedTemp := RichTextFixed{
		tdCommon: tdCommon{Type: "richTextFixed"},
		Text:     text,
	}

	return &richTextFixedTemp
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

// NewRichTextURL creates a new RichTextURL
//
// @param text Text
// @param uRL URL
func NewRichTextURL(text RichText, uRL string) *RichTextURL {
	richTextURLTemp := RichTextURL{
		tdCommon: tdCommon{Type: "richTextUrl"},
		Text:     text,
		URL:      uRL,
	}

	return &richTextURLTemp
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

// NewRichTextEmailAddress creates a new RichTextEmailAddress
//
// @param text Text
// @param emailAddress Email address
func NewRichTextEmailAddress(text RichText, emailAddress string) *RichTextEmailAddress {
	richTextEmailAddressTemp := RichTextEmailAddress{
		tdCommon:     tdCommon{Type: "richTextEmailAddress"},
		Text:         text,
		EmailAddress: emailAddress,
	}

	return &richTextEmailAddressTemp
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

// NewRichTexts creates a new RichTexts
//
// @param texts Texts
func NewRichTexts(texts []RichText) *RichTexts {
	richTextsTemp := RichTexts{
		tdCommon: tdCommon{Type: "richTexts"},
		Texts:    texts,
	}

	return &richTextsTemp
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

// NewPageBlockTitle creates a new PageBlockTitle
//
// @param title Title
func NewPageBlockTitle(title RichText) *PageBlockTitle {
	pageBlockTitleTemp := PageBlockTitle{
		tdCommon: tdCommon{Type: "pageBlockTitle"},
		Title:    title,
	}

	return &pageBlockTitleTemp
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

// NewPageBlockSubtitle creates a new PageBlockSubtitle
//
// @param subtitle Subtitle
func NewPageBlockSubtitle(subtitle RichText) *PageBlockSubtitle {
	pageBlockSubtitleTemp := PageBlockSubtitle{
		tdCommon: tdCommon{Type: "pageBlockSubtitle"},
		Subtitle: subtitle,
	}

	return &pageBlockSubtitleTemp
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

// NewPageBlockAuthorDate creates a new PageBlockAuthorDate
//
// @param author Author
// @param publishDate Point in time (Unix timestamp) when the article was published; 0 if unknown
func NewPageBlockAuthorDate(author RichText, publishDate int32) *PageBlockAuthorDate {
	pageBlockAuthorDateTemp := PageBlockAuthorDate{
		tdCommon:    tdCommon{Type: "pageBlockAuthorDate"},
		Author:      author,
		PublishDate: publishDate,
	}

	return &pageBlockAuthorDateTemp
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

// NewPageBlockHeader creates a new PageBlockHeader
//
// @param header Header
func NewPageBlockHeader(header RichText) *PageBlockHeader {
	pageBlockHeaderTemp := PageBlockHeader{
		tdCommon: tdCommon{Type: "pageBlockHeader"},
		Header:   header,
	}

	return &pageBlockHeaderTemp
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

// NewPageBlockSubheader creates a new PageBlockSubheader
//
// @param subheader Subheader
func NewPageBlockSubheader(subheader RichText) *PageBlockSubheader {
	pageBlockSubheaderTemp := PageBlockSubheader{
		tdCommon:  tdCommon{Type: "pageBlockSubheader"},
		Subheader: subheader,
	}

	return &pageBlockSubheaderTemp
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

// NewPageBlockParagraph creates a new PageBlockParagraph
//
// @param text Paragraph text
func NewPageBlockParagraph(text RichText) *PageBlockParagraph {
	pageBlockParagraphTemp := PageBlockParagraph{
		tdCommon: tdCommon{Type: "pageBlockParagraph"},
		Text:     text,
	}

	return &pageBlockParagraphTemp
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

// NewPageBlockPreformatted creates a new PageBlockPreformatted
//
// @param text Paragraph text
// @param language Programming language for which the text should be formatted
func NewPageBlockPreformatted(text RichText, language string) *PageBlockPreformatted {
	pageBlockPreformattedTemp := PageBlockPreformatted{
		tdCommon: tdCommon{Type: "pageBlockPreformatted"},
		Text:     text,
		Language: language,
	}

	return &pageBlockPreformattedTemp
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

// NewPageBlockFooter creates a new PageBlockFooter
//
// @param footer Footer
func NewPageBlockFooter(footer RichText) *PageBlockFooter {
	pageBlockFooterTemp := PageBlockFooter{
		tdCommon: tdCommon{Type: "pageBlockFooter"},
		Footer:   footer,
	}

	return &pageBlockFooterTemp
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

// NewPageBlockDivider creates a new PageBlockDivider
//
func NewPageBlockDivider() *PageBlockDivider {
	pageBlockDividerTemp := PageBlockDivider{
		tdCommon: tdCommon{Type: "pageBlockDivider"},
	}

	return &pageBlockDividerTemp
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

// NewPageBlockAnchor creates a new PageBlockAnchor
//
// @param name Name of the anchor
func NewPageBlockAnchor(name string) *PageBlockAnchor {
	pageBlockAnchorTemp := PageBlockAnchor{
		tdCommon: tdCommon{Type: "pageBlockAnchor"},
		Name:     name,
	}

	return &pageBlockAnchorTemp
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

// NewPageBlockList creates a new PageBlockList
//
// @param items Texts
// @param isOrdered True, if the items should be marked with numbers
func NewPageBlockList(items []RichText, isOrdered bool) *PageBlockList {
	pageBlockListTemp := PageBlockList{
		tdCommon:  tdCommon{Type: "pageBlockList"},
		Items:     items,
		IsOrdered: isOrdered,
	}

	return &pageBlockListTemp
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

// NewPageBlockBlockQuote creates a new PageBlockBlockQuote
//
// @param text Quote text
// @param caption Quote caption
func NewPageBlockBlockQuote(text RichText, caption RichText) *PageBlockBlockQuote {
	pageBlockBlockQuoteTemp := PageBlockBlockQuote{
		tdCommon: tdCommon{Type: "pageBlockBlockQuote"},
		Text:     text,
		Caption:  caption,
	}

	return &pageBlockBlockQuoteTemp
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

// NewPageBlockPullQuote creates a new PageBlockPullQuote
//
// @param text Quote text
// @param caption Quote caption
func NewPageBlockPullQuote(text RichText, caption RichText) *PageBlockPullQuote {
	pageBlockPullQuoteTemp := PageBlockPullQuote{
		tdCommon: tdCommon{Type: "pageBlockPullQuote"},
		Text:     text,
		Caption:  caption,
	}

	return &pageBlockPullQuoteTemp
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
	Animation    *Animation `json:"animation"`     // Animation file; may be null
	Caption      RichText   `json:"caption"`       // Animation caption
	NeedAutoplay bool       `json:"need_autoplay"` // True, if the animation should be played automatically
}

// MessageType return the string telegram-type of PageBlockAnimation
func (pageBlockAnimation *PageBlockAnimation) MessageType() string {
	return "pageBlockAnimation"
}

// NewPageBlockAnimation creates a new PageBlockAnimation
//
// @param animation Animation file; may be null
// @param caption Animation caption
// @param needAutoplay True, if the animation should be played automatically
func NewPageBlockAnimation(animation *Animation, caption RichText, needAutoplay bool) *PageBlockAnimation {
	pageBlockAnimationTemp := PageBlockAnimation{
		tdCommon:     tdCommon{Type: "pageBlockAnimation"},
		Animation:    animation,
		Caption:      caption,
		NeedAutoplay: needAutoplay,
	}

	return &pageBlockAnimationTemp
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
		Animation    *Animation `json:"animation"`     // Animation file; may be null
		NeedAutoplay bool       `json:"need_autoplay"` // True, if the animation should be played automatically
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
	Audio   *Audio   `json:"audio"`   // Audio file; may be null
	Caption RichText `json:"caption"` // Audio file caption
}

// MessageType return the string telegram-type of PageBlockAudio
func (pageBlockAudio *PageBlockAudio) MessageType() string {
	return "pageBlockAudio"
}

// NewPageBlockAudio creates a new PageBlockAudio
//
// @param audio Audio file; may be null
// @param caption Audio file caption
func NewPageBlockAudio(audio *Audio, caption RichText) *PageBlockAudio {
	pageBlockAudioTemp := PageBlockAudio{
		tdCommon: tdCommon{Type: "pageBlockAudio"},
		Audio:    audio,
		Caption:  caption,
	}

	return &pageBlockAudioTemp
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
		Audio *Audio `json:"audio"` // Audio file; may be null

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
	Photo   *Photo   `json:"photo"`   // Photo file; may be null
	Caption RichText `json:"caption"` // Photo caption
}

// MessageType return the string telegram-type of PageBlockPhoto
func (pageBlockPhoto *PageBlockPhoto) MessageType() string {
	return "pageBlockPhoto"
}

// NewPageBlockPhoto creates a new PageBlockPhoto
//
// @param photo Photo file; may be null
// @param caption Photo caption
func NewPageBlockPhoto(photo *Photo, caption RichText) *PageBlockPhoto {
	pageBlockPhotoTemp := PageBlockPhoto{
		tdCommon: tdCommon{Type: "pageBlockPhoto"},
		Photo:    photo,
		Caption:  caption,
	}

	return &pageBlockPhotoTemp
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
		Photo *Photo `json:"photo"` // Photo file; may be null

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
	Video        *Video   `json:"video"`         // Video file; may be null
	Caption      RichText `json:"caption"`       // Video caption
	NeedAutoplay bool     `json:"need_autoplay"` // True, if the video should be played automatically
	IsLooped     bool     `json:"is_looped"`     // True, if the video should be looped
}

// MessageType return the string telegram-type of PageBlockVideo
func (pageBlockVideo *PageBlockVideo) MessageType() string {
	return "pageBlockVideo"
}

// NewPageBlockVideo creates a new PageBlockVideo
//
// @param video Video file; may be null
// @param caption Video caption
// @param needAutoplay True, if the video should be played automatically
// @param isLooped True, if the video should be looped
func NewPageBlockVideo(video *Video, caption RichText, needAutoplay bool, isLooped bool) *PageBlockVideo {
	pageBlockVideoTemp := PageBlockVideo{
		tdCommon:     tdCommon{Type: "pageBlockVideo"},
		Video:        video,
		Caption:      caption,
		NeedAutoplay: needAutoplay,
		IsLooped:     isLooped,
	}

	return &pageBlockVideoTemp
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
		Video        *Video `json:"video"`         // Video file; may be null
		NeedAutoplay bool   `json:"need_autoplay"` // True, if the video should be played automatically
		IsLooped     bool   `json:"is_looped"`     // True, if the video should be looped
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

// NewPageBlockCover creates a new PageBlockCover
//
// @param cover Cover
func NewPageBlockCover(cover PageBlock) *PageBlockCover {
	pageBlockCoverTemp := PageBlockCover{
		tdCommon: tdCommon{Type: "pageBlockCover"},
		Cover:    cover,
	}

	return &pageBlockCoverTemp
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
	PosterPhoto    *Photo   `json:"poster_photo"`    // Poster photo, if available; may be null
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

// NewPageBlockEmbedded creates a new PageBlockEmbedded
//
// @param uRL Web page URL, if available
// @param hTML HTML-markup of the embedded page
// @param posterPhoto Poster photo, if available; may be null
// @param width Block width
// @param height Block height
// @param caption Block caption
// @param isFullWidth True, if the block should be full width
// @param allowScrolling True, if scrolling should be allowed
func NewPageBlockEmbedded(uRL string, hTML string, posterPhoto *Photo, width int32, height int32, caption RichText, isFullWidth bool, allowScrolling bool) *PageBlockEmbedded {
	pageBlockEmbeddedTemp := PageBlockEmbedded{
		tdCommon:       tdCommon{Type: "pageBlockEmbedded"},
		URL:            uRL,
		HTML:           hTML,
		PosterPhoto:    posterPhoto,
		Width:          width,
		Height:         height,
		Caption:        caption,
		IsFullWidth:    isFullWidth,
		AllowScrolling: allowScrolling,
	}

	return &pageBlockEmbeddedTemp
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
		PosterPhoto    *Photo `json:"poster_photo"`    // Poster photo, if available; may be null
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
	AuthorPhoto *Photo      `json:"author_photo"` // Post author photo
	Date        int32       `json:"date"`         // Point in time (Unix timestamp) when the post was created; 0 if unknown
	PageBlocks  []PageBlock `json:"page_blocks"`  // Post content
	Caption     RichText    `json:"caption"`      // Post caption
}

// MessageType return the string telegram-type of PageBlockEmbeddedPost
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) MessageType() string {
	return "pageBlockEmbeddedPost"
}

// NewPageBlockEmbeddedPost creates a new PageBlockEmbeddedPost
//
// @param uRL Web page URL
// @param author Post author
// @param authorPhoto Post author photo
// @param date Point in time (Unix timestamp) when the post was created; 0 if unknown
// @param pageBlocks Post content
// @param caption Post caption
func NewPageBlockEmbeddedPost(uRL string, author string, authorPhoto *Photo, date int32, pageBlocks []PageBlock, caption RichText) *PageBlockEmbeddedPost {
	pageBlockEmbeddedPostTemp := PageBlockEmbeddedPost{
		tdCommon:    tdCommon{Type: "pageBlockEmbeddedPost"},
		URL:         uRL,
		Author:      author,
		AuthorPhoto: authorPhoto,
		Date:        date,
		PageBlocks:  pageBlocks,
		Caption:     caption,
	}

	return &pageBlockEmbeddedPostTemp
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
		AuthorPhoto *Photo      `json:"author_photo"` // Post author photo
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

// NewPageBlockCollage creates a new PageBlockCollage
//
// @param pageBlocks Collage item contents
// @param caption Block caption
func NewPageBlockCollage(pageBlocks []PageBlock, caption RichText) *PageBlockCollage {
	pageBlockCollageTemp := PageBlockCollage{
		tdCommon:   tdCommon{Type: "pageBlockCollage"},
		PageBlocks: pageBlocks,
		Caption:    caption,
	}

	return &pageBlockCollageTemp
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

// NewPageBlockSlideshow creates a new PageBlockSlideshow
//
// @param pageBlocks Slideshow item contents
// @param caption Block caption
func NewPageBlockSlideshow(pageBlocks []PageBlock, caption RichText) *PageBlockSlideshow {
	pageBlockSlideshowTemp := PageBlockSlideshow{
		tdCommon:   tdCommon{Type: "pageBlockSlideshow"},
		PageBlocks: pageBlocks,
		Caption:    caption,
	}

	return &pageBlockSlideshowTemp
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
	Title    string     `json:"title"`    // Chat title
	Photo    *ChatPhoto `json:"photo"`    // Chat photo; may be null
	Username string     `json:"username"` // Chat username, by which all other information about the chat should be resolved
}

// MessageType return the string telegram-type of PageBlockChatLink
func (pageBlockChatLink *PageBlockChatLink) MessageType() string {
	return "pageBlockChatLink"
}

// NewPageBlockChatLink creates a new PageBlockChatLink
//
// @param title Chat title
// @param photo Chat photo; may be null
// @param username Chat username, by which all other information about the chat should be resolved
func NewPageBlockChatLink(title string, photo *ChatPhoto, username string) *PageBlockChatLink {
	pageBlockChatLinkTemp := PageBlockChatLink{
		tdCommon: tdCommon{Type: "pageBlockChatLink"},
		Title:    title,
		Photo:    photo,
		Username: username,
	}

	return &pageBlockChatLinkTemp
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

// NewWebPageInstantView creates a new WebPageInstantView
//
// @param pageBlocks Content of the web page
// @param isFull True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
func NewWebPageInstantView(pageBlocks []PageBlock, isFull bool) *WebPageInstantView {
	webPageInstantViewTemp := WebPageInstantView{
		tdCommon:   tdCommon{Type: "webPageInstantView"},
		PageBlocks: pageBlocks,
		IsFull:     isFull,
	}

	return &webPageInstantViewTemp
}

// WebPage Describes a web page preview
type WebPage struct {
	tdCommon
	URL            string     `json:"url"`              // Original URL of the link
	DisplayURL     string     `json:"display_url"`      // URL to display
	Type           string     `json:"type"`             // Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
	SiteName       string     `json:"site_name"`        // Short name of the site (e.g., Google Docs, App Store)
	Title          string     `json:"title"`            // Title of the content
	Description    string     `json:"description"`      //
	Photo          *Photo     `json:"photo"`            // Image representing the content; may be null
	EmbedURL       string     `json:"embed_url"`        // URL to show in the embedded preview
	EmbedType      string     `json:"embed_type"`       // MIME type of the embedded preview, (e.g., text/html or video/mp4)
	EmbedWidth     int32      `json:"embed_width"`      // Width of the embedded preview
	EmbedHeight    int32      `json:"embed_height"`     // Height of the embedded preview
	Duration       int32      `json:"duration"`         // Duration of the content, in seconds
	Author         string     `json:"author"`           // Author of the content
	Animation      *Animation `json:"animation"`        // Preview of the content as an animation, if available; may be null
	Audio          *Audio     `json:"audio"`            // Preview of the content as an audio file, if available; may be null
	Document       *Document  `json:"document"`         // Preview of the content as a document, if available (currently only available for small PDF files and ZIP archives); may be null
	Sticker        *Sticker   `json:"sticker"`          // Preview of the content as a sticker for small WEBP files, if available; may be null
	Video          *Video     `json:"video"`            // Preview of the content as a video, if available; may be null
	VideoNote      *VideoNote `json:"video_note"`       // Preview of the content as a video note, if available; may be null
	VoiceNote      *VoiceNote `json:"voice_note"`       // Preview of the content as a voice note, if available; may be null
	HasInstantView bool       `json:"has_instant_view"` // True, if the web page has an instant view
}

// MessageType return the string telegram-type of WebPage
func (webPage *WebPage) MessageType() string {
	return "webPage"
}

// NewWebPage creates a new WebPage
//
// @param uRL Original URL of the link
// @param displayURL URL to display
// @param typeParam Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
// @param siteName Short name of the site (e.g., Google Docs, App Store)
// @param title Title of the content
// @param description
// @param photo Image representing the content; may be null
// @param embedURL URL to show in the embedded preview
// @param embedType MIME type of the embedded preview, (e.g., text/html or video/mp4)
// @param embedWidth Width of the embedded preview
// @param embedHeight Height of the embedded preview
// @param duration Duration of the content, in seconds
// @param author Author of the content
// @param animation Preview of the content as an animation, if available; may be null
// @param audio Preview of the content as an audio file, if available; may be null
// @param document Preview of the content as a document, if available (currently only available for small PDF files and ZIP archives); may be null
// @param sticker Preview of the content as a sticker for small WEBP files, if available; may be null
// @param video Preview of the content as a video, if available; may be null
// @param videoNote Preview of the content as a video note, if available; may be null
// @param voiceNote Preview of the content as a voice note, if available; may be null
// @param hasInstantView True, if the web page has an instant view
func NewWebPage(uRL string, displayURL string, typeParam string, siteName string, title string, description string, photo *Photo, embedURL string, embedType string, embedWidth int32, embedHeight int32, duration int32, author string, animation *Animation, audio *Audio, document *Document, sticker *Sticker, video *Video, videoNote *VideoNote, voiceNote *VoiceNote, hasInstantView bool) *WebPage {
	webPageTemp := WebPage{
		tdCommon:       tdCommon{Type: "webPage"},
		URL:            uRL,
		DisplayURL:     displayURL,
		Type:           typeParam,
		SiteName:       siteName,
		Title:          title,
		Description:    description,
		Photo:          photo,
		EmbedURL:       embedURL,
		EmbedType:      embedType,
		EmbedWidth:     embedWidth,
		EmbedHeight:    embedHeight,
		Duration:       duration,
		Author:         author,
		Animation:      animation,
		Audio:          audio,
		Document:       document,
		Sticker:        sticker,
		Video:          video,
		VideoNote:      videoNote,
		VoiceNote:      voiceNote,
		HasInstantView: hasInstantView,
	}

	return &webPageTemp
}

// Address Describes an address
type Address struct {
	tdCommon
	CountryCode string `json:"country_code"` // A two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line of the address
	StreetLine2 string `json:"street_line2"` // Second line of the address
	PostalCode  string `json:"postal_code"`  // Address postal code
}

// MessageType return the string telegram-type of Address
func (address *Address) MessageType() string {
	return "address"
}

// NewAddress creates a new Address
//
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
// @param state State, if applicable
// @param city City
// @param streetLine1 First line of the address
// @param streetLine2 Second line of the address
// @param postalCode Address postal code
func NewAddress(countryCode string, state string, city string, streetLine1 string, streetLine2 string, postalCode string) *Address {
	addressTemp := Address{
		tdCommon:    tdCommon{Type: "address"},
		CountryCode: countryCode,
		State:       state,
		City:        city,
		StreetLine1: streetLine1,
		StreetLine2: streetLine2,
		PostalCode:  postalCode,
	}

	return &addressTemp
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

// NewLabeledPricePart creates a new LabeledPricePart
//
// @param label Label for this portion of the product price
// @param amount Currency amount in minimal quantity of the currency
func NewLabeledPricePart(label string, amount int64) *LabeledPricePart {
	labeledPricePartTemp := LabeledPricePart{
		tdCommon: tdCommon{Type: "labeledPricePart"},
		Label:    label,
		Amount:   amount,
	}

	return &labeledPricePartTemp
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

// NewInvoice creates a new Invoice
//
// @param currency ISO 4217 currency code
// @param priceParts A list of objects used to calculate the total price of the product
// @param isTest True, if the payment is a test payment
// @param needName True, if the user's name is needed for payment
// @param needPhoneNumber True, if the user's phone number is needed for payment
// @param needEmailAddress True, if the user's email address is needed for payment
// @param needShippingAddress True, if the user's shipping address is needed for payment
// @param sendPhoneNumberToProvider True, if the user's phone number will be sent to the provider
// @param sendEmailAddressToProvider True, if the user's email address will be sent to the provider
// @param isFlexible True, if the total price depends on the shipping method
func NewInvoice(currency string, priceParts []LabeledPricePart, isTest bool, needName bool, needPhoneNumber bool, needEmailAddress bool, needShippingAddress bool, sendPhoneNumberToProvider bool, sendEmailAddressToProvider bool, isFlexible bool) *Invoice {
	invoiceTemp := Invoice{
		tdCommon:                   tdCommon{Type: "invoice"},
		Currency:                   currency,
		PriceParts:                 priceParts,
		IsTest:                     isTest,
		NeedName:                   needName,
		NeedPhoneNumber:            needPhoneNumber,
		NeedEmailAddress:           needEmailAddress,
		NeedShippingAddress:        needShippingAddress,
		SendPhoneNumberToProvider:  sendPhoneNumberToProvider,
		SendEmailAddressToProvider: sendEmailAddressToProvider,
		IsFlexible:                 isFlexible,
	}

	return &invoiceTemp
}

// OrderInfo Order information
type OrderInfo struct {
	tdCommon
	Name            string   `json:"name"`             // Name of the user
	PhoneNumber     string   `json:"phone_number"`     // Phone number of the user
	EmailAddress    string   `json:"email_address"`    // Email address of the user
	ShippingAddress *Address `json:"shipping_address"` // Shipping address for this order; may be null
}

// MessageType return the string telegram-type of OrderInfo
func (orderInfo *OrderInfo) MessageType() string {
	return "orderInfo"
}

// NewOrderInfo creates a new OrderInfo
//
// @param name Name of the user
// @param phoneNumber Phone number of the user
// @param emailAddress Email address of the user
// @param shippingAddress Shipping address for this order; may be null
func NewOrderInfo(name string, phoneNumber string, emailAddress string, shippingAddress *Address) *OrderInfo {
	orderInfoTemp := OrderInfo{
		tdCommon:        tdCommon{Type: "orderInfo"},
		Name:            name,
		PhoneNumber:     phoneNumber,
		EmailAddress:    emailAddress,
		ShippingAddress: shippingAddress,
	}

	return &orderInfoTemp
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

// NewShippingOption creates a new ShippingOption
//
// @param iD Shipping option identifier
// @param title Option title
// @param priceParts A list of objects used to calculate the total shipping costs
func NewShippingOption(iD string, title string, priceParts []LabeledPricePart) *ShippingOption {
	shippingOptionTemp := ShippingOption{
		tdCommon:   tdCommon{Type: "shippingOption"},
		ID:         iD,
		Title:      title,
		PriceParts: priceParts,
	}

	return &shippingOptionTemp
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

// NewSavedCredentials creates a new SavedCredentials
//
// @param iD Unique identifier of the saved credentials
// @param title Title of the saved credentials
func NewSavedCredentials(iD string, title string) *SavedCredentials {
	savedCredentialsTemp := SavedCredentials{
		tdCommon: tdCommon{Type: "savedCredentials"},
		ID:       iD,
		Title:    title,
	}

	return &savedCredentialsTemp
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

// NewInputCredentialsSaved creates a new InputCredentialsSaved
//
// @param savedCredentialsID Identifier of the saved credentials
func NewInputCredentialsSaved(savedCredentialsID string) *InputCredentialsSaved {
	inputCredentialsSavedTemp := InputCredentialsSaved{
		tdCommon:           tdCommon{Type: "inputCredentialsSaved"},
		SavedCredentialsID: savedCredentialsID,
	}

	return &inputCredentialsSavedTemp
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

// NewInputCredentialsNew creates a new InputCredentialsNew
//
// @param data Contains JSON-encoded data with a credential identifier from the payment provider
// @param allowSave True, if the credential identifier can be saved on the server side
func NewInputCredentialsNew(data string, allowSave bool) *InputCredentialsNew {
	inputCredentialsNewTemp := InputCredentialsNew{
		tdCommon:  tdCommon{Type: "inputCredentialsNew"},
		Data:      data,
		AllowSave: allowSave,
	}

	return &inputCredentialsNewTemp
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

// NewInputCredentialsAndroidPay creates a new InputCredentialsAndroidPay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsAndroidPay(data string) *InputCredentialsAndroidPay {
	inputCredentialsAndroidPayTemp := InputCredentialsAndroidPay{
		tdCommon: tdCommon{Type: "inputCredentialsAndroidPay"},
		Data:     data,
	}

	return &inputCredentialsAndroidPayTemp
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

// NewInputCredentialsApplePay creates a new InputCredentialsApplePay
//
// @param data JSON-encoded data with the credential identifier
func NewInputCredentialsApplePay(data string) *InputCredentialsApplePay {
	inputCredentialsApplePayTemp := InputCredentialsApplePay{
		tdCommon: tdCommon{Type: "inputCredentialsApplePay"},
		Data:     data,
	}

	return &inputCredentialsApplePayTemp
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

// NewPaymentsProviderStripe creates a new PaymentsProviderStripe
//
// @param publishableKey Stripe API publishable key
// @param needCountry True, if the user country must be provided
// @param needPostalCode True, if the user ZIP/postal code must be provided
// @param needCardholderName True, if the cardholder name must be provided
func NewPaymentsProviderStripe(publishableKey string, needCountry bool, needPostalCode bool, needCardholderName bool) *PaymentsProviderStripe {
	paymentsProviderStripeTemp := PaymentsProviderStripe{
		tdCommon:           tdCommon{Type: "paymentsProviderStripe"},
		PublishableKey:     publishableKey,
		NeedCountry:        needCountry,
		NeedPostalCode:     needPostalCode,
		NeedCardholderName: needCardholderName,
	}

	return &paymentsProviderStripeTemp
}

// PaymentForm Contains information about an invoice payment form
type PaymentForm struct {
	tdCommon
	Invoice            *Invoice                `json:"invoice"`              // Full information of the invoice
	URL                string                  `json:"url"`                  // Payment form URL
	PaymentsProvider   *PaymentsProviderStripe `json:"payments_provider"`    // Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
	SavedOrderInfo     *OrderInfo              `json:"saved_order_info"`     // Saved server-side order information; may be null
	SavedCredentials   *SavedCredentials       `json:"saved_credentials"`    // Contains information about saved card credentials; may be null
	CanSaveCredentials bool                    `json:"can_save_credentials"` // True, if the user can choose to save credentials
	NeedPassword       bool                    `json:"need_password"`        // True, if the user will be able to save credentials protected by a password they set up
}

// MessageType return the string telegram-type of PaymentForm
func (paymentForm *PaymentForm) MessageType() string {
	return "paymentForm"
}

// NewPaymentForm creates a new PaymentForm
//
// @param invoice Full information of the invoice
// @param uRL Payment form URL
// @param paymentsProvider Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
// @param savedOrderInfo Saved server-side order information; may be null
// @param savedCredentials Contains information about saved card credentials; may be null
// @param canSaveCredentials True, if the user can choose to save credentials
// @param needPassword True, if the user will be able to save credentials protected by a password they set up
func NewPaymentForm(invoice *Invoice, uRL string, paymentsProvider *PaymentsProviderStripe, savedOrderInfo *OrderInfo, savedCredentials *SavedCredentials, canSaveCredentials bool, needPassword bool) *PaymentForm {
	paymentFormTemp := PaymentForm{
		tdCommon:           tdCommon{Type: "paymentForm"},
		Invoice:            invoice,
		URL:                uRL,
		PaymentsProvider:   paymentsProvider,
		SavedOrderInfo:     savedOrderInfo,
		SavedCredentials:   savedCredentials,
		CanSaveCredentials: canSaveCredentials,
		NeedPassword:       needPassword,
	}

	return &paymentFormTemp
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

// NewValidatedOrderInfo creates a new ValidatedOrderInfo
//
// @param orderInfoID Temporary identifier of the order information
// @param shippingOptions Available shipping options
func NewValidatedOrderInfo(orderInfoID string, shippingOptions []ShippingOption) *ValidatedOrderInfo {
	validatedOrderInfoTemp := ValidatedOrderInfo{
		tdCommon:        tdCommon{Type: "validatedOrderInfo"},
		OrderInfoID:     orderInfoID,
		ShippingOptions: shippingOptions,
	}

	return &validatedOrderInfoTemp
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

// NewPaymentResult creates a new PaymentResult
//
// @param success True, if the payment request was successful; otherwise the verification_url will be not empty
// @param verificationURL URL for additional payment credentials verification
func NewPaymentResult(success bool, verificationURL string) *PaymentResult {
	paymentResultTemp := PaymentResult{
		tdCommon:        tdCommon{Type: "paymentResult"},
		Success:         success,
		VerificationURL: verificationURL,
	}

	return &paymentResultTemp
}

// PaymentReceipt Contains information about a successful payment
type PaymentReceipt struct {
	tdCommon
	Date                   int32           `json:"date"`                      // Point in time (Unix timestamp) when the payment was made
	PaymentsProviderUserID int32           `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	Invoice                *Invoice        `json:"invoice"`                   // Contains information about the invoice
	OrderInfo              *OrderInfo      `json:"order_info"`                // Contains order information; may be null
	ShippingOption         *ShippingOption `json:"shipping_option"`           // Chosen shipping option; may be null
	CredentialsTitle       string          `json:"credentials_title"`         // Title of the saved credentials
}

// MessageType return the string telegram-type of PaymentReceipt
func (paymentReceipt *PaymentReceipt) MessageType() string {
	return "paymentReceipt"
}

// NewPaymentReceipt creates a new PaymentReceipt
//
// @param date Point in time (Unix timestamp) when the payment was made
// @param paymentsProviderUserID User identifier of the payment provider bot
// @param invoice Contains information about the invoice
// @param orderInfo Contains order information; may be null
// @param shippingOption Chosen shipping option; may be null
// @param credentialsTitle Title of the saved credentials
func NewPaymentReceipt(date int32, paymentsProviderUserID int32, invoice *Invoice, orderInfo *OrderInfo, shippingOption *ShippingOption, credentialsTitle string) *PaymentReceipt {
	paymentReceiptTemp := PaymentReceipt{
		tdCommon:               tdCommon{Type: "paymentReceipt"},
		Date:                   date,
		PaymentsProviderUserID: paymentsProviderUserID,
		Invoice:                invoice,
		OrderInfo:              orderInfo,
		ShippingOption:         shippingOption,
		CredentialsTitle:       credentialsTitle,
	}

	return &paymentReceiptTemp
}

// DatedFile File with the date it was uploaded
type DatedFile struct {
	tdCommon
	File *File `json:"file"` // The file
	Date int32 `json:"date"` // Point in time (Unix timestamp) when the file was uploaded
}

// MessageType return the string telegram-type of DatedFile
func (datedFile *DatedFile) MessageType() string {
	return "datedFile"
}

// NewDatedFile creates a new DatedFile
//
// @param file The file
// @param date Point in time (Unix timestamp) when the file was uploaded
func NewDatedFile(file *File, date int32) *DatedFile {
	datedFileTemp := DatedFile{
		tdCommon: tdCommon{Type: "datedFile"},
		File:     file,
		Date:     date,
	}

	return &datedFileTemp
}

// PassportElementTypePersonalDetails A Telegram Passport element containing the user's personal details
type PassportElementTypePersonalDetails struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePersonalDetails
func (passportElementTypePersonalDetails *PassportElementTypePersonalDetails) MessageType() string {
	return "passportElementTypePersonalDetails"
}

// NewPassportElementTypePersonalDetails creates a new PassportElementTypePersonalDetails
//
func NewPassportElementTypePersonalDetails() *PassportElementTypePersonalDetails {
	passportElementTypePersonalDetailsTemp := PassportElementTypePersonalDetails{
		tdCommon: tdCommon{Type: "passportElementTypePersonalDetails"},
	}

	return &passportElementTypePersonalDetailsTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePersonalDetails *PassportElementTypePersonalDetails) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePersonalDetailsType
}

// PassportElementTypePassport A Telegram Passport element containing the user's passport
type PassportElementTypePassport struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePassport
func (passportElementTypePassport *PassportElementTypePassport) MessageType() string {
	return "passportElementTypePassport"
}

// NewPassportElementTypePassport creates a new PassportElementTypePassport
//
func NewPassportElementTypePassport() *PassportElementTypePassport {
	passportElementTypePassportTemp := PassportElementTypePassport{
		tdCommon: tdCommon{Type: "passportElementTypePassport"},
	}

	return &passportElementTypePassportTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePassport *PassportElementTypePassport) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePassportType
}

// PassportElementTypeDriverLicense A Telegram Passport element containing the user's driver license
type PassportElementTypeDriverLicense struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeDriverLicense
func (passportElementTypeDriverLicense *PassportElementTypeDriverLicense) MessageType() string {
	return "passportElementTypeDriverLicense"
}

// NewPassportElementTypeDriverLicense creates a new PassportElementTypeDriverLicense
//
func NewPassportElementTypeDriverLicense() *PassportElementTypeDriverLicense {
	passportElementTypeDriverLicenseTemp := PassportElementTypeDriverLicense{
		tdCommon: tdCommon{Type: "passportElementTypeDriverLicense"},
	}

	return &passportElementTypeDriverLicenseTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeDriverLicense *PassportElementTypeDriverLicense) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeDriverLicenseType
}

// PassportElementTypeIDentityCard A Telegram Passport element containing the user's identity card
type PassportElementTypeIDentityCard struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeIDentityCard
func (passportElementTypeIDentityCard *PassportElementTypeIDentityCard) MessageType() string {
	return "passportElementTypeIdentityCard"
}

// NewPassportElementTypeIDentityCard creates a new PassportElementTypeIDentityCard
//
func NewPassportElementTypeIDentityCard() *PassportElementTypeIDentityCard {
	passportElementTypeIDentityCardTemp := PassportElementTypeIDentityCard{
		tdCommon: tdCommon{Type: "passportElementTypeIdentityCard"},
	}

	return &passportElementTypeIDentityCardTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeIDentityCard *PassportElementTypeIDentityCard) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeIDentityCardType
}

// PassportElementTypeInternalPassport A Telegram Passport element containing the user's internal passport
type PassportElementTypeInternalPassport struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeInternalPassport
func (passportElementTypeInternalPassport *PassportElementTypeInternalPassport) MessageType() string {
	return "passportElementTypeInternalPassport"
}

// NewPassportElementTypeInternalPassport creates a new PassportElementTypeInternalPassport
//
func NewPassportElementTypeInternalPassport() *PassportElementTypeInternalPassport {
	passportElementTypeInternalPassportTemp := PassportElementTypeInternalPassport{
		tdCommon: tdCommon{Type: "passportElementTypeInternalPassport"},
	}

	return &passportElementTypeInternalPassportTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeInternalPassport *PassportElementTypeInternalPassport) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeInternalPassportType
}

// PassportElementTypeAddress A Telegram Passport element containing the user's address
type PassportElementTypeAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeAddress
func (passportElementTypeAddress *PassportElementTypeAddress) MessageType() string {
	return "passportElementTypeAddress"
}

// NewPassportElementTypeAddress creates a new PassportElementTypeAddress
//
func NewPassportElementTypeAddress() *PassportElementTypeAddress {
	passportElementTypeAddressTemp := PassportElementTypeAddress{
		tdCommon: tdCommon{Type: "passportElementTypeAddress"},
	}

	return &passportElementTypeAddressTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeAddress *PassportElementTypeAddress) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeAddressType
}

// PassportElementTypeUtilityBill A Telegram Passport element containing the user's utility bill
type PassportElementTypeUtilityBill struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeUtilityBill
func (passportElementTypeUtilityBill *PassportElementTypeUtilityBill) MessageType() string {
	return "passportElementTypeUtilityBill"
}

// NewPassportElementTypeUtilityBill creates a new PassportElementTypeUtilityBill
//
func NewPassportElementTypeUtilityBill() *PassportElementTypeUtilityBill {
	passportElementTypeUtilityBillTemp := PassportElementTypeUtilityBill{
		tdCommon: tdCommon{Type: "passportElementTypeUtilityBill"},
	}

	return &passportElementTypeUtilityBillTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeUtilityBill *PassportElementTypeUtilityBill) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeUtilityBillType
}

// PassportElementTypeBankStatement A Telegram Passport element containing the user's bank statement
type PassportElementTypeBankStatement struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeBankStatement
func (passportElementTypeBankStatement *PassportElementTypeBankStatement) MessageType() string {
	return "passportElementTypeBankStatement"
}

// NewPassportElementTypeBankStatement creates a new PassportElementTypeBankStatement
//
func NewPassportElementTypeBankStatement() *PassportElementTypeBankStatement {
	passportElementTypeBankStatementTemp := PassportElementTypeBankStatement{
		tdCommon: tdCommon{Type: "passportElementTypeBankStatement"},
	}

	return &passportElementTypeBankStatementTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeBankStatement *PassportElementTypeBankStatement) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeBankStatementType
}

// PassportElementTypeRentalAgreement A Telegram Passport element containing the user's rental agreement
type PassportElementTypeRentalAgreement struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeRentalAgreement
func (passportElementTypeRentalAgreement *PassportElementTypeRentalAgreement) MessageType() string {
	return "passportElementTypeRentalAgreement"
}

// NewPassportElementTypeRentalAgreement creates a new PassportElementTypeRentalAgreement
//
func NewPassportElementTypeRentalAgreement() *PassportElementTypeRentalAgreement {
	passportElementTypeRentalAgreementTemp := PassportElementTypeRentalAgreement{
		tdCommon: tdCommon{Type: "passportElementTypeRentalAgreement"},
	}

	return &passportElementTypeRentalAgreementTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeRentalAgreement *PassportElementTypeRentalAgreement) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeRentalAgreementType
}

// PassportElementTypePassportRegistration A Telegram Passport element containing the registration page of the user's passport
type PassportElementTypePassportRegistration struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePassportRegistration
func (passportElementTypePassportRegistration *PassportElementTypePassportRegistration) MessageType() string {
	return "passportElementTypePassportRegistration"
}

// NewPassportElementTypePassportRegistration creates a new PassportElementTypePassportRegistration
//
func NewPassportElementTypePassportRegistration() *PassportElementTypePassportRegistration {
	passportElementTypePassportRegistrationTemp := PassportElementTypePassportRegistration{
		tdCommon: tdCommon{Type: "passportElementTypePassportRegistration"},
	}

	return &passportElementTypePassportRegistrationTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePassportRegistration *PassportElementTypePassportRegistration) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePassportRegistrationType
}

// PassportElementTypeTemporaryRegistration A Telegram Passport element containing the user's temporary registration
type PassportElementTypeTemporaryRegistration struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeTemporaryRegistration
func (passportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration) MessageType() string {
	return "passportElementTypeTemporaryRegistration"
}

// NewPassportElementTypeTemporaryRegistration creates a new PassportElementTypeTemporaryRegistration
//
func NewPassportElementTypeTemporaryRegistration() *PassportElementTypeTemporaryRegistration {
	passportElementTypeTemporaryRegistrationTemp := PassportElementTypeTemporaryRegistration{
		tdCommon: tdCommon{Type: "passportElementTypeTemporaryRegistration"},
	}

	return &passportElementTypeTemporaryRegistrationTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeTemporaryRegistrationType
}

// PassportElementTypePhoneNumber A Telegram Passport element containing the user's phone number
type PassportElementTypePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePhoneNumber
func (passportElementTypePhoneNumber *PassportElementTypePhoneNumber) MessageType() string {
	return "passportElementTypePhoneNumber"
}

// NewPassportElementTypePhoneNumber creates a new PassportElementTypePhoneNumber
//
func NewPassportElementTypePhoneNumber() *PassportElementTypePhoneNumber {
	passportElementTypePhoneNumberTemp := PassportElementTypePhoneNumber{
		tdCommon: tdCommon{Type: "passportElementTypePhoneNumber"},
	}

	return &passportElementTypePhoneNumberTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePhoneNumber *PassportElementTypePhoneNumber) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePhoneNumberType
}

// PassportElementTypeEmailAddress A Telegram Passport element containing the user's email address
type PassportElementTypeEmailAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeEmailAddress
func (passportElementTypeEmailAddress *PassportElementTypeEmailAddress) MessageType() string {
	return "passportElementTypeEmailAddress"
}

// NewPassportElementTypeEmailAddress creates a new PassportElementTypeEmailAddress
//
func NewPassportElementTypeEmailAddress() *PassportElementTypeEmailAddress {
	passportElementTypeEmailAddressTemp := PassportElementTypeEmailAddress{
		tdCommon: tdCommon{Type: "passportElementTypeEmailAddress"},
	}

	return &passportElementTypeEmailAddressTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeEmailAddress *PassportElementTypeEmailAddress) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeEmailAddressType
}

// Date Represents a date according to the Gregorian calendar
type Date struct {
	tdCommon
	Day   int32 `json:"day"`   // Day of the month, 1-31
	Month int32 `json:"month"` // Month, 1-12
	Year  int32 `json:"year"`  // Year, 1-9999
}

// MessageType return the string telegram-type of Date
func (date *Date) MessageType() string {
	return "date"
}

// NewDate creates a new Date
//
// @param day Day of the month, 1-31
// @param month Month, 1-12
// @param year Year, 1-9999
func NewDate(day int32, month int32, year int32) *Date {
	dateTemp := Date{
		tdCommon: tdCommon{Type: "date"},
		Day:      day,
		Month:    month,
		Year:     year,
	}

	return &dateTemp
}

// PersonalDetails Contains the user's personal details
type PersonalDetails struct {
	tdCommon
	FirstName            string `json:"first_name"`             // First name of the user written in English; 1-255 characters
	MiddleName           string `json:"middle_name"`            // Middle name of the user written in English; 0-255 characters
	LastName             string `json:"last_name"`              // Last name of the user written in English; 1-255 characters
	NativeFirstName      string `json:"native_first_name"`      // Native first name of the user; 1-255 characters
	NativeMiddleName     string `json:"native_middle_name"`     // Native middle name of the user; 0-255 characters
	NativeLastName       string `json:"native_last_name"`       // Native last name of the user; 1-255 characters
	Birthdate            *Date  `json:"birthdate"`              // Birthdate of the user
	Gender               string `json:"gender"`                 // Gender of the user, "male" or "female"
	CountryCode          string `json:"country_code"`           // A two-letter ISO 3166-1 alpha-2 country code of the user's country
	ResidenceCountryCode string `json:"residence_country_code"` // A two-letter ISO 3166-1 alpha-2 country code of the user's residence country
}

// MessageType return the string telegram-type of PersonalDetails
func (personalDetails *PersonalDetails) MessageType() string {
	return "personalDetails"
}

// NewPersonalDetails creates a new PersonalDetails
//
// @param firstName First name of the user written in English; 1-255 characters
// @param middleName Middle name of the user written in English; 0-255 characters
// @param lastName Last name of the user written in English; 1-255 characters
// @param nativeFirstName Native first name of the user; 1-255 characters
// @param nativeMiddleName Native middle name of the user; 0-255 characters
// @param nativeLastName Native last name of the user; 1-255 characters
// @param birthdate Birthdate of the user
// @param gender Gender of the user, "male" or "female"
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code of the user's country
// @param residenceCountryCode A two-letter ISO 3166-1 alpha-2 country code of the user's residence country
func NewPersonalDetails(firstName string, middleName string, lastName string, nativeFirstName string, nativeMiddleName string, nativeLastName string, birthdate *Date, gender string, countryCode string, residenceCountryCode string) *PersonalDetails {
	personalDetailsTemp := PersonalDetails{
		tdCommon:             tdCommon{Type: "personalDetails"},
		FirstName:            firstName,
		MiddleName:           middleName,
		LastName:             lastName,
		NativeFirstName:      nativeFirstName,
		NativeMiddleName:     nativeMiddleName,
		NativeLastName:       nativeLastName,
		Birthdate:            birthdate,
		Gender:               gender,
		CountryCode:          countryCode,
		ResidenceCountryCode: residenceCountryCode,
	}

	return &personalDetailsTemp
}

// IDentityDocument An identity document
type IDentityDocument struct {
	tdCommon
	Number      string      `json:"number"`       // Document number; 1-24 characters
	ExpiryDate  *Date       `json:"expiry_date"`  // Document expiry date; may be null
	FrontSide   *DatedFile  `json:"front_side"`   // Front side of the document
	ReverseSide *DatedFile  `json:"reverse_side"` // Reverse side of the document; only for driver license and identity card
	Selfie      *DatedFile  `json:"selfie"`       // Selfie with the document; may be null
	Translation []DatedFile `json:"translation"`  // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of IDentityDocument
func (iDentityDocument *IDentityDocument) MessageType() string {
	return "identityDocument"
}

// NewIDentityDocument creates a new IDentityDocument
//
// @param number Document number; 1-24 characters
// @param expiryDate Document expiry date; may be null
// @param frontSide Front side of the document
// @param reverseSide Reverse side of the document; only for driver license and identity card
// @param selfie Selfie with the document; may be null
// @param translation List of files containing a certified English translation of the document
func NewIDentityDocument(number string, expiryDate *Date, frontSide *DatedFile, reverseSide *DatedFile, selfie *DatedFile, translation []DatedFile) *IDentityDocument {
	iDentityDocumentTemp := IDentityDocument{
		tdCommon:    tdCommon{Type: "identityDocument"},
		Number:      number,
		ExpiryDate:  expiryDate,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
	}

	return &iDentityDocumentTemp
}

// InputIDentityDocument An identity document to be saved to Telegram Passport
type InputIDentityDocument struct {
	tdCommon
	Number      string      `json:"number"`       // Document number; 1-24 characters
	ExpiryDate  *Date       `json:"expiry_date"`  // Document expiry date, if available
	FrontSide   InputFile   `json:"front_side"`   // Front side of the document
	ReverseSide InputFile   `json:"reverse_side"` // Reverse side of the document; only for driver license and identity card
	Selfie      InputFile   `json:"selfie"`       // Selfie with the document, if available
	Translation []InputFile `json:"translation"`  // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of InputIDentityDocument
func (inputIDentityDocument *InputIDentityDocument) MessageType() string {
	return "inputIdentityDocument"
}

// NewInputIDentityDocument creates a new InputIDentityDocument
//
// @param number Document number; 1-24 characters
// @param expiryDate Document expiry date, if available
// @param frontSide Front side of the document
// @param reverseSide Reverse side of the document; only for driver license and identity card
// @param selfie Selfie with the document, if available
// @param translation List of files containing a certified English translation of the document
func NewInputIDentityDocument(number string, expiryDate *Date, frontSide InputFile, reverseSide InputFile, selfie InputFile, translation []InputFile) *InputIDentityDocument {
	inputIDentityDocumentTemp := InputIDentityDocument{
		tdCommon:    tdCommon{Type: "inputIdentityDocument"},
		Number:      number,
		ExpiryDate:  expiryDate,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
	}

	return &inputIDentityDocumentTemp
}

// UnmarshalJSON unmarshal to json
func (inputIDentityDocument *InputIDentityDocument) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Number      string      `json:"number"`      // Document number; 1-24 characters
		ExpiryDate  *Date       `json:"expiry_date"` // Document expiry date, if available
		Translation []InputFile `json:"translation"` // List of files containing a certified English translation of the document
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputIDentityDocument.tdCommon = tempObj.tdCommon
	inputIDentityDocument.Number = tempObj.Number
	inputIDentityDocument.ExpiryDate = tempObj.ExpiryDate
	inputIDentityDocument.Translation = tempObj.Translation

	fieldFrontSide, _ := unmarshalInputFile(objMap["front_side"])
	inputIDentityDocument.FrontSide = fieldFrontSide

	fieldReverseSide, _ := unmarshalInputFile(objMap["reverse_side"])
	inputIDentityDocument.ReverseSide = fieldReverseSide

	fieldSelfie, _ := unmarshalInputFile(objMap["selfie"])
	inputIDentityDocument.Selfie = fieldSelfie

	return nil
}

// PersonalDocument A personal document, containing some information about a user
type PersonalDocument struct {
	tdCommon
	Files       []DatedFile `json:"files"`       // List of files containing the pages of the document
	Translation []DatedFile `json:"translation"` // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of PersonalDocument
func (personalDocument *PersonalDocument) MessageType() string {
	return "personalDocument"
}

// NewPersonalDocument creates a new PersonalDocument
//
// @param files List of files containing the pages of the document
// @param translation List of files containing a certified English translation of the document
func NewPersonalDocument(files []DatedFile, translation []DatedFile) *PersonalDocument {
	personalDocumentTemp := PersonalDocument{
		tdCommon:    tdCommon{Type: "personalDocument"},
		Files:       files,
		Translation: translation,
	}

	return &personalDocumentTemp
}

// InputPersonalDocument A personal document to be saved to Telegram Passport
type InputPersonalDocument struct {
	tdCommon
	Files       []InputFile `json:"files"`       // List of files containing the pages of the document
	Translation []InputFile `json:"translation"` // List of files containing a certified English translation of the document
}

// MessageType return the string telegram-type of InputPersonalDocument
func (inputPersonalDocument *InputPersonalDocument) MessageType() string {
	return "inputPersonalDocument"
}

// NewInputPersonalDocument creates a new InputPersonalDocument
//
// @param files List of files containing the pages of the document
// @param translation List of files containing a certified English translation of the document
func NewInputPersonalDocument(files []InputFile, translation []InputFile) *InputPersonalDocument {
	inputPersonalDocumentTemp := InputPersonalDocument{
		tdCommon:    tdCommon{Type: "inputPersonalDocument"},
		Files:       files,
		Translation: translation,
	}

	return &inputPersonalDocumentTemp
}

// PassportElementPersonalDetails A Telegram Passport element containing the user's personal details
type PassportElementPersonalDetails struct {
	tdCommon
	PersonalDetails *PersonalDetails `json:"personal_details"` // Personal details of the user
}

// MessageType return the string telegram-type of PassportElementPersonalDetails
func (passportElementPersonalDetails *PassportElementPersonalDetails) MessageType() string {
	return "passportElementPersonalDetails"
}

// NewPassportElementPersonalDetails creates a new PassportElementPersonalDetails
//
// @param personalDetails Personal details of the user
func NewPassportElementPersonalDetails(personalDetails *PersonalDetails) *PassportElementPersonalDetails {
	passportElementPersonalDetailsTemp := PassportElementPersonalDetails{
		tdCommon:        tdCommon{Type: "passportElementPersonalDetails"},
		PersonalDetails: personalDetails,
	}

	return &passportElementPersonalDetailsTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPersonalDetails *PassportElementPersonalDetails) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPersonalDetailsType
}

// PassportElementPassport A Telegram Passport element containing the user's passport
type PassportElementPassport struct {
	tdCommon
	Passport *IDentityDocument `json:"passport"` // Passport
}

// MessageType return the string telegram-type of PassportElementPassport
func (passportElementPassport *PassportElementPassport) MessageType() string {
	return "passportElementPassport"
}

// NewPassportElementPassport creates a new PassportElementPassport
//
// @param passport Passport
func NewPassportElementPassport(passport *IDentityDocument) *PassportElementPassport {
	passportElementPassportTemp := PassportElementPassport{
		tdCommon: tdCommon{Type: "passportElementPassport"},
		Passport: passport,
	}

	return &passportElementPassportTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPassport *PassportElementPassport) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPassportType
}

// PassportElementDriverLicense A Telegram Passport element containing the user's driver license
type PassportElementDriverLicense struct {
	tdCommon
	DriverLicense *IDentityDocument `json:"driver_license"` // Driver license
}

// MessageType return the string telegram-type of PassportElementDriverLicense
func (passportElementDriverLicense *PassportElementDriverLicense) MessageType() string {
	return "passportElementDriverLicense"
}

// NewPassportElementDriverLicense creates a new PassportElementDriverLicense
//
// @param driverLicense Driver license
func NewPassportElementDriverLicense(driverLicense *IDentityDocument) *PassportElementDriverLicense {
	passportElementDriverLicenseTemp := PassportElementDriverLicense{
		tdCommon:      tdCommon{Type: "passportElementDriverLicense"},
		DriverLicense: driverLicense,
	}

	return &passportElementDriverLicenseTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementDriverLicense *PassportElementDriverLicense) GetPassportElementEnum() PassportElementEnum {
	return PassportElementDriverLicenseType
}

// PassportElementIDentityCard A Telegram Passport element containing the user's identity card
type PassportElementIDentityCard struct {
	tdCommon
	IDentityCard *IDentityDocument `json:"identity_card"` // Identity card
}

// MessageType return the string telegram-type of PassportElementIDentityCard
func (passportElementIDentityCard *PassportElementIDentityCard) MessageType() string {
	return "passportElementIdentityCard"
}

// NewPassportElementIDentityCard creates a new PassportElementIDentityCard
//
// @param iDentityCard Identity card
func NewPassportElementIDentityCard(iDentityCard *IDentityDocument) *PassportElementIDentityCard {
	passportElementIDentityCardTemp := PassportElementIDentityCard{
		tdCommon:     tdCommon{Type: "passportElementIdentityCard"},
		IDentityCard: iDentityCard,
	}

	return &passportElementIDentityCardTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementIDentityCard *PassportElementIDentityCard) GetPassportElementEnum() PassportElementEnum {
	return PassportElementIDentityCardType
}

// PassportElementInternalPassport A Telegram Passport element containing the user's internal passport
type PassportElementInternalPassport struct {
	tdCommon
	InternalPassport *IDentityDocument `json:"internal_passport"` // Internal passport
}

// MessageType return the string telegram-type of PassportElementInternalPassport
func (passportElementInternalPassport *PassportElementInternalPassport) MessageType() string {
	return "passportElementInternalPassport"
}

// NewPassportElementInternalPassport creates a new PassportElementInternalPassport
//
// @param internalPassport Internal passport
func NewPassportElementInternalPassport(internalPassport *IDentityDocument) *PassportElementInternalPassport {
	passportElementInternalPassportTemp := PassportElementInternalPassport{
		tdCommon:         tdCommon{Type: "passportElementInternalPassport"},
		InternalPassport: internalPassport,
	}

	return &passportElementInternalPassportTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementInternalPassport *PassportElementInternalPassport) GetPassportElementEnum() PassportElementEnum {
	return PassportElementInternalPassportType
}

// PassportElementAddress A Telegram Passport element containing the user's address
type PassportElementAddress struct {
	tdCommon
	Address *Address `json:"address"` // Address
}

// MessageType return the string telegram-type of PassportElementAddress
func (passportElementAddress *PassportElementAddress) MessageType() string {
	return "passportElementAddress"
}

// NewPassportElementAddress creates a new PassportElementAddress
//
// @param address Address
func NewPassportElementAddress(address *Address) *PassportElementAddress {
	passportElementAddressTemp := PassportElementAddress{
		tdCommon: tdCommon{Type: "passportElementAddress"},
		Address:  address,
	}

	return &passportElementAddressTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementAddress *PassportElementAddress) GetPassportElementEnum() PassportElementEnum {
	return PassportElementAddressType
}

// PassportElementUtilityBill A Telegram Passport element containing the user's utility bill
type PassportElementUtilityBill struct {
	tdCommon
	UtilityBill *PersonalDocument `json:"utility_bill"` // Utility bill
}

// MessageType return the string telegram-type of PassportElementUtilityBill
func (passportElementUtilityBill *PassportElementUtilityBill) MessageType() string {
	return "passportElementUtilityBill"
}

// NewPassportElementUtilityBill creates a new PassportElementUtilityBill
//
// @param utilityBill Utility bill
func NewPassportElementUtilityBill(utilityBill *PersonalDocument) *PassportElementUtilityBill {
	passportElementUtilityBillTemp := PassportElementUtilityBill{
		tdCommon:    tdCommon{Type: "passportElementUtilityBill"},
		UtilityBill: utilityBill,
	}

	return &passportElementUtilityBillTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementUtilityBill *PassportElementUtilityBill) GetPassportElementEnum() PassportElementEnum {
	return PassportElementUtilityBillType
}

// PassportElementBankStatement A Telegram Passport element containing the user's bank statement
type PassportElementBankStatement struct {
	tdCommon
	BankStatement *PersonalDocument `json:"bank_statement"` // Bank statement
}

// MessageType return the string telegram-type of PassportElementBankStatement
func (passportElementBankStatement *PassportElementBankStatement) MessageType() string {
	return "passportElementBankStatement"
}

// NewPassportElementBankStatement creates a new PassportElementBankStatement
//
// @param bankStatement Bank statement
func NewPassportElementBankStatement(bankStatement *PersonalDocument) *PassportElementBankStatement {
	passportElementBankStatementTemp := PassportElementBankStatement{
		tdCommon:      tdCommon{Type: "passportElementBankStatement"},
		BankStatement: bankStatement,
	}

	return &passportElementBankStatementTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementBankStatement *PassportElementBankStatement) GetPassportElementEnum() PassportElementEnum {
	return PassportElementBankStatementType
}

// PassportElementRentalAgreement A Telegram Passport element containing the user's rental agreement
type PassportElementRentalAgreement struct {
	tdCommon
	RentalAgreement *PersonalDocument `json:"rental_agreement"` // Rental agreement
}

// MessageType return the string telegram-type of PassportElementRentalAgreement
func (passportElementRentalAgreement *PassportElementRentalAgreement) MessageType() string {
	return "passportElementRentalAgreement"
}

// NewPassportElementRentalAgreement creates a new PassportElementRentalAgreement
//
// @param rentalAgreement Rental agreement
func NewPassportElementRentalAgreement(rentalAgreement *PersonalDocument) *PassportElementRentalAgreement {
	passportElementRentalAgreementTemp := PassportElementRentalAgreement{
		tdCommon:        tdCommon{Type: "passportElementRentalAgreement"},
		RentalAgreement: rentalAgreement,
	}

	return &passportElementRentalAgreementTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementRentalAgreement *PassportElementRentalAgreement) GetPassportElementEnum() PassportElementEnum {
	return PassportElementRentalAgreementType
}

// PassportElementPassportRegistration A Telegram Passport element containing the user's passport registration pages
type PassportElementPassportRegistration struct {
	tdCommon
	PassportRegistration *PersonalDocument `json:"passport_registration"` // Passport registration pages
}

// MessageType return the string telegram-type of PassportElementPassportRegistration
func (passportElementPassportRegistration *PassportElementPassportRegistration) MessageType() string {
	return "passportElementPassportRegistration"
}

// NewPassportElementPassportRegistration creates a new PassportElementPassportRegistration
//
// @param passportRegistration Passport registration pages
func NewPassportElementPassportRegistration(passportRegistration *PersonalDocument) *PassportElementPassportRegistration {
	passportElementPassportRegistrationTemp := PassportElementPassportRegistration{
		tdCommon:             tdCommon{Type: "passportElementPassportRegistration"},
		PassportRegistration: passportRegistration,
	}

	return &passportElementPassportRegistrationTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPassportRegistration *PassportElementPassportRegistration) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPassportRegistrationType
}

// PassportElementTemporaryRegistration A Telegram Passport element containing the user's temporary registration
type PassportElementTemporaryRegistration struct {
	tdCommon
	TemporaryRegistration *PersonalDocument `json:"temporary_registration"` // Temporary registration
}

// MessageType return the string telegram-type of PassportElementTemporaryRegistration
func (passportElementTemporaryRegistration *PassportElementTemporaryRegistration) MessageType() string {
	return "passportElementTemporaryRegistration"
}

// NewPassportElementTemporaryRegistration creates a new PassportElementTemporaryRegistration
//
// @param temporaryRegistration Temporary registration
func NewPassportElementTemporaryRegistration(temporaryRegistration *PersonalDocument) *PassportElementTemporaryRegistration {
	passportElementTemporaryRegistrationTemp := PassportElementTemporaryRegistration{
		tdCommon:              tdCommon{Type: "passportElementTemporaryRegistration"},
		TemporaryRegistration: temporaryRegistration,
	}

	return &passportElementTemporaryRegistrationTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementTemporaryRegistration *PassportElementTemporaryRegistration) GetPassportElementEnum() PassportElementEnum {
	return PassportElementTemporaryRegistrationType
}

// PassportElementPhoneNumber A Telegram Passport element containing the user's phone number
type PassportElementPhoneNumber struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // Phone number
}

// MessageType return the string telegram-type of PassportElementPhoneNumber
func (passportElementPhoneNumber *PassportElementPhoneNumber) MessageType() string {
	return "passportElementPhoneNumber"
}

// NewPassportElementPhoneNumber creates a new PassportElementPhoneNumber
//
// @param phoneNumber Phone number
func NewPassportElementPhoneNumber(phoneNumber string) *PassportElementPhoneNumber {
	passportElementPhoneNumberTemp := PassportElementPhoneNumber{
		tdCommon:    tdCommon{Type: "passportElementPhoneNumber"},
		PhoneNumber: phoneNumber,
	}

	return &passportElementPhoneNumberTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPhoneNumber *PassportElementPhoneNumber) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPhoneNumberType
}

// PassportElementEmailAddress A Telegram Passport element containing the user's email address
type PassportElementEmailAddress struct {
	tdCommon
	EmailAddress string `json:"email_address"` // Email address
}

// MessageType return the string telegram-type of PassportElementEmailAddress
func (passportElementEmailAddress *PassportElementEmailAddress) MessageType() string {
	return "passportElementEmailAddress"
}

// NewPassportElementEmailAddress creates a new PassportElementEmailAddress
//
// @param emailAddress Email address
func NewPassportElementEmailAddress(emailAddress string) *PassportElementEmailAddress {
	passportElementEmailAddressTemp := PassportElementEmailAddress{
		tdCommon:     tdCommon{Type: "passportElementEmailAddress"},
		EmailAddress: emailAddress,
	}

	return &passportElementEmailAddressTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementEmailAddress *PassportElementEmailAddress) GetPassportElementEnum() PassportElementEnum {
	return PassportElementEmailAddressType
}

// InputPassportElementPersonalDetails A Telegram Passport element to be saved containing the user's personal details
type InputPassportElementPersonalDetails struct {
	tdCommon
	PersonalDetails *PersonalDetails `json:"personal_details"` // Personal details of the user
}

// MessageType return the string telegram-type of InputPassportElementPersonalDetails
func (inputPassportElementPersonalDetails *InputPassportElementPersonalDetails) MessageType() string {
	return "inputPassportElementPersonalDetails"
}

// NewInputPassportElementPersonalDetails creates a new InputPassportElementPersonalDetails
//
// @param personalDetails Personal details of the user
func NewInputPassportElementPersonalDetails(personalDetails *PersonalDetails) *InputPassportElementPersonalDetails {
	inputPassportElementPersonalDetailsTemp := InputPassportElementPersonalDetails{
		tdCommon:        tdCommon{Type: "inputPassportElementPersonalDetails"},
		PersonalDetails: personalDetails,
	}

	return &inputPassportElementPersonalDetailsTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPersonalDetails *InputPassportElementPersonalDetails) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPersonalDetailsType
}

// InputPassportElementPassport A Telegram Passport element to be saved containing the user's passport
type InputPassportElementPassport struct {
	tdCommon
	Passport *InputIDentityDocument `json:"passport"` // The passport to be saved
}

// MessageType return the string telegram-type of InputPassportElementPassport
func (inputPassportElementPassport *InputPassportElementPassport) MessageType() string {
	return "inputPassportElementPassport"
}

// NewInputPassportElementPassport creates a new InputPassportElementPassport
//
// @param passport The passport to be saved
func NewInputPassportElementPassport(passport *InputIDentityDocument) *InputPassportElementPassport {
	inputPassportElementPassportTemp := InputPassportElementPassport{
		tdCommon: tdCommon{Type: "inputPassportElementPassport"},
		Passport: passport,
	}

	return &inputPassportElementPassportTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPassport *InputPassportElementPassport) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPassportType
}

// InputPassportElementDriverLicense A Telegram Passport element to be saved containing the user's driver license
type InputPassportElementDriverLicense struct {
	tdCommon
	DriverLicense *InputIDentityDocument `json:"driver_license"` // The driver license to be saved
}

// MessageType return the string telegram-type of InputPassportElementDriverLicense
func (inputPassportElementDriverLicense *InputPassportElementDriverLicense) MessageType() string {
	return "inputPassportElementDriverLicense"
}

// NewInputPassportElementDriverLicense creates a new InputPassportElementDriverLicense
//
// @param driverLicense The driver license to be saved
func NewInputPassportElementDriverLicense(driverLicense *InputIDentityDocument) *InputPassportElementDriverLicense {
	inputPassportElementDriverLicenseTemp := InputPassportElementDriverLicense{
		tdCommon:      tdCommon{Type: "inputPassportElementDriverLicense"},
		DriverLicense: driverLicense,
	}

	return &inputPassportElementDriverLicenseTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementDriverLicense *InputPassportElementDriverLicense) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementDriverLicenseType
}

// InputPassportElementIDentityCard A Telegram Passport element to be saved containing the user's identity card
type InputPassportElementIDentityCard struct {
	tdCommon
	IDentityCard *InputIDentityDocument `json:"identity_card"` // The identity card to be saved
}

// MessageType return the string telegram-type of InputPassportElementIDentityCard
func (inputPassportElementIDentityCard *InputPassportElementIDentityCard) MessageType() string {
	return "inputPassportElementIdentityCard"
}

// NewInputPassportElementIDentityCard creates a new InputPassportElementIDentityCard
//
// @param iDentityCard The identity card to be saved
func NewInputPassportElementIDentityCard(iDentityCard *InputIDentityDocument) *InputPassportElementIDentityCard {
	inputPassportElementIDentityCardTemp := InputPassportElementIDentityCard{
		tdCommon:     tdCommon{Type: "inputPassportElementIdentityCard"},
		IDentityCard: iDentityCard,
	}

	return &inputPassportElementIDentityCardTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementIDentityCard *InputPassportElementIDentityCard) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementIDentityCardType
}

// InputPassportElementInternalPassport A Telegram Passport element to be saved containing the user's internal passport
type InputPassportElementInternalPassport struct {
	tdCommon
	InternalPassport *InputIDentityDocument `json:"internal_passport"` // The internal passport to be saved
}

// MessageType return the string telegram-type of InputPassportElementInternalPassport
func (inputPassportElementInternalPassport *InputPassportElementInternalPassport) MessageType() string {
	return "inputPassportElementInternalPassport"
}

// NewInputPassportElementInternalPassport creates a new InputPassportElementInternalPassport
//
// @param internalPassport The internal passport to be saved
func NewInputPassportElementInternalPassport(internalPassport *InputIDentityDocument) *InputPassportElementInternalPassport {
	inputPassportElementInternalPassportTemp := InputPassportElementInternalPassport{
		tdCommon:         tdCommon{Type: "inputPassportElementInternalPassport"},
		InternalPassport: internalPassport,
	}

	return &inputPassportElementInternalPassportTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementInternalPassport *InputPassportElementInternalPassport) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementInternalPassportType
}

// InputPassportElementAddress A Telegram Passport element to be saved containing the user's address
type InputPassportElementAddress struct {
	tdCommon
	Address *Address `json:"address"` // The address to be saved
}

// MessageType return the string telegram-type of InputPassportElementAddress
func (inputPassportElementAddress *InputPassportElementAddress) MessageType() string {
	return "inputPassportElementAddress"
}

// NewInputPassportElementAddress creates a new InputPassportElementAddress
//
// @param address The address to be saved
func NewInputPassportElementAddress(address *Address) *InputPassportElementAddress {
	inputPassportElementAddressTemp := InputPassportElementAddress{
		tdCommon: tdCommon{Type: "inputPassportElementAddress"},
		Address:  address,
	}

	return &inputPassportElementAddressTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementAddress *InputPassportElementAddress) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementAddressType
}

// InputPassportElementUtilityBill A Telegram Passport element to be saved containing the user's utility bill
type InputPassportElementUtilityBill struct {
	tdCommon
	UtilityBill *InputPersonalDocument `json:"utility_bill"` // The utility bill to be saved
}

// MessageType return the string telegram-type of InputPassportElementUtilityBill
func (inputPassportElementUtilityBill *InputPassportElementUtilityBill) MessageType() string {
	return "inputPassportElementUtilityBill"
}

// NewInputPassportElementUtilityBill creates a new InputPassportElementUtilityBill
//
// @param utilityBill The utility bill to be saved
func NewInputPassportElementUtilityBill(utilityBill *InputPersonalDocument) *InputPassportElementUtilityBill {
	inputPassportElementUtilityBillTemp := InputPassportElementUtilityBill{
		tdCommon:    tdCommon{Type: "inputPassportElementUtilityBill"},
		UtilityBill: utilityBill,
	}

	return &inputPassportElementUtilityBillTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementUtilityBill *InputPassportElementUtilityBill) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementUtilityBillType
}

// InputPassportElementBankStatement A Telegram Passport element to be saved containing the user's bank statement
type InputPassportElementBankStatement struct {
	tdCommon
	BankStatement *InputPersonalDocument `json:"bank_statement"` // The bank statement to be saved
}

// MessageType return the string telegram-type of InputPassportElementBankStatement
func (inputPassportElementBankStatement *InputPassportElementBankStatement) MessageType() string {
	return "inputPassportElementBankStatement"
}

// NewInputPassportElementBankStatement creates a new InputPassportElementBankStatement
//
// @param bankStatement The bank statement to be saved
func NewInputPassportElementBankStatement(bankStatement *InputPersonalDocument) *InputPassportElementBankStatement {
	inputPassportElementBankStatementTemp := InputPassportElementBankStatement{
		tdCommon:      tdCommon{Type: "inputPassportElementBankStatement"},
		BankStatement: bankStatement,
	}

	return &inputPassportElementBankStatementTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementBankStatement *InputPassportElementBankStatement) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementBankStatementType
}

// InputPassportElementRentalAgreement A Telegram Passport element to be saved containing the user's rental agreement
type InputPassportElementRentalAgreement struct {
	tdCommon
	RentalAgreement *InputPersonalDocument `json:"rental_agreement"` // The rental agreement to be saved
}

// MessageType return the string telegram-type of InputPassportElementRentalAgreement
func (inputPassportElementRentalAgreement *InputPassportElementRentalAgreement) MessageType() string {
	return "inputPassportElementRentalAgreement"
}

// NewInputPassportElementRentalAgreement creates a new InputPassportElementRentalAgreement
//
// @param rentalAgreement The rental agreement to be saved
func NewInputPassportElementRentalAgreement(rentalAgreement *InputPersonalDocument) *InputPassportElementRentalAgreement {
	inputPassportElementRentalAgreementTemp := InputPassportElementRentalAgreement{
		tdCommon:        tdCommon{Type: "inputPassportElementRentalAgreement"},
		RentalAgreement: rentalAgreement,
	}

	return &inputPassportElementRentalAgreementTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementRentalAgreement *InputPassportElementRentalAgreement) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementRentalAgreementType
}

// InputPassportElementPassportRegistration A Telegram Passport element to be saved containing the user's passport registration
type InputPassportElementPassportRegistration struct {
	tdCommon
	PassportRegistration *InputPersonalDocument `json:"passport_registration"` // The passport registration page to be saved
}

// MessageType return the string telegram-type of InputPassportElementPassportRegistration
func (inputPassportElementPassportRegistration *InputPassportElementPassportRegistration) MessageType() string {
	return "inputPassportElementPassportRegistration"
}

// NewInputPassportElementPassportRegistration creates a new InputPassportElementPassportRegistration
//
// @param passportRegistration The passport registration page to be saved
func NewInputPassportElementPassportRegistration(passportRegistration *InputPersonalDocument) *InputPassportElementPassportRegistration {
	inputPassportElementPassportRegistrationTemp := InputPassportElementPassportRegistration{
		tdCommon:             tdCommon{Type: "inputPassportElementPassportRegistration"},
		PassportRegistration: passportRegistration,
	}

	return &inputPassportElementPassportRegistrationTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPassportRegistration *InputPassportElementPassportRegistration) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPassportRegistrationType
}

// InputPassportElementTemporaryRegistration A Telegram Passport element to be saved containing the user's temporary registration
type InputPassportElementTemporaryRegistration struct {
	tdCommon
	TemporaryRegistration *InputPersonalDocument `json:"temporary_registration"` // The temporary registration document to be saved
}

// MessageType return the string telegram-type of InputPassportElementTemporaryRegistration
func (inputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration) MessageType() string {
	return "inputPassportElementTemporaryRegistration"
}

// NewInputPassportElementTemporaryRegistration creates a new InputPassportElementTemporaryRegistration
//
// @param temporaryRegistration The temporary registration document to be saved
func NewInputPassportElementTemporaryRegistration(temporaryRegistration *InputPersonalDocument) *InputPassportElementTemporaryRegistration {
	inputPassportElementTemporaryRegistrationTemp := InputPassportElementTemporaryRegistration{
		tdCommon:              tdCommon{Type: "inputPassportElementTemporaryRegistration"},
		TemporaryRegistration: temporaryRegistration,
	}

	return &inputPassportElementTemporaryRegistrationTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementTemporaryRegistrationType
}

// InputPassportElementPhoneNumber A Telegram Passport element to be saved containing the user's phone number
type InputPassportElementPhoneNumber struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // The phone number to be saved
}

// MessageType return the string telegram-type of InputPassportElementPhoneNumber
func (inputPassportElementPhoneNumber *InputPassportElementPhoneNumber) MessageType() string {
	return "inputPassportElementPhoneNumber"
}

// NewInputPassportElementPhoneNumber creates a new InputPassportElementPhoneNumber
//
// @param phoneNumber The phone number to be saved
func NewInputPassportElementPhoneNumber(phoneNumber string) *InputPassportElementPhoneNumber {
	inputPassportElementPhoneNumberTemp := InputPassportElementPhoneNumber{
		tdCommon:    tdCommon{Type: "inputPassportElementPhoneNumber"},
		PhoneNumber: phoneNumber,
	}

	return &inputPassportElementPhoneNumberTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPhoneNumber *InputPassportElementPhoneNumber) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPhoneNumberType
}

// InputPassportElementEmailAddress A Telegram Passport element to be saved containing the user's email address
type InputPassportElementEmailAddress struct {
	tdCommon
	EmailAddress string `json:"email_address"` // The email address to be saved
}

// MessageType return the string telegram-type of InputPassportElementEmailAddress
func (inputPassportElementEmailAddress *InputPassportElementEmailAddress) MessageType() string {
	return "inputPassportElementEmailAddress"
}

// NewInputPassportElementEmailAddress creates a new InputPassportElementEmailAddress
//
// @param emailAddress The email address to be saved
func NewInputPassportElementEmailAddress(emailAddress string) *InputPassportElementEmailAddress {
	inputPassportElementEmailAddressTemp := InputPassportElementEmailAddress{
		tdCommon:     tdCommon{Type: "inputPassportElementEmailAddress"},
		EmailAddress: emailAddress,
	}

	return &inputPassportElementEmailAddressTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementEmailAddress *InputPassportElementEmailAddress) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementEmailAddressType
}

// PassportElements Contains information about saved Telegram Passport elements
type PassportElements struct {
	tdCommon
	Elements []PassportElement `json:"elements"` // Telegram Passport elements
}

// MessageType return the string telegram-type of PassportElements
func (passportElements *PassportElements) MessageType() string {
	return "passportElements"
}

// NewPassportElements creates a new PassportElements
//
// @param elements Telegram Passport elements
func NewPassportElements(elements []PassportElement) *PassportElements {
	passportElementsTemp := PassportElements{
		tdCommon: tdCommon{Type: "passportElements"},
		Elements: elements,
	}

	return &passportElementsTemp
}

// PassportElementErrorSourceUnspecified The element contains an error in an unspecified place. The error will be considered resolved when new data is added
type PassportElementErrorSourceUnspecified struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceUnspecified
func (passportElementErrorSourceUnspecified *PassportElementErrorSourceUnspecified) MessageType() string {
	return "passportElementErrorSourceUnspecified"
}

// NewPassportElementErrorSourceUnspecified creates a new PassportElementErrorSourceUnspecified
//
func NewPassportElementErrorSourceUnspecified() *PassportElementErrorSourceUnspecified {
	passportElementErrorSourceUnspecifiedTemp := PassportElementErrorSourceUnspecified{
		tdCommon: tdCommon{Type: "passportElementErrorSourceUnspecified"},
	}

	return &passportElementErrorSourceUnspecifiedTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceUnspecified *PassportElementErrorSourceUnspecified) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceUnspecifiedType
}

// PassportElementErrorSourceDataField One of the data fields contains an error. The error will be considered resolved when the value of the field changes
type PassportElementErrorSourceDataField struct {
	tdCommon
	FieldName string `json:"field_name"` // Field name
}

// MessageType return the string telegram-type of PassportElementErrorSourceDataField
func (passportElementErrorSourceDataField *PassportElementErrorSourceDataField) MessageType() string {
	return "passportElementErrorSourceDataField"
}

// NewPassportElementErrorSourceDataField creates a new PassportElementErrorSourceDataField
//
// @param fieldName Field name
func NewPassportElementErrorSourceDataField(fieldName string) *PassportElementErrorSourceDataField {
	passportElementErrorSourceDataFieldTemp := PassportElementErrorSourceDataField{
		tdCommon:  tdCommon{Type: "passportElementErrorSourceDataField"},
		FieldName: fieldName,
	}

	return &passportElementErrorSourceDataFieldTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceDataField *PassportElementErrorSourceDataField) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceDataFieldType
}

// PassportElementErrorSourceFrontSide The front side of the document contains an error. The error will be considered resolved when the file with the front side changes
type PassportElementErrorSourceFrontSide struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceFrontSide
func (passportElementErrorSourceFrontSide *PassportElementErrorSourceFrontSide) MessageType() string {
	return "passportElementErrorSourceFrontSide"
}

// NewPassportElementErrorSourceFrontSide creates a new PassportElementErrorSourceFrontSide
//
func NewPassportElementErrorSourceFrontSide() *PassportElementErrorSourceFrontSide {
	passportElementErrorSourceFrontSideTemp := PassportElementErrorSourceFrontSide{
		tdCommon: tdCommon{Type: "passportElementErrorSourceFrontSide"},
	}

	return &passportElementErrorSourceFrontSideTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFrontSide *PassportElementErrorSourceFrontSide) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFrontSideType
}

// PassportElementErrorSourceReverseSide The reverse side of the document contains an error. The error will be considered resolved when the file with the reverse side changes
type PassportElementErrorSourceReverseSide struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceReverseSide
func (passportElementErrorSourceReverseSide *PassportElementErrorSourceReverseSide) MessageType() string {
	return "passportElementErrorSourceReverseSide"
}

// NewPassportElementErrorSourceReverseSide creates a new PassportElementErrorSourceReverseSide
//
func NewPassportElementErrorSourceReverseSide() *PassportElementErrorSourceReverseSide {
	passportElementErrorSourceReverseSideTemp := PassportElementErrorSourceReverseSide{
		tdCommon: tdCommon{Type: "passportElementErrorSourceReverseSide"},
	}

	return &passportElementErrorSourceReverseSideTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceReverseSide *PassportElementErrorSourceReverseSide) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceReverseSideType
}

// PassportElementErrorSourceSelfie The selfie with the document contains an error. The error will be considered resolved when the file with the selfie changes
type PassportElementErrorSourceSelfie struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceSelfie
func (passportElementErrorSourceSelfie *PassportElementErrorSourceSelfie) MessageType() string {
	return "passportElementErrorSourceSelfie"
}

// NewPassportElementErrorSourceSelfie creates a new PassportElementErrorSourceSelfie
//
func NewPassportElementErrorSourceSelfie() *PassportElementErrorSourceSelfie {
	passportElementErrorSourceSelfieTemp := PassportElementErrorSourceSelfie{
		tdCommon: tdCommon{Type: "passportElementErrorSourceSelfie"},
	}

	return &passportElementErrorSourceSelfieTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceSelfie *PassportElementErrorSourceSelfie) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceSelfieType
}

// PassportElementErrorSourceTranslationFile One of files with the translation of the document contains an error. The error will be considered resolved when the file changes
type PassportElementErrorSourceTranslationFile struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceTranslationFile
func (passportElementErrorSourceTranslationFile *PassportElementErrorSourceTranslationFile) MessageType() string {
	return "passportElementErrorSourceTranslationFile"
}

// NewPassportElementErrorSourceTranslationFile creates a new PassportElementErrorSourceTranslationFile
//
func NewPassportElementErrorSourceTranslationFile() *PassportElementErrorSourceTranslationFile {
	passportElementErrorSourceTranslationFileTemp := PassportElementErrorSourceTranslationFile{
		tdCommon: tdCommon{Type: "passportElementErrorSourceTranslationFile"},
	}

	return &passportElementErrorSourceTranslationFileTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceTranslationFile *PassportElementErrorSourceTranslationFile) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceTranslationFileType
}

// PassportElementErrorSourceTranslationFiles The translation of the document contains an error. The error will be considered resolved when the list of translation files changes
type PassportElementErrorSourceTranslationFiles struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceTranslationFiles
func (passportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles) MessageType() string {
	return "passportElementErrorSourceTranslationFiles"
}

// NewPassportElementErrorSourceTranslationFiles creates a new PassportElementErrorSourceTranslationFiles
//
func NewPassportElementErrorSourceTranslationFiles() *PassportElementErrorSourceTranslationFiles {
	passportElementErrorSourceTranslationFilesTemp := PassportElementErrorSourceTranslationFiles{
		tdCommon: tdCommon{Type: "passportElementErrorSourceTranslationFiles"},
	}

	return &passportElementErrorSourceTranslationFilesTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceTranslationFiles *PassportElementErrorSourceTranslationFiles) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceTranslationFilesType
}

// PassportElementErrorSourceFile The file contains an error. The error will be considered resolved when the file changes
type PassportElementErrorSourceFile struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceFile
func (passportElementErrorSourceFile *PassportElementErrorSourceFile) MessageType() string {
	return "passportElementErrorSourceFile"
}

// NewPassportElementErrorSourceFile creates a new PassportElementErrorSourceFile
//
func NewPassportElementErrorSourceFile() *PassportElementErrorSourceFile {
	passportElementErrorSourceFileTemp := PassportElementErrorSourceFile{
		tdCommon: tdCommon{Type: "passportElementErrorSourceFile"},
	}

	return &passportElementErrorSourceFileTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFile *PassportElementErrorSourceFile) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFileType
}

// PassportElementErrorSourceFiles The list of attached files contains an error. The error will be considered resolved when the list of files changes
type PassportElementErrorSourceFiles struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementErrorSourceFiles
func (passportElementErrorSourceFiles *PassportElementErrorSourceFiles) MessageType() string {
	return "passportElementErrorSourceFiles"
}

// NewPassportElementErrorSourceFiles creates a new PassportElementErrorSourceFiles
//
func NewPassportElementErrorSourceFiles() *PassportElementErrorSourceFiles {
	passportElementErrorSourceFilesTemp := PassportElementErrorSourceFiles{
		tdCommon: tdCommon{Type: "passportElementErrorSourceFiles"},
	}

	return &passportElementErrorSourceFilesTemp
}

// GetPassportElementErrorSourceEnum return the enum type of this object
func (passportElementErrorSourceFiles *PassportElementErrorSourceFiles) GetPassportElementErrorSourceEnum() PassportElementErrorSourceEnum {
	return PassportElementErrorSourceFilesType
}

// PassportElementError Contains the description of an error in a Telegram Passport element
type PassportElementError struct {
	tdCommon
	Type    PassportElementType        `json:"type"`    // Type of the Telegram Passport element which has the error
	Message string                     `json:"message"` // Error message
	Source  PassportElementErrorSource `json:"source"`  // Error source
}

// MessageType return the string telegram-type of PassportElementError
func (passportElementError *PassportElementError) MessageType() string {
	return "passportElementError"
}

// NewPassportElementError creates a new PassportElementError
//
// @param typeParam Type of the Telegram Passport element which has the error
// @param message Error message
// @param source Error source
func NewPassportElementError(typeParam PassportElementType, message string, source PassportElementErrorSource) *PassportElementError {
	passportElementErrorTemp := PassportElementError{
		tdCommon: tdCommon{Type: "passportElementError"},
		Type:     typeParam,
		Message:  message,
		Source:   source,
	}

	return &passportElementErrorTemp
}

// UnmarshalJSON unmarshal to json
func (passportElementError *PassportElementError) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Message string `json:"message"` // Error message

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	passportElementError.tdCommon = tempObj.tdCommon
	passportElementError.Message = tempObj.Message

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	passportElementError.Type = fieldType

	fieldSource, _ := unmarshalPassportElementErrorSource(objMap["source"])
	passportElementError.Source = fieldSource

	return nil
}

// PassportSuitableElement Contains information about a Telegram Passport element that was requested by a service
type PassportSuitableElement struct {
	tdCommon
	Type                  PassportElementType `json:"type"`                    // Type of the element
	IsSelfieRequired      bool                `json:"is_selfie_required"`      // True, if a selfie is required with the identity document
	IsTranslationRequired bool                `json:"is_translation_required"` // True, if a certified English translation is required with the document
	IsNativeNameRequired  bool                `json:"is_native_name_required"` // True, if personal details must include the user's name in the language of their country of residence
}

// MessageType return the string telegram-type of PassportSuitableElement
func (passportSuitableElement *PassportSuitableElement) MessageType() string {
	return "passportSuitableElement"
}

// NewPassportSuitableElement creates a new PassportSuitableElement
//
// @param typeParam Type of the element
// @param isSelfieRequired True, if a selfie is required with the identity document
// @param isTranslationRequired True, if a certified English translation is required with the document
// @param isNativeNameRequired True, if personal details must include the user's name in the language of their country of residence
func NewPassportSuitableElement(typeParam PassportElementType, isSelfieRequired bool, isTranslationRequired bool, isNativeNameRequired bool) *PassportSuitableElement {
	passportSuitableElementTemp := PassportSuitableElement{
		tdCommon:              tdCommon{Type: "passportSuitableElement"},
		Type:                  typeParam,
		IsSelfieRequired:      isSelfieRequired,
		IsTranslationRequired: isTranslationRequired,
		IsNativeNameRequired:  isNativeNameRequired,
	}

	return &passportSuitableElementTemp
}

// UnmarshalJSON unmarshal to json
func (passportSuitableElement *PassportSuitableElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		IsSelfieRequired      bool `json:"is_selfie_required"`      // True, if a selfie is required with the identity document
		IsTranslationRequired bool `json:"is_translation_required"` // True, if a certified English translation is required with the document
		IsNativeNameRequired  bool `json:"is_native_name_required"` // True, if personal details must include the user's name in the language of their country of residence
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	passportSuitableElement.tdCommon = tempObj.tdCommon
	passportSuitableElement.IsSelfieRequired = tempObj.IsSelfieRequired
	passportSuitableElement.IsTranslationRequired = tempObj.IsTranslationRequired
	passportSuitableElement.IsNativeNameRequired = tempObj.IsNativeNameRequired

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	passportSuitableElement.Type = fieldType

	return nil
}

// PassportRequiredElement Contains a description of the required Telegram Passport element that was requested by a service
type PassportRequiredElement struct {
	tdCommon
	SuitableElements []PassportSuitableElement `json:"suitable_elements"` // List of Telegram Passport elements any of which is enough to provide
}

// MessageType return the string telegram-type of PassportRequiredElement
func (passportRequiredElement *PassportRequiredElement) MessageType() string {
	return "passportRequiredElement"
}

// NewPassportRequiredElement creates a new PassportRequiredElement
//
// @param suitableElements List of Telegram Passport elements any of which is enough to provide
func NewPassportRequiredElement(suitableElements []PassportSuitableElement) *PassportRequiredElement {
	passportRequiredElementTemp := PassportRequiredElement{
		tdCommon:         tdCommon{Type: "passportRequiredElement"},
		SuitableElements: suitableElements,
	}

	return &passportRequiredElementTemp
}

// PassportAuthorizationForm Contains information about a Telegram Passport authorization form that was requested
type PassportAuthorizationForm struct {
	tdCommon
	ID               int32                     `json:"id"`                 // Unique identifier of the authorization form
	RequiredElements []PassportRequiredElement `json:"required_elements"`  // Information about the Telegram Passport elements that need to be provided to complete the form
	Elements         []PassportElement         `json:"elements"`           // Already available Telegram Passport elements
	Errors           []PassportElementError    `json:"errors"`             // Errors in the elements that are already available
	PrivacyPolicyURL string                    `json:"privacy_policy_url"` // URL for the privacy policy of the service; can be empty
}

// MessageType return the string telegram-type of PassportAuthorizationForm
func (passportAuthorizationForm *PassportAuthorizationForm) MessageType() string {
	return "passportAuthorizationForm"
}

// NewPassportAuthorizationForm creates a new PassportAuthorizationForm
//
// @param iD Unique identifier of the authorization form
// @param requiredElements Information about the Telegram Passport elements that need to be provided to complete the form
// @param elements Already available Telegram Passport elements
// @param errors Errors in the elements that are already available
// @param privacyPolicyURL URL for the privacy policy of the service; can be empty
func NewPassportAuthorizationForm(iD int32, requiredElements []PassportRequiredElement, elements []PassportElement, errors []PassportElementError, privacyPolicyURL string) *PassportAuthorizationForm {
	passportAuthorizationFormTemp := PassportAuthorizationForm{
		tdCommon:         tdCommon{Type: "passportAuthorizationForm"},
		ID:               iD,
		RequiredElements: requiredElements,
		Elements:         elements,
		Errors:           errors,
		PrivacyPolicyURL: privacyPolicyURL,
	}

	return &passportAuthorizationFormTemp
}

// EncryptedCredentials Contains encrypted Telegram Passport data credentials
type EncryptedCredentials struct {
	tdCommon
	Data   []byte `json:"data"`   // The encrypted credentials
	Hash   []byte `json:"hash"`   // The decrypted data hash
	Secret []byte `json:"secret"` // Secret for data decryption, encrypted with the service's public key
}

// MessageType return the string telegram-type of EncryptedCredentials
func (encryptedCredentials *EncryptedCredentials) MessageType() string {
	return "encryptedCredentials"
}

// NewEncryptedCredentials creates a new EncryptedCredentials
//
// @param data The encrypted credentials
// @param hash The decrypted data hash
// @param secret Secret for data decryption, encrypted with the service's public key
func NewEncryptedCredentials(data []byte, hash []byte, secret []byte) *EncryptedCredentials {
	encryptedCredentialsTemp := EncryptedCredentials{
		tdCommon: tdCommon{Type: "encryptedCredentials"},
		Data:     data,
		Hash:     hash,
		Secret:   secret,
	}

	return &encryptedCredentialsTemp
}

// EncryptedPassportElement Contains information about an encrypted Telegram Passport element; for bots only
type EncryptedPassportElement struct {
	tdCommon
	Type        PassportElementType `json:"type"`         // Type of Telegram Passport element
	Data        []byte              `json:"data"`         // Encrypted JSON-encoded data about the user
	FrontSide   *DatedFile          `json:"front_side"`   // The front side of an identity document
	ReverseSide *DatedFile          `json:"reverse_side"` // The reverse side of an identity document; may be null
	Selfie      *DatedFile          `json:"selfie"`       // Selfie with the document; may be null
	Translation []DatedFile         `json:"translation"`  // List of files containing a certified English translation of the document
	Files       []DatedFile         `json:"files"`        // List of attached files
	Value       string              `json:"value"`        // Unencrypted data, phone number or email address
	Hash        string              `json:"hash"`         // Hash of the entire element
}

// MessageType return the string telegram-type of EncryptedPassportElement
func (encryptedPassportElement *EncryptedPassportElement) MessageType() string {
	return "encryptedPassportElement"
}

// NewEncryptedPassportElement creates a new EncryptedPassportElement
//
// @param typeParam Type of Telegram Passport element
// @param data Encrypted JSON-encoded data about the user
// @param frontSide The front side of an identity document
// @param reverseSide The reverse side of an identity document; may be null
// @param selfie Selfie with the document; may be null
// @param translation List of files containing a certified English translation of the document
// @param files List of attached files
// @param value Unencrypted data, phone number or email address
// @param hash Hash of the entire element
func NewEncryptedPassportElement(typeParam PassportElementType, data []byte, frontSide *DatedFile, reverseSide *DatedFile, selfie *DatedFile, translation []DatedFile, files []DatedFile, value string, hash string) *EncryptedPassportElement {
	encryptedPassportElementTemp := EncryptedPassportElement{
		tdCommon:    tdCommon{Type: "encryptedPassportElement"},
		Type:        typeParam,
		Data:        data,
		FrontSide:   frontSide,
		ReverseSide: reverseSide,
		Selfie:      selfie,
		Translation: translation,
		Files:       files,
		Value:       value,
		Hash:        hash,
	}

	return &encryptedPassportElementTemp
}

// UnmarshalJSON unmarshal to json
func (encryptedPassportElement *EncryptedPassportElement) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Data        []byte      `json:"data"`         // Encrypted JSON-encoded data about the user
		FrontSide   *DatedFile  `json:"front_side"`   // The front side of an identity document
		ReverseSide *DatedFile  `json:"reverse_side"` // The reverse side of an identity document; may be null
		Selfie      *DatedFile  `json:"selfie"`       // Selfie with the document; may be null
		Translation []DatedFile `json:"translation"`  // List of files containing a certified English translation of the document
		Files       []DatedFile `json:"files"`        // List of attached files
		Value       string      `json:"value"`        // Unencrypted data, phone number or email address
		Hash        string      `json:"hash"`         // Hash of the entire element
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	encryptedPassportElement.tdCommon = tempObj.tdCommon
	encryptedPassportElement.Data = tempObj.Data
	encryptedPassportElement.FrontSide = tempObj.FrontSide
	encryptedPassportElement.ReverseSide = tempObj.ReverseSide
	encryptedPassportElement.Selfie = tempObj.Selfie
	encryptedPassportElement.Translation = tempObj.Translation
	encryptedPassportElement.Files = tempObj.Files
	encryptedPassportElement.Value = tempObj.Value
	encryptedPassportElement.Hash = tempObj.Hash

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	encryptedPassportElement.Type = fieldType

	return nil
}

// InputPassportElementErrorSourceUnspecified The element contains an error in an unspecified place. The error will be considered resolved when new data is added
type InputPassportElementErrorSourceUnspecified struct {
	tdCommon
	ElementHash []byte `json:"element_hash"` // Current hash of the entire element
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceUnspecified
func (inputPassportElementErrorSourceUnspecified *InputPassportElementErrorSourceUnspecified) MessageType() string {
	return "inputPassportElementErrorSourceUnspecified"
}

// NewInputPassportElementErrorSourceUnspecified creates a new InputPassportElementErrorSourceUnspecified
//
// @param elementHash Current hash of the entire element
func NewInputPassportElementErrorSourceUnspecified(elementHash []byte) *InputPassportElementErrorSourceUnspecified {
	inputPassportElementErrorSourceUnspecifiedTemp := InputPassportElementErrorSourceUnspecified{
		tdCommon:    tdCommon{Type: "inputPassportElementErrorSourceUnspecified"},
		ElementHash: elementHash,
	}

	return &inputPassportElementErrorSourceUnspecifiedTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceUnspecified *InputPassportElementErrorSourceUnspecified) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceUnspecifiedType
}

// InputPassportElementErrorSourceDataField A data field contains an error. The error is considered resolved when the field's value changes
type InputPassportElementErrorSourceDataField struct {
	tdCommon
	FieldName string `json:"field_name"` // Field name
	DataHash  []byte `json:"data_hash"`  // Current data hash
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceDataField
func (inputPassportElementErrorSourceDataField *InputPassportElementErrorSourceDataField) MessageType() string {
	return "inputPassportElementErrorSourceDataField"
}

// NewInputPassportElementErrorSourceDataField creates a new InputPassportElementErrorSourceDataField
//
// @param fieldName Field name
// @param dataHash Current data hash
func NewInputPassportElementErrorSourceDataField(fieldName string, dataHash []byte) *InputPassportElementErrorSourceDataField {
	inputPassportElementErrorSourceDataFieldTemp := InputPassportElementErrorSourceDataField{
		tdCommon:  tdCommon{Type: "inputPassportElementErrorSourceDataField"},
		FieldName: fieldName,
		DataHash:  dataHash,
	}

	return &inputPassportElementErrorSourceDataFieldTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceDataField *InputPassportElementErrorSourceDataField) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceDataFieldType
}

// InputPassportElementErrorSourceFrontSide The front side of the document contains an error. The error is considered resolved when the file with the front side of the document changes
type InputPassportElementErrorSourceFrontSide struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the front side
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFrontSide
func (inputPassportElementErrorSourceFrontSide *InputPassportElementErrorSourceFrontSide) MessageType() string {
	return "inputPassportElementErrorSourceFrontSide"
}

// NewInputPassportElementErrorSourceFrontSide creates a new InputPassportElementErrorSourceFrontSide
//
// @param fileHash Current hash of the file containing the front side
func NewInputPassportElementErrorSourceFrontSide(fileHash []byte) *InputPassportElementErrorSourceFrontSide {
	inputPassportElementErrorSourceFrontSideTemp := InputPassportElementErrorSourceFrontSide{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceFrontSide"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceFrontSideTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFrontSide *InputPassportElementErrorSourceFrontSide) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFrontSideType
}

// InputPassportElementErrorSourceReverseSide The reverse side of the document contains an error. The error is considered resolved when the file with the reverse side of the document changes
type InputPassportElementErrorSourceReverseSide struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the reverse side
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceReverseSide
func (inputPassportElementErrorSourceReverseSide *InputPassportElementErrorSourceReverseSide) MessageType() string {
	return "inputPassportElementErrorSourceReverseSide"
}

// NewInputPassportElementErrorSourceReverseSide creates a new InputPassportElementErrorSourceReverseSide
//
// @param fileHash Current hash of the file containing the reverse side
func NewInputPassportElementErrorSourceReverseSide(fileHash []byte) *InputPassportElementErrorSourceReverseSide {
	inputPassportElementErrorSourceReverseSideTemp := InputPassportElementErrorSourceReverseSide{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceReverseSide"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceReverseSideTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceReverseSide *InputPassportElementErrorSourceReverseSide) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceReverseSideType
}

// InputPassportElementErrorSourceSelfie The selfie contains an error. The error is considered resolved when the file with the selfie changes
type InputPassportElementErrorSourceSelfie struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the selfie
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceSelfie
func (inputPassportElementErrorSourceSelfie *InputPassportElementErrorSourceSelfie) MessageType() string {
	return "inputPassportElementErrorSourceSelfie"
}

// NewInputPassportElementErrorSourceSelfie creates a new InputPassportElementErrorSourceSelfie
//
// @param fileHash Current hash of the file containing the selfie
func NewInputPassportElementErrorSourceSelfie(fileHash []byte) *InputPassportElementErrorSourceSelfie {
	inputPassportElementErrorSourceSelfieTemp := InputPassportElementErrorSourceSelfie{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceSelfie"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceSelfieTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceSelfie *InputPassportElementErrorSourceSelfie) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceSelfieType
}

// InputPassportElementErrorSourceTranslationFile One of the files containing the translation of the document contains an error. The error is considered resolved when the file with the translation changes
type InputPassportElementErrorSourceTranslationFile struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file containing the translation
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceTranslationFile
func (inputPassportElementErrorSourceTranslationFile *InputPassportElementErrorSourceTranslationFile) MessageType() string {
	return "inputPassportElementErrorSourceTranslationFile"
}

// NewInputPassportElementErrorSourceTranslationFile creates a new InputPassportElementErrorSourceTranslationFile
//
// @param fileHash Current hash of the file containing the translation
func NewInputPassportElementErrorSourceTranslationFile(fileHash []byte) *InputPassportElementErrorSourceTranslationFile {
	inputPassportElementErrorSourceTranslationFileTemp := InputPassportElementErrorSourceTranslationFile{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceTranslationFile"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceTranslationFileTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceTranslationFile *InputPassportElementErrorSourceTranslationFile) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceTranslationFileType
}

// InputPassportElementErrorSourceTranslationFiles The translation of the document contains an error. The error is considered resolved when the list of files changes
type InputPassportElementErrorSourceTranslationFiles struct {
	tdCommon
	FileHashes [][]byte `json:"file_hashes"` // Current hashes of all files with the translation
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceTranslationFiles
func (inputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles) MessageType() string {
	return "inputPassportElementErrorSourceTranslationFiles"
}

// NewInputPassportElementErrorSourceTranslationFiles creates a new InputPassportElementErrorSourceTranslationFiles
//
// @param fileHashes Current hashes of all files with the translation
func NewInputPassportElementErrorSourceTranslationFiles(fileHashes [][]byte) *InputPassportElementErrorSourceTranslationFiles {
	inputPassportElementErrorSourceTranslationFilesTemp := InputPassportElementErrorSourceTranslationFiles{
		tdCommon:   tdCommon{Type: "inputPassportElementErrorSourceTranslationFiles"},
		FileHashes: fileHashes,
	}

	return &inputPassportElementErrorSourceTranslationFilesTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceTranslationFiles *InputPassportElementErrorSourceTranslationFiles) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceTranslationFilesType
}

// InputPassportElementErrorSourceFile The file contains an error. The error is considered resolved when the file changes
type InputPassportElementErrorSourceFile struct {
	tdCommon
	FileHash []byte `json:"file_hash"` // Current hash of the file which has the error
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFile
func (inputPassportElementErrorSourceFile *InputPassportElementErrorSourceFile) MessageType() string {
	return "inputPassportElementErrorSourceFile"
}

// NewInputPassportElementErrorSourceFile creates a new InputPassportElementErrorSourceFile
//
// @param fileHash Current hash of the file which has the error
func NewInputPassportElementErrorSourceFile(fileHash []byte) *InputPassportElementErrorSourceFile {
	inputPassportElementErrorSourceFileTemp := InputPassportElementErrorSourceFile{
		tdCommon: tdCommon{Type: "inputPassportElementErrorSourceFile"},
		FileHash: fileHash,
	}

	return &inputPassportElementErrorSourceFileTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFile *InputPassportElementErrorSourceFile) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFileType
}

// InputPassportElementErrorSourceFiles The list of attached files contains an error. The error is considered resolved when the file list changes
type InputPassportElementErrorSourceFiles struct {
	tdCommon
	FileHashes [][]byte `json:"file_hashes"` // Current hashes of all attached files
}

// MessageType return the string telegram-type of InputPassportElementErrorSourceFiles
func (inputPassportElementErrorSourceFiles *InputPassportElementErrorSourceFiles) MessageType() string {
	return "inputPassportElementErrorSourceFiles"
}

// NewInputPassportElementErrorSourceFiles creates a new InputPassportElementErrorSourceFiles
//
// @param fileHashes Current hashes of all attached files
func NewInputPassportElementErrorSourceFiles(fileHashes [][]byte) *InputPassportElementErrorSourceFiles {
	inputPassportElementErrorSourceFilesTemp := InputPassportElementErrorSourceFiles{
		tdCommon:   tdCommon{Type: "inputPassportElementErrorSourceFiles"},
		FileHashes: fileHashes,
	}

	return &inputPassportElementErrorSourceFilesTemp
}

// GetInputPassportElementErrorSourceEnum return the enum type of this object
func (inputPassportElementErrorSourceFiles *InputPassportElementErrorSourceFiles) GetInputPassportElementErrorSourceEnum() InputPassportElementErrorSourceEnum {
	return InputPassportElementErrorSourceFilesType
}

// InputPassportElementError Contains the description of an error in a Telegram Passport element; for bots only
type InputPassportElementError struct {
	tdCommon
	Type    PassportElementType             `json:"type"`    // Type of Telegram Passport element that has the error
	Message string                          `json:"message"` // Error message
	Source  InputPassportElementErrorSource `json:"source"`  // Error source
}

// MessageType return the string telegram-type of InputPassportElementError
func (inputPassportElementError *InputPassportElementError) MessageType() string {
	return "inputPassportElementError"
}

// NewInputPassportElementError creates a new InputPassportElementError
//
// @param typeParam Type of Telegram Passport element that has the error
// @param message Error message
// @param source Error source
func NewInputPassportElementError(typeParam PassportElementType, message string, source InputPassportElementErrorSource) *InputPassportElementError {
	inputPassportElementErrorTemp := InputPassportElementError{
		tdCommon: tdCommon{Type: "inputPassportElementError"},
		Type:     typeParam,
		Message:  message,
		Source:   source,
	}

	return &inputPassportElementErrorTemp
}

// UnmarshalJSON unmarshal to json
func (inputPassportElementError *InputPassportElementError) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Message string `json:"message"` // Error message

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	inputPassportElementError.tdCommon = tempObj.tdCommon
	inputPassportElementError.Message = tempObj.Message

	fieldType, _ := unmarshalPassportElementType(objMap["type"])
	inputPassportElementError.Type = fieldType

	fieldSource, _ := unmarshalInputPassportElementErrorSource(objMap["source"])
	inputPassportElementError.Source = fieldSource

	return nil
}

// MessageText A text message
type MessageText struct {
	tdCommon
	Text    *FormattedText `json:"text"`     // Text of the message
	WebPage *WebPage       `json:"web_page"` // A preview of the web page that's mentioned in the text; may be null
}

// MessageType return the string telegram-type of MessageText
func (messageText *MessageText) MessageType() string {
	return "messageText"
}

// NewMessageText creates a new MessageText
//
// @param text Text of the message
// @param webPage A preview of the web page that's mentioned in the text; may be null
func NewMessageText(text *FormattedText, webPage *WebPage) *MessageText {
	messageTextTemp := MessageText{
		tdCommon: tdCommon{Type: "messageText"},
		Text:     text,
		WebPage:  webPage,
	}

	return &messageTextTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageText *MessageText) GetMessageContentEnum() MessageContentEnum {
	return MessageTextType
}

// MessageAnimation An animation message (GIF-style).
type MessageAnimation struct {
	tdCommon
	Animation *Animation     `json:"animation"` // Message content
	Caption   *FormattedText `json:"caption"`   // Animation caption
	IsSecret  bool           `json:"is_secret"` // True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
}

// MessageType return the string telegram-type of MessageAnimation
func (messageAnimation *MessageAnimation) MessageType() string {
	return "messageAnimation"
}

// NewMessageAnimation creates a new MessageAnimation
//
// @param animation Message content
// @param caption Animation caption
// @param isSecret True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
func NewMessageAnimation(animation *Animation, caption *FormattedText, isSecret bool) *MessageAnimation {
	messageAnimationTemp := MessageAnimation{
		tdCommon:  tdCommon{Type: "messageAnimation"},
		Animation: animation,
		Caption:   caption,
		IsSecret:  isSecret,
	}

	return &messageAnimationTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageAnimation *MessageAnimation) GetMessageContentEnum() MessageContentEnum {
	return MessageAnimationType
}

// MessageAudio An audio message
type MessageAudio struct {
	tdCommon
	Audio   *Audio         `json:"audio"`   // Message content
	Caption *FormattedText `json:"caption"` // Audio caption
}

// MessageType return the string telegram-type of MessageAudio
func (messageAudio *MessageAudio) MessageType() string {
	return "messageAudio"
}

// NewMessageAudio creates a new MessageAudio
//
// @param audio Message content
// @param caption Audio caption
func NewMessageAudio(audio *Audio, caption *FormattedText) *MessageAudio {
	messageAudioTemp := MessageAudio{
		tdCommon: tdCommon{Type: "messageAudio"},
		Audio:    audio,
		Caption:  caption,
	}

	return &messageAudioTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageAudio *MessageAudio) GetMessageContentEnum() MessageContentEnum {
	return MessageAudioType
}

// MessageDocument A document message (general file)
type MessageDocument struct {
	tdCommon
	Document *Document      `json:"document"` // Message content
	Caption  *FormattedText `json:"caption"`  // Document caption
}

// MessageType return the string telegram-type of MessageDocument
func (messageDocument *MessageDocument) MessageType() string {
	return "messageDocument"
}

// NewMessageDocument creates a new MessageDocument
//
// @param document Message content
// @param caption Document caption
func NewMessageDocument(document *Document, caption *FormattedText) *MessageDocument {
	messageDocumentTemp := MessageDocument{
		tdCommon: tdCommon{Type: "messageDocument"},
		Document: document,
		Caption:  caption,
	}

	return &messageDocumentTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageDocument *MessageDocument) GetMessageContentEnum() MessageContentEnum {
	return MessageDocumentType
}

// MessagePhoto A photo message
type MessagePhoto struct {
	tdCommon
	Photo    *Photo         `json:"photo"`     // Message content
	Caption  *FormattedText `json:"caption"`   // Photo caption
	IsSecret bool           `json:"is_secret"` // True, if the photo must be blurred and must be shown only while tapped
}

// MessageType return the string telegram-type of MessagePhoto
func (messagePhoto *MessagePhoto) MessageType() string {
	return "messagePhoto"
}

// NewMessagePhoto creates a new MessagePhoto
//
// @param photo Message content
// @param caption Photo caption
// @param isSecret True, if the photo must be blurred and must be shown only while tapped
func NewMessagePhoto(photo *Photo, caption *FormattedText, isSecret bool) *MessagePhoto {
	messagePhotoTemp := MessagePhoto{
		tdCommon: tdCommon{Type: "messagePhoto"},
		Photo:    photo,
		Caption:  caption,
		IsSecret: isSecret,
	}

	return &messagePhotoTemp
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

// NewMessageExpiredPhoto creates a new MessageExpiredPhoto
//
func NewMessageExpiredPhoto() *MessageExpiredPhoto {
	messageExpiredPhotoTemp := MessageExpiredPhoto{
		tdCommon: tdCommon{Type: "messageExpiredPhoto"},
	}

	return &messageExpiredPhotoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredPhoto *MessageExpiredPhoto) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredPhotoType
}

// MessageSticker A sticker message
type MessageSticker struct {
	tdCommon
	Sticker *Sticker `json:"sticker"` // Message content
}

// MessageType return the string telegram-type of MessageSticker
func (messageSticker *MessageSticker) MessageType() string {
	return "messageSticker"
}

// NewMessageSticker creates a new MessageSticker
//
// @param sticker Message content
func NewMessageSticker(sticker *Sticker) *MessageSticker {
	messageStickerTemp := MessageSticker{
		tdCommon: tdCommon{Type: "messageSticker"},
		Sticker:  sticker,
	}

	return &messageStickerTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageSticker *MessageSticker) GetMessageContentEnum() MessageContentEnum {
	return MessageStickerType
}

// MessageVideo A video message
type MessageVideo struct {
	tdCommon
	Video    *Video         `json:"video"`     // Message content
	Caption  *FormattedText `json:"caption"`   // Video caption
	IsSecret bool           `json:"is_secret"` // True, if the video thumbnail must be blurred and the video must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideo
func (messageVideo *MessageVideo) MessageType() string {
	return "messageVideo"
}

// NewMessageVideo creates a new MessageVideo
//
// @param video Message content
// @param caption Video caption
// @param isSecret True, if the video thumbnail must be blurred and the video must be shown only while tapped
func NewMessageVideo(video *Video, caption *FormattedText, isSecret bool) *MessageVideo {
	messageVideoTemp := MessageVideo{
		tdCommon: tdCommon{Type: "messageVideo"},
		Video:    video,
		Caption:  caption,
		IsSecret: isSecret,
	}

	return &messageVideoTemp
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

// NewMessageExpiredVideo creates a new MessageExpiredVideo
//
func NewMessageExpiredVideo() *MessageExpiredVideo {
	messageExpiredVideoTemp := MessageExpiredVideo{
		tdCommon: tdCommon{Type: "messageExpiredVideo"},
	}

	return &messageExpiredVideoTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageExpiredVideo *MessageExpiredVideo) GetMessageContentEnum() MessageContentEnum {
	return MessageExpiredVideoType
}

// MessageVideoNote A video note message
type MessageVideoNote struct {
	tdCommon
	VideoNote *VideoNote `json:"video_note"` // Message content
	IsViewed  bool       `json:"is_viewed"`  // True, if at least one of the recipients has viewed the video note
	IsSecret  bool       `json:"is_secret"`  // True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideoNote
func (messageVideoNote *MessageVideoNote) MessageType() string {
	return "messageVideoNote"
}

// NewMessageVideoNote creates a new MessageVideoNote
//
// @param videoNote Message content
// @param isViewed True, if at least one of the recipients has viewed the video note
// @param isSecret True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
func NewMessageVideoNote(videoNote *VideoNote, isViewed bool, isSecret bool) *MessageVideoNote {
	messageVideoNoteTemp := MessageVideoNote{
		tdCommon:  tdCommon{Type: "messageVideoNote"},
		VideoNote: videoNote,
		IsViewed:  isViewed,
		IsSecret:  isSecret,
	}

	return &messageVideoNoteTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVideoNote *MessageVideoNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVideoNoteType
}

// MessageVoiceNote A voice note message
type MessageVoiceNote struct {
	tdCommon
	VoiceNote  *VoiceNote     `json:"voice_note"`  // Message content
	Caption    *FormattedText `json:"caption"`     // Voice note caption
	IsListened bool           `json:"is_listened"` // True, if at least one of the recipients has listened to the voice note
}

// MessageType return the string telegram-type of MessageVoiceNote
func (messageVoiceNote *MessageVoiceNote) MessageType() string {
	return "messageVoiceNote"
}

// NewMessageVoiceNote creates a new MessageVoiceNote
//
// @param voiceNote Message content
// @param caption Voice note caption
// @param isListened True, if at least one of the recipients has listened to the voice note
func NewMessageVoiceNote(voiceNote *VoiceNote, caption *FormattedText, isListened bool) *MessageVoiceNote {
	messageVoiceNoteTemp := MessageVoiceNote{
		tdCommon:   tdCommon{Type: "messageVoiceNote"},
		VoiceNote:  voiceNote,
		Caption:    caption,
		IsListened: isListened,
	}

	return &messageVoiceNoteTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVoiceNote *MessageVoiceNote) GetMessageContentEnum() MessageContentEnum {
	return MessageVoiceNoteType
}

// MessageLocation A message with a location
type MessageLocation struct {
	tdCommon
	Location   *Location `json:"location"`    // Message content
	LivePeriod int32     `json:"live_period"` // Time relative to the message sent date until which the location can be updated, in seconds
	ExpiresIn  int32     `json:"expires_in"`  // Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
}

// MessageType return the string telegram-type of MessageLocation
func (messageLocation *MessageLocation) MessageType() string {
	return "messageLocation"
}

// NewMessageLocation creates a new MessageLocation
//
// @param location Message content
// @param livePeriod Time relative to the message sent date until which the location can be updated, in seconds
// @param expiresIn Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
func NewMessageLocation(location *Location, livePeriod int32, expiresIn int32) *MessageLocation {
	messageLocationTemp := MessageLocation{
		tdCommon:   tdCommon{Type: "messageLocation"},
		Location:   location,
		LivePeriod: livePeriod,
		ExpiresIn:  expiresIn,
	}

	return &messageLocationTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageLocation *MessageLocation) GetMessageContentEnum() MessageContentEnum {
	return MessageLocationType
}

// MessageVenue A message with information about a venue
type MessageVenue struct {
	tdCommon
	Venue *Venue `json:"venue"` // Message content
}

// MessageType return the string telegram-type of MessageVenue
func (messageVenue *MessageVenue) MessageType() string {
	return "messageVenue"
}

// NewMessageVenue creates a new MessageVenue
//
// @param venue Message content
func NewMessageVenue(venue *Venue) *MessageVenue {
	messageVenueTemp := MessageVenue{
		tdCommon: tdCommon{Type: "messageVenue"},
		Venue:    venue,
	}

	return &messageVenueTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageVenue *MessageVenue) GetMessageContentEnum() MessageContentEnum {
	return MessageVenueType
}

// MessageContact A message with a user contact
type MessageContact struct {
	tdCommon
	Contact *Contact `json:"contact"` // Message content
}

// MessageType return the string telegram-type of MessageContact
func (messageContact *MessageContact) MessageType() string {
	return "messageContact"
}

// NewMessageContact creates a new MessageContact
//
// @param contact Message content
func NewMessageContact(contact *Contact) *MessageContact {
	messageContactTemp := MessageContact{
		tdCommon: tdCommon{Type: "messageContact"},
		Contact:  contact,
	}

	return &messageContactTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageContact *MessageContact) GetMessageContentEnum() MessageContentEnum {
	return MessageContactType
}

// MessageGame A message with a game
type MessageGame struct {
	tdCommon
	Game *Game `json:"game"` // Game
}

// MessageType return the string telegram-type of MessageGame
func (messageGame *MessageGame) MessageType() string {
	return "messageGame"
}

// NewMessageGame creates a new MessageGame
//
// @param game Game
func NewMessageGame(game *Game) *MessageGame {
	messageGameTemp := MessageGame{
		tdCommon: tdCommon{Type: "messageGame"},
		Game:     game,
	}

	return &messageGameTemp
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
	Photo               *Photo `json:"photo"`                 // Product photo; may be null
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

// NewMessageInvoice creates a new MessageInvoice
//
// @param title Product title
// @param description
// @param photo Product photo; may be null
// @param currency Currency for the product price
// @param totalAmount Product total price in the minimal quantity of the currency
// @param startParameter Unique invoice bot start_parameter. To share an invoice use the URL https://t.me/{bot_username}?start={start_parameter}
// @param isTest True, if the invoice is a test invoice
// @param needShippingAddress True, if the shipping address should be specified
// @param receiptMessageID The identifier of the message with the receipt, after the product has been purchased
func NewMessageInvoice(title string, description string, photo *Photo, currency string, totalAmount int64, startParameter string, isTest bool, needShippingAddress bool, receiptMessageID int64) *MessageInvoice {
	messageInvoiceTemp := MessageInvoice{
		tdCommon:            tdCommon{Type: "messageInvoice"},
		Title:               title,
		Description:         description,
		Photo:               photo,
		Currency:            currency,
		TotalAmount:         totalAmount,
		StartParameter:      startParameter,
		IsTest:              isTest,
		NeedShippingAddress: needShippingAddress,
		ReceiptMessageID:    receiptMessageID,
	}

	return &messageInvoiceTemp
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

// NewMessageCall creates a new MessageCall
//
// @param discardReason Reason why the call was discarded
// @param duration Call duration, in seconds
func NewMessageCall(discardReason CallDiscardReason, duration int32) *MessageCall {
	messageCallTemp := MessageCall{
		tdCommon:      tdCommon{Type: "messageCall"},
		DiscardReason: discardReason,
		Duration:      duration,
	}

	return &messageCallTemp
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

// NewMessageBasicGroupChatCreate creates a new MessageBasicGroupChatCreate
//
// @param title Title of the basic group
// @param memberUserIDs User identifiers of members in the basic group
func NewMessageBasicGroupChatCreate(title string, memberUserIDs []int32) *MessageBasicGroupChatCreate {
	messageBasicGroupChatCreateTemp := MessageBasicGroupChatCreate{
		tdCommon:      tdCommon{Type: "messageBasicGroupChatCreate"},
		Title:         title,
		MemberUserIDs: memberUserIDs,
	}

	return &messageBasicGroupChatCreateTemp
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

// NewMessageSupergroupChatCreate creates a new MessageSupergroupChatCreate
//
// @param title Title of the supergroup or channel
func NewMessageSupergroupChatCreate(title string) *MessageSupergroupChatCreate {
	messageSupergroupChatCreateTemp := MessageSupergroupChatCreate{
		tdCommon: tdCommon{Type: "messageSupergroupChatCreate"},
		Title:    title,
	}

	return &messageSupergroupChatCreateTemp
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

// NewMessageChatChangeTitle creates a new MessageChatChangeTitle
//
// @param title New chat title
func NewMessageChatChangeTitle(title string) *MessageChatChangeTitle {
	messageChatChangeTitleTemp := MessageChatChangeTitle{
		tdCommon: tdCommon{Type: "messageChatChangeTitle"},
		Title:    title,
	}

	return &messageChatChangeTitleTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageChatChangeTitle *MessageChatChangeTitle) GetMessageContentEnum() MessageContentEnum {
	return MessageChatChangeTitleType
}

// MessageChatChangePhoto An updated chat photo
type MessageChatChangePhoto struct {
	tdCommon
	Photo *Photo `json:"photo"` // New chat photo
}

// MessageType return the string telegram-type of MessageChatChangePhoto
func (messageChatChangePhoto *MessageChatChangePhoto) MessageType() string {
	return "messageChatChangePhoto"
}

// NewMessageChatChangePhoto creates a new MessageChatChangePhoto
//
// @param photo New chat photo
func NewMessageChatChangePhoto(photo *Photo) *MessageChatChangePhoto {
	messageChatChangePhotoTemp := MessageChatChangePhoto{
		tdCommon: tdCommon{Type: "messageChatChangePhoto"},
		Photo:    photo,
	}

	return &messageChatChangePhotoTemp
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

// NewMessageChatDeletePhoto creates a new MessageChatDeletePhoto
//
func NewMessageChatDeletePhoto() *MessageChatDeletePhoto {
	messageChatDeletePhotoTemp := MessageChatDeletePhoto{
		tdCommon: tdCommon{Type: "messageChatDeletePhoto"},
	}

	return &messageChatDeletePhotoTemp
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

// NewMessageChatAddMembers creates a new MessageChatAddMembers
//
// @param memberUserIDs User identifiers of the new members
func NewMessageChatAddMembers(memberUserIDs []int32) *MessageChatAddMembers {
	messageChatAddMembersTemp := MessageChatAddMembers{
		tdCommon:      tdCommon{Type: "messageChatAddMembers"},
		MemberUserIDs: memberUserIDs,
	}

	return &messageChatAddMembersTemp
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

// NewMessageChatJoinByLink creates a new MessageChatJoinByLink
//
func NewMessageChatJoinByLink() *MessageChatJoinByLink {
	messageChatJoinByLinkTemp := MessageChatJoinByLink{
		tdCommon: tdCommon{Type: "messageChatJoinByLink"},
	}

	return &messageChatJoinByLinkTemp
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

// NewMessageChatDeleteMember creates a new MessageChatDeleteMember
//
// @param userID User identifier of the deleted chat member
func NewMessageChatDeleteMember(userID int32) *MessageChatDeleteMember {
	messageChatDeleteMemberTemp := MessageChatDeleteMember{
		tdCommon: tdCommon{Type: "messageChatDeleteMember"},
		UserID:   userID,
	}

	return &messageChatDeleteMemberTemp
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

// NewMessageChatUpgradeTo creates a new MessageChatUpgradeTo
//
// @param supergroupID Identifier of the supergroup to which the basic group was upgraded
func NewMessageChatUpgradeTo(supergroupID int32) *MessageChatUpgradeTo {
	messageChatUpgradeToTemp := MessageChatUpgradeTo{
		tdCommon:     tdCommon{Type: "messageChatUpgradeTo"},
		SupergroupID: supergroupID,
	}

	return &messageChatUpgradeToTemp
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

// NewMessageChatUpgradeFrom creates a new MessageChatUpgradeFrom
//
// @param title Title of the newly created supergroup
// @param basicGroupID The identifier of the original basic group
func NewMessageChatUpgradeFrom(title string, basicGroupID int32) *MessageChatUpgradeFrom {
	messageChatUpgradeFromTemp := MessageChatUpgradeFrom{
		tdCommon:     tdCommon{Type: "messageChatUpgradeFrom"},
		Title:        title,
		BasicGroupID: basicGroupID,
	}

	return &messageChatUpgradeFromTemp
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

// NewMessagePinMessage creates a new MessagePinMessage
//
// @param messageID Identifier of the pinned message, can be an identifier of a deleted message
func NewMessagePinMessage(messageID int64) *MessagePinMessage {
	messagePinMessageTemp := MessagePinMessage{
		tdCommon:  tdCommon{Type: "messagePinMessage"},
		MessageID: messageID,
	}

	return &messagePinMessageTemp
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

// NewMessageScreenshotTaken creates a new MessageScreenshotTaken
//
func NewMessageScreenshotTaken() *MessageScreenshotTaken {
	messageScreenshotTakenTemp := MessageScreenshotTaken{
		tdCommon: tdCommon{Type: "messageScreenshotTaken"},
	}

	return &messageScreenshotTakenTemp
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

// NewMessageChatSetTTL creates a new MessageChatSetTTL
//
// @param tTL New TTL
func NewMessageChatSetTTL(tTL int32) *MessageChatSetTTL {
	messageChatSetTTLTemp := MessageChatSetTTL{
		tdCommon: tdCommon{Type: "messageChatSetTtl"},
		TTL:      tTL,
	}

	return &messageChatSetTTLTemp
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

// NewMessageCustomServiceAction creates a new MessageCustomServiceAction
//
// @param text Message text to be shown in the chat
func NewMessageCustomServiceAction(text string) *MessageCustomServiceAction {
	messageCustomServiceActionTemp := MessageCustomServiceAction{
		tdCommon: tdCommon{Type: "messageCustomServiceAction"},
		Text:     text,
	}

	return &messageCustomServiceActionTemp
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

// NewMessageGameScore creates a new MessageGameScore
//
// @param gameMessageID Identifier of the message with the game, can be an identifier of a deleted message
// @param gameID Identifier of the game, may be different from the games presented in the message with the game
// @param score New score
func NewMessageGameScore(gameMessageID int64, gameID JSONInt64, score int32) *MessageGameScore {
	messageGameScoreTemp := MessageGameScore{
		tdCommon:      tdCommon{Type: "messageGameScore"},
		GameMessageID: gameMessageID,
		GameID:        gameID,
		Score:         score,
	}

	return &messageGameScoreTemp
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

// NewMessagePaymentSuccessful creates a new MessagePaymentSuccessful
//
// @param invoiceMessageID Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
// @param currency Currency for the price of the product
// @param totalAmount Total price for the product, in the minimal quantity of the currency
func NewMessagePaymentSuccessful(invoiceMessageID int64, currency string, totalAmount int64) *MessagePaymentSuccessful {
	messagePaymentSuccessfulTemp := MessagePaymentSuccessful{
		tdCommon:         tdCommon{Type: "messagePaymentSuccessful"},
		InvoiceMessageID: invoiceMessageID,
		Currency:         currency,
		TotalAmount:      totalAmount,
	}

	return &messagePaymentSuccessfulTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePaymentSuccessful *MessagePaymentSuccessful) GetMessageContentEnum() MessageContentEnum {
	return MessagePaymentSuccessfulType
}

// MessagePaymentSuccessfulBot A payment has been completed; for bots only
type MessagePaymentSuccessfulBot struct {
	tdCommon
	InvoiceMessageID        int64      `json:"invoice_message_id"`         // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
	Currency                string     `json:"currency"`                   // Currency for price of the product
	TotalAmount             int64      `json:"total_amount"`               // Total price for the product, in the minimal quantity of the currency
	InvoicePayload          []byte     `json:"invoice_payload"`            // Invoice payload
	ShippingOptionID        string     `json:"shipping_option_id"`         // Identifier of the shipping option chosen by the user; may be empty if not applicable
	OrderInfo               *OrderInfo `json:"order_info"`                 // Information about the order; may be null
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"` // Provider payment identifier
}

// MessageType return the string telegram-type of MessagePaymentSuccessfulBot
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) MessageType() string {
	return "messagePaymentSuccessfulBot"
}

// NewMessagePaymentSuccessfulBot creates a new MessagePaymentSuccessfulBot
//
// @param invoiceMessageID Identifier of the message with the corresponding invoice; can be an identifier of a deleted message
// @param currency Currency for price of the product
// @param totalAmount Total price for the product, in the minimal quantity of the currency
// @param invoicePayload Invoice payload
// @param shippingOptionID Identifier of the shipping option chosen by the user; may be empty if not applicable
// @param orderInfo Information about the order; may be null
// @param telegramPaymentChargeID Telegram payment identifier
// @param providerPaymentChargeID Provider payment identifier
func NewMessagePaymentSuccessfulBot(invoiceMessageID int64, currency string, totalAmount int64, invoicePayload []byte, shippingOptionID string, orderInfo *OrderInfo, telegramPaymentChargeID string, providerPaymentChargeID string) *MessagePaymentSuccessfulBot {
	messagePaymentSuccessfulBotTemp := MessagePaymentSuccessfulBot{
		tdCommon:                tdCommon{Type: "messagePaymentSuccessfulBot"},
		InvoiceMessageID:        invoiceMessageID,
		Currency:                currency,
		TotalAmount:             totalAmount,
		InvoicePayload:          invoicePayload,
		ShippingOptionID:        shippingOptionID,
		OrderInfo:               orderInfo,
		TelegramPaymentChargeID: telegramPaymentChargeID,
		ProviderPaymentChargeID: providerPaymentChargeID,
	}

	return &messagePaymentSuccessfulBotTemp
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

// NewMessageContactRegistered creates a new MessageContactRegistered
//
func NewMessageContactRegistered() *MessageContactRegistered {
	messageContactRegisteredTemp := MessageContactRegistered{
		tdCommon: tdCommon{Type: "messageContactRegistered"},
	}

	return &messageContactRegisteredTemp
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

// NewMessageWebsiteConnected creates a new MessageWebsiteConnected
//
// @param domainName Domain name of the connected website
func NewMessageWebsiteConnected(domainName string) *MessageWebsiteConnected {
	messageWebsiteConnectedTemp := MessageWebsiteConnected{
		tdCommon:   tdCommon{Type: "messageWebsiteConnected"},
		DomainName: domainName,
	}

	return &messageWebsiteConnectedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messageWebsiteConnected *MessageWebsiteConnected) GetMessageContentEnum() MessageContentEnum {
	return MessageWebsiteConnectedType
}

// MessagePassportDataSent Telegram Passport data has been sent
type MessagePassportDataSent struct {
	tdCommon
	Types []PassportElementType `json:"types"` // List of Telegram Passport element types sent
}

// MessageType return the string telegram-type of MessagePassportDataSent
func (messagePassportDataSent *MessagePassportDataSent) MessageType() string {
	return "messagePassportDataSent"
}

// NewMessagePassportDataSent creates a new MessagePassportDataSent
//
// @param typeParams List of Telegram Passport element types sent
func NewMessagePassportDataSent(typeParams []PassportElementType) *MessagePassportDataSent {
	messagePassportDataSentTemp := MessagePassportDataSent{
		tdCommon: tdCommon{Type: "messagePassportDataSent"},
		Types:    typeParams,
	}

	return &messagePassportDataSentTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePassportDataSent *MessagePassportDataSent) GetMessageContentEnum() MessageContentEnum {
	return MessagePassportDataSentType
}

// MessagePassportDataReceived Telegram Passport data has been received; for bots only
type MessagePassportDataReceived struct {
	tdCommon
	Elements    []EncryptedPassportElement `json:"elements"`    // List of received Telegram Passport elements
	Credentials *EncryptedCredentials      `json:"credentials"` // Encrypted data credentials
}

// MessageType return the string telegram-type of MessagePassportDataReceived
func (messagePassportDataReceived *MessagePassportDataReceived) MessageType() string {
	return "messagePassportDataReceived"
}

// NewMessagePassportDataReceived creates a new MessagePassportDataReceived
//
// @param elements List of received Telegram Passport elements
// @param credentials Encrypted data credentials
func NewMessagePassportDataReceived(elements []EncryptedPassportElement, credentials *EncryptedCredentials) *MessagePassportDataReceived {
	messagePassportDataReceivedTemp := MessagePassportDataReceived{
		tdCommon:    tdCommon{Type: "messagePassportDataReceived"},
		Elements:    elements,
		Credentials: credentials,
	}

	return &messagePassportDataReceivedTemp
}

// GetMessageContentEnum return the enum type of this object
func (messagePassportDataReceived *MessagePassportDataReceived) GetMessageContentEnum() MessageContentEnum {
	return MessagePassportDataReceivedType
}

// MessageUnsupported Message content that is not supported by the client
type MessageUnsupported struct {
	tdCommon
}

// MessageType return the string telegram-type of MessageUnsupported
func (messageUnsupported *MessageUnsupported) MessageType() string {
	return "messageUnsupported"
}

// NewMessageUnsupported creates a new MessageUnsupported
//
func NewMessageUnsupported() *MessageUnsupported {
	messageUnsupportedTemp := MessageUnsupported{
		tdCommon: tdCommon{Type: "messageUnsupported"},
	}

	return &messageUnsupportedTemp
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

// NewTextEntityTypeMention creates a new TextEntityTypeMention
//
func NewTextEntityTypeMention() *TextEntityTypeMention {
	textEntityTypeMentionTemp := TextEntityTypeMention{
		tdCommon: tdCommon{Type: "textEntityTypeMention"},
	}

	return &textEntityTypeMentionTemp
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

// NewTextEntityTypeHashtag creates a new TextEntityTypeHashtag
//
func NewTextEntityTypeHashtag() *TextEntityTypeHashtag {
	textEntityTypeHashtagTemp := TextEntityTypeHashtag{
		tdCommon: tdCommon{Type: "textEntityTypeHashtag"},
	}

	return &textEntityTypeHashtagTemp
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

// NewTextEntityTypeCashtag creates a new TextEntityTypeCashtag
//
func NewTextEntityTypeCashtag() *TextEntityTypeCashtag {
	textEntityTypeCashtagTemp := TextEntityTypeCashtag{
		tdCommon: tdCommon{Type: "textEntityTypeCashtag"},
	}

	return &textEntityTypeCashtagTemp
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

// NewTextEntityTypeBotCommand creates a new TextEntityTypeBotCommand
//
func NewTextEntityTypeBotCommand() *TextEntityTypeBotCommand {
	textEntityTypeBotCommandTemp := TextEntityTypeBotCommand{
		tdCommon: tdCommon{Type: "textEntityTypeBotCommand"},
	}

	return &textEntityTypeBotCommandTemp
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

// NewTextEntityTypeURL creates a new TextEntityTypeURL
//
func NewTextEntityTypeURL() *TextEntityTypeURL {
	textEntityTypeURLTemp := TextEntityTypeURL{
		tdCommon: tdCommon{Type: "textEntityTypeUrl"},
	}

	return &textEntityTypeURLTemp
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

// NewTextEntityTypeEmailAddress creates a new TextEntityTypeEmailAddress
//
func NewTextEntityTypeEmailAddress() *TextEntityTypeEmailAddress {
	textEntityTypeEmailAddressTemp := TextEntityTypeEmailAddress{
		tdCommon: tdCommon{Type: "textEntityTypeEmailAddress"},
	}

	return &textEntityTypeEmailAddressTemp
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

// NewTextEntityTypeBold creates a new TextEntityTypeBold
//
func NewTextEntityTypeBold() *TextEntityTypeBold {
	textEntityTypeBoldTemp := TextEntityTypeBold{
		tdCommon: tdCommon{Type: "textEntityTypeBold"},
	}

	return &textEntityTypeBoldTemp
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

// NewTextEntityTypeItalic creates a new TextEntityTypeItalic
//
func NewTextEntityTypeItalic() *TextEntityTypeItalic {
	textEntityTypeItalicTemp := TextEntityTypeItalic{
		tdCommon: tdCommon{Type: "textEntityTypeItalic"},
	}

	return &textEntityTypeItalicTemp
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

// NewTextEntityTypeCode creates a new TextEntityTypeCode
//
func NewTextEntityTypeCode() *TextEntityTypeCode {
	textEntityTypeCodeTemp := TextEntityTypeCode{
		tdCommon: tdCommon{Type: "textEntityTypeCode"},
	}

	return &textEntityTypeCodeTemp
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

// NewTextEntityTypePre creates a new TextEntityTypePre
//
func NewTextEntityTypePre() *TextEntityTypePre {
	textEntityTypePreTemp := TextEntityTypePre{
		tdCommon: tdCommon{Type: "textEntityTypePre"},
	}

	return &textEntityTypePreTemp
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

// NewTextEntityTypePreCode creates a new TextEntityTypePreCode
//
// @param language Programming language of the code; as defined by the sender
func NewTextEntityTypePreCode(language string) *TextEntityTypePreCode {
	textEntityTypePreCodeTemp := TextEntityTypePreCode{
		tdCommon: tdCommon{Type: "textEntityTypePreCode"},
		Language: language,
	}

	return &textEntityTypePreCodeTemp
}

// GetTextEntityTypeEnum return the enum type of this object
func (textEntityTypePreCode *TextEntityTypePreCode) GetTextEntityTypeEnum() TextEntityTypeEnum {
	return TextEntityTypePreCodeType
}

// TextEntityTypeTextURL A text description shown instead of a raw URL
type TextEntityTypeTextURL struct {
	tdCommon
	URL string `json:"url"` // HTTP or tg:// URL to be opened when the link is clicked
}

// MessageType return the string telegram-type of TextEntityTypeTextURL
func (textEntityTypeTextURL *TextEntityTypeTextURL) MessageType() string {
	return "textEntityTypeTextUrl"
}

// NewTextEntityTypeTextURL creates a new TextEntityTypeTextURL
//
// @param uRL HTTP or tg:// URL to be opened when the link is clicked
func NewTextEntityTypeTextURL(uRL string) *TextEntityTypeTextURL {
	textEntityTypeTextURLTemp := TextEntityTypeTextURL{
		tdCommon: tdCommon{Type: "textEntityTypeTextUrl"},
		URL:      uRL,
	}

	return &textEntityTypeTextURLTemp
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

// NewTextEntityTypeMentionName creates a new TextEntityTypeMentionName
//
// @param userID Identifier of the mentioned user
func NewTextEntityTypeMentionName(userID int32) *TextEntityTypeMentionName {
	textEntityTypeMentionNameTemp := TextEntityTypeMentionName{
		tdCommon: tdCommon{Type: "textEntityTypeMentionName"},
		UserID:   userID,
	}

	return &textEntityTypeMentionNameTemp
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

// NewTextEntityTypePhoneNumber creates a new TextEntityTypePhoneNumber
//
func NewTextEntityTypePhoneNumber() *TextEntityTypePhoneNumber {
	textEntityTypePhoneNumberTemp := TextEntityTypePhoneNumber{
		tdCommon: tdCommon{Type: "textEntityTypePhoneNumber"},
	}

	return &textEntityTypePhoneNumberTemp
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

// NewInputThumbnail creates a new InputThumbnail
//
// @param thumbnail Thumbnail file to send. Sending thumbnails by file_id is currently not supported
// @param width Thumbnail width, usually shouldn't exceed 90. Use 0 if unknown
// @param height Thumbnail height, usually shouldn't exceed 90. Use 0 if unknown
func NewInputThumbnail(thumbnail InputFile, width int32, height int32) *InputThumbnail {
	inputThumbnailTemp := InputThumbnail{
		tdCommon:  tdCommon{Type: "inputThumbnail"},
		Thumbnail: thumbnail,
		Width:     width,
		Height:    height,
	}

	return &inputThumbnailTemp
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
	Text                  *FormattedText `json:"text"`                     // Formatted text to be sent; 1-GetOption("message_text_length_max") characters. Only Bold, Italic, Code, Pre, PreCode and TextUrl entities are allowed to be specified manually
	DisableWebPagePreview bool           `json:"disable_web_page_preview"` // True, if rich web page previews for URLs in the message text should be disabled
	ClearDraft            bool           `json:"clear_draft"`              // True, if a chat message draft should be deleted
}

// MessageType return the string telegram-type of InputMessageText
func (inputMessageText *InputMessageText) MessageType() string {
	return "inputMessageText"
}

// NewInputMessageText creates a new InputMessageText
//
// @param text Formatted text to be sent; 1-GetOption("message_text_length_max") characters. Only Bold, Italic, Code, Pre, PreCode and TextUrl entities are allowed to be specified manually
// @param disableWebPagePreview True, if rich web page previews for URLs in the message text should be disabled
// @param clearDraft True, if a chat message draft should be deleted
func NewInputMessageText(text *FormattedText, disableWebPagePreview bool, clearDraft bool) *InputMessageText {
	inputMessageTextTemp := InputMessageText{
		tdCommon:              tdCommon{Type: "inputMessageText"},
		Text:                  text,
		DisableWebPagePreview: disableWebPagePreview,
		ClearDraft:            clearDraft,
	}

	return &inputMessageTextTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageText *InputMessageText) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageTextType
}

// InputMessageAnimation An animation message (GIF-style).
type InputMessageAnimation struct {
	tdCommon
	Animation InputFile       `json:"animation"` // Animation file to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"` // Animation thumbnail, if available
	Duration  int32           `json:"duration"`  // Duration of the animation, in seconds
	Width     int32           `json:"width"`     // Width of the animation; may be replaced by the server
	Height    int32           `json:"height"`    // Height of the animation; may be replaced by the server
	Caption   *FormattedText  `json:"caption"`   // Animation caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageAnimation
func (inputMessageAnimation *InputMessageAnimation) MessageType() string {
	return "inputMessageAnimation"
}

// NewInputMessageAnimation creates a new InputMessageAnimation
//
// @param animation Animation file to be sent
// @param thumbnail Animation thumbnail, if available
// @param duration Duration of the animation, in seconds
// @param width Width of the animation; may be replaced by the server
// @param height Height of the animation; may be replaced by the server
// @param caption Animation caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageAnimation(animation InputFile, thumbnail *InputThumbnail, duration int32, width int32, height int32, caption *FormattedText) *InputMessageAnimation {
	inputMessageAnimationTemp := InputMessageAnimation{
		tdCommon:  tdCommon{Type: "inputMessageAnimation"},
		Animation: animation,
		Thumbnail: thumbnail,
		Duration:  duration,
		Width:     width,
		Height:    height,
		Caption:   caption,
	}

	return &inputMessageAnimationTemp
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
		Thumbnail *InputThumbnail `json:"thumbnail"` // Animation thumbnail, if available
		Duration  int32           `json:"duration"`  // Duration of the animation, in seconds
		Width     int32           `json:"width"`     // Width of the animation; may be replaced by the server
		Height    int32           `json:"height"`    // Height of the animation; may be replaced by the server
		Caption   *FormattedText  `json:"caption"`   // Animation caption; 0-GetOption("message_caption_length_max") characters
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
	Audio               InputFile       `json:"audio"`                 // Audio file to be sent
	AlbumCoverThumbnail *InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album, if available
	Duration            int32           `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
	Title               string          `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
	Performer           string          `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
	Caption             *FormattedText  `json:"caption"`               // Audio caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageAudio
func (inputMessageAudio *InputMessageAudio) MessageType() string {
	return "inputMessageAudio"
}

// NewInputMessageAudio creates a new InputMessageAudio
//
// @param audio Audio file to be sent
// @param albumCoverThumbnail Thumbnail of the cover for the album, if available
// @param duration Duration of the audio, in seconds; may be replaced by the server
// @param title Title of the audio; 0-64 characters; may be replaced by the server
// @param performer Performer of the audio; 0-64 characters, may be replaced by the server
// @param caption Audio caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageAudio(audio InputFile, albumCoverThumbnail *InputThumbnail, duration int32, title string, performer string, caption *FormattedText) *InputMessageAudio {
	inputMessageAudioTemp := InputMessageAudio{
		tdCommon:            tdCommon{Type: "inputMessageAudio"},
		Audio:               audio,
		AlbumCoverThumbnail: albumCoverThumbnail,
		Duration:            duration,
		Title:               title,
		Performer:           performer,
		Caption:             caption,
	}

	return &inputMessageAudioTemp
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
		AlbumCoverThumbnail *InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album, if available
		Duration            int32           `json:"duration"`              // Duration of the audio, in seconds; may be replaced by the server
		Title               string          `json:"title"`                 // Title of the audio; 0-64 characters; may be replaced by the server
		Performer           string          `json:"performer"`             // Performer of the audio; 0-64 characters, may be replaced by the server
		Caption             *FormattedText  `json:"caption"`               // Audio caption; 0-GetOption("message_caption_length_max") characters
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
	Document  InputFile       `json:"document"`  // Document to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"` // Document thumbnail, if available
	Caption   *FormattedText  `json:"caption"`   // Document caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageDocument
func (inputMessageDocument *InputMessageDocument) MessageType() string {
	return "inputMessageDocument"
}

// NewInputMessageDocument creates a new InputMessageDocument
//
// @param document Document to be sent
// @param thumbnail Document thumbnail, if available
// @param caption Document caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageDocument(document InputFile, thumbnail *InputThumbnail, caption *FormattedText) *InputMessageDocument {
	inputMessageDocumentTemp := InputMessageDocument{
		tdCommon:  tdCommon{Type: "inputMessageDocument"},
		Document:  document,
		Thumbnail: thumbnail,
		Caption:   caption,
	}

	return &inputMessageDocumentTemp
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
		Thumbnail *InputThumbnail `json:"thumbnail"` // Document thumbnail, if available
		Caption   *FormattedText  `json:"caption"`   // Document caption; 0-GetOption("message_caption_length_max") characters
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
	Photo               InputFile       `json:"photo"`                  // Photo to send
	Thumbnail           *InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent, this is sent to the other party in secret chats only
	AddedStickerFileIDs []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
	Width               int32           `json:"width"`                  // Photo width
	Height              int32           `json:"height"`                 // Photo height
	Caption             *FormattedText  `json:"caption"`                // Photo caption; 0-GetOption("message_caption_length_max") characters
	TTL                 int32           `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessagePhoto
func (inputMessagePhoto *InputMessagePhoto) MessageType() string {
	return "inputMessagePhoto"
}

// NewInputMessagePhoto creates a new InputMessagePhoto
//
// @param photo Photo to send
// @param thumbnail Photo thumbnail to be sent, this is sent to the other party in secret chats only
// @param addedStickerFileIDs File identifiers of the stickers added to the photo, if applicable
// @param width Photo width
// @param height Photo height
// @param caption Photo caption; 0-GetOption("message_caption_length_max") characters
// @param tTL Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
func NewInputMessagePhoto(photo InputFile, thumbnail *InputThumbnail, addedStickerFileIDs []int32, width int32, height int32, caption *FormattedText, tTL int32) *InputMessagePhoto {
	inputMessagePhotoTemp := InputMessagePhoto{
		tdCommon:            tdCommon{Type: "inputMessagePhoto"},
		Photo:               photo,
		Thumbnail:           thumbnail,
		AddedStickerFileIDs: addedStickerFileIDs,
		Width:               width,
		Height:              height,
		Caption:             caption,
		TTL:                 tTL,
	}

	return &inputMessagePhotoTemp
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
		Thumbnail           *InputThumbnail `json:"thumbnail"`              // Photo thumbnail to be sent, this is sent to the other party in secret chats only
		AddedStickerFileIDs []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable
		Width               int32           `json:"width"`                  // Photo width
		Height              int32           `json:"height"`                 // Photo height
		Caption             *FormattedText  `json:"caption"`                // Photo caption; 0-GetOption("message_caption_length_max") characters
		TTL                 int32           `json:"ttl"`                    // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
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
	Sticker   InputFile       `json:"sticker"`   // Sticker to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"` // Sticker thumbnail, if available
	Width     int32           `json:"width"`     // Sticker width
	Height    int32           `json:"height"`    // Sticker height
}

// MessageType return the string telegram-type of InputMessageSticker
func (inputMessageSticker *InputMessageSticker) MessageType() string {
	return "inputMessageSticker"
}

// NewInputMessageSticker creates a new InputMessageSticker
//
// @param sticker Sticker to be sent
// @param thumbnail Sticker thumbnail, if available
// @param width Sticker width
// @param height Sticker height
func NewInputMessageSticker(sticker InputFile, thumbnail *InputThumbnail, width int32, height int32) *InputMessageSticker {
	inputMessageStickerTemp := InputMessageSticker{
		tdCommon:  tdCommon{Type: "inputMessageSticker"},
		Sticker:   sticker,
		Thumbnail: thumbnail,
		Width:     width,
		Height:    height,
	}

	return &inputMessageStickerTemp
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
		Thumbnail *InputThumbnail `json:"thumbnail"` // Sticker thumbnail, if available
		Width     int32           `json:"width"`     // Sticker width
		Height    int32           `json:"height"`    // Sticker height
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
	Video               InputFile       `json:"video"`                  // Video to be sent
	Thumbnail           *InputThumbnail `json:"thumbnail"`              // Video thumbnail, if available
	AddedStickerFileIDs []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
	Duration            int32           `json:"duration"`               // Duration of the video, in seconds
	Width               int32           `json:"width"`                  // Video width
	Height              int32           `json:"height"`                 // Video height
	SupportsStreaming   bool            `json:"supports_streaming"`     // True, if the video should be tried to be streamed
	Caption             *FormattedText  `json:"caption"`                // Video caption; 0-GetOption("message_caption_length_max") characters
	TTL                 int32           `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessageVideo
func (inputMessageVideo *InputMessageVideo) MessageType() string {
	return "inputMessageVideo"
}

// NewInputMessageVideo creates a new InputMessageVideo
//
// @param video Video to be sent
// @param thumbnail Video thumbnail, if available
// @param addedStickerFileIDs File identifiers of the stickers added to the video, if applicable
// @param duration Duration of the video, in seconds
// @param width Video width
// @param height Video height
// @param supportsStreaming True, if the video should be tried to be streamed
// @param caption Video caption; 0-GetOption("message_caption_length_max") characters
// @param tTL Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
func NewInputMessageVideo(video InputFile, thumbnail *InputThumbnail, addedStickerFileIDs []int32, duration int32, width int32, height int32, supportsStreaming bool, caption *FormattedText, tTL int32) *InputMessageVideo {
	inputMessageVideoTemp := InputMessageVideo{
		tdCommon:            tdCommon{Type: "inputMessageVideo"},
		Video:               video,
		Thumbnail:           thumbnail,
		AddedStickerFileIDs: addedStickerFileIDs,
		Duration:            duration,
		Width:               width,
		Height:              height,
		SupportsStreaming:   supportsStreaming,
		Caption:             caption,
		TTL:                 tTL,
	}

	return &inputMessageVideoTemp
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
		Thumbnail           *InputThumbnail `json:"thumbnail"`              // Video thumbnail, if available
		AddedStickerFileIDs []int32         `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
		Duration            int32           `json:"duration"`               // Duration of the video, in seconds
		Width               int32           `json:"width"`                  // Video width
		Height              int32           `json:"height"`                 // Video height
		SupportsStreaming   bool            `json:"supports_streaming"`     // True, if the video should be tried to be streamed
		Caption             *FormattedText  `json:"caption"`                // Video caption; 0-GetOption("message_caption_length_max") characters
		TTL                 int32           `json:"ttl"`                    // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
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
	VideoNote InputFile       `json:"video_note"` // Video note to be sent
	Thumbnail *InputThumbnail `json:"thumbnail"`  // Video thumbnail, if available
	Duration  int32           `json:"duration"`   // Duration of the video, in seconds
	Length    int32           `json:"length"`     // Video width and height; must be positive and not greater than 640
}

// MessageType return the string telegram-type of InputMessageVideoNote
func (inputMessageVideoNote *InputMessageVideoNote) MessageType() string {
	return "inputMessageVideoNote"
}

// NewInputMessageVideoNote creates a new InputMessageVideoNote
//
// @param videoNote Video note to be sent
// @param thumbnail Video thumbnail, if available
// @param duration Duration of the video, in seconds
// @param length Video width and height; must be positive and not greater than 640
func NewInputMessageVideoNote(videoNote InputFile, thumbnail *InputThumbnail, duration int32, length int32) *InputMessageVideoNote {
	inputMessageVideoNoteTemp := InputMessageVideoNote{
		tdCommon:  tdCommon{Type: "inputMessageVideoNote"},
		VideoNote: videoNote,
		Thumbnail: thumbnail,
		Duration:  duration,
		Length:    length,
	}

	return &inputMessageVideoNoteTemp
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
		Thumbnail *InputThumbnail `json:"thumbnail"` // Video thumbnail, if available
		Duration  int32           `json:"duration"`  // Duration of the video, in seconds
		Length    int32           `json:"length"`    // Video width and height; must be positive and not greater than 640
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
	VoiceNote InputFile      `json:"voice_note"` // Voice note to be sent
	Duration  int32          `json:"duration"`   // Duration of the voice note, in seconds
	Waveform  []byte         `json:"waveform"`   // Waveform representation of the voice note, in 5-bit format
	Caption   *FormattedText `json:"caption"`    // Voice note caption; 0-GetOption("message_caption_length_max") characters
}

// MessageType return the string telegram-type of InputMessageVoiceNote
func (inputMessageVoiceNote *InputMessageVoiceNote) MessageType() string {
	return "inputMessageVoiceNote"
}

// NewInputMessageVoiceNote creates a new InputMessageVoiceNote
//
// @param voiceNote Voice note to be sent
// @param duration Duration of the voice note, in seconds
// @param waveform Waveform representation of the voice note, in 5-bit format
// @param caption Voice note caption; 0-GetOption("message_caption_length_max") characters
func NewInputMessageVoiceNote(voiceNote InputFile, duration int32, waveform []byte, caption *FormattedText) *InputMessageVoiceNote {
	inputMessageVoiceNoteTemp := InputMessageVoiceNote{
		tdCommon:  tdCommon{Type: "inputMessageVoiceNote"},
		VoiceNote: voiceNote,
		Duration:  duration,
		Waveform:  waveform,
		Caption:   caption,
	}

	return &inputMessageVoiceNoteTemp
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
		Duration int32          `json:"duration"` // Duration of the voice note, in seconds
		Waveform []byte         `json:"waveform"` // Waveform representation of the voice note, in 5-bit format
		Caption  *FormattedText `json:"caption"`  // Voice note caption; 0-GetOption("message_caption_length_max") characters
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
	Location   *Location `json:"location"`    // Location to be sent
	LivePeriod int32     `json:"live_period"` // Period for which the location can be updated, in seconds; should bebetween 60 and 86400 for a live location and 0 otherwise
}

// MessageType return the string telegram-type of InputMessageLocation
func (inputMessageLocation *InputMessageLocation) MessageType() string {
	return "inputMessageLocation"
}

// NewInputMessageLocation creates a new InputMessageLocation
//
// @param location Location to be sent
// @param livePeriod Period for which the location can be updated, in seconds; should bebetween 60 and 86400 for a live location and 0 otherwise
func NewInputMessageLocation(location *Location, livePeriod int32) *InputMessageLocation {
	inputMessageLocationTemp := InputMessageLocation{
		tdCommon:   tdCommon{Type: "inputMessageLocation"},
		Location:   location,
		LivePeriod: livePeriod,
	}

	return &inputMessageLocationTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageLocation *InputMessageLocation) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageLocationType
}

// InputMessageVenue A message with information about a venue
type InputMessageVenue struct {
	tdCommon
	Venue *Venue `json:"venue"` // Venue to send
}

// MessageType return the string telegram-type of InputMessageVenue
func (inputMessageVenue *InputMessageVenue) MessageType() string {
	return "inputMessageVenue"
}

// NewInputMessageVenue creates a new InputMessageVenue
//
// @param venue Venue to send
func NewInputMessageVenue(venue *Venue) *InputMessageVenue {
	inputMessageVenueTemp := InputMessageVenue{
		tdCommon: tdCommon{Type: "inputMessageVenue"},
		Venue:    venue,
	}

	return &inputMessageVenueTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageVenue *InputMessageVenue) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageVenueType
}

// InputMessageContact A message containing a user contact
type InputMessageContact struct {
	tdCommon
	Contact *Contact `json:"contact"` // Contact to send
}

// MessageType return the string telegram-type of InputMessageContact
func (inputMessageContact *InputMessageContact) MessageType() string {
	return "inputMessageContact"
}

// NewInputMessageContact creates a new InputMessageContact
//
// @param contact Contact to send
func NewInputMessageContact(contact *Contact) *InputMessageContact {
	inputMessageContactTemp := InputMessageContact{
		tdCommon: tdCommon{Type: "inputMessageContact"},
		Contact:  contact,
	}

	return &inputMessageContactTemp
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

// NewInputMessageGame creates a new InputMessageGame
//
// @param botUserID User identifier of the bot that owns the game
// @param gameShortName Short name of the game
func NewInputMessageGame(botUserID int32, gameShortName string) *InputMessageGame {
	inputMessageGameTemp := InputMessageGame{
		tdCommon:      tdCommon{Type: "inputMessageGame"},
		BotUserID:     botUserID,
		GameShortName: gameShortName,
	}

	return &inputMessageGameTemp
}

// GetInputMessageContentEnum return the enum type of this object
func (inputMessageGame *InputMessageGame) GetInputMessageContentEnum() InputMessageContentEnum {
	return InputMessageGameType
}

// InputMessageInvoice A message with an invoice; can be used only by bots and only in private chats
type InputMessageInvoice struct {
	tdCommon
	Invoice        *Invoice `json:"invoice"`         // Invoice
	Title          string   `json:"title"`           // Product title; 1-32 characters
	Description    string   `json:"description"`     //
	PhotoURL       string   `json:"photo_url"`       // Product photo URL; optional
	PhotoSize      int32    `json:"photo_size"`      // Product photo size
	PhotoWidth     int32    `json:"photo_width"`     // Product photo width
	PhotoHeight    int32    `json:"photo_height"`    // Product photo height
	Payload        []byte   `json:"payload"`         // The invoice payload
	ProviderToken  string   `json:"provider_token"`  // Payment provider token
	ProviderData   string   `json:"provider_data"`   // JSON-encoded data about the invoice, which will be shared with the payment provider
	StartParameter string   `json:"start_parameter"` // Unique invoice bot start_parameter for the generation of this invoice
}

// MessageType return the string telegram-type of InputMessageInvoice
func (inputMessageInvoice *InputMessageInvoice) MessageType() string {
	return "inputMessageInvoice"
}

// NewInputMessageInvoice creates a new InputMessageInvoice
//
// @param invoice Invoice
// @param title Product title; 1-32 characters
// @param description
// @param photoURL Product photo URL; optional
// @param photoSize Product photo size
// @param photoWidth Product photo width
// @param photoHeight Product photo height
// @param payload The invoice payload
// @param providerToken Payment provider token
// @param providerData JSON-encoded data about the invoice, which will be shared with the payment provider
// @param startParameter Unique invoice bot start_parameter for the generation of this invoice
func NewInputMessageInvoice(invoice *Invoice, title string, description string, photoURL string, photoSize int32, photoWidth int32, photoHeight int32, payload []byte, providerToken string, providerData string, startParameter string) *InputMessageInvoice {
	inputMessageInvoiceTemp := InputMessageInvoice{
		tdCommon:       tdCommon{Type: "inputMessageInvoice"},
		Invoice:        invoice,
		Title:          title,
		Description:    description,
		PhotoURL:       photoURL,
		PhotoSize:      photoSize,
		PhotoWidth:     photoWidth,
		PhotoHeight:    photoHeight,
		Payload:        payload,
		ProviderToken:  providerToken,
		ProviderData:   providerData,
		StartParameter: startParameter,
	}

	return &inputMessageInvoiceTemp
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

// NewInputMessageForwarded creates a new InputMessageForwarded
//
// @param fromChatID Identifier for the chat this forwarded message came from
// @param messageID Identifier of the message to forward
// @param inGameShare True, if a game message should be shared within a launched game; applies only to game messages
func NewInputMessageForwarded(fromChatID int64, messageID int64, inGameShare bool) *InputMessageForwarded {
	inputMessageForwardedTemp := InputMessageForwarded{
		tdCommon:    tdCommon{Type: "inputMessageForwarded"},
		FromChatID:  fromChatID,
		MessageID:   messageID,
		InGameShare: inGameShare,
	}

	return &inputMessageForwardedTemp
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

// NewSearchMessagesFilterEmpty creates a new SearchMessagesFilterEmpty
//
func NewSearchMessagesFilterEmpty() *SearchMessagesFilterEmpty {
	searchMessagesFilterEmptyTemp := SearchMessagesFilterEmpty{
		tdCommon: tdCommon{Type: "searchMessagesFilterEmpty"},
	}

	return &searchMessagesFilterEmptyTemp
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

// NewSearchMessagesFilterAnimation creates a new SearchMessagesFilterAnimation
//
func NewSearchMessagesFilterAnimation() *SearchMessagesFilterAnimation {
	searchMessagesFilterAnimationTemp := SearchMessagesFilterAnimation{
		tdCommon: tdCommon{Type: "searchMessagesFilterAnimation"},
	}

	return &searchMessagesFilterAnimationTemp
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

// NewSearchMessagesFilterAudio creates a new SearchMessagesFilterAudio
//
func NewSearchMessagesFilterAudio() *SearchMessagesFilterAudio {
	searchMessagesFilterAudioTemp := SearchMessagesFilterAudio{
		tdCommon: tdCommon{Type: "searchMessagesFilterAudio"},
	}

	return &searchMessagesFilterAudioTemp
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

// NewSearchMessagesFilterDocument creates a new SearchMessagesFilterDocument
//
func NewSearchMessagesFilterDocument() *SearchMessagesFilterDocument {
	searchMessagesFilterDocumentTemp := SearchMessagesFilterDocument{
		tdCommon: tdCommon{Type: "searchMessagesFilterDocument"},
	}

	return &searchMessagesFilterDocumentTemp
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

// NewSearchMessagesFilterPhoto creates a new SearchMessagesFilterPhoto
//
func NewSearchMessagesFilterPhoto() *SearchMessagesFilterPhoto {
	searchMessagesFilterPhotoTemp := SearchMessagesFilterPhoto{
		tdCommon: tdCommon{Type: "searchMessagesFilterPhoto"},
	}

	return &searchMessagesFilterPhotoTemp
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

// NewSearchMessagesFilterVideo creates a new SearchMessagesFilterVideo
//
func NewSearchMessagesFilterVideo() *SearchMessagesFilterVideo {
	searchMessagesFilterVideoTemp := SearchMessagesFilterVideo{
		tdCommon: tdCommon{Type: "searchMessagesFilterVideo"},
	}

	return &searchMessagesFilterVideoTemp
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

// NewSearchMessagesFilterVoiceNote creates a new SearchMessagesFilterVoiceNote
//
func NewSearchMessagesFilterVoiceNote() *SearchMessagesFilterVoiceNote {
	searchMessagesFilterVoiceNoteTemp := SearchMessagesFilterVoiceNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVoiceNote"},
	}

	return &searchMessagesFilterVoiceNoteTemp
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

// NewSearchMessagesFilterPhotoAndVideo creates a new SearchMessagesFilterPhotoAndVideo
//
func NewSearchMessagesFilterPhotoAndVideo() *SearchMessagesFilterPhotoAndVideo {
	searchMessagesFilterPhotoAndVideoTemp := SearchMessagesFilterPhotoAndVideo{
		tdCommon: tdCommon{Type: "searchMessagesFilterPhotoAndVideo"},
	}

	return &searchMessagesFilterPhotoAndVideoTemp
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

// NewSearchMessagesFilterURL creates a new SearchMessagesFilterURL
//
func NewSearchMessagesFilterURL() *SearchMessagesFilterURL {
	searchMessagesFilterURLTemp := SearchMessagesFilterURL{
		tdCommon: tdCommon{Type: "searchMessagesFilterUrl"},
	}

	return &searchMessagesFilterURLTemp
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

// NewSearchMessagesFilterChatPhoto creates a new SearchMessagesFilterChatPhoto
//
func NewSearchMessagesFilterChatPhoto() *SearchMessagesFilterChatPhoto {
	searchMessagesFilterChatPhotoTemp := SearchMessagesFilterChatPhoto{
		tdCommon: tdCommon{Type: "searchMessagesFilterChatPhoto"},
	}

	return &searchMessagesFilterChatPhotoTemp
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

// NewSearchMessagesFilterCall creates a new SearchMessagesFilterCall
//
func NewSearchMessagesFilterCall() *SearchMessagesFilterCall {
	searchMessagesFilterCallTemp := SearchMessagesFilterCall{
		tdCommon: tdCommon{Type: "searchMessagesFilterCall"},
	}

	return &searchMessagesFilterCallTemp
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

// NewSearchMessagesFilterMissedCall creates a new SearchMessagesFilterMissedCall
//
func NewSearchMessagesFilterMissedCall() *SearchMessagesFilterMissedCall {
	searchMessagesFilterMissedCallTemp := SearchMessagesFilterMissedCall{
		tdCommon: tdCommon{Type: "searchMessagesFilterMissedCall"},
	}

	return &searchMessagesFilterMissedCallTemp
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

// NewSearchMessagesFilterVideoNote creates a new SearchMessagesFilterVideoNote
//
func NewSearchMessagesFilterVideoNote() *SearchMessagesFilterVideoNote {
	searchMessagesFilterVideoNoteTemp := SearchMessagesFilterVideoNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVideoNote"},
	}

	return &searchMessagesFilterVideoNoteTemp
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

// NewSearchMessagesFilterVoiceAndVideoNote creates a new SearchMessagesFilterVoiceAndVideoNote
//
func NewSearchMessagesFilterVoiceAndVideoNote() *SearchMessagesFilterVoiceAndVideoNote {
	searchMessagesFilterVoiceAndVideoNoteTemp := SearchMessagesFilterVoiceAndVideoNote{
		tdCommon: tdCommon{Type: "searchMessagesFilterVoiceAndVideoNote"},
	}

	return &searchMessagesFilterVoiceAndVideoNoteTemp
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

// NewSearchMessagesFilterMention creates a new SearchMessagesFilterMention
//
func NewSearchMessagesFilterMention() *SearchMessagesFilterMention {
	searchMessagesFilterMentionTemp := SearchMessagesFilterMention{
		tdCommon: tdCommon{Type: "searchMessagesFilterMention"},
	}

	return &searchMessagesFilterMentionTemp
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

// NewSearchMessagesFilterUnreadMention creates a new SearchMessagesFilterUnreadMention
//
func NewSearchMessagesFilterUnreadMention() *SearchMessagesFilterUnreadMention {
	searchMessagesFilterUnreadMentionTemp := SearchMessagesFilterUnreadMention{
		tdCommon: tdCommon{Type: "searchMessagesFilterUnreadMention"},
	}

	return &searchMessagesFilterUnreadMentionTemp
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

// NewChatActionTyping creates a new ChatActionTyping
//
func NewChatActionTyping() *ChatActionTyping {
	chatActionTypingTemp := ChatActionTyping{
		tdCommon: tdCommon{Type: "chatActionTyping"},
	}

	return &chatActionTypingTemp
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

// NewChatActionRecordingVideo creates a new ChatActionRecordingVideo
//
func NewChatActionRecordingVideo() *ChatActionRecordingVideo {
	chatActionRecordingVideoTemp := ChatActionRecordingVideo{
		tdCommon: tdCommon{Type: "chatActionRecordingVideo"},
	}

	return &chatActionRecordingVideoTemp
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

// NewChatActionUploadingVideo creates a new ChatActionUploadingVideo
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVideo(progress int32) *ChatActionUploadingVideo {
	chatActionUploadingVideoTemp := ChatActionUploadingVideo{
		tdCommon: tdCommon{Type: "chatActionUploadingVideo"},
		Progress: progress,
	}

	return &chatActionUploadingVideoTemp
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

// NewChatActionRecordingVoiceNote creates a new ChatActionRecordingVoiceNote
//
func NewChatActionRecordingVoiceNote() *ChatActionRecordingVoiceNote {
	chatActionRecordingVoiceNoteTemp := ChatActionRecordingVoiceNote{
		tdCommon: tdCommon{Type: "chatActionRecordingVoiceNote"},
	}

	return &chatActionRecordingVoiceNoteTemp
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

// NewChatActionUploadingVoiceNote creates a new ChatActionUploadingVoiceNote
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVoiceNote(progress int32) *ChatActionUploadingVoiceNote {
	chatActionUploadingVoiceNoteTemp := ChatActionUploadingVoiceNote{
		tdCommon: tdCommon{Type: "chatActionUploadingVoiceNote"},
		Progress: progress,
	}

	return &chatActionUploadingVoiceNoteTemp
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

// NewChatActionUploadingPhoto creates a new ChatActionUploadingPhoto
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingPhoto(progress int32) *ChatActionUploadingPhoto {
	chatActionUploadingPhotoTemp := ChatActionUploadingPhoto{
		tdCommon: tdCommon{Type: "chatActionUploadingPhoto"},
		Progress: progress,
	}

	return &chatActionUploadingPhotoTemp
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

// NewChatActionUploadingDocument creates a new ChatActionUploadingDocument
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingDocument(progress int32) *ChatActionUploadingDocument {
	chatActionUploadingDocumentTemp := ChatActionUploadingDocument{
		tdCommon: tdCommon{Type: "chatActionUploadingDocument"},
		Progress: progress,
	}

	return &chatActionUploadingDocumentTemp
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

// NewChatActionChoosingLocation creates a new ChatActionChoosingLocation
//
func NewChatActionChoosingLocation() *ChatActionChoosingLocation {
	chatActionChoosingLocationTemp := ChatActionChoosingLocation{
		tdCommon: tdCommon{Type: "chatActionChoosingLocation"},
	}

	return &chatActionChoosingLocationTemp
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

// NewChatActionChoosingContact creates a new ChatActionChoosingContact
//
func NewChatActionChoosingContact() *ChatActionChoosingContact {
	chatActionChoosingContactTemp := ChatActionChoosingContact{
		tdCommon: tdCommon{Type: "chatActionChoosingContact"},
	}

	return &chatActionChoosingContactTemp
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

// NewChatActionStartPlayingGame creates a new ChatActionStartPlayingGame
//
func NewChatActionStartPlayingGame() *ChatActionStartPlayingGame {
	chatActionStartPlayingGameTemp := ChatActionStartPlayingGame{
		tdCommon: tdCommon{Type: "chatActionStartPlayingGame"},
	}

	return &chatActionStartPlayingGameTemp
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

// NewChatActionRecordingVideoNote creates a new ChatActionRecordingVideoNote
//
func NewChatActionRecordingVideoNote() *ChatActionRecordingVideoNote {
	chatActionRecordingVideoNoteTemp := ChatActionRecordingVideoNote{
		tdCommon: tdCommon{Type: "chatActionRecordingVideoNote"},
	}

	return &chatActionRecordingVideoNoteTemp
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

// NewChatActionUploadingVideoNote creates a new ChatActionUploadingVideoNote
//
// @param progress Upload progress, as a percentage
func NewChatActionUploadingVideoNote(progress int32) *ChatActionUploadingVideoNote {
	chatActionUploadingVideoNoteTemp := ChatActionUploadingVideoNote{
		tdCommon: tdCommon{Type: "chatActionUploadingVideoNote"},
		Progress: progress,
	}

	return &chatActionUploadingVideoNoteTemp
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

// NewChatActionCancel creates a new ChatActionCancel
//
func NewChatActionCancel() *ChatActionCancel {
	chatActionCancelTemp := ChatActionCancel{
		tdCommon: tdCommon{Type: "chatActionCancel"},
	}

	return &chatActionCancelTemp
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

// NewUserStatusEmpty creates a new UserStatusEmpty
//
func NewUserStatusEmpty() *UserStatusEmpty {
	userStatusEmptyTemp := UserStatusEmpty{
		tdCommon: tdCommon{Type: "userStatusEmpty"},
	}

	return &userStatusEmptyTemp
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

// NewUserStatusOnline creates a new UserStatusOnline
//
// @param expires Point in time (Unix timestamp) when the user's online status will expire
func NewUserStatusOnline(expires int32) *UserStatusOnline {
	userStatusOnlineTemp := UserStatusOnline{
		tdCommon: tdCommon{Type: "userStatusOnline"},
		Expires:  expires,
	}

	return &userStatusOnlineTemp
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

// NewUserStatusOffline creates a new UserStatusOffline
//
// @param wasOnline Point in time (Unix timestamp) when the user was last online
func NewUserStatusOffline(wasOnline int32) *UserStatusOffline {
	userStatusOfflineTemp := UserStatusOffline{
		tdCommon:  tdCommon{Type: "userStatusOffline"},
		WasOnline: wasOnline,
	}

	return &userStatusOfflineTemp
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

// NewUserStatusRecently creates a new UserStatusRecently
//
func NewUserStatusRecently() *UserStatusRecently {
	userStatusRecentlyTemp := UserStatusRecently{
		tdCommon: tdCommon{Type: "userStatusRecently"},
	}

	return &userStatusRecentlyTemp
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

// NewUserStatusLastWeek creates a new UserStatusLastWeek
//
func NewUserStatusLastWeek() *UserStatusLastWeek {
	userStatusLastWeekTemp := UserStatusLastWeek{
		tdCommon: tdCommon{Type: "userStatusLastWeek"},
	}

	return &userStatusLastWeekTemp
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

// NewUserStatusLastMonth creates a new UserStatusLastMonth
//
func NewUserStatusLastMonth() *UserStatusLastMonth {
	userStatusLastMonthTemp := UserStatusLastMonth{
		tdCommon: tdCommon{Type: "userStatusLastMonth"},
	}

	return &userStatusLastMonthTemp
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

// NewStickers creates a new Stickers
//
// @param stickers List of stickers
func NewStickers(stickers []Sticker) *Stickers {
	stickersTemp := Stickers{
		tdCommon: tdCommon{Type: "stickers"},
		Stickers: stickers,
	}

	return &stickersTemp
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

// NewStickerEmojis creates a new StickerEmojis
//
// @param emojis List of emojis
func NewStickerEmojis(emojis []string) *StickerEmojis {
	stickerEmojisTemp := StickerEmojis{
		tdCommon: tdCommon{Type: "stickerEmojis"},
		Emojis:   emojis,
	}

	return &stickerEmojisTemp
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

// NewStickerSet creates a new StickerSet
//
// @param iD Identifier of the sticker set
// @param title Title of the sticker set
// @param name Name of the sticker set
// @param isInstalled True, if the sticker set has been installed by the current user
// @param isArchived True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
// @param isOfficial True, if the sticker set is official
// @param isMasks True, if the stickers in the set are masks
// @param isViewed True for already viewed trending sticker sets
// @param stickers List of stickers in this set
// @param emojis A list of emoji corresponding to the stickers in the same order
func NewStickerSet(iD JSONInt64, title string, name string, isInstalled bool, isArchived bool, isOfficial bool, isMasks bool, isViewed bool, stickers []Sticker, emojis []StickerEmojis) *StickerSet {
	stickerSetTemp := StickerSet{
		tdCommon:    tdCommon{Type: "stickerSet"},
		ID:          iD,
		Title:       title,
		Name:        name,
		IsInstalled: isInstalled,
		IsArchived:  isArchived,
		IsOfficial:  isOfficial,
		IsMasks:     isMasks,
		IsViewed:    isViewed,
		Stickers:    stickers,
		Emojis:      emojis,
	}

	return &stickerSetTemp
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

// NewStickerSetInfo creates a new StickerSetInfo
//
// @param iD Identifier of the sticker set
// @param title Title of the sticker set
// @param name Name of the sticker set
// @param isInstalled True, if the sticker set has been installed by current user
// @param isArchived True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously
// @param isOfficial True, if the sticker set is official
// @param isMasks True, if the stickers in the set are masks
// @param isViewed True for already viewed trending sticker sets
// @param size Total number of stickers in the set
// @param covers Contains up to the first 5 stickers from the set, depending on the context. If the client needs more stickers the full set should be requested
func NewStickerSetInfo(iD JSONInt64, title string, name string, isInstalled bool, isArchived bool, isOfficial bool, isMasks bool, isViewed bool, size int32, covers []Sticker) *StickerSetInfo {
	stickerSetInfoTemp := StickerSetInfo{
		tdCommon:    tdCommon{Type: "stickerSetInfo"},
		ID:          iD,
		Title:       title,
		Name:        name,
		IsInstalled: isInstalled,
		IsArchived:  isArchived,
		IsOfficial:  isOfficial,
		IsMasks:     isMasks,
		IsViewed:    isViewed,
		Size:        size,
		Covers:      covers,
	}

	return &stickerSetInfoTemp
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

// NewStickerSets creates a new StickerSets
//
// @param totalCount Approximate total number of sticker sets found
// @param sets List of sticker sets
func NewStickerSets(totalCount int32, sets []StickerSetInfo) *StickerSets {
	stickerSetsTemp := StickerSets{
		tdCommon:   tdCommon{Type: "stickerSets"},
		TotalCount: totalCount,
		Sets:       sets,
	}

	return &stickerSetsTemp
}

// CallDiscardReasonEmpty The call wasn't discarded, or the reason is unknown
type CallDiscardReasonEmpty struct {
	tdCommon
}

// MessageType return the string telegram-type of CallDiscardReasonEmpty
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) MessageType() string {
	return "callDiscardReasonEmpty"
}

// NewCallDiscardReasonEmpty creates a new CallDiscardReasonEmpty
//
func NewCallDiscardReasonEmpty() *CallDiscardReasonEmpty {
	callDiscardReasonEmptyTemp := CallDiscardReasonEmpty{
		tdCommon: tdCommon{Type: "callDiscardReasonEmpty"},
	}

	return &callDiscardReasonEmptyTemp
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

// NewCallDiscardReasonMissed creates a new CallDiscardReasonMissed
//
func NewCallDiscardReasonMissed() *CallDiscardReasonMissed {
	callDiscardReasonMissedTemp := CallDiscardReasonMissed{
		tdCommon: tdCommon{Type: "callDiscardReasonMissed"},
	}

	return &callDiscardReasonMissedTemp
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

// NewCallDiscardReasonDeclined creates a new CallDiscardReasonDeclined
//
func NewCallDiscardReasonDeclined() *CallDiscardReasonDeclined {
	callDiscardReasonDeclinedTemp := CallDiscardReasonDeclined{
		tdCommon: tdCommon{Type: "callDiscardReasonDeclined"},
	}

	return &callDiscardReasonDeclinedTemp
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

// NewCallDiscardReasonDisconnected creates a new CallDiscardReasonDisconnected
//
func NewCallDiscardReasonDisconnected() *CallDiscardReasonDisconnected {
	callDiscardReasonDisconnectedTemp := CallDiscardReasonDisconnected{
		tdCommon: tdCommon{Type: "callDiscardReasonDisconnected"},
	}

	return &callDiscardReasonDisconnectedTemp
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

// NewCallDiscardReasonHungUp creates a new CallDiscardReasonHungUp
//
func NewCallDiscardReasonHungUp() *CallDiscardReasonHungUp {
	callDiscardReasonHungUpTemp := CallDiscardReasonHungUp{
		tdCommon: tdCommon{Type: "callDiscardReasonHungUp"},
	}

	return &callDiscardReasonHungUpTemp
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

// NewCallProtocol creates a new CallProtocol
//
// @param uDPP2p True, if UDP peer-to-peer connections are supported
// @param uDPReflector True, if connection through UDP reflectors is supported
// @param minLayer Minimum supported API layer; use 65
// @param maxLayer Maximum supported API layer; use 65
func NewCallProtocol(uDPP2p bool, uDPReflector bool, minLayer int32, maxLayer int32) *CallProtocol {
	callProtocolTemp := CallProtocol{
		tdCommon:     tdCommon{Type: "callProtocol"},
		UDPP2p:       uDPP2p,
		UDPReflector: uDPReflector,
		MinLayer:     minLayer,
		MaxLayer:     maxLayer,
	}

	return &callProtocolTemp
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

// NewCallConnection creates a new CallConnection
//
// @param iD Reflector identifier
// @param iP IPv4 reflector address
// @param iPv6 IPv6 reflector address
// @param port Reflector port number
// @param peerTag Connection peer tag
func NewCallConnection(iD JSONInt64, iP string, iPv6 string, port int32, peerTag []byte) *CallConnection {
	callConnectionTemp := CallConnection{
		tdCommon: tdCommon{Type: "callConnection"},
		ID:       iD,
		IP:       iP,
		IPv6:     iPv6,
		Port:     port,
		PeerTag:  peerTag,
	}

	return &callConnectionTemp
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

// NewCallID creates a new CallID
//
// @param iD Call identifier
func NewCallID(iD int32) *CallID {
	callIDTemp := CallID{
		tdCommon: tdCommon{Type: "callId"},
		ID:       iD,
	}

	return &callIDTemp
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

// NewCallStatePending creates a new CallStatePending
//
// @param isCreated True, if the call has already been created by the server
// @param isReceived True, if the call has already been received by the other party
func NewCallStatePending(isCreated bool, isReceived bool) *CallStatePending {
	callStatePendingTemp := CallStatePending{
		tdCommon:   tdCommon{Type: "callStatePending"},
		IsCreated:  isCreated,
		IsReceived: isReceived,
	}

	return &callStatePendingTemp
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

// NewCallStateExchangingKeys creates a new CallStateExchangingKeys
//
func NewCallStateExchangingKeys() *CallStateExchangingKeys {
	callStateExchangingKeysTemp := CallStateExchangingKeys{
		tdCommon: tdCommon{Type: "callStateExchangingKeys"},
	}

	return &callStateExchangingKeysTemp
}

// GetCallStateEnum return the enum type of this object
func (callStateExchangingKeys *CallStateExchangingKeys) GetCallStateEnum() CallStateEnum {
	return CallStateExchangingKeysType
}

// CallStateReady The call is ready to use
type CallStateReady struct {
	tdCommon
	Protocol      *CallProtocol    `json:"protocol"`       // Call protocols supported by the peer
	Connections   []CallConnection `json:"connections"`    // Available UDP reflectors
	Config        string           `json:"config"`         // A JSON-encoded call config
	EncryptionKey []byte           `json:"encryption_key"` // Call encryption key
	Emojis        []string         `json:"emojis"`         // Encryption key emojis fingerprint
}

// MessageType return the string telegram-type of CallStateReady
func (callStateReady *CallStateReady) MessageType() string {
	return "callStateReady"
}

// NewCallStateReady creates a new CallStateReady
//
// @param protocol Call protocols supported by the peer
// @param connections Available UDP reflectors
// @param config A JSON-encoded call config
// @param encryptionKey Call encryption key
// @param emojis Encryption key emojis fingerprint
func NewCallStateReady(protocol *CallProtocol, connections []CallConnection, config string, encryptionKey []byte, emojis []string) *CallStateReady {
	callStateReadyTemp := CallStateReady{
		tdCommon:      tdCommon{Type: "callStateReady"},
		Protocol:      protocol,
		Connections:   connections,
		Config:        config,
		EncryptionKey: encryptionKey,
		Emojis:        emojis,
	}

	return &callStateReadyTemp
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

// NewCallStateHangingUp creates a new CallStateHangingUp
//
func NewCallStateHangingUp() *CallStateHangingUp {
	callStateHangingUpTemp := CallStateHangingUp{
		tdCommon: tdCommon{Type: "callStateHangingUp"},
	}

	return &callStateHangingUpTemp
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

// NewCallStateDiscarded creates a new CallStateDiscarded
//
// @param reason The reason, why the call has ended
// @param needRating True, if the call rating should be sent to the server
// @param needDebugInformation True, if the call debug information should be sent to the server
func NewCallStateDiscarded(reason CallDiscardReason, needRating bool, needDebugInformation bool) *CallStateDiscarded {
	callStateDiscardedTemp := CallStateDiscarded{
		tdCommon:             tdCommon{Type: "callStateDiscarded"},
		Reason:               reason,
		NeedRating:           needRating,
		NeedDebugInformation: needDebugInformation,
	}

	return &callStateDiscardedTemp
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
	Error *Error `json:"error"` // Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
}

// MessageType return the string telegram-type of CallStateError
func (callStateError *CallStateError) MessageType() string {
	return "callStateError"
}

// NewCallStateError creates a new CallStateError
//
// @param error Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
func NewCallStateError(error *Error) *CallStateError {
	callStateErrorTemp := CallStateError{
		tdCommon: tdCommon{Type: "callStateError"},
		Error:    error,
	}

	return &callStateErrorTemp
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

// NewCall creates a new Call
//
// @param iD Call identifier, not persistent
// @param userID Peer user identifier
// @param isOutgoing True, if the call is outgoing
// @param state Call state
func NewCall(iD int32, userID int32, isOutgoing bool, state CallState) *Call {
	callTemp := Call{
		tdCommon:   tdCommon{Type: "call"},
		ID:         iD,
		UserID:     userID,
		IsOutgoing: isOutgoing,
		State:      state,
	}

	return &callTemp
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

// NewAnimations creates a new Animations
//
// @param animations List of animations
func NewAnimations(animations []Animation) *Animations {
	animationsTemp := Animations{
		tdCommon:   tdCommon{Type: "animations"},
		Animations: animations,
	}

	return &animationsTemp
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

// NewImportedContacts creates a new ImportedContacts
//
// @param userIDs User identifiers of the imported contacts in the same order as they were specified in the request; 0 if the contact is not yet a registered user
// @param importerCount The number of users that imported the corresponding contact; 0 for already registered users or if unavailable
func NewImportedContacts(userIDs []int32, importerCount []int32) *ImportedContacts {
	importedContactsTemp := ImportedContacts{
		tdCommon:      tdCommon{Type: "importedContacts"},
		UserIDs:       userIDs,
		ImporterCount: importerCount,
	}

	return &importedContactsTemp
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

// NewInputInlineQueryResultAnimatedGif creates a new InputInlineQueryResultAnimatedGif
//
// @param iD Unique identifier of the query result
// @param title Title of the query result
// @param thumbnailURL URL of the static result thumbnail (JPEG or GIF), if it exists
// @param gifURL The URL of the GIF-file (file size must not exceed 1MB)
// @param gifDuration Duration of the GIF, in seconds
// @param gifWidth Width of the GIF
// @param gifHeight Height of the GIF
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultAnimatedGif(iD string, title string, thumbnailURL string, gifURL string, gifDuration int32, gifWidth int32, gifHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultAnimatedGif {
	inputInlineQueryResultAnimatedGifTemp := InputInlineQueryResultAnimatedGif{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultAnimatedGif"},
		ID:                  iD,
		Title:               title,
		ThumbnailURL:        thumbnailURL,
		GifURL:              gifURL,
		GifDuration:         gifDuration,
		GifWidth:            gifWidth,
		GifHeight:           gifHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultAnimatedGifTemp
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

// NewInputInlineQueryResultAnimatedMpeg4 creates a new InputInlineQueryResultAnimatedMpeg4
//
// @param iD Unique identifier of the query result
// @param title Title of the result
// @param thumbnailURL URL of the static result thumbnail (JPEG or GIF), if it exists
// @param mpeg4URL The URL of the MPEG4-file (file size must not exceed 1MB)
// @param mpeg4Duration Duration of the video, in seconds
// @param mpeg4Width Width of the video
// @param mpeg4Height Height of the video
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultAnimatedMpeg4(iD string, title string, thumbnailURL string, mpeg4URL string, mpeg4Duration int32, mpeg4Width int32, mpeg4Height int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultAnimatedMpeg4 {
	inputInlineQueryResultAnimatedMpeg4Temp := InputInlineQueryResultAnimatedMpeg4{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultAnimatedMpeg4"},
		ID:                  iD,
		Title:               title,
		ThumbnailURL:        thumbnailURL,
		Mpeg4URL:            mpeg4URL,
		Mpeg4Duration:       mpeg4Duration,
		Mpeg4Width:          mpeg4Width,
		Mpeg4Height:         mpeg4Height,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultAnimatedMpeg4Temp
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

// NewInputInlineQueryResultArticle creates a new InputInlineQueryResultArticle
//
// @param iD Unique identifier of the query result
// @param uRL URL of the result, if it exists
// @param hideURL True, if the URL must be not shown
// @param title Title of the result
// @param description
// @param thumbnailURL URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultArticle(iD string, uRL string, hideURL bool, title string, description string, thumbnailURL string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultArticle {
	inputInlineQueryResultArticleTemp := InputInlineQueryResultArticle{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultArticle"},
		ID:                  iD,
		URL:                 uRL,
		HideURL:             hideURL,
		Title:               title,
		Description:         description,
		ThumbnailURL:        thumbnailURL,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultArticleTemp
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

// NewInputInlineQueryResultAudio creates a new InputInlineQueryResultAudio
//
// @param iD Unique identifier of the query result
// @param title Title of the audio file
// @param performer Performer of the audio file
// @param audioURL The URL of the audio file
// @param audioDuration Audio file duration, in seconds
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAudio, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultAudio(iD string, title string, performer string, audioURL string, audioDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultAudio {
	inputInlineQueryResultAudioTemp := InputInlineQueryResultAudio{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultAudio"},
		ID:                  iD,
		Title:               title,
		Performer:           performer,
		AudioURL:            audioURL,
		AudioDuration:       audioDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultAudioTemp
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
	Contact             *Contact            `json:"contact"`               // User contact
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

// NewInputInlineQueryResultContact creates a new InputInlineQueryResultContact
//
// @param iD Unique identifier of the query result
// @param contact User contact
// @param thumbnailURL URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultContact(iD string, contact *Contact, thumbnailURL string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultContact {
	inputInlineQueryResultContactTemp := InputInlineQueryResultContact{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultContact"},
		ID:                  iD,
		Contact:             contact,
		ThumbnailURL:        thumbnailURL,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultContactTemp
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
		ID              string   `json:"id"`               // Unique identifier of the query result
		Contact         *Contact `json:"contact"`          // User contact
		ThumbnailURL    string   `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32    `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32    `json:"thumbnail_height"` // Thumbnail height, if known

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

// NewInputInlineQueryResultDocument creates a new InputInlineQueryResultDocument
//
// @param iD Unique identifier of the query result
// @param title Title of the resulting file
// @param description
// @param documentURL URL of the file
// @param mimeType MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
// @param thumbnailURL The URL of the file thumbnail, if it exists
// @param thumbnailWidth Width of the thumbnail
// @param thumbnailHeight Height of the thumbnail
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageDocument, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultDocument(iD string, title string, description string, documentURL string, mimeType string, thumbnailURL string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultDocument {
	inputInlineQueryResultDocumentTemp := InputInlineQueryResultDocument{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultDocument"},
		ID:                  iD,
		Title:               title,
		Description:         description,
		DocumentURL:         documentURL,
		MimeType:            mimeType,
		ThumbnailURL:        thumbnailURL,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultDocumentTemp
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

// NewInputInlineQueryResultGame creates a new InputInlineQueryResultGame
//
// @param iD Unique identifier of the query result
// @param gameShortName Short name of the game
// @param replyMarkup Message reply markup. Must be of type replyMarkupInlineKeyboard or null
func NewInputInlineQueryResultGame(iD string, gameShortName string, replyMarkup ReplyMarkup) *InputInlineQueryResultGame {
	inputInlineQueryResultGameTemp := InputInlineQueryResultGame{
		tdCommon:      tdCommon{Type: "inputInlineQueryResultGame"},
		ID:            iD,
		GameShortName: gameShortName,
		ReplyMarkup:   replyMarkup,
	}

	return &inputInlineQueryResultGameTemp
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
	Location            *Location           `json:"location"`              // Location result
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

// NewInputInlineQueryResultLocation creates a new InputInlineQueryResultLocation
//
// @param iD Unique identifier of the query result
// @param location Location result
// @param livePeriod Amount of time relative to the message sent time until the location can be updated, in seconds
// @param title Title of the result
// @param thumbnailURL URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultLocation(iD string, location *Location, livePeriod int32, title string, thumbnailURL string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultLocation {
	inputInlineQueryResultLocationTemp := InputInlineQueryResultLocation{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultLocation"},
		ID:                  iD,
		Location:            location,
		LivePeriod:          livePeriod,
		Title:               title,
		ThumbnailURL:        thumbnailURL,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultLocationTemp
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
		ID              string    `json:"id"`               // Unique identifier of the query result
		Location        *Location `json:"location"`         // Location result
		LivePeriod      int32     `json:"live_period"`      // Amount of time relative to the message sent time until the location can be updated, in seconds
		Title           string    `json:"title"`            // Title of the result
		ThumbnailURL    string    `json:"thumbnail_url"`    // URL of the result thumbnail, if it exists
		ThumbnailWidth  int32     `json:"thumbnail_width"`  // Thumbnail width, if known
		ThumbnailHeight int32     `json:"thumbnail_height"` // Thumbnail height, if known

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

// NewInputInlineQueryResultPhoto creates a new InputInlineQueryResultPhoto
//
// @param iD Unique identifier of the query result
// @param title Title of the result, if known
// @param description
// @param thumbnailURL URL of the photo thumbnail, if it exists
// @param photoURL The URL of the JPEG photo (photo size must not exceed 5MB)
// @param photoWidth Width of the photo
// @param photoHeight Height of the photo
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessagePhoto, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultPhoto(iD string, title string, description string, thumbnailURL string, photoURL string, photoWidth int32, photoHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultPhoto {
	inputInlineQueryResultPhotoTemp := InputInlineQueryResultPhoto{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultPhoto"},
		ID:                  iD,
		Title:               title,
		Description:         description,
		ThumbnailURL:        thumbnailURL,
		PhotoURL:            photoURL,
		PhotoWidth:          photoWidth,
		PhotoHeight:         photoHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultPhotoTemp
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

// NewInputInlineQueryResultSticker creates a new InputInlineQueryResultSticker
//
// @param iD Unique identifier of the query result
// @param thumbnailURL URL of the sticker thumbnail, if it exists
// @param stickerURL The URL of the WEBP sticker (sticker file size must not exceed 5MB)
// @param stickerWidth Width of the sticker
// @param stickerHeight Height of the sticker
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, inputMessageSticker, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultSticker(iD string, thumbnailURL string, stickerURL string, stickerWidth int32, stickerHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultSticker {
	inputInlineQueryResultStickerTemp := InputInlineQueryResultSticker{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultSticker"},
		ID:                  iD,
		ThumbnailURL:        thumbnailURL,
		StickerURL:          stickerURL,
		StickerWidth:        stickerWidth,
		StickerHeight:       stickerHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultStickerTemp
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
	Venue               *Venue              `json:"venue"`                 // Venue result
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

// NewInputInlineQueryResultVenue creates a new InputInlineQueryResultVenue
//
// @param iD Unique identifier of the query result
// @param venue Venue result
// @param thumbnailURL URL of the result thumbnail, if it exists
// @param thumbnailWidth Thumbnail width, if known
// @param thumbnailHeight Thumbnail height, if known
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultVenue(iD string, venue *Venue, thumbnailURL string, thumbnailWidth int32, thumbnailHeight int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVenue {
	inputInlineQueryResultVenueTemp := InputInlineQueryResultVenue{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVenue"},
		ID:                  iD,
		Venue:               venue,
		ThumbnailURL:        thumbnailURL,
		ThumbnailWidth:      thumbnailWidth,
		ThumbnailHeight:     thumbnailHeight,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVenueTemp
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
		Venue           *Venue `json:"venue"`            // Venue result
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

// NewInputInlineQueryResultVideo creates a new InputInlineQueryResultVideo
//
// @param iD Unique identifier of the query result
// @param title Title of the result
// @param description
// @param thumbnailURL The URL of the video thumbnail (JPEG), if it exists
// @param videoURL URL of the embedded video player or video file
// @param mimeType MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
// @param videoWidth Width of the video
// @param videoHeight Height of the video
// @param videoDuration Video duration, in seconds
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVideo, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultVideo(iD string, title string, description string, thumbnailURL string, videoURL string, mimeType string, videoWidth int32, videoHeight int32, videoDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVideo {
	inputInlineQueryResultVideoTemp := InputInlineQueryResultVideo{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVideo"},
		ID:                  iD,
		Title:               title,
		Description:         description,
		ThumbnailURL:        thumbnailURL,
		VideoURL:            videoURL,
		MimeType:            mimeType,
		VideoWidth:          videoWidth,
		VideoHeight:         videoHeight,
		VideoDuration:       videoDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVideoTemp
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

// NewInputInlineQueryResultVoiceNote creates a new InputInlineQueryResultVoiceNote
//
// @param iD Unique identifier of the query result
// @param title Title of the voice note
// @param voiceNoteURL The URL of the voice note file
// @param voiceNoteDuration Duration of the voice note, in seconds
// @param replyMarkup The message reply markup. Must be of type replyMarkupInlineKeyboard or null
// @param inputMessageContent The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVoiceNote, InputMessageLocation, InputMessageVenue or InputMessageContact
func NewInputInlineQueryResultVoiceNote(iD string, title string, voiceNoteURL string, voiceNoteDuration int32, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) *InputInlineQueryResultVoiceNote {
	inputInlineQueryResultVoiceNoteTemp := InputInlineQueryResultVoiceNote{
		tdCommon:            tdCommon{Type: "inputInlineQueryResultVoiceNote"},
		ID:                  iD,
		Title:               title,
		VoiceNoteURL:        voiceNoteURL,
		VoiceNoteDuration:   voiceNoteDuration,
		ReplyMarkup:         replyMarkup,
		InputMessageContent: inputMessageContent,
	}

	return &inputInlineQueryResultVoiceNoteTemp
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
	ID          string     `json:"id"`          // Unique identifier of the query result
	URL         string     `json:"url"`         // URL of the result, if it exists
	HideURL     bool       `json:"hide_url"`    // True, if the URL must be not shown
	Title       string     `json:"title"`       // Title of the result
	Description string     `json:"description"` //
	Thumbnail   *PhotoSize `json:"thumbnail"`   // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultArticle
func (inlineQueryResultArticle *InlineQueryResultArticle) MessageType() string {
	return "inlineQueryResultArticle"
}

// NewInlineQueryResultArticle creates a new InlineQueryResultArticle
//
// @param iD Unique identifier of the query result
// @param uRL URL of the result, if it exists
// @param hideURL True, if the URL must be not shown
// @param title Title of the result
// @param description
// @param thumbnail Result thumbnail; may be null
func NewInlineQueryResultArticle(iD string, uRL string, hideURL bool, title string, description string, thumbnail *PhotoSize) *InlineQueryResultArticle {
	inlineQueryResultArticleTemp := InlineQueryResultArticle{
		tdCommon:    tdCommon{Type: "inlineQueryResultArticle"},
		ID:          iD,
		URL:         uRL,
		HideURL:     hideURL,
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
	}

	return &inlineQueryResultArticleTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultArticle *InlineQueryResultArticle) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultArticleType
}

// InlineQueryResultContact Represents a user contact
type InlineQueryResultContact struct {
	tdCommon
	ID        string     `json:"id"`        // Unique identifier of the query result
	Contact   *Contact   `json:"contact"`   // A user contact
	Thumbnail *PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultContact
func (inlineQueryResultContact *InlineQueryResultContact) MessageType() string {
	return "inlineQueryResultContact"
}

// NewInlineQueryResultContact creates a new InlineQueryResultContact
//
// @param iD Unique identifier of the query result
// @param contact A user contact
// @param thumbnail Result thumbnail; may be null
func NewInlineQueryResultContact(iD string, contact *Contact, thumbnail *PhotoSize) *InlineQueryResultContact {
	inlineQueryResultContactTemp := InlineQueryResultContact{
		tdCommon:  tdCommon{Type: "inlineQueryResultContact"},
		ID:        iD,
		Contact:   contact,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultContactTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultContact *InlineQueryResultContact) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultContactType
}

// InlineQueryResultLocation Represents a point on the map
type InlineQueryResultLocation struct {
	tdCommon
	ID        string     `json:"id"`        // Unique identifier of the query result
	Location  *Location  `json:"location"`  // Location result
	Title     string     `json:"title"`     // Title of the result
	Thumbnail *PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultLocation
func (inlineQueryResultLocation *InlineQueryResultLocation) MessageType() string {
	return "inlineQueryResultLocation"
}

// NewInlineQueryResultLocation creates a new InlineQueryResultLocation
//
// @param iD Unique identifier of the query result
// @param location Location result
// @param title Title of the result
// @param thumbnail Result thumbnail; may be null
func NewInlineQueryResultLocation(iD string, location *Location, title string, thumbnail *PhotoSize) *InlineQueryResultLocation {
	inlineQueryResultLocationTemp := InlineQueryResultLocation{
		tdCommon:  tdCommon{Type: "inlineQueryResultLocation"},
		ID:        iD,
		Location:  location,
		Title:     title,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultLocationTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultLocation *InlineQueryResultLocation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultLocationType
}

// InlineQueryResultVenue Represents information about a venue
type InlineQueryResultVenue struct {
	tdCommon
	ID        string     `json:"id"`        // Unique identifier of the query result
	Venue     *Venue     `json:"venue"`     // Venue result
	Thumbnail *PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultVenue
func (inlineQueryResultVenue *InlineQueryResultVenue) MessageType() string {
	return "inlineQueryResultVenue"
}

// NewInlineQueryResultVenue creates a new InlineQueryResultVenue
//
// @param iD Unique identifier of the query result
// @param venue Venue result
// @param thumbnail Result thumbnail; may be null
func NewInlineQueryResultVenue(iD string, venue *Venue, thumbnail *PhotoSize) *InlineQueryResultVenue {
	inlineQueryResultVenueTemp := InlineQueryResultVenue{
		tdCommon:  tdCommon{Type: "inlineQueryResultVenue"},
		ID:        iD,
		Venue:     venue,
		Thumbnail: thumbnail,
	}

	return &inlineQueryResultVenueTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVenue *InlineQueryResultVenue) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVenueType
}

// InlineQueryResultGame Represents information about a game
type InlineQueryResultGame struct {
	tdCommon
	ID   string `json:"id"`   // Unique identifier of the query result
	Game *Game  `json:"game"` // Game result
}

// MessageType return the string telegram-type of InlineQueryResultGame
func (inlineQueryResultGame *InlineQueryResultGame) MessageType() string {
	return "inlineQueryResultGame"
}

// NewInlineQueryResultGame creates a new InlineQueryResultGame
//
// @param iD Unique identifier of the query result
// @param game Game result
func NewInlineQueryResultGame(iD string, game *Game) *InlineQueryResultGame {
	inlineQueryResultGameTemp := InlineQueryResultGame{
		tdCommon: tdCommon{Type: "inlineQueryResultGame"},
		ID:       iD,
		Game:     game,
	}

	return &inlineQueryResultGameTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultGame *InlineQueryResultGame) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultGameType
}

// InlineQueryResultAnimation Represents an animation file
type InlineQueryResultAnimation struct {
	tdCommon
	ID        string     `json:"id"`        // Unique identifier of the query result
	Animation *Animation `json:"animation"` // Animation file
	Title     string     `json:"title"`     // Animation title
}

// MessageType return the string telegram-type of InlineQueryResultAnimation
func (inlineQueryResultAnimation *InlineQueryResultAnimation) MessageType() string {
	return "inlineQueryResultAnimation"
}

// NewInlineQueryResultAnimation creates a new InlineQueryResultAnimation
//
// @param iD Unique identifier of the query result
// @param animation Animation file
// @param title Animation title
func NewInlineQueryResultAnimation(iD string, animation *Animation, title string) *InlineQueryResultAnimation {
	inlineQueryResultAnimationTemp := InlineQueryResultAnimation{
		tdCommon:  tdCommon{Type: "inlineQueryResultAnimation"},
		ID:        iD,
		Animation: animation,
		Title:     title,
	}

	return &inlineQueryResultAnimationTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAnimation *InlineQueryResultAnimation) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAnimationType
}

// InlineQueryResultAudio Represents an audio file
type InlineQueryResultAudio struct {
	tdCommon
	ID    string `json:"id"`    // Unique identifier of the query result
	Audio *Audio `json:"audio"` // Audio file
}

// MessageType return the string telegram-type of InlineQueryResultAudio
func (inlineQueryResultAudio *InlineQueryResultAudio) MessageType() string {
	return "inlineQueryResultAudio"
}

// NewInlineQueryResultAudio creates a new InlineQueryResultAudio
//
// @param iD Unique identifier of the query result
// @param audio Audio file
func NewInlineQueryResultAudio(iD string, audio *Audio) *InlineQueryResultAudio {
	inlineQueryResultAudioTemp := InlineQueryResultAudio{
		tdCommon: tdCommon{Type: "inlineQueryResultAudio"},
		ID:       iD,
		Audio:    audio,
	}

	return &inlineQueryResultAudioTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultAudio *InlineQueryResultAudio) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultAudioType
}

// InlineQueryResultDocument Represents a document
type InlineQueryResultDocument struct {
	tdCommon
	ID          string    `json:"id"`          // Unique identifier of the query result
	Document    *Document `json:"document"`    // Document
	Title       string    `json:"title"`       // Document title
	Description string    `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultDocument
func (inlineQueryResultDocument *InlineQueryResultDocument) MessageType() string {
	return "inlineQueryResultDocument"
}

// NewInlineQueryResultDocument creates a new InlineQueryResultDocument
//
// @param iD Unique identifier of the query result
// @param document Document
// @param title Document title
// @param description
func NewInlineQueryResultDocument(iD string, document *Document, title string, description string) *InlineQueryResultDocument {
	inlineQueryResultDocumentTemp := InlineQueryResultDocument{
		tdCommon:    tdCommon{Type: "inlineQueryResultDocument"},
		ID:          iD,
		Document:    document,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultDocumentTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultDocument *InlineQueryResultDocument) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultDocumentType
}

// InlineQueryResultPhoto Represents a photo
type InlineQueryResultPhoto struct {
	tdCommon
	ID          string `json:"id"`          // Unique identifier of the query result
	Photo       *Photo `json:"photo"`       // Photo
	Title       string `json:"title"`       // Title of the result, if known
	Description string `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultPhoto
func (inlineQueryResultPhoto *InlineQueryResultPhoto) MessageType() string {
	return "inlineQueryResultPhoto"
}

// NewInlineQueryResultPhoto creates a new InlineQueryResultPhoto
//
// @param iD Unique identifier of the query result
// @param photo Photo
// @param title Title of the result, if known
// @param description
func NewInlineQueryResultPhoto(iD string, photo *Photo, title string, description string) *InlineQueryResultPhoto {
	inlineQueryResultPhotoTemp := InlineQueryResultPhoto{
		tdCommon:    tdCommon{Type: "inlineQueryResultPhoto"},
		ID:          iD,
		Photo:       photo,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultPhotoTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultPhoto *InlineQueryResultPhoto) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultPhotoType
}

// InlineQueryResultSticker Represents a sticker
type InlineQueryResultSticker struct {
	tdCommon
	ID      string   `json:"id"`      // Unique identifier of the query result
	Sticker *Sticker `json:"sticker"` // Sticker
}

// MessageType return the string telegram-type of InlineQueryResultSticker
func (inlineQueryResultSticker *InlineQueryResultSticker) MessageType() string {
	return "inlineQueryResultSticker"
}

// NewInlineQueryResultSticker creates a new InlineQueryResultSticker
//
// @param iD Unique identifier of the query result
// @param sticker Sticker
func NewInlineQueryResultSticker(iD string, sticker *Sticker) *InlineQueryResultSticker {
	inlineQueryResultStickerTemp := InlineQueryResultSticker{
		tdCommon: tdCommon{Type: "inlineQueryResultSticker"},
		ID:       iD,
		Sticker:  sticker,
	}

	return &inlineQueryResultStickerTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultSticker *InlineQueryResultSticker) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultStickerType
}

// InlineQueryResultVideo Represents a video
type InlineQueryResultVideo struct {
	tdCommon
	ID          string `json:"id"`          // Unique identifier of the query result
	Video       *Video `json:"video"`       // Video
	Title       string `json:"title"`       // Title of the video
	Description string `json:"description"` //
}

// MessageType return the string telegram-type of InlineQueryResultVideo
func (inlineQueryResultVideo *InlineQueryResultVideo) MessageType() string {
	return "inlineQueryResultVideo"
}

// NewInlineQueryResultVideo creates a new InlineQueryResultVideo
//
// @param iD Unique identifier of the query result
// @param video Video
// @param title Title of the video
// @param description
func NewInlineQueryResultVideo(iD string, video *Video, title string, description string) *InlineQueryResultVideo {
	inlineQueryResultVideoTemp := InlineQueryResultVideo{
		tdCommon:    tdCommon{Type: "inlineQueryResultVideo"},
		ID:          iD,
		Video:       video,
		Title:       title,
		Description: description,
	}

	return &inlineQueryResultVideoTemp
}

// GetInlineQueryResultEnum return the enum type of this object
func (inlineQueryResultVideo *InlineQueryResultVideo) GetInlineQueryResultEnum() InlineQueryResultEnum {
	return InlineQueryResultVideoType
}

// InlineQueryResultVoiceNote Represents a voice note
type InlineQueryResultVoiceNote struct {
	tdCommon
	ID        string     `json:"id"`         // Unique identifier of the query result
	VoiceNote *VoiceNote `json:"voice_note"` // Voice note
	Title     string     `json:"title"`      // Title of the voice note
}

// MessageType return the string telegram-type of InlineQueryResultVoiceNote
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) MessageType() string {
	return "inlineQueryResultVoiceNote"
}

// NewInlineQueryResultVoiceNote creates a new InlineQueryResultVoiceNote
//
// @param iD Unique identifier of the query result
// @param voiceNote Voice note
// @param title Title of the voice note
func NewInlineQueryResultVoiceNote(iD string, voiceNote *VoiceNote, title string) *InlineQueryResultVoiceNote {
	inlineQueryResultVoiceNoteTemp := InlineQueryResultVoiceNote{
		tdCommon:  tdCommon{Type: "inlineQueryResultVoiceNote"},
		ID:        iD,
		VoiceNote: voiceNote,
		Title:     title,
	}

	return &inlineQueryResultVoiceNoteTemp
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

// NewInlineQueryResults creates a new InlineQueryResults
//
// @param inlineQueryID Unique identifier of the inline query
// @param nextOffset The offset for the next request. If empty, there are no more results
// @param results Results of the query
// @param switchPmText If non-empty, this text should be shown on the button, which opens a private chat with the bot and sends the bot a start message with the switch_pm_parameter
// @param switchPmParameter Parameter for the bot start message
func NewInlineQueryResults(inlineQueryID JSONInt64, nextOffset string, results []InlineQueryResult, switchPmText string, switchPmParameter string) *InlineQueryResults {
	inlineQueryResultsTemp := InlineQueryResults{
		tdCommon:          tdCommon{Type: "inlineQueryResults"},
		InlineQueryID:     inlineQueryID,
		NextOffset:        nextOffset,
		Results:           results,
		SwitchPmText:      switchPmText,
		SwitchPmParameter: switchPmParameter,
	}

	return &inlineQueryResultsTemp
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

// NewCallbackQueryPayloadData creates a new CallbackQueryPayloadData
//
// @param data Data that was attached to the callback button
func NewCallbackQueryPayloadData(data []byte) *CallbackQueryPayloadData {
	callbackQueryPayloadDataTemp := CallbackQueryPayloadData{
		tdCommon: tdCommon{Type: "callbackQueryPayloadData"},
		Data:     data,
	}

	return &callbackQueryPayloadDataTemp
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

// NewCallbackQueryPayloadGame creates a new CallbackQueryPayloadGame
//
// @param gameShortName A short name of the game that was attached to the callback button
func NewCallbackQueryPayloadGame(gameShortName string) *CallbackQueryPayloadGame {
	callbackQueryPayloadGameTemp := CallbackQueryPayloadGame{
		tdCommon:      tdCommon{Type: "callbackQueryPayloadGame"},
		GameShortName: gameShortName,
	}

	return &callbackQueryPayloadGameTemp
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

// NewCallbackQueryAnswer creates a new CallbackQueryAnswer
//
// @param text Text of the answer
// @param showAlert True, if an alert should be shown to the user instead of a toast notification
// @param uRL URL to be opened
func NewCallbackQueryAnswer(text string, showAlert bool, uRL string) *CallbackQueryAnswer {
	callbackQueryAnswerTemp := CallbackQueryAnswer{
		tdCommon:  tdCommon{Type: "callbackQueryAnswer"},
		Text:      text,
		ShowAlert: showAlert,
		URL:       uRL,
	}

	return &callbackQueryAnswerTemp
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

// NewCustomRequestResult creates a new CustomRequestResult
//
// @param result A JSON-serialized result
func NewCustomRequestResult(result string) *CustomRequestResult {
	customRequestResultTemp := CustomRequestResult{
		tdCommon: tdCommon{Type: "customRequestResult"},
		Result:   result,
	}

	return &customRequestResultTemp
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

// NewGameHighScore creates a new GameHighScore
//
// @param position Position in the high score table
// @param userID User identifier
// @param score User score
func NewGameHighScore(position int32, userID int32, score int32) *GameHighScore {
	gameHighScoreTemp := GameHighScore{
		tdCommon: tdCommon{Type: "gameHighScore"},
		Position: position,
		UserID:   userID,
		Score:    score,
	}

	return &gameHighScoreTemp
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

// NewGameHighScores creates a new GameHighScores
//
// @param scores A list of game high scores
func NewGameHighScores(scores []GameHighScore) *GameHighScores {
	gameHighScoresTemp := GameHighScores{
		tdCommon: tdCommon{Type: "gameHighScores"},
		Scores:   scores,
	}

	return &gameHighScoresTemp
}

// ChatEventMessageEdited A message was edited
type ChatEventMessageEdited struct {
	tdCommon
	OldMessage *Message `json:"old_message"` // The original message before the edit
	NewMessage *Message `json:"new_message"` // The message after it was edited
}

// MessageType return the string telegram-type of ChatEventMessageEdited
func (chatEventMessageEdited *ChatEventMessageEdited) MessageType() string {
	return "chatEventMessageEdited"
}

// NewChatEventMessageEdited creates a new ChatEventMessageEdited
//
// @param oldMessage The original message before the edit
// @param newMessage The message after it was edited
func NewChatEventMessageEdited(oldMessage *Message, newMessage *Message) *ChatEventMessageEdited {
	chatEventMessageEditedTemp := ChatEventMessageEdited{
		tdCommon:   tdCommon{Type: "chatEventMessageEdited"},
		OldMessage: oldMessage,
		NewMessage: newMessage,
	}

	return &chatEventMessageEditedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageEdited *ChatEventMessageEdited) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageEditedType
}

// ChatEventMessageDeleted A message was deleted
type ChatEventMessageDeleted struct {
	tdCommon
	Message *Message `json:"message"` // Deleted message
}

// MessageType return the string telegram-type of ChatEventMessageDeleted
func (chatEventMessageDeleted *ChatEventMessageDeleted) MessageType() string {
	return "chatEventMessageDeleted"
}

// NewChatEventMessageDeleted creates a new ChatEventMessageDeleted
//
// @param message Deleted message
func NewChatEventMessageDeleted(message *Message) *ChatEventMessageDeleted {
	chatEventMessageDeletedTemp := ChatEventMessageDeleted{
		tdCommon: tdCommon{Type: "chatEventMessageDeleted"},
		Message:  message,
	}

	return &chatEventMessageDeletedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventMessageDeleted *ChatEventMessageDeleted) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventMessageDeletedType
}

// ChatEventMessagePinned A message was pinned
type ChatEventMessagePinned struct {
	tdCommon
	Message *Message `json:"message"` // Pinned message
}

// MessageType return the string telegram-type of ChatEventMessagePinned
func (chatEventMessagePinned *ChatEventMessagePinned) MessageType() string {
	return "chatEventMessagePinned"
}

// NewChatEventMessagePinned creates a new ChatEventMessagePinned
//
// @param message Pinned message
func NewChatEventMessagePinned(message *Message) *ChatEventMessagePinned {
	chatEventMessagePinnedTemp := ChatEventMessagePinned{
		tdCommon: tdCommon{Type: "chatEventMessagePinned"},
		Message:  message,
	}

	return &chatEventMessagePinnedTemp
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

// NewChatEventMessageUnpinned creates a new ChatEventMessageUnpinned
//
func NewChatEventMessageUnpinned() *ChatEventMessageUnpinned {
	chatEventMessageUnpinnedTemp := ChatEventMessageUnpinned{
		tdCommon: tdCommon{Type: "chatEventMessageUnpinned"},
	}

	return &chatEventMessageUnpinnedTemp
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

// NewChatEventMemberJoined creates a new ChatEventMemberJoined
//
func NewChatEventMemberJoined() *ChatEventMemberJoined {
	chatEventMemberJoinedTemp := ChatEventMemberJoined{
		tdCommon: tdCommon{Type: "chatEventMemberJoined"},
	}

	return &chatEventMemberJoinedTemp
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

// NewChatEventMemberLeft creates a new ChatEventMemberLeft
//
func NewChatEventMemberLeft() *ChatEventMemberLeft {
	chatEventMemberLeftTemp := ChatEventMemberLeft{
		tdCommon: tdCommon{Type: "chatEventMemberLeft"},
	}

	return &chatEventMemberLeftTemp
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

// NewChatEventMemberInvited creates a new ChatEventMemberInvited
//
// @param userID New member user identifier
// @param status New member status
func NewChatEventMemberInvited(userID int32, status ChatMemberStatus) *ChatEventMemberInvited {
	chatEventMemberInvitedTemp := ChatEventMemberInvited{
		tdCommon: tdCommon{Type: "chatEventMemberInvited"},
		UserID:   userID,
		Status:   status,
	}

	return &chatEventMemberInvitedTemp
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

// NewChatEventMemberPromoted creates a new ChatEventMemberPromoted
//
// @param userID Chat member user identifier
// @param oldStatus Previous status of the chat member
// @param newStatus New status of the chat member
func NewChatEventMemberPromoted(userID int32, oldStatus ChatMemberStatus, newStatus ChatMemberStatus) *ChatEventMemberPromoted {
	chatEventMemberPromotedTemp := ChatEventMemberPromoted{
		tdCommon:  tdCommon{Type: "chatEventMemberPromoted"},
		UserID:    userID,
		OldStatus: oldStatus,
		NewStatus: newStatus,
	}

	return &chatEventMemberPromotedTemp
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

// NewChatEventMemberRestricted creates a new ChatEventMemberRestricted
//
// @param userID Chat member user identifier
// @param oldStatus Previous status of the chat member
// @param newStatus New status of the chat member
func NewChatEventMemberRestricted(userID int32, oldStatus ChatMemberStatus, newStatus ChatMemberStatus) *ChatEventMemberRestricted {
	chatEventMemberRestrictedTemp := ChatEventMemberRestricted{
		tdCommon:  tdCommon{Type: "chatEventMemberRestricted"},
		UserID:    userID,
		OldStatus: oldStatus,
		NewStatus: newStatus,
	}

	return &chatEventMemberRestrictedTemp
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

// NewChatEventTitleChanged creates a new ChatEventTitleChanged
//
// @param oldTitle Previous chat title
// @param newTitle New chat title
func NewChatEventTitleChanged(oldTitle string, newTitle string) *ChatEventTitleChanged {
	chatEventTitleChangedTemp := ChatEventTitleChanged{
		tdCommon: tdCommon{Type: "chatEventTitleChanged"},
		OldTitle: oldTitle,
		NewTitle: newTitle,
	}

	return &chatEventTitleChangedTemp
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

// NewChatEventDescriptionChanged creates a new ChatEventDescriptionChanged
//
// @param oldDescription Previous chat description
// @param newDescription New chat description
func NewChatEventDescriptionChanged(oldDescription string, newDescription string) *ChatEventDescriptionChanged {
	chatEventDescriptionChangedTemp := ChatEventDescriptionChanged{
		tdCommon:       tdCommon{Type: "chatEventDescriptionChanged"},
		OldDescription: oldDescription,
		NewDescription: newDescription,
	}

	return &chatEventDescriptionChangedTemp
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

// NewChatEventUsernameChanged creates a new ChatEventUsernameChanged
//
// @param oldUsername Previous chat username
// @param newUsername New chat username
func NewChatEventUsernameChanged(oldUsername string, newUsername string) *ChatEventUsernameChanged {
	chatEventUsernameChangedTemp := ChatEventUsernameChanged{
		tdCommon:    tdCommon{Type: "chatEventUsernameChanged"},
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}

	return &chatEventUsernameChangedTemp
}

// GetChatEventActionEnum return the enum type of this object
func (chatEventUsernameChanged *ChatEventUsernameChanged) GetChatEventActionEnum() ChatEventActionEnum {
	return ChatEventUsernameChangedType
}

// ChatEventPhotoChanged The chat photo was changed
type ChatEventPhotoChanged struct {
	tdCommon
	OldPhoto *ChatPhoto `json:"old_photo"` // Previous chat photo value; may be null
	NewPhoto *ChatPhoto `json:"new_photo"` // New chat photo value; may be null
}

// MessageType return the string telegram-type of ChatEventPhotoChanged
func (chatEventPhotoChanged *ChatEventPhotoChanged) MessageType() string {
	return "chatEventPhotoChanged"
}

// NewChatEventPhotoChanged creates a new ChatEventPhotoChanged
//
// @param oldPhoto Previous chat photo value; may be null
// @param newPhoto New chat photo value; may be null
func NewChatEventPhotoChanged(oldPhoto *ChatPhoto, newPhoto *ChatPhoto) *ChatEventPhotoChanged {
	chatEventPhotoChangedTemp := ChatEventPhotoChanged{
		tdCommon: tdCommon{Type: "chatEventPhotoChanged"},
		OldPhoto: oldPhoto,
		NewPhoto: newPhoto,
	}

	return &chatEventPhotoChangedTemp
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

// NewChatEventInvitesToggled creates a new ChatEventInvitesToggled
//
// @param anyoneCanInvite New value of anyone_can_invite
func NewChatEventInvitesToggled(anyoneCanInvite bool) *ChatEventInvitesToggled {
	chatEventInvitesToggledTemp := ChatEventInvitesToggled{
		tdCommon:        tdCommon{Type: "chatEventInvitesToggled"},
		AnyoneCanInvite: anyoneCanInvite,
	}

	return &chatEventInvitesToggledTemp
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

// NewChatEventSignMessagesToggled creates a new ChatEventSignMessagesToggled
//
// @param signMessages New value of sign_messages
func NewChatEventSignMessagesToggled(signMessages bool) *ChatEventSignMessagesToggled {
	chatEventSignMessagesToggledTemp := ChatEventSignMessagesToggled{
		tdCommon:     tdCommon{Type: "chatEventSignMessagesToggled"},
		SignMessages: signMessages,
	}

	return &chatEventSignMessagesToggledTemp
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

// NewChatEventStickerSetChanged creates a new ChatEventStickerSetChanged
//
// @param oldStickerSetID Previous identifier of the chat sticker set; 0 if none
// @param newStickerSetID New identifier of the chat sticker set; 0 if none
func NewChatEventStickerSetChanged(oldStickerSetID JSONInt64, newStickerSetID JSONInt64) *ChatEventStickerSetChanged {
	chatEventStickerSetChangedTemp := ChatEventStickerSetChanged{
		tdCommon:        tdCommon{Type: "chatEventStickerSetChanged"},
		OldStickerSetID: oldStickerSetID,
		NewStickerSetID: newStickerSetID,
	}

	return &chatEventStickerSetChangedTemp
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

// NewChatEventIsAllHistoryAvailableToggled creates a new ChatEventIsAllHistoryAvailableToggled
//
// @param isAllHistoryAvailable New value of is_all_history_available
func NewChatEventIsAllHistoryAvailableToggled(isAllHistoryAvailable bool) *ChatEventIsAllHistoryAvailableToggled {
	chatEventIsAllHistoryAvailableToggledTemp := ChatEventIsAllHistoryAvailableToggled{
		tdCommon:              tdCommon{Type: "chatEventIsAllHistoryAvailableToggled"},
		IsAllHistoryAvailable: isAllHistoryAvailable,
	}

	return &chatEventIsAllHistoryAvailableToggledTemp
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

// NewChatEvent creates a new ChatEvent
//
// @param iD Chat event identifier
// @param date Point in time (Unix timestamp) when the event happened
// @param userID Identifier of the user who performed the action that triggered the event
// @param action Action performed by the user
func NewChatEvent(iD JSONInt64, date int32, userID int32, action ChatEventAction) *ChatEvent {
	chatEventTemp := ChatEvent{
		tdCommon: tdCommon{Type: "chatEvent"},
		ID:       iD,
		Date:     date,
		UserID:   userID,
		Action:   action,
	}

	return &chatEventTemp
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

// NewChatEvents creates a new ChatEvents
//
// @param events List of events
func NewChatEvents(events []ChatEvent) *ChatEvents {
	chatEventsTemp := ChatEvents{
		tdCommon: tdCommon{Type: "chatEvents"},
		Events:   events,
	}

	return &chatEventsTemp
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

// NewChatEventLogFilters creates a new ChatEventLogFilters
//
// @param messageEdits True, if message edits should be returned
// @param messageDeletions True, if message deletions should be returned
// @param messagePins True, if pin/unpin events should be returned
// @param memberJoins True, if members joining events should be returned
// @param memberLeaves True, if members leaving events should be returned
// @param memberInvites True, if invited member events should be returned
// @param memberPromotions True, if member promotion/demotion events should be returned
// @param memberRestrictions True, if member restricted/unrestricted/banned/unbanned events should be returned
// @param infoChanges True, if changes in chat information should be returned
// @param settingChanges True, if changes in chat settings should be returned
func NewChatEventLogFilters(messageEdits bool, messageDeletions bool, messagePins bool, memberJoins bool, memberLeaves bool, memberInvites bool, memberPromotions bool, memberRestrictions bool, infoChanges bool, settingChanges bool) *ChatEventLogFilters {
	chatEventLogFiltersTemp := ChatEventLogFilters{
		tdCommon:           tdCommon{Type: "chatEventLogFilters"},
		MessageEdits:       messageEdits,
		MessageDeletions:   messageDeletions,
		MessagePins:        messagePins,
		MemberJoins:        memberJoins,
		MemberLeaves:       memberLeaves,
		MemberInvites:      memberInvites,
		MemberPromotions:   memberPromotions,
		MemberRestrictions: memberRestrictions,
		InfoChanges:        infoChanges,
		SettingChanges:     settingChanges,
	}

	return &chatEventLogFiltersTemp
}

// LanguagePackStringValueOrdinary An ordinary language pack string
type LanguagePackStringValueOrdinary struct {
	tdCommon
	Value string `json:"value"` // String value
}

// MessageType return the string telegram-type of LanguagePackStringValueOrdinary
func (languagePackStringValueOrdinary *LanguagePackStringValueOrdinary) MessageType() string {
	return "languagePackStringValueOrdinary"
}

// NewLanguagePackStringValueOrdinary creates a new LanguagePackStringValueOrdinary
//
// @param value String value
func NewLanguagePackStringValueOrdinary(value string) *LanguagePackStringValueOrdinary {
	languagePackStringValueOrdinaryTemp := LanguagePackStringValueOrdinary{
		tdCommon: tdCommon{Type: "languagePackStringValueOrdinary"},
		Value:    value,
	}

	return &languagePackStringValueOrdinaryTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValueOrdinary *LanguagePackStringValueOrdinary) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValueOrdinaryType
}

// LanguagePackStringValuePluralized A language pack string which has different forms based on the number of some object it mentions
type LanguagePackStringValuePluralized struct {
	tdCommon
	ZeroValue  string `json:"zero_value"`  // Value for zero objects
	OneValue   string `json:"one_value"`   // Value for one object
	TwoValue   string `json:"two_value"`   // Value for two objects
	FewValue   string `json:"few_value"`   // Value for few objects
	ManyValue  string `json:"many_value"`  // Value for many objects
	OtherValue string `json:"other_value"` // Default value
}

// MessageType return the string telegram-type of LanguagePackStringValuePluralized
func (languagePackStringValuePluralized *LanguagePackStringValuePluralized) MessageType() string {
	return "languagePackStringValuePluralized"
}

// NewLanguagePackStringValuePluralized creates a new LanguagePackStringValuePluralized
//
// @param zeroValue Value for zero objects
// @param oneValue Value for one object
// @param twoValue Value for two objects
// @param fewValue Value for few objects
// @param manyValue Value for many objects
// @param otherValue Default value
func NewLanguagePackStringValuePluralized(zeroValue string, oneValue string, twoValue string, fewValue string, manyValue string, otherValue string) *LanguagePackStringValuePluralized {
	languagePackStringValuePluralizedTemp := LanguagePackStringValuePluralized{
		tdCommon:   tdCommon{Type: "languagePackStringValuePluralized"},
		ZeroValue:  zeroValue,
		OneValue:   oneValue,
		TwoValue:   twoValue,
		FewValue:   fewValue,
		ManyValue:  manyValue,
		OtherValue: otherValue,
	}

	return &languagePackStringValuePluralizedTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValuePluralized *LanguagePackStringValuePluralized) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValuePluralizedType
}

// LanguagePackStringValueDeleted A deleted language pack string, the value should be taken from the built-in english language pack
type LanguagePackStringValueDeleted struct {
	tdCommon
}

// MessageType return the string telegram-type of LanguagePackStringValueDeleted
func (languagePackStringValueDeleted *LanguagePackStringValueDeleted) MessageType() string {
	return "languagePackStringValueDeleted"
}

// NewLanguagePackStringValueDeleted creates a new LanguagePackStringValueDeleted
//
func NewLanguagePackStringValueDeleted() *LanguagePackStringValueDeleted {
	languagePackStringValueDeletedTemp := LanguagePackStringValueDeleted{
		tdCommon: tdCommon{Type: "languagePackStringValueDeleted"},
	}

	return &languagePackStringValueDeletedTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValueDeleted *LanguagePackStringValueDeleted) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValueDeletedType
}

// LanguagePackString Represents one language pack string
type LanguagePackString struct {
	tdCommon
	Key   string                  `json:"key"`   // String key
	Value LanguagePackStringValue `json:"value"` // String value
}

// MessageType return the string telegram-type of LanguagePackString
func (languagePackString *LanguagePackString) MessageType() string {
	return "languagePackString"
}

// NewLanguagePackString creates a new LanguagePackString
//
// @param key String key
// @param value String value
func NewLanguagePackString(key string, value LanguagePackStringValue) *LanguagePackString {
	languagePackStringTemp := LanguagePackString{
		tdCommon: tdCommon{Type: "languagePackString"},
		Key:      key,
		Value:    value,
	}

	return &languagePackStringTemp
}

// UnmarshalJSON unmarshal to json
func (languagePackString *LanguagePackString) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Key string `json:"key"` // String key

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	languagePackString.tdCommon = tempObj.tdCommon
	languagePackString.Key = tempObj.Key

	fieldValue, _ := unmarshalLanguagePackStringValue(objMap["value"])
	languagePackString.Value = fieldValue

	return nil
}

// LanguagePackStrings Contains a list of language pack strings
type LanguagePackStrings struct {
	tdCommon
	Strings []LanguagePackString `json:"strings"` // A list of language pack strings
}

// MessageType return the string telegram-type of LanguagePackStrings
func (languagePackStrings *LanguagePackStrings) MessageType() string {
	return "languagePackStrings"
}

// NewLanguagePackStrings creates a new LanguagePackStrings
//
// @param strings A list of language pack strings
func NewLanguagePackStrings(strings []LanguagePackString) *LanguagePackStrings {
	languagePackStringsTemp := LanguagePackStrings{
		tdCommon: tdCommon{Type: "languagePackStrings"},
		Strings:  strings,
	}

	return &languagePackStringsTemp
}

// LanguagePackInfo Contains information about a language pack
type LanguagePackInfo struct {
	tdCommon
	ID               string `json:"id"`                 // Unique language pack identifier
	Name             string `json:"name"`               // Language name
	NativeName       string `json:"native_name"`        // Name of the language in that language
	LocalStringCount int32  `json:"local_string_count"` // Total number of non-deleted strings from the language pack available locally
}

// MessageType return the string telegram-type of LanguagePackInfo
func (languagePackInfo *LanguagePackInfo) MessageType() string {
	return "languagePackInfo"
}

// NewLanguagePackInfo creates a new LanguagePackInfo
//
// @param iD Unique language pack identifier
// @param name Language name
// @param nativeName Name of the language in that language
// @param localStringCount Total number of non-deleted strings from the language pack available locally
func NewLanguagePackInfo(iD string, name string, nativeName string, localStringCount int32) *LanguagePackInfo {
	languagePackInfoTemp := LanguagePackInfo{
		tdCommon:         tdCommon{Type: "languagePackInfo"},
		ID:               iD,
		Name:             name,
		NativeName:       nativeName,
		LocalStringCount: localStringCount,
	}

	return &languagePackInfoTemp
}

// LocalizationTargetInfo Contains information about the current localization target
type LocalizationTargetInfo struct {
	tdCommon
	LanguagePacks []LanguagePackInfo `json:"language_packs"` // List of available language packs for this application
}

// MessageType return the string telegram-type of LocalizationTargetInfo
func (localizationTargetInfo *LocalizationTargetInfo) MessageType() string {
	return "localizationTargetInfo"
}

// NewLocalizationTargetInfo creates a new LocalizationTargetInfo
//
// @param languagePacks List of available language packs for this application
func NewLocalizationTargetInfo(languagePacks []LanguagePackInfo) *LocalizationTargetInfo {
	localizationTargetInfoTemp := LocalizationTargetInfo{
		tdCommon:      tdCommon{Type: "localizationTargetInfo"},
		LanguagePacks: languagePacks,
	}

	return &localizationTargetInfoTemp
}

// DeviceTokenGoogleCloudMessaging A token for Google Cloud Messaging
type DeviceTokenGoogleCloudMessaging struct {
	tdCommon
	Token string `json:"token"` // Device registration token; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenGoogleCloudMessaging
func (deviceTokenGoogleCloudMessaging *DeviceTokenGoogleCloudMessaging) MessageType() string {
	return "deviceTokenGoogleCloudMessaging"
}

// NewDeviceTokenGoogleCloudMessaging creates a new DeviceTokenGoogleCloudMessaging
//
// @param token Device registration token; may be empty to de-register a device
func NewDeviceTokenGoogleCloudMessaging(token string) *DeviceTokenGoogleCloudMessaging {
	deviceTokenGoogleCloudMessagingTemp := DeviceTokenGoogleCloudMessaging{
		tdCommon: tdCommon{Type: "deviceTokenGoogleCloudMessaging"},
		Token:    token,
	}

	return &deviceTokenGoogleCloudMessagingTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenGoogleCloudMessaging *DeviceTokenGoogleCloudMessaging) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenGoogleCloudMessagingType
}

// DeviceTokenApplePush A token for Apple Push Notification service
type DeviceTokenApplePush struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token; may be empty to de-register a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePush
func (deviceTokenApplePush *DeviceTokenApplePush) MessageType() string {
	return "deviceTokenApplePush"
}

// NewDeviceTokenApplePush creates a new DeviceTokenApplePush
//
// @param deviceToken Device token; may be empty to de-register a device
// @param isAppSandbox True, if App Sandbox is enabled
func NewDeviceTokenApplePush(deviceToken string, isAppSandbox bool) *DeviceTokenApplePush {
	deviceTokenApplePushTemp := DeviceTokenApplePush{
		tdCommon:     tdCommon{Type: "deviceTokenApplePush"},
		DeviceToken:  deviceToken,
		IsAppSandbox: isAppSandbox,
	}

	return &deviceTokenApplePushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePush *DeviceTokenApplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushType
}

// DeviceTokenApplePushVoIP A token for Apple Push Notification service VoIP notifications
type DeviceTokenApplePushVoIP struct {
	tdCommon
	DeviceToken  string `json:"device_token"`   // Device token; may be empty to de-register a device
	IsAppSandbox bool   `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePushVoIP
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) MessageType() string {
	return "deviceTokenApplePushVoIP"
}

// NewDeviceTokenApplePushVoIP creates a new DeviceTokenApplePushVoIP
//
// @param deviceToken Device token; may be empty to de-register a device
// @param isAppSandbox True, if App Sandbox is enabled
func NewDeviceTokenApplePushVoIP(deviceToken string, isAppSandbox bool) *DeviceTokenApplePushVoIP {
	deviceTokenApplePushVoIPTemp := DeviceTokenApplePushVoIP{
		tdCommon:     tdCommon{Type: "deviceTokenApplePushVoIP"},
		DeviceToken:  deviceToken,
		IsAppSandbox: isAppSandbox,
	}

	return &deviceTokenApplePushVoIPTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenApplePushVoIPType
}

// DeviceTokenWindowsPush A token for Windows Push Notification Services
type DeviceTokenWindowsPush struct {
	tdCommon
	AccessToken string `json:"access_token"` // The access token that will be used to send notifications; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenWindowsPush
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) MessageType() string {
	return "deviceTokenWindowsPush"
}

// NewDeviceTokenWindowsPush creates a new DeviceTokenWindowsPush
//
// @param accessToken The access token that will be used to send notifications; may be empty to de-register a device
func NewDeviceTokenWindowsPush(accessToken string) *DeviceTokenWindowsPush {
	deviceTokenWindowsPushTemp := DeviceTokenWindowsPush{
		tdCommon:    tdCommon{Type: "deviceTokenWindowsPush"},
		AccessToken: accessToken,
	}

	return &deviceTokenWindowsPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWindowsPushType
}

// DeviceTokenMicrosoftPush A token for Microsoft Push Notification Service
type DeviceTokenMicrosoftPush struct {
	tdCommon
	ChannelURI string `json:"channel_uri"` // Push notification channel URI; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPush
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) MessageType() string {
	return "deviceTokenMicrosoftPush"
}

// NewDeviceTokenMicrosoftPush creates a new DeviceTokenMicrosoftPush
//
// @param channelURI Push notification channel URI; may be empty to de-register a device
func NewDeviceTokenMicrosoftPush(channelURI string) *DeviceTokenMicrosoftPush {
	deviceTokenMicrosoftPushTemp := DeviceTokenMicrosoftPush{
		tdCommon:   tdCommon{Type: "deviceTokenMicrosoftPush"},
		ChannelURI: channelURI,
	}

	return &deviceTokenMicrosoftPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushType
}

// DeviceTokenMicrosoftPushVoIP A token for Microsoft Push Notification Service VoIP channel
type DeviceTokenMicrosoftPushVoIP struct {
	tdCommon
	ChannelURI string `json:"channel_uri"` // Push notification channel URI; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPushVoIP
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) MessageType() string {
	return "deviceTokenMicrosoftPushVoIP"
}

// NewDeviceTokenMicrosoftPushVoIP creates a new DeviceTokenMicrosoftPushVoIP
//
// @param channelURI Push notification channel URI; may be empty to de-register a device
func NewDeviceTokenMicrosoftPushVoIP(channelURI string) *DeviceTokenMicrosoftPushVoIP {
	deviceTokenMicrosoftPushVoIPTemp := DeviceTokenMicrosoftPushVoIP{
		tdCommon:   tdCommon{Type: "deviceTokenMicrosoftPushVoIP"},
		ChannelURI: channelURI,
	}

	return &deviceTokenMicrosoftPushVoIPTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenMicrosoftPushVoIPType
}

// DeviceTokenWebPush A token for web Push API
type DeviceTokenWebPush struct {
	tdCommon
	Endpoint        string `json:"endpoint"`         // Absolute URL exposed by the push service where the application server can send push messages; may be empty to de-register a device
	P256dhBase64url string `json:"p256dh_base64url"` // Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
	AuthBase64url   string `json:"auth_base64url"`   // Base64url-encoded authentication secret
}

// MessageType return the string telegram-type of DeviceTokenWebPush
func (deviceTokenWebPush *DeviceTokenWebPush) MessageType() string {
	return "deviceTokenWebPush"
}

// NewDeviceTokenWebPush creates a new DeviceTokenWebPush
//
// @param endpoint Absolute URL exposed by the push service where the application server can send push messages; may be empty to de-register a device
// @param p256dhBase64url Base64url-encoded P-256 elliptic curve Diffie-Hellman public key
// @param authBase64url Base64url-encoded authentication secret
func NewDeviceTokenWebPush(endpoint string, p256dhBase64url string, authBase64url string) *DeviceTokenWebPush {
	deviceTokenWebPushTemp := DeviceTokenWebPush{
		tdCommon:        tdCommon{Type: "deviceTokenWebPush"},
		Endpoint:        endpoint,
		P256dhBase64url: p256dhBase64url,
		AuthBase64url:   authBase64url,
	}

	return &deviceTokenWebPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenWebPush *DeviceTokenWebPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenWebPushType
}

// DeviceTokenSimplePush A token for Simple Push API for Firefox OS
type DeviceTokenSimplePush struct {
	tdCommon
	Endpoint string `json:"endpoint"` // Absolute URL exposed by the push service where the application server can send push messages; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenSimplePush
func (deviceTokenSimplePush *DeviceTokenSimplePush) MessageType() string {
	return "deviceTokenSimplePush"
}

// NewDeviceTokenSimplePush creates a new DeviceTokenSimplePush
//
// @param endpoint Absolute URL exposed by the push service where the application server can send push messages; may be empty to de-register a device
func NewDeviceTokenSimplePush(endpoint string) *DeviceTokenSimplePush {
	deviceTokenSimplePushTemp := DeviceTokenSimplePush{
		tdCommon: tdCommon{Type: "deviceTokenSimplePush"},
		Endpoint: endpoint,
	}

	return &deviceTokenSimplePushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenSimplePush *DeviceTokenSimplePush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenSimplePushType
}

// DeviceTokenUbuntuPush A token for Ubuntu Push Client service
type DeviceTokenUbuntuPush struct {
	tdCommon
	Token string `json:"token"` // Token; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenUbuntuPush
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) MessageType() string {
	return "deviceTokenUbuntuPush"
}

// NewDeviceTokenUbuntuPush creates a new DeviceTokenUbuntuPush
//
// @param token Token; may be empty to de-register a device
func NewDeviceTokenUbuntuPush(token string) *DeviceTokenUbuntuPush {
	deviceTokenUbuntuPushTemp := DeviceTokenUbuntuPush{
		tdCommon: tdCommon{Type: "deviceTokenUbuntuPush"},
		Token:    token,
	}

	return &deviceTokenUbuntuPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenUbuntuPushType
}

// DeviceTokenBlackBerryPush A token for BlackBerry Push Service
type DeviceTokenBlackBerryPush struct {
	tdCommon
	Token string `json:"token"` // Token; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenBlackBerryPush
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) MessageType() string {
	return "deviceTokenBlackBerryPush"
}

// NewDeviceTokenBlackBerryPush creates a new DeviceTokenBlackBerryPush
//
// @param token Token; may be empty to de-register a device
func NewDeviceTokenBlackBerryPush(token string) *DeviceTokenBlackBerryPush {
	deviceTokenBlackBerryPushTemp := DeviceTokenBlackBerryPush{
		tdCommon: tdCommon{Type: "deviceTokenBlackBerryPush"},
		Token:    token,
	}

	return &deviceTokenBlackBerryPushTemp
}

// GetDeviceTokenEnum return the enum type of this object
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) GetDeviceTokenEnum() DeviceTokenEnum {
	return DeviceTokenBlackBerryPushType
}

// DeviceTokenTizenPush A token for Tizen Push Service
type DeviceTokenTizenPush struct {
	tdCommon
	RegID string `json:"reg_id"` // Push service registration identifier; may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenTizenPush
func (deviceTokenTizenPush *DeviceTokenTizenPush) MessageType() string {
	return "deviceTokenTizenPush"
}

// NewDeviceTokenTizenPush creates a new DeviceTokenTizenPush
//
// @param regID Push service registration identifier; may be empty to de-register a device
func NewDeviceTokenTizenPush(regID string) *DeviceTokenTizenPush {
	deviceTokenTizenPushTemp := DeviceTokenTizenPush{
		tdCommon: tdCommon{Type: "deviceTokenTizenPush"},
		RegID:    regID,
	}

	return &deviceTokenTizenPushTemp
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

// NewWallpaper creates a new Wallpaper
//
// @param iD Unique persistent wallpaper identifier
// @param sizes Available variants of the wallpaper in different sizes. These photos can only be downloaded; they can't be sent in a message
// @param color Main color of the wallpaper in RGB24 format; should be treated as background color if no photos are specified
func NewWallpaper(iD int32, sizes []PhotoSize, color int32) *Wallpaper {
	wallpaperTemp := Wallpaper{
		tdCommon: tdCommon{Type: "wallpaper"},
		ID:       iD,
		Sizes:    sizes,
		Color:    color,
	}

	return &wallpaperTemp
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

// NewWallpapers creates a new Wallpapers
//
// @param wallpapers A list of wallpapers
func NewWallpapers(wallpapers []Wallpaper) *Wallpapers {
	wallpapersTemp := Wallpapers{
		tdCommon:   tdCommon{Type: "wallpapers"},
		Wallpapers: wallpapers,
	}

	return &wallpapersTemp
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

// NewHashtags creates a new Hashtags
//
// @param hashtags A list of hashtags
func NewHashtags(hashtags []string) *Hashtags {
	hashtagsTemp := Hashtags{
		tdCommon: tdCommon{Type: "hashtags"},
		Hashtags: hashtags,
	}

	return &hashtagsTemp
}

// CheckChatUsernameResultOk The username can be set
type CheckChatUsernameResultOk struct {
	tdCommon
}

// MessageType return the string telegram-type of CheckChatUsernameResultOk
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) MessageType() string {
	return "checkChatUsernameResultOk"
}

// NewCheckChatUsernameResultOk creates a new CheckChatUsernameResultOk
//
func NewCheckChatUsernameResultOk() *CheckChatUsernameResultOk {
	checkChatUsernameResultOkTemp := CheckChatUsernameResultOk{
		tdCommon: tdCommon{Type: "checkChatUsernameResultOk"},
	}

	return &checkChatUsernameResultOkTemp
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

// NewCheckChatUsernameResultUsernameInvalid creates a new CheckChatUsernameResultUsernameInvalid
//
func NewCheckChatUsernameResultUsernameInvalid() *CheckChatUsernameResultUsernameInvalid {
	checkChatUsernameResultUsernameInvalidTemp := CheckChatUsernameResultUsernameInvalid{
		tdCommon: tdCommon{Type: "checkChatUsernameResultUsernameInvalid"},
	}

	return &checkChatUsernameResultUsernameInvalidTemp
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

// NewCheckChatUsernameResultUsernameOccupied creates a new CheckChatUsernameResultUsernameOccupied
//
func NewCheckChatUsernameResultUsernameOccupied() *CheckChatUsernameResultUsernameOccupied {
	checkChatUsernameResultUsernameOccupiedTemp := CheckChatUsernameResultUsernameOccupied{
		tdCommon: tdCommon{Type: "checkChatUsernameResultUsernameOccupied"},
	}

	return &checkChatUsernameResultUsernameOccupiedTemp
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

// NewCheckChatUsernameResultPublicChatsTooMuch creates a new CheckChatUsernameResultPublicChatsTooMuch
//
func NewCheckChatUsernameResultPublicChatsTooMuch() *CheckChatUsernameResultPublicChatsTooMuch {
	checkChatUsernameResultPublicChatsTooMuchTemp := CheckChatUsernameResultPublicChatsTooMuch{
		tdCommon: tdCommon{Type: "checkChatUsernameResultPublicChatsTooMuch"},
	}

	return &checkChatUsernameResultPublicChatsTooMuchTemp
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

// NewCheckChatUsernameResultPublicGroupsUnavailable creates a new CheckChatUsernameResultPublicGroupsUnavailable
//
func NewCheckChatUsernameResultPublicGroupsUnavailable() *CheckChatUsernameResultPublicGroupsUnavailable {
	checkChatUsernameResultPublicGroupsUnavailableTemp := CheckChatUsernameResultPublicGroupsUnavailable{
		tdCommon: tdCommon{Type: "checkChatUsernameResultPublicGroupsUnavailable"},
	}

	return &checkChatUsernameResultPublicGroupsUnavailableTemp
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

// NewOptionValueBoolean creates a new OptionValueBoolean
//
// @param value The value of the option
func NewOptionValueBoolean(value bool) *OptionValueBoolean {
	optionValueBooleanTemp := OptionValueBoolean{
		tdCommon: tdCommon{Type: "optionValueBoolean"},
		Value:    value,
	}

	return &optionValueBooleanTemp
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

// NewOptionValueEmpty creates a new OptionValueEmpty
//
func NewOptionValueEmpty() *OptionValueEmpty {
	optionValueEmptyTemp := OptionValueEmpty{
		tdCommon: tdCommon{Type: "optionValueEmpty"},
	}

	return &optionValueEmptyTemp
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

// NewOptionValueInteger creates a new OptionValueInteger
//
// @param value The value of the option
func NewOptionValueInteger(value int32) *OptionValueInteger {
	optionValueIntegerTemp := OptionValueInteger{
		tdCommon: tdCommon{Type: "optionValueInteger"},
		Value:    value,
	}

	return &optionValueIntegerTemp
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

// NewOptionValueString creates a new OptionValueString
//
// @param value The value of the option
func NewOptionValueString(value string) *OptionValueString {
	optionValueStringTemp := OptionValueString{
		tdCommon: tdCommon{Type: "optionValueString"},
		Value:    value,
	}

	return &optionValueStringTemp
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

// NewUserPrivacySettingRuleAllowAll creates a new UserPrivacySettingRuleAllowAll
//
func NewUserPrivacySettingRuleAllowAll() *UserPrivacySettingRuleAllowAll {
	userPrivacySettingRuleAllowAllTemp := UserPrivacySettingRuleAllowAll{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowAll"},
	}

	return &userPrivacySettingRuleAllowAllTemp
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

// NewUserPrivacySettingRuleAllowContacts creates a new UserPrivacySettingRuleAllowContacts
//
func NewUserPrivacySettingRuleAllowContacts() *UserPrivacySettingRuleAllowContacts {
	userPrivacySettingRuleAllowContactsTemp := UserPrivacySettingRuleAllowContacts{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowContacts"},
	}

	return &userPrivacySettingRuleAllowContactsTemp
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

// NewUserPrivacySettingRuleAllowUsers creates a new UserPrivacySettingRuleAllowUsers
//
// @param userIDs The user identifiers
func NewUserPrivacySettingRuleAllowUsers(userIDs []int32) *UserPrivacySettingRuleAllowUsers {
	userPrivacySettingRuleAllowUsersTemp := UserPrivacySettingRuleAllowUsers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleAllowUsers"},
		UserIDs:  userIDs,
	}

	return &userPrivacySettingRuleAllowUsersTemp
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

// NewUserPrivacySettingRuleRestrictAll creates a new UserPrivacySettingRuleRestrictAll
//
func NewUserPrivacySettingRuleRestrictAll() *UserPrivacySettingRuleRestrictAll {
	userPrivacySettingRuleRestrictAllTemp := UserPrivacySettingRuleRestrictAll{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictAll"},
	}

	return &userPrivacySettingRuleRestrictAllTemp
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

// NewUserPrivacySettingRuleRestrictContacts creates a new UserPrivacySettingRuleRestrictContacts
//
func NewUserPrivacySettingRuleRestrictContacts() *UserPrivacySettingRuleRestrictContacts {
	userPrivacySettingRuleRestrictContactsTemp := UserPrivacySettingRuleRestrictContacts{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictContacts"},
	}

	return &userPrivacySettingRuleRestrictContactsTemp
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

// NewUserPrivacySettingRuleRestrictUsers creates a new UserPrivacySettingRuleRestrictUsers
//
// @param userIDs The user identifiers
func NewUserPrivacySettingRuleRestrictUsers(userIDs []int32) *UserPrivacySettingRuleRestrictUsers {
	userPrivacySettingRuleRestrictUsersTemp := UserPrivacySettingRuleRestrictUsers{
		tdCommon: tdCommon{Type: "userPrivacySettingRuleRestrictUsers"},
		UserIDs:  userIDs,
	}

	return &userPrivacySettingRuleRestrictUsersTemp
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

// NewUserPrivacySettingRules creates a new UserPrivacySettingRules
//
// @param rules A list of rules
func NewUserPrivacySettingRules(rules []UserPrivacySettingRule) *UserPrivacySettingRules {
	userPrivacySettingRulesTemp := UserPrivacySettingRules{
		tdCommon: tdCommon{Type: "userPrivacySettingRules"},
		Rules:    rules,
	}

	return &userPrivacySettingRulesTemp
}

// UserPrivacySettingShowStatus A privacy setting for managing whether the user's online status is visible
type UserPrivacySettingShowStatus struct {
	tdCommon
}

// MessageType return the string telegram-type of UserPrivacySettingShowStatus
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) MessageType() string {
	return "userPrivacySettingShowStatus"
}

// NewUserPrivacySettingShowStatus creates a new UserPrivacySettingShowStatus
//
func NewUserPrivacySettingShowStatus() *UserPrivacySettingShowStatus {
	userPrivacySettingShowStatusTemp := UserPrivacySettingShowStatus{
		tdCommon: tdCommon{Type: "userPrivacySettingShowStatus"},
	}

	return &userPrivacySettingShowStatusTemp
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

// NewUserPrivacySettingAllowChatInvites creates a new UserPrivacySettingAllowChatInvites
//
func NewUserPrivacySettingAllowChatInvites() *UserPrivacySettingAllowChatInvites {
	userPrivacySettingAllowChatInvitesTemp := UserPrivacySettingAllowChatInvites{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowChatInvites"},
	}

	return &userPrivacySettingAllowChatInvitesTemp
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

// NewUserPrivacySettingAllowCalls creates a new UserPrivacySettingAllowCalls
//
func NewUserPrivacySettingAllowCalls() *UserPrivacySettingAllowCalls {
	userPrivacySettingAllowCallsTemp := UserPrivacySettingAllowCalls{
		tdCommon: tdCommon{Type: "userPrivacySettingAllowCalls"},
	}

	return &userPrivacySettingAllowCallsTemp
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

// NewAccountTTL creates a new AccountTTL
//
// @param days Number of days of inactivity before the account will be flagged for deletion; should range from 30-366 days
func NewAccountTTL(days int32) *AccountTTL {
	accountTTLTemp := AccountTTL{
		tdCommon: tdCommon{Type: "accountTtl"},
		Days:     days,
	}

	return &accountTTLTemp
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

// NewSession creates a new Session
//
// @param iD Session identifier
// @param isCurrent True, if this session is the current session
// @param aPIID Telegram API identifier, as provided by the application
// @param applicationName Name of the application, as provided by the application
// @param applicationVersion The version of the application, as provided by the application
// @param isOfficialApplication True, if the application is an official application or uses the api_id of an official application
// @param deviceModel Model of the device the application has been run or is running on, as provided by the application
// @param platform Operating system the application has been run or is running on, as provided by the application
// @param systemVersion Version of the operating system the application has been run or is running on, as provided by the application
// @param logInDate Point in time (Unix timestamp) when the user has logged in
// @param lastActiveDate Point in time (Unix timestamp) when the session was last used
// @param iP IP address from which the session was created, in human-readable format
// @param country A two-letter country code for the country from which the session was created, based on the IP address
// @param region Region code from which the session was created, based on the IP address
func NewSession(iD JSONInt64, isCurrent bool, aPIID int32, applicationName string, applicationVersion string, isOfficialApplication bool, deviceModel string, platform string, systemVersion string, logInDate int32, lastActiveDate int32, iP string, country string, region string) *Session {
	sessionTemp := Session{
		tdCommon:              tdCommon{Type: "session"},
		ID:                    iD,
		IsCurrent:             isCurrent,
		APIID:                 aPIID,
		ApplicationName:       applicationName,
		ApplicationVersion:    applicationVersion,
		IsOfficialApplication: isOfficialApplication,
		DeviceModel:           deviceModel,
		Platform:              platform,
		SystemVersion:         systemVersion,
		LogInDate:             logInDate,
		LastActiveDate:        lastActiveDate,
		IP:                    iP,
		Country:               country,
		Region:                region,
	}

	return &sessionTemp
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

// NewSessions creates a new Sessions
//
// @param sessions List of sessions
func NewSessions(sessions []Session) *Sessions {
	sessionsTemp := Sessions{
		tdCommon: tdCommon{Type: "sessions"},
		Sessions: sessions,
	}

	return &sessionsTemp
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

// NewConnectedWebsite creates a new ConnectedWebsite
//
// @param iD Website identifier
// @param domainName The domain name of the website
// @param botUserID User identifier of a bot linked with the website
// @param browser The version of a browser used to log in
// @param platform Operating system the browser is running on
// @param logInDate Point in time (Unix timestamp) when the user was logged in
// @param lastActiveDate Point in time (Unix timestamp) when obtained authorization was last used
// @param iP IP address from which the user was logged in, in human-readable format
// @param location Human-readable description of a country and a region, from which the user was logged in, based on the IP address
func NewConnectedWebsite(iD JSONInt64, domainName string, botUserID int32, browser string, platform string, logInDate int32, lastActiveDate int32, iP string, location string) *ConnectedWebsite {
	connectedWebsiteTemp := ConnectedWebsite{
		tdCommon:       tdCommon{Type: "connectedWebsite"},
		ID:             iD,
		DomainName:     domainName,
		BotUserID:      botUserID,
		Browser:        browser,
		Platform:       platform,
		LogInDate:      logInDate,
		LastActiveDate: lastActiveDate,
		IP:             iP,
		Location:       location,
	}

	return &connectedWebsiteTemp
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

// NewConnectedWebsites creates a new ConnectedWebsites
//
// @param websites List of connected websites
func NewConnectedWebsites(websites []ConnectedWebsite) *ConnectedWebsites {
	connectedWebsitesTemp := ConnectedWebsites{
		tdCommon: tdCommon{Type: "connectedWebsites"},
		Websites: websites,
	}

	return &connectedWebsitesTemp
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

// NewChatReportSpamState creates a new ChatReportSpamState
//
// @param canReportSpam True, if a prompt with the "Report spam" action should be shown to the user
func NewChatReportSpamState(canReportSpam bool) *ChatReportSpamState {
	chatReportSpamStateTemp := ChatReportSpamState{
		tdCommon:      tdCommon{Type: "chatReportSpamState"},
		CanReportSpam: canReportSpam,
	}

	return &chatReportSpamStateTemp
}

// ChatReportReasonSpam The chat contains spam messages
type ChatReportReasonSpam struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonSpam
func (chatReportReasonSpam *ChatReportReasonSpam) MessageType() string {
	return "chatReportReasonSpam"
}

// NewChatReportReasonSpam creates a new ChatReportReasonSpam
//
func NewChatReportReasonSpam() *ChatReportReasonSpam {
	chatReportReasonSpamTemp := ChatReportReasonSpam{
		tdCommon: tdCommon{Type: "chatReportReasonSpam"},
	}

	return &chatReportReasonSpamTemp
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

// NewChatReportReasonViolence creates a new ChatReportReasonViolence
//
func NewChatReportReasonViolence() *ChatReportReasonViolence {
	chatReportReasonViolenceTemp := ChatReportReasonViolence{
		tdCommon: tdCommon{Type: "chatReportReasonViolence"},
	}

	return &chatReportReasonViolenceTemp
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

// NewChatReportReasonPornography creates a new ChatReportReasonPornography
//
func NewChatReportReasonPornography() *ChatReportReasonPornography {
	chatReportReasonPornographyTemp := ChatReportReasonPornography{
		tdCommon: tdCommon{Type: "chatReportReasonPornography"},
	}

	return &chatReportReasonPornographyTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonPornography *ChatReportReasonPornography) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonPornographyType
}

// ChatReportReasonCopyright The chat contains copyrighted content
type ChatReportReasonCopyright struct {
	tdCommon
}

// MessageType return the string telegram-type of ChatReportReasonCopyright
func (chatReportReasonCopyright *ChatReportReasonCopyright) MessageType() string {
	return "chatReportReasonCopyright"
}

// NewChatReportReasonCopyright creates a new ChatReportReasonCopyright
//
func NewChatReportReasonCopyright() *ChatReportReasonCopyright {
	chatReportReasonCopyrightTemp := ChatReportReasonCopyright{
		tdCommon: tdCommon{Type: "chatReportReasonCopyright"},
	}

	return &chatReportReasonCopyrightTemp
}

// GetChatReportReasonEnum return the enum type of this object
func (chatReportReasonCopyright *ChatReportReasonCopyright) GetChatReportReasonEnum() ChatReportReasonEnum {
	return ChatReportReasonCopyrightType
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

// NewChatReportReasonCustom creates a new ChatReportReasonCustom
//
// @param text Report text
func NewChatReportReasonCustom(text string) *ChatReportReasonCustom {
	chatReportReasonCustomTemp := ChatReportReasonCustom{
		tdCommon: tdCommon{Type: "chatReportReasonCustom"},
		Text:     text,
	}

	return &chatReportReasonCustomTemp
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

// NewPublicMessageLink creates a new PublicMessageLink
//
// @param link Message link
// @param hTML HTML-code for embedding the message
func NewPublicMessageLink(link string, hTML string) *PublicMessageLink {
	publicMessageLinkTemp := PublicMessageLink{
		tdCommon: tdCommon{Type: "publicMessageLink"},
		Link:     link,
		HTML:     hTML,
	}

	return &publicMessageLinkTemp
}

// FileTypeNone The data is not a file
type FileTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeNone
func (fileTypeNone *FileTypeNone) MessageType() string {
	return "fileTypeNone"
}

// NewFileTypeNone creates a new FileTypeNone
//
func NewFileTypeNone() *FileTypeNone {
	fileTypeNoneTemp := FileTypeNone{
		tdCommon: tdCommon{Type: "fileTypeNone"},
	}

	return &fileTypeNoneTemp
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

// NewFileTypeAnimation creates a new FileTypeAnimation
//
func NewFileTypeAnimation() *FileTypeAnimation {
	fileTypeAnimationTemp := FileTypeAnimation{
		tdCommon: tdCommon{Type: "fileTypeAnimation"},
	}

	return &fileTypeAnimationTemp
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

// NewFileTypeAudio creates a new FileTypeAudio
//
func NewFileTypeAudio() *FileTypeAudio {
	fileTypeAudioTemp := FileTypeAudio{
		tdCommon: tdCommon{Type: "fileTypeAudio"},
	}

	return &fileTypeAudioTemp
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

// NewFileTypeDocument creates a new FileTypeDocument
//
func NewFileTypeDocument() *FileTypeDocument {
	fileTypeDocumentTemp := FileTypeDocument{
		tdCommon: tdCommon{Type: "fileTypeDocument"},
	}

	return &fileTypeDocumentTemp
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

// NewFileTypePhoto creates a new FileTypePhoto
//
func NewFileTypePhoto() *FileTypePhoto {
	fileTypePhotoTemp := FileTypePhoto{
		tdCommon: tdCommon{Type: "fileTypePhoto"},
	}

	return &fileTypePhotoTemp
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

// NewFileTypeProfilePhoto creates a new FileTypeProfilePhoto
//
func NewFileTypeProfilePhoto() *FileTypeProfilePhoto {
	fileTypeProfilePhotoTemp := FileTypeProfilePhoto{
		tdCommon: tdCommon{Type: "fileTypeProfilePhoto"},
	}

	return &fileTypeProfilePhotoTemp
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

// NewFileTypeSecret creates a new FileTypeSecret
//
func NewFileTypeSecret() *FileTypeSecret {
	fileTypeSecretTemp := FileTypeSecret{
		tdCommon: tdCommon{Type: "fileTypeSecret"},
	}

	return &fileTypeSecretTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecret *FileTypeSecret) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretType
}

// FileTypeSecretThumbnail The file is a thumbnail of a file from a secret chat
type FileTypeSecretThumbnail struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecretThumbnail
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) MessageType() string {
	return "fileTypeSecretThumbnail"
}

// NewFileTypeSecretThumbnail creates a new FileTypeSecretThumbnail
//
func NewFileTypeSecretThumbnail() *FileTypeSecretThumbnail {
	fileTypeSecretThumbnailTemp := FileTypeSecretThumbnail{
		tdCommon: tdCommon{Type: "fileTypeSecretThumbnail"},
	}

	return &fileTypeSecretThumbnailTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecretThumbnailType
}

// FileTypeSecure The file is a file from Secure storage used for storing Telegram Passport files
type FileTypeSecure struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSecure
func (fileTypeSecure *FileTypeSecure) MessageType() string {
	return "fileTypeSecure"
}

// NewFileTypeSecure creates a new FileTypeSecure
//
func NewFileTypeSecure() *FileTypeSecure {
	fileTypeSecureTemp := FileTypeSecure{
		tdCommon: tdCommon{Type: "fileTypeSecure"},
	}

	return &fileTypeSecureTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeSecure *FileTypeSecure) GetFileTypeEnum() FileTypeEnum {
	return FileTypeSecureType
}

// FileTypeSticker The file is a sticker
type FileTypeSticker struct {
	tdCommon
}

// MessageType return the string telegram-type of FileTypeSticker
func (fileTypeSticker *FileTypeSticker) MessageType() string {
	return "fileTypeSticker"
}

// NewFileTypeSticker creates a new FileTypeSticker
//
func NewFileTypeSticker() *FileTypeSticker {
	fileTypeStickerTemp := FileTypeSticker{
		tdCommon: tdCommon{Type: "fileTypeSticker"},
	}

	return &fileTypeStickerTemp
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

// NewFileTypeThumbnail creates a new FileTypeThumbnail
//
func NewFileTypeThumbnail() *FileTypeThumbnail {
	fileTypeThumbnailTemp := FileTypeThumbnail{
		tdCommon: tdCommon{Type: "fileTypeThumbnail"},
	}

	return &fileTypeThumbnailTemp
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

// NewFileTypeUnknown creates a new FileTypeUnknown
//
func NewFileTypeUnknown() *FileTypeUnknown {
	fileTypeUnknownTemp := FileTypeUnknown{
		tdCommon: tdCommon{Type: "fileTypeUnknown"},
	}

	return &fileTypeUnknownTemp
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

// NewFileTypeVideo creates a new FileTypeVideo
//
func NewFileTypeVideo() *FileTypeVideo {
	fileTypeVideoTemp := FileTypeVideo{
		tdCommon: tdCommon{Type: "fileTypeVideo"},
	}

	return &fileTypeVideoTemp
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

// NewFileTypeVideoNote creates a new FileTypeVideoNote
//
func NewFileTypeVideoNote() *FileTypeVideoNote {
	fileTypeVideoNoteTemp := FileTypeVideoNote{
		tdCommon: tdCommon{Type: "fileTypeVideoNote"},
	}

	return &fileTypeVideoNoteTemp
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

// NewFileTypeVoiceNote creates a new FileTypeVoiceNote
//
func NewFileTypeVoiceNote() *FileTypeVoiceNote {
	fileTypeVoiceNoteTemp := FileTypeVoiceNote{
		tdCommon: tdCommon{Type: "fileTypeVoiceNote"},
	}

	return &fileTypeVoiceNoteTemp
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

// NewFileTypeWallpaper creates a new FileTypeWallpaper
//
func NewFileTypeWallpaper() *FileTypeWallpaper {
	fileTypeWallpaperTemp := FileTypeWallpaper{
		tdCommon: tdCommon{Type: "fileTypeWallpaper"},
	}

	return &fileTypeWallpaperTemp
}

// GetFileTypeEnum return the enum type of this object
func (fileTypeWallpaper *FileTypeWallpaper) GetFileTypeEnum() FileTypeEnum {
	return FileTypeWallpaperType
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

// NewStorageStatisticsByFileType creates a new StorageStatisticsByFileType
//
// @param fileType File type
// @param size Total size of the files
// @param count Total number of files
func NewStorageStatisticsByFileType(fileType FileType, size int64, count int32) *StorageStatisticsByFileType {
	storageStatisticsByFileTypeTemp := StorageStatisticsByFileType{
		tdCommon: tdCommon{Type: "storageStatisticsByFileType"},
		FileType: fileType,
		Size:     size,
		Count:    count,
	}

	return &storageStatisticsByFileTypeTemp
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

// NewStorageStatisticsByChat creates a new StorageStatisticsByChat
//
// @param chatID Chat identifier; 0 if none
// @param size Total size of the files in the chat
// @param count Total number of files in the chat
// @param byFileType Statistics split by file types
func NewStorageStatisticsByChat(chatID int64, size int64, count int32, byFileType []StorageStatisticsByFileType) *StorageStatisticsByChat {
	storageStatisticsByChatTemp := StorageStatisticsByChat{
		tdCommon:   tdCommon{Type: "storageStatisticsByChat"},
		ChatID:     chatID,
		Size:       size,
		Count:      count,
		ByFileType: byFileType,
	}

	return &storageStatisticsByChatTemp
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

// NewStorageStatistics creates a new StorageStatistics
//
// @param size Total size of files
// @param count Total number of files
// @param byChat Statistics split by chats
func NewStorageStatistics(size int64, count int32, byChat []StorageStatisticsByChat) *StorageStatistics {
	storageStatisticsTemp := StorageStatistics{
		tdCommon: tdCommon{Type: "storageStatistics"},
		Size:     size,
		Count:    count,
		ByChat:   byChat,
	}

	return &storageStatisticsTemp
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

// NewStorageStatisticsFast creates a new StorageStatisticsFast
//
// @param filesSize Approximate total size of files
// @param fileCount Approximate number of files
// @param databaseSize Size of the database
func NewStorageStatisticsFast(filesSize int64, fileCount int32, databaseSize int64) *StorageStatisticsFast {
	storageStatisticsFastTemp := StorageStatisticsFast{
		tdCommon:     tdCommon{Type: "storageStatisticsFast"},
		FilesSize:    filesSize,
		FileCount:    fileCount,
		DatabaseSize: databaseSize,
	}

	return &storageStatisticsFastTemp
}

// NetworkTypeNone The network is not available
type NetworkTypeNone struct {
	tdCommon
}

// MessageType return the string telegram-type of NetworkTypeNone
func (networkTypeNone *NetworkTypeNone) MessageType() string {
	return "networkTypeNone"
}

// NewNetworkTypeNone creates a new NetworkTypeNone
//
func NewNetworkTypeNone() *NetworkTypeNone {
	networkTypeNoneTemp := NetworkTypeNone{
		tdCommon: tdCommon{Type: "networkTypeNone"},
	}

	return &networkTypeNoneTemp
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

// NewNetworkTypeMobile creates a new NetworkTypeMobile
//
func NewNetworkTypeMobile() *NetworkTypeMobile {
	networkTypeMobileTemp := NetworkTypeMobile{
		tdCommon: tdCommon{Type: "networkTypeMobile"},
	}

	return &networkTypeMobileTemp
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

// NewNetworkTypeMobileRoaming creates a new NetworkTypeMobileRoaming
//
func NewNetworkTypeMobileRoaming() *NetworkTypeMobileRoaming {
	networkTypeMobileRoamingTemp := NetworkTypeMobileRoaming{
		tdCommon: tdCommon{Type: "networkTypeMobileRoaming"},
	}

	return &networkTypeMobileRoamingTemp
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

// NewNetworkTypeWiFi creates a new NetworkTypeWiFi
//
func NewNetworkTypeWiFi() *NetworkTypeWiFi {
	networkTypeWiFiTemp := NetworkTypeWiFi{
		tdCommon: tdCommon{Type: "networkTypeWiFi"},
	}

	return &networkTypeWiFiTemp
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

// NewNetworkTypeOther creates a new NetworkTypeOther
//
func NewNetworkTypeOther() *NetworkTypeOther {
	networkTypeOtherTemp := NetworkTypeOther{
		tdCommon: tdCommon{Type: "networkTypeOther"},
	}

	return &networkTypeOtherTemp
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

// NewNetworkStatisticsEntryFile creates a new NetworkStatisticsEntryFile
//
// @param fileType Type of the file the data is part of
// @param networkType Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
// @param sentBytes Total number of bytes sent
// @param receivedBytes Total number of bytes received
func NewNetworkStatisticsEntryFile(fileType FileType, networkType NetworkType, sentBytes int64, receivedBytes int64) *NetworkStatisticsEntryFile {
	networkStatisticsEntryFileTemp := NetworkStatisticsEntryFile{
		tdCommon:      tdCommon{Type: "networkStatisticsEntryFile"},
		FileType:      fileType,
		NetworkType:   networkType,
		SentBytes:     sentBytes,
		ReceivedBytes: receivedBytes,
	}

	return &networkStatisticsEntryFileTemp
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

// NewNetworkStatisticsEntryCall creates a new NetworkStatisticsEntryCall
//
// @param networkType Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
// @param sentBytes Total number of bytes sent
// @param receivedBytes Total number of bytes received
// @param duration Total call duration, in seconds
func NewNetworkStatisticsEntryCall(networkType NetworkType, sentBytes int64, receivedBytes int64, duration float64) *NetworkStatisticsEntryCall {
	networkStatisticsEntryCallTemp := NetworkStatisticsEntryCall{
		tdCommon:      tdCommon{Type: "networkStatisticsEntryCall"},
		NetworkType:   networkType,
		SentBytes:     sentBytes,
		ReceivedBytes: receivedBytes,
		Duration:      duration,
	}

	return &networkStatisticsEntryCallTemp
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

// NewNetworkStatistics creates a new NetworkStatistics
//
// @param sinceDate Point in time (Unix timestamp) when the app began collecting statistics
// @param entries Network statistics entries
func NewNetworkStatistics(sinceDate int32, entries []NetworkStatisticsEntry) *NetworkStatistics {
	networkStatisticsTemp := NetworkStatistics{
		tdCommon:  tdCommon{Type: "networkStatistics"},
		SinceDate: sinceDate,
		Entries:   entries,
	}

	return &networkStatisticsTemp
}

// ConnectionStateWaitingForNetwork Currently waiting for the network to become available. Use SetNetworkType to change the available network type
type ConnectionStateWaitingForNetwork struct {
	tdCommon
}

// MessageType return the string telegram-type of ConnectionStateWaitingForNetwork
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) MessageType() string {
	return "connectionStateWaitingForNetwork"
}

// NewConnectionStateWaitingForNetwork creates a new ConnectionStateWaitingForNetwork
//
func NewConnectionStateWaitingForNetwork() *ConnectionStateWaitingForNetwork {
	connectionStateWaitingForNetworkTemp := ConnectionStateWaitingForNetwork{
		tdCommon: tdCommon{Type: "connectionStateWaitingForNetwork"},
	}

	return &connectionStateWaitingForNetworkTemp
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

// NewConnectionStateConnectingToProxy creates a new ConnectionStateConnectingToProxy
//
func NewConnectionStateConnectingToProxy() *ConnectionStateConnectingToProxy {
	connectionStateConnectingToProxyTemp := ConnectionStateConnectingToProxy{
		tdCommon: tdCommon{Type: "connectionStateConnectingToProxy"},
	}

	return &connectionStateConnectingToProxyTemp
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

// NewConnectionStateConnecting creates a new ConnectionStateConnecting
//
func NewConnectionStateConnecting() *ConnectionStateConnecting {
	connectionStateConnectingTemp := ConnectionStateConnecting{
		tdCommon: tdCommon{Type: "connectionStateConnecting"},
	}

	return &connectionStateConnectingTemp
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

// NewConnectionStateUpdating creates a new ConnectionStateUpdating
//
func NewConnectionStateUpdating() *ConnectionStateUpdating {
	connectionStateUpdatingTemp := ConnectionStateUpdating{
		tdCommon: tdCommon{Type: "connectionStateUpdating"},
	}

	return &connectionStateUpdatingTemp
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

// NewConnectionStateReady creates a new ConnectionStateReady
//
func NewConnectionStateReady() *ConnectionStateReady {
	connectionStateReadyTemp := ConnectionStateReady{
		tdCommon: tdCommon{Type: "connectionStateReady"},
	}

	return &connectionStateReadyTemp
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

// NewTopChatCategoryUsers creates a new TopChatCategoryUsers
//
func NewTopChatCategoryUsers() *TopChatCategoryUsers {
	topChatCategoryUsersTemp := TopChatCategoryUsers{
		tdCommon: tdCommon{Type: "topChatCategoryUsers"},
	}

	return &topChatCategoryUsersTemp
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

// NewTopChatCategoryBots creates a new TopChatCategoryBots
//
func NewTopChatCategoryBots() *TopChatCategoryBots {
	topChatCategoryBotsTemp := TopChatCategoryBots{
		tdCommon: tdCommon{Type: "topChatCategoryBots"},
	}

	return &topChatCategoryBotsTemp
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

// NewTopChatCategoryGroups creates a new TopChatCategoryGroups
//
func NewTopChatCategoryGroups() *TopChatCategoryGroups {
	topChatCategoryGroupsTemp := TopChatCategoryGroups{
		tdCommon: tdCommon{Type: "topChatCategoryGroups"},
	}

	return &topChatCategoryGroupsTemp
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

// NewTopChatCategoryChannels creates a new TopChatCategoryChannels
//
func NewTopChatCategoryChannels() *TopChatCategoryChannels {
	topChatCategoryChannelsTemp := TopChatCategoryChannels{
		tdCommon: tdCommon{Type: "topChatCategoryChannels"},
	}

	return &topChatCategoryChannelsTemp
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

// NewTopChatCategoryInlineBots creates a new TopChatCategoryInlineBots
//
func NewTopChatCategoryInlineBots() *TopChatCategoryInlineBots {
	topChatCategoryInlineBotsTemp := TopChatCategoryInlineBots{
		tdCommon: tdCommon{Type: "topChatCategoryInlineBots"},
	}

	return &topChatCategoryInlineBotsTemp
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

// NewTopChatCategoryCalls creates a new TopChatCategoryCalls
//
func NewTopChatCategoryCalls() *TopChatCategoryCalls {
	topChatCategoryCallsTemp := TopChatCategoryCalls{
		tdCommon: tdCommon{Type: "topChatCategoryCalls"},
	}

	return &topChatCategoryCallsTemp
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

// NewTMeURLTypeUser creates a new TMeURLTypeUser
//
// @param userID Identifier of the user
func NewTMeURLTypeUser(userID int32) *TMeURLTypeUser {
	tMeURLTypeUserTemp := TMeURLTypeUser{
		tdCommon: tdCommon{Type: "tMeUrlTypeUser"},
		UserID:   userID,
	}

	return &tMeURLTypeUserTemp
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

// NewTMeURLTypeSupergroup creates a new TMeURLTypeSupergroup
//
// @param supergroupID Identifier of the supergroup or channel
func NewTMeURLTypeSupergroup(supergroupID int64) *TMeURLTypeSupergroup {
	tMeURLTypeSupergroupTemp := TMeURLTypeSupergroup{
		tdCommon:     tdCommon{Type: "tMeUrlTypeSupergroup"},
		SupergroupID: supergroupID,
	}

	return &tMeURLTypeSupergroupTemp
}

// GetTMeURLTypeEnum return the enum type of this object
func (tMeURLTypeSupergroup *TMeURLTypeSupergroup) GetTMeURLTypeEnum() TMeURLTypeEnum {
	return TMeURLTypeSupergroupType
}

// TMeURLTypeChatInvite A chat invite link
type TMeURLTypeChatInvite struct {
	tdCommon
	Info *ChatInviteLinkInfo `json:"info"` // Chat invite link info
}

// MessageType return the string telegram-type of TMeURLTypeChatInvite
func (tMeURLTypeChatInvite *TMeURLTypeChatInvite) MessageType() string {
	return "tMeUrlTypeChatInvite"
}

// NewTMeURLTypeChatInvite creates a new TMeURLTypeChatInvite
//
// @param info Chat invite link info
func NewTMeURLTypeChatInvite(info *ChatInviteLinkInfo) *TMeURLTypeChatInvite {
	tMeURLTypeChatInviteTemp := TMeURLTypeChatInvite{
		tdCommon: tdCommon{Type: "tMeUrlTypeChatInvite"},
		Info:     info,
	}

	return &tMeURLTypeChatInviteTemp
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

// NewTMeURLTypeStickerSet creates a new TMeURLTypeStickerSet
//
// @param stickerSetID Identifier of the sticker set
func NewTMeURLTypeStickerSet(stickerSetID JSONInt64) *TMeURLTypeStickerSet {
	tMeURLTypeStickerSetTemp := TMeURLTypeStickerSet{
		tdCommon:     tdCommon{Type: "tMeUrlTypeStickerSet"},
		StickerSetID: stickerSetID,
	}

	return &tMeURLTypeStickerSetTemp
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

// NewTMeURL creates a new TMeURL
//
// @param uRL URL
// @param typeParam Type of the URL
func NewTMeURL(uRL string, typeParam TMeURLType) *TMeURL {
	tMeURLTemp := TMeURL{
		tdCommon: tdCommon{Type: "tMeUrl"},
		URL:      uRL,
		Type:     typeParam,
	}

	return &tMeURLTemp
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

// NewTMeURLs creates a new TMeURLs
//
// @param uRLs List of URLs
func NewTMeURLs(uRLs []TMeURL) *TMeURLs {
	tMeURLsTemp := TMeURLs{
		tdCommon: tdCommon{Type: "tMeUrls"},
		URLs:     uRLs,
	}

	return &tMeURLsTemp
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

// NewCount creates a new Count
//
// @param count Count
func NewCount(count int32) *Count {
	countTemp := Count{
		tdCommon: tdCommon{Type: "count"},
		Count:    count,
	}

	return &countTemp
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

// NewText creates a new Text
//
// @param text Text
func NewText(text string) *Text {
	textTemp := Text{
		tdCommon: tdCommon{Type: "text"},
		Text:     text,
	}

	return &textTemp
}

// Seconds Contains a value representing a number of seconds
type Seconds struct {
	tdCommon
	Seconds float64 `json:"seconds"` // Number of seconds
}

// MessageType return the string telegram-type of Seconds
func (seconds *Seconds) MessageType() string {
	return "seconds"
}

// NewSeconds creates a new Seconds
//
// @param seconds Number of seconds
func NewSeconds(seconds float64) *Seconds {
	secondsTemp := Seconds{
		tdCommon: tdCommon{Type: "seconds"},
		Seconds:  seconds,
	}

	return &secondsTemp
}

// DeepLinkInfo Contains information about a tg:// deep link
type DeepLinkInfo struct {
	tdCommon
	Text                  *FormattedText `json:"text"`                    // Text to be shown to the user
	NeedUpdateApplication bool           `json:"need_update_application"` // True, if user should be asked to update the application
}

// MessageType return the string telegram-type of DeepLinkInfo
func (deepLinkInfo *DeepLinkInfo) MessageType() string {
	return "deepLinkInfo"
}

// NewDeepLinkInfo creates a new DeepLinkInfo
//
// @param text Text to be shown to the user
// @param needUpdateApplication True, if user should be asked to update the application
func NewDeepLinkInfo(text *FormattedText, needUpdateApplication bool) *DeepLinkInfo {
	deepLinkInfoTemp := DeepLinkInfo{
		tdCommon:              tdCommon{Type: "deepLinkInfo"},
		Text:                  text,
		NeedUpdateApplication: needUpdateApplication,
	}

	return &deepLinkInfoTemp
}

// TextParseModeMarkdown The text should be parsed in markdown-style
type TextParseModeMarkdown struct {
	tdCommon
}

// MessageType return the string telegram-type of TextParseModeMarkdown
func (textParseModeMarkdown *TextParseModeMarkdown) MessageType() string {
	return "textParseModeMarkdown"
}

// NewTextParseModeMarkdown creates a new TextParseModeMarkdown
//
func NewTextParseModeMarkdown() *TextParseModeMarkdown {
	textParseModeMarkdownTemp := TextParseModeMarkdown{
		tdCommon: tdCommon{Type: "textParseModeMarkdown"},
	}

	return &textParseModeMarkdownTemp
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

// NewTextParseModeHTML creates a new TextParseModeHTML
//
func NewTextParseModeHTML() *TextParseModeHTML {
	textParseModeHTMLTemp := TextParseModeHTML{
		tdCommon: tdCommon{Type: "textParseModeHTML"},
	}

	return &textParseModeHTMLTemp
}

// GetTextParseModeEnum return the enum type of this object
func (textParseModeHTML *TextParseModeHTML) GetTextParseModeEnum() TextParseModeEnum {
	return TextParseModeHTMLType
}

// ProxyTypeSocks5 A SOCKS5 proxy server
type ProxyTypeSocks5 struct {
	tdCommon
	Username string `json:"username"` // Username for logging in; may be empty
	Password string `json:"password"` // Password for logging in; may be empty
}

// MessageType return the string telegram-type of ProxyTypeSocks5
func (proxyTypeSocks5 *ProxyTypeSocks5) MessageType() string {
	return "proxyTypeSocks5"
}

// NewProxyTypeSocks5 creates a new ProxyTypeSocks5
//
// @param username Username for logging in; may be empty
// @param password Password for logging in; may be empty
func NewProxyTypeSocks5(username string, password string) *ProxyTypeSocks5 {
	proxyTypeSocks5Temp := ProxyTypeSocks5{
		tdCommon: tdCommon{Type: "proxyTypeSocks5"},
		Username: username,
		Password: password,
	}

	return &proxyTypeSocks5Temp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeSocks5 *ProxyTypeSocks5) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeSocks5Type
}

// ProxyTypeHttp A HTTP transparent proxy server
type ProxyTypeHttp struct {
	tdCommon
	Username string `json:"username"`  // Username for logging in; may be empty
	Password string `json:"password"`  // Password for logging in; may be empty
	HttpOnly bool   `json:"http_only"` // Pass true, if the proxy supports only HTTP requests and doesn't support transparent TCP connections via HTTP CONNECT method
}

// MessageType return the string telegram-type of ProxyTypeHttp
func (proxyTypeHttp *ProxyTypeHttp) MessageType() string {
	return "proxyTypeHttp"
}

// NewProxyTypeHttp creates a new ProxyTypeHttp
//
// @param username Username for logging in; may be empty
// @param password Password for logging in; may be empty
// @param httpOnly Pass true, if the proxy supports only HTTP requests and doesn't support transparent TCP connections via HTTP CONNECT method
func NewProxyTypeHttp(username string, password string, httpOnly bool) *ProxyTypeHttp {
	proxyTypeHttpTemp := ProxyTypeHttp{
		tdCommon: tdCommon{Type: "proxyTypeHttp"},
		Username: username,
		Password: password,
		HttpOnly: httpOnly,
	}

	return &proxyTypeHttpTemp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeHttp *ProxyTypeHttp) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeHttpType
}

// ProxyTypeMtproto An MTProto proxy server
type ProxyTypeMtproto struct {
	tdCommon
	Secret string `json:"secret"` // The proxy's secret in hexadecimal encoding
}

// MessageType return the string telegram-type of ProxyTypeMtproto
func (proxyTypeMtproto *ProxyTypeMtproto) MessageType() string {
	return "proxyTypeMtproto"
}

// NewProxyTypeMtproto creates a new ProxyTypeMtproto
//
// @param secret The proxy's secret in hexadecimal encoding
func NewProxyTypeMtproto(secret string) *ProxyTypeMtproto {
	proxyTypeMtprotoTemp := ProxyTypeMtproto{
		tdCommon: tdCommon{Type: "proxyTypeMtproto"},
		Secret:   secret,
	}

	return &proxyTypeMtprotoTemp
}

// GetProxyTypeEnum return the enum type of this object
func (proxyTypeMtproto *ProxyTypeMtproto) GetProxyTypeEnum() ProxyTypeEnum {
	return ProxyTypeMtprotoType
}

// Proxy Contains information about a proxy server
type Proxy struct {
	tdCommon
	ID           int32     `json:"id"`             // Unique identifier of the proxy
	Server       string    `json:"server"`         // Proxy server IP address
	Port         int32     `json:"port"`           // Proxy server port
	LastUsedDate int32     `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
	IsEnabled    bool      `json:"is_enabled"`     // True, if the proxy is enabled now
	Type         ProxyType `json:"type"`           // Type of the proxy
}

// MessageType return the string telegram-type of Proxy
func (proxy *Proxy) MessageType() string {
	return "proxy"
}

// NewProxy creates a new Proxy
//
// @param iD Unique identifier of the proxy
// @param server Proxy server IP address
// @param port Proxy server port
// @param lastUsedDate Point in time (Unix timestamp) when the proxy was last used; 0 if never
// @param isEnabled True, if the proxy is enabled now
// @param typeParam Type of the proxy
func NewProxy(iD int32, server string, port int32, lastUsedDate int32, isEnabled bool, typeParam ProxyType) *Proxy {
	proxyTemp := Proxy{
		tdCommon:     tdCommon{Type: "proxy"},
		ID:           iD,
		Server:       server,
		Port:         port,
		LastUsedDate: lastUsedDate,
		IsEnabled:    isEnabled,
		Type:         typeParam,
	}

	return &proxyTemp
}

// UnmarshalJSON unmarshal to json
func (proxy *Proxy) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID           int32  `json:"id"`             // Unique identifier of the proxy
		Server       string `json:"server"`         // Proxy server IP address
		Port         int32  `json:"port"`           // Proxy server port
		LastUsedDate int32  `json:"last_used_date"` // Point in time (Unix timestamp) when the proxy was last used; 0 if never
		IsEnabled    bool   `json:"is_enabled"`     // True, if the proxy is enabled now

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	proxy.tdCommon = tempObj.tdCommon
	proxy.ID = tempObj.ID
	proxy.Server = tempObj.Server
	proxy.Port = tempObj.Port
	proxy.LastUsedDate = tempObj.LastUsedDate
	proxy.IsEnabled = tempObj.IsEnabled

	fieldType, _ := unmarshalProxyType(objMap["type"])
	proxy.Type = fieldType

	return nil
}

// Proxies Represents a list of proxy servers
type Proxies struct {
	tdCommon
	Proxies []Proxy `json:"proxies"` // List of proxy servers
}

// MessageType return the string telegram-type of Proxies
func (proxies *Proxies) MessageType() string {
	return "proxies"
}

// NewProxies creates a new Proxies
//
// @param proxies List of proxy servers
func NewProxies(proxies []Proxy) *Proxies {
	proxiesTemp := Proxies{
		tdCommon: tdCommon{Type: "proxies"},
		Proxies:  proxies,
	}

	return &proxiesTemp
}

// InputSticker Describes a sticker that should be added to a sticker set
type InputSticker struct {
	tdCommon
	PngSticker   InputFile     `json:"png_sticker"`   // PNG image with the sticker; must be up to 512 kB in size and fit in a 512x512 square
	Emojis       string        `json:"emojis"`        // Emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position"` // For masks, position where the mask should be placed; may be null
}

// MessageType return the string telegram-type of InputSticker
func (inputSticker *InputSticker) MessageType() string {
	return "inputSticker"
}

// NewInputSticker creates a new InputSticker
//
// @param pngSticker PNG image with the sticker; must be up to 512 kB in size and fit in a 512x512 square
// @param emojis Emoji corresponding to the sticker
// @param maskPosition For masks, position where the mask should be placed; may be null
func NewInputSticker(pngSticker InputFile, emojis string, maskPosition *MaskPosition) *InputSticker {
	inputStickerTemp := InputSticker{
		tdCommon:     tdCommon{Type: "inputSticker"},
		PngSticker:   pngSticker,
		Emojis:       emojis,
		MaskPosition: maskPosition,
	}

	return &inputStickerTemp
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
		Emojis       string        `json:"emojis"`        // Emoji corresponding to the sticker
		MaskPosition *MaskPosition `json:"mask_position"` // For masks, position where the mask should be placed; may be null
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

// NewUpdateAuthorizationState creates a new UpdateAuthorizationState
//
// @param authorizationState New authorization state
func NewUpdateAuthorizationState(authorizationState AuthorizationState) *UpdateAuthorizationState {
	updateAuthorizationStateTemp := UpdateAuthorizationState{
		tdCommon:           tdCommon{Type: "updateAuthorizationState"},
		AuthorizationState: authorizationState,
	}

	return &updateAuthorizationStateTemp
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
	Message             *Message `json:"message"`              // The new message
	DisableNotification bool     `json:"disable_notification"` // True, if this message must not generate a notification
	ContainsMention     bool     `json:"contains_mention"`     // True, if the message contains a mention of the current user
}

// MessageType return the string telegram-type of UpdateNewMessage
func (updateNewMessage *UpdateNewMessage) MessageType() string {
	return "updateNewMessage"
}

// NewUpdateNewMessage creates a new UpdateNewMessage
//
// @param message The new message
// @param disableNotification True, if this message must not generate a notification
// @param containsMention True, if the message contains a mention of the current user
func NewUpdateNewMessage(message *Message, disableNotification bool, containsMention bool) *UpdateNewMessage {
	updateNewMessageTemp := UpdateNewMessage{
		tdCommon:            tdCommon{Type: "updateNewMessage"},
		Message:             message,
		DisableNotification: disableNotification,
		ContainsMention:     containsMention,
	}

	return &updateNewMessageTemp
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

// NewUpdateMessageSendAcknowledged creates a new UpdateMessageSendAcknowledged
//
// @param chatID The chat identifier of the sent message
// @param messageID A temporary message identifier
func NewUpdateMessageSendAcknowledged(chatID int64, messageID int64) *UpdateMessageSendAcknowledged {
	updateMessageSendAcknowledgedTemp := UpdateMessageSendAcknowledged{
		tdCommon:  tdCommon{Type: "updateMessageSendAcknowledged"},
		ChatID:    chatID,
		MessageID: messageID,
	}

	return &updateMessageSendAcknowledgedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendAcknowledgedType
}

// UpdateMessageSendSucceeded A message has been successfully sent
type UpdateMessageSendSucceeded struct {
	tdCommon
	Message      *Message `json:"message"`        // Information about the sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change
	OldMessageID int64    `json:"old_message_id"` // The previous temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendSucceeded
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) MessageType() string {
	return "updateMessageSendSucceeded"
}

// NewUpdateMessageSendSucceeded creates a new UpdateMessageSendSucceeded
//
// @param message Information about the sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change
// @param oldMessageID The previous temporary message identifier
func NewUpdateMessageSendSucceeded(message *Message, oldMessageID int64) *UpdateMessageSendSucceeded {
	updateMessageSendSucceededTemp := UpdateMessageSendSucceeded{
		tdCommon:     tdCommon{Type: "updateMessageSendSucceeded"},
		Message:      message,
		OldMessageID: oldMessageID,
	}

	return &updateMessageSendSucceededTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) GetUpdateEnum() UpdateEnum {
	return UpdateMessageSendSucceededType
}

// UpdateMessageSendFailed A message failed to send. Be aware that some messages being sent can be irrecoverably deleted, in which case updateDeleteMessages will be received instead of this update
type UpdateMessageSendFailed struct {
	tdCommon
	Message      *Message `json:"message"`        // Contains information about the message that failed to send
	OldMessageID int64    `json:"old_message_id"` // The previous temporary message identifier
	ErrorCode    int32    `json:"error_code"`     // An error code
	ErrorMessage string   `json:"error_message"`  // Error message
}

// MessageType return the string telegram-type of UpdateMessageSendFailed
func (updateMessageSendFailed *UpdateMessageSendFailed) MessageType() string {
	return "updateMessageSendFailed"
}

// NewUpdateMessageSendFailed creates a new UpdateMessageSendFailed
//
// @param message Contains information about the message that failed to send
// @param oldMessageID The previous temporary message identifier
// @param errorCode An error code
// @param errorMessage Error message
func NewUpdateMessageSendFailed(message *Message, oldMessageID int64, errorCode int32, errorMessage string) *UpdateMessageSendFailed {
	updateMessageSendFailedTemp := UpdateMessageSendFailed{
		tdCommon:     tdCommon{Type: "updateMessageSendFailed"},
		Message:      message,
		OldMessageID: oldMessageID,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}

	return &updateMessageSendFailedTemp
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

// NewUpdateMessageContent creates a new UpdateMessageContent
//
// @param chatID Chat identifier
// @param messageID Message identifier
// @param newContent New message content
func NewUpdateMessageContent(chatID int64, messageID int64, newContent MessageContent) *UpdateMessageContent {
	updateMessageContentTemp := UpdateMessageContent{
		tdCommon:   tdCommon{Type: "updateMessageContent"},
		ChatID:     chatID,
		MessageID:  messageID,
		NewContent: newContent,
	}

	return &updateMessageContentTemp
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

// NewUpdateMessageEdited creates a new UpdateMessageEdited
//
// @param chatID Chat identifier
// @param messageID Message identifier
// @param editDate Point in time (Unix timestamp) when the message was edited
// @param replyMarkup New message reply markup; may be null
func NewUpdateMessageEdited(chatID int64, messageID int64, editDate int32, replyMarkup ReplyMarkup) *UpdateMessageEdited {
	updateMessageEditedTemp := UpdateMessageEdited{
		tdCommon:    tdCommon{Type: "updateMessageEdited"},
		ChatID:      chatID,
		MessageID:   messageID,
		EditDate:    editDate,
		ReplyMarkup: replyMarkup,
	}

	return &updateMessageEditedTemp
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

// NewUpdateMessageViews creates a new UpdateMessageViews
//
// @param chatID Chat identifier
// @param messageID Message identifier
// @param views New value of the view count
func NewUpdateMessageViews(chatID int64, messageID int64, views int32) *UpdateMessageViews {
	updateMessageViewsTemp := UpdateMessageViews{
		tdCommon:  tdCommon{Type: "updateMessageViews"},
		ChatID:    chatID,
		MessageID: messageID,
		Views:     views,
	}

	return &updateMessageViewsTemp
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

// NewUpdateMessageContentOpened creates a new UpdateMessageContentOpened
//
// @param chatID Chat identifier
// @param messageID Message identifier
func NewUpdateMessageContentOpened(chatID int64, messageID int64) *UpdateMessageContentOpened {
	updateMessageContentOpenedTemp := UpdateMessageContentOpened{
		tdCommon:  tdCommon{Type: "updateMessageContentOpened"},
		ChatID:    chatID,
		MessageID: messageID,
	}

	return &updateMessageContentOpenedTemp
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

// NewUpdateMessageMentionRead creates a new UpdateMessageMentionRead
//
// @param chatID Chat identifier
// @param messageID Message identifier
// @param unreadMentionCount The new number of unread mention messages left in the chat
func NewUpdateMessageMentionRead(chatID int64, messageID int64, unreadMentionCount int32) *UpdateMessageMentionRead {
	updateMessageMentionReadTemp := UpdateMessageMentionRead{
		tdCommon:           tdCommon{Type: "updateMessageMentionRead"},
		ChatID:             chatID,
		MessageID:          messageID,
		UnreadMentionCount: unreadMentionCount,
	}

	return &updateMessageMentionReadTemp
}

// GetUpdateEnum return the enum type of this object
func (updateMessageMentionRead *UpdateMessageMentionRead) GetUpdateEnum() UpdateEnum {
	return UpdateMessageMentionReadType
}

// UpdateNewChat A new chat has been loaded/created. This update is guaranteed to come before the chat identifier is returned to the client. The chat field changes will be reported through separate updates
type UpdateNewChat struct {
	tdCommon
	Chat *Chat `json:"chat"` // The chat
}

// MessageType return the string telegram-type of UpdateNewChat
func (updateNewChat *UpdateNewChat) MessageType() string {
	return "updateNewChat"
}

// NewUpdateNewChat creates a new UpdateNewChat
//
// @param chat The chat
func NewUpdateNewChat(chat *Chat) *UpdateNewChat {
	updateNewChatTemp := UpdateNewChat{
		tdCommon: tdCommon{Type: "updateNewChat"},
		Chat:     chat,
	}

	return &updateNewChatTemp
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

// NewUpdateChatTitle creates a new UpdateChatTitle
//
// @param chatID Chat identifier
// @param title The new chat title
func NewUpdateChatTitle(chatID int64, title string) *UpdateChatTitle {
	updateChatTitleTemp := UpdateChatTitle{
		tdCommon: tdCommon{Type: "updateChatTitle"},
		ChatID:   chatID,
		Title:    title,
	}

	return &updateChatTitleTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatTitle *UpdateChatTitle) GetUpdateEnum() UpdateEnum {
	return UpdateChatTitleType
}

// UpdateChatPhoto A chat photo was changed
type UpdateChatPhoto struct {
	tdCommon
	ChatID int64      `json:"chat_id"` // Chat identifier
	Photo  *ChatPhoto `json:"photo"`   // The new chat photo; may be null
}

// MessageType return the string telegram-type of UpdateChatPhoto
func (updateChatPhoto *UpdateChatPhoto) MessageType() string {
	return "updateChatPhoto"
}

// NewUpdateChatPhoto creates a new UpdateChatPhoto
//
// @param chatID Chat identifier
// @param photo The new chat photo; may be null
func NewUpdateChatPhoto(chatID int64, photo *ChatPhoto) *UpdateChatPhoto {
	updateChatPhotoTemp := UpdateChatPhoto{
		tdCommon: tdCommon{Type: "updateChatPhoto"},
		ChatID:   chatID,
		Photo:    photo,
	}

	return &updateChatPhotoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatPhoto *UpdateChatPhoto) GetUpdateEnum() UpdateEnum {
	return UpdateChatPhotoType
}

// UpdateChatLastMessage The last message of a chat was changed. If last_message is null then the last message in the chat became unknown. Some new unknown messages might be added to the chat in this case
type UpdateChatLastMessage struct {
	tdCommon
	ChatID      int64     `json:"chat_id"`      // Chat identifier
	LastMessage *Message  `json:"last_message"` // The new last message in the chat; may be null
	Order       JSONInt64 `json:"order"`        // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatLastMessage
func (updateChatLastMessage *UpdateChatLastMessage) MessageType() string {
	return "updateChatLastMessage"
}

// NewUpdateChatLastMessage creates a new UpdateChatLastMessage
//
// @param chatID Chat identifier
// @param lastMessage The new last message in the chat; may be null
// @param order New value of the chat order
func NewUpdateChatLastMessage(chatID int64, lastMessage *Message, order JSONInt64) *UpdateChatLastMessage {
	updateChatLastMessageTemp := UpdateChatLastMessage{
		tdCommon:    tdCommon{Type: "updateChatLastMessage"},
		ChatID:      chatID,
		LastMessage: lastMessage,
		Order:       order,
	}

	return &updateChatLastMessageTemp
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

// NewUpdateChatOrder creates a new UpdateChatOrder
//
// @param chatID Chat identifier
// @param order New value of the order
func NewUpdateChatOrder(chatID int64, order JSONInt64) *UpdateChatOrder {
	updateChatOrderTemp := UpdateChatOrder{
		tdCommon: tdCommon{Type: "updateChatOrder"},
		ChatID:   chatID,
		Order:    order,
	}

	return &updateChatOrderTemp
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

// NewUpdateChatIsPinned creates a new UpdateChatIsPinned
//
// @param chatID Chat identifier
// @param isPinned New value of is_pinned
// @param order New value of the chat order
func NewUpdateChatIsPinned(chatID int64, isPinned bool, order JSONInt64) *UpdateChatIsPinned {
	updateChatIsPinnedTemp := UpdateChatIsPinned{
		tdCommon: tdCommon{Type: "updateChatIsPinned"},
		ChatID:   chatID,
		IsPinned: isPinned,
		Order:    order,
	}

	return &updateChatIsPinnedTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsPinned *UpdateChatIsPinned) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsPinnedType
}

// UpdateChatIsMarkedAsUnread A chat was marked as unread or was read
type UpdateChatIsMarkedAsUnread struct {
	tdCommon
	ChatID           int64 `json:"chat_id"`             // Chat identifier
	IsMarkedAsUnread bool  `json:"is_marked_as_unread"` // New value of is_marked_as_unread
}

// MessageType return the string telegram-type of UpdateChatIsMarkedAsUnread
func (updateChatIsMarkedAsUnread *UpdateChatIsMarkedAsUnread) MessageType() string {
	return "updateChatIsMarkedAsUnread"
}

// NewUpdateChatIsMarkedAsUnread creates a new UpdateChatIsMarkedAsUnread
//
// @param chatID Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func NewUpdateChatIsMarkedAsUnread(chatID int64, isMarkedAsUnread bool) *UpdateChatIsMarkedAsUnread {
	updateChatIsMarkedAsUnreadTemp := UpdateChatIsMarkedAsUnread{
		tdCommon:         tdCommon{Type: "updateChatIsMarkedAsUnread"},
		ChatID:           chatID,
		IsMarkedAsUnread: isMarkedAsUnread,
	}

	return &updateChatIsMarkedAsUnreadTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsMarkedAsUnread *UpdateChatIsMarkedAsUnread) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsMarkedAsUnreadType
}

// UpdateChatIsSponsored A chat's is_sponsored field has changed
type UpdateChatIsSponsored struct {
	tdCommon
	ChatID      int64     `json:"chat_id"`      // Chat identifier
	IsSponsored bool      `json:"is_sponsored"` // New value of is_sponsored
	Order       JSONInt64 `json:"order"`        // New value of chat order
}

// MessageType return the string telegram-type of UpdateChatIsSponsored
func (updateChatIsSponsored *UpdateChatIsSponsored) MessageType() string {
	return "updateChatIsSponsored"
}

// NewUpdateChatIsSponsored creates a new UpdateChatIsSponsored
//
// @param chatID Chat identifier
// @param isSponsored New value of is_sponsored
// @param order New value of chat order
func NewUpdateChatIsSponsored(chatID int64, isSponsored bool, order JSONInt64) *UpdateChatIsSponsored {
	updateChatIsSponsoredTemp := UpdateChatIsSponsored{
		tdCommon:    tdCommon{Type: "updateChatIsSponsored"},
		ChatID:      chatID,
		IsSponsored: isSponsored,
		Order:       order,
	}

	return &updateChatIsSponsoredTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatIsSponsored *UpdateChatIsSponsored) GetUpdateEnum() UpdateEnum {
	return UpdateChatIsSponsoredType
}

// UpdateChatDefaultDisableNotification The value of the default disable_notification parameter, used when a message is sent to the chat, was changed
type UpdateChatDefaultDisableNotification struct {
	tdCommon
	ChatID                     int64 `json:"chat_id"`                      // Chat identifier
	DefaultDisableNotification bool  `json:"default_disable_notification"` // The new default_disable_notification value
}

// MessageType return the string telegram-type of UpdateChatDefaultDisableNotification
func (updateChatDefaultDisableNotification *UpdateChatDefaultDisableNotification) MessageType() string {
	return "updateChatDefaultDisableNotification"
}

// NewUpdateChatDefaultDisableNotification creates a new UpdateChatDefaultDisableNotification
//
// @param chatID Chat identifier
// @param defaultDisableNotification The new default_disable_notification value
func NewUpdateChatDefaultDisableNotification(chatID int64, defaultDisableNotification bool) *UpdateChatDefaultDisableNotification {
	updateChatDefaultDisableNotificationTemp := UpdateChatDefaultDisableNotification{
		tdCommon:                   tdCommon{Type: "updateChatDefaultDisableNotification"},
		ChatID:                     chatID,
		DefaultDisableNotification: defaultDisableNotification,
	}

	return &updateChatDefaultDisableNotificationTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatDefaultDisableNotification *UpdateChatDefaultDisableNotification) GetUpdateEnum() UpdateEnum {
	return UpdateChatDefaultDisableNotificationType
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

// NewUpdateChatReadInbox creates a new UpdateChatReadInbox
//
// @param chatID Chat identifier
// @param lastReadInboxMessageID Identifier of the last read incoming message
// @param unreadCount The number of unread messages left in the chat
func NewUpdateChatReadInbox(chatID int64, lastReadInboxMessageID int64, unreadCount int32) *UpdateChatReadInbox {
	updateChatReadInboxTemp := UpdateChatReadInbox{
		tdCommon:               tdCommon{Type: "updateChatReadInbox"},
		ChatID:                 chatID,
		LastReadInboxMessageID: lastReadInboxMessageID,
		UnreadCount:            unreadCount,
	}

	return &updateChatReadInboxTemp
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

// NewUpdateChatReadOutbox creates a new UpdateChatReadOutbox
//
// @param chatID Chat identifier
// @param lastReadOutboxMessageID Identifier of last read outgoing message
func NewUpdateChatReadOutbox(chatID int64, lastReadOutboxMessageID int64) *UpdateChatReadOutbox {
	updateChatReadOutboxTemp := UpdateChatReadOutbox{
		tdCommon:                tdCommon{Type: "updateChatReadOutbox"},
		ChatID:                  chatID,
		LastReadOutboxMessageID: lastReadOutboxMessageID,
	}

	return &updateChatReadOutboxTemp
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

// NewUpdateChatUnreadMentionCount creates a new UpdateChatUnreadMentionCount
//
// @param chatID Chat identifier
// @param unreadMentionCount The number of unread mention messages left in the chat
func NewUpdateChatUnreadMentionCount(chatID int64, unreadMentionCount int32) *UpdateChatUnreadMentionCount {
	updateChatUnreadMentionCountTemp := UpdateChatUnreadMentionCount{
		tdCommon:           tdCommon{Type: "updateChatUnreadMentionCount"},
		ChatID:             chatID,
		UnreadMentionCount: unreadMentionCount,
	}

	return &updateChatUnreadMentionCountTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) GetUpdateEnum() UpdateEnum {
	return UpdateChatUnreadMentionCountType
}

// UpdateChatNotificationSettings Notification settings for a chat were changed
type UpdateChatNotificationSettings struct {
	tdCommon
	ChatID               int64                     `json:"chat_id"`               // Chat identifier
	NotificationSettings *ChatNotificationSettings `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateChatNotificationSettings
func (updateChatNotificationSettings *UpdateChatNotificationSettings) MessageType() string {
	return "updateChatNotificationSettings"
}

// NewUpdateChatNotificationSettings creates a new UpdateChatNotificationSettings
//
// @param chatID Chat identifier
// @param notificationSettings The new notification settings
func NewUpdateChatNotificationSettings(chatID int64, notificationSettings *ChatNotificationSettings) *UpdateChatNotificationSettings {
	updateChatNotificationSettingsTemp := UpdateChatNotificationSettings{
		tdCommon:             tdCommon{Type: "updateChatNotificationSettings"},
		ChatID:               chatID,
		NotificationSettings: notificationSettings,
	}

	return &updateChatNotificationSettingsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatNotificationSettings *UpdateChatNotificationSettings) GetUpdateEnum() UpdateEnum {
	return UpdateChatNotificationSettingsType
}

// UpdateScopeNotificationSettings Notification settings for some type of chats were updated
type UpdateScopeNotificationSettings struct {
	tdCommon
	Scope                NotificationSettingsScope  `json:"scope"`                 // Types of chats for which notification settings were updated
	NotificationSettings *ScopeNotificationSettings `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateScopeNotificationSettings
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) MessageType() string {
	return "updateScopeNotificationSettings"
}

// NewUpdateScopeNotificationSettings creates a new UpdateScopeNotificationSettings
//
// @param scope Types of chats for which notification settings were updated
// @param notificationSettings The new notification settings
func NewUpdateScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) *UpdateScopeNotificationSettings {
	updateScopeNotificationSettingsTemp := UpdateScopeNotificationSettings{
		tdCommon:             tdCommon{Type: "updateScopeNotificationSettings"},
		Scope:                scope,
		NotificationSettings: notificationSettings,
	}

	return &updateScopeNotificationSettingsTemp
}

// UnmarshalJSON unmarshal to json
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		NotificationSettings *ScopeNotificationSettings `json:"notification_settings"` // The new notification settings
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	updateScopeNotificationSettings.tdCommon = tempObj.tdCommon
	updateScopeNotificationSettings.NotificationSettings = tempObj.NotificationSettings

	fieldScope, _ := unmarshalNotificationSettingsScope(objMap["scope"])
	updateScopeNotificationSettings.Scope = fieldScope

	return nil
}

// GetUpdateEnum return the enum type of this object
func (updateScopeNotificationSettings *UpdateScopeNotificationSettings) GetUpdateEnum() UpdateEnum {
	return UpdateScopeNotificationSettingsType
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

// NewUpdateChatReplyMarkup creates a new UpdateChatReplyMarkup
//
// @param chatID Chat identifier
// @param replyMarkupMessageID Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
func NewUpdateChatReplyMarkup(chatID int64, replyMarkupMessageID int64) *UpdateChatReplyMarkup {
	updateChatReplyMarkupTemp := UpdateChatReplyMarkup{
		tdCommon:             tdCommon{Type: "updateChatReplyMarkup"},
		ChatID:               chatID,
		ReplyMarkupMessageID: replyMarkupMessageID,
	}

	return &updateChatReplyMarkupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateChatReplyMarkup *UpdateChatReplyMarkup) GetUpdateEnum() UpdateEnum {
	return UpdateChatReplyMarkupType
}

// UpdateChatDraftMessage A chat draft has changed. Be aware that the update may come in the currently opened chat but with old content of the draft. If the user has changed the content of the draft, this update shouldn't be applied
type UpdateChatDraftMessage struct {
	tdCommon
	ChatID       int64         `json:"chat_id"`       // Chat identifier
	DraftMessage *DraftMessage `json:"draft_message"` // The new draft message; may be null
	Order        JSONInt64     `json:"order"`         // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatDraftMessage
func (updateChatDraftMessage *UpdateChatDraftMessage) MessageType() string {
	return "updateChatDraftMessage"
}

// NewUpdateChatDraftMessage creates a new UpdateChatDraftMessage
//
// @param chatID Chat identifier
// @param draftMessage The new draft message; may be null
// @param order New value of the chat order
func NewUpdateChatDraftMessage(chatID int64, draftMessage *DraftMessage, order JSONInt64) *UpdateChatDraftMessage {
	updateChatDraftMessageTemp := UpdateChatDraftMessage{
		tdCommon:     tdCommon{Type: "updateChatDraftMessage"},
		ChatID:       chatID,
		DraftMessage: draftMessage,
		Order:        order,
	}

	return &updateChatDraftMessageTemp
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

// NewUpdateDeleteMessages creates a new UpdateDeleteMessages
//
// @param chatID Chat identifier
// @param messageIDs Identifiers of the deleted messages
// @param isPermanent True, if the messages are permanently deleted by a user (as opposed to just becoming unaccessible)
// @param fromCache True, if the messages are deleted only from the cache and can possibly be retrieved again in the future
func NewUpdateDeleteMessages(chatID int64, messageIDs []int64, isPermanent bool, fromCache bool) *UpdateDeleteMessages {
	updateDeleteMessagesTemp := UpdateDeleteMessages{
		tdCommon:    tdCommon{Type: "updateDeleteMessages"},
		ChatID:      chatID,
		MessageIDs:  messageIDs,
		IsPermanent: isPermanent,
		FromCache:   fromCache,
	}

	return &updateDeleteMessagesTemp
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

// NewUpdateUserChatAction creates a new UpdateUserChatAction
//
// @param chatID Chat identifier
// @param userID Identifier of a user performing an action
// @param action The action description
func NewUpdateUserChatAction(chatID int64, userID int32, action ChatAction) *UpdateUserChatAction {
	updateUserChatActionTemp := UpdateUserChatAction{
		tdCommon: tdCommon{Type: "updateUserChatAction"},
		ChatID:   chatID,
		UserID:   userID,
		Action:   action,
	}

	return &updateUserChatActionTemp
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

// NewUpdateUserStatus creates a new UpdateUserStatus
//
// @param userID User identifier
// @param status New status of the user
func NewUpdateUserStatus(userID int32, status UserStatus) *UpdateUserStatus {
	updateUserStatusTemp := UpdateUserStatus{
		tdCommon: tdCommon{Type: "updateUserStatus"},
		UserID:   userID,
		Status:   status,
	}

	return &updateUserStatusTemp
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
	User *User `json:"user"` // New data about the user
}

// MessageType return the string telegram-type of UpdateUser
func (updateUser *UpdateUser) MessageType() string {
	return "updateUser"
}

// NewUpdateUser creates a new UpdateUser
//
// @param user New data about the user
func NewUpdateUser(user *User) *UpdateUser {
	updateUserTemp := UpdateUser{
		tdCommon: tdCommon{Type: "updateUser"},
		User:     user,
	}

	return &updateUserTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUser *UpdateUser) GetUpdateEnum() UpdateEnum {
	return UpdateUserType
}

// UpdateBasicGroup Some data of a basic group has changed. This update is guaranteed to come before the basic group identifier is returned to the client
type UpdateBasicGroup struct {
	tdCommon
	BasicGroup *BasicGroup `json:"basic_group"` // New data about the group
}

// MessageType return the string telegram-type of UpdateBasicGroup
func (updateBasicGroup *UpdateBasicGroup) MessageType() string {
	return "updateBasicGroup"
}

// NewUpdateBasicGroup creates a new UpdateBasicGroup
//
// @param basicGroup New data about the group
func NewUpdateBasicGroup(basicGroup *BasicGroup) *UpdateBasicGroup {
	updateBasicGroupTemp := UpdateBasicGroup{
		tdCommon:   tdCommon{Type: "updateBasicGroup"},
		BasicGroup: basicGroup,
	}

	return &updateBasicGroupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroup *UpdateBasicGroup) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupType
}

// UpdateSupergroup Some data of a supergroup or a channel has changed. This update is guaranteed to come before the supergroup identifier is returned to the client
type UpdateSupergroup struct {
	tdCommon
	Supergroup *Supergroup `json:"supergroup"` // New data about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroup
func (updateSupergroup *UpdateSupergroup) MessageType() string {
	return "updateSupergroup"
}

// NewUpdateSupergroup creates a new UpdateSupergroup
//
// @param supergroup New data about the supergroup
func NewUpdateSupergroup(supergroup *Supergroup) *UpdateSupergroup {
	updateSupergroupTemp := UpdateSupergroup{
		tdCommon:   tdCommon{Type: "updateSupergroup"},
		Supergroup: supergroup,
	}

	return &updateSupergroupTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroup *UpdateSupergroup) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupType
}

// UpdateSecretChat Some data of a secret chat has changed. This update is guaranteed to come before the secret chat identifier is returned to the client
type UpdateSecretChat struct {
	tdCommon
	SecretChat *SecretChat `json:"secret_chat"` // New data about the secret chat
}

// MessageType return the string telegram-type of UpdateSecretChat
func (updateSecretChat *UpdateSecretChat) MessageType() string {
	return "updateSecretChat"
}

// NewUpdateSecretChat creates a new UpdateSecretChat
//
// @param secretChat New data about the secret chat
func NewUpdateSecretChat(secretChat *SecretChat) *UpdateSecretChat {
	updateSecretChatTemp := UpdateSecretChat{
		tdCommon:   tdCommon{Type: "updateSecretChat"},
		SecretChat: secretChat,
	}

	return &updateSecretChatTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSecretChat *UpdateSecretChat) GetUpdateEnum() UpdateEnum {
	return UpdateSecretChatType
}

// UpdateUserFullInfo Some data from userFullInfo has been changed
type UpdateUserFullInfo struct {
	tdCommon
	UserID       int32         `json:"user_id"`        // User identifier
	UserFullInfo *UserFullInfo `json:"user_full_info"` // New full information about the user
}

// MessageType return the string telegram-type of UpdateUserFullInfo
func (updateUserFullInfo *UpdateUserFullInfo) MessageType() string {
	return "updateUserFullInfo"
}

// NewUpdateUserFullInfo creates a new UpdateUserFullInfo
//
// @param userID User identifier
// @param userFullInfo New full information about the user
func NewUpdateUserFullInfo(userID int32, userFullInfo *UserFullInfo) *UpdateUserFullInfo {
	updateUserFullInfoTemp := UpdateUserFullInfo{
		tdCommon:     tdCommon{Type: "updateUserFullInfo"},
		UserID:       userID,
		UserFullInfo: userFullInfo,
	}

	return &updateUserFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUserFullInfo *UpdateUserFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateUserFullInfoType
}

// UpdateBasicGroupFullInfo Some data from basicGroupFullInfo has been changed
type UpdateBasicGroupFullInfo struct {
	tdCommon
	BasicGroupID       int32               `json:"basic_group_id"`        // Identifier of a basic group
	BasicGroupFullInfo *BasicGroupFullInfo `json:"basic_group_full_info"` // New full information about the group
}

// MessageType return the string telegram-type of UpdateBasicGroupFullInfo
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) MessageType() string {
	return "updateBasicGroupFullInfo"
}

// NewUpdateBasicGroupFullInfo creates a new UpdateBasicGroupFullInfo
//
// @param basicGroupID Identifier of a basic group
// @param basicGroupFullInfo New full information about the group
func NewUpdateBasicGroupFullInfo(basicGroupID int32, basicGroupFullInfo *BasicGroupFullInfo) *UpdateBasicGroupFullInfo {
	updateBasicGroupFullInfoTemp := UpdateBasicGroupFullInfo{
		tdCommon:           tdCommon{Type: "updateBasicGroupFullInfo"},
		BasicGroupID:       basicGroupID,
		BasicGroupFullInfo: basicGroupFullInfo,
	}

	return &updateBasicGroupFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateBasicGroupFullInfoType
}

// UpdateSupergroupFullInfo Some data from supergroupFullInfo has been changed
type UpdateSupergroupFullInfo struct {
	tdCommon
	SupergroupID       int32               `json:"supergroup_id"`        // Identifier of the supergroup or channel
	SupergroupFullInfo *SupergroupFullInfo `json:"supergroup_full_info"` // New full information about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroupFullInfo
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) MessageType() string {
	return "updateSupergroupFullInfo"
}

// NewUpdateSupergroupFullInfo creates a new UpdateSupergroupFullInfo
//
// @param supergroupID Identifier of the supergroup or channel
// @param supergroupFullInfo New full information about the supergroup
func NewUpdateSupergroupFullInfo(supergroupID int32, supergroupFullInfo *SupergroupFullInfo) *UpdateSupergroupFullInfo {
	updateSupergroupFullInfoTemp := UpdateSupergroupFullInfo{
		tdCommon:           tdCommon{Type: "updateSupergroupFullInfo"},
		SupergroupID:       supergroupID,
		SupergroupFullInfo: supergroupFullInfo,
	}

	return &updateSupergroupFullInfoTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) GetUpdateEnum() UpdateEnum {
	return UpdateSupergroupFullInfoType
}

// UpdateServiceNotification Service notification from the server. Upon receiving this the client must show a popup with the content of the notification
type UpdateServiceNotification struct {
	tdCommon
	Type    string         `json:"type"`    // Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" should be shown under notification; if user presses the second, all local data should be destroyed using Destroy method
	Content MessageContent `json:"content"` // Notification content
}

// MessageType return the string telegram-type of UpdateServiceNotification
func (updateServiceNotification *UpdateServiceNotification) MessageType() string {
	return "updateServiceNotification"
}

// NewUpdateServiceNotification creates a new UpdateServiceNotification
//
// @param typeParam Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" should be shown under notification; if user presses the second, all local data should be destroyed using Destroy method
// @param content Notification content
func NewUpdateServiceNotification(typeParam string, content MessageContent) *UpdateServiceNotification {
	updateServiceNotificationTemp := UpdateServiceNotification{
		tdCommon: tdCommon{Type: "updateServiceNotification"},
		Type:     typeParam,
		Content:  content,
	}

	return &updateServiceNotificationTemp
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
		Type string `json:"type"` // Notification type. If type begins with "AUTH_KEY_DROP_", then two buttons "Cancel" and "Log out" should be shown under notification; if user presses the second, all local data should be destroyed using Destroy method

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
	File *File `json:"file"` // New data about the file
}

// MessageType return the string telegram-type of UpdateFile
func (updateFile *UpdateFile) MessageType() string {
	return "updateFile"
}

// NewUpdateFile creates a new UpdateFile
//
// @param file New data about the file
func NewUpdateFile(file *File) *UpdateFile {
	updateFileTemp := UpdateFile{
		tdCommon: tdCommon{Type: "updateFile"},
		File:     file,
	}

	return &updateFileTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFile *UpdateFile) GetUpdateEnum() UpdateEnum {
	return UpdateFileType
}

// UpdateFileGenerationStart The file generation process needs to be started by the client
type UpdateFileGenerationStart struct {
	tdCommon
	GenerationID    JSONInt64 `json:"generation_id"`    // Unique identifier for the generation process
	OriginalPath    string    `json:"original_path"`    // The path to a file from which a new file is generated; may be empty
	DestinationPath string    `json:"destination_path"` // The path to a file that should be created and where the new file should be generated
	Conversion      string    `json:"conversion"`       // String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains an HTTP/HTTPS URL of a file, which should be downloaded by the client
}

// MessageType return the string telegram-type of UpdateFileGenerationStart
func (updateFileGenerationStart *UpdateFileGenerationStart) MessageType() string {
	return "updateFileGenerationStart"
}

// NewUpdateFileGenerationStart creates a new UpdateFileGenerationStart
//
// @param generationID Unique identifier for the generation process
// @param originalPath The path to a file from which a new file is generated; may be empty
// @param destinationPath The path to a file that should be created and where the new file should be generated
// @param conversion String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains an HTTP/HTTPS URL of a file, which should be downloaded by the client
func NewUpdateFileGenerationStart(generationID JSONInt64, originalPath string, destinationPath string, conversion string) *UpdateFileGenerationStart {
	updateFileGenerationStartTemp := UpdateFileGenerationStart{
		tdCommon:        tdCommon{Type: "updateFileGenerationStart"},
		GenerationID:    generationID,
		OriginalPath:    originalPath,
		DestinationPath: destinationPath,
		Conversion:      conversion,
	}

	return &updateFileGenerationStartTemp
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

// NewUpdateFileGenerationStop creates a new UpdateFileGenerationStop
//
// @param generationID Unique identifier for the generation process
func NewUpdateFileGenerationStop(generationID JSONInt64) *UpdateFileGenerationStop {
	updateFileGenerationStopTemp := UpdateFileGenerationStop{
		tdCommon:     tdCommon{Type: "updateFileGenerationStop"},
		GenerationID: generationID,
	}

	return &updateFileGenerationStopTemp
}

// GetUpdateEnum return the enum type of this object
func (updateFileGenerationStop *UpdateFileGenerationStop) GetUpdateEnum() UpdateEnum {
	return UpdateFileGenerationStopType
}

// UpdateCall New call was created or information about a call was updated
type UpdateCall struct {
	tdCommon
	Call *Call `json:"call"` // New data about a call
}

// MessageType return the string telegram-type of UpdateCall
func (updateCall *UpdateCall) MessageType() string {
	return "updateCall"
}

// NewUpdateCall creates a new UpdateCall
//
// @param call New data about a call
func NewUpdateCall(call *Call) *UpdateCall {
	updateCallTemp := UpdateCall{
		tdCommon: tdCommon{Type: "updateCall"},
		Call:     call,
	}

	return &updateCallTemp
}

// GetUpdateEnum return the enum type of this object
func (updateCall *UpdateCall) GetUpdateEnum() UpdateEnum {
	return UpdateCallType
}

// UpdateUserPrivacySettingRules Some privacy setting rules have been changed
type UpdateUserPrivacySettingRules struct {
	tdCommon
	Setting UserPrivacySetting       `json:"setting"` // The privacy setting
	Rules   *UserPrivacySettingRules `json:"rules"`   // New privacy rules
}

// MessageType return the string telegram-type of UpdateUserPrivacySettingRules
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) MessageType() string {
	return "updateUserPrivacySettingRules"
}

// NewUpdateUserPrivacySettingRules creates a new UpdateUserPrivacySettingRules
//
// @param setting The privacy setting
// @param rules New privacy rules
func NewUpdateUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) *UpdateUserPrivacySettingRules {
	updateUserPrivacySettingRulesTemp := UpdateUserPrivacySettingRules{
		tdCommon: tdCommon{Type: "updateUserPrivacySettingRules"},
		Setting:  setting,
		Rules:    rules,
	}

	return &updateUserPrivacySettingRulesTemp
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
		Rules *UserPrivacySettingRules `json:"rules"` // New privacy rules
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

// NewUpdateUnreadMessageCount creates a new UpdateUnreadMessageCount
//
// @param unreadCount Total number of unread messages
// @param unreadUnmutedCount Total number of unread messages in unmuted chats
func NewUpdateUnreadMessageCount(unreadCount int32, unreadUnmutedCount int32) *UpdateUnreadMessageCount {
	updateUnreadMessageCountTemp := UpdateUnreadMessageCount{
		tdCommon:           tdCommon{Type: "updateUnreadMessageCount"},
		UnreadCount:        unreadCount,
		UnreadUnmutedCount: unreadUnmutedCount,
	}

	return &updateUnreadMessageCountTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUnreadMessageCount *UpdateUnreadMessageCount) GetUpdateEnum() UpdateEnum {
	return UpdateUnreadMessageCountType
}

// UpdateUnreadChatCount Number of unread chats, i.e. with unread messages or marked as unread, has changed. This update is sent only if a message database is used
type UpdateUnreadChatCount struct {
	tdCommon
	UnreadCount                int32 `json:"unread_count"`                   // Total number of unread chats
	UnreadUnmutedCount         int32 `json:"unread_unmuted_count"`           // Total number of unread unmuted chats
	MarkedAsUnreadCount        int32 `json:"marked_as_unread_count"`         // Total number of chats marked as unread
	MarkedAsUnreadUnmutedCount int32 `json:"marked_as_unread_unmuted_count"` // Total number of unmuted chats marked as unread
}

// MessageType return the string telegram-type of UpdateUnreadChatCount
func (updateUnreadChatCount *UpdateUnreadChatCount) MessageType() string {
	return "updateUnreadChatCount"
}

// NewUpdateUnreadChatCount creates a new UpdateUnreadChatCount
//
// @param unreadCount Total number of unread chats
// @param unreadUnmutedCount Total number of unread unmuted chats
// @param markedAsUnreadCount Total number of chats marked as unread
// @param markedAsUnreadUnmutedCount Total number of unmuted chats marked as unread
func NewUpdateUnreadChatCount(unreadCount int32, unreadUnmutedCount int32, markedAsUnreadCount int32, markedAsUnreadUnmutedCount int32) *UpdateUnreadChatCount {
	updateUnreadChatCountTemp := UpdateUnreadChatCount{
		tdCommon:                   tdCommon{Type: "updateUnreadChatCount"},
		UnreadCount:                unreadCount,
		UnreadUnmutedCount:         unreadUnmutedCount,
		MarkedAsUnreadCount:        markedAsUnreadCount,
		MarkedAsUnreadUnmutedCount: markedAsUnreadUnmutedCount,
	}

	return &updateUnreadChatCountTemp
}

// GetUpdateEnum return the enum type of this object
func (updateUnreadChatCount *UpdateUnreadChatCount) GetUpdateEnum() UpdateEnum {
	return UpdateUnreadChatCountType
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

// NewUpdateOption creates a new UpdateOption
//
// @param name The option name
// @param value The new option value
func NewUpdateOption(name string, value OptionValue) *UpdateOption {
	updateOptionTemp := UpdateOption{
		tdCommon: tdCommon{Type: "updateOption"},
		Name:     name,
		Value:    value,
	}

	return &updateOptionTemp
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

// NewUpdateInstalledStickerSets creates a new UpdateInstalledStickerSets
//
// @param isMasks True, if the list of installed mask sticker sets was updated
// @param stickerSetIDs The new list of installed ordinary sticker sets
func NewUpdateInstalledStickerSets(isMasks bool, stickerSetIDs []JSONInt64) *UpdateInstalledStickerSets {
	updateInstalledStickerSetsTemp := UpdateInstalledStickerSets{
		tdCommon:      tdCommon{Type: "updateInstalledStickerSets"},
		IsMasks:       isMasks,
		StickerSetIDs: stickerSetIDs,
	}

	return &updateInstalledStickerSetsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateInstalledStickerSets *UpdateInstalledStickerSets) GetUpdateEnum() UpdateEnum {
	return UpdateInstalledStickerSetsType
}

// UpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed
type UpdateTrendingStickerSets struct {
	tdCommon
	StickerSets *StickerSets `json:"sticker_sets"` // The new list of trending sticker sets
}

// MessageType return the string telegram-type of UpdateTrendingStickerSets
func (updateTrendingStickerSets *UpdateTrendingStickerSets) MessageType() string {
	return "updateTrendingStickerSets"
}

// NewUpdateTrendingStickerSets creates a new UpdateTrendingStickerSets
//
// @param stickerSets The new list of trending sticker sets
func NewUpdateTrendingStickerSets(stickerSets *StickerSets) *UpdateTrendingStickerSets {
	updateTrendingStickerSetsTemp := UpdateTrendingStickerSets{
		tdCommon:    tdCommon{Type: "updateTrendingStickerSets"},
		StickerSets: stickerSets,
	}

	return &updateTrendingStickerSetsTemp
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

// NewUpdateRecentStickers creates a new UpdateRecentStickers
//
// @param isAttached True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated
// @param stickerIDs The new list of file identifiers of recently used stickers
func NewUpdateRecentStickers(isAttached bool, stickerIDs []int32) *UpdateRecentStickers {
	updateRecentStickersTemp := UpdateRecentStickers{
		tdCommon:   tdCommon{Type: "updateRecentStickers"},
		IsAttached: isAttached,
		StickerIDs: stickerIDs,
	}

	return &updateRecentStickersTemp
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

// NewUpdateFavoriteStickers creates a new UpdateFavoriteStickers
//
// @param stickerIDs The new list of file identifiers of favorite stickers
func NewUpdateFavoriteStickers(stickerIDs []int32) *UpdateFavoriteStickers {
	updateFavoriteStickersTemp := UpdateFavoriteStickers{
		tdCommon:   tdCommon{Type: "updateFavoriteStickers"},
		StickerIDs: stickerIDs,
	}

	return &updateFavoriteStickersTemp
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

// NewUpdateSavedAnimations creates a new UpdateSavedAnimations
//
// @param animationIDs The new list of file identifiers of saved animations
func NewUpdateSavedAnimations(animationIDs []int32) *UpdateSavedAnimations {
	updateSavedAnimationsTemp := UpdateSavedAnimations{
		tdCommon:     tdCommon{Type: "updateSavedAnimations"},
		AnimationIDs: animationIDs,
	}

	return &updateSavedAnimationsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateSavedAnimations *UpdateSavedAnimations) GetUpdateEnum() UpdateEnum {
	return UpdateSavedAnimationsType
}

// UpdateLanguagePackStrings Some language pack strings have been updated
type UpdateLanguagePackStrings struct {
	tdCommon
	LocalizationTarget string               `json:"localization_target"` // Localization target to which the language pack belongs
	LanguagePackID     string               `json:"language_pack_id"`    // Identifier of the updated language pack
	Strings            []LanguagePackString `json:"strings"`             // List of changed language pack strings
}

// MessageType return the string telegram-type of UpdateLanguagePackStrings
func (updateLanguagePackStrings *UpdateLanguagePackStrings) MessageType() string {
	return "updateLanguagePackStrings"
}

// NewUpdateLanguagePackStrings creates a new UpdateLanguagePackStrings
//
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackID Identifier of the updated language pack
// @param strings List of changed language pack strings
func NewUpdateLanguagePackStrings(localizationTarget string, languagePackID string, strings []LanguagePackString) *UpdateLanguagePackStrings {
	updateLanguagePackStringsTemp := UpdateLanguagePackStrings{
		tdCommon:           tdCommon{Type: "updateLanguagePackStrings"},
		LocalizationTarget: localizationTarget,
		LanguagePackID:     languagePackID,
		Strings:            strings,
	}

	return &updateLanguagePackStringsTemp
}

// GetUpdateEnum return the enum type of this object
func (updateLanguagePackStrings *UpdateLanguagePackStrings) GetUpdateEnum() UpdateEnum {
	return UpdateLanguagePackStringsType
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

// NewUpdateConnectionState creates a new UpdateConnectionState
//
// @param state The new connection state
func NewUpdateConnectionState(state ConnectionState) *UpdateConnectionState {
	updateConnectionStateTemp := UpdateConnectionState{
		tdCommon: tdCommon{Type: "updateConnectionState"},
		State:    state,
	}

	return &updateConnectionStateTemp
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

// UpdateTermsOfService New terms of service must be accepted by the user. If the terms of service are declined, then the deleteAccount method should be called with the reason "Decline ToS update"
type UpdateTermsOfService struct {
	tdCommon
	TermsOfServiceID string          `json:"terms_of_service_id"` // Identifier of the terms of service
	TermsOfService   *TermsOfService `json:"terms_of_service"`    // The new terms of service
}

// MessageType return the string telegram-type of UpdateTermsOfService
func (updateTermsOfService *UpdateTermsOfService) MessageType() string {
	return "updateTermsOfService"
}

// NewUpdateTermsOfService creates a new UpdateTermsOfService
//
// @param termsOfServiceID Identifier of the terms of service
// @param termsOfService The new terms of service
func NewUpdateTermsOfService(termsOfServiceID string, termsOfService *TermsOfService) *UpdateTermsOfService {
	updateTermsOfServiceTemp := UpdateTermsOfService{
		tdCommon:         tdCommon{Type: "updateTermsOfService"},
		TermsOfServiceID: termsOfServiceID,
		TermsOfService:   termsOfService,
	}

	return &updateTermsOfServiceTemp
}

// GetUpdateEnum return the enum type of this object
func (updateTermsOfService *UpdateTermsOfService) GetUpdateEnum() UpdateEnum {
	return UpdateTermsOfServiceType
}

// UpdateNewInlineQuery A new incoming inline query; for bots only
type UpdateNewInlineQuery struct {
	tdCommon
	ID           JSONInt64 `json:"id"`             // Unique query identifier
	SenderUserID int32     `json:"sender_user_id"` // Identifier of the user who sent the query
	UserLocation *Location `json:"user_location"`  // User location, provided by the client; may be null
	Query        string    `json:"query"`          // Text of the query
	Offset       string    `json:"offset"`         // Offset of the first entry to return
}

// MessageType return the string telegram-type of UpdateNewInlineQuery
func (updateNewInlineQuery *UpdateNewInlineQuery) MessageType() string {
	return "updateNewInlineQuery"
}

// NewUpdateNewInlineQuery creates a new UpdateNewInlineQuery
//
// @param iD Unique query identifier
// @param senderUserID Identifier of the user who sent the query
// @param userLocation User location, provided by the client; may be null
// @param query Text of the query
// @param offset Offset of the first entry to return
func NewUpdateNewInlineQuery(iD JSONInt64, senderUserID int32, userLocation *Location, query string, offset string) *UpdateNewInlineQuery {
	updateNewInlineQueryTemp := UpdateNewInlineQuery{
		tdCommon:     tdCommon{Type: "updateNewInlineQuery"},
		ID:           iD,
		SenderUserID: senderUserID,
		UserLocation: userLocation,
		Query:        query,
		Offset:       offset,
	}

	return &updateNewInlineQueryTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewInlineQuery *UpdateNewInlineQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewInlineQueryType
}

// UpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only
type UpdateNewChosenInlineResult struct {
	tdCommon
	SenderUserID    int32     `json:"sender_user_id"`    // Identifier of the user who sent the query
	UserLocation    *Location `json:"user_location"`     // User location, provided by the client; may be null
	Query           string    `json:"query"`             // Text of the query
	ResultID        string    `json:"result_id"`         // Identifier of the chosen result
	InlineMessageID string    `json:"inline_message_id"` // Identifier of the sent inline message, if known
}

// MessageType return the string telegram-type of UpdateNewChosenInlineResult
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) MessageType() string {
	return "updateNewChosenInlineResult"
}

// NewUpdateNewChosenInlineResult creates a new UpdateNewChosenInlineResult
//
// @param senderUserID Identifier of the user who sent the query
// @param userLocation User location, provided by the client; may be null
// @param query Text of the query
// @param resultID Identifier of the chosen result
// @param inlineMessageID Identifier of the sent inline message, if known
func NewUpdateNewChosenInlineResult(senderUserID int32, userLocation *Location, query string, resultID string, inlineMessageID string) *UpdateNewChosenInlineResult {
	updateNewChosenInlineResultTemp := UpdateNewChosenInlineResult{
		tdCommon:        tdCommon{Type: "updateNewChosenInlineResult"},
		SenderUserID:    senderUserID,
		UserLocation:    userLocation,
		Query:           query,
		ResultID:        resultID,
		InlineMessageID: inlineMessageID,
	}

	return &updateNewChosenInlineResultTemp
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

// NewUpdateNewCallbackQuery creates a new UpdateNewCallbackQuery
//
// @param iD Unique query identifier
// @param senderUserID Identifier of the user who sent the query
// @param chatID Identifier of the chat, in which the query was sent
// @param messageID Identifier of the message, from which the query originated
// @param chatInstance Identifier that uniquely corresponds to the chat to which the message was sent
// @param payload Query payload
func NewUpdateNewCallbackQuery(iD JSONInt64, senderUserID int32, chatID int64, messageID int64, chatInstance JSONInt64, payload CallbackQueryPayload) *UpdateNewCallbackQuery {
	updateNewCallbackQueryTemp := UpdateNewCallbackQuery{
		tdCommon:     tdCommon{Type: "updateNewCallbackQuery"},
		ID:           iD,
		SenderUserID: senderUserID,
		ChatID:       chatID,
		MessageID:    messageID,
		ChatInstance: chatInstance,
		Payload:      payload,
	}

	return &updateNewCallbackQueryTemp
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

// NewUpdateNewInlineCallbackQuery creates a new UpdateNewInlineCallbackQuery
//
// @param iD Unique query identifier
// @param senderUserID Identifier of the user who sent the query
// @param inlineMessageID Identifier of the inline message, from which the query originated
// @param chatInstance An identifier uniquely corresponding to the chat a message was sent to
// @param payload Query payload
func NewUpdateNewInlineCallbackQuery(iD JSONInt64, senderUserID int32, inlineMessageID string, chatInstance JSONInt64, payload CallbackQueryPayload) *UpdateNewInlineCallbackQuery {
	updateNewInlineCallbackQueryTemp := UpdateNewInlineCallbackQuery{
		tdCommon:        tdCommon{Type: "updateNewInlineCallbackQuery"},
		ID:              iD,
		SenderUserID:    senderUserID,
		InlineMessageID: inlineMessageID,
		ChatInstance:    chatInstance,
		Payload:         payload,
	}

	return &updateNewInlineCallbackQueryTemp
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
	ID              JSONInt64 `json:"id"`               // Unique query identifier
	SenderUserID    int32     `json:"sender_user_id"`   // Identifier of the user who sent the query
	InvoicePayload  string    `json:"invoice_payload"`  // Invoice payload
	ShippingAddress *Address  `json:"shipping_address"` // User shipping address
}

// MessageType return the string telegram-type of UpdateNewShippingQuery
func (updateNewShippingQuery *UpdateNewShippingQuery) MessageType() string {
	return "updateNewShippingQuery"
}

// NewUpdateNewShippingQuery creates a new UpdateNewShippingQuery
//
// @param iD Unique query identifier
// @param senderUserID Identifier of the user who sent the query
// @param invoicePayload Invoice payload
// @param shippingAddress User shipping address
func NewUpdateNewShippingQuery(iD JSONInt64, senderUserID int32, invoicePayload string, shippingAddress *Address) *UpdateNewShippingQuery {
	updateNewShippingQueryTemp := UpdateNewShippingQuery{
		tdCommon:        tdCommon{Type: "updateNewShippingQuery"},
		ID:              iD,
		SenderUserID:    senderUserID,
		InvoicePayload:  invoicePayload,
		ShippingAddress: shippingAddress,
	}

	return &updateNewShippingQueryTemp
}

// GetUpdateEnum return the enum type of this object
func (updateNewShippingQuery *UpdateNewShippingQuery) GetUpdateEnum() UpdateEnum {
	return UpdateNewShippingQueryType
}

// UpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout
type UpdateNewPreCheckoutQuery struct {
	tdCommon
	ID               JSONInt64  `json:"id"`                 // Unique query identifier
	SenderUserID     int32      `json:"sender_user_id"`     // Identifier of the user who sent the query
	Currency         string     `json:"currency"`           // Currency for the product price
	TotalAmount      int64      `json:"total_amount"`       // Total price for the product, in the minimal quantity of the currency
	InvoicePayload   []byte     `json:"invoice_payload"`    // Invoice payload
	ShippingOptionID string     `json:"shipping_option_id"` // Identifier of a shipping option chosen by the user; may be empty if not applicable
	OrderInfo        *OrderInfo `json:"order_info"`         // Information about the order; may be null
}

// MessageType return the string telegram-type of UpdateNewPreCheckoutQuery
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) MessageType() string {
	return "updateNewPreCheckoutQuery"
}

// NewUpdateNewPreCheckoutQuery creates a new UpdateNewPreCheckoutQuery
//
// @param iD Unique query identifier
// @param senderUserID Identifier of the user who sent the query
// @param currency Currency for the product price
// @param totalAmount Total price for the product, in the minimal quantity of the currency
// @param invoicePayload Invoice payload
// @param shippingOptionID Identifier of a shipping option chosen by the user; may be empty if not applicable
// @param orderInfo Information about the order; may be null
func NewUpdateNewPreCheckoutQuery(iD JSONInt64, senderUserID int32, currency string, totalAmount int64, invoicePayload []byte, shippingOptionID string, orderInfo *OrderInfo) *UpdateNewPreCheckoutQuery {
	updateNewPreCheckoutQueryTemp := UpdateNewPreCheckoutQuery{
		tdCommon:         tdCommon{Type: "updateNewPreCheckoutQuery"},
		ID:               iD,
		SenderUserID:     senderUserID,
		Currency:         currency,
		TotalAmount:      totalAmount,
		InvoicePayload:   invoicePayload,
		ShippingOptionID: shippingOptionID,
		OrderInfo:        orderInfo,
	}

	return &updateNewPreCheckoutQueryTemp
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

// NewUpdateNewCustomEvent creates a new UpdateNewCustomEvent
//
// @param event A JSON-serialized event
func NewUpdateNewCustomEvent(event string) *UpdateNewCustomEvent {
	updateNewCustomEventTemp := UpdateNewCustomEvent{
		tdCommon: tdCommon{Type: "updateNewCustomEvent"},
		Event:    event,
	}

	return &updateNewCustomEventTemp
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

// NewUpdateNewCustomQuery creates a new UpdateNewCustomQuery
//
// @param iD The query identifier
// @param data JSON-serialized query data
// @param timeout Query timeout
func NewUpdateNewCustomQuery(iD JSONInt64, data string, timeout int32) *UpdateNewCustomQuery {
	updateNewCustomQueryTemp := UpdateNewCustomQuery{
		tdCommon: tdCommon{Type: "updateNewCustomQuery"},
		ID:       iD,
		Data:     data,
		Timeout:  timeout,
	}

	return &updateNewCustomQueryTemp
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

// NewTestInt creates a new TestInt
//
// @param value Number
func NewTestInt(value int32) *TestInt {
	testIntTemp := TestInt{
		tdCommon: tdCommon{Type: "testInt"},
		Value:    value,
	}

	return &testIntTemp
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

// NewTestString creates a new TestString
//
// @param value String
func NewTestString(value string) *TestString {
	testStringTemp := TestString{
		tdCommon: tdCommon{Type: "testString"},
		Value:    value,
	}

	return &testStringTemp
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

// NewTestBytes creates a new TestBytes
//
// @param value Bytes
func NewTestBytes(value []byte) *TestBytes {
	testBytesTemp := TestBytes{
		tdCommon: tdCommon{Type: "testBytes"},
		Value:    value,
	}

	return &testBytesTemp
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

// NewTestVectorInt creates a new TestVectorInt
//
// @param value Vector of numbers
func NewTestVectorInt(value []int32) *TestVectorInt {
	testVectorIntTemp := TestVectorInt{
		tdCommon: tdCommon{Type: "testVectorInt"},
		Value:    value,
	}

	return &testVectorIntTemp
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

// NewTestVectorIntObject creates a new TestVectorIntObject
//
// @param value Vector of objects
func NewTestVectorIntObject(value []TestInt) *TestVectorIntObject {
	testVectorIntObjectTemp := TestVectorIntObject{
		tdCommon: tdCommon{Type: "testVectorIntObject"},
		Value:    value,
	}

	return &testVectorIntObjectTemp
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

// NewTestVectorString creates a new TestVectorString
//
// @param value Vector of strings
func NewTestVectorString(value []string) *TestVectorString {
	testVectorStringTemp := TestVectorString{
		tdCommon: tdCommon{Type: "testVectorString"},
		Value:    value,
	}

	return &testVectorStringTemp
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

// NewTestVectorStringObject creates a new TestVectorStringObject
//
// @param value Vector of objects
func NewTestVectorStringObject(value []TestString) *TestVectorStringObject {
	testVectorStringObjectTemp := TestVectorStringObject{
		tdCommon: tdCommon{Type: "testVectorStringObject"},
		Value:    value,
	}

	return &testVectorStringObjectTemp
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

func unmarshalChatMembersFilter(rawMsg *json.RawMessage) (ChatMembersFilter, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ChatMembersFilterEnum(objMap["@type"].(string)) {
	case ChatMembersFilterAdministratorsType:
		var chatMembersFilterAdministrators ChatMembersFilterAdministrators
		err := json.Unmarshal(*rawMsg, &chatMembersFilterAdministrators)
		return &chatMembersFilterAdministrators, err

	case ChatMembersFilterMembersType:
		var chatMembersFilterMembers ChatMembersFilterMembers
		err := json.Unmarshal(*rawMsg, &chatMembersFilterMembers)
		return &chatMembersFilterMembers, err

	case ChatMembersFilterRestrictedType:
		var chatMembersFilterRestricted ChatMembersFilterRestricted
		err := json.Unmarshal(*rawMsg, &chatMembersFilterRestricted)
		return &chatMembersFilterRestricted, err

	case ChatMembersFilterBannedType:
		var chatMembersFilterBanned ChatMembersFilterBanned
		err := json.Unmarshal(*rawMsg, &chatMembersFilterBanned)
		return &chatMembersFilterBanned, err

	case ChatMembersFilterBotsType:
		var chatMembersFilterBots ChatMembersFilterBots
		err := json.Unmarshal(*rawMsg, &chatMembersFilterBots)
		return &chatMembersFilterBots, err

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
	case NotificationSettingsScopePrivateChatsType:
		var notificationSettingsScopePrivateChats NotificationSettingsScopePrivateChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopePrivateChats)
		return &notificationSettingsScopePrivateChats, err

	case NotificationSettingsScopeGroupChatsType:
		var notificationSettingsScopeGroupChats NotificationSettingsScopeGroupChats
		err := json.Unmarshal(*rawMsg, &notificationSettingsScopeGroupChats)
		return &notificationSettingsScopeGroupChats, err

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

func unmarshalPassportElementType(rawMsg *json.RawMessage) (PassportElementType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementTypeEnum(objMap["@type"].(string)) {
	case PassportElementTypePersonalDetailsType:
		var passportElementTypePersonalDetails PassportElementTypePersonalDetails
		err := json.Unmarshal(*rawMsg, &passportElementTypePersonalDetails)
		return &passportElementTypePersonalDetails, err

	case PassportElementTypePassportType:
		var passportElementTypePassport PassportElementTypePassport
		err := json.Unmarshal(*rawMsg, &passportElementTypePassport)
		return &passportElementTypePassport, err

	case PassportElementTypeDriverLicenseType:
		var passportElementTypeDriverLicense PassportElementTypeDriverLicense
		err := json.Unmarshal(*rawMsg, &passportElementTypeDriverLicense)
		return &passportElementTypeDriverLicense, err

	case PassportElementTypeIDentityCardType:
		var passportElementTypeIDentityCard PassportElementTypeIDentityCard
		err := json.Unmarshal(*rawMsg, &passportElementTypeIDentityCard)
		return &passportElementTypeIDentityCard, err

	case PassportElementTypeInternalPassportType:
		var passportElementTypeInternalPassport PassportElementTypeInternalPassport
		err := json.Unmarshal(*rawMsg, &passportElementTypeInternalPassport)
		return &passportElementTypeInternalPassport, err

	case PassportElementTypeAddressType:
		var passportElementTypeAddress PassportElementTypeAddress
		err := json.Unmarshal(*rawMsg, &passportElementTypeAddress)
		return &passportElementTypeAddress, err

	case PassportElementTypeUtilityBillType:
		var passportElementTypeUtilityBill PassportElementTypeUtilityBill
		err := json.Unmarshal(*rawMsg, &passportElementTypeUtilityBill)
		return &passportElementTypeUtilityBill, err

	case PassportElementTypeBankStatementType:
		var passportElementTypeBankStatement PassportElementTypeBankStatement
		err := json.Unmarshal(*rawMsg, &passportElementTypeBankStatement)
		return &passportElementTypeBankStatement, err

	case PassportElementTypeRentalAgreementType:
		var passportElementTypeRentalAgreement PassportElementTypeRentalAgreement
		err := json.Unmarshal(*rawMsg, &passportElementTypeRentalAgreement)
		return &passportElementTypeRentalAgreement, err

	case PassportElementTypePassportRegistrationType:
		var passportElementTypePassportRegistration PassportElementTypePassportRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTypePassportRegistration)
		return &passportElementTypePassportRegistration, err

	case PassportElementTypeTemporaryRegistrationType:
		var passportElementTypeTemporaryRegistration PassportElementTypeTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTypeTemporaryRegistration)
		return &passportElementTypeTemporaryRegistration, err

	case PassportElementTypePhoneNumberType:
		var passportElementTypePhoneNumber PassportElementTypePhoneNumber
		err := json.Unmarshal(*rawMsg, &passportElementTypePhoneNumber)
		return &passportElementTypePhoneNumber, err

	case PassportElementTypeEmailAddressType:
		var passportElementTypeEmailAddress PassportElementTypeEmailAddress
		err := json.Unmarshal(*rawMsg, &passportElementTypeEmailAddress)
		return &passportElementTypeEmailAddress, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalPassportElement(rawMsg *json.RawMessage) (PassportElement, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementEnum(objMap["@type"].(string)) {
	case PassportElementPersonalDetailsType:
		var passportElementPersonalDetails PassportElementPersonalDetails
		err := json.Unmarshal(*rawMsg, &passportElementPersonalDetails)
		return &passportElementPersonalDetails, err

	case PassportElementPassportType:
		var passportElementPassport PassportElementPassport
		err := json.Unmarshal(*rawMsg, &passportElementPassport)
		return &passportElementPassport, err

	case PassportElementDriverLicenseType:
		var passportElementDriverLicense PassportElementDriverLicense
		err := json.Unmarshal(*rawMsg, &passportElementDriverLicense)
		return &passportElementDriverLicense, err

	case PassportElementIDentityCardType:
		var passportElementIDentityCard PassportElementIDentityCard
		err := json.Unmarshal(*rawMsg, &passportElementIDentityCard)
		return &passportElementIDentityCard, err

	case PassportElementInternalPassportType:
		var passportElementInternalPassport PassportElementInternalPassport
		err := json.Unmarshal(*rawMsg, &passportElementInternalPassport)
		return &passportElementInternalPassport, err

	case PassportElementAddressType:
		var passportElementAddress PassportElementAddress
		err := json.Unmarshal(*rawMsg, &passportElementAddress)
		return &passportElementAddress, err

	case PassportElementUtilityBillType:
		var passportElementUtilityBill PassportElementUtilityBill
		err := json.Unmarshal(*rawMsg, &passportElementUtilityBill)
		return &passportElementUtilityBill, err

	case PassportElementBankStatementType:
		var passportElementBankStatement PassportElementBankStatement
		err := json.Unmarshal(*rawMsg, &passportElementBankStatement)
		return &passportElementBankStatement, err

	case PassportElementRentalAgreementType:
		var passportElementRentalAgreement PassportElementRentalAgreement
		err := json.Unmarshal(*rawMsg, &passportElementRentalAgreement)
		return &passportElementRentalAgreement, err

	case PassportElementPassportRegistrationType:
		var passportElementPassportRegistration PassportElementPassportRegistration
		err := json.Unmarshal(*rawMsg, &passportElementPassportRegistration)
		return &passportElementPassportRegistration, err

	case PassportElementTemporaryRegistrationType:
		var passportElementTemporaryRegistration PassportElementTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTemporaryRegistration)
		return &passportElementTemporaryRegistration, err

	case PassportElementPhoneNumberType:
		var passportElementPhoneNumber PassportElementPhoneNumber
		err := json.Unmarshal(*rawMsg, &passportElementPhoneNumber)
		return &passportElementPhoneNumber, err

	case PassportElementEmailAddressType:
		var passportElementEmailAddress PassportElementEmailAddress
		err := json.Unmarshal(*rawMsg, &passportElementEmailAddress)
		return &passportElementEmailAddress, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputPassportElement(rawMsg *json.RawMessage) (InputPassportElement, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputPassportElementEnum(objMap["@type"].(string)) {
	case InputPassportElementPersonalDetailsType:
		var inputPassportElementPersonalDetails InputPassportElementPersonalDetails
		err := json.Unmarshal(*rawMsg, &inputPassportElementPersonalDetails)
		return &inputPassportElementPersonalDetails, err

	case InputPassportElementPassportType:
		var inputPassportElementPassport InputPassportElementPassport
		err := json.Unmarshal(*rawMsg, &inputPassportElementPassport)
		return &inputPassportElementPassport, err

	case InputPassportElementDriverLicenseType:
		var inputPassportElementDriverLicense InputPassportElementDriverLicense
		err := json.Unmarshal(*rawMsg, &inputPassportElementDriverLicense)
		return &inputPassportElementDriverLicense, err

	case InputPassportElementIDentityCardType:
		var inputPassportElementIDentityCard InputPassportElementIDentityCard
		err := json.Unmarshal(*rawMsg, &inputPassportElementIDentityCard)
		return &inputPassportElementIDentityCard, err

	case InputPassportElementInternalPassportType:
		var inputPassportElementInternalPassport InputPassportElementInternalPassport
		err := json.Unmarshal(*rawMsg, &inputPassportElementInternalPassport)
		return &inputPassportElementInternalPassport, err

	case InputPassportElementAddressType:
		var inputPassportElementAddress InputPassportElementAddress
		err := json.Unmarshal(*rawMsg, &inputPassportElementAddress)
		return &inputPassportElementAddress, err

	case InputPassportElementUtilityBillType:
		var inputPassportElementUtilityBill InputPassportElementUtilityBill
		err := json.Unmarshal(*rawMsg, &inputPassportElementUtilityBill)
		return &inputPassportElementUtilityBill, err

	case InputPassportElementBankStatementType:
		var inputPassportElementBankStatement InputPassportElementBankStatement
		err := json.Unmarshal(*rawMsg, &inputPassportElementBankStatement)
		return &inputPassportElementBankStatement, err

	case InputPassportElementRentalAgreementType:
		var inputPassportElementRentalAgreement InputPassportElementRentalAgreement
		err := json.Unmarshal(*rawMsg, &inputPassportElementRentalAgreement)
		return &inputPassportElementRentalAgreement, err

	case InputPassportElementPassportRegistrationType:
		var inputPassportElementPassportRegistration InputPassportElementPassportRegistration
		err := json.Unmarshal(*rawMsg, &inputPassportElementPassportRegistration)
		return &inputPassportElementPassportRegistration, err

	case InputPassportElementTemporaryRegistrationType:
		var inputPassportElementTemporaryRegistration InputPassportElementTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &inputPassportElementTemporaryRegistration)
		return &inputPassportElementTemporaryRegistration, err

	case InputPassportElementPhoneNumberType:
		var inputPassportElementPhoneNumber InputPassportElementPhoneNumber
		err := json.Unmarshal(*rawMsg, &inputPassportElementPhoneNumber)
		return &inputPassportElementPhoneNumber, err

	case InputPassportElementEmailAddressType:
		var inputPassportElementEmailAddress InputPassportElementEmailAddress
		err := json.Unmarshal(*rawMsg, &inputPassportElementEmailAddress)
		return &inputPassportElementEmailAddress, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalPassportElementErrorSource(rawMsg *json.RawMessage) (PassportElementErrorSource, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementErrorSourceEnum(objMap["@type"].(string)) {
	case PassportElementErrorSourceUnspecifiedType:
		var passportElementErrorSourceUnspecified PassportElementErrorSourceUnspecified
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceUnspecified)
		return &passportElementErrorSourceUnspecified, err

	case PassportElementErrorSourceDataFieldType:
		var passportElementErrorSourceDataField PassportElementErrorSourceDataField
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceDataField)
		return &passportElementErrorSourceDataField, err

	case PassportElementErrorSourceFrontSideType:
		var passportElementErrorSourceFrontSide PassportElementErrorSourceFrontSide
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFrontSide)
		return &passportElementErrorSourceFrontSide, err

	case PassportElementErrorSourceReverseSideType:
		var passportElementErrorSourceReverseSide PassportElementErrorSourceReverseSide
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceReverseSide)
		return &passportElementErrorSourceReverseSide, err

	case PassportElementErrorSourceSelfieType:
		var passportElementErrorSourceSelfie PassportElementErrorSourceSelfie
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceSelfie)
		return &passportElementErrorSourceSelfie, err

	case PassportElementErrorSourceTranslationFileType:
		var passportElementErrorSourceTranslationFile PassportElementErrorSourceTranslationFile
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceTranslationFile)
		return &passportElementErrorSourceTranslationFile, err

	case PassportElementErrorSourceTranslationFilesType:
		var passportElementErrorSourceTranslationFiles PassportElementErrorSourceTranslationFiles
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceTranslationFiles)
		return &passportElementErrorSourceTranslationFiles, err

	case PassportElementErrorSourceFileType:
		var passportElementErrorSourceFile PassportElementErrorSourceFile
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFile)
		return &passportElementErrorSourceFile, err

	case PassportElementErrorSourceFilesType:
		var passportElementErrorSourceFiles PassportElementErrorSourceFiles
		err := json.Unmarshal(*rawMsg, &passportElementErrorSourceFiles)
		return &passportElementErrorSourceFiles, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}

func unmarshalInputPassportElementErrorSource(rawMsg *json.RawMessage) (InputPassportElementErrorSource, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputPassportElementErrorSourceEnum(objMap["@type"].(string)) {
	case InputPassportElementErrorSourceUnspecifiedType:
		var inputPassportElementErrorSourceUnspecified InputPassportElementErrorSourceUnspecified
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceUnspecified)
		return &inputPassportElementErrorSourceUnspecified, err

	case InputPassportElementErrorSourceDataFieldType:
		var inputPassportElementErrorSourceDataField InputPassportElementErrorSourceDataField
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceDataField)
		return &inputPassportElementErrorSourceDataField, err

	case InputPassportElementErrorSourceFrontSideType:
		var inputPassportElementErrorSourceFrontSide InputPassportElementErrorSourceFrontSide
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFrontSide)
		return &inputPassportElementErrorSourceFrontSide, err

	case InputPassportElementErrorSourceReverseSideType:
		var inputPassportElementErrorSourceReverseSide InputPassportElementErrorSourceReverseSide
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceReverseSide)
		return &inputPassportElementErrorSourceReverseSide, err

	case InputPassportElementErrorSourceSelfieType:
		var inputPassportElementErrorSourceSelfie InputPassportElementErrorSourceSelfie
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceSelfie)
		return &inputPassportElementErrorSourceSelfie, err

	case InputPassportElementErrorSourceTranslationFileType:
		var inputPassportElementErrorSourceTranslationFile InputPassportElementErrorSourceTranslationFile
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceTranslationFile)
		return &inputPassportElementErrorSourceTranslationFile, err

	case InputPassportElementErrorSourceTranslationFilesType:
		var inputPassportElementErrorSourceTranslationFiles InputPassportElementErrorSourceTranslationFiles
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceTranslationFiles)
		return &inputPassportElementErrorSourceTranslationFiles, err

	case InputPassportElementErrorSourceFileType:
		var inputPassportElementErrorSourceFile InputPassportElementErrorSourceFile
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFile)
		return &inputPassportElementErrorSourceFile, err

	case InputPassportElementErrorSourceFilesType:
		var inputPassportElementErrorSourceFiles InputPassportElementErrorSourceFiles
		err := json.Unmarshal(*rawMsg, &inputPassportElementErrorSourceFiles)
		return &inputPassportElementErrorSourceFiles, err

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

	case MessagePassportDataSentType:
		var messagePassportDataSent MessagePassportDataSent
		err := json.Unmarshal(*rawMsg, &messagePassportDataSent)
		return &messagePassportDataSent, err

	case MessagePassportDataReceivedType:
		var messagePassportDataReceived MessagePassportDataReceived
		err := json.Unmarshal(*rawMsg, &messagePassportDataReceived)
		return &messagePassportDataReceived, err

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

func unmarshalLanguagePackStringValue(rawMsg *json.RawMessage) (LanguagePackStringValue, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LanguagePackStringValueEnum(objMap["@type"].(string)) {
	case LanguagePackStringValueOrdinaryType:
		var languagePackStringValueOrdinary LanguagePackStringValueOrdinary
		err := json.Unmarshal(*rawMsg, &languagePackStringValueOrdinary)
		return &languagePackStringValueOrdinary, err

	case LanguagePackStringValuePluralizedType:
		var languagePackStringValuePluralized LanguagePackStringValuePluralized
		err := json.Unmarshal(*rawMsg, &languagePackStringValuePluralized)
		return &languagePackStringValuePluralized, err

	case LanguagePackStringValueDeletedType:
		var languagePackStringValueDeleted LanguagePackStringValueDeleted
		err := json.Unmarshal(*rawMsg, &languagePackStringValueDeleted)
		return &languagePackStringValueDeleted, err

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

	case ChatReportReasonCopyrightType:
		var chatReportReasonCopyright ChatReportReasonCopyright
		err := json.Unmarshal(*rawMsg, &chatReportReasonCopyright)
		return &chatReportReasonCopyright, err

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

	case FileTypeSecretThumbnailType:
		var fileTypeSecretThumbnail FileTypeSecretThumbnail
		err := json.Unmarshal(*rawMsg, &fileTypeSecretThumbnail)
		return &fileTypeSecretThumbnail, err

	case FileTypeSecureType:
		var fileTypeSecure FileTypeSecure
		err := json.Unmarshal(*rawMsg, &fileTypeSecure)
		return &fileTypeSecure, err

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

func unmarshalProxyType(rawMsg *json.RawMessage) (ProxyType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ProxyTypeEnum(objMap["@type"].(string)) {
	case ProxyTypeSocks5Type:
		var proxyTypeSocks5 ProxyTypeSocks5
		err := json.Unmarshal(*rawMsg, &proxyTypeSocks5)
		return &proxyTypeSocks5, err

	case ProxyTypeHttpType:
		var proxyTypeHttp ProxyTypeHttp
		err := json.Unmarshal(*rawMsg, &proxyTypeHttp)
		return &proxyTypeHttp, err

	case ProxyTypeMtprotoType:
		var proxyTypeMtproto ProxyTypeMtproto
		err := json.Unmarshal(*rawMsg, &proxyTypeMtproto)
		return &proxyTypeMtproto, err

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

	case UpdateChatIsMarkedAsUnreadType:
		var updateChatIsMarkedAsUnread UpdateChatIsMarkedAsUnread
		err := json.Unmarshal(*rawMsg, &updateChatIsMarkedAsUnread)
		return &updateChatIsMarkedAsUnread, err

	case UpdateChatIsSponsoredType:
		var updateChatIsSponsored UpdateChatIsSponsored
		err := json.Unmarshal(*rawMsg, &updateChatIsSponsored)
		return &updateChatIsSponsored, err

	case UpdateChatDefaultDisableNotificationType:
		var updateChatDefaultDisableNotification UpdateChatDefaultDisableNotification
		err := json.Unmarshal(*rawMsg, &updateChatDefaultDisableNotification)
		return &updateChatDefaultDisableNotification, err

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

	case UpdateChatNotificationSettingsType:
		var updateChatNotificationSettings UpdateChatNotificationSettings
		err := json.Unmarshal(*rawMsg, &updateChatNotificationSettings)
		return &updateChatNotificationSettings, err

	case UpdateScopeNotificationSettingsType:
		var updateScopeNotificationSettings UpdateScopeNotificationSettings
		err := json.Unmarshal(*rawMsg, &updateScopeNotificationSettings)
		return &updateScopeNotificationSettings, err

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

	case UpdateUnreadChatCountType:
		var updateUnreadChatCount UpdateUnreadChatCount
		err := json.Unmarshal(*rawMsg, &updateUnreadChatCount)
		return &updateUnreadChatCount, err

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

	case UpdateLanguagePackStringsType:
		var updateLanguagePackStrings UpdateLanguagePackStrings
		err := json.Unmarshal(*rawMsg, &updateLanguagePackStrings)
		return &updateLanguagePackStrings, err

	case UpdateConnectionStateType:
		var updateConnectionState UpdateConnectionState
		err := json.Unmarshal(*rawMsg, &updateConnectionState)
		return &updateConnectionState, err

	case UpdateTermsOfServiceType:
		var updateTermsOfService UpdateTermsOfService
		err := json.Unmarshal(*rawMsg, &updateTermsOfService)
		return &updateTermsOfService, err

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
