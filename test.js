const checkFn = (exArray, str) => {
  return exArray.includes(str);
};

const test = (include, exclude, qsn) => {
  const ins = ["IN", "US"];
  const exc = ["KA-IN", "CH-TN-IN"];

  const qsnStr = "SE-TN-IN";
  const qsnArray = ["SE-TN-IN", "TN-IN", "IN"];

  const inEx = (ele) => {
    return exc.includes(ele);
  }
  const inIn = (ele) => {
    return ins.includes(ele);
  }

  console.log(qsnArray.some(inEx))
  console.log(qsnArray.some(inIn))
};


const saveDistributor = ()

////test();