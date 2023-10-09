### 思路
- 以下为一组
  - 取左边界到右边界，以upper为行，i为列从left到right 将上边界下移++upper
  - 取上边界到下边界，以right为列，i为行从upper到down 将右边界左移--right
  - 取右边界到左边界，以down为行，i为列从right到left 将下边界上移--down
  - 取下边界到上边界，以left为列，i为行从down到upper 将左边界右移++left
- 当左边界大于右边界or上边界大于下边界，结束！