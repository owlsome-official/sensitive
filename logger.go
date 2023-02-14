package sensitive

import "log"

type Logger struct {
	Enable bool
}

type LoggerStructure interface {
	Print(v ...any)
	Printf(format string, v ...any)
}

func NewLogger(enable bool) LoggerStructure {
	return &Logger{
		Enable: enable,
	}
}

func (l *Logger) Print(v ...any) {
	if l.Enable {
		log.Println(v...)
	}
}
func (l *Logger) Printf(format string, v ...any) {
	if l.Enable {
		log.Printf(format, v...)
	}
}
