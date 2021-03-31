#!/usr/bin/env python3

import unittest

import array_strings

class TestArrayStrings(unittest.TestCase):

    def test_Unique(self):
        r = array_strings.e1_1_unique("abc")
        self.assertTrue(r)

    def test_IsPermutation(self):
        r = array_strings.e1_2_is_permutation("abc","bca")
        self.assertTrue(r)

    def test_IsPermutation2(self):
        r = array_strings.e1_2_is_permutation2("abc","bca")
        self.assertTrue(r)

    def test_ReplaceSpace(self):
        r = array_strings.e1_3_replace_string1("Mr. John Smith      ", 13)
        self.assertEquals(r, "Mr.%20John%20Smith")

    def test_ReplaceSpace2(self):
        ba = bytearray(b'Mr John Smith    ')
        array_strings.e1_3_replace_string2(ba, 13)
        result = b'Mr%20John%20Smith'
        self.assertTrue(str(ba == result))

    def test_PermutationAna(self):
        r = array_strings.e1_4_permutation("Tact coa")
        self.assertTrue(r!=None)

    def test_OneAway(self):
        self.assertTrue(array_strings.e1_5_oneaway("TALE", "TALES"))
        self.assertTrue(array_strings.e1_5_oneaway("TALE", "TLE"))
        self.assertFalse(array_strings.e1_5_oneaway("TALE", "TLES"))
        self.assertFalse(array_strings.e1_5_oneaway("PALE", "PLES"))
        self.assertTrue(array_strings.e1_5_oneaway("ALE", "ARE"))
        self.assertFalse(array_strings.e1_5_oneaway("AYLE", "AYLE"))
        self.assertTrue(array_strings.e1_5_oneaway("LOT", "LO"))
        self.assertTrue(array_strings.e1_5_oneaway("PARLE", "PALE"))
        self.assertTrue(array_strings.e1_5_oneaway("A", "P"))
        self.assertTrue(array_strings.e1_5_oneaway("A", ""))
        self.assertTrue(array_strings.e1_5_oneaway("AB", "B"))
        self.assertFalse(array_strings.e1_5_oneaway("AAB", "B"))
        self.assertFalse(array_strings.e1_5_oneaway("", ""))
        self.assertFalse(array_strings.e1_5_oneaway("PALE", "BAKE"))

    def test_StringCompression(self):
        self.assertTrue(array_strings.e1_6_stringcompression("aabcccccaaa") == "a2b1c5a3")
        self.assertTrue(array_strings.e1_6_stringcompression("") == "")
        self.assertTrue(array_strings.e1_6_stringcompression("aaaaaaaaaaaaaa") == "a14")
        self.assertTrue(array_strings.e1_6_stringcompression("abc") == "abc")
        self.assertTrue(array_strings.e1_6_stringcompression("abcaaaaa") == "a1b1c1a5")
        self.assertEquals(array_strings.e1_6_stringcompression("abcaaa"),"abcaaa")


    def test_ZeroColumnRow(self):
        m = 3
        n = 2
        mat = [[0 for j in range(n)] for i in range(m)]
        mat[0][0] = 1
        mat[0][1] = 2
        mat[1][0] = 3
        mat[1][1] = 4
        mat[2][0] = 0
        mat[2][1] = 6
        r = array_strings.e1_9_zerocolumnrows(mat)
        print(r)

        # print('\n'.join([''.join(['{:4}'.format(item) for item in row]) for row in mat]))

    def test_RotationStr(self):
        self.assertTrue(array_strings.e1_9_isrotation1("waterbottle","erbottlewat"))
        self.assertTrue(array_strings.e1_9_isrotation2("waterbottle","erbottlewat"))

if __name__ == '__main__':
    unittest.main()
