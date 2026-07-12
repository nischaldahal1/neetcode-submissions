func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    // Initialize the queue with the root node
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        // Pop the front element from the queue
        curr := queue[0]
        queue = queue[1:]

        // Swap the left and right children
        curr.Left, curr.Right = curr.Right, curr.Left

        // If children exist, add them to the queue to be processed later
        if curr.Left != nil {
            queue = append(queue, curr.Left)
        }
        if curr.Right != nil {
            queue = append(queue, curr.Right)
        }
    }

    return root
}