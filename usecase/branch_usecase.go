package usecase

import (
	"context"
	"golang-laundry-app/model/web/branch"
	"golang-laundry-app/model/web/response"
)

type BranchUsecase interface {
	FindAll(ctx context.Context) []response.BranchResponse
	FindById(ctx context.Context, branchId int) response.BranchResponse
	Create(ctx context.Context, branchCreateRequest *branch.CreateRequestBranch) response.BranchResponse
	Update(ctx context.Context, branchUpdateRequest *branch.UpdateRequestBranch) response.BranchResponse
	Delete(ctx context.Context, branchId int)
}
