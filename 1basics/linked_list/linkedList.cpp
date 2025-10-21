#include<iostream>

using  namespace std;

class Node{
    public:
    int data;
    Node *next=NULL;
};

class Arr{
    public:
    int *arr;
    int size;
};

class LinkedList{
    public:
    Node *head = NULL;
    void AddAtStarting(int value){
        if (head == NULL){
            // It means that the linked list is empty.
            Node *first = new Node();
            first->data = value;
            // return first;
            this->head = first;
        } else {
            // Means there are some elements already in LL
            Node *aNode=  new Node();
            aNode->data = value;
            this->head->next = aNode;
        }
    }

    void PrintFirstElementOfLL(){
        cout << head->data << endl;
    }

    void ConstructLLfromArray(){
        
    }
};

int main(){
    LinkedList newlist;
    newlist.AddAtStarting(21);
    newlist.PrintFirstElementOfLL();
    newlist.AddAtStarting(25);
    newlist.PrintFirstElementOfLL();
}