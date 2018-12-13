#include <fstream>
#include <iostream>
#include <string>
#include <iterator>
#include <map>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <regex>

using namespace std;
map<string, map<string, map<string, int> > > cities_map_new;
map<string, map<string, map<string, map<string, map<string, int> > > > > dmap_new;


int read_cities(); // Read the Cities.txt file
int read_distributor(); // Read the Distributors.txt file which has all the Access Rules
int region_chk(); // Take Input from User for Access Check

string* dist;
vector < string > distributor_dependency_chk(string);

int main()
{
    read_cities();
    read_distributor();
    region_chk();
}

int read_cities()
{
    int i;
    string line;
    ifstream fin;
    fin.open("cities.txt");

    // Execute a loop until EOF (End of File)
    while (fin)
    {
        // Read a Line from File
        getline(fin, line);
        int n = line.length();
        // declaring character array
        char char_array[n + 1];

        // copying the contents of the
        // string to char array
        strcpy(char_array, line.c_str());
        char delim[] = ",";
        char *token = strtok(char_array, delim);
        char *city[40];
        int j = 0;

        while (token)
        {
            city[j] = token;
            token = strtok(NULL, delim);
            j++;
        }

        cities_map_new[city[2]][city[1]][city[0]] = 1;
    }
    // Close the file
    fin.close();
    
    return 0;
}

int read_distributor()
{
    string s;
    ifstream fin;
    string distributor;
    smatch match;
    string i_or_e_or_p; // Flag for Include and Exclude
    string country, province, city;
    string d[10];


    // by default open mode = ios::in mode
    fin.open("Distributors.txt");

    // Regex for Distributor Rules
    regex b_p("^Permissions for ([^ ]+)$");
    regex c_p("^Permissions for ([^ ]+) < ([^ ]+)$");
    regex d_p("^Permissions for ([^ ]+) < ([^ ]+) < ([^ ]+)$");

    regex b_i("^INCLUDE: ([^ ]+)$");
    regex c_i("^INCLUDE: ([^ ]+)-([^ ]+)$");
    regex d_i("^INCLUDE: ([^ ]+)-([^ ]+)-([^ ]+)$");

    regex b_e("^EXCLUDE: ([^ ]+)$");
    regex c_e("^EXCLUDE: ([^ ]+)-([^ ]+)$");
    regex d_e("^EXCLUDE: ([^ ]+)-([^ ]+)-([^ ]+)$");

    // Execute a loop until EOF (End of File)
    while (fin)
    {

        // Read a Line from File
        getline(fin, s);
        // cout << s << endl;
        if (regex_match(s, b_p))
        {
            if (regex_search(s, match, b_p))
            {

                distributor = match[1];
                i_or_e_or_p = "p";
            }
        }
        else if (regex_match(s, c_p))
        {
            if (regex_search(s, match, c_p))
            {
                distributor = match[1];
                i_or_e_or_p = "p";
            }
        }
        else if (regex_match(s, d_p))
        {
            if (regex_search(s, match, d_p))
            {
                distributor = match[1];
                i_or_e_or_p = "p";
            }
        }
        else if (regex_match(s, d_i))
        {
            if (regex_search(s, match, d_i))
            {
                city = match[1];
                i_or_e_or_p = "i";
                dmap_new[distributor][i_or_e_or_p][match[1]][match[2]][match[3]] = 1;
            }
        }
        else if (regex_match(s, c_i))
        {
            if (regex_search(s, match, c_i))
            {
                province = match[1];
                i_or_e_or_p = "i";
                dmap_new[distributor][i_or_e_or_p]["abc"][match[1]][match[2]] = 1;
            }
        }
        else if (regex_match(s, b_i))
        {
            if (regex_search(s, match, b_i))
            {
                country = match[1];
                i_or_e_or_p = "i";
                dmap_new[distributor][i_or_e_or_p]["abc"]["abc"][match[1]] = 1;
            }
        }
        else if (regex_match(s, d_e))
        {
            if (regex_search(s, match, d_e))
            {
                city = match[1];
                i_or_e_or_p = "e";
                dmap_new[distributor][i_or_e_or_p][match[1]][match[2]][match[3]] = 1;
            }
        }
        else if (regex_match(s, c_e))
        {
            if (regex_search(s, match, c_e))
            {
                province = match[1];
                i_or_e_or_p = "e";
                dmap_new[distributor][i_or_e_or_p]["abc"][match[1]][match[2]] = 1;
            }
        }
        else if (regex_match(s, b_e))
        {
            if (regex_search(s, match, b_e))
            {
                country = match[1];
                i_or_e_or_p = "e";
                dmap_new[distributor][i_or_e_or_p]["abc"]["abc"][match[1]] = 1;
            }
        }
        // Below else condition to handle invalid lines 
        // else
        //     cout << "Invalid line in the Distributor Auth File. " << endl;
    }
    fin.close();
    return 0;
}


