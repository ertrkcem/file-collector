package main

import (
	"fmt"
	"os"

	"file-collector/config"
	ftpclient "file-collector/utils"
)

func main() {
	// 1. Load configurations from a JSON file.
	configs, err := config.LoadConfigurations("config/config.json")
	if err != nil {
		fmt.Println("Error loading configurations:", err)
		os.Exit(1)
	}

	// 2. Loop through loaded configurations and process each server.
	for _, conf := range configs {
		// 3. Based on server_type, establish connection (implement logic for FTP/SFTP).
		if conf.ServerType == "ftp" {
			conn, err := ftpclient.EstablishConnection(conf)
			if err != nil {
				fmt.Printf("Error connecting to server %s: %v\n", conf.Host, conf.ServerType)
				continue
			}

			// quit connection at the end
			conn.Quit()
		} else {
			fmt.Printf("Skipping")
		}

		// 4. Download files from the specified path.
		// ... (code to download files from the server)

		// 5. Loop through each target API in the config.
		// for _, target := range conf.Targets {
		// 	// 6. Upload downloaded files to the target API.
		// 	// ... (code to upload files to the API)
		// }
	}

	fmt.Println("File collection completed successfully!")
}
