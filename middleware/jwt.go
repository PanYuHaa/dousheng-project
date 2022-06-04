package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	Id   int64  `json:"userId"`
	Name string `json:"name"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	// MyJWTKey 自定义的token秘钥
	MyJWTKey = []byte("SuperKey")
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
	//换票区间
	bufferTime = int64(5 * time.Minute)
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	var err error
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	claims.Issuer = "douyin"
	// 生成token，并签名生成JWT
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MyJWTKey)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return token
}

// ParseToken 解析Token
func ParseToken(tokenString string) *UserClaims {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyJWTKey, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}

// JwtVerify 验证token
func JwtVerify(c *gin.Context) {
	// 过滤是否验证token，暂时空着
	token := c.Query("token")
	if token == "" {
		panic("token not exist !")
	}
	// 如果到了换票的时间区间
	claims := ParseToken(token)
	if claims.ExpiresAt-time.Now().Unix() < bufferTime {
		token = Refresh(claims)
	}
	// 验证token，并存储在请求中
	c.Set("userClaim", ParseToken(token))
}

// Refresh 刷新token
func Refresh(claims *UserClaims) string {
	//jwt.TimeFunc = func() time.Time {
	//	return time.Unix(0, 0)
	//}
	//claims := ParseToken(tokenString)
	//jwt.TimeFunc = time.Now
	//claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}
