public class Program
{
    private static void Main(string[] args)
    {
        Console.WriteLine(LongestPalindromeDP02("abadba"));
    }

    static string LongestPalindromeDP02(string s)
    {
        int n = s.Length;
        int startIndex = 0;
        int maxlen = 0;
        bool[,] table = new bool[n, n];

        for (int g = 0; g < n; g++)
        {
            for (int i = 0, j = g; j < n; i++, j++)
            {
                Solve(table, i, j, s);
                if (table[i, j])
                {
                    if (j - i + 1 > maxlen)
                    {
                        startIndex = i;
                        maxlen = j - i + 1;
                    }
                }
            }
        }
        return s.Substring(startIndex, maxlen);
    }

    static void Solve(bool[,] table, int i, int j, string s)
    {
        if (i == j)
        {
            table[i, j] = true;
            return;
        }
        if (j - i == 1)
        {
            table[i, j] = s[i] == s[j];
            return;
        }
        table[i, j] = s[i] == s[j] && table[i + 1, j - 1];
    }

    static string LongestPalindrome(string s)
    {
        string result = "";

        for (int i = 0; i < s.Length; i++)
        {
            string temp = s.Substring(i);
            int tempLen = temp.Length;
            while (tempLen > 0)
            {
                if (IsPalindrome(temp))
                {
                    if (temp.Length > result.Length)
                    {
                        result = temp;
                    }
                    break;
                }
                else
                {
                    tempLen--;
                    temp = temp.Substring(0, tempLen);
                }
            }
        }
        return result;
    }

    static bool IsPalindrome(string s)
    {
        for (int i = 0; i < s.Length / 2; i++)
        {
            if (s[i] != s[s.Length - 1 - i])
            {
                return false;
            }
        }
        return true;
    }
}