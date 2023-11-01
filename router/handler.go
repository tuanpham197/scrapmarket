package router

import (
	"sendo/composers"
	apiAuth "sendo/internal/auth/controller/httpapi"
	authMiddleware "sendo/internal/auth/controller/middleware"
	mysqlRepoAuth "sendo/internal/auth/infras/mysql"
	authServeAuth "sendo/internal/auth/service"
	shopMiddleware "sendo/internal/shop/controller/middleware"

	"github.com/redis/go-redis/v9"
	sctx "github.com/viettranx/service-context"

	apiShop "sendo/internal/shop/controller/httpapi"
	mysqlRepoShop "sendo/internal/shop/infras/mysql"
	serveShop "sendo/internal/shop/service"
	"sendo/pkg/common"

	apiFile "sendo/internal/file/controller/httpapi"
	mysqlRepoFile "sendo/internal/file/infras/mysql"
	serveFile "sendo/internal/file/service"

	apiRole "sendo/internal/permission/controller/httpapi"
	mysqlRole "sendo/internal/permission/infras/mysql"
	serveRole "sendo/internal/permission/service"

	apiShopCategory "sendo/internal/shop_category/controller/httpapi"
	mysqlRepoShopCategory "sendo/internal/shop_category/infras/mysql"
	serveShopCategory "sendo/internal/shop_category/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.RouterGroup, db *gorm.DB, rdb *redis.Client, serviceCtx sctx.ServiceContext) {

	authRoute := r.Group("auth")
	repositoryAuth := mysqlRepoAuth.NewMySQLRepo(db, rdb)
	hasher := new(common.Hasher)
	serviceAuth := authServeAuth.NewService(repositoryAuth, hasher)
	controllerAuth := apiAuth.NewAPIController(serviceAuth)
	controllerAuth.SetUpRoute(authRoute)

	// Setup Category
	categoryAPIService := composers.ComposeCategoryAPIService(serviceCtx, db)
	categoryRoute := r.Group("categories")
	{
		categoryRoute.GET("", categoryAPIService.GetListCategory())
		categoryRoute.GET("/:id", categoryAPIService.GetCategoryById())
		categoryRoute.GET("/search/:id", categoryAPIService.SearchCategory())
		categoryRoute.GET("/list-child", categoryAPIService.GetListAndChild())
		categoryRoute.POST("", categoryAPIService.AddCategory())
	}

	// Setup Shop
	shopRoute := r.Group("shops")
	shopRepository := mysqlRepoShop.NewMySQLRepo(db)
	shopService := serveShop.NewShopService(shopRepository)
	shopController := apiShop.NewAPIController(shopService)
	shopController.SetUpRoute(shopRoute)

	// Setup file
	fileRoute := r.Group("files")
	fileRepository := mysqlRepoFile.NewMySQLRepo(db)
	fileService := serveFile.NewService(fileRepository)
	fileController := apiFile.NewAPIController(fileService)
	fileController.SetUpRoute(fileRoute)

	// Setup product
	productApiService := composers.ComposeProductAPIService(serviceCtx, db)
	productRoute := r.Group("products")
	{
		productRoute.GET("", productApiService.GetListProduct())
		productRoute.GET("/:id", productApiService.GetDetail())
		productRoute.GET("/config", productApiService.GetConfig())
		productRoute.Use(authMiddleware.TokenVerificationMiddleware)
		productRoute.Use(shopMiddleware.ShopMiddleware)
		productRoute.POST("/create", productApiService.InsertProduct())
	}

	// Setup product
	shopCategoryRoute := r.Group("shop-categories")
	shopCategoryRepository := mysqlRepoShopCategory.NewMySQLRepo(db)
	shopCategoryService := serveShopCategory.NewShopService(shopCategoryRepository)
	shopCategoryController := apiShopCategory.NewAPIController(shopCategoryService)
	shopCategoryController.SetUpRoute(shopCategoryRoute)

	// Setup permission - role
	roleRoute := r.Group("roles")
	roleRepository := mysqlRole.NewMySQLRepo(db)
	permissionRepository := mysqlRole.NewMySQLRepoPermission((db))
	roleService := serveRole.NewService(roleRepository, permissionRepository)
	roleController := apiRole.NewAPIController(roleService)
	roleController.SetUpRoute(roleRoute)

}
