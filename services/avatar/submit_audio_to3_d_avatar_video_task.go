package avatar

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

// SubmitAudioTo3DAvatarVideoTask invokes the avatar.SubmitAudioTo3DAvatarVideoTask API synchronously
func (client *Client) SubmitAudioTo3DAvatarVideoTask(request *SubmitAudioTo3DAvatarVideoTaskRequest) (response *SubmitAudioTo3DAvatarVideoTaskResponse, err error) {
	response = CreateSubmitAudioTo3DAvatarVideoTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAudioTo3DAvatarVideoTaskWithChan invokes the avatar.SubmitAudioTo3DAvatarVideoTask API asynchronously
func (client *Client) SubmitAudioTo3DAvatarVideoTaskWithChan(request *SubmitAudioTo3DAvatarVideoTaskRequest) (<-chan *SubmitAudioTo3DAvatarVideoTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitAudioTo3DAvatarVideoTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAudioTo3DAvatarVideoTask(request)
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

// SubmitAudioTo3DAvatarVideoTaskWithCallback invokes the avatar.SubmitAudioTo3DAvatarVideoTask API asynchronously
func (client *Client) SubmitAudioTo3DAvatarVideoTaskWithCallback(request *SubmitAudioTo3DAvatarVideoTaskRequest, callback func(response *SubmitAudioTo3DAvatarVideoTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAudioTo3DAvatarVideoTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitAudioTo3DAvatarVideoTask(request)
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

// SubmitAudioTo3DAvatarVideoTaskRequest is the request struct for api SubmitAudioTo3DAvatarVideoTask
type SubmitAudioTo3DAvatarVideoTaskRequest struct {
	*requests.RpcRequest
	App               SubmitAudioTo3DAvatarVideoTaskApp        `position:"Query" name:"App"  type:"Struct"`
	AvatarInfo        SubmitAudioTo3DAvatarVideoTaskAvatarInfo `position:"Query" name:"AvatarInfo"  type:"Struct"`
	Title             string                                   `position:"Query" name:"Title"`
	ExtParams         string                                   `position:"Query" name:"ExtParams"`
	Url               string                                   `position:"Query" name:"Url"`
	VideoInfo         SubmitAudioTo3DAvatarVideoTaskVideoInfo  `position:"Query" name:"VideoInfo"  type:"Struct"`
	CallbackParams    string                                   `position:"Query" name:"CallbackParams"`
	TenantId          requests.Integer                         `position:"Query" name:"TenantId"`
	Callback          requests.Boolean                         `position:"Query" name:"Callback"`
	ExtParamsCLS      string                                   `position:"Query" name:"ExtParams_CLS"`
	CallbackParamsCLS string                                   `position:"Query" name:"CallbackParams_CLS"`
}

// SubmitAudioTo3DAvatarVideoTaskApp is a repeated param struct in SubmitAudioTo3DAvatarVideoTaskRequest
type SubmitAudioTo3DAvatarVideoTaskApp struct {
	AppId string `name:"AppId"`
}

// SubmitAudioTo3DAvatarVideoTaskAvatarInfo is a repeated param struct in SubmitAudioTo3DAvatarVideoTaskRequest
type SubmitAudioTo3DAvatarVideoTaskAvatarInfo struct {
	Code         string `name:"Code"`
	Locate       string `name:"Locate"`
	Angle        string `name:"Angle"`
	IndustryCode string `name:"IndustryCode"`
}

// SubmitAudioTo3DAvatarVideoTaskVideoInfo is a repeated param struct in SubmitAudioTo3DAvatarVideoTaskRequest
type SubmitAudioTo3DAvatarVideoTaskVideoInfo struct {
	IsAlpha            string `name:"IsAlpha"`
	BackgroundImageUrl string `name:"BackgroundImageUrl"`
	Resolution         string `name:"Resolution"`
	AlphaFormat        string `name:"AlphaFormat"`
}

// SubmitAudioTo3DAvatarVideoTaskResponse is the response struct for api SubmitAudioTo3DAvatarVideoTask
type SubmitAudioTo3DAvatarVideoTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitAudioTo3DAvatarVideoTaskRequest creates a request to invoke SubmitAudioTo3DAvatarVideoTask API
func CreateSubmitAudioTo3DAvatarVideoTaskRequest() (request *SubmitAudioTo3DAvatarVideoTaskRequest) {
	request = &SubmitAudioTo3DAvatarVideoTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "SubmitAudioTo3DAvatarVideoTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitAudioTo3DAvatarVideoTaskResponse creates a response to parse from SubmitAudioTo3DAvatarVideoTask response
func CreateSubmitAudioTo3DAvatarVideoTaskResponse() (response *SubmitAudioTo3DAvatarVideoTaskResponse) {
	response = &SubmitAudioTo3DAvatarVideoTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
