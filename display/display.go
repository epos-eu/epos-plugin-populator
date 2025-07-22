// Package display provides colored console output functions for different message types.
// It uses ANSI color codes to display formatted messages with visual distinction
// based on severity or purpose.
//
// The package includes functions for common logging levels and workflow steps:
//   - Error: Red colored error messages
//   - Warn: Yellow colored warning messages
//   - Info: Blue colored informational messages
//   - Step: Cyan colored step/process messages
//   - Done: Green colored completion messages
//
// All functions follow printf-style formatting and automatically include
// appropriate prefixes and color reset sequences.
//
// Example usage:
//
//	display.Info("Starting process with %d items", count)
//	display.Step("Processing item %s", itemName)
//	display.Done("Successfully processed %d items", processed)
//	display.Error("Failed to process item: %v", err)
package display

import (
	"fmt"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func Error(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s[ERROR]    %s%s\n", colorRed, message, colorReset)
}

func Warn(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s[WARNING]  %s%s\n", colorYellow, message, colorReset)
}

func Info(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s[INFO]     %s%s\n", colorBlue, message, colorReset)
}

func Step(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s[STEP]     %s%s\n", colorCyan, message, colorReset)
}

func Done(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s[DONE]     %s%s\n", colorGreen, message, colorReset)
}
