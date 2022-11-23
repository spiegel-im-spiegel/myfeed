package feed

import (
	"io"
	"strings"
	"time"

	"github.com/goark/errs"
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/atom"
	"github.com/spiegel-im-spiegel/myfeed/ecode"
	"github.com/spiegel-im-spiegel/myfeed/metadata"
)

// DecodeFeed create metadata.Metadata from feed data.
func DecodeFeed(t FeedType, r io.Reader) (*metadata.Metadata, error) {
	switch t {
	case TypeFlickr:
		return decodeFlickrFeed(r)
	case TypeBlog, TypeZenn:
		return decodeBlogFeed(r)
	}
	return nil, errs.Wrap(ecode.ErrNotSupportTarget, errs.WithContext("type", t.String()))
}

func decodeFlickrFeed(r io.Reader) (*metadata.Metadata, error) {
	f, err := (&atom.Parser{}).Parse(r)
	if err != nil {
		return nil, err
	}
	data := &metadata.Metadata{
		Title:       f.Title,
		Description: f.Subtitle,
	}
	for _, lnk := range f.Links {
		switch {
		case strings.EqualFold(lnk.Rel, "self"):
			data.FeedLink = lnk.Href
		case strings.EqualFold(lnk.Rel, "alternate"):
			data.Link = lnk.Href
		}
	}
	items := []*metadata.Item{}
	for _, i := range f.Entries {
		item := &metadata.Item{
			Title:     i.Title,
			Published: i.PublishedParsed,
			Updated:   i.UpdatedParsed,
		}
		image := &metadata.Image{}
		for _, lnk := range i.Links {
			switch {
			case strings.EqualFold(lnk.Rel, "alternate"):
				item.Link = lnk.Href
			case strings.EqualFold(lnk.Rel, "enclosure"):
				image.MimeType = lnk.Type
				image.URL = lnk.Href
			}
		}
		if i.Extensions != nil {
			if flickr, ok := i.Extensions["flickr"]; ok {
				if taken, ok := flickr["date_taken"]; ok && len(taken) > 0 {
					if parsedTaken, err := time.Parse(time.RFC3339, taken[0].Value); err == nil {
						parsedTaken = parsedTaken.In(time.UTC)
						image.Taken = &parsedTaken
					}
				}
			}
		}
		item.Images = []*metadata.Image{image}
		authors := []*metadata.Author{}
		for _, a := range i.Authors {
			authors = append(authors, &metadata.Author{Name: a.Name, URL: a.URI})
		}
		item.Authors = authors
		items = append(items, item)
	}
	data.Items = items
	return data, nil
}

func decodeBlogFeed(r io.Reader) (*metadata.Metadata, error) {
	f, err := gofeed.NewParser().Parse(r)
	if err != nil {
		return nil, err
	}
	data := &metadata.Metadata{
		FeedLink:    f.FeedLink,
		Title:       f.Title,
		Description: f.Description,
		Link:        f.Link,
	}
	authors := []*metadata.Author{}
	for _, a := range f.Authors {
		authors = append(authors, &metadata.Author{Name: a.Name})
	}
	data.Authors = authors
	items := []*metadata.Item{}
	for _, i := range f.Items {
		item := &metadata.Item{
			Title:       i.Title,
			Description: i.Description,
			Link:        i.Link,
			Published:   i.PublishedParsed,
			Updated:     i.UpdatedParsed,
		}
		authors := []*metadata.Author{}
		for _, a := range i.Authors {
			authors = append(authors, &metadata.Author{Name: a.Name})
		}
		item.Authors = authors
		if i.Image != nil {
			item.Images = []*metadata.Image{{Title: i.Image.Title, URL: i.Image.URL}}
		}
		items = append(items, item)
	}
	data.Items = items
	return data, nil
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
