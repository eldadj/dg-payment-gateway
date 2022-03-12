package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	auth "github.com/eldadj/dgpg/dto/merchant/authenticate"
	eutil "github.com/eldadj/dgpg/internal/errors"
	mutil "github.com/eldadj/dgpg/internal/merchant"
	"gorm.io/gorm"
)

type Merchant struct {
	MerchantId   int64 `gorm:"primaryKey"`
	Fullname     string
	Username     string `gorm:"column:user_name"`
	PasswordHash string `gorm:"column:pwd_hash"`
	//Password   string `gorm:"-"`
}

func (*Merchant) TableName() string {
	return "merchant"
}

func Authenticate(req auth.Request) (jwtToken string, err error) {
	if req.Username == "" {
		return "", eutil.ErrFieldValueRequired("username")
	}
	if req.Password == "" {
		return "", eutil.ErrFieldValueRequired("password")
	}

	err = ExecDBFunc(func(tx *gorm.DB) error {
		var err error
		var m Merchant
		/*stmt := tx.Session(&gorm.Session{DryRun: true}).Find(&m, "user_name = ?", req.Username).Statement
		println(stmt.SQL.String())*/
		result := tx.Find(&m, "user_name = ?", req.Username)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
			return eutil.ErrMerchantAuthenticationFailed
		}
		if !mutil.IsPasswordHashValid(req.Password, m.PasswordHash) {
			return eutil.ErrMerchantAuthenticationFailed
		}

		claims := jwt.MapClaims{}
		// store the merchantid for quick lookup later
		//TODO: refactor to store something else or use redis for storage and lookup
		claims["mid"] = m.MerchantId

		jwtToken, err = mutil.Generate(claims)

		return err
	})

	return
}
