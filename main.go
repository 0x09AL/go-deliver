package main
import (
	"go-deliver/terminal"
	"go-deliver/model"
	"go-deliver/servers"
	"log"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"gopkg.in/gcfg.v1"
	"fmt"
	b64 "encoding/base64"
)


func main() {

	var b64_banner = "DQogIF9fX19fX19fICAgICAgICAgICAgICAgICBfX19fX19fXyAgICAgICAgIC5fXyAgLl9fICAgICAgICAgICAgICAgICAgICANCiAvICBfX19fXy8gIF9fX18gICAgICAgICAgIFxfX19fX18gXCAgIF9fX18gfCAgfCB8X198X18gIF9fIF9fX19fX19fX19fIA0KLyAgIFwgIF9fXyAvICBfIFwgICBfX19fX18gIHwgICAgfCAgXF8vIF9fIFx8ICB8IHwgIFwgIFwvIC8vIF9fIFxfICBfXyBcDQpcICAgIFxfXCAgKCAgPF8+ICkgL19fX19fLyAgfCAgICBgICAgXCAgX19fL3wgIHxffCAgfFwgICAvXCAgX19fL3wgIHwgXC8NCiBcX19fX19fICAvXF9fX18vICAgICAgICAgIC9fX19fX19fICAvXF9fXyAgPl9fX18vX198IFxfLyAgXF9fXyAgPl9ffCAgIA0KICAgICAgICBcLyAgICAgICAgICAgICAgICAgICAgICAgICBcLyAgICAgXC8gICAgICAgICAgICAgICAgICAgXC8gICAgICAgDQo="
	banner, _ := b64.StdEncoding.DecodeString(b64_banner)
	fmt.Println(string(banner))

	Configuration := model.CFG{}

	err := gcfg.ReadFileInto(&Configuration,"config.conf")

	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	if(Configuration.Http.Enable == "true"){
		log.Println(fmt.Sprintf("Starting http server on port %d",Configuration.Http.Port))
		go servers.StartHTTPListener(Configuration)
	}
	if(Configuration.Https.Enable == "true"){
		log.Println(fmt.Sprintf("Starting https server on port %d",Configuration.Https.Port))
		go servers.StartHTTPSListener(Configuration)
	}


	// Start the terminal
	terminal.StartTerminal()


}