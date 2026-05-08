#include<iostream>
#include<bits/stdc++.h>

using namespace std;
int main(){
  unordered_map<int, list<int>> m;
  m[1]={2, 4};
  m[2]={1, 3, 4};
  m[3]={2, 4};
  m[4]={1, 3, 2};

  int starts_from = 1;
  queue<int> q;
  q.push(starts_from);

  unordered_map<int,int> visited_vertices = {{1, 0}, {2, 0}, {3, 0}, {4, 0}};

  while(!q.empty()){
    int current = q.front();
    cout << current << " " ;

    for (auto i : m[current]){
      q.push(i);
      m[current] = {};
    };

    q.pop();
  }

}
