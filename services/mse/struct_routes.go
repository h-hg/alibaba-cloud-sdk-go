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

// Routes is a nested struct in mse response
type Routes struct {
	Id                 int64               `json:"Id" xml:"Id"`
	Name               string              `json:"Name" xml:"Name"`
	GatewayUniqueId    string              `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	GatewayId          int64               `json:"GatewayId" xml:"GatewayId"`
	Status             int                 `json:"Status" xml:"Status"`
	RouteOrder         int                 `json:"RouteOrder" xml:"RouteOrder"`
	DefaultServiceName string              `json:"DefaultServiceName" xml:"DefaultServiceName"`
	DefaultServiceId   int64               `json:"DefaultServiceId" xml:"DefaultServiceId"`
	Predicates         string              `json:"Predicates" xml:"Predicates"`
	GmtCreate          string              `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified        string              `json:"GmtModified" xml:"GmtModified"`
	Services           string              `json:"Services" xml:"Services"`
	DomainId           int64               `json:"DomainId" xml:"DomainId"`
	DomainName         string              `json:"DomainName" xml:"DomainName"`
	DestinationType    string              `json:"DestinationType" xml:"DestinationType"`
	DomainIdList       []int64             `json:"DomainIdList" xml:"DomainIdList"`
	DomainNameList     []string            `json:"DomainNameList" xml:"DomainNameList"`
	RoutePredicates    RoutePredicates     `json:"RoutePredicates" xml:"RoutePredicates"`
	DirectResponse     DirectResponse      `json:"DirectResponse" xml:"DirectResponse"`
	Redirect           Redirect            `json:"Redirect" xml:"Redirect"`
	RouteServices      []RouteServicesItem `json:"RouteServices" xml:"RouteServices"`
}
