package request_test

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/x-vanio/client-and-server-api/pkg/request"
)

var (
	jsonStr = `{ 
		"USDBRL": { 
			"code": "USD", 
			"codein": 
			"BRL", 
			"name": "Dólar Americano/Real Brasileiro",
			"high": "4.75",
			"low": "4.6963",
			"varBid": "-0.0095",
			"pctChange": "-0.2",
			"bid": "4.7314",
			"ask": "4.7344",
			"timestamp": "1690577990",
			"create_date": "2023-07-28 17:59:50"
		}
	}`
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestByDollarQuoteSuccess(t *testing.T) {
	mockClient := &MockHTTPClient{}

	response := &http.Response{
		Body: io.NopCloser(strings.NewReader(jsonStr)),
	}

	mockClient.On("Do", mock.Anything).Return(response, nil)

	client := request.NewClient(mockClient, time.Millisecond*200)

	resp, err := client.ByDollarQuote("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestByDollarQuoteErrBadRequest(t *testing.T) {
	mockClient := &MockHTTPClient{}

	mockClient.On("Do", mock.Anything).Return(&http.Response{}, errors.New("bad request"))

	client := request.NewClient(mockClient, time.Millisecond*0)

	resp, err := client.ByDollarQuote("#")

	assert.Error(t, err)
	assert.Nil(t, resp)
}
