package usecase

import (
	"context"
	"golang-laundry-app/model/web/employee"
	"golang-laundry-app/model/web/response"
)

type EmployeeUsecase interface {
	FindAll(ctx context.Context) []response.EmployeeResponse
	FindById(ctx context.Context, employeeId int) response.EmployeeResponse
	Create(ctx context.Context, createRequestEmployee *employee.CreateRequestEmployee) response.EmployeeResponse
	Update(ctx context.Context, updateRequestEmployee *employee.UpdateRequestEmployee) response.EmployeeResponse
	Delete(ctx context.Context, employeeId int)
}
