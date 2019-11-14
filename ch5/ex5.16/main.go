package main

import "fmt"

type typeString struct {
}

func main() {
	str := new(typeString)
	fmt.Println(str.Join("ciao", "sono", "enrico"))

}

func (s *typeString) Join(inputs ...string) string {
	temp := ""
	for _, v := range inputs {
		temp = fmt.Sprintf("%s %s", temp, v)
	}
	return temp
}
