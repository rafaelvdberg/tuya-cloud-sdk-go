package device

import (
	"fmt"

	"github.com/rafaelvdberg/tuya-cloud-sdk-go/api/common"
)

type GetDeviceFactoryReq struct {
	DeviceID string
}

func (t *GetDeviceFactoryReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceFactoryReq) API() string {
	return fmt.Sprintf("/v1.0/devices/factory-infos?device_ids=%s", t.DeviceID)
}

// GetDeviceFactory Get the factory information based on the device id
func GetDeviceFactory(deviceID string) (*GetDeviceFactoryResponse, error) {
	a := &GetDeviceFactoryReq{DeviceID: deviceID}
	resp := &GetDeviceFactoryResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

//GetDeviceFactoryResponse -- response to API
type GetDeviceFactoryResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		ID         string `json:"id"`
		MacAddress string `json:"mac"`
		Serial     string `json:"sn"`
		UUID       string `json:"uuid"`
	} `json:"result"`
	//  error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
