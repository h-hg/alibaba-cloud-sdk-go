package dbfs

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

// ListTagValues invokes the dbfs.ListTagValues API synchronously
func (client *Client) ListTagValues(request *ListTagValuesRequest) (response *ListTagValuesResponse, err error) {
	response = CreateListTagValuesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagValuesWithChan invokes the dbfs.ListTagValues API asynchronously
func (client *Client) ListTagValuesWithChan(request *ListTagValuesRequest) (<-chan *ListTagValuesResponse, <-chan error) {
	responseChan := make(chan *ListTagValuesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagValues(request)
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

// ListTagValuesWithCallback invokes the dbfs.ListTagValues API asynchronously
func (client *Client) ListTagValuesWithCallback(request *ListTagValuesRequest, callback func(response *ListTagValuesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagValuesResponse
		var err error
		defer close(result)
		response, err = client.ListTagValues(request)
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

// ListTagValuesRequest is the request struct for api ListTagValues
type ListTagValuesRequest struct {
	*requests.RpcRequest
	TagKey string `position:"Query" name:"TagKey"`
}

// ListTagValuesResponse is the response struct for api ListTagValues
type ListTagValuesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	TagValues []string `json:"TagValues" xml:"TagValues"`
}

// CreateListTagValuesRequest creates a request to invoke ListTagValues API
func CreateListTagValuesRequest() (request *ListTagValuesRequest) {
	request = &ListTagValuesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "ListTagValues", "", "")
	request.Method = requests.POST
	return
}

// CreateListTagValuesResponse creates a response to parse from ListTagValues response
func CreateListTagValuesResponse() (response *ListTagValuesResponse) {
	response = &ListTagValuesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
