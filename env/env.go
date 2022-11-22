package env

import (
	"os"
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

func EnableLogFile() bool {
	return strings.EqualFold(os.Getenv("ENABLE_LOGFILE"), "true")
}

func LogDir() string {
	dir := os.Getenv("LOG_DIR")
	if len(dir) == 0 {
		// directory name is ${XDG_CACHE_HOME}/${ServiceName}
		dir = cache.Dir(ServiceName)
	}
	return dir
}

func LogLevel() LoggerLevel {
	return getLogLevel(os.Getenv("LOGLEVEL"))
}

func ZerologLevel() zerolog.Level {
	return LogLevel().ZerlogLevel()
}

func EmailConfig() (*emailconf.Config, error) {
	path := os.Getenv("ERRORMAIL_FILE")
	if len(path) == 0 {
		return &emailconf.Config{}, errs.Wrap(ecode.ErrInvalidConfFile)
	}
	cfg, err := emailconf.ImportFile(path)
	if err != nil {
		return &emailconf.Config{}, errs.Wrap(err, errs.WithContext("path", path))
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
