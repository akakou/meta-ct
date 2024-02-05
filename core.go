package metact

type MetaCT struct {
	AppId       string
	AccessToken string
}

func NewCT(appId string, accessToken string) *MetaCT {
	return &MetaCT{
		AppId:       appId,
		AccessToken: accessToken,
	}
}
