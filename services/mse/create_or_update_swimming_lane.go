package mse

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

// CreateOrUpdateSwimmingLane invokes the mse.CreateOrUpdateSwimmingLane API synchronously
func (client *Client) CreateOrUpdateSwimmingLane(request *CreateOrUpdateSwimmingLaneRequest) (response *CreateOrUpdateSwimmingLaneResponse, err error) {
	response = CreateCreateOrUpdateSwimmingLaneResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateSwimmingLaneWithChan invokes the mse.CreateOrUpdateSwimmingLane API asynchronously
func (client *Client) CreateOrUpdateSwimmingLaneWithChan(request *CreateOrUpdateSwimmingLaneRequest) (<-chan *CreateOrUpdateSwimmingLaneResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateSwimmingLaneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateSwimmingLane(request)
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

// CreateOrUpdateSwimmingLaneWithCallback invokes the mse.CreateOrUpdateSwimmingLane API asynchronously
func (client *Client) CreateOrUpdateSwimmingLaneWithCallback(request *CreateOrUpdateSwimmingLaneRequest, callback func(response *CreateOrUpdateSwimmingLaneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateSwimmingLaneResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateSwimmingLane(request)
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

// CreateOrUpdateSwimmingLaneRequest is the request struct for api CreateOrUpdateSwimmingLane
type CreateOrUpdateSwimmingLaneRequest struct {
	*requests.RpcRequest
	Source         string                                  `position:"Query" name:"Source"`
	GmtModified    string                                  `position:"Query" name:"GmtModified"`
	UserId         string                                  `position:"Query" name:"UserId"`
	LicenseKey     string                                  `position:"Query" name:"LicenseKey"`
	EntryRule      string                                  `position:"Query" name:"EntryRule"`
	Enable         requests.Boolean                        `position:"Query" name:"Enable"`
	Id             requests.Integer                        `position:"Query" name:"Id"`
	Tag            string                                  `position:"Query" name:"Tag"`
	EntryRules     *[]CreateOrUpdateSwimmingLaneEntryRules `position:"Query" name:"EntryRules"  type:"Repeated"`
	GroupId        requests.Integer                        `position:"Query" name:"GroupId"`
	GmtCreate      string                                  `position:"Query" name:"GmtCreate"`
	EnableRules    requests.Boolean                        `position:"Query" name:"EnableRules"`
	Name           string                                  `position:"Query" name:"Name"`
	AcceptLanguage string                                  `position:"Query" name:"AcceptLanguage"`
	Status         requests.Integer                        `position:"Query" name:"Status"`
}

// CreateOrUpdateSwimmingLaneEntryRules is a repeated param struct in CreateOrUpdateSwimmingLaneRequest
type CreateOrUpdateSwimmingLaneEntryRules struct {
	RestItems *[]CreateOrUpdateSwimmingLaneEntryRulesRestItems `name:"RestItems" type:"Repeated"`
	Path      string                                           `name:"Path"`
	Condition string                                           `name:"Condition"`
	Enable    string                                           `name:"Enable"`
	Priority  string                                           `name:"Priority"`
}

// CreateOrUpdateSwimmingLaneEntryRulesRestItems is a repeated param struct in CreateOrUpdateSwimmingLaneRequest
type CreateOrUpdateSwimmingLaneEntryRulesRestItems struct {
	Datum     string    `name:"Datum"`
	Divisor   string    `name:"Divisor"`
	Rate      string    `name:"Rate"`
	NameList  *[]string `name:"NameList" type:"Repeated"`
	Name      string    `name:"Name"`
	Type      string    `name:"Type"`
	Cond      string    `name:"Cond"`
	Remainder string    `name:"Remainder"`
	Value     string    `name:"Value"`
	Operator  string    `name:"Operator"`
}

// CreateOrUpdateSwimmingLaneResponse is the response struct for api CreateOrUpdateSwimmingLane
type CreateOrUpdateSwimmingLaneResponse struct {
	*responses.BaseResponse
}

// CreateCreateOrUpdateSwimmingLaneRequest creates a request to invoke CreateOrUpdateSwimmingLane API
func CreateCreateOrUpdateSwimmingLaneRequest() (request *CreateOrUpdateSwimmingLaneRequest) {
	request = &CreateOrUpdateSwimmingLaneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateOrUpdateSwimmingLane", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOrUpdateSwimmingLaneResponse creates a response to parse from CreateOrUpdateSwimmingLane response
func CreateCreateOrUpdateSwimmingLaneResponse() (response *CreateOrUpdateSwimmingLaneResponse) {
	response = &CreateOrUpdateSwimmingLaneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
