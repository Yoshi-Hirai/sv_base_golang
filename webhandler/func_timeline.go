// webhandler webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"time"
)

// ---- const

// ---- struct
// タイムラインのポスト情報
type TimeLinePost struct {
	BoardId     int       `json:"boardid"`
	AccountId   int       `json:"accountid"`
	AccountName string    `json:"accountname"`
	PostTime    time.Time `json:"posttime"`
	Text        string    `json:"text"`
	CaptionUrl  string    `json:"captionurl"`
}

// ---- Package Global Variable
var postData []TimeLinePost

// ---- public function ----

// ---- private function ----
// TimeLine処理
func handlerTimelinePost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //オプションを解析します。デフォルトでは解析しません。
	slog.Info("Request", "method", r.Method, "form", r.Form, "path", r.URL.Path, "scheme", r.URL.Scheme, "url_long", r.Form["url_long"])
	if r.Method == http.MethodPost {

		// Body が空かどうか確認
		if r.Body == nil {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Body を読み取る
		body, err := ioutil.ReadAll(r.Body)
		if err != nil || len(body) == 0 {
			http.Error(w, "Failed to read body or empty body", http.StatusBadRequest)
			return
		}
		// JSON デコードを試行
		var params map[string]interface{}
		err = json.Unmarshal(body, &params)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		// パラメータが存在するかをチェック
		if len(params) == 0 {
			http.Error(w, "Empty JSON object", http.StatusBadRequest)
			return
		}
		slog.Info("post", "parms", params)

		var single TimeLinePost
		single.AccountId = int(params["accountId"].(float64))
		single.AccountName = "dummy"
		single.BoardId = int(params["boardId"].(float64))
		single.CaptionUrl = params["captionUrl"].(string)
		single.PostTime = time.Now()
		single.Text = params["text"].(string)
		postData = append(postData, single)

		// 正常なレスポンス
		w.WriteHeader(http.StatusCreated)
		output, err := json.Marshal(&postData)
		if err != nil {
			return
		}
		fmt.Fprintf(w, string(output)) //ここでwに入るものがクライアントに出力されます。

	} else {
		output, err := json.Marshal(&postData)
		if err != nil {
			return
		}
		fmt.Fprintf(w, string(output)) //ここでwに入るものがクライアントに出力されます。
	}
}
