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
 */

package sdk

import (
	"encoding/hex"
	"github.com/aberic/fabric-client/config"
	"github.com/aberic/fabric-client/geneses"
	"github.com/aberic/fabric-client/grpc/proto/generate"
	"github.com/aberic/gnomon"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/resource"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"golang.org/x/protobuf/proto"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

const (
	leagueDomain   = "league01.com"
	orderName      = "od"
	orderDomain    = "order.com"
	order0NodeName = "order0"
	order1NodeName = "order1"
	org1Name       = "org1"
	org2Name       = "org2"
	org3Name       = "org3"
	org1Domain     = "one.com"
	org2Domain     = "two.com"
	org3Domain     = "three.com"
	node1          = "node1"
	node2          = "node2"
	admin          = "Admin"
	user2          = "User2"
	channelID      = "mychannel01"

	cryptoType    = generate.CryptoType_ECDSA
	signAlgorithm = generate.SignAlgorithm_ECDSAWithSHA256
)

var (
	algorithm = &generate.ReqKeyConfig_EccAlgorithm{
		EccAlgorithm: generate.EccAlgorithm_p256,
	}
)

func TestPemConfig_GenerateCrypto(t *testing.T) {
	pemConfig := &geneses.PemConfig{KeyConfig: &generate.ReqKeyConfig{
		CryptoType: cryptoType,
		Algorithm:  algorithm,
	}}
	t.Log(pemConfig.GenerateCrypto())
}

func TestPemConfig_GenerateCryptos(t *testing.T) {
	for i := 0; i < 30; i++ {
		pemConfig := &geneses.PemConfig{KeyConfig: &generate.ReqKeyConfig{
			CryptoType: cryptoType,
			Algorithm:  algorithm,
		}}
		t.Log(pemConfig.GenerateCrypto())
	}
}

func TestGenerateConfig_CreateLeague(t *testing.T) {
	priData, err := ioutil.ReadFile("/tmp/354535000/pri.key")
	if nil != err {
		t.Error(err)
	}
	priTlsData, err := ioutil.ReadFile("/tmp/364725000/pri.key")
	if nil != err {
		t.Error(err)
	}
	gc := &geneses.GenerateConfig{}
	t.Log(gc.CreateLeague(&generate.ReqCreateLeague{
		Domain:     leagueDomain,
		PriData:    priData,
		PriTlsData: priTlsData,
		Csr: &generate.CSR{
			Country:      []string{"CN"},
			Organization: []string{"league01"},
			Locality:     []string{"Beijing"},
			Province:     []string{"Beijing"},
			CommonName:   leagueDomain,
		},
		SignAlgorithm: signAlgorithm,
	}))
}

func TestGenerateConfig_CreateOrg(t *testing.T) {
	gc := &geneses.GenerateConfig{}
	t.Log(gc.CreateOrg(&generate.ReqCreateOrg{
		OrgType:      generate.OrgType_Order,
		LeagueDomain: leagueDomain,
		Name:         orderName,
		Domain:       orderDomain,
	}))
	t.Log(gc.CreateOrg(&generate.ReqCreateOrg{
		OrgType:      generate.OrgType_Peer,
		LeagueDomain: leagueDomain,
		Name:         org1Name,
		Domain:       org1Domain,
	}))
	t.Log(gc.CreateOrg(&generate.ReqCreateOrg{
		OrgType:      generate.OrgType_Peer,
		LeagueDomain: leagueDomain,
		Name:         org2Name,
		Domain:       org2Domain,
	}))
	t.Log(gc.CreateOrg(&generate.ReqCreateOrg{
		OrgType:      generate.OrgType_Peer,
		LeagueDomain: leagueDomain,
		Name:         org3Name,
		Domain:       org3Domain,
	}))
}

func TestGenerateConfig_CAGenesisAddAffiliation(t *testing.T) {
	genesisAddAffiliation("admin", "adminpw", leagueDomain, orderDomain, orderName, geneses.MspID(orderName), "rootCA", orderName, "http://10.0.61.22:7054", t)
	genesisAddAffiliation("admin", "adminpw", leagueDomain, org1Domain, org1Name, geneses.MspID(org1Name), "rootCA", org1Name, "http://10.0.61.22:7054", t)
	genesisAddAffiliation("admin", "adminpw", leagueDomain, org2Domain, org2Name, geneses.MspID(org2Name), "rootCA", org2Name, "http://10.0.61.22:7054", t)
	genesisAddAffiliation("admin", "adminpw", leagueDomain, org3Domain, org3Name, geneses.MspID(org3Name), "rootCA", org3Name, "http://10.0.61.22:7054", t)
}

func TestGenerateConfig_CAGenesisRegister(t *testing.T) {
	// 获取各组织根CA用于启动fabric-ca
	genesisRegister("admin", "adminpw", leagueDomain, orderDomain, orderName, geneses.MspID(orderName), "rootCA", orderName, "orderer,user", "orderer,user", strings.Split(geneses.CertUserCAName(orderName, orderDomain, admin), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, orderDomain, orderName, geneses.MspID(orderName), "rootCA", orderName, "orderer,user", "orderer,user", strings.Split(geneses.CertNodeCAName(orderName, orderDomain, order0NodeName), "-")[0], "adminpw", "orderer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, orderDomain, orderName, geneses.MspID(orderName), "rootCA", orderName, "orderer,user", "orderer,user", strings.Split(geneses.CertNodeCAName(orderName, orderDomain, order1NodeName), "-")[0], "adminpw", "orderer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org1Domain, org1Name, geneses.MspID(org1Name), "rootCA", org1Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org1Name, org1Domain, admin), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org1Domain, org1Name, geneses.MspID(org1Name), "rootCA", org1Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org1Name, org1Domain, user2), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org1Domain, org1Name, geneses.MspID(org1Name), "rootCA", org1Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org1Name, org1Domain, node1), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org1Domain, org1Name, geneses.MspID(org1Name), "rootCA", org1Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org1Name, org1Domain, node2), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org2Domain, org2Name, geneses.MspID(org2Name), "rootCA", org2Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org2Name, org2Domain, admin), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org2Domain, org2Name, geneses.MspID(org2Name), "rootCA", org2Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org2Name, org2Domain, user2), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org2Domain, org2Name, geneses.MspID(org2Name), "rootCA", org2Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org2Name, org2Domain, node1), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org2Domain, org2Name, geneses.MspID(org2Name), "rootCA", org2Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org2Name, org2Domain, node2), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org3Domain, org3Name, geneses.MspID(org3Name), "rootCA", org3Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org3Name, org3Domain, admin), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org3Domain, org3Name, geneses.MspID(org3Name), "rootCA", org3Name, "client,peer,user", "peer,user", strings.Split(geneses.CertUserCAName(org3Name, org3Domain, user2), "-")[0], "adminpw", "user", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org3Domain, org3Name, geneses.MspID(org3Name), "rootCA", org3Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org3Name, org3Domain, node1), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
	genesisRegister("admin", "adminpw", leagueDomain, org3Domain, org3Name, geneses.MspID(org3Name), "rootCA", org3Name, "client,peer,user", "peer,user", strings.Split(geneses.CertNodeCAName(org3Name, org3Domain, node2), "-")[0], "adminpw", "peer", "http://10.0.61.22:7054", t)
}

