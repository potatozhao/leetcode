# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def first(self, node, str_list):
        if node == None:
            str_list.append("null")
            return
        str_list.append(str(node.val))
        self.first(node.left, str_list)
        self.first(node.right, str_list)
        return

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        str_list = []
        self.first(root, str_list)
        return ",".join(str_list)
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if len(data) ==0:
            return None
        str_list = data.split(",")
        root = self.last(str_list)
        return root

    def last(self, str_list):
        if str_list[0] == 'null':
            str_list.pop(0)
            return None
        node = TreeNode(str_list[0])
        str_list.pop(0)
        node.left = self.last(str_list)
        node.right = self.last(str_list)
        return node

        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))