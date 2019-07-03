const DistributorCtrl = require('./controllers/DistributorController');

module.exports = (router)=>{
    router.post('/api/v1/distributor/checkpermission', DistributorCtrl.checkPermission);
}