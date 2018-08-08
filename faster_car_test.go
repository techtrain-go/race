package race

import (
	"testing"
)

func TestFasterCar_Go(t *testing.T) {
	car := NewCar(m)
	path := car.Go("A", "D")

	expected := []string{"A", "C", "D"}

	for i, city := range expected {
		if city != path[i] {
			t.Failed()
		}
	}

	t.Log(path)
}

func BenchmarkFasterCar_Go(b *testing.B) {
	car := NewFasterCar(m)
	for i := 0; i < b.N; i++ {
		car.Go("A", "D")
	}
}
