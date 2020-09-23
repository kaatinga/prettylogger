package bufferedlogger

import (
	"bytes"
	"github.com/rs/zerolog"
	"io"
)

const (
	LoggerFormat = "§ 02.01.2006 15:04"
)

type Logger struct {
	Title     *zerolog.Logger
	SubMsg    *zerolog.Logger
	SubSubMsg *zerolog.Logger
}

type BufferedLogger struct {
	Logger
	Data bytes.Buffer
}

func InitLog(writer io.Writer) (log Logger) {

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	defaultWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: LoggerFormat}
	subMessageWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: "└ §"}
	subSubMessageWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: "  └ §"}

	// Setting up logger
	title := zerolog.New(defaultWriter).With().Timestamp().Logger()
	subMsg := zerolog.New(subMessageWriter).With().Timestamp().Logger()
	subSubMsg := zerolog.New(subSubMessageWriter).With().Timestamp().Logger()

	log.Title = &title
	log.SubMsg = &subMsg
	log.SubSubMsg = &subSubMsg

	return
}

func InitBufferedLog() (log BufferedLogger) {

	log.Data = bytes.Buffer{}
	log.Logger = InitLog(&log.Data)

	//log.subSubMsg.Warn().Msg("Buffered Warning")

	return
}
