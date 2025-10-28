/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *current = head;
        if (head.next == NULL) {
            return false
        }
        else {
            ListNode* slow = head;
            ListNode* fast = head;
            for (ListNode *current = head; (slow == fast)||(current == NULL); { current = current->next } ) {
                
            }
        }
    }
};