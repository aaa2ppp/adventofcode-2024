package main

import "fmt"

type button struct {
	x, y  int
	price int
}

type prize struct {
	x, y int
}

func main() {
	a := button{42, 78, 3}
	b := button{49, 91, 1}

	n := (10000000000000 + 7*50) / 7

	p := prize{a.x * n, a.y * n}
	i := 1000000

	for i > 0 {
		p.x += 7
		p.y += 13
		if  p.x*a.y == p.y*a.x && p.x*b.y == p.y*b.x && p.x%a.x != 0 && p.x%b.x != 0{
			break
		}
		i--
	}

	if i > 0 {
		fmt.Println(p)
	} else {
		fmt.Println("oops!..")
	}

}
