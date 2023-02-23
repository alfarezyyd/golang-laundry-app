package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-laundry-app/controller"
	"golang-laundry-app/exception"
)

type (
	EmployeeController  controller.WebController
	BranchController    controller.WebController
	InventoryController controller.WebController
	ServiceController   controller.WebController
	PromoController     controller.WebController
	UserController      controller.WebController
	AdminController     controller.WebController
	OrderController     controller.WebController
)

func NewSetupRouter(
	employeeController EmployeeController,
	branchController BranchController,
	inventoryController InventoryController,
	serviceController ServiceController,
	promoController PromoController,
	userController UserController,
	adminController AdminController,
	orderController OrderController,
	authController controller.AuthController,
) *httprouter.Router {
	router := httprouter.New()

	// Panic Handler
	router.PanicHandler = exception.ErrorHandler

	// Authentication User
	router.POST("/login", authController.Login)
	router.GET("/logout", authController.Logout)

	// Branch Route
	router.GET("/api/branchs", branchController.FindAll)
	router.GET("/api/branchs/:branchId/detail", branchController.FindById)
	router.POST("/api/branchs", branchController.Create)
	router.PUT("/api/branchs/:branchId/edit", branchController.Update)
	router.DELETE("/api/branchs/:branchId/delete", branchController.Delete)

	// Employee Route
	router.GET("/api/employees", employeeController.FindAll)
	router.GET("/api/employees/:employeeId/detail", employeeController.FindById)
	router.POST("/api/employees", employeeController.Create)
	router.PUT("/api/employees/:employeeId/edit", employeeController.Update)
	router.DELETE("/api/employees/:employeeId/delete", employeeController.Delete)

	// Inventory Route
	router.GET("/api/inventories", inventoryController.FindAll)
	router.GET("/api/inventories/:inventoryId/detail", inventoryController.FindById)
	router.POST("/api/inventories", inventoryController.Create)
	router.PUT("/api/inventories/:inventoryId/edit", inventoryController.Update)
	router.DELETE("/api/inventories/:inventoryId/delete", inventoryController.Delete)

	// Service Route
	router.GET("/api/services", serviceController.FindAll)
	router.GET("/api/services/:serviceId/detail", serviceController.FindById)
	router.POST("/api/services", serviceController.Create)
	router.PUT("/api/services/:serviceId/edit", serviceController.Update)
	router.DELETE("/api/services/:serviceId/delete", serviceController.Delete)

	// Promo Route
	router.GET("/api/promos", promoController.FindAll)
	router.GET("/api/promos/:promoId/detail", promoController.FindById)
	router.POST("/api/promos", promoController.Create)
	router.PUT("/api/promos/:promoId/edit", promoController.Update)
	router.DELETE("/api/promos/:promoId/delete", promoController.Delete)

	// User Route
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId/detail", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId/edit", userController.Update)
	router.DELETE("/api/users/:userId/delete", userController.Delete)

	// Admin Route
	router.GET("/api/admins", adminController.FindAll)
	router.GET("/api/admins/:adminId/detail", adminController.FindById)
	router.POST("/api/admins", adminController.Create)
	router.PUT("/api/admins/:adminId/edit", adminController.Update)
	router.DELETE("/api/admins/:adminId/delete", adminController.Delete)

	// Order Route
	router.GET("/api/orders", orderController.FindAll)
	router.GET("/api/orders/:orderId/detail", orderController.FindById)
	router.POST("/api/orders", orderController.Create)
	router.PUT("/api/orders/:orderId/edit", orderController.Update)
	router.DELETE("/api/orders/:orderId/delete", orderController.Delete)

	return router
}
