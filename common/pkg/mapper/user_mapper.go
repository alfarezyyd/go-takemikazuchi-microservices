package mapper

import (
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	grpcUser "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"googlemaps.github.io/maps"
	"net/http"
	"time"
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

func MapUserModelIntoUserResponse(userModel *model.User) *dto.UserResponseDto {
	var userResponse dto.UserResponseDto
	err := mapstructure.Decode(userModel, &userResponse)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, nil))
	return &userResponse
}

func MapUserResponseIntoQueryUserResponse(userResponse *dto.UserResponseDto) *grpcUser.QueryUserResponse {
	var queryUserResponse grpcUser.QueryUserResponse
	err := mapstructure.Decode(userResponse, &queryUserResponse)
	queryUserResponse.CreatedAt = userResponse.CreatedAt.Format(time.RFC3339)
	queryUserResponse.UpdatedAt = userResponse.UpdatedAt.Format(time.RFC3339)
	fmt.Println(err)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, nil))
	return &queryUserResponse
}

func MapUserJwtClaimIntoUserJwtClaimGrpc(jwtClaim *dto.JwtClaimDto) *grpcUser.UserJwtClaim {
	var userJwtClaim grpcUser.UserJwtClaim
	err := mapstructure.Decode(jwtClaim, &userJwtClaim)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, nil))
	return &userJwtClaim
}

func MapUserJwtClaimGrpcIntoUserJwtClaim(userJwtClaim *grpcUser.UserJwtClaim) *dto.JwtClaimDto {
	var jwtClaimDto dto.JwtClaimDto
	err := mapstructure.Decode(userJwtClaim, &jwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, nil))
	return &jwtClaimDto
}

func MapReverseGeocodingIntoUserAddresses(geocodingResult *maps.GeocodingResult,
	userAddress *model.UserAddress,
	userId uint64,
	addressAdditionalAddress string) {
	userAddress.FormattedAddress = geocodingResult.FormattedAddress
	userAddress.PlaceId = geocodingResult.PlaceID
	userAddress.UserId = userId
	userAddress.Longitude = geocodingResult.Geometry.Location.Lng
	userAddress.Latitude = geocodingResult.Geometry.Location.Lat
	userAddress.AdditionalInformation = addressAdditionalAddress
	for _, addressComponent := range geocodingResult.AddressComponents {
		switch addressComponent.Types[0] {
		case "street_address":
		case "street_number":
			userAddress.StreetNumber += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "route":
			userAddress.Route += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "administrative_area_level_4":
			userAddress.Village += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "administrative_area_level_3":
			userAddress.District += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "administrative_area_level_2":
			userAddress.City += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "administrative_area_level_1":
			userAddress.Province += fmt.Sprintf("%s ", addressComponent.LongName)
			break
		case "country":
			userAddress.Country = addressComponent.LongName
			break
		case "postal_code":
			userAddress.PostalCode = addressComponent.LongName
			break
		}
	}
}
