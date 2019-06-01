"""Module which contains the user Interface"""
import sys
from lib import *

_distributors = [] # Varible to sore the distributor details.

# User Interface starts here.
def user_input():
    """To take user inputs"""
    while True:
        choice = input('Enter the choice:\n' \
                '1. To create the Primary Distributor\n' \
                '2. To create the sub-distributor \n'\
                '3. To check the permissions\n'\
                '4. To see all the distributors\n'\
                '5. To exit\n')
        if choice == '1':
            create_primary_distributor()
            print('\n\n')
        elif choice == '2':
            create_sub_distributor()
            print('\n\n')
        elif choice == '3':
            check_the_permission()
            print('\n\n')
        elif choice == '4':
            see_all_the_distributors()
            print('\n\n')
        elif choice == '5':
            print('Exiting....')
            sys.exit()
        else:
            print("Sorry, Wrong choice")
            print('\n\n')

def create_primary_distributor():
    """A function to take user input and creates the Distributor."""
    locations_to_include = input('Please enter the list(comma seperated) locations to include in \'-\' seperated format\n'
                                 'Example: PUNCH-JK-IN,TN-IN,SUTAC-JUN-PE,SORAS-AYA-PE\n'
                                 '++++++++++++++++++++++++++++++++++++++++++++++++++++\n')
    locations_to_exclude = input('Please enter the list(comma seperated) locations to exclude in \'-\' seperated format\n')
    invalid_code = check_code_validation(locations_to_include, locations_to_exclude)
    if invalid_code:
        print('%s code is invalid, no city exists with that data' %invalid_code)
        return
    permissions = {'include' : locations_to_include.split(',') if locations_to_include != '' else [],
                   'exclude': locations_to_exclude.split(',') if locations_to_exclude != '' else []}
    distributor = Distributor(permissions)
    _distributors.append(distributor)
    print('Distributor created successfully with id: %d and the following permissions' %(distributor.id))
    print(distributor.permissions_hash)


def create_sub_distributor():
    """A function to take user input and creates the sub Distributor."""
    parent_distributor_id = input('Enter the parent distributor id:\n')
    try:
        parent_distributor = _distributors[int(parent_distributor_id)]
    except IndexError:
        print('Sorry, No distributor exists with that id.')
        return
    except ValueError:
        print('Sorry, Invalid ID type.')
        return
    locations_to_include = input('Please enter the list(comma seperated) locations to include in \'-\' seperated format\n'
                                 'Example: PUNCH-JK-IN,TN-IN,SUTAC-JUN-PE,SORAS-AYA-PE\n'
                                 '++++++++++++++++++++++++++++++++++++++++++++++++++++\n')
    locations_to_exclude = input('Please enter the list(comma seperated) locations to exclude in \'-\' seperated format\n')
    invalid_code = check_code_validation(locations_to_include, locations_to_exclude)
    if invalid_code:
        print('%s code is invalid, no city exists with that data.' %invalid_code)
        return
    locations_to_include_s = []
    for i in locations_to_include.split(','):
        if not parent_distributor.has_permission(i):
            print('Sorry the distributor dont have permissions to assign this location: \t' + i)
            return
        else:
            locations_to_include_s.append(i)

    permissions = {'include' : locations_to_include_s, 'exclude': locations_to_exclude.split(',') if locations_to_exclude != '' else []}
    distributor = Distributor(permissions, parent_distributor)
    _distributors.append(distributor)
    print('Distributor created successfully with the id: %d and following permissions' %(distributor.id))
    print(distributor.permissions_hash)

def check_the_permission():
    """A function to check the Distributor's permission."""
    distributor_id = input('Please enter the distributor id: \n')
    try:
        distributor = _distributors[int(distributor_id)]
    except IndexError:
        print('Sorry, No distributor exists with that id.')
        return
    except ValueError:
        print('Sorry, Invalid ID type.')
        return
    permission_code = input('Please enter the location to check the permissions with \'-\' seperated:\n'\
                            'Example: PUNCH-JK-IN or TN-IN\n'\
                            '++++++++++++++++++++++++++++++++++++\n')
    invalid_code = check_code_validation(permission_code, '')
    if invalid_code:
        print('%s code is invalid, no city exists with that data' %invalid_code)
        return
    print(distributor.has_permission(permission_code))

def see_all_the_distributors():
    """A function to print all the distributors"""
    for distributor in _distributors:
        print('Distributor ID: %d \n' %(distributor.id))
        print('Distributor permissions: %s \n' %(distributor.permissions_hash))
        print('Distributor parent ID: %s\n' %(distributor.parent.id if distributor.parent else None))
        print('++++++++++++++++++++++++++++++++++++++++++++++++++++')

if __name__ == "__main__":
    while True:
        try:
            print('Loading the Data from CSV.')
            read_data_from_csv()
            user_input()
        except Exception as error:
            print(error)
            print('Error occurred, Try again!')
            print('\n\n')
