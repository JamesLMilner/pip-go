package pip_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/JamesMilnerUK/pip-go"
)

func TestPip(t *testing.T) {

	rectangle := pip.Polygon{
		Points: []pip.Point{
			pip.Point{X: 1.0, Y: 1.0},
			pip.Point{X: 1.0, Y: 2.0},
			pip.Point{X: 2.0, Y: 2.0},
			pip.Point{X: 2.0, Y: 1.0},
		},
	}

	pt1 := pip.Point{X: 1.1, Y: 1.1} // Should be true
	pt2 := pip.Point{X: 1.2, Y: 1.2} // Should be true
	pt3 := pip.Point{X: 1.3, Y: 1.3} // Should be true
	pt4 := pip.Point{X: 1.4, Y: 1.4} // Should be true
	pt5 := pip.Point{X: 1.5, Y: 1.5} // Should be true
	pt6 := pip.Point{X: 1.6, Y: 1.6} // Should be true
	pt7 := pip.Point{X: 1.7, Y: 1.7} // Should be true
	pt8 := pip.Point{X: 1.8, Y: 1.8} // Should be true

	pt9 := pip.Point{X: -4.9, Y: 1.2}    // Should be false
	pt10 := pip.Point{X: 10.0, Y: 10.0}  // Should be false
	pt11 := pip.Point{X: -5.0, Y: -6.0}  // Should be false
	pt12 := pip.Point{X: -13.0, Y: 1.0}  // Should be false
	pt13 := pip.Point{X: 4.9, Y: -1.2}   // Should be false
	pt14 := pip.Point{X: 10.0, Y: -10.0} // Should be false
	pt15 := pip.Point{X: 5.0, Y: 6.0}    // Should be false
	pt16 := pip.Point{X: -13.0, Y: 1.0}  // Should be false

	assert(pip.PointInPolygon(pt1, rectangle), true, t)
	assert(pip.PointInPolygon(pt2, rectangle), true, t)
	assert(pip.PointInPolygon(pt3, rectangle), true, t)
	assert(pip.PointInPolygon(pt4, rectangle), true, t)
	assert(pip.PointInPolygon(pt5, rectangle), true, t)
	assert(pip.PointInPolygon(pt6, rectangle), true, t)
	assert(pip.PointInPolygon(pt7, rectangle), true, t)
	assert(pip.PointInPolygon(pt8, rectangle), true, t)

	assert(pip.PointInPolygon(pt9, rectangle), false, t)
	assert(pip.PointInPolygon(pt10, rectangle), false, t)
	assert(pip.PointInPolygon(pt11, rectangle), false, t)
	assert(pip.PointInPolygon(pt12, rectangle), false, t)
	assert(pip.PointInPolygon(pt13, rectangle), false, t)
	assert(pip.PointInPolygon(pt14, rectangle), false, t)
	assert(pip.PointInPolygon(pt15, rectangle), false, t)
	assert(pip.PointInPolygon(pt16, rectangle), false, t)

	t.Log("Finished")

}

