package main

import "fmt"

type Course struct {
	CourseId    int     `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Website   string `json:"website"`
}

func main() {
	fmt.Println("hello")
}
