package zlog

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	// Teal ...
	Teal = Color("\033[1;36m%s\033[0m")
	// Yellow ...
	Yellow = Color("\033[35m%s\033[0m")
)

const (
	EnvLocal = "local"

	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite

	colorBold     = 1
	colorDarkGray = 90
)

var Logger zerolog.Logger

// Color ...
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// Graylog 的錯誤等級
const (
	levelEmerg   = int8(0)
	levelAlert   = int8(1)
	levelCrit    = int8(2)
	levelErr     = int8(3)
	levelWarning = int8(4)
	levelNotice  = int8(5)
	levelInfo    = int8(6)
	levelDebug   = int8(7)
)

type callerHook struct {
	enableCaller   bool
	callerMinLevel int8
}

func (h callerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {

	if h.enableCaller && level >= zerolog.Level(h.callerMinLevel) {
		if _, file, line, ok := runtime.Caller(3); ok {
			e.Str("file", fmt.Sprintf("%s:%d", file, line))
		}
	}
}

func Setup(config *Config) {
	zerolog.DisableSampling(true)
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	level := zerolog.Level(config.Level)

	var logger zerolog.Logger
	if config.Env == EnvLocal {
		output := zerolog.ConsoleWriter{Out: os.Stdout}
		output.FormatLevel = func(i interface{}) string {
			var l string
			if ll, ok := i.(string); ok {
				switch ll {
				case "trace":
					l = colorize("TRACE", colorMagenta)
				case "debug":
					l = colorize("DEBUG", colorYellow)
				case "info":
					l = colorize("INFO ", colorGreen)
				case "warn":
					l = colorize("WARN ", colorRed)
				case "error":
					l = colorize(colorize("ERROR", colorRed), colorBold)
				case "fatal":
					l = colorize(colorize("FATAL", colorRed), colorBold)
				case "panic":
					l = colorize(colorize("PANIC", colorRed), colorBold)
				default:
					l = colorize("???", colorBold)
				}
			} else {
				l = "???"
			}
			return fmt.Sprintf("|%-10s|", l)
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%-50s", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s = ", Teal(i))
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatTimestamp = func(i interface{}) string {
			return fmt.Sprintf("%s", Yellow(i))
		}
		output.PartsOrder = []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
			zerolog.CallerFieldName,
		}
		logger = zerolog.New(output)
	} else {
		logger = zerolog.New(os.Stdout)
	}

	// 設定基礎日誌欄位
	fields := make(map[string]interface{})
	if config.AppID != "" {
		fields["app_id"] = config.AppID
	}
	if config.Env != "" {
		fields["env"] = config.Env
	}

	zctx := logger.Hook(callerHook{
		enableCaller:   config.EnableCaller,
		callerMinLevel: config.CallerMinLevel,
	}).With().Timestamp()

	if len(fields) > 0 {
		zctx = zctx.Fields(fields)
	}

	log.Logger = zctx.Logger().Level(level)
}

// colorize returns the string s wrapped in ANSI code c, unless disabled is true.
func colorize(s interface{}, c int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}
