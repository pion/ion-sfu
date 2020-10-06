// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sfu

import (
	"sync"
)

// Ensure, that RouterMock does implement Router.
// If this is not the case, regenerate this file with moq.
var _ Router = &RouterMock{}

// RouterMock is a mock implementation of Router.
//
//     func TestSomethingThatUsesRouter(t *testing.T) {
//
//         // make and configure a mocked Router
//         mockedRouter := &RouterMock{
//             AddReceiverFunc: func(recv Receiver)  {
// 	               panic("mock out the AddReceiver method")
//             },
//             AddSenderFunc: func(p *WebRTCTransport) error {
// 	               panic("mock out the AddSender method")
//             },
//             ConfigFunc: func() RouterConfig {
// 	               panic("mock out the Config method")
//             },
//             GetReceiverFunc: func(layer uint8) Receiver {
// 	               panic("mock out the GetReceiver method")
//             },
//             IDFunc: func() string {
// 	               panic("mock out the ID method")
//             },
//             SwitchSpatialLayerFunc: func(targetLayer uint8, sub Sender) bool {
// 	               panic("mock out the SwitchSpatialLayer method")
//             },
//         }
//
//         // use mockedRouter in code that requires Router
//         // and then make assertions.
//
//     }
type RouterMock struct {
	// AddReceiverFunc mocks the AddReceiver method.
	AddReceiverFunc func(recv Receiver)

	// AddSenderFunc mocks the AddSender method.
	AddSenderFunc func(p *WebRTCTransport) error

	// ConfigFunc mocks the Config method.
	ConfigFunc func() RouterConfig

	// GetReceiverFunc mocks the GetReceiver method.
	GetReceiverFunc func(layer uint8) Receiver

	// IDFunc mocks the ID method.
	IDFunc func() string

	// SwitchSpatialLayerFunc mocks the SwitchSpatialLayer method.
	SwitchSpatialLayerFunc func(targetLayer uint8, sub Sender) bool

	// calls tracks calls to the methods.
	calls struct {
		// AddReceiver holds details about calls to the AddReceiver method.
		AddReceiver []struct {
			// Recv is the recv argument value.
			Recv Receiver
		}
		// AddSender holds details about calls to the AddSender method.
		AddSender []struct {
			// P is the p argument value.
			P *WebRTCTransport
		}
		// Config holds details about calls to the Config method.
		Config []struct {
		}
		// GetReceiver holds details about calls to the GetReceiver method.
		GetReceiver []struct {
			// Layer is the layer argument value.
			Layer uint8
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// SwitchSpatialLayer holds details about calls to the SwitchSpatialLayer method.
		SwitchSpatialLayer []struct {
			// TargetLayer is the targetLayer argument value.
			TargetLayer uint8
			// Sub is the sub argument value.
			Sub Sender
		}
	}
	lockAddReceiver        sync.RWMutex
	lockAddSender          sync.RWMutex
	lockConfig             sync.RWMutex
	lockGetReceiver        sync.RWMutex
	lockID                 sync.RWMutex
	lockSwitchSpatialLayer sync.RWMutex
}

// AddReceiver calls AddReceiverFunc.
func (mock *RouterMock) AddReceiver(recv Receiver) {
	if mock.AddReceiverFunc == nil {
		panic("RouterMock.AddReceiverFunc: method is nil but Router.AddReceiver was just called")
	}
	callInfo := struct {
		Recv Receiver
	}{
		Recv: recv,
	}
	mock.lockAddReceiver.Lock()
	mock.calls.AddReceiver = append(mock.calls.AddReceiver, callInfo)
	mock.lockAddReceiver.Unlock()
	mock.AddReceiverFunc(recv)
}

// AddReceiverCalls gets all the calls that were made to AddReceiver.
// Check the length with:
//     len(mockedRouter.AddReceiverCalls())
func (mock *RouterMock) AddReceiverCalls() []struct {
	Recv Receiver
} {
	var calls []struct {
		Recv Receiver
	}
	mock.lockAddReceiver.RLock()
	calls = mock.calls.AddReceiver
	mock.lockAddReceiver.RUnlock()
	return calls
}

// AddSender calls AddSenderFunc.
func (mock *RouterMock) AddSender(p *WebRTCTransport) error {
	if mock.AddSenderFunc == nil {
		panic("RouterMock.AddSenderFunc: method is nil but Router.AddSender was just called")
	}
	callInfo := struct {
		P *WebRTCTransport
	}{
		P: p,
	}
	mock.lockAddSender.Lock()
	mock.calls.AddSender = append(mock.calls.AddSender, callInfo)
	mock.lockAddSender.Unlock()
	return mock.AddSenderFunc(p)
}

// AddSenderCalls gets all the calls that were made to AddSender.
// Check the length with:
//     len(mockedRouter.AddSenderCalls())
func (mock *RouterMock) AddSenderCalls() []struct {
	P *WebRTCTransport
} {
	var calls []struct {
		P *WebRTCTransport
	}
	mock.lockAddSender.RLock()
	calls = mock.calls.AddSender
	mock.lockAddSender.RUnlock()
	return calls
}

// Config calls ConfigFunc.
func (mock *RouterMock) Config() RouterConfig {
	if mock.ConfigFunc == nil {
		panic("RouterMock.ConfigFunc: method is nil but Router.Config was just called")
	}
	callInfo := struct {
	}{}
	mock.lockConfig.Lock()
	mock.calls.Config = append(mock.calls.Config, callInfo)
	mock.lockConfig.Unlock()
	return mock.ConfigFunc()
}

// ConfigCalls gets all the calls that were made to Config.
// Check the length with:
//     len(mockedRouter.ConfigCalls())
func (mock *RouterMock) ConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockConfig.RLock()
	calls = mock.calls.Config
	mock.lockConfig.RUnlock()
	return calls
}

