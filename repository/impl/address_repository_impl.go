package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-laundry-app/helper"
	"golang-laundry-app/model/domain"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepositoryImpl() *AddressRepositoryImpl {
	return &AddressRepositoryImpl{}
}

func (addressRepository *AddressRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, addressId int) (domain.Address, error) {
	SQL := "SELECT id, building_name, building_number, street, village, neighbourhood_number, hamlet_number, sub_district, district, province, postal_code, description FROM addresses WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, addressId)
	defer rows.Close()
	helper.PanicIfError(err)
	var addressData domain.Address
	if rows.Next() {
		err = rows.Scan(
			&addressData.Id,
			&addressData.BuildingName,
			&addressData.BuildingNumber,
			&addressData.Street,
			&addressData.Village,
			&addressData.NeighbourhoodNumber,
			&addressData.HamletNumber,
			&addressData.SubDistrict,
			&addressData.District,
			&addressData.Province,
			&addressData.PostalCode,
			&addressData.Description,
		)
		helper.PanicIfError(err)
		return addressData, nil
	} else {
		return addressData, errors.New("address not found")
	}
}

func (addressRepository *AddressRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, address *domain.Address) {
	SQL := "INSERT INTO addresses(building_name, building_number, street, village, neighbourhood_number, hamlet_number, sub_district, district, province, postal_code, description) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(
		ctx,
		SQL,
		address.BuildingName,
		address.BuildingNumber,
		address.Street,
		address.Village,
		address.NeighbourhoodNumber,
		address.HamletNumber,
		address.SubDistrict,
		address.District,
		address.Province,
		address.PostalCode,
		address.Description,
	)
	helper.PanicIfError(err)

	newAddressId, err := result.LastInsertId()
	helper.PanicIfError(err)

	address.Id = int(newAddressId)
}

func (addressRepository *AddressRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, address *domain.Address) {
	SQL := "UPDATE addresses SET building_name = ?, building_number = ?, street = ?, village = ?, neighbourhood_number = ?, hamlet_number = ?, sub_district = ?, district = ?, province = ?, postal_code = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		address.BuildingName,
		address.BuildingNumber,
		address.Street,
		address.Village,
		address.NeighbourhoodNumber,
		address.HamletNumber,
		address.SubDistrict,
		address.District,
		address.Province,
		address.PostalCode,
		address.Description,
		address.Id,
	)
	helper.PanicIfError(err)
}

func (addressRepository *AddressRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, addressId int) {
	SQL := "DELETE FROM addresses WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, addressId)
	helper.PanicIfError(err)
}
