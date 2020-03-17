package main

/*
(gdb) info frame
Stack level 0, frame at 0xc00002e748:
Arglist at 0xc00002e698, args: x=100, y= []uint8 = {...}
Locals at 0xc00002e698, Previous frame's sp is 0xc00002e748 (gdb) x/3xg &x
0xc00002e748: 0x0000000000000064 0x000000c000068010 0xc00002e758: 0x0000000000000003
(gdb) x/3xb 0x000000c000068010 0xc000068010: 0x01 0x02 0x03
*/
func test(x int, y ...byte) {
	println(y)
}

func main() {
	test(100, 1, 2, 3)
}
