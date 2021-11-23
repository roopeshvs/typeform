package gotypeform

type UpdateRequestBody struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
