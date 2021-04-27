package application

import (
	"log"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-common/domain/entity/exception"
	"github.com/smarest/smarest-common/infrastructure/persistence"
)

// LoginService to get user info from cookie
type LoginService struct {
	LoginURL  string
	TokenName string
	*persistence.LoginRepository
}

// NewLoginService create LoginService
func NewLoginService(loginURL string, tokenName string, loginRepository *persistence.LoginRepository) *LoginService {
	return &LoginService{loginURL, tokenName, loginRepository}
}

func (s *LoginService) CheckCookie(c *gin.Context) (interface{}, *exception.Error) {
	cookie, err := c.Cookie(s.TokenName)
	if err != nil {
		log.Print(err)
		return "", exception.CreateError(exception.CodeNotFound, "Cookie not found.")
	}
	user, cErr := s.LoginRepository.GetUserByCookie(cookie)
	if cErr != nil {
		log.Print(cErr.ErrorMessage)
		return "", exception.CreateError(exception.CodeUnknown, "client has error")
	}
	return user, nil

}

func (s *LoginService) SetCookie(c *gin.Context) {
	UID := c.Params.ByName("uid")
	log.Print(s.TokenName + ":" + UID)
	c.SetCookie(s.TokenName, UID, 60*60*24, "", "", http.SameSiteDefaultMode, false, false)
}

func (s *LoginService) GetLoginUrl(c *gin.Context) string {
	return fmt.Sprintf("%s?frm=http://%s%s", s.LoginURL, c.Request.Host, c.Request.URL)
}
