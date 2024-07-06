using System.Text;

internal class Program
{
    private static void Main(string[] args)
    {
        int num = 3999;

        string roman = IntToRoman02(num);
        System.Console.WriteLine(roman);
    }

    public static string IntToRoman(int num)
    {
        Dictionary<int, string> romans = new() {
            {1, "I"},
            {2, "II"},
            {3, "III"},
            {4, "IV"},
            {5, "V"},
            {6, "VI"},
            {7, "VII"},
            {8, "VIII"},
            {9, "IX"},
            {10, "X"},
            {40, "XL"},
            {50, "L"},
            {90, "XC"},
            {100, "C"},
            {400, "CD"},
            {500, "D"},
            {900, "CM"},
            {1000, "M"}
        };
        StringBuilder result = new();

        int step = 1;
        while (num > 0)
        {
            int remainder = num % (10 * step);
            if (romans.ContainsKey(remainder))
            {
                result.Insert(0, romans[remainder]);
                num -= remainder;
            }
            else
            {
                num -= remainder;
                while (remainder > 0)
                {
                    int temp = remainder - step;
                    if (romans.ContainsKey(temp))
                    {
                        result.Insert(0, romans[temp] + romans[step]);
                        remainder -= temp + step;
                    }
                    else
                    {
                        result.Insert(0, romans[step]);
                        remainder -= step;
                    }
                }
            }
            step *= 10;
        }
        return result.ToString();
    }

    public static string IntToRoman02(int num)
    {
        Dictionary<int, string> romans = new() {
            {1000, "M"},
            {900, "CM"},
            {500, "D"},
            {400, "CD"},
            {100, "C"},
            {90, "XC"},
            {50, "L"},
            {40, "XL"},
            {10, "X"},
            {9, "IX"},
            {5, "V"},
            {4, "IV"},
            {1, "I"}
        };
        StringBuilder result = new();

        foreach (var item in romans.Keys)
        {
            while (num >= item)
            {
                result.Append(romans[item]);
                num -= item;
            }
        }
        return result.ToString();
    }
}