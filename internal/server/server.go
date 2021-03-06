/*
 * Copyright 2019 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"github.com/BBVA/kapow/internal/server/control"
	"github.com/BBVA/kapow/internal/server/data"
	"github.com/BBVA/kapow/internal/server/user"
)

// StartServer Starts one instance of each server in a goroutine and remains listening on a channel for trace events generated by them
func StartServer(controlBindAddr, dataBindAddr, userBindAddr string) {
	go control.Run(controlBindAddr)
	go data.Run(dataBindAddr)
	go user.Run(userBindAddr)

	// Wait for ever
	select {}
}
