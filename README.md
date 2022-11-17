# sport-programming-cli
CLI Application for easy copying from library

## Installation
Download app from Releases page. Add it to PATH and use anywhere.

## Usage
- ### spc update - update library from  [sport-programming-library](https://github.com/0wol/sport-programming-library)
- ### spc get - get function or file by name
  - #### Flags:
    - ```filter [function, file]```
  - #### Example: 
  ```
  $ spc get lcm -filter function
  template<typename T>
  T lcm(T a, T b) {
      return (a / gcd<T>(a, b)) * b;
  }
  $ spc get lcm -filter file
  template<typename T>
  T gcd(T a, T b) {
      while (a != 0)  {
          b %= a;
          swap(a, b);
      }
      return b;
  }
  
  template<typename T>
  T lcm(T a, T b) {
      return (a / gcd<T>(a, b)) * b;
  }
  ```
  
