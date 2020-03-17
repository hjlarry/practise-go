package main

func add(x, y int) int {
	z := x + y
	return z
}

func main() {
	a, b := 0x11, 0x22
	s := add(a, b)
	println(s)
}

/*
由调用方分配参数和返回值内存
go build -gcflags "-l -N" call.go

go tool objdump -s "main\.main" call

TEXT main.main(SB) /root/.mac/gocode/call.go
  call.go:8		0x452370		64488b0c25f8ffffff	MOVQ FS:0xfffffff8, CX
  call.go:8		0x452379		483b6110		CMPQ 0x10(CX), SP
  call.go:8		0x45237d		7668			JBE 0x4523e7
  call.go:8		0x45237f		4883ec38		SUBQ $0x38, SP
  call.go:8		0x452383		48896c2430		MOVQ BP, 0x30(SP)
  call.go:8		0x452388		488d6c2430		LEAQ 0x30(SP), BP
  call.go:9		0x45238d		48c744242811000000	MOVQ $0x11, 0x28(SP)
  call.go:9		0x452396		48c744242022000000	MOVQ $0x22, 0x20(SP)
  call.go:10		0x45239f		488b442428		MOVQ 0x28(SP), AX
  call.go:10		0x4523a4		48890424		MOVQ AX, 0(SP)
  call.go:10		0x4523a8		48c744240822000000	MOVQ $0x22, 0x8(SP)
  call.go:10		0x4523b1		e87affffff		CALL main.add(SB)
  call.go:10		0x4523b6		488b442410		MOVQ 0x10(SP), AX
  call.go:10		0x4523bb		4889442418		MOVQ AX, 0x18(SP)
  call.go:11		0x4523c0		e8eb30fdff		CALL runtime.printlock(SB)
  call.go:11		0x4523c5		488b442418		MOVQ 0x18(SP), AX
  call.go:11		0x4523ca		48890424		MOVQ AX, 0(SP)
  call.go:11		0x4523ce		e85d38fdff		CALL runtime.printint(SB)
  call.go:11		0x4523d3		e86833fdff		CALL runtime.printnl(SB)
  call.go:11		0x4523d8		e85331fdff		CALL runtime.printunlock(SB)
  call.go:12		0x4523dd		488b6c2430		MOVQ 0x30(SP), BP
  call.go:12		0x4523e2		4883c438		ADDQ $0x38, SP
  call.go:12		0x4523e6		c3			RET
  call.go:8		0x4523e7		e8b47affff		CALL runtime.morestack_noctxt(SB)
  call.go:8		0x4523ec		eb82			JMP main.main(SB)


go tool objdump -s "main\.add" call

TEXT main.add(SB) /root/.mac/gocode/call.go
  call.go:3		0x452330		4883ec10		SUBQ $0x10, SP
  call.go:3		0x452334		48896c2408		MOVQ BP, 0x8(SP)
  call.go:3		0x452339		488d6c2408		LEAQ 0x8(SP), BP
  call.go:3		0x45233e		48c744242800000000	MOVQ $0x0, 0x28(SP)
  call.go:4		0x452347		488b442418		MOVQ 0x18(SP), AX
  call.go:4		0x45234c		4803442420		ADDQ 0x20(SP), AX
  call.go:4		0x452351		48890424		MOVQ AX, 0(SP)
  call.go:5		0x452355		4889442428		MOVQ AX, 0x28(SP)
  call.go:5		0x45235a		488b6c2408		MOVQ 0x8(SP), BP
  call.go:5		0x45235f		4883c410		ADDQ $0x10, SP
  call.go:5		0x452363		c3			RET
*/
