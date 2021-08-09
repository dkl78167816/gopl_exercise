package main

// import "fmt"

type bird struct {
	age   int
	color int
}

var ketty bird

func check1(d *bird) *bird {
	return d
}

func check2(d bird) *bird {
	return &d
}

func check3(d *bird) bird {
	return *d
}

func check4(d bird) bird {
	return d
}

func main() {
	ketty.age = 10
	ketty.color = 1
	check1(&ketty).age = 1
	check2(ketty).age = 1
	// check3(&ketty).age = 1 error
	// check4(ketty).age = 1 error
	d := check4(ketty)
	d.age = 1
}
