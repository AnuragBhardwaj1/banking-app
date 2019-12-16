package services

import (
    "banking-app/providers"
    "banking-app/repositories"
)

type sessionsService struct {
    userRepo     UserRepository
    authProvider AuthProvider
}

func NewSessionsService() *sessionsService {
    authProvider := providers.NewAuthProvider()
    return &sessionsService{userRepo: repositories.NewUserRepository(), authProvider: authProvider}
}

func (sessionsService *sessionsService) AcquireSession(loginRequest LoginRequest) (string, error) {
    credentials := providers.Credentials{Email: loginRequest.Email, Password: loginRequest.Password}
    err, apiKey := sessionsService.authProvider.AcquireSession(credentials)
    if err != nil {
        return "", err
    }
    return apiKey, nil
}
