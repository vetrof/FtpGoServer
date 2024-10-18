package main

import (
	"log"
	"os"

	"github.com/goftp/file-driver"
	"github.com/goftp/server"
)

func main() {
	// Create a directory for the FTP server if it doesn't exist
	rootPath := "ftp"
	log.Println("Checking if root directory exists...")
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		log.Fatalf("Failed to create root directory: %v", err)
	}
	log.Println("Root directory is set to:", rootPath)

	// Set up a simple file driver
	log.Println("Setting up file driver...")
	factory := &filedriver.FileDriverFactory{
		RootPath: rootPath,
	}
	log.Println("File driver set up with root path:", rootPath)

	// Set up simple authentication
	log.Println("Setting up authentication...")
	auth := &server.SimpleAuth{
		Name:     "user",
		Password: "12345",
	}
	log.Println("Authentication set up with user:", auth.Name)

	// Create FTP server options with debug logging enabled
	log.Println("Creating FTP server options...")
	opts := &server.ServerOpts{
		Factory:  factory,
		Auth:     auth,
		Port:     2121,
		Hostname: "0.0.0.0",
		Logger:   new(server.StdLogger), // Enable debug logging
	}
	log.Println("FTP server options created with port:", opts.Port)

	// Create the FTP server
	log.Println("Creating FTP server...")
	ftpServer := server.NewServer(opts)

	// Start the FTP server
	log.Println("Starting FTP server on port 2121...")
	if err := ftpServer.ListenAndServe(); err != nil {
		log.Fatal("Error starting FTP server:", err)
	}
}
