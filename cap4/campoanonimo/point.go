package main

// sem anonimo

//type Point struct {
//	X, Y int
//}
//
//type Circle struct {
//	Center Point
//	Radius int
//}
//
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel

	// sem anonimo
	//w.Circle.Center.X = 8
	//w.Circle.Center.Y = 8
	//w.Circle.Radius = 5
	//w.Spokes = 20

	// com anonimo
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

}
