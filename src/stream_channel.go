package tradeking

func NewStreamChannel() (sc *StreamChannel) {
    channel := make(chan string)
    alive := true
    sc = &StreamChannel{
        Channel: channel,
        alive: alive}
    return
}

type StreamChannel struct {
    Channel chan string
    alive bool
}

func (s *StreamChannel) IsAlive() (bool) {
    return s.alive
}

func (s *StreamChannel) Write(buf string) {
    s.Channel <- buf
}

func (s *StreamChannel) Close() {
    close(s.Channel)
    s.alive = false
}