func TestGenerateConfig_CreateCsr(t *testing.T) {
	t.Log(createCsr("/tmp/263926000/pri.key", orderName, orderDomain, admin, false, t))
	t.Log(createCsr("/tmp/275001000/pri.key", org1Name, org1Domain, admin, false, t))
	t.Log(createCsr("/tmp/275937000/pri.key", org1Name, org1Domain, user2, false, t))
	t.Log(createCsr("/tmp/276880000/pri.key", org2Name, org2Domain, admin, false, t))
	t.Log(createCsr("/tmp/277731000/pri.key", org2Name, org2Domain, user2, false, t))
	t.Log(createCsr("/tmp/186602000/pri.key", org3Name, org3Domain, admin, false, t))
	t.Log(createCsr("/tmp/198066000/pri.key", org3Name, org3Domain, user2, false, t))

	t.Log(createCsr("/tmp/278330000/pri.key", orderName, orderDomain, order0NodeName, true, t))
	t.Log(createCsr("/tmp/278916000/pri.key", orderName, orderDomain, order1NodeName, true, t))
	t.Log(createCsr("/tmp/279511000/pri.key", org1Name, org1Domain, node1, true, t))
	t.Log(createCsr("/tmp/280104000/pri.key", org1Name, org1Domain, node2, true, t))
	t.Log(createCsr("/tmp/280660000/pri.key", org2Name, org2Domain, node1, true, t))
	t.Log(createCsr("/tmp/281219000/pri.key", org2Name, org2Domain, node2, true, t))
	t.Log(createCsr("/tmp/198937000/pri.key", org3Name, org3Domain, node1, true, t))
	t.Log(createCsr("/tmp/199644000/pri.key", org3Name, org3Domain, node2, true, t))
}

