package cmd

import (
	"github.com/desmos-labs/desmos/v2/app"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/hephaestus/network"

	"github.com/desmos-labs/hephaestus/bot"
	"github.com/desmos-labs/hephaestus/types"
)

// StartCmd returns a Cobra command allowing to start the bot
func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start [config-file]",
		Short: "Starts the bot using the provided configuration file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read the configuration
			cfg, err := types.Parse(args[0])
			if err != nil {
				return err
			}

			encodingConfig := app.MakeTestEncodingConfig()

			// Build the networks clients
			var mainnet, testnet *network.Client
			if cfg.Networks.Testnet != nil {
				testnet, err = network.NewClient(cfg.Networks.Testnet, encodingConfig)
				if err != nil {
					return err
				}
			}
			if cfg.Networks.Mainnet != nil {
				mainnet, err = network.NewClient(cfg.Networks.Mainnet, encodingConfig)
				if err != nil {
					return err
				}
			}

			// Create the bot
			hephaestus, err := bot.Create(cfg.BotConfig, testnet, mainnet)
			if err != nil {
				return err
			}

			// Start the bot
			hephaestus.Start()

			return nil
		},
	}
}
