type Node struct{
    Key int
    Val int
    NextNode *Node
}

type LRUCache struct {
    MaxNum int
    Chache map[int]*Node
    ListRoot *Node
    LastNode *Node
}


func Constructor(capacity int) LRUCache {
    ret := LRUCache{}
    ret.MaxNum = capacity
    ret.Chache = make(map[int]*Node, capacity)
    ret.ListRoot = &Node{}
    ret.LastNode = ret.ListRoot
    return ret
}

func (this* LRUCache) New(key, value int) *Node{
    tmp := &Node{
        Key:key,
        Val:value,
    }
    this.LastNode.NextNode = tmp
    this.LastNode = tmp
    return tmp
}

func (this* LRUCache) Swap(node *Node){
    if node.NextNode == nil{
        return
    }
    a := node.NextNode
    node.NextNode = a.NextNode
    a.Key, node.Key = node.Key, a.Key
    a.Val, node.Val = node.Val, a.Val
    this.Chache[node.Key], this.Chache[a.Key] = node,a
    a.NextNode = nil
    if node.NextNode == nil{
        this.LastNode = node
    }
    this.LastNode.NextNode = a
    this.LastNode = a
    return
}

func (this *LRUCache)Del(){
    if this.ListRoot.NextNode == nil{
        return
    }
    this.ListRoot = this.ListRoot.NextNode
    delete(this.Chache, this.ListRoot.Key)
}

func (this *LRUCache) Get(key int) int {
    if ret, ok := this.Chache[key];ok{
        this.Swap(ret)
        return this.Chache[key].Val
    }else{
        return -1
    }
}


func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.Chache[key]; ok{
        node.Val = value
        this.Swap(node)
        return
    }
    if len(this.Chache) >= this.MaxNum{
        this.Del()
    }
   this.Chache[key] = this.New(key, value)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */