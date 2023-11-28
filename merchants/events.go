package merchants

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MerchantsEventHandler struct {
	instance   *MerchantRegistryFilterer
	ArweaveIDs <-chan string
}

var arweaveIDs chan string = make(chan string, 3) //TODO: play with pointers?
var blockNumber uint64 = uint64(9510955)          //TODO: make it configurable?

func NewMerchantsEventHandler(rpcUrl string) *MerchantsEventHandler {

	ethClient, err := ethclient.Dial(rpcUrl)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	address := common.HexToAddress("0x3FB3633b64fbe861e6Ee9Cb07Db07597278A1587")
	instance, err := NewMerchantRegistryFilterer(address, ethClient)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		filterOpts := &bind.FilterOpts{Start: blockNumber, End: nil, Context: context.Background()}
		itr, err := instance.FilterNewMerchantRegistered(filterOpts, nil)
		if err != nil {
			log.Fatal(err)
		}

		for itr.Next() {
			event := itr.Event
			log.Println("Got event with merchant id", event.MerchantId)
			arweaveIDs <- event.ArweaveID
			blockNumber = event.Raw.BlockNumber
		}

		go func() {
			sink := make(chan *MerchantRegistryNewMerchantRegistered, 3)

			watchOpts := &bind.WatchOpts{Start: &blockNumber, Context: context.Background()}
			sub, err := instance.WatchNewMerchantRegistered(watchOpts, sink, nil)

			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Listening for new events from block: %d\n", blockNumber)

			for {
				select {
				case err := <-sub.Err():
					log.Fatal(err)
				case event := <-sink:
					log.Printf("Got new event with merchant id: %s", event.MerchantId)
					arweaveIDs <- event.ArweaveID
				}
			}
		}()
	}()

	return &MerchantsEventHandler{instance, arweaveIDs}
}
