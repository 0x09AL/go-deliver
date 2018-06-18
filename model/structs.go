package model


type Payload struct {
	Id int
	Name string
	Content_type string
	Host_blacklist	string
	Host_whitelist	string
	Data_file	string
	Data_b64	string
	Ptype 		string
	Type_id		int
	Guid		string
	Filename	string
	//one_liner	string

}

type Host struct {
	Id int
	Name string
	Htype string
	Data string

}


type PayloadType struct {
	Type_id int
	Type_name string
	Type_template string
	Content_type string
}

type CFG struct {
	Http struct {
		Enable string
		Port int
		Template404 string
	}
	Https struct {
		Enable string
		Port int
		Template404 string
		Publickey string
		Privatekey string
	}

}