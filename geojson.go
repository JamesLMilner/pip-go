package pip

func PointInSlice(x float64, y float64, poly [][]float64) bool {

	j := 0
	intersect := false

	for i := 1; i < len(poly); i++ {

		if ((poly[i][1] > y) != (poly[j][1] > y)) &&
			(x < (poly[j][0]-poly[i][0])*(y-poly[i][1])/
				(poly[j][1]-poly[i][1])+poly[i][0]) {
			intersect = !intersect
		}

		j = i

	}

	return intersect
}

func PointInGeoJsonPolygon(pt []float64, poly [][][]float64) bool {

	x := pt[0]
	y := pt[1]
	outline := poly[0]

	// Check if inside GeoJson hole
	for i, poly := range poly {
		if i > 1 {
			if PointInSlice(x, y, poly) {
				return false
			}
		}
	}

	bb := GetBoundingBoxFromGeoJson(outline) // Get the bounding box of the polygon in question

	// If point not in bounding box return false immediately
	if !PointInGeoJsonBoundingBox(x, y, bb) {
		return false
	}

	// If the point is in the bounding box then we need to check the polygon
	return PointInSlice(x, y, outline)

}

func PointInGeoJsonBoundingBox(x float64, y float64, bb []float64) bool {
	// Check if point is in bounding box

	// Bottom Left is the smallest and x and y value
	// Top Right is the largest x and y value
	return x >= bb[0] && x <= bb[2] && y >= bb[1] && y <= bb[3]
}

func GetBoundingBoxFromGeoJson(poly [][]float64) []float64 {

	var maxX, maxY, minX, minY float64

	for i := 0; i < len(poly); i++ {
		side := poly[i]

		if side[0] > maxX || maxX == 0.0 {
			maxX = side[0]
		}
		if side[1] > maxY || maxY == 0.0 {
			maxY = side[1]
		}
		if side[0] < minX || minX == 0.0 {
			minX = side[0]
		}
		if side[1] < minY || minY == 0.0 {
			minY = side[1]
		}

	}

	return []float64{minX, minY, maxX, maxY}

}
