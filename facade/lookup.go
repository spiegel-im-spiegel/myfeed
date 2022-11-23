package facade

import (
	"context"
	"io"
	"os"
	"os/signal"

	"github.com/goark/errs"
	"github.com/goark/gocli/config"
	"github.com/goark/gocli/rwi"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/myfeed/env"
	"github.com/spiegel-im-spiegel/myfeed/feed"
	"github.com/spiegel-im-spiegel/myfeed/loggr"
	"github.com/spiegel-im-spiegel/myfeed/lookup"
)

// newVersionCmd returns cobra.Command instance for show sub-command
func newLookupCmd(ui *rwi.RWI) *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:     "lookup [flags] [<URL string>]",
		Aliases: []string{"look", "l"},
		Short:   "lookup feed data",
		Long:    "lookup feed data from web",
		RunE: func(cmd *cobra.Command, args []string) error {
			//load ${XDG_CONFIG_HOME}/${ServiceName}/env file
			if err := godotenv.Load(config.Path(env.ServiceName, "env")); err != nil {
				//load .env file
				_ = godotenv.Load()
			}
			//Options
			rawflag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return debugPrint(ui, errs.New("error in --raw option", errs.WithCause(err)))
			}
			flickrId, err := cmd.Flags().GetString("flickr-id")
			if err != nil {
				return debugPrint(ui, errs.New("error in --flickr-id option", errs.WithCause(err)))
			}
			zennId, err := cmd.Flags().GetString("zenn-id")
			if err != nil {
				return debugPrint(ui, errs.New("error in --zenn-id option", errs.WithCause(err)))
			}
			if len(flickrId) == 0 && len(zennId) == 0 && len(args) == 0 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))
			} else if len(flickrId) > 0 && len(zennId) > 0 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))
			} else if (len(flickrId) > 0 || len(zennId) > 0) && len(args) > 0 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))
			}
			// cancel event by Ctrl+C interrupt
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()
			// logger
			if batchFlag {
				logger = loggr.New(env.QuietLog()) // quiet mode
			}

			var rc io.ReadCloser
			if len(flickrId) > 0 {
				rc, err = lookup.Lookup(ctx, feed.TypeFlickr, flickrId, rawflag)
			} else if len(zennId) > 0 {
				rc, err = lookup.Lookup(ctx, feed.TypeZenn, zennId, rawflag)
			} else {
				rc, err = lookup.Lookup(ctx, feed.TypeBlog, args[0], rawflag)
			}
			if err != nil {
				return debugPrint(ui, errs.Wrap(err))
			}
			defer rc.Close()

			_, err = io.Copy(ui.Writer(), rc)
			return debugPrint(ui, errs.Wrap(err))
		},
	}
	lookupCmd.Flags().StringP("flickr-id", "i", "", "Flickr user ID")
	lookupCmd.Flags().StringP("zenn-id", "z", "", "Zenn user ID")
	lookupCmd.Flags().BoolP("raw", "r", false, "Output raw data")

	return lookupCmd
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
