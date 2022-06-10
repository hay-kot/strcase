package strcase

import (
	"testing"
)

func toPascal(tb testing.TB) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"AnyKind.of_string13_37", "AnyKindOfString13_37"},
		{"AnyKind.of_string13_37 test", "AnyKindOfString13_37Test"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"ID", "Id"},
	}
	for _, i := range cases {
		in, out := i[0], i[1]
		result := ToPascal(in, New())
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func TestToPascal(t *testing.T) {
	toPascal(t)
}

func BenchmarkToPascal(b *testing.B) {
	benchmarkCamelTest(b, toPascal)
}

func toCamel(tb testing.TB) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
		{"AnyKind.of-string", "anyKindOfString"},
		{"ID", "id"},
		{"AnyKind.of_string13_37", "anyKindOfString13_37"},
		{"AnyKind.of_string13_37test", "anyKindOfString13_37Test"},
		{"some string", "someString"},
		{" some string", "someString"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamel(in, New())
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func TestToCamel(t *testing.T) {
	toCamel(t)
}

func TestCustomAcronymsToPascal(t *testing.T) {
	tests := []struct {
		name         string
		acronymKey   string
		acronymValue string
		expected     string
	}{
		{
			name:         "API Custom Acronym",
			acronymKey:   "API",
			acronymValue: "api",
			expected:     "Api",
		},
		{
			name:         "ABCDACME Custom Acroynm",
			acronymKey:   "ABCDACME",
			acronymValue: "AbcdAcme",
			expected:     "AbcdAcme",
		},
		{
			name:         "PostgreSQL Custom Acronym",
			acronymKey:   "PostgreSQL",
			acronymValue: "PostgreSQL",
			expected:     "PostgreSQL",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := New()
			conf[test.acronymKey] = test.acronymValue

			if result := ToPascal(test.acronymKey, conf); result != test.expected {
				t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
			}
		})
	}
}

func TestCustomAcronymsToCamel(t *testing.T) {
	tests := []struct {
		name         string
		acronymKey   string
		acronymValue string
		expected     string
	}{
		{
			name:         "API Custom Acronym",
			acronymKey:   "API",
			acronymValue: "api",
			expected:     "api",
		},
		{
			name:         "ABCDACME Custom Acroynm",
			acronymKey:   "ABCDACME",
			acronymValue: "AbcdAcme",
			expected:     "abcdAcme",
		},
		{
			name:         "PostgreSQL Custom Acronym",
			acronymKey:   "PostgreSQL",
			acronymValue: "PostgreSQL",
			expected:     "postgreSQL",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := New()
			conf[test.acronymKey] = test.acronymValue

			if result := ToCamel(test.acronymKey, conf); result != test.expected {
				t.Errorf("expected custom acronym result %s, got %s", test.expected, result)
			}
		})
	}
}

func BenchmarkToLowerCamel(b *testing.B) {
	benchmarkCamelTest(b, toCamel)
}

func benchmarkCamelTest(b *testing.B, fn func(testing.TB)) {
	for n := 0; n < b.N; n++ {
		fn(b)
	}
}
