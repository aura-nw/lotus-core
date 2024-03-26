package config

type Config struct {
	Server  ServerInfo  `toml:"server"`
	DB      DBInfo      `toml:"db"`
	Bitcoin BitcoinInfo `toml:"bitcoin"`
	Evm     EvmInfo     `toml:"evm"`
	// Address of bitcoin multsig wallet
	BitcoinMultisig string `toml:"bitcoin-multisig"`
	BridgeInterval  int64  `toml:"bridge-interval"`
	RedeemScript    string `toml:"redeem-script"`
	PrivateKey      string `toml:"private-key"`
}

type EvmInfo struct {
	Url           string       `toml:"url"`
	ChainID       int64        `toml:"chain-id"`
	QueryInterval int64        `toml:"query-interval"`
	MinConfirms   int64        `toml:"min-confirms"`
	Contracts     EvmContracts `toml:"contracts"`
	PrivateKey    string       `toml:"private-key"`
	CallTimeout   uint64       `toml:"call-timeout"`
}

type EvmContracts struct {
	WrappedBtcAddr string `toml:"wrapped-btc-addr"`
	GatewayAddr    string `toml:"gateway-addr"`
}

type BitcoinInfo struct {
	Network       string `toml:"network"`
	Host          string `toml:"host"`
	User          string `toml:"user"`
	Password      string `toml:"password"`
	QueryInterval int64  `toml:"query-interval"`
	MinConfirms   int64  `toml:"min-confirms"`
}
type ServerInfo struct {
	Port string `toml:"port"`
}

func GetBitcoinInfoTestnet() BitcoinInfo {
	return BitcoinInfo{
		Network:       "testnet3",
		Host:          "103.154.187.179:18332",
		User:          "user",
		Password:      "123@abcD",
		MinConfirms:   3,
		QueryInterval: 60,
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
		Name:     "btc-bridge",
		User:     "aurauser",
		Password: "aurapassword",
	}
}
