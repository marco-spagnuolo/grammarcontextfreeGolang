package cfg

import (
	"strings"
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
