package main
import (
    "fmt"
    "os"
)

type HEADER_SECTION struct {
	Raw		[]byte
	Vars 	map[string][]byte
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
		Raw: b,
		Vars: map[string][]byte{
			"e_magic":		b[0:2],
			"e_cblp":		b[2:4],
			"e_cp":			b[4:6],
			"e_crlc":		b[6:8], 
			"e_cparhdr":	b[8:10],
			"e_minalloc":	b[10:12],
			"e_maxalloc":	b[12:14],
			"e_ss":			b[14:16],
			"e_sp":			b[16:18],
			"e_csum":		b[18:20],
			"e_ip":			b[20:22],
			"e_cs":			b[22:24],
			"e_lfarlc":		b[24:26],
			"e_ovno":		b[26:28],
			"e_res":		b[28:36],
			"e_oemid":		b[36:38],
			"e_oeminfo":	b[38:40],
			"e_res2":		b[40:60],
			"e_lfanew":		b[60:64],
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
    fmt.Println(PrintBytes(h.Raw))
    for k, v := range h.Vars {
        fmt.Printf("[+] %s:\t%s,\t\\x%02x\n", k, v, v)
    }
}

func ResetF(f *os.File) int64 {
    o, err := f.Seek(0,0)
    if  err != nil {
        panic(err)
    } 
    return o
}
