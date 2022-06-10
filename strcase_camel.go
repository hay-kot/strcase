package strcase

import (
	"strings"
)

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool, acronyms AcronymsConf) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if acronyms != nil {
		if a, ok := acronyms[s]; ok {
			s = a
		}
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	prevWasNum := false
	sBytes := []byte(s)
	sLen := len(sBytes)
	for i, v := range sBytes {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
			prevWasNum = false
		} else if isNum(v) {
			n.WriteByte(v)
			prevWasNum = true
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
			if capNext && prevWasNum && (sLen >= i+1 && isNum([]byte(s)[i+1])) {
				n.WriteByte('_')
				prevWasNum = false
			}
		}
	}
	return n.String()
}

// ToPascal converts a string to CamelCase
func ToPascal(s string, acronyms AcronymsConf) string {
	return toCamelInitCase(s, true, acronyms)
}

// ToCamel converts a string to lowerCamelCase
func ToCamel(s string, acronyms AcronymsConf) string {
	return toCamelInitCase(s, false, acronyms)
}

func isNum(v byte) bool {
	return v >= '0' && v <= '9'
}
