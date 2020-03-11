#include <iostream>

using namespace std;

// Definition for a Node.
class Node {
	public:
		int val;
		vector<Node*> children;

		Node() {}

		Node(int _val) {
			val = _val;
		}

		Node(int _val, vector<Node*> _children) {
			val = _val;
			children = _children;
		}
};

class Solution {
	public:
		vector<int> postorder(Node* root) {
			vector<int> ans;
			stack<Node*> s;

			if (root == NULL) {
				return ans;
			}

			s.push(root);
			while (!s.empty()) {
				Node* one = s.top();
				s.pop();

				for (int j = 0; j < one->children.size(); j++) {
					s.push(one->children[j]);
				}

				ans.push_back(one->val);
			}
			reverse(ans.begin(), ans.end());
			return ans;
		}
};
