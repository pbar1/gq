# Introduction

`gq` is a command line [filter][1] (à la sed/awk) used for rendering [Go templates][2],
along with a few other useful data transformation features. It is heavily inspired
by [Kubectl][4] and [Helm][5].

## Features

- Go template rendering ([Sprig][3] functions included)
- JSONPath querying ([using Kubectl syntax][6])
- Data format transformation
- Line-based processing and whole-input processing

## Install

### Manual

Download the binary from the [GitHub Releases][7] page and place it on your path.

### Go

Run the following command. Make sure you have `$GOPATH/bin` on your `$PATH`.

```sh
go get github.com/pbar1/gq
```

### Docker

TODO

## Quick Start

Fetch a GitHub repo's tags and parse the response:

```sh
curl --silent https://api.github.com/repos/torvalds/linux/git/refs/tags \
| gq -r 'printf "%s -> %s\n" (.ref | replace "refs/tags/" "") .object.sha'
```

<!-- Sources -->

[1]: https://en.wikipedia.org/wiki/Filter_(software)
[2]: https://golang.org/pkg/text/template/
[3]: http://masterminds.github.io/sprig/
[4]: https://kubectl.docs.kubernetes.io/
[5]: https://helm.sh/
[6]: https://kubernetes.io/docs/reference/kubectl/jsonpath/
[7]: https://github.com/pbar1/gq/releases
