package tests

import (
	"chainlink-trial/contract"
	"chainlink-trial/environment"
	"chainlink-trial/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"math/big"
)

var _ = Describe("price feed test", func() {
	var (
		client       *ethclient.Client
		feed         []float64
		address      common.Address
		rounds       int
		fluxInstance *contract.FluxAggregator
		err          error
	)
	DescribeTable("Verify aggregated median is range bound",
		func(network string, contractAddress string, numberOfRounds int, deviation float64) {
			// set up
			address = common.HexToAddress(contractAddress)
			rounds = numberOfRounds
			feed = make([]float64, rounds)

			// connect to etherium mainnet
			client, err = environment.ConnectToMainNet(network)
			Expect(err).ShouldNot(HaveOccurred(), "Connection to mainnet failed")

			// get instance of deployed smart contract
			fluxInstance, err = contract.NewFluxAggregator(address, client)
			Expect(err).ShouldNot(HaveOccurred(), "Failed to create an instance of price feed contract")

			// aggregate round data
			round, err := fluxInstance.LatestRound(&bind.CallOpts{})
			Expect(err).ShouldNot(HaveOccurred(), "Failed to read the latest round id")
			for i := 0; i < rounds; i++ {
				r := big.NewInt(int64(i))
				data, err := fluxInstance.GetRoundData(&bind.CallOpts{}, big.NewInt(0).Sub(round, r))
				log.Info().Int("round", i).Str("feed", data.Answer.String()).Msg("")
				Expect(err).ShouldNot(HaveOccurred(), "Unable to get the round data for round id")
				feed[i], _ = util.ToNormalizedFloat(data.Answer, 8)
			}
			// compare individual price with median and assert if it is range bound
			median := util.Median(feed)
			log.Info().Float64("median", median).Msg(fmt.Sprintf("median for total %v rounds", rounds))
			result := util.MeasureDeviation(feed, median, deviation)
			Expect(result).ShouldNot(BeFalse(), "Deviation has crossed boundary")
		},
		Entry(nil, "infura", "0xF570deEffF684D964dc3E15E1F9414283E3f7419", 5, float64(10)),
		Entry(nil, "infura", "0xF570deEffF684D964dc3E15E1F9414283E3f7419", 10, float64(15)),
		Entry(nil, "infura", "0xF570deEffF684D964dc3E15E1F9414283E3f7419", 3, float64(20)),
	)
})
