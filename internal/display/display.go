package display

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ImageDisplay struct {
	useImgcat bool
}

func New() (*ImageDisplay, error) {
	d := &ImageDisplay{}

	// Check if imgcat is available
	_, err := exec.LookPath("imgcat")
	if err == nil {
		d.useImgcat = true
		return d, nil
	}

	// Imgcat not found
	return nil, fmt.Errorf("imgcat command not found. Please install it first")
}

func (d *ImageDisplay) Display(imageData []byte) error {
	return d.DisplayWithSize(imageData, "")
}

func (d *ImageDisplay) DisplayWithSize(imageData []byte, width string) error {
	if !d.useImgcat {
		return fmt.Errorf("no suitable image display method available")
	}

	args := []string{}
	if width != "" && width != "auto" {
		args = append(args, "-W", width)
	}

	cmd := exec.Command("imgcat", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start imgcat: %w", err)
	}

	if _, err := io.WriteString(stdin, string(imageData)); err != nil {
		stdin.Close()
		cmd.Wait()
		return fmt.Errorf("failed to write image data: %w", err)
	}

	stdin.Close()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("imgcat failed: %w", err)
	}

	return nil
}

func (d *ImageDisplay) ClearScreen() {
	// Move cursor to home position
	fmt.Print("\033[H")
	// Clear from cursor to end of screen
	fmt.Print("\033[J")
}

func (d *ImageDisplay) MoveCursorHome() {
	// Move cursor to home position
	fmt.Print("\033[H")
}