package main

import (
	c "app/config"
	"app/libs/logger"
	"app/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	apiName := c.Env.App.Name
	e := echo.New()
	e.Server.ReadTimeout = time.Duration(c.Env.App.ServerTimeout) * time.Minute
	e.Server.WriteTimeout = time.Duration(c.Env.App.ServerTimeout) * time.Minute
	e.HideBanner = true
	e.HidePort = true

	log := logger.New(logger.DEAFULT_LOGGER, logger.LogProps{
		LogLevel: logger.DEBUG,
	})

	for _, group := range routes.AppRouting {
		eGroup := e.Group(group.Prefix, group.Middlewares...)

		for _, route := range group.Routes {
			log.Debug(logger.LogInfo{Key: apiName, Value: "Endpoint " + group.Prefix + route.Path})
			eGroup.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
		}
	}

	log.Debug(logger.LogInfo{Key: apiName, Value: "Server START"})
	url := fmt.Sprintf("%s:%d", c.Env.App.Address, c.Env.App.Port)
	log.Debug(logger.LogInfo{Key: apiName, Value: "Server STARTED and RUNNING at: " + url})

	go func() {
		if err := e.Start(url); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	message := "Server is STARTING SHUTDOWN PROCESS"
	log.Info(logger.LogInfo{Key: apiName, Value: message})

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Env.App.ExitTimeout)*time.Second)
	defer cancel()

	message = "Server is SHUTTING DOWN NOW"
	log.Info(logger.LogInfo{Key: apiName, Value: message})

	if err := e.Shutdown(ctx); err != nil {
		message = fmt.Sprintf("Server was SHUTTED DOWN SUDDENLY - Error: %v", err.Error())
		log.Error(logger.LogInfo{Key: apiName, Value: message})
		return
	}

	message = "Server has been SHUTTED DOWN NICELLY"
	log.Info(logger.LogInfo{Key: apiName, Value: message})

}
