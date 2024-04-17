#!/usr/bin/env bash

rm -rf ./clients/evm/contracts/*_abi.go

abigen -abi ./clients/evm/abis/AuraWrappedBitcoin.json -pkg contracts -type AuraWrappedBitcoin -out ./clients/evm/contracts/aura_wrapped_bitcoin_abi.go
abigen -abi ./clients/evm/abis/AuraWrappedBrc20.json -pkg contracts -type AuraWrappedBrc20 -out ./clients/evm/contracts/aura_wrapped_brc20_abi.go
abigen -abi ./clients/evm/abis/Gateway.json -pkg contracts -type Gateway -out ./clients/evm/contracts/gateway_abi.go
abigen -abi ./clients/evm/abis/Brc20Gateway.json -pkg contracts -type Brc20Gateway -out ./clients/evm/contracts/brc20_gateway_abi.go
