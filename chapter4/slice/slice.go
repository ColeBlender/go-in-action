package slice

import "fmt"

func Slice() {
	s1 := make([]string, 5)
	fmt.Println(s1)

	s2 := make([]int, 3, 5)
	fmt.Println(s2)

	s3 := []string{"red", "blue", "green"}
	fmt.Println(s3)

	s4 := []string{99: ""}
	fmt.Println(s4)

	s5 := []int{1, 2, 3, 4, 5}
	fmt.Println(s5[1:3])
}
