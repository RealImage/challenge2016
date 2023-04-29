
const findInclusionInImmediateParent = (immediateParent, includes, geographyTreeI, end) => {


    if (end < 0) {

        return { status: false, message: `Yout Parent doesnot have permission to in include ${includes}` }

    }


    if (!geographyTreeI[includes[end]]) {

        return { status: false, message: `${includes[end]} not found in out geography DB` }



    } else if (geographyTreeI[includes[end]].includes.includes(immediateParent)) {


        return { status: true, message: 'SuccessFully Validated' }

    } else {

        const childs = geographyTreeI[includes[end]].childs


        return findInclusionInImmediateParent(immediateParent, includes, childs, --end)



    }

}



function helperExclusion(allParents, includes, geographyTreeI, end) {


    for (let i = 0; i < allParents.length; i++) {


        const bool = geographyTreeI[includes[end]].excludes.includes(allParents[i])

        if (bool) {

            return { status: true, message: `${includes.join('-')} is excluded in yout parent ${allParents[i]}` }

        }


    }


    return { status: false, message: 'not exluded' }







}


const findExclusionInAllParents = (allParents, includes, geographyTreeI, end) => {

    if (end < 0) {

        return { status: false, message: `Validate successFully` }

    }


    if (!geographyTreeI[includes[end]]) {

        return { status: false, message: `${includes[end]} not found in out geography DB` }



    } else {

        const helperStatus = helperExclusion(allParents, includes, geographyTreeI, end)

        if (helperStatus.status) {

            return { status: true, message: helperStatus.message }

        }

        else {
            const childs = geographyTreeI[includes[end]].childs;

            return findExclusionInAllParents(allParents, includes, childs, --end)



        }

    }
}



const validateGeography = (distributersI, geographyTreeI, includes, excludes) => {


    if (distributersI.length == 1) {

        return { status: true, message: 'Successfully validated' }
    }

    const immediateParent = distributersI[1];

    let includedInImmediatedParent = false;

    for (let i = 0; i < includes.length; i++) {

        const includeStatus = findInclusionInImmediateParent(immediateParent, includes[i].split('-'), geographyTreeI, includes[i].split('-').length - 1)
        if (includeStatus.status) {
            includedInImmediatedParent = true;
            break;
        }
    }

    if (includedInImmediatedParent == false) {

        return { status: false, message: `Your parent dont have permission to include ${includes.join('-')} ` }

    }

    let excludedd = false;

    for (let i = 0; i < includes.length; i++) {

        const exCludeStatus = findExclusionInAllParents(distributersI.slice(1), includes[i].split('-'), geographyTreeI, includes[i].split('-').length - 1)
        if (exCludeStatus.status) {
            excludedd = true;;

            return { status: false, message: exCludeStatus.message }
        }
    }




    return { status: true, message: 'Successfully validated' }



}


function insertIncluesDistributorGeographyTree(distributers, includes, geographyTree) {


    const mainDistributer = distributers[0];

    function insertSingle(includesI, end, geographyTreeI) {

        if (end == 0) {
            return geographyTreeI[includesI[end]].includes.push(mainDistributer);
        }
        const child = geographyTreeI[includesI[end]].childs

        return insertSingle(includesI, --end, child)

    }


    for (let i = 0; i < includes.length; i++) {

        insertSingle(includes[i].split('-'), includes[i].split('-').length - 1, geographyTree)

    }



}


function insertExcludedDistributorGeographyTree(distributers, excludes, geographyTree) {


    const mainDistributer = distributers[0];

    function insertSingle(excludesI, end, geographyTreeI) {

        if (end == 0) {
            return geographyTreeI[excludesI[end]].excludes.push(mainDistributer);
        }

        const child = geographyTreeI[excludesI[end]].childs

        return insertSingle(excludesI, --end, child)

    }


    for (let i = 0; i < excludes.length; i++) {

        insertSingle(excludes[i].split('-'), excludes[i].split('-').length - 1, geographyTree)

    }




}


const ifGeographyExist = (distributersI, geographyTreeI, include) => {


    const immediateParent = distributersI[0];

    let includedInImmediatedParent = false;



    const includeStatus = findInclusionInImmediateParent(immediateParent, include.split('-'), geographyTreeI, include.split('-').length - 1)
    if (includeStatus.status) {
        includedInImmediatedParent = true;

    }


    if (includedInImmediatedParent == false) {

        return { status: false, message: `You dont have permission to access ${include}` }

    }

    let excludedd = false;


    const exCludeStatus = findExclusionInAllParents(distributersI, include.split('-'), geographyTreeI, include.split('-').length - 1)
    if (exCludeStatus.status) {
        excludedd = true;;

        return { status: false, message: exCludeStatus.message }
    }

    return { status: true, message: 'Successfully validated' }

}

module.exports = { findInclusionInImmediateParent, findExclusionInAllParents, insertIncluesDistributorGeographyTree, insertExcludedDistributorGeographyTree, validateGeography, ifGeographyExist }