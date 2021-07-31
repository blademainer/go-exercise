package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
	"time"
)

func main() {
	// TOTP基于HOTP，附带的时间戳用于替代递增计数器。
	//
	// 通过定义一个纪元（epoch, T0）的起点并以时间步骤（time step, TS）为单位进行计数，当前时间戳被转换成一个整形数值的时间计数器（time-counter, TC）。举一个例子：
	// TC = 最低值((unixtime(当前时间)−unixtime(T0))/TS)，
	// TOTP = HOTP(密钥, TC),
	// 	TOTP-值 = TOTP mod 10d，d 是所需的一次性密码的位数。
	key := []byte("hello totp")
	token, err := NewToken(key, 6)
	if err != nil {
		panic(err.Error())
	}
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	i := 0
	for range t.C {
		i++
		fmt.Println(token.Gen())
		// fmt.Println(token.GenByTime(t0))
	}
}

const step = 30

var t0 time.Time

func init() {
	var err error
	t0, err = time.Parse(time.RFC3339, "2021-01-02T00:00:00+08:00")
	if err != nil {
		panic(err.Error())
	}
}

type Token struct {
	h   func() hash.Hash
	len int
}

func NewToken(key []byte, len int) (*Token, error) {
	t := &Token{
		h: func() hash.Hash {
			return hmac.New(sha256.New, key)
		},
		len: len,
	}

	return t, nil
}

func (t *Token) Gen() string {
	return t.GenByTime(time.Now())
}

func (t *Token) GenByTime(now time.Time) string {
	since := now.Sub(t0)
	i := int(since.Seconds())
	f := flow(i, step)
	rs := t.sum(f)
	l := len(rs)
	rs = rs[l-t.len : l]
	return rs
}

func (t *Token) sum(f int64) string {
	h := t.h()
	err := binary.Write(h, binary.LittleEndian, f)
	if err != nil {
		panic(err)
	}
	sum := h.Sum(nil)
	rs := hex.EncodeToString(sum)
	return strings.ToUpper(rs)
}

func flow(i, step int) int64 {
	return int64(i / step)
}
