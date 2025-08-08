package helpers

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	// Set default log level to Info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "\033[90m15:04:05\033[0m", // Gray timestamp
		NoColor:    false,
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
		FormatLevel: func(i interface{}) string {
			level := "????"
			if l, ok := i.(string); ok {
				switch l {
				case "debug":
					level = "\033[36mDEBUG\033[0m" // Cyan
				case "info":
					level = "\033[32mINFO \033[0m" // Green
				case "warn":
					level = "\033[33mWARN \033[0m" // Yellow
				case "error":
					level = "\033[31mERROR\033[0m" // Red
				case "fatal":
					level = "\033[31mFATAL\033[0m" // Red
				}
			}
			return level
		},
		FormatMessage: func(i interface{}) string {
			if i == nil {
				return ""
			}
			return "  " + i.(string) // Double space after level
		},
		FormatFieldName: func(i interface{}) string {
			return " \033[2m" + i.(string) + "=\033[0m" // Dim
		},
		FormatFieldValue: func(i interface{}) string {
			return "\033[1m" + i.(string) + "\033[0m" // Bold
		},
	}

	Log = zerolog.New(output).With().Timestamp().Logger()
}
