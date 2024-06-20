package sessionmanger

import (
	"container/list"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session struct {
	sid          string                 // sessionID
	timeAccessed time.Time              // 最后访问时间
	value        map[string]interface{} // 存放一些session必要的数据
}

type SessionManger struct {
	cookieName  string
	lock        sync.Mutex // 保护SessionManger整个变量
	maxLifeTime int64
	sessions    map[string]*list.Element
	list        *list.List //用来做GC
}

// NewSessionManger creates a new SessionManger
func NewSessionManger(cookiename string, maxLifeTime int64) *SessionManger {
	return &SessionManger{
		cookieName:  cookiename,
		maxLifeTime: maxLifeTime,
		list:        list.New(),
		sessions:    make(map[string]*list.Element, 0),
	}
}

func (sm *SessionManger) SessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (sm *SessionManger) SessionStart(w http.ResponseWriter, _ *http.Request) (session *Session) {
	sid := sm.SessionID()
	session = sm.sessionInit(sid)
	newcookie := http.Cookie{
		Name:     sm.cookieName,
		Value:    url.QueryEscape(sid),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(sm.maxLifeTime),
	}
	http.SetCookie(w, &newcookie)
	return
}

// SessionInit 返回一个Session变量，调用sessionInit前需加锁
func (sm *SessionManger) sessionInit(sid string) (session *Session) {
	session = &Session{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        make(map[string]interface{}, 0),
	}
	sm.lock.Lock()
	element := sm.list.PushBack(session)
	sm.sessions[sid] = element
	sm.lock.Unlock()
	return
}

// SessionGet 根据sid返回对应的Session，如果不存在，调用SessionInit()函数返回session
func (sm *SessionManger) SessionGet(w http.ResponseWriter, r *http.Request) *Session {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	cookie, err := r.Cookie(sm.cookieName)
	if err == nil && cookie.Value != "" {
		sid, _ := url.QueryUnescape(cookie.Value)
		if element, ok := sm.sessions[sid]; ok {
			return element.Value.(*Session)
		}
	}
	return nil
}

// SessionDestroy 销毁sid对应的Session
func (sm *SessionManger) SessionDestroy(sid string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if element, ok := sm.sessions[sid]; ok {
		delete(sm.sessions, sid)
		sm.list.Remove(element)
	}
}

// SessionGC 根据maxLifeTime来删除过期的Session
func (sm *SessionManger) SessionGC() {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	for {
		element := sm.list.Front()
		if element == nil {
			break
		}
		if element.Value.(*Session).timeAccessed.Unix()+sm.maxLifeTime < time.Now().Unix() {
			sm.list.Remove(element)
			delete(sm.sessions, element.Value.(*Session).sid)
		} else {
			break
		}
	}
	time.AfterFunc(time.Duration(sm.maxLifeTime), func() { sm.SessionGC() })
}

func (sm *SessionManger) SessionUpdate(sid string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if element, ok := sm.sessions[sid]; ok {
		element.Value.(*Session).timeAccessed = time.Now()
		sm.list.MoveToBack(element)
	}
}

func (s *Session) SessionId() string {
	return s.sid
}
