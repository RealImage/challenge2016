import re

DISTRIBUTOR_PATTERN = re.compile("DISTRIBUTOR(\d+)", flags=re.IGNORECASE)
MATCH_INCLUDE = re.compile("^INCLUDE")
MATCH_EXCLUDE = re.compile("^EXCLUDE")
MATCH_LOCATION_CHAIN = re.compile("(?<=[-:\s])\w+")

MATCH_LOCATION_CHAIN = re.compile("(?<=[-:->])\w+")
def match_distributor_identifiers_in_string(string):
    pattern = DISTRIBUTOR_PATTERN
    return re.findall(pattern, string)

def parse_permission_statement(statement):
    locations = re.findall(MATCH_LOCATION_CHAIN, statement)
    locations = [location for location in reversed(locations)]
    if re.match(MATCH_INCLUDE, statement):
        return (1, locations)
    else:
        return (-1, locations)


def parse_query_statement(statement):
    locations = re.findall(MATCH_LOCATION_CHAIN, statement)
    locations = [location for location in reversed(locations)]
    distributor_id = re.findall(DISTRIBUTOR_PATTERN, statement)[0]
    return (distributor_id, locations)
