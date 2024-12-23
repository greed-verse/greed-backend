package shared

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	core *zerolog.Logger
}

var once sync.Once
var log zerolog.Logger

func NewLogger() *Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		// burstSampler := &zerolog.BurstSampler{
		// 	Burst:       3,
		// 	Period:      1 * time.Second,
		// 	NextSampler: &zerolog.BasicSampler{N: 5},
		// }

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel) // default to INFO
		}

		//TODO: Add lumberjack log rolling integration
		// if os.Getenv("APP_ENV") != "development" {}

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		//TODO: Add Level Sampling
		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})
	return &Logger{
		core: &log,
	}
}

func (l *Logger) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		defer func() {
			fmt.Println(time.Since(start).Nanoseconds())
			l.core.Info().
				Str("method", c.Method()).
				Str("time_elapsed", time.Since(start).String()).
				Str("route", c.Route().Path).
				Send()
		}()
		return c.Next()
	}
}

func (l *Logger) Core() *zerolog.Logger {
	return l.core
}
