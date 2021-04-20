#!/usr/bin/env python3

import unittest

import recursive_dynamic

class TestRecursiveDynamic(unittest.TestCase):

    def test_stairways(self):
        r = recursive_dynamic.e8_1_stairway(1)
        self.assertEquals(r, 1)

        r = recursive_dynamic.e8_1_stairway(2)
        self.assertEquals(r, 2)

        r = recursive_dynamic.e8_1_stairway(3)
        self.assertEquals(r, 4)

        r = recursive_dynamic.e8_1_stairway(6)
        self.assertEquals(r, 24)

        r = recursive_dynamic.e8_1_stairway(333)
        self.assertEquals(r, 24)


if __name__ == '__main__':
    unittest.main()
