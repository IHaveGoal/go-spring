/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package guava

import "github.com/go-spring/spring-base/fastdev"

func init() {
	fastdev.RegisterProtocol(fastdev.APCU, &proto{})
}

type proto struct{}

func (p *proto) ShouldDiff() bool {
	return true
}

func (p *proto) GetLabel(data string) string {
	return data[:4]
}

func (p *proto) FlatRequest(data string) (map[string]string, error) {
	return nil, nil
}

func (p *proto) FlatResponse(data string) (map[string]string, error) {
	return nil, nil
}
