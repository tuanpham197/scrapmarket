package main

import (
	"log"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
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

	// Set timezone
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")

	db, err := connectDBWithRetry(5)

	if err != nil {
		log.Fatalln(err)
	}
	engine := gin.Default()

	// SETUP CORS
	configCors(engine)

	// DOCS SWAGGER
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Recover panic error
	engine.Use(common.Recover())

	v1 := engine.Group("/api/v1")
	router.InitRouter(v1, db)
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}

}

func configCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{FE_URL}                         // Specify the allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"} // Specify the allowed HTTP methods
	config.AllowHeaders = []string{"*"}
	engine.Use(cors.New(config))
}

func connectDBWithRetry(times int) (*gorm.DB, error) {
	var e error

	for i := 1; i <= times; i++ {
		dsn := os.Getenv("MYSQL_DSN")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			// sqlDB, err := db.DB()
			// if err != nil {
			// 	return nil, err
			// }

			// // Check max connection mysql allowed: select @@max_connections
			// // Default: select @@max_connections = 151
			// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			// sqlDB.SetMaxIdleConns(10)

			// // SetMaxOpenConns sets the maximum number of open connections to the database.
			// sqlDB.SetMaxOpenConns(100)

			return db, nil
		}

		e = err

		time.Sleep(time.Second * 2)
	}

	return nil, e
}
