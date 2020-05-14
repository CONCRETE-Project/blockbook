package concretecoin

import (
	"encoding/json"

	"github.com/CONCRETE-Project/blockbook/bchain"
	"github.com/CONCRETE-Project/blockbook/bchain/coins/btc"
	"github.com/golang/glog"
)

type ConcreteCoinRPC struct {
	*btc.BitcoinRPC
}

func NewConcreteCoinRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}
	s := &ConcreteCoinRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateFee = true
	s.ChainConfig.SupportsEstimateSmartFee = false
	return s, nil
}

func (b *ConcreteCoinRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain
	params := GetChainParams(chainName)
	b.Parser = NewConcreteCoinParser(params, b.ChainConfig)
	b.Testnet = false
	b.Network = "livenet"
	glog.Info("rpc: block chain ", params.Name)
	return nil
}
