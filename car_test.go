package race

import (
	"testing"
)

var m = `A - 3 - B
A - 2 - C
A - 4 - E
B - 8 - C
C - 1 - D
A - 4 - E
E - 3 - D`

func TestCar_Go(t *testing.T) {
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

func BenchmarkCar_Go(b *testing.B) {
	car := NewCar(m)
	for i := 0; i < b.N; i++ {
		car.Go("A", "D")
	}
}
