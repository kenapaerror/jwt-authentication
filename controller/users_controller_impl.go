package controller

import (
	"net/http"
	"yt-users-service/helper"
	"yt-users-service/model/web"
	"yt-users-service/service"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UsersCreateRequest{}
	helper.BodyToRequest(request, &userCreateRequest)

	response := controller.UserService.Create(request.Context(), userCreateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UsersUpdateRequest{}
	helper.BodyToRequest(request, &userUpdateRequest)

	userUpdateRequest.Id = params.ByName("user_id")

	response := controller.UserService.Update(request.Context(), userUpdateRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")

	controller.UserService.Delete(request.Context(), userId)

	webResponse := web.Response{
		Status: "OK",
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("user_id")

	response := controller.UserService.FindById(request.Context(), userId)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := controller.UserService.FindAll(request.Context())

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) Auth(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userAuthRequest := web.UserAuthRequest{}
	helper.BodyToRequest(request, &userAuthRequest)

	response := controller.UserService.Auth(request.Context(), userAuthRequest)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}

func (controller *UserControllerImpl) CreateWithRefreshToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")

	response := controller.UserService.CreateWithRefreshToken(request.Context(), token)

	webResponse := web.Response{
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(writer, webResponse)
}
