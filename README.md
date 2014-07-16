gistup
======

CLI tool for uploading gists. Uses Github OAuth2 token.

```sh
$ export GITHUB_TOKEN="tokenhere"
```

The files to be uploaded are are space seperated.

```sh
# Public upload
$ gistup -d "Some description of the gists" -f "file1.txt file2.txt" -p
```

```sh
# Private upload
$ gistup -d "Some description of the gists" -f "file1.txt file2.txt"
```
## Installing

For go users

```sh
$ go get github.com/domluna/gistup
```

For everyone else

Binaries can be found under [release](https://github.com/domluna/gistup/releases) for Windows, Linux and OSX.

