/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */
function addTwoNumbers(l1, l2) {
    let current = new ListNode();
    let result = current;
    let sum = 0;
    while (l1 != null || l2 != null) {
        if (l1 != null) {
            sum += l1.val;
            l1 = l1.next;
        }
        if (l2 != null) {
            sum += l2.val;
            l2 = l2.next;
        }
        current.val = sum % 10;
        sum = Math.floor(sum / 10);
        if (l1 != null || l2 != null) {
            current.next = new ListNode();
            current = current.next;
        }
    }
    if (sum > 0) {
        current.next = new ListNode(sum);
        current = current.next;
    }
    return result;
}
;
class ListNode {
    constructor(val, next) {
        this.val = (val === undefined ? 0 : val);
        this.next = (next === undefined ? null : next);
    }
}
const l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
const l2 = new ListNode(5, new ListNode(6, new ListNode(4)));
console.log(addTwoNumbers(l1, l2)); // ListNode { val: 7, next: ListNode { val: 0, next: ListNode { val: 8, next: null } } }
