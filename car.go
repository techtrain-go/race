package race

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/techtrain-go/race/solver"
)

type Car struct {
	reader io.Reader
}

func NewCar(reader io.Reader) *Car {
	return &Car{
		reader: reader,
	}
}

type Point struct {
	Left  string
	Edge  int
	Right string
}

func (c *Car) Go(start, finish string) ([]string, error) {
	re := regexp.MustCompile(`(?P<Left>[A-Z1-9]+) - (?P<Edge>[0-9]+) - (?P<Right>[A-Z1-9]+)`)
	rawData, err := ioutil.ReadAll(c.reader)
	if err != nil {
		return nil, err
	}

	if len(rawData) == 0 {
		return nil, io.EOF
	}

	data := string(rawData)
	l := len(strings.Split(data, "\n")) - 1

	var all []map[string]string
	for i := 0; i < l; i++ {
		line := strings.Split(data, "\n")[i]
		match := re.FindStringSubmatch(line)
		params := map[string]string{}
		for i, name := range re.SubexpNames() {
			if i > 0 && i <= len(match) {
				params[name] = match[i]
			}
		}

		all = append(all, params)
	}

	var points []Point
	for _, p := range all {
		edge, _ := strconv.Atoi(p["Edge"])
		points = append(points, Point{
			Left:  p["Left"],
			Edge:  edge,
			Right: p["Right"],
		})
	}

	// Convert our points array to int->string table
	i := 1
	namesToCodes := map[string]int{}
	for _, p := range points {
		if _, ok := namesToCodes[p.Left]; ok {
			continue
		}

		namesToCodes[p.Left] = i
		i++
	}

	// Add right side too
	for _, p := range points {
		if _, ok := namesToCodes[p.Right]; ok {
			continue
		}

		namesToCodes[p.Right] = i
		i++
	}

	// To map return to original
	codesToNames := map[int]string{}
	for name, code := range namesToCodes {
		//fmt.Println(name, code)
		codesToNames[code] = name
	}

	var spoints []solver.Point
	for _, value := range points {
		leftCode, ok := namesToCodes[value.Left]
		if !ok {
			return nil, fmt.Errorf("failed to find '%s' for left code from incoming path in original map", value.Left)
		}

		rightCode, ok := namesToCodes[value.Right]
		if !ok {
			return nil, fmt.Errorf("failed to find '%s' for right code from incoming path in original map", value.Right)
		}
		spoints = append(spoints, solver.Point{
			Left:     leftCode,
			Distance: int64(value.Edge),
			Right:    rightCode,
		})
	}

	path, err := solver.Solve(spoints, namesToCodes[start], namesToCodes[finish])
	if err != nil {
		return nil, err
	}

	var resultPath []string
	for _, value := range path.Path {
		names, ok := codesToNames[value]
		if !ok {
			return nil, fmt.Errorf("failed to find code %d from result path in original map", value)
		}

		resultPath = append(resultPath, names)
	}

	return resultPath, nil
}
