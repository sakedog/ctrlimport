# ctrlimport

handle import path rule

## install

```sh
go install github.com/sakedog/ctrlimport/cmd/ctrlimport@latest
```

## how to use

```sh
$ go vet -vettool=$(which ctrlimport) \
	-ctrlimport.ng=github.com/foo/frameworks/... \       # can't import github.com/foo/frameworks/...
	-ctrlimport.ng=github.com/foo/adaptors/... \         # can't import github.com/foo/adaptors/...
	-ctrlimport.ok=github.com/foo/frameworks/testutils \ # CAN import github.com/foo/frameworks/testutils
	./entities/...
```
