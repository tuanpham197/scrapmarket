package router

import (
	apiAuth "sendo/internal/auth/controller/httpapi"
	mysqlRepoAuth "sendo/internal/auth/infras/mysql"
	authServeAuth "sendo/internal/auth/service"
	apicategory "sendo/internal/category/controller/httpapi"
	mysqlRepocategory "sendo/internal/category/infras/mysql"
	serveCategory "sendo/internal/category/service"
	apiShop "sendo/internal/shop/controller/httpapi"
	mysqlRepoShop "sendo/internal/shop/infras/mysql"
	serveShop "sendo/internal/shop/service"
	"sendo/pkg/common"

	apiFile "sendo/internal/file/controller/httpapi"
	mysqlRepoFile "sendo/internal/file/infras/mysql"
	serveFile "sendo/internal/file/service"

	apiProduct "sendo/internal/product/controller/httpapi"
	mysqlRepoProduct "sendo/internal/product/infras/mysql"
	serveProduct "sendo/internal/product/service"

	apiRole "sendo/internal/permission/controller/httpapi"
	mysqlRole "sendo/internal/permission/infras/mysql"
	serveRole "sendo/internal/permission/service"

	apiShopCategory "sendo/internal/shop_category/controller/httpapi"
	mysqlRepoShopCategory "sendo/internal/shop_category/infras/mysql"
	serveShopCategory "sendo/internal/shop_category/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.RouterGroup, db *gorm.DB) {

	authRoute := r.Group("/auth")
	repositoryAuth := mysqlRepoAuth.NewMySQLRepo(db)
	hasher := &common.Hasher{}
	serviceAuth := authServeAuth.NewService(repositoryAuth, hasher)
	controllerAuth := apiAuth.NewAPIController(serviceAuth)
	controllerAuth.SetUpRoute(authRoute)

	// Setup Category
	categoryRoute := r.Group("/categories")
	categoryRepository := mysqlRepocategory.NewMySQLRepo(db)
	categoryService := serveCategory.NewService(categoryRepository)
	categoryController := apicategory.NewAPIController(categoryService)
	categoryController.SetUpRoute(categoryRoute)

	// Setup Shop
	shopRoute := r.Group("/shops")
	shopRepository := mysqlRepoShop.NewMySQLRepo(db)
	shopService := serveShop.NewShopService(shopRepository)
	shopController := apiShop.NewAPIController(shopService)
	shopController.SetUpRoute(shopRoute)

	// Setup file
	fileRoute := r.Group("/files")
	fileRepository := mysqlRepoFile.NewMySQLRepo(db)
	fileService := serveFile.NewService(fileRepository)
	fileController := apiFile.NewAPIController(fileService)
	fileController.SetUpRoute(fileRoute)

	// Setup product
	productRoute := r.Group("/products")
	productRepository := mysqlRepoProduct.NewProductRepo(db)
	productService := serveProduct.NewProductService(productRepository)
	productController := apiProduct.NewAPIController(productService)
	productController.SetUpRoute(productRoute)

	// Setup product
	shopCategoryRoute := r.Group("/shop-categories")
	shopCategoryRepository := mysqlRepoShopCategory.NewMySQLRepo(db)
	shopCategoryService := serveShopCategory.NewShopService(shopCategoryRepository)
	shopCategoryController := apiShopCategory.NewAPIController(shopCategoryService)
	shopCategoryController.SetUpRoute(shopCategoryRoute)

	// Setup permission - role
	roleRoute := r.Group("/roles")
	roleRepository := mysqlRole.NewMySQLRepo(db)
	permissionRepository := mysqlRole.NewMySQLRepoPermission((db))
	roleService := serveRole.NewService(roleRepository, permissionRepository)
	roleController := apiRole.NewAPIController(roleService)
	roleController.SetUpRoute(roleRoute)

}
