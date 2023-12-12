package main

const (
    BOOT_ADDRESS = 0
    BOOT_SECTOR  = 0
    SECTOR_SIZE  = 0
)

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        new(CPU),
        new(Memory),
        new(HardDrive),
    }
}

type ComputerFacade struct {
    processor *CPU
    ram       *Memory
    hd        *HardDrive
}

func (c *ComputerFacade) start() {
    c.processor.Freeze()
    c.ram.Load(BOOT_ADDRESS, c.hd.Read(BOOT_SECTOR, SECTOR_SIZE))
    c.processor.Jump(BOOT_ADDRESS)
    c.processor.Execute()
}
