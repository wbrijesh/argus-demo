package main

import (
	"log/slog"

	"github.com/wbrijesh/argus_client"
)

func main() {
	client := argus_client.NewClient(
		argus_client.ClientConfig{
			ApiKey:  "argus_jW35MW2m-dNdtSbhFVSrudlSIpQCA99apo_2mRvdbjU",
			BaseUrl: "http://localhost:8080",
		},
	)

	argusHandler := argus_client.NewArgusHandler(client)
	logger := slog.New(argusHandler)

	logger.Info("Info Log using slog handler")
	logger.Error("Error Log using slog handler")
	logger.Debug("Debug Log using slog handler")
	logger.Warn("Warn Log using slog handler")
}
