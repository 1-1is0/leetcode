class Solution:
    def largestGoodInteger(self, num: str) -> str:
        largest = -1
        for i in range(0, len(num)-2):
            n = num[i:i+3]
            if n.count(n[0]) == 3:
                if int(n) >= largest:
                    largest = int(n)
                print()
        return f"{largest:03}" if largest >=0 else ""







        
num = "6777133339"
num = "2300019"
num = "423523233118"
sol = Solution()
print(sol.largestGoodInteger(num))