package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// TODO: Read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	myState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	myCommands := commands{
		allCommands: make(map[string]func(*state, command) error),
	}
	myCommands.register("login", handlerLogin)
	myCommands.register("register", handlerRegister)
	myCommands.register("reset", handlerReset)
	myCommands.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Println("no command")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = myCommands.run(&myState, cmd)
	if err != nil {
		fmt.Println("error running command", err)
		os.Exit(1)
	}

	// TODO: Set the current user to your name and update the config file
	/*err = cfg.SetUser("devinfinity")
	//if err != nil {
	//	fmt.Println("error setting user:", err)
	//	return
	}
	*/

	// TODO: Read the config file again and print it
	cfg, err = config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
