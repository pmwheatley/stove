package bnet

import (
	// "github.com/HearthSim/hs-proto/go"
	// "github.com/golang/protobuf/proto"
	"log"
)

type ChallengeNotifyServiceBinder struct{}

func (ChallengeNotifyServiceBinder) Bind(sess *Session) Service {
	service := &ChallengeNotifyService{sess}
	if sess != nil {
		go service.Run()
	}
	return service
}

type ChallengeNotifyService struct {
	sess *Session
}

func (s *ChallengeNotifyService) Name() string {
	return "bnet.protocol.challenge.ChallengeNotify"
}

func (s *ChallengeNotifyService) Methods() []string {
	return []string{
		"",
		"ChallengeUser",
		"ChallengeResult",
		"OnExternalChallenge",
		"OnExternalChallengeResult",
	}
}

func (s *ChallengeNotifyService) Invoke(method int, body []byte) (resp []byte, err error) {
	log.Panicf("ChallengeNotify is a client export, not a server export")
	return
}

func (s *ChallengeNotifyService) Run() {
	s.sess.WaitForTransition(StateLoggingIn)
	log.Println("ChallengeNotify received LoggingIn event")
}
