#include <iostream>

using namespace std;

template <class T>
T sum(T a, T b)
{
    return a + b;
}

template <class T>
void withCallBack(T cb)
{
    cb();
}

int main()
{
    std::string myName = "Oussama Tailba";

    cout << sum(99.99, 199.99) << endl;
    cout << sum(26, 1) << endl;

    withCallBack([myName]()
                 { std::cout << myName << " is Him" << endl; });

    return 0;
}