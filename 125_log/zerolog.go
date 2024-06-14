package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AddFieldHook struct {
}

func (AddFieldHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("name", "dj")
}

func main() {
	hooked := log.Hook(AddFieldHook{})
	hooked.Debug().Msg("hooked debug")
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Debug().Msg("hooked debug")
}
