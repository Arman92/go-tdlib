package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tg-tdlib/tdjson"
	"tg-tdlib/types"
	"time"
)

func main() {
	tdjson.SetLogVerbosityLevel(1)
	tdjson.SetFilePath("./errors.txt")

	// Create new instance of client
	client := tdjson.NewClient(tdjson.TdlibConfig{
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
		DatabaseDirectory:   "tdlib-db",
		FileDirectory:       "tdlib-files",
		IgnoreFileNames:     false,
	})

	// Handle Ctrl+C
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		client.Destroy()
		os.Exit(1)
	}()

	for {
		currentState, _ := client.Authorize()
		if currentState == tdjson.AuthorizationStateWaitPhoneNumber {
			fmt.Print("Enter phone: ")
			var number string
			fmt.Scanln(&number)
			client.SendPhoneNumber(number)
		} else if currentState == tdjson.AuthorizationStateWaitCode {
			fmt.Print("Enter code: ")
			var code string
			fmt.Scanln(&code)
			client.SendAuthCode(code)
		} else if currentState == tdjson.AuthorizationStateWaitPassword {
			fmt.Print("Enter Password: ")
			var password string
			fmt.Scanln(&password)
			client.SendAuthPassword(password)
		} else if currentState == tdjson.AuthorizationStateReady {
			break
		}
	}

	go func() {
		time.Sleep(1 * time.Second)

		for {
			res, err := client.SendAndCatch(tdjson.Update{
				"@type": "getAuthorizationState",
			})
			fmt.Println(res)
			fmt.Println(err)

			result, err := client.SendAndCatchBytes(tdjson.Update{
				"@type":          "getChats",
				"offset_order":   9223372036854775807,
				"offset_chat_id": 0,
				"limit":          200,
			})

			if err != nil {
				panic("Fuck")
			}

			chatIDs := types.Chats{}
			json.Unmarshal(result, &chatIDs)

			var chats []types.Chat
			chats = make([]types.Chat, 0, 1)

			for _, chatID := range chatIDs.ChatIDs {
				result, err = client.SendAndCatchBytes(tdjson.Update{
					"@type":   "getChat",
					"chat_id": chatID,
				})

				var update tdjson.Update
				json.Unmarshal(result, &update)
				fmt.Println(update)

				var chat types.Chat
				json.Unmarshal(result, &chat)
				chats = append(chats, chat)
			}

			result, err = client.SendAndCatchBytes(tdjson.Update{
				"@type":   "getChatMember",
				"chat_id": -1001178287687,
				"user_id": 41507975,
			})

			var update tdjson.Update
			json.Unmarshal(result, &update)
			fmt.Println(update)

			time.Sleep(3 * time.Second)
		}
	}()

	// Main loop
	for update := range client.Updates {
		// Show all updates
		fmt.Println(update)
		fmt.Print("\n\n")

	}
}
