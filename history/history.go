package history

import (
	"fmt"
)

func HistoryPrinter(history []string) {
	// FIXME: this check is useless since it doesn't give you an error if we don't have any element
	// this checks if any valid operation has been done or not.
	// try avoiding "else" since it's not a best practice to use them.
	// https://dev.to/dglsparsons/write-better-code-and-be-a-better-programmer-by-never-using-else-statements-4dbl#:~:text=Else%20statements%20are%20problematic%20as,through%2C%20and%20hurts%20the%20readability.
	// https://developer20.com/elseless/
	if len(history) > 0 {
		for indexNum, indexVal := range history {
			fmt.Println(indexNum+1, ". ", indexVal)
		}
	} else {
		fmt.Println("\n\nNo valid operation done!")
	}
}
