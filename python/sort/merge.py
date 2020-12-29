#!/usr/bin/python3

def merge(l1, l2, l):
    i = j = 0
    while i + j < len(l):
        if j == len(l2) or (i < len(l1) and l1[i] < l2[j]):
            l[i+j] = l1[i]
            i += 1
        else:
            l[i+j] = l2[j]
            j += 1


def merge_sort(l):
    n = len(l)
    if n < 2:
        return l
    mid = n // 2
    l1 = l[0:mid]
    l2 = l[mid:n]
    merge_sort(l1)
    merge_sort(l2)
    merge(l1, l2, l)


if __name__ == '__main__':
    l = [85, 24, 63, 45, 17, 31, 96, 50]
    merge_sort(l)
    print(l)
