package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// read meta.json
	metaFile, err := os.Open("meta.json")
	if err != nil {
		fmt.Println("Error opening meta.json")
	}
	defer metaFile.Close()
	metafileContent, _ := io.ReadAll(metaFile)
	// parse json
	var config_file map[string]interface{}
	json.Unmarshal(metafileContent, &config_file)
	var trackers_index Trackers
	json.Unmarshal(metafileContent, &trackers_index)
	// fmt.Println(trackers_index.Trackers[0].Name)
	// print basic debug vars stuff
	debug_println(trackers_index, config_file, metafileContent)
}
