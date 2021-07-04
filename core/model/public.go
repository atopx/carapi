package model

const PublicConstFile = `package public

const (
	DayTimeFormat    = "2006-01-02"
	SecondTimeFormat = "2006-01-02 15:04:05"
)
`
const PublicHandleFile = `package public

import (
	"app/config"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Cfg *config.Config
)
`

const PublicUtilsFile = `package public

import (
	"crypto/sha1"
	"encoding/hex"
	"reflect"
	"strings"
	"unsafe"
)

// StrToBytes 字符串转字节
func StrToBytes(s string) []byte {
	var sp = (*reflect.StringHeader)(unsafe.Pointer(&s))
	var bp = reflect.SliceHeader{Data: sp.Data, Len: sp.Len, Cap: sp.Len}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

// BytesToStr 字节转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Hash 生成sha1
func Hash(str string) string {
	var hash = sha1.New()
	hash.Write(StrToBytes(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// HashValid sha1校验
func HashValid(str, hex string) bool {
	return Hash(str) == hex
}

// LikeQueryJoin 数据库like匹配
func LikeQueryJoin(value string) string {
	return strings.Join([]string{"%", value, "%"}, "")
}
`
