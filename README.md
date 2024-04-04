# Lotus Core

Bridge bitcoin, ordinals, brc-20,... to Aura Network.

Lotus-core is a bridge designed to seamlessly connect various blockchain assets with the Aura Network. This opens a world of possibilities for users who want to leverage the scalability, agility, and user-friendliness of Aura network for their Bitcoin (including Ordinals), BRC-20 tokens, and potentially more in the future.

## 1. Modules

* `bridge`
* `config`
* `clients`
* `database`
* `types`
* `utils`

## 2. Config

The bridge service need a config file located at `./bridge.toml` for default. Here is describe of fields in config:

a. General:

* `bridge-interval`: This line sets the interval (in seconds) at which the bridge checks for new transactions on the source chain (likely Bitcoin) to be transferred to Aura Network.

b. Server:

* `port`: This defines the port number on which the bridge server listens for incoming connections.

c. Database:

* `host`: The IP address or hostname of the database server.
* `port`:  The port number used to connect to the database.
* `name`: The name of the database used by lotus-core.
* `user`: The username for accessing the database.
* `password`: The password for the database user.

d. Bitcoin:

* `network`: Specifies the Bitcoin network to connect to (likely "testnet3" for a test network in this case).

* `host`: The IP address and port of the Bitcoin node used for communication.
* `user`: The username for authentication with the Bitcoin node (if required).
* `password`: The password for the Bitcoin node user (if required).
* `query-interval`: The interval (in seconds) at which the bridge queries the Bitcoin node for new transactions.
* `min-confirms`: The minimum number of confirmations required for a Bitcoin transaction before it's considered for bridging.
* `bitcoin-multisig`: The multisignature address used for Bitcoin transactions on the bridge.
* `private-key`: The private key associated with the bridge's multisignature address (likely obfuscated for security reasons).
* `redeem-script`: The redeem script for the multisignature address (likely obfuscated).

e. Evm:

* `url`: The URL of the Aura Network JSON RPC endpoint for communication.
* `chain-id`: The chain ID of the Aura Network used by the bridge.
* `query-interval`: The interval (in seconds) at which the bridge queries Aura Network for transaction confirmations.
* `min-confirms`: The minimum number of confirmations required for an Aura Network transaction before it's considered finalized.
* `private-key`: The private key used by the bridge for signing transactions on Aura Network (likely obfuscated).
* `call-timeout`: The timeout value (in seconds) for making calls to the Aura Network JSON RPC endpoint.

## 3. Run

After editing config properly. Run the service using command:

```bash
make run
```