func TestGenerateConfig_CreateOrgNode(t *testing.T) {
	t.Log(createOrgNode(generate.OrgType_Order, "/tmp/285000000/pub.key", orderName, orderDomain, order0NodeName, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Order, "/tmp/285711000/pub.key", orderName, orderDomain, order1NodeName, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/286275000/pub.key", org1Name, org1Domain, node1, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/286836000/pub.key", org1Name, org1Domain, node2, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/287420000/pub.key", org2Name, org2Domain, node1, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/287960000/pub.key", org2Name, org2Domain, node2, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/201468000/pub.key", org3Name, org3Domain, node1, "http://10.0.61.22:7054", t))
	t.Log(createOrgNode(generate.OrgType_Peer, "/tmp/202170000/pub.key", org3Name, org3Domain, node2, "http://10.0.61.22:7054", t))
}

func TestGenerateConfig_CreateOrgUser(t *testing.T) {
	t.Log(createOrgUser(generate.OrgType_Order, "/tmp/281845000/pub.key", orderName, orderDomain, admin, "http://10.0.61.22:7054", true, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/282485000/pub.key", org1Name, org1Domain, admin, "http://10.0.61.22:7054", true, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/283058000/pub.key", org1Name, org1Domain, user2, "http://10.0.61.22:7054", false, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/283628000/pub.key", org2Name, org2Domain, admin, "http://10.0.61.22:7054", true, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/284287000/pub.key", org2Name, org2Domain, user2, "http://10.0.61.22:7054", false, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/200170000/pub.key", org3Name, org3Domain, admin, "http://10.0.61.22:7054", true, t))
	t.Log(createOrgUser(generate.OrgType_Peer, "/tmp/200816000/pub.key", org3Name, org3Domain, user2, "http://10.0.61.22:7054", false, t))
}

func TestGenerateConfig_GenesisBlock(t *testing.T) {
	genesis := geneses.Genesis{
		Info: &generate.ReqGenesis{
			League: &generate.LeagueInBlock{
				Domain: leagueDomain,
				Addresses: []string{
					strings.Join([]string{order0NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
					strings.Join([]string{order1NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
				},
				BatchTimeout: 2,
				BatchSize: &generate.BatchSize{
					MaxMessageCount:   1000,
					AbsoluteMaxBytes:  10 * 1024 * 1024,
					PreferredMaxBytes: 2 * 1024 * 1024,
				},
				Kafka: &generate.Kafka{
					Brokers: []string{"kafka1:9090", "kafka2:9091", "kafka3:9092", "kafka4:9093"},
				},
				MaxChannels: 1000,
			},
			Orgs: []*generate.OrgInBlock{
				{Domain: orderDomain, Name: orderName, Type: generate.OrgType_Order},
				{Domain: org1Domain, Name: org1Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org1Name, org1Domain}, "."), Port: 7051},
				}},
				{Domain: org2Domain, Name: org2Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org2Name, org2Domain}, "."), Port: 7061},
				}},
			},
		},
	}
	genesis.Init()
	if err := genesis.CreateGenesisBlock("default"); nil != err {
		t.Error(err)
	} else {
		t.Log("create genesis block success")
	}
}

func TestGenerateConfig_InspectGenesisBlock(t *testing.T) {
	data, err := ioutil.ReadFile(geneses.GenesisBlockFilePath(leagueDomain))
	//data, err := ioutil.ReadFile("/Users/aberic/Documents/path/go/src/github.com/aberic/fabric-client/geneses/example/test/channel-artifacts/genesis.block")
	if nil != err {
		t.Error(err)
	}
	str, err := resource.InspectBlock(data)
	if nil != err {
		t.Error(err)
	}
	t.Log(str)
}

func TestGenerateConfig_CreateChannelTx(t *testing.T) {
	genesis := geneses.Genesis{
		Info: &generate.ReqGenesis{
			League: &generate.LeagueInBlock{
				Domain: leagueDomain,
				Addresses: []string{
					strings.Join([]string{order0NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
					strings.Join([]string{order1NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
				},
				BatchTimeout: 2,
				BatchSize: &generate.BatchSize{
					MaxMessageCount:   1000,
					AbsoluteMaxBytes:  10 * 1024 * 1024,
					PreferredMaxBytes: 2 * 1024 * 1024,
				},
				Kafka: &generate.Kafka{
					Brokers: []string{"kafka1:9090", "kafka2:9091", "kafka3:9092", "kafka4:9093"},
				},
				MaxChannels: 1000,
			},
			Orgs: []*generate.OrgInBlock{
				{Domain: orderDomain, Name: orderName, Type: generate.OrgType_Order},
				{Domain: org1Domain, Name: org1Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org1Name, org1Domain}, "."), Port: 7051},
				}},
				{Domain: org2Domain, Name: org2Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org2Name, org2Domain}, "."), Port: 7061},
				}},
			},
		},
	}
	genesis.Init()
	if err := genesis.CreateChannelCreateTx("default", channelID); nil != err {
		t.Error(err)
	} else {
		t.Log("create genesis block success")
	}
}

func TestGenerateConfig_InspectChannelTx(t *testing.T) {
	data, err := ioutil.ReadFile(geneses.ChannelTXFilePath(leagueDomain, channelID))
	//data, err := ioutil.ReadFile("/Users/aberic/Documents/path/go/src/github.com/aberic/fabric-client/geneses/example/test/channel-artifacts/channel01.tx")
	if nil != err {
		t.Error(err)
	}
	str, err := resource.InspectChannelCreateTx(data)
	if nil != err {
		t.Error(err)
	}
	t.Log(str)
}

