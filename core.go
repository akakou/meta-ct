package metactapi

type MetaCTApi struct {
	AppId       string
	AccessToken string
}

func NewCT(appId string, accessToken string) *MetaCTApi {
	return &MetaCTApi{
		AppId:       appId,
		AccessToken: accessToken,
	}
}
