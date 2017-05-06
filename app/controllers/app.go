package controllers

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"strconv"
	//"strings"
	//    "io"
	"io/ioutil"
	"path"
	//"os"
)

const (
	PmB = iota
	PmG
	PvB
	PvG
	PsB
	PsG
)

type Choice struct {
	Answer string
	Mark   string
}

type Question struct {
	Question string
	Group    int
	Choices  [2]Choice
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	question_json := path.Join(revel.BasePath, "utils/questions_json")
	data, err := ioutil.ReadFile(question_json)
	if err != nil {
		fmt.Print(err.Error())
	}
	questions := [48]Question{}
	json.Unmarshal(data, &questions)

	return c.Render(questions)
}

func (c App) LOTest() revel.Result {
	myresult := map[int]int{
		PmB: 0,
		PmG: 0,
		PvB: 0,
		PvG: 0,
		PsB: 0,
		PsG: 0,
	}

	question_json := path.Join(revel.BasePath, "utils/questions_json")
	data, err := ioutil.ReadFile(question_json)
	if err != nil {
		fmt.Print(err.Error())
	}
	questions := [48]Question{}
	json.Unmarshal(data, &questions)
	for k, v := range questions {
		mark := c.Params.Form.Get(strconv.Itoa(k))
		//if mark = v.Choices
		switch v.Group {
		case PmB:
			if mark == "i" {
				myresult[PmB] += 1
			}
		case PmG:
			if mark == "i" {
				myresult[PmG] += 1
			}
		case PvB:
			if mark == "i" {
				myresult[PvB] += 1
			}
		case PvG:
			if mark == "i" {
				myresult[PvG] += 1
			}
		case PsB:
			if mark == "i" {
				myresult[PsB] += 1
			}
		case PsG:
			if mark == "i" {
				myresult[PsG] += 1
			}
		default:
			panic("unkonw choice")
		}
	}

	return c.Render(myresult)
}
