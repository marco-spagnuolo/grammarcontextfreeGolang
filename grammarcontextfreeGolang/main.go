
package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)


//NewBWT

//use bwt to find all the strings that match a pattern
func (cfg *ContextFreeGrammar) Find(pattern string) []string {
	var result []string
	//the bwt
	bwt := NewBWT(cfg)
	//the pattern
	p := NewPattern(pattern)
	//the matches
	matches := bwt.Find(p)
	//the strings
	for _, match := range matches {
		result = append(result, match.String())
	}
	return result
}

//use bwt to find all the strings that match a pattern
func (cfg *ContextFreeGrammar) FindAll(pattern string) []string {
	var result []string
	//the bwt
	bwt := NewBWT(cfg)
	//the pattern
	p := NewPattern(pattern)
	//the matches
	matches := bwt.FindAll(p)
	//the strings
	for _, match := range matches {
		result = append(result, match.String())
	}
	return result
}

//construct a context free grammar from a file
func NewContextFreeGrammarFromFile(filename string) *ContextFreeGrammar {
	//the result
	var result = &ContextFreeGrammar{
		rules:        make(map[string][]string),
		nonterminals: make(map[string]bool),
		terminals:    make(map[string]bool),
	}
	//the lines
	var lines []string
	//the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//the scanner
	scanner := bufio.NewScanner(file)
	//read the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//the start
	result.start = lines[0]
	//the rules
	for _, line := range lines[1:] {
		//the parts
		parts := strings.Split(line, " -> ")
		//the left
		left := parts[0]
		//the right
		right := strings.Split(parts[1], " | ")
		//add the rule
		result.rules[left] = right
		//add the nonterminal
		result.nonterminals[left] = true
		//add the terminals
		for _, r := range right {
			for _, s := range strings.Split(r, " ") {
				if !result.nonterminals[s] {
					result.terminals[s] = true
				}
			}
		}
	}
	return result
}

//construct a context free grammar from a string

func NewContextFreeGrammarFromString(s string) *ContextFreeGrammar {
	//the result
	var result = &ContextFreeGrammar{
		rules:        make(map[string][]string),
		nonterminals: make(map[string]bool),
		terminals:    make(map[string]bool),
	}
	//the lines
	var lines []string
	//the scanner
	scanner := bufio.NewScanner(strings.NewReader(s))
	//read the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//the start
	result.start = lines[0]
	//the rules
	for _, line := range lines[1:] {
		//the parts
		parts := strings.Split(line, " -> ")
		//the left
		left := parts[0]
		//the right
		right := strings.Split(parts[1], " | ")
		//add the rule
		result.rules[left] = right
		//add the nonterminal
		result.nonterminals[left] = true
		//add the terminals
		for _, r := range right {
			for _, s := range strings.Split(r, " ") {
				if !result.nonterminals[s] {
					result.terminals[s] = true
				}
			}
		}
	}
	return result
}

//generate a file with  100 random sequence of 4 letters A, C, G, T
func GenerateRandomDNAFile(filename string) {
	//the file
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//the writer
	w := bufio.NewWriter(file)
	//the letters
	letters := []string{"A", "C", "G", "T"}
	//the random
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//write the lines
	for i := 0; i < 100; i++ {
		//the line
		var line string
		//the length
		length := r.Intn(100) + 1
		//the letters
		for j := 0; j < length; j++ {
			line += letters[r.Intn(4)]
		}
		//write the line
		w.WriteString(line + "	" + line + "	" + line + "	" + line)
	}
	//flush
	w.Flush()
}

