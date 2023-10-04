package pe

import "os"

const (
    IMAGE_DOS_HEADER_SIZE   = 64
    IMAGE_DOS_HEADER_OFFSET = 0
    IMAGE_DOS_STUB_SIZE     = 64
    IMAGE_DOS_STUB_OFFSET   = 64
    e_magic = iota
    e_cblp
    e_cp
    e_crlc
    e_cparhdr
    e_minalloc
    e_maxalloc
    e_ss
    e_sp
    e_csum
    e_ip
    e_cs
    e_lfarlc
    e_ovno
    e_res
    e_oemid
    e_oeminfo
    e_res2
    e_elfanew
)

type IMAGE_DOS_HEADER struct {
    Title   string
    Raw     []byte
}

func GetDOSH(f *os.File) []byte {
    var h IMAGE_DOS_HEADER
    var b make([]byte, IMAGE_DOS_HEADER_SIZE)
    f.Seek(IMAGE_DOS_HEADER_OFFSET, 0)
    f.Read(b)
    return b
}

func GetDOSH_emagic() []byte {
    
}

func GetDOSHeaderElements() []HeaderElement {
    return []HeaderElement{
       HeaderElement{
            Name: "e_magic",
            Offset: 0,
            Size:   2,
        },
        HeaderElement{
            Name: "e_cblp",
            Offset: 2,
            Size:   2,
        },
        HeaderElement{
            Name: "e_cp",
            Offset: 4,
            Size:   2,
        },
        HeaderElement{
            Name: "e_crlc",
            Offset: 6,
            Size:   2,
        },
        HeaderElement{
            Name: "e_cparhdr",
            Offset: 8,
            Size:   2,
        },
        HeaderElement{
            Name: "e_minalloc",
            Offset: 10,
            Size:   2,
        },
        HeaderElement{
            Name: "e_maxalloc",
            Offset: 12,
            Size:   2,
        },
        HeaderElement{
            Name: "e_ss",
            Offset: 14,
            Size:   2,
        },
        HeaderElement{
            Name: "e_sp",
            Offset: 16,
            Size:   2,
        },
        HeaderElement{
            Name: "e_csum",
            Offset: 18,
            Size:   2,
        },
        HeaderElement{
            Name: "e_ip",
            Offset: 20,
            Size:   2,
        },
        HeaderElement{
            Name: "e_cs",
            Offset: 22,
            Size:   2,
        },
        HeaderElement{
            Name: "e_lfarlc",
            Offset: 24,
            Size:   2,
        },
        HeaderElement{
            Name: "e_ovno",
            Offset: 26,
            Size:   2,
        },
        HeaderElement{
            Name: "e_res",
            Offset: 28,
            Size:   8,
        },
        HeaderElement{
            Name: "e_oemid",
            Offset: 36,
            Size:   2,
        },
        HeaderElement{
            Name: "e_oeminfo",
            Offset: 38,
            Size:   2,
        },
        HeaderElement{
            Name: "e_res2",
            Offset: 40,
            Size:   20,
        },
        HeaderElement{
            Name: "e_lfanew",
            Offset: 60,
            Size:   4,
        },
    }
}

func GetDOSStub(f *os.File) ImageHeader {
   h := ImageHeader{
        Title:      "IMAGE_DOS_STUB",
        Raw:        make([]byte, IMAGE_DOS_STUB_SIZE),
        Offset:     IMAGE_DOS_STUB_OFFSET,
        Size:       IMAGE_DOS_STUB_SIZE,
        Elements:   GetDOSStubElements(),
    }
    f.Seek(IMAGE_DOS_STUB_OFFSET, 0)
    f.Read(h.Raw)
    return h
}

func GetDOSStubElements() []HeaderElement {
    return []HeaderElement{
        HeaderElement{
            Name: "pre_dos_msg",
            Offset: 0,
            Size:   14,
        },
        HeaderElement{
            Name: "dos_msg",
            Offset: 14,
            Size:   39,
        },
        HeaderElement{
            Name: "post_dos_msg",
            Offset: 53,
            Size:   11,
        },
    }
}
