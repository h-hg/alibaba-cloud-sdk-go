package config

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

// ListDiscoveredResources invokes the config.ListDiscoveredResources API synchronously
func (client *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (response *ListDiscoveredResourcesResponse, err error) {
	response = CreateListDiscoveredResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDiscoveredResourcesWithChan invokes the config.ListDiscoveredResources API asynchronously
func (client *Client) ListDiscoveredResourcesWithChan(request *ListDiscoveredResourcesRequest) (<-chan *ListDiscoveredResourcesResponse, <-chan error) {
	responseChan := make(chan *ListDiscoveredResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDiscoveredResources(request)
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

// ListDiscoveredResourcesWithCallback invokes the config.ListDiscoveredResources API asynchronously
func (client *Client) ListDiscoveredResourcesWithCallback(request *ListDiscoveredResourcesRequest, callback func(response *ListDiscoveredResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDiscoveredResourcesResponse
		var err error
		defer close(result)
		response, err = client.ListDiscoveredResources(request)
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

// ListDiscoveredResourcesRequest is the request struct for api ListDiscoveredResources
type ListDiscoveredResourcesRequest struct {
	*requests.RpcRequest
	ResourceDeleted requests.Integer `position:"Query" name:"ResourceDeleted"`
	MultiAccount    requests.Boolean `position:"Query" name:"MultiAccount"`
	Regions         string           `position:"Query" name:"Regions"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ComplianceType  string           `position:"Query" name:"ComplianceType"`
	ResourceId      string           `position:"Query" name:"ResourceId"`
	ResourceTypes   string           `position:"Query" name:"ResourceTypes"`
	MemberId        requests.Integer `position:"Query" name:"MemberId"`
}

// ListDiscoveredResourcesResponse is the response struct for api ListDiscoveredResources
type ListDiscoveredResourcesResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	DiscoveredResourceProfiles DiscoveredResourceProfiles `json:"DiscoveredResourceProfiles" xml:"DiscoveredResourceProfiles"`
}

// CreateListDiscoveredResourcesRequest creates a request to invoke ListDiscoveredResources API
func CreateListDiscoveredResourcesRequest() (request *ListDiscoveredResourcesRequest) {
	request = &ListDiscoveredResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "ListDiscoveredResources", "", "")
	request.Method = requests.POST
	return
}

// CreateListDiscoveredResourcesResponse creates a response to parse from ListDiscoveredResources response
func CreateListDiscoveredResourcesResponse() (response *ListDiscoveredResourcesResponse) {
	response = &ListDiscoveredResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
