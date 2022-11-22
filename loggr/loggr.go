package loggr

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spiegel-im-spiegel/myfeed/env"

	"github.com/goark/errs"
	"github.com/rs/zerolog"
)

func New(quiet bool) *zerolog.Logger {
	logger := zerolog.Nop()
	if env.ZerologLevel() == zerolog.NoLevel {
		return &logger
	}
	if env.EnableLogFile() {
		// make log directory
		dir := env.LogDir()
		_ = os.MkdirAll(dir, 0700) // エラーは無視
		// log file is $LOG_DIR/access.YYYYMMDD.log
		logpath := filepath.Join(dir, fmt.Sprintf("access.%s.log", time.Now().Local().Format("20060102")))
		// create logger
		if file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600); err != nil {
			logger = zerolog.New(os.Stdout).Level(env.ZerologLevel()).With().Timestamp().Logger()
			logger.Error().Interface("error", errs.Wrap(err)).Str("logpath", logpath).Msg("error in opening logfile")
		} else if quiet {
			logger = zerolog.New(file).Level(env.ZerologLevel()).With().Timestamp().Logger()
		} else {
			logger = zerolog.New(io.MultiWriter(
				file,
				zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
			)).Level(env.ZerologLevel()).With().Timestamp().Logger()
		}
		return &logger
	} else if quiet {
		return nil
	}
	logger = zerolog.New(os.Stderr).Level(env.ZerologLevel()).With().Timestamp().Logger()
	return &logger
}
