package broker

import (
	"sync"
	"testing"
)

func TestNewSupervisor(t *testing.T) {
	supervisor := NewSupervisor()
	if supervisor.clientsMapping == nil {
		t.Fatalf("error, expected clientsMapping to be a nil pointer, instead got %v", supervisor.clientsMapping)
	}
	if supervisor.clientsMutex == nil {
		t.Fatalf("error, expected clientsMutex to be a nil pointer, instead got %v", supervisor.clientsMutex)
	}
	if supervisor.IsRunning != false {
		t.Fatalf("error, expected IsRunning to be a false, instead got %v", supervisor.IsRunning)
	}
	if supervisor.clientOperationChannel == nil {
		t.Fatalf("error, expected clientOperationChannel to be a nil pointer, instead got %v", supervisor.clientOperationChannel)
	}
	if supervisor.stoppedSignal == nil {
		t.Fatalf("error, expected stoppedSignal to be a nil pointer, instead got %v", supervisor.stoppedSignal)
	}
	if supervisor.startedSignal == nil {
		t.Fatalf("error, expected startedSignal to be a nil pointer, instead got %v", supervisor.startedSignal)
	}
	if supervisor.mainWg == nil {
		t.Fatalf("error, expected mainWg to be a nil pointer, instead got %v", supervisor.mainWg)
	}
	if supervisor.clientCount != 0 {
		t.Fatalf("error, expected clientCount to be a zero, instead got %v", supervisor.clientCount)
	}
}

func TestRun(t *testing.T) {
	var wg sync.WaitGroup
	supervisor := NewSupervisor()
	wg.Add(1)
	go func() {
		supervisor.START()
		wg.Done()
	}()
	supervisor.WaitUntilReady()
	if supervisor.IsRunning != true {
		t.Fatalf("error, expected IsRunnint to be True, instead got %v", supervisor.IsRunning)
	}
	supervisor.STOP()
	wg.Wait()
	if supervisor.IsRunning != false {
		t.Fatalf("error, expected IsRunnint to be False, instead got %v", supervisor.IsRunning)
	}
}

func TestAddClient(t *testing.T) {
	var wg sync.WaitGroup
	supervisor := NewSupervisor()
	wg.Add(1)
	go func() {
		supervisor.START()
		wg.Done()
	}()
	supervisor.WaitUntilReady()
	client := Client{
		Id: "1",
	}
	supervisor.SignalClientAddition(&client)
	supervisor.STOP()
	wg.Wait()
	clientCount := supervisor.GetClientCount()
	if clientCount != 1 {
		t.Fatalf("error, expected GetClientCount to return 1, instead got %v", clientCount)
	}
	result := supervisor.clientsMapping[client.Id]
	if result == nil {
		t.Fatalf("error, expected result not to be a poiner")
	}
	if *result != client {
		t.Fatal("error, expected result to point to the same location address ass the defined client")
	}
}

func TestRemoveClient(t *testing.T) {
	var wg sync.WaitGroup
	supervisor := NewSupervisor()
	wg.Add(1)
	go func() {
		supervisor.START()
		wg.Done()
	}()
	supervisor.WaitUntilReady()
	client := Client{
		Id: "1",
	}
	_ = supervisor.addClient(&client)
	clientCount := supervisor.GetClientCount()
	if clientCount != 1 {
		t.Fatalf("error, expected GetClientCount to return 1, instead got %v", clientCount)
	}
	supervisor.SignalClientRemoval(&client)
	supervisor.STOP()
	wg.Wait()
	clientCount = supervisor.GetClientCount()
	if clientCount != 0 {
		t.Fatalf("error, expected GetClientCount to return 0, instead got %v", clientCount)
	}
	result := supervisor.clientsMapping[client.Id]
	if result != nil {
		t.Fatalf("error, expected result to be a nil pointer, instead got %v", result)
	}
}
