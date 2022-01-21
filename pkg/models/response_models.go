package models

// swagger:model
type CurrentResponse struct {
	//current
	Number int `json:"current"`

	//error
	Err error `json:"error,omitempty"`
}


// swagger:model
type NextResponse struct {
	//current
	Number int `json:"next"`

	//error
	Err error `json:"error,omitempty"`
}


// swagger:model
type PreviousResponse struct {
	//current
	Number int `json:"previous"`

	//error
	Err error `json:"error,omitempty"`
}
