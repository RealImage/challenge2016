const validateDistributers = (distributersI, end, treeI) => {



    if (end == 0) {

        return { status: true, message: 'Successfully validated' }


    }


    if (!treeI || !treeI[distributersI[end]]) {

        return { status: false, message: 'Distributer not Found  ->' + distributersI[end] }

    }

    const childs = treeI[distributersI[end]].childs


    return validateDistributers(distributersI, --end, childs)


}



const insertDistributer = (distributers, end, distributiontTree, includes, excludes) => {


    if (end == 0) {
        distributiontTree[distributers[end]] = {
            'includes': includes,
            'excludes': excludes,
            'childs': {
            }
        }

        return;
    }
    const childs = distributiontTree[distributers[end]].childs;

    return insertDistributer(distributers, --end, childs, includes, excludes);



}

module.exports = {
    validateDistributers,
    insertDistributer

}