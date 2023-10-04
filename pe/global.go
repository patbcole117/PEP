package pe

import (
	"fmt"
	"errors"
)

var (
	errWriteTooLarge = errors.New("write too large")
	PRINT_COLS	= 8
)

func PrintBytes(b []byte) string {
	var s string
	 for i := 0; i < len(b); i++ {
		 if i % PRINT_COLS == 0 {
			 s = s + "\n"
		 } else if i % (PRINT_COLS/2) == 0 {
			 s = s + " "
		 }
		 s = s + fmt.Sprintf("\\x%02x", b[i])
	 }
	 s = s + "\n"
	 return s   
 }