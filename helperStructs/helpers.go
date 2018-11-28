package helpers

type Vcn struct {
	CidrBlock   *string `json:"cidrBlock"`
	DisplayName *string `json:"displayName"`
}

type Instance struct {
	DisplayName        *string `json:"displayName"`
	Shape              *string `json:"shape"`
	AvailabilityDomain *string `json:"availabilityDomain"`
}
