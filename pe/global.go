package pe

import (
    "encoding/binary"
    "fmt"
    "errors"
)

var (
    errWriteTooBig     = errors.New("write too large")
    errInvalidParam    = errors.New("param is invalid")
    errInvalidMachine  = errors.New("machine type is invalid")
    PRINT_COLS         = 8
)

func DWORDToI64(b []byte) int64 {
    return int64(binary.LittleEndian.Uint32(b))
}

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