package snapshothashes

import (
	_ "embed"
)

//go:embed mainnet.toml
var Mainnet []byte

//go:embed goerli.toml
var Goerli []byte

//go:embed sepolia.toml
var Sepolia []byte

//go:embed mumbai.toml
var Mumbai []byte

//go:embed amoy.toml
var Amoy []byte

//go:embed bor-mainnet.toml
var BorMainnet []byte

//go:embed gnosis.toml
var Gnosis []byte

//go:embed chiado.toml
var Chiado []byte

//go:embed bsc.toml
var Bsc []byte

//go:embed chapel.toml
var Chapel []byte

//go:embed history/bsc.toml
var BscHistory []byte

//go:embed history/chapel.toml
var ChapelHistory []byte
