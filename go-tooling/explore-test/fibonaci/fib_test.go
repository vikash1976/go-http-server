package fibonaci

import "testing"

type fibTest struct {
	input    int
	expected int
}

var result int

var fibTests = []fibTest{
	{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

var fibTestsBigger = []fibTest{
	{10, 55}, {20, 6765}, {30, 832040}, {40, 102334155}, {10, 55}, {20, 6765}, {30, 832040}, {40, 102334155},
}

func TestFib10(t *testing.T) {
	t.Parallel()
	for _, tt := range fibTests {
		actual := Fib(tt.input)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}

}

func TestFib40(t *testing.T) {
	t.Parallel()
	for _, tt := range fibTestsBigger {
		actual := Fib(tt.input)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}

}

func benchmarkFib(num int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(num)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
