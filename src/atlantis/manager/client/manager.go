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

func OutputRoleReply(reply *ManagerRoleReply) error {
	Log("-> Status: %s", reply.Status)
	if reply.Manager != nil {
		Log("-> Manager:")
		Log("->   Region:         %s", reply.Manager.Region)
		Log("->   Host:           %s", reply.Manager.Host)
		Log("->   Registry CName: %s", reply.Manager.RegistryCName)
		Log("->   Manager CName:  %s", reply.Manager.ManagerCName)
		Log("->   Roles:")
		for role, typeMap := range reply.Manager.Roles {
			Log("->     %s", role)
			for typeName, val := range typeMap {
				Log("->       %s : %t", typeName, val)
			}
		}
	}
	return Output(map[string]interface{}{"status": reply.Status, "manager": reply.Manager}, nil, nil)
}

type AddRoleCommand struct {
	Region string `short:"r" long:"region" description:"the region to add a role for"`
	Host   string `short:"H" long:"host" description:"the host to add a role for"`
	Role   string `short:"l" long:"role" description:"the role to add"`
	Type   string `short:"t" long:"type" description:"the type to add"`
}

func (c *AddRoleCommand) Execute(args []string) error {
	err := Init()
	if err != nil {
		return OutputError(err)
	}
	Log("Add Role...")
	args = ExtractArgs([]*string{&c.Region, &c.Host}, args)
	arg := ManagerRoleArg{ManagerAuthArg: dummyAuthArg, Region: c.Region, Host: c.Host, Role: c.Role, Type: c.Type}
	var reply ManagerRoleReply
	err = rpcClient.CallAuthed("AddRole", &arg, &reply)
	if err != nil {
		return OutputError(err)
	}
	return OutputRoleReply(&reply)
}

type RemoveRoleCommand struct {
	Region string `short:"r" long:"region" description:"the region to remove a role for"`
	Host   string `short:"H" long:"host" description:"the host to remove a role for"`
	Role   string `short:"l" long:"role" description:"the role to remove"`
	Type   string `short:"t" long:"type" description:"the type to remove"`
}

func (c *RemoveRoleCommand) Execute(args []string) error {
	err := Init()
	if err != nil {
		return OutputError(err)
	}
	Log("Remove Role...")
	args = ExtractArgs([]*string{&c.Region, &c.Host}, args)
	arg := ManagerRoleArg{ManagerAuthArg: dummyAuthArg, Region: c.Region, Host: c.Host, Role: c.Role, Type: c.Type}
	var reply ManagerRoleReply
	err = rpcClient.CallAuthed("RemoveRole", &arg, &reply)
	if err != nil {
		return OutputError(err)
	}
	return OutputRoleReply(&reply)
}

type HasRoleCommand struct {
	Region string `short:"r" long:"region" description:"the region to check a role for"`
	Host   string `short:"H" long:"host" description:"the host to check a role for"`
	Role   string `short:"l" long:"role" description:"the role to check"`
	Type   string `short:"t" long:"type" description:"the type to check"`
}

func (c *HasRoleCommand) Execute(args []string) error {
	err := Init()
	if err != nil {
		return OutputError(err)
	}
	Log("Has Role...")
	args = ExtractArgs([]*string{&c.Region, &c.Host}, args)
	arg := ManagerRoleArg{ManagerAuthArg: dummyAuthArg, Region: c.Region, Host: c.Host, Role: c.Role, Type: c.Type}
	var reply ManagerHasRoleReply
	err = rpcClient.CallAuthed("HasRole", &arg, &reply)
	if err != nil {
		return OutputError(err)
	}
	Log("-> Status: %s", reply.Status)
	Log("-> HasRole: %t", reply.HasRole)
	return Output(map[string]interface{}{"status": reply.Status, "hasRole": reply.HasRole}, nil, nil)
}
