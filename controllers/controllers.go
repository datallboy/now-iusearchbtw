package controllers

import (
	"net/http"
	"now-iusearchbtw/config"
	"now-iusearchbtw/responses"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
)

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Pong!"}})
	}
}

func NewContainer(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		containerID, err := createContainer(config)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err}})
		}

		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"containerID": containerID}})
	}
}

func KillContainer(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		containerID := c.QueryParam("containerID")

		if containerID == "" {
			return c.JSON(http.StatusBadRequest, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": "Specify container ID in query parameters."}})
		}

		err := config.Client.ContainerStop(config.Context, containerID, container.StopOptions{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err}})
		}

		err = config.Client.ContainerRemove(config.Context, containerID, types.ContainerRemoveOptions{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err}})
		}

		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"containerID": containerID}})
	}
}

func createContainer(c *config.Config) (string, error) {
	container, err := c.Client.ContainerCreate(c.Context, &container.Config{
		Image:        "archlinux",
		Cmd:          []string{"/bin/bash"},
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		Tty:          true,
	}, nil, nil, nil, "")
	if err != nil {
		return "", err
	}

	if err := c.Client.ContainerStart(c.Context, container.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return container.ID[:10], nil
}
