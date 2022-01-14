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

var allChats []*tdlib.Chat
var haveFullChatList bool

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

	// Handle Ctrl+C , Gracefully exit and shutdown tdlib
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		client.DestroyInstance()
		os.Exit(1)
	}()

	// Wait while we get AuthorizationReady!
	// Note: See authorization example for complete auhtorization sequence example
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	// get at most 1000 chats list
	getChatList(client, 1000)
	fmt.Printf("got %d chats\n", len(allChats))

	for _, chat := range allChats {
		fmt.Printf("Chat title: %s \n", chat.Title)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func getChatList(client *client.Client, limit int) error {
	var chatList = tdlib.NewChatListMain()

	chats, err := client.GetChats(chatList, int32(limit))
	if err != nil {
		return err
	}

	for len(chats.ChatIDs) != limit {
		// get chats (ids) from tdlib
		_, err := client.LoadChats(chatList, int32(limit-len(chats.ChatIDs)))
		if err != nil {
			if err.(tdlib.RequestError).Code != 404 {
				chats, err = client.GetChats(chatList, int32(limit))
				break
			}
			return err
		}

		chats, err = client.GetChats(chatList, int32(limit))
	}

	for _, chatID := range chats.ChatIDs {
		// get chat info from tdlib
		chat, err := client.GetChat(chatID)
		if err == nil {
			allChats = append(allChats, chat)
		} else {
			return err
		}
	}

	return nil
}
