# PoC embed binary

This is a proof of concept that combines two executables together into a single executable that gets extracted when it's
used for the first time.

### Example

Run `make build` to build 3 executables (`entrypoint/main`, `plugin/main` and `main`). Notice that size of `main` is
around 6 MB at this point. When running `main` it extracts the other two executables, one overwrites the existing `main`
executable, the other is extracted into `plugins/plug`. Notice how the new main now uses only around 2MB, since the
original was overwritten. The new extracted `main` is then called with the same environment and arguments as the
original `main`, all stdout and stderr outputs are redirected to the original outputs.

The output of the first run:
```
$ ./main
extracting embedded files
overwritten main, running the new main

hello main!
Hello from embedded!
bye!
```

The output of the second run (the extracting step is skipped):
```
$ ./main
hello main!
Hello from embedded!
bye!
```
