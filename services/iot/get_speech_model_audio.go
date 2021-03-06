package iot

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

// GetSpeechModelAudio invokes the iot.GetSpeechModelAudio API synchronously
func (client *Client) GetSpeechModelAudio(request *GetSpeechModelAudioRequest) (response *GetSpeechModelAudioResponse, err error) {
	response = CreateGetSpeechModelAudioResponse()
	err = client.DoAction(request, response)
	return
}

// GetSpeechModelAudioWithChan invokes the iot.GetSpeechModelAudio API asynchronously
func (client *Client) GetSpeechModelAudioWithChan(request *GetSpeechModelAudioRequest) (<-chan *GetSpeechModelAudioResponse, <-chan error) {
	responseChan := make(chan *GetSpeechModelAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSpeechModelAudio(request)
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

// GetSpeechModelAudioWithCallback invokes the iot.GetSpeechModelAudio API asynchronously
func (client *Client) GetSpeechModelAudioWithCallback(request *GetSpeechModelAudioRequest, callback func(response *GetSpeechModelAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSpeechModelAudioResponse
		var err error
		defer close(result)
		response, err = client.GetSpeechModelAudio(request)
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

// GetSpeechModelAudioRequest is the request struct for api GetSpeechModelAudio
type GetSpeechModelAudioRequest struct {
	*requests.RpcRequest
	ApiProduct          string    `position:"Body" name:"ApiProduct"`
	ApiRevision         string    `position:"Body" name:"ApiRevision"`
	SpeechModelCodeList *[]string `position:"Body" name:"SpeechModelCodeList"  type:"Repeated"`
}

// GetSpeechModelAudioResponse is the response struct for api GetSpeechModelAudio
type GetSpeechModelAudioResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Success      bool                      `json:"Success" xml:"Success"`
	Code         string                    `json:"Code" xml:"Code"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetSpeechModelAudio `json:"Data" xml:"Data"`
}

// CreateGetSpeechModelAudioRequest creates a request to invoke GetSpeechModelAudio API
func CreateGetSpeechModelAudioRequest() (request *GetSpeechModelAudioRequest) {
	request = &GetSpeechModelAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetSpeechModelAudio", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSpeechModelAudioResponse creates a response to parse from GetSpeechModelAudio response
func CreateGetSpeechModelAudioResponse() (response *GetSpeechModelAudioResponse) {
	response = &GetSpeechModelAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
