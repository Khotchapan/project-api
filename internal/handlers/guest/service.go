package guest

import (
	"context"
	"errors"
	"log"

	"github.com/khotchapan/KonLakRod-api/internal/core/connection"
	"github.com/khotchapan/KonLakRod-api/internal/core/bcrypt"
	"github.com/khotchapan/KonLakRod-api/internal/entities"
	"github.com/khotchapan/KonLakRod-api/internal/handlers/token"
	"github.com/khotchapan/KonLakRod-api/internal/core/mongodb/user"
	"github.com/labstack/echo/v4"
)

type GuestInterface interface {
	LoginUsers(c echo.Context, request *LoginUsersForm) (*entities.Token, error)
	Login(c echo.Context, request *LoginUsersForm) (*string, error)
}

type Service struct {
	con        *connection.Connection
	collection *connection.Collection
	//tokenService *token.Service
	tokenService token.ServiceInterface
}

func NewService(app, collection context.Context) *Service {
	return &Service{
		con:          connection.GetConnect(app, connection.ConnectionInit),
		collection:   connection.GetCollection(collection, connection.CollectionInit),
		tokenService: token.NewService(app, collection),
	}
}

func (s *Service) LoginUsers(c echo.Context, request *LoginUsersForm) (*entities.Token, error) {
	log.Println("========STEP2========")
	log.Println("request", *request.Username)
	log.Println("request", *request.Password)

	us := &user.Users{}
	err := s.collection.Users.FindOneByName(request.Username, us)
	if err != nil {
		return nil, err
	}
	log.Println("========STEP3========")
	log.Println("us", us)
	log.Println("us.PasswordHash:", us.PasswordHash)
	if !bcrypt.ComparePassword(*request.Password, us.PasswordHash) {
		log.Println("check")
		return nil, errors.New("password is incorrect")
	}
	//token, err := s.collection.TokenService.Create(c, us)
	log.Println("========STEP3.2========")
	token, err := s.tokenService.Create(c, us)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Service) Login(c echo.Context, request *LoginUsersForm) (*string, error) {
	// Throws unauthorized error
	// if *request.Username != "jon" || *request.Password != "shhh!" {
	// 	//return echo.ErrUnauthorized
	// 	return nil, errors.New("math: square root of negative number")
	// }

	// Set custom claims
	// claims := &jwtCustomClaims{
	// 	"Jon Snow",
	// 	true,
	// 	jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	// 	},
	// }
	us := &user.Users{}
	err := s.collection.Users.FindOneByName(request.Username, us)
	if err != nil {
		return nil, err
	}
	if !bcrypt.ComparePassword(*request.Password, us.PasswordHash) {
		log.Println("check")
		return nil, errors.New("password is incorrect")
	}
	token, err := s.tokenService.Create2(c, us)
	if err != nil {
		return nil, err
	}

	return token, nil
}
