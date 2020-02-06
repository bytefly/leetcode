#include <iostream>
using namespace std;

int NUM;
// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num) {
    if (num == NUM) return 0;
    if (num < NUM) return -1;
    if (num > NUM) return 1;
}

class Solution {
    public:
        int guessNumber(int n) {
            uint32_t start = 1;
            uint32_t end = n;

            while (1) {
                uint32_t num = (start+end)/2;
                int t = guess(num);
                switch (t) {
                    case -1:
                        start = num+1;
                        break;
                    case 1:
                        end = num-1;
                        break;
                    case 0:
                        return num;
                }
            }
        }
};

int main(int argc, char **argv) {
    NUM = atoi(argv[1]);
    Solution s;
    cout << s.guessNumber(atoi(argv[2])) << endl;
}