// Take input from User for <Distributor Region>
// and Validate the Access
int region_chk()
{
    string s, distributor, region, country, province, city;
    string *p;
    vector<string> dist_vector;
    int flag = 1;

    cout << "Enter the input in the format <Distributor Region>: " << endl;
    getline(cin, s);
    cout << "You entered: " << s << endl;
    smatch match;
    int length;

    // Regex patterns to split Dist and Region Details
    regex b_e("([^ ]+) ([^ ]+)$");
    regex c_e("([^ ]+) ([^ ]+)-([^ ]+)$");
    regex d_e("([^ ]+) ([^ ]+)-([^ ]+)-([^ ]+)$");

    if (regex_match(s, d_e))
    {
        if (regex_search(s, match, d_e)) // Check for <Dist City-Province-Country>
        {
            distributor = match[1];
            city = match[2];
            province = match[3];
            country = match[4];

            if (cities_map_new.count(country) > 0)
            {
                cout << "Country Exists" << endl;
                if (cities_map_new[country].count(province) > 0)
                {
                    cout << "Province Exists" << endl;
                    if (cities_map_new[country][province].count(city) > 0)
                    {
                        cout << "City Exists" << endl;

                        dist_vector = distributor_dependency_chk(distributor); // Check for Distributor Dependency

                        for (int i = 0; i < dist_vector.size(); ++i)
                        {
                            if (flag == 1)
                            {

                                if (dmap_new[distributor]["e"]["abc"]["abc"].count(country) > 0)
                                {
                                    cout << "No Access in 1 - " << i << endl;
                                    flag = 0;
                                }
                                else if (dmap_new[distributor]["e"]["abc"][province].count(country))
                                {
                                    cout << "No Access in 2 - " << i << endl;
                                    flag = 0;
                                }
                                else if (dmap_new[distributor]["e"][city][province].count(country))
                                {
                                    cout << "No Access in 3 - " << i << endl;
                                    flag = 0;
                                }
                                else if (dmap_new[distributor]["i"]["abc"]["abc"].count(country) > 0)
                                {
                                    cout << "Access Granted in 1 - " << i << endl;
                                    flag = 1;
                                }
                                else if (dmap_new[distributor]["i"]["abc"][province].count(country) > 0)
                                {
                                    cout << "Access Granted in 2 - " << i << endl;
                                    flag = 1;
                                }
                                else if (dmap_new[distributor]["i"][city][province].count(country))
                                {
                                    cout << "Access Granted in 3 - " << i << endl;
                                    flag = 1;
                                }
                                else
                                {
                                    cout << "No access " << i << endl;
                                    flag = 0;
                                }
                            }
                        }

                        if (flag == 1)
                        {
                            cout << "YES " << endl;
                        }
                        else
                        {
                            cout << "NO " << endl;
                        }
                    }
                    else
                    {
                        cout << "City does not Exists" << endl;
                    }
                }
                else
                {
                    cout << "Province does not Exists" << endl;
                }
            }
            else
            {
                cout << "Country does not Exist" << endl;
            }
        }
    }
    else if (regex_match(s, c_e)) // Check for <Dist Province-Country>
    {
        if (regex_search(s, match, c_e))
        {
            distributor = match[1];
            province = match[2];
            country = match[3];
            if (cities_map_new.count(country) > 0)
            {
                cout << "Country Exists 2 " << endl;
                if (cities_map_new[country].count(province) > 0)
                {
                    cout << "Province Exists 2 " << endl;

                    dist_vector = distributor_dependency_chk(distributor); // Check for Distributor Dependency

                    for (int i = 0; i < dist_vector.size(); ++i)
                    {
                        if (flag == 1)
                        {

                            if (dmap_new[dist_vector[i]]["e"]["abc"]["abc"].count(country) > 0)
                            {
                                cout << "No Access in 1 - " << i << endl;
                                flag = 0;
                            }
                            else if (dmap_new[dist_vector[i]]["e"]["abc"][province].count(country) > 0)
                            {
                                cout << "No Access in 2 - " << i << endl;
                                flag = 0;
                            }
                            else if (dmap_new[dist_vector[i]]["i"]["abc"]["abc"].count(country) > 0)
                            {
                                cout << "Access Granted in 1 - " << i << endl;
                                flag = 1;
                            }
                            else if (dmap_new[dist_vector[i]]["i"]["abc"][province].count(country) > 0)
                            {
                                cout << "Access Granted in 2 - " << i << endl;
                                flag = 1;
                            }
                            else
                            {
                                cout << "No access " << endl;
                                flag = 0;
                            }
                        }
                    }

                    if (flag == 1)
                    {
                        cout << "YES " << endl;
                    }
                    else
                    {
                        cout << "NO " << endl;
                    }
                }
                else
                {
                    cout << "Province does not Exists" << endl;
                }
            }
            else
            {
                cout << "Country does not Exist" << endl;
            }
        }
    }
    else if (regex_match(s, b_e)) // Check for <Dist Country>
    {
        if (regex_search(s, match, b_e))
        {
            distributor = match[1];
            country = match[2];
            if (cities_map_new.count(country) > 0)
            {
                cout << "Country Exists 3 " << endl;

                dist_vector = distributor_dependency_chk(distributor); // Check for Distributor Dependency

                for (int i = 0; i < dist_vector.size(); i++)
                {
                    if (flag == 1)
                    {
                        if (dmap_new[dist_vector[i]]["e"]["abc"]["abc"].count(country) > 0)
                        {
                            cout << "No Access in 1 - " << i << endl;
                            flag = 0;
                        }
                        else if (dmap_new[dist_vector[i]]["i"]["abc"]["abc"].count(country) > 0)
                        {
                            cout << "Access Granted in 1 - " << i << endl;
                            flag = 1;
                        }
                        else
                        {
                            cout << "No Rule. So No Access - " << i << endl;
                            flag = 0;
                        }
                    }
                }
            }

            else
            {
                cout << "Country does not Exist" << endl;
            }

            if (flag == 1)
            {
                cout << "YES " << endl;
            }
            else
            {
                cout << "NO " << endl;
            }
        }
    }
    return 0;
}

