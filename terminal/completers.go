package terminal

import (
	"github.com/chzyer/readline"
	"go-deliver/database"
)

var MainCompleter = readline.NewPrefixCompleter(
	readline.PcItem("payload",
		readline.PcItem("add",
			readline.PcItemDynamic(database.GetPayloadTypeCompleter(),
				),
		),
		readline.PcItem("delete",
			readline.PcItemDynamic(database.GetPayloadDeleteCompleter()),
		),
		readline.PcItem("list"),

	),
	readline.PcItem("host",
		readline.PcItem("add"),
		readline.PcItem("list"),
		readline.PcItem("delete",
			readline.PcItemDynamic(database.GetHostNameCompleter()),),
	),
	//	readline.PcItem("listeners") To be implemented later .
)

var PayloadCompleter = readline.NewPrefixCompleter(
	readline.PcItem("set",
		readline.PcItem("name"),
		readline.PcItem("content_type"),
		readline.PcItem("host_blacklist",
			readline.PcItemDynamic(database.GetHostNameCompleter()),
				),
		readline.PcItem("host_whitelist",
			readline.PcItemDynamic(database.GetHostNameCompleter()),
				),
		readline.PcItem("data_file"),
		readline.PcItem("data_b64"),
		readline.PcItem("ptype"),
		readline.PcItem("filename"),
		//readline.PcItem("listener"), // This is will be implemented later.
	),
	readline.PcItem("unset",
		readline.PcItem("content_type"),
		readline.PcItem("host_blacklist"),
		readline.PcItem("host_whitelist"),
		readline.PcItem("data_file"),
		readline.PcItem("data_b64"),
		readline.PcItem("type"),
		//readline.PcItem("listener"), // This is will be implemented later.
	),
	readline.PcItem("options"),
	readline.PcItem("create"),
	readline.PcItem("back"),

)


var HostCompleter = readline.NewPrefixCompleter(
	readline.PcItem("set",
		readline.PcItem("name"),
		readline.PcItem("htype",
			readline.PcItem("ip"),
			readline.PcItem("subnet"),
		),
		readline.PcItem("restriction_type",
			readline.PcItem("whitelist"),
			readline.PcItem("blacklist"),
		),
		readline.PcItem("data"),
	),
	readline.PcItem("options"),
	readline.PcItem("create"),
	readline.PcItem("back"),
)