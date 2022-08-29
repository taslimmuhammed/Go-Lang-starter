package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file"

	file , err := os.Create("./mylcogofile.txt")
    
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is :", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	
	fmt.Println("text inside the data file is\n" , string(databyte))
}

func checkNilErr(err error){
	if err!= nil {
		panic(err)
	}
}