import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class QubeCinemaApp {
/*
Permissions for DISTRIBUTOR1
INCLUDE: INDIA
INCLUDE: UNITEDSTATES
EXCLUDE: KARNATAKA-INDIA
EXCLUDE: CHENNAI-TAMILNADU-INDIA

Permissions for DISTRIBUTOR2 < DISTRIBUTOR1
INCLUDE: INDIA
EXCLUDE: TAMILNADU-INDIA
*/
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        DistributorPermissionService dps = new DistributorPermissionService();
        while(true) {
            System.out.println("Enter 1 for adding permission");
            System.out.println("Enter 2 for checking permission");
            System.out.println("Enter 3 for quit");
            int input = sc.nextInt();
            sc.nextLine();
            switch (input) {
                case 1:
                    System.out.println("Enter your input in correct format");
                    sc.next();sc.next();
                    String permissionFor = sc.nextLine();
                    while(true) {
                        try {
                            System.out.println("Enter permission(INCLUDE/EXCLUDE): REGION or press 4 for Main menu");
                            String permission = sc.next();
                            if (permission.trim().equalsIgnoreCase("4")) {
                                break;
                            }
                            String region = sc.nextLine();
                            boolean include = permission.startsWith("INCLUDE");
                            boolean exclude = permission.startsWith("EXCLUDE");
                            if (!(include || exclude)) {
                                System.out.println("invalid input, please try again !");
                                continue;
                            }
                            String[] dists = permissionFor.split("<");
                            String parentDist = null, childDist = null;
                            if (dists.length >= 1) {
                                childDist = dists[0];
                            }

                            if (dists.length > 1) {
                                parentDist = dists[1].trim();
                            }
                            dps.includePermission(parentDist, childDist.trim(), region.trim(), include);
                        } catch (RuntimeException re) {
                            System.out.println("Error: " + re.getMessage());
                        }
                    }
                    break;

                case 2:
                    while(true) {
                        try {
                            System.out.println("Enter Distributor & Region(i.e. DistributorName1 BANGLORE-KARNATAKA-INDIA) or press 4 for Main menu");
                            String dist = sc.next().trim();
                            if (dist.trim().equalsIgnoreCase("4")) {
                                break;
                            }
                            String region = sc.nextLine().trim();
                            boolean result = dps.hasPermission2Distribute(dist, region);
                            System.out.println("Can `" + dist + "` distribute to `" + region + "` : " + (result ? "Yes" : "No"));
                        } catch (RuntimeException re) {
                            System.out.println("Error : " + re.getMessage());
                        }
                    }
                    break;

                case 3:
                    System.exit(0);
                default:
                    System.out.println("Invalid Input, Please enter 1/2/3 Options !");
            }
        }
    }
}

class DistributorPermissionService {
    Map<String, Node> countriesMap = new HashMap<>();
    Map<String, String> distributorRelationshipMap = new HashMap<>();

    public void includePermission(String parentDist, String childDist, String region, boolean include) {
        String[] regions = region.split("-");
        String country = null, province = null, city = null;
        if(regions.length > 3 && regions.length < 1) {
            throw new RuntimeException("Invalid input, Atleast one region needs to be present");
        }

        if(parentDist != null && !distributorRelationshipMap.containsKey(parentDist)) {
            throw new RuntimeException("No distributor was found with name `"+parentDist+"`");
        }

        if(regions.length == 3) {
            country = regions[2];
            province = regions[1];
            city = regions[0];
        }

        if(regions.length == 2) {
            province = regions[0];
            country = regions[1];
        }
        if(regions.length == 1) {
            country = regions[0];
        }
        distributorRelationshipMap.putIfAbsent(childDist, parentDist);

        Node countryRegion = null, provinceRegion = null, cityRegion = null;
        countryRegion = countriesMap.getOrDefault(country, new Node(Type.COUNTRY, country, null));
        countriesMap.putIfAbsent(country, countryRegion);

        if(province == null) {
            // add inc/exc here for country
            if(countryRegion.hasPermission(parentDist)) {
                addPermission(countryRegion, include, childDist);
                return;
            }
            System.out.println(parentDist + " doesn't have permission for ->" + country);
            return;
        } else {
            provinceRegion = countryRegion.childMap.getOrDefault(province, new Node(Type.PROVINCE, province, countryRegion));
            countryRegion.addChildIfAbsent(provinceRegion);
        }

        if(city == null) {
            // add inc/exc here for province
            if(provinceRegion.hasPermission(parentDist)) {
                addPermission(provinceRegion, include, childDist);
                return;
            }
            System.out.println(parentDist + " doesn't have permission for ->" + province + "-" + country);
            return;
        } else {
            // add inc/exc here for city
            cityRegion = provinceRegion.childMap.getOrDefault(city, new Node(Type.CITY, city, provinceRegion));
            provinceRegion.addChildIfAbsent(cityRegion);
            if(cityRegion.hasPermission(parentDist)) {
                addPermission(cityRegion, include, childDist);
                return;
            }
            System.out.println(parentDist + " doesn't have permission for ->" + city + "-" + province + "-" + country);
        }
    }

