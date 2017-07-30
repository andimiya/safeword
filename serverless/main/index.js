'use strict';

const Alexa = require('alexa-sdk');
const request = require('request');
const handlers = {
  LaunchRequest: function() {
    this.emit('SafeIntent');
  },
  SafeIntent: function() {
    request.get(process.env.PI_ENDPOINT, (err, res) => {
      if (err) {
        console.log('ERROR VIDEO', err);
      }
      request.post(process.env.TEXT_ENDPOINT, (err, res) => {
        if (err) {
          console.log('ERROR TEXT', err);
        }
        this.emit(':tell', '');
      });
    });
  },
};

exports.handler = (event, context) => {
  const alexa = Alexa.handler(event, context);
  alexa.appId = process.env.APP_ID;
  alexa.registerHandlers(handlers);
  alexa.execute();
};
