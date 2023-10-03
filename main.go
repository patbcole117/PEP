package main
import (
    _ "io"
    "fmt"
    "os"
)

func main() {
    var args = os.Args
    if len(args) < 2 || len(args) > 2 {
        Help()
        os.Exit(1)
    }

    f, err := os.Open(args[1])
    if err != nil {
        PrintErr(err)
    }
    defer f.Close()

    MS_DOS_HEADER := GetMS_DOS_HEADER(f)
    MS_DOS_STUB := GetMS_DOS_STUB(f)
    fmt.Println(PrintBytes(MS_DOS_HEADER))
    fmt.Println(string(MS_DOS_HEADER))
    fmt.Println(PrintBytes(MS_DOS_STUB))
    fmt.Println(string(MS_DOS_STUB))
    
}

func Help() {
    fmt.Println("[?] Please provide a file to parse.")
}

func PrintErr(err error) {
    fmt.Println("[!] Error:", err.Error())
    os.Exit(1)
}

func PrintBytes(b []byte) string {
   var s string
    for i := 0; i < len(b); i++ {
        if i % 16 == 0 {
            s = s + "\n"
        }
        s = s + fmt.Sprintf("\\x%02x", b[i])
    }
    s = s + "\n"
    return s   
}

func ResetF(f *os.File) int64 {
    o, err := f.Seek(0,0)
    if  err != nil {
        panic(err)
    } 
    return o
}

func GetMS_DOS_HEADER(f *os.File) []byte {
    var b = make([]byte, 64)
    f.Read(b)
    ResetF(f)
    return b
}

func GetMS_DOS_STUB(f *os.File) []byte {
    var b = make([]byte, 64)
    f.Seek(64, 0)
    f.Read(b)
    ResetF(f)
    return b
}
