package live

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

// DescribeLiveAsrConfig invokes the live.DescribeLiveAsrConfig API synchronously
func (client *Client) DescribeLiveAsrConfig(request *DescribeLiveAsrConfigRequest) (response *DescribeLiveAsrConfigResponse, err error) {
	response = CreateDescribeLiveAsrConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveAsrConfigWithChan invokes the live.DescribeLiveAsrConfig API asynchronously
func (client *Client) DescribeLiveAsrConfigWithChan(request *DescribeLiveAsrConfigRequest) (<-chan *DescribeLiveAsrConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveAsrConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveAsrConfig(request)
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

// DescribeLiveAsrConfigWithCallback invokes the live.DescribeLiveAsrConfig API asynchronously
func (client *Client) DescribeLiveAsrConfigWithCallback(request *DescribeLiveAsrConfigRequest, callback func(response *DescribeLiveAsrConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveAsrConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveAsrConfig(request)
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

// DescribeLiveAsrConfigRequest is the request struct for api DescribeLiveAsrConfig
type DescribeLiveAsrConfigRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveAsrConfigResponse is the response struct for api DescribeLiveAsrConfig
type DescribeLiveAsrConfigResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	LiveAsrConfig LiveAsrConfig `json:"LiveAsrConfig" xml:"LiveAsrConfig"`
}

// CreateDescribeLiveAsrConfigRequest creates a request to invoke DescribeLiveAsrConfig API
func CreateDescribeLiveAsrConfigRequest() (request *DescribeLiveAsrConfigRequest) {
	request = &DescribeLiveAsrConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveAsrConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveAsrConfigResponse creates a response to parse from DescribeLiveAsrConfig response
func CreateDescribeLiveAsrConfigResponse() (response *DescribeLiveAsrConfigResponse) {
	response = &DescribeLiveAsrConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}