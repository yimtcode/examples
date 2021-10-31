# 演示

> 常用设计模式、算法、语言系统库使用演示，不支持PR。

## 1.数据结构

|   标题   |               实现               | 备注 |
| :------: | :------------------------------: | :--: |
|   数组   | [Go](golang/datastructure/array) |      |
| 双端队列 | [Go](golang/datastructure/deque) |      |
|   堆栈   | [Go](golang/datastructure/stack) |      |

## 2.算法

**排序算法**

|     名称     |                          实现                          |
| :----------: | :----------------------------------------------------: |
| 折半插入排序 |  [Go](golang/algorithm/binary-insertion-sort/main.go)  |
| 直接插入排序 | [GO](golang/algorithm/straight-insertion-sort/main.go) |
|   选择排序   |     [Go](golang/algorithm/selection-sort/main.go)      |
|   归并排序   |       [Go](golang/algorithm/merge-sort/main.go)        |

**寻路算法**

|      名称      |                        实现                         | 演示     | 场景                 |
| :------------: | :-------------------------------------------------: | -------- | -------------------- |
|  广度优先搜索  | [Go](golang/algorithm/breadth-first-search/main.go) | 查找节点 | 查找节点、寻路       |
| 狄克斯特拉算法 | [Go](golang/algorithm/dijkstras-algorithm/main.go)  | 查找节点 | 带权重查找节点、寻路 |

**递推算法**

|     名称     |                             实现                             |  备注  |
| :----------: | :----------------------------------------------------------: | :----: |
| 斐波那契数列 | [Go](golang/algorithm/recursive-algorithm/fibonacci_sequence.go) | 顺推法 |

**其他**

|   名称   |                        实现                        |           演示            | 场景                     |
| :------: | :------------------------------------------------: | :-----------------------: | ------------------------ |
| 动态规划 | [Go](golang/algorithm/dynamic-programming/main.go) | 1.相同子串 2.相同子串长度 | 关键字匹配、单词模糊查找 |
| 贪心算法 |  [Go](golang/algorithm/greedy-algorithm/main.go)   |          找零钱           | 找零钱                   |

## 3.设计模式

**创建型模式**

**结构型模式**

|    名称    |                     实现                     |
| :--------: | :------------------------------------------: |
| 适配器模式 | [Go](golang/designpattern/structure/adapter) |

**行为型模式**

|    名称    |                     实现                     |
| :--------: | :------------------------------------------: |
| 观察者模式 | [Go](golang/designpattern/behavior/observer) |

