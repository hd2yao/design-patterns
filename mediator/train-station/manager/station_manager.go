package manager

import "design-patterns/mediator/train-station/interf"

type stationManager struct {
    isPlatformFree bool
    trainQueue     []interf.Train
}

func NewStationManager() *stationManager {
    return &stationManager{isPlatformFree: true}
}

func (manager *stationManager) CanArrive(train interf.Train) bool {
    if manager.isPlatformFree {
        manager.isPlatformFree = false
        return true
    }
    manager.trainQueue = append(manager.trainQueue, train)
    return false
}

func (manager *stationManager) NotifyAboutDeparture() {
    if !manager.isPlatformFree {
        manager.isPlatformFree = true
    }
    if len(manager.trainQueue) > 0 {
        firstTrainInQueue := manager.trainQueue[0]
        manager.trainQueue = manager.trainQueue[1:]
        firstTrainInQueue.PermitArrival()
    }
}
