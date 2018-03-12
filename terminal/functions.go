package terminal


import(
	"fmt"
	"github.com/chzyer/readline"
	"strings"
	"io"
)

var context string = "main"
var prompt string = "go-deliver (\033[0;32m%s\033[0;0m)\033[31m >> \033[0;0m"

type Payload struct {
	id int
	name string
	content_type string
	host_blacklist	string
	host_whitelist	string
	data_file	string
	data_b64	string
	ptype 		string
	one_liner	string

}

var MainCompleter = readline.NewPrefixCompleter(
	readline.PcItem("payload",
		readline.PcItem("add",
			readline.PcItem("mshta"),
			readline.PcItem("regsrv32"),
			readline.PcItem("powershell"),
			readline.PcItem("javascript"),
		),
		readline.PcItem("delete"),
		readline.PcItem("list"),
	),
	readline.PcItem("host"),
	//	readline.PcItem("listeners") To be implemented later .
)

var PayloadCompleter = readline.NewPrefixCompleter(
	readline.PcItem("set",
		readline.PcItem("content_type"),
		readline.PcItem("host_blacklist"),
		readline.PcItem("host_whitelist"),
		readline.PcItem("data_file"),
		readline.PcItem("data_b64"),
		readline.PcItem("type"),
		//readline.PcItem("listener"), // This is to be implemented later.
		),
	readline.PcItem("unset",
		readline.PcItem("content_type"),
		readline.PcItem("host_blacklist"),
		readline.PcItem("host_whitelist"),
		readline.PcItem("data_file"),
		readline.PcItem("data_b64"),
		readline.PcItem("type"),
		//readline.PcItem("listener"), // This is to be implemented later.
	),
	readline.PcItem("options"),
)

func handlePayloadCreation(ptype string, l *readline.Instance)  {
	payload := Payload{}
	payload.ptype = ptype
	fmt.Println(fmt.Sprintf("Will create a payload named with the ptype %s",payload.ptype))
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
		if line == "exit"{
			backMain(l)
			break
		}else{
			fmt.Println(line)
		}
	}
}



func backMain(l *readline.Instance){
	context = "main"
	l.SetPrompt(fmt.Sprintf(prompt,"main"))
}

func handleInput(line string ,l *readline.Instance)  {


	line = strings.TrimSpace(line)
	temp := strings.Split(line," ")

	if len(temp) > 2 {

		command := temp[1]
		switch {

		// Handle the payload functions
		case strings.HasPrefix(line, "payload "):

			var ptype string = temp[2]
			switch  command{
			case "add":
				handlePayloadCreation(ptype,l)
			case "delete":
				fmt.Println("Remove a payload")
			default:
				fmt.Println("Invalid command")
			}

		// Handle the Hosts functions
		case strings.HasPrefix(line, "host "):
			switch command {
			case "add":
				fmt.Println("Add a host")
			case "delete":
				fmt.Println("Remove a host")
			default:
				fmt.Println("Invalid command")
			}

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



