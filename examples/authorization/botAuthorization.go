package main

import (
	"fmt"

	"github.com/Arman92/go-tdlib/v2/client"
	"github.com/Arman92/go-tdlib/v2/tdlib"
)

const botToken = "<your bot token>"

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

	for {
		currentState, _ := client.Authorize()
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			_, err := client.CheckAuthenticationBotToken(botToken)
			if err != nil {
				fmt.Printf("Error check bot token: %v", err)
				return
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready! Let's rock")
			break
		}
	}

	// rawUpdates gets all updates comming from tdlib
	rawUpdates := client.GetRawUpdatesChannel(100)
	for update := range rawUpdates {
		// Show all updates
		fmt.Println(update.Data)
		fmt.Print("\n\n")
	}
}
