package mapper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/internal/model"
	dto2 "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func MapUserDtoIntoUserModel[T *dto2.CreateUserDto](userTransferObject T) *model.User {
	var userModel model.User
	err := mapstructure.Decode(userTransferObject, &userModel)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 14)
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userModel.Password = string(hashedPassword)
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	return &userModel
}

func MapJwtClaimIntoUserClaim(jwtClaim jwt.MapClaims) (*dto2.JwtClaimDto, error) {
	var userClaim dto2.JwtClaimDto
	err := mapstructure.Decode(jwtClaim, &userClaim)
	if err != nil {
		return nil, err
	}
	return &userClaim, nil
}
