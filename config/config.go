package config

type BridgeConfig struct {
	DBInfo      *DBInfo      `json:"db_info"`
	BitcoinInfo *BitcoinInfo `json:"bitcoin_info"`
	EvmInfo     *EvmInfo     `json:"evm_info"`
}

type OperatorConfig struct {
	BitcoinInfo *BitcoinInfo `json:"bitcoin_info"`
	EvmInfo     *EvmInfo     `json:"evm_info"`
}

type EvmInfo struct {
	Url     string
	ChainId string
}

type BitcoinInfo struct {
	Network  string
	Host     string
	User     string
	Password string
}

func GetBitcoinInfoTestnet() BitcoinInfo {
	return BitcoinInfo{
		Network:  "testnet3",
		Host:     "103.154.187.179:18332",
		User:     "user",
		Password: "123@abcD",
	}
}

// DBInfo configures the postgres database
type DBInfo struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func DBConfigForTest() DBInfo {
	return DBInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "bitcoin-bridge",
		User:     "aurauser",
		Password: "aurapassword",
	}
}
