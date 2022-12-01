with open("./sample.txt") as f:
    elves = sorted([sum([int(line) for line in group.split('\n')]) for group in f.read().split("\n\n")])
    print(elves[-1], sum(elves[-3:]))