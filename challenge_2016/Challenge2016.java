package com.qube.Challenge;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Scanner;


public class Challenge2016 {

	public static List<String> Distributors = new ArrayList<String>();
	public static Map<String,List<String>> IncludeMap = new HashMap<>();
	public static Map<String,List<String>> ExcludeMap = new HashMap<>();
	
	public static List<String> cities = new ArrayList<String>();
	
	public static final String PermissionConstant = "Permissions";
	public static Scanner sc = new Scanner(System.in);
	public static Boolean isValidInclude = false;
	public static Boolean isValidExclude = false;
	
	public static Boolean isCode = false;
	
	public static List<String> includePerm = new ArrayList<String>();
	public static List<String> excludePerm = new ArrayList<String>();
	
	public static Boolean[] subDistributorValid = new Boolean[2];
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int input = -1;
		
		try {
			retrieveCitiesFromCSV();
			System.out.println("Select either code or name for permission(sample code: US, Name: united States): ");
			while(true) {
				String permissionOption = sc.nextLine();
				
				if(permissionOption.trim().equalsIgnoreCase("code") || permissionOption.trim().equalsIgnoreCase("name")) {
					
					isCode = true;
					break;
				}
				System.out.println("enter either code or name");
				continue;
				
			}
			while(input !=6) {
				System.out.println("Please select your action");
				System.out.println("1. View Distributor\n2. Add Distributor\n3. Update Distributor"
						+ "\n4. Check Distributor permissions\n5. Delete\n6. Exit Process");
				
				input = Integer.parseInt(sc.nextLine());
				switch(input) {
				case 1: 
					viewDistributor();
					break;
				case 2:
					setDistributor();
					break;
				case 3: 
					updateDistributor();
					break;
				case 4:
					checkPermission();
					break;
				case 5:
					deleteDistributor();
					break;
				case 6:
					System.out.println("Existing Process");
					break;
				default:
					System.out.println("Please select the valid action");
				}
			}
		}catch(Exception e) {
			System.out.println("exception caught: "+e);
			
		}
		
		

	}
	
	/*
	 * To Add new and sub distributor
	 */
	public static void setDistributor() throws Exception {
		isValidInclude = false;
		isValidExclude = false;
		System.out.println("Adding Distributor Permission");
		System.out.println("Enter Distributor Name: ");
		String name = sc.nextLine();
		if(name.contains("<")) {
			System.out.println("name: "+name);
			String[] nameSplit = name.split("<");
//			String subDistributor = nameSplit[0].trim();
			String parentDistributor = setParentDistributorForSub(nameSplit);
//			System.out.println("parent fomr subdis set: "+parentDistributor);
			
//			Boolean valid = true;
			if(Distributors.contains(parentDistributor)) {
				if(checkPermission()) {
					if(includePerm.size()==0 && excludePerm.size()==0) {
						System.out.println("Permission format is not correct");
					}else {
						List<String> parentIncPerm = IncludeMap.get(parentDistributor);
						List<String> parentExcPerm = ExcludeMap.get(parentDistributor);
						
						
						List<String> tempIncludePerm = new ArrayList<String>();
						includePerm.stream().forEach(i -> tempIncludePerm.add(i));
						for (String string : includePerm) {
							checkSubDistributorPermission(parentIncPerm,string,0);
							checkSubDistributorPermission(parentExcPerm,string,1);
							if(!subDistributorValid[0] || subDistributorValid[1]) {
								tempIncludePerm.remove(string);
								System.out.println("Permission denied for "+string);
//								valid = false;

							}
						}
						IncludeMap.put(name, tempIncludePerm);
						ExcludeMap.put(name,excludePerm);
						Distributors.add(name);
						System.out.println("Permission added");
					}
				}
				
				
			}else {
				System.out.println("Distributor "+parentDistributor+" not found");
			}
		}else {
			if(!Distributors.contains(name)) {

				if(checkPermission()) {
					if(includePerm.size()==0 && excludePerm.size()==0) {
						System.out.println("Permission format is not correct");
					}else {
						IncludeMap.put(name, includePerm);
						ExcludeMap.put(name,excludePerm);
						Distributors.add(name);
						System.out.println("Permission added");
					}	
				}else {
					System.out.println("Please enter valid Place");
				}					
			}else {
				System.out.println("distributor is already present, Please use update option");
			}
		}
		
	}

	
	/*
	 * To split the parent distributor from Sub distributor
	 */
	public static String setParentDistributorForSub(String[] nameSplit) throws Exception{
		String parentDistributor = "";
		for(int i =1; i<nameSplit.length ; i++) {
			if(i !=(nameSplit.length-1)) {
				parentDistributor +=nameSplit[i].trim()+" < ";
			}else {
				parentDistributor += nameSplit[i].trim();
			}
		}
		return parentDistributor;
	}
	
	/*
	 * To view all the distributor which got added
	 */
	public static void viewDistributor() throws Exception {
		if(Distributors.size()==0) {
			System.out.println("No Distributors added");
		}else {
			for(String distributor : Distributors) {
				System.out.println("Distributor: "+distributor);
				System.out.println("Include: "+IncludeMap.get(distributor));
				System.out.println("Exclude: "+ExcludeMap.get(distributor));
			}
		}
	}
	/*
	 * To Update the existing Distributor
	 */
	public static void updateDistributor() throws Exception{
		System.out.println("Enter Distributor which need to be updated: ");
		String key = sc.nextLine();
		if(Distributors.contains(key)) {
			System.out.println("enter the Permission Include Exclude with format<include/Exclude>: <place> separated by ,");
			List<String> permList = permissionSplit();
			List<String> includeList = IncludeMap.get(key);
			List<String> excludeList = ExcludeMap.get(key);
			List<String> newIncludePerm = new ArrayList<String>();
			List<String> newExcludePerm = new ArrayList<String>();
			for(String perm: permList) {
				if(perm.contains("Include")) {
					newIncludePerm.add(perm);
				}else if(perm.contains("Exclude")) {
					newExcludePerm.add(perm);
				}
			}
			newIncludePerm = includeSplitPermission(newIncludePerm);
			newExcludePerm = excludeSplitPermission(newExcludePerm);
			int newIncludePermSize = newIncludePerm.size() ;
			int newExcludePermSize = newExcludePerm.size();
			if(newIncludePermSize!=0) {
				isValidInclude = validateWithCitiesCSV(newIncludePerm);
			}else if(newIncludePermSize ==0 ) {
				isValidInclude = true;
			}
			if(newExcludePermSize!=0) {
				isValidExclude = validateWithCitiesCSV(newExcludePerm);
			}else if(newExcludePermSize ==0 ) {
				isValidExclude = true;
			}
			if(isValidExclude && isValidInclude) {
				for (String perm : newIncludePerm) {
					if(!includeList.contains(perm)) {
						includeList.add(perm);
					}
					if(excludeList.contains(perm)) {
						excludeList.remove(perm);
					}
				}
				for (String perm : newExcludePerm) {
					if(!excludeList.contains(perm)) {
						excludeList.add(perm);
					}
					if(includeList.contains(perm)) {
						includeList.remove(perm);
					}
				}
				
				List<String> newIncludeList = new ArrayList<>();
				List<String> newExcludeList = new ArrayList<String>();
				includeList.stream().forEach(i -> newIncludeList.add(i));
				excludeList.stream().forEach(i -> newExcludeList.add(i) );

				List<String> subDistributorList = new ArrayList<String>();
				Distributors.stream().filter(i -> i.matches("[[A-Za-z0-9]+ < ]+"+key)).forEach((i) -> {subDistributorList.add(i);});
				
				
				//While the update itself is sub Distributor
				if(key.contains("<")) {
					String[] nameSplit = key.split("<");
					String parentDistributor = setParentDistributorForSub(nameSplit);
					checkSubDisPermWhileUpdate(key,newIncludeList,newExcludeList,parentDistributor);
					
				}else { //when the update is main DisTributor
					IncludeMap.put(key, newIncludeList);
					ExcludeMap.put(key,newExcludeList);
					System.out.println("Permission Updated for "+key);
				}
				
				//If the update have any SubDistributor
				if(subDistributorList.size()!=0) {
					updateSubDistributor(subDistributorList,key);
				}
				
			}else {
				System.out.println("Please enter valid Place to update");
			}
		}else {
			System.out.println("Distributor not found");
		}
	}
	
	/*
	 * To Delete Existing Distributor
	 */
	public static void deleteDistributor() throws Exception{
		System.out.println("Enter Distributor to delete");
		String distributor = sc.nextLine();
		if(Distributors.contains(distributor)) {
			List<String> tempDistributor = new ArrayList<String>();
			Distributors.stream().forEach(i -> tempDistributor.add(i));
			tempDistributor.stream().filter(i -> i.matches("[[A-Za-z0-9]+ < ]+"+distributor)).forEach((i) -> {System.out.println("dis: "+i);Distributors.remove(i);});
			Distributors.remove(distributor);
			System.out.println("removed: "+Distributors.toString());
			
		}else {
			System.out.println("No distributor found");
		}
	}
	
	/*
	 * This will check the permission of subDistributor during update action
	 */
	public static void checkSubDisPermWhileUpdate(String subDistributor, List<String> newIncludeList
			,List<String> newExcludeList,String parentDistributor) throws Exception{
		
		List<String> parentIncPerm = IncludeMap.get(parentDistributor);
		List<String> parentExcPerm = ExcludeMap.get(parentDistributor);
		
		List<String> tempIncludePerm = new ArrayList<String>();
		newIncludeList.stream().forEach(i -> tempIncludePerm.add(i));
		
		for (String string : newIncludeList) {
			checkSubDistributorPermission(parentIncPerm,string,0);
			checkSubDistributorPermission(parentExcPerm,string,1);
			if(!subDistributorValid[0] || subDistributorValid[1]) {
				tempIncludePerm.remove(string);
				System.out.println("Permission denied for "+string);
			}else {
				IncludeMap.put(subDistributor, tempIncludePerm);
				ExcludeMap.put(subDistributor,newExcludeList);
				System.out.println("Permission Updated for "+subDistributor);
			}
		}
	}
	
	/*
	 * Prep fields for checking sub Distributor permission
	 */
	public static void updateSubDistributor(List<String> subDistributorList,String parentDistributor) throws Exception{
		for (String string : subDistributorList) {
			List<String> subInclude = IncludeMap.get(string);
			List<String> subExclude = ExcludeMap.get(string);
			checkSubDisPermWhileUpdate(string,subInclude,subExclude,parentDistributor);
		}
	}
	
	/*
	 * To check if the places in include and exclude is valid
	 */
	public static Boolean checkPermission() throws Exception{
		System.out.println("Enter Include Exclude with format <Include/Exclude>: <place> separated by ,");
		List<String> permList = permissionSplit();
		
		includePerm = includeSplitPermission(permList);
		excludePerm = excludeSplitPermission(permList);
		int includePermSize = includePerm.size() ;
		int excludePermSize = excludePerm.size();
		if(includePermSize!=0) {
			isValidInclude = validateWithCitiesCSV(includePerm);
		}else if(includePermSize ==0 ) {
			isValidInclude = true;
		}
		if(excludePermSize!=0) {
			isValidExclude = validateWithCitiesCSV(excludePerm);
		}else if(excludePermSize ==0 ) {
			isValidExclude = true;
		}
		if(isValidExclude && isValidInclude) {
			return true;
		}
		return false;
	}
	
	/*
	 * To check if the subdistributor permission is the subset of main distributor
	 */
	public static void checkSubDistributorPermission(List<String> parentPermission,String subPermission,int index) throws Exception{
		System.out.println("index: "+index);
		Boolean valid = false;
		String[] subPermissionSplit = {};
		if(subPermission.contains("-")) {
			  subPermissionSplit = subPermission.split("-");
		}
		if(parentPermission.contains(subPermission)) {
			valid = true;
			subDistributorValid[index] = valid;
			return;
		}else {
			for (String string : parentPermission) {
				if(string.contains("-")) {
					String[] parentPermissionSplit = string.split("-");
					if(parentPermissionSplit.length==2 && subPermissionSplit.length==3) {
						if(parentPermissionSplit[0].equals(subPermissionSplit[1]) && parentPermissionSplit[1].equals(subPermissionSplit[2])) {
							valid = true;
							break;
						}
					}
				}else {
					if(subPermissionSplit.length==2) {
						if(string.equals(subPermissionSplit[1])) {
							valid = true;
							break;
						}
					}
				}
			}
		}
		subDistributorValid[index] = valid;
		
	}
	/*
	 * Extracting Include place from the given input
	 */
	public static List<String> includeSplitPermission(List<String> permList) throws Exception {
		List<String> includeList = new  ArrayList<String>();
		
		permList.stream().filter(i -> i.contains("Include: "))
								.map(i -> i.split(": ")[1]).forEach(a -> includeList.add(a));
		return includeList;
	}
	/*
	 * Extracting Exclude place from give input
	 */
	public static List<String> excludeSplitPermission(List<String> permList) throws Exception {
		
		List<String> excludeList = new  ArrayList<String>();
		
		
		permList.stream().filter(i -> i.contains("Exclude: "))
								.map(i -> i.split(": ")[1]).forEach(a -> excludeList.add(a));

		return excludeList;
	}
	/*
	 * splitting Include and Exclude from input
	 */
	public static List<String> permissionSplit() throws Exception{
		String permission = sc.nextLine();
		List<String> permList = new ArrayList<String>();
		Arrays.stream(permission.split(","))
		.forEach(a -> {a.trim();
						permList.add(a);});
		return permList;
	}
	/*
	 * Fetching places to validate from cities.csv
	 */
	public static void retrieveCitiesFromCSV() throws Exception{
		String line = "";
		
		try{
			BufferedReader br = new BufferedReader(new FileReader("./Resource/cities.csv"));
			while ((line = br.readLine()) != null)   
			{  

				cities.add(line.replace(",", ", "));
			}  
		}catch(Exception ex) {
			System.out.println("Error reading csv");
		}
	}
	/*
	 * Starting point to validate the entered place with CSV
	 * splitting permission for further validation
	 */
	public static Boolean validateWithCitiesCSV(List<String> permission) throws Exception{
		Boolean valid = false;
		for (String perm : permission) {
			if(perm.contains("-")) {
				String[] permsplit = perm.split("-");
				if(permsplit.length==3) {
					valid = checkCityBasedOnCodeOrName(permsplit[0], permsplit[1], permsplit[2]);
				}
				else if(permsplit.length==2) {
					valid = checkCityBasedOnCodeOrName("",permsplit[0],permsplit[1]);
				}
			}else {
				valid = checkCityBasedOnCodeOrName("","",perm);
			}
			
		}
		return valid;
		
	}
	/*
	 * Checks if the user entered code or name of the place for further validation
	 * If selected code the end index of code is 2 and name end index is 5 from city String[]
	 */
	public static Boolean checkCityBasedOnCodeOrName(String permCity,String permProvince, String permCountry) throws Exception{
		Boolean valid = false;
		for (String string : cities) {
			String[] city = string.split(", ");
			if(isCode) {
				valid = validateCity(permCity, permProvince, permCountry, 2, city);
				if(valid) 
					break;
			}else {
				valid = validateCity(permCity, permProvince, permCountry, 5, city);
				if(valid)
					break;
			}
		}
		return valid;
	}
	
	/*
	 * Check the input city with csv city and returns boolean
	 */
	public static Boolean validateCity(String permCity,String permProvince, String permCountry
			,int position,String[] city) throws Exception{
		int cityindex = position-2;
		int proviceIndex= position-1;
		int countryIndex = position;
		if(!permCity.equals("") && permProvince != "" && permCountry != "") {
			if( city[countryIndex].equals(permCountry)  &&city[proviceIndex].equals(permProvince)
					&& city[cityindex].equals(permCity)) {
				return true;
			}
		}
		if(!permProvince.equals("") && permCountry != "" && permCity.equals("")) {
			if(city[proviceIndex].equals(permProvince) && city[countryIndex].equals(permCountry)){
				return true;
			}
		}
		if(permCountry != "" && permProvince.equals("") && permCity.equals("")) {
			if(city[countryIndex].equals(permCountry)) {
				return true;
			}
		}
		return false;
	}
}
