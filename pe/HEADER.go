package pe

import "fmt"

type HeaderElement struct {
    Name    string
    Offset  uint64
    Size    uint64
}

type ImageHeader struct {
    Title       string
    Raw         []byte
    Offset      uint64
    Size        uint64
    Elements    []HeaderElement 
}

func (h *ImageHeader) Print() {
    fmt.Printf("\n\n[#] %s [#]\n", h.Title)
    fmt.Println(PrintBytes(h.Raw))
    for _, e := range h.Elements {
        fmt.Printf("[+] %s:", e.Name)
        v := h.Raw[e.Offset:e.Offset + e.Size]
        if e.Size > 8 {
            fmt.Printf("\n%s\n%s\n", v, PrintBytes(v))
        }else {
            fmt.Printf("\t%s\t%s", v, PrintBytes(v))
        }
    }
}

func PrintBytes(b []byte) string {
    var s string
    for i := 0; i < len(b); i++ {
        if i % 8 == 0 && i != 0 {
            s = s + "\n"
        } else if i % 4 == 0 && i != 0 {
            s = s + " "
        }
        s = s + fmt.Sprintf("\\x%02x", b[i])
    }
    s = s + "\n"
    return s
}
