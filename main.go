package main

import (
	"fmt"
	"git.sr.ht/~hjertnes/doing/config"
	"git.sr.ht/~hjertnes/doing/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func help(){
	fmt.Println("doing is a small utility to append lines to your org-roam daily log")
	fmt.Println("Usage:")
	fmt.Println("\t doing text you want to append")
	fmt.Println("Configuration")
	fmt.Println("\t config file is stored at ~/.doing.yml")
	fmt.Println("\t currently has only one key Path that need to be pointed at your org-roam folder")
}

func success(){
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
func main(){
	if len(os.Args) > 1{
		success()
	} else {
		help()
	}
}
