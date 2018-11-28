// this file was generated by protoc-gen-gotemplate
package jsonclient

import (
	"context"
	"encoding/json"

	"berty.tech/core/api/client"
	"berty.tech/core/api/node"
	"berty.tech/core/api/p2p"
	"berty.tech/core/api/protobuf/graphql"
	"berty.tech/core/entity"
	"berty.tech/core/network"
	"go.uber.org/zap"
)

func init() {
	registerUnary("berty.node.ID", NodeID)
	registerServerStream("berty.node.EventStream", NodeEventStream)
	registerServerStream("berty.node.EventList", NodeEventList)
	registerUnary("berty.node.GetEvent", NodeGetEvent)
	registerUnary("berty.node.EventSeen", NodeEventSeen)
	registerUnary("berty.node.ContactRequest", NodeContactRequest)
	registerUnary("berty.node.ContactAcceptRequest", NodeContactAcceptRequest)
	registerUnary("berty.node.ContactRemove", NodeContactRemove)
	registerUnary("berty.node.ContactUpdate", NodeContactUpdate)
	registerServerStream("berty.node.ContactList", NodeContactList)
	registerUnary("berty.node.GetContact", NodeGetContact)
	registerUnary("berty.node.ConversationCreate", NodeConversationCreate)
	registerServerStream("berty.node.ConversationList", NodeConversationList)
	registerUnary("berty.node.ConversationInvite", NodeConversationInvite)
	registerUnary("berty.node.ConversationExclude", NodeConversationExclude)
	registerUnary("berty.node.ConversationAddMessage", NodeConversationAddMessage)
	registerUnary("berty.node.GetConversation", NodeGetConversation)
	registerUnary("berty.node.GetConversationMember", NodeGetConversationMember)
	registerUnary("berty.node.ConversationRead", NodeConversationRead)
	registerUnary("berty.node.HandleEvent", NodeHandleEvent)
	registerUnary("berty.node.GenerateFakeData", NodeGenerateFakeData)
	registerUnary("berty.node.RunIntegrationTests", NodeRunIntegrationTests)
	registerUnary("berty.node.DebugPing", NodeDebugPing)
	registerUnary("berty.node.DebugRequeueEvent", NodeDebugRequeueEvent)
	registerUnary("berty.node.DebugRequeueAll", NodeDebugRequeueAll)
	registerUnary("berty.node.DeviceInfos", NodeDeviceInfos)
	registerUnary("berty.node.AppVersion", NodeAppVersion)
	registerUnary("berty.node.Peers", NodePeers)
	registerUnary("berty.node.Protocols", NodeProtocols)
	registerServerStream("berty.node.LogStream", NodeLogStream)
	registerServerStream("berty.node.LogfileList", NodeLogfileList)
	registerServerStream("berty.node.LogfileRead", NodeLogfileRead)
	registerServerStream("berty.node.MonitorBandwidth", NodeMonitorBandwidth)
	registerServerStream("berty.node.MonitorPeers", NodeMonitorPeers)
	registerUnary("berty.node.Panic", NodePanic)
}
func NodeID(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ID"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ID(ctx, &typedInput)
}
func NodeEventStream(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "EventStream"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.EventStreamInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().EventStream(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeEventList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "EventList"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.EventListInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().EventList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeGetEvent(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "GetEvent"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput p2p.Event
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().GetEvent(ctx, &typedInput)
}
func NodeEventSeen(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "EventSeen"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput graphql.Node
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().EventSeen(ctx, &typedInput)
}
func NodeContactRequest(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactRequest"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ContactRequestInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactRequest(ctx, &typedInput)
}
func NodeContactAcceptRequest(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactAcceptRequest"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactAcceptRequest(ctx, &typedInput)
}
func NodeContactRemove(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactRemove"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactRemove(ctx, &typedInput)
}
func NodeContactUpdate(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactUpdate"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ContactUpdate(ctx, &typedInput)
}
func NodeContactList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ContactList"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ContactListInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().ContactList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeGetContact(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "GetContact"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.Contact
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().GetContact(ctx, &typedInput)
}
func NodeConversationCreate(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationCreate"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ConversationCreateInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationCreate(ctx, &typedInput)
}
func NodeConversationList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationList"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ConversationListInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().ConversationList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeConversationInvite(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationInvite"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ConversationManageMembersInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationInvite(ctx, &typedInput)
}
func NodeConversationExclude(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationExclude"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ConversationManageMembersInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationExclude(ctx, &typedInput)
}
func NodeConversationAddMessage(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationAddMessage"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.ConversationAddMessageInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationAddMessage(ctx, &typedInput)
}
func NodeGetConversation(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "GetConversation"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.Conversation
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().GetConversation(ctx, &typedInput)
}
func NodeGetConversationMember(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "GetConversationMember"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput entity.ConversationMember
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().GetConversationMember(ctx, &typedInput)
}
func NodeConversationRead(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "ConversationRead"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput graphql.Node
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().ConversationRead(ctx, &typedInput)
}
func NodeHandleEvent(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "HandleEvent"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput p2p.Event
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().HandleEvent(ctx, &typedInput)
}
func NodeGenerateFakeData(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "GenerateFakeData"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().GenerateFakeData(ctx, &typedInput)
}
func NodeRunIntegrationTests(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "RunIntegrationTests"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.IntegrationTestInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().RunIntegrationTests(ctx, &typedInput)
}
func NodeDebugPing(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "DebugPing"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.PingDestination
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().DebugPing(ctx, &typedInput)
}
func NodeDebugRequeueEvent(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "DebugRequeueEvent"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.EventIDInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().DebugRequeueEvent(ctx, &typedInput)
}
func NodeDebugRequeueAll(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "DebugRequeueAll"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().DebugRequeueAll(ctx, &typedInput)
}
func NodeDeviceInfos(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "DeviceInfos"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().DeviceInfos(ctx, &typedInput)
}
func NodeAppVersion(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "AppVersion"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().AppVersion(ctx, &typedInput)
}
func NodePeers(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "Peers"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().Peers(ctx, &typedInput)
}
func NodeProtocols(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "Protocols"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput network.Peer
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().Protocols(ctx, &typedInput)
}
func NodeLogStream(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "LogStream"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.LogStreamInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().LogStream(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeLogfileList(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "LogfileList"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().LogfileList(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeLogfileRead(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "LogfileRead"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.LogfileReadInput
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().LogfileRead(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeMonitorBandwidth(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "MonitorBandwidth"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput network.BandwidthStats
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().MonitorBandwidth(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodeMonitorPeers(client *client.Client, ctx context.Context, jsonInput []byte) (GenericServerStreamClient, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "MonitorPeers"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	stream, err := client.Node().MonitorPeers(ctx, &typedInput)
	if err != nil {
		return nil, err
	}
	// start a stream proxy
	streamProxy := newGenericServerStreamProxy()
	go func() {
		for {
			data, err := stream.Recv()
			streamProxy.queue <- genericStreamEntry{data: data, err: err}
			if err != nil {
				break
			}
		}
		// FIXME: wait for queue to be empty, then close chan
	}()
	return streamProxy, nil
}
func NodePanic(client *client.Client, ctx context.Context, jsonInput []byte) (interface{}, error) {
	logger().Debug("client call",
		zap.String("service", "Service"),
		zap.String("method", "Panic"),
		zap.String("input", string(jsonInput)),
	)
	var typedInput node.Void
	if err := json.Unmarshal(jsonInput, &typedInput); err != nil {
		return nil, err
	}
	return client.Node().Panic(ctx, &typedInput)
}
