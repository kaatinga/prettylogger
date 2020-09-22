package bufferedlogger

import (
	"bytes"
	"github.com/rs/zerolog"
	"io"
)

const (
	LoggerFormat = "02.01.2006 15:04\n§"
)

type Logger struct {
	title     *zerolog.Logger
	subMsg    *zerolog.Logger
	subSubMsg *zerolog.Logger
}

type BufferedLogger struct {
	Logger
	data bytes.Buffer
}

func InitLog(writer io.Writer) (log Logger) {

	defaultWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: LoggerFormat}
	subMessageWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: "└ §"}
	subSubMessageWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: "  └ §"}

	// Setting up logger
	title := zerolog.New(defaultWriter).With().Timestamp().Logger()
	subMsg := zerolog.New(subMessageWriter).With().Timestamp().Logger()
	subSubMsg := zerolog.New(subSubMessageWriter).With().Timestamp().Logger()

	log.title = &title
	log.subMsg = &subMsg
	log.subSubMsg = &subSubMsg

	return
}

func InitBufferedLog() (log BufferedLogger) {

	log.data = bytes.Buffer{}
	log.Logger = InitLog(&log.data)

	//log.subSubMsg.Warn().Msg("Buffered Warning")

	return
}
