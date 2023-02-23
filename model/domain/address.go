package domain

type Address struct {
	Id                  int
	BuildingName        string
	BuildingNumber      int
	Street              string
	Village             string
	NeighbourhoodNumber int
	HamletNumber        int
	SubDistrict         string
	District            string
	Province            string
	PostalCode          int
	Description         string
}
