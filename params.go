package aliyun_email

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type Params map[string]interface{}

func (p Params) PutString(key, val string) {
	p[key] = val
}
func (p Params) PutBoolean(key string, val bool) {
	p[key] = val
}
func (p Params) PutInt(key string, val int) {
	p[key] = val
}
func (p Params) Get(key string) interface{} {
	return p[key]
}
func (p Params) Keys() []string {
	keys := make([]string, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}
	return keys
}
func (p Params) Remove(key string) {
	delete(p, key)
}

func (p Params) SortedKeys() []string {
	keys := p.Keys()
	sort.Strings(keys)
	return keys
}

func (p Params) Sign(method, accessKeySecret string) {
	p.Remove("Signature")
	keys := p.SortedKeys()
	buf := bytes.NewBufferString("")
	for index, key := range keys {
		buf.WriteString(key)
		buf.WriteString("=")
		switch v := p.Get(key).(type) {
		case string:
			buf.WriteString(p.UrlEncode(v))
		case int:
			buf.WriteString(p.UrlEncode(strconv.Itoa(v)))
		case bool:
			buf.WriteString(p.UrlEncode(strconv.FormatBool(v)))
		}
		if index != len(keys)-1 {
			buf.WriteString("&")
		}
	}
	strToSign := method + "&" + p.UrlEncode("/") + "&" + p.UrlEncode(buf.String())
	mac := hmac.New(sha1.New, []byte(accessKeySecret))
	mac.Write([]byte(strToSign))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	p.PutString("Signature", sign)
}

func (p Params) UrlEncode(val string) string {
	val = url.QueryEscape(val)
	val = strings.Replace(val, "+", "%20", -1)
	val = strings.Replace(val, "*", "%2A", -1)
	val = strings.Replace(val, "%7E", "~", -1)
	return val
}

func (p Params) ToUrlValues() url.Values {
	values := url.Values{}
	for key := range p {
		switch v := p.Get(key).(type) {
		case string:
			values[key] = []string{v}
		case int:
			values[key] = []string{strconv.Itoa(v)}
		case bool:
			values[key] = []string{strconv.FormatBool(v)}
		}
	}
	return values
}
