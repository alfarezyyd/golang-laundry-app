package helper

import (
	"golang-laundry-app/model/domain"
	"golang-laundry-app/model/web/response"
)

func ConvertToAddressResponse(address *domain.Address) response.AddressResponse {
	return response.AddressResponse{
		Id:                  address.Id,
		BuildingName:        address.BuildingName,
		BuildingNumber:      address.BuildingNumber,
		Street:              address.Street,
		Village:             address.Village,
		NeighbourhoodNumber: address.NeighbourhoodNumber,
		HamletNumber:        address.HamletNumber,
		SubDistrict:         address.SubDistrict,
		District:            address.District,
		Province:            address.Province,
		PostalCode:          address.PostalCode,
		Description:         address.Description,
	}
}

func ConvertToBranchResponse(branch *domain.Branch, address *response.AddressResponse) response.BranchResponse {
	return response.BranchResponse{
		Id:              branch.Id,
		Name:            branch.Name,
		TelephoneNumber: branch.TelephoneNumber,
		CreatedAt:       &branch.CreatedAt,
		UpdatedAt:       &branch.UpdatedAt,
		Address:         address,
	}
}

func ConvertToEmployeeResponse(employee *domain.Employee, address *response.AddressResponse, branchName *string) response.EmployeeResponse {
	return response.EmployeeResponse{
		Id:              employee.Id,
		Address:         address,
		BranchName:      branchName,
		Code:            employee.Code,
		Level:           employee.Level,
		FullName:        employee.FullName,
		Password:        employee.Password,
		Email:           employee.Email,
		TelephoneNumber: employee.TelephoneNumber,
		Photo:           employee.Photo,
		Status:          employee.Status,
		CreatedAt:       &employee.CreatedAt,
		UpdatedAt:       &employee.UpdatedAt,
	}
}

func ConvertToInventoryResponse(inventory *domain.Inventory, employee *domain.Employee) response.InventoryResponse {
	employeeResponse := ConvertToEmployeeResponse(employee, nil, nil)

	return response.InventoryResponse{
		Id:        inventory.Id,
		Code:      inventory.Code,
		Commodity: inventory.Commodity,
		Variant:   inventory.Variant,
		Quantity:  inventory.Quantity,
		Price:     inventory.Price,
		Supplier:  inventory.Supplier,
		CreatedAt: &inventory.CreatedAt,
		UpdatedAt: &inventory.UpdatedAt,
		Employee:  &employeeResponse,
	}
}

func ConvertToServiceResponse(service *domain.Service) response.ServiceResponse {
	return response.ServiceResponse{
		Id:        service.Id,
		Code:      service.Code,
		Name:      service.Name,
		Price:     service.Price,
		Duration:  service.Duration,
		CreatedAt: service.CreatedAt,
		UpdatedAt: service.UpdatedAt,
	}
}

func ConvertToPromoResponse(promo *domain.Promo) response.PromoResponse {
	return response.PromoResponse{
		Id:          promo.Id,
		Code:        promo.Code,
		Name:        promo.Name,
		Discount:    promo.Discount,
		Description: promo.Description,
		Status:      promo.Status,
		Photo:       promo.Photo,
		Start:       &promo.Start,
		End:         &promo.End,
	}
}

func ConvertToUserResponse(user *domain.User, address *response.AddressResponse) response.UserResponse {
	return response.UserResponse{
		Id:              user.Id,
		Level:           user.Level,
		FullName:        user.FullName,
		Gender:          user.Gender,
		Password:        user.Password,
		Email:           user.Email,
		TelephoneNumber: user.TelephoneNumber,
		Photo:           user.Photo,
		EmailVerifiedAt: user.EmailVerifiedAt,
		CreatedAt:       &user.CreatedAt,
		UpdatedAt:       &user.UpdatedAt,
		Address:         address,
	}
}

func ConvertToAdminResponse(admin *domain.Admin, address *response.AddressResponse, branchName *string) response.AdminResponse {
	return response.AdminResponse{
		Id:              admin.Id,
		Code:            admin.Code,
		Level:           admin.Level,
		FullName:        admin.FullName,
		Password:        admin.Password,
		Email:           admin.Email,
		TelephoneNumber: admin.TelephoneNumber,
		Photo:           admin.Photo,
		Status:          admin.Status,
		CreatedAt:       &admin.CreatedAt,
		UpdatedAt:       &admin.UpdatedAt,
		Address:         address,
		BranchName:      *branchName,
	}
}

func ConvertToOrderResponse(order *domain.Order, inventory []*response.InventoryResponse, promo []*response.PromoResponse, userName, employeeName, serviceName *string) response.OrderResponse {
	return response.OrderResponse{
		Id:                order.Id,
		UserName:          userName,
		EmployeeName:      employeeName,
		ServiceName:       serviceName,
		Code:              order.Code,
		Type:              order.Type,
		Price:             order.Price,
		Weight:            order.Weight,
		Payment:           order.Payment,
		Description:       order.Description,
		Status:            order.Status,
		Entry:             &order.Entry,
		Out:               &order.Out,
		InventoryResponse: inventory,
		PromoResponse:     promo,
		CreatedAt:         &order.CreatedAt,
		UpdatedAt:         &order.UpdatedAt,
	}
}
