/*
面试题 16.25. LRU缓存
难度
中等

23





设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:

LRUCache cache = new LRUCache( 2 缓存容量 );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
*/


struct Node{
    int val;
    int key;
    Node* next;
    Node* prev;
};
class LRUCache {
public:
    LRUCache(int capacity) {
        this->capacity = capacity;
        this->thead = new Node;
        this->last = new(Node);
        this->thead->next = this->last;
        this->last->prev = this->thead;
    }
    
    int get(int key) {
        map<int, Node*>::iterator it = list_map.find(key);
        if(it == this->list_map.end()){
            return -1;
        }
        Node *tmp = it->second;
        get_node(tmp);
        put_front(tmp);
        return tmp->val;
    }
    
    void put(int key, int value) {
        map<int, Node*>::iterator it = list_map.find(key);
        if(it != list_map.end()){
            it->second->val = value;
            get_node(it->second);
            put_front(it->second);
            return;
        }
        if(list_map.size() == this->capacity){
            Node* deleted_node = delete_map_key(delete_last());
            if(deleted_node){
                delete deleted_node;
            }
        }
        Node *tmp = new Node;
        tmp->val = value;
        tmp->key = key;
        list_map[key] = tmp;
        put_front(tmp);
        return;
    }

    // 从链表取出节点
    Node* get_node(Node *node){
        node->next->prev = node->prev;
        node->prev->next = node->next;
        return node;
    }

    // 放回链表头部。
    void put_front(Node *node){
        node->next = thead->next;
        node->prev = thead;
        thead->next->prev = node;
        thead->next = node;
        return;
    }

    int delete_last(){
        Node* tmp = last->prev;
        return get_node(tmp)->key;
    }

    Node* delete_map_key(int key){
        map<int, Node*>::iterator it = list_map.find(key);
        if(it == list_map.end()){
            return NULL;
        }
        Node* tmp = it->second;
        list_map.erase(it);
        return tmp;
    }

    int capacity;
    Node* thead;// 链表头
    Node* last;// 链表尾
    map<int, Node*> list_map;
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */