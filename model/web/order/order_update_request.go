package order

type UpdateRequestOrder struct {
	Id            int    `json:"id,omitempty"`
	IdUser        int    `json:"id_user,omitempty"`
	IdEmployee    int    `json:"id_employee,omitempty"`
	IdService     int    `json:"id_service,omitempty"`
	IdPromos      []int  `json:"id_promos,omitempty"`
	IdInventories []int  `json:"id_inventories,omitempty"`
	Type          string `json:"type,omitempty"`
	Price         int    `json:"price,omitempty"`
	Weight        int    `json:"weight,omitempty"`
	Payment       string `json:"payment,omitempty"`
	Description   string `json:"description,omitempty"`
	Status        string `json:"status,omitempty"`
	Entry         string `json:"entry,omitempty"`
	Out           string `json:"out,omitempty"`
}
