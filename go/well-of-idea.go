package main

func Well(x []string) (r string) {
	c := 0
	for _, v := range x {
		if v == "good" {
			c++
		}
	}

	switch c {
	case 0:
		r = "Fail!"
	case 1, 2:
		r = "Publish!"
	default:
		r = "I smell a series!"
	}
	return
}
