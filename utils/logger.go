package utils

import (
	"github.com/fatih/color"
)

type logger struct{}

// Colors
var blue = color.New(color.FgCyan).PrintlnFunc()
var red = color.New(color.FgRed).PrintlnFunc()
var orange = color.New(color.FgYellow).PrintlnFunc()
var green = color.New(color.FgGreen).PrintlnFunc()

/**
 * Info	Log info message
 */
func (l *logger) Info(s string) {
	blue("[info] " + s)
}

/**
 * Warn	Log warning message
 */
func (l *logger) Warn(s string) {
	orange("[warn] " + s)
}

/**
 * Error	Log error message
 */
func (l *logger) Error(s string) {
	red("[error] " + s)
}

/**
 * Success	Log success message
 */
func (l *logger) Success(s string) {
	green("[success] " + s)
}

/**
 * LoggerService	Default logger instance.
 */
var LoggerService = &logger{}
