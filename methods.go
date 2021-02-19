package tdlib

import (
	"encoding/json"
	"fmt"
)

// GetAuthorizationState Returns the current authorization state; this is an offline request. For informational purposes only. Use updateAuthorizationState instead to maintain the current authorization state. Can be called before initialization
func (client *Client) GetAuthorizationState() (AuthorizationState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAuthorizationState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch AuthorizationStateEnum(result.Data["@type"].(string)) {

	case AuthorizationStateWaitTdlibParametersType:
		var authorizationState AuthorizationStateWaitTdlibParameters
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitEncryptionKeyType:
		var authorizationState AuthorizationStateWaitEncryptionKey
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPhoneNumberType:
		var authorizationState AuthorizationStateWaitPhoneNumber
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitCodeType:
		var authorizationState AuthorizationStateWaitCode
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitOtherDeviceConfirmationType:
		var authorizationState AuthorizationStateWaitOtherDeviceConfirmation
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitRegistrationType:
		var authorizationState AuthorizationStateWaitRegistration
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPasswordType:
		var authorizationState AuthorizationStateWaitPassword
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateReadyType:
		var authorizationState AuthorizationStateReady
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateLoggingOutType:
		var authorizationState AuthorizationStateLoggingOut
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosingType:
		var authorizationState AuthorizationStateClosing
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosedType:
		var authorizationState AuthorizationStateClosed
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
// @param parameters Parameters
func (client *Client) SetTdlibParameters(parameters *TdlibParameters) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setTdlibParameters",
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckDatabaseEncryptionKey Checks the database encryption key for correctness. Works only when the current authorization state is authorizationStateWaitEncryptionKey
// @param encryptionKey Encryption key to check or set up
func (client *Client) CheckDatabaseEncryptionKey(encryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "checkDatabaseEncryptionKey",
		"encryption_key": encryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number
func (client *Client) SetAuthenticationPhoneNumber(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setAuthenticationPhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResendAuthenticationCode Re-sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode and the next_code_type of the result is not null
func (client *Client) ResendAuthenticationCode() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendAuthenticationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode
// @param code The verification code received via SMS, Telegram message, phone call, or flash call
func (client *Client) CheckAuthenticationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RequestQrCodeAuthentication Requests QR code authentication by scanning a QR code on another logged in device. Works only when the current authorization state is authorizationStateWaitPhoneNumber,
// @param otherUserIDs List of user identifiers of other users currently using the application
func (client *Client) RequestQrCodeAuthentication(otherUserIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "requestQrCodeAuthentication",
		"other_user_ids": otherUserIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RegisterUser Finishes user registration. Works only when the current authorization state is authorizationStateWaitRegistration
// @param firstName The first name of the user; 1-64 characters
// @param lastName The last name of the user; 0-64 characters
func (client *Client) RegisterUser(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "registerUser",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationPassword Checks the authentication password for correctness. Works only when the current authorization state is authorizationStateWaitPassword
// @param password The password to check
func (client *Client) CheckAuthenticationPassword(password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RequestAuthenticationPasswordRecovery Requests to send a password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (client *Client) RequestAuthenticationPasswordRecovery() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestAuthenticationPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RecoverAuthenticationPassword Recovers the password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
func (client *Client) RecoverAuthenticationPassword(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverAuthenticationPassword",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in
// @param token The bot token
func (client *Client) CheckAuthenticationBotToken(token string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationBotToken",
		"token": token,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) LogOut() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "logOut",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent. Can be called before initialization
func (client *Client) Close() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "close",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed. After the destruction completes updateAuthorizationState with authorizationStateClosed will be sent. Can be called before authorization
func (client *Client) Destroy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "destroy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success
// @param link A link from a QR code. The link must be scanned by the in-app camera
func (client *Client) ConfirmQrCodeAuthentication(link string) (*Session, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "confirmQrCodeAuthentication",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var session Session
	err = json.Unmarshal(result.Raw, &session)
	return &session, err

}

// GetCurrentState Returns all updates needed to restore current TDLib state, i.e. all actual UpdateAuthorizationState/UpdateUser/UpdateNewChat and others. This is especially useful if TDLib is run in a separate process. Can be called before initialization
func (client *Client) GetCurrentState() (*Updates, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCurrentState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var updates Updates
	err = json.Unmarshal(result.Raw, &updates)
	return &updates, err

}

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain
// @param newEncryptionKey New encryption key
func (client *Client) SetDatabaseEncryptionKey(newEncryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setDatabaseEncryptionKey",
		"new_encryption_key": newEncryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPasswordState Returns the current state of 2-step verification
func (client *Client) GetPasswordState() (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// SetPassword Changes the password for the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed
// @param oldPassword Previous password of the user
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
// @param setRecoveryEmailAddress Pass true if the recovery email address should be changed
// @param newRecoveryEmailAddress New recovery email address; may be empty
func (client *Client) SetPassword(oldPassword string, newPassword string, newHint string, setRecoveryEmailAddress bool, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setPassword",
		"old_password":               oldPassword,
		"new_password":               newPassword,
		"new_hint":                   newHint,
		"set_recovery_email_address": setRecoveryEmailAddress,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// GetRecoveryEmailAddress Returns a 2-step verification recovery email address that was previously set up. This method can be used to verify a password provided by the user
// @param password The password for the current user
func (client *Client) GetRecoveryEmailAddress(password string) (*RecoveryEmailAddress, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecoveryEmailAddress",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var recoveryEmailAddress RecoveryEmailAddress
	err = json.Unmarshal(result.Raw, &recoveryEmailAddress)
	return &recoveryEmailAddress, err

}

// SetRecoveryEmailAddress Changes the 2-step verification recovery email address of the user. If a new recovery email address is specified, then the change will not be applied until the new recovery email address is confirmed.
// @param password
// @param newRecoveryEmailAddress
func (client *Client) SetRecoveryEmailAddress(password string, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setRecoveryEmailAddress",
		"password":                   password,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// CheckRecoveryEmailAddressCode Checks the 2-step verification recovery email address verification code
// @param code Verification code
func (client *Client) CheckRecoveryEmailAddressCode(code string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkRecoveryEmailAddressCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// ResendRecoveryEmailAddressCode Resends the 2-step verification recovery email address verification code
func (client *Client) ResendRecoveryEmailAddressCode() (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendRecoveryEmailAddressCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// RequestPasswordRecovery Requests to send a password recovery code to an email address that was previously set up
func (client *Client) RequestPasswordRecovery() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// RecoverPassword Recovers the password using a recovery code sent to an email address that was previously set up
// @param recoveryCode Recovery code to check
func (client *Client) RecoverPassword(recoveryCode string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverPassword",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// CreateTemporaryPassword Creates a new temporary password for processing payments
// @param password Persistent user password
// @param validFor Time during which the temporary password will be valid, in seconds; should be between 60 and 86400
func (client *Client) CreateTemporaryPassword(password string, validFor int32) (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "createTemporaryPassword",
		"password":  password,
		"valid_for": validFor,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}

// GetTemporaryPasswordState Returns information about the current temporary password
func (client *Client) GetTemporaryPasswordState() (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTemporaryPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}

// GetMe Returns the current user
func (client *Client) GetMe() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMe",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

}

// GetUser Returns information about a user by their identifier. This is an offline request if the current user is not a bot
// @param userID User identifier
func (client *Client) GetUser(userID int32) (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUser",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userDummy User
	err = json.Unmarshal(result.Raw, &userDummy)
	return &userDummy, err

}

// GetUserFullInfo Returns full information about a user by their identifier
// @param userID User identifier
func (client *Client) GetUserFullInfo(userID int32) (*UserFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserFullInfo",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userFullInfo UserFullInfo
	err = json.Unmarshal(result.Raw, &userFullInfo)
	return &userFullInfo, err

}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline request if the current user is not a bot
// @param basicGroupID Basic group identifier
func (client *Client) GetBasicGroup(basicGroupID int32) (*BasicGroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroup",
		"basic_group_id": basicGroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupDummy BasicGroup
	err = json.Unmarshal(result.Raw, &basicGroupDummy)
	return &basicGroupDummy, err

}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier
// @param basicGroupID Basic group identifier
func (client *Client) GetBasicGroupFullInfo(basicGroupID int32) (*BasicGroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroupFullInfo",
		"basic_group_id": basicGroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupFullInfo BasicGroupFullInfo
	err = json.Unmarshal(result.Raw, &basicGroupFullInfo)
	return &basicGroupFullInfo, err

}

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline request if the current user is not a bot
// @param supergroupID Supergroup or channel identifier
func (client *Client) GetSupergroup(supergroupID int32) (*Supergroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroup",
		"supergroup_id": supergroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupDummy Supergroup
	err = json.Unmarshal(result.Raw, &supergroupDummy)
	return &supergroupDummy, err

}

// GetSupergroupFullInfo Returns full information about a supergroup or a channel by its identifier, cached for up to 1 minute
// @param supergroupID Supergroup or channel identifier
func (client *Client) GetSupergroupFullInfo(supergroupID int32) (*SupergroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupFullInfo",
		"supergroup_id": supergroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupFullInfo SupergroupFullInfo
	err = json.Unmarshal(result.Raw, &supergroupFullInfo)
	return &supergroupFullInfo, err

}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline request
// @param secretChatID Secret chat identifier
func (client *Client) GetSecretChat(secretChatID int32) (*SecretChat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getSecretChat",
		"secret_chat_id": secretChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var secretChatDummy SecretChat
	err = json.Unmarshal(result.Raw, &secretChatDummy)
	return &secretChatDummy, err

}

// GetChat Returns information about a chat by its identifier, this is an offline request if the current user is not a bot
// @param chatID Chat identifier
func (client *Client) GetChat(chatID int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// GetMessage Returns information about a message
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message to get
func (client *Client) GetMessage(chatID int64, messageID int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessage",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetMessageLocally Returns information about a message, if it is available locally without sending network request. This is an offline request
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message to get
func (client *Client) GetMessageLocally(chatID int64, messageID int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageLocally",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetRepliedMessage Returns information about a message that is replied by a given message. Also returns the pinned message, the game message, and the invoice message for messages of the types messagePinMessage, messageGameScore, and messagePaymentSuccessful respectively
// @param chatID Identifier of the chat the message belongs to
// @param messageID Identifier of the message reply to which to get
func (client *Client) GetRepliedMessage(chatID int64, messageID int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getRepliedMessage",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetChatPinnedMessage Returns information about a newest pinned message in the chat
// @param chatID Identifier of the chat the message belongs to
func (client *Client) GetChatPinnedMessage(chatID int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatPinnedMessage",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetCallbackQueryMessage Returns information about a message with the callback button that originated a callback query; for bots only
// @param chatID Identifier of the chat the message belongs to
// @param messageID Message identifier
// @param callbackQueryID Identifier of the callback query
func (client *Client) GetCallbackQueryMessage(chatID int64, messageID int64, callbackQueryID JSONInt64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getCallbackQueryMessage",
		"chat_id":           chatID,
		"message_id":        messageID,
		"callback_query_id": callbackQueryID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result
// @param chatID Identifier of the chat the messages belong to
// @param messageIDs Identifiers of the messages to get
func (client *Client) GetMessages(chatID int64, messageIDs []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessages",
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessageThread Returns information about a message thread. Can be used only if message.can_get_message_thread == true
// @param chatID Chat identifier
// @param messageID Identifier of the message
func (client *Client) GetMessageThread(chatID int64, messageID int64) (*MessageThreadInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageThread",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageThreadInfo MessageThreadInfo
	err = json.Unmarshal(result.Raw, &messageThreadInfo)
	return &messageThreadInfo, err

}

// GetFile Returns information about a file; this is an offline request
// @param fileID Identifier of the file to get
func (client *Client) GetFile(fileID int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFile",
		"file_id": fileID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetRemoteFile Returns information about a file by its remote ID; this is an offline request. Can be used to register a URL as a file for further uploading, or sending as a message. Even the request succeeds, the file can be used only if it is still accessible to the user.
// @param remoteFileID Remote identifier of the file to get
// @param fileType File type, if known
func (client *Client) GetRemoteFile(remoteFileID string, fileType FileType) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getRemoteFile",
		"remote_file_id": remoteFileID,
		"file_type":      fileType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetChats Returns an ordered list of chats in a chat list. Chats are sorted by the pair (chat.position.order, chat.id) in descending order. (For example, to get a list of chats from the beginning, the offset_order should be equal to a biggest signed 64-bit number 9223372036854775807 == 2^63 - 1).
// @param chatList The chat list in which to return chats
// @param offsetOrder Chat order to return chats from
// @param offsetChatID Chat identifier to return chats from
// @param limit The maximum number of chats to be returned. It is possible that fewer chats than the limit are returned even if the end of the list is not reached
func (client *Client) GetChats(chatList ChatList, offsetOrder JSONInt64, offsetChatID int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChats",
		"chat_list":      chatList,
		"offset_order":   offsetOrder,
		"offset_chat_id": offsetChatID,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchPublicChat Searches a public chat by its username. Currently only private chats, supergroups and channels can be public. Returns the chat if found; otherwise an error is returned
// @param username Username to be resolved
func (client *Client) SearchPublicChat(username string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchPublicChat",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently only private chats, supergroups and channels can be public. Returns a meaningful number of results. Returns nothing if the length of the searched username prefix is less than 5. Excludes private chats with contacts and chats from the chat list from the results
// @param query Query to search for
func (client *Client) SearchPublicChats(query string) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchPublicChats",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChats Searches for the specified query in the title and username of already known chats, this is an offline request. Returns chats in the order seen in the main chat list
// @param query Query to search for. If the query is empty, returns up to 20 recently found chats
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChats(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChats",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the main chat list
// @param query Query to search for
// @param limit The maximum number of chats to be returned
func (client *Client) SearchChatsOnServer(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChatsOnServer",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsNearby Returns a list of users and location-based supergroups nearby. The list of users nearby will be updated for 60 seconds after the request by the updates updateUsersNearby. The request should be sent again every 25 seconds with adjusted location to not miss new chats
// @param location Current user location
func (client *Client) SearchChatsNearby(location *Location) (*ChatsNearby, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchChatsNearby",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatsNearby ChatsNearby
	err = json.Unmarshal(result.Raw, &chatsNearby)
	return &chatsNearby, err

}

// GetTopChats Returns a list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of chats to be returned
// @param limit The maximum number of chats to be returned; up to 30
func (client *Client) GetTopChats(category TopChatCategory, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getTopChats",
		"category": category,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of frequently used chats
// @param chatID Chat identifier
func (client *Client) RemoveTopChat(category TopChatCategory, chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeTopChat",
		"category": category,
		"chat_id":  chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first
// @param chatID Identifier of the chat to add
func (client *Client) AddRecentlyFoundChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addRecentlyFoundChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats
// @param chatID Identifier of the chat to be removed
func (client *Client) RemoveRecentlyFoundChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentlyFoundChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (client *Client) ClearRecentlyFoundChats() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearRecentlyFoundChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckChatUsername Checks whether a username can be set for a chat
// @param chatID Chat identifier; should be identifier of a supergroup chat, or a channel chat, or a private chat with self, or zero if chat is being created
// @param username Username to be checked
func (client *Client) CheckChatUsername(chatID int64, username string) (CheckChatUsernameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkChatUsername",
		"chat_id":  chatID,
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckChatUsernameResultEnum(result.Data["@type"].(string)) {

	case CheckChatUsernameResultOkType:
		var checkChatUsernameResult CheckChatUsernameResultOk
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameInvalid
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameOccupied
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResult CheckChatUsernameResultPublicChatsTooMuch
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResult CheckChatUsernameResultPublicGroupsUnavailable
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetCreatedPublicChats Returns a list of public chats of the specified type, owned by the user
// @param typeParam Type of the public chats to return
func (client *Client) GetCreatedPublicChats(typeParam PublicChatType) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCreatedPublicChats",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// CheckCreatedPublicChatsLimit Checks whether the maximum number of owned public chats has been reached. Returns corresponding error if the limit was reached
// @param typeParam Type of the public chats, for which to check the limit
func (client *Client) CheckCreatedPublicChatsLimit(typeParam PublicChatType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkCreatedPublicChatsLimit",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSuitableDiscussionChats Returns a list of basic group and supergroup chats, which can be used as a discussion group for a channel. Returned basic group chats must be first upgraded to supergroups before they can be set as a discussion group. To set a returned supergroup as a discussion group, access to its old messages must be enabled using toggleSupergroupIsAllHistoryAvailable first
func (client *Client) GetSuitableDiscussionChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSuitableDiscussionChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetInactiveSupergroupChats Returns a list of recently inactive supergroups and channels. Can be used when user reaches limit on the number of joined supergroups and channels and receives CHANNELS_TOO_MUCH error
func (client *Client) GetInactiveSupergroupChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInactiveSupergroupChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetGroupsInCommon Returns a list of common group chats with a given user. Chats are sorted by their type and creation date
// @param userID User identifier
// @param offsetChatID Chat identifier starting from which to return chats; use 0 for the first request
// @param limit The maximum number of chats to be returned; up to 100
func (client *Client) GetGroupsInCommon(userID int32, offsetChatID int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getGroupsInCommon",
		"user_id":        userID,
		"offset_chat_id": offsetChatID,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetChatHistory Returns messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id).
// @param chatID Chat identifier
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyLocal If true, returns only messages that are available locally without sending network requests
func (client *Client) GetChatHistory(chatID int64, fromMessageID int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatHistory",
		"chat_id":         chatID,
		"from_message_id": fromMessageID,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessageThreadHistory Returns messages in a message thread of a message. Can be used only if message.can_get_message_thread == true. Message thread of a channel message is in the channel's linked supergroup.
// @param chatID Chat identifier
// @param messageID Message identifier, which thread history needs to be returned
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset up to 99 to get additionally some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than or equal to -offset. Fewer messages may be returned than specified by the limit, even if the end of the message thread history has not been reached
func (client *Client) GetMessageThreadHistory(chatID int64, messageID int64, fromMessageID int64, offset int32, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getMessageThreadHistory",
		"chat_id":         chatID,
		"message_id":      messageID,
		"from_message_id": fromMessageID,
		"offset":          offset,
		"limit":           limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// DeleteChatHistory Deletes all messages in the chat. Use Chat.can_be_deleted_only_for_self and Chat.can_be_deleted_for_all_users fields to find whether and how the method can be applied to the chat
// @param chatID Chat identifier
// @param removeFromChatList Pass true if the chat should be removed from the chat list
// @param revoke Pass true to try to delete chat history for all users
func (client *Client) DeleteChatHistory(chatID int64, removeFromChatList bool, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "deleteChatHistory",
		"chat_id":               chatID,
		"remove_from_chat_list": removeFromChatList,
		"revoke":                revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// DeleteChat Deletes a chat along with all messages in the corresponding chat for all chat members; requires owner privileges. For group chats this will release the username and remove all members. Chats with more than 1000 members can't be deleted using this method
// @param chatID Chat identifier
func (client *Client) DeleteChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query
// @param chatID Identifier of the chat in which to search messages
// @param query Query to search for
// @param sender If not null, only messages sent by the specified sender will be returned. Not supported in secret chats
// @param fromMessageID Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter Filter for message content in the search results
// @param messageThreadID If not 0, only messages in the specified thread will be returned; supergroups only
func (client *Client) SearchChatMessages(chatID int64, query string, sender MessageSender, fromMessageID int64, offset int32, limit int32, filter SearchMessagesFilter, messageThreadID int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchChatMessages",
		"chat_id":           chatID,
		"query":             query,
		"sender":            sender,
		"from_message_id":   fromMessageID,
		"offset":            offset,
		"limit":             limit,
		"filter":            filter,
		"message_thread_id": messageThreadID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)).
// @param chatList Chat list in which to search messages; pass null to search in all chats regardless of their chat list
// @param query Query to search for
// @param offsetDate The date of the message starting from which the results should be fetched. Use 0 or any date in the future to get results from the last message
// @param offsetChatID The chat identifier of the last found message, or 0 for the first request
// @param offsetMessageID The message identifier of the last found message, or 0 for the first request
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter Filter for message content in the search results; searchMessagesFilterCall, searchMessagesFilterMissedCall, searchMessagesFilterMention, searchMessagesFilterUnreadMention, searchMessagesFilterFailedToSend and searchMessagesFilterPinned are unsupported in this function
// @param minDate If not 0, the minimum date of the messages to return
// @param maxDate If not 0, the maximum date of the messages to return
func (client *Client) SearchMessages(chatList ChatList, query string, offsetDate int32, offsetChatID int64, offsetMessageID int64, limit int32, filter SearchMessagesFilter, minDate int32, maxDate int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchMessages",
		"chat_list":         chatList,
		"query":             query,
		"offset_date":       offsetDate,
		"offset_chat_id":    offsetChatID,
		"offset_message_id": offsetMessageID,
		"limit":             limit,
		"filter":            filter,
		"min_date":          minDate,
		"max_date":          maxDate,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance the number of returned messages is chosen by the library
// @param chatID Identifier of the chat in which to search. Specify 0 to search in all secret chats
// @param query Query to search for. If empty, searchChatMessages should be used instead
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter A filter for message content in the search results
func (client *Client) SearchSecretMessages(chatID int64, query string, offset string, limit int32, filter SearchMessagesFilter) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchSecretMessages",
		"chat_id": chatID,
		"query":   query,
		"offset":  offset,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err

}

// SearchCallMessages Searches for call messages. Returns the results in reverse chronological order (i. e., in order of decreasing message_id). For optimal performance the number of returned messages is chosen by the library
// @param fromMessageID Identifier of the message from which to search; use 0 to get results from the last message
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyMissed If true, returns only messages with missed calls
func (client *Client) SearchCallMessages(fromMessageID int64, limit int32, onlyMissed bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchCallMessages",
		"from_message_id": fromMessageID,
		"limit":           limit,
		"only_missed":     onlyMissed,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// DeleteAllCallMessages Deletes all call messages
// @param revoke Pass true to delete the messages for all users
func (client *Client) DeleteAllCallMessages(revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAllCallMessages",
		"revoke": revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user
// @param chatID Chat identifier
// @param limit The maximum number of messages to be returned
func (client *Client) SearchChatRecentLocationMessages(chatID int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatRecentLocationMessages",
		"chat_id": chatID,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetActiveLiveLocationMessages Returns all active live locations that should be updated by the application. The list is persistent across application restarts only if the message database is used
func (client *Client) GetActiveLiveLocationMessages() (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveLiveLocationMessages",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date
// @param chatID Chat identifier
// @param date Point in time (Unix timestamp) relative to which to search for messages
func (client *Client) GetChatMessageByDate(chatID int64, date int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMessageByDate",
		"chat_id": chatID,
		"date":    date,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat
// @param chatID Identifier of the chat in which to count messages
// @param filter Filter for message content; searchMessagesFilterEmpty is unsupported in this function
// @param returnLocal If true, returns count that is available locally without sending network requests, returning -1 if the number of messages is unknown
func (client *Client) GetChatMessageCount(chatID int64, filter SearchMessagesFilter, returnLocal bool) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getChatMessageCount",
		"chat_id":      chatID,
		"filter":       filter,
		"return_local": returnLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// GetChatScheduledMessages Returns all scheduled messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id)
// @param chatID Chat identifier
func (client *Client) GetChatScheduledMessages(chatID int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatScheduledMessages",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetMessagePublicForwards Returns forwarded copies of a channel message to different public channels. For optimal performance the number of returned messages is chosen by the library
// @param chatID Chat identifier of the message
// @param messageID Message identifier
// @param offset Offset of the first entry to return as received from the previous request; use empty string to get first chunk of results
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. Fewer messages may be returned than specified by the limit, even if the end of the list has not been reached
func (client *Client) GetMessagePublicForwards(chatID int64, messageID int64, offset string, limit int32) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessagePublicForwards",
		"chat_id":    chatID,
		"message_id": messageID,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err

}

// RemoveNotification Removes an active notification from notification list. Needs to be called only if the notification is removed by the current user
// @param notificationGroupID Identifier of notification group to which the notification belongs
// @param notificationID Identifier of removed notification
func (client *Client) RemoveNotification(notificationGroupID int32, notificationID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotification",
		"notification_group_id": notificationGroupID,
		"notification_id":       notificationID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveNotificationGroup Removes a group of active notifications. Needs to be called only if the notification group is removed by the current user
// @param notificationGroupID Notification group identifier
// @param maxNotificationID The maximum identifier of removed notifications
func (client *Client) RemoveNotificationGroup(notificationGroupID int32, maxNotificationID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "removeNotificationGroup",
		"notification_group_id": notificationGroupID,
		"max_notification_id":   maxNotificationID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels. This is an offline request
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatID int64, messageID int64, forAlbum bool, forComment bool) (*MessageLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessageLink",
		"chat_id":     chatID,
		"message_id":  messageID,
		"for_album":   forAlbum,
		"for_comment": forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLink MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err

}

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only for messages in supergroups and channels with a username
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param forAlbum Pass true to return an HTML code for embedding of the whole media album
func (client *Client) GetMessageEmbeddingCode(chatID int64, messageID int64, forAlbum bool) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageEmbeddingCode",
		"chat_id":    chatID,
		"message_id": messageID,
		"for_album":  forAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetMessageLinkInfo Returns information about a public or private message link
// @param uRL The message link in the format "https://t.me/c/...", or "tg://privatepost?...", or "https://t.me/username/...", or "tg://resolve?..."
func (client *Client) GetMessageLinkInfo(uRL string) (*MessageLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMessageLinkInfo",
		"url":   uRL,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLinkInfo MessageLinkInfo
	err = json.Unmarshal(result.Raw, &messageLinkInfo)
	return &messageLinkInfo, err

}

// SendMessage Sends a message. Returns the sent message
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageID Identifier of the message to reply to or 0
// @param options Options to be used to send the message
// @param replyMarkup Markup for replying to the message; for bots only
// @param inputMessageContent The content of the message to be sent
func (client *Client) SendMessage(chatID int64, messageThreadID int64, replyToMessageID int64, options *MessageSendOptions, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "sendMessage",
		"chat_id":               chatID,
		"message_thread_id":     messageThreadID,
		"reply_to_message_id":   replyToMessageID,
		"options":               options,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SendMessageAlbum Sends 2-10 messages grouped together into an album. Currently only audio, document, photo and video messages can be grouped into an album. Documents and audio files can be only grouped in an album with messages of the same type. Returns sent messages
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the messages will be sent
// @param replyToMessageID Identifier of a message to reply to or 0
// @param options Options to be used to send the messages
// @param inputMessageContents Contents of messages to be sent. At most 10 messages can be added to an album
func (client *Client) SendMessageAlbum(chatID int64, messageThreadID int64, replyToMessageID int64, options *MessageSendOptions, inputMessageContents []InputMessageContent) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "sendMessageAlbum",
		"chat_id":                chatID,
		"message_thread_id":      messageThreadID,
		"reply_to_message_id":    replyToMessageID,
		"options":                options,
		"input_message_contents": inputMessageContents,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command. Bots can't be invited to a private chat other than the chat with the bot. Bots can't be invited to channels (although they can be added as admins) and secret chats. Returns the sent message
// @param botUserID Identifier of the bot
// @param chatID Identifier of the target chat
// @param parameter A hidden parameter sent to the bot for deep linking purposes (https://core.telegram.org/bots#deep-linking)
func (client *Client) SendBotStartMessage(botUserID int32, chatID int64, parameter string) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserID,
		"chat_id":     chatID,
		"parameter":   parameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
// @param chatID Target chat
// @param messageThreadID If not 0, a message thread identifier in which the message will be sent
// @param replyToMessageID Identifier of a message to reply to or 0
// @param options Options to be used to send the message
// @param queryID Identifier of the inline query
// @param resultID Identifier of the inline result
// @param hideViaBot If true, there will be no mention of a bot, via which the message is sent. Can be used only for bots GetOption("animation_search_bot_username"), GetOption("photo_search_bot_username") and GetOption("venue_search_bot_username")
func (client *Client) SendInlineQueryResultMessage(chatID int64, messageThreadID int64, replyToMessageID int64, options *MessageSendOptions, queryID JSONInt64, resultID string, hideViaBot bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "sendInlineQueryResultMessage",
		"chat_id":             chatID,
		"message_thread_id":   messageThreadID,
		"reply_to_message_id": replyToMessageID,
		"options":             options,
		"query_id":            queryID,
		"result_id":           resultID,
		"hide_via_bot":        hideViaBot,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
// @param chatID Identifier of the chat to which to forward messages
// @param fromChatID Identifier of the chat from which to forward messages
// @param messageIDs Identifiers of the messages to forward. Message identifiers must be in a strictly increasing order. At most 100 messages can be forwarded simultaneously
// @param options Options to be used to send the messages
// @param sendCopy True, if content of the messages needs to be copied without links to the original messages. Always true if the messages are forwarded to a secret chat
// @param removeCaption True, if media caption of message copies needs to be removed. Ignored if send_copy is false
func (client *Client) ForwardMessages(chatID int64, fromChatID int64, messageIDs []int64, options *MessageSendOptions, sendCopy bool, removeCaption bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "forwardMessages",
		"chat_id":        chatID,
		"from_chat_id":   fromChatID,
		"message_ids":    messageIDs,
		"options":        options,
		"send_copy":      sendCopy,
		"remove_caption": removeCaption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// ResendMessages Resends messages which failed to send. Can be called only for messages for which messageSendingStateFailed.can_retry is true and after specified in messageSendingStateFailed.retry_after time passed.
// @param chatID Identifier of the chat to send messages
// @param messageIDs Identifiers of the messages to resend. Message identifiers must be in a strictly increasing order
func (client *Client) ResendMessages(chatID int64, messageIDs []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "resendMessages",
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendChatSetTTLMessage Changes the current TTL setting (sets a new self-destruct timer) in a secret chat and sends the corresponding message
// @param chatID Chat identifier
// @param tTL New TTL value, in seconds
func (client *Client) SendChatSetTTLMessage(chatID int64, tTL int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatSetTtlMessage",
		"chat_id": chatID,
		"ttl":     tTL,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendChatScreenshotTakenNotification Sends a notification about a screenshot taken in a chat. Supported only in private and secret chats
// @param chatID Chat identifier
func (client *Client) SendChatScreenshotTakenNotification(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatScreenshotTakenNotification",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
// @param chatID Target chat
// @param sender The sender sender of the message
// @param replyToMessageID Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message
// @param inputMessageContent The content of the message to be added
func (client *Client) AddLocalMessage(chatID int64, sender MessageSender, replyToMessageID int64, disableNotification bool, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "addLocalMessage",
		"chat_id":               chatID,
		"sender":                sender,
		"reply_to_message_id":   replyToMessageID,
		"disable_notification":  disableNotification,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// DeleteMessages Deletes messages
// @param chatID Chat identifier
// @param messageIDs Identifiers of the messages to be deleted
// @param revoke Pass true to try to delete messages for all chat members. Always true for supergroups, channels and secret chats
func (client *Client) DeleteMessages(chatID int64, messageIDs []int64, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteMessages",
		"chat_id":     chatID,
		"message_ids": messageIDs,
		"revoke":      revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// DeleteChatMessagesFromUser Deletes all messages sent by the specified user to a chat. Supported only for supergroups; requires can_delete_messages administrator privileges
// @param chatID Chat identifier
// @param userID User identifier
func (client *Client) DeleteChatMessagesFromUser(chatID int64, userID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChatMessagesFromUser",
		"chat_id": chatID,
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New text content of the message. Should be of type InputMessageText
func (client *Client) EditMessageText(chatID int64, messageID int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageText",
		"chat_id":               chatID,
		"message_id":            messageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param location New location content of the message; may be null. Pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditMessageLiveLocation(chatID int64, messageID int64, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editMessageLiveLocation",
		"chat_id":                chatID,
		"message_id":             messageID,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video. The media in the message can't be replaced if the message was set to self-destruct. Media can't be replaced by self-destructing media. Media in an album can be edited only to contain a photo or a video. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: InputMessageAnimation, InputMessageAudio, InputMessageDocument, InputMessagePhoto or InputMessageVideo
func (client *Client) EditMessageMedia(chatID int64, messageID int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageMedia",
		"chat_id":               chatID,
		"message_id":            messageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditMessageCaption(chatID int64, messageID int64, replyMarkup ReplyMarkup, caption *FormattedText) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageCaption",
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": replyMarkup,
		"caption":      caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param replyMarkup The new message reply markup
func (client *Client) EditMessageReplyMarkup(chatID int64, messageID int64, replyMarkup ReplyMarkup) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageReplyMarkup",
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
// @param inlineMessageID Inline message identifier
// @param replyMarkup The new message reply markup
// @param inputMessageContent New text content of the message. Should be of type InputMessageText
func (client *Client) EditInlineMessageText(inlineMessageID string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageText",
		"inline_message_id":     inlineMessageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
// @param inlineMessageID Inline message identifier
// @param replyMarkup The new message reply markup
// @param location New location content of the message; may be null. Pass null to stop sharing the live location
// @param heading The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
// @param proximityAlertRadius The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the notification is disabled
func (client *Client) EditInlineMessageLiveLocation(inlineMessageID string, replyMarkup ReplyMarkup, location *Location, heading int32, proximityAlertRadius int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "editInlineMessageLiveLocation",
		"inline_message_id":      inlineMessageID,
		"reply_markup":           replyMarkup,
		"location":               location,
		"heading":                heading,
		"proximity_alert_radius": proximityAlertRadius,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
// @param inlineMessageID Inline message identifier
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: InputMessageAnimation, InputMessageAudio, InputMessageDocument, InputMessagePhoto or InputMessageVideo
func (client *Client) EditInlineMessageMedia(inlineMessageID string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageMedia",
		"inline_message_id":     inlineMessageID,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
// @param inlineMessageID Inline message identifier
// @param replyMarkup The new message reply markup
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditInlineMessageCaption(inlineMessageID string, replyMarkup ReplyMarkup, caption *FormattedText) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageCaption",
		"inline_message_id": inlineMessageID,
		"reply_markup":      replyMarkup,
		"caption":           caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
// @param inlineMessageID Inline message identifier
// @param replyMarkup The new message reply markup
func (client *Client) EditInlineMessageReplyMarkup(inlineMessageID string, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageReplyMarkup",
		"inline_message_id": inlineMessageID,
		"reply_markup":      replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditMessageSchedulingState Edits the time when a scheduled message will be sent. Scheduling state of all messages in the same album or forwarded together with the message will be also changed
// @param chatID The chat the message belongs to
// @param messageID Identifier of the message
// @param schedulingState The new message scheduling state. Pass null to send the message immediately
func (client *Client) EditMessageSchedulingState(chatID int64, messageID int64, schedulingState MessageSchedulingState) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "editMessageSchedulingState",
		"chat_id":          chatID,
		"message_id":       messageID,
		"scheduling_state": schedulingState,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) contained in the text. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*TextEntities, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var textEntities TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err

}

// ParseTextEntities Parses Bold, Italic, Underline, Strikethrough, Code, Pre, PreCode, TextUrl and MentionName entities contained in the text. Can be called synchronously
// @param text The text to parse
// @param parseMode Text parse mode
func (client *Client) ParseTextEntities(text string, parseMode TextParseMode) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "parseTextEntities",
		"text":       text,
		"parse_mode": parseMode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// ParseMarkdown Parses Markdown entities in a human-friendly format, ignoring markup errors. Can be called synchronously
// @param text The text to parse. For example, "__italic__ ~~strikethrough~~ **bold** `code` ```pre``` __[italic__ text_url](telegram.org) __italic**bold italic__bold**"
func (client *Client) ParseMarkdown(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "parseMarkdown",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetMarkdownText Replaces text entities with Markdown formatting in a human-friendly format. Entities that can't be represented in Markdown unambiguously are kept as is. Can be called synchronously
// @param text The text
func (client *Client) GetMarkdownText(text *FormattedText) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMarkdownText",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously
// @param fileName The name of the file or path to the file
func (client *Client) GetFileMimeType(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileMimeType",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously
// @param mimeType The MIME type of the file
func (client *Client) GetFileExtension(mimeType string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileExtension",
		"mime_type": mimeType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// CleanFileName Removes potentially dangerous characters from the name of a file. The encoding of the file name is supposed to be UTF-8. Returns an empty string on failure. Can be called synchronously
// @param fileName File name or path to the file
func (client *Client) CleanFileName(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "cleanFileName",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. Can be called synchronously
// @param languagePackDatabasePath Path to the language pack database in which strings are stored
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackID Language pack identifier
// @param key Language pack key of the string to be returned
func (client *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackID string, key string) (LanguagePackStringValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                       "getLanguagePackString",
		"language_pack_database_path": languagePackDatabasePath,
		"localization_target":         localizationTarget,
		"language_pack_id":            languagePackID,
		"key":                         key,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LanguagePackStringValueEnum(result.Data["@type"].(string)) {

	case LanguagePackStringValueOrdinaryType:
		var languagePackStringValue LanguagePackStringValueOrdinary
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValuePluralizedType:
		var languagePackStringValue LanguagePackStringValuePluralized
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValueDeletedType:
		var languagePackStringValue LanguagePackStringValueDeleted
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetJsonValue Converts a JSON-serialized string to corresponding JsonValue object. Can be called synchronously
// @param jsonstring The JSON-serialized string
func (client *Client) GetJsonValue(jsonstring string) (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getJsonValue",
		"json":  jsonstring,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously
// @param jsonValue The JsonValue object
func (client *Client) GetJsonString(jsonValue JsonValue) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getJsonString",
		"json_value": jsonValue,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SetPollAnswer Changes the user answer to a poll. A poll in quiz mode can be answered only once
// @param chatID Identifier of the chat to which the poll belongs
// @param messageID Identifier of the message containing the poll
// @param optionIDs 0-based identifiers of answer options, chosen by the user. User can choose more than 1 answer option only is the poll allows multiple answers
func (client *Client) SetPollAnswer(chatID int64, messageID int64, optionIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setPollAnswer",
		"chat_id":    chatID,
		"message_id": messageID,
		"option_ids": optionIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPollVoters Returns users voted for the specified option in a non-anonymous polls. For the optimal performance the number of returned users is chosen by the library
// @param chatID Identifier of the chat to which the poll belongs
// @param messageID Identifier of the message containing the poll
// @param optionID 0-based identifier of the answer option
// @param offset Number of users to skip in the result; must be non-negative
// @param limit The maximum number of users to be returned; must be positive and can't be greater than 50. Fewer users may be returned than specified by the limit, even if the end of the voter list has not been reached
func (client *Client) GetPollVoters(chatID int64, messageID int64, optionID int32, offset int32, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPollVoters",
		"chat_id":    chatID,
		"message_id": messageID,
		"option_id":  optionID,
		"offset":     offset,
		"limit":      limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// StopPoll Stops a poll. A poll in a message can be stopped when the message has can_be_edited flag set
// @param chatID Identifier of the chat to which the poll belongs
// @param messageID Identifier of the message containing the poll
// @param replyMarkup The new message reply markup; for bots only
func (client *Client) StopPoll(chatID int64, messageID int64, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "stopPoll",
		"chat_id":      chatID,
		"message_id":   messageID,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// HideSuggestedAction Hides a suggested action
// @param action Suggested action to hide
func (client *Client) HideSuggestedAction(action SuggestedAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "hideSuggestedAction",
		"action": action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLoginURLInfo Returns information about a button of type inlineKeyboardButtonTypeLoginUrl. The method needs to be called when the user presses the button
// @param chatID Chat identifier of the message with the button
// @param messageID Message identifier of the message with the button
// @param buttonID Button identifier
func (client *Client) GetLoginURLInfo(chatID int64, messageID int64, buttonID int32) (LoginURLInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLoginUrlInfo",
		"chat_id":    chatID,
		"message_id": messageID,
		"button_id":  buttonID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LoginURLInfoEnum(result.Data["@type"].(string)) {

	case LoginURLInfoOpenType:
		var loginURLInfo LoginURLInfoOpen
		err = json.Unmarshal(result.Raw, &loginURLInfo)
		return &loginURLInfo, err

	case LoginURLInfoRequestConfirmationType:
		var loginURLInfo LoginURLInfoRequestConfirmation
		err = json.Unmarshal(result.Raw, &loginURLInfo)
		return &loginURLInfo, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetLoginURL Returns an HTTP URL which can be used to automatically authorize the user on a website after clicking an inline button of type inlineKeyboardButtonTypeLoginUrl.
// @param chatID Chat identifier of the message with the button
// @param messageID Message identifier of the message with the button
// @param buttonID Button identifier
// @param allowWriteAccess True, if the user allowed the bot to send them messages
func (client *Client) GetLoginURL(chatID int64, messageID int64, buttonID int32, allowWriteAccess bool) (*HttpURL, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "getLoginUrl",
		"chat_id":            chatID,
		"message_id":         messageID,
		"button_id":          buttonID,
		"allow_write_access": allowWriteAccess,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpURL HttpURL
	err = json.Unmarshal(result.Raw, &httpURL)
	return &httpURL, err

}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param botUserID The identifier of the target bot
// @param chatID Identifier of the chat where the query was sent
// @param userLocation Location of the user, only if needed
// @param query Text of the query
// @param offset Offset of the first entry to return
func (client *Client) GetInlineQueryResults(botUserID int32, chatID int64, userLocation *Location, query string, offset string) (*InlineQueryResults, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getInlineQueryResults",
		"bot_user_id":   botUserID,
		"chat_id":       chatID,
		"user_location": userLocation,
		"query":         query,
		"offset":        offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var inlineQueryResults InlineQueryResults
	err = json.Unmarshal(result.Raw, &inlineQueryResults)
	return &inlineQueryResults, err

}

// AnswerInlineQuery Sets the result of an inline query; for bots only
// @param inlineQueryID Identifier of the inline query
// @param isPersonal True, if the result of the query can be cached for the specified user
// @param results The results of the query
// @param cacheTime Allowed time to cache the results of the query, in seconds
// @param nextOffset Offset for the next inline query; pass an empty string if there are no more results
// @param switchPmText If non-empty, this text should be shown on the button that opens a private chat with the bot and sends a start message to the bot with the parameter switch_pm_parameter
// @param switchPmParameter The parameter for the bot start message
func (client *Client) AnswerInlineQuery(inlineQueryID JSONInt64, isPersonal bool, results []InputInlineQueryResult, cacheTime int32, nextOffset string, switchPmText string, switchPmParameter string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "answerInlineQuery",
		"inline_query_id":     inlineQueryID,
		"is_personal":         isPersonal,
		"results":             results,
		"cache_time":          cacheTime,
		"next_offset":         nextOffset,
		"switch_pm_text":      switchPmText,
		"switch_pm_parameter": switchPmParameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param chatID Identifier of the chat with the message
// @param messageID Identifier of the message from which the query originated
// @param payload Query payload
func (client *Client) GetCallbackQueryAnswer(chatID int64, messageID int64, payload CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatID,
		"message_id": messageID,
		"payload":    payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callbackQueryAnswer CallbackQueryAnswer
	err = json.Unmarshal(result.Raw, &callbackQueryAnswer)
	return &callbackQueryAnswer, err

}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
// @param callbackQueryID Identifier of the callback query
// @param text Text of the answer
// @param showAlert If true, an alert should be shown to the user instead of a toast notification
// @param uRL URL to be opened
// @param cacheTime Time during which the result of the query can be cached, in seconds
func (client *Client) AnswerCallbackQuery(callbackQueryID JSONInt64, text string, showAlert bool, uRL string, cacheTime int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerCallbackQuery",
		"callback_query_id": callbackQueryID,
		"text":              text,
		"show_alert":        showAlert,
		"url":               uRL,
		"cache_time":        cacheTime,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerShippingQuery Sets the result of a shipping query; for bots only
// @param shippingQueryID Identifier of the shipping query
// @param shippingOptions Available shipping options
// @param errorMessage An error message, empty on success
func (client *Client) AnswerShippingQuery(shippingQueryID JSONInt64, shippingOptions []ShippingOption, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerShippingQuery",
		"shipping_query_id": shippingQueryID,
		"shipping_options":  shippingOptions,
		"error_message":     errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only
// @param preCheckoutQueryID Identifier of the pre-checkout query
// @param errorMessage An error message, empty on success
func (client *Client) AnswerPreCheckoutQuery(preCheckoutQueryID JSONInt64, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "answerPreCheckoutQuery",
		"pre_checkout_query_id": preCheckoutQueryID,
		"error_message":         errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGameScore Updates the game score of the specified user in the game; for bots only
// @param chatID The chat to which the message with the game belongs
// @param messageID Identifier of the message
// @param editMessage True, if the message should be edited
// @param userID User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetGameScore(chatID int64, messageID int64, editMessage bool, userID int32, score int32, force bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setGameScore",
		"chat_id":      chatID,
		"message_id":   messageID,
		"edit_message": editMessage,
		"user_id":      userID,
		"score":        score,
		"force":        force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
// @param inlineMessageID Inline message identifier
// @param editMessage True, if the message should be edited
// @param userID User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetInlineGameScore(inlineMessageID string, editMessage bool, userID int32, score int32, force bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setInlineGameScore",
		"inline_message_id": inlineMessageID,
		"edit_message":      editMessage,
		"user_id":           userID,
		"score":             score,
		"force":             force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only
// @param chatID The chat that contains the message with the game
// @param messageID Identifier of the message
// @param userID User identifier
func (client *Client) GetGameHighScores(chatID int64, messageID int64, userID int32) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getGameHighScores",
		"chat_id":    chatID,
		"message_id": messageID,
		"user_id":    userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only
// @param inlineMessageID Inline message identifier
// @param userID User identifier
func (client *Client) GetInlineGameHighScores(inlineMessageID string, userID int32) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getInlineGameHighScores",
		"inline_message_id": inlineMessageID,
		"user_id":           userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a ForceReply reply markup has been used. UpdateChatReplyMarkup will be sent if the reply markup will be changed
// @param chatID Chat identifier
// @param messageID The message identifier of the used keyboard
func (client *Client) DeleteChatReplyMarkup(chatID int64, messageID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "deleteChatReplyMarkup",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendChatAction Sends a notification about user activity in a chat
// @param chatID Chat identifier
// @param messageThreadID If not 0, a message thread identifier in which the action was performed
// @param action The action description
func (client *Client) SendChatAction(chatID int64, messageThreadID int64, action ChatAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendChatAction",
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
		"action":            action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenChat Informs TDLib that the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats)
// @param chatID Chat identifier
func (client *Client) OpenChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "openChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CloseChat Informs TDLib that the chat is closed by the user. Many useful activities depend on the chat being opened or closed
// @param chatID Chat identifier
func (client *Client) CloseChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "closeChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewMessages Informs TDLib that messages are being viewed by the user. Many useful activities depend on whether the messages are currently being viewed or not (e.g., marking messages as read, incrementing a view counter, updating a view counter, removing deleted messages in supergroups and channels)
// @param chatID Chat identifier
// @param messageThreadID If not 0, a message thread identifier in which the messages are being viewed
// @param messageIDs The identifiers of the messages being viewed
// @param forceRead True, if messages in closed chats should be marked as read by the request
func (client *Client) ViewMessages(chatID int64, messageThreadID int64, messageIDs []int64, forceRead bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "viewMessages",
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
		"message_ids":       messageIDs,
		"force_read":        forceRead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenMessageContent Informs TDLib that the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message). An updateMessageContentOpened update will be generated if something has changed
// @param chatID Chat identifier of the message
// @param messageID Identifier of the message with the opened content
func (client *Client) OpenMessageContent(chatID int64, messageID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "openMessageContent",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReadAllChatMentions Marks all mentions in a chat as read
// @param chatID Chat identifier
func (client *Client) ReadAllChatMentions(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readAllChatMentions",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreatePrivateChat Returns an existing chat corresponding to a given user
// @param userID User identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreatePrivateChat(userID int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createPrivateChat",
		"user_id": userID,
		"force":   force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group
// @param basicGroupID Basic group identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateBasicGroupChat(basicGroupID int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createBasicGroupChat",
		"basic_group_id": basicGroupID,
		"force":          force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel
// @param supergroupID Supergroup or channel identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateSupergroupChat(supergroupID int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "createSupergroupChat",
		"supergroup_id": supergroupID,
		"force":         force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat
// @param secretChatID Secret chat identifier
func (client *Client) CreateSecretChat(secretChatID int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createSecretChat",
		"secret_chat_id": secretChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns the newly created chat
// @param userIDs Identifiers of users to be added to the basic group
// @param title Title of the new basic group; 1-128 characters
func (client *Client) CreateNewBasicGroupChat(userIDs []int32, title string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewBasicGroupChat",
		"user_ids": userIDs,
		"title":    title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
// @param title Title of the new chat; 1-128 characters
// @param isChannel True, if a channel chat needs to be created
// @param description
// @param location Chat location if a location-based supergroup is being created
// @param forImport True, if the supergroup is created for importing messages using importMessage
func (client *Client) CreateNewSupergroupChat(title string, isChannel bool, description string, location *ChatLocation, forImport bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "createNewSupergroupChat",
		"title":       title,
		"is_channel":  isChannel,
		"description": description,
		"location":    location,
		"for_import":  forImport,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat
// @param userID Identifier of the target user
func (client *Client) CreateNewSecretChat(userID int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createNewSecretChat",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires creator privileges. Deactivates the original basic group
// @param chatID Identifier of the chat to upgrade
func (client *Client) UpgradeBasicGroupChatToSupergroupChat(chatID int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// GetChatListsToAddChat Returns chat lists to which the chat can be added. This is an offline request
// @param chatID Chat identifier
func (client *Client) GetChatListsToAddChat(chatID int64) (*ChatLists, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatListsToAddChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatLists ChatLists
	err = json.Unmarshal(result.Raw, &chatLists)
	return &chatLists, err

}

// AddChatToList Adds a chat to a chat list. A chat can't be simultaneously in Main and Archive chat lists, so it is automatically removed from another one if needed
// @param chatID Chat identifier
// @param chatList The chat list. Use getChatListsToAddChat to get suitable chat lists
func (client *Client) AddChatToList(chatID int64, chatList ChatList) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addChatToList",
		"chat_id":   chatID,
		"chat_list": chatList,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatFilter Returns information about a chat filter by its identifier
// @param chatFilterID Chat filter identifier
func (client *Client) GetChatFilter(chatFilterID int32) (*ChatFilter, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChatFilter",
		"chat_filter_id": chatFilterID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterDummy ChatFilter
	err = json.Unmarshal(result.Raw, &chatFilterDummy)
	return &chatFilterDummy, err

}

// CreateChatFilter Creates new chat filter. Returns information about the created chat filter
// @param filter Chat filter
func (client *Client) CreateChatFilter(filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "createChatFilter",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}

// EditChatFilter Edits existing chat filter. Returns information about the edited chat filter
// @param chatFilterID Chat filter identifier
// @param filter The edited chat filter
func (client *Client) EditChatFilter(chatFilterID int32, filter *ChatFilter) (*ChatFilterInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "editChatFilter",
		"chat_filter_id": chatFilterID,
		"filter":         filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatFilterInfo ChatFilterInfo
	err = json.Unmarshal(result.Raw, &chatFilterInfo)
	return &chatFilterInfo, err

}

// DeleteChatFilter Deletes existing chat filter
// @param chatFilterID Chat filter identifier
func (client *Client) DeleteChatFilter(chatFilterID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "deleteChatFilter",
		"chat_filter_id": chatFilterID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReorderChatFilters Changes the order of chat filters
// @param chatFilterIDs Identifiers of chat filters in the new correct order
func (client *Client) ReorderChatFilters(chatFilterIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderChatFilters",
		"chat_filter_ids": chatFilterIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecommendedChatFilters Returns recommended chat filters for the current user
func (client *Client) GetRecommendedChatFilters() (*RecommendedChatFilters, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecommendedChatFilters",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var recommendedChatFilters RecommendedChatFilters
	err = json.Unmarshal(result.Raw, &recommendedChatFilters)
	return &recommendedChatFilters, err

}

// GetChatFilterDefaultIconName Returns default icon name for a filter. Can be called synchronously
// @param filter Chat filter
func (client *Client) GetChatFilterDefaultIconName(filter *ChatFilter) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getChatFilterDefaultIconName",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatID Chat identifier
// @param title New title of the chat; 1-128 characters
func (client *Client) SetChatTitle(chatID int64, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatTitle",
		"chat_id": chatID,
		"title":   title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires can_change_info administrator right
// @param chatID Chat identifier
// @param photo New chat photo. Pass null to delete the chat photo
func (client *Client) SetChatPhoto(chatID int64, photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatPhoto",
		"chat_id": chatID,
		"photo":   photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatPermissions Changes the chat members permissions. Supported only for basic groups and supergroups. Requires can_restrict_members administrator right
// @param chatID Chat identifier
// @param permissions New non-administrator members permissions in the chat
func (client *Client) SetChatPermissions(chatID int64, permissions *ChatPermissions) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatPermissions",
		"chat_id":     chatID,
		"permissions": permissions,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDraftMessage Changes the draft message in a chat
// @param chatID Chat identifier
// @param messageThreadID If not 0, a message thread identifier in which the draft was changed
// @param draftMessage New draft message; may be null
func (client *Client) SetChatDraftMessage(chatID int64, messageThreadID int64, draftMessage *DraftMessage) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setChatDraftMessage",
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
		"draft_message":     draftMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatNotificationSettings Changes the notification settings of a chat. Notification settings of a chat with the current user (Saved Messages) can't be changed
// @param chatID Chat identifier
// @param notificationSettings New notification settings for the chat. If the chat is muted for more than 1 week, it is considered to be muted forever
func (client *Client) SetChatNotificationSettings(chatID int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setChatNotificationSettings",
		"chat_id":               chatID,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat
// @param chatID Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func (client *Client) ToggleChatIsMarkedAsUnread(chatID int64, isMarkedAsUnread bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleChatIsMarkedAsUnread",
		"chat_id":             chatID,
		"is_marked_as_unread": isMarkedAsUnread,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat
// @param chatID Chat identifier
// @param defaultDisableNotification New value of default_disable_notification
func (client *Client) ToggleChatDefaultDisableNotification(chatID int64, defaultDisableNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                        "toggleChatDefaultDisableNotification",
		"chat_id":                      chatID,
		"default_disable_notification": defaultDisableNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatClientData Changes application-specific data associated with a chat
// @param chatID Chat identifier
// @param clientData New value of client_data
func (client *Client) SetChatClientData(chatID int64, clientData string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatClientData",
		"chat_id":     chatID,
		"client_data": clientData,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDescription Changes information about a chat. Available for basic groups, supergroups, and channels. Requires can_change_info administrator right
// @param chatID Identifier of the chat
// @param description
func (client *Client) SetChatDescription(chatID int64, description string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatDescription",
		"chat_id":     chatID,
		"description": description,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDiscussionGroup Changes the discussion group of a channel chat; requires can_change_info administrator right in the channel if it is specified
// @param chatID Identifier of the channel chat. Pass 0 to remove a link from the supergroup passed in the second argument to a linked channel chat (requires can_pin_messages rights in the supergroup)
// @param discussionChatID Identifier of a new channel's discussion group. Use 0 to remove the discussion group.
func (client *Client) SetChatDiscussionGroup(chatID int64, discussionChatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setChatDiscussionGroup",
		"chat_id":            chatID,
		"discussion_chat_id": discussionChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatLocation Changes the location of a chat. Available only for some location-based supergroups, use supergroupFullInfo.can_set_location to check whether the method is allowed to use
// @param chatID Chat identifier
// @param location New location for the chat; must be valid and not null
func (client *Client) SetChatLocation(chatID int64, location *ChatLocation) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setChatLocation",
		"chat_id":  chatID,
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatSlowModeDelay Changes the slow mode delay of a chat. Available only for supergroups; requires can_restrict_members rights
// @param chatID Chat identifier
// @param slowModeDelay New slow mode delay for the chat; must be one of 0, 10, 30, 60, 300, 900, 3600
func (client *Client) SetChatSlowModeDelay(chatID int64, slowModeDelay int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "setChatSlowModeDelay",
		"chat_id":         chatID,
		"slow_mode_delay": slowModeDelay,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// PinChatMessage Pins a message in a chat; requires can_pin_messages rights or can_edit_messages rights in the channel
// @param chatID Identifier of the chat
// @param messageID Identifier of the new pinned message
// @param disableNotification True, if there should be no notification about the pinned message. Notifications are always disabled in channels and private chats
// @param onlyForSelf True, if the message needs to be pinned for one side only; private chats only
func (client *Client) PinChatMessage(chatID int64, messageID int64, disableNotification bool, onlyForSelf bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "pinChatMessage",
		"chat_id":              chatID,
		"message_id":           messageID,
		"disable_notification": disableNotification,
		"only_for_self":        onlyForSelf,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnpinChatMessage Removes a pinned message from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatID Identifier of the chat
// @param messageID Identifier of the removed pinned message
func (client *Client) UnpinChatMessage(chatID int64, messageID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "unpinChatMessage",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnpinAllChatMessages Removes all pinned messages from a chat; requires can_pin_messages rights in the group or can_edit_messages rights in the channel
// @param chatID Identifier of the chat
func (client *Client) UnpinAllChatMessages(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "unpinAllChatMessages",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// JoinChat Adds the current user as a new member to a chat. Private and secret chats can't be joined using this method
// @param chatID Chat identifier
func (client *Client) JoinChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "joinChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LeaveChat Removes the current user from chat members. Private and secret chats can't be left using this method
// @param chatID Chat identifier
func (client *Client) LeaveChat(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "leaveChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMember Adds a new member to a chat. Members can't be added to private or secret chats
// @param chatID Chat identifier
// @param userID Identifier of the user
// @param forwardLimit The number of earlier messages from the chat to be forwarded to the new member; up to 100. Ignored for supergroups and channels
func (client *Client) AddChatMember(chatID int64, userID int32, forwardLimit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "addChatMember",
		"chat_id":       chatID,
		"user_id":       userID,
		"forward_limit": forwardLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMembers Adds multiple new members to a chat. Currently this method is only available for supergroups and channels. This method can't be used to join a chat. Members can't be added to a channel if it has more than 200 members
// @param chatID Chat identifier
// @param userIDs Identifiers of the users to be added to the chat. The maximum number of added users is 20 for supergroups and 100 for channels
func (client *Client) AddChatMembers(chatID int64, userIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "addChatMembers",
		"chat_id":  chatID,
		"user_ids": userIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatMemberStatus Changes the status of a chat member, needs appropriate privileges. This function is currently not suitable for adding new members to the chat and transferring chat ownership; instead, use addChatMember or transferChatOwnership
// @param chatID Chat identifier
// @param userID User identifier
// @param status The new status of the member in the chat
func (client *Client) SetChatMemberStatus(chatID int64, userID int32, status ChatMemberStatus) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatMemberStatus",
		"chat_id": chatID,
		"user_id": userID,
		"status":  status,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// BanChatMember Bans a member in a chat. Members can't be banned in private or secret chats. In supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first
// @param chatID Chat identifier
// @param userID Identifier of the user
// @param bannedUntilDate Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever. Ignored in basic groups
// @param revokeMessages Pass true to delete all messages in the chat for the user. Always true for supergroups and channels
func (client *Client) BanChatMember(chatID int64, userID int32, bannedUntilDate int32, revokeMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "banChatMember",
		"chat_id":           chatID,
		"user_id":           userID,
		"banned_until_date": bannedUntilDate,
		"revoke_messages":   revokeMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// CanTransferOwnership Checks whether the current session can be used to transfer a chat ownership to another user
func (client *Client) CanTransferOwnership() (CanTransferOwnershipResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "canTransferOwnership",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CanTransferOwnershipResultEnum(result.Data["@type"].(string)) {

	case CanTransferOwnershipResultOkType:
		var canTransferOwnershipResult CanTransferOwnershipResultOk
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordNeededType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordNeeded
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultPasswordTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultPasswordTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	case CanTransferOwnershipResultSessionTooFreshType:
		var canTransferOwnershipResult CanTransferOwnershipResultSessionTooFresh
		err = json.Unmarshal(result.Raw, &canTransferOwnershipResult)
		return &canTransferOwnershipResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// TransferChatOwnership Changes the owner of a chat. The current user must be a current owner of the chat. Use the method canTransferOwnership to check whether the ownership can be transferred from the current session. Available only for supergroups and channel chats
// @param chatID Chat identifier
// @param userID Identifier of the user to which transfer the ownership. The ownership can't be transferred to a bot or to a deleted user
// @param password The password of the current user
func (client *Client) TransferChatOwnership(chatID int64, userID int32, password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "transferChatOwnership",
		"chat_id":  chatID,
		"user_id":  userID,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatMember Returns information about a single member of a chat
// @param chatID Chat identifier
// @param userID User identifier
func (client *Client) GetChatMember(chatID int64, userID int32) (*ChatMember, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMember",
		"chat_id": chatID,
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMember ChatMember
	err = json.Unmarshal(result.Raw, &chatMember)
	return &chatMember, err

}

// SearchChatMembers Searches for a specified query in the first name, last name and username of the members of a specified chat. Requires administrator rights in channels
// @param chatID Chat identifier
// @param query Query to search for
// @param limit The maximum number of users to be returned
// @param filter The type of users to return. By default, chatMembersFilterMembers
func (client *Client) SearchChatMembers(chatID int64, query string, limit int32, filter ChatMembersFilter) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatMembers",
		"chat_id": chatID,
		"query":   query,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// GetChatAdministrators Returns a list of administrators of the chat with their custom titles
// @param chatID Chat identifier
func (client *Client) GetChatAdministrators(chatID int64) (*ChatAdministrators, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatAdministrators",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatAdministrators ChatAdministrators
	err = json.Unmarshal(result.Raw, &chatAdministrators)
	return &chatAdministrators, err

}

// ClearAllDraftMessages Clears draft messages in all chats
// @param excludeSecretChats If true, local draft messages in secret chats will not be cleared
func (client *Client) ClearAllDraftMessages(excludeSecretChats bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "clearAllDraftMessages",
		"exclude_secret_chats": excludeSecretChats,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatNotificationSettingsExceptions Returns list of chats with non-default notification settings
// @param scope If specified, only chats from the specified scope will be returned
// @param compareSound If true, also chats with non-default sound will be returned
func (client *Client) GetChatNotificationSettingsExceptions(scope NotificationSettingsScope, compareSound bool) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatNotificationSettingsExceptions",
		"scope":         scope,
		"compare_sound": compareSound,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type
// @param scope Types of chats for which to return the notification settings information
func (client *Client) GetScopeNotificationSettings(scope NotificationSettingsScope) (*ScopeNotificationSettings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getScopeNotificationSettings",
		"scope": scope,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var scopeNotificationSettings ScopeNotificationSettings
	err = json.Unmarshal(result.Raw, &scopeNotificationSettings)
	return &scopeNotificationSettings, err

}

// SetScopeNotificationSettings Changes notification settings for chats of a given type
// @param scope Types of chats for which to change the notification settings
// @param notificationSettings The new notification settings for the given scope
func (client *Client) SetScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setScopeNotificationSettings",
		"scope":                 scope,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetAllNotificationSettings Resets all notification settings to their default values. By default, all chats are unmuted, the sound is set to "default" and message previews are shown
func (client *Client) ResetAllNotificationSettings() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetAllNotificationSettings",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsPinned Changes the pinned state of a chat. There can be up to GetOption("pinned_chat_count_max")/GetOption("pinned_archived_chat_count_max") pinned non-secret chats and the same number of secret chats in the main/arhive chat list
// @param chatList Chat list in which to change the pinned state of the chat
// @param chatID Chat identifier
// @param isPinned True, if the chat is pinned
func (client *Client) ToggleChatIsPinned(chatList ChatList, chatID int64, isPinned bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "toggleChatIsPinned",
		"chat_list": chatList,
		"chat_id":   chatID,
		"is_pinned": isPinned,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPinnedChats Changes the order of pinned chats
// @param chatList Chat list in which to change the order of pinned chats
// @param chatIDs The new list of pinned chats
func (client *Client) SetPinnedChats(chatList ChatList, chatIDs []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setPinnedChats",
		"chat_list": chatList,
		"chat_ids":  chatIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DownloadFile Downloads a file from the cloud. Download progress and completion of the download will be notified through updateFile updates
// @param fileID Identifier of the file to download
// @param priority Priority of the download (1-32). The higher the priority, the earlier the file will be downloaded. If the priorities of two files are equal, then the last one for which downloadFile was called will be downloaded first
// @param offset The starting position from which the file should be downloaded
// @param limit Number of bytes which should be downloaded starting from the "offset" position before the download will be automatically cancelled; use 0 to download without a limit
// @param synchronous If false, this request returns file state just after the download has been started. If true, this request returns file state only after
func (client *Client) DownloadFile(fileID int32, priority int32, offset int32, limit int32, synchronous bool) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "downloadFile",
		"file_id":     fileID,
		"priority":    priority,
		"offset":      offset,
		"limit":       limit,
		"synchronous": synchronous,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetFileDownloadedPrefixSize Returns file downloaded prefix size from a given offset
// @param fileID Identifier of the file
// @param offset Offset from which downloaded prefix size should be calculated
func (client *Client) GetFileDownloadedPrefixSize(fileID int32, offset int32) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFileDownloadedPrefixSize",
		"file_id": fileID,
		"offset":  offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing
// @param fileID Identifier of a file to stop downloading
// @param onlyIfPending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (client *Client) CancelDownloadFile(fileID int32, onlyIfPending bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "cancelDownloadFile",
		"file_id":         fileID,
		"only_if_pending": onlyIfPending,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UploadFile Asynchronously uploads a file to the cloud without sending it in a message. updateFile will be used to notify about upload progress and successful completion of the upload. The file will not have a persistent remote identifier until it will be sent in a message
// @param file File to upload
// @param fileType File type
// @param priority Priority of the upload (1-32). The higher the priority, the earlier the file will be uploaded. If the priorities of two files are equal, then the first one for which uploadFile was called will be uploaded first
func (client *Client) UploadFile(file InputFile, fileType FileType, priority int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "uploadFile",
		"file":      file,
		"file_type": fileType,
		"priority":  priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// CancelUploadFile Stops the uploading of a file. Supported only for files uploaded by using uploadFile. For other files the behavior is undefined
// @param fileID Identifier of the file to stop uploading
func (client *Client) CancelUploadFile(fileID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "cancelUploadFile",
		"file_id": fileID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// WriteGeneratedFilePart Writes a part of a generated file. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct write to the destination file
// @param generationID The identifier of the generation process
// @param offset The offset from which to write the data to the file
// @param data The data to write
func (client *Client) WriteGeneratedFilePart(generationID JSONInt64, offset int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "writeGeneratedFilePart",
		"generation_id": generationID,
		"offset":        offset,
		"data":          data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetFileGenerationProgress Informs TDLib on a file generation progress
// @param generationID The identifier of the generation process
// @param expectedSize Expected size of the generated file, in bytes; 0 if unknown
// @param localPrefixSize The number of bytes already generated
func (client *Client) SetFileGenerationProgress(generationID JSONInt64, expectedSize int32, localPrefixSize int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setFileGenerationProgress",
		"generation_id":     generationID,
		"expected_size":     expectedSize,
		"local_prefix_size": localPrefixSize,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// FinishFileGeneration Finishes the file generation
// @param generationID The identifier of the generation process
// @param error If set, means that file generation has failed and should be terminated
func (client *Client) FinishFileGeneration(generationID JSONInt64, error *Error) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "finishFileGeneration",
		"generation_id": generationID,
		"error":         error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReadFilePart Reads a part of a file from the TDLib file cache and returns read bytes. This method is intended to be used only if the application has no direct access to TDLib's file system, because it is usually slower than a direct read from the file
// @param fileID Identifier of the file. The file must be located in the TDLib file cache
// @param offset The offset from which to read the file
// @param count Number of bytes to read. An error will be returned if there are not enough bytes available in the file from the specified position. Pass 0 to read all available data from the specified position
func (client *Client) ReadFilePart(fileID int32, offset int32, count int32) (*FilePart, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readFilePart",
		"file_id": fileID,
		"offset":  offset,
		"count":   count,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var filePart FilePart
	err = json.Unmarshal(result.Raw, &filePart)
	return &filePart, err

}

// DeleteFile Deletes a file from the TDLib file cache
// @param fileID Identifier of the file to delete
func (client *Client) DeleteFile(fileID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteFile",
		"file_id": fileID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMessageFileType Returns information about a file with messages exported from another app
// @param messageFileHead Beginning of the message file; up to 100 first lines
func (client *Client) GetMessageFileType(messageFileHead string) (MessageFileType, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getMessageFileType",
		"message_file_head": messageFileHead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch MessageFileTypeEnum(result.Data["@type"].(string)) {

	case MessageFileTypePrivateType:
		var messageFileType MessageFileTypePrivate
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeGroupType:
		var messageFileType MessageFileTypeGroup
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	case MessageFileTypeUnknownType:
		var messageFileType MessageFileTypeUnknown
		err = json.Unmarshal(result.Raw, &messageFileType)
		return &messageFileType, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// ImportMessages Imports messages exported from another app
// @param chatID Identifier of a chat to which the messages will be imported. It must be an identifier of a private chat with a mutual contact or an identifier of a supergroup chat with can_change_info administrator right
// @param messageFile File with messages to import. Only inputFileLocal and inputFileGenerated are supported. The file must not be previously uploaded
// @param attachedFiles Files used in the imported messages. Only inputFileLocal and inputFileGenerated are supported. The files must not be previously uploaded
func (client *Client) ImportMessages(chatID int64, messageFile InputFile, attachedFiles []InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "importMessages",
		"chat_id":        chatID,
		"message_file":   messageFile,
		"attached_files": attachedFiles,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReplacePermanentChatInviteLink Replaces current permanent invite link for a chat with a new permanent invite link. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right
// @param chatID Chat identifier
func (client *Client) ReplacePermanentChatInviteLink(chatID int64) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "replacePermanentChatInviteLink",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked; must begin with "https://t.me/joinchat/", "https://telegram.me/joinchat/", or "https://telegram.dog/joinchat/"
func (client *Client) CheckChatInviteLink(inviteLink string) (*ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkInfo ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err

}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible
// @param inviteLink Invite link to import; must begin with "https://t.me/joinchat/", "https://telegram.me/joinchat/", or "https://telegram.dog/joinchat/"
func (client *Client) JoinChatByInviteLink(inviteLink string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "joinChatByInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateCall Creates a new call
// @param userID Identifier of the user to be called
// @param protocol Description of the call protocols supported by the application
// @param isVideo True, if a video call needs to be created
func (client *Client) CreateCall(userID int32, protocol *CallProtocol, isVideo bool) (*CallID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createCall",
		"user_id":  userID,
		"protocol": protocol,
		"is_video": isVideo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callID CallID
	err = json.Unmarshal(result.Raw, &callID)
	return &callID, err

}

// AcceptCall Accepts an incoming call
// @param callID Call identifier
// @param protocol Description of the call protocols supported by the application
func (client *Client) AcceptCall(callID int32, protocol *CallProtocol) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "acceptCall",
		"call_id":  callID,
		"protocol": protocol,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallSignalingData Sends call signaling data
// @param callID Call identifier
// @param data The data
func (client *Client) SendCallSignalingData(callID int32, data []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendCallSignalingData",
		"call_id": callID,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DiscardCall Discards a call
// @param callID Call identifier
// @param isDisconnected True, if the user was disconnected
// @param duration The call duration, in seconds
// @param isVideo True, if the call was a video call
// @param connectionID Identifier of the connection used during the call
func (client *Client) DiscardCall(callID int32, isDisconnected bool, duration int32, isVideo bool, connectionID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "discardCall",
		"call_id":         callID,
		"is_disconnected": isDisconnected,
		"duration":        duration,
		"is_video":        isVideo,
		"connection_id":   connectionID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallRating Sends a call rating
// @param callID Call identifier
// @param rating Call rating; 1-5
// @param comment An optional user comment if the rating is less than 5
// @param problems List of the exact types of problems with the call, specified by the user
func (client *Client) SendCallRating(callID int32, rating int32, comment string, problems []CallProblem) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "sendCallRating",
		"call_id":  callID,
		"rating":   rating,
		"comment":  comment,
		"problems": problems,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallDebugInformation Sends debug information for a call
// @param callID Call identifier
// @param debugInformation Debug information in application-specific format
func (client *Client) SendCallDebugInformation(callID int32, debugInformation string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendCallDebugInformation",
		"call_id":           callID,
		"debug_information": debugInformation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreateVoiceChat Creates a voice chat (a group call bound to a chat). Available only for basic groups and supergroups; requires can_manage_voice_chats rights
// @param chatID Chat identifier
func (client *Client) CreateVoiceChat(chatID int64) (*GroupCallID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createVoiceChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallID GroupCallID
	err = json.Unmarshal(result.Raw, &groupCallID)
	return &groupCallID, err

}

// GetGroupCall Returns information about a group call
// @param groupCallID Group call identifier
func (client *Client) GetGroupCall(groupCallID int32) (*GroupCall, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getGroupCall",
		"group_call_id": groupCallID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallDummy GroupCall
	err = json.Unmarshal(result.Raw, &groupCallDummy)
	return &groupCallDummy, err

}

// JoinGroupCall Joins a group call
// @param groupCallID Group call identifier
// @param payload Group join payload, received from tgcalls. Use null to cancel previous joinGroupCall request
// @param source Caller synchronization source identifier; received from tgcalls
// @param isMuted True, if the user's microphone is muted
func (client *Client) JoinGroupCall(groupCallID int32, payload *GroupCallPayload, source int32, isMuted bool) (*GroupCallJoinResponse, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "joinGroupCall",
		"group_call_id": groupCallID,
		"payload":       payload,
		"source":        source,
		"is_muted":      isMuted,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var groupCallJoinResponse GroupCallJoinResponse
	err = json.Unmarshal(result.Raw, &groupCallJoinResponse)
	return &groupCallJoinResponse, err

}

// ToggleGroupCallMuteNewParticipants Toggles whether new participants of a group call can be unmuted only by administrators of the group call. Requires groupCall.can_change_mute_new_participants group call flag
// @param groupCallID Group call identifier
// @param muteNewParticipants New value of the mute_new_participants setting
func (client *Client) ToggleGroupCallMuteNewParticipants(groupCallID int32, muteNewParticipants bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "toggleGroupCallMuteNewParticipants",
		"group_call_id":         groupCallID,
		"mute_new_participants": muteNewParticipants,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// InviteGroupCallParticipants Invites users to a group call. Sends a service message of type messageInviteToGroupCall for voice chats
// @param groupCallID Group call identifier
// @param userIDs User identifiers. At most 10 users can be invited simultaneously
func (client *Client) InviteGroupCallParticipants(groupCallID int32, userIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "inviteGroupCallParticipants",
		"group_call_id": groupCallID,
		"user_ids":      userIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGroupCallParticipantIsSpeaking Informs TDLib that a group call participant speaking state has changed
// @param groupCallID Group call identifier
// @param source Group call participant's synchronization source identifier, or 0 for the current user
// @param isSpeaking True, if the user is speaking
func (client *Client) SetGroupCallParticipantIsSpeaking(groupCallID int32, source int32, isSpeaking bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallParticipantIsSpeaking",
		"group_call_id": groupCallID,
		"source":        source,
		"is_speaking":   isSpeaking,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleGroupCallParticipantIsMuted Toggles whether a group call participant is muted, unmuted, or allowed to unmute themself
// @param groupCallID Group call identifier
// @param userID User identifier
// @param isMuted Pass true if the user must be muted and false otherwise
func (client *Client) ToggleGroupCallParticipantIsMuted(groupCallID int32, userID int32, isMuted bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleGroupCallParticipantIsMuted",
		"group_call_id": groupCallID,
		"user_id":       userID,
		"is_muted":      isMuted,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGroupCallParticipantVolumeLevel Changes a group call participant's volume level. If the current user can manage the group call, then the participant's volume level will be changed for all users with default volume level
// @param groupCallID Group call identifier
// @param userID User identifier
// @param volumeLevel New participant's volume level; 1-20000 in hundreds of percents
func (client *Client) SetGroupCallParticipantVolumeLevel(groupCallID int32, userID int32, volumeLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setGroupCallParticipantVolumeLevel",
		"group_call_id": groupCallID,
		"user_id":       userID,
		"volume_level":  volumeLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LoadGroupCallParticipants Loads more group call participants. The loaded participants will be received through updates. Use the field groupCall.loaded_all_participants to check whether all participants has already been loaded
// @param groupCallID Group call identifier. The group call must be previously received through getGroupCall and must be joined or being joined
// @param limit Maximum number of participants to load
func (client *Client) LoadGroupCallParticipants(groupCallID int32, limit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "loadGroupCallParticipants",
		"group_call_id": groupCallID,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LeaveGroupCall Leaves a group call
// @param groupCallID Group call identifier
func (client *Client) LeaveGroupCall(groupCallID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "leaveGroupCall",
		"group_call_id": groupCallID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DiscardGroupCall Discards a group call. Requires groupCall.can_be_managed
// @param groupCallID Group call identifier
func (client *Client) DiscardGroupCall(groupCallID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "discardGroupCall",
		"group_call_id": groupCallID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleMessageSenderIsBlocked Changes the block state of a message sender. Currently, only users and supergroup chats can be blocked
// @param sender Message Sender
// @param isBlocked New value of is_blocked
func (client *Client) ToggleMessageSenderIsBlocked(sender MessageSender, isBlocked bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "toggleMessageSenderIsBlocked",
		"sender":     sender,
		"is_blocked": isBlocked,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// BlockMessageSenderFromReplies Blocks an original sender of a message in the Replies chat
// @param messageID The identifier of an incoming message in the Replies chat
// @param deleteMessage Pass true if the message must be deleted
// @param deleteAllMessages Pass true if all messages from the same sender must be deleted
// @param reportSpam Pass true if the sender must be reported to the Telegram moderators
func (client *Client) BlockMessageSenderFromReplies(messageID int64, deleteMessage bool, deleteAllMessages bool, reportSpam bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "blockMessageSenderFromReplies",
		"message_id":          messageID,
		"delete_message":      deleteMessage,
		"delete_all_messages": deleteAllMessages,
		"report_spam":         reportSpam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetBlockedMessageSenders Returns users and chats that were blocked by the current user
// @param offset Number of users and chats to skip in the result; must be non-negative
// @param limit The maximum number of users and chats to return; up to 100
func (client *Client) GetBlockedMessageSenders(offset int32, limit int32) (*MessageSenders, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getBlockedMessageSenders",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageSenders MessageSenders
	err = json.Unmarshal(result.Raw, &messageSenders)
	return &messageSenders, err

}

// AddContact Adds a user to the contact list or edits an existing contact by their user identifier
// @param contact The contact to add or edit; phone number can be empty and needs to be specified only if known, vCard is ignored
// @param sharePhoneNumber True, if the new contact needs to be allowed to see current user's phone number. A corresponding rule to userPrivacySettingShowPhoneNumber will be added if needed. Use the field UserFullInfo.need_phone_number_privacy_exception to check whether the current user needs to be asked to share their phone number
func (client *Client) AddContact(contact *Contact, sharePhoneNumber bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "addContact",
		"contact":            contact,
		"share_phone_number": sharePhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
// @param contacts The list of contacts to import or edit; contacts' vCard are ignored and are not imported
func (client *Client) ImportContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "importContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// GetContacts Returns all user contacts
func (client *Client) GetContacts() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
// @param query Query to search for; may be empty to return all contacts
// @param limit The maximum number of users to be returned
func (client *Client) SearchContacts(query string, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchContacts",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// RemoveContacts Removes users from the contact list
// @param userIDs Identifiers of users to be deleted
func (client *Client) RemoveContacts(userIDs []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeContacts",
		"user_ids": userIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetImportedContactCount Returns the total number of imported contacts
func (client *Client) GetImportedContactCount() (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getImportedContactCount",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// ChangeImportedContacts Changes imported contacts using the list of contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts.
// @param contacts
func (client *Client) ChangeImportedContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "changeImportedContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// ClearImportedContacts Clears all imported contacts, contact list remains unchanged
func (client *Client) ClearImportedContacts() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearImportedContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SharePhoneNumber Shares the phone number of the current user with a mutual contact. Supposed to be called when the user clicks on chatActionBarSharePhoneNumber
// @param userID Identifier of the user with whom to share the phone number. The user must be a mutual contact
func (client *Client) SharePhoneNumber(userID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sharePhoneNumber",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetUserProfilePhotos Returns the profile photos of a user. The result of this query may be outdated: some photos might have been deleted already
// @param userID User identifier
// @param offset The number of photos to skip; must be non-negative
// @param limit The maximum number of photos to be returned; up to 100
func (client *Client) GetUserProfilePhotos(userID int32, offset int32, limit int32) (*ChatPhotos, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserProfilePhotos",
		"user_id": userID,
		"offset":  offset,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatPhotos ChatPhotos
	err = json.Unmarshal(result.Raw, &chatPhotos)
	return &chatPhotos, err

}

// GetStickers Returns stickers from the installed sticker sets that correspond to a given emoji. If the emoji is not empty, favorite and recently used stickers may also be returned
// @param emoji String representation of emoji. If empty, returns all known installed stickers
// @param limit The maximum number of stickers to be returned
func (client *Client) GetStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// SearchStickers Searches for stickers from public sticker sets that correspond to a given emoji
// @param emoji String representation of emoji; must be non-empty
// @param limit The maximum number of stickers to be returned
func (client *Client) SearchStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// GetInstalledStickerSets Returns a list of installed sticker sets
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
func (client *Client) GetInstalledStickerSets(isMasks bool) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getInstalledStickerSets",
		"is_masks": isMasks,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetArchivedStickerSets Returns a list of archived sticker sets
// @param isMasks Pass true to return mask stickers sets; pass false to return ordinary sticker sets
// @param offsetStickerSetID Identifier of the sticker set from which to return the result
// @param limit The maximum number of sticker sets to return
func (client *Client) GetArchivedStickerSets(isMasks bool, offsetStickerSetID JSONInt64, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "getArchivedStickerSets",
		"is_masks":              isMasks,
		"offset_sticker_set_id": offsetStickerSetID,
		"limit":                 limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetTrendingStickerSets Returns a list of trending sticker sets. For the optimal performance the number of returned sticker sets is chosen by the library
// @param offset The offset from which to return the sticker sets; must be non-negative
// @param limit The maximum number of sticker sets to be returned; must be non-negative. Fewer sticker sets may be returned than specified by the limit, even if the end of the list has not been reached
func (client *Client) GetTrendingStickerSets(offset int32, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getTrendingStickerSets",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file. Currently only photos and videos can have attached sticker sets
// @param fileID File identifier
func (client *Client) GetAttachedStickerSets(fileID int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getAttachedStickerSets",
		"file_id": fileID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetStickerSet Returns information about a sticker set by its identifier
// @param setID Identifier of the sticker set
func (client *Client) GetStickerSet(setID JSONInt64) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getStickerSet",
		"set_id": setID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchStickerSet Searches for a sticker set by its name
// @param name Name of the sticker set
func (client *Client) SearchStickerSet(name string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSet",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
// @param query Query to search for
// @param limit The maximum number of sticker sets to return
func (client *Client) SearchInstalledStickerSets(isMasks bool, query string, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchInstalledStickerSets",
		"is_masks": isMasks,
		"query":    query,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// SearchStickerSets Searches for ordinary sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
// @param query Query to search for
func (client *Client) SearchStickerSets(query string) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSets",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set
// @param setID Identifier of the sticker set
// @param isInstalled The new value of is_installed
// @param isArchived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (client *Client) ChangeStickerSet(setID JSONInt64, isInstalled bool, isArchived bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changeStickerSet",
		"set_id":       setID,
		"is_installed": isInstalled,
		"is_archived":  isArchived,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user
// @param stickerSetIDs Identifiers of viewed trending sticker sets
func (client *Client) ViewTrendingStickerSets(stickerSetIDs []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "viewTrendingStickerSets",
		"sticker_set_ids": stickerSetIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReorderInstalledStickerSets Changes the order of installed sticker sets
// @param isMasks Pass true to change the order of mask sticker sets; pass false to change the order of ordinary sticker sets
// @param stickerSetIDs Identifiers of installed sticker sets in the new correct order
func (client *Client) ReorderInstalledStickerSets(isMasks bool, stickerSetIDs []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderInstalledStickerSets",
		"is_masks":        isMasks,
		"sticker_set_ids": stickerSetIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentStickers Returns a list of recently used stickers
// @param isAttached Pass true to return stickers and masks that were recently attached to photos or video files; pass false to return recently sent stickers
func (client *Client) GetRecentStickers(isAttached bool) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddRecentSticker Manually adds a new sticker to the list of recently used stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param isAttached Pass true to add the sticker to the list of stickers recently attached to photo or video files; pass false to add the sticker to the list of recently sent stickers
// @param sticker Sticker file to add
func (client *Client) AddRecentSticker(isAttached bool, sticker InputFile) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "addRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers
// @param isAttached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers
// @param sticker Sticker file to delete
func (client *Client) RemoveRecentSticker(isAttached bool, sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "removeRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentStickers Clears the list of recently used stickers
// @param isAttached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (client *Client) ClearRecentStickers(isAttached bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "clearRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetFavoriteStickers Returns favorite stickers
func (client *Client) GetFavoriteStickers() (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getFavoriteStickers",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param sticker Sticker file to add
func (client *Client) AddFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers
// @param sticker Sticker file to delete from the list
func (client *Client) RemoveFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetStickerEmojis Returns emoji corresponding to a sticker. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object
// @param sticker Sticker file identifier
func (client *Client) GetStickerEmojis(sticker InputFile) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStickerEmojis",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err

}

// SearchEmojis Searches for emojis by keywords. Supported only if the file database is enabled
// @param text Text to search for
// @param exactMatch True, if only emojis, which exactly match text needs to be returned
// @param inputLanguageCodes List of possible IETF language tags of the user's input language; may be empty if unknown
func (client *Client) SearchEmojis(text string, exactMatch bool, inputLanguageCodes []string) (*Emojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "searchEmojis",
		"text":                 text,
		"exact_match":          exactMatch,
		"input_language_codes": inputLanguageCodes,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emojis Emojis
	err = json.Unmarshal(result.Raw, &emojis)
	return &emojis, err

}

// GetEmojiSuggestionsURL Returns an HTTP URL which can be used to automatically log in to the translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
// @param languageCode Language code for which the emoji replacements will be suggested
func (client *Client) GetEmojiSuggestionsURL(languageCode string) (*HttpURL, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getEmojiSuggestionsUrl",
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpURL HttpURL
	err = json.Unmarshal(result.Raw, &httpURL)
	return &httpURL, err

}

// GetSavedAnimations Returns saved animations
func (client *Client) GetSavedAnimations() (*Animations, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedAnimations",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var animations Animations
	err = json.Unmarshal(result.Raw, &animations)
	return &animations, err

}

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first. Only non-secret video animations with MIME type "video/mp4" can be added to the list
// @param animation The animation file to be added. Only animations known to the server (i.e. successfully sent via a message) can be added to the list
func (client *Client) AddSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveSavedAnimation Removes an animation from the list of saved animations
// @param animation Animation file to be removed
func (client *Client) RemoveSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "removeSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (client *Client) GetRecentInlineBots() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentInlineBots",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchHashtags Searches for recently used hashtags by their prefix
// @param prefix Hashtag prefix to search for
// @param limit The maximum number of hashtags to be returned
func (client *Client) SearchHashtags(prefix string, limit int32) (*Hashtags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "searchHashtags",
		"prefix": prefix,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var hashtags Hashtags
	err = json.Unmarshal(result.Raw, &hashtags)
	return &hashtags, err

}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags
// @param hashtag Hashtag to delete
func (client *Client) RemoveRecentHashtag(hashtag string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentHashtag",
		"hashtag": hashtag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetWebPagePreview Returns a web page preview by the text of the message. Do not call this function too often. Returns a 404 error if the web page has no preview
// @param text Message text with formatting
func (client *Client) GetWebPagePreview(text *FormattedText) (*WebPage, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getWebPagePreview",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPage WebPage
	err = json.Unmarshal(result.Raw, &webPage)
	return &webPage, err

}

// GetWebPageInstantView Returns an instant view version of a web page if available. Returns a 404 error if the web page has no instant view page
// @param uRL The web page URL
// @param forceFull If true, the full instant view for the web page will be returned
func (client *Client) GetWebPageInstantView(uRL string, forceFull bool) (*WebPageInstantView, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getWebPageInstantView",
		"url":        uRL,
		"force_full": forceFull,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPageInstantView WebPageInstantView
	err = json.Unmarshal(result.Raw, &webPageInstantView)
	return &webPageInstantView, err

}

// SetProfilePhoto Changes a profile photo for the current user
// @param photo Profile photo to set
func (client *Client) SetProfilePhoto(photo InputChatPhoto) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setProfilePhoto",
		"photo": photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteProfilePhoto Deletes a profile photo
// @param profilePhotoID Identifier of the profile photo to delete
func (client *Client) DeleteProfilePhoto(profilePhotoID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteProfilePhoto",
		"profile_photo_id": profilePhotoID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetName Changes the first and last name of the current user
// @param firstName The new value of the first name for the user; 1-64 characters
// @param lastName The new value of the optional last name for the user; 0-64 characters
func (client *Client) SetName(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setName",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBio Changes the bio of the current user
// @param bio The new value of the user bio; 0-70 characters without line feeds
func (client *Client) SetBio(bio string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setBio",
		"bio":   bio,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetUsername Changes the username of the current user
// @param username The new value of the username. Use an empty string to remove the username
func (client *Client) SetUsername(username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setUsername",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetLocation Changes the location of the current user. Needs to be called if GetOption("is_location_visible") is true and location changes for more than 1 kilometer
// @param location The new location of the user
func (client *Client) SetLocation(location *Location) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setLocation",
		"location": location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ChangePhoneNumber Changes the phone number of the user and sends an authentication code to the user's new phone number. On success, returns information about the sent code
// @param phoneNumber The new phone number of the user in international format
// @param settings Settings for the authentication of the user's phone number
func (client *Client) ChangePhoneNumber(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changePhoneNumber",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendChangePhoneNumberCode Re-sends the authentication code sent to confirm a new phone number for the user. Works only if the previously received authenticationCodeInfo next_code_type was not null
func (client *Client) ResendChangePhoneNumberCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendChangePhoneNumberCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckChangePhoneNumberCode Checks the authentication code sent to confirm a new phone number of the user
// @param code Verification code received by SMS, phone call or flash call
func (client *Client) CheckChangePhoneNumberCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkChangePhoneNumberCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCommands Sets the list of commands supported by the bot; for bots only
// @param commands List of the bot's commands
func (client *Client) SetCommands(commands []BotCommand) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setCommands",
		"commands": commands,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetActiveSessions Returns all active sessions of the current user
func (client *Client) GetActiveSessions() (*Sessions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sessions Sessions
	err = json.Unmarshal(result.Raw, &sessions)
	return &sessions, err

}

// TerminateSession Terminates a session of the current user
// @param sessionID Session identifier
func (client *Client) TerminateSession(sessionID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "terminateSession",
		"session_id": sessionID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (client *Client) TerminateAllOtherSessions() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "terminateAllOtherSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (client *Client) GetConnectedWebsites() (*ConnectedWebsites, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getConnectedWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var connectedWebsites ConnectedWebsites
	err = json.Unmarshal(result.Raw, &connectedWebsites)
	return &connectedWebsites, err

}

// DisconnectWebsite Disconnects website from the current user's Telegram account
// @param websiteID Website identifier
func (client *Client) DisconnectWebsite(websiteID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "disconnectWebsite",
		"website_id": websiteID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (client *Client) DisconnectAllWebsites() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disconnectAllWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupUsername Changes the username of a supergroup or channel, requires owner privileges in the supergroup or channel
// @param supergroupID Identifier of the supergroup or channel
// @param username New value of the username. Use an empty string to remove the username
func (client *Client) SetSupergroupUsername(supergroupID int32, username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setSupergroupUsername",
		"supergroup_id": supergroupID,
		"username":      username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires can_change_info administrator right
// @param supergroupID Identifier of the supergroup
// @param stickerSetID New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (client *Client) SetSupergroupStickerSet(supergroupID int32, stickerSetID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setSupergroupStickerSet",
		"supergroup_id":  supergroupID,
		"sticker_set_id": stickerSetID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupSignMessages Toggles sender signatures messages sent in a channel; requires can_change_info administrator right
// @param supergroupID Identifier of the channel
// @param signMessages New value of sign_messages
func (client *Client) ToggleSupergroupSignMessages(supergroupID int32, signMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupSignMessages",
		"supergroup_id": supergroupID,
		"sign_messages": signMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires can_change_info administrator right
// @param supergroupID The identifier of the supergroup
// @param isAllHistoryAvailable The new value of is_all_history_available
func (client *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupID int32, isAllHistoryAvailable bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "toggleSupergroupIsAllHistoryAvailable",
		"supergroup_id":            supergroupID,
		"is_all_history_available": isAllHistoryAvailable,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportSupergroupSpam Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
// @param supergroupID Supergroup identifier
// @param userID User identifier
// @param messageIDs Identifiers of messages sent in the supergroup by the user. This list must be non-empty
func (client *Client) ReportSupergroupSpam(supergroupID int32, userID int32, messageIDs []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "reportSupergroupSpam",
		"supergroup_id": supergroupID,
		"user_id":       userID,
		"message_ids":   messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if SupergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
// @param supergroupID Identifier of the supergroup or channel
// @param filter The type of users to return. By default, supergroupMembersFilterRecent
// @param offset Number of users to skip
// @param limit The maximum number of users be returned; up to 200
func (client *Client) GetSupergroupMembers(supergroupID int32, filter SupergroupMembersFilter, offset int32, limit int32) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupMembers",
		"supergroup_id": supergroupID,
		"filter":        filter,
		"offset":        offset,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// CloseSecretChat Closes a secret chat, effectively transferring its state to secretChatStateClosed
// @param secretChatID Secret chat identifier
func (client *Client) CloseSecretChat(secretChatID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "closeSecretChat",
		"secret_chat_id": secretChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only for supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i. e., in order of decreasing event_id)
// @param chatID Chat identifier
// @param query Search query by which to filter events
// @param fromEventID Identifier of an event from which to return results. Use 0 to get results from the latest events
// @param limit The maximum number of events to return; up to 100
// @param filters The types of events to return. By default, all types will be returned
// @param userIDs User identifiers by which to filter events. By default, events relating to all users will be returned
func (client *Client) GetChatEventLog(chatID int64, query string, fromEventID JSONInt64, limit int32, filters *ChatEventLogFilters, userIDs []int32) (*ChatEvents, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatEventLog",
		"chat_id":       chatID,
		"query":         query,
		"from_event_id": fromEventID,
		"limit":         limit,
		"filters":       filters,
		"user_ids":      userIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatEvents ChatEvents
	err = json.Unmarshal(result.Raw, &chatEvents)
	return &chatEvents, err

}

// GetPaymentForm Returns an invoice payment form. This method should be called when the user presses inlineKeyboardButtonBuy
// @param chatID Chat identifier of the Invoice message
// @param messageID Message identifier
func (client *Client) GetPaymentForm(chatID int64, messageID int64) (*PaymentForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentForm",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentForm PaymentForm
	err = json.Unmarshal(result.Raw, &paymentForm)
	return &paymentForm, err

}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
// @param chatID Chat identifier of the Invoice message
// @param messageID Message identifier
// @param orderInfo The order information, provided by the user
// @param allowSave True, if the order information can be saved
func (client *Client) ValidateOrderInfo(chatID int64, messageID int64, orderInfo *OrderInfo, allowSave bool) (*ValidatedOrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "validateOrderInfo",
		"chat_id":    chatID,
		"message_id": messageID,
		"order_info": orderInfo,
		"allow_save": allowSave,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var validatedOrderInfo ValidatedOrderInfo
	err = json.Unmarshal(result.Raw, &validatedOrderInfo)
	return &validatedOrderInfo, err

}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
// @param chatID Chat identifier of the Invoice message
// @param messageID Message identifier
// @param orderInfoID Identifier returned by ValidateOrderInfo, or an empty string
// @param shippingOptionID Identifier of a chosen shipping option, if applicable
// @param credentials The credentials chosen by user for payment
func (client *Client) SendPaymentForm(chatID int64, messageID int64, orderInfoID string, shippingOptionID string, credentials InputCredentials) (*PaymentResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "sendPaymentForm",
		"chat_id":            chatID,
		"message_id":         messageID,
		"order_info_id":      orderInfoID,
		"shipping_option_id": shippingOptionID,
		"credentials":        credentials,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentResult PaymentResult
	err = json.Unmarshal(result.Raw, &paymentResult)
	return &paymentResult, err

}

// GetPaymentReceipt Returns information about a successful payment
// @param chatID Chat identifier of the PaymentSuccessful message
// @param messageID Message identifier
func (client *Client) GetPaymentReceipt(chatID int64, messageID int64) (*PaymentReceipt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentReceipt",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentReceipt PaymentReceipt
	err = json.Unmarshal(result.Raw, &paymentReceipt)
	return &paymentReceipt, err

}

// GetSavedOrderInfo Returns saved order info, if any
func (client *Client) GetSavedOrderInfo() (*OrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var orderInfo OrderInfo
	err = json.Unmarshal(result.Raw, &orderInfo)
	return &orderInfo, err

}

// DeleteSavedOrderInfo Deletes saved order info
func (client *Client) DeleteSavedOrderInfo() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (client *Client) DeleteSavedCredentials() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedCredentials",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupportUser Returns a user that can be contacted to get support
func (client *Client) GetSupportUser() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSupportUser",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

}

// GetBackgrounds Returns backgrounds installed by the user
// @param forDarkTheme True, if the backgrounds must be ordered for dark theme
func (client *Client) GetBackgrounds(forDarkTheme bool) (*Backgrounds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBackgrounds",
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgrounds Backgrounds
	err = json.Unmarshal(result.Raw, &backgrounds)
	return &backgrounds, err

}

// GetBackgroundURL Constructs a persistent HTTP URL for a background
// @param name Background name
// @param typeParam Background type
func (client *Client) GetBackgroundURL(name string, typeParam BackgroundType) (*HttpURL, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getBackgroundUrl",
		"name":  name,
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpURL HttpURL
	err = json.Unmarshal(result.Raw, &httpURL)
	return &httpURL, err

}

// SearchBackground Searches for a background by its name
// @param name The name of the background
func (client *Client) SearchBackground(name string) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchBackground",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var background Background
	err = json.Unmarshal(result.Raw, &background)
	return &background, err

}

// SetBackground Changes the background selected by the user; adds background to the list of installed backgrounds
// @param background The input background to use, null for filled backgrounds
// @param typeParam Background type; null for default background. The method will return error 404 if type is null
// @param forDarkTheme True, if the background is chosen for dark theme
func (client *Client) SetBackground(background InputBackground, typeParam BackgroundType, forDarkTheme bool) (*Background, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setBackground",
		"background":     background,
		"type":           typeParam,
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgroundDummy Background
	err = json.Unmarshal(result.Raw, &backgroundDummy)
	return &backgroundDummy, err

}

// RemoveBackground Removes background from the list of installed backgrounds
// @param backgroundID The background identifier
func (client *Client) RemoveBackground(backgroundID JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "removeBackground",
		"background_id": backgroundID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetBackgrounds Resets list of installed backgrounds to its default value
func (client *Client) ResetBackgrounds() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetBackgrounds",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline request if only_local is true. Can be called before authorization
// @param onlyLocal If true, returns only locally available information without sending network requests
func (client *Client) GetLocalizationTargetInfo(onlyLocal bool) (*LocalizationTargetInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLocalizationTargetInfo",
		"only_local": onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var localizationTargetInfo LocalizationTargetInfo
	err = json.Unmarshal(result.Raw, &localizationTargetInfo)
	return &localizationTargetInfo, err

}

// GetLanguagePackInfo Returns information about a language pack. Returned language pack identifier may be different from a provided one. Can be called before authorization
// @param languagePackID Language pack identifier
func (client *Client) GetLanguagePackInfo(languagePackID string) (*LanguagePackInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackInfo",
		"language_pack_id": languagePackID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackInfo LanguagePackInfo
	err = json.Unmarshal(result.Raw, &languagePackInfo)
	return &languagePackInfo, err

}

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys. Can be called before authorization
// @param languagePackID Language pack identifier of the strings to be returned
// @param keys Language pack keys of the strings to be returned; leave empty to request all available strings
func (client *Client) GetLanguagePackStrings(languagePackID string, keys []string) (*LanguagePackStrings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackStrings",
		"language_pack_id": languagePackID,
		"keys":             keys,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackStrings LanguagePackStrings
	err = json.Unmarshal(result.Raw, &languagePackStrings)
	return &languagePackStrings, err

}

// SynchronizeLanguagePack Fetches the latest versions of all strings from a language pack in the current localization target from the server. This method shouldn't be called explicitly for the current used/base language packs. Can be called before authorization
// @param languagePackID Language pack identifier
func (client *Client) SynchronizeLanguagePack(languagePackID string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "synchronizeLanguagePack",
		"language_pack_id": languagePackID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddCustomServerLanguagePack Adds a custom server language pack to the list of installed language packs in current localization target. Can be called before authorization
// @param languagePackID Identifier of a language pack to be added; may be different from a name that is used in an "https://t.me/setlanguage/" link
func (client *Client) AddCustomServerLanguagePack(languagePackID string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "addCustomServerLanguagePack",
		"language_pack_id": languagePackID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCustomLanguagePack Adds or changes a custom local language pack to the current localization target
// @param info Information about the language pack. Language pack ID must start with 'X', consist only of English letters, digits and hyphens, and must not exceed 64 characters. Can be called before authorization
// @param strings Strings of the new language pack
func (client *Client) SetCustomLanguagePack(info *LanguagePackInfo, strings []LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setCustomLanguagePack",
		"info":    info,
		"strings": strings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditCustomLanguagePackInfo Edits information about a custom local language pack in the current localization target. Can be called before authorization
// @param info New information about the custom local language pack
func (client *Client) EditCustomLanguagePackInfo(info *LanguagePackInfo) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "editCustomLanguagePackInfo",
		"info":  info,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom local language pack. Can be called before authorization
// @param languagePackID Identifier of a previously added custom local language pack in the current localization target
// @param newString New language pack string
func (client *Client) SetCustomLanguagePackString(languagePackID string, newString *LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "setCustomLanguagePackString",
		"language_pack_id": languagePackID,
		"new_string":       newString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack which is currently in use (including base language pack) or is being synchronized can't be deleted. Can be called before authorization
// @param languagePackID Identifier of the language pack to delete
func (client *Client) DeleteLanguagePack(languagePackID string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteLanguagePack",
		"language_pack_id": languagePackID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription
// @param deviceToken Device token
// @param otherUserIDs List of user identifiers of other users currently using the application
func (client *Client) RegisterDevice(deviceToken DeviceToken, otherUserIDs []int32) (*PushReceiverID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverID PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}

// ProcessPushNotification Handles a push notification. Returns error with code 406 if the push notification is not supported and connection to the server is required to fetch new data. Can be called before authorization
// @param payload JSON-encoded push notification payload with all fields sent by the server, and "google.sent_time" and "google.notification.sound" fields added
func (client *Client) ProcessPushNotification(payload string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "processPushNotification",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPushReceiverID Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously
// @param payload JSON-encoded push notification payload
func (client *Client) GetPushReceiverID(payload string) (*PushReceiverID, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getPushReceiverId",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pushReceiverID PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}

// GetRecentlyVisitedTMeURLs Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeURLs(referrer string) (*TMeURLs, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tMeURLs TMeURLs
	err = json.Unmarshal(result.Raw, &tMeURLs)
	return &tMeURLs, err

}

// SetUserPrivacySettingRules Changes user privacy settings
// @param setting The privacy setting
// @param rules The new privacy rules
func (client *Client) SetUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setUserPrivacySettingRules",
		"setting": setting,
		"rules":   rules,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetUserPrivacySettingRules Returns the current privacy settings
// @param setting The privacy setting
func (client *Client) GetUserPrivacySettingRules(setting UserPrivacySetting) (*UserPrivacySettingRules, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserPrivacySettingRules",
		"setting": setting,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userPrivacySettingRules UserPrivacySettingRules
	err = json.Unmarshal(result.Raw, &userPrivacySettingRules)
	return &userPrivacySettingRules, err

}

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization
// @param name The name of the option
func (client *Client) GetOption(name string) (OptionValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getOption",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch OptionValueEnum(result.Data["@type"].(string)) {

	case OptionValueBooleanType:
		var optionValue OptionValueBoolean
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueEmptyType:
		var optionValue OptionValueEmpty
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueIntegerType:
		var optionValue OptionValueInteger
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueStringType:
		var optionValue OptionValueString
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
// @param name The name of the option
// @param value The new value of the option
func (client *Client) SetOption(name string, value OptionValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setOption",
		"name":  name,
		"value": value,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAccountTTL Changes the period of inactivity after which the account of the current user will automatically be deleted
// @param tTL New account TTL
func (client *Client) SetAccountTTL(tTL *AccountTTL) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setAccountTtl",
		"ttl":   tTL,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetAccountTTL Returns the period of inactivity after which the account of the current user will automatically be deleted
func (client *Client) GetAccountTTL() (*AccountTTL, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAccountTtl",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountTTL AccountTTL
	err = json.Unmarshal(result.Raw, &accountTTL)
	return &accountTTL, err

}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account. Can be called before authorization when the current authorization state is authorizationStateWaitPassword
// @param reason The reason why the account was deleted; optional
func (client *Client) DeleteAccount(reason string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAccount",
		"reason": reason,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveChatActionBar Removes a chat action bar without any other action
// @param chatID Chat identifier
func (client *Client) RemoveChatActionBar(chatID int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeChatActionBar",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportChat Reports a chat to the Telegram moderators. A chat can be reported only from the chat action bar, or if this is a private chat with a bot, a private chat with a user sharing their location, a supergroup, or a channel, since other chats can't be checked by moderators
// @param chatID Chat identifier
// @param reason The reason for reporting the chat
// @param messageIDs Identifiers of reported messages, if any
func (client *Client) ReportChat(chatID int64, reason ChatReportReason, messageIDs []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "reportChat",
		"chat_id":     chatID,
		"reason":      reason,
		"message_ids": messageIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatStatisticsURL Returns an HTTP URL with the chat statistics. Currently this method of getting the statistics are disabled and can be deleted in the future
// @param chatID Chat identifier
// @param parameters Parameters from "tg://statsrefresh?params=******" link
// @param isDark Pass true if a URL with the dark theme must be returned
func (client *Client) GetChatStatisticsURL(chatID int64, parameters string, isDark bool) (*HttpURL, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getChatStatisticsUrl",
		"chat_id":    chatID,
		"parameters": parameters,
		"is_dark":    isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var httpURL HttpURL
	err = json.Unmarshal(result.Raw, &httpURL)
	return &httpURL, err

}

// GetChatStatistics Returns detailed statistics about a chat. Currently this method can be used only for supergroups and channels. Can be used only if SupergroupFullInfo.can_get_statistics == true
// @param chatID Chat identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetChatStatistics(chatID int64, isDark bool) (ChatStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatStatistics",
		"chat_id": chatID,
		"is_dark": isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch ChatStatisticsEnum(result.Data["@type"].(string)) {

	case ChatStatisticsSupergroupType:
		var chatStatistics ChatStatisticsSupergroup
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	case ChatStatisticsChannelType:
		var chatStatistics ChatStatisticsChannel
		err = json.Unmarshal(result.Raw, &chatStatistics)
		return &chatStatistics, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetMessageStatistics Returns detailed statistics about a message. Can be used only if Message.can_get_statistics == true
// @param chatID Chat identifier
// @param messageID Message identifier
// @param isDark Pass true if a dark theme is used by the application
func (client *Client) GetMessageStatistics(chatID int64, messageID int64, isDark bool) (*MessageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessageStatistics",
		"chat_id":    chatID,
		"message_id": messageID,
		"is_dark":    isDark,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageStatistics MessageStatistics
	err = json.Unmarshal(result.Raw, &messageStatistics)
	return &messageStatistics, err

}

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph
// @param chatID Chat identifier
// @param token The token for graph loading
// @param x X-value for zoomed in graph or 0 otherwise
func (client *Client) GetStatisticalGraph(chatID int64, token string, x int64) (StatisticalGraph, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStatisticalGraph",
		"chat_id": chatID,
		"token":   token,
		"x":       x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch StatisticalGraphEnum(result.Data["@type"].(string)) {

	case StatisticalGraphDataType:
		var statisticalGraph StatisticalGraphData
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphAsyncType:
		var statisticalGraph StatisticalGraphAsync
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case StatisticalGraphErrorType:
		var statisticalGraph StatisticalGraphError
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetStorageStatistics Returns storage usage statistics. Can be called before authorization
// @param chatLimit The maximum number of chats with the largest storage usage for which separate statistics should be returned. All other chats will be grouped in entries with chat_id == 0. If the chat info database is not used, the chat_limit is ignored and is always set to 0
func (client *Client) GetStorageStatistics(chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getStorageStatistics",
		"chat_limit": chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics. Can be called before authorization
func (client *Client) GetStorageStatisticsFast() (*StorageStatisticsFast, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStorageStatisticsFast",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatisticsFast StorageStatisticsFast
	err = json.Unmarshal(result.Raw, &storageStatisticsFast)
	return &storageStatisticsFast, err

}

// GetDatabaseStatistics Returns database statistics
func (client *Client) GetDatabaseStatistics() (*DatabaseStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDatabaseStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var databaseStatistics DatabaseStatistics
	err = json.Unmarshal(result.Raw, &databaseStatistics)
	return &databaseStatistics, err

}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
// @param size Limit on the total size of files after deletion. Pass -1 to use the default limit
// @param tTL Limit on the time that has passed since the last time a file was accessed (or creation time for some filesystems). Pass -1 to use the default limit
// @param count Limit on the total count of files after deletion. Pass -1 to use the default limit
// @param immunityDelay The amount of time after the creation of a file during which it can't be deleted, in seconds. Pass -1 to use the default value
// @param fileTypes If not empty, only files with the given type(s) are considered. By default, all types except thumbnails, profile photos, stickers and wallpapers are deleted
// @param chatIDs If not empty, only files from the given chats are considered. Use 0 as chat identifier to delete files not belonging to any chat (e.g., profile photos)
// @param excludeChatIDs If not empty, files from the given chats are excluded. Use 0 as chat identifier to exclude all files not belonging to any chat (e.g., profile photos)
// @param returnDeletedFileStatistics Pass true if statistics about the files that were deleted must be returned instead of the whole storage usage statistics. Affects only returned statistics
// @param chatLimit Same as in getStorageStatistics. Affects only returned statistics
func (client *Client) OptimizeStorage(size int64, tTL int32, count int32, immunityDelay int32, fileTypes []FileType, chatIDs []int64, excludeChatIDs []int64, returnDeletedFileStatistics bool, chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                          "optimizeStorage",
		"size":                           size,
		"ttl":                            tTL,
		"count":                          count,
		"immunity_delay":                 immunityDelay,
		"file_types":                     fileTypes,
		"chat_ids":                       chatIDs,
		"exclude_chat_ids":               excludeChatIDs,
		"return_deleted_file_statistics": returnDeletedFileStatistics,
		"chat_limit":                     chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks, so it should be called whenever the network is changed, even if the network type remains the same.
// @param typeParam
func (client *Client) SetNetworkType(typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setNetworkType",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization
// @param onlyCurrent If true, returns only data for the current library launch
func (client *Client) GetNetworkStatistics(onlyCurrent bool) (*NetworkStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getNetworkStatistics",
		"only_current": onlyCurrent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var networkStatistics NetworkStatistics
	err = json.Unmarshal(result.Raw, &networkStatistics)
	return &networkStatistics, err

}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization
// @param entry The network statistics entry with the data to be added to statistics
func (client *Client) AddNetworkStatistics(entry NetworkStatisticsEntry) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "addNetworkStatistics",
		"entry": entry,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (client *Client) ResetNetworkStatistics() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetNetworkStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetAutoDownloadSettingsPresets Returns auto-download settings presets for the current user
func (client *Client) GetAutoDownloadSettingsPresets() (*AutoDownloadSettingsPresets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAutoDownloadSettingsPresets",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var autoDownloadSettingsPresets AutoDownloadSettingsPresets
	err = json.Unmarshal(result.Raw, &autoDownloadSettingsPresets)
	return &autoDownloadSettingsPresets, err

}

// SetAutoDownloadSettings Sets auto-download settings
// @param settings New user auto-download settings
// @param typeParam Type of the network for which the new settings are applied
func (client *Client) SetAutoDownloadSettings(settings *AutoDownloadSettings, typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setAutoDownloadSettings",
		"settings": settings,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetBankCardInfo Returns information about a bank card
// @param bankCardNumber The bank card number
func (client *Client) GetBankCardInfo(bankCardNumber string) (*BankCardInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getBankCardInfo",
		"bank_card_number": bankCardNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var bankCardInfo BankCardInfo
	err = json.Unmarshal(result.Raw, &bankCardInfo)
	return &bankCardInfo, err

}

// GetPassportElement Returns one of the available Telegram Passport elements
// @param typeParam Telegram Passport element type
// @param password Password of the current user
func (client *Client) GetPassportElement(typeParam PassportElementType, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getPassportElement",
		"type":     typeParam,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIDentityCardType:
		var passportElement PassportElementIDentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetAllPassportElements Returns all available Telegram Passport elements
// @param password Password of the current user
func (client *Client) GetAllPassportElements(password string) (*PassportElements, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getAllPassportElements",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElements PassportElements
	err = json.Unmarshal(result.Raw, &passportElements)
	return &passportElements, err

}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
// @param element Input Telegram Passport element
// @param password Password of the current user
func (client *Client) SetPassportElement(element InputPassportElement, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setPassportElement",
		"element":  element,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIDentityCardType:
		var passportElement PassportElementIDentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// DeletePassportElement Deletes a Telegram Passport element
// @param typeParam Element type
func (client *Client) DeletePassportElement(typeParam PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deletePassportElement",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed
// @param userID User identifier
// @param errors The errors
func (client *Client) SetPassportElementErrors(userID int32, errors []InputPassportElementError) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setPassportElementErrors",
		"user_id": userID,
		"errors":  errors,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which should be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
func (client *Client) GetPreferredCountryLanguage(countryCode string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getPreferredCountryLanguage",
		"country_code": countryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SendPhoneNumberVerificationCode Sends a code to verify a phone number to be added to a user's Telegram Passport
// @param phoneNumber The phone number of the user, in international format
// @param settings Settings for the authentication of the user's phone number
func (client *Client) SendPhoneNumberVerificationCode(phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "sendPhoneNumberVerificationCode",
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberVerificationCode Re-sends the code to verify a phone number to be added to a user's Telegram Passport
func (client *Client) ResendPhoneNumberVerificationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberVerificationCode Checks the phone number verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckPhoneNumberVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport
// @param emailAddress Email address
func (client *Client) SendEmailAddressVerificationCode(emailAddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "sendEmailAddressVerificationCode",
		"email_address": emailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// ResendEmailAddressVerificationCode Re-sends the code to verify an email address to be added to a user's Telegram Passport
func (client *Client) ResendEmailAddressVerificationCode() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendEmailAddressVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckEmailAddressVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkEmailAddressVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPassportAuthorizationForm Returns a Telegram Passport authorization form for sharing data with a service
// @param botUserID User identifier of the service's bot
// @param scope Telegram Passport element types requested by the service
// @param publicKey Service's public_key
// @param nonce Authorization form nonce provided by the service
func (client *Client) GetPassportAuthorizationForm(botUserID int32, scope string, publicKey string, nonce string) (*PassportAuthorizationForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getPassportAuthorizationForm",
		"bot_user_id": botUserID,
		"scope":       scope,
		"public_key":  publicKey,
		"nonce":       nonce,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportAuthorizationForm PassportAuthorizationForm
	err = json.Unmarshal(result.Raw, &passportAuthorizationForm)
	return &passportAuthorizationForm, err

}

// GetPassportAuthorizationFormAvailableElements Returns already available Telegram Passport elements suitable for completing a Telegram Passport authorization form. Result can be received only once for each authorization form
// @param autorizationFormID Authorization form identifier
// @param password Password of the current user
func (client *Client) GetPassportAuthorizationFormAvailableElements(autorizationFormID int32, password string) (*PassportElementsWithErrors, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "getPassportAuthorizationFormAvailableElements",
		"autorization_form_id": autorizationFormID,
		"password":             password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElementsWithErrors PassportElementsWithErrors
	err = json.Unmarshal(result.Raw, &passportElementsWithErrors)
	return &passportElementsWithErrors, err

}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service. This method must be called after getPassportAuthorizationFormAvailableElements if some previously available elements are going to be reused
// @param autorizationFormID Authorization form identifier
// @param typeParams Types of Telegram Passport elements chosen by user to complete the authorization form
func (client *Client) SendPassportAuthorizationForm(autorizationFormID int32, typeParams []PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "sendPassportAuthorizationForm",
		"autorization_form_id": autorizationFormID,
		"types":                typeParams,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendPhoneNumberConfirmationCode Sends phone number confirmation code. Should be called when user presses "https://t.me/confirmphone?phone=*******&hash=**********" or "tg://confirmphone?phone=*******&hash=**********" link
// @param hash Value of the "hash" parameter from the link
// @param phoneNumber Value of the "phone" parameter from the link
// @param settings Settings for the authentication of the user's phone number
func (client *Client) SendPhoneNumberConfirmationCode(hash string, phoneNumber string, settings *PhoneNumberAuthenticationSettings) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "sendPhoneNumberConfirmationCode",
		"hash":         hash,
		"phone_number": phoneNumber,
		"settings":     settings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberConfirmationCode Resends phone number confirmation code
func (client *Client) ResendPhoneNumberConfirmationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberConfirmationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberConfirmationCode Checks phone number confirmation code
// @param code The phone number confirmation code
func (client *Client) CheckPhoneNumberConfirmationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberConfirmationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
// @param pendingUpdateCount The number of pending updates
// @param errorMessage The last error message
func (client *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "setBotUpdatesStatus",
		"pending_update_count": pendingUpdateCount,
		"error_message":        errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UploadStickerFile Uploads a PNG image with a sticker; for bots only; returns the uploaded file
// @param userID Sticker file owner
// @param pngSticker PNG image with the sticker; must be up to 512 KB in size and fit in 512x512 square
func (client *Client) UploadStickerFile(userID int32, pngSticker InputFile) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "uploadStickerFile",
		"user_id":     userID,
		"png_sticker": pngSticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// CreateNewStickerSet Creates a new sticker set; for bots only. Returns the newly created sticker set
// @param userID Sticker set owner
// @param title Sticker set title; 1-64 characters
// @param name Sticker set name. Can contain only English letters, digits and underscores. Must end with *"_by_<bot username>"* (*<bot_username>* is case insensitive); 1-64 characters
// @param isMasks True, if stickers are masks. Animated stickers can't be masks
// @param stickers List of stickers to be added to the set; must be non-empty. All stickers must be of the same type
func (client *Client) CreateNewStickerSet(userID int32, title string, name string, isMasks bool, stickers []InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewStickerSet",
		"user_id":  userID,
		"title":    title,
		"name":     name,
		"is_masks": isMasks,
		"stickers": stickers,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// AddStickerToSet Adds a new sticker to a set; for bots only. Returns the sticker set
// @param userID Sticker set owner
// @param name Sticker set name
// @param sticker Sticker to add to the set
func (client *Client) AddStickerToSet(userID int32, name string, sticker InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addStickerToSet",
		"user_id": userID,
		"name":    name,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerSetThumbnail Sets a sticker set thumbnail; for bots only. Returns the sticker set
// @param userID Sticker set owner
// @param name Sticker set name
// @param thumbnail Thumbnail to set in PNG or TGS format. Animated thumbnail must be set for animated sticker sets and only for them. Pass a zero InputFileId to delete the thumbnail
func (client *Client) SetStickerSetThumbnail(userID int32, name string, thumbnail InputFile) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "setStickerSetThumbnail",
		"user_id":   userID,
		"name":      name,
		"thumbnail": thumbnail,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
// @param position New position of the sticker in the set, zero-based
func (client *Client) SetStickerPositionInSet(sticker InputFile, position int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setStickerPositionInSet",
		"sticker":  sticker,
		"position": position,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
func (client *Client) RemoveStickerFromSet(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeStickerFromSet",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
// @param location Location of the map center
// @param zoom Map zoom level; 13-20
// @param width Map width in pixels before applying scale; 16-1024
// @param height Map height in pixels before applying scale; 16-1024
// @param scale Map scale; 1-3
// @param chatID Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
func (client *Client) GetMapThumbnailFile(location *Location, zoom int32, width int32, height int32, scale int32, chatID int64) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getMapThumbnailFile",
		"location": location,
		"zoom":     zoom,
		"width":    width,
		"height":   height,
		"scale":    scale,
		"chat_id":  chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// AcceptTermsOfService Accepts Telegram terms of services
// @param termsOfServiceID Terms of service identifier
func (client *Client) AcceptTermsOfService(termsOfServiceID string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "acceptTermsOfService",
		"terms_of_service_id": termsOfServiceID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCustomRequest Sends a custom request; for bots only
// @param method The method name
// @param parameters JSON-serialized method parameters
func (client *Client) SendCustomRequest(method string, parameters string) (*CustomRequestResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "sendCustomRequest",
		"method":     method,
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var customRequestResult CustomRequestResult
	err = json.Unmarshal(result.Raw, &customRequestResult)
	return &customRequestResult, err

}

// AnswerCustomQuery Answers a custom query; for bots only
// @param customQueryID Identifier of a custom query
// @param data JSON-serialized answer to the query
func (client *Client) AnswerCustomQuery(customQueryID JSONInt64, data string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "answerCustomQuery",
		"custom_query_id": customQueryID,
		"data":            data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before initialization
// @param seconds Number of seconds before the function returns
func (client *Client) SetAlarm(seconds float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setAlarm",
		"seconds": seconds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCountries Returns information about existing countries. Can be called before authorization
func (client *Client) GetCountries() (*Countries, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountries",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var countries Countries
	err = json.Unmarshal(result.Raw, &countries)
	return &countries, err

}

// GetCountryCode Uses the current IP address to find the current country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (client *Client) GetCountryCode() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountryCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetPhoneNumberInfo Returns information about a phone number by its prefix. Can be called before authorization
// @param phoneNumberPrefix The phone number prefix
func (client *Client) GetPhoneNumberInfo(phoneNumberPrefix string) (*PhoneNumberInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "getPhoneNumberInfo",
		"phone_number_prefix": phoneNumberPrefix,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var phoneNumberInfo PhoneNumberInfo
	err = json.Unmarshal(result.Raw, &phoneNumberInfo)
	return &phoneNumberInfo, err

}

// GetInviteText Returns the default text for invitation messages to be used as a placeholder when the current user invites friends to Telegram
func (client *Client) GetInviteText() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInviteText",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization
// @param link The link
func (client *Client) GetDeepLinkInfo(link string) (*DeepLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDeepLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var deepLinkInfo DeepLinkInfo
	err = json.Unmarshal(result.Raw, &deepLinkInfo)
	return &deepLinkInfo, err

}

// GetApplicationConfig Returns application config, provided by the server. Can be called before authorization
func (client *Client) GetApplicationConfig() (JsonValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getApplicationConfig",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch JsonValueEnum(result.Data["@type"].(string)) {

	case JsonValueNullType:
		var jsonValue JsonValueNull
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueBooleanType:
		var jsonValue JsonValueBoolean
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueNumberType:
		var jsonValue JsonValueNumber
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueStringType:
		var jsonValue JsonValueString
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueArrayType:
		var jsonValue JsonValueArray
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	case JsonValueObjectType:
		var jsonValue JsonValueObject
		err = json.Unmarshal(result.Raw, &jsonValue)
		return &jsonValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SaveApplicationLogEvent Saves application log event on the server. Can be called before authorization
// @param typeParam Event type
// @param chatID Optional chat identifier, associated with the event
// @param data The log event data
func (client *Client) SaveApplicationLogEvent(typeParam string, chatID int64, data JsonValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "saveApplicationLogEvent",
		"type":    typeParam,
		"chat_id": chatID,
		"data":    data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy should be enabled
// @param typeParam Proxy type
func (client *Client) AddProxy(server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "addProxy",
		"server": server,
		"port":   port,
		"enable": enable,
		"type":   typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxy Proxy
	err = json.Unmarshal(result.Raw, &proxy)
	return &proxy, err

}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
// @param proxyID Proxy identifier
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy should be enabled
// @param typeParam Proxy type
func (client *Client) EditProxy(proxyID int32, server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "editProxy",
		"proxy_id": proxyID,
		"server":   server,
		"port":     port,
		"enable":   enable,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxyDummy Proxy
	err = json.Unmarshal(result.Raw, &proxyDummy)
	return &proxyDummy, err

}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization
// @param proxyID Proxy identifier
func (client *Client) EnableProxy(proxyID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "enableProxy",
		"proxy_id": proxyID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (client *Client) DisableProxy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disableProxy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveProxy Removes a proxy server. Can be called before authorization
// @param proxyID Proxy identifier
func (client *Client) RemoveProxy(proxyID int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeProxy",
		"proxy_id": proxyID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetProxies Returns list of proxies that are currently set up. Can be called before authorization
func (client *Client) GetProxies() (*Proxies, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getProxies",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxies Proxies
	err = json.Unmarshal(result.Raw, &proxies)
	return &proxies, err

}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization
// @param proxyID Proxy identifier
func (client *Client) GetProxyLink(proxyID int32) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getProxyLink",
		"proxy_id": proxyID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization
// @param proxyID Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (client *Client) PingProxy(proxyID int32) (*Seconds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "pingProxy",
		"proxy_id": proxyID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var seconds Seconds
	err = json.Unmarshal(result.Raw, &seconds)
	return &seconds, err

}

// SetLogStream Sets new log stream for internal logging of TDLib. Can be called synchronously
// @param logStream New log stream
func (client *Client) SetLogStream(logStream LogStream) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setLogStream",
		"log_stream": logStream,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogStream Returns information about currently used log stream for internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogStream() (LogStream, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogStream",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LogStreamEnum(result.Data["@type"].(string)) {

	case LogStreamDefaultType:
		var logStream LogStreamDefault
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamFileType:
		var logStream LogStreamFile
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamEmptyType:
		var logStream LogStreamEmpty
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib. Can be called synchronously
// @param newVerbosityLevel New value of the verbosity level for logging. Value 0 corresponds to fatal errors, value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings, value 3 corresponds to informational, value 4 corresponds to debug, value 5 corresponds to verbose debug, value greater than 5 and up to 1023 can be used to enable even more logging
func (client *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of TDLib. Can be called synchronously
func (client *Client) GetLogVerbosityLevel() (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogVerbosityLevel",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// GetLogTags Returns list of available TDLib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. Can be called synchronously
func (client *Client) GetLogTags() (*LogTags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTags",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logTags LogTags
	err = json.Unmarshal(result.Raw, &logTags)
	return &logTags, err

}

// SetLogTagVerbosityLevel Sets the verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
// @param newVerbosityLevel New verbosity level; 1-1024
func (client *Client) SetLogTagVerbosityLevel(tag string, newVerbosityLevel int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "setLogTagVerbosityLevel",
		"tag":                 tag,
		"new_verbosity_level": newVerbosityLevel,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetLogTagVerbosityLevel Returns current verbosity level for a specified TDLib internal log tag. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) GetLogTagVerbosityLevel(tag string) (*LogVerbosityLevel, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getLogTagVerbosityLevel",
		"tag":   tag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// AddLogMessage Adds a message to TDLib internal log. Can be called synchronously
// @param verbosityLevel The minimum verbosity level needed for the message to be logged; 0-1023
// @param text Text of a message to log
func (client *Client) AddLogMessage(verbosityLevel int32, text string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "addLogMessage",
		"verbosity_level": verbosityLevel,
		"text":            text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestCallEmpty Does nothing; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestCallEmpty() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallEmpty",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization
// @param x String to return
func (client *Client) TestCallString(x string) (*TestString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testString TestString
	err = json.Unmarshal(result.Raw, &testString)
	return &testString, err

}

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization
// @param x Bytes to return
func (client *Client) TestCallBytes(x []byte) (*TestBytes, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallBytes",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testBytes TestBytes
	err = json.Unmarshal(result.Raw, &testBytes)
	return &testBytes, err

}

// TestCallVectorInt Returns the received vector of numbers; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of numbers to return
func (client *Client) TestCallVectorInt(x []int32) (*TestVectorInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorInt TestVectorInt
	err = json.Unmarshal(result.Raw, &testVectorInt)
	return &testVectorInt, err

}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorIntObject(x []TestInt) (*TestVectorIntObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorIntObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorIntObject TestVectorIntObject
	err = json.Unmarshal(result.Raw, &testVectorIntObject)
	return &testVectorIntObject, err

}

// TestCallVectorString Returns the received vector of strings; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of strings to return
func (client *Client) TestCallVectorString(x []string) (*TestVectorString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorString TestVectorString
	err = json.Unmarshal(result.Raw, &testVectorString)
	return &testVectorString, err

}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only. This is an offline method. Can be called before authorization
// @param x Vector of objects to return
func (client *Client) TestCallVectorStringObject(x []TestString) (*TestVectorStringObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorStringObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorStringObject TestVectorStringObject
	err = json.Unmarshal(result.Raw, &testVectorStringObject)
	return &testVectorStringObject, err

}

// TestSquareInt Returns the squared received number; for testing only. This is an offline method. Can be called before authorization
// @param x Number to square
func (client *Client) TestSquareInt(x int32) (*TestInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testSquareInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testInt TestInt
	err = json.Unmarshal(result.Raw, &testInt)
	return &testInt, err

}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only. Can be called before authorization
func (client *Client) TestNetwork() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testNetwork",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestProxy Sends a simple network request to the Telegram servers via proxy; for testing only. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param typeParam Proxy type
// @param dcID Identifier of a datacenter, with which to test connection
// @param timeout The maximum overall timeout for the request
func (client *Client) TestProxy(server string, port int32, typeParam ProxyType, dcID int32, timeout float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "testProxy",
		"server":  server,
		"port":    port,
		"type":    typeParam,
		"dc_id":   dcID,
		"timeout": timeout,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (client *Client) TestGetDifference() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testGetDifference",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only. This is an offline method. Can be called before authorization
func (client *Client) TestUseUpdate() (Update, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testUseUpdate",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch UpdateEnum(result.Data["@type"].(string)) {

	case UpdateAuthorizationStateType:
		var update UpdateAuthorizationState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewMessageType:
		var update UpdateNewMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendAcknowledgedType:
		var update UpdateMessageSendAcknowledged
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendSucceededType:
		var update UpdateMessageSendSucceeded
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendFailedType:
		var update UpdateMessageSendFailed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentType:
		var update UpdateMessageContent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageEditedType:
		var update UpdateMessageEdited
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageIsPinnedType:
		var update UpdateMessageIsPinned
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageInteractionInfoType:
		var update UpdateMessageInteractionInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentOpenedType:
		var update UpdateMessageContentOpened
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageMentionReadType:
		var update UpdateMessageMentionRead
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageLiveLocationViewedType:
		var update UpdateMessageLiveLocationViewed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatType:
		var update UpdateNewChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatTitleType:
		var update UpdateChatTitle
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPhotoType:
		var update UpdateChatPhoto
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPermissionsType:
		var update UpdateChatPermissions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatLastMessageType:
		var update UpdateChatLastMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPositionType:
		var update UpdateChatPosition
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsMarkedAsUnreadType:
		var update UpdateChatIsMarkedAsUnread
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsBlockedType:
		var update UpdateChatIsBlocked
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatHasScheduledMessagesType:
		var update UpdateChatHasScheduledMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatVoiceChatType:
		var update UpdateChatVoiceChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDefaultDisableNotificationType:
		var update UpdateChatDefaultDisableNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadInboxType:
		var update UpdateChatReadInbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadOutboxType:
		var update UpdateChatReadOutbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatUnreadMentionCountType:
		var update UpdateChatUnreadMentionCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatNotificationSettingsType:
		var update UpdateChatNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateScopeNotificationSettingsType:
		var update UpdateScopeNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatActionBarType:
		var update UpdateChatActionBar
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReplyMarkupType:
		var update UpdateChatReplyMarkup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDraftMessageType:
		var update UpdateChatDraftMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatFiltersType:
		var update UpdateChatFilters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatOnlineMemberCountType:
		var update UpdateChatOnlineMemberCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationType:
		var update UpdateNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNotificationGroupType:
		var update UpdateNotificationGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateActiveNotificationsType:
		var update UpdateActiveNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateHavePendingNotificationsType:
		var update UpdateHavePendingNotifications
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDeleteMessagesType:
		var update UpdateDeleteMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserChatActionType:
		var update UpdateUserChatAction
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserStatusType:
		var update UpdateUserStatus
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserType:
		var update UpdateUser
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupType:
		var update UpdateBasicGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupType:
		var update UpdateSupergroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSecretChatType:
		var update UpdateSecretChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserFullInfoType:
		var update UpdateUserFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupFullInfoType:
		var update UpdateBasicGroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupFullInfoType:
		var update UpdateSupergroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateServiceNotificationType:
		var update UpdateServiceNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileType:
		var update UpdateFile
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStartType:
		var update UpdateFileGenerationStart
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStopType:
		var update UpdateFileGenerationStop
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateCallType:
		var update UpdateCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallType:
		var update UpdateGroupCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateGroupCallParticipantType:
		var update UpdateGroupCallParticipant
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallSignalingDataType:
		var update UpdateNewCallSignalingData
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserPrivacySettingRulesType:
		var update UpdateUserPrivacySettingRules
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadMessageCountType:
		var update UpdateUnreadMessageCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadChatCountType:
		var update UpdateUnreadChatCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateOptionType:
		var update UpdateOption
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateStickerSetType:
		var update UpdateStickerSet
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateInstalledStickerSetsType:
		var update UpdateInstalledStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTrendingStickerSetsType:
		var update UpdateTrendingStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateRecentStickersType:
		var update UpdateRecentStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFavoriteStickersType:
		var update UpdateFavoriteStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSavedAnimationsType:
		var update UpdateSavedAnimations
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSelectedBackgroundType:
		var update UpdateSelectedBackground
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateLanguagePackStringsType:
		var update UpdateLanguagePackStrings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateConnectionStateType:
		var update UpdateConnectionState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTermsOfServiceType:
		var update UpdateTermsOfService
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUsersNearbyType:
		var update UpdateUsersNearby
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDiceEmojisType:
		var update UpdateDiceEmojis
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateAnimationSearchParametersType:
		var update UpdateAnimationSearchParameters
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSuggestedActionsType:
		var update UpdateSuggestedActions
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineQueryType:
		var update UpdateNewInlineQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChosenInlineResultType:
		var update UpdateNewChosenInlineResult
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallbackQueryType:
		var update UpdateNewCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineCallbackQueryType:
		var update UpdateNewInlineCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewShippingQueryType:
		var update UpdateNewShippingQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewPreCheckoutQueryType:
		var update UpdateNewPreCheckoutQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomEventType:
		var update UpdateNewCustomEvent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomQueryType:
		var update UpdateNewCustomQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollType:
		var update UpdatePoll
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdatePollAnswerType:
		var update UpdatePollAnswer
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// TestReturnError Returns the specified error and ensures that the Error object is used; for testing only. Can be called synchronously
// @param error The error to be returned
func (client *Client) TestReturnError(error *Error) (*Error, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testReturnError",
		"error": error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var errorDummy Error
	err = json.Unmarshal(result.Raw, &errorDummy)
	return &errorDummy, err

}
