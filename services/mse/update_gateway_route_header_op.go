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

// UpdateGatewayRouteHeaderOp invokes the mse.UpdateGatewayRouteHeaderOp API synchronously
func (client *Client) UpdateGatewayRouteHeaderOp(request *UpdateGatewayRouteHeaderOpRequest) (response *UpdateGatewayRouteHeaderOpResponse, err error) {
	response = CreateUpdateGatewayRouteHeaderOpResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayRouteHeaderOpWithChan invokes the mse.UpdateGatewayRouteHeaderOp API asynchronously
func (client *Client) UpdateGatewayRouteHeaderOpWithChan(request *UpdateGatewayRouteHeaderOpRequest) (<-chan *UpdateGatewayRouteHeaderOpResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayRouteHeaderOpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayRouteHeaderOp(request)
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

// UpdateGatewayRouteHeaderOpWithCallback invokes the mse.UpdateGatewayRouteHeaderOp API asynchronously
func (client *Client) UpdateGatewayRouteHeaderOpWithCallback(request *UpdateGatewayRouteHeaderOpRequest, callback func(response *UpdateGatewayRouteHeaderOpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayRouteHeaderOpResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayRouteHeaderOp(request)
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

// UpdateGatewayRouteHeaderOpRequest is the request struct for api UpdateGatewayRouteHeaderOp
type UpdateGatewayRouteHeaderOpRequest struct {
	*requests.RpcRequest
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	HeaderOpJSON    string           `position:"Query" name:"HeaderOpJSON"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	Id              requests.Integer `position:"Query" name:"Id"`
	GatewayId       requests.Integer `position:"Query" name:"GatewayId"`
}

// UpdateGatewayRouteHeaderOpResponse is the response struct for api UpdateGatewayRouteHeaderOp
type UpdateGatewayRouteHeaderOpResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayRouteHeaderOpRequest creates a request to invoke UpdateGatewayRouteHeaderOp API
func CreateUpdateGatewayRouteHeaderOpRequest() (request *UpdateGatewayRouteHeaderOpRequest) {
	request = &UpdateGatewayRouteHeaderOpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayRouteHeaderOp", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayRouteHeaderOpResponse creates a response to parse from UpdateGatewayRouteHeaderOp response
func CreateUpdateGatewayRouteHeaderOpResponse() (response *UpdateGatewayRouteHeaderOpResponse) {
	response = &UpdateGatewayRouteHeaderOpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
