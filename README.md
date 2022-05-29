# Project Document

The project is a lighter version of the chainlink integrations framework and is heavily inspired by its design and choice of libraries.

## About the tests

- Tests are fetching BTC/USD price feed data from a deployed smart contract at location `0xF570deEffF684D964dc3E15E1F9414283E3f7419` (address can be parameterized)
- Using infura to connect to etherium mainnet
- number of rounds, contract address, deviation percentage is parameterized tests
- test obtains price feed data from parameterized number of rounds, calculates median and verifies if the deviation is range bound. Range defined by parameterized deviation percentage.
- median calculation is happening off chain as the deployed contract did not expose any method to calculate median in place

## Project Stack

- Golang as programming language
- Ginkgo and Gomega as test framework and assertions
- zero log for logging
- viper for config management
- infura to connect to mainnet
- solana 6.12 to generate go smart contract binding

## Running Tests

- All project dependencies are defined in `go.mod` file
- Infura project key is added in the network yaml file
- run ginkgo test

```
cd tests && ginkgo
```

> The last round was fetched on Dec 8, 2021 and there is no feed data after that.
> The test is fetching the latest round id and working backwards to find previous data.
> Both round data and deviation percentage is parameterized.


