package broker

import (
	"fmt"
	"sync"
)

type Supervisor struct {
	IsRunning              bool
	clientCount            uint
	clientsMutex           *sync.Mutex
	clientsMapping         map[string]*Client
	clientOperationChannel chan ClientOperation
	stoppedSignal          chan struct{}
	startedSignal          chan struct{}
	mainWg                 *sync.WaitGroup
}

func NewSupervisor() *Supervisor {
	return &Supervisor{
		clientCount:            0,
		IsRunning:              false,
		clientsMutex:           &sync.Mutex{},
		clientsMapping:         make(map[string]*Client),
		clientOperationChannel: make(chan ClientOperation),
		stoppedSignal:          make(chan struct{}),
		startedSignal:          make(chan struct{}),
		mainWg:                 &sync.WaitGroup{},
	}
}

func (s *Supervisor) START() {
	s.mainWg.Add(1)
	go s.waitForClientOperations()
	s.IsRunning = true
	close(s.startedSignal)
	s.mainWg.Wait()
	s.IsRunning = false

}

func (s *Supervisor) waitForClientOperations() {
	defer s.mainWg.Done()
	for {
		select {
		case o := <-s.clientOperationChannel:
			switch o.Operation {
			case Add:
				s.addClient(o.Client)
			case Remove:
				s.removeClient(o.Client.Id)
			default:
				fmt.Println("no operation")
			}
		case <-s.stoppedSignal:
			return
		}
	}

}

func (s *Supervisor) addClient(c *Client) error {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	s.clientsMapping[c.Id] = c
	s.clientCount += 1
	return nil
}

func (s *Supervisor) removeClient(clientId string) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	delete(s.clientsMapping, clientId)
	s.clientCount -= 1
}

func (s *Supervisor) GetClientCount() uint {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	return s.clientCount

}

func (s *Supervisor) SignalClientAddition(c *Client) {
	s.clientOperationChannel <- ClientOperation{Add, c}
}

func (s *Supervisor) SignalClientRemoval(c *Client) {
	s.clientOperationChannel <- ClientOperation{Remove, c}
}

func (s *Supervisor) WaitUntilReady() {
	<-s.startedSignal
}

func (s *Supervisor) STOP() {
	close(s.stoppedSignal)
}
