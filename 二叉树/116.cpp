/*
116. 填充每个节点的下一个右侧节点指针
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        if(root == nullptr){
            return nullptr;
        }
        vector<Node*>* a = new vector<Node*>;
        vector<Node*>* b = new vector<Node*>;
        a->push_back(root);
        while(a->size()>0){
            swap(a, b);
            for(int i = 0; i< b->size(); i++){
                if(i < b->size()-1){
                    b->at(i)->next = b->at(i+1);
                }
                if (b->at(i)->left != nullptr){
                    a->push_back(b->at(i)->left);
                }
                if(b->at(i)->right != nullptr){
                    a->push_back(b->at(i)->right);
                }
            }
            b->clear();
        }
        return root;
    }
};