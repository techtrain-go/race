package solver

import (
	"github.com/RyanCarrier/dijkstra"
	"github.com/pkg/errors"
)

type Point struct {
	Left     int
	Distance int64
	Right    int
}

func Solve(points []Point, start, end int) (dijkstra.BestPath, error) {
	graph := dijkstra.NewGraph()

	visited := map[int]struct{}{}
	for _, value := range points {
		if _, ok := visited[value.Left]; ok {
			continue
		}

		visited[value.Left] = struct{}{}
		graph.AddVertex(value.Left)
	}

	for _, value := range points {
		if _, ok := visited[value.Right]; ok {
			continue
		}

		visited[value.Right] = struct{}{}
		graph.AddVertex(value.Right)
	}

	for _, value := range points {
		//fmt.Println(value.Left, value.Right, value.Distance)
		err := graph.AddArc(value.Left, value.Right, value.Distance)
		if err != nil {
			return dijkstra.BestPath{}, errors.Wrapf(err, "failed adding arcs %d->%d %d", value.Left, value.Right, value.Distance)
		}
	}

	return graph.Shortest(start, end)
}
