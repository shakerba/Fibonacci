package models

// swagger:model
type CurrentResponse struct {
	//current
	Number uint64 `json:"current"`

	//error
	Err error `json:"error,omitempty"`
}


// swagger:model
type NextResponse struct {
	//current
	Number uint64 `json:"next"`

	//error
	Err error `json:"error,omitempty"`
}


// swagger:model
type PreviousResponse struct {
	//current
	Number uint64 `json:"previous"`

	//error
	Err error `json:"error,omitempty"`
}
