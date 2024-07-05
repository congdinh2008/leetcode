#include <stdio.h>
#include <stdbool.h>

bool IsPalindrome(int x)
{
    if (x < 0)
        return false;

    if (x == 0)
        return true;

    if (x % 10 == 0)
        return false;

    int temp = x;
    int reversed = 0;

    while (temp > 0)
    {
        reversed = reversed * 10 + temp % 10;
        temp /= 10;
    }

    return x == reversed;
}

int main()
{
    int x = 121;
    bool result = IsPalindrome(x);

    printf("Is %d a palindrome number? %s\n", x, result ? "Yes" : "No");
    return 0;
}