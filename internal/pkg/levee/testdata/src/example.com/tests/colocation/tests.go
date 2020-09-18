// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package colocation

import (
	"example.com/core"
)

func TestTaintIsPropagatedToOperandAfterNonArrayAlloc(s core.Source, ip *core.Innocuous) {
	i := core.Innocuous{}
	taintColocated(s, &i, ip)
	core.Sink(ip) // want "a source has reached a sink"
}

func taintColocated(a interface{}, i *core.Innocuous, c interface{}) interface{} {
	return nil
}