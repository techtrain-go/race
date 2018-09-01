package race

import (
	"os"
	"testing"
)

func TestFasterCar_Go(t *testing.T) {
	routes, err := os.Open("./routes")
	if err != nil {
		t.Fatal(err)
	}

	car := NewCar(routes)
	path, err := car.Go("A", "D")
	if err != nil {
		t.Fatal(err)
	}

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
