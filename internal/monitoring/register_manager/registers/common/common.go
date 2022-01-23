package common

type Register interface {
	//Just increases value of register
	Inc()

	//Just decreases value of register
	Dec()
}
