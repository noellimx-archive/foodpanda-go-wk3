# Heap Heap Hoorah

Heaps are rooted graphs that have a wide range of applications, mainly for quick ranking of an unordered set of elements or nodes. It consists of nodes and edges, with a parent node unidirectionally connected from itself to its children nodes. All children nodes satisfies an inequality invariant with respect to its parent.

There are a variety of heaps . Some heap structures are theoretically more efficient, but exists only in literature.

## Example - Max Binary-Heap

An simple example will be an integer max [binary]-heap, which has the invariant property of children nodes must be greater than [or equal to] its parent.

    Heap H
                                             Value                                                Index

Level 0 (Root)                                10                                                    0
                                            /    \                                                /    \
Level 1                                    7      4                                              1      2
                                         /  \    /  \                                          /  \    /  \
Level 2                                 2    5  1    3                                        3    4  5    6
                                                                                            /   \
                                                                                           7    ................


### Positions

For implementation of heaps, node index notation starts from the root. A child node can derive the index of the parent and vice versa.

idx_p = (idx_c) => floor((idx_c - 1) / 2)

idx_child_left = (idx_p) => idx_p * 2 + 1
idx_child_right = (idx_p) => idx_p * 2 + 2

### In Comparison with Binary Search Tree

#### Ordering

The ordering is preserved top-down (or bottom-up) instead of left-right. i.e left-most branch: 10 -> 7 -> 2    ===>     10 > 7 > 2

#### Height

Height of the tree is similar to a tree in the order of O(log_n k) where n is maximum no. of children and k is no. of nodes.

### Operations

# Node Insertion

A new node alters the tree without violation of invariant between parent and children nodes.

First, the new node is placed from the left-most position in the lowest incomplete level and a provisional edge is formed from the parent (see (positions)[#Positions]). The new node may violate the heap structure. Consider adding a new node with value 9 at index 7 in Heap H - the altered branch from root is 10 -> 7 -> 2 -> 9 with heap invariant violated, i.e 10 > 7 > 2 !> [9].


    Heap H

Level 0 (Root)                                10                                                   
                                            /    \                                              
Level 1                                    7      4                                            
                                         /  \    /  \                                          
Level 2                                 2    5  1    3
                                       /
[Level 3]                            [9]


Second, a "bubble-up" operation is performed, where the new node will swap with its parent if heap invariant does not hold. The new node will promote until a parent relationship satifies heap-invariant.

    Heap H (left-most branch)

                                         2 !> 9                        7 !> 9                    10 > 9  (OK)

Level 0 (Root)                               10                          10                        10                             
                                            /                           /                         /
Level 1                                    7   ...           =>        7   ...        =>         [9]  ...         
                                         /   \                        /  \                      /   \
Level 2                                 2      5                    [9]   5                    7     5
                                       /                            /                         /
[Level 3]                            [9]                           2                         2




The insertion operation will be bounded by the number of comparisons during the walk of the new node, in the order of O(log_n k).