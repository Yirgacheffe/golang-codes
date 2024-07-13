package main

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"http-nofiles/middleware"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/pure/v5"
	ut "github.com/go-playground/universal-translator"
)

var (
	tmpls  *template.Template
	uTrans *ut.UniversalTranslator
)

var transKey = struct {
	name string
}{
	name: "transKey",
}

type Translator interface {
	locales.Translator

	T(key interface{}, params ...string) string
	C(key interface{}, num float64, digits uint64, param string) string
	O(key interface{}, num float64, digits uint64, param string) string

	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string
	Currency() currency.Type
}

type translator struct {
	locales.Translator
	trans ut.Translator
}

var _ Translator = (*translator)(nil) // assertion

func (t *translator) T(key interface{}, param ...string) string {
	s, err := t.trans.T(key, param...)
	if err != nil {
		log.Printf("issue translating key: %v error: %s", key, err)
	}

	return s
}

func (t *translator) C(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating cardinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) O(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating ordinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string {

	s, err := t.trans.R(key, num1, digits1, num2, digits2, param1, param2)
	if err != nil {
		log.Printf("issue translating range key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) Currency() currency.Type {
	switch t.Locale() {
	case "en":
		return currency.USD
	case "fr":
		return currency.EUR
	case "zh":
		return currency.CNY
	default:
		return currency.USD
	}
}

func main() {
	en := en.New()
	uTrans = ut.New(en, en, fr.New(), zh.New())

	setup()

	var e error
	tmpls, e = template.ParseFiles("index.tmpl")
	if e != nil {
		panic(e)
	}

	r := pure.New()
	r.Use(middleware.LoggingAndRecovery(true), translatorMiddleware)
	r.Get("/", index)

	log.Println("running on port: 9091")
	log.Println("try: http://localhost:9091/?locale=en and http://localhost:9091/?locale=fr")

	http.ListenAndServe(":9091", r.Serve())
}

func setup() {
	en, _ := uTrans.FindTranslator("en")
	en.AddCardinal("days-left", "There is {0} day  left", locales.PluralRuleOne, false)
	en.AddCardinal("days-left", "There is {0} days left", locales.PluralRuleOther, false)

	fr, _ := uTrans.FindTranslator("fr")
	fr.AddCardinal("days-left", "Il reste {0} jour", locales.PluralRuleOne, false)
	fr.AddCardinal("days-left", "Il reste {0} jours", locales.PluralRuleOther, false)

	zh, _ := uTrans.FindTranslator("zh")
	zh.AddCardinal("days-left", "还剩于 {0} 天", locales.PluralRuleOne, false)
	zh.AddCardinal("days-left", "还剩于 {0} 天", locales.PluralRuleOther, false)

	if err := uTrans.VerifyTranslations(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t := r.Context().Value(transKey).(Translator)

	s := struct {
		Trans       Translator
		Now         time.Time
		PositiveNum float64
		NegativeNum float64
		Percent     float64
	}{
		Trans:       t,
		Now:         time.Now(),
		PositiveNum: 12345.22,
		NegativeNum: -8374.23,
		Percent:     90.76,
	}

	if err := tmpls.ExecuteTemplate(w, "index", s); err != nil {
		log.Fatal(err)
	}
}

func translatorMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		locale := params.Get("locale")
		var t ut.Translator

		if len(locale) > 0 {
			var found bool
			if t, found = uTrans.GetTranslator(locale); found {
				goto END
			}
		}
		t, _ = uTrans.FindTranslator(pure.AcceptedLanguages(r)...)

	END:
		r = r.WithContext(context.WithValue(r.Context(), transKey, &translator{trans: t, Translator: t.(locales.Translator)}))

		next(w, r)
	}
}
