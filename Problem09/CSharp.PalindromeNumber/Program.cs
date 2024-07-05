internal class Program
{
    private static void Main(string[] args)
    {
        int x = 121;
        bool result = IsPalindromeString(x);
        Console.WriteLine(result);
    }

    public static bool IsPalindrome(int x)
    {
        if (x < 0) return false;

        if (x == 0) return true;

        if (x % 10 == 0) return false;

        int temp = x;
        int reversed = 0;

        while (temp > 0)
        {
            reversed = reversed * 10 + temp % 10;
            temp /= 10;
        }

        return x == reversed;
    }

    public static bool IsPalindromeString(int x)
    {
        if (x < 0) return false;

        if (x == 0) return true;

        if (x % 10 == 0) return false;

        string s = x.ToString();
        int left = 0;
        int right = s.Length - 1;

        while (left < right)
        {
            if (s[left] != s[right]) return false;

            left++;
            right--;
        }

        return true;
    }
}