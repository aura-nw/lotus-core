package config

type BridgeConfig struct {
	Server  BridgeServer `toml:"server"`
	DB      DBInfo       `toml:"db"`
	Bitcoin BitcoinInfo  `toml:"bitcoin"`
	Evm     EvmInfo      `toml:"evm"`
	// List addresses of operators on EVM
	EvmOperators []string `toml:"evm-operators"`
	// Address of bitcoin multsig wallet
	BitcoinMultisig string `toml:"bitcoin-multisig"`
}

type OperatorConfig struct {
	Server    OperatorServer `toml:"server"`
	BridgeUrl string         `toml:"bridge-url"`
	Bitcoin   BitcoinInfo    `toml:"bitcoin"`
	Evm       EvmInfo        `toml:"evm"`
}

type EvmInfo struct {
	Url     string `toml:"url"`
	ChainID int64  `toml:"chain-id"`
}

type BitcoinInfo struct {
	Network  string `toml:"network"`
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type OperatorServer struct {
	GrpcPort string `toml:"grpc-port"`
}

type BridgeServer struct {
	Port string `toml:"port"`
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
