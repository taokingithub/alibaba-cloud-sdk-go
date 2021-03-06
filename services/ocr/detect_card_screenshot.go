package ocr

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

// DetectCardScreenshot invokes the ocr.DetectCardScreenshot API synchronously
func (client *Client) DetectCardScreenshot(request *DetectCardScreenshotRequest) (response *DetectCardScreenshotResponse, err error) {
	response = CreateDetectCardScreenshotResponse()
	err = client.DoAction(request, response)
	return
}

// DetectCardScreenshotWithChan invokes the ocr.DetectCardScreenshot API asynchronously
func (client *Client) DetectCardScreenshotWithChan(request *DetectCardScreenshotRequest) (<-chan *DetectCardScreenshotResponse, <-chan error) {
	responseChan := make(chan *DetectCardScreenshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectCardScreenshot(request)
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

// DetectCardScreenshotWithCallback invokes the ocr.DetectCardScreenshot API asynchronously
func (client *Client) DetectCardScreenshotWithCallback(request *DetectCardScreenshotRequest, callback func(response *DetectCardScreenshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectCardScreenshotResponse
		var err error
		defer close(result)
		response, err = client.DetectCardScreenshot(request)
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

// DetectCardScreenshotRequest is the request struct for api DetectCardScreenshot
type DetectCardScreenshotRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// DetectCardScreenshotResponse is the response struct for api DetectCardScreenshot
type DetectCardScreenshotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectCardScreenshotRequest creates a request to invoke DetectCardScreenshot API
func CreateDetectCardScreenshotRequest() (request *DetectCardScreenshotRequest) {
	request = &DetectCardScreenshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "DetectCardScreenshot", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectCardScreenshotResponse creates a response to parse from DetectCardScreenshot response
func CreateDetectCardScreenshotResponse() (response *DetectCardScreenshotResponse) {
	response = &DetectCardScreenshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
