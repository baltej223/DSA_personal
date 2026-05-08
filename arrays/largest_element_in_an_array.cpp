#include<iostream>
using namespace std;

int main(){
  int array[] = {1,2,52,5363,63};
  int max = 0;

  for (int i=1; i < 5; i++){
    if (array[i] > array[max]){
      max = i;
    }
  }
  cout << "Max element which is " << array[max] << "\n at index, " << max << endl;
}
