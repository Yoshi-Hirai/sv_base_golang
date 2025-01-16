// webhandler webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

// ---- const

// ---- struct
// タイムラインのポスト情報
type TimeLinePost struct {
	BoardId     int    `"json:"boardid"`
	AccountId   int    `json:"accountid"`
	AccountName string `json:"accountname"`
	PostTime    string `json:"posttime"`
	Text        string `json:"text"`
	CaptionUrl  string `json:"captionurl"`
}

// ---- Package Global Variable

// ---- public function ----

// ---- private function ----
// TimeLine処理
func handlerTimelinePost(w http.ResponseWriter, r *http.Request) {

	post := TimeLinePost{
		BoardId:     1,
		AccountId:   19,
		AccountName: "BanNam",
		PostTime:    "2025/01/15 14:00:00",
		Text:        "こんにちわ。",
		CaptionUrl:  "",
	}

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

	output, err := json.Marshal(&post)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(output)) //ここでwに入るものがクライアントに出力されます。
}
