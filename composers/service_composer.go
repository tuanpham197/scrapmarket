package composers

import (
	apiCategory "sendo/internal/category/controller/httpapi"
	mysqlRepocategory "sendo/internal/category/infras/mysql"
	rpcClient "sendo/internal/category/infras/rpc"
	serveCategory "sendo/internal/category/service"
	apiProduct "sendo/internal/product/controller/httpapi"
	mysqlRepoProduct "sendo/internal/product/infras/mysql"
	serveProduct "sendo/internal/product/service"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"gorm.io/gorm"
)

type CategoryService interface {
	AddCategory() func(*gin.Context)
	GetListCategory() func(*gin.Context)
	GetCategoryById() func(*gin.Context)
	SearchCategory() func(*gin.Context)
	GetListAndChild() func(*gin.Context)
}

type ProductService interface {
	InsertProduct() func(*gin.Context)
	GetListProduct() func(*gin.Context)
	GetDetail() func(*gin.Context)
	GetConfig() func(ctx *gin.Context)
}

func ComposeCategoryAPIService(serviceCtx sctx.ServiceContext, db *gorm.DB) CategoryService {
	authRepo := mysqlRepocategory.NewMySQLRepo(db)

	repoRPC := rpcClient.NewRPCClient(ComposeHelloRPCClient(serviceCtx))
	biz := serveCategory.NewService(authRepo, repoRPC)
	serviceAPI := apiCategory.NewAPIController(biz)

	return serviceAPI
}

func ComposeProductAPIService(serviceCtx sctx.ServiceContext, db *gorm.DB) ProductService {
	productRepository := mysqlRepoProduct.NewProductRepo(db)
	productService := serveProduct.NewProductService(productRepository)
	productController := apiProduct.NewAPIController(productService)

	return productController
}
