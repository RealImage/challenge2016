"""Module to solve real image 2016 Challenge."""

_cities_data = {} # Variable to store the cities data from the CSV.

def read_data_from_csv():
    """Read the Cities Data from the CSV file.

    Read the cities csv file and store the content in a Python Dictinoary(Hash).

    Returns:
    Dictionary/Hash with content in the following format. For example
    { 'country':{
          'city': { [list of cities]}
          }
    }
    """
    with open('cities.csv', 'r') as file:
        cities = file.read()
    cities_list = cities.split('\n')

    for city in cities_list[1:len(cities_list)-1]:
        city_codes = city.split(',')[0:3]
        city_codes.reverse()
        country = _cities_data.get(city_codes[0], 0)
        if country != 0:
            try:
                country[city_codes[1]].append(city_codes[2])
            except KeyError:
                country[city_codes[1]] = [city_codes[2]]
        else:
            _cities_data[city_codes[0]] = {city_codes[1] : [city_codes[2]]}
    return _cities_data


def create_permissions(permissions, permission_type, city_codes):
    """Creates the include or exclude permissions to the Distributor.

    Creates the permissions for the distributor based on the city codes passed.
    Added to the existing Country/Province exists else Creates a new one.

    Args:
      permissions: Distributor permissions hash reference.
      permissions_type: Type of the permission either 'include or exclude'
      city_codes: List of city codes to be added to the permissions. For example
            [ 'PUNCH-JK-IN', 'TN-IN', 'SUTAC-JUN-PE', 'SORAS-AYA-PE' ]

    Returns:
      Distributor's permissions Hash. Example:
      {'exclude': {},
       'include': {'PE': {'AYA': ['SORAS'], 'JUN': ['SUTAC']}, 'IN': {'JK': ['PUNCH'], 'TN': []}}}
    """
    for city_code in city_codes:
        codes_list = city_code.split('-')
        codes_list.reverse()
        country = permissions[permission_type].get(codes_list[0], 0) if len(codes_list) > 0 else 0
        if country != 0:
            if len(codes_list) > 1:
                city_code = [codes_list[2]] if len(codes_list) > 2 else [] # Setting the city codes.
                try:
                    # Helps in case of including whole city instead of single city.
                    # e.g. While moving from C-P-C to P-C.
                    country[codes_list[1]].clear() if not city_code else country[codes_list[1]].extend(city_code)
                except KeyError:
                    country[codes_list[1]] = city_code
            else:
                pass
        else:
            permissions[permission_type][codes_list[0]] = {codes_list[1]: [codes_list[2]] if len(codes_list) > 2 else []} if len(codes_list) > 1 else {}
    return permissions


def authorize(permissions, city_codes_list, exclude=False):
    """Authorize the permission of distributor.

    Checks the city codes recursively against the distribtor permissions passed.

    Args:
        permissions: Distributor permissions hash reference.
        city_code: Code of the city to check the permission.
        exclde: True/False based on the type of permisision.

    Returns:
        True/False: Part of the permission or not.
    """
    
    count = 1
    for city_code in city_codes_list:  
        if not permissions:
            return True
        elif count < 3:
            try:
                permissions = permissions[city_code]
            except KeyError:
                return False
        else:
            return city_code in permissions
        count += 1
    if (exclude and count < 4 and permissions): #Condition to allow the superset locations in exclude.
        return False
    return True

def check_code_validation(locations_to_include, locations_to_exclude):
    """Checks the code exists in the core data.

    Split the passed codes and checks whether that exists or not.

    Args:
        locations_to_include: Comma(,) seperated list of location.
        locations_to_exclude: Comma(,) seperated list of location.

    Returns:
        True/False: True for non existence, False for existence.

    """
    codes = locations_to_include.split(',')
    codes.extend(locations_to_exclude.split(',')) if locations_to_exclude != '' else 0
    for code in codes:
        code = code.split('-')
        code.reverse()
        if not check_code_exists(code, _cities_data):
            return '-'.join(reversed(code))
    return False

def check_code_exists(codes, cities_data):
    """Support method to check the existence of cities

    Args:
        codes: List (Splitted to code to check the existency).
        cities_data: Data passed recursively to compare with codes.

    Return:
        Code: If the passed code doesn't exist. (or)
        False: If code exists.
    """

    if len(codes) == 0:
        return True
    elif isinstance(cities_data, list):
        return codes[0] in cities_data
    else:
        try:
            city_code = cities_data[codes[0]]
            return check_code_exists(codes[1:], city_code)
        except KeyError:
            return False

class Distributor(object):
    """Class to hold all the information of a Distributor.

    Attributes:
        permissions: List of permissions in the following format.
            [ 'PUNCH-JK-IN', 'TN-IN', 'SUTAC-JUN-PE', 'SORAS-AYA-PE' ]
        parent: Parent distributor object if exists.
    """
    distributor_id = 0
    def __init__(self, permissions, parent=None):
        super(Distributor, self).__init__()
        self.permissions_hash = {'include': {}, 'exclude': {}}
        self.set_permissions(permissions)
        self.id = Distributor.distributor_id
        self.parent = parent
        Distributor.distributor_id += 1

    def set_permissions(self, permissions):
        """Sets the permissions for the Distributor."""
        create_permissions(self.permissions_hash, 'include', permissions['include'])
        create_permissions(self.permissions_hash, 'exclude', permissions['exclude'])
        return self.permissions_hash

    def has_permission(self, city_code):
        """Checks the distributor has the permission for the code passed."""
        if self.parent:
            if not self.parent.has_permission(city_code):
                return False
        city_codes_list = city_code.split('-')
        city_codes_list.reverse()
        included = authorize(self.permissions_hash['include'], city_codes_list)
        if included:
            excluded = authorize(self.permissions_hash['exclude'], city_codes_list, True) if self.permissions_hash['exclude'] != {} else False
            return not excluded
        return included
