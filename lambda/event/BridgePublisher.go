package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"lambda/model"
)

type BridgePublisher struct {
	svc *eventbridge.EventBridge
}

func (p *BridgePublisher) PublishPersona(ctx context.Context, persona model.Persona) error {
	pe, err := json.Marshal(persona)
	fmt.Printf("PublishPersona %s\n", string(pe))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Firing Events")
	output, err := p.svc.PutEventsWithContext(ctx, &eventbridge.PutEventsInput{
		Entries: []*eventbridge.PutEventsRequestEntry{
			{
				Source:     aws.String("my.person.service"),
				DetailType: aws.String("PersonCreated"),
				Detail:     aws.String(string(pe)),
			},
		},
	})
	if output != nil {
		fmt.Printf("PutEventsWithContext Output %s\n", output.String())
	}
	if err != nil {
		fmt.Printf("Error After PutEventsWithContenxt %s\n", err.Error())
	}
	return err
}

func NewBridgePublisher(sess *session.Session) *BridgePublisher {
	return &BridgePublisher{
		svc: eventbridge.New(sess),
	}
}
