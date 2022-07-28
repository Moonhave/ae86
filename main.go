package main

import (
	"ae86/cmd/app"
	"ae86/pkg/banner"
	"ae86/pkg/logger"
	"runtime"
	"time"
)

const bannerTemplate = `
{{ .ansiColor.Red }}  █████╗ ███████╗ █████╗  ██████╗  {{ .ansiColor.Green }} Now:      {{ .ansiColor.Blue }} {{ .now }}
{{ .ansiColor.Red }} ██╔══██╗██╔════╝██╔══██╗██╔════╝  {{ .ansiColor.Green }} NumCPU:   {{ .ansiColor.Blue }} {{ .numCPU }}
{{ .ansiColor.Red }} ███████║█████╗  ╚█████╔╝███████╗  {{ .ansiColor.Green }} GOOS:     {{ .ansiColor.Blue }} {{ .GOOS }}
{{ .ansiColor.Red }} ██╔══██║██╔══╝  ██╔══██╗██╔═══██╗ {{ .ansiColor.Green }} GOARCH:   {{ .ansiColor.Blue }} {{ .GOARCH }}
{{ .ansiColor.Red }} ██║  ██║███████╗╚█████╔╝╚██████╔╝ {{ .ansiColor.Green }} Compiler: {{ .ansiColor.Blue }} {{ .Compiler }}
{{ .ansiColor.Red }} ╚═╝  ╚═╝╚══════╝ ╚════╝  ╚═════╝  {{ .ansiColor.Default }}
`

func main() {
	banner.Default(bannerTemplate, map[string]interface{}{
		"now":      time.Now().Format(time.ANSIC),
		"numCPU":   runtime.NumCPU(),
		"GOOS":     runtime.GOOS,
		"GOARCH":   runtime.GOARCH,
		"Compiler": runtime.Compiler,
	})

	logger.Init()
	app.Run()
}
