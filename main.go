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

    f, err := os.OpenFile(args[1], os.O_RDWR, 0644)
    if err != nil {
        PrintErr(err)
    }
    defer f.Close()

    pe.Print_IMAGE_DOS_HEADER(f)
}

func Help() {
    fmt.Println("[?] Please provide a file to parse.")
}

func PrintErr(err error) {
    fmt.Println("[!] Error:", err.Error())
    os.Exit(1)
}



func ResetF(f *os.File) int64 {
    o, err := f.Seek(0,0)
    if  err != nil {
        panic(err)
    } 
    return o
}
