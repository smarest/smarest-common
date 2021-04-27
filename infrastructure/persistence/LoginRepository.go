package persistence

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smarest/smarest-common/client"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type GetUserNameByCookieResponse struct {
	UserName string `json:"userName"`
}

type LoginRepository struct {
	Client *client.LoginClient
}

func NewLoginRepository(client *client.LoginClient) *LoginRepository {
	return &LoginRepository{Client: client}
}

func (r *LoginRepository) GetUserByCookie(cookie string) (interface{}, *exception.Error) {
	resp, err := r.Client.GetUserByCookie(cookie)
	if err != nil {
		return "", exception.CreateError(exception.CodeUnknown, "client has error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResponse exception.Error
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return "", exception.CreateError(exception.CodeUnknown, err.Error())
		}
		if errResponse.ErrorCode == exception.CodeSignatureInvalid {
			return "", exception.CreateError(exception.CodeSignatureInvalid, "Cookie is invalid.")
		}
		return "", exception.CreateError(exception.CodeUnknown, fmt.Sprintf("api has error. [%s]", errResponse.ErrorMessage))
	}

	var response GetUserNameByCookieResponse
	resErr := json.NewDecoder(resp.Body).Decode(&response)
	if resErr != nil {
		return "", exception.CreateError(exception.CodeUnknown, fmt.Sprintf("can not decode response. [%s]", resErr.Error()))
	}

	if response.UserName == "" {
		return "", exception.CreateError(exception.CodeNotFound, "UserName not found")
	}

	return response.UserName, nil
}
