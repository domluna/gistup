gistup
======

Simple CLI for uploading files as gists.

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
