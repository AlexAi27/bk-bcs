/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"fmt"
	"regexp"
)

// ValidateIDName validate normal ID or name
func ValidateIDName(content, kind string) (bool, string) {
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	isMatched := re.MatchString(content)
	if !isMatched {
		return isMatched, fmt.Sprintf(`%s must be form of ^[a-zA-Z0-9-_]+$`, kind)
	}
	return isMatched, ""
}

// ValidateIPv4Cidr validate ipv4 subnet cidr
func ValidateIPv4Cidr(cidr string) (bool, string) {
	re := regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?$`)
	isMatched := re.MatchString(cidr)
	if !isMatched {
		return isMatched, "subnet cidr is invalid"
	}
	return isMatched, ""
}
