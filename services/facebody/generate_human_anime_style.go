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

// GenerateHumanAnimeStyle invokes the facebody.GenerateHumanAnimeStyle API synchronously
// api document: https://help.aliyun.com/api/facebody/generatehumananimestyle.html
func (client *Client) GenerateHumanAnimeStyle(request *GenerateHumanAnimeStyleRequest) (response *GenerateHumanAnimeStyleResponse, err error) {
	response = CreateGenerateHumanAnimeStyleResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateHumanAnimeStyleWithChan invokes the facebody.GenerateHumanAnimeStyle API asynchronously
// api document: https://help.aliyun.com/api/facebody/generatehumananimestyle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateHumanAnimeStyleWithChan(request *GenerateHumanAnimeStyleRequest) (<-chan *GenerateHumanAnimeStyleResponse, <-chan error) {
	responseChan := make(chan *GenerateHumanAnimeStyleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateHumanAnimeStyle(request)
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

// GenerateHumanAnimeStyleWithCallback invokes the facebody.GenerateHumanAnimeStyle API asynchronously
// api document: https://help.aliyun.com/api/facebody/generatehumananimestyle.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateHumanAnimeStyleWithCallback(request *GenerateHumanAnimeStyleRequest, callback func(response *GenerateHumanAnimeStyleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateHumanAnimeStyleResponse
		var err error
		defer close(result)
		response, err = client.GenerateHumanAnimeStyle(request)
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

// GenerateHumanAnimeStyleRequest is the request struct for api GenerateHumanAnimeStyle
type GenerateHumanAnimeStyleRequest struct {
	*requests.RpcRequest
	AlgoType string `position:"Query" name:"AlgoType"`
	ImageURL string `position:"Query" name:"ImageURL"`
}

// GenerateHumanAnimeStyleResponse is the response struct for api GenerateHumanAnimeStyle
type GenerateHumanAnimeStyleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGenerateHumanAnimeStyleRequest creates a request to invoke GenerateHumanAnimeStyle API
func CreateGenerateHumanAnimeStyleRequest() (request *GenerateHumanAnimeStyleRequest) {
	request = &GenerateHumanAnimeStyleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "GenerateHumanAnimeStyle", "", "")
	return
}

// CreateGenerateHumanAnimeStyleResponse creates a response to parse from GenerateHumanAnimeStyle response
func CreateGenerateHumanAnimeStyleResponse() (response *GenerateHumanAnimeStyleResponse) {
	response = &GenerateHumanAnimeStyleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
