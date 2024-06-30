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

var defaultHeader = map[string]string{
	"Content-Type":                "application/json",
	"Access-Control-Allow-Origin": "*",
}

func (h *Handler) HandlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request Received")
	switch request.HTTPMethod {
	case "POST":
		return h.createPersona(ctx, request)
	case "GET":
		if id, ok := request.PathParameters["id"]; ok {
			return h.getPersonaByID(ctx, id)
		}
		return h.getAllPersonas(ctx)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Method Not Allowed",
			Headers:    defaultHeader,
		}, nil
	}
}
func (h *Handler) createPersona(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Unmarshalling Request Body")
	var persona model.Persona
	err := json.Unmarshal([]byte(request.Body), &persona)
	if err != nil {
		fmt.Printf("Unmarshalling Error %s\n", err.Error())
		jsonError, _ := json.Marshal(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: string(jsonError), Headers: defaultHeader}, nil
	}
	persona.ID = uuid.New().String()
	fmt.Println("Saving Persona in Storage")
	err = h.storage.SavePersona(ctx, &persona)
	if err != nil {
		fmt.Printf("Storage Error:%s\n\n", err.Error())
		jsonError, _ := json.Marshal(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: string(jsonError), Headers: defaultHeader}, nil
	}

	fmt.Println("Publish Persona in Event Log")
	err = h.publisher.PublishPersona(ctx, persona)
	if err != nil {
		fmt.Printf("Publisher Error:%s\n\n", err.Error())
		jsonError, _ := json.Marshal(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: string(jsonError), Headers: defaultHeader}, nil
	}

	fmt.Println("Marshaling Response")
	responseBody, err := json.Marshal(persona)
	if err != nil {
		fmt.Printf("Marshalling Error: %s\n\n", err.Error())
		jsonError, _ := json.Marshal(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: string(jsonError), Headers: defaultHeader}, nil
	}

	fmt.Println("Alles Goede")
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(responseBody),
		Headers:    defaultHeader,
	}, nil
}

func (h *Handler) getPersonaByID(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {
	persona, err := h.storage.GetPersonaByID(ctx, id)
	if err != nil {
		fmt.Printf("Storage Error:%s\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error(), Headers: defaultHeader}, nil
	}

	if persona == nil {
		fmt.Printf("Persona with id: %s coudn't be found \n", id)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "Persona not found", Headers: defaultHeader}, nil
	}

	responseBody, err := json.Marshal(persona)
	if err != nil {
		fmt.Printf("Marshalling Error: %s\n\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error(), Headers: defaultHeader}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers:    defaultHeader,
	}, nil
}

func (h *Handler) getAllPersonas(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	personas, err := h.storage.GetAllPersonas(ctx)
	if err != nil {
		fmt.Printf("Storage Error:%s\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error(), Headers: defaultHeader}, nil
	}

	responseBody, err := json.Marshal(personas)
	if err != nil {
		fmt.Printf("Marshalling Error: %s\n\n", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error(), Headers: defaultHeader}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers:    defaultHeader,
	}, nil
}
