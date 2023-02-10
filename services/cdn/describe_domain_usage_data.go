package cdn

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

// DescribeDomainUsageData invokes the cdn.DescribeDomainUsageData API synchronously
func (client *Client) DescribeDomainUsageData(request *DescribeDomainUsageDataRequest) (response *DescribeDomainUsageDataResponse, err error) {
	response = CreateDescribeDomainUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainUsageDataWithChan invokes the cdn.DescribeDomainUsageData API asynchronously
func (client *Client) DescribeDomainUsageDataWithChan(request *DescribeDomainUsageDataRequest) (<-chan *DescribeDomainUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainUsageData(request)
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

// DescribeDomainUsageDataWithCallback invokes the cdn.DescribeDomainUsageData API asynchronously
func (client *Client) DescribeDomainUsageDataWithCallback(request *DescribeDomainUsageDataRequest, callback func(response *DescribeDomainUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainUsageData(request)
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

// DescribeDomainUsageDataRequest is the request struct for api DescribeDomainUsageData
type DescribeDomainUsageDataRequest struct {
	*requests.RpcRequest
	Area         string `position:"Query" name:"Area"`
	Field        string `position:"Query" name:"Field"`
	DomainName   string `position:"Query" name:"DomainName"`
	EndTime      string `position:"Query" name:"EndTime"`
	Interval     string `position:"Query" name:"Interval"`
	StartTime    string `position:"Query" name:"StartTime"`
	Type         string `position:"Query" name:"Type"`
	DataProtocol string `position:"Query" name:"DataProtocol"`
}

// DescribeDomainUsageDataResponse is the response struct for api DescribeDomainUsageData
type DescribeDomainUsageDataResponse struct {
	*responses.BaseResponse
	EndTime              string               `json:"EndTime" xml:"EndTime"`
	Type                 string               `json:"Type" xml:"Type"`
	StartTime            string               `json:"StartTime" xml:"StartTime"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Area                 string               `json:"Area" xml:"Area"`
	DomainName           string               `json:"DomainName" xml:"DomainName"`
	DataInterval         string               `json:"DataInterval" xml:"DataInterval"`
	UsageDataPerInterval UsageDataPerInterval `json:"UsageDataPerInterval" xml:"UsageDataPerInterval"`
}

// CreateDescribeDomainUsageDataRequest creates a request to invoke DescribeDomainUsageData API
func CreateDescribeDomainUsageDataRequest() (request *DescribeDomainUsageDataRequest) {
	request = &DescribeDomainUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainUsageData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainUsageDataResponse creates a response to parse from DescribeDomainUsageData response
func CreateDescribeDomainUsageDataResponse() (response *DescribeDomainUsageDataResponse) {
	response = &DescribeDomainUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
