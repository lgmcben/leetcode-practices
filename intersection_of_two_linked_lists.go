// Challenge link:
// https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/785/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    
    mapA := make(map[*ListNode]*ListNode)

     
    // Make a map of list A
      current := headA
      for current != nil {
            // fmt.Println(current.Val)

            

            mapA[current] = current

            
            // fmt.Println(mapA)
            current = current.Next
      }
    
    
   // Iterate through list B, check for duplicate address
      current = headB
      for current != nil {
            // fmt.Println(current.Val)

          intersectNode, ok := mapA[current]
          
          if ok {
              return intersectNode
          }

          current = current.Next
      } 
    

    return nil
}
