#! /usr/bin/env python

import re

pattern = re.compile(r'(\d+)\-(\d+) (\w): (.*)')

total = 0

with open('input.txt', 'r') as f:
    for line in f:
        least, most, ch, password = pattern.findall(line.strip())[0]
        if int(least) <= password.count(ch) <= int(most):
            total += 1

print(total)
