package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
}

type Anything interface{}

type Hateoas struct {
	Anything
	Links map[string]string `json:"_links"`
}

func MarshalHateoas(subject interface{}) ([]byte, error) {
	h := &Hateoas{subject, make(map[string]string)}
	switch s := subject.(type) {
	case *User:
		h.Links["self"] = fmt.Sprintf("http://user/%d", s.Id)
	case *Session:
		h.Links["self"] = fmt.Sprintf("http://session/%d", s.Id)
	}
	return json.MarshalIndent(h, "", "    ")
}

func main() {
	u := &User{123, "James Dean"}
	s := &Session{456, 123}
	json, err := MarshalHateoas(u)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(json))
	}
	json, err = MarshalHateoas(s)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(json))
	}
}
