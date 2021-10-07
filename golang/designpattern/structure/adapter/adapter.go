package main

import "strings"

// HDMIToVGA Adapter: HDMI to VGA
type HDMIToVGA struct {
	displayContent string
}

func (h *HDMIToVGA) VGAOut() string {
	return "VGA output: " + h.displayContent
}

// NewAdapterHDMIToVGA adapter factory function
func NewAdapterHDMIToVGA(hdmi HDMI) VGA {
	hdmiContent := hdmi.HDMIOut()
	// convert content
	displayContent := strings.Split(hdmiContent, ": ")[1]
	return &HDMIToVGA{displayContent: displayContent}
}
