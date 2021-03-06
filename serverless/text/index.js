const AWS = require('aws-sdk');
const sns = new AWS.SNS();

module.exports.handler = (event, context, callback) => {
  const params = {
    Message: `It's me, ${process.env
      .NAME}. I'm in danger, please help me. I'm at home.`,
    MessageAttributes: {
      someKey: {
        DataType: 'String',
        StringValue: 'String',
      },
    },
    MessageStructure: 'String',
    PhoneNumber: process.env.PHONE_NUMBER,
    Subject: 'Emergency',
  };

  sns.publish(params, (err, data) => {
    if (err) {
      return callback(err, {
        statusCode: 500,
      });
    }
    return callback(null, {
      statusCode: 200,
    });
  });
};
