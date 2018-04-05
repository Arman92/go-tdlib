package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"tg-tdlib/tdlib"
	"time"
)

var allChats []tdlib.Chat

func main() {
	tdlib.SetLogVerbosityLevel(1)
	tdlib.SetFilePath("./errors.txt")

	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
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

	// Handle Ctrl+C , Gracefully exit and shutdown tdlib
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		client.DestroyInstance()
		os.Exit(1)
	}()

	// Wait while we get AuthorizationReady!
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	// Main loop
	go func() {
		// Just fetch updates so the Updates channel won't cause the main routine to get blocked
		for update := range client.RawUpdates {
			_ = update
		}
	}()

	// Get chats
	// see https://stackoverflow.com/questions/37782348/how-to-use-getchats-in-tdlib
	chats, err := client.GetChats(math.MaxInt64, 0, 100)
	allChats = make([]tdlib.Chat, 0, 1)
	if err != nil {
		fmt.Printf("Error getting chats, err: %v\n", err)
	} else {
		for _, chatID := range chats.ChatIDs {
			chat, err := client.GetChat(chatID)
			if err == nil {
				fmt.Println("Got chat info: ", *chat)
				allChats = append(allChats, *chat)
			}
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
