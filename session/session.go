package session

import "awesomeProject/utils"

const COOKIE = "sessionId"

type sessionData struct {
	username string
}

type Session struct {
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)
	s.data = make(map[string]*sessionData)
	return s
}

func (s *Session) Init(username string) string {
	sessionId := utils.GenerateId()
	data := &sessionData{username: username}
	s.data[sessionId] = data
	return sessionId
}
func (s *Session) Get(sessionId string) string {
	data := s.data[sessionId]
	if data == nil {
		return ""
	}
	return data.username

}
