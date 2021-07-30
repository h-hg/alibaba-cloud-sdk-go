package dyvmsapi

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

// ReportVoipProblems invokes the dyvmsapi.ReportVoipProblems API synchronously
func (client *Client) ReportVoipProblems(request *ReportVoipProblemsRequest) (response *ReportVoipProblemsResponse, err error) {
	response = CreateReportVoipProblemsResponse()
	err = client.DoAction(request, response)
	return
}

// ReportVoipProblemsWithChan invokes the dyvmsapi.ReportVoipProblems API asynchronously
func (client *Client) ReportVoipProblemsWithChan(request *ReportVoipProblemsRequest) (<-chan *ReportVoipProblemsResponse, <-chan error) {
	responseChan := make(chan *ReportVoipProblemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportVoipProblems(request)
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

// ReportVoipProblemsWithCallback invokes the dyvmsapi.ReportVoipProblems API asynchronously
func (client *Client) ReportVoipProblemsWithCallback(request *ReportVoipProblemsRequest, callback func(response *ReportVoipProblemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportVoipProblemsResponse
		var err error
		defer close(result)
		response, err = client.ReportVoipProblems(request)
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

// ReportVoipProblemsRequest is the request struct for api ReportVoipProblems
type ReportVoipProblemsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Title                string           `position:"Query" name:"Title"`
	VoipId               string           `position:"Query" name:"VoipId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ChannelId            string           `position:"Query" name:"ChannelId"`
	Desc                 string           `position:"Query" name:"Desc"`
}

// ReportVoipProblemsResponse is the response struct for api ReportVoipProblems
type ReportVoipProblemsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Module    string `json:"Module" xml:"Module"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateReportVoipProblemsRequest creates a request to invoke ReportVoipProblems API
func CreateReportVoipProblemsRequest() (request *ReportVoipProblemsRequest) {
	request = &ReportVoipProblemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "ReportVoipProblems", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportVoipProblemsResponse creates a response to parse from ReportVoipProblems response
func CreateReportVoipProblemsResponse() (response *ReportVoipProblemsResponse) {
	response = &ReportVoipProblemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
