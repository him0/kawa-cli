package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Live     bool
	Interval float64
	URL      string
	Width    string
}

func Parse() (*Config, error) {
	cfg := &Config{
		URL:   "https://cam.wni.co.jp/taikobashi/camera.jpg",
		Width: "80%",
	}

	flag.BoolVar(&cfg.Live, "live", false, "Enable live mode")
	flag.BoolVar(&cfg.Live, "l", false, "Enable live mode (shorthand)")
	flag.Float64Var(&cfg.Interval, "interval", 60.0, "Update interval in seconds (live mode only)")
	flag.Float64Var(&cfg.Interval, "i", 60.0, "Update interval in seconds (shorthand)")
	flag.StringVar(&cfg.Width, "width", "80%", "Image width (e.g., 80%, 100px, auto)")
	flag.StringVar(&cfg.Width, "w", "80%", "Image width (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nDisplay live camera image from Meguro River\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s                    # Display image once\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --live             # Live mode (60 second interval)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -l -i 30           # Live mode (30 second interval)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 100%%           # Full width display\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 50%%            # Half width display\n", os.Args[0])
	}

	flag.Parse()

	if cfg.Interval < 0.1 {
		return nil, fmt.Errorf("interval must be at least 0.1 seconds")
	}

	if cfg.Interval > 3600 {
		return nil, fmt.Errorf("interval must be less than 3600 seconds")
	}

	return cfg, nil
}