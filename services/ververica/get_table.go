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

// GetTable invokes the ververica.GetTable API synchronously
func (client *Client) GetTable(request *GetTableRequest) (response *GetTableResponse, err error) {
	response = CreateGetTableResponse()
	err = client.DoAction(request, response)
	return
}

// GetTableWithChan invokes the ververica.GetTable API asynchronously
func (client *Client) GetTableWithChan(request *GetTableRequest) (<-chan *GetTableResponse, <-chan error) {
	responseChan := make(chan *GetTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTable(request)
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

// GetTableWithCallback invokes the ververica.GetTable API asynchronously
func (client *Client) GetTableWithCallback(request *GetTableRequest, callback func(response *GetTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTableResponse
		var err error
		defer close(result)
		response, err = client.GetTable(request)
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

// GetTableRequest is the request struct for api GetTable
type GetTableRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Database  string `position:"Query" name:"database"`
	Cat       string `position:"Path" name:"cat"`
	Namespace string `position:"Path" name:"namespace"`
	Table     string `position:"Query" name:"table"`
}

// GetTableResponse is the response struct for api GetTable
type GetTableResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateGetTableRequest creates a request to invoke GetTable API
func CreateGetTableRequest() (request *GetTableRequest) {
	request = &GetTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "GetTable", "/pop/workspaces/[workspace]/catalog/v1beta2/namespaces/[namespace]/catalogs/[cat]:getTable", "", "")
	request.Method = requests.GET
	return
}

// CreateGetTableResponse creates a response to parse from GetTable response
func CreateGetTableResponse() (response *GetTableResponse) {
	response = &GetTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