func TestGenerateConfig_CreateChannel(t *testing.T) {
	conf := configGenerateConfig(org2Name, org2Domain, node1, admin)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	result, err := Create(orderName, admin, "grpcs://10.0.61.23:7050", org2Name, admin, channelID,
		filepath.Join(geneses.ChannelArtifactsPath(leagueDomain), strings.Join([]string{channelID, "tx"}, ".")), // "/Users/aberic/Documents/path/go/src/github.com/aberic/fabric-client/example/config/channel-artifacts/cc6519b67c4177fc1110.tx",
		confData)
	t.Log("test query result", result)
}

func TestGenerateConfig_Join(t *testing.T) {
	//genesisJoinChannel(org1Name, org1Domain, node1, admin, t)
	//genesisJoinChannel(org2Name, org2Domain, node1, admin, t)
	genesisJoinChannel(org3Name, org3Domain, node1, admin, t)
}

func TestGenerateConfig_Channels(t *testing.T) {
	//channels(org1Name, org1Domain, node1, admin, t)
	//channels(org2Name, org2Domain, node1, admin, t)
	channels(org3Name, org3Domain, node1, admin, t)
}

func TestGenerateConfig_QueryConfigBlock(t *testing.T) {
	conf := configGenerateConfig(org3Name, org3Domain, node1, admin)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	result := QueryConfigBlock(channelID, org3Name, admin, geneses.NodeDomain(org3Name, org3Domain, node1), confData)
	block := result.Data.(*common.Block)
	data, err := proto.Marshal(block)
	str, err := resource.InspectBlock(data)
	if nil != err {
		t.Error(err)
	}
	t.Log(str)
}

func TestGenerateConfig_AddGroup(t *testing.T) {
	conf := configGenerateConfig(org1Name, org1Domain, node1, admin)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	result := QueryConfigBlock(channelID, org1Name, admin, geneses.NodeDomain(org1Name, org1Domain, node1), confData)
	block := result.Data.(*common.Block)
	//data, err := proto.Marshal(block)
	//_, err = resource.InspectBlock(data)
	//if nil != err {
	//	t.Error(err)
	//}
	//t.Log(str)

	blockData, err := obtainGenesisBlock()
	updateBlock := &common.Block{}
	if err = proto.Unmarshal(blockData, updateBlock); nil != err {
		t.Error(err)
	}
	configUpdateEnvBytes, err := AddGroup(block, updateBlock, channelID, "default", "org3")
	if nil != err {
		t.Error(err)
	}
	//t.Log(string(configUpdateBytes))
	//data, err = proto.Marshal(block)
	//if nil != err {
	//	t.Error(err)
	//}
	//_, err = resource.InspectBlock(data)
	//if nil != err {
	//	t.Error(err)
	//}
	//t.Log(str)
	//resource.CreateConfigEnvelope()

	//resmgmtClient, sdk, err := resMgmtOrgClient(org1Name, admin, confData)
	//if nil != err {
	//	t.Error(err)
	//}
	err = CreateConfigUpdateBytes(configUpdateEnvBytes, leagueDomain, channelID)
	if nil != err {
		t.Error(err)
	}
	//err = UpdateChannel(sdk, resmgmtClient, block, configUpdateBytes, leagueDomain, "grpcs://10.0.61.23:7050", channelID)
	//if nil != err {
	//	t.Error(err)
	//}
}

func TestGenerateConfig_UpdateChannel(t *testing.T) {
	conf := configGenerateConfig(org1Name, org1Domain, node1, admin)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	resmgmtClient, _, err := resMgmtOrgClient(org1Name, admin, confData)
	if nil != err {
		t.Error(err)
	}
	channelUpdateFilePath := geneses.ChannelUpdateTXFilePath(leagueDomain, channelID)
	envelopeBytes, err := ioutil.ReadFile(channelUpdateFilePath)
	if nil != err {
		t.Error(err)
	}
	err = UpdateChannel(resmgmtClient, envelopeBytes, "grpcs://10.0.61.23:7050", channelID)
	if nil != err {
		t.Error(err)
	}
}

func TestGenerateConfig_Sign(t *testing.T) {
	blockSign(leagueDomain, org1Name, org1Domain, node1, admin, channelID, t)
}

func blockSign(leagueDomain, orgName, orgDomain, nodeName, orgUser, channelID string, t *testing.T) {
	conf := configGenerateConfig(orgName, orgDomain, nodeName, orgUser)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	channelUpdateFilePath := geneses.ChannelUpdateTXFilePath(leagueDomain, channelID)
	envelopeBytes, err := ioutil.ReadFile(channelUpdateFilePath)
	if nil != err {
		t.Error(err)
	}
	err = Sign(confData, envelopeBytes, leagueDomain, orgName, orgUser, channelID)
}

