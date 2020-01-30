#include <iostream>
#include <vector>
#include <map>
using namespace std;

class Node {
    public:
        int val;
        Node* next;
        Node* random;

        Node(int _val) {
            val = _val;
            next = NULL;
            random = NULL;
        }

        void Dump() {
            Node *p = this;
            cout << "[";
            while (p != NULL) {
                cout << "[" << p->val << ",";
                if (p->random == NULL) {
                    cout << " null]";
                } else {
                    cout << " " << p->random->val << "]";
                }
                if (p->next != NULL) {
                    cout << ", ";
                }
                p = p->next;
            }
            cout << "]" << endl;
        }
};

class Solution {
    public:
        Node* copyRandomList(Node* head) {
            vector<Node*> v; 
            map<uint64_t, int> m;
            Node *p = head;
            int index = 0;
            while (p != NULL) {
                Node *q = new Node(p->val);
                v.push_back(q);
                m[(uint64_t)p] = index;

                p = p->next;
                index++;
            }

            p = head;
            index = 0;
            while (p != NULL) {
                if (index < v.size() - 1) {
                    v[index]->next = v[index+1];
                }
                if (p->random != NULL) {
                    v[index]->random = v[m[(uint64_t)p->random]];
                }

                p = p->next;
                index++;
            }

            if (v.size() == 0) {
                return NULL;
            }
            return v[0];
        }
};

int main(int argc, char **argv) {
    //case 0
    Node *p1 = new Node(7);
    Node *p2 = new Node(13);
    Node *p3 = new Node(11);
    Node *p4 = new Node(10);
    Node *p5 = new Node(1);
    p1->next = p2;
    p2->next = p3;
    p3->next = p4;
    p4->next = p5;
    p2->random = p1;
    p3->random = p5;
    p4->random = p3;
    p5->random = p1;
/*
    //case 1
    Node *p1 = new Node(1);
    Node *p2 = new Node(2);
    p1->next = p2;
    p1->random = p2;
    p2->random = p2;
    //case 2
    Node *p1 = new Node(3);
    Node *p2 = new Node(3);
    Node *p3 = new Node(3);
    p1->next = p2;
    p2->next = p3;
    p2->random = p1;
*/
    Solution s;
    Node *q = s.copyRandomList(p1);
    q->Dump();
    return 0;
}
