package feed

import (
	"sort"
	"strings"
)

// FeedType is type of feed
type FeedType int

const (
	TypeUnknown FeedType = iota
	TypeBlog
	TypeFlickr
	TypeZenn
)

var typeMap = map[string]FeedType{
	"blog":   TypeBlog,
	"flickr": TypeFlickr,
	"zenn":   TypeZenn,
}

func TypeList() []string {
	var lst []string
	for k := range typeMap {
		lst = append(lst, k)
	}
	sort.Strings(lst)
	return lst
}

func GetTypeFrom(s string) FeedType {
	if t, ok := typeMap[strings.ToLower(s)]; ok {
		return t
	}
	return TypeUnknown
}

func (t FeedType) String() string {
	for k, v := range typeMap {
		if t == v {
			return k
		}
	}
	return "Unknown"
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
