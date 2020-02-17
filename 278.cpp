#include <iostream>
using namespace std;

int badVersion = 1;

bool isBadVersion(int version) {
    if (version >= badVersion) {
        return true;
    }
    return false;
}

class Solution {
    public:
        int firstBadVersion(int n) {
            int left = 1;
            int right = n;
            while (left <= right) {
                int mid = left+(right-left)/2;
                if (isBadVersion(mid)) {
                    right = mid-1;
                } else {
                    left = mid+1;
                }
            }
            return left;
        }
};

int main(int argc, char **argv) {
    Solution s;
    badVersion = atoi(argv[1]);
    cout<<s.firstBadVersion(5)<<endl;
}
