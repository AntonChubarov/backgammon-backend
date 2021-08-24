package auth
//
//import (
//	"backgammon/config"
//	domainAuth "backgammon/domain/authdomain"
//	"backgammon/utils"
//	"github.com/dlclark/regexp2"
//)
//
//type UserAuthService struct {
//	storage            domainAuth.UserStorage
//	mainSessionStorage SessionStorage
//	config             *config.ServerConfig
//	hasher             StringHasher
//	tokenGenerator     TokenGenerator
//	usernameRegexp *regexp2.Regexp
//	passwordRegexp *regexp2.Regexp
//}
//
//func NewUserAuthService(storage domainAuth.UserStorage,
//	mainSessionStorage SessionStorage,
//	config *config.ServerConfig,
//	tokenGenerator TokenGenerator) *UserAuthService {
//	return &UserAuthService{storage: storage,
//		mainSessionStorage: mainSessionStorage,
//		config:             config,
//		hasher:             NewHasherSHA256(),
//		tokenGenerator:     tokenGenerator,
//		usernameRegexp: regexp2.MustCompile("^(?=.{6,20}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$", 0),
//		passwordRegexp: regexp2.MustCompile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$", 0),
//	}
//}
//
//func (uas *UserAuthService) RegisterNewUser(data domainAuth.UserData) error {
//	var isMatch bool
//
//	if isMatch, _ = uas.usernameRegexp.MatchString(string(data.UserName)); !isMatch {
//		return ErrorPoorUsername
//	}
//	if isMatch, _ = uas.passwordRegexp.MatchString(string(data.Password)); !isMatch {
//		return ErrorPoorPassword
//	}
//
//	_, err := uas.storage.GetUserByUsername(data.UserName)
//	if err == nil {
//		return ErrorUserExists
//	}
//
//	data.UUID = utils.GenerateUUID()
//
//	passwordHash, err := uas.hasher.HashString(string(data.Password))
//	if err != nil {
//		//log.Println("In app.RegisterNewUser", err)
//		return err
//	}
//	data.Password = domainAuth.Password(passwordHash)
//
//	err = uas.storage.AddNewUser(data)
//	if err != nil {
//		//log.Println("In app.RegisterNewUser", err)
//		return err
//	}
//
//	return nil
//}
//
//func (uas *UserAuthService) AuthorizeUser(data domainAuth.UserData) (token string, err error) {
//	token = ""
//	var user domainAuth.UserData
//
//	user, err = uas.storage.GetUserByUsername(data.UserName)
//	if err != nil {
//		return
//	}
//
//	var passwordHash string
//	passwordHash, err = uas.hasher.HashString(string(data.Password))
//	if err != nil {
//		return
//	}
//
//	if passwordHash != string(user.Password) {
//		err = ErrorInvalidPassword
//		return
//	}
//
//	var wasFound bool
//	token, wasFound = uas.mainSessionStorage.GetTokenByUUID(string(user.UUID))
//	if wasFound {
//		return
//	}
//	token = uas.tokenGenerator.GenerateToken()
//	uas.mainSessionStorage.AddNewUser(&UserSessionData{Token: token, UserUUID: string(user.UUID)})
//	return
//}
//
//func (uas *UserAuthService) CheckToken(token string) error {
//	if !uas.mainSessionStorage.IsTokenValid(token) {
//		return ErrorInvalidToken
//	}
//	return nil
//}
