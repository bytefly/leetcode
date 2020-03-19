#include <iostream>
#include <vector>

using namespace std;

// This is the custom function interface.
// You should not implement it, or speculate about its implementation
class CustomFunction {
	public:
		// Returns f(x, y) for any given positive integers x and y.
		// Note that f(x, y) is increasing with respect to both x and y.
		// i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
		int f(int x, int y) {
			return x+y;
		}
};

class Solution {
	public:
		vector<vector<int>> findSolution(CustomFunction& customfunction, int z) {
			int x, y;
			vector<vector<int>> ans;

			x = 1;
			y = 1000;

			while (x <= 1000 && y >= 1) {
				int t = customfunction.f(x, y);
				if (t == z) {
					vector<int> item{x, y};
					ans.push_back(item);
					x++;
					y--;
				} else if (t > z) {
					y--;
				} else {
					x++;
				}
			}

			return ans;
		}
};

int main() {
	Solution s;
	CustomFunction c;
	vector<vector<int>> v = s.findSolution(c, 5);
	cout << "[ ";
	for (int i = 0; i < v.size(); i++) {
		cout<<"[" << v[i][0] << "," << v[i][1] << "] ";
	}
	cout << "]" << endl;
}
