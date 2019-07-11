## Go Examples

Repository of Go snippets for reference purposes

Every subdirectory contains a `Makefile` which implements a `run` rule to run the code in that directory

```
$ make run
```

Some directories differ from the above statement. How to run the examples from those directories are described below

### cross-compile

```
$ make release
```

### gRPC

```
$ make run-server
$ make run-client
```

