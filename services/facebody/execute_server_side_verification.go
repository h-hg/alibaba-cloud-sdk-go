package facebody

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

// ExecuteServerSideVerification invokes the facebody.ExecuteServerSideVerification API synchronously
func (client *Client) ExecuteServerSideVerification(request *ExecuteServerSideVerificationRequest) (response *ExecuteServerSideVerificationResponse, err error) {
	response = CreateExecuteServerSideVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteServerSideVerificationWithChan invokes the facebody.ExecuteServerSideVerification API asynchronously
func (client *Client) ExecuteServerSideVerificationWithChan(request *ExecuteServerSideVerificationRequest) (<-chan *ExecuteServerSideVerificationResponse, <-chan error) {
	responseChan := make(chan *ExecuteServerSideVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteServerSideVerification(request)
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

// ExecuteServerSideVerificationWithCallback invokes the facebody.ExecuteServerSideVerification API asynchronously
func (client *Client) ExecuteServerSideVerificationWithCallback(request *ExecuteServerSideVerificationRequest, callback func(response *ExecuteServerSideVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteServerSideVerificationResponse
		var err error
		defer close(result)
		response, err = client.ExecuteServerSideVerification(request)
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

// ExecuteServerSideVerificationRequest is the request struct for api ExecuteServerSideVerification
type ExecuteServerSideVerificationRequest struct {
	*requests.RoaRequest
	FacialPictureData string `position:"Body" name:"facialPictureData"`
	SceneType         string `position:"Body" name:"sceneType"`
	CertificateNumber string `position:"Body" name:"certificateNumber"`
	CertificateName   string `position:"Body" name:"certificateName"`
	FacialPictureUrl  string `position:"Body" name:"facialPictureUrl"`
}

// ExecuteServerSideVerificationResponse is the response struct for api ExecuteServerSideVerification
type ExecuteServerSideVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateExecuteServerSideVerificationRequest creates a request to invoke ExecuteServerSideVerification API
func CreateExecuteServerSideVerificationRequest() (request *ExecuteServerSideVerificationRequest) {
	request = &ExecuteServerSideVerificationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("facebody", "2020-09-10", "ExecuteServerSideVerification", "/viapi/thirdparty/realperson/execServerSideVerification", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExecuteServerSideVerificationResponse creates a response to parse from ExecuteServerSideVerification response
func CreateExecuteServerSideVerificationResponse() (response *ExecuteServerSideVerificationResponse) {
	response = &ExecuteServerSideVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
