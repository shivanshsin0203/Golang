package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct{
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice int `json:"coursePrice"`
	Author *Author `json:"author"`
}
type Author struct{
	FullName string `json:"fullName"`
	WebSite string `json:"webSite"`
}
var courses []Course

func (c *Course) isEmpty() bool{
	return c.CourseName == "" 
}
func main(){
	fmt.Println("Making api in go")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Java", CoursePrice: 100, Author: &Author{FullName: "John Doe", WebSite: "www.johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{FullName: "Jane Doe", WebSite: "www.janedoe.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Go", CoursePrice: 300, Author: &Author{FullName: "Jane Doe", WebSite: "www.janedoe.com"}})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}
// Controller function
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for _, item := range courses{
		if item.CourseId == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}
func createCourse(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")
	if(r.Body == nil){
		json.NewEncoder(w).Encode("No data found")
		return
	}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil{
		json.NewEncoder(w).Encode("Invalid data")
		return
	}
	if course.isEmpty(){
		json.NewEncoder(w).Encode("Invalid data")
		return
	}
	
	course.CourseId =strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	
}
func updateCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item := range courses{
		if item.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil{
				json.NewEncoder(w).Encode("Invalid data")
				return
			}
			if course.isEmpty(){
				json.NewEncoder(w).Encode("Invalid data")
				return
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}
func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)
	for index, item := range courses{
		if item.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}