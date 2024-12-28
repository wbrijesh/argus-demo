package main

import (
	"log/slog"
	"time"

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

	defer func() {
		if r := recover(); r != nil {
			argusHandler.Flush()
		}
	}()

	logger.Info("Info Log using slog handler using buffers")
	logger.Error("Error Log using slog handler using buffers")

	time.Sleep(3 * time.Second)

	logger.Debug("Debug Log using slog handler using buffers")
	logger.Warn("Warn Log using slog handler using buffers")

	panic("simulated panic")
}
