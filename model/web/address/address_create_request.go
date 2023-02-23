package address

type CreateRequestAddress struct {
	BuildingName        string `json:"building_name,omitempty" validate:"required,max=255,min=10"`
	BuildingNumber      int    `json:"building_number,omitempty" validate:"required"`
	Street              string `json:"street,omitempty" validate:"required,max=255,min=10"`
	Village             string `json:"village,omitempty" validate:"required,max=255,min=10"`
	NeighbourhoodNumber int    `json:"neighbourhood_number,omitempty" validate:"required"`
	HamletNumber        int    `json:"hamlet_number,omitempty" validate:"required"`
	SubDistrict         string `json:"sub_district,omitempty" validate:"required,max=255,min=10"`
	District            string `json:"district,omitempty" validate:"required,max=255,min=10"`
	Province            string `json:"province,omitempty" validate:"required,max=255,min=10"`
	PostalCode          int    `json:"postal_code,omitempty" validate:"required"`
	Description         string `json:"description,omitempty" validate:"max=255,min=10"`
}
