const checkFn = (exArray, str) => {
  return exArray.includes(str);
};

const checkDistributorAccess = qsn => {
  const dist1 = {
    name: "d1",
    includedPlaces: ["IN", "US"],
    excludedPlaces: ["KA-IN", "CH-TN-IN"]
  };

  const qsnSplit = qsn.split("-").reverse();
  const qsnArray = [];

  let initS = "";

  for (const i of qsnSplit) {
    if (qsnSplit.indexOf(i) == 0) {
      initS = i;
    } else {
      initS = `${i}-${initS}`;
    }
    qsnArray.push(i);
  }

  const inEx = ele => {
    return dist1.excludedPlaces.includes(ele);
  };
  const inIn = ele => {
    return dist1.includedPlaces.includes(ele);
  };

  if (qsnArray.some(inIn) && !qsnArray.some(inEx)) {
    console.log("The given area is allocated for d1");
  }
  console.log("The given area is not allocated for d1");
};

checkDistributorAccess("TN-IN");
