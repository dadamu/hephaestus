module github.com/desmos-labs/discord-bot

go 1.15

require (
	github.com/andersfylling/disgord v0.26.1
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/desmos-labs/desmos v0.16.0
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/pelletier/go-toml v1.8.1
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
