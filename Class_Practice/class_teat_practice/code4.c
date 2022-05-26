#include <stdio.h>

int main()
{
    int sum = 0;
    for (int i = 1; i < 31; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            sum += i;
        }
    }
    printf("Sum: %d", sum);
}