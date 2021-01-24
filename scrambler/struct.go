package scrambler

type Scrambler struct {
	opts ConstructorOpts
	start          uint32
	cur            uint32
	highestBitMask uint32
	feedbackMask   uint32
	outOffset      uint32
	atStart        bool
	atEnd          bool
	n              int
}

type ConstructorOpts struct {
	Maximum uint32
	Seed int
	OutOffset int
	Loop bool
}
