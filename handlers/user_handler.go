package handlers

import (
	v1 "echo-demo/clients/userclient"
	"echo-demo/clients/userclient/dtos"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

var UserClient v1.UserClient

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := UserClient.DeleteUser(c.Request().Context(), &v1.DeleteUserRequest{
		Id: int32(id),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	success, errMsg := resp.GetSuccess(), resp.GetError()
	if !success {
		return echo.NewHTTPError(http.StatusInternalServerError, errMsg)
	}
	return c.JSON(http.StatusOK, success)
}

func UpdateUser(c echo.Context) error {
	userRequest := dtos.UserRequest{}
	err := c.Bind(&userRequest)
	if err != nil {
		return err
	}
	resp, err := UserClient.UpdateUser(c.Request().Context(), &v1.UpdateUserRequest{
		UserDto: &v1.UserDto{
			Id:               userRequest.Id,
			Name:             userRequest.Name,
			Age:              userRequest.Age,
			Email:            userRequest.Email,
			CurrentAddress:   userRequest.CurrenAddress,
			PermanentAddress: userRequest.PermanentAddress,
		},
	})
	if err != nil {
		return err
	}
	userClientResponse := resp.GetUserResponse()
	userResponse := convertTo(userClientResponse.GetUserDto())
	userResponse.Id = resp.GetId()
	if userClientResponse.GetError() != "" {
		return echo.NewHTTPError(http.StatusInternalServerError, userClientResponse.GetError())
	}
	return c.JSON(http.StatusOK, userResponse)
}

func CreatUser(c echo.Context) error {
	userRequest := dtos.UserRequest{}
	err := c.Bind(&userRequest)
	if err != nil {
		return err
	}
	resp, err := UserClient.CreateUser(c.Request().Context(), &v1.CreateUserRequest{
		User: &v1.UserDto{
			Name:             userRequest.Name,
			Age:              userRequest.Age,
			Email:            userRequest.Email,
			CurrentAddress:   userRequest.CurrenAddress,
			PermanentAddress: userRequest.PermanentAddress,
		},
	})
	if err != nil {
		return err
	}
	userClientResponse := resp.GetUserResponse()
	userResponse := convertTo(userClientResponse.GetUserDto())
	if userClientResponse.GetError() != "" {
		return echo.NewHTTPError(http.StatusInternalServerError, userClientResponse.GetError())
	}
	return c.JSON(http.StatusOK, userResponse)
}

func GetAllUsers(c echo.Context) error {
	resp, err := UserClient.ListUser(c.Request().Context(), &v1.ListUserRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userResponseArray := make([]dtos.UserResponse, 0)
	for _, userDto := range resp.GetUserDtos() {
		userResponseArray = append(userResponseArray, convertTo(userDto))
	}
	return c.JSON(http.StatusOK, dtos.UsersResponse{Users: userResponseArray})
}

func convertTo(userDto *v1.UserDto) dtos.UserResponse {
	return dtos.UserResponse{Id: userDto.GetId(), Name: userDto.GetName(), Email: userDto.GetEmail(),
		CurrentAddress: userDto.GetCurrentAddress(), PermanentAddress: userDto.GetPermanentAddress(), Age: userDto.GetAge()}
}

func GetUsers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := UserClient.GetUser(c.Request().Context(), &v1.GetUserRequest{
		Id: int32(id),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	userDto := resp.GetUserDto()
	return c.JSON(http.StatusOK, convertTo(userDto))
}
