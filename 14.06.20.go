package main

import (
	"errors"
	"fmt"
	"math"
)

type kvadrat struct {

	width, height int
}

func (r* kvadrat) area() int {

	return r.height * r.width
}

func (r kvadrat) perimeter() int{

	return ( 2 * r.height ) + ( 2 * r.width )
}

func (r* kvadrat) double(){

	r.height *= 2

}

func (r kvadrat) dd(){

	r.height *= 2

}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func f1( arg int) (int, error){

	if arg < 0 {

		return -1, errors.New("Can't Work With Neg Numbers")
	}

	return 1, nil
}


type argErr struct{
	arg int
	prob string
}

func (e *argErr) Error() string{

	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {

	if arg > 0 {

		return -1, &argErr{arg, "Can't Work With Positive Numbers"}
	}
	return 1, nil
}

func main() {

	vari := kvadrat { width:10 , height:5 }

	fmt.Println( vari.area() )
	fmt.Println( vari.perimeter() )

	ref := &vari

	fmt.Println( ref.area() )
	fmt.Println( ref.perimeter() )

	vari.dd()
	fmt.Println(vari.height)

	vari.double()
	fmt.Println(vari.height)

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{-16, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argErr); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}