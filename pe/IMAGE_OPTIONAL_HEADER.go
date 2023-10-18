package pe

import (
    "os"
)

func Get_IMAGE_OPTIONAL_HEADER(f *os.File) []byte {
    b := make([]byte, WORDToI64(Get_SizeOfOptionalHeader(f)))
    f.Seek(DWORDToI64(Get_e_lfanew(f)) + PE_SIGNATURE_SIZE + IMAGE_FILE_HEADER_SIZE, 0)
    f.Read(b)
    return b
}