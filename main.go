package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	kvstore "github.com/varigg/kvstore/internal"
)

var store *kvstore.KVStore
var showTxDepth bool

// checkArgs verifies that the correct number of arguments have been provided
func checkArgs(args []string, min int, max int) error {
	if max != -1 && len(args) > max {
		return fmt.Errorf("unexpected parameters in command: %v", args[max:])
	}
	if len(args) < min {
		return fmt.Errorf("command requires at least %d parameters", min)
	}
	return nil
}

// main is just used to setup and execute the run function
// that way it is easier to test the main logic
func main() {
	// cobra would be overkill for one flag
	flag.BoolVar(&showTxDepth, "showTransactionDepth", false, "modifies the prompt to give an indication of nested transaction depth")
	flag.Parse()

	if err := run(os.Args, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// prompt will either print 1 '>' or a number of '>' matching the number of nested transactions
// on the stack if the `showTransactionDepth` flag has been set
// It's not in the requirements, but I think it's a useful feature
func prompt() {

	if showTxDepth {
		for i := 0; i <= store.GetStackLevel(); i++ {
			fmt.Print(">")
		}
	} else {
		fmt.Print(">")
	}
}

// run contains the main read-eval-print loop
// os level inputs/outputs have been abstracted for easier testing
func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {

	store = kvstore.NewKVStore()

	in := bufio.NewReader(stdin)
	out := bufio.NewWriter(stdout)
	errOut := bufio.NewWriter(stderr)

	for {
		// writes directly to os.Stdout so test doesn't get cluttered
		prompt()
		input, err := in.ReadString('\n')
		if err != nil {
			return err
		}
		if len(input) == 1 {
			continue
		}
		args := strings.Split(strings.Trim(input, " \n"), " ")
		var result interface{}

		switch args[0] {
		case "READ":
			err = checkArgs(args[1:], 1, 1)
			if err == nil {
				result = store.Read(args[1])
				if result == nil {
					err = fmt.Errorf("Key not found: " + args[1])
				} else {
					_, err = out.WriteString(fmt.Sprintln(result))
					if err != nil {
						return err
					}
					err = out.Flush()
					if err != nil {
						return err
					}
				}
			}
		case "WRITE":
			err = checkArgs(args[1:], 2, -1)
			if err == nil {
				store.Write(args[1], strings.Join((args[2:]), " "))
			}
		case "START":
			err = checkArgs(args[1:], 0, 0)
			if err == nil {
				store.Start()
			}
		case "ABORT":
			err = checkArgs(args[1:], 0, 0)
			if err == nil {
				err = store.Abort()
			}
		case "COMMIT":
			err = checkArgs(args[1:], 0, 0)
			if err == nil {
				err = store.Commit()
			}
		case "DELETE":
			err = checkArgs(args[1:], 1, 1)
			if err == nil {
				store.DeleteKey(args[1])
			}
		case "QUIT":
			return nil
		default:
			_, err = errOut.WriteString(fmt.Sprintf("Unknown command: %s\n", args[0]))
			if err != nil {
				return err
			}
			err = errOut.Flush()
			if err != nil {
				return err
			}
		}

		if err != nil {
			_, err = errOut.WriteString(err.Error() + "\n")
			if err != nil {
				return err
			}
			err = errOut.Flush()
			if err != nil {
				return err
			}
		}
	}

}
