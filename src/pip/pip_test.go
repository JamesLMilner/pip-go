package pip_test

import (
	"testing"
	"pip"
	"strconv"

)

func TestPip(t *testing.T) {

	rectangle := pip.Polygon{
		Points : []pip.Point {
			pip.Point{X : 1.0, Y : 1.0},
			pip.Point{X : 1.0, Y : 2.0},
			pip.Point{X : 2.0, Y : 2.0},
			pip.Point{X : 2.0, Y : 1.0},
		},
	}

	pt1 := pip.Point{X : 1.5, Y : 1.5} // Should be true
	pt2 := pip.Point{X : 4.9, Y : 1.2} // Should be false
	pt3 := pip.Point{X : 1.8, Y : 1.1} // Should be true
	pt4 := pip.Point{X : 1.5, Y : 1.5} // Should be true
	pt5 := pip.Point{X : 1.0, Y : 2.0} // Should be true (on the line)
	pt6 := pip.Point{X : 10.0, Y : 10.0} // Should be false

	assert(pip.PointInPolygon(pt1, rectangle), true, t)
	assert(pip.PointInPolygon(pt2, rectangle), false, t)
	assert(pip.PointInPolygon(pt3, rectangle), true, t)
	assert(pip.PointInPolygon(pt4, rectangle), true, t)
	assert(pip.PointInPolygon(pt5, rectangle), false, t)
	assert(pip.PointInPolygon(pt6, rectangle), false, t)
	t.Log("Finished")

}

func assert(a bool, b bool, t *testing.T) bool {
	test :=  a == b
	t.Log("The point was correctly identified " + strconv.FormatBool(test))
	return a == b
}
