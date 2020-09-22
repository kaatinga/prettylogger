package bufferedlogger

import (
	"bytes"
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	unbufferedLog := InitLog(os.Stdout)
	buffer := bytes.Buffer{}
	bufferedLog := InitLog(&buffer)

	defer func() {
		_, err := os.Stdout.Write(buffer.Bytes())
		if err != nil {
			unbufferedLog.SubSubMsg.Error().Msg("Incorrect Stdout")
		}
	}()

	bufferedLog.Title.Info().Msg("Buffered Main Message")
	unbufferedLog.Title.Info().Msg("Unbuffered Main Message")
	unbufferedLog.SubMsg.Info().Msg("Unbuffered Sub Message")
	bufferedLog.SubMsg.Info().Msg("Buffered Sub Message")
	bufferedLog.SubSubMsg.Debug().Str("aKey", "aValue").Msg("Buffered Additional Dataset")
	unbufferedLog.SubSubMsg.Debug().Str("aKey", "aValue").Msg("Unbuffered Additional Dataset")
	unbufferedLog.SubSubMsg.Warn().Msg("Unbuffered Warning")
	bufferedLog.SubSubMsg.Warn().Msg("Buffered Warning")
}
