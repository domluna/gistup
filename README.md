gistup
======

Simple CLI for uploading files as gists.

### Install

```sh
$ go get github.com/domluna/gistup
```

### Usage

Public gist.

```sh
$ gistup -d "Awesome gist" file1.txt file2.txt
```

Secret Gist.

```sh
$ gistup -d "secret upload" -s secret.txt
```

### Help

```sh
gistup uploads files to Github as gists

Usage:
  gistup [flags]
Flags:
  -d, --description="": description of gist
  -h, --help=false: help for gistup
  -s, --secret=false: gist set to secret
  -t, --token="": use this github token
```
