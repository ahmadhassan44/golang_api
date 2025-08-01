package main

import "fmt"

type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Website   string `json:"website"`
}

var courses []Course

//middlewares

func IsEmpty(course *Course) bool {
	return course.CourseId == "" && course.CourseName == ""
}

func main() {
	fmt.Println("hello")
}
