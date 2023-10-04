package main
import (
    "fmt"
    "os"
    "github.com/patbcole117/PEP/pe"
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

    h := pe.GetDOSHeader(f)
    h.Print()
    
    h = pe.GetDOSStub(f)
    h.Print()
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
        } else if i % 8 == 0 {
            s = s + " "
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
