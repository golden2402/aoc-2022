import sys
import re

line = sys.stdin.readline().rstrip()

# part 1:
print(next(i + 1 for i in range(3, len(line))
      if len(set(line[i - 3: i + 1])) == 4))
# part 2:
print(next(i + 1 for i in range(13, len(line))
      if len(set(line[i - 13: i + 1])) == 14))
