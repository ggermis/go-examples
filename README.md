## Go Examples

Repository of Go snippets for reference purposes

Every subdirectory contains a `Makefile` which implements a `run` rule to run the code in that directory

```
$ make run
```

Some directories differ from the above statement. How to run the examples from those directories is described below

### cross-compile

```
$ make release
```

### gRPC

```
$ make run-server
$ make run-client
```

### webassembly

To view the output of the main.wasm file, look at the javascript console in your browser using the inspector

```
$ make run
$ make open
```