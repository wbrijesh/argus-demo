package main

import "github.com/wbrijesh/argus_client"

func main() {
	client := argus_client.NewClient(
		argus_client.ClientConfig{
			ApiKey:  "argus_dSHLPT5hZfs3-7G7R8X8aO9WtMhaTtg52Nnn96RQf7s",
			BaseUrl: "http://localhost:8080",
		},
	)

	logs := []argus_client.LogEntry{
		{Level: argus_client.LevelInfo, Message: "From Demo App: This is an info message"},
		{Level: argus_client.LevelError, Message: "From Demo App: This is an error message"},
		{Level: argus_client.LevelDebug, Message: "From Demo App: This is a debug message"},
		{Level: argus_client.LevelWarn, Message: "From Demo App: This is a warning message"},
		{Level: argus_client.LevelFatal, Message: "From Demo App: This is a fatal message"},
	}

	err := client.SendLogs(logs)

	if err != nil {
		panic(err)
	}
}
