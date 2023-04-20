<div align=center>

# Linked List

![](./Linked_Cyclic_List.jpg?raw=true)

</div>

A [linked list](https://en.wikipedia.org/wiki/Linked_list) is a data structure, which is a linear chain of data elements whose order is not specified by their physical placement in memory. This structure consists of nodes that point to an element ahead or behind the given node (depending on the type of the linked list).

# Types of Linked List implemented within this repository

* **Singly Linked List** is a linked list in which a node points only to the next element. Some of the main uses of a singly linked list are to build fundamental data structures, such as stacks and queues.

* **Doubly Linked List** is a linked list in which a node points to both an element ahead of and behind the node. Any application that requires forward and backward tracking of information uses a doubly linked list. For example, the undo and redo functions are implemented with these doubly linked lists.

* **Looped Linked Lists** are singly or doubly-linked that chase their own tail:  
A points to B, B points to C, C points to D, and D points to A. 
They are better suited for cyclic data, such as train schedules.  
These lists lack the first and last elements.  
Therefore, it is necessary to introduce the notion of the current position.