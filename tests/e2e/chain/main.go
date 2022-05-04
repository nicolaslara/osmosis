package chain

func Init(id, dataDir string, numVal int) (*Chain, error) {
	chain, err := new(id, dataDir)
	if err != nil {
		return nil, err
	}
	if err := initNodes(chain, numVal); err != nil {
		return nil, err
	}
	if err := initGenesis(chain); err != nil {
		return nil, err
	}
	if err := initValidatorConfigs(chain); err != nil {
		return nil, err
	}
	return chain.export(), nil
}
