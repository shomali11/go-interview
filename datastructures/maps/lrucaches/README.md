### Description

Implement an LRU Cache. Similar to a Map in storing and retrieving data efficiently, but an LRU (Least Recently Used) cache evicts the least recently used entry when it is full.

### Example:

```
Size 3

Input: Put ("key1", "1"), Put ("key2", "2"), Put ("key3", "3"), Get "key1", Put ("key4", "4"), GetEntries
Output: [("key4", "4"), ("key1", "1"), ("key3", "3")]
```