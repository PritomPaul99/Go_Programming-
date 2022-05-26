#include <stdio.h>

int main()
{
    int a, b;
    scanf("%d %d", &a, &b);
    if (a + b >= 180 || a <= 0 || b <= 0)
    {
        printf("Error! Enter valid angles!");
    }
    printf("3rd angle: %d", (180 - (a + b)));
    return 0;
}