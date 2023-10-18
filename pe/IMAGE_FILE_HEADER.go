package pe

import (
    "encoding/binary"
    "os"
    "time"
)

const (
    IMAGE_FILE_HEADER_SIZE    = 20
    PE_SIGNATURE_SIZE         = 4

    Machine_SIZE                    = 2
    Machine_OFFSET                  = 0
    NumberOfSections_SIZE           = 2
    NumberOfSections_OFFSET         = 2
    TimeDateStamp_SIZE              = 4
    TimeDateStamp_OFFSET            = 4
    PointerToSymbolTable_SIZE       = 4
    PointerToSymbolTable_OFFSET     = 8
    NumberOfSymbols_SIZE            = 4
    NumberOfSymbols_OFFSET          = 12
    SizeOfOptionalHeader_SIZE       = 2
    SizeOfOptionalHeader_OFFSET     = 16
    Characteristics_SIZE            = 2
    Characteristics_OFFSET          = 18

    IMAGE_FILE_MACHINE_UNKNOWN      uint16 = 0
    IMAGE_FILE_MACHINE_ALPHA        uint16 = 388
    IMAGE_FILE_MACHINE_ALPHA64      uint16 = 644
    IMAGE_FILE_MACHINE_AM33         uint16 = 467
    IMAGE_FILE_MACHINE_AMD64        uint16 = 34404
    IMAGE_FILE_MACHINE_ARM          uint16 = 448
    IMAGE_FILE_MACHINE_ARM64        uint16 = 43620
    IMAGE_FILE_MACHINE_ARMNT        uint16 = 452
    IMAGE_FILE_MACHINE_AXP64        uint16 = 644
    IMAGE_FILE_MACHINE_EBC          uint16 = 3772
    IMAGE_FILE_MACHINE_I386         uint16 = 332
    IMAGE_FILE_MACHINE_IA64         uint16 = 512
    IMAGE_FILE_MACHINE_LOONGARCH32  uint16 = 25138
    IMAGE_FILE_MACHINE_LOONGARCH64  uint16 = 25188
    IMAGE_FILE_MACHINE_M32R         uint16 = 36929
    IMAGE_FILE_MACHINE_MIPS16       uint16 = 614
    IMAGE_FILE_MACHINE_MIPSFPU      uint16 = 870
    IMAGE_FILE_MACHINE_MIPSFPU16    uint16 = 1126
    IMAGE_FILE_MACHINE_POWERPC      uint16 = 496
    IMAGE_FILE_MACHINE_POWERPCFP    uint16 = 497
    IMAGE_FILE_MACHINE_R4000        uint16 = 358
    IMAGE_FILE_MACHINE_RISCV32      uint16 = 20530
    IMAGE_FILE_MACHINE_RISCV64      uint16 = 20580
    IMAGE_FILE_MACHINE_RISCV128     uint16 = 20776
    IMAGE_FILE_MACHINE_SH3          uint16 = 418
    IMAGE_FILE_MACHINE_SH3DSP       uint16 = 419
    IMAGE_FILE_MACHINE_SH4          uint16 = 422
    IMAGE_FILE_MACHINE_SH5          uint16 = 424
    IMAGE_FILE_MACHINE_THUMB        uint16 = 450
    IMAGE_FILE_MACHINE_WCEMIPSV2    uint16 = 361

    IMAGE_FILE_RELOCS_STRIPPED          uint16 = 1
    IMAGE_FILE_EXECUTABLE_IMAGE         uint16 = 2
    IMAGE_FILE_NUMS_STRIPPED            uint16 = 4
    IMAGE_FILE_LOCAL_SYMS_STRIPPED      uint16 = 8
    IMAGE_FILE_AGGRESSIVE_WS_TRIM       uint16 = 16
    IMAGE_FILE_LARGE_ADDRESS_AWARE      uint16 = 32
    IMAGE_FILE_BYTES_REVERSED_LO        uint16 = 128
    IMAGE_FILE_32BIT_MACHINE            uint16 = 256
    IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP  uint16 = 1024
    IMAGE_FILE_NET_RUN_FROM_SWAP        uint16 = 2048
    IMAGE_FILE_SYSTEM                   uint16 = 4096
    IMAGE_FILE_FILE_DLL                 uint16 = 8192
    IMAGE_FILE_UP_SYSTEM_ONLY           uint16 = 16384
    IMAGE_FILE_BYTES_REVERSED_HI        uint16 = 32768
)

func Get_PE_SIGNATURE(f *os.File) []byte {
    b := make([]byte, PE_SIGNATURE_SIZE)
    f.Seek(DWORDToI64(Get_e_lfanew(f)), 0)
    f.Read(b)
    return b
}

