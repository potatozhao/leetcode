/*
208. 实现 Trie (前缀树)
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");   
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

type Trie struct {
    NodeList []*Trie
    IsEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{NodeList:make([]*Trie, 26)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    nextNode := this
    for i := range word{
        num := int(word[i] - 'a')
        if nextNode.NodeList[num] == nil{
            nextNode.NodeList[num] = new(Trie)
            nextNode.NodeList[num].NodeList = make([]*Trie, 26)
        }
        nextNode = nextNode.NodeList[num]
    }
    nextNode.IsEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    nextNode := this
    for i := range word{
        num := int(word[i] - 'a')
        if nextNode.NodeList[num] == nil{
            return false
        }
        nextNode = nextNode.NodeList[num]
    }
    if nextNode.IsEnd {
        return true
    }
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    nextNode := this
    for i := range prefix{
        num := int(prefix[i] - 'a')
        if nextNode.NodeList[num] == nil{
            return false
        }
        nextNode = nextNode.NodeList[num]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */