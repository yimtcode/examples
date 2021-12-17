# 演示

> 常用设计模式、算法、语言系统库使用演示，不支持PR。

## 1.数据结构

|   标题   |                   实现                   | 备注 |
| :------: | :--------------------------------------: | :--: |
|   数组   |     [Go](golang/datastructure/array)     |      |
| 双端队列 |     [Go](golang/datastructure/deque)     |      |
|   堆栈   |     [Go](golang/datastructure/stack)     |      |
|  AVL树   | [Go](golang/datastructure/tree/avl-tree) |      |

## 2.算法

**排序算法**

|     名称     |                          实现                          |                             参考                             |
| :----------: | :----------------------------------------------------: | :----------------------------------------------------------: |
| 折半插入排序 |  [Go](golang/algorithm/binary-insertion-sort/main.go)  |                                                              |
| 直接插入排序 | [GO](golang/algorithm/straight-insertion-sort/main.go) |                                                              |
|   选择排序   |     [Go](golang/algorithm/selection-sort/main.go)      |                                                              |
|   归并排序   |       [Go](golang/algorithm/merge-sort/main.go)        |                                                              |
|   希尔排序   |       [Go](golang/algorithm/shell-sort/main.go)        | [图解排序算法(二)之希尔排序](https://www.cnblogs.com/chengxiao/p/6104371.html)、[希尔排序](https://zh.wikipedia.org/zh-cn/希尔排序#Go) |

**寻路算法**

|      名称      |                        实现                         | 演示     | 场景                 |
| :------------: | :-------------------------------------------------: | -------- | -------------------- |
|  广度优先搜索  | [Go](golang/algorithm/breadth-first-search/main.go) | 查找节点 | 查找节点、寻路       |
| 狄克斯特拉算法 | [Go](golang/algorithm/dijkstras-algorithm/main.go)  | 查找节点 | 带权重查找节点、寻路 |

**递推算法**

|     名称      |                             实现                             |  备注  |
| :-----------: | :----------------------------------------------------------: | :----: |
| 斐波那契数列  | [Go](golang/algorithm/recursive-algorithm/fibonacci_sequence.go) | 顺推法 |
| 10进制转2进制 | [Go](golang/algorithm/recursive-algorithm/decimal_to_bianry.go) | 顺推法 |
|  母年生小牛   |      [Go](golang/algorithm/recursive-algorithm/cow.go)       | 顺推法 |
|   猴子摘桃    | [Go](golang/algorithm/recursive-algorithm/monkey_pick_peach.go) | 逆推法 |
|  该存多少钱   |    [Go](golang/algorithm/recursive-algorithm/deposit.go)     | 逆推法 |

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

## 4.参考

- 《图解算法》

- 《C/C++函数与算法速查手册》

