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
			log.SubSubMsg.Error().Msg("Incorrect Stdout")
		}
	}()

	bufferedLog.Title.Info().Msg("Buffered Main Message")
	log.Title.Info().Msg("Unbuffered Main Message")
	log.SubMsg.Info().Msg("Unbuffered Sub Message")
	bufferedLog.SubMsg.Info().Msg("Buffered Sub Message")
	bufferedLog.SubSubMsg.Debug().Str("aKey", "aValue").Msg("Buffered Additional Dataset")
	log.SubSubMsg.Debug().Str("aKey", "aValue").Msg("Unbuffered Additional Dataset")
	log.SubSubMsg.Warn().Msg("Unbuffered Warning")
	bufferedLog.SubSubMsg.Warn().Msg("Buffered Warning")
}
