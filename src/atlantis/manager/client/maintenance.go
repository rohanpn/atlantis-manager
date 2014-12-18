/* Copyright 2014 Ooyala, Inc. All rights reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is
 * distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package client

import (
	. "atlantis/manager/rpc/types"
)

type ContainerMaintenanceCommand struct {
	ContainerID string `short:"c" long:"container" description:"the container to set maintenance for"`
	Maintenance bool   `short:"m" long:"maintenance" description:"true to set maintenance mode"`
	Arg         ManagerContainerMaintenanceArg
	Reply       ManagerContainerMaintenanceReply
}

type IdleCommand struct {
	Properties string `noauth:"true"`
	Arg        ManagerIdleArg
	Reply      ManagerIdleReply
}
