
package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

type ContextFreeGrammar struct {
	//the start symbol
	start string
	//the rules
	rules map[string][]string

	//the nonterminals
	nonterminals map[string]bool
	//the terminals
	terminals map[string]bool
}

func (cfg *ContextFreeGrammar) String() string {
	var sb strings.Builder
	sb.WriteString("start: ")
	sb.WriteString(cfg.start)
	sb.WriteString(" rules: ")
	for k, v := range cfg.rules {
		sb.WriteString(k)
		sb.WriteString(" -> ")
		sb.WriteString(strings.Join(v, " | "))
		sb.WriteString("	")
	}
	return sb.String()
}

func (cfg *ContextFreeGrammar) IsTerminal(s string) bool {
	return cfg.terminals[s]
}

func (cfg *ContextFreeGrammar) IsNonterminal(s string) bool {
	return cfg.nonterminals[s]
}

func (cfg *ContextFreeGrammar) GetRules(s string) []string {
	return cfg.rules[s]
}

func (cfg *ContextFreeGrammar) GetStart() string {
	return cfg.start
}

func (cfg *ContextFreeGrammar) GetTerminals() []string {
	var result []string
	for k, _ := range cfg.terminals {
		result = append(result, k)
	}
	return result
}

func (cfg *ContextFreeGrammar) GetNonterminals() []string {
	var result []string
	for k, _ := range cfg.nonterminals {
		result = append(result, k)
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForNonterminal(s string) []string {
	return cfg.rules[s]
}

func (cfg *ContextFreeGrammar) GetRulesForTerminal(s string) []string {
	return cfg.rules[s]
}

func (cfg *ContextFreeGrammar) GetRulesForSymbol(s string) []string {
	return cfg.rules[s]
}

func (cfg *ContextFreeGrammar) GetRulesForSymbols(s []string) []string {
	var result []string
	for _, v := range s {
		result = append(result, cfg.rules[v]...)
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithNonterminal(s []string) []string {
	var result []string
	for _, v := range s {
		if cfg.IsNonterminal(v) {
			result = append(result, cfg.rules[v]...)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithTerminal(s []string) []string {
	var result []string
	for _, v := range s {
		if cfg.IsTerminal(v) {
			result = append(result, cfg.rules[v]...)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithSymbol(s []string) []string {
	var result []string
	for _, v := range s {
		result = append(result, cfg.rules[v]...)
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithSymbols(s []string) []string {
	var result []string
	for _, v := range s {
		result = append(result, cfg.rules[v]...)
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithNonterminalAndTerminal(s []string) []string {
	var result []string
	for _, v := range s {
		if cfg.IsNonterminal(v) {
			result = append(result, cfg.rules[v]...)
		}
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithNonterminalAndSymbol(s []string) []string {
	var result []string
	for _, v := range s {
		if cfg.IsNonterminal(v) {
			result = append(result, cfg.rules[v]...)
		}
	}
	return result
}

func (cfg *ContextFreeGrammar) GetRulesForSymbolsWithNonterminalAndSymbols(s []string) []string {
	var result []string
	for _, v := range s {
		if cfg.IsNonterminal(v) {
			result = append(result, cfg.rules[v]...)
		}
	}
	return result
}

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

