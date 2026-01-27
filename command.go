package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

type command struct {
	name string
	args []string
}

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type commands struct {
	allCommands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username")
	}
	//cmd.args[0] to update s.cfg
	u, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == sql.ErrNoRows {
		fmt.Println("user doesn't exist")
		return errors.New("user doesn't exist")
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("username set to %s\n", s.cfg.CurrentUserName)
	fmt.Println("logged in user:")
	printUser(u)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		current := ""
		if s.cfg.CurrentUserName == u.Name {
			current = " (current)"
		}
		fmt.Printf("* %s%s\n", u.Name, current)
	}

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username")
	}
	u, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != sql.ErrNoRows {
		fmt.Println("user exists")
		return errors.New("user exists")
	}

	u, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("logged in user:")
	printUser(u)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.allCommands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.allCommands[name] = f
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
