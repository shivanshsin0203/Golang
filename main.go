package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func main(){
	fmt.Println("Making api in go")
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