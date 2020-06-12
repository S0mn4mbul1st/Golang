package main

import (
	"fmt"
	"time"
)

	func plus(a int, b int) int{

	return a + b
	}

	func plusPlus(a, b, c int) int {
	return a + b + c
	}


func vals() (int, int) {
	return 3, 7
}

func main() {

	fmt.Println("hello world")
	fmt.Println("Aykhan" + "Imanov")

	fmt.Println("8.0/2.0=", 8.0/2.0)
	fmt.Println("8.0/2=", 8.0/2)
	fmt.Println("8/2.0=", 8/2.0)
	fmt.Println("8/2=", 8/2)

	fmt.Println("7.0/2.0=", 7.0/2.0)
	fmt.Println("7.0/2=", 7.0/2)
	fmt.Println("7/2.0=", 7/2.0)
	fmt.Println("7/2=", 7/2)

	fmt.Println("36%11=", 36%11)

	var str = "Hello, World!"
	fmt.Println(str)

	var b, c int = 1,2
	fmt.Println("B = " , b , "\nC = ", c )

	var num int = 11
	if  num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "Apple"
	fmt.Println(f)

	i := 2

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("\nIt's the weekend")
	default:
		fmt.Println("\nIt's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}


	whatAmI := func(i interface{}) {

		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")


	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	var twoD [5][3]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	aa := make([]string, len(s))
	copy(aa, s)
	fmt.Println("cpy:", aa)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	dd := []string{"g", "h", "i"}
	fmt.Println("dcl:", dd)

	twoM := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoM[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoM[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoM)

	m := make ( map [string] int )

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m) )

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)


	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	aA, bB := vals()
	fmt.Println(aA)
	fmt.Println(bB)

	_, cC := vals()
	fmt.Println(cC)



}