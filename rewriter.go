package rewriter

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Source      string
	Destination string
}

type RegexRule struct {
	Source      *regexp.Regexp
	Destination string
}

type Rewriter struct {
	Rules      []Rule
	RegexRules []RegexRule
}

func NewRewriter(rules []Rule) Rewriter {
	regexRules := []RegexRule{}
	for _, rule := range rules {
		source := rule.RegexRule()
		regexRules = append(regexRules, RegexRule{source, rule.Destination})
	}

	return Rewriter{
		Rules:      rules,
		RegexRules: regexRules,
	}
}

func (rule Rule) RegexRule() *regexp.Regexp {
	k := rule.Source
	k = regexp.QuoteMeta(k)
	k = strings.Replace(k, `\*`, "(.*?)", -1)
	if strings.HasPrefix(k, `\^`) {
		k = strings.Replace(k, `\^`, "^", -1)
	}
	k = k + "$"
	return regexp.MustCompile(k)
}

func (rewriter Rewriter) Rewrite(path string) (string, error) {
	for _, regexRule := range rewriter.RegexRules {
		groups := regexRule.Source.FindAllStringSubmatch(path, -1)
		if groups != nil {
			values := groups[0][1:]
			replace := make([]string, 2*len(values))
			for i, v := range values {
				j := 2 * i
				replace[j] = "$" + strconv.Itoa(i+1)
				replace[j+1] = v
			}
			replacer := strings.NewReplacer(replace...)

			if replacer != nil {
				replaced := replacer.Replace(regexRule.Destination)
				return replaced, nil
			}
		}
	}
	return path, nil
}

func (rewriter Rewriter) MustRewrite(path string) string {
	replaced, err := rewriter.Rewrite(path)
	if err != nil {
		log.Fatal(err)
	}
	return replaced
}
