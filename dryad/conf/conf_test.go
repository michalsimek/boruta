/*
 *  Copyright (c) 2017-2018 Samsung Electronics Co., Ltd All Rights Reserved
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License
 */

package conf_test

import (
	"bytes"
	"strings"

	"github.com/SamsungSLAV/boruta"
	. "github.com/SamsungSLAV/boruta/dryad/conf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conf", func() {
	marshaled := `listen_address = ":7175"
boruta_address = ""
ssh_address = ":22"
sdcard = "/dev/sdX"
stm_path = "/run/stm.socket"

[caps]

[user]
  name = "boruta-user"
  groups = []
`
	unmarshaled := &General{
		Address:   ":7175",
		SSHAdress: ":22",
		Caps:      boruta.Capabilities(map[string]string{}),
		User: &User{
			Name:   "boruta-user",
			Groups: []string{},
		},
		SDcard:    "/dev/sdX",
		STMsocket: "/run/stm.socket",
	}
	var g *General

	BeforeEach(func() {
		g = NewConf()
	})

	It("should initially have default configuration", func() {
		Expect(g).To(Equal(unmarshaled))
	})

	It("should encode default configuration", func() {
		var w bytes.Buffer
		g.Marshal(&w)
		result := w.String()
		Expect(result).ToNot(BeEmpty())
		Expect(result).To(Equal(marshaled))
	})

	It("should decode default configuration", func() {
		g = new(General)
		g.Unmarshal(strings.NewReader(marshaled))
		Expect(g).To(Equal(unmarshaled))
	})
})
