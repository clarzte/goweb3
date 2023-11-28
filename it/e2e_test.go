package it_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"cryptom.app/config"
	"cryptom.app/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite
	port    int
	baseUrl string
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &E2ETestSuite{})
}

func (s *E2ETestSuite) SetupSuite() {
	config.Load()

	s.port = config.ServerPort()
	s.baseUrl = fmt.Sprintf("http://localhost:%d", s.port)

	serverReady := make(chan bool)

	server := &server.Server{
		Port:            config.ServerPort(),
		MongoDbUri:      config.MongoDbConnectionUrl(),
		RpcUrl:          config.RpcUrl(),
		AllowedOrigins:  config.AllowedOrigins(),
		ServerReady:     serverReady,
		StorageHostname: config.StorageHostname(),
	}

	go server.Start()
	<-serverReady
}

func (s *E2ETestSuite) Test_EndToEnd_GetMerchantsWrongParams() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/merchants", s.baseUrl), nil)
	s.NoError(err)

	client := &http.Client{}

	res, err := client.Do(req)
	s.NoError(err)

	byteBody, err := io.ReadAll(res.Body)
	s.NoError(err)

	defer res.Body.Close()

	s.Equal(http.StatusBadRequest, res.StatusCode)
	s.Equal(`{"error":"Invalid query parameters"}`, strings.Trim(string(byteBody), "\n"))
}

func (s *E2ETestSuite) Test_EndToEnd_GetMerchantNearMe() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/merchants?radius=3000&lat=52.367979&lng=4.869938", s.baseUrl), nil)
	s.NoError(err)

	client := &http.Client{}

	res, err := client.Do(req)
	s.NoError(err)

	byteBody, err := io.ReadAll(res.Body)
	s.NoError(err)

	defer res.Body.Close()

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(`[{"merchantAddress":"0x1234567890123456789012345678901234567890","name":"jumbo","street":"Main St.","number":"10A","postcode":"12345","country":"Testland","lat":52.3791,"lng":4.873455,"description":"Your local shop"}]`, strings.Trim(string(byteBody), "\n"))
}

func (s *E2ETestSuite) Test_EndToEnd_GetMerchantsFromFarAway() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/merchants?radius=30000000&lat=52.367979&lng=4.869938", s.baseUrl), nil)
	s.NoError(err)

	client := &http.Client{}

	res, err := client.Do(req)
	s.NoError(err)

	byteBody, err := io.ReadAll(res.Body)
	s.NoError(err)

	defer res.Body.Close()

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(`[{"merchantAddress":"0x1234567890123456789012345678901234567890","name":"jumbo","street":"Main St.","number":"10A","postcode":"12345","country":"Testland","lat":52.3791,"lng":4.873455,"description":"Your local shop"},{"merchantAddress":"0x1224567890123456789012345678901234567890","name":"coffeeshop","street":"Main St.","number":"10A","postcode":"12345","country":"Testland","lat":19.96,"lng":22,"description":"The best warm drinks in town!"}]`, strings.Trim(string(byteBody), "\n"))
}

func (s *E2ETestSuite) Test_EndToEnd_GetMerchantByKeywordJum() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/merchants?radius=30000000&lat=52.367979&lng=4.869938&keyword=jum", s.baseUrl), nil)
	s.NoError(err)

	client := &http.Client{}

	res, err := client.Do(req)
	s.NoError(err)

	byteBody, err := io.ReadAll(res.Body)
	s.NoError(err)

	defer res.Body.Close()

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(`[{"merchantAddress":"0x1234567890123456789012345678901234567890","name":"jumbo","street":"Main St.","number":"10A","postcode":"12345","country":"Testland","lat":52.3791,"lng":4.873455,"description":"Your local shop"}]`, strings.Trim(string(byteBody), "\n"))
}

func (s *E2ETestSuite) Test_EndToEnd_GetMerchantsByKeywordDrink() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/merchants?radius=30000000&lat=52.367979&lng=4.869938&keyword=drink", s.baseUrl), nil)
	s.NoError(err)

	client := &http.Client{}

	res, err := client.Do(req)
	s.NoError(err)

	byteBody, err := io.ReadAll(res.Body)
	s.NoError(err)

	defer res.Body.Close()

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal(`[{"merchantAddress":"0x1224567890123456789012345678901234567890","name":"coffeeshop","street":"Main St.","number":"10A","postcode":"12345","country":"Testland","lat":19.96,"lng":22,"description":"The best warm drinks in town!"}]`, strings.Trim(string(byteBody), "\n"))
}
