package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	// accessing struct fields
	diblert.Salary -= 5000

	// taking the address and accessing through a pointer
	// because otherwise, the string wouldn't be mutable
	position := &dilbert.Position
	&position = "Senior " + *position

	// also works with a pointer to a whole struct
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
}

func structLiterals() {
	// struct typ ecan be written using a struct literal
	type Point struct{ X, Y int }

	p := Point{1, 2}
}

func comparingStruct() {
	type Point struct {
		X int
		Y int
	}

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"
	// structs are comparable if the fields are comparable

	// comparable struct types may be used as the key type of a map
	type address struct {
		hostname string
		port     int
	}

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

func embeddingStruct() {
	// first implementation
	type Circle1 struct {
		X, Y, Radius int
	}

	type Wheel1 struct {
		X, Y, Radius, Spokes int
	}

	var w Wheel1
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	// second implementation
	type Point2 struct {
		X, Y int
	}

	type Circle2 struct {
		Center Point
		Radius int
	}

	type Wheel2 struct {
		Circle Circle
		Spokes int
	}

	var w2 Wheel2
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	// problem with second implementation is that
	// it makes the accesing the fields of a Wheel more verbose

	// third implementation
	type Circle3 struct {
		Point2
		Radius int
	}

	type Wheel3 struct {
		Circle3
		Spokes int
	}

	// now the implementation allows for easy accessing the fields
	var w3 Wheel3
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20

	// There's not corresponding shorthand for the struct literal syntax
	w4 = Wheel3{8, 8, 5, 20}                       // compile error: unknown fields
	w5 = Wheel3{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	// to declare you need to use one of the two options below
	w6 = Wheel3{Circle{Point{8, 8}, 5}, 20}
	w7 = Wheel3{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20, // trailing comma is necessary
	}
}
