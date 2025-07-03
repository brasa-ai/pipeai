package helpers

import (
	"os"

	"github.com/rs/zerolog"
)

var Log = zerolog.New(os.Stderr).With().Timestamp().Logger()
