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

// JoinDeviceMeeting invokes the aliyuncvc.JoinDeviceMeeting API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joindevicemeeting.html
func (client *Client) JoinDeviceMeeting(request *JoinDeviceMeetingRequest) (response *JoinDeviceMeetingResponse, err error) {
	response = CreateJoinDeviceMeetingResponse()
	err = client.DoAction(request, response)
	return
}

// JoinDeviceMeetingWithChan invokes the aliyuncvc.JoinDeviceMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joindevicemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinDeviceMeetingWithChan(request *JoinDeviceMeetingRequest) (<-chan *JoinDeviceMeetingResponse, <-chan error) {
	responseChan := make(chan *JoinDeviceMeetingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinDeviceMeeting(request)
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

// JoinDeviceMeetingWithCallback invokes the aliyuncvc.JoinDeviceMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joindevicemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinDeviceMeetingWithCallback(request *JoinDeviceMeetingRequest, callback func(response *JoinDeviceMeetingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinDeviceMeetingResponse
		var err error
		defer close(result)
		response, err = client.JoinDeviceMeeting(request)
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

// JoinDeviceMeetingRequest is the request struct for api JoinDeviceMeeting
type JoinDeviceMeetingRequest struct {
	*requests.RpcRequest
	Token       string `position:"Body" name:"Token"`
	Password    string `position:"Body" name:"Password"`
	MeetingCode string `position:"Body" name:"MeetingCode"`
	SN          string `position:"Body" name:"SN"`
}

// JoinDeviceMeetingResponse is the response struct for api JoinDeviceMeeting
type JoinDeviceMeetingResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Device    Device `json:"Device" xml:"Device"`
}

// CreateJoinDeviceMeetingRequest creates a request to invoke JoinDeviceMeeting API
func CreateJoinDeviceMeetingRequest() (request *JoinDeviceMeetingRequest) {
	request = &JoinDeviceMeetingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "JoinDeviceMeeting", "aliyuncvc", "openAPI")
	return
}

// CreateJoinDeviceMeetingResponse creates a response to parse from JoinDeviceMeeting response
func CreateJoinDeviceMeetingResponse() (response *JoinDeviceMeetingResponse) {
	response = &JoinDeviceMeetingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
