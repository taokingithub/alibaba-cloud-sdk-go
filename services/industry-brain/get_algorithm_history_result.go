package industry_brain

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

// GetAlgorithmHistoryResult invokes the industry_brain.GetAlgorithmHistoryResult API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getalgorithmhistoryresult.html
func (client *Client) GetAlgorithmHistoryResult(request *GetAlgorithmHistoryResultRequest) (response *GetAlgorithmHistoryResultResponse, err error) {
	response = CreateGetAlgorithmHistoryResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetAlgorithmHistoryResultWithChan invokes the industry_brain.GetAlgorithmHistoryResult API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getalgorithmhistoryresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAlgorithmHistoryResultWithChan(request *GetAlgorithmHistoryResultRequest) (<-chan *GetAlgorithmHistoryResultResponse, <-chan error) {
	responseChan := make(chan *GetAlgorithmHistoryResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAlgorithmHistoryResult(request)
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

// GetAlgorithmHistoryResultWithCallback invokes the industry_brain.GetAlgorithmHistoryResult API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getalgorithmhistoryresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAlgorithmHistoryResultWithCallback(request *GetAlgorithmHistoryResultRequest, callback func(response *GetAlgorithmHistoryResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAlgorithmHistoryResultResponse
		var err error
		defer close(result)
		response, err = client.GetAlgorithmHistoryResult(request)
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

// GetAlgorithmHistoryResultRequest is the request struct for api GetAlgorithmHistoryResult
type GetAlgorithmHistoryResultRequest struct {
	*requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	ServiceId string `position:"Query" name:"ServiceId"`
}

// GetAlgorithmHistoryResultResponse is the response struct for api GetAlgorithmHistoryResult
type GetAlgorithmHistoryResultResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Code      string   `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	TraceId   string   `json:"TraceId" xml:"TraceId"`
	Data      []Region `json:"Data" xml:"Data"`
}

// CreateGetAlgorithmHistoryResultRequest creates a request to invoke GetAlgorithmHistoryResult API
func CreateGetAlgorithmHistoryResultRequest() (request *GetAlgorithmHistoryResultRequest) {
	request = &GetAlgorithmHistoryResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2019-06-30", "GetAlgorithmHistoryResult", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAlgorithmHistoryResultResponse creates a response to parse from GetAlgorithmHistoryResult response
func CreateGetAlgorithmHistoryResultResponse() (response *GetAlgorithmHistoryResultResponse) {
	response = &GetAlgorithmHistoryResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
