package main

import "fmt"

func main() {
	// arrays
	var a [5]int
	fmt.Println("emp:", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//slices
	//test := []string{}
	s := make([]string, 1)
	fmt.Println("emp:", s)

	s[0] = "a"

	fmt.Println("set:", s)
	fmt.Println("get:", s[0])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// range ussage
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// twoD := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	//     innerLen := i + 1
	//     twoD[i] = make([]int, innerLen)
	//     for j := 0; j < innerLen; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	//fmt.Println("2d: ", twoD)
}
