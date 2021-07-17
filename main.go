package prettylogger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	titleTimeFormat = "§ 02.01.2006 15:04"
	subMessageFormat = "└ §"
	subSubMessageFormat = "  └ §"
)

type Logger struct {
	Title     *zerolog.Logger
	SubMsg    *zerolog.Logger
	SubSubMsg *zerolog.Logger
}

func InitLogger(level zerolog.Level, colour, customLogger bool) *Logger {

	writer := os.Stdout

	globalWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: " -", NoColor: !colour}

	// Settings for the global logger
	log.Logger = zerolog.New(globalWriter).Level(level).With().Timestamp().Logger()

	log.Info().Msg("global logger format is altered")

	if customLogger {

		titleWriter, subMessageWriter, subSubMessageWriter := globalWriter, globalWriter, globalWriter

		titleWriter.TimeFormat = titleTimeFormat
		subMessageWriter.TimeFormat = subMessageFormat
		subSubMessageWriter.TimeFormat = subSubMessageFormat

		// Setting up logger
		title := zerolog.New(titleWriter).Level(level).With().Timestamp().Logger()
		subMsg := zerolog.New(subMessageWriter).Level(level).With().Timestamp().Logger()
		subSubMsg := zerolog.New(subSubMessageWriter).Level(level).With().Timestamp().Logger()

		title.Info().Msg("custom logger is enabled")
		subMsg.Debug().Bool("colour", colour).Send()
		subMsg.Debug().Str("log level", level.String()).Send()
		subMsg.Debug().Str("submessages begins with prefix", subMessageFormat).Send()
		subSubMsg.Debug().Str("subsubmessages begins with prefix", subSubMessageFormat).Send()

		return &Logger{
			Title:     &title,
			SubMsg:    &subMsg,
			SubSubMsg: &subSubMsg,
		}
	}

	return nil
}
