package tdlib

//#cgo linux CFLAGS: -I/usr/local/include
//#cgo darwin CFLAGS: -I/usr/local/include
//#cgo windows CFLAGS: -IC:/src/td -IC:/src/td/build
//#cgo linux LDFLAGS: -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lc++ -lssl -lcrypto -ldl -lz -lm
//#cgo darwin LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lc++ -lssl -lcrypto -ldl -lz -lm
//#cgo windows LDFLAGS: -LC:/src/td/build/Debug -ltdjson
//#include <stdlib.h>
//#include <td/telegram/td_json_client.h>
//#include <td/telegram/td_log.h>
import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

// UpdateData alias for use in UpdateMsg
type UpdateData map[string]interface{}

// UpdateMsg is used to unmarshal recieved json strings into
type UpdateMsg struct {
	Data UpdateData
	Raw  []byte
}

// EventFilterFunc used to filter out unwanted messages in receiver channels
type EventFilterFunc func(msg *TdMessage) bool

// EventReceiver used to retreive updates from tdlib to user
type EventReceiver struct {
	Instance   TdMessage
	Chan       chan TdMessage
	FilterFunc EventFilterFunc
}

// Client is the Telegram TdLib client
type Client struct {
	Client       unsafe.Pointer
	Config       Config
	rawUpdates   chan UpdateMsg
	receivers    []EventReceiver
	waiters      map[string]chan UpdateMsg
	receiverLock *sync.Mutex
	waitersLock  *sync.RWMutex
}

// Config holds tdlibParameters
type Config struct {
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
// Client itself and RawUpdates channel
func NewClient(config Config) *Client {
	// Seed rand with time
	rand.Seed(time.Now().UnixNano())

	client := Client{Client: C.td_json_client_create()}
	client.receivers = make([]EventReceiver, 0, 1)
	client.receiverLock = &sync.Mutex{}
	client.waitersLock = &sync.RWMutex{}
	client.Config = config
	client.waiters = make(map[string]chan UpdateMsg)

	go func() {
		for {
			// get update
			updateBytes := client.Receive(10)
			var updateData UpdateData
			json.Unmarshal(updateBytes, &updateData)

			// does new update has @extra field?
			if extra, hasExtra := updateData["@extra"].(string); hasExtra {

				client.waitersLock.RLock()
				waiter, found := client.waiters[extra]
				client.waitersLock.RUnlock()

				// trying to load update with this salt
				if found {
					// found? send it to waiter channel
					waiter <- UpdateMsg{Data: updateData, Raw: updateBytes}

					// trying to prevent memory leak
					close(waiter)
				}
			} else {
				// does new updates has @type field?
				if msgType, hasType := updateData["@type"]; hasType {

					if client.rawUpdates != nil {
						// if rawUpdates is initialized, send the update in rawUpdates channel
						client.rawUpdates <- UpdateMsg{Data: updateData, Raw: updateBytes}
					}

					client.receiverLock.Lock()
					for _, receiver := range client.receivers {
						if msgType == receiver.Instance.MessageType() {
							var newMsg TdMessage
							newMsg = reflect.New(reflect.ValueOf(receiver.Instance).Elem().Type()).Interface().(TdMessage)

							err := json.Unmarshal(updateBytes, &newMsg)
							if err != nil {
								fmt.Printf("Error unmarhaling to type %v", err)
							}
							if receiver.FilterFunc(&newMsg) {
								receiver.Chan <- newMsg
							}
						}
					}
					client.receiverLock.Unlock()
				}
			}
		}
	}()

	return &client
}

// GetRawUpdatesChannel creates a general channel that fetches every update comming from tdlib
func (client *Client) GetRawUpdatesChannel(capacity int) chan UpdateMsg {
	client.rawUpdates = make(chan UpdateMsg, capacity)
	return client.rawUpdates
}

// AddEventReceiver adds a new receiver to be subscribed in receiver channels
func (client *Client) AddEventReceiver(msgInstance TdMessage, filterFunc EventFilterFunc, channelCapacity int) EventReceiver {
	receiver := EventReceiver{
		Instance:   msgInstance,
		Chan:       make(chan TdMessage, channelCapacity),
		FilterFunc: filterFunc,
	}

	client.receiverLock.Lock()
	defer client.receiverLock.Unlock()
	client.receivers = append(client.receivers, receiver)

	return receiver
}

// DestroyInstance Destroys the TDLib client instance.
// After this is called the client instance shouldn't be used anymore.
func (client *Client) DestroyInstance() {
	C.td_json_client_destroy(client.Client)
}

// Send Sends request to the TDLib client.
// You can provide string or UpdateData.
func (client *Client) Send(jsonQuery interface{}) {
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(client.Client, query)
}

// Receive Receives incoming updates and request responses from the TDLib client.
// You can provide string or UpdateData.
func (client *Client) Receive(timeout float64) []byte {
	result := C.td_json_client_receive(client.Client, C.double(timeout))

	return []byte(C.GoString(result))
}

// Execute Synchronously executes TDLib request.
// Only a few requests can be executed synchronously.
func (client *Client) Execute(jsonQuery interface{}) UpdateMsg {
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	result := C.td_json_client_execute(client.Client, query)

	var update UpdateData
	json.Unmarshal([]byte(C.GoString(result)), &update)
	return UpdateMsg{Data: update, Raw: []byte(C.GoString(result))}
}

// SetFilePath Sets the path to the file to where the internal TDLib log will be written.
// By default TDLib writes logs to stderr or an OS specific log.
// Use this method to write the log to a file instead.
func SetFilePath(path string) {
	bytes, _ := json.Marshal(UpdateData{
		"@type": "setLogStream",
		"log_stream": UpdateData{
			"@type":         "logStreamFile",
			"path":          path,
			"max_file_size": 10485760,
		},
	})

	query := C.CString(string(bytes))
	C.td_json_client_execute(nil, query)
	C.free(unsafe.Pointer(query))
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib.
// By default the TDLib uses a verbosity level of 5 for logging.
func SetLogVerbosityLevel(level int) {
	bytes, _ := json.Marshal(UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": level,
	})

	query := C.CString(string(bytes))
	C.td_json_client_execute(nil, query)
	C.free(unsafe.Pointer(query))
}

// SendAndCatch Sends request to the TDLib client and catches the result in updates channel.
// You can provide string or UpdateData.
func (client *Client) SendAndCatch(jsonQuery interface{}) (UpdateMsg, error) {
	var update UpdateData

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case UpdateData:
		update = jsonQuery.(UpdateData)
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
	waiter := make(chan UpdateMsg, 1)

	client.waitersLock.Lock()
	client.waiters[randomString] = waiter
	client.waitersLock.Unlock()

	// send it through already implemented method
	client.Send(update)

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()

		return response, nil
		// or timeout
	case <-time.After(10 * time.Second):
		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()

		return UpdateMsg{}, errors.New("timeout")
	}
}