func obtainGenesisBlock() ([]byte, error) {
	genesis := geneses.Genesis{
		Info: &generate.ReqGenesis{
			League: &generate.LeagueInBlock{
				Domain: leagueDomain,
				Addresses: []string{
					strings.Join([]string{order0NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
					strings.Join([]string{order1NodeName, ".", orderName, ".", orderDomain, ":7050"}, ""),
				},
				BatchTimeout: 2,
				BatchSize: &generate.BatchSize{
					MaxMessageCount:   1000,
					AbsoluteMaxBytes:  10 * 1024 * 1024,
					PreferredMaxBytes: 2 * 1024 * 1024,
				},
				Kafka: &generate.Kafka{
					Brokers: []string{"kafka1:9090", "kafka2:9091", "kafka3:9092", "kafka4:9093"},
				},
				MaxChannels: 1000,
			},
			Orgs: []*generate.OrgInBlock{
				{Domain: orderDomain, Name: orderName, Type: generate.OrgType_Order},
				{Domain: org1Domain, Name: org1Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org1Name, org1Domain}, "."), Port: 7051},
				}},
				{Domain: org2Domain, Name: org2Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org2Name, org2Domain}, "."), Port: 7061},
				}},
				{Domain: org3Domain, Name: org3Name, Type: generate.OrgType_Peer, AnchorPeers: []*generate.AnchorPeer{
					{Host: strings.Join([]string{node1, org3Name, org3Domain}, "."), Port: 7071},
				}},
			},
		},
	}
	genesis.Init()
	return genesis.ObtainGenesisBlockData("default")
}

func genesisJoinChannel(orgName, orgDomain, nodeName, orgUserName string, t *testing.T) {
	conf := configGenerateConfig(orgName, orgDomain, nodeName, orgUserName)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	result := Join("grpcs://10.0.61.23:7050", orgName, orgUserName, channelID, geneses.NodeDomain(orgName, orgDomain, nodeName), confData)
	t.Log("test query result", result)
}

func channels(orgName, orgDomain, nodeName, orgUserName string, t *testing.T) {
	conf := configGenerateConfig(orgName, orgDomain, nodeName, orgUserName)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error(err)
	}
	result, err := Channels(orgName, orgUserName, geneses.NodeDomain(orgName, orgDomain, nodeName), confData)
	t.Log(result)
}

func genesisAddAffiliation(enrollID, enrollSecret, leagueDomain, orgDomain, orgName, orgMspID, caName, affiliationName, url string, t *testing.T) {
	conf := genesisCAConfigCustom(enrollID, enrollSecret, leagueDomain, orgDomain, orgName, orgMspID, caName, url)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error("yaml", err)
	}
	result, err := AddAffiliation(orgName, &msp.AffiliationRequest{
		Name:   affiliationName, // Name of the affiliation
		Force:  true,            // Creates parent affiliations if they do not exist
		CAName: caName,          // Name of the CA
	}, confData)
	if err != nil {
		t.Error("TestAddAffiliation", err)
	}
	t.Log("test ca TestAddAffiliation, result = ", result)
}

func genesisRegister(enrollID, enrollSecret, leagueDomain, orgDomain, orgName, orgMspID, caName, affiliationName, roles, delegateRoles, name, secret, registerCAType, url string, t *testing.T) {
	conf := genesisCAConfigCustom(enrollID, enrollSecret, leagueDomain, orgDomain, orgName, orgMspID, caName, url)
	confData, err := yaml.Marshal(&conf)
	if err != nil {
		t.Error("yaml", err)
	}
	result, err := Register(orgName, &msp.RegistrationRequest{
		Name:           name,
		Type:           registerCAType,  // (e.g. "client, orderer, peer, app, user")
		MaxEnrollments: -1,              // if omitted, this defaults to max_enrollments configured on the server
		Affiliation:    affiliationName, // The identity's affiliation e.g. org1.department1
		Attributes: []msp.Attribute{ // Optional attributes associated with this identity.
			// Attribute defines additional attributes that may be passed along during registration
			{Name: "hf.Registrar.Roles", Value: roles},                 // "client,orderer,peer,user" 注册者允许管理的角色列表
			{Name: "hf.Registrar.DelegateRoles", Value: delegateRoles}, // "client,orderer,peer,user" 注册者可以赋予被注册身份的hf.Registrar.Roles属性的角色列表
			{Name: "hf.Registrar.Attributes", Value: "*"},              // 注册者允许注册的属性列表
			{Name: "hf.GenCRL", Value: "true"},                         // 要注册的身份是否可以生成CRL
			{Name: "hf.Revoker", Value: "true"},                        // 要注册的身份是否可以回收证书
			{Name: "hf.AffiliationMgr", Value: "false"},                // 要注册的身份是否可以管理联盟
			{Name: "hf.IntermediateCA", Value: "true"},
		},
		CAName: caName,
		Secret: secret,
	}, confData)
	if err != nil {
		t.Error("Register", err)
	}
	t.Log("test ca Register, result = ", result)
}

