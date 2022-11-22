# [myfeed] -- Manage my feed

[![check vulns](https://github.com/spiegel-im-spiegel/myfeed/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/myfeed/actions)
[![lint status](https://github.com/spiegel-im-spiegel/myfeed/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/myfeed/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/myfeed/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/myfeed.svg)](https://github.com/spiegel-im-spiegel/myfeed/releases/latest)

## Download and Build

```
$ go install github.com/spiegel-im-spiegel/myfeed@latest
```

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/myfeed/releases/latest).

## Usage

```
$ feed -h
Manage my feed

Usage:
  myfeed [flags]
  myfeed [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     print the version number

Flags:
      --debug   for debug
  -h, --help    help for myfeed

Use "myfeed [command] --help" for more information about a command.
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[myfeed]: https://github.com/spiegel-im-spiegel/myfeed "spiegel-im-spiegel/myfeed: Manage my feed"
