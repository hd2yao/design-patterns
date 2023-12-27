package main

type LightState interface {
	// 亮起当前状态的交通灯
	Light()
	// 转换到新状态的时候，调用的方法
	EnterState()
	// 设置一个状态要转变的状态
	NextLight(light *TrafficLight)
	// 检测车速
	CarPassingSpeed(*TrafficLight, int, string)
}
