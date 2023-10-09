### 非递归
1. 准备一个`pre` 空节点 指向`head`
2. 再用一个指针`tmp`，指针后面的两个节点互换
   - `tmp`->`start`->`end`
   - 变为`tmp`->`end`->`start`
3. 指针指向`start`，重复2操作

### 递归
- head.next.next递归调用方法
- head节点连接调用的结果；next连接head
- 直到head==nil 或 head.next==nil
- 返回next