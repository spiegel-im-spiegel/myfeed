package env

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/gocli/cache"
	"github.com/rs/zerolog"
	"github.com/spiegel-im-spiegel/myfeed/ecode"
	"github.com/spiegel-im-spiegel/myfeed/env/emailconf"
)

const (
	ServiceName = "myfeed"
)

// LogDir returns path for log files.
func DataDir() string {
	dir := os.Getenv("DATA_DIR")
	if len(dir) == 0 {
		// directory name is ${XDG_CACHE_HOME}/${ServiceName}/data
		dir = filepath.Join(cache.Dir(ServiceName), "data")
	}
	return dir
}

// EnableLogFile returns true if environment value ENABLE_LOGFILE is "true".
func EnableLogFile() bool {
	return strings.EqualFold(os.Getenv("ENABLE_LOGFILE"), "true")
}

// QuietLog returns true if environment value QUIET_LOG is "true".
func QuietLog() bool {
	return strings.EqualFold(os.Getenv("QUIET_LOG"), "true")
}

// LogDir returns path for log files.
func LogDir() string {
	dir := os.Getenv("LOG_DIR")
	if len(dir) == 0 {
		// directory name is ${XDG_CACHE_HOME}/${ServiceName}/log
		dir = filepath.Join(cache.Dir(ServiceName), "log")
	}
	return dir
}

// LogLevel returns log level.
func LogLevel() LoggerLevel {
	return getLogLevel(os.Getenv("LOGLEVEL"))
}

// ZerologLevel returns log level for zerolog.
func ZerologLevel() zerolog.Level {
	return LogLevel().ZerlogLevel()
}

// EmailConfig returns configuration for error email.
func EmailConfig() (*emailconf.Config, error) {
	s := os.Getenv("ERRORMAIL_CONFIG")
	if len(s) == 0 {
		return nil, nil
	}
	cfg, err := emailconf.Import([]byte(s))
	if err != nil {
		return nil, errs.Wrap(ecode.ErrInvalidConfFile, errs.WithCause(err), errs.WithContext("ERRORMAIL_CONFIG", s))
	}
	return cfg, nil
}

/* Copyright 2022 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
