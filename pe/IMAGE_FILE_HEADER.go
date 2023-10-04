package pe

import (
    "encoding/binary"
    "os"
)

const (
    IMAGE_FILE_HEADER_SIZE = 20
)

func GetFileHeader(f *os.File) ImageHeader {
    h := ImageHeader{
        Title:      "IMAGE_FILE_HEADER",
        Raw:        make([]byte, IMAGE_FILE_HEADER_SIZE),
        Size:       IMAGE_FILE_HEADER_SIZE,
        Elements:   GetFileHeaderelements()
    }
    b := GetDOSHeaderValue(f, "e_lfanew")
    h.Offset = binary.LittleEndian.Uint16
    f.Seek(h.Offset, 0)
    f.Read(h.Raw)
    return h
}

func GetFileHeaderElements() []Headerelements {
    return []HeaderElements{
        HeaderElement{
            Name: "Machine",
            Offset: 0,
            Size:   2,
        },
        HeaderElement{
            Name: "NumberOfSections",
            Offset: 2,
            Size:   2,
        },
        HeaderElement{
            Name: "TimeDateStamp",
            Offset: 4,
            Size:   4,
        },
        HeaderElement{
            Name: "PointerToSymbolTable",
            Offset: 8,
            Size:   4,
        },
        HeaderElement{
            Name: "NumberOfSymbols",
            Offset: 12,
            Size:   4,
        },
        HeaderElement{
            Name: "SizeOfOptionalHeader",
            Offset: 16,
            Size:   2,
        },
        HeaderElement{
            Name: "Characteristics",
            Offset: 18,
            Size:   2,
        },
    }
}
