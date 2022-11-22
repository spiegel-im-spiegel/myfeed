package facade

import (
	"context"
	"io"
	"net/url"
	"os"
	"os/signal"

	"github.com/goark/errs"
	"github.com/goark/gocli/config"
	"github.com/goark/gocli/rwi"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/myfeed/env"
	"github.com/spiegel-im-spiegel/myfeed/fetch"
	"github.com/spiegel-im-spiegel/myfeed/loggr"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newFetchCmd(ui *rwi.RWI) *cobra.Command {
	fetchCmd := &cobra.Command{
		Use:     "fetch [flags] [<URL string>]",
		Aliases: []string{"f"},
		Short:   "fetch feed data",
		Long:    "fetch feed data from web",
		RunE: func(cmd *cobra.Command, args []string) error {
			//load ${XDG_CONFIG_HOME}/${ServiceName}/env file
			if err := godotenv.Load(config.Path(env.ServiceName, "env")); err != nil {
				//load .env file
				_ = godotenv.Load()
			}
			//Options
			flickrId, err := cmd.Flags().GetString("flickr-id")
			if err != nil {
				return debugPrint(ui, errs.New("error in --flickr-id option", errs.WithCause(err)))
			}
			if len(flickrId) == 0 && len(args) == 0 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))

			}
			// cancel event by Ctrl+C interrupt
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()
			// logger
			logger = loggr.New(true) // quiet mode

			var u *url.URL
			if len(flickrId) > 0 {
				u, err = fetch.FlickrFeedURL(flickrId)
			} else {
				u, err = url.Parse(args[0])
			}
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}

			rc, err := fetch.Feed(ctx, u)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			defer rc.Close()

			_, err = io.Copy(ui.Writer(), rc)
			return debugPrint(ui, errs.Wrap(err))
		},
	}
	fetchCmd.Flags().StringP("flickr-id", "i", "", "Flickr user ID")

	return fetchCmd
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
