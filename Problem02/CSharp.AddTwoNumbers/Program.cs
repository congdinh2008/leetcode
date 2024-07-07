namespace CSharp.AddTwoNumbers;
internal class Program
{
    private static void Main(string[] args)
    {
        ListNode listNode01 = new(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))))));
        ListNode listNode02 = new(9, new ListNode(9, new ListNode(9, new ListNode(9))));

        // ListNode listNode01 = new(2, new ListNode(4, new ListNode(3)));
        // ListNode listNode02 = new(5, new ListNode(6, new ListNode(4)));

        ListNode? sum = AddTwoNumbers(listNode01, listNode02);

        while (sum != null)
        {
            Console.Write(sum.val + " ");
            sum = sum.next;
        }
    }

    public static ListNode AddTwoNumbers(ListNode l1, ListNode l2)
    {
        ListNode result = new(0);
        ListNode current = result;
        int sum = 0;

        while (l1 != null || l2 != null)
        {
            if (l1 != null)
            {
                sum += l1.val;
                l1 = l1.next;
            }

            if (l2 != null)
            {
                sum += l2.val;
                l2 = l2.next;
            }
            result.val = sum % 10;
            sum /= 10;

            if (l1 != null || l2 != null)
            {
                result.next = new ListNode(0);
                result = result.next;
            }
        }

        if (sum > 0)
        {
            result.next = new ListNode(0);
            result = result.next;
            result.val = sum;
        }

        return current;
    }
}

public class ListNode(int val = 0, ListNode next = null)
{
    public int val = val;
    public ListNode next = next;
}