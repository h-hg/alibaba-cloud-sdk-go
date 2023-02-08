package sae

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

// GetChangeOrderMetric invokes the sae.GetChangeOrderMetric API synchronously
func (client *Client) GetChangeOrderMetric(request *GetChangeOrderMetricRequest) (response *GetChangeOrderMetricResponse, err error) {
	response = CreateGetChangeOrderMetricResponse()
	err = client.DoAction(request, response)
	return
}

// GetChangeOrderMetricWithChan invokes the sae.GetChangeOrderMetric API asynchronously
func (client *Client) GetChangeOrderMetricWithChan(request *GetChangeOrderMetricRequest) (<-chan *GetChangeOrderMetricResponse, <-chan error) {
	responseChan := make(chan *GetChangeOrderMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetChangeOrderMetric(request)
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

// GetChangeOrderMetricWithCallback invokes the sae.GetChangeOrderMetric API asynchronously
func (client *Client) GetChangeOrderMetricWithCallback(request *GetChangeOrderMetricRequest, callback func(response *GetChangeOrderMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetChangeOrderMetricResponse
		var err error
		defer close(result)
		response, err = client.GetChangeOrderMetric(request)
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

// GetChangeOrderMetricRequest is the request struct for api GetChangeOrderMetric
type GetChangeOrderMetricRequest struct {
	*requests.RoaRequest
	CreateTime string           `position:"Path" name:"CreateTime"`
	Limit      requests.Integer `position:"Path" name:"Limit"`
	OrderBy    string           `position:"Path" name:"OrderBy"`
}

// GetChangeOrderMetricResponse is the response struct for api GetChangeOrderMetric
type GetChangeOrderMetricResponse struct {
	*responses.BaseResponse
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Code      string                 `json:"Code" xml:"Code"`
	Success   bool                   `json:"Success" xml:"Success"`
	Data      []ChangeOrderMetricDto `json:"Data" xml:"Data"`
}

// CreateGetChangeOrderMetricRequest creates a request to invoke GetChangeOrderMetric API
func CreateGetChangeOrderMetricRequest() (request *GetChangeOrderMetricRequest) {
	request = &GetChangeOrderMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "GetChangeOrderMetric", "/pop/v1/sam/getChangeOrderMetric", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetChangeOrderMetricResponse creates a response to parse from GetChangeOrderMetric response
func CreateGetChangeOrderMetricResponse() (response *GetChangeOrderMetricResponse) {
	response = &GetChangeOrderMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
