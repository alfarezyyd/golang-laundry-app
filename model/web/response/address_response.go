package response

type AddressResponse struct {
	Id                  int    `json:"id,omitempty"`
	BuildingName        string `json:"building_name,omitempty"`
	BuildingNumber      int    `json:"building_number,omitempty"`
	Street              string `json:"street,omitempty"`
	Village             string `json:"village,omitempty"`
	NeighbourhoodNumber int    `json:"neighbourhood_number,omitempty"`
	HamletNumber        int    `json:"hamlet_number,omitempty"`
	SubDistrict         string `json:"sub_district,omitempty"`
	District            string `json:"district,omitempty"`
	Province            string `json:"province,omitempty"`
	PostalCode          int    `json:"postal_code,omitempty"`
	Description         string `json:"description,omitempty"`
}
