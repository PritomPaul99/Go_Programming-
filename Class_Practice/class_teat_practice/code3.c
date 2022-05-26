#include <stdio.h>

int main()
{
    int a;
    scanf("%d", &a);

    int hr = a / 60;

    int min = a - hr * 60;

    printf("Hr: %d Min: %d", hr, min);
}