#include <iostream>
#include <vector>
#include <map>
#include <queue>
using namespace std;

// Employee info
class Employee {
    public:
        // It's the unique ID of each node.
        // unique id of this employee
        int id;
        // the importance value of this employee
        int importance;
        // the id of direct subordinates
        vector<int> subordinates;

    public:
        Employee(int _id, int _importance, vector<int> _sub) {
            id = _id;
            importance = _importance;
            subordinates = _sub;
        }
};

class Solution {
    public:
        int getImportance(vector<Employee*> employees, int id) {
            map<int, int> m;
            queue<int> q;
            Employee *leader;
            int ans = 0;

            for (int i = 0; i < employees.size(); i++) {
                m[employees[i]->id] = i;
                if (employees[i]->id == id) {
                    q.push(employees[i]->id);
                }
            }
            while (!q.empty()) {
                leader = employees[m[q.front()]];
                ans += leader->importance;
                for (int i = 0; i < leader->subordinates.size(); i++) {
                    q.push(leader->subordinates[i]);
                }
                q.pop();
            }
            return ans;
        }
};

int main(int argc, char **argv) {
    vector<Employee*> v;
    vector<int> sub1;
    sub1.push_back(2);
    sub1.push_back(3);
    v.push_back(new Employee(1, 5, sub1));
    vector<int> sub2;
    v.push_back(new Employee(2, 3, sub2));
    vector<int> sub3;
    v.push_back(new Employee(3, 3, sub3));

    Solution s;
    cout<<s.getImportance(v, 1)<<endl;
}
