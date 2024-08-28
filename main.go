package main

import (
	"fmt"
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