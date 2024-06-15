package main

import (
	"encoding/json"
	"io"
	"os"
)

type Student struct {
	Rating []int
}
type Group struct {
	Students []Student
}
type Rating struct {
	Average float32
}

func main() {
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var dat Group
	if err := json.Unmarshal(data, &dat); err != nil {
		panic(err)
	}
	var rt, st float32
	stud := dat.Students
	for i := range stud {
		st += 1
		for range stud[i].Rating {
			rt += 1
		}
	}
	result := rt / st
	var aver Rating
	aver.Average = result
	finish, _ := json.MarshalIndent(aver, "", "    ")
	io.WriteString(os.Stdout, string(finish))
}
