package forwarded

import (
	"testing"
)

func TestParse(t *testing.T) {
	text := "for=192.168.0.1;host=1234567890.execute-api.ap-northeast-1.amazonaws.com;proto=https"
	forwarded := Parse(text)
	if len(forwarded.For) != 1 {
		t.Errorf("expected = 1 actual = %d", len(forwarded.For))
	}
	expected := "192.168.0.1"
	if expected != forwarded.For[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[0])
	}
	expected = "1234567890.execute-api.ap-northeast-1.amazonaws.com"
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = "https"
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	actual := forwarded.String()
	if text != actual {
		t.Errorf("expected = %s actual = %s", text, actual)
	}
}

func TestParse2(t *testing.T) {
	text := "for=\"_mdn\""
	forwarded := Parse(text)
	if len(forwarded.For) != 1 {
		t.Errorf("expected = 1 actual = %d", len(forwarded.For))
	}
	expected := "\"_mdn\""
	if expected != forwarded.For[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[0])
	}
	expected = ""
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = ""
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	actual := forwarded.String()
	if text != actual {
		t.Errorf("expected = %s actual = %s", text, actual)
	}
}

func TestParse3(t *testing.T) {
	text := "For=\"[2001:db8:cafe::17]:4711\""
	forwarded := Parse(text)
	if len(forwarded.For) != 1 {
		t.Errorf("expected = 1 actual = %d", len(forwarded.For))
	}
	expected := "\"[2001:db8:cafe::17]:4711\""
	if expected != forwarded.For[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[0])
	}
	expected = ""
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = ""
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	expected = "for=\"[2001:db8:cafe::17]:4711\""
	actual := forwarded.String()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
}

func TestParse4(t *testing.T) {
	text := "for=192.0.2.60; proto=http; by=203.0.113.43"
	forwarded := Parse(text)
	if len(forwarded.For) != 1 {
		t.Errorf("expected = 1 actual = %d", len(forwarded.For))
	}
	expected := "192.0.2.60"
	if expected != forwarded.For[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[0])
	}
	if len(forwarded.By) != 1 {
		t.Errorf("expected = 1 actual = %d", len(forwarded.For))
	}
	expected = "203.0.113.43"
	if expected != forwarded.By[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.By[0])
	}
	expected = ""
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = "http"
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	expected = "for=192.0.2.60;by=203.0.113.43;proto=http"
	actual := forwarded.String()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
}

func TestParse5(t *testing.T) {
	text := "for=192.0.2.43, for=198.51.100.17 "
	forwarded := Parse(text)
	if len(forwarded.For) != 2 {
		t.Errorf("expected = 2 actual = %d", len(forwarded.For))
		return
	}
	expected := "192.0.2.43"
	if expected != forwarded.For[0] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[0])
	}
	expected = "198.51.100.17"
	if expected != forwarded.For[1] {
		t.Errorf("expected = %s actual = %s", expected, forwarded.For[1])
	}
	expected = "192.0.2.43"
	actual := forwarded.ClientIP()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
	expected = ""
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = ""
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	expected = "for=192.0.2.43,for=198.51.100.17"
	actual = forwarded.String()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
}

func TestParseEmpty(t *testing.T) {
	text := ""
	forwarded := Parse(text)
	if len(forwarded.For) != 0 {
		t.Errorf("expected = 0 actual = %d", len(forwarded.For))
		return
	}
	expected := ""
	actual := forwarded.ClientIP()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
	expected = ""
	if expected != forwarded.Host {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Host)
	}
	expected = ""
	if expected != forwarded.Proto {
		t.Errorf("expected = %s actual = %s", expected, forwarded.Proto)
	}
	expected = ""
	actual = forwarded.String()
	if expected != actual {
		t.Errorf("expected = %s actual = %s", expected, actual)
	}
}

func BenchmarkParse(b *testing.B) {
	text := "for=192.168.0.1;host=1234567890.execute-api.ap-northeast-1.amazonaws.com;proto=https"
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		_ = Parse(text)
	}
}
