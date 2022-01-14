package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Arman92/go-tdlib/v2/client"
	"github.com/Arman92/go-tdlib/v2/tdlib"
)

var allChats []*tdlib.Chat
var haveFullChatList bool

// TdInstance ...
type TdInstance struct {
	AcountName          string         `json:"AcountName"`
	TdlibDbDirectory    string         `json:"TdlibDbDirectory"`
	TdlibFilesDirectory string         `json:"TdlibFilesDirectory"`
	TdlibClient         *client.Client `json:"-"`
}

var botInstances []TdInstance

func main() {

	var reconfig bool
	flag.BoolVar(&reconfig, "reconfig", false, "Pass true if you want to reconfigure the accounts")
	flag.Parse()

	client.SetLogVerbosityLevel(1)

	// if reconfig flag is set true, we will re-configure the telegram accounts, otherwise we just read the previous config and add the instances
	if reconfig {
		reader := bufio.NewReader(os.Stdin)

		for {
			var tdInstance TdInstance
			fmt.Println("Enter a name for this telegram account so you'll remeber it")
			text, _ := reader.ReadString('\n')
			tdInstance.AcountName = text[:len(text)-1]

			fmt.Println("Enter account's tdlib client Db directory (empty for default: ./tddata/{accName}-tdlib-db)")
			text, _ = reader.ReadString('\n')
			if text == "\n" {
				tdInstance.TdlibDbDirectory = "./tddata/" + tdInstance.AcountName + "-tdlib-db"
			} else {
				tdInstance.TdlibFilesDirectory = text[:len(text)-1]
			}

			fmt.Println("Enter accounts's tdlib client Files directory (empty for default: ./tddata/{accName}-tdlib-files)")
			text, _ = reader.ReadString('\n')
			if text == "\n" {
				tdInstance.TdlibDbDirectory = "./tddata/" + tdInstance.AcountName + "-tdlib-files"
			} else {
				tdInstance.TdlibFilesDirectory = text[:len(text)-1]
			}

			botInstances = append(botInstances, tdInstance)
			fmt.Println("Adding the TdInstance tdlib client, may took a while...")
			tdInstance.addTdlibClient()

			fmt.Println("If you are want add more telegram accounts, enter y(es), otherwise if you want to save the configs and go after your dreams, just hit enter")
			text, _ = reader.ReadString('\n')
			if strings.HasPrefix(strings.ToLower(text), "y") {
				continue
			} else {
				json, _ := json.Marshal(botInstances)
				f, err := os.Create("configs.json")
				if err != nil {
					log.Fatalf("Failed to open file configs.json: %v", err)
				}

				_, err = io.Copy(f, bytes.NewReader(json))
				f.Close()

				break
			}
		}
	} else {
		f, err := os.Open("configs.json")
		if err != nil {
			log.Fatalf("Failed to open file configs.json: %v", err)
		}
		err = json.NewDecoder(f).Decode(&botInstances)
		f.Close()
		if err != nil {
			log.Fatalf("Failed to unmarshal configs: %v", err)
		}

		for i := range botInstances {
			botInstances[i].addTdlibClient()
		}
	}

}

func (tdInstance *TdInstance) addTdlibClient() {
	// Create new instance of client
	client := client.NewClient(client.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tdInstance.TdlibDbDirectory,
		FileDirectory:       tdInstance.TdlibFilesDirectory,
		IgnoreFileNames:     false,
	})
	tdInstance.TdlibClient = client

	// Main loop
	go func() {
		// Just fetch updates so the Updates channel won't cause the main routine to get blocked
		rawUpdates := client.GetRawUpdatesChannel(100)
		for update := range rawUpdates {
			// Show all updates
			_ = update
			// fmt.Println(update.Data)
			// fmt.Print("\n\n")
		}
	}()

	for {
		currentState, err := client.Authorize()
		if err != nil {
			fmt.Printf("Error getting current state: %v", err)
			continue
		}
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			fmt.Print("Enter phone: ")
			var number string
			fmt.Scanln(&number)
			_, err := client.SendPhoneNumber(number)
			if err != nil {
				fmt.Printf("Error sending phone number: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			fmt.Print("Enter code: ")
			var code string
			fmt.Scanln(&code)
			_, err := client.SendAuthCode(code)
			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
			fmt.Print("Enter Password: ")
			var password string
			fmt.Scanln(&password)
			_, err := client.SendAuthPassword(password)
			if err != nil {
				fmt.Printf("Error sending auth password: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready! Let's rock")
			break
		}
	}

}
