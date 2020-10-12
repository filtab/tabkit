package mine

type FiltabMineClient struct {
}

func (f *FiltabMineClient) Info() (FiltabMineInfo, error) {
	return FiltabMineInfo{}, nil
}

func (f *FiltabMineClient) GetBalance(coinType coinType) (FiltabMineBalance, error) {
	return FiltabMineBalance{}, nil
}

func (f *FiltabMineClient) WithdrawTx(coinType coinType, volume uint64, address string, options ...Option) (FiltabMineInfo, error) {
	return FiltabMineInfo{}, nil
}
