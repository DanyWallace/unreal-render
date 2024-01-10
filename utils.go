package main

import (
	"fmt"
	"os"
)

type Configvars struct {
	Master_url string `json:"master_url"`
	User       string `json:"user"`
}

type Trackers struct {
	Trackers []Trackernode `json:"nodes"`
}

type Trackernode struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	SshPort int    `json:"ssh_port"`
}

func test_example() {
	fmt.Println("example func")
}

func debug_println(trackers_index Trackers, config_file map[string]interface{}, metafileContent []byte) {
	if len(os.Args) > 1 && os.Args[1] == "vars" {
		// print out all trackers
		fmt.Println("Hosts:")
		var master_tracker Trackernode
		for _, tracker := range trackers_index.Trackers {
			fmt.Println(tracker.Name, "-", tracker.Url, "- ssh-port:", tracker.SshPort)
			if tracker.Name == "master" { // isolate master tracker
				master_tracker = tracker
			}
		}

		fmt.Println("Master url:", config_file["master_url"])
		fmt.Println("metafileContent:", string(metafileContent))
		fmt.Println("master:", master_tracker.Url, "- ssh-port:", master_tracker.SshPort)
	}
}
