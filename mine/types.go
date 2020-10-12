package mine

type coinType string

const (
	FIL coinType = "filecoin"
	ETH coinType = "ethereum"
	BTC coinType = "bitcoin"
)

type FiltabMineInfo struct {
}

type FiltabMineBalance struct {
	coinType
}

type Option struct {
}
