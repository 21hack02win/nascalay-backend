package handler

import (
	"net/http"

	"github.com/21hack02win/nascalay-backend/usecases/repository"
	"github.com/labstack/echo/v4"
)

func (h *handler) JoinRoom(c echo.Context) error {
	req := new(JoinRoomJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room, err := h.r.JoinRoom(&repository.JoinRoomArgs{
		Avatar:   req.Avatar,
		RoomId:   req.RoomId,
		Username: req.Username,
	})
	if err != nil {
		return newEchoHTTPError(err)
	}

	return echo.NewHTTPError(http.StatusOK, refillRoom(room, room.Members[len(room.Members)-1].Id)) // TODO: sessionでとる
}

func (h *handler) CreateRoom(c echo.Context) error {
	req := new(CreateRoomJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room, err := h.r.CreateRoom(&repository.CreateRoomArgs{
		Avatar:   req.Avatar,
		Capacity: req.Capacity,
		Username: req.Username,
	})
	if err != nil {
		return newEchoHTTPError(err)
	}

	return echo.NewHTTPError(http.StatusCreated, refillRoom(room, room.HostId))
}

func (h *handler) GetRoom(c echo.Context, roomId RoomId) error {
	return nil
}
