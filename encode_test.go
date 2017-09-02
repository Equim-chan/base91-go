package base91

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

var (
	encodeSpec = map[string]string{
		"1":                                        "xA",
		"Hello World!":                             ">OwJh>Io0Tv!8PE",
		"「さやかちゃん、大好きだ！(*^ω^*)」":                    "tC??dBPUBX|xqnB@VEC%qCXQ{+WB|9~5]PIlN+\";B`%tx34t0c.;[Gf6W0WBUG",
		"\x3e\xeb\xa0\x34\x10\x01\x9d\x96\x5e":     "5fNOkLP/rav",
		"\xf2\x8e\x88\x31\x1a\xf0\x68\xce\x7a\x3f": "EquimSayaka~A",
	}
)

func TestEncode(t *testing.T) {
	for k, v := range encodeSpec {
		if actual := EncodeToString([]byte(k)); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func TestEncoder(t *testing.T) {
	for k, v := range encodeSpec {
		buf := new(bytes.Buffer)
		e := NewEncoder(buf)
		if _, err := e.Write([]byte(k)); err != nil {
			t.Fatal(err)
		}
		if err := e.Close(); err != nil {
			t.Fatal(err)
		}
		if actual := string(buf.Bytes()); actual != v {
			t.Fatalf("expected `%s`, got `%s`", v, actual)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s := make([]byte, 1024*1024)
		if _, err := io.ReadFull(rand.Reader, s); err != nil {
			b.Fatal(err)
		}

		b.StartTimer()

		EncodeToString(s)
	}
}

func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		buf := new(bytes.Buffer)
		buf.Grow(1024 * 1024)
		e := NewEncoder(buf)
		defer e.Close()

		b.StartTimer()

		if _, err := io.CopyN(e, rand.Reader, 1024*1024); err != nil {
			b.Fatal(err)
		}
	}
}
