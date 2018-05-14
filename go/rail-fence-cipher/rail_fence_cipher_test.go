package railfence

import "testing"

func testCases(op func(string, int) string, cases []testCase, t *testing.T) {
	for _, tc := range cases {
		if actual := op(tc.message, tc.rails); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %q\nActual: %q", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func TestEncode(t *testing.T) { testCases(Encode, encodeTests, t) }
func TestDecode(t *testing.T) { testCases(Decode, decodeTests, t) }

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range encodeTests {
			Encode(test.message, test.rails)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range decodeTests {
			Decode(test.message, test.rails)
		}
	}
}
