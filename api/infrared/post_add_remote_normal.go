package infrared

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rafaelvdberg/tuya-cloud-sdk-go/api/common"
)

type PostAddRemoteNormalReq struct {
	DeviceID     string
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	BrandID      string `json:"brand_id"`
	BrandName    string `json:"brand_name"`
	RemoteIndex  string `json:"remote_index"`
	RemoteName   string `json:"remote_name"`
}

func (t *PostAddRemoteNormalReq) Method() string {
	return common.RequestPost
}

func (t *PostAddRemoteNormalReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/normal/add-remote", t.DeviceID)
}

func (t *PostAddRemoteNormalReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostAddRemoteNormal adds remote to device
func PostAddRemoteNormal(deviceID string, categoryId int, categoryName string, brandID int, brandName string, remoteIndex int, remoteName string) (*PostAddRemoteResponse, error) {
	a := &PostAddRemoteNormalReq{
		DeviceID:    deviceID,
		CategoryId:  strconv.Itoa(categoryId),
		BrandID:     strconv.Itoa(brandID),
		BrandName:   brandName,
		RemoteIndex: strconv.Itoa(remoteIndex),
		RemoteName:  remoteName,
	}
	resp := &PostAddRemoteResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}
