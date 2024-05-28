package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helper
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName==""
}
func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r:=mux.NewRouter()
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	r.HandleFunc("/",serverHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))

}
func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by LCO"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-type","application/json")
	//grab id from request
	params :=mux.Vars(r)
	//looping through courses
	for _,course:= range courses{
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
func CreateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("content-type","application/json")
	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some data")

	}

	//what about - {}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//generating a unique id string
	rand.Seed(time.Now().UnixNano())
	
	course.CourseId= strconv.Itoa(rand.Intn(100))
	courses=append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
	
}
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating the course")
	w.Header().Set("content-type","application/json")
	
	params:=mux.Vars(r)

	for index,course:= range courses{
		if course.CourseId==params["id"]{
			courses:=append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}




	
}
func DeleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete One course")
	w.Header().Set("content-type","application/json")
	params:=mux.Vars(r)

	for index,course:= range courses{
		if course.CourseId == params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			json.NewEncoder(w).Encode("Succesfully deleted")
			break
		}
	}
	json.NewEncoder(w).Encode("Not deleted")
	return
}