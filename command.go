package main

import (
	"errors"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
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
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("username set to %s\n", s.cfg.CurrentUserName)
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
