package mapper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func MapUserDtoIntoUserModel[T *dto.CreateUserDto](userTransferObject T) *model.User {
	var userModel model.User
	err := mapstructure.Decode(userTransferObject, &userModel)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 14)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userModel.Password = string(hashedPassword)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &userModel
}

func MapJwtClaimIntoUserClaim(jwtClaim jwt.MapClaims) (*dto.JwtClaimDto, error) {
	var userClaim dto.JwtClaimDto
	err := mapstructure.Decode(jwtClaim, &userClaim)
	if err != nil {
		return nil, err
	}
	return &userClaim, nil
}
