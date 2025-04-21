package gapstone

import (
	"bytes"
	"unicode"
)

// NormalizeString удаляет незначимые символы из строки.
func NormalizeString(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		if isSemanticallySignificant(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// NormalizeBytes удаляет незначимые символы из []byte.
func NormalizeBytes(b []byte) []byte {
	var buf bytes.Buffer
	for _, r := range bytes.Runes(b) {
		if isSemanticallySignificant(r) {
			buf.WriteRune(r)
		}
	}
	return buf.Bytes()
}

// Для более строгого сравнения дополнительные фильтры.
func isSemanticallySignificant(r rune) bool {
	return !unicode.IsSpace(r) &&
		!unicode.IsControl(r) &&
		r != '\uFEFF' && // BOM
		r != '\u200B' // Zero-width space
}

// CompareNormalized надежно сравнивает строки/байты после нормализации.
//
// Usage examples:
// CompareNormalized("hello\nworld", "hello world"))  // true
// CompareNormalized([]byte("go\tlang"), "golang"))   // true
// CompareNormalized(bytes.NewBufferString(" data "), "data") // true
// CompareNormalized("one", "two"))                   // false
func CompareNormalized(a, b any) bool {
	switch v1 := a.(type) {
	case string:
		if bv2, ok := b.(*bytes.Buffer); ok {
			return NormalizeString(v1) == NormalizeString(bv2.String())
		}
		v2, ok := b.(string)
		if !ok {
			if v3, ok := b.([]byte); ok {
				v2 = string(v3)
			} else {
				return false
			}
		}
		return NormalizeString(v1) == NormalizeString(v2)

	case []byte:
		if bv2, ok := b.(*bytes.Buffer); ok {
			return bytes.Equal(NormalizeBytes(v1), NormalizeBytes(bv2.Bytes()))
		}
		v2, ok := b.([]byte)
		if !ok {
			if v3, ok := b.(string); ok {
				v2 = []byte(v3)
			} else {
				return false
			}
		}
		return bytes.Equal(NormalizeBytes(v1), NormalizeBytes(v2))

	case *bytes.Buffer:
		if bv2, ok := b.(*bytes.Buffer); ok {
			return bytes.Equal(NormalizeBytes(v1.Bytes()), NormalizeBytes(bv2.Bytes()))
		}
		if sv2, ok := b.(string); ok {
			return bytes.Equal(NormalizeBytes(v1.Bytes()), NormalizeBytes([]byte(sv2)))
		}
		if bv2, ok := b.([]byte); ok {
			return bytes.Equal(NormalizeBytes(v1.Bytes()), NormalizeBytes(bv2))
		}
		return false

	default:
		return false
	}
}
