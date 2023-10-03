package main
import (
    "fmt"
    "os"
)

type HEADER_SECTION struct {
    Title       string
	Raw		    []byte
	VarNames 	[]string
    VarValues   [][]byte
}

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

    h := GetIMAGE_DOS_HEADER(f)
    h.Print()
    
}

func GetIMAGE_DOS_HEADER(f *os.File) HEADER_SECTION {
    var b = make([]byte, 64)
    f.Read(b)
    ResetF(f)
	return HEADER_SECTION {
        Title: "IMAGE_DOS_HEADER",
		Raw: b,
		VarNames: []string{
			"e_magic",
			"e_cblp",
			"e_cp",
			"e_crlc",
			"e_cparhdr",
			"e_minalloc",
			"e_maxalloc",
			"e_ss",
			"e_sp",
			"e_csum",
			"e_ip",
			"e_cs",
			"e_lfarlc",
			"e_ovno",
			"e_res",
			"e_oemid",
			"e_oeminfo",
			"e_res2",
			"e_lfanew",
		},
        VarValues: [][]byte{
            b[0:2],
            b[2:4],
            b[4:6],
            b[6:8], 
            b[8:10],
            b[10:12],
            b[12:14],
            b[14:16],
            b[16:18],
            b[18:20],
            b[20:22],
            b[22:24],
            b[24:26],
            b[26:28],
            b[28:36],
            b[36:38],
            b[38:40],
            b[40:60],
            b[60:64],
        },
	}
}

func GetMSDOSSTUB(f *os.File) []byte {
    var b = make([]byte, 64)
    f.Seek(64, 0)
    f.Read(b)
    ResetF(f)
    return b
}

// Helpers

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

func (h *HEADER_SECTION) Print() {
    fmt.Println(h.Title)
    fmt.Println(PrintBytes(h.Raw))
    for i := range h.VarNames {
        fmt.Printf("[+] %s:\t%s,\t\\x%02x\n", h.VarNames[i], h.VarValues[i], h.VarValues[i])
    }
}

func ResetF(f *os.File) int64 {
    o, err := f.Seek(0,0)
    if  err != nil {
        panic(err)
    } 
    return o
}
