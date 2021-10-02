package messagescache

import (
	"bytes"
	"goserver/pkg/apiobjects"
	"sort"
	"sync"
)

var (
	usersconvsmap = make(map[string]string)
	convsmap      = make(map[string]*PrivateConversation)
	mutex         = sync.Mutex{}
)

type PrivateConversation struct {
	messageschan   chan *apiobjects.Message
	conversationid string
	users          [2]string
}

func sortUserIdToAlphabeticString(users ...string) string {
	sort.Strings(users)
	var b bytes.Buffer
	for _, v := range users {
		b.WriteString(v)
	}
	return b.String()
}

func FetchPrivateConversationByUsers(user1 string, user2 string) string {
	mutex.Lock()
	defer mutex.Unlock()
	return usersconvsmap[sortUserIdToAlphabeticString(user1, user2)]
}

func createNewPrivateConversationContext() {

}

func AddMessageToConversation(message apiobjects.Message) {
	mutex.Lock()
	defer mutex.Unlock()

}
