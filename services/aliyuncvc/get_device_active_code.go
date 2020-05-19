package aliyuncvc

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

// GetDeviceActiveCode invokes the aliyuncvc.GetDeviceActiveCode API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/getdeviceactivecode.html
func (client *Client) GetDeviceActiveCode(request *GetDeviceActiveCodeRequest) (response *GetDeviceActiveCodeResponse, err error) {
	response = CreateGetDeviceActiveCodeResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceActiveCodeWithChan invokes the aliyuncvc.GetDeviceActiveCode API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/getdeviceactivecode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceActiveCodeWithChan(request *GetDeviceActiveCodeRequest) (<-chan *GetDeviceActiveCodeResponse, <-chan error) {
	responseChan := make(chan *GetDeviceActiveCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceActiveCode(request)
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

// GetDeviceActiveCodeWithCallback invokes the aliyuncvc.GetDeviceActiveCode API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/getdeviceactivecode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceActiveCodeWithCallback(request *GetDeviceActiveCodeRequest, callback func(response *GetDeviceActiveCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceActiveCodeResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceActiveCode(request)
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

// GetDeviceActiveCodeRequest is the request struct for api GetDeviceActiveCode
type GetDeviceActiveCodeRequest struct {
	*requests.RpcRequest
	SN string `position:"Body" name:"SN"`
}

// GetDeviceActiveCodeResponse is the response struct for api GetDeviceActiveCode
type GetDeviceActiveCodeResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Devices   []Data `json:"Devices" xml:"Devices"`
}

// CreateGetDeviceActiveCodeRequest creates a request to invoke GetDeviceActiveCode API
func CreateGetDeviceActiveCodeRequest() (request *GetDeviceActiveCodeRequest) {
	request = &GetDeviceActiveCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "GetDeviceActiveCode", "aliyuncvc", "openAPI")
	return
}

// CreateGetDeviceActiveCodeResponse creates a response to parse from GetDeviceActiveCode response
func CreateGetDeviceActiveCodeResponse() (response *GetDeviceActiveCodeResponse) {
	response = &GetDeviceActiveCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
