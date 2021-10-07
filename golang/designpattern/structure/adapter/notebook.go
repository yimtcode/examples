package main

type Notebook struct {
	DisplayContent string
}

func (n *Notebook) HDMIOut() string {
	return "HDMI output: " + n.DisplayContent
}
