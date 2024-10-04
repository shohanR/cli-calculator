package history

import (
	"fmt"
)

// FIXME: this function should be moved to another file
// FIXME: it prints 4/0=undefined that should not be printed.
func HistoryPrinter(history []string) {
	// FIXME: this check is useless since it doesn't give you an error if we don't have any element
	// this checks if any valid operation has been done or not.
	if len(history) > 0 {
		for indexNum, indexVal := range history {
			fmt.Println(indexNum+1, ". ", indexVal)
		}
	} else {
		fmt.Println("\n\nNo valid operation done!")
	}
	// TODO: print the history without closing the calculator. The function name is "HistoryPrinter" so it doesn't have to close the whole program
}
