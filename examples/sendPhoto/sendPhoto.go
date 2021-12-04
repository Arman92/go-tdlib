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
	// Should get chatID somehow, check out "getChats" example
	chatID := int64(1055350095) // Foursquare bot chat id

	// Note: file path ("./bunny.jpg") should be relative to your execution directory.
	inputMsg := tdlib.NewInputMessagePhoto(tdlib.NewInputFileLocal("./bunny.jpg"), nil, nil, 400, 400,
		tdlib.NewFormattedText("A photo sent from go-tdlib!", nil), 0)
	sentMsg, err := client.SendMessage(chatID, 0, 0, nil, nil, inputMsg)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Wait until the photo gets uploaded/sent.
	for {
		statMsg, err := client.GetMessage(sentMsg.ChatID, sentMsg.ID)
		if err == nil {
			fmt.Println(statMsg.SendingState)
		} else {
			fmt.Println(err)
			return
		}

		time.Sleep(2 * time.Second)
	}
}
