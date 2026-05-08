#include<bits/stdc++.h>
#include<iostream>

using namespace std;

int main(){
    unordered_set<int> s;
    
    s.insert(10);
    s.insert(20);
    s.insert(10);
    s.insert(30);

    for (auto x : s){
        cout << x << " ";
        // No garenteed order.
        cout << endl;
    }

    // We have function like .begin() and .end(), these returns iterators overs the set.
    // What do you mean by iterators, iterators here just means, pointer to first and last element.
    // s.begin() : pointer to first element, which is stored in set, not in the way we inserted.
    // Sets using hashmap at the back.

    for (auto i = s.begin();  i != s.end(); i++){
        cout<< *i << endl;
    }
    return 0;
}