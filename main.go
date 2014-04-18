package main

import (
	"flag"
	"fmt"

	ipupdate "iptodate/lib"

	"github.com/BurntSushi/toml"
)

var (
	config_flag        = flag.String("c", "/etc/ipupdate.conf", "Config file for CF/Domain info")
	find_id_flag       = flag.Bool("l", false, "Find the ID that corresponds to the record")
	add_record_flag    = flag.Bool("a", false, "Add a record with the configured name")
	update_record_flag = flag.Bool("u", false, "Update the record with the configured data and address")
	dns_zone           = flag.String("h", "", "Host (ie: example.com) to update")
	subdomain          = flag.String("s", "", "Subdomain (ie: home.example.com would be 'home' to update")
)

func main() {
	flag.Parse()
	var record_id, zone, sub string
	var ass bool
	var config Config
	_, err := toml.DecodeFile(*config_flag, &config)
	if err != nil {
		panic(err)
	}
	zone, ass = AssignVar(*dns_zone, config.Domain)
	if !ass {
		fmt.Println("I need a DNS zone in the config file or as a flag!")
	}
	sub, ass = AssignVar(*subdomain, config.Name)
	if !ass {
		fmt.Println("I need a subdomain, either in the config file or as a flag!")
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
		fmt.Println(err)
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

func AssignVar(first string, second string) (string, bool) {
	if first == "" {
		if second == "" {
			return "", false
		}
		return first, true
	}
	if second == "" {
		return "", false
	}
	return second, true
}

type Config struct {
	APIKey   string
	APIEmail string
	ID       string
	Domain   string
	Name     string
	IPAPI    string
}
