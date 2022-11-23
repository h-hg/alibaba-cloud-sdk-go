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

// DeleteDeployment invokes the ververica.DeleteDeployment API synchronously
func (client *Client) DeleteDeployment(request *DeleteDeploymentRequest) (response *DeleteDeploymentResponse, err error) {
	response = CreateDeleteDeploymentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeploymentWithChan invokes the ververica.DeleteDeployment API asynchronously
func (client *Client) DeleteDeploymentWithChan(request *DeleteDeploymentRequest) (<-chan *DeleteDeploymentResponse, <-chan error) {
	responseChan := make(chan *DeleteDeploymentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDeployment(request)
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

// DeleteDeploymentWithCallback invokes the ververica.DeleteDeployment API asynchronously
func (client *Client) DeleteDeploymentWithCallback(request *DeleteDeploymentRequest, callback func(response *DeleteDeploymentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeploymentResponse
		var err error
		defer close(result)
		response, err = client.DeleteDeployment(request)
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

// DeleteDeploymentRequest is the request struct for api DeleteDeployment
type DeleteDeploymentRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	DeploymentId string `position:"Path" name:"deploymentId"`
	Namespace    string `position:"Path" name:"namespace"`
}

// DeleteDeploymentResponse is the response struct for api DeleteDeployment
type DeleteDeploymentResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateDeleteDeploymentRequest creates a request to invoke DeleteDeployment API
func CreateDeleteDeploymentRequest() (request *DeleteDeploymentRequest) {
	request = &DeleteDeploymentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "DeleteDeployment", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteDeploymentResponse creates a response to parse from DeleteDeployment response
func CreateDeleteDeploymentResponse() (response *DeleteDeploymentResponse) {
	response = &DeleteDeploymentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