func Get_IMAGE_FILE_HEADER(f *os.File) []byte {
    b := make([]byte, IMAGE_FILE_HEADER_SIZE)
    f.Seek(DWORDToI64(Get_e_lfanew(f)) + PE_SIGNATURE_SIZE, 0)
    f.Read(b)
    return b
}

func Get_Machine(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[Machine_OFFSET:Machine_OFFSET + Machine_SIZE]
}

func Get_NumberOfSections(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[NumberOfSections_OFFSET:NumberOfSections_OFFSET + NumberOfSections_SIZE]
}

func Get_TimeDateStamp(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[TimeDateStamp_OFFSET:TimeDateStamp_OFFSET + TimeDateStamp_SIZE]
}

func Get_PointerToSymbolTable(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[PointerToSymbolTable_OFFSET:PointerToSymbolTable_OFFSET + PointerToSymbolTable_SIZE]
}

func Get_NumberOfSymbols(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[NumberOfSymbols_OFFSET:NumberOfSymbols_OFFSET + NumberOfSymbols_SIZE]
}

func Get_SizeOfOptionalHeader(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[SizeOfOptionalHeader_OFFSET:SizeOfOptionalHeader_OFFSET + SizeOfOptionalHeader_SIZE]
}

func Get_Characteristics(f *os.File) []byte {
    return Get_IMAGE_FILE_HEADER(f)[Characteristics_OFFSET:Characteristics_OFFSET + Characteristics_SIZE]
}

func Decode_Machine(b []byte) (string, error) {
    if len(b) != Machine_SIZE {
        return "",  errInvalidParam
    }
    m := binary.LittleEndian.Uint16(b)
    switch m {
    case IMAGE_FILE_MACHINE_UNKNOWN:
        return "IMAGE_FILE_MACHINE_UNKNOWN", nil
    case IMAGE_FILE_MACHINE_ALPHA:
        return "IMAGE_FILE_MACHINE_ALPHA", nil
    case IMAGE_FILE_MACHINE_ALPHA64:
        return "IMAGE_FILE_MACHINE_ALPHA64/AXP64", nil
    case IMAGE_FILE_MACHINE_AM33:
        return "IMAGE_FILE_MACHINE_AM33", nil
    case IMAGE_FILE_MACHINE_AMD64:
        return "IMAGE_FILE_MACHINE_AMD64", nil
    case IMAGE_FILE_MACHINE_ARM:
        return "IMAGE_FILE_MACHINE_ARM", nil
    case IMAGE_FILE_MACHINE_ARM64:
        return "IMAGE_FILE_MACHINE_ARM64", nil
    case IMAGE_FILE_MACHINE_ARMNT:
        return "IMAGE_FILE_MACHINE_ARMNT", nil
    case IMAGE_FILE_MACHINE_EBC:
        return "IMAGE_FILE_MACHINE_EBC", nil
    case IMAGE_FILE_MACHINE_I386:
        return "IMAGE_FILE_MACHINE_I386", nil
    case IMAGE_FILE_MACHINE_IA64:
        return "IMAGE_FILE_MACHINE_IA64", nil
    case IMAGE_FILE_MACHINE_LOONGARCH32:
        return "IMAGE_FILE_MACHINE_LOONGARCH32", nil
    case IMAGE_FILE_MACHINE_LOONGARCH64:
        return "IMAGE_FILE_MACHINE_LOONGARCH64", nil
    case IMAGE_FILE_MACHINE_M32R:
        return "IMAGE_FILE_MACHINE_M32R", nil
    case IMAGE_FILE_MACHINE_MIPS16:
        return "IMAGE_FILE_MACHINE_MIPS16", nil
    case IMAGE_FILE_MACHINE_MIPSFPU:
        return "IMAGE_FILE_MACHINE_MIPSFPU", nil
    case IMAGE_FILE_MACHINE_MIPSFPU16:
        return "IMAGE_FILE_MACHINE_MIPSFPU16", nil
    case IMAGE_FILE_MACHINE_POWERPC:
        return "IMAGE_FILE_MACHINE_POWERPC", nil
    case IMAGE_FILE_MACHINE_POWERPCFP:
        return "IMAGE_FILE_MACHINE_POWERPCFP", nil
    case IMAGE_FILE_MACHINE_R4000:
        return "IMAGE_FILE_MACHINE_R4000", nil
    case IMAGE_FILE_MACHINE_RISCV32:
        return "IMAGE_FILE_MACHINE_RISCV32", nil
    case IMAGE_FILE_MACHINE_RISCV64:
        return "IMAGE_FILE_MACHINE_RISCV64", nil
    case IMAGE_FILE_MACHINE_RISCV128:
        return "IMAGE_FILE_MACHINE_RISCV128", nil
    case IMAGE_FILE_MACHINE_SH3:
        return "IMAGE_FILE_MACHINE_SH3", nil
    case IMAGE_FILE_MACHINE_SH3DSP:
        return "IMAGE_FILE_MACHINE_SH3DSP", nil
    case IMAGE_FILE_MACHINE_SH4:
        return "IMAGE_FILE_MACHINE_SH4", nil
    case IMAGE_FILE_MACHINE_SH5:
        return "IMAGE_FILE_MACHINE_SH5", nil
    case IMAGE_FILE_MACHINE_THUMB:
        return "IMAGE_FILE_MACHINE_THUMB", nil
    case IMAGE_FILE_MACHINE_WCEMIPSV2:
        return "IMAGE_FILE_MACHINE_WCEMIPSV2", nil
    default:
        return "", errInvalidMachine
    }
}

