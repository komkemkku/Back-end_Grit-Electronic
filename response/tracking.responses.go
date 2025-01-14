package response

type TrackingResponse struct {
	TrackingNumber string `json:"tracking_number"`
	Status         string `json:"status"`
	Location       string `json:"location"`
	Timestamp      string `json:"timestamp"`
}