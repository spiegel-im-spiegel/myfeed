package env

import "github.com/rs/zerolog"

type LoggerLevel int

const (
	LevelNop LoggerLevel = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

var levelMap = map[string]LoggerLevel{
	"nop":   LevelNop,
	"error": LevelError,
	"warn":  LevelWarn,
	"info":  LevelInfo,
	"debug": LevelDebug,
	"trace": LevelTrace,
}

var zerologLevelMap = map[LoggerLevel]zerolog.Level{
	LevelNop:   zerolog.NoLevel,
	LevelError: zerolog.ErrorLevel,
	LevelWarn:  zerolog.WarnLevel,
	LevelInfo:  zerolog.InfoLevel,
	LevelDebug: zerolog.DebugLevel,
	LevelTrace: zerolog.TraceLevel,
}

func getLogLevel(s string) LoggerLevel {
	if lvl, ok := levelMap[s]; ok {
		return lvl
	}
	return LevelInfo
}

func (lvl LoggerLevel) ZerlogLevel() zerolog.Level {
	if l, ok := zerologLevelMap[lvl]; ok {
		return l
	}
	return zerolog.InfoLevel
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
