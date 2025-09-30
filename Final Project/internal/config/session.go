package appconfig

import (
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
)

func SessionInit() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Store = redisstore.New(RedisInit())
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = len(os.Getenv("REDIS")) > 0
	return session
}
