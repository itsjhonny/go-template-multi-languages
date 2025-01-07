package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Translation map[string]string

var supportedLanguages = map[string]string{
	"en-us": "en",
	"pt-br": "pt_br",
	"es-es": "es",
}

func loadTranslations(locale string) (Translation, error) {
	file, err := os.Open(fmt.Sprintf("locales/%s.json", locale))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var translations Translation
	if err := json.NewDecoder(file).Decode(&translations); err != nil {
		return nil, err
	}

	return translations, nil
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    localeParam := "en-us"

    if len(parts) > 0 {
        if _, exists := supportedLanguages[parts[0]]; exists {
            localeParam = parts[0]
        }
    }

    locale := supportedLanguages[localeParam]

    translations, err := loadTranslations(locale)
    if err != nil {
        http.Error(w, "Failed to load translations", http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.New("index").ParseFiles(
        "templates/index.html",
        "templates/head.html",
        "templates/selectLang.html",
    ))

    data := map[string]interface{}{
        "Lang":         locale,
        "LocaleParam":  localeParam,
        "Translations": translations,
    }

    if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func setLanguageHandler(w http.ResponseWriter, r *http.Request) {
    locale := r.URL.Query().Get("lang")

    if _, exists := supportedLanguages[locale]; exists {
        http.SetCookie(w, &http.Cookie{
            Name:  "lang",
            Value: locale,
            Path:  "/",
        })
        http.Redirect(w, r, fmt.Sprintf("/%s/calc", locale), http.StatusFound)
        return
    }


    http.Redirect(w, r, "/en-us/calc", http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/en-us/calc", http.StatusFound)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/set-lang", setLanguageHandler)
	http.HandleFunc("/en-us/calc", calcHandler)
	http.HandleFunc("/pt-br/calc", calcHandler)
	http.HandleFunc("/es-es/calc", calcHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
