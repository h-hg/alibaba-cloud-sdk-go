package oceanbasepro

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

// DescribeOasSlowSQLList invokes the oceanbasepro.DescribeOasSlowSQLList API synchronously
func (client *Client) DescribeOasSlowSQLList(request *DescribeOasSlowSQLListRequest) (response *DescribeOasSlowSQLListResponse, err error) {
	response = CreateDescribeOasSlowSQLListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOasSlowSQLListWithChan invokes the oceanbasepro.DescribeOasSlowSQLList API asynchronously
func (client *Client) DescribeOasSlowSQLListWithChan(request *DescribeOasSlowSQLListRequest) (<-chan *DescribeOasSlowSQLListResponse, <-chan error) {
	responseChan := make(chan *DescribeOasSlowSQLListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOasSlowSQLList(request)
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

// DescribeOasSlowSQLListWithCallback invokes the oceanbasepro.DescribeOasSlowSQLList API asynchronously
func (client *Client) DescribeOasSlowSQLListWithCallback(request *DescribeOasSlowSQLListRequest, callback func(response *DescribeOasSlowSQLListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOasSlowSQLListResponse
		var err error
		defer close(result)
		response, err = client.DescribeOasSlowSQLList(request)
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

// DescribeOasSlowSQLListRequest is the request struct for api DescribeOasSlowSQLList
type DescribeOasSlowSQLListRequest struct {
	*requests.RpcRequest
	StartTime       string           `position:"Body" name:"StartTime"`
	SearchRule      string           `position:"Body" name:"SearchRule"`
	MergeDynamicSql requests.Boolean `position:"Body" name:"MergeDynamicSql"`
	DynamicSql      requests.Boolean `position:"Body" name:"DynamicSql"`
	SqlTextLength   requests.Integer `position:"Body" name:"SqlTextLength"`
	TenantId        string           `position:"Body" name:"TenantId"`
	SearchValue     string           `position:"Body" name:"SearchValue"`
	SqlId           string           `position:"Body" name:"SqlId"`
	FilterCondition string           `position:"Body" name:"FilterCondition"`
	SearchParam     string           `position:"Body" name:"SearchParam"`
	EndTime         string           `position:"Body" name:"EndTime"`
	NodeIp          string           `position:"Body" name:"NodeIp"`
	InstanceId      string           `position:"Body" name:"InstanceId"`
	DbName          string           `position:"Body" name:"DbName"`
	SearchKeyWord   string           `position:"Body" name:"SearchKeyWord"`
	AcceptLanguage  string           `position:"Body" name:"AcceptLanguage"`
}

// DescribeOasSlowSQLListResponse is the response struct for api DescribeOasSlowSQLList
type DescribeOasSlowSQLListResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Data      []DataItemInDescribeOasSlowSQLList `json:"Data" xml:"Data"`
}

// CreateDescribeOasSlowSQLListRequest creates a request to invoke DescribeOasSlowSQLList API
func CreateDescribeOasSlowSQLListRequest() (request *DescribeOasSlowSQLListRequest) {
	request = &DescribeOasSlowSQLListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeOasSlowSQLList", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOasSlowSQLListResponse creates a response to parse from DescribeOasSlowSQLList response
func CreateDescribeOasSlowSQLListResponse() (response *DescribeOasSlowSQLListResponse) {
	response = &DescribeOasSlowSQLListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
