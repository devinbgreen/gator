package command

import (
	"errors"
	"fmt"
	"gator/internal/config"
)

type command struct {
	name string
	args []string
}

type commands struct {
	allCommands map[string]func(*state, command) error
}

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username")
	}
	//cmd.args[0] to update s.cfg
	s.cfg.CurrentUserName = cmd.args[0]
	fmt.Printf("username set to %s\n", s.cfg.CurrentUserName)
	return nil
}
