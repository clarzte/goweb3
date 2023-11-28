package merchants

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MerchantDTO struct {
	MerchantAddress string  `json:"merchantAddress"`
	Name            string  `json:"name"`
	Street          string  `json:"street"`
	Number          string  `json:"number"`
	PostCode        string  `json:"postcode"`
	Country         string  `json:"country"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Description     string  `json:"description"`
}

type Storage struct {
	httpClient *http.Client
	url        string
}

type StorageActions interface {
	GetMerchantInfo(arweaveID string) (MerchantDTO, error)
}

func NewMerchantsStorage(url string) *Storage {
	return &Storage{&http.Client{}, url}
}

func (a *Storage) GetMerchantInfo(arweaveID string) (MerchantDTO, error) {
	req, err := http.NewRequest(echo.GET, a.url+arweaveID, nil)

	if err != nil {
		return MerchantDTO{}, err
	}

	res, err := a.httpClient.Do(req)

	if err != nil {
		return MerchantDTO{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println("Error reading response body:", err)
		return MerchantDTO{}, err
	}

	var merchant MerchantDTO

	if err = json.Unmarshal(body, &merchant); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return MerchantDTO{}, err
	}

	log.Printf("Got merchant info for arweaveID: %s\n", arweaveID)

	return merchant, nil
}
