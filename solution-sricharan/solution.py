import csv
import re
import utils

INCLUDE = 1
EXCLUDE = -1


# Each Region includes a code and a list of sub regions.
class Region:

    def __init__(self, code):
        self.code = code
        self.sub_regions = {}

    def add_sub_region(self, child_region):
        self.sub_regions[child_region.code] = child_region

    def search_in_next_lower_level(self, code):
        region = self.sub_regions.get(code)
        return region

    @staticmethod
    def create_root_node():
        return Region(code='ROOT')

    @staticmethod
    def create_region_and_sub_regions(region_codes):
        current_node = root_region
        for region_code in region_codes:
            region = current_node.search_in_next_lower_level(region_code)
            if region is None:
                region = Region(region_code)
                current_node.add_sub_region(region)
            current_node = region


class Distributor:
    def __init__(self, identifier):
        self.identifier = identifier
        self.included_regions = {}
        self.excluded_regions = {}
        self.child_distributors = {}

    def associate_distributor_as_child(self, distributor):
        self.child_distributors[distributor.identifier] = distributor

    def add_to_included_regions(self, region):
        self.included_regions[region.code] = region

    def add_to_excluded_regions(self, region):
        self.excluded_regions[region.code] = region

    def is_region_included(self, location_code):
        region = self.included_regions.get(location_code)
        if region:
            return True
        else:
            return False


    def is_region_excluded(self, location_code):
        region = self.excluded_regions.get(location_code)
        if region:
            return True
        else:
            return False


    @staticmethod
    def create_distributor(identifier):
        distributor = Distributor(int(identifier))
        return distributor

    @staticmethod
    def create_root_node():
        return Distributor(identifier='ROOT')

def get_region_data_and_add_to_tree():
    with open('cities.csv', 'rb') as csvfile:
        reader = csv.reader(csvfile)
        for index, row in enumerate(reader):
            # Not considering the column headers
            if index == 0:
                continue
            Region.create_region_and_sub_regions(reversed(row[:3]))


def create_distributor_hierarchy(identifiers):
    current_node = root_node
    distributors_in_path = []
    for identifier in identifiers:
        distributor = current_node.child_distributors.get(int(identifier))
        if distributor is None:
            distributor = Distributor.create_distributor(identifier)
            current_node.child_distributors[distributor.identifier] = distributor
        distributors_in_path.append(distributor)
        current_node = distributor
    return distributors_in_path


def handle_non_hierarchical_permission_assignment(locations, include_or_exclude, distributor):
    current_node = root_region
    for location in locations:
        region = current_node.search_in_next_lower_level(location)
        current_node = region
    if include_or_exclude == INCLUDE:
        distributor.included_regions[current_node.code] = current_node
    else:
        distributor.excluded_regions[current_node.code] = current_node


def find_region_hierarchy(locations):
    current_node = root_region
    for location in locations:
        region = current_node.sub_regions.get(location)
        current_node = region
    return current_node


def list_nodes_in_path(target_id, node):
    if node.identifier == target_id:
        return (True, [node])
    else:
        for identifier, distributor in node.child_distributors.items():
            status, path_list = list_nodes_in_path(target_id, distributor)
            if status:
                if node.identifier == "ROOT":
                    return (status, path_list)
                else:
                    return (status, path_list + [node])
        return (False, None)


def greedy_lookup(locations, distributors, is_querying=False):
    is_accessible = False
    if not is_querying:
        distributor_slice = distributors[:len(distributors) - 1]
    else:
        distributor_slice = distributors
    for distributor in distributor_slice:
        for location in locations:
            inclusion_status = distributor.is_region_included(location)
            exclusion_status = distributor.is_region_excluded(location)
            if exclusion_status:
                return False
            elif inclusion_status:
                is_accessible = True
    return is_accessible


def hierarchical_querying(distributors, locations):
    is_accessible = greedy_lookup(locations, distributors, is_querying=True)
    if is_accessible:
        return True
    else:
        return False


def handle_hierarchical_permission_assignment(locations, include_or_exclude, distributors):
    is_accessible = greedy_lookup(locations, distributors)
    if is_accessible:
        region = find_region_hierarchy(locations)
        if include_or_exclude == INCLUDE:
            distributors[len(distributors) - 1].add_to_included_regions(region)
        else:
            distributors[len(distributors) - 1].add_to_excluded_regions(region)
        return True
    else:
        return False

def provide_permissions_for_distributor():
    lines = []
    line_index = 0
    while True:
        line = raw_input()
        if not line:
            break
        else:
            lines.append(line)
    is_a_hierarchical_permission_sequence = False
    matched_identifiers = utils.match_distributor_identifiers_in_string(
        lines[0]
    )
    distributors_in_path = create_distributor_hierarchy(reversed(matched_identifiers))
    if len(distributors_in_path) > 1:
        is_a_hierarchical_permission_sequence = True
    for line in lines[1:]:
        include_or_exclude, locations = utils.parse_permission_statement(line.replace(" ", ""))
        if not is_a_hierarchical_permission_sequence:
            handle_non_hierarchical_permission_assignment(
                locations,
                include_or_exclude,
                distributors_in_path[0]
            )
        else:
            status = handle_hierarchical_permission_assignment(
                locations,
                include_or_exclude,
                distributors_in_path
            )
            if not status:
                print "Operation not permitted."
            else:
                print "Operation done."


def non_hierarchical_querying(distributor, locations):
    current_node = root_region
    is_accessible = False
    for index, location in enumerate(locations):
        region = current_node.sub_regions[location]
        inclusion_status = distributor.is_region_included(location)
        exclusion_status = distributor.is_region_excluded(location)
        if exclusion_status:
            return False
        elif inclusion_status:
            is_accessible = True
        current_node = region
    if not is_accessible:
        return False
    else:
        return True


def query_distributor_access():
    statement = raw_input("Enter the query")
    distributor_id, locations = utils.parse_query_statement(statement)
    status, path = list_nodes_in_path(int(distributor_id), root_node)
    if len(path) == 1:
        is_accessible = non_hierarchical_querying(path[0], locations)
    else:
        is_accessible = hierarchical_querying(path, locations)
    if is_accessible:
        print "Yes, allowed to distribute."
    else:
        print "Not allowed to distribute."

def main():
    option_mappings = {1: provide_permissions_for_distributor, 2: query_distributor_access}
    get_region_data_and_add_to_tree()
    while(1):
        print "1. Create and provide permissions of Distributor."
        print "2. Query Distributor Access."
        print "Please provide your choice."
        choice = input()
        option_mappings[choice]()

new_root_region = Region.create_root_node()
new_root_distributor = Distributor.create_root_node()

root_region = new_root_region
root_node = new_root_distributor

if __name__== "__main__":
  main()
