'use strict';

const https = require('https');

module.exports.handler = (event, context, callback) => {
  https.get(process.env.PI_ENDPOINT, res => {
    const { statusCode } = res;
    let error;

    if (statusCode !== 200) {
      error = new Error('Request Failed.\n' + `Status Code: ${statusCode}`);
      return callback(error, {
        statusCode,
      });
    }
    return callback(error, {
      statusCode,
    });
  });
};
