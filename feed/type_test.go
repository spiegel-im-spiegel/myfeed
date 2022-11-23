package feed_test

import (
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/myfeed/feed"
)

func TestTypeList(t *testing.T) {
	res := strings.Join([]string{"blog", "flickr", "zenn"}, ",")
	s := strings.Join(feed.TypeList(), ",")
	if s != res {
		t.Errorf("TypeList() = \"%v\", want \"%v\".", s, res)
	}
}

func TestGetTypeFrom(t *testing.T) {
	testCases := []struct {
		s string
		t feed.FeedType
	}{
		{s: "", t: feed.TypeUnknown},
		{s: "blog", t: feed.TypeBlog},
		{s: "flickr", t: feed.TypeFlickr},
		{s: "zenn", t: feed.TypeZenn},
		{s: "Blog", t: feed.TypeBlog},
		{s: "Flickr", t: feed.TypeFlickr},
		{s: "Zenn", t: feed.TypeZenn},
		{s: "foo", t: feed.TypeUnknown},
	}

	for _, tc := range testCases {
		tp := feed.GetTypeFrom(tc.s)
		if tp != tc.t {
			t.Errorf("GetTypeFrom(%v) = \"%v\", want \"%v\".", tc.s, tp, tc.t)
		}
	}
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
