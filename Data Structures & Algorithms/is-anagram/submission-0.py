class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        # make sure both strings are the same size or return False
        if len(s) != len(t):
            return False
        array = []
        # Fill array with first string characters
        for i in s:
            array.append(i)
        # Loop through second string and if a letter is found in array, pull it out.
        # If the array is empty at the end without trying to pull from the empty array, it
        # should return True.
        for i in t:
            if i in array:
                array.remove(i)
        if not array:
            return True
        if array:
            return False



        