package environment

import (
	"chainlink-trial/config"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

func ConnectToMainNet(provider string) (*ethclient.Client, error) {
	log.Info().Msg(fmt.Sprintf("connecting to etherium mainnet through %v", provider))
	conf, err := config.LoadConfig(provider)
	if err != nil {
		log.Error().Err(err).Msg("unable to load network config")
		return nil, err
	}
	client, err := ethclient.Dial(conf.Url + conf.Key)
	if err != nil {
		log.Error().Err(err).Msg("failed to establish connection to mainnet")
		return nil, err
	}
	log.Info().Msg("established connection to mainnet")
	return client, nil
}
