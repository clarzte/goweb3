# Web3 pet project using Go

## Description
This is a pet project to learn Go and Web3.
It's a simple backend application which listens on new registration events on the smart contract and then downloads the merchants data from [arweave](https://www.arweave.org/) and puts them into MongoDB which acts as a cache. MongoDB then allows for easy querying of the data using geolocation.

## Stack
- Go
- MongoDB
- Docker
- arweave
- echo
- viper
- go-ethereum