// GetReceiver calls GetReceiverFunc.
func (mock *RouterMock) GetReceiver(layer uint8) Receiver {
	if mock.GetReceiverFunc == nil {
		panic("RouterMock.GetReceiverFunc: method is nil but Router.GetReceiver was just called")
	}
	callInfo := struct {
		Layer uint8
	}{
		Layer: layer,
	}
	mock.lockGetReceiver.Lock()
	mock.calls.GetReceiver = append(mock.calls.GetReceiver, callInfo)
	mock.lockGetReceiver.Unlock()
	return mock.GetReceiverFunc(layer)
}

// GetReceiverCalls gets all the calls that were made to GetReceiver.
// Check the length with:
//     len(mockedRouter.GetReceiverCalls())
func (mock *RouterMock) GetReceiverCalls() []struct {
	Layer uint8
} {
	var calls []struct {
		Layer uint8
	}
	mock.lockGetReceiver.RLock()
	calls = mock.calls.GetReceiver
	mock.lockGetReceiver.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *RouterMock) ID() string {
	if mock.IDFunc == nil {
		panic("RouterMock.IDFunc: method is nil but Router.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//     len(mockedRouter.IDCalls())
func (mock *RouterMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// SwitchSpatialLayer calls SwitchSpatialLayerFunc.
func (mock *RouterMock) SwitchSpatialLayer(targetLayer uint8, sub Sender) bool {
	if mock.SwitchSpatialLayerFunc == nil {
		panic("RouterMock.SwitchSpatialLayerFunc: method is nil but Router.SwitchSpatialLayer was just called")
	}
	callInfo := struct {
		TargetLayer uint8
		Sub         Sender
	}{
		TargetLayer: targetLayer,
		Sub:         sub,
	}
	mock.lockSwitchSpatialLayer.Lock()
	mock.calls.SwitchSpatialLayer = append(mock.calls.SwitchSpatialLayer, callInfo)
	mock.lockSwitchSpatialLayer.Unlock()
	return mock.SwitchSpatialLayerFunc(targetLayer, sub)
}

// SwitchSpatialLayerCalls gets all the calls that were made to SwitchSpatialLayer.
// Check the length with:
//     len(mockedRouter.SwitchSpatialLayerCalls())
func (mock *RouterMock) SwitchSpatialLayerCalls() []struct {
	TargetLayer uint8
	Sub         Sender
} {
	var calls []struct {
		TargetLayer uint8
		Sub         Sender
	}
	mock.lockSwitchSpatialLayer.RLock()
	calls = mock.calls.SwitchSpatialLayer
	mock.lockSwitchSpatialLayer.RUnlock()
	return calls
}
