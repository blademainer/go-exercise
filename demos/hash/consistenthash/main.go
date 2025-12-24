package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
)

// ConsistentHash represents a consistent hash ring
type ConsistentHash struct {
	hashRing     map[uint32]string // hash value -> node name
	sortedHashes []uint32          // sorted hash values
	virtualNodes int               // number of virtual nodes per physical node
}

// NewConsistentHash creates a new consistent hash instance
func NewConsistentHash(virtualNodes int) *ConsistentHash {
	return &ConsistentHash{
		hashRing:     make(map[uint32]string),
		sortedHashes: make([]uint32, 0),
		virtualNodes: virtualNodes,
	}
}

// hash calculates hash value for a given key
func (ch *ConsistentHash) hash(key string) uint32 {
	h := sha256.New()
	h.Write([]byte(key))
	hashBytes := h.Sum(nil)
	// Use first 4 bytes to create uint32
	return uint32(hashBytes[0])<<24 | uint32(hashBytes[1])<<16 | uint32(hashBytes[2])<<8 | uint32(hashBytes[3])
}

// AddNode adds a physical node to the hash ring
func (ch *ConsistentHash) AddNode(nodeName string) {
	// Add virtual nodes for this physical node
	for i := 0; i < ch.virtualNodes; i++ {
		virtualKey := nodeName + "#" + strconv.Itoa(i)
		hashValue := ch.hash(virtualKey)
		ch.hashRing[hashValue] = nodeName
		ch.sortedHashes = append(ch.sortedHashes, hashValue)
	}
	// Sort the hash values
	sort.Slice(ch.sortedHashes, func(i, j int) bool {
		return ch.sortedHashes[i] < ch.sortedHashes[j]
	})
	fmt.Printf("Node added: %s (with %d virtual nodes)\n", nodeName, ch.virtualNodes)
}

// RemoveNode removes a physical node from the hash ring
func (ch *ConsistentHash) RemoveNode(nodeName string) {
	// Remove all virtual nodes for this physical node
	for i := 0; i < ch.virtualNodes; i++ {
		virtualKey := nodeName + "#" + strconv.Itoa(i)
		hashValue := ch.hash(virtualKey)
		delete(ch.hashRing, hashValue)
		// Remove from sorted hashes
		for idx, hash := range ch.sortedHashes {
			if hash == hashValue {
				ch.sortedHashes = append(ch.sortedHashes[:idx], ch.sortedHashes[idx+1:]...)
				break
			}
		}
	}
	fmt.Printf("Node removed: %s\n", nodeName)
}

// GetNode returns the node responsible for the given key
func (ch *ConsistentHash) GetNode(key string) string {
	if len(ch.sortedHashes) == 0 {
		return ""
	}
	
	hashValue := ch.hash(key)
	
	// Binary search to find the first node with hash >= key's hash
	idx := sort.Search(len(ch.sortedHashes), func(i int) bool {
		return ch.sortedHashes[i] >= hashValue
	})
	
	// If we've gone past the end, wrap around to the first node
	if idx == len(ch.sortedHashes) {
		idx = 0
	}
	
	return ch.hashRing[ch.sortedHashes[idx]]
}

// PrintDistribution shows the distribution of keys across nodes
func (ch *ConsistentHash) PrintDistribution(keys []string) {
	distribution := make(map[string]int)
	
	for _, key := range keys {
		node := ch.GetNode(key)
		distribution[node]++
	}
	
	fmt.Println("\n=== Key Distribution ===")
	for node, count := range distribution {
		percentage := float64(count) / float64(len(keys)) * 100
		fmt.Printf("Node %s: %d keys (%.2f%%)\n", node, count, percentage)
	}
	fmt.Println()
}

// GetNodeCount returns the number of physical nodes
func (ch *ConsistentHash) GetNodeCount() int {
	nodeSet := make(map[string]bool)
	for _, node := range ch.hashRing {
		nodeSet[node] = true
	}
	return len(nodeSet)
}

func main() {
	fmt.Println("=== Consistent Hash Demo ===")
	
	// Create a consistent hash with 150 virtual nodes per physical node
	ch := NewConsistentHash(150)
	
	// Simulate some data keys
	keys := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		keys = append(keys, fmt.Sprintf("key_%d", i))
	}
	
	fmt.Println("\n--- Step 1: Add initial nodes ---")
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")
	ch.PrintDistribution(keys)
	
	// Show some key mappings
	fmt.Println("Sample key mappings:")
	for i := 0; i < 5; i++ {
		key := keys[i]
		node := ch.GetNode(key)
		fmt.Printf("  %s -> %s\n", key, node)
	}
	
	fmt.Println("\n--- Step 2: Add a new node ---")
	ch.AddNode("node4")
	ch.PrintDistribution(keys)
	
	// Show how mappings changed
	fmt.Println("Sample key mappings after adding node4:")
	for i := 0; i < 5; i++ {
		key := keys[i]
		node := ch.GetNode(key)
		fmt.Printf("  %s -> %s\n", key, node)
	}
	
	fmt.Println("\n--- Step 3: Remove a node ---")
	ch.RemoveNode("node2")
	ch.PrintDistribution(keys)
	
	// Show how mappings changed
	fmt.Println("Sample key mappings after removing node2:")
	for i := 0; i < 5; i++ {
		key := keys[i]
		node := ch.GetNode(key)
		fmt.Printf("  %s -> %s\n", key, node)
	}
	
	fmt.Println("\n--- Step 4: Add multiple nodes ---")
	ch.AddNode("node5")
	ch.AddNode("node6")
	ch.PrintDistribution(keys)
	
	fmt.Println("\n--- Step 5: Remove multiple nodes ---")
	ch.RemoveNode("node4")
	ch.RemoveNode("node5")
	ch.PrintDistribution(keys)
	
	fmt.Printf("\nFinal node count: %d\n", ch.GetNodeCount())
	fmt.Println("\n=== Demo Complete ===")
}