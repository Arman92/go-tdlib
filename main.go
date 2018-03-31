package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// Get API_ID and API_HASH from env vars
	apiId := os.Getenv("API_ID")
	apiId = "187786"
	if apiId == "" {
		log.Fatal("API_ID env variable not specified")
	}
	apiHash := os.Getenv("API_HASH")
	apiHash = "e782045df67ba48e441ccb105da8fc85"
	if apiHash == "" {
		log.Fatal("API_HASH env variable not specified")
	}

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

	go func() {
		time.Sleep(1 * time.Second)

		for {

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

		// Authorization block
		if update["@type"].(string) == "updateAuthorizationState" {
			if authorizationState, ok := update["authorization_state"].(tdjson.Update)["@type"].(string); ok {
				res, err := client.Auth(authorizationState, apiId, apiHash)
				if err != nil {
					log.Println(err)
				}
				log.Println(res)
			}
		} else {
			// data := map[string]interface{}{
			// 	"@type":          "getChats",
			// 	"offset_order":   0,
			// 	"offset_chat_id": 0,
			// 	"limit":          2000,
			// }
			// // result, err := client.SendAndCatch("{'@type': 'getChats', 'extra': 'bb123aa', 'offset_order': 0, 'offset_chat_id': 0, 'limit': 1000}")
			// result, err := client.SendAndCatch(data)
			// fmt.Printf("Result is : %s", result)
			// if err != nil {
			// 	fmt.Printf("errrror")
			// }

			// result, err := client.SendAndCatch("{'@type': 'getChats', 'extra': 'bb123aa', 'offset_order': 0, 'offset_chat_id': 0, 'limit': 1000}")
			// result := client.Execute(`{
			// 	"@type": "getTextEntities",
			// 	"text": "@telegram /test_command https://telegram.org telegram.me",
			// 	"@extra": ["5", 7.0]
			// }`)

		}
	}
}
