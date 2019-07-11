package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"
	"log"
)

type phone struct {
	number string
}

func (p *phone) send(m string) error {
	params := &sns.PublishInput{
		Message:     aws.String(m),
		PhoneNumber: aws.String(p.number),
	}

	conn := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))

	_, err := sns.New(conn).Publish(params)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Sending SMS failed [number: '%s', message: '%s']", p.number, m))
	}

	return nil
}

func main() {
	p := phone{number: "+32479707308"}
	err := p.send("test SMS sending")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
