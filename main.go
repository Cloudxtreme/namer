package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	dat, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err)
	}
	names := strings.Split(string(dat), " ")
	clean_names := []string{}
	for _, s := range names {
		if strings.TrimSpace(s) != "" {
			clean_names = append(clean_names, s)
		}
	}
	fmt.Println("Name:", clean_names[rand.Intn(len(clean_names))])
}
