#include <iostream>
#include <vector>
#include <map>

void isValid(std::string c);

int main()
{
    std::map<char, char> withClosing = {
        {'{', '}'},
        {'[', ']'},
        {'(', ')'},
    };

    std::string case1 = "{{[[()]]}}";
    std::string case2 = "{[]}";
    std::string case3 = "{[()]";

    isValid(case1);
}

void isValid(std::string c)
{
    std::vector<char> brakets;

    for (char b : c)
    {

        brakets.push_back(b);
    }
}