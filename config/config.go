package config

type Config struct {
	Server         ServerInfo  `toml:"server"`
	DB             DBInfo      `toml:"db"`
	Bitcoin        BitcoinInfo `toml:"bitcoin"`
	Evm            EvmInfo     `toml:"evm"`
	BridgeInterval uint32      `toml:"bridge-interval"`
}

type EvmInfo struct {
	Url           string       `toml:"url"`
	ChainID       int64        `toml:"chain-id"`
	QueryInterval uint32       `toml:"query-interval"`
	MinConfirms   int64        `toml:"min-confirms"`
	Contracts     EvmContracts `toml:"contracts"`
	PrivateKey    string       `toml:"private-key"`
	CallTimeout   uint32       `toml:"call-timeout"`
}

type EvmContracts struct {
	WrappedBtcAddr string `toml:"wrapped-btc-addr"`
	GatewayAddr    string `toml:"gateway-addr"`
}

type BitcoinInfo struct {
	Network         string `toml:"network"`
	Host            string `toml:"host"`
	User            string `toml:"user"`
	Password        string `toml:"password"`
	QueryInterval   uint32 `toml:"query-interval"`
	MinConfirms     int64  `toml:"min-confirms"`
	BitcoinMultisig string `toml:"bitcoin-multisig"`
	RedeemScript    string `toml:"redeem-script"`
	PrivateKey      string `toml:"private-key"`
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

func GetEvmInfoForTest() EvmInfo {
	return EvmInfo{
		Url:           "https://jsonrpc.dev.aura.network",
		ChainID:       1235,
		QueryInterval: 6,
		MinConfirms:   1,
		Contracts: EvmContracts{
			WrappedBtcAddr: "0xe10Ce7eAb22643a24CEeaAc30e38F89E8bfeDC8e",
			GatewayAddr:    "0x7316e5511784EA3A2735c0Fe8a99324231E49030",
		},
		PrivateKey:  "7cf06d80baec389b7c63daff9eccfae043a2068a2c37a3d1e8be8976eec8928b",
		CallTimeout: 10,
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
