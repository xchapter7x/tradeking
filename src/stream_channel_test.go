package tradeking

import (
    "testing"
)

func Test_IsAliveWhenAlive(t *testing.T) {
    stream := NewStreamChannel()
    defer stream.Close()
    aliveValue := stream.IsAlive()
    aliveControl := true

    if aliveControl != aliveValue {
        t.Errorf("%s != %s; values should be the same", aliveControl, aliveValue)
    }
}
func Test_IsAliveReturnsFalseOnClose(t *testing.T) {
    stream := NewStreamChannel()
    stream.Close()
    aliveValue := stream.IsAlive()
    aliveControl := false

    if aliveControl != aliveValue {
        t.Errorf("%s != %s; values should be the same", aliveControl, aliveValue)
    }
}

func Test_WriteStringToBuffer(t *testing.T) {
    stream := NewStreamChannel()
    defer stream.Close()
    controlString := "testString"
    go func() { stream.Write(controlString) }()
    streamValue := <-stream.Channel

    if controlString != streamValue {
        t.Errorf("%s != %s; values should be the same", controlString, streamValue)
    }
}


func Test_CloseClosesChannel(t *testing.T) {
    stream := NewStreamChannel()
    stream.Close()
    _, channelOpen := <-stream.Channel

    if channelOpen {
        t.Errorf("the channel should have been closed")
    }
}
