package command

import (
	"errors"
	"gator/internal/config"
)

type command struct {
	name string
	args []string
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
	return nil
}
