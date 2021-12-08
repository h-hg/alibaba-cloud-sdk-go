package dyplsapi

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

// QueryRecordFileDownloadUrl invokes the dyplsapi.QueryRecordFileDownloadUrl API synchronously
func (client *Client) QueryRecordFileDownloadUrl(request *QueryRecordFileDownloadUrlRequest) (response *QueryRecordFileDownloadUrlResponse, err error) {
	response = CreateQueryRecordFileDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecordFileDownloadUrlWithChan invokes the dyplsapi.QueryRecordFileDownloadUrl API asynchronously
func (client *Client) QueryRecordFileDownloadUrlWithChan(request *QueryRecordFileDownloadUrlRequest) (<-chan *QueryRecordFileDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *QueryRecordFileDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecordFileDownloadUrl(request)
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

// QueryRecordFileDownloadUrlWithCallback invokes the dyplsapi.QueryRecordFileDownloadUrl API asynchronously
func (client *Client) QueryRecordFileDownloadUrlWithCallback(request *QueryRecordFileDownloadUrlRequest, callback func(response *QueryRecordFileDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecordFileDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.QueryRecordFileDownloadUrl(request)
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

// QueryRecordFileDownloadUrlRequest is the request struct for api QueryRecordFileDownloadUrl
type QueryRecordFileDownloadUrlRequest struct {
	*requests.RpcRequest
	CallId               string           `position:"Query" name:"CallId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProductType          string           `position:"Query" name:"ProductType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CallTime             string           `position:"Query" name:"CallTime"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// QueryRecordFileDownloadUrlResponse is the response struct for api QueryRecordFileDownloadUrl
type QueryRecordFileDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Code        string `json:"Code" xml:"Code"`
	Message     string `json:"Message" xml:"Message"`
	DownloadUrl string `json:"DownloadUrl" xml:"DownloadUrl"`
}

// CreateQueryRecordFileDownloadUrlRequest creates a request to invoke QueryRecordFileDownloadUrl API
func CreateQueryRecordFileDownloadUrlRequest() (request *QueryRecordFileDownloadUrlRequest) {
	request = &QueryRecordFileDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "QueryRecordFileDownloadUrl", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryRecordFileDownloadUrlResponse creates a response to parse from QueryRecordFileDownloadUrl response
func CreateQueryRecordFileDownloadUrlResponse() (response *QueryRecordFileDownloadUrlResponse) {
	response = &QueryRecordFileDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
