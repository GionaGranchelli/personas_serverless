package api

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lambda/model"
	"testing"
)

// Mock Storage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetAllPersonas(ctx context.Context) ([]*model.Persona, error) {
	args := m.Called(ctx)
	var personas []*model.Persona
	personas = append(personas, &model.Persona{
		ID:          "123",
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Address:     "123 Main St",
	})
	return personas, args.Error(1)
}
func (m *MockStorage) SavePersona(ctx context.Context, persona *model.Persona) error {
	args := m.Called(ctx, persona)
	return args.Error(0)
}
func (m *MockStorage) GetPersonaByID(ctx context.Context, id string) (*model.Persona, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Persona), args.Error(1)
}

// Mock Storage
type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PublishPersona(ctx context.Context, persona model.Persona) error {
	args := m.Called(ctx, persona)
	return args.Error(0)
}

func TestCreatePersona(t *testing.T) {
	mockStorage := new(MockStorage)
	mockPublisher := new(MockPublisher) // Similar to MockStorage
	handler := NewHandler(mockStorage, mockPublisher)

	persona := model.Persona{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Address:     "123 Main St",
	}

	body, _ := json.Marshal(persona)
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       string(body),
	}

	mockStorage.On("SavePersona", mock.Anything, mock.Anything).Return(nil)
	mockPublisher.On("PublishPersona", mock.Anything, mock.Anything).Return(nil)

	response, err := handler.HandlerRequest(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, 201, response.StatusCode)
	mockStorage.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestGetPersonaByID(t *testing.T) {
	mockStorage := new(MockStorage)
	mockPublisher := new(MockPublisher)
	handler := NewHandler(mockStorage, mockPublisher)

	persona := &model.Persona{
		ID:          "123",
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456789",
		Address:     "123 Main St",
	}

	mockStorage.On("GetPersonaByID", mock.Anything, "123").Return(persona, nil)

	request := events.APIGatewayProxyRequest{
		HTTPMethod:     "GET",
		PathParameters: map[string]string{"id": "123"},
	}

	response, err := handler.HandlerRequest(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	var responseBody model.Persona
	err = json.Unmarshal([]byte(response.Body), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, persona, &responseBody)

	mockStorage.AssertExpectations(t)
}

func TestGetAllPersonas(t *testing.T) {
	mockStorage := new(MockStorage)
	mockPublisher := new(MockPublisher)
	handler := NewHandler(mockStorage, mockPublisher)

	personas := []*model.Persona{
		{
			ID:          "123",
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "123456789",
			Address:     "123 Main St",
		},
	}

	mockStorage.On("GetAllPersonas", mock.Anything).Return(personas, nil)

	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
	}

	response, err := handler.HandlerRequest(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	var responseBody []*model.Persona
	err = json.Unmarshal([]byte(response.Body), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, personas, responseBody)

	mockStorage.AssertExpectations(t)
}
