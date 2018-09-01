package race

import (
	"os"
	"testing"
)

var m = `
A - 3 - B
A - 2 - C
A - 4 - E
B - 8 - C
C - 1 - D
E - 3 - B
B - 5 - D
E - 4 - L
L - 3 - F
D - 6 - G
G - 7 - H
F - 7 - D
L - 2 - H
G - 9 - L
E - 5 - M
E - 7 - N
M - 3 - N
L - 9 - K
H - 4 - I
K - 8 - H
`

func TestCar_Go(t *testing.T) {
	routes, err := os.Open("./routes")
	if err != nil {
		t.Fatal(err)
	}

	car := NewCar(routes)
	path, err := car.Go("A", "I")
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

func BenchmarkCar_Go(b *testing.B) {
	routes, err := os.Open("./routes")
	if err != nil {
		b.Fatal(err)
	}

	car := NewCar(routes)
	for i := 0; i < b.N; i++ {
		car.Go("A", "I")
	}
}
