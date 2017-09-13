package logger

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"github.com/rickychandra/goboil/common/stringutil"
)

//Log to stderr.
var (
	std = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
)

//Init inits with necessary data.
var Init = func() {
}

//Error logs error message.
func Error(ctx context.Context, err error) {
	std.Printf("[%s] %s", "ERROR", err.Error())
}

func Info(ctx context.Context, msg *string) {
	std.Printf("[%s] %s", "INFO", stringutil.GetOrDefault(msg))
}

func Warn(ctx context.Context, msg *string) {
	std.Printf("[%s] %s", "WARN", stringutil.GetOrDefault(msg))
}
