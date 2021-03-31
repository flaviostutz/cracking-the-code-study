
def e1_1_unique(str_array):
    m = {}
    for s in str_array:
        if s in m:
            return False
        else:
            m[s] = s
    return True

def e1_2_is_permutation(str1, str2):
    c = 0
    for s1 in str1:
        for s2 in str2:
            if s1 == s2:
                c = c + 1
    return len(str1) == len(str2) and len(str1) == c


def e1_2_is_permutation2(str1, str2):
    n = {}
    for s1 in str1:
        if s1 in str2:
            n[s1] = 1
        else:
            return False
    return len(str1) == len(str2) and len(n) == len(str1)

def e1_3_replace_string1(str, size):
    r = ""
    str = str.strip()
    for s in str:
        if s == " ":
            r = r + "%20"
        else:
            r = r + s
    return r

def e1_3_replace_string2(byte_array, size):
    # print(byte_array)
    ic = size - 1
    oc = len(byte_array) - 1
    for c in range(0,len(byte_array)):
        if c < (len(byte_array) - size):
            byte_array[oc] = byte_array[ic]
            oc = oc - 1
        else:
            if byte_array[ic] == 32:
                #add '%20'
                byte_array[oc] = 48
                byte_array[oc-1] = 50
                byte_array[oc-2] = 37
                oc = oc - 3
            else:
                byte_array[oc] = byte_array[ic]
                oc = oc - 1
        ic = ic - 1


import math

def e1_4_permutation(str):
    return e1_4_permutation2(str.lower(), "")

def e1_4_permutation2(str, prefix):
    n = len(str)
    if n == 0:
        if e1_4_is_palindrome(prefix):
            # print(prefix + " PALINDROME!")
            return prefix
    else:
        for i in range (0, n):
            p = e1_4_permutation2(str[0:i] + str[i+1:n], prefix + str[i])
#             print(p)
            if p != None:
                return p

def e1_4_is_palindrome(str2):
    str = ""
    for i in range(0, len(str2)):
        if str2[i] != " ":
            str = str + str2[i]
        
    for i in range(0, int(math.floor(len(str)/2))):
        if str[i] != str[len(str)-i-1]:
            return False
    return True

def e1_5_oneaway(str1, str2):
    ec = 0
    if len(str1) == len(str2):
        for i in range(0, len(str1)):
            if str1[i] == str2[i]:
                ec = ec + 1
        return (ec == (len(str1)-1))
    
    elif len(str1) > len(str2):
        i2 = 0
        for i in range(0, len(str1)):
            if i2 < len(str2) and str1[i] == str2[i2]:
                ec = ec + 1
                i2 = i2 + 1
        return (ec == (len(str1)-1))

    elif len(str1) < len(str2):
        i2 = 0
        for i in range(0, len(str2)):
            if i < len(str1) and str1[i] == str2[i2]:
                ec = ec + 1
                i2 = i2 + 1
        return (ec == (len(str2)-1))

    return False

def e1_6_stringcompression(str1):
    #aabcccccaaa
    if len(str1) == 0:
        return ""
    r = ""
    prev = ""
    c = 0
    for cur in str1:
        # print(cur)
        if prev == cur:
            # print("prev==cur")
            c = c + 1
        else:
            # print("prev!=cur")
            if prev != "":
                r = r + prev + str(c)
                # print("ADD")
            prev = cur
            c = 1

    comp = r + cur + str(c)

    if len(comp) > len(str1):
        return str1

    return comp
        
def e1_9_zerocolumnrows(mat):
    zeroes = []
    for i in range(len(mat)):
        # print("x=", x)
        for j in range(len(mat[i])):
            if mat[i][j] == 0:
                zeroes.append((i,j))

    for z in zeroes:
        a = z[0]
        b = z[1]
        for i in range(len(mat)):
            mat[i][b] = 0

        for j in range(len(mat[i])):
            mat[a][j] = 0

    return mat


def e1_9_isrotation1(s1, s2):
    # print(">>>>>>>>")
    ss2 = s2 + s2
    s1i = 0
    matches = 0
    for i in range(len(ss2)):
        if s1i > len(s1) - 1:
            break
        # print(ss2[i], s1[s1i])
        if ss2[i] == s1[s1i]:
            # print("match")
            s1i = s1i + 1
            matches = matches + 1
        else:
            # print("nomatch")
            matches = 0
            s1i = 0
    return matches == len(s1)

def e1_9_isrotation2(s1, s2):
    ss2 = s2 + s2
    return ss2.find(s2) != -1