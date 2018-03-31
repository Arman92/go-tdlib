package tdjson

//#cgo linux CFLAGS: -I/usr/local/include
//#cgo windows CFLAGS: -IC:/src/td -IC:/src/td/build
//#cgo linux LDFLAGS: -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
//#cgo windows LDFLAGS: -LC:/src/td/build/Debug -ltdjson
//#include <stdlib.h>
//#include <td/telegram/td_json_client.h>
//#include <td/telegram/td_log.h>
import "C"

import (
	"encoding/json"
	"errors"
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

// Update is used to unmarshal recieved json strings into
type Update = map[string]interface{}

// AuthorizationState is used to indicate Authorization State
type AuthorizationState uint8

// AuthorizationStates
const (
	AuthorizationStateClosed = iota
	AuthorizationStateClosing
	AuthorizationStateLoggingOut
	AuthorizationStateReady
	AuthorizationStateWaitCode
	AuthorizationStateWaitEncryptionKey
	AuthorizationStateWaitPassword
	AuthorizationStateWaitPhoneNumber
	AuthorizationStateWaitTdlibParameters
)

// Client is the Telegram TdLib client
type Client struct {
	Client     unsafe.Pointer
	Updates    chan Update
	waiters    sync.Map
	rawWaiters sync.Map
}

// TdlibConfig holds tdlibParameters
type TdlibConfig struct {
	APIID              string // Application identifier for Telegram API access, which can be obtained at https://my.telegram.org   --- must be non-empty..
	APIHash            string // Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org  --- must be non-empty..
	SystemLanguageCode string // IETF language tag of the user's operating system language; must be non-empty.
	DeviceModel        string // Model of the device the application is being run on; must be non-empty.
	SystemVersion      string // Version of the operating system the application is being run on; must be non-empty.
	ApplicationVersion string // Application version; must be non-empty.
	// Optional fields
	UseTestDataCenter      bool   // if set to true, the Telegram test environment will be used instead of the production environment.
	DatabaseDirectory      string // The path to the directory for the persistent database; if empty, the current working directory will be used.
	FileDirectory          string // The path to the directory for storing files; if empty, database_directory will be used.
	UseFileDatabase        bool   // If set to true, information about downloaded and uploaded files will be saved between application restarts.
	UseChatInfoDatabase    bool   // If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database.
	UseMessageDatabase     bool   // If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database.
	UseSecretChats         bool   // If set to true, support for secret chats will be enabled.
	EnableStorageOptimizer bool   // If set to true, old files will automatically be deleted.
	IgnoreFileNames        bool   // If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name.
}

// NewClient Creates a new instance of TDLib.
// Has two public fields:
// Client itself and Updates channel
func NewClient(config TdlibConfig) *Client {
	// Seed rand with time
	rand.Seed(time.Now().UnixNano())

	client := Client{Client: C.td_json_client_create()}
	client.Updates = make(chan Update, 100)

	go func() {
		for {
			// get update
			updateBytes := client.Receive(10)
			var update Update
			json.Unmarshal(updateBytes, &update)

			// does new update has @extra field?
			if extra, hasExtra := update["@extra"].(string); hasExtra {

				// trying to load update with this salt
				if waiter, found := client.waiters.Load(extra); found {
					// found? send it to waiter channel
					waiter.(chan Update) <- update

					// trying to prevent memory leak
					close(waiter.(chan Update))
				}

				// trying to load update with this salt
				if rawWaiter, found := client.rawWaiters.Load(extra); found {
					// found? send it to rawWaiter channel
					rawWaiter.(chan []byte) <- updateBytes

					// trying to prevent memory leak
					close(rawWaiter.(chan []byte))
				}
			} else {
				// does new updates has @type field?
				if _, hasType := update["@type"]; hasType {
					// if yes, send it in main channel
					client.Updates <- update
				}
			}
		}
	}()

	client.Send(Update{
		"@type": "setTdlibParameters",
		"parameters": Update{
			"@type":                    "tdlibParameters",
			"use_message_database":     config.UseMessageDatabase,
			"api_id":                   config.APIID,
			"api_hash":                 config.APIHash,
			"system_language_code":     config.SystemLanguageCode,
			"device_model":             config.DeviceModel,
			"system_version":           config.SystemVersion,
			"application_version":      config.ApplicationVersion,
			"enable_storage_optimizer": config.EnableStorageOptimizer,
		},
	})

	return &client
}

// Destroy Destroys the TDLib client instance.
// After this is called the client instance shouldn't be used anymore.
func (c *Client) Destroy() {
	C.td_json_client_destroy(c.Client)
}

// Send Sends request to the TDLib client.
// You can provide string or Update.
func (c *Client) Send(jsonQuery interface{}) {
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case Update:
		jsonBytes, _ := json.Marshal(jsonQuery.(Update))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(c.Client, query)
}

// Receive Receives incoming updates and request responses from the TDLib client.
// You can provide string or Update.
func (c *Client) Receive(timeout float64) []byte {
	result := C.td_json_client_receive(c.Client, C.double(timeout))

	// var update Update
	// json.Unmarshal([]byte(C.GoString(result)), &update)
	return []byte(C.GoString(result))
}

// Execute Synchronously executes TDLib request.
// Only a few requests can be executed synchronously.
func (c *Client) Execute(jsonQuery interface{}) Update {
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case Update:
		jsonBytes, _ := json.Marshal(jsonQuery.(Update))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	result := C.td_json_client_execute(c.Client, query)

	var update Update
	json.Unmarshal([]byte(C.GoString(result)), &update)
	return update
}

// SetFilePath Sets the path to the file to where the internal TDLib log will be written.
// By default TDLib writes logs to stderr or an OS specific log.
// Use this method to write the log to a file instead.
func SetFilePath(path string) {
	query := C.CString(path)
	defer C.free(unsafe.Pointer(query))

	C.td_set_log_file_path(query)
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib.
// By default the TDLib uses a verbosity level of 5 for logging.
func SetLogVerbosityLevel(level int) {
	C.td_set_log_verbosity_level(C.int(level))
}

// SendAndCatch Sends request to the TDLib client and catches the result in updates channel.
// You can provide string or Update.
func (c *Client) SendAndCatch(jsonQuery interface{}) (Update, error) {
	var update Update

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case Update:
		update = jsonQuery.(Update)
	}

	// letters for generating random string
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// generate random string for @extra field
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	randomString := string(b)

	// set @extra field
	update["@extra"] = randomString

	// create waiter chan and save it in Waiters
	waiter := make(chan Update, 1)
	c.waiters.Store(randomString, waiter)

	// send it through already implemented method
	c.Send(update)

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		return response, nil
		// or timeout
	case <-time.After(10 * time.Second):
		c.waiters.Delete(randomString)
		return Update{}, errors.New("timeout")
	}
}

