package facebody

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

// CountCrowd invokes the facebody.CountCrowd API synchronously
// api document: https://help.aliyun.com/api/facebody/countcrowd.html
func (client *Client) CountCrowd(request *CountCrowdRequest) (response *CountCrowdResponse, err error) {
	response = CreateCountCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// CountCrowdWithChan invokes the facebody.CountCrowd API asynchronously
// api document: https://help.aliyun.com/api/facebody/countcrowd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CountCrowdWithChan(request *CountCrowdRequest) (<-chan *CountCrowdResponse, <-chan error) {
	responseChan := make(chan *CountCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountCrowd(request)
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

// CountCrowdWithCallback invokes the facebody.CountCrowd API asynchronously
// api document: https://help.aliyun.com/api/facebody/countcrowd.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CountCrowdWithCallback(request *CountCrowdRequest, callback func(response *CountCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountCrowdResponse
		var err error
		defer close(result)
		response, err = client.CountCrowd(request)
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

// CountCrowdRequest is the request struct for api CountCrowd
type CountCrowdRequest struct {
	*requests.RpcRequest
	IsShow   requests.Boolean `position:"Body" name:"IsShow"`
	ImageURL string           `position:"Body" name:"ImageURL"`
}

// CountCrowdResponse is the response struct for api CountCrowd
type CountCrowdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCountCrowdRequest creates a request to invoke CountCrowd API
func CreateCountCrowdRequest() (request *CountCrowdRequest) {
	request = &CountCrowdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "CountCrowd", "", "")
	return
}

// CreateCountCrowdResponse creates a response to parse from CountCrowd response
func CreateCountCrowdResponse() (response *CountCrowdResponse) {
	response = &CountCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
