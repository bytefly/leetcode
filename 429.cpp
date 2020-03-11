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
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int>> ans;
        queue<Node*> s;

        if (root == NULL) {
            return ans;
        }

        s.push(root);
        while (!s.empty()) {
			int size = s.size();
			vector<int> sub;
			for (int i = 0; i < size; i++) {
				Node* one = s.front();
				s.pop();
				sub.push_back(one->val);

				for (int i = 0; i < one->children.size(); i++) {
					s.push(one->children[i]);
				}
			}
			ans.push_back(sub);
        }
        return ans;
    }
};
