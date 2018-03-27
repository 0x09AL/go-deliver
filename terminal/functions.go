package terminal


import(
	"fmt"
	"github.com/chzyer/readline"
	"strings"
	"io"
	"go-deliver/model"
	"go-deliver/database"
	"encoding/json"
	"log"

)

var context string = "main"
var prompt string = "go-deliver (\033[0;32m%s\033[0;0m)\033[31m >> \033[0;0m"

func handleHostCreation(name string, l *readline.Instance){
	host := model.Host{}
	host.Name = name
	l.Config.AutoComplete = HostCompleter
	l.SetPrompt(fmt.Sprintf(prompt,"host-options"))
	for {line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		temp := strings.Split(line," ")
		command := temp[0]
		switch command{
		case "back":
			backMain(l)
			return
		case "options":
			// To be fixed
			data, err := json.MarshalIndent(host,"", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)

		case "set":
			if len(temp) == 3{
				key := temp[1]
				value := temp[2]
				switch key{
				case "name":
					host.Name = value
				case "htype":
					host.Htype = value
				case "data":
					host.Data = value
				}

			}
		case "unset":
			if len(temp) == 2{
				key := temp[1]
				switch key{
				case "name":
					host.Name = ""
				case "htype":
					host.Htype = ""
				case "data":
					host.Data = ""
				}
			}
		case "create":
			database.CreateHost(host)
		default:


		}
	}
}

func handlePayloadCreation(ptype string, l *readline.Instance)  {

	payload := model.Payload{}
	payload.Ptype = ptype
	payload.Type_id,payload.Content_type = database.GetTypeid(ptype)


	l.Config.AutoComplete = PayloadCompleter
	l.SetPrompt(fmt.Sprintf(prompt,"payload-options"))

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		temp := strings.Split(line," ")
		command := temp[0]
		switch command{
		case "back":
			backMain(l)
			return
		case "options":
			// To be fixed
			data, err := json.MarshalIndent(payload,"", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)

		case "set":
			if len(temp) == 3{
				key := temp[1]
				value := temp[2]
				switch key{
				case "name":
					payload.Name = value
				case "content_type":
					payload.Content_type = value
				case "host_whitelist":
					payload.Host_whitelist = value
				case "host_blacklist":
					payload.Host_blacklist = value
				case "data_file":
					payload.Data_file = value
				case "data_b64":
					payload.Data_b64 = value
				case "ptype":
					payload.Ptype = value
				}

			}
		case "unset":
			if len(temp) == 2{
				key := temp[1]
				switch key{
				case "name":
					payload.Name = ""
				case "content_type":
					payload.Content_type = ""
				case "host_whitelist":
					payload.Host_whitelist = ""
				case "host_blacklist":
					payload.Host_blacklist = ""
				case "data_file":
					payload.Data_file = ""
				case "data_b64":
					payload.Data_b64 = ""
				case "ptype":
					payload.Ptype = ""
				}
			}
		case "create":
			database.InsertPayload(payload)
		default:


		}
	}
}



func backMain(l *readline.Instance){
	context = "main"
	l.SetPrompt(fmt.Sprintf(prompt,"main"))
	l.Config.AutoComplete = MainCompleter
}

func handleInput(line string ,l *readline.Instance)  {

	var ptype string
	var command string
	line = strings.TrimSpace(line)
	temp := strings.Split(line," ")
	if len(temp) >1 {
		command = temp[1]
	}

	switch {

	// Handle the payload functions
	case strings.HasPrefix(line, "payload "):

		switch  command{
		case "add":
			if len(temp) > 2{
				ptype = temp[2]
				handlePayloadCreation(ptype,l)
			}
		case "delete":
			if len(temp) > 2{
				name := temp[2]
				database.DeletePayload(name)
			}
		case "list":
			log.Println("Listing payloads")
			database.GetPayloads()
		default:
			fmt.Println("Invalid command")
		}

		// Handle the Hosts functions
	case strings.HasPrefix(line, "host "):
		switch command {
		case "add":
			if len(temp) > 2{
				name := temp[2]
				handleHostCreation(name,l)
			}
		case "list":
			log.Println("Listing hosts")
			database.ListHosts()
		case "delete":
			if len(temp) > 2{
				name := temp[2]
				database.DeleteHost(name)
			}
		default:
			fmt.Println("Invalid command")
		}

	}



}

func StartTerminal()  {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          fmt.Sprintf(prompt,"main"),
		HistoryFile:     "history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:	 MainCompleter,

	})


	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {

		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		handleInput(line,l)


	}
}


