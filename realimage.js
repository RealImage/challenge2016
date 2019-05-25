var fs = require('fs');
fs.writeFileSync('output.txt','');
fs.readFile('./cities.csv',function (citiesErr, citiesBuf) {
    if(citiesErr){
        console.error(citiesErr);
    } else{
        const cites=citiesBuf.toString();
        fs.readFile('./input.txt',function (inpErr, inpBuf) {
            if(inpErr){
                console.error(citiesErr);
            } else{
                const inputStr = inpBuf.toString();
                let distributor = inputStr.split('Permissions for ');
                let sampleval={};
                for(let i=1;i<distributor.length;i++){
                    let disData=distributor[i].split('\n');
                    //const disname=distributor[i].split('\n')[0].trim();
                    let include=[];
                    let exclude=[];
                    let check=[];
                    let disname='';
                    let parentDis=''; 
                    for(let j in disData){
                        if(j == 0){
                            disname=disData[j].split('<');
                            if(disname.length>1){
                               parentDis = disname[1].trim();
                               disname = disname[0].trim();
                            } else{
                                disname = disname[0].trim();
                            }
                        }
                        if(disData[j].includes('INCLUDE: ')){
                            include.push(disData[j].split('INCLUDE: ')[1].trim());
                        }
                        if(disData[j].includes('EXCLUDE: ')){
                            exclude.push(disData[j].split('EXCLUDE: ')[1].trim());
                        } 
                        if(disData[j].includes('Check: ')){
                            check=disData[j].split('Check: ')[1].split(',');
                        }
                    }      
                    for(j in check){
                        check[j]=check[j].trim();
                    }
                    let checkFlag=true;
                    if(parentDis!=''){
                        if(sampleval[parentDis]!=undefined){
                            const parInclude=sampleval[parentDis].include;
                            const parExclude=sampleval[parentDis].exclude;
                            for(j in include){
                                if(parInclude.includes(include[j])){
                                    checkFlag=true;
                                } else if(parExclude.includes(include[j])){
                                    checkFlag=false;
                                    break;
                                } else {
                                    var tempCheck=[];
                                    if(cites.split(','+include[j]+',').length>1){
                                        tempCheck=cites.split(','+include[j]+',')[1].split(',');
                                    } else{
                                        tempCheck=cites.split(include[j]+',')[1].split(',');
                                    }
                                    if(!(parInclude.includes(tempCheck[0])||parInclude.includes(tempCheck[1]))){
                                        checkFlag=false;
                                        break;
                                    }  else if(parExclude.includes(tempCheck[0])||parExclude.includes(tempCheck[1])){
                                        checkFlag=false;
                                        break;
                                    } else if(parInclude.includes(tempCheck[0])||parInclude.includes(tempCheck[1])){
                                        checkFlag=true;
                                    } else{
                                        checkFlag=false;
                                        break;
                                    }
                                }
                            }
                        }else{
                            fs.appendFileSync('output.txt', 'No corrrect record for '+parentDis+"\n");
                            checkFlag=false;
                        }
                    }    
                    if(checkFlag){
                        if(parentDis){
                            const parExclude=sampleval[parentDis].exclude;
                            exclude=exclude.concat(parExclude);
                        }
                        sampleval[disname]={'include':include,'exclude':exclude};
                        fs.appendFileSync('output.txt', disname+' permission'+"\n");
                        for(j in check){
                            if(exclude.includes(check[j])){
                                fs.appendFileSync('output.txt', "No "+check[j]+"\n");
                            } else if(include.includes(check[j])){
                                fs.appendFileSync('output.txt', "Yes "+check[j]+"\n");
                            } else{
                                let citiesVal=[];
                                if(cites.split(','+check[j]+',').length>1){
                                    citiesVal=cites.split(','+check[j]+',')[1].split('\n')[0].split(',');
                                } else{
                                    citiesVal=cites.split(check[j]+',')[1].split('\n')[0].split(',');
                                }
                                const couCode = citiesVal[1];
                                const stateCode = citiesVal[0];
                                if(exclude.includes(stateCode)||exclude.includes(couCode)){
                                    fs.appendFileSync('output.txt', "No "+check[j]+"\n");
                                } else if(include.includes(stateCode)|| include.includes(couCode)){
                                    fs.appendFileSync('output.txt', "Yes "+check[j]+"\n");
                                } else {
                                    fs.appendFileSync('output.txt', "No "+check[j]+"\n");
                                }
                            }
                        }
                    } else{
                        fs.appendFileSync('output.txt', parentDis + ' cannot authorize '+ disname +' with a region that they themselves do not have access to.'+"\n");
                    }
            }
        }
        })
    }
})