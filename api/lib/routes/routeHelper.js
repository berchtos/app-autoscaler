var logger = require('../log/logger');
var models = require('../models')();
var HttpStatus = require('http-status-codes');

exports.createOrUpdatePolicy = function(req,schedulerResponse,callback) {
/*  Create policy will only be called in the async waterfall when we do not 
get any error during the schedule creation/update. */
  models.policy_json.findOrCreate({ where:{ app_id: req.params.app_id },
    defaults: { app_id: req.params.app_id,
    policy_json: req.body } })
    .spread(function(result, created) {
      if(created) {
        logger.info('No policy exists, creating policy..',{ 'app id': req.params.app_id });  
        callback(null, { 'statusCode':HttpStatus.CREATED,'response':result });
      }
      else {
        logger.info('Updating the existing policy',{ 'app id': req.params.app_id });
        models.policy_json.update({
          app_id: req.params.app_id,
          policy_json: req.body
        },{ where: { app_id: req.params.app_id } ,returning:true }).then(function(result) {
          callback(null, { 'statusCode':HttpStatus.OK,'response':result[1] });
        }).catch(function(error) {
          logger.error ('Failed to update policy',
             { 'app id': req.params.app_id,'error':error });
          callback(error,{ 'statusCode':HttpStatus.INTERNAL_SERVER_ERROR });
        });
      }
    }).catch(function(error) {
      logger.error ('Failed to create policy', { 'app id': req.params.app_id,'error':error });
      callback(error,{ 'statusCode':HttpStatus.INTERNAL_SERVER_ERROR });
    });
}