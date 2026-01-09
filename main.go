package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	// TODO: Read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		return
	}

	// TODO: Set the current user to your name and update the config file
	err = cfg.SetUser("devinfinity")
	if err != nil {
		fmt.Println("error setting user:", err)
		return
	}
	// TODO: Read the config file again and print it
	cfg, err = config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		return
	}
	fmt.Println(cfg)
}
