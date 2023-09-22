# Distributor Access Management Program

This program manages distributor access to various regions using a JSON file containing city data. Distributors can be granted or denied access to specific regions, and their permissions can be delegated to other distributors.

## Usage

```bash
./qube --help [options]
```

### Options

- `-input` (default: "cities.json"): Specifies the JSON file containing city data.

- `-distributor`: Specifies the distributor's name.

- `-include`: Regions to include (comma-separated).

- `-exclude`: Regions to exclude (comma-separated).

- `-from`: Distributor to delegate permissions from.

- `-to`: Distributor to delegate permissions to.

- `-access`: Display distributor access.

- `-region`: Display distributors for a specific region.

- `-delete`: Delete a distributor and its children.

### Examples

#### 1. Add Distributor Access

Add a distributor with access to specific regions.

```bash
./qube -input cities.json -distributor newtheatre -include IN
```

This command grants the "newtheatre" distributor access to all regions in India.

#### 2. Remove Distributor Access

Remove access to specific regions for a distributor.

```bash
./qube -input cities.json -distributor newtheatre -exclude TN,IN
```

This command removes access to Tamil Nadu (TN) and all regions in India (IN) for the "newtheatre" distributor.

#### 3. Delegate Permissions

Delegate permissions from one distributor to another for specific regions.

```bash
./qube -input cities.json -from newtheatre -to oldtheatre -include JK,IN
```

This command delegates permissions from the "newtheatre" distributor to the "oldtheatre" distributor for Jammu and Kashmir (JK) and all regions in India (IN).

#### 4. Display Distributor Access

Display the regions to which a distributor has access.

```bash
./qube -access -distributor newtheatre
```

This command shows that the "newtheatre" distributor has access to the United States (US) and all regions in India (IN).

#### 5. Display Distributors for a Specific Region

Display the distributors that have access to a specific region.

```bash
./qube -access -region IN
```

This command displays that India (IN) is accessible by the "newtheatre" distributor.

#### 6. Delete Distributor

Delete a distributor and its children.

```bash
./qube -delete newtheatre
```

This command deletes the "newtheatre" distributor and any child distributors it may have.

## Explanation

- The program loads city data from the input JSON file, which contains details about cities, provinces, and countries.

- Distributor access is managed using a data structure called `DistributorData`, which holds information about distributors and their access.

- Distributors can have access to regions specified by their inclusion and exclusion lists.

- Permissions can be delegated from one distributor to another if the source distributor has access to the specified regions.

- The program can display distributor access to regions and show which distributors have access to a specific region.

- Distributors and their children can be deleted from the system.

---