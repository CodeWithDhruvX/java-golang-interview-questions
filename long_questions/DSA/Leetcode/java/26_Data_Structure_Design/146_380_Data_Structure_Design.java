import java.util.*;

public class DataStructureDesign {
    
    // 146. LRU Cache
    // Time: O(1) for get and put, Space: O(capacity)
    public static class LRUCache {
        private int capacity;
        private Map<Integer, Node> cache;
        private Node head;
        private Node tail;
        
        private class Node {
            int key;
            int value;
            Node prev;
            Node next;
            
            Node(int key, int value) {
                this.key = key;
                this.value = value;
            }
        }
        
        public LRUCache(int capacity) {
            this.capacity = capacity;
            this.cache = new HashMap<>();
            this.head = new Node(0, 0);
            this.tail = new Node(0, 0);
            head.next = tail;
            tail.prev = head;
        }
        
        public int get(int key) {
            if (!cache.containsKey(key)) {
                return -1;
            }
            
            Node node = cache.get(key);
            moveToHead(node);
            return node.value;
        }
        
        public void put(int key, int value) {
            if (cache.containsKey(key)) {
                Node node = cache.get(key);
                node.value = value;
                moveToHead(node);
            } else {
                Node newNode = new Node(key, value);
                cache.put(key, newNode);
                addToHead(newNode);
                
                if (cache.size() > capacity) {
                    Node tailNode = removeTail();
                    cache.remove(tailNode.key);
                }
            }
        }
        
        private void addToHead(Node node) {
            node.prev = head;
            node.next = head.next;
            head.next.prev = node;
            head.next = node;
        }
        
        private void removeNode(Node node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
        }
        
        private void moveToHead(Node node) {
            removeNode(node);
            addToHead(node);
        }
        
        private Node removeTail() {
            Node lastNode = tail.prev;
            removeNode(lastNode);
            return lastNode;
        }
    }

    // 380. Insert Delete GetRandom O(1)
    // Time: O(1) for all operations, Space: O(N)
    public static class RandomizedSet {
        private List<Integer> list;
        private Map<Integer, Integer> map;
        private Random random;
        
        public RandomizedSet() {
            list = new ArrayList<>();
            map = new HashMap<>();
            random = new Random();
        }
        
        public boolean insert(int val) {
            if (map.containsKey(val)) {
                return false;
            }
            
            map.put(val, list.size());
            list.add(val);
            return true;
        }
        
        public boolean remove(int val) {
            if (!map.containsKey(val)) {
                return false;
            }
            
            int index = map.get(val);
            int lastElement = list.get(list.size() - 1);
            
            // Move last element to the position of element to remove
            list.set(index, lastElement);
            map.put(lastElement, index);
            
            // Remove last element
            list.remove(list.size() - 1);
            map.remove(val);
            
            return true;
        }
        
        public int getRandom() {
            return list.get(random.nextInt(list.size()));
        }
    }

    public static void main(String[] args) {
        // Test cases for LRUCache
        System.out.println("LRU Cache Operations:");
        LRUCache lruCache = new LRUCache(2);
        lruCache.put(1, 1);
        lruCache.put(2, 2);
        System.out.printf("Get 1: %d\n", lruCache.get(1)); // 1
        lruCache.put(3, 3); // evicts key 2
        System.out.printf("Get 2: %d\n", lruCache.get(2)); // -1 (not found)
        lruCache.put(4, 4); // evicts key 1
        System.out.printf("Get 1: %d\n", lruCache.get(1)); // -1 (not found)
        System.out.printf("Get 3: %d\n", lruCache.get(3)); // 3
        System.out.printf("Get 4: %d\n", lruCache.get(4)); // 4
        
        // Test cases for RandomizedSet
        System.out.println("\nRandomizedSet Operations:");
        RandomizedSet randomSet = new RandomizedSet();
        System.out.printf("Insert 1: %b\n", randomSet.insert(1)); // true
        System.out.printf("Remove 2: %b\n", randomSet.remove(2)); // false
        System.out.printf("Insert 2: %b\n", randomSet.insert(2)); // true
        System.out.printf("GetRandom: %d\n", randomSet.getRandom()); // 1 or 2
        System.out.printf("Remove 1: %b\n", randomSet.remove(1)); // true
        System.out.printf("Insert 2: %b\n", randomSet.insert(2)); // false
        System.out.printf("GetRandom: %d\n", randomSet.getRandom()); // 2
        
        // Additional test for LRUCache
        System.out.println("\nAdditional LRU Cache Test:");
        LRUCache lruCache2 = new LRUCache(3);
        lruCache2.put(1, 1);
        lruCache2.put(2, 2);
        lruCache2.put(3, 3);
        System.out.printf("Get 2: %d\n", lruCache2.get(2)); // 2
        lruCache2.put(4, 4); // evicts key 3
        System.out.printf("Get 1: %d\n", lruCache2.get(1)); // 1
        System.out.printf("Get 3: %d\n", lruCache2.get(3)); // -1 (not found)
        System.out.printf("Get 4: %d\n", lruCache2.get(4)); // 4
        
        // Additional test for RandomizedSet
        System.out.println("\nAdditional RandomizedSet Test:");
        RandomizedSet randomSet2 = new RandomizedSet();
        randomSet2.insert(10);
        randomSet2.insert(20);
        randomSet2.insert(30);
        System.out.printf("Remove 20: %b\n", randomSet2.remove(20)); // true
        System.out.printf("Remove 20: %b\n", randomSet2.remove(20)); // false
        System.out.printf("Insert 20: %b\n", randomSet2.insert(20)); // true
        System.out.printf("GetRandom (multiple calls): ");
        for (int i = 0; i < 5; i++) {
            System.out.printf("%d ", randomSet2.getRandom());
        }
        System.out.println();
    }
}
