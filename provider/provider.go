package provider

import "github.com/aws/aws-sdk-go/aws/session"

type SessionProvider interface {
	//Get user will create the session and send it to us for the use
	Get() (*session.Session, error)
}

type DefaultSession struct{}

func (defaultSession *DefaultSession) DefaultSessionProvider() (*session.Session, error) {
	return session.NewSession()
}

func (defaultSession *DefaultSession) Get() (*session.Session, error) {
	sess, err := session.NewSession()
	return sess, err
}
