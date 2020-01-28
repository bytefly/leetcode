#include <iostream>

class Node {
public:
    int val;
    Node *prev;
    Node *next;
    Node *child;
};

class Solution {
public:
    Node *flatten(Node *head) {
        Node *p = head;
        while (p != NULL) {
            if (p->child != NULL) {
                Node *m = p->next;
                p->next = p->child;
                p->child->prev = p;

                Node *n = flatten(p->child);
                while (n != NULL && n->next != NULL) {
                    n = n->next;
                }
                if (n != NULL) {
                    n->next = m;
                }
                if (m != NULL) {
                    m->prev = n;
                }
                p->child = NULL;
                p=m;
            } else {
                p = p->next;
            }
        }
        return head;
    }
};

int main(int argc, char **argv) {
    int vals[3][6] = {1, 2, 3, 4, 5, 6,
                    -1, -1, 7, 8, 9, 10,
                    -1, -1, -1, 11, 12, -1};
    Node *nodes[3][6]; 
    for (int i=0; i< 3;i++) {
        for (int j=0; j<6;j++) {
            if (vals[i][j] != -1) {
                Node *p = new Node;
                p->val = vals[i][j];
                p->prev = p->next = p->child = NULL;
                nodes[i][j] = p;
                if (j > 0 && vals[i][j-1] != -1) {
                    p->prev = nodes[i][j-1];
                    nodes[i][j-1]->next = p;
                } 
            }
        }
    }
    nodes[0][2]->child = nodes[1][2];
    nodes[1][3]->child = nodes[2][3];

    Solution s;
    Node *ans = s.flatten(nodes[0][0]);
    for (Node *p = ans; p != NULL; p = p->next) {
        std::cout<<p->val<<" ";
    }
    std::cout<<std::endl;
    return 0;
}
