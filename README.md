gistup
======

Simple CLI for uploading files as gists.

### Install

```
$ go get github.com/domluna/gistup
```

### Usage

Public gist.

```
$ gistup -d "Awesome gist" file1.txt file2.txt
```

Secret Gist.

```
$ gistup -d "secret upload" -s secret.txt
```

### Help

```
gistup uploads files to Github as gists

Usage:
  gistup [flags]
Flags:
  -d, --description="": description of gist
  -h, --help=false: help for gistup
  -s, --secret=false: gist set to secret
  -t, --token="": use this github token
```
