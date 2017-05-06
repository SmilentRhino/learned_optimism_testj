package main

import (
    "encoding/json"
    "fmt"
//    "strings"
    //    "io"
    "io/ioutil"
//    "os"
)

const (
	PmB = iota
	PmG
	PvB
	PvG
	PsB
	PsG
)

type Choice struct{
    Answer string
    Mark string
}

type Question struct{
    Question string
    Group int
    Choices [2]Choice 
}

func main() {

    data, _ := ioutil.ReadFile("./questions_json")
    questions := [48]Question {}
    json.Unmarshal(data,&questions)
    for i :=range questions {
	fmt.Println(questions[i].Question)
    }
}
