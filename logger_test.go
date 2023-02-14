package sensitive

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		name   string
		logger Logger
		args   string
	}{
		{name: "Test Print", logger: Logger{Enable: true}, args: "some string"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer func() {
				log.SetOutput(os.Stderr)
			}()
			l := tt.logger
			l.Print(tt.args)
			assert.NotEmpty(t, buf.String())
			l.Printf(tt.args)
			assert.NotEmpty(t, buf.String())
		})
	}
}
