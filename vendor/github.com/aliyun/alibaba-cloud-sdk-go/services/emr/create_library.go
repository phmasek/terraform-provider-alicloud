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

// CreateLibrary invokes the emr.CreateLibrary API synchronously
func (client *Client) CreateLibrary(request *CreateLibraryRequest) (response *CreateLibraryResponse, err error) {
	response = CreateCreateLibraryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLibraryWithChan invokes the emr.CreateLibrary API asynchronously
func (client *Client) CreateLibraryWithChan(request *CreateLibraryRequest) (<-chan *CreateLibraryResponse, <-chan error) {
	responseChan := make(chan *CreateLibraryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLibrary(request)
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

// CreateLibraryWithCallback invokes the emr.CreateLibrary API asynchronously
func (client *Client) CreateLibraryWithCallback(request *CreateLibraryRequest, callback func(response *CreateLibraryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLibraryResponse
		var err error
		defer close(result)
		response, err = client.CreateLibrary(request)
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

// CreateLibraryRequest is the request struct for api CreateLibrary
type CreateLibraryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LibraryVersion  string           `position:"Query" name:"LibraryVersion"`
	Type            string           `position:"Query" name:"Type"`
	Scope           string           `position:"Query" name:"Scope"`
	Name            string           `position:"Query" name:"Name"`
	SourceType      string           `position:"Query" name:"SourceType"`
	Properties      string           `position:"Query" name:"Properties"`
	SourceLocation  string           `position:"Query" name:"SourceLocation"`
}

// CreateLibraryResponse is the response struct for api CreateLibrary
type CreateLibraryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateLibraryRequest creates a request to invoke CreateLibrary API
func CreateCreateLibraryRequest() (request *CreateLibraryRequest) {
	request = &CreateLibraryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateLibrary", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLibraryResponse creates a response to parse from CreateLibrary response
func CreateCreateLibraryResponse() (response *CreateLibraryResponse) {
	response = &CreateLibraryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
