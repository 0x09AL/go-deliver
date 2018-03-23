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
	//one_liner	string

}

type Host struct {
	name string
	htype string
	data string
}

