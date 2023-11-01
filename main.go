package main

import (
	"fmt"
	"log"
	"os"
	"sendo/cmd"
	"sendo/db/connections"
	"sendo/pkg/queue"
	"time"

	sctx "github.com/viettranx/service-context"

	"github.com/redis/go-redis/v9"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "sendo/docs"
	"sendo/pkg/common"

	// "sendo/pkg/queue"
	// "sendo/pkg/utils/response"
	"sendo/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

const FE_URL = "*"
const DEFAULT_RETRY_CONNECT = 5

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("Demo Microservices"),
		sctx.WithComponent(cmd.NewConfig()),
	)
}

// @title           Sendo fake
// @version         1.0
// @description     Sendo

// @contact.name   Scrap team
// @contact.url    https://twitter.com/sendo
// @contact.email  sendo_fake@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4040
// @BasePath  /api/v1
func main() {
	// Init service context
	serviceCtx := newServiceCtx()
	db, err := connectDBWithRetry(DEFAULT_RETRY_CONNECT)
	if err != nil {
		log.Fatalln(err)
	}

	engine := gin.Default()

	queue.SetupRabbitMQConnectionChannel()

	// SETUP CORS
	configCors(engine)

	// DOCS SWAGGER
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Recover panic error
	engine.Use(common.Recover())

	v1 := engine.Group("/api/v1")
	rdb, _ := connectRedis()
	router.InitRouter(v1, db, rdb, serviceCtx)
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}

}

func configCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{FE_URL} // Specify the allowed origins
	config.AllowHeaders = []string{"*"}
	engine.Use(cors.New(config))
}

func connectDBWithRetry(times int) (*gorm.DB, error) {
	var e error

	for i := 1; i <= times; i++ {
		db, err := connections.GetInstance()

		if err == nil {
			return db, nil
		}
		e = err
		time.Sleep(time.Second * 2)
	}

	return nil, e
}

func connectRedis() (*redis.Client, error) {
	strConn := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     strConn,
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	if rdb == nil {
		log.Fatal("error connect redis")
		return nil, nil
	}
	return rdb, nil
}
