package main

import (
	"log"
	"strings"
	"bufio"
	"os"
)

func main() {   
   scanner := bufio.NewScanner(os.Stdin)
   input := ""
	
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	
	 parts := strings.Split(input, " ")
	 result := ""
	 
	 for i := 0; i < len(parts); i++ {
	 	result += strings.ToLower(parts[i])
		 if len(parts)-1 != i {
		 	result += "-"
		 }
	 }
	 
    
	log.Printf("%v\n", result)
}