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

// DataInExtractPedestrianFeatureAttribute is a nested struct in facebody response
type DataInExtractPedestrianFeatureAttribute struct {
	Hair            string    `json:"Hair" xml:"Hair"`
	UpperType       string    `json:"UpperType" xml:"UpperType"`
	UpperTypeScore  float64   `json:"UpperTypeScore" xml:"UpperTypeScore"`
	LowerColor      string    `json:"LowerColor" xml:"LowerColor"`
	QualityScore    float64   `json:"QualityScore" xml:"QualityScore"`
	Gender          string    `json:"Gender" xml:"Gender"`
	Feature         string    `json:"Feature" xml:"Feature"`
	UpperColorScore float64   `json:"UpperColorScore" xml:"UpperColorScore"`
	GenderScore     float64   `json:"GenderScore" xml:"GenderScore"`
	LowerColorScore float64   `json:"LowerColorScore" xml:"LowerColorScore"`
	ObjType         string    `json:"ObjType" xml:"ObjType"`
	LowerTypeScore  float64   `json:"LowerTypeScore" xml:"LowerTypeScore"`
	HairScore       float64   `json:"HairScore" xml:"HairScore"`
	UpperColor      string    `json:"UpperColor" xml:"UpperColor"`
	LowerType       string    `json:"LowerType" xml:"LowerType"`
	AgeScore        float64   `json:"AgeScore" xml:"AgeScore"`
	ObjTypeScore    float64   `json:"ObjTypeScore" xml:"ObjTypeScore"`
	Age             string    `json:"Age" xml:"Age"`
	Elements        []Element `json:"Elements" xml:"Elements"`
}
