package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"lambda/event"
	"lambda/model"
	"lambda/storage"
	"net/http"
)

type Handler struct {
	storage   storage.Storage
	publisher event.Publisher
}

func NewHandler(storage storage.Storage, publisher event.Publisher) *Handler {
	return &Handler{
		storage:   storage,
		publisher: publisher,
	}
}

func (h *Handler) HandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request Received")
	if request.HTTPMethod != "POST" {
		fmt.Println("Request Received But it is not a POST")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}
	fmt.Println("Unmarshaling Request Body")
	var persona model.Persona
	err := json.Unmarshal([]byte(request.Body), &persona)
	if err != nil {
		fmt.Printf("Unmarshalling Error %s\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}, nil
	}
	persona.ID = uuid.New().String()
	fmt.Println("Saving Persona in Storage")
	err = h.storage.SavePersona(ctx, &persona)
	if err != nil {
		fmt.Printf("Storage Error:%s\n\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
	}
	fmt.Println("Publishing Persona")
	err = h.publisher.PublishPersona(ctx, persona)
	if err != nil {
		fmt.Printf("Publisher Error:%s\n\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
	}
	response, err := json.Marshal(persona)
	fmt.Println("Mashaling Response")
	if err != nil {
		fmt.Printf("Marshalling Error: %s\n\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
	}
	fmt.Println("Alles Goede")
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(response),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
