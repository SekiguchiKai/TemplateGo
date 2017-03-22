package main

import (
	"html/template"
	"net/http"
	"log"
)
// メイン関数
func main() {
	// ハンドラをバンドル
	http.HandleFunc("/", Handler)

	// リスナー
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// テンプレートを生成
func Handler(w http.ResponseWriter, r *http.Request) {
	// 構造体Programを生成
	p := &Program{Languages:[]string{"Go", "Java", "Python", "JavaScript"},}

	// 関数を名前をつけてtemplateのFuncMapに登録
	funcMap := template.FuncMap{
		"judgeLangType" : JudgeLangType,
	}

	// テンプレート部分
	text := `
	Go : {{ index .Languages 0 | judgeLangType}}
	Java : {{ index .Languages 1 | judgeLangType}}
	Python : {{ index .Languages 2 | judgeLangType}}
	JavaScript : {{ index .Languages 3 | judgeLangType}}
	`

	// テンプレートをパースして、http.ResponseWriterに書き込み
	tmpl := template.Must(template.New("calculator").Funcs(funcMap).Parse(text))
	tmpl.Execute(w, p)
}

// 構造体
type Program struct {
	Languages []string
}

// 言語を引数に受けて、その言語の型を返す
func JudgeLangType(lang string) string{
	switch lang {
	case "Go", "Java":
		return "Static"
	case "Python", "JavaScript":
		return "Dynamic"
	default:
		return "unknown"
	}
}
