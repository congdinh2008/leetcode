internal class Program
{
    private static void Main(string[] args)
    {
        string s = "CXIX";
        int result = RomanToInt02(s);
        System.Console.WriteLine(result);
    }

    public static int RomanToInt02(string s)
    {
        Dictionary<string, int> romans = new Dictionary<string, int>() {
            {"I", 1},
            {"V", 5},
            {"X", 10},
            {"L", 50},
            {"C", 100},
            {"D", 500},
            {"M", 1000}
        };

        int number = 0;

        for (int i = 0; i < s.Length - 1; i++)
        {
            if (romans[s[i].ToString()] < romans[s[i + 1].ToString()])
            {
                number -= romans[s[i].ToString()];
            }
            else
            {
                number += romans[s[i].ToString()];
            }
        }

        return number + romans[s[^1].ToString()];
    }

    public static int RomanToInt(string s)
    {
        Dictionary<string, int> romans = new Dictionary<string, int>() {
            {"I", 1},
            {"II", 2},
            {"III", 3},
            {"IV", 4},
            {"V", 5},
            {"VI", 6},
            {"VII", 7},
            {"VIII", 8},
            {"IX", 9},
            {"X", 10},
            {"XL", 40},
            {"L", 50},
            {"XC", 90},
            {"C", 100},
            {"CD", 400},
            {"D", 500},
            {"CM", 900},
            {"M", 1000}
        };

        if (romans.ContainsKey(s))
        {
            return romans[s];
        }
        else
        {
            string temp = string.Empty;
            int number = 0;
            int i = s.Length - 1;
            bool consecutive = false;
            while (i >= 0)
            {
                temp = s[i] + temp;
                if (romans.ContainsKey(temp))
                {
                    if (consecutive)
                    {
                        number = romans[temp];
                        consecutive = true;

                    }
                    else
                    {
                        number += romans[s[i].ToString()];
                        consecutive = true;

                        if (i > 0)
                        {
                            switch (s[i - 1].ToString() + s[i].ToString())
                            {
                                case "CM":
                                case "CD":
                                    number -= 100;
                                    consecutive = false;
                                    temp = string.Empty;
                                    i--;
                                    break;
                                case "XC":
                                case "XL":
                                    number -= 10;
                                    i--;
                                    consecutive = false;
                                    temp = string.Empty;
                                    break;
                            }
                        }
                    }
                }
                else
                {
                    consecutive = false;
                    number += romans[s[i].ToString()];
                    if (i > 0)
                    {
                        switch (s[i - 1].ToString() + s[i].ToString())
                        {
                            case "CM":
                            case "CD":
                                number -= 100;
                                i--;
                                break;
                            case "XC":
                            case "XL":
                                number -= 10;
                                i--;
                                break;
                        }
                    }
                    temp = string.Empty;
                }
                i--;
            }
            return number;
        }
    }
}