func createCsr(priKeyFilePath, orgName, orgDomain, childName string, isNode bool, t *testing.T) error {
	priData, err := ioutil.ReadFile(priKeyFilePath)
	if nil != err {
		t.Error(err)
	}
	var commonName string
	if isNode {
		commonName = strings.Split(geneses.CertNodeCAName(orgName, orgDomain, childName), "-")[0]
	} else {
		commonName = strings.Split(geneses.CertUserCAName(orgName, orgDomain, childName), "-")[0]
	}
	gc := &geneses.GenerateConfig{}
	return gc.CreateCsr(&generate.ReqCreateCsr{
		PriKey:       priData,
		LeagueDomain: leagueDomain,
		OrgName:      orgName,
		OrgDomain:    orgDomain,
		Name: &generate.CSR{
			Country:      []string{"CN"},
			Organization: []string{orgName},
			Locality:     []string{childName},
			Province:     []string{orgName},
			CommonName:   commonName,
		},
		SignAlgorithm: signAlgorithm,
	})
}

func createOrgNode(orgType generate.OrgType, pubTlsKeyFilePath, orgName, orgDomain, nodeName, caURL string, t *testing.T) error {
	commonName := strings.Split(geneses.CertNodeCAName(orgName, orgDomain, nodeName), "-")[0]
	csrPem, err := ioutil.ReadFile(geneses.CsrFilePath(leagueDomain, orgName, orgDomain, commonName))
	if nil != err {
		t.Error(err)
	}
	pubTlsData, err := ioutil.ReadFile(pubTlsKeyFilePath)
	if nil != err {
		t.Error(err)
	}
	gc := &geneses.GenerateConfig{}
	return gc.CreateOrgNode(&generate.ReqCreateOrgNode{
		OrgType: orgType,
		OrgChild: &generate.OrgChild{
			LeagueDomain: leagueDomain,
			OrgName:      orgName,
			OrgDomain:    orgDomain,
			Name:         nodeName,
			PubTlsData:   pubTlsData,
			EnrollInfo: &generate.EnrollInfo{
				CsrPem:            csrPem,
				FabricCaServerURL: caURL,
				EnrollRequest: &generate.EnrollRequest{
					EnrollID: commonName,
					Secret:   "adminpw",
					//Profile:  "tls",
					Hosts: []string{commonName}, // []string{"hello.cn"}
					Name: &generate.CSR{
						Country:      []string{"CN"},
						Organization: []string{orgName},
						Locality:     []string{nodeName},
						Province:     []string{orgName},
						CommonName:   commonName,
					},
				},
				NotAfter:  50000,
				NotBefore: 0,
			},
			SignAlgorithm: signAlgorithm,
		},
	})
}

func createOrgUser(orgType generate.OrgType, pubTlsKeyFilePath, orgName, orgDomain, nodeName, caURL string, isAdmin bool, t *testing.T) error {
	commonName := strings.Split(geneses.CertUserCAName(orgName, orgDomain, nodeName), "-")[0]
	csrPem, err := ioutil.ReadFile(geneses.CsrFilePath(leagueDomain, orgName, orgDomain, commonName))
	if nil != err {
		t.Error(err)
	}
	pubTlsData, err := ioutil.ReadFile(pubTlsKeyFilePath)
	if nil != err {
		t.Error(err)
	}
	gc := &geneses.GenerateConfig{}
	return gc.CreateOrgUser(&generate.ReqCreateOrgUser{
		IsAdmin: isAdmin,
		OrgType: orgType,
		OrgChild: &generate.OrgChild{
			LeagueDomain: leagueDomain,
			OrgName:      orgName,
			OrgDomain:    orgDomain,
			Name:         nodeName,
			PubTlsData:   pubTlsData,
			EnrollInfo: &generate.EnrollInfo{
				CsrPem:            csrPem,
				FabricCaServerURL: caURL,
				EnrollRequest: &generate.EnrollRequest{
					EnrollID: commonName,
					Secret:   "adminpw",
					//Profile:  "tls",
					Hosts: []string{commonName}, // []string{"hello.cn"}
					Name: &generate.CSR{
						Country:      []string{"CN"},
						Organization: []string{orgName},
						Locality:     []string{nodeName},
						Province:     []string{orgName},
						CommonName:   commonName,
					},
				},
				NotAfter:  50000,
				NotBefore: 0,
			},
			SignAlgorithm: signAlgorithm,
		},
	})
}

