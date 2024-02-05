package metactapi

type MetaCTApi struct {
	AppId       int
	AccessToken string
}

func NewCT(appId int, credentials string) *MetaCTApi {
	return &MetaCTApi{
		AppId:       appId,
		AccessToken: credentials,
	}
}
