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

// CreateLive invokes the aliyuncvc.CreateLive API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createlive.html
func (client *Client) CreateLive(request *CreateLiveRequest) (response *CreateLiveResponse, err error) {
	response = CreateCreateLiveResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLiveWithChan invokes the aliyuncvc.CreateLive API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createlive.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLiveWithChan(request *CreateLiveRequest) (<-chan *CreateLiveResponse, <-chan error) {
	responseChan := make(chan *CreateLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLive(request)
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

// CreateLiveWithCallback invokes the aliyuncvc.CreateLive API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/createlive.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateLiveWithCallback(request *CreateLiveRequest, callback func(response *CreateLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveResponse
		var err error
		defer close(result)
		response, err = client.CreateLive(request)
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

// CreateLiveRequest is the request struct for api CreateLive
type CreateLiveRequest struct {
	*requests.RpcRequest
	Memo             string           `position:"Body" name:"Memo"`
	UserId           string           `position:"Body" name:"UserId"`
	OpenPasswordFlag requests.Boolean `position:"Body" name:"OpenPasswordFlag"`
	Password         string           `position:"Body" name:"Password"`
	LiveName         string           `position:"Body" name:"LiveName"`
}

// CreateLiveResponse is the response struct for api CreateLive
type CreateLiveResponse struct {
	*responses.BaseResponse
	ErrorCode int      `json:"ErrorCode" xml:"ErrorCode"`
	Message   string   `json:"Message" xml:"Message"`
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	LiveInfo  LiveInfo `json:"LiveInfo" xml:"LiveInfo"`
}

// CreateCreateLiveRequest creates a request to invoke CreateLive API
func CreateCreateLiveRequest() (request *CreateLiveRequest) {
	request = &CreateLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "CreateLive", "aliyuncvc", "openAPI")
	return
}

// CreateCreateLiveResponse creates a response to parse from CreateLive response
func CreateCreateLiveResponse() (response *CreateLiveResponse) {
	response = &CreateLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
