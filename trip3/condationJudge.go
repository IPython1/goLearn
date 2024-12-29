package main

func main() {
	sign := true
	if sign {
		println("true")
	} else {
		println("false")
	}

	a, b := 1, 2
	sign = a < b
	switch sign {
	case true:
		println("true")
	case false:
		println("false")
	default:
		println("default")
	}

}
