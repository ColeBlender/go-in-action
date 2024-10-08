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
	s6 := s5[1:3]
	// fmt.Println(s5[1:3])

	s6 = append(s6, 69)
	fmt.Println(s5, s6)

	s7 := []int{1, 2, 3, 4, 5, 6, 7}
	s8 := s7[1:2:3]
	fmt.Println(s8)
	s9 := append(s7, s8...)
	fmt.Println(s9)

	for i, x := range s9 {
		fmt.Println(i, x)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
