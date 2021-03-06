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

// QuerySpeechModelPushDeviceJob invokes the iot.QuerySpeechModelPushDeviceJob API synchronously
func (client *Client) QuerySpeechModelPushDeviceJob(request *QuerySpeechModelPushDeviceJobRequest) (response *QuerySpeechModelPushDeviceJobResponse, err error) {
	response = CreateQuerySpeechModelPushDeviceJobResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySpeechModelPushDeviceJobWithChan invokes the iot.QuerySpeechModelPushDeviceJob API asynchronously
func (client *Client) QuerySpeechModelPushDeviceJobWithChan(request *QuerySpeechModelPushDeviceJobRequest) (<-chan *QuerySpeechModelPushDeviceJobResponse, <-chan error) {
	responseChan := make(chan *QuerySpeechModelPushDeviceJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySpeechModelPushDeviceJob(request)
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

// QuerySpeechModelPushDeviceJobWithCallback invokes the iot.QuerySpeechModelPushDeviceJob API asynchronously
func (client *Client) QuerySpeechModelPushDeviceJobWithCallback(request *QuerySpeechModelPushDeviceJobRequest, callback func(response *QuerySpeechModelPushDeviceJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySpeechModelPushDeviceJobResponse
		var err error
		defer close(result)
		response, err = client.QuerySpeechModelPushDeviceJob(request)
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

// QuerySpeechModelPushDeviceJobRequest is the request struct for api QuerySpeechModelPushDeviceJob
type QuerySpeechModelPushDeviceJobRequest struct {
	*requests.RpcRequest
	PageId      requests.Integer `position:"Body" name:"PageId"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	JobCode     string           `position:"Body" name:"JobCode"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
	DeviceName  string           `position:"Body" name:"DeviceName"`
	Status      string           `position:"Body" name:"Status"`
}

// QuerySpeechModelPushDeviceJobResponse is the response struct for api QuerySpeechModelPushDeviceJob
type QuerySpeechModelPushDeviceJobResponse struct {
	*responses.BaseResponse
	RequestId    string                              `json:"RequestId" xml:"RequestId"`
	Success      bool                                `json:"Success" xml:"Success"`
	Code         string                              `json:"Code" xml:"Code"`
	ErrorMessage string                              `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySpeechModelPushDeviceJob `json:"Data" xml:"Data"`
}

// CreateQuerySpeechModelPushDeviceJobRequest creates a request to invoke QuerySpeechModelPushDeviceJob API
func CreateQuerySpeechModelPushDeviceJobRequest() (request *QuerySpeechModelPushDeviceJobRequest) {
	request = &QuerySpeechModelPushDeviceJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySpeechModelPushDeviceJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySpeechModelPushDeviceJobResponse creates a response to parse from QuerySpeechModelPushDeviceJob response
func CreateQuerySpeechModelPushDeviceJobResponse() (response *QuerySpeechModelPushDeviceJobResponse) {
	response = &QuerySpeechModelPushDeviceJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
