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

// Model for course - file
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

//middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func (c *Course) IsEmptyWhileCreate() bool {
	return c.CourseName == ""
}

//fake DB
var courses []Course

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Alll courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One courses")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One courses")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about data like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmptyWhileCreate() {
		json.NewEncoder(w).Encode("Please send some data, no data inside")
		return
	}

	//TODO: check only if the title is duplicate , steps - loop through and matched the title and if found then return error

	//generate unique id and convert to string

	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	//this sends response back to to the client
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One courses")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop, get id data, remove from array, update with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			// request body have course which is decoded into course type
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send a response when ID is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One courses")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop, get id data, remove from array

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			break
		}
	}

	

	//TODO: send a response when ID is not found


}

func main() {

	fmt.Println("API - RV.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 100, Author: &Author{Fullname: "Habibi Wallahi", Website: "rv.in"}})
	courses = append(courses, Course{"2", "Angular", 300, &Author{"Rohit Verma", "rv.in"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}
