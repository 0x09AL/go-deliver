package database

import (
	"go-deliver/model"
	"fmt"
	"log"
	"github.com/olekukonko/tablewriter"
	"os"

)

func CreateHost(host model.Host)  {
	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare(model.CreateHostQuery)
	fmt.Println("")
	if err_stmt != nil {
		log.Fatal(err_stmt)
	}
	_, err := stmt.Exec(host.Name,host.Htype,host.Data)
	tx.Commit()
	if err != nil{
		log.Println("ERROR: Error inserting host.")
	}else{

		log.Println(fmt.Sprintf("Host with name %s created successfully.",host.Name)) // Fix the URL output
	}
}

func ListHosts()  {

	rows, err := db.Query(model.GetHostsQuery)

	host := model.Host{}
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Host Type", "Data"})

	for rows.Next() {
		err := rows.Scan(&host.Id, &host.Name, &host.Htype , &host.Data)
		table.Append([]string{ host.Name,host.Htype,host.Data})
		if err != nil {
			log.Fatal(err)
		}
	}
	table.Render()

}

func DeleteHost(name string)  {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(model.DeleteHostQuery)
	_, err := stmt.Exec(name)
	if err != nil{
		log.Panic(err)
	}
	tx.Commit()
	log.Println(fmt.Sprintf("Success : Host %s deleted .",name))
}

func GetData(name string) (string, string){

	host := model.Host{}
	err := db.QueryRow(model.GetHostDataQuery, name).Scan(&host.Htype,&host.Data)
	if err != nil {
		panic(err)
	}
	return host.Htype, host.Data
}


func GetHostNameCompleter()  func(string) []string{

	return func(line string) []string {
		var HostNames []string
		var temp string
		rows, err := db.Query(model.GetHostNamesCompleterQuery)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			err := rows.Scan(&temp)
			HostNames = append(HostNames,temp)
			if err != nil {
				log.Fatal(err)
			}
		}
		return HostNames
	}
}