func Decode_TimeDateStamp(b []byte) (string, error) {
    if len(b) != TimeDateStamp_SIZE {
        return "",  errInvalidParam
    }

    temp := int64(binary.LittleEndian.Uint32(b))
    l, err := time.LoadLocation("UTC")
    if err != nil {
        return "", err
    }
    t := time.Unix(temp, 0).In(l)

    return t.Format(time.RFC1123Z), nil
}

func Decode_Characteristics(b []byte) ([]string, error) {
    var ret []string
    if len(b) != Machine_SIZE {
        return nil,  errInvalidParam
    }
    c := binary.LittleEndian.Uint16(b)
    if c & IMAGE_FILE_RELOCS_STRIPPED == IMAGE_FILE_RELOCS_STRIPPED {
        ret = append(ret, "IMAGE_FILE_RELOCS_STRIPPED")
    }
    if c & IMAGE_FILE_EXECUTABLE_IMAGE == IMAGE_FILE_EXECUTABLE_IMAGE {
        ret = append(ret, "IMAGE_FILE_EXECUTABLE_IMAGE")
    }
    if c & IMAGE_FILE_NUMS_STRIPPED == IMAGE_FILE_NUMS_STRIPPED {
        ret = append(ret, "IMAGE_FILE_NUMS_STRIPPED")
    }
    if c & IMAGE_FILE_LOCAL_SYMS_STRIPPED == IMAGE_FILE_LOCAL_SYMS_STRIPPED {
        ret = append(ret, "IMAGE_FILE_LOCAL_SYMS_STRIPPED")
    }
    if c & IMAGE_FILE_AGGRESSIVE_WS_TRIM == IMAGE_FILE_AGGRESSIVE_WS_TRIM {
        ret = append(ret, "IMAGE_FILE_AGGRESSIVE_WS_TRIM")
    }
    if c & IMAGE_FILE_LARGE_ADDRESS_AWARE == IMAGE_FILE_LARGE_ADDRESS_AWARE {
        ret = append(ret, "IMAGE_FILE_LARGE_ADDRESS_AWARE")
    }
    if c & IMAGE_FILE_BYTES_REVERSED_LO == IMAGE_FILE_BYTES_REVERSED_LO {
        ret = append(ret, "IMAGE_FILE_BYTES_REVERSED_LO")
    }
    if c & IMAGE_FILE_32BIT_MACHINE == IMAGE_FILE_32BIT_MACHINE {
        ret = append(ret, "IMAGE_FILE_32BIT_MACHINE")
    }
    if c & IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP == IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP {
        ret = append(ret, "IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP")
    }
    if c & IMAGE_FILE_NET_RUN_FROM_SWAP == IMAGE_FILE_NET_RUN_FROM_SWAP {
        ret = append(ret, "IMAGE_FILE_NET_RUN_FROM_SWAP")
    }
    if c & IMAGE_FILE_SYSTEM == IMAGE_FILE_SYSTEM {
        ret = append(ret, "IMAGE_FILE_SYSTEM")
    }
    if c & IMAGE_FILE_FILE_DLL == IMAGE_FILE_FILE_DLL {
        ret = append(ret, "IMAGE_FILE_FILE_DLL")
    }
    if c & IMAGE_FILE_UP_SYSTEM_ONLY == IMAGE_FILE_UP_SYSTEM_ONLY {
        ret = append(ret, "IMAGE_FILE_UP_SYSTEM_ONLY")
    }
    if c & IMAGE_FILE_BYTES_REVERSED_HI == IMAGE_FILE_BYTES_REVERSED_HI {
        ret = append(ret, "IMAGE_FILE_BYTES_REVERSED_HI")
    }
    return ret, nil
}