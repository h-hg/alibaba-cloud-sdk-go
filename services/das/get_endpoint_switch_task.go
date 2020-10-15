package das

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

// GetEndpointSwitchTask invokes the das.GetEndpointSwitchTask API synchronously
func (client *Client) GetEndpointSwitchTask(request *GetEndpointSwitchTaskRequest) (response *GetEndpointSwitchTaskResponse, err error) {
	response = CreateGetEndpointSwitchTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetEndpointSwitchTaskWithChan invokes the das.GetEndpointSwitchTask API asynchronously
func (client *Client) GetEndpointSwitchTaskWithChan(request *GetEndpointSwitchTaskRequest) (<-chan *GetEndpointSwitchTaskResponse, <-chan error) {
	responseChan := make(chan *GetEndpointSwitchTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEndpointSwitchTask(request)
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

// GetEndpointSwitchTaskWithCallback invokes the das.GetEndpointSwitchTask API asynchronously
func (client *Client) GetEndpointSwitchTaskWithCallback(request *GetEndpointSwitchTaskRequest, callback func(response *GetEndpointSwitchTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEndpointSwitchTaskResponse
		var err error
		defer close(result)
		response, err = client.GetEndpointSwitchTask(request)
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

// GetEndpointSwitchTaskRequest is the request struct for api GetEndpointSwitchTask
type GetEndpointSwitchTaskRequest struct {
	*requests.RpcRequest
	SkipAuth  string `position:"Query" name:"skipAuth"`
	Context   string `position:"Query" name:"__context"`
	Signature string `position:"Query" name:"signature"`
	UserId    string `position:"Query" name:"UserId"`
	Uid       string `position:"Query" name:"Uid"`
	AccessKey string `position:"Query" name:"accessKey"`
	TaskId    string `position:"Query" name:"TaskId"`
	Timestamp string `position:"Query" name:"timestamp"`
}

// GetEndpointSwitchTaskResponse is the response struct for api GetEndpointSwitchTask
type GetEndpointSwitchTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Synchro   string `json:"Synchro" xml:"Synchro"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetEndpointSwitchTaskRequest creates a request to invoke GetEndpointSwitchTask API
func CreateGetEndpointSwitchTaskRequest() (request *GetEndpointSwitchTaskRequest) {
	request = &GetEndpointSwitchTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetEndpointSwitchTask", "hdm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEndpointSwitchTaskResponse creates a response to parse from GetEndpointSwitchTask response
func CreateGetEndpointSwitchTaskResponse() (response *GetEndpointSwitchTaskResponse) {
	response = &GetEndpointSwitchTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
