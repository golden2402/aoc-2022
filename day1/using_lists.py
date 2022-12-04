import sys

elves = sorted([sum([int(line) for line in group.split('\n')])
               for group in sys.stdin.read().split("\n\n")])
print(elves[-1], sum(elves[-3:]))
