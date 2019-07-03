const HttpStatus = require('http-status');
const _ = require('underscore');

 const distributorPermissions = [
        {
            distributor:"Distributor1",
            includes:[{country: "INDIA"},{country: "UNITEDSTATES"},{country: "PASKISTAN"}],
            excludes:[{country:"INDIA", state:"KARNATAKA"},{country:"INDIA", state:"TAMILNADU", city:"CHENNAI"}],
            under: null
        },
        {
            distributor:"Distributor2",
            includes:[{country: "INDIA"}],
            excludes:[{state:"TAMILNADU", country:"INDIA"}],
            under: "Distributor1"
        },
        {
            distributor:"Distributor3",
            includes:[{country:"INDIA", state:"KARNATAKA", city:"HUBLI"}],
            excludes:[],
            under: "Distributor2"
        }
    ]

exports.checkPermission = (req, res)=>{
    var reqBody = req.body;

    if(reqBody.distributor){
        const distributor = _.find(distributorPermissions, (x)=>{
            return x.distributor == reqBody.distributor;
        })

        if(distributor){
            const includes = distributor.includes;

            const checkExcludes = ()=>{
                return new Promise((resolve, reject)=>{
                    const loop = (distributor)=>{
                        let excludes = distributor.excludes;
                        this.isExcluded(excludes, reqBody).then(()=>{
                            if(distributor.under){
                                loop(this.getDistributor(distributor.under));
                            }else{
                                console.log("Inside ")
                                resolve();
                            }
                        }).catch((e)=>{
                            console.log(e)
                            reject();
                        })
                    }
                    loop(distributor);
                })
            }

            const checkIncludes = ()=>{
                return new Promise((resolve, reject)=>{
                    if(includes.length>0){
                        const data =  _.find(includes, (x)=>{
                            if((!x.city) && (!x.state) && x.country == reqBody.country){
                               return true
                            }else if(!x.city && reqBody.state == x.state && x.country == reqBody.country){
                                return true
                            }else if(x.city == reqBody.city && reqBody.state == x.state && reqBody.country == x.country){
                                return true
                            }
                        });
                        if(data){
                            resolve()
                        }else{
                            reject();
                        }
                    }else{
                        reject();
                    }
                })
            }
            
            Promise.all([checkExcludes(), checkIncludes()]).then((val)=>{
                res.send("Yes");
            }).catch((e)=>{
                console.log(e)
                res.send("No");
            })
        }else{
            res.send("invalid Distributor");
        }
    }else{
        res.send("Distributor required");
    }
}

exports.getDistributor = (distributor)=>{
    const dist =  _.find(distributorPermissions, (x)=>{
        return x.distributor == distributor;
    });

    return dist;
}

exports.isExcluded = (excludes, reqBody)=>{
    return new Promise((resolve, reject)=>{
        if(excludes.length>0){
            const data =_.find(excludes, (x)=>{
                if((!x.city) && (!x.state) && x.country == reqBody.country){
                    return true
                }else if(!x.city && reqBody.state == x.state && x.country == reqBody.country){
                    return true
                }else if(x.city == reqBody.city && reqBody.state == x.state && reqBody.country == x.country){
                    return true
                }
            });
            if(data){
                reject();
            }else{
                resolve();
            }
        }else{
            resolve();
        }
    })
}