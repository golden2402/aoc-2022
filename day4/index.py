import sys
import re

def contained(x1, x2, y1, y2):
    return x1 >= y1 and x2 <= y2 or x1 <= y1 and x2 >= y2

def intersect(x1, x2, y1, y2):
    return x2 >= y1 and x1 <= y2

groups = [[int(field) for field in re.split("-|,", line.rstrip('\n'))] for line in sys.stdin]

print(sum([1 for group in groups if contained(*group)]))
print(sum([1 for group in groups if intersect(*group)]))