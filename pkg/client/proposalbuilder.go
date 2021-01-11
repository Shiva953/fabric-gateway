/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	gateway "github.com/hyperledger/fabric-gateway/protos"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/protoutil"
)

type proposalBuilder struct {
	client          gateway.GatewayClient
	signingID       *signingIdentity
	channelName     string
	chaincodeID     string
	transactionName string
	transient       map[string][]byte
	args            [][]byte
}

func (builder *proposalBuilder) build() (*Proposal, error) {
	proposalProto, transactionID, err := builder.newProposalProto()
	if err != nil {
		return nil, fmt.Errorf("Failed to create Proposal protobuf: %w", err)
	}

	proposalBytes, err := proto.Marshal(proposalProto)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshall Proposal protobuf: %w", err)
	}

	signedProposalProto := &peer.SignedProposal{
		ProposalBytes: proposalBytes,
	}

	proposedTransactionProto := &gateway.ProposedTransaction{
		Proposal:  signedProposalProto,
		TxId:      transactionID,
		ChannelId: builder.channelName,
	}

	proposal := &Proposal{
		client:              builder.client,
		signingID:           builder.signingID,
		proposedTransaction: proposedTransactionProto,
	}
	return proposal, nil
}

func (builder *proposalBuilder) newProposalProto() (*peer.Proposal, string, error) {
	invocationSpec := &peer.ChaincodeInvocationSpec{
		ChaincodeSpec: &peer.ChaincodeSpec{
			Type:        peer.ChaincodeSpec_NODE,
			ChaincodeId: &peer.ChaincodeID{Name: builder.chaincodeID},
			Input:       &peer.ChaincodeInput{Args: builder.chaincodeArgs()},
		},
	}

	creator, err := builder.signingID.Creator()
	if err != nil {
		return nil, "", fmt.Errorf("Failed to serialize identity: %w", err)
	}

	result, transactionID, err := protoutil.CreateChaincodeProposalWithTransient(
		common.HeaderType_ENDORSER_TRANSACTION,
		builder.channelName,
		invocationSpec,
		creator,
		builder.transient,
	)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to create chaincode proposal: %w", err)
	}

	return result, transactionID, nil
}

func (builder *proposalBuilder) chaincodeArgs() [][]byte {
	result := make([][]byte, len(builder.args)+1)

	result[0] = []byte(builder.transactionName)
	copy(result[1:], builder.args)

	return result
}

// ProposalOption implements an option for a transaction proposal.
type ProposalOption = func(builder *proposalBuilder) error

// WithArguments appends to the transaction function arguments associated with a transaction proposal.
func WithArguments(args ...[]byte) ProposalOption {
	return func(builder *proposalBuilder) error {
		builder.args = append(builder.args, args...)
		return nil
	}
}

// WithStringArguments appends to the transaction function arguments associated with a transaction proposal.
func WithStringArguments(args ...string) ProposalOption {
	return WithArguments(stringsAsBytes(args)...)
}

func stringsAsBytes(strings []string) [][]byte {
	results := make([][]byte, 0, len(strings))

	for _, v := range strings {
		results = append(results, []byte(v))
	}

	return results
}

// WithTransient specifies the transient data associated with a transaction proposal.
func WithTransient(transient map[string][]byte) ProposalOption {
	return func(builder *proposalBuilder) error {
		builder.transient = transient
		return nil
	}
}