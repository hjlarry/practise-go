#include <stdio.h>

// gcc -g -O2 -o recursion recursion.c
// objdump -d -M intel recursion | less
// 观察到factorial的汇编是跳转循环指令
__attribute__((noinline)) int factorial(int n, int total)
{
    if (n == 1)
    {
        return total;
    }
    return factorial(n - 1, n * total);
}

__attribute__((noinline)) int main()
{
    printf("%d\n", factorial(5, 1));
    return 0;
}