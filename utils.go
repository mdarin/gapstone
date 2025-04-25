package gapstone

import (
	"bytes"
	"fmt"
	"math/bits"
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

// PrettyPrintHex форматирует число в hex с автоматическим выбором полного представления.
// Примеры вывода
//
//	PrettyPrintHex(0xc)           -> 0xc
//	PrettyPrintHex(0x6)           -> 0x6
//	PrettyPrintHex(0xfffffffffffffdd0) -> 0xfffffffffffffdd0
//	PrettyPrintHex(255)           -> 0xff
//	PrettyPrintHex(int64(-1))     -> 0xffffffffffffffff (полное представление)
//	PrettyPrintHex("test")      -> "test" (неподдерживаемый тип)
func PrettyPrintHex(num any) string {
	var n uint64

	// Проверяем, является ли значение строкой (особый случай)
	if s, ok := num.(string); ok {
		return s // Возвращаем строку как есть
	}

	// Приводим число к uint64 (поддерживаем основные целочисленные типы)
	switch v := num.(type) {
	case int:
		n = uint64(v)
	case int8:
		n = uint64(v)
	case int16:
		n = uint64(v)
	case int32:
		n = uint64(v)
	case int64:
		n = uint64(v)
	case uint:
		n = uint64(v)
	case uint8:
		n = uint64(v)
	case uint16:
		n = uint64(v)
	case uint32:
		n = uint64(v)
	case uint64:
		n = v
	default:
		return fmt.Sprintf("%v", num) // Для неподдерживаемых типов выводим как есть
	}

	// Определяем, нужно ли полное представление (если есть ведущие нули в 64-битном представлении)
	leadingZeros := bits.LeadingZeros64(n)
	isLong := leadingZeros < 6 // Примерный порог: если меньше 6 ведущих нулей, считаем "длинным"

	// %#x — форматирует число в hex с префиксом 0x.
	// 0x%x — то же самое, но префикс добавлен вручную.
	// %016x — гарантирует вывод 16 символов, дополняя нулями слева (но в данном случае число уже занимает 16 hex-цифр).
	if isLong {
		return fmt.Sprintf("0x%016x", n) // Полное 64-битное представление
	}
	return fmt.Sprintf("0x%x", n) // Короткий формат
}
