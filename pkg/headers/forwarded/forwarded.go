package forwarded

import "strings"

type Forwarded struct {
	For   []string
	By    []string
	Host  string
	Proto string
}

func New() *Forwarded {
	return &Forwarded{
		For:   []string{},
		By:    []string{},
		Host:  "",
		Proto: "",
	}
}
func (fw *Forwarded) ClientIP() string {
	if fw.For != nil && len(fw.For) > 0 {
		return fw.For[0]
	}
	return ""
}
func (fw *Forwarded) addFor(v string) {
	if fw.For == nil {
		fw.For = make([]string, 0, 1)
	}
	fw.For = append(fw.For, v)
}
func (fw *Forwarded) addBy(v string) {
	if fw.By == nil {
		fw.By = make([]string, 0, 1)
	}
	fw.By = append(fw.By, v)
}
func (fw *Forwarded) String() string {
	builder, ok := appendString(&strings.Builder{}, false, "for", fw.For...)
	builder, ok = appendString(builder, ok, "by", fw.By...)
	builder, ok = appendString(builder, ok, "host", fw.Host)
	builder, ok = appendString(builder, ok, "proto", fw.Proto)
	return builder.String()
}

func appendString(buf *strings.Builder, ok bool, key string, values ...string) (*strings.Builder, bool) {
	found := false
	for i, v := range values {
		if v == "" {
			continue
		}
		if i > 0 {
			buf.WriteRune(',')
		} else {
			if ok {
				buf.WriteRune(';')
			}
			found = true
		}
		buf.WriteString(key)
		buf.WriteRune('=')
		buf.WriteString(v)
	}
	if !found {
		found = ok
	}
	return buf, found
}

func setValue(forwarded *Forwarded, key, value string) {
	if key != "" {
		key = strings.ToLower(key)
		value = strings.Trim(value, " ")
		switch key {
		case "for":
			forwarded.addFor(value)
		case "by":
			forwarded.addBy(value)
		case "host":
			forwarded.Host = value
		case "proto":
			forwarded.Proto = value
		}
	}
}

func Parse(text string) *Forwarded {
	result := New()
	head := -1
	var key string
	for i, r := range text {
		switch r {
		case ';', ',':
			setValue(result, key, text[head:i])
			key = ""
			head = -1
		case ' ':
			// ignore
		case '=':
			key = text[head:i]
			head = i + 1
		default:
			if head < 0 {
				head = i
			}
		}
	}
	if key != "" && head >= 0 {
		setValue(result, key, text[head:])
	}
	return result
}
