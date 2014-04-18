package main

import (
	"flag"
	"fmt"

	ipupdate "github.com/nickvanw/iptodate/lib"

	"github.com/BurntSushi/toml"
)

var (
	config_flag        = flag.String("c", "/etc/ipupdate.conf", "Config file for CF/Domain info")
	find_id_flag       = flag.Bool("l", false, "Find the ID that corresponds to the record")
	add_record_flag    = flag.Bool("a", false, "Add a record with the configured name")
	update_record_flag = flag.Bool("u", false, "Update the record with the configured data and address")
)

func main() {
	flag.Parse()
	var record_id string
	var config Config
	_, err := toml.DecodeFile(*config_flag, &config)
	if err != nil {
		panic(err)
	}
	if !*add_record_flag && !*find_id_flag && !*update_record_flag {
		fmt.Println("I didn't get any arguments! Run with the -h command to see what they are")
	}
	if *add_record_flag {
		_, err := ipupdate.GetId(config.Domain, config.APIEmail, config.APIKey, config.Name)
		if err == nil {
			fmt.Println("There's already a record with that name!, run this program without the -a flag to change it")
			return
		}
		ip, err := ipupdate.FetchIP(config.IPAPI)
		if err != nil {
			fmt.Println("Error fetching IP:", err)
			return
		}
		id, err := ipupdate.CreateRecord(ip, config.Domain, config.Name, config.APIEmail, config.APIKey)
		if err != nil {
			fmt.Println("Error creating record: ", err)
			return
		}
		fmt.Println("Created record with ID:", id)
		record_id = id
	}
	if *find_id_flag {
		id, err := ipupdate.GetId(config.Domain, config.APIEmail, config.APIKey, config.Name)
		if err != nil {
			fmt.Println("Error finding ID: ", err)
			return
		}
		record_id = id
		fmt.Println("Found a record with the id", record_id)
	}
	if record_id == "" {
		record_id = config.ID
	}
	if *update_record_flag {
		ip, err := ipupdate.FetchIP(config.IPAPI)
		if err != nil {
			fmt.Println("Error fetching IP:", err)
			return
		}
		upd, err := ipupdate.SetAddress(ip, config.Domain, config.Name, config.APIEmail, record_id, config.APIKey)
		if err != nil || upd == false {
			fmt.Println("Error updating Address", err)
			return
		}
		fmt.Println("IP Updated successfully to", ip)
	}
}

type Config struct {
	APIKey   string
	APIEmail string
	ID       string
	Domain   string
	Name     string
	IPAPI    string
}
