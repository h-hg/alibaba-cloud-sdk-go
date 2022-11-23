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

// ListDatabases invokes the ververica.ListDatabases API synchronously
func (client *Client) ListDatabases(request *ListDatabasesRequest) (response *ListDatabasesResponse, err error) {
	response = CreateListDatabasesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDatabasesWithChan invokes the ververica.ListDatabases API asynchronously
func (client *Client) ListDatabasesWithChan(request *ListDatabasesRequest) (<-chan *ListDatabasesResponse, <-chan error) {
	responseChan := make(chan *ListDatabasesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDatabases(request)
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

// ListDatabasesWithCallback invokes the ververica.ListDatabases API asynchronously
func (client *Client) ListDatabasesWithCallback(request *ListDatabasesRequest, callback func(response *ListDatabasesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDatabasesResponse
		var err error
		defer close(result)
		response, err = client.ListDatabases(request)
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

// ListDatabasesRequest is the request struct for api ListDatabases
type ListDatabasesRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Cat       string `position:"Path" name:"cat"`
	Namespace string `position:"Path" name:"namespace"`
}

// ListDatabasesResponse is the response struct for api ListDatabases
type ListDatabasesResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateListDatabasesRequest creates a request to invoke ListDatabases API
func CreateListDatabasesRequest() (request *ListDatabasesRequest) {
	request = &ListDatabasesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListDatabases", "/pop/workspaces/[workspace]/catalog/v1beta2/namespaces/[namespace]/catalogs/[cat]:listDatabases", "", "")
	request.Method = requests.GET
	return
}

// CreateListDatabasesResponse creates a response to parse from ListDatabases response
func CreateListDatabasesResponse() (response *ListDatabasesResponse) {
	response = &ListDatabasesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
