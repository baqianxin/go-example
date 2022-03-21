package base

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	s "strings"
	"testing"
	"unicode/utf8"
)

// 标准strings库的使用

var p = fmt.Println

// 特殊字符 - 中文字符 - rune读取
func runeChart() {
	const cn_zh = "aaaa"
	p("字符串长度:", utf8.RuneCountInString(cn_zh))
	for i, w := 0, 0; i < len(cn_zh); i += w {
		runeValue, width := utf8.DecodeRuneInString(cn_zh[i:])
		p(runeValue)
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func TestStrings(t *testing.T) {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
	runeChart()
}

// A URL represents a parsed URL (technically, a URI reference).
//
// The general form represented is:
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//
// URLs that do not start with a slash after the scheme are interpreted as:
//
//	scheme:opaque[?query][#fragment]
//
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
// A consequence is that it is impossible to tell which slashes in the Path were
// slashes in the raw URL and which were %2f. This distinction is rarely important,
// but when it is, the code should use RawPath, an optional field which only gets
// set if the default encoding is different from Path.
//
// URL's String method uses the EscapedPath method to obtain the path. See the
// EscapedPath method for more details.
// type URL struct {
// 	Scheme      string
// 	Opaque      string    // encoded opaque data
// 	User        *Userinfo // username and password information
// 	Host        string    // host or host:port
// 	Path        string    // path (relative paths may omit leading slash)
// 	RawPath     string    // encoded path hint (see EscapedPath method)
// 	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
// 	RawQuery    string    // encoded query values, without '?'
// 	Fragment    string    // fragment for references, without '#'
// 	RawFragment string    // encoded fragment hint (see EscapedFragment method)
// }

// Url 解析
func TestUrlParse(t *testing.T) {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawFragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}

// 字符串 hash 取值
func TestSha1(t *testing.T) {
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// 标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串。
// base64.StdEncoding.EncodeToString
/// ///  URLEncoding
func TestHash(t *testing.T) {
	data := "abc123!?$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
