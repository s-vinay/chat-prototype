package userconncache

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var (
	wsconnusermap = make(map[string]*ThreadSafeWSConn)
	staticusers   = []string{"user-1", "user-2", "user-3", "user-4", "user-5", "user-6", "user-107", "user-108", "user-109", "user-110", "user-111"}
	statictoken   = "abcd1234pqr"
)

// Helper to make Gorilla Websockets threadsafe
type ThreadSafeWSConn struct {
	conn   *websocket.Conn
	mutex  sync.Mutex
	userid string
	open   bool
}

func checkValidUserWSLogin(userid string, token string) bool {
	log.Printf("Checking login for user %s \n", userid)
	if token != statictoken {
		log.Printf("Error, Invalid token for user %s \n", userid)
		return false
	}
	for _, v := range staticusers {
		if v == userid {
			return true
		}
	}
	log.Printf("Error, Invalid userid for user %s \n", userid)
	return false
}

func NewThreadSafeWSConn(conn *websocket.Conn, userid string, token string) (*ThreadSafeWSConn, error) {
	valid := checkValidUserWSLogin(userid, token)
	if !valid {
		log.Printf("Error, Invalid login for user %s \n", userid)
		return nil, errors.New("invalid login details")
	}
	t := &ThreadSafeWSConn{conn, sync.Mutex{}, userid, true}
	t.conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Error, closed ws session for user %s due to code %d, reason %s \n", userid, code, text)
		t.Close()
		return nil
	})
	wsconnusermap[userid] = t
	return t, nil
}

func IsUserOnline(userid string) bool {
	if conn := wsconnusermap[userid]; conn != nil {
		return conn.open
	}
	log.Printf("Error, User offline for user %s \n", userid)
	return false
}

func (t *ThreadSafeWSConn) ReadWSConn() *websocket.Conn {
	return t.conn
}

func (t *ThreadSafeWSConn) WriteJSON(v interface{}) {
	if !t.open {
		log.Printf("Error, sending message failed for conn closed already for %s \n", t.userid)
		return
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if err := t.conn.WriteJSON(v); err != nil {
		log.Printf("Error, sending message failed for user %s \n", t.userid)
		t.conn.Close()
		t.open = false
		delete(wsconnusermap, t.userid)
	}
}

func (t *ThreadSafeWSConn) Close() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.conn.Close()
	t.open = false
	delete(wsconnusermap, t.userid)
}
