## KVStore

A simple command line utility simulating a key-value store.

### Build
The project comes with a [drone](drone.io) pipeline that will produce a linux binary. You can also run `make build` to build it on your platform.

### Usage
`make run` will build and start the program. You can invoke the program directly after building it. This program take an optional parameter `-showTransactionDepth`. Setting it will modify the prompt to visualize the number of nested transactions.

Commands
* READ Reads and prints, to stdout, the val associated with key. If the value is not present an error
is printed to stderr.
*  WRITE Stores val in key.
*  DELETE Removes a key from store. Future READ commands on that key will return an error.
*  START Start a transaction.
*  COMMIT Commit a transaction. All actions in the current transaction are committed to the
parent transaction or the root store. If there is no current transaction an error is output to
stderr.
*  ABORT Abort a transaction. All actions in the current transaction are discarded.
*  QUIT Exit the REPL cleanly. 