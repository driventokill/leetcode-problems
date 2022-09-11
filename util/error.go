package util

import "strings"

type Partial interface {
	Code() string
	Msg() string
}

type P struct {
	code, msg string
}

func (p P) Code() string {
	return p.code
}

func (p P) Msg() string {
	return p.msg
}

type BizError struct {
	Parts   []Partial
	isFinal bool // 校验当前是否已经是完整的错误码
	Partial
}

func (err BizError) Code() string {
	if !err.isFinal {
		panic("error is not finalized")
	}
	return join(err.Parts, func(p Partial) string {
		return p.Code()
	}, "")
}

func (err BizError) Msg() string {
	if !err.isFinal {
		panic("error is not finalized")
	}
	return join(err.Parts, func(p Partial) string {
		return p.Msg()
	}, "|")
}

func join(parts []Partial, f func(p Partial) string, sep string) string {
	if len(parts) == 0 {
		return ""
	}

	var elem []string
	for _, v := range parts {
		elem = append(elem, f(v))
	}
	return strings.Join(elem, sep)
}

func NewPartial(p ...Partial) *BizError {
	return &BizError{
		Parts:   p,
		isFinal: false,
	}
}

func New(err *BizError, p Partial) *BizError {
	var isFinal bool
	if len(err.Parts) == 4 {
		isFinal = true
	}
	return &BizError{
		Parts:   append(err.Parts, p),
		isFinal: isFinal,
	}
}

// --- 猜喜 ----
var FavoritePartialError = NewPartial(P{"Warn_1", "Warn"}, P{"01", "猜喜"})
var FavoriteInsertPartialError = New(FavoritePartialError, P{"01", "强插"})

var ForceGetConfigErr = New(FavoriteInsertPartialError, P{"0001", "强插配置查询失败"})
