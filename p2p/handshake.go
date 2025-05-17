package p2p


type HanshakeFunc func(any) error 
func NOHandshakeFunc (any) error { return nil}
