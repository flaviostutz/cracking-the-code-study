class Node:
    prev = None
    nxt = None
    value = None
    key = None

class LRUCache:

    # @param capacity, an integer
    def __init__(self, capacity):
        self.cap = capacity
        self.head = None
        self.tail = None
        self.nodes = {}

    # @return an integer
    def get(self, key):
        if key in self.nodes:
            v = self.nodes[key]
            self.deleteNode(v)
            self.insertNode(v)
            return v.value
        return -1

    # @param key, an integer
    # @param value, an integer
    # @return nothing
    def set(self, key, value):
        if key in self.nodes:
            v = self.nodes[key]
            self.deleteNode(v)
            self.insertNode(v)
            v.value = value
        else:
            if self.tail is not None and len(self.nodes) >= self.cap:
                # print("EVICT ENTRY ", self.nodes.keys(), self.tail, self.cap, len(self.nodes))
                self.nodes.pop(self.tail.key, None)
                self.deleteNode(self.tail)
            
            n = Node()
            n.value = value
            n.key = key
            self.insertNode(n)
            self.nodes[key] = n

    def deleteNode(self, node):
        if node.prev is None:
            self.head = node.nxt
        else:
            node.prev.nxt = node.nxt

        if node.nxt is None:
            self.tail = node.prev
        else:
            node.nxt.prev = node.prev

        node.prev = None
        node.nxt = None
        
    def insertNode(self, node):
        node.nxt = self.head
        node.prev = None
        if self.head is not None:
            self.head.prev = node
        self.head = node
        if self.tail is None:
            self.tail = node


