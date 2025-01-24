package main

import (
	"log/slog"
	"net/http"
	"sync"

	"sv_funapp_websocket/config"
	"sv_funapp_websocket/log"

	"github.com/gorilla/websocket"
)

// クライアントを管理する構造体
type ChatRoom struct {
	clients   map[*websocket.Conn]bool // チャットルームに接続しているクライアント管理情報(keyはWebSocket接続(*websocket.Conn))
	broadcast chan []byte              // メッセージをブロードキャストするためのチャネル
	mu        sync.Mutex
}

// ---- Global Variable
var IsDebug bool = true

// ---- Package Global Variable

// ---- public function ----

// ---- private function ----
// 初期化処理関数
func initialize() bool {

	// ロガー生成
	jsonLogger := log.GetInstance()
	jsonLogger.Info("main")

	// Config読み込み
	errConfig := config.ReadConfigInformation()
	if errConfig != nil {
		slog.Error("Read Config Failed")
		jsonLogger.Error("Read Config Failed", slog.String("error", errConfig.Error()))
		return false
	}
	return true
}

// WebSocketアップグレーダー(HTTP接続をWebSocket接続にアップグレードする)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,                                       // 読み取りバッファのサイズ
	WriteBufferSize: 1024,                                       // 書き込みバッファのサイズ
	CheckOrigin:     func(r *http.Request) bool { return true }, // 接続リクエストの「オリジン」を検証
}

// チャットルームを作成
func newChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

// クライアントを登録
func (room *ChatRoom) registerClient(conn *websocket.Conn) {
	room.mu.Lock()            // Mutexでロック（スレッドセーフにするため）
	defer room.mu.Unlock()    // 関数が終了したらロックを解除
	room.clients[conn] = true // クライアントをマップに追加
	slog.Info("New client connected", "Addr", conn.RemoteAddr().String())
}

// クライアントを削除
func (room *ChatRoom) unregisterClient(conn *websocket.Conn) {
	room.mu.Lock()             // Mutexでロック（スレッドセーフにするため）
	defer room.mu.Unlock()     // 関数が終了したらロックを解除
	delete(room.clients, conn) // クライアントをマップから削除
	conn.Close()               // WebSocket接続を閉じる
	slog.Info("Client disconnected", "Addr", conn.RemoteAddr().String())
}

// メッセージを全クライアントにブロードキャスト
func (room *ChatRoom) broadcastMessages() {
	for {
		message := <-room.broadcast      // データをチャンネルから受け取る
		room.mu.Lock()                   // Mutexでロック（スレッドセーフにするため）
		for conn := range room.clients { // 接続クライアント送信
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				room.unregisterClient(conn)
			}
		}
		room.mu.Unlock() // ロックを解除
	}
}

// WebSocketハンドラー
func (room *ChatRoom) handleConnections(w http.ResponseWriter, r *http.Request) {

	// WebSocket接続へのアップグレード
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("Error upgrading to WebSocket:", "error", err)
		return
	}
	defer room.unregisterClient(conn) // 接続が終了した時のクリーンアップ

	room.registerClient(conn) //　クライアント登録

	for { // WebSocket接続が維持されている限り、このループを実行
		_, message, err := conn.ReadMessage()
		if err != nil { // クライアント接続が切れた
			slog.Info("Connection error:", "error", err)
			break
		}
		room.broadcast <- message // データをチャンネルに送る
		slog.Info("BC", "message", message)
	}
}

//

func main() {

	isSuccess := initialize()
	if isSuccess == false {
		slog.Error("Initialize Failed")
		return
	}

	// jsonファイル読み書きテスト
	//test.TestJsonFileReadWrite()
	// csvファイル読み書きテスト
	//test.TestCsvFileReadWrite()

	// Redisへの接続テスト
	//test.TestRedisConnect()
	// PSQLへの接続テスト
	//test.TestPostgresConnect()

	// convertパッケージに用意された関数を実行するサンプル処理
	//test.TestConvertPackage()

	chatRoom := newChatRoom()
	go chatRoom.broadcastMessages()

	mux := http.NewServeMux()
	mux.Handle("/ws", http.HandlerFunc(chatRoom.handleConnections))

	slog.Info("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
