package service

import "dousheng-demo/repository"

func Init() {
	startId = repository.TimeLimitAmount(9999999999)
}
