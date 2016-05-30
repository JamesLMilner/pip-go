# :dart: Point in Polygon - Go

Detect if a point (X, Y) is within a given polygon (Set of X, Ys) .

## Usage

    import pip

    // Polygons are self closing (i.e. you don't need to add first point at end)
    rectangle := pip.Polygon{
      Points : []pip.Point {
        pip.Point{X : 1.0, Y : 1.0},
        pip.Point{X : 1.0, Y : 2.0},
        pip.Point{X : 2.0, Y : 2.0},
        pip.Point{X : 2.0, Y : 1.0},
      },
    }

    pt1 := pip.Point{X : 1.1,  Y : 1.1}

    pip.PointInPolygon(pt1, rectangle) // Test - Should return true

## Caveats

* Currently no support for holes although this should be easy to account for if the hole is tested as a separate polygon.
* Currently no support for points that reside on edges (returns outside, i.e. false).
* Beware of roundoff errors
* The direction that you list the vertices (clockwise or counterclockwise) does not matter.

## Credit

Based on the example code given in [this article]( https://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html) by W. Randolph Franklin. The method is named PNPOLY.


## License
MIT License
Original C code : Copyright (c) 1970-2003, Wm. Randolph Franklin
Modified Go code: Copyright 2016 James Milner