    public void excludePermission() {
    }

    public boolean hasPermission2Distribute(String dist, String region) {
        String[] regions = region.split("-");
        String country = null, province = null, city = null;
        if(regions.length > 3 && regions.length < 1) {
            throw new RuntimeException("Invalid input, Atleast one region needs to be present");
        }
        if(!distributorRelationshipMap.containsKey(dist)) {
            throw new RuntimeException("No distributor was found with name `"+dist+"`");
        }


        if(regions.length == 3) {
            country = regions[2];
            province = regions[1];
            city = regions[0];
        }

        if(regions.length == 2) {
            province = regions[0];
            country = regions[1];
        }
        if(regions.length == 1) {
            country = regions[0];
        }
        boolean result = true;
        do {
            result = result && hasPermission2Distribute(dist, country, province, city);
            dist = distributorRelationshipMap.get(dist);
        } while(dist != null && result);
        return result;
    }

    private boolean hasPermission2Distribute(String dist, String country, String province, String city) {
        Node countryRegion = countriesMap.get(country);
        if(countryRegion == null) {
            return false;
        }

        if(countryRegion.isPermissionExcluded(dist)) {
            return false;
        }

        boolean hasCountryIncluded = countryRegion.hasPermission(dist);
        if(province == null) {
            return hasCountryIncluded;
        }

        Node provinceRegion = countryRegion.childMap.get(province);
        if(provinceRegion == null) {
            return hasCountryIncluded;
        }
        if(provinceRegion.isPermissionExcluded(dist)) {
            return false;
        }

        boolean hasProvinceIncluded = provinceRegion.hasPermission(dist);
        if(city == null) {
            return (hasCountryIncluded || hasProvinceIncluded);
        }

        Node cityRegion = provinceRegion.childMap.get(city);
        if(cityRegion == null) {
            return (hasCountryIncluded || hasProvinceIncluded);
        }

        if(cityRegion.isPermissionExcluded(dist)) {
            return false;
        }
        return true;
    }

    public void addPermission(Node node, boolean include, String dist) {
        if(include) {
            node.includeRegion4Dist(dist);
        } else {
            node.excludeRegion4Dist(dist);
        }
    }
}

class Node {
    Type type;
    Node parent;
    String name;
    Map<String, Node> childMap = new HashMap<>();
    Map<String, String> includeACLs = new HashMap<>();
    Map<String, String> excludeACLs = new HashMap<>();

    public Node(Type type, String name, Node parent) {
        this.type = type;
        this.name = name;
        this.parent = parent;
    }

    public void includeRegion4Dist(String dist) {
        excludeACLs.remove(dist);
        includeACLs.put(dist, dist);
    }

    public void excludeRegion4Dist(String dist) {
        includeACLs.remove(dist);
        excludeACLs.put(dist, dist);
    }

    public void addChildIfAbsent(Node child) {
        childMap.putIfAbsent(child.name, child);
    }

    public boolean hasPermission(String dist) {
        if(dist == null || isPermissionIncluded(dist)) {
            return true;
        }
        if(isPermissionExcluded(dist)) {
            return false;
        }
        return isPermissionInherited(dist);
    }

    private boolean isPermissionIncluded(String dist) {
        return includeACLs.containsKey(dist);
    }

    private boolean isPermissionInherited(String dist) {
        //todo needs to work on this tomorrow morning....
        Node parent = this.parent;

        while(parent != null) {
            if(parent.isPermissionIncluded(dist)) {
                return true;
            }
            if(parent.isPermissionExcluded(dist)) {
                return false;
            }
            parent = parent.parent;
        }
        return false;
    }

    public boolean isPermissionExcluded(String dist) {
        return excludeACLs.containsKey(dist);
    }
}

enum Type {
    COUNTRY, PROVINCE, CITY
}
