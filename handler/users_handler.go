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

//----------------------------------------------------------

func (this *usersHandler) CreateUser(ctx *gin.Context) {
    var registrationDetails Registration
    if err := ctx.ShouldBindJSON(&registrationDetails); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
        return
    }
    createUserRequest := this.buildCreateUserRequest(&registrationDetails)

    err := this.userService.CreateUser(createUserRequest)
    if err != nil {
        ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("+%v", registrationDetails)})
    }
    ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (usersHandler *usersHandler) buildCreateUserRequest(registrationDetails *Registration) *services.CreateUserRequest {
    return &services.CreateUserRequest{
        Email:    registrationDetails.Email,
        Password: registrationDetails.Password,
        UserName: registrationDetails.UserName,
    }
}

func (this *usersHandler) GetUsers(ctx *gin.Context) {
    var apiKeyHeader ApiKeyHeader

    if err := ctx.ShouldBindHeader(&apiKeyHeader); err != nil {
        ctx.JSON(200, err)
        return
    }

    usersdao, err := this.userService.GetUsers(apiKeyHeader.ApiKey)
    if err != nil {
        ctx.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("+%v", err)})
        return
    }
    usersDto := this.mapToDto(usersdao)
    ctx.JSON(http.StatusOK,ListUserDao{Users: usersDto})
    return
}

func (this *usersHandler) GetUser(ctx *gin.Context)    {}
func (this *usersHandler) PatchUser(ctx *gin.Context)  {}
func (this *usersHandler) DeleteUser(ctx *gin.Context) {}

func (this *usersHandler) mapToDto(usersdao []*services.User) []UserInfo {
    var usersDto []UserInfo
    for _, user := range usersdao {
        user := UserInfo{
            Name:   user.Name,
            Email:  user.Email,
            ApiKey: user.ApiKey,
        }
        usersDto = append(usersDto, user)
    }
    return usersDto
}

//--------------------------------------------------

