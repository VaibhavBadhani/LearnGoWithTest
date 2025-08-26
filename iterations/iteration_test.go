package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("A")
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("expected repeation was '%q' but we got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/************************************** Result **************************************

	go test -bench=. -benchmem


	goos: darwin
	goarch: arm64
	pkg: learn-go-with-test/iterations
	cpu: Apple M1 Pro
	BenchmarkRepeat-10      58892335                17.09 ns/op            8 B/op          1 allocs/op
	PASS
	ok      learn-go-with-test/iterations   1.792s

    func Benchmark(b *testing.B) {
		//... setup ...
		for b.Loop() {
			//... code to measure ...
		}
		//... cleanup ...
    }

    goos: Your operating system (darwin for macOS).
	goarch: Your architecture (arm64 since you’re on an M1 Pro).
	pkg: Which package is being tested (learn-go-with-test/iterations).
	cpu: The CPU info Go detected.
	B/op: the number of bytes allocated per iteration
    allocs/op: the number of memory allocations per iteration

************************************** END **************************************/
