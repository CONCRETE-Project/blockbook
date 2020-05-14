package concretecoin

import (
	"github.com/CONCRETE-Project/blockbook/bchain"
	"github.com/CONCRETE-Project/blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xbd6b0cbf
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{55}
	MainNetParams.ScriptHashAddrID = []byte{56}
}

type ConcreteCoinParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewConcreteCoinParser(params *chaincfg.Params, c *btc.Configuration) *ConcreteCoinParser {
	return &ConcreteCoinParser{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

func (p *ConcreteCoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

func (p *ConcreteCoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
