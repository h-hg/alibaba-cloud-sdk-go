package mse

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

// UpdateNacosService invokes the mse.UpdateNacosService API synchronously
func (client *Client) UpdateNacosService(request *UpdateNacosServiceRequest) (response *UpdateNacosServiceResponse, err error) {
	response = CreateUpdateNacosServiceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNacosServiceWithChan invokes the mse.UpdateNacosService API asynchronously
func (client *Client) UpdateNacosServiceWithChan(request *UpdateNacosServiceRequest) (<-chan *UpdateNacosServiceResponse, <-chan error) {
	responseChan := make(chan *UpdateNacosServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNacosService(request)
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

// UpdateNacosServiceWithCallback invokes the mse.UpdateNacosService API asynchronously
func (client *Client) UpdateNacosServiceWithCallback(request *UpdateNacosServiceRequest, callback func(response *UpdateNacosServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNacosServiceResponse
		var err error
		defer close(result)
		response, err = client.UpdateNacosService(request)
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

// UpdateNacosServiceRequest is the request struct for api UpdateNacosService
type UpdateNacosServiceRequest struct {
	*requests.RpcRequest
	ClusterId        string `position:"Query" name:"ClusterId"`
	GroupName        string `position:"Query" name:"GroupName"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	NamespaceId      string `position:"Query" name:"NamespaceId"`
	AcceptLanguage   string `position:"Query" name:"AcceptLanguage"`
	ServiceName      string `position:"Query" name:"ServiceName"`
	ProtectThreshold string `position:"Query" name:"ProtectThreshold"`
}

// UpdateNacosServiceResponse is the response struct for api UpdateNacosService
type UpdateNacosServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateUpdateNacosServiceRequest creates a request to invoke UpdateNacosService API
func CreateUpdateNacosServiceRequest() (request *UpdateNacosServiceRequest) {
	request = &UpdateNacosServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateNacosService", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateNacosServiceResponse creates a response to parse from UpdateNacosService response
func CreateUpdateNacosServiceResponse() (response *UpdateNacosServiceResponse) {
	response = &UpdateNacosServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
