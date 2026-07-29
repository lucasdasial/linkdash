package main

import "github.com/charmbracelet/lipgloss"

var (
	primary = lipgloss.Color("#e738ed")
	accent  = lipgloss.Color("#04B575")
	muted   = lipgloss.Color("#6E6E6E")
	warn    = lipgloss.Color("#FF5F5F")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(primary).
			Padding(0, 2).
			MarginBottom(1)

	promptStyle = lipgloss.NewStyle().
			Foreground(primary).
			Bold(true)

	hintStyle = lipgloss.NewStyle().
			Foreground(muted).
			Italic(true)

	resultLabelStyle = lipgloss.NewStyle().
				Foreground(muted)

	resultBoxStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(accent).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accent).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(warn)

	byeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			MarginTop(1)
)
