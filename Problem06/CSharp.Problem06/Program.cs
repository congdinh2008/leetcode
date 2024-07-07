using System.Text;

internal class Program
{
    private static void Main(string[] args)
    {
        string s = "PAYPALISHIRING";
        int numRows = 3;

        Console.WriteLine(Convert(s, numRows));
    }

    public static string Convert(string s, int numRows)
    {
        if (numRows == 1 || numRows >= s.Length)
        {
            return s;
        }

        var result = new StringBuilder();
        // The term is the number of characters in one zigzag cycle
        int term = 2 * numRows - 2;

        for (int i = 0; i < numRows; i++)
        {
            for (int j = i; j < s.Length; j += term)
            {
                // Always add the current character in the row
                result.Append(s[j]);
                // For non-first and non-last rows, add the character that appears diagonally in the zigzag
                if (i != 0 && i != numRows - 1)
                {
                    int diagIndex = j + term - 2 * i;
                    if (diagIndex < s.Length)
                    {
                        result.Append(s[diagIndex]);
                    }
                }
            }
        }

        return result.ToString()
    }
}