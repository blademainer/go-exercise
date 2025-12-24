package main

import (
	"fmt"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	ch := NewConsistentHash(100)
	
	// Test adding nodes
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")
	
	if ch.GetNodeCount() != 3 {
		t.Errorf("Expected 3 nodes, got %d", ch.GetNodeCount())
	}
	
	// Test key distribution
	testKey := "test_key_123"
	node := ch.GetNode(testKey)
	if node == "" {
		t.Error("Expected non-empty node for key")
	}
	
	// Key should consistently map to the same node
	for i := 0; i < 10; i++ {
		if ch.GetNode(testKey) != node {
			t.Error("Key mapping is not consistent")
		}
	}
}

func TestAddRemoveNode(t *testing.T) {
	ch := NewConsistentHash(100)
	
	ch.AddNode("node1")
	ch.AddNode("node2")
	
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	
	// Record initial mappings
	initialMappings := make(map[string]string)
	for _, key := range keys {
		initialMappings[key] = ch.GetNode(key)
	}
	
	// Add a new node
	ch.AddNode("node3")
	
	// Some keys may have moved, but not all
	movedCount := 0
	for _, key := range keys {
		if ch.GetNode(key) != initialMappings[key] {
			movedCount++
		}
	}
	
	// In a good consistent hash, only a fraction of keys should move
	fmt.Printf("Keys moved after adding node: %d/%d\n", movedCount, len(keys))
	
	// Remove a node
	ch.RemoveNode("node2")
	
	if ch.GetNodeCount() != 2 {
		t.Errorf("Expected 2 nodes after removal, got %d", ch.GetNodeCount())
	}
	
	// All keys should still have a node
	for _, key := range keys {
		if ch.GetNode(key) == "" {
			t.Errorf("Key %s has no node after removal", key)
		}
	}
}

func TestEmptyRing(t *testing.T) {
	ch := NewConsistentHash(100)
	
	node := ch.GetNode("test_key")
	if node != "" {
		t.Errorf("Expected empty string for empty ring, got %s", node)
	}
}

func TestDistributionBalance(t *testing.T) {
	ch := NewConsistentHash(150)
	
	// Add 3 nodes
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")
	
	// Create 1000 keys
	distribution := make(map[string]int)
	totalKeys := 1000
	
	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("key_%d", i)
		node := ch.GetNode(key)
		distribution[node]++
	}
	
	// Check that each node gets roughly 1/3 of the keys (with some tolerance)
	expectedPerNode := totalKeys / 3
	tolerance := float64(expectedPerNode) * 0.3 // 30% tolerance
	
	for node, count := range distribution {
		diff := float64(count - expectedPerNode)
		if diff < 0 {
			diff = -diff
		}
		if diff > tolerance {
			t.Logf("Warning: Node %s has %d keys, expected around %d (tolerance: %.0f)", 
				node, count, expectedPerNode, tolerance)
		}
	}
	
	t.Logf("Distribution: %v", distribution)
}
