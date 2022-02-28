package languages

import (
	"golang.org/x/text/language"
	"sort"
	"strconv"
	"strings"
)

func Parse(str string) (Languages, error) {
	result := make(Languages, 0, 5)
	begin := 0
	for i := strings.Index(str, ","); i > 0; i = strings.Index(str[begin:], ",") {
		i += begin
		v := str[begin:i]
		if lang, err := parse(v); err != nil {
			return nil, err
		} else {
			result = append(result, lang)
		}
		begin = i + 1
	}
	if len(str) > begin {
		v := str[begin:]
		if lang, err := parse(v); err != nil {
			return nil, err
		} else {
			result = append(result, lang)
		}
	}
	return result, nil
}

func qualityValue(v string) int {
	if v == "" {
		return 10
	}
	v = strings.TrimPrefix(v, "q=")
	if n, err := strconv.ParseFloat(v, 64); err != nil {
		return 0
	} else if n <= 1 {
		return int(n * 10)
	}
	return 10
}

func parse(v string) (Language, error) {
	if j := strings.Index(v, ";"); j > 0 {
		if tag, err := language.Parse(v[0:j]); err != nil {
			return Language{}, err
		} else {
			return Language{
				Tag: tag,
				q:   qualityValue(v[j+1:]),
			}, nil
		}
	} else {
		if tag, err := language.Parse(v); err != nil {
			return Language{}, err
		} else {
			return Language{
				Tag: tag,
				q:   10,
			}, nil
		}
	}
}

type Language struct {
	language.Tag
	q int
}

type Languages []Language

func (lang Languages) Sort() Languages {
	sort.Slice(lang, func(i, j int) bool {
		return lang[i].q > lang[j].q
	})
	return lang
}
func (lang Languages) Filter(supports ...language.Tag) Languages {
	return lang.filter(false, supports...)
}

func (lang Languages) FilterS(supports ...language.Tag) Languages {
	return lang.filter(true, supports...)
}

func (lang Languages) filter(convert bool, supports ...language.Tag) Languages {
	result := make(Languages, 0, len(supports))
	for _, v := range lang {
		b1, s1, r1 := v.Raw()
		looseScript := false
		if s1.String() == "Zzzz" {
			looseScript = true
		}
		looseRegion := false
		if r1.String() == "ZZ" {
			looseRegion = true
		}
		for _, s := range supports {
			b2, s2, r2 := s.Raw()
			if b1.ISO3() == b2.ISO3() &&
				compareScript(looseScript, s1, s2) &&
				compareRegion(looseRegion, r1, r2) {
				result = appendList(result, v, s, convert)
				break
			}
		}
	}
	return result
}

func appendList(list Languages, lang Language, tag language.Tag, convert bool) Languages {
	if convert {
		for _, v := range list {
			if v.Tag == tag {
				return list
			}
		}
		return append(list, Language{
			Tag: tag,
			q:   lang.q,
		})
	}
	return append(list, lang)
}

func compareRegion(loose bool, r1, r2 language.Region) bool {
	if loose {
		return true
	}
	if r2.ISO3() == "ZZZ" {
		return true
	}
	return r1.ISO3() == r2.ISO3()
}

func compareScript(loose bool, s1, s2 language.Script) bool {
	if loose {
		return true
	}
	if s2.String() == "Zzzz" {
		return true
	}
	return s1.String() == s2.String()
}
