package lrg

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

// GetVersion invokes the lrg.GetVersion API synchronously
// api document: https://help.aliyun.com/api/lrg/getversion.html
func (client *Client) GetVersion(request *GetVersionRequest) (response *GetVersionResponse, err error) {
	response = CreateGetVersionResponse()
	err = client.DoAction(request, response)
	return
}

// GetVersionWithChan invokes the lrg.GetVersion API asynchronously
// api document: https://help.aliyun.com/api/lrg/getversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVersionWithChan(request *GetVersionRequest) (<-chan *GetVersionResponse, <-chan error) {
	responseChan := make(chan *GetVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVersion(request)
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

// GetVersionWithCallback invokes the lrg.GetVersion API asynchronously
// api document: https://help.aliyun.com/api/lrg/getversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVersionWithCallback(request *GetVersionRequest, callback func(response *GetVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVersionResponse
		var err error
		defer close(result)
		response, err = client.GetVersion(request)
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

// GetVersionRequest is the request struct for api GetVersion
type GetVersionRequest struct {
	*requests.RoaRequest
	VersionName string `position:"Path" name:"versionName"`
}

// GetVersionResponse is the response struct for api GetVersion
type GetVersionResponse struct {
	*responses.BaseResponse
	Code    int    `json:"code" xml:"code"`
	Success bool   `json:"success" xml:"success"`
	Message string `json:"message" xml:"message"`
	Data    Data   `json:"data" xml:"data"`
}

// CreateGetVersionRequest creates a request to invoke GetVersion API
func CreateGetVersionRequest() (request *GetVersionRequest) {
	request = &GetVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "GetVersion", "/api/v2/version/[versionName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetVersionResponse creates a response to parse from GetVersion response
func CreateGetVersionResponse() (response *GetVersionResponse) {
	response = &GetVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
