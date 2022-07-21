package domain

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

// QueryDomainByInstanceId invokes the domain.QueryDomainByInstanceId API synchronously
func (client *Client) QueryDomainByInstanceId(request *QueryDomainByInstanceIdRequest) (response *QueryDomainByInstanceIdResponse, err error) {
	response = CreateQueryDomainByInstanceIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainByInstanceIdWithChan invokes the domain.QueryDomainByInstanceId API asynchronously
func (client *Client) QueryDomainByInstanceIdWithChan(request *QueryDomainByInstanceIdRequest) (<-chan *QueryDomainByInstanceIdResponse, <-chan error) {
	responseChan := make(chan *QueryDomainByInstanceIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainByInstanceId(request)
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

// QueryDomainByInstanceIdWithCallback invokes the domain.QueryDomainByInstanceId API asynchronously
func (client *Client) QueryDomainByInstanceIdWithCallback(request *QueryDomainByInstanceIdRequest, callback func(response *QueryDomainByInstanceIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainByInstanceIdResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainByInstanceId(request)
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

// QueryDomainByInstanceIdRequest is the request struct for api QueryDomainByInstanceId
type QueryDomainByInstanceIdRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDomainByInstanceIdResponse is the response struct for api QueryDomainByInstanceId
type QueryDomainByInstanceIdResponse struct {
	*responses.BaseResponse
	Email                        string                           `json:"Email" xml:"Email"`
	RegistrationDate             string                           `json:"RegistrationDate" xml:"RegistrationDate"`
	RegistrationDateLong         int64                            `json:"RegistrationDateLong" xml:"RegistrationDateLong"`
	RealNameStatus               string                           `json:"RealNameStatus" xml:"RealNameStatus"`
	Premium                      bool                             `json:"Premium" xml:"Premium"`
	DomainNameVerificationStatus string                           `json:"DomainNameVerificationStatus" xml:"DomainNameVerificationStatus"`
	ExpirationDateLong           int64                            `json:"ExpirationDateLong" xml:"ExpirationDateLong"`
	TransferOutStatus            string                           `json:"TransferOutStatus" xml:"TransferOutStatus"`
	ZhRegistrantOrganization     string                           `json:"ZhRegistrantOrganization" xml:"ZhRegistrantOrganization"`
	EmailVerificationClientHold  bool                             `json:"EmailVerificationClientHold" xml:"EmailVerificationClientHold"`
	EmailVerificationStatus      int                              `json:"EmailVerificationStatus" xml:"EmailVerificationStatus"`
	RegistrantOrganization       string                           `json:"RegistrantOrganization" xml:"RegistrantOrganization"`
	TransferProhibitionLock      string                           `json:"TransferProhibitionLock" xml:"TransferProhibitionLock"`
	DomainNameProxyService       bool                             `json:"DomainNameProxyService" xml:"DomainNameProxyService"`
	RegistrantType               string                           `json:"RegistrantType" xml:"RegistrantType"`
	RegistrantUpdatingStatus     string                           `json:"RegistrantUpdatingStatus" xml:"RegistrantUpdatingStatus"`
	RequestId                    string                           `json:"RequestId" xml:"RequestId"`
	DomainName                   string                           `json:"DomainName" xml:"DomainName"`
	InstanceId                   string                           `json:"InstanceId" xml:"InstanceId"`
	ZhRegistrantName             string                           `json:"ZhRegistrantName" xml:"ZhRegistrantName"`
	ExpirationDate               string                           `json:"ExpirationDate" xml:"ExpirationDate"`
	RegistrantName               string                           `json:"RegistrantName" xml:"RegistrantName"`
	UserId                       string                           `json:"UserId" xml:"UserId"`
	UpdateProhibitionLock        string                           `json:"UpdateProhibitionLock" xml:"UpdateProhibitionLock"`
	DomainGroupId                int64                            `json:"DomainGroupId" xml:"DomainGroupId"`
	Remark                       string                           `json:"Remark" xml:"Remark"`
	DomainGroupName              string                           `json:"DomainGroupName" xml:"DomainGroupName"`
	ExpirationDateStatus         string                           `json:"ExpirationDateStatus" xml:"ExpirationDateStatus"`
	ExpirationCurrDateDiff       int                              `json:"ExpirationCurrDateDiff" xml:"ExpirationCurrDateDiff"`
	DomainType                   string                           `json:"DomainType" xml:"DomainType"`
	DomainStatus                 string                           `json:"DomainStatus" xml:"DomainStatus"`
	DnsList                      DnsListInQueryDomainByInstanceId `json:"DnsList" xml:"DnsList"`
}

// CreateQueryDomainByInstanceIdRequest creates a request to invoke QueryDomainByInstanceId API
func CreateQueryDomainByInstanceIdRequest() (request *QueryDomainByInstanceIdRequest) {
	request = &QueryDomainByInstanceIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainByInstanceId", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDomainByInstanceIdResponse creates a response to parse from QueryDomainByInstanceId response
func CreateQueryDomainByInstanceIdResponse() (response *QueryDomainByInstanceIdResponse) {
	response = &QueryDomainByInstanceIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
