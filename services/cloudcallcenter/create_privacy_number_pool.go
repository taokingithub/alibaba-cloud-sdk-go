package cloudcallcenter

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

// CreatePrivacyNumberPool invokes the cloudcallcenter.CreatePrivacyNumberPool API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberpool.html
func (client *Client) CreatePrivacyNumberPool(request *CreatePrivacyNumberPoolRequest) (response *CreatePrivacyNumberPoolResponse, err error) {
	response = CreateCreatePrivacyNumberPoolResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrivacyNumberPoolWithChan invokes the cloudcallcenter.CreatePrivacyNumberPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePrivacyNumberPoolWithChan(request *CreatePrivacyNumberPoolRequest) (<-chan *CreatePrivacyNumberPoolResponse, <-chan error) {
	responseChan := make(chan *CreatePrivacyNumberPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrivacyNumberPool(request)
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

// CreatePrivacyNumberPoolWithCallback invokes the cloudcallcenter.CreatePrivacyNumberPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePrivacyNumberPoolWithCallback(request *CreatePrivacyNumberPoolRequest, callback func(response *CreatePrivacyNumberPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrivacyNumberPoolResponse
		var err error
		defer close(result)
		response, err = client.CreatePrivacyNumberPool(request)
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

// CreatePrivacyNumberPoolRequest is the request struct for api CreatePrivacyNumberPool
type CreatePrivacyNumberPoolRequest struct {
	*requests.RpcRequest
	ProviderId string `position:"Query" name:"ProviderId"`
	BizId      string `position:"Query" name:"BizId"`
	Name       string `position:"Query" name:"Name"`
	Type       string `position:"Query" name:"Type"`
}

// CreatePrivacyNumberPoolResponse is the response struct for api CreatePrivacyNumberPool
type CreatePrivacyNumberPoolResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreatePrivacyNumberPoolRequest creates a request to invoke CreatePrivacyNumberPool API
func CreateCreatePrivacyNumberPoolRequest() (request *CreatePrivacyNumberPoolRequest) {
	request = &CreatePrivacyNumberPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreatePrivacyNumberPool", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePrivacyNumberPoolResponse creates a response to parse from CreatePrivacyNumberPool response
func CreateCreatePrivacyNumberPoolResponse() (response *CreatePrivacyNumberPoolResponse) {
	response = &CreatePrivacyNumberPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}