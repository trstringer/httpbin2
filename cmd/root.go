/*
Copyright © 2023 Thomas Stringer <thomas@trstringer.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/trstringer/httpbin2/server"
)

var (
	port int

	message         string
	messageHostName bool

	statusCode     int
	statusCodeRate int
	delayMin       int
	delayMax       int
	delayRate      int

	relays []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpbin2",
	Short: "HTTP server for testing, demo'ing, and learning",
	Run: func(cmd *cobra.Command, args []string) {
		serverMessage, err := generateServerMessage(message, messageHostName)
		if err != nil {
			fmt.Printf("Error getting server message: %v\n", err)
			os.Exit(1)
		}
		svr := server.New(
			server.WithPort(port),
			server.WithQOS(statusCode, statusCodeRate, delayMin, delayMax, delayRate),
			server.WithMessage(serverMessage),
			server.WithRelays(relays),
		)
		svr.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Base server config.
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Listening port")

	// Message returning config.
	rootCmd.Flags().StringVar(&message, "message", "", "Message for server to return")
	rootCmd.Flags().BoolVar(&messageHostName, "message-hostname", false, "Return hostname as message to requests")

	// Mimic poor quality of service config.
	rootCmd.Flags().IntVar(&statusCode, "status-code", 200, "HTTP status  code to return (default 200)")
	rootCmd.Flags().IntVar(&statusCodeRate, "status-code-rate", 100, "Status code rate (default 100)")
	rootCmd.Flags().IntVar(&delayMin, "delay-min-ms", 0, "Low end for response delay (default 0)")
	rootCmd.Flags().IntVar(&delayMax, "delay-max-ms", 0, "High end for response delay (default 0)")
	rootCmd.Flags().IntVar(&delayRate, "delay-rate", 100, "Delay rate (default 100)")

	// Request relay trigger.
	rootCmd.Flags().StringSliceVar(&relays, "relay", []string{}, "Relay routes to request on trigger (can be specified multiple times)")
}

func generateServerMessage(message string, addHostname bool) (string, error) {
	output := message
	if addHostname {
		hostname, err := os.Hostname()
		if err != nil {
			return "", fmt.Errorf("error getting hostname: %w", err)
		}
		output = fmt.Sprintf("(%s) %s", hostname, message)
	}
	return output, nil
}
