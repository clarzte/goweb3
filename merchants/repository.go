package merchants

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MerchantsRepository struct {
	merchants *mongo.Collection
}

type MerchantMongo struct {
	MerchantAddress string   `bson:"merchantAddress"`
	Name            string   `bson:"name"`
	Street          string   `bson:"street"`
	Number          string   `bson:"number"`
	PostCode        string   `bson:"postcode"`
	Country         string   `bson:"country"`
	Location        Location `bson:"location"`
	Description     string   `bson:"description"`
	ArweaveID       string   `bson:"arweaveID"`
}

type Location struct {
	Type        string    `bson:"type" default:"Point"`
	Coordinates []float64 `bson:"coordinates"`
}

type MerchantsRepoActions interface {
	FindAll() ([]MerchantMongo, error)
	FindBy(lat float64, lng float64, radius float64, keyword string) ([]MerchantMongo, error)
	FindOne(id string) (*MerchantMongo, error)
	InsertOne(merchant MerchantMongo) error
	InsertMany(merchants []MerchantMongo) error
	GetDifference(arweaveIDs []string) ([]string, error)
}

func NewMerchantsRepository(uri string) *MerchantsRepository {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatalf("Failed to connect to the MongoDB client: %v", err)
	}

	return &MerchantsRepository{
		merchants: client.Database("merchants").Collection("merchants"),
	}
}

func (r *MerchantsRepository) FindAll() ([]MerchantMongo, error) {
	var merchants []MerchantMongo

	cursor, err := r.merchants.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var merchant MerchantMongo
		if err := cursor.Decode(&merchant); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func (r *MerchantsRepository) FindBy(lat float64, lng float64, radius float64, keyword string) ([]MerchantMongo, error) {
	var merchants []MerchantMongo

	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
				"$maxDistance": radius,
			},
		},
	}

	if keyword != "" {
		filter["$or"] = []bson.M{
			{"name": bson.M{"$regex": keyword, "$options": "i"}},
			{"description": bson.M{"$regex": keyword, "$options": "i"}},
		}
	}

	cursor, err := r.merchants.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var merchant MerchantMongo
		if err := cursor.Decode(&merchant); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func (r *MerchantsRepository) FindOne(id string) (*MerchantMongo, error) {
	var merchant MerchantMongo

	if err := r.merchants.FindOne(context.Background(), bson.M{"_id": id}).Decode(&merchant); err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *MerchantsRepository) InsertMany(merchants []MerchantMongo) error {
	var documents []interface{}

	for _, merchant := range merchants {
		documents = append(documents, merchant)
	}

	_, err := r.merchants.InsertMany(context.Background(), documents)

	if err != nil {
		return err
	}

	log.Printf("Inserted %d merchants to db'\n", len(merchants))

	return nil
}

func (r *MerchantsRepository) GetDifference(arweaveIDs []string) ([]string, error) { //checks for existing arweaveIDs in the db,
	filter := bson.M{ //returns the ones that are not in the db
		"arweaveID": bson.M{
			"$in": arweaveIDs,
		},
	}

	type ArweaveID struct {
		ArweaveID string `bson:"arweaveID"`
	}

	cursor, err := r.merchants.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var existingArweaveIDs []string
	for cursor.Next(context.Background()) {
		var result ArweaveID
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
			return nil, err
		}
		existingArweaveIDs = append(existingArweaveIDs, result.ArweaveID)
	}

	return getDifference(arweaveIDs, existingArweaveIDs), nil
}

func getDifference(slice1, slice2 []string) []string {
	diff := make([]string, 0)

	lookup := make(map[string]bool)
	for _, item := range slice2 {
		lookup[item] = true
	}

	for _, item := range slice1 {
		if _, exists := lookup[item]; !exists {
			diff = append(diff, item)
		}
	}

	return diff
}

func (r *MerchantsRepository) InsertOne(merchant MerchantMongo) error {
	_, err := r.merchants.InsertOne(context.Background(), merchant)

	if err != nil {
		return err
	}

	return nil
}
