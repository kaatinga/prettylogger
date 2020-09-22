package bufferedlogger

import (
	"bytes"
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	log := InitLog(os.Stdout)
	buffer := bytes.Buffer{}
	bufferedLog := InitLog(&buffer)

	defer func() {
		_, err := os.Stdout.Write(buffer.Bytes())
		if err != nil {
			log.subSubMsg.Error().Msg("Incorrect Stdout")
		}
	}()

	bufferedLog.title.Info().Msg("Buffered Main Message")
	log.title.Info().Msg("Unbuffered Main Message")
	log.subMsg.Info().Msg("Unbuffered Sub Message")
	bufferedLog.subMsg.Info().Msg("Buffered Sub Message")
	bufferedLog.subSubMsg.Debug().Str("aKey", "aValue").Msg("Buffered Additional Dataset")
	log.subSubMsg.Debug().Str("aKey", "aValue").Msg("Unbuffered Additional Dataset")
	log.subSubMsg.Warn().Msg("Unbuffered Warning")
	bufferedLog.subSubMsg.Warn().Msg("Buffered Warning")


}
