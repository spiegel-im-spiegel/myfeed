package fetch

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/goark/errs"
	ftch "github.com/goark/fetch"
)

// Feed fetches feed data from URL.
func Feed(ctx context.Context, u *url.URL) (io.ReadCloser, error) {
	resp, err := ftch.New().Get(u, ftch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	return resp.Body(), nil
}

// FlickrFeedURL makes URL from Flickr ID.
func FlickrFeedURL(flickrId string) (*url.URL, error) {
	u, err := url.Parse("https://www.flickr.com/services/feeds/photos_public.gne?format=atom")
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("flickrId", flickrId))
	}
	q := u.Query()
	q.Add("id", flickrId)
	u.RawQuery = q.Encode()
	return u, nil
}

func ZennFeedURL(zennId string) (*url.URL, error) {
	return url.Parse(strings.Join([]string{"https://zenn.dev", zennId, "feed"}, "/"))
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
