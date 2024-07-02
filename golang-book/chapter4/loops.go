package main

import "fmt"

func main() {
	i := 1

	for i <= 5 {
		if i % 2 == 0 {
		    fmt.Println(i, `even`)
		} else {
			fmt.Println(i, `odd`)
		}
		i += 1
	}

	for i := 6; i <= 10; i++ {
		if i % 2 == 0 {
		    fmt.Println(i, `even`)
		} else {
			fmt.Println(i, `odd`)
		}
	}

	var no int

	fmt.Println(`Enter a no from 1 to 5 to get its text based form : `)
	fmt.Scanf("%d", &no)

	switch no {
	case 1: fmt.Println(`one`)
	case 2: fmt.Println(`two`)
	case 3: fmt.Println(`three`)
	case 4: fmt.Println(`four`)
	case 5: fmt.Println(`five`)
	default: fmt.Println(`Wrong no entered!!! Enter the no from 1 to 5.`)
	}
}