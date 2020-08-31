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

// GetPointReckoningValue invokes the industry_brain.GetPointReckoningValue API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getpointreckoningvalue.html
func (client *Client) GetPointReckoningValue(request *GetPointReckoningValueRequest) (response *GetPointReckoningValueResponse, err error) {
	response = CreateGetPointReckoningValueResponse()
	err = client.DoAction(request, response)
	return
}

// GetPointReckoningValueWithChan invokes the industry_brain.GetPointReckoningValue API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getpointreckoningvalue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPointReckoningValueWithChan(request *GetPointReckoningValueRequest) (<-chan *GetPointReckoningValueResponse, <-chan error) {
	responseChan := make(chan *GetPointReckoningValueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPointReckoningValue(request)
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

// GetPointReckoningValueWithCallback invokes the industry_brain.GetPointReckoningValue API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getpointreckoningvalue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPointReckoningValueWithCallback(request *GetPointReckoningValueRequest, callback func(response *GetPointReckoningValueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPointReckoningValueResponse
		var err error
		defer close(result)
		response, err = client.GetPointReckoningValue(request)
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

// GetPointReckoningValueRequest is the request struct for api GetPointReckoningValue
type GetPointReckoningValueRequest struct {
	*requests.RpcRequest
	TimeType   string `position:"Query" name:"TimeType"`
	TemplateId string `position:"Query" name:"TemplateId"`
}

// GetPointReckoningValueResponse is the response struct for api GetPointReckoningValue
type GetPointReckoningValueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateGetPointReckoningValueRequest creates a request to invoke GetPointReckoningValue API
func CreateGetPointReckoningValueRequest() (request *GetPointReckoningValueRequest) {
	request = &GetPointReckoningValueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2019-06-30", "GetPointReckoningValue", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPointReckoningValueResponse creates a response to parse from GetPointReckoningValue response
func CreateGetPointReckoningValueResponse() (response *GetPointReckoningValueResponse) {
	response = &GetPointReckoningValueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
