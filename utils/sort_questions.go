package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "strings"
    //    "io"
    "io/ioutil"
    "os"
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

    data, _ := ioutil.ReadFile("./raw_questions.text")
    result := strings.Split(string(data), "\n")
    questions := [48]Question {}
    for i :=range result {
	if i%4 == 0 {
	    questions[i/4].Question = result[i]
	    questions[i/4].Choices[0].Answer = result[i+1]
	    questions[i/4].Choices[1].Answer = result[i+2]
	}
        fmt.Println(result[i])
    }
    for i :=range questions {
	fmt.Println(questions[i].Question)
	fmt.Print("Please input type of the question, 0PmB, 1PmG, 2PvB," +
	    "3PvG, 4PsB, 5PsG:")
	fmt.Scanf("%d", &questions[i].Group)
	fmt.Println(questions[i].Choices[0].Answer)
	fmt.Print("Please input Mark of the Answer:")
	fmt.Scanf("%s", &questions[i].Choices[0].Mark)
	fmt.Println(questions[i].Choices[1].Answer)
	fmt.Print("Please input Mark of the Answer:")
	fmt.Scanf("%s", &questions[i].Choices[1].Mark)
    }
    question_json, _ := json.Marshal(questions)
    fmt.Print(string(question_json))
    f, _ := os.Create("questions_json")
    defer f.Close()
    w := bufio.NewWriter(f)
    n, _ := w.Write(question_json)
    fmt.Print("%d bytes written", n)
    w.Flush()
}
