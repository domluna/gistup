gistup
======

CLI tool for uploading gists. Uses github OAuth2 token.

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
$ gistup -d "Some description of the gists" -f "file1.txt file2.txt" -p
```

