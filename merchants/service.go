package merchants

import (
	"context"
	"log"
	"time"
)

type MerchantsService struct {
	repo       MerchantsRepoActions
	storage    StorageActions
	arweaveIDs <-chan string
}

type Merchant struct {
	MerchantAddress string  `json:"merchantAddress"`
	Name            string  `json:"name"`
	Street          string  `json:"street"`
	Number          string  `json:"number"`
	PostCode        string  `json:"postcode"`
	Country         string  `json:"country"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	Description     string  `json:"description"`
}

type MerchantsServiceActions interface {
	GetMerchants(lat float64, lng float64, radius float64, keyword string) ([]Merchant, error)
}

func NewMerchantsService(repo MerchantsRepoActions, storage StorageActions, arweaveIDs <-chan string) *MerchantsService {
	client := &MerchantsService{repo, storage, arweaveIDs}

	go client.addNewMerchants()

	return client
}

func (s *MerchantsService) GetMerchants(lat float64, lng float64, radius float64, keyword string) ([]Merchant, error) {
	merchantsMongo, err := s.repo.FindBy(lat, lng, radius, keyword)

	var merchants []Merchant = make([]Merchant, 0)

	for _, merchant := range merchantsMongo {
		merchants = append(merchants, Merchant{
			MerchantAddress: merchant.MerchantAddress,
			Name:            merchant.Name,
			Street:          merchant.Street,
			Number:          merchant.Number,
			PostCode:        merchant.PostCode,
			Country:         merchant.Country,
			Lat:             merchant.Location.Coordinates[1],
			Lng:             merchant.Location.Coordinates[0],
			Description:     merchant.Description,
		})
	}

	return merchants, err
}

func (s *MerchantsService) addNewMerchants() {
	arweaveIDsBuffer := []string{}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		select {
		case arweaveID := <-s.arweaveIDs:
			log.Println("Service received arweave ID: ", arweaveID)
			arweaveIDsBuffer = append(arweaveIDsBuffer, arweaveID)
			cancel()
		case <-ctx.Done(): //FIXME: it triggers once and errors upon startup
			if len(arweaveIDsBuffer) == 0 {
				continue
			}
			newArweaveIDs, err := s.repo.GetDifference(arweaveIDsBuffer)
			if err != nil {
				log.Printf("Error getting the arweaveIDs difference: %s\n", err)
			}
			arweaveIDsBuffer = make([]string, 0)
			log.Printf("%d merchants does not exist, adding to repo\n", len(newArweaveIDs))

			merchantDTOs := make(map[string](MerchantDTO), 0)

			for _, arweaveID := range newArweaveIDs {
				merchantDTO, err := s.storage.GetMerchantInfo(arweaveID)

				if err != nil {
					log.Printf("Error getting merchant info: %s\n", err)
				}

				merchantDTOs[arweaveID] = merchantDTO
			}

			merchantMongos := make([]MerchantMongo, 0)

			for arweaveID, merchantDTO := range merchantDTOs {
				merchantMongos = append(merchantMongos, MerchantMongo{
					ArweaveID:       arweaveID,
					Name:            merchantDTO.Name,
					Location:        Location{Type: "Point", Coordinates: []float64{merchantDTO.Longitude, merchantDTO.Latitude}},
					MerchantAddress: merchantDTO.MerchantAddress,
					Street:          merchantDTO.Street,
					Number:          merchantDTO.Number,
					PostCode:        merchantDTO.PostCode,
					Country:         merchantDTO.Country,
					Description:     merchantDTO.Description,
				})
			}

			ex := s.repo.InsertMany(merchantMongos)

			if ex != nil {
				log.Printf("Error inserting merchants: %s\n", ex)
			}

			log.Println("Successfully added new merchants to repo")
		}
	}
}
