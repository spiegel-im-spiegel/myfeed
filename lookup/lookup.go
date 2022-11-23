package lookup

import (
	"context"
	"io"
	"net/url"

	"github.com/goark/errs"
	"github.com/spiegel-im-spiegel/myfeed/feed"
	"github.com/spiegel-im-spiegel/myfeed/fetch"
	"github.com/spiegel-im-spiegel/myfeed/metadata"
)

// Lookup returs feed data (raw data or metadata encoded JSON).
func Lookup(ctx context.Context, t feed.FeedType, id string, rawflag bool) (io.ReadCloser, error) {
	var u *url.URL
	var err error
	switch t {
	case feed.TypeBlog:
		u, err = url.Parse(id)
	case feed.TypeFlickr:
		u, err = fetch.FlickrFeedURL(id)
	case feed.TypeZenn:
		u, err = fetch.ZennFeedURL(id)
	}
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("type", t.String()), errs.WithContext("id", id), errs.WithContext("rawflag", rawflag))
	}
	rc, err := fetch.Feed(ctx, u)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("type", t.String()), errs.WithContext("id", id), errs.WithContext("rawflag", rawflag))
	}
	if rawflag {
		return rc, nil
	}
	defer rc.Close()
	data, err := feed.DecodeFeed(t, rc)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("type", t.String()), errs.WithContext("id", id), errs.WithContext("rawflag", rawflag))
	}
	if t == feed.TypeFlickr || t == feed.TypeZenn {
		data.ID = id
	}
	buf := metadata.NewBuffer()
	if err := buf.EncodeMetadata(data); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("type", t.String()), errs.WithContext("id", id), errs.WithContext("rawflag", rawflag))
	}
	return buf, nil
}