func BenchmarkPip(b *testing.B) {

	rectangle := pip.Polygon{
		Points: []pip.Point{
			pip.Point{X: 1.0, Y: 1.0},
			pip.Point{X: 1.0, Y: 2.0},
			pip.Point{X: 2.0, Y: 2.0},
			pip.Point{X: 2.0, Y: 1.0},
		},
	}

	pt1 := pip.Point{X: 1.1, Y: 1.1} // Should be true
	pt2 := pip.Point{X: 1.2, Y: 1.2} // Should be true
	pt3 := pip.Point{X: 1.3, Y: 1.3} // Should be true
	pt4 := pip.Point{X: 1.4, Y: 1.4} // Should be true
	pt5 := pip.Point{X: 1.5, Y: 1.5} // Should be true
	pt6 := pip.Point{X: 1.6, Y: 1.6} // Should be true
	pt7 := pip.Point{X: 1.7, Y: 1.7} // Should be true
	pt8 := pip.Point{X: 1.8, Y: 1.8} // Should be true

	pt9 := pip.Point{X: -4.9, Y: 1.2}    // Should be false
	pt10 := pip.Point{X: 10.0, Y: 10.0}  // Should be false
	pt11 := pip.Point{X: -5.0, Y: -6.0}  // Should be false
	pt12 := pip.Point{X: -13.0, Y: 1.0}  // Should be false
	pt13 := pip.Point{X: 4.9, Y: -1.2}   // Should be false
	pt14 := pip.Point{X: 10.0, Y: -10.0} // Should be false
	pt15 := pip.Point{X: 5.0, Y: 6.0}    // Should be false
	pt16 := pip.Point{X: -13.0, Y: 1.0}  // Should be false

	pip.PointInPolygon(pt1, rectangle)
	pip.PointInPolygon(pt2, rectangle)
	pip.PointInPolygon(pt3, rectangle)
	pip.PointInPolygon(pt4, rectangle)
	pip.PointInPolygon(pt5, rectangle)
	pip.PointInPolygon(pt6, rectangle)
	pip.PointInPolygon(pt7, rectangle)
	pip.PointInPolygon(pt8, rectangle)

	pip.PointInPolygon(pt9, rectangle)
	pip.PointInPolygon(pt10, rectangle)
	pip.PointInPolygon(pt11, rectangle)
	pip.PointInPolygon(pt12, rectangle)
	pip.PointInPolygon(pt13, rectangle)
	pip.PointInPolygon(pt14, rectangle)
	pip.PointInPolygon(pt15, rectangle)
	pip.PointInPolygon(pt16, rectangle)

}

func BenchmarkPipOneMillion(b *testing.B) {

	// Setup, let's not time this
	b.StopTimer()
	polygon := pip.Polygon{
		Points: []pip.Point{
			pip.Point{X: 0.0, Y: 0.0},
			pip.Point{X: 30.0, Y: 50.0},
			pip.Point{X: 0.0, Y: 100.0},
			pip.Point{X: 50.0, Y: 70.0},
			pip.Point{X: 100.0, Y: 100.0},
			pip.Point{X: 70.0, Y: 50.0},
			pip.Point{X: 100.0, Y: 0.0},
			pip.Point{X: 50.0, Y: 30.0},
		},
	}

	var x float64
	var y float64
	var pts []pip.Point

	for i := 0; i < 1000000; i++ {
		x = 100.0 * rand.Float64()
		y = 100.0 * rand.Float64()
		pts = append(pts, pip.Point{X: x, Y: y})
	}

	b.StartTimer()
	// Actually test the function
	for n := 0; n < b.N; n++ {
		for j := 0; j < len(pts); j++ {
			pip.PointInPolygon(pts[j], polygon)
		}
	}

}

func BenchmarkPipParallelOneMillion(b *testing.B) {

	// Setup, let's not time this
	b.StopTimer()
	polygon := pip.Polygon{
		Points: []pip.Point{
			pip.Point{X: 0.0, Y: 0.0},
			pip.Point{X: 30.0, Y: 50.0},
			pip.Point{X: 0.0, Y: 100.0},
			pip.Point{X: 50.0, Y: 70.0},
			pip.Point{X: 100.0, Y: 100.0},
			pip.Point{X: 70.0, Y: 50.0},
			pip.Point{X: 100.0, Y: 0.0},
			pip.Point{X: 50.0, Y: 30.0},
		},
	}

	var x float64
	var y float64
	var pts []pip.Point

	for i := 0; i < 1000000; i++ {
		x = 100.0 * rand.Float64()
		y = 100.0 * rand.Float64()
		pts = append(pts, pip.Point{X: x, Y: y})
	}

	b.StartTimer()
	// Actually test the function
	for n := 0; n < b.N; n++ {
		pip.PointInPolygonParallel(pts, polygon, 7)
	}

}

func assert(a bool, b bool, t *testing.T) bool {
	test := a == b
	t.Log("Was the point was correctly identified? " + strconv.FormatBool(test))
	return a == b
}
