package questions;

import lombok.AccessLevel;
import lombok.NonNull;
import lombok.RequiredArgsConstructor;

/**
 * Created by user3301 on 9/28/2017.
 */

@RequiredArgsConstructor
class ListNode {
    @NonNull
    int val;
    ListNode next;
}

public class DeleteNodeInALinkedList {

    public void deleteNode(ListNode node) {
        node.val = node.next.val;
        node.next = node.next.next;
    }
}
