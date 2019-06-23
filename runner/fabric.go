/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	pb "github.com/ennoo/fabric-client/grpc/proto"
	"github.com/ennoo/fabric-client/grpc/server"
	"github.com/ennoo/fabric-client/route"
	"github.com/ennoo/rivet"
	"google.golang.org/grpc"
	"net"
)

func main() {
	go httpListener()
	grpcListener()
}

func httpListener() {
	rivet.Initialize(false, false, false)
	// rivet.UseDiscovery(discovery.ComponentConsul, "127.0.0.1:8500", "test", "127.0.0.1", 8081)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine: rivet.SetupRouter(
			route.Config,
			route.Channel,
			route.Order,
			route.ChainCode,
		),
		DefaultPort: "19865",
	})
}

func grpcListener() {
	var (
		listener net.Listener
		err      error
	)
	//  创建server端监听端口
	if listener, err = net.Listen("tcp", ":19877"); nil != err {
		panic(err)
	}
	//  创建grpc的server
	rpcServer := grpc.NewServer()

	//  注册我们自定义的helloworld服务
	pb.RegisterLedgerConfigServer(rpcServer, &server.ConfigServer{})
	pb.RegisterLedgerChannelServer(rpcServer, &server.ChannelServer{})
	pb.RegisterLedgerChainCodeServer(rpcServer, &server.ChainCodeServer{})
	pb.RegisterLedgerServer(rpcServer, &server.LedgerServer{})

	//  启动grpc服务
	if err = rpcServer.Serve(listener); nil != err {
		panic(err)
	}
}
