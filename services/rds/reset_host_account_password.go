package rds

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

// ResetHostAccountPassword invokes the rds.ResetHostAccountPassword API synchronously
func (client *Client) ResetHostAccountPassword(request *ResetHostAccountPasswordRequest) (response *ResetHostAccountPasswordResponse, err error) {
	response = CreateResetHostAccountPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetHostAccountPasswordWithChan invokes the rds.ResetHostAccountPassword API asynchronously
func (client *Client) ResetHostAccountPasswordWithChan(request *ResetHostAccountPasswordRequest) (<-chan *ResetHostAccountPasswordResponse, <-chan error) {
	responseChan := make(chan *ResetHostAccountPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetHostAccountPassword(request)
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

// ResetHostAccountPasswordWithCallback invokes the rds.ResetHostAccountPassword API asynchronously
func (client *Client) ResetHostAccountPasswordWithCallback(request *ResetHostAccountPasswordRequest, callback func(response *ResetHostAccountPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetHostAccountPasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetHostAccountPassword(request)
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

// ResetHostAccountPasswordRequest is the request struct for api ResetHostAccountPassword
type ResetHostAccountPasswordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
}

// ResetHostAccountPasswordResponse is the response struct for api ResetHostAccountPassword
type ResetHostAccountPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetHostAccountPasswordRequest creates a request to invoke ResetHostAccountPassword API
func CreateResetHostAccountPasswordRequest() (request *ResetHostAccountPasswordRequest) {
	request = &ResetHostAccountPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ResetHostAccountPassword", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetHostAccountPasswordResponse creates a response to parse from ResetHostAccountPassword response
func CreateResetHostAccountPasswordResponse() (response *ResetHostAccountPasswordResponse) {
	response = &ResetHostAccountPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
