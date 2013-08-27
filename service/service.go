package service

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
