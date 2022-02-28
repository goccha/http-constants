package languages

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestParse(t *testing.T) {
	str := "ja,en-US;q=0.7,en;q=0.3,zh-Hant-TW;q=0.9,zh-Hans-TW;q=0.9"
	if v, err := Parse(str); err != nil {
		t.Error(err)
	} else {
		v = v.Sort()
		b, s, r := v[0].Raw()
		assert.Equal(t, "ja", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "ZZ", r.String())

		b, s, r = v[1].Raw()
		assert.Equal(t, "zh", b.String())
		assert.Equal(t, "Hant", s.String())
		assert.Equal(t, "TW", r.String())

		b, s, r = v[2].Raw()
		assert.Equal(t, "zh", b.String())
		assert.Equal(t, "Hans", s.String())
		assert.Equal(t, "TW", r.String())

		b, s, r = v[3].Raw()
		assert.Equal(t, "en", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "US", r.String())

		b, s, r = v[4].Raw()
		assert.Equal(t, "en", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "ZZ", r.String())
	}
}

func TestLanguages_Filter(t *testing.T) {
	str := "ja;q=0.8,en-US;q=0.7,en;q=0.3,zh-Hant-TW;q=0.9,zh-Hans-TW;q=0.9"
	if v, err := Parse(str); err != nil {
		t.Error(err)
	} else {
		filters := []language.Tag{
			language.Make("ja-JP"),
			language.Make("en-US"),
			language.Make("zh-CN"),
			language.Make("zh-TW"),
			language.Make("ko-KR"),
		}
		v = v.Filter(filters...).Sort()
		b, s, r := v[2].Raw()
		assert.Equal(t, "ja", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "ZZ", r.String())

		b, s, r = v[0].Raw()
		assert.Equal(t, "zh", b.String())
		assert.Equal(t, "Hant", s.String())
		assert.Equal(t, "TW", r.String())

		b, s, r = v[1].Raw()
		assert.Equal(t, "zh", b.String())
		assert.Equal(t, "Hans", s.String())
		assert.Equal(t, "TW", r.String())

		b, s, r = v[3].Raw()
		assert.Equal(t, "en", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "US", r.String())

		b, s, r = v[4].Raw()
		assert.Equal(t, "en", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "ZZ", r.String())

	}
}

func TestLanguages_FilterS(t *testing.T) {
	str := "ja;q=0.8,en-US;q=0.7,en;q=0.3,zh-Hant-TW;q=0.9,zh-Hans-TW;q=0.9"
	if v, err := Parse(str); err != nil {
		t.Error(err)
	} else {
		filters := []language.Tag{
			language.Make("ja-JP"),
			language.Make("en-US"),
			language.Make("zh-CN"),
			language.Make("zh-TW"),
			language.Make("ko-KR"),
		}
		v = v.FilterS(filters...).Sort()
		assert.Equal(t, 3, len(v))

		b, s, r := v[1].Raw()
		assert.Equal(t, "ja", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "JP", r.String())

		b, s, r = v[0].Raw()
		assert.Equal(t, "zh", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "TW", r.String())

		b, s, r = v[2].Raw()
		assert.Equal(t, "en", b.String())
		assert.Equal(t, "Zzzz", s.String())
		assert.Equal(t, "US", r.String())

	}
}

func BenchmarkParse(b *testing.B) {
	str := "ja,en-US;q=0.7,en;q=0.3,zh-Hant-TW;q=0.9,zh-Hans-TW;q=0.9"
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		_, _ = Parse(str)
	}
}

func BenchmarkFilter(b *testing.B) {
	filters := []language.Tag{
		language.Make("ja-JP"),
		language.Make("en-US"),
		language.Make("zh-CN"),
		language.Make("zh-TW"),
		language.Make("ko-KR"),
	}
	str := "ja;q=0.8,en-US;q=0.7,en;q=0.3,zh-Hant-TW;q=0.9,zh-Hans-TW;q=0.9"
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		v, _ := Parse(str)
		_ = v.Filter(filters...)
	}
}
