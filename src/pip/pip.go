package pip

type Point struct {
	// A point
	X float64
	Y float64
}

type Polygon struct {
	// A polygon
	Points []Point
}

type BoundingBox struct {
	BottomLeft Point
	TopRight Point
}

func PointInPolygon(pt Point, poly Polygon) bool {
	// Checks if point is inside polygon

	bb := GetBoundingBox(poly)
	if PointInBoundingBox(pt, bb) {
		// If the point is in the bounding box then we need to check the polygon

		nverts := len(poly.Points)
		intersect := false

		verts := poly.Points
		j := 0

		for i := 1; i < nverts; {

			if (
			 (verts[i].Y > pt.Y) != (verts[j].Y > pt.Y)) &&
			 (pt.X < (verts[j].X - verts[i].X) * (pt.Y - verts[i].Y) / (verts[j].Y - verts[i].Y) + verts[i].X) {
				intersect = !intersect
			}

			j = i
			i++

		}

		return intersect

	} else {
		// Else we can just return false because it's outside the bounding box
		return false
	}


}

func PointInBoundingBox(pt Point, bb BoundingBox) bool {

	bbMaxX := bb.TopRight.X
	bbMaxY := bb.TopRight.Y
	bbMinX := bb.BottomLeft.X
	bbMinY := bb.BottomLeft.Y

	return pt.X < bbMaxX && pt.X > bbMinX &&
	 	   pt.Y < bbMaxY && pt.Y > bbMinY

}

func GetBoundingBox(poly Polygon) BoundingBox {

	maxX := 0.0
	maxY := 0.0
	minX := 0.0
	minY := 0.0

	for i := 0; i < len(poly.Points); i++ {
		side := poly.Points[i]

		if side.X > maxX { maxX = side.X }
		if side.Y > maxX { maxY = side.Y }
		if side.X < minX { minX = side.X }
		if side.Y < minY { minY = side.Y }
	}

	return BoundingBox{
			BottomLeft : Point{ X: minX, Y : minY},
			TopRight     : Point{ X: maxX, Y : maxY},
		}

}