func genesisCAConfigCustom(enrollID, enrollSecret, leagueDomain, orgDomain, orgName, orgMspID, caName, url string) *config.Config {
	conf := config.Config{}
	orgPath := geneses.CryptoUserTmpPath(leagueDomain, orgDomain, orgName)
	conf.InitCustomClient(false, orgName, "debug", "", "", "",
		nil, nil, nil, nil,
		&config.ClientCredentialStore{
			Path:        orgPath,
			CryptoStore: &config.ClientCredentialStoreCryptoStore{Path: orgPath},
		},
		&config.ClientBCCSP{
			Security: &config.ClientBCCSPSecurity{
				Enabled:       true,
				HashAlgorithm: "SHA2",
				SoftVerify:    true,
				Level:         256,
				Default:       &config.ClientBCCSPSecurityDefault{Provider: "SW"},
			},
		})
	conf.AddOrSetOrgForOrganizations(orgName, orgMspID, "/Users/aberic/Documents/code/ca/demo/msp", map[string]string{}, []string{}, []string{caName})
	conf.AddOrSetCertificateAuthority(caName, url,
		"/Users/aberic/Documents/path/go/src/github.com/aberic/gnomon/tmp/example/ca/pemp384/rootCA.crt",
		"/Users/aberic/Documents/path/go/src/github.com/aberic/gnomon/tmp/example/ca/pemp256/rootCA.key",
		"/Users/aberic/Documents/path/go/src/github.com/aberic/gnomon/tmp/example/ca/pemp256/rootCA.crt",
		caName, enrollID, enrollSecret)
	return &conf
}

