#!/usr/bin/python

import random

random_number = random.randrange(1, 10)

if random_number % 2 == 0:
    print("#yes#")
else:
    print("#no#")
