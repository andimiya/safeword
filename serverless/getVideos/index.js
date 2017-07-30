'use strict';

const AWS = require('aws-sdk');
const s3 = new AWS.S3();

module.exports.handler = (event, context, callback) => {
  const params = {
    Bucket: process.env.BUCKET_NAME,
    StartAfter: 'videos',
  };
  s3.listObjectsV2(params, (err, data) => {
    callback(err, {
      statusCode: 200,
      headers: {
        'Access-Control-Allow-Origin': '*',
      },
      body: JSON.stringify(data.Contents),
    });
  });
};
