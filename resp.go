package metactapi

type SubscribeList struct {
	Data []struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

type OnlyStatusResp struct {
	Success bool `json:"success"`
}
