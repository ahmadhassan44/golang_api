package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

func (course Course) IsEmpty() bool {
	return course.CourseName == ""
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

func seedCourses() {
	// Create 5 sample courses with different authors
	courses = []Course{
		{
			CourseId:    uuid.New().String(),
			CourseName:  "Introduction to Computer Science",
			CoursePrice: 99,
			Author: &Author{
				FirstName: "John",
				LastName:  "Smith",
				Website:   "johnsmith.dev",
			},
		},
		{
			CourseId:    uuid.New().String(),
			CourseName:  "Web Development Fundamentals",
			CoursePrice: 149,
			Author: &Author{
				FirstName: "Sarah",
				LastName:  "Johnson",
				Website:   "sarahjohnson.com",
			},
		},
		{
			CourseId:    uuid.New().String(),
			CourseName:  "Go Programming Language",
			CoursePrice: 199,
			Author: &Author{
				FirstName: "Mike",
				LastName:  "Chen",
				Website:   "mikechen.io",
			},
		},
		{
			CourseId:    uuid.New().String(),
			CourseName:  "Database Design and Management",
			CoursePrice: 129,
			Author: &Author{
				FirstName: "Emily",
				LastName:  "Davis",
				Website:   "emilydavis.tech",
			},
		},
		{
			CourseId:    uuid.New().String(),
			CourseName:  "RESTful API Development",
			CoursePrice: 179,
			Author: &Author{
				FirstName: "Alex",
				LastName:  "Wilson",
				Website:   "alexwilson.dev",
			},
		},
	}
}

func main() {
	seedCourses()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	router.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	router.HandleFunc("/courses/{id}", updateCourse).Methods("PATCH")
	router.HandleFunc("/courses", addCourse).Methods("PUT")

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my golang api</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Course not found"})
}
func addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Please send course body to add course!"})
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Please send course body to add course!"})
		return
	}
	course.CourseId = uuid.New().String()
	courses = append(courses, course)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var cour Course
			_ = json.NewDecoder(r.Body).Decode(&cour)
			cour.CourseId = params["id"]
			courses = append(courses, cour)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cour)
			return
		}
	}
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "No ID provided"})
		return
	}
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
