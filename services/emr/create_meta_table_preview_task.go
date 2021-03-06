package emr

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

// CreateMetaTablePreviewTask invokes the emr.CreateMetaTablePreviewTask API synchronously
func (client *Client) CreateMetaTablePreviewTask(request *CreateMetaTablePreviewTaskRequest) (response *CreateMetaTablePreviewTaskResponse, err error) {
	response = CreateCreateMetaTablePreviewTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMetaTablePreviewTaskWithChan invokes the emr.CreateMetaTablePreviewTask API asynchronously
func (client *Client) CreateMetaTablePreviewTaskWithChan(request *CreateMetaTablePreviewTaskRequest) (<-chan *CreateMetaTablePreviewTaskResponse, <-chan error) {
	responseChan := make(chan *CreateMetaTablePreviewTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMetaTablePreviewTask(request)
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

// CreateMetaTablePreviewTaskWithCallback invokes the emr.CreateMetaTablePreviewTask API asynchronously
func (client *Client) CreateMetaTablePreviewTaskWithCallback(request *CreateMetaTablePreviewTaskRequest, callback func(response *CreateMetaTablePreviewTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMetaTablePreviewTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateMetaTablePreviewTask(request)
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

// CreateMetaTablePreviewTaskRequest is the request struct for api CreateMetaTablePreviewTask
type CreateMetaTablePreviewTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Password        string           `position:"Query" name:"Password"`
	TableId         string           `position:"Query" name:"TableId"`
	DatabaseId      string           `position:"Query" name:"DatabaseId"`
	User            string           `position:"Query" name:"User"`
}

// CreateMetaTablePreviewTaskResponse is the response struct for api CreateMetaTablePreviewTask
type CreateMetaTablePreviewTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateMetaTablePreviewTaskRequest creates a request to invoke CreateMetaTablePreviewTask API
func CreateCreateMetaTablePreviewTaskRequest() (request *CreateMetaTablePreviewTaskRequest) {
	request = &CreateMetaTablePreviewTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateMetaTablePreviewTask", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMetaTablePreviewTaskResponse creates a response to parse from CreateMetaTablePreviewTask response
func CreateCreateMetaTablePreviewTaskResponse() (response *CreateMetaTablePreviewTaskResponse) {
	response = &CreateMetaTablePreviewTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
