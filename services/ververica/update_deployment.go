package ververica

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

// UpdateDeployment invokes the ververica.UpdateDeployment API synchronously
func (client *Client) UpdateDeployment(request *UpdateDeploymentRequest) (response *UpdateDeploymentResponse, err error) {
	response = CreateUpdateDeploymentResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDeploymentWithChan invokes the ververica.UpdateDeployment API asynchronously
func (client *Client) UpdateDeploymentWithChan(request *UpdateDeploymentRequest) (<-chan *UpdateDeploymentResponse, <-chan error) {
	responseChan := make(chan *UpdateDeploymentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDeployment(request)
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

// UpdateDeploymentWithCallback invokes the ververica.UpdateDeployment API asynchronously
func (client *Client) UpdateDeploymentWithCallback(request *UpdateDeploymentRequest, callback func(response *UpdateDeploymentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDeploymentResponse
		var err error
		defer close(result)
		response, err = client.UpdateDeployment(request)
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

// UpdateDeploymentRequest is the request struct for api UpdateDeployment
type UpdateDeploymentRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	ParamsJson   string `position:"Body" name:"paramsJson"`
	DeploymentId string `position:"Path" name:"deploymentId"`
	Namespace    string `position:"Path" name:"namespace"`
}

// UpdateDeploymentResponse is the response struct for api UpdateDeployment
type UpdateDeploymentResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      int    `json:"data" xml:"data"`
}

// CreateUpdateDeploymentRequest creates a request to invoke UpdateDeployment API
func CreateUpdateDeploymentRequest() (request *UpdateDeploymentRequest) {
	request = &UpdateDeploymentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "UpdateDeployment", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]/patch", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateDeploymentResponse creates a response to parse from UpdateDeployment response
func CreateUpdateDeploymentResponse() (response *UpdateDeploymentResponse) {
	response = &UpdateDeploymentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
