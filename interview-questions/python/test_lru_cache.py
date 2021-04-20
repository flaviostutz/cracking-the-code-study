#!/usr/bin/env python3

import unittest

import lru_cache

class TestArrayStrings(unittest.TestCase):

    def test_LRU1(self):
        lc = lru_cache.LRUCache(9)
        lc.set(1,1)
        self.assertEqual(1, lc.get(1))
        lc.set(1,2)
        self.assertEqual(2, lc.get(1))
        lc.set(1,3)
        self.assertEqual(3, lc.get(1))
        lc.set(2,4)
        self.assertEqual(4, lc.get(2))
        lc.set(3,5)
        self.assertEqual(5, lc.get(3))

    def test_LRU2(self):
        lc = lru_cache.LRUCache(2)
        lc.set(1,1)
        self.assertEqual(1, lc.get(1))
        lc.set(1,2)
        self.assertEqual(2, lc.get(1))
        lc.set(1,3)
        self.assertEqual(3, lc.get(1))
        lc.set(2,4)
        self.assertEqual(4, lc.get(2))
        lc.set(3,5)
        self.assertEqual(5, lc.get(3))
        self.assertEqual(-1, lc.get(1))
        self.assertEqual(4, lc.get(2))
        lc.set(4,6)
        self.assertEqual(6, lc.get(4))
        self.assertEqual(-1, lc.get(2))

        self.assertEqual(-1, lc.get(1))
        self.assertEqual(-1, lc.get(2))
        self.assertEqual(5, lc.get(3))
        self.assertEqual(6, lc.get(4))
        
        lc.set(4,7)
        self.assertEqual(7, lc.get(4))

    def test_LRU2(self):
        lc2 = lru_cache.LRUCache(7)
        self.assertEqual(-1, lc2.get(2))
        lc2.set(2,6)
        self.assertEqual(-1, lc2.get(1))
        lc2.set(1,5)
        lc2.set(1,2)
        self.assertEqual(2, lc2.get(1))
        self.assertEqual(6, lc2.get(2))

if __name__ == '__main__':
    unittest.main()
