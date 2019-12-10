package handler

import (
    "banking-app/services"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    _ "github.com/go-errors/errors"
)


type usersHandler struct {
    userService UserService
}

func NewUsersHandler() *usersHandler {
    userService := services.NewUserService()
    return &usersHandler{userService: userService}
}

func (this *usersHandler) GetUsers(ctx *gin.Context) {}

func (this *usersHandler) GetUser(ctx *gin.Context) {}

func (this *usersHandler) CreateUser(ctx *gin.Context) {
    var registrationDetails *Registration
    if err := ctx.ShouldBind(registrationDetails); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
        return
    }
    createUserRequest := this.buildCreateUserRequest(registrationDetails)

    this.userService.CreateUser(createUserRequest)
    ctx.JSON(http.StatusOK, gin.H{"hello": fmt.Sprintf("+%v", registrationDetails)})
}

func (this *usersHandler) PatchUser(ctx *gin.Context) {}
func (this *usersHandler) DeleteUser(ctx *gin.Context) {}

func (usersHandler *usersHandler) buildCreateUserRequest(registrationDetails *registrationDetails) services.CreateUserRequest {
 return &services.CreateUserRequest {
    email: registrationDetails.Email,
    password: registrationDetails.Password,
 }
}
