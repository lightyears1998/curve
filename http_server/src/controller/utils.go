package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

func checkAndGetUid(c *gin.Context) (int, error) {
	secretTokenString, err := c.Cookie(model.KeyForTokenInCookies)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	uid, err := GlobalTokenManager.GetUidFromSecretTokenString(secretTokenString)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return uid, nil
}