vector<string> distributor_dependency_chk(string dist)
{
    // string* d[3];
    string *d = new string[3];
    vector<string> dist_vector;
    smatch match;
    string s;
    ifstream fin;
    fin.open("Distributors.txt");

    // string* names = new string[3];

    regex b_p("^Permissions for ([^ ]+)$");
    regex c_p("^Permissions for ([^ ]+) < ([^ ]+)$");
    regex d_p("^Permissions for ([^ ]+) < ([^ ]+) < ([^ ]+)$");
    while (fin)
    {

        // Read a Line from File
        getline(fin, s);
        // cout << s << endl;
        if (regex_match(s, b_p))
        {

            if (regex_search(s, match, b_p))
            {
                if (dist == match[1])
                {
                    d[0] = match[1];
                    dist_vector.push_back(match[1]);
                    cout << "Main Dist - " << s << endl;

                    return dist_vector;
                }
            }
        }
        else if (regex_match(s, c_p))
        {
            if (regex_search(s, match, c_p))
            {
                if (dist == match[1])
                {
                    d[0] = match[1];
                    d[1] = match[2];
                    dist_vector.push_back(match[2]);
                    dist_vector.push_back(match[1]);

                    cout << "Sub Dist - " << s << endl;

                    return dist_vector;
                }
            }
        }
        else if (regex_match(s, d_p))
        {
            if (regex_search(s, match, d_p))
            {
                if (dist == match[1])
                {
                    d[0] = match[1];
                    d[1] = match[2];
                    d[2] = match[3];

                    dist_vector.push_back(match[3]);
                    dist_vector.push_back(match[2]);
                    dist_vector.push_back(match[1]);

                    cout << "Sub Sub Dist - " << s << endl;

                    return dist_vector;
                }
            }
        }
        // else
        //     cout << "Invalid Distributor. " << endl;
    }
    fin.close();
    return dist_vector;
}
