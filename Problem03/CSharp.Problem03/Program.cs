
using System.Security.AccessControl;

internal class Program
{
    private static void Main(string[] args)
    {
        string s = "abceeabdcxbb";
        System.Console.WriteLine(LengthOfLongestSubstring(s));
    }

    private static int LengthOfLongestSubstring(string s)
    {
        Dictionary<char, int> map = [];
        int max = 0;
        int left = 0;

        for (int i = 0; i < s.Length; i++)
        {
            if (map.ContainsKey(s[i]) && map[s[i]] >= left)
            {
                left = map[s[i]] + 1;
            }
            map[s[i]] = i;
            int len = i - left + 1;
            if (len > max)
            {
                max = len;
            }
        }

        return max;
    }
}