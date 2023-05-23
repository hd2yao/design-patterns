package manage_tower

import "design-patterns/mediator/airport/interf"

type manageTower struct {
    isRunwayFree bool
    airportQueue []interf.Airplane
}

func NewManageTower() *manageTower {
    return &manageTower{
        isRunwayFree: true,
    }
}

func (tower *manageTower) CanLanding(airplane interf.Airplane) bool {
    if tower.isRunwayFree {
        // 跑道空闲，允许降落，同时把状态变为繁忙
        tower.isRunwayFree = false
        return true
    }
    // 跑道繁忙，把飞机加入等待通知的队列
    tower.airportQueue = append(tower.airportQueue, airplane)
    return false
}

func (tower *manageTower) NotifyAboutDeparture() {
    if !tower.isRunwayFree {
        // 一架飞机起飞后，改变跑道状态为空闲，与 CanLanding 中修改状态对应
        tower.isRunwayFree = true
    }
    // 一架飞机起飞后，通知等待队列中的第一架飞机;如果队列为空，就无需通知
    if len(tower.airportQueue) > 0 {
        firstPlaneInWaitingQueue := tower.airportQueue[0]
        tower.airportQueue = tower.airportQueue[1:]
        firstPlaneInWaitingQueue.PermitLanding()
    }
}
