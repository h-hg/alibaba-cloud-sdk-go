package hitsdb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetLindormInstance invokes the hitsdb.GetLindormInstance API synchronously
func (client *Client) GetLindormInstance(request *GetLindormInstanceRequest) (response *GetLindormInstanceResponse, err error) {
	response = CreateGetLindormInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetLindormInstanceWithChan invokes the hitsdb.GetLindormInstance API asynchronously
func (client *Client) GetLindormInstanceWithChan(request *GetLindormInstanceRequest) (<-chan *GetLindormInstanceResponse, <-chan error) {
	responseChan := make(chan *GetLindormInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLindormInstance(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetLindormInstanceWithCallback invokes the hitsdb.GetLindormInstance API asynchronously
func (client *Client) GetLindormInstanceWithCallback(request *GetLindormInstanceRequest, callback func(response *GetLindormInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLindormInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetLindormInstance(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetLindormInstanceRequest is the request struct for api GetLindormInstance
type GetLindormInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// GetLindormInstanceResponse is the response struct for api GetLindormInstance
type GetLindormInstanceResponse struct {
	*responses.BaseResponse
	VpcId                string   `json:"VpcId" xml:"VpcId"`
	VswitchId            string   `json:"VswitchId" xml:"VswitchId"`
	CreateTime           string   `json:"CreateTime" xml:"CreateTime"`
	PayType              string   `json:"PayType" xml:"PayType"`
	NetworkType          string   `json:"NetworkType" xml:"NetworkType"`
	ServiceType          string   `json:"ServiceType" xml:"ServiceType"`
	EnableKms            bool     `json:"EnableKms" xml:"EnableKms"`
	EnableML             bool     `json:"EnableML" xml:"EnableML"`
	DiskUsage            string   `json:"DiskUsage" xml:"DiskUsage"`
	DiskCategory         string   `json:"DiskCategory" xml:"DiskCategory"`
	RequestId            string   `json:"RequestId" xml:"RequestId"`
	ColdStorage          int      `json:"ColdStorage" xml:"ColdStorage"`
	ExpiredMilliseconds  int64    `json:"ExpiredMilliseconds" xml:"ExpiredMilliseconds"`
	EngineType           int      `json:"EngineType" xml:"EngineType"`
	ExpireTime           string   `json:"ExpireTime" xml:"ExpireTime"`
	AutoRenew            bool     `json:"AutoRenew" xml:"AutoRenew"`
	DeletionProtection   string   `json:"DeletionProtection" xml:"DeletionProtection"`
	InstanceStorage      string   `json:"InstanceStorage" xml:"InstanceStorage"`
	AliUid               int64    `json:"AliUid" xml:"AliUid"`
	InstanceId           string   `json:"InstanceId" xml:"InstanceId"`
	RegionId             string   `json:"RegionId" xml:"RegionId"`
	EnableFS             bool     `json:"EnableFS" xml:"EnableFS"`
	CreateMilliseconds   int64    `json:"CreateMilliseconds" xml:"CreateMilliseconds"`
	InstanceAlias        string   `json:"InstanceAlias" xml:"InstanceAlias"`
	EnableBDS            bool     `json:"EnableBDS" xml:"EnableBDS"`
	EnablePhoenix        bool     `json:"EnablePhoenix" xml:"EnablePhoenix"`
	DiskThreshold        string   `json:"DiskThreshold" xml:"DiskThreshold"`
	ZoneId               string   `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus       string   `json:"InstanceStatus" xml:"InstanceStatus"`
	EnableCompute        bool     `json:"EnableCompute" xml:"EnableCompute"`
	EnableSSL            bool     `json:"EnableSSL" xml:"EnableSSL"`
	EnableMLCtrl         bool     `json:"EnableMLCtrl" xml:"EnableMLCtrl"`
	EnableCdc            bool     `json:"EnableCdc" xml:"EnableCdc"`
	EnableStream         bool     `json:"EnableStream" xml:"EnableStream"`
	EnableLTS            bool     `json:"EnableLTS" xml:"EnableLTS"`
	EnableShs            bool     `json:"EnableShs" xml:"EnableShs"`
	EnableBlob           bool     `json:"EnableBlob" xml:"EnableBlob"`
	MaintainStartTime    string   `json:"MaintainStartTime" xml:"MaintainStartTime"`
	MaintainEndTime      string   `json:"MaintainEndTime" xml:"MaintainEndTime"`
	ResourceGroupId      string   `json:"ResourceGroupId" xml:"ResourceGroupId"`
	LocalCloudCategory   string   `json:"LocalCloudCategory" xml:"LocalCloudCategory"`
	LocalCloudStorage    int      `json:"LocalCloudStorage" xml:"LocalCloudStorage"`
	PrimaryZoneId        string   `json:"PrimaryZoneId" xml:"PrimaryZoneId"`
	StandbyZoneId        string   `json:"StandbyZoneId" xml:"StandbyZoneId"`
	ArbiterZoneId        string   `json:"ArbiterZoneId" xml:"ArbiterZoneId"`
	PrimaryVSwitchId     string   `json:"PrimaryVSwitchId" xml:"PrimaryVSwitchId"`
	StandbyVSwitchId     string   `json:"StandbyVSwitchId" xml:"StandbyVSwitchId"`
	ArbiterVSwitchId     string   `json:"ArbiterVSwitchId" xml:"ArbiterVSwitchId"`
	MultiZoneCombination string   `json:"MultiZoneCombination" xml:"MultiZoneCombination"`
	CoreDiskCategory     string   `json:"CoreDiskCategory" xml:"CoreDiskCategory"`
	CoreSpec             string   `json:"CoreSpec" xml:"CoreSpec"`
	CoreNum              int      `json:"CoreNum" xml:"CoreNum"`
	CoreSingleStorage    int      `json:"CoreSingleStorage" xml:"CoreSingleStorage"`
	LogDiskCategory      string   `json:"LogDiskCategory" xml:"LogDiskCategory"`
	LogSpec              string   `json:"LogSpec" xml:"LogSpec"`
	LogNum               int      `json:"LogNum" xml:"LogNum"`
	LogSingleStorage     int      `json:"LogSingleStorage" xml:"LogSingleStorage"`
	ArchVersion          string   `json:"ArchVersion" xml:"ArchVersion"`
	EngineList           []Engine `json:"EngineList" xml:"EngineList"`
}

// CreateGetLindormInstanceRequest creates a request to invoke GetLindormInstance API
func CreateGetLindormInstanceRequest() (request *GetLindormInstanceRequest) {
	request = &GetLindormInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "GetLindormInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLindormInstanceResponse creates a response to parse from GetLindormInstance response
func CreateGetLindormInstanceResponse() (response *GetLindormInstanceResponse) {
	response = &GetLindormInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
