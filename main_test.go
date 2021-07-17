package prettylogger

import (
	"github.com/rs/zerolog"
	"testing"
)

func TestInitLogger(t *testing.T) {

	tests := []struct {
		name string
		level zerolog.Level
		colour bool
		customLogger bool
	}{
		{ "ok 1", zerolog.TraceLevel, true, true},
		{ "ok 2", zerolog.InfoLevel, true, true},
		{ "ok 3", zerolog.InfoLevel, false, true},
		{ "ok 4", zerolog.InfoLevel, false, false},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gottenLogger := InitLogger(tt.level, tt.colour, tt.customLogger)
			if (gottenLogger != nil) != tt.customLogger {
				t.Errorf("InitLogger() = %v, want %v", gottenLogger, "a logger")
			}
		})
	}
}