// SendAndCatchBytes Sends request to the TDLib client and catches the result (bytes) in updates channel.
// You can provide string or Update.
func (c *Client) SendAndCatchBytes(jsonQuery interface{}) ([]byte, error) {
	var update Update

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case Update:
		update = jsonQuery.(Update)
	}

	// letters for generating random string
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// generate random string for @extra field
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	randomString := string(b)

	// set @extra field
	update["@extra"] = randomString

	// create waiter chan and save it in Waiters
	waiter := make(chan []byte, 1)
	c.rawWaiters.Store(randomString, waiter)

	// send it through already implemented method
	c.Send(update)

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		return response, nil
		// or timeout
	case <-time.After(10 * time.Second):
		c.rawWaiters.Delete(randomString)
		return nil, errors.New("timeout")
	}
}

// GetAuthorizationState returns authorization state
func (c *Client) GetAuthorizationState() (AuthorizationState, error) {
	res, err := c.SendAndCatch(Update{
		"@type": "getAuthorizationState",
	})

	switch res["@type"].(string) {
	case "authorizationStateClosed":
		return AuthorizationStateClosed, err
	case "authorizationStateClosing":
		return AuthorizationStateClosing, err
	case "authorizationStateLoggingOut":
		return AuthorizationStateLoggingOut, err
	case "authorizationStateReady":
		return AuthorizationStateReady, err
	case "authorizationStateWaitCode":
		return AuthorizationStateWaitCode, err
	case "authorizationStateWaitEncryptionKey":
		return AuthorizationStateWaitEncryptionKey, err
	case "authorizationStateWaitPassword":
		return AuthorizationStateWaitPassword, err
	case "authorizationStateWaitPhoneNumber":
		return AuthorizationStateWaitPhoneNumber, err
	case "authorizationStateWaitTdlibParameters":
		return AuthorizationStateWaitTdlibParameters, err
	default:
		return AuthorizationStateClosed, err
	}
}

// Authorize is used to authorize the users
func (c *Client) Authorize() (AuthorizationState, error) {
	if state, _ := c.GetAuthorizationState(); state == AuthorizationStateWaitEncryptionKey {
		_, err := c.SendAndCatch(Update{
			"@type": "checkDatabaseEncryptionKey",
		})

		if err != nil {
			return AuthorizationStateClosed, err
		}
	}

	return c.GetAuthorizationState()
}

// SendPhoneNumber sends phone number to tdlib
func (c *Client) SendPhoneNumber(phoneNumber string) (AuthorizationState, error) {
	_, err := c.SendAndCatch(Update{
		"@type":        "setAuthenticationPhoneNumber",
		"phone_number": phoneNumber,
	})

	if err != nil {
		return AuthorizationStateClosed, err
	}

	return c.GetAuthorizationState()
}

// SendAuthCode sends auth code to tdlib
func (c *Client) SendAuthCode(code string) (AuthorizationState, error) {
	_, err := c.SendAndCatch(Update{
		"@type": "checkAuthenticationCode",
		"code":  code,
	})

	if err != nil {
		return AuthorizationStateClosed, err
	}

	return c.GetAuthorizationState()
}

// SendAuthPassword sends two-step verification password (user defined)to tdlib
func (c *Client) SendAuthPassword(password string) (AuthorizationState, error) {
	_, err := c.SendAndCatch(Update{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})

	if err != nil {
		return AuthorizationStateClosed, err
	}

	return c.GetAuthorizationState()
}
