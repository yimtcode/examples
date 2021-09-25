package main

// HDMI HDMI输出接口
type HDMI interface {
	HDMIOut() string
}

// VGA VGA输出接口
type VGA interface {
	VGAOut() string
}