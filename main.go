package main

import (
	"github.com/selfharm-enj/otus_m2_hw12/internal/model"
	"github.com/selfharm-enj/otus_m2_hw12/internal/repository"
	"github.com/selfharm-enj/otus_m2_hw12/internal/service"
)

func main() {
	dataCh := make(chan model.IDReader)
	service.StartService(dataCh)

	go func() {
		for data := range dataCh {
			repository.AddData(data)
		}
	}()

	go service.LogChanges()
	select {}
}
