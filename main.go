package main

import (
	"fmt"
	config "git.sr.ht/~hjertnes/doing/config"
	"git.sr.ht/~hjertnes/doing/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main(){
	conf, err := config.Read()
	if err != nil{
		panic(err)
	}

	text := strings.Join(os.Args[1:], " ")
	today := time.Now().Format("2006-01-02")

	filename := utils.ReplaceTilde(fmt.Sprintf("%s/%s.org", conf.Path, today))

	content, err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}



	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == ""{
		lines = lines[0:len(lines)-2]
	}
	lines = append(lines, fmt.Sprintf("- %s", text))

	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), os.ModePerm)
	if err != nil{
		panic(err)
	}
}
