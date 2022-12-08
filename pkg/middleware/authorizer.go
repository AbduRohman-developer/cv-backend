package middleware

import (
	"fmt"
	"github.com/AbduRohman-developer/cv-backend/config"
	. "github.com/AbduRohman-developer/cv-backend/pkg/jwt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
)

func Authorize(c *gin.Context) {
	var (
		claims = jwt.MapClaims{}
		err    error
		cfg    = config.Get()
	)
	tokenString, exists := c.Get("Authorization")
	if exists {
		claims, err = ExtractWithClaims(tokenString.(string), []byte(cfg.AccessSigningKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})

			return
		}
	}

	enforcer := newEnforcer(cfg)
	if ok, err := enforcer.Enforce(claims["role"], c.FullPath(), c.GetHeader("method")); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("you are not allowed, %v", err),
		})
	}

	c.Next()
}

func newEnforcer(cfg *config.Config) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer(cfg.CasbinModelConfigPath, cfg.CasbinModelPath)
	if err != nil {
		log.Fatal(err)
	}

	return enforcer
}
