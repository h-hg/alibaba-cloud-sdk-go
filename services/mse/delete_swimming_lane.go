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

// DeleteSwimmingLane invokes the mse.DeleteSwimmingLane API synchronously
func (client *Client) DeleteSwimmingLane(request *DeleteSwimmingLaneRequest) (response *DeleteSwimmingLaneResponse, err error) {
	response = CreateDeleteSwimmingLaneResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSwimmingLaneWithChan invokes the mse.DeleteSwimmingLane API asynchronously
func (client *Client) DeleteSwimmingLaneWithChan(request *DeleteSwimmingLaneRequest) (<-chan *DeleteSwimmingLaneResponse, <-chan error) {
	responseChan := make(chan *DeleteSwimmingLaneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSwimmingLane(request)
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

// DeleteSwimmingLaneWithCallback invokes the mse.DeleteSwimmingLane API asynchronously
func (client *Client) DeleteSwimmingLaneWithCallback(request *DeleteSwimmingLaneRequest, callback func(response *DeleteSwimmingLaneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSwimmingLaneResponse
		var err error
		defer close(result)
		response, err = client.DeleteSwimmingLane(request)
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

// DeleteSwimmingLaneRequest is the request struct for api DeleteSwimmingLane
type DeleteSwimmingLaneRequest struct {
	*requests.RpcRequest
	LaneId         requests.Integer `position:"Query" name:"LaneId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// DeleteSwimmingLaneResponse is the response struct for api DeleteSwimmingLane
type DeleteSwimmingLaneResponse struct {
	*responses.BaseResponse
}

// CreateDeleteSwimmingLaneRequest creates a request to invoke DeleteSwimmingLane API
func CreateDeleteSwimmingLaneRequest() (request *DeleteSwimmingLaneRequest) {
	request = &DeleteSwimmingLaneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteSwimmingLane", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSwimmingLaneResponse creates a response to parse from DeleteSwimmingLane response
func CreateDeleteSwimmingLaneResponse() (response *DeleteSwimmingLaneResponse) {
	response = &DeleteSwimmingLaneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
