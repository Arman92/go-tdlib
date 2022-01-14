package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Arman92/go-tdlib/v2/client"
	"github.com/Arman92/go-tdlib/v2/tdlib"
)

func main() {
	client.SetLogVerbosityLevel(1)
	client.SetFilePath("./errors.txt")

	// Create new instance of client
	client := client.NewClient(client.Config{
		APIID:               "187786",
		APIHash:             "e782045df67ba48e441ccb105da8fc85",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   "./tdlib-db",
		FileDirectory:       "./tdlib-files",
		IgnoreFileNames:     false,
	})

	// You can set user-name and password to empty of don't need it
	// Socks5
	client.AddProxy("127.0.0.1", 1234, true, tdlib.NewProxyTypeSocks5("user-name", "password"))
	// HTTP - HTTPS proxy
	client.AddProxy("127.0.0.1", 1234, true, tdlib.NewProxyTypeHttp("user-name", "password", false))
	// MtProto Proxy
	client.AddProxy("127.0.0.1", 1234, true, tdlib.NewProxyTypeMtproto("MTPROTO-SECRET"))

	// Handle Ctrl+C
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		client.DestroyInstance()
		os.Exit(1)
	}()

	// Wait while we get AuthorizationReady!
	// Note: See authorization example for complete authorization sequence example
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	// Send "/start" text every 5 seconds to Forsquare bot chat
	go func() {
		// Should get chatID somehow, check out "getChats" example
		chatID := int64(198529620) // Foursquare bot chat id

		inputMsgTxt := tdlib.NewInputMessageText(tdlib.NewFormattedText("/start", nil), true, true)
		client.SendMessage(chatID, 0, 0, nil, nil, inputMsgTxt)

		time.Sleep(5 * time.Second)
	}()

	// rawUpdates gets all updates comming from tdlib
	rawUpdates := client.GetRawUpdatesChannel(100)
	for update := range rawUpdates {
		// Show all updates
		fmt.Println(update.Data)
		fmt.Print("\n\n")
	}

}
