package live

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

// SetLiveStreamPreloadTasks invokes the live.SetLiveStreamPreloadTasks API synchronously
func (client *Client) SetLiveStreamPreloadTasks(request *SetLiveStreamPreloadTasksRequest) (response *SetLiveStreamPreloadTasksResponse, err error) {
	response = CreateSetLiveStreamPreloadTasksResponse()
	err = client.DoAction(request, response)
	return
}

// SetLiveStreamPreloadTasksWithChan invokes the live.SetLiveStreamPreloadTasks API asynchronously
func (client *Client) SetLiveStreamPreloadTasksWithChan(request *SetLiveStreamPreloadTasksRequest) (<-chan *SetLiveStreamPreloadTasksResponse, <-chan error) {
	responseChan := make(chan *SetLiveStreamPreloadTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLiveStreamPreloadTasks(request)
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

// SetLiveStreamPreloadTasksWithCallback invokes the live.SetLiveStreamPreloadTasks API asynchronously
func (client *Client) SetLiveStreamPreloadTasksWithCallback(request *SetLiveStreamPreloadTasksRequest, callback func(response *SetLiveStreamPreloadTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLiveStreamPreloadTasksResponse
		var err error
		defer close(result)
		response, err = client.SetLiveStreamPreloadTasks(request)
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

// SetLiveStreamPreloadTasksRequest is the request struct for api SetLiveStreamPreloadTasks
type SetLiveStreamPreloadTasksRequest struct {
	*requests.RpcRequest
	PreloadedStartTime string           `position:"Query" name:"PreloadedStartTime"`
	Area               string           `position:"Query" name:"Area"`
	PreloadedEndTime   string           `position:"Query" name:"PreloadedEndTime"`
	DomainName         string           `position:"Query" name:"DomainName"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	PlayUrl            string           `position:"Query" name:"PlayUrl"`
}

// SetLiveStreamPreloadTasksResponse is the response struct for api SetLiveStreamPreloadTasks
type SetLiveStreamPreloadTasksResponse struct {
	*responses.BaseResponse
	Status               string               `json:"Status" xml:"Status"`
	FailedURL            int                  `json:"FailedURL" xml:"FailedURL"`
	TotalURL             int                  `json:"TotalURL" xml:"TotalURL"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	SuccessURL           int                  `json:"SuccessURL" xml:"SuccessURL"`
	PreloadTasksMessages PreloadTasksMessages `json:"PreloadTasksMessages" xml:"PreloadTasksMessages"`
}

// CreateSetLiveStreamPreloadTasksRequest creates a request to invoke SetLiveStreamPreloadTasks API
func CreateSetLiveStreamPreloadTasksRequest() (request *SetLiveStreamPreloadTasksRequest) {
	request = &SetLiveStreamPreloadTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetLiveStreamPreloadTasks", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLiveStreamPreloadTasksResponse creates a response to parse from SetLiveStreamPreloadTasks response
func CreateSetLiveStreamPreloadTasksResponse() (response *SetLiveStreamPreloadTasksResponse) {
	response = &SetLiveStreamPreloadTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
