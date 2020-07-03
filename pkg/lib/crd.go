/*
 * Copyright 2019-2020 VMware, Inc.
 * All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package lib

import (
	akocrd "ako/pkg/client/clientset/versioned"
	akoinformer "ako/pkg/client/informers/externalversions/ako/v1alpha1"
)

var CRDClientset *akocrd.Clientset

// crd "ako/pkg/client/clientset/versioned/typed/ako/v1alpha1"
func SetCRDClientset(cs *akocrd.Clientset) {
	CRDClientset = cs
}

func GetCRDClientset() *akocrd.Clientset {
	return CRDClientset
}

var CRDInformers *AKOCrdInformers

type AKOCrdInformers struct {
	HostRuleInformer  akoinformer.HostRuleInformer
	HTTPRuleInformer akoinformer.HTTPRuleInformer
}

func SetCRDInformers(c *AKOCrdInformers) {
	CRDInformers = c
}

func GetCRDInformers() *AKOCrdInformers {
	return CRDInformers
}
