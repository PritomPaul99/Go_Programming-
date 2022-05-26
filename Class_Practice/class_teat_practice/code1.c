#include <stdio.h>
#include <math.h>

int main()
{
    int a, b;
    scanf("%d %d", &a, &b);
    int sqr = pow((a + b), 2);
    printf("Squire of  (a+b): %d", sqr);
    return 0;
}