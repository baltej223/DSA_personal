#include<bits/stdc++.h>

using namespace std;

int main(){

    vector<int> v1;

    v1.push_back(1);

    for (auto i:v1){
        cout << i <<endl;
    }
    
    v1.erase(v1.begin()); // .erase always takes iterator.
   
     for (auto i:v1){
        cout << i <<endl;
    }

    v1.pop_back();
    
    bool isEmpty = v1.empty();
    
    return 0;
}