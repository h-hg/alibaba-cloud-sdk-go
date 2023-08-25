package oceanbasepro

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

// ListProjects invokes the oceanbasepro.ListProjects API synchronously
func (client *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
	response = CreateListProjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectsWithChan invokes the oceanbasepro.ListProjects API asynchronously
func (client *Client) ListProjectsWithChan(request *ListProjectsRequest) (<-chan *ListProjectsResponse, <-chan error) {
	responseChan := make(chan *ListProjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjects(request)
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

// ListProjectsWithCallback invokes the oceanbasepro.ListProjects API asynchronously
func (client *Client) ListProjectsWithCallback(request *ListProjectsRequest, callback func(response *ListProjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectsResponse
		var err error
		defer close(result)
		response, err = client.ListProjects(request)
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

// ListProjectsRequest is the request struct for api ListProjects
type ListProjectsRequest struct {
	*requests.RpcRequest
	SearchKey           string           `position:"Body" name:"SearchKey"`
	Type                string           `position:"Body" name:"Type"`
	VisibleSubProject   requests.Boolean `position:"Body" name:"VisibleSubProject"`
	PageNumber          requests.Integer `position:"Body" name:"PageNumber"`
	SinkEndpointTypes   *[]string        `position:"Body" name:"SinkEndpointTypes"  type:"Json"`
	PageSize            requests.Integer `position:"Body" name:"PageSize"`
	Order               string           `position:"Body" name:"Order"`
	SourceEndpointTypes *[]string        `position:"Body" name:"SourceEndpointTypes"  type:"Json"`
	SortField           string           `position:"Body" name:"SortField"`
	LabelIds            *[]string        `position:"Body" name:"LabelIds"  type:"Json"`
	Status              *[]string        `position:"Body" name:"Status"  type:"Json"`
}

// ListProjectsResponse is the response struct for api ListProjects
type ListProjectsResponse struct {
	*responses.BaseResponse
}

// CreateListProjectsRequest creates a request to invoke ListProjects API
func CreateListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ListProjects", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListProjectsResponse creates a response to parse from ListProjects response
func CreateListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
