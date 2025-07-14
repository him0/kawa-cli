package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/him0/kawa-cli/internal/config"
	"github.com/him0/kawa-cli/internal/display"
	"github.com/him0/kawa-cli/internal/fetcher"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	disp, err := display.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Please install imgcat: brew install eddieantonio/eddieantonio/imgcat\n")
		os.Exit(1)
	}

	fetch := fetcher.New()

	if cfg.Live {
		runLiveMode(cfg, fetch, disp)
	} else {
		runOnceMode(cfg, fetch, disp)
	}
}

func runOnceMode(cfg *config.Config, fetch *fetcher.ImageFetcher, disp *display.ImageDisplay) {
	imageData, err := fetch.Fetch(cfg.URL)
	if err != nil {
		log.Fatalf("Failed to fetch image: %v", err)
	}

	if err := disp.DisplayWithSize(imageData, cfg.Width); err != nil {
		log.Fatalf("Failed to display image: %v", err)
	}
}

func runLiveMode(cfg *config.Config, fetch *fetcher.ImageFetcher, disp *display.ImageDisplay) {
	fmt.Printf("Starting live mode (interval: %.1fs). Press Ctrl+C to stop.\n", cfg.Interval)
	time.Sleep(1 * time.Second)

	// Clear screen for live mode
	disp.ClearScreen()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Channel for image data
	imageChan := make(chan []byte, 1)
	errorChan := make(chan error, 1)

	// WaitGroup for goroutines
	var wg sync.WaitGroup

	// Ticker for updates
	ticker := time.NewTicker(time.Duration(cfg.Interval * float64(time.Second)))
	defer ticker.Stop()

	// Flag to signal shutdown
	shutdown := make(chan struct{})

	// Fetch first image immediately
	wg.Add(1)
	go func() {
		defer wg.Done()
		imageData, err := fetch.Fetch(cfg.URL)
		if err != nil {
			errorChan <- err
			return
		}
		imageChan <- imageData
	}()

	// Main loop
	for {
		select {
		case <-sigChan:
			fmt.Println("\nShutting down...")
			close(shutdown)
			wg.Wait()
			return

		case imageData := <-imageChan:
			disp.MoveCursorHome()
			if err := disp.DisplayWithSize(imageData, cfg.Width); err != nil {
				log.Printf("Failed to display image: %v", err)
			}

		case err := <-errorChan:
			log.Printf("Failed to fetch image: %v", err)

		case <-ticker.C:
			// Start fetching next image
			select {
			case <-shutdown:
				return
			default:
				wg.Add(1)
				go func() {
					defer wg.Done()
					imageData, err := fetch.Fetch(cfg.URL)
					if err != nil {
						select {
						case errorChan <- err:
						case <-shutdown:
						}
						return
					}
					select {
					case imageChan <- imageData:
					case <-shutdown:
					}
				}()
			}
		}
	}
}