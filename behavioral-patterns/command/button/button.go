package main

// 请求者

type Button struct {
    command Command
}

func (b *Button) press() {
    b.command.execute()
}