func configGenerateConfig(orgName, orgDomain, nodeName, orgUserName string) *config.Config {
	rootPath := geneses.CryptoConfigPath(leagueDomain)
	//rootPath := "/Users/admin/Documents/code/git/go/src/github.com/aberic/fabric-client/example"
	conf := config.Config{}
	_, orgUserPath := geneses.CryptoOrgAndNodePath(leagueDomain, orgDomain, orgName, orgUserName, true, geneses.CcnAdmin)
	conf.InitCustomClient(true, orgName, "debug", rootPath, // rootPath+"/config/crypto-config"
		filepath.Join(orgUserPath, "tls", "client.key"), // rootPath+"/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/tls/client.key",
		filepath.Join(orgUserPath, "tls", "client.crt"), // rootPath+"/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/tls/client.crt")
		&config.ClientPeer{
			Timeout: &config.ClientPeerTimeout{Connection: "10s", Response: "180s",
				Discovery: &config.ClientPeerTimeoutDiscovery{GreyListExpiry: "10s"},
			},
		},
		&config.ClientEventService{Timeout: &config.ClientEventServiceTimeout{RegistrationResponse: "15s"}},
		&config.ClientOrder{Timeout: &config.ClientOrderTimeout{Connection: "15s", Response: "15s"}},
		&config.ClientGlobal{
			Timeout: &config.ClientGlobalTimeout{
				Query:   "180s",
				Execute: "180s",
				Resmgmt: "180s",
			},
			Cache: &config.ClientGlobalCache{
				ConnectionIdle:    "30s",
				EventServiceIdle:  "2m",
				ChannelConfig:     "30m",
				ChannelMembership: "30s",
				Discovery:         "10s",
				Selection:         "10m",
			},
		},
		&config.ClientCredentialStore{
			Path:        path.Join(orgUserPath, "msp", "signcerts"),
			CryptoStore: &config.ClientCredentialStoreCryptoStore{Path: path.Join(orgUserPath, "msp")},
		},
		&config.ClientBCCSP{
			Security: &config.ClientBCCSPSecurity{
				Enabled:       true,
				HashAlgorithm: "SHA2",
				SoftVerify:    true,
				Level:         256,
				Default:       &config.ClientBCCSPSecurityDefault{Provider: "SW"},
			},
		},
	)
	//conf.AddOrSetPeerForChannel("cc6519b67c4177fc11", "peer0",
	//	true, true, true, true)
	conf.AddOrSetPeerForChannel(channelID, geneses.NodeDomain(orgName, orgDomain, nodeName),
		true, true, true, true)
	conf.AddOrSetQueryChannelPolicyForChannel(channelID, "500ms", "5s",
		1, 1, 5, 2.0)
	conf.AddOrSetDiscoveryPolicyForChannel(channelID, "500ms", "5s",
		2, 4, 2.0)
	conf.AddOrSetEventServicePolicyForChannel(channelID, "PreferOrg", "RoundRobin",
		"6s", 5, 8)
	_, orderUserPath := geneses.CryptoOrgAndNodePath(leagueDomain, orderDomain, orderName, admin, false, geneses.CcnAdmin)
	conf.AddOrSetOrdererForOrganizations(orderName, strings.Join([]string{orderName, "MSP"}, ""),
		path.Join(orderUserPath, "msp"), // rootPath+"/config/crypto-config/ordererOrganizations/20de78630ef6a411/users/Admin@20de78630ef6a411/msp",
		map[string]string{
			//"Admin": rootPath + "/config/crypto-config/ordererOrganizations/20de78630ef6a411/users/Admin@20de78630ef6a411/msp/signcerts/Admin@20de78630ef6a411-cert.pem",
			admin: filepath.Join(orderUserPath, "msp", "signcerts", geneses.CertUserCAName(orderName, orderDomain, admin)),
		},
	)

	_, org1UserPath := geneses.CryptoOrgAndNodePath(leagueDomain, org1Domain, org1Name, admin, true, geneses.CcnAdmin)
	_, org2UserPath := geneses.CryptoOrgAndNodePath(leagueDomain, org2Domain, org2Name, admin, true, geneses.CcnAdmin)
	_, org3UserPath := geneses.CryptoOrgAndNodePath(leagueDomain, org3Domain, org3Name, admin, true, geneses.CcnAdmin)
	conf.AddOrSetOrgForOrganizations(org1Name, strings.Join([]string{org1Name, "MSP"}, ""),
		path.Join(org1UserPath, "msp"), // rootPath+"/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp",
		map[string]string{
			//"Admin": rootPath + "/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp/signcerts/Admin@20de78630ef6a411-org1-cert.pem",
			admin: filepath.Join(org1UserPath, "msp", "signcerts", geneses.CertUserCAName(org1Name, org1Domain, admin)),
		},
		[]string{geneses.NodeDomain(org1Name, org1Domain, node1), geneses.NodeDomain(org1Name, org1Domain, node2)}, //"peer0", "peer1",
		[]string{},
	)
	conf.AddOrSetOrgForOrganizations(org2Name, strings.Join([]string{org2Name, "MSP"}, ""),
		path.Join(org2UserPath, "msp"), // rootPath+"/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp",
		map[string]string{
			//"Admin": rootPath + "/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp/signcerts/Admin@20de78630ef6a411-org1-cert.pem",
			admin: filepath.Join(org2UserPath, "msp", "signcerts", geneses.CertUserCAName(org2Name, org2Domain, admin)),
		},
		[]string{geneses.NodeDomain(org2Name, org2Domain, node1), geneses.NodeDomain(org2Name, org2Domain, node2)}, //"peer0", "peer1",
		[]string{},
	)
	conf.AddOrSetOrgForOrganizations(org3Name, strings.Join([]string{org3Name, "MSP"}, ""),
		path.Join(org3UserPath, "msp"), // rootPath+"/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp",
		map[string]string{
			//"Admin": rootPath + "/config/crypto-config/peerOrganizations/20de78630ef6a411-org1/users/Admin@20de78630ef6a411-org1/msp/signcerts/Admin@20de78630ef6a411-org1-cert.pem",
			admin: filepath.Join(org3UserPath, "msp", "signcerts", geneses.CertUserCAName(org3Name, org3Domain, admin)),
		},
		[]string{geneses.NodeDomain(org3Name, org3Domain, node1), geneses.NodeDomain(org3Name, org3Domain, node2)}, //"peer0", "peer1",
		[]string{},
	)
	tlsRootCaFilePath := filepath.Join(geneses.CryptoRootTLSCAPath(leagueDomain), geneses.CertRootTLSCAName(leagueDomain))
	conf.AddOrSetOrderer(orderName, "grpcs://10.0.61.23:7050",
		strings.Join([]string{order0NodeName, orderName, orderDomain}, "."), "0s", "20s",
		tlsRootCaFilePath, // rootPath+"/config/crypto-config/ordererOrganizations/20de78630ef6a411/tlsca/tlsca.20de78630ef6a411-cert.pem",
		false, false, false)
	conf.AddOrSetPeer(geneses.NodeDomain(org1Name, org1Domain, node1), "grpcs://10.0.61.23:7051",
		"grpcs://10.0.61.23:7052", strings.Join([]string{node1, org1Name, org1Domain}, "."),
		"0s", "20s",
		tlsRootCaFilePath,
		false, false, false)
	conf.AddOrSetPeer(geneses.NodeDomain(org2Name, org2Domain, node1), "grpcs://10.0.61.23:7061",
		"grpcs://10.0.61.23:7062", strings.Join([]string{node1, org2Name, org2Domain}, "."),
		"0s", "20s",
		tlsRootCaFilePath,
		false, false, false)
	conf.AddOrSetPeer(geneses.NodeDomain(org3Name, org3Domain, node1), "grpcs://10.0.61.23:7071",
		"grpcs://10.0.61.23:7072", strings.Join([]string{node1, org3Name, org3Domain}, "."),
		"0s", "20s",
		tlsRootCaFilePath,
		false, false, false)
	return &conf
}

func TestGenerateConfig_eccSKI(t *testing.T) {
	//priKey, err := gnomon.CryptoECC().LoadPriPemFP("/tmp/366428000/pri.key")
	priKey, err := gnomon.CryptoECC().LoadPriPemFP("/tmp/276880000/pri.key")
	if nil != err {
		t.Error(err)
	}
	gc := &geneses.GenerateConfig{}
	data := gc.EccSKI(priKey)
	t.Log(data)
	t.Log(hex.EncodeToString(data))
}
