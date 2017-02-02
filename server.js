/*jshint esversion: 6 */
  app =  require('./config/express'),
  config =  require('./config/env');


app.listen(config.port, () => {
    console.log(`API Server started on port ${config.port} (${config.env})`);
});