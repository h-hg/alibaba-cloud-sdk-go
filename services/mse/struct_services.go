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

// Services is a nested struct in mse response
type Services struct {
	Name                  string         `json:"Name" xml:"Name"`
	ServiceNameInRegistry string         `json:"ServiceNameInRegistry" xml:"ServiceNameInRegistry"`
	GatewayId             int64          `json:"GatewayId" xml:"GatewayId"`
	MetaInfo              string         `json:"MetaInfo" xml:"MetaInfo"`
	HealehStatus          string         `json:"HealehStatus" xml:"HealehStatus"`
	GmtModified           string         `json:"GmtModified" xml:"GmtModified"`
	ServicePort           int64          `json:"ServicePort" xml:"ServicePort"`
	Id                    int64          `json:"Id" xml:"Id"`
	GroupName             string         `json:"GroupName" xml:"GroupName"`
	GatewayUniqueId       string         `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	SourceId              int64          `json:"SourceId" xml:"SourceId"`
	GmtCreate             string         `json:"GmtCreate" xml:"GmtCreate"`
	SourceType            string         `json:"SourceType" xml:"SourceType"`
	ServiceProtocol       string         `json:"ServiceProtocol" xml:"ServiceProtocol"`
	Namespace             string         `json:"Namespace" xml:"Namespace"`
	Ips                   []string       `json:"Ips" xml:"Ips"`
	Versions              []VersionsItem `json:"Versions" xml:"Versions"`
}
