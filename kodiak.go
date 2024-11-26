package protocols

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Protocol = (*KodiakLPPriceProvider)(nil)

type KodiakLPPriceProvider struct {
	address common.Address
	abiPath string
}

func NewKodiakLPPriceProvider(address common.Address, abiPath string) {
	return &KodiakLPPriceProvider{
		address: address,
		abiPath: abiPath,
	}
}

func (k *KodiakLPPriceProvider) Initailize(ctx context.Context, client *ethclient.Client) error {
	return nil
}

func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context, client *ethclient.Client) (uint64, error) {
	return 0, nil
}

func (k *KodiakLPPriceProvider) TVL(ctx context.Context, client *ethclient.Client) (uint64, error) {
	return 0, nil
}
