package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var sum int
	for _, score := range s.score {
		sum += score
	}
	return float64(sum) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score > max {
			max = score
			name = s.name[i]
		}
	}
	return max, name
}
func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is: ", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is: , "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is: , "+nameMin+" (", scoreMin, ")")

}
