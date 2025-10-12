package musicall

import (
	"log"
)

// LogError must be used to log error on the terminal. Indeed, for a
// more clear code, some functions that manipulate the music data don't
// return an error, even if an error occurs (for example if trying to
// get the note of name "Ré" while it is registered with the name "Re",
// without accent). Instead, we call the LogError function to print the
// error message. By default, the print uses log.Fatal that interupts the
// process (and then you know explicitly the error and can fix it).
var LogError = log.Fatalf

//var LogError = fmt.Printf
