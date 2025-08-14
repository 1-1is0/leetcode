# Definition for singly-linked list.
from typing import Optional
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        num1 = self.to_num(l1)
        num2 = self.to_num(l2)
        val = num1 + num2
        l = self.create_list(val)
        # print(val)
        # print(self.to_num(l))
        return l

        
    def to_num(self, l1: Optional[ListNode]):
        num = str(l1.val)
        while l1.next:
            l1 = l1.next
            num += str(l1.val)
        return int(num[::-1])

    def create_list(self, val: int):
        prev = None
        for i in str(val):
            prev = ListNode(int(i), prev)
        return prev

class Solution2:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = ListNode()
        current = head
        carry = 0
        while l1 != None or l2 != None or carry!=0:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0
            result = val1 + val2 + carry
            carry = result // 10
            next_node = ListNode(result % 10, None)
            current.next = next_node
            current = next_node
            # go to the next element
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return head.next

    def to_num(self, l1: Optional[ListNode]):
        num = str(l1.val)
        while l1.next:
            l1 = l1.next
            num += str(l1.val)
        return int(num[::-1])



        


        

l1 = ListNode(2, ListNode(4, ListNode(3)))
l2 = ListNode(5, ListNode(6, ListNode(4)))
sol = Solution2()
print(sol.addTwoNumbers(l1, l2))