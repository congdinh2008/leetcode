internal class Program
{
    private static void Main(string[] args)
    {
        int[] nums1 = [3, 4];
        int[] nums2 = []; // Example: new int[] {1,2,3,4,5,9};
        Console.WriteLine(FindMedianSortedArrays(nums1, nums2));
    }

    public static double FindMedianSortedArrays(int[] nums1, int[] nums2)
    {
        int len1 = nums1.Length, len2 = nums2.Length;
        int resultLen = len1 + len2;
        int[] result = new int[resultLen];

        int i = 0, j = 0, k = 0;
        while (i < len1 || j < len2)
        {
            if (i < len1 && (j >= len2 || nums1[i] <= nums2[j]))
            {
                result[k++] = nums1[i++];
            }
            else if (j < len2)
            {
                result[k++] = nums2[j++];
            }
        }

        if (resultLen % 2 == 1)
        {
            return result[resultLen / 2];
        }
        else
        {
            return (result[resultLen / 2] + result[resultLen / 2 - 1]) / 2.0;
        }
    }
}