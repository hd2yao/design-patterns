package main

/*
   我们把电脑拥有的 CPU、RAM 和硬盘视为子系统
   调用方启动电脑前需要分别启动这三个子系统
   因此我们在子系统上增加一个外观对象
   让调用方直接调用外观对象，由外观对象再去分别对接子系统最终完成电脑的启动
*/

func main() {
    computer := NewComputerFacade()
    computer.start()
}
