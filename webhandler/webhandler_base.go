// webhandler webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"sv_base/auth"
)

// ---- Package Global Variable

// ---- public function ----
// WebHandlerInit (public)Webhandlerシステムの初期化関数。handlerを登録し指定されたポートでサーバーを立ち上げる
func WebHandlerInit(port string) error {
	slog.Info("WebHandler Open", "port", port)

	http.Handle("/sayhello", auth.AuthMiddleware(http.HandlerFunc(handleSayHelloName)))
	http.HandleFunc("/jsonSampleResponse", handleJsonSampleResponse)

	http.HandleFunc("/login", HandlerLogin)
	http.HandleFunc("/protected", HandlerProtected)

	webhandlerErr := http.ListenAndServe(":"+port, nil)
	if webhandlerErr != nil {
		return webhandlerErr
	}

	return nil
}

// ---- private function ----
// 固定テキストを応答するサンプルWebhandle
func handleSayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //オプションを解析します。デフォルトでは解析しません。
	slog.Info("Request", "form", r.Form, "path", r.URL.Path, "scheme", r.URL.Scheme, "url_long", r.Form["url_long"])
	str := ""
	for k, v := range r.Form {
		str = str + k + " "
		str = str + strings.Join(v, "") + " "
	}

	len := r.ContentLength
	body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
	r.Body.Read(body)         // byte 配列にリクエストボディを読み込む
	fmt.Fprintf(w, string(body))
	fmt.Fprintf(w, "Hello Brown!") //ここでwに入るものがクライアントに出力されます。
}

// Jsonを応答するサンプルwebhandler
type Post struct {
	Id      int      `json:"id"`
	Content string   `json:"content"`
	Member  []Member `json:"member"`
	Result  []Result `json:"result"`
}
type Member struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Result struct {
	Id     int    `json:"id"`
	Rank   int    `json:"rank"`
	Race   string `json:"race"`
	Driver string `json:"driber"`
}

func handleJsonSampleResponse(w http.ResponseWriter, r *http.Request) {
	post := Post{
		Id:      1,
		Content: "RB Formula1",
		Member: []Member{
			Member{
				Id:   1,
				Name: "Yuki",
			},
			Member{
				Id:   2,
				Name: "Liam",
			},
		},
		Result: []Result{
			Result{
				Id:     22,
				Rank:   9,
				Race:   "LVG",
				Driver: "Yuki",
			},
			Result{
				Id:     22,
				Rank:   16,
				Race:   "LVG",
				Driver: "Liam",
			},
		},
	}

	output, err := json.Marshal(&post)
	if err != nil {
		return
	}
	fmt.Fprint(w, string(output))
}
