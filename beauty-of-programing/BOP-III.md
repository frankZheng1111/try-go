# 编程之美第三章节思路小结

## 索引
- [3.1 字符串移位包含问题](#31-字符串移位包含问题)
- [3.4 从无头单链表中删除节点](#34-从无头单链表中删除节点)
- [3.6 编程判断两个链表是否相交](#36-编程判断两个链表是否相交)
l [3.7 队列中取最大值操作问题](#37-队列中取最大值操作问题)
l [3.8 求二叉树节点的最大距离](#38-求二叉树节点的最大距离)

## 3.1 字符串移位包含问题
给定字符串s1和s2，要求判定s2是否能够被s1作循环移位(如abcd 左移1位变成bcda)的字符串所包含

- 解法1: 遍历所有循环移位的结果并判定是否包含，时间O(N^2)
- 解法2: 找寻规律，若保留所有循环位移的结果，则循环位移的字符串位abcdabcd，判定字符串长度不超过原长度且是否包含在新字符串内即可

## 3.4 从无头单链表中删除节点
假设有一个没有头指针的单链表，现有一个指针指向链表中的一个节点(不是首尾节点)，请将该节点从链表中删除

- 解法1: 因为没有头指针，所以无法从头节点遍历至要删除的节点的前一个节点，传统的将前一个节点的指针指向删除节点的下一个节点无法实现。可以选择删除需要被删除节点(n1)的下一个节点(n2)，并将n2的值付给n1,达到伪删除n1的效果

## 3.6 编程判断两个链表是否相交
假设两个链表都不带环, 设长度分别为l1, 和l2

- 解法0: 最直观的枚举第一个链表的点是否在第二个链表中，时间复杂度O(l1\*l2)
- 解法1: 空间换时间, 用一个Hash表记录第一个链表中的点，在遍历第二个链表是否出现在hash表, 时间复杂度O(l1+l2) = O(N), 空间复杂度O(l1) = O(N)
- 解法2: 将链表1接在链表2后，判断链表是否有环，入环点即是相交点O(N)
- 解法3: 遍历两个链表，判断两个链表最后一个节点是否是同一个, 时间复杂度O(l1+l2)

## 3.7 队列中取最大值操作问题
该队列需要提供如下方法: 入队列, 出队列，返回队列最大值，要求时间复杂度尽可能低

- 解法0: 传统队列，入队列出队列时间复杂度为O(1), 返回最大值为O(N)
- 解法1: 维护一个最大堆(通过指针指向后一个元素)，入队列出队列时间复杂O(logN), 返回最大值为O(1)
- 解法2: 维护两个最大栈(一共四个栈)来模仿一个队列, 一个栈进入队列，一个栈出队列(出队列操作时将入队栈的元素全部push到出对栈中再出栈),两个最大栈中的最大值较大的值为队列最大值, 时间复杂度都是O(1), 出队列因为每个元素移动次数都是3次，所以平均时间复杂度是O(1)

## 3.8 求二叉树节点的最大距离
即求整棵二叉树中路径距离最长的两个节点

- 解法思路: 相距最远的节点一定是两个叶子节点或者是一个叶子节点和一个根节点。
- 解法1: 对于任意一个节点，以该节点为根，假设这个根有k个孩子节点，假设相距最远的两个节点u和节点v的路径和根节点的关系是:
  - 若路径经过根节点，则u和v属于不同子数且它们都是该子数中到根节点中距离最远
  - 若未经过根节点，它们一定属于根节点的k个子树中最远的两个节点
- 因此问题就转化为子树上的解，可动态规划求解

