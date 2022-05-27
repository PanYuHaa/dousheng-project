package controller

import "dousheng-demo/service"

func Init() {
	userIdSequence = service.GetLastUserId()
}
