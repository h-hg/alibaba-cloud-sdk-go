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

// GetAvailabilityMetric invokes the sae.GetAvailabilityMetric API synchronously
func (client *Client) GetAvailabilityMetric(request *GetAvailabilityMetricRequest) (response *GetAvailabilityMetricResponse, err error) {
	response = CreateGetAvailabilityMetricResponse()
	err = client.DoAction(request, response)
	return
}

// GetAvailabilityMetricWithChan invokes the sae.GetAvailabilityMetric API asynchronously
func (client *Client) GetAvailabilityMetricWithChan(request *GetAvailabilityMetricRequest) (<-chan *GetAvailabilityMetricResponse, <-chan error) {
	responseChan := make(chan *GetAvailabilityMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAvailabilityMetric(request)
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

// GetAvailabilityMetricWithCallback invokes the sae.GetAvailabilityMetric API asynchronously
func (client *Client) GetAvailabilityMetricWithCallback(request *GetAvailabilityMetricRequest, callback func(response *GetAvailabilityMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAvailabilityMetricResponse
		var err error
		defer close(result)
		response, err = client.GetAvailabilityMetric(request)
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

// GetAvailabilityMetricRequest is the request struct for api GetAvailabilityMetric
type GetAvailabilityMetricRequest struct {
	*requests.RoaRequest
	Limit requests.Integer `position:"Path" name:"Limit"`
}

// GetAvailabilityMetricResponse is the response struct for api GetAvailabilityMetric
type GetAvailabilityMetricResponse struct {
	*responses.BaseResponse
	Message   string                  `json:"Message" xml:"Message"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Code      string                  `json:"Code" xml:"Code"`
	Success   bool                    `json:"Success" xml:"Success"`
	Data      []AvailabilityMetricDto `json:"Data" xml:"Data"`
}

// CreateGetAvailabilityMetricRequest creates a request to invoke GetAvailabilityMetric API
func CreateGetAvailabilityMetricRequest() (request *GetAvailabilityMetricRequest) {
	request = &GetAvailabilityMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "GetAvailabilityMetric", "/pop/v1/sam/getAvailabilityMetric", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetAvailabilityMetricResponse creates a response to parse from GetAvailabilityMetric response
func CreateGetAvailabilityMetricResponse() (response *GetAvailabilityMetricResponse) {
	response = &GetAvailabilityMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
