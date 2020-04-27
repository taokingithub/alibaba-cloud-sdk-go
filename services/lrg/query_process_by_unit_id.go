package lrg

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

// QueryProcessByUnitId invokes the lrg.QueryProcessByUnitId API synchronously
// api document: https://help.aliyun.com/api/lrg/queryprocessbyunitid.html
func (client *Client) QueryProcessByUnitId(request *QueryProcessByUnitIdRequest) (response *QueryProcessByUnitIdResponse, err error) {
	response = CreateQueryProcessByUnitIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryProcessByUnitIdWithChan invokes the lrg.QueryProcessByUnitId API asynchronously
// api document: https://help.aliyun.com/api/lrg/queryprocessbyunitid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProcessByUnitIdWithChan(request *QueryProcessByUnitIdRequest) (<-chan *QueryProcessByUnitIdResponse, <-chan error) {
	responseChan := make(chan *QueryProcessByUnitIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryProcessByUnitId(request)
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

// QueryProcessByUnitIdWithCallback invokes the lrg.QueryProcessByUnitId API asynchronously
// api document: https://help.aliyun.com/api/lrg/queryprocessbyunitid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProcessByUnitIdWithCallback(request *QueryProcessByUnitIdRequest, callback func(response *QueryProcessByUnitIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryProcessByUnitIdResponse
		var err error
		defer close(result)
		response, err = client.QueryProcessByUnitId(request)
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

// QueryProcessByUnitIdRequest is the request struct for api QueryProcessByUnitId
type QueryProcessByUnitIdRequest struct {
	*requests.RoaRequest
	DeployUnitName string `position:"Path" name:"deployUnitName"`
	ProductName    string `position:"Path" name:"productName"`
}

// QueryProcessByUnitIdResponse is the response struct for api QueryProcessByUnitId
type QueryProcessByUnitIdResponse struct {
	*responses.BaseResponse
	Code    int                      `json:"code" xml:"code"`
	Success bool                     `json:"success" xml:"success"`
	Message string                   `json:"message" xml:"message"`
	Data    []map[string]interface{} `json:"data" xml:"data"`
}

// CreateQueryProcessByUnitIdRequest creates a request to invoke QueryProcessByUnitId API
func CreateQueryProcessByUnitIdRequest() (request *QueryProcessByUnitIdRequest) {
	request = &QueryProcessByUnitIdRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "QueryProcessByUnitId", "/api/v2/[productName]/deploy-unit/[deployUnitName]/process", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryProcessByUnitIdResponse creates a response to parse from QueryProcessByUnitId response
func CreateQueryProcessByUnitIdResponse() (response *QueryProcessByUnitIdResponse) {
	response = &QueryProcessByUnitIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
