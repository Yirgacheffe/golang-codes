package main

import (
	"fmt"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_CA"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/nl"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
)

var universalTranslator *ut.UniversalTranslator

func main() {

	e := en.New()
	universalTranslator = ut.New(e, e, en_CA.New(), nl.New(), fr.New(), zh.New())

	fr, _ := universalTranslator.GetTranslator("fr")

	fmt.Println("Cardinal:", fr.PluralsCardinal())
	fmt.Println("Ordinal:", fr.PluralsOrdinal())
	fmt.Println("Range:", fr.PluralsRange())

	//
	fr.Add("welcome", "欢迎 {0} to our test code", false)

	greetings, _ := fr.T("welcome", "Viper")
	fmt.Println(greetings)

	//
	fr.AddCardinal("days", "You have {0} day left to register", locales.PluralRuleOne, false)
	fr.AddCardinal("days", "You have {0} days left to register", locales.PluralRuleOther, false)

	oneDay, _ := fr.C("days", 1, 0, fr.FmtNumber(1, 0))
	others, _ := fr.C("days", 2, 0, fr.FmtNumber(2, 0))
	digs, _ := fr.C("days", 114.25, 2, fr.FmtNumber(114.25, 2))

	fmt.Println(oneDay)
	fmt.Println(others)
	fmt.Println(digs)

	//
	fr.AddOrdinal("day-of-month", "{0}st", locales.PluralRuleOne, false)
	fr.AddOrdinal("day-of-month", "{0}nd", locales.PluralRuleTwo, false)
	fr.AddOrdinal("day-of-month", "{0}rd", locales.PluralRuleFew, false)
	fr.AddOrdinal("day-of-month", "{0}th", locales.PluralRuleOther, false)

	fmt.Println(fr.O("day-of-month", 1, 0, fr.FmtNumber(1, 0)))
	fmt.Println(fr.O("day-of-month", 2, 0, fr.FmtNumber(2, 0)))
	fmt.Println(fr.O("day-of-month", 3, 0, fr.FmtNumber(3, 0)))
	fmt.Println(fr.O("day-of-month", 4, 0, fr.FmtNumber(4, 0)))
	fmt.Println(fr.O("day-of-month", 1046.20, 0, fr.FmtNumber(1046.20, 0)))

	//
	fr.AddRange("between", "It's {0}-{1} day away", locales.PluralRuleOne, false)
	fr.AddRange("between", "It's {0}-{1} days away", locales.PluralRuleOther, false)

	fmt.Println(fr.R("between", 0, 0, 1, 0, fr.FmtNumber(0, 0), fr.FmtNumber(1, 0)))
	fmt.Println(fr.R("between", 1, 0, 2, 0, fr.FmtNumber(1, 0), fr.FmtNumber(2, 0)))
	fmt.Println(fr.R("between", 1, 0, 100, 0, fr.FmtNumber(1, 0), fr.FmtNumber(100, 0)))
}
