internal class Program
{
    private static void Main(string[] args)
    {
        int[] nums1 = [3, 4];
        int[] nums2 = []; // Example: new int[] {1,2,3,4,5,9};
        Console.WriteLine(FindMedianSortedArraysTwoPointer(nums1, nums2));
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

    public static double FindMedianSortedArraysTwoPointer(int[] nums1, int[] nums2)
    {
        int len1 = nums1.Length, len2 = nums2.Length;
        int total = len1 + len2;

        int i = 0, j = 0;
        double median1 = 0, median2 = 0;
        for (int counter = 0; counter <= total / 2; counter++)
        {
            median2 = median1;
            if (i == len1 || (j < len2 && nums1[i] >= nums2[j])){
                median1 = nums2[j];
                j++;
            } else {
                median1 = nums1[i];
                i++;
            }
        }

        if (total % 2 == 1)
        {
            return median1;
        }
        else
        {
            return (median1 + median2) / 2;
        }
    }
}