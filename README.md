# go-sesa

Golang implementation of sesanetwork Distributed Network followed by this [Whitepaper](https://sesa.network/docs/sesanetworkWhitepaper.pdf) based on various open-source and documentations:
- [Hashgraph](https://arxiv.org/pdf/1907.02900.pdf)
- [Hashgraph Sharding](https://www.mdpi.com/2076-3417/13/15/8726)
- [TEE Directed Acyclic Graph](https://www.mdpi.com/2079-9292/12/11/2393)
- [Lachesis](https://arxiv.org/abs/2108.01900)
- [Delegated Proof-of-Stake](https://www.mdpi.com/1099-4300/25/9/1320)
- [Proof-of-Elapsed-Time](https://ieeexplore.ieee.org/document/9472787)
- [Ethereum](https://github.com/ethereum/go-ethereum)
- [Hedera](https://github.com/hashgraph/hedera-services)
- [Opera](https://github.com/Fantom-foundation/go-opera)
- [Erigon](https://github.com/ledgerwatch/erigon)

## Building the source

Building `sesa` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make sesa
```
The build output is ```build/sesa``` executable.

## Running `sesa`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `sesa` instance.

### Launching a network

Launching `sesa` readonly (non-validator) node for network specified by the genesis file:

```shell
$ sesa --genesis file.g
```

### Configuration

As an alternative to passing the numerous flags to the `sesa` binary, you can also pass a
configuration file via:

```shell
$ sesa --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ sesa --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `sesa validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ sesa --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`sesa` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ sesa --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ sesa --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ sesa --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ sesa --datadir /path/to/datadir account new # create new account
$ sesa --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Client has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/sesanetwork/go-sesa/app	0.033s
?   	github.com/sesanetwork/go-sesa/cmd/cmdtest	[no test files]
ok  	github.com/sesanetwork/go-sesa/cmd/sesa	13.890s
?   	github.com/sesanetwork/go-sesa/cmd/sesa/metrics	[no test files]
?   	github.com/sesanetwork/go-sesa/cmd/sesa/tracing	[no test files]
?   	github.com/sesanetwork/go-sesa/crypto	[no test files]
?   	github.com/sesanetwork/go-sesa/debug	[no test files]
?   	github.com/sesanetwork/go-sesa/ethapi	[no test files]
?   	github.com/sesanetwork/go-sesa/eventcheck	[no test files]
?   	github.com/sesanetwork/go-sesa/eventcheck/basiccheck	[no test files]
?   	github.com/sesanetwork/go-sesa/eventcheck/gaspowercheck	[no test files]
?   	github.com/sesanetwork/go-sesa/eventcheck/heavycheck	[no test files]
?   	github.com/sesanetwork/go-sesa/eventcheck/parentscheck	[no test files]
ok  	github.com/sesanetwork/go-sesa/evmcore	6.322s
?   	github.com/sesanetwork/go-sesa/gossip	[no test files]
?   	github.com/sesanetwork/go-sesa/gossip/emitter	[no test files]
ok  	github.com/sesanetwork/go-sesa/gossip/filters	1.250s
?   	github.com/sesanetwork/go-sesa/gossip/gasprice	[no test files]
?   	github.com/sesanetwork/go-sesa/gossip/occuredtxs	[no test files]
?   	github.com/sesanetwork/go-sesa/gossip/piecefunc	[no test files]
ok  	github.com/sesanetwork/go-sesa/integration	21.640s
```

Also it is tested with [fuzzing](./FUZZING.md).


### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ sesa --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ sesa --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ sesa --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ sesa --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the sesa processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
Check README.md in the demo directory for more information.
