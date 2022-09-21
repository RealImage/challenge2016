from pickle import FALSE, TRUE
import sys


DEFAULT_MSG_TEMPLATE = "\tAuthorization Checking Tool\t"
DEFAULT_GEO_FILE_NAME = "cities.csv"
filename = DEFAULT_GEO_FILE_NAME

lst_geo_find_authority = []


list_geo_included_input_regions = []
list_geo_excluded_input_regions = []
list_geo_search_authority_regions = []
list_geo_all_regions = []
lst_geo_diff = set([])
lst_all_distributor_input_cmd_line = []
dict_distributor_search_frm_cmd_line = {}

def format_design_layout(ch='*', num= 10, msg= ""):
    print(f"********{msg}****************\n")

def input_command_line_options():
    print("\n pyhton3 app.py \n -I <included geographical area/ [Place name must be same with CSV file[ [Please follow README-SOL.md]/each place ends with ':']> \n " \
           " -E <excluded geographical area / [Place name must be same with CSV file [Please follow README-SOL.md]/each place ends with ':']>\n" \
           " -D <distributors name / [Each distributor name must end with ':']> \n" \
           " -F <name of csv file>\n")

def chk_command_line_option_only():
    if len(sys.argv) < 1 and len(sys.argv) > 8:
        print("Invalid number of arguments\n")
        format_design_layout("Correct Format should be \n")
        input_command_line_options
        return False
    if  str(sys.argv[1]).capitalize != "-I" and \
        str(sys.argv[3]).capitalize != "-E" and \
        str(sys.argv[5]).capitalize != "-D" and \
        str(sys.argv[7]).capitalize != "-F" :
        print("Invalid options\n")
        format_design_layout("Correct Format should be \n")
        input_command_line_options
        return False
    return TRUE

def chk_geo_frm_csv_match_input_geo(geo_to_match_input_param = "", choice = "include_region"):
    total_cnt = geo_to_match_input_param.count("-")
    cnt = 0
    flag1 = FALSE
    while cnt < total_cnt:
        if filename.endswith(".csv") or filename.endswith(".CSV") :
            with open(filename) as geo_details:
                # for each_line_frm_csv in geo_details:
                flag2 = FALSE
                while (line := geo_details.readline().rstrip()):
                    geo_to_match = geo_to_match_input_param.split("-")
                    geo_find = str(geo_to_match[cnt]).lower()

                    tokens = list(line.split(","))
                    list_geo_all_regions.append(line)
                    for each_token in tokens:
                        if choice == "include_region" and (tokens[len(tokens) - 1].strip().replace(" ", "").lower().startswith(geo_find.rstrip(",")) \
                            or each_token.strip().replace(" ", "").lower() == geo_find.rstrip(",")):
                            list_geo_included_input_regions.append(line)
                            flag2 = TRUE
                            break
                        elif choice == "exclude_region" and (tokens[len(tokens) - 1].strip().replace(" ", "").lower().startswith(geo_find.rstrip(",")) \
                            or each_token.strip().replace(" ", "").lower().startswith(geo_find.rstrip(","))):
                            list_geo_excluded_input_regions.append(line)
                            flag2 = TRUE
                            break
                        elif choice == "find_authority_region" and (tokens[len(tokens) - 1].strip().replace(" ", "").lower().startswith(geo_find.rstrip(",")) \
                            or each_token.strip().replace(" ", "").lower().startswith(geo_find.rstrip(","))):
                            list_geo_search_authority_regions.append(line)
                            flag2 = TRUE
                            break
                    if flag2 == TRUE:
                        flag1 = TRUE
                        break
        if flag1 == TRUE:
            cnt += 2
            break
        else:
            cnt += 2


def main(args = [], distributor_name = "", geo_find_authority = ""):
    list_geo_included_input_regions.clear()
    list_geo_excluded_input_regions.clear()
    list_geo_search_authority_regions.clear()
    list_geo_all_regions.clear()

    chk_geo_frm_csv_match_input_geo(str(args[2]).replace("\"", "").replace(":", "-"), "include_region")
    chk_geo_frm_csv_match_input_geo(str(args[4]).replace("\"", "").replace(":", "-"), "exclude_region")
    chk_geo_frm_csv_match_input_geo(str(geo_find_authority), "find_authority_region")

    lst_filter = list(set(list_geo_included_input_regions) - set(list_geo_excluded_input_regions))
    if  len(lst_filter) > 0 and set(lst_filter).issubset(set(list_geo_search_authority_regions)):
        print(f"Yes.\t The distributor {distributor_name} has PERMISSION to release film in these input regions [{geo_find_authority}")
    else:
        print(f"No.\t The distributor {distributor_name} has no PERMISSION to release film in these input regions [{geo_find_authority}")

def populate_lst_distributors_frm_cmd_line(args = ""):
    lst_distributors = args.replace("\"", "").split(":")
    last_index = len(lst_distributors) - 1
    index = 0
    while last_index >= 0:
        dict_distributor_search_frm_cmd_line[index] = lst_distributors[last_index]
        index += 1
        last_index -= 1

if __name__ == "__main__":
    format_design_layout(f"Welcome {DEFAULT_MSG_TEMPLATE}")
    input_command_line_options()
    if not input_command_line_options and not chk_command_line_option_only:
       format_design_layout("Bye from", DEFAULT_MSG_TEMPLATE)
       sys.exit(0)
    ch = 'y'
    lst_geo_find_authority = []
    lst_distributor = []
    while(ch == 'y'):
        lst_distributor.append(input(f"Please enter the distributor name now[No Spaces in each name/ Name must end with ':'] \t:"))
        lst_geo_find_authority.append(input(f"Please enter the geographical place where distributor's PERMISSION [No Spaces in each name/ Name must end with ':'] \t:"))
        index = 0
        populate_lst_distributors_frm_cmd_line(str(sys.argv[6]))
        index = 0
        for value in dict_distributor_search_frm_cmd_line.values():
            if str(lst_distributor[0]).replace("\"", "").rstrip(":").lower() == str(value).rstrip(":").lower():
                main( sys.argv,  str(value).rstrip(":"), str(lst_geo_find_authority[0]).\
                     replace("\"", "").replace(":", "-"))
                index += 1
                break
        if index == 0:
            format_design_layout("Sorry ",DEFAULT_MSG_TEMPLATE," Invalid Inputs")
        lst_geo_find_authority.clear()
        lst_distributor.clear()
        dict_distributor_search_frm_cmd_line.clear()
        ch = input("Press y to continue entering geo locations.Any key to exit \t:")

