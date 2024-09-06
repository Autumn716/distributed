package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "6000"
	ServiceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       ServiceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: ServiceAddress + "/services",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}

	if logProvier, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvier)
		log.SetClientLogger(logProvier, r.ServiceName)
	}

	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}
