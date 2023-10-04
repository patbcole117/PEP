package pe

import (
    "fmt"
    "os"
)

const (
    IMAGE_DOS_HEADER_SIZE   = 64
    IMAGE_DOS_HEADER_OFFSET = 0
    IMAGE_DOS_STUB_SIZE     = 64
    IMAGE_DOS_STUB_OFFSET   = 64

    e_magic_SIZE        = 2
    e_magic_OFFSET      = IMAGE_DOS_HEADER_OFFSET + 0
    e_cblp_SIZE         = 2
    e_cblp_OFFSET       = IMAGE_DOS_HEADER_OFFSET + 2
    e_cp_SIZE           = 2
    e_cp_OFFSET         = IMAGE_DOS_HEADER_OFFSET + 4
    e_crlc_SIZE         = 2
    e_crlc_OFFSET       = IMAGE_DOS_HEADER_OFFSET + 6
    e_cparhdr_SIZE      = 2
    e_cparhdr_OFFSET    = IMAGE_DOS_HEADER_OFFSET + 8
    e_minalloc_SIZE     = 2
    e_minalloc_OFFSET   = IMAGE_DOS_HEADER_OFFSET + 10
    e_maxalloc_SIZE     = 2
    e_maxalloc_OFFSET   = IMAGE_DOS_HEADER_OFFSET + 12
    e_ss_SIZE           = 2
    e_ss_OFFSET         = IMAGE_DOS_HEADER_OFFSET + 14
    e_sp_SIZE           = 2
    e_sp_OFFSET         = IMAGE_DOS_HEADER_OFFSET + 16
    e_csum_SIZE         = 2
    e_csum_OFFSET       = IMAGE_DOS_HEADER_OFFSET + 18
    e_ip_SIZE           = 2
    e_ip_OFFSET         = IMAGE_DOS_HEADER_OFFSET + 20
    e_cs_SIZE           = 2
    e_cs_OFFSET         = IMAGE_DOS_HEADER_OFFSET + 22
    e_lfarlc_SIZE       = 2
    e_lfarlc_OFFSET     = IMAGE_DOS_HEADER_OFFSET + 24
    e_ovno_SIZE         = 2
    e_ovno_OFFSET       = IMAGE_DOS_HEADER_OFFSET + 26
    e_res_SIZE          = 8
    e_res_LEN           = 4
    e_res_OFFSET        = IMAGE_DOS_HEADER_OFFSET + 28
    e_oemid_SIZE        = 2
    e_oemid_OFFSET      = IMAGE_DOS_HEADER_OFFSET + 36
    e_oeminfo_SIZE      = 2
    e_oeminfo_OFFSET    = IMAGE_DOS_HEADER_OFFSET + 38
    e_res2_SIZE         = 20
    e_res2_LEN          = 10
    e_res2_OFFSET       = IMAGE_DOS_HEADER_OFFSET + 40
    e_lfanew_SIZE       = 4
    e_lfanew_OFFSET     = IMAGE_DOS_HEADER_OFFSET + 60
)

func Get_IMAGE_DOS_HEADER(f *os.File) []byte {
    b := make([]byte, IMAGE_DOS_HEADER_SIZE)
    f.Seek(IMAGE_DOS_HEADER_OFFSET, 0)
    f.Read(b)
    return b
}
func Get_e_magic(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_magic_OFFSET:e_magic_OFFSET + e_magic_SIZE]
}
func Set_e_magic(f *os.File, b []byte) (int, error) {
    if len(b) > e_magic_SIZE {
        return 0, errWriteTooLarge
    }
    return f.WriteAt(b, e_magic_OFFSET)
}
func Get_e_cblp(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_cblp_OFFSET:e_cblp_OFFSET + e_cblp_SIZE]
}
func Get_e_cp(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_cp_OFFSET:e_cp_OFFSET + e_cp_SIZE]
}
func Get_e_crlc(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_crlc_OFFSET:e_crlc_OFFSET + e_crlc_SIZE]
}
func Get_e_cparhdr(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_cparhdr_OFFSET:e_cparhdr_OFFSET + e_cparhdr_SIZE]
}
func Get_e_minalloc(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_minalloc_OFFSET:e_minalloc_OFFSET + e_minalloc_SIZE]
}
func Get_e_maxalloc(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_maxalloc_OFFSET:e_maxalloc_OFFSET + e_maxalloc_SIZE]
}
func Get_e_sp(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_sp_OFFSET:e_sp_OFFSET + e_sp_SIZE]
}
func Get_e_ss(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_ss_OFFSET:e_ss_OFFSET + e_ss_SIZE]
}
func Get_e_csum(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_csum_OFFSET:e_csum_OFFSET + e_csum_SIZE]
}
func Get_e_ip(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_ip_OFFSET:e_ip_OFFSET + e_ip_SIZE]
}
func Get_e_cs(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_cs_OFFSET:e_cs_OFFSET + e_cs_SIZE]
}
func Get_e_lfarlc(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_lfarlc_OFFSET:e_lfarlc_OFFSET + e_lfarlc_SIZE]
}
func Get_e_ovno(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_ovno_OFFSET:e_ovno_OFFSET + e_ovno_SIZE]
}
func Get_e_res(f *os.File) [][]byte {
    b := Get_IMAGE_DOS_HEADER(f)[e_res_OFFSET:e_res_OFFSET + e_res_SIZE]
    ret := make([][]byte, e_res_LEN)
    s := e_res_SIZE / e_res_LEN
    for i := range ret {
        ret[i] = b[i*s:i*s+s]
    }
    return ret
}
func Get_e_oemid(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_oemid_OFFSET:e_oemid_OFFSET + e_oemid_SIZE]
}
func Get_e_oeminfo(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_oeminfo_OFFSET:e_oeminfo_OFFSET + e_oeminfo_SIZE]
}
func Get_e_res2(f *os.File) [][]byte {
    b := Get_IMAGE_DOS_HEADER(f)[e_res2_OFFSET:e_res2_OFFSET + e_res2_SIZE]
    ret := make([][]byte, e_res2_LEN)
    s := e_res2_SIZE / e_res2_LEN
    for i := range ret {
        ret[i] = b[i*s:i*s+s]
    }
    return ret
}
func Get_e_lfanew(f *os.File) []byte {
    return Get_IMAGE_DOS_HEADER(f)[e_lfanew_OFFSET:e_lfanew_OFFSET + e_lfanew_SIZE]
}

func Print_IMAGE_DOS_HEADER(f *os.File) {
    b := Get_IMAGE_DOS_HEADER(f)
    fmt.Printf("\n[+] IMAGE_DOS_HEADER\n")
    fmt.Println(PrintBytes(b))
}
