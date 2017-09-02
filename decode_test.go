package base91

import (
	"bytes"
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"
)

var (
	decodeSpec = map[string]string{
		"xA":                                                              "1",
		">OwJh>Io0Tv!8PE":                                                 "Hello World!",
		"tC??dBPUBX|xqnB@VEC%qCXQ{+WB|9~5]PIlN+\";B`%tx34t0c.;[Gf6W0WBUG": "「さやかちゃん、大好きだ！(*^ω^*)」",
		"5fNOkLP/rav":   "\x3e\xeb\xa0\x34\x10\x01\x9d\x96\x5e",
		"EquimSayaka~A": "\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f",
	}
)

func TestDecode(t *testing.T) {
	for k, v := range decodeSpec {
		if actual := string(DecodeString(k)); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func TestDecoder(t *testing.T) {
	for k, v := range decodeSpec {
		d := NewDecoder(bytes.NewReader([]byte(k)))
		actual, err := ioutil.ReadAll(d)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s := make([]byte, 1024*1024)
		if _, err := io.ReadFull(rand.Reader, s); err != nil {
			b.Fatal(err)
		}
		encoded := EncodeToString(s)

		b.StartTimer()

		DecodeString(encoded)
	}
}

func BenchmarkDecoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		buf := new(bytes.Buffer)
		buf.Grow(1024 * 1024)
		d := NewDecoder(rand.Reader)

		b.StartTimer()

		if _, err := io.CopyN(buf, d, 1024*1024); err != nil {
			b.Fatal(err)
		}
	}
}