// Authorize is used to authorize the users
func (client *Client) Authorize() (AuthorizationState, error) {
	state, err := client.GetAuthorizationState()
	if err != nil {
		return nil, err
	}

	if state.GetAuthorizationStateEnum() == AuthorizationStateWaitEncryptionKeyType {
		ok, err := client.CheckDatabaseEncryptionKey(nil)

		if ok == nil || err != nil {
			return nil, err
		}
	} else if state.GetAuthorizationStateEnum() == AuthorizationStateWaitTdlibParametersType {
		client.sendTdLibParams()
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

func (client *Client) sendTdLibParams() {
	client.Send(UpdateData{
		"@type": "setTdlibParameters",
		"parameters": UpdateData{
			"@type":                    "tdlibParameters",
			"use_test_dc":              client.Config.UseTestDataCenter,
			"database_directory":       client.Config.DatabaseDirectory,
			"files_directory":          client.Config.FileDirectory,
			"use_file_database":        client.Config.UseFileDatabase,
			"use_chat_info_database":   client.Config.UseChatInfoDatabase,
			"use_message_database":     client.Config.UseMessageDatabase,
			"use_secret_chats":         client.Config.UseSecretChats,
			"api_id":                   client.Config.APIID,
			"api_hash":                 client.Config.APIHash,
			"system_language_code":     client.Config.SystemLanguageCode,
			"device_model":             client.Config.DeviceModel,
			"system_version":           client.Config.SystemVersion,
			"application_version":      client.Config.ApplicationVersion,
			"enable_storage_optimizer": client.Config.EnableStorageOptimizer,
			"ignore_file_names":        client.Config.IgnoreFileNames,
		},
	})
}

// SendPhoneNumber sends phone number to tdlib
func (client *Client) SendPhoneNumber(phoneNumber string) (AuthorizationState, error) {
	phoneNumberConfig := PhoneNumberAuthenticationSettings{AllowFlashCall: false, IsCurrentPhoneNumber: false, AllowSmsRetrieverAPI: false}
	_, err := client.SetAuthenticationPhoneNumber(phoneNumber, &phoneNumberConfig)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthCode sends auth code to tdlib
func (client *Client) SendAuthCode(code string) (AuthorizationState, error) {
	_, err := client.CheckAuthenticationCode(code)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthPassword sends two-step verification password (user defined)to tdlib
func (client *Client) SendAuthPassword(password string) (AuthorizationState, error) {
	_, err := client.CheckAuthenticationPassword(password)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}
