package main

type OffCommand struct {
    device Device
}

func (o *OffCommand) execute() {
    o.device.off()